package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ugiordan/architecture-analyzer/pkg/annotator"
	"github.com/ugiordan/architecture-analyzer/pkg/arch"
	"github.com/ugiordan/architecture-analyzer/pkg/builder"
	"github.com/ugiordan/architecture-analyzer/pkg/dataflow"
	"github.com/ugiordan/architecture-analyzer/pkg/domains"
	"github.com/ugiordan/architecture-analyzer/pkg/extractor"
	"github.com/ugiordan/architecture-analyzer/pkg/graph"
	"github.com/ugiordan/architecture-analyzer/pkg/linker"
	"github.com/ugiordan/architecture-analyzer/pkg/query"
	"github.com/ugiordan/architecture-analyzer/pkg/sarif"
)

func cmdScan(args []string) error {
	fs := flag.NewFlagSet("scan", flag.ExitOnError)
	outputFile := fs.String("output", "", "Output findings JSON file (default: stdout)")
	format := fs.String("format", "text", "Output format: text, json, sarif")
	domainList := fs.String("domains", "", "Comma-separated domains to run (default: all registered)")
	withArch := fs.Bool("with-arch", false, "Cross-reference with architecture data")
	importSARIF := fs.String("import-sarif", "", "Comma-separated SARIF files to ingest after building graph")
	fs.Parse(args)

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: arch-analyzer scan <repo-path> [--output file] [--format text|json|sarif] [--domains sec,test] [--with-arch] [--import-sarif files]")
	}

	repoPath := fs.Arg(0)

	cpg, err := buildCPG(repoPath)
	if err != nil {
		return err
	}

	var archData *domains.ArchitectureData
	if *withArch {
		archData = prepareArchData(repoPath, cpg, "")
	}

	findings, err := runSecurityScan(cpg, *domainList, archData)
	if err != nil {
		return err
	}

	if *importSARIF != "" {
		externalFindings, sarifErr := ingestSARIFFiles(cpg, *importSARIF)
		if sarifErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: SARIF ingestion: %v\n", sarifErr)
		}
		findings = append(findings, externalFindings...)
	}

	registeredDomains := domains.Names()
	switch *format {
	case "text":
		printFindings(cpg, findings)
	case "json":
		if len(registeredDomains) > 0 {
			return outputJSON(*outputFile, domainGroupedJSON(findings))
		}
		return outputJSON(*outputFile, findings)
	case "sarif":
		return outputSARIF(*outputFile, findings)
	default:
		return fmt.Errorf("unknown format: %s", *format)
	}
	return nil
}

// cmdDomains lists all registered analysis domains.
func cmdDomains() error {
	registered := domains.All()
	if len(registered) == 0 {
		fmt.Println("No domains registered.")
		return nil
	}
	fmt.Printf("%d registered domain(s):\n", len(registered))
	for _, d := range registered {
		fmt.Printf("  %-12s languages: %s", d.Name(), strings.Join(d.SupportedLanguages(), ", "))
		deps := d.Dependencies()
		if len(deps) > 0 {
			fmt.Printf("  deps: %s", strings.Join(deps, ", "))
		}
		fmt.Printf("  queries: %d\n", len(d.Queries()))
	}
	return nil
}

// cmdGraph exports the code property graph.
func cmdGraph(args []string) error {
	fs := flag.NewFlagSet("graph", flag.ExitOnError)
	format := fs.String("format", "json", "Output format: json, dot")
	outputFile := fs.String("output", "", "Output file (default: stdout)")
	fs.Parse(args)

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: arch-analyzer graph <repo-path> [--format json|dot] [--output file]")
	}

	cpg, err := buildCPG(fs.Arg(0))
	if err != nil {
		return err
	}

	var content []byte
	switch *format {
	case "json":
		output := map[string]interface{}{
			"schema_version": graph.SchemaVersion,
			"nodes":          cpg.Nodes(),
			"edges":          cpg.Edges(),
		}
		content, err = json.MarshalIndent(output, "", "  ")
		if err != nil {
			return err
		}
		content = append(content, '\n')
	case "dot":
		content = []byte(renderDOT(cpg))
	default:
		return fmt.Errorf("unknown format: %s", *format)
	}

	if *outputFile != "" {
		return os.WriteFile(*outputFile, content, 0o644)
	}
	_, err = os.Stdout.Write(content)
	return err
}

// cmdIngest ingests SARIF findings into an optional existing code graph.
func cmdIngest(args []string) error {
	fs := flag.NewFlagSet("ingest", flag.ExitOnError)
	graphFile := fs.String("graph", "", "Existing code-graph.json to enrich")
	outputFile := fs.String("output", "", "Output file for enriched graph (default: stdout)")
	fs.Parse(args)

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: arch-analyzer ingest <sarif-file> [--graph code-graph.json] [--output file]")
	}

	sarifPath := fs.Arg(0)
	f, err := os.Open(sarifPath)
	if err != nil {
		return fmt.Errorf("opening SARIF file: %w", err)
	}
	defer f.Close()

	report, err := sarif.Parse(f)
	if err != nil {
		return fmt.Errorf("parsing SARIF: %w", err)
	}

	if *graphFile == "" {
		// Standalone mode: create a fresh CPG, ingest, output only the finding nodes
		cpg := graph.NewCPG()
		result, err := sarif.Ingest(cpg, report, "")
		if err != nil {
			return fmt.Errorf("ingesting SARIF: %w", err)
		}

		output := map[string]interface{}{
			"nodes":   cpg.NodesByKind(graph.NodeExternalFinding),
			"summary": result,
		}
		return outputJSON(*outputFile, output)
	}

	// Graph mode: load existing CPG from JSON, ingest, output enriched graph
	cpg, err := loadCPGFromJSON(*graphFile)
	if err != nil {
		return fmt.Errorf("loading graph: %w", err)
	}

	result, err := sarif.Ingest(cpg, report, "")
	if err != nil {
		return fmt.Errorf("ingesting SARIF: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Ingested: %d findings (%d linked, %d unlinked) from %s\n",
		result.FindingsTotal, result.FindingsLinked, result.FindingsUnlinked, result.ToolSummary())

	enriched := map[string]interface{}{
		"schema_version": 3,
		"nodes":          cpg.Nodes(),
		"edges":          cpg.Edges(),
	}
	return outputJSON(*outputFile, enriched)
}

// loadCPGFromJSON reconstructs a CPG from a code-graph.json file.
func loadCPGFromJSON(path string) (*graph.CPG, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var raw struct {
		SchemaVersion int          `json:"schema_version"`
		Nodes         []graph.Node `json:"nodes"`
		Edges         []graph.Edge `json:"edges"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}

	cpg := graph.NewCPG()
	for i := range raw.Nodes {
		if err := cpg.AddNode(&raw.Nodes[i]); err != nil {
			return nil, fmt.Errorf("loading node: %w", err)
		}
	}
	for i := range raw.Edges {
		cpg.AddEdge(&raw.Edges[i])
	}
	return cpg, nil
}

// cmdDiff computes a structural diff between two code-graph.json snapshots.

func runSecurityScan(cpg *graph.CPG, domainList string, archData *domains.ArchitectureData) ([]query.Finding, error) {
	// Resolve domain analyzers
	var analyzers []domains.DomainAnalyzer
	if domainList != "" {
		names := strings.Split(domainList, ",")
		resolved, resolveErr := domains.ResolveDependencies(names)
		if resolveErr != nil {
			return nil, resolveErr
		}
		var err error
		analyzers, err = domains.Get(resolved)
		if err != nil {
			return nil, err
		}
	} else {
		analyzers = domains.All()
	}

	// Phase 1: Run all domain annotations
	var orch *domains.Orchestrator
	if len(analyzers) > 0 {
		orch = domains.NewOrchestrator(analyzers)
		if err := orch.AnnotateAll(cpg, "go", archData); err != nil {
			return nil, fmt.Errorf("domain annotation: %w", err)
		}
	}

	// Phase 2: Run taint propagation engine
	te := dataflow.NewTaintEngine()
	taintEdges := te.Run(cpg)
	for _, e := range taintEdges {
		cpg.AddEdge(e)
	}
	if len(taintEdges) > 0 {
		fmt.Printf("Taint engine: %d taint paths found\n", len(taintEdges))
	}

	// Phase 3: Run legacy queries (CGA-002 now uses EdgeTaint)
	engine := query.NewEngine()
	findings := engine.RunAll(cpg)

	// Phase 4: Run domain queries
	if orch != nil {
		results, runErr := orch.RunQueries(cpg)
		if runErr != nil {
			return nil, fmt.Errorf("domain queries: %w", runErr)
		}
		for _, dr := range results {
			fmt.Printf("Domain %s: %d findings\n", dr.Domain, len(dr.Findings))
			findings = append(findings, dr.Findings...)
		}
	}

	return findings, nil
}

// prepareArchData extracts architecture data from a repo and sets it on the CPG.
func prepareArchData(repoPath string, cpg *graph.CPG, org string) *domains.ArchitectureData {
	opts := &extractor.ExtractOptions{Org: org}
	archResult, err := extractor.ExtractAll(repoPath, opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: architecture extraction failed: %v\n", err)
		return nil
	}

	raw, mErr := json.Marshal(archResult)
	if mErr != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to marshal architecture data: %v\n", mErr)
		return nil
	}
	var data map[string]interface{}
	if err := json.Unmarshal(raw, &data); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to unmarshal architecture data: %v\n", err)
		return nil
	}
	archData := &domains.ArchitectureData{}

	parsed, parseErr := arch.Parse(data)
	if parseErr != nil {
		fmt.Fprintf(os.Stderr, "Warning: architecture data parsing failed: %v\n", parseErr)
	} else {
		cpg.ArchData = parsed
	}

	return archData
}

// ingestSARIFFiles processes comma-separated SARIF file paths, ingesting each
// into the CPG independently. Returns all ExternalFinding nodes converted to
// query.Finding. If some files fail, findings from successful files are still
// returned alongside a warning error. If ALL files fail, returns an error.
func ingestSARIFFiles(cpg *graph.CPG, paths string) ([]query.Finding, error) {
	files := strings.Split(paths, ",")

	// Record existing ExternalFinding node IDs so we only convert new ones.
	existingIDs := make(map[string]bool)
	for _, n := range cpg.NodesByKind(graph.NodeExternalFinding) {
		existingIDs[n.ID] = true
	}

	var errors []string
	succeeded := 0

	for _, raw := range files {
		p := strings.TrimSpace(raw)
		if p == "" {
			continue
		}

		f, err := os.Open(p)
		if err != nil {
			errors = append(errors, fmt.Sprintf("%s: %v", p, err))
			continue
		}

		report, err := sarif.Parse(f)
		f.Close()
		if err != nil {
			errors = append(errors, fmt.Sprintf("%s: parse error: %v", p, err))
			continue
		}

		result, err := sarif.Ingest(cpg, report, "")
		if err != nil {
			errors = append(errors, fmt.Sprintf("%s: ingest error: %v", p, err))
			continue
		}

		fmt.Fprintf(os.Stderr, "SARIF %s: %d findings (%d linked, %d unlinked) from %s\n",
			filepath.Base(p), result.FindingsTotal, result.FindingsLinked, result.FindingsUnlinked, result.ToolSummary())
		succeeded++
	}

	// Convert all new ExternalFinding nodes to query.Finding.
	var findings []query.Finding
	for _, n := range cpg.NodesByKind(graph.NodeExternalFinding) {
		if existingIDs[n.ID] {
			continue
		}
		findings = append(findings, query.Finding{
			RuleID:      n.RuleID,
			Severity:    n.Severity,
			Message:     n.Message,
			File:        n.File,
			Line:        n.Line,
			NodeID:      n.ID,
			Domain:      "external/" + n.ToolName,
			ToolName:    n.ToolName,
			ToolVersion: n.ToolVersion,
		})
	}

	if succeeded == 0 && len(errors) > 0 {
		return nil, fmt.Errorf("all SARIF files failed: %s", strings.Join(errors, "; "))
	}
	if len(errors) > 0 {
		return findings, fmt.Errorf("some SARIF files failed: %s", strings.Join(errors, "; "))
	}
	return findings, nil
}

// buildCPG constructs a Code Property Graph from a repo directory.
func buildCPG(repoPath string) (*graph.CPG, error) {
	b := builder.NewBuilder()
	cpg, err := b.BuildFromDir(repoPath)
	if err != nil {
		return nil, fmt.Errorf("building code graph: %w", err)
	}

	sl := linker.NewStorageLinker()
	linked := sl.Link(cpg)

	sa := annotator.NewSecurityAnnotator()
	sa.Annotate(cpg)

	fmt.Printf("Graph: %d nodes, %d edges, %d storage links\n",
		len(cpg.Nodes()), len(cpg.Edges()), linked)
	fmt.Printf("  Functions: %d, Call sites: %d, HTTP handlers: %d, DB ops: %d\n",
		len(cpg.NodesByKind(graph.NodeFunction)),
		len(cpg.NodesByKind(graph.NodeCallSite)),
		len(cpg.NodesByKind(graph.NodeHTTPEndpoint)),
		len(cpg.NodesByKind(graph.NodeDBOperation)))

	return cpg, nil
}

func printFindings(cpg *graph.CPG, findings []query.Finding) {
	if len(findings) == 0 {
		fmt.Println("No security findings.")
		return
	}
	fmt.Printf("\n%d security finding(s):\n", len(findings))
	externalByTool := make(map[string]int)
	for _, f := range findings {
		prefix := ""
		if strings.HasPrefix(f.Domain, "external/") {
			prefix = fmt.Sprintf("[ext/%s] ", f.ToolName)
			externalByTool[f.ToolName]++
		} else if f.Domain != "" && f.Domain != "legacy" {
			prefix = fmt.Sprintf("[%s] ", f.Domain)
		}
		fmt.Printf("  %s[%s] %s: %s (%s:%d)\n", prefix, f.Severity, f.RuleID, f.Message, f.File, f.Line)
	}
	if len(externalByTool) > 0 {
		fmt.Printf("\nExternal findings summary:\n")
		total := 0
		for tool, count := range externalByTool {
			fmt.Printf("  %s: %d finding(s)\n", tool, count)
			total += count
		}
		fmt.Printf("  Total external: %d of %d findings\n", total, len(findings))

		// Cross-reference: internal findings corroborated by external scanners
		correlated := sarif.CrossReference(cpg, findings)
		if output := sarif.FormatCorrelations(correlated); output != "" {
			fmt.Print(output)
		}
	}
}

func renderDOT(cpg *graph.CPG) string {
	var b strings.Builder
	b.WriteString("digraph CPG {\n")
	b.WriteString("  rankdir=LR;\n")
	for _, n := range cpg.Nodes() {
		label := fmt.Sprintf("%s\\n(%s)", n.Name, n.Kind)
		fmt.Fprintf(&b, "  %q [label=%q];\n", n.ID, label)
	}
	for _, e := range cpg.Edges() {
		fmt.Fprintf(&b, "  %q -> %q [label=%q];\n", e.From, e.To, e.Kind)
	}
	b.WriteString("}\n")
	return b.String()
}

func outputSARIF(path string, findings []query.Finding) error {
	// Group findings by domain for per-domain SARIF runs
	grouped := make(map[string][]query.Finding)
	for _, f := range findings {
		domain := f.Domain
		if domain == "" {
			domain = "legacy"
		}
		grouped[domain] = append(grouped[domain], f)
	}

	var runs []map[string]interface{}
	for domain, domainFindings := range grouped {
		toolName := "arch-analyzer/" + domain
		toolVersion := version
		// For external findings, preserve the original tool identity
		if strings.HasPrefix(domain, "external/") && len(domainFindings) > 0 {
			toolName = domainFindings[0].ToolName
			if domainFindings[0].ToolVersion != "" {
				toolVersion = domainFindings[0].ToolVersion
			}
		}
		runs = append(runs, map[string]interface{}{
			"tool": map[string]interface{}{
				"driver": map[string]interface{}{
					"name":    toolName,
					"version": toolVersion,
				},
			},
			"results": sarifResults(domainFindings),
		})
	}

	sarif := map[string]interface{}{
		"$schema": "https://raw.githubusercontent.com/oasis-tcs/sarif-spec/main/sarif-2.1/schema/sarif-schema-2.1.0.json",
		"version": "2.1.0",
		"runs":    runs,
	}
	if path != "" {
		return writeJSON(path, sarif)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(sarif)
}

func domainGroupedJSON(findings []query.Finding) map[string]interface{} {
	grouped := make(map[string][]query.Finding)
	externalByTool := make(map[string]int)
	for _, f := range findings {
		domain := f.Domain
		if domain == "" {
			domain = "legacy"
		}
		grouped[domain] = append(grouped[domain], f)
		if strings.HasPrefix(f.Domain, "external/") {
			externalByTool[f.ToolName]++
		}
	}

	domainResults := make(map[string]interface{})
	for domain, domainFindings := range grouped {
		domainResults[domain] = map[string]interface{}{
			"findings": domainFindings,
			"count":    len(domainFindings),
		}
	}

	result := map[string]interface{}{
		"domains":        domainResults,
		"total_findings": len(findings),
	}
	if len(externalByTool) > 0 {
		result["external_tools"] = externalByTool
	}
	return result
}

func sarifResults(findings []query.Finding) []map[string]interface{} {
	results := make([]map[string]interface{}, 0, len(findings))
	for _, f := range findings {
		r := map[string]interface{}{
			"ruleId": f.RuleID,
			"level":  sarifLevel(f.Severity),
			"message": map[string]string{
				"text": f.Message,
			},
			"locations": []map[string]interface{}{
				{
					"physicalLocation": map[string]interface{}{
						"artifactLocation": map[string]string{
							"uri": f.File,
						},
						"region": map[string]int{
							"startLine": f.Line,
						},
					},
				},
			},
		}
		results = append(results, r)
	}
	return results
}

func sarifLevel(severity string) string {
	switch strings.ToLower(severity) {
	case "critical", "high":
		return "error"
	case "medium":
		return "warning"
	default:
		return "note"
	}
}

func outputJSON(path string, data interface{}) error {
	if path != "" {
		return writeJSON(path, data)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}

func writeJSON(path string, data interface{}) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
		return fmt.Errorf("creating directory for %s: %w", path, err)
	}
	raw, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(raw, '\n'), 0o640)
}

// parseAliases splits a comma-separated aliases string into a slice,
// trimming whitespace. Returns nil for empty input.
func parseAliases(raw string) []string {
	if raw == "" {
		return nil
	}
	parts := strings.Split(raw, ",")
	var out []string
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func loadJSON(path string) (map[string]interface{}, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}
	return data, nil
}

func writeDiagrams(dir string, diagrams map[string]string) error {
	if err := os.MkdirAll(dir, 0o750); err != nil {
		return fmt.Errorf("creating diagram dir: %w", err)
	}
	for filename, content := range diagrams {
		path := filepath.Join(dir, filename)
		if err := os.WriteFile(path, []byte(content), 0o640); err != nil {
			return fmt.Errorf("writing diagram %s: %w", path, err)
		}
	}
	return nil
}
