package agentcontext

import (
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	testingdomain "github.com/ugiordan/architecture-analyzer/pkg/domains/testing"
	"github.com/ugiordan/architecture-analyzer/pkg/extractor"
	"github.com/ugiordan/architecture-analyzer/pkg/graph"
)

const SchemaVersion = 1

const (
	maxSourceAreas       = 200
	maxSymbolsPerArea    = 12
	maxEntryPoints       = 500
	maxCallRelationships = 500
	maxTestRelationships = 500
	maxOperations        = 250
	maxTaintFlows        = 250
	maxReadmeDescription = 1200
)

var (
	readmeImageRE = regexp.MustCompile(`!\[[^\]]*\]\([^)]*\)`)
	readmeLinkRE  = regexp.MustCompile(`\[([^\]]+)\]\([^)]*\)`)
	readmeCodeRE  = regexp.MustCompile("`([^`]*)`")
	readmeHTMLRE  = regexp.MustCompile(`<[^>]*>`)
)

type Context struct {
	SchemaVersion              int                `json:"schema_version"`
	Component                  string             `json:"component,omitempty"`
	Repo                       string             `json:"repo,omitempty"`
	CommitSHA                  string             `json:"commit_sha,omitempty"`
	AnalyzerVersion            string             `json:"analyzer_version,omitempty"`
	Summary                    string             `json:"summary,omitempty"`
	RepositoryDescription      string             `json:"repository_description,omitempty"`
	DescriptionSource          string             `json:"description_source,omitempty"`
	DescriptionTruncated       bool               `json:"description_truncated,omitempty"`
	GraphStatus                string             `json:"graph_status"`
	Graph                      GraphSummary       `json:"graph"`
	Languages                  []LanguageSummary  `json:"languages,omitempty"`
	SourceAreas                []SourceArea       `json:"source_areas,omitempty"`
	SourceAreasTruncated       int                `json:"source_areas_truncated,omitempty"`
	EntryPoints                []EntryPoint       `json:"entry_points,omitempty"`
	EntryPointsTruncated       int                `json:"entry_points_truncated,omitempty"`
	Controllers                []ControllerFlow   `json:"controllers,omitempty"`
	CallRelationships          []CallRelationship `json:"call_relationships,omitempty"`
	CallRelationshipsTruncated int                `json:"call_relationships_truncated,omitempty"`
	TestRelationships          []TestRelationship `json:"test_relationships,omitempty"`
	TestRelationshipsTruncated int                `json:"test_relationships_truncated,omitempty"`
	Operations                 []Operation        `json:"operations,omitempty"`
	OperationsTruncated        int                `json:"operations_truncated,omitempty"`
	TaintFlows                 []TaintFlow        `json:"taint_flows,omitempty"`
	TaintFlowsTruncated        int                `json:"taint_flows_truncated,omitempty"`
	Coverage                   map[string]string  `json:"coverage,omitempty"`
}

type GraphSummary struct {
	SourceFiles       int `json:"source_files"`
	Functions         int `json:"functions"`
	TestFunctions     int `json:"test_functions"`
	CallSites         int `json:"call_sites"`
	ResolvedCalls     int `json:"resolved_calls"`
	HTTPEndpoints     int `json:"http_endpoints"`
	StorageOperations int `json:"storage_operations"`
	ExternalCalls     int `json:"external_calls"`
	TaintFlows        int `json:"taint_flows"`
}

type LanguageSummary struct {
	Name          string `json:"name"`
	SourceFiles   int    `json:"source_files"`
	Functions     int    `json:"functions"`
	TestFunctions int    `json:"test_functions"`
}

type SourceArea struct {
	Path             string   `json:"path"`
	SourceFiles      int      `json:"source_files"`
	Functions        int      `json:"functions"`
	TestFunctions    int      `json:"test_functions"`
	Languages        []string `json:"languages,omitempty"`
	Symbols          []Symbol `json:"symbols,omitempty"`
	SymbolsTruncated int      `json:"symbols_truncated,omitempty"`
}

type Symbol struct {
	Name        string   `json:"name"`
	Language    string   `json:"language,omitempty"`
	Source      string   `json:"source"`
	Complexity  int      `json:"complexity,omitempty"`
	Callers     int      `json:"callers"`
	Annotations []string `json:"annotations,omitempty"`
}

type EntryPoint struct {
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Method     string `json:"method,omitempty"`
	Path       string `json:"path,omitempty"`
	TrustLevel string `json:"trust_level,omitempty"`
	Source     string `json:"source"`
}

type ControllerFlow struct {
	Name    string            `json:"name"`
	Source  string            `json:"source,omitempty"`
	Watches []ControllerWatch `json:"watches,omitempty"`
	Steps   []ReconcileStep   `json:"steps,omitempty"`
}

type ControllerWatch struct {
	Type        string              `json:"type"`
	GVK         string              `json:"gvk"`
	Source      string              `json:"source"`
	ResourceOps []ResourceOperation `json:"resource_ops,omitempty"`
}

type ResourceOperation struct {
	Kind      string `json:"kind"`
	Group     string `json:"group,omitempty"`
	Verb      string `json:"verb"`
	Condition string `json:"condition,omitempty"`
	Source    string `json:"source"`
}

type ReconcileStep struct {
	Order     int    `json:"order"`
	Method    string `json:"method"`
	Component string `json:"component,omitempty"`
	Condition string `json:"condition,omitempty"`
	Source    string `json:"source"`
}

type CallRelationship struct {
	Caller         string `json:"caller"`
	CallerSource   string `json:"caller_source"`
	Callee         string `json:"callee"`
	CalleeSource   string `json:"callee_source"`
	CallSiteSource string `json:"call_site_source"`
	Confidence     string `json:"confidence,omitempty"`
}

type TestRelationship struct {
	Test         string `json:"test"`
	TestSource   string `json:"test_source"`
	Target       string `json:"target"`
	TargetSource string `json:"target_source"`
	Confidence   string `json:"confidence,omitempty"`
}

type Operation struct {
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Operation string `json:"operation,omitempty"`
	Target    string `json:"target,omitempty"`
	Source    string `json:"source"`
}

type TaintFlow struct {
	Label      string `json:"label"`
	Confidence string `json:"confidence,omitempty"`
	Source     string `json:"source"`
	Sink       string `json:"sink"`
}

type areaAccumulator struct {
	files     map[string]bool
	languages map[string]bool
	functions int
	tests     int
	symbols   []rankedSymbol
}

type rankedSymbol struct {
	symbol Symbol
	root   bool
}

type rankedCall struct {
	relation CallRelationship
	root     bool
	callers  int
}

func Build(repoPath string, arch *extractor.ComponentArchitecture, cpg *graph.CPG) Context {
	ctx := Context{
		SchemaVersion: SchemaVersion,
		GraphStatus:   "unavailable",
		Coverage:      map[string]string{},
	}
	if arch != nil {
		ctx.Component = arch.Component
		ctx.Repo = arch.Repo
		ctx.CommitSHA = arch.CommitSHA
		ctx.AnalyzerVersion = arch.AnalyzerVersion
		ctx.Summary = arch.Summary
		for key, value := range arch.DataCoverage {
			ctx.Coverage[key] = value
		}
	}
	ctx.RepositoryDescription, ctx.DescriptionSource, ctx.DescriptionTruncated = readmeDescription(repoPath)
	ctx.Controllers = collectControllers(arch)
	if cpg == nil {
		return ctx
	}

	ctx.GraphStatus = "available"
	nodes := cpg.Nodes()
	nodeByID := make(map[string]*graph.Node, len(nodes))
	sourceFiles := make(map[string]bool)
	languageFiles := make(map[string]map[string]bool)
	languageFunctions := make(map[string]int)
	languageTests := make(map[string]int)
	areas := make(map[string]*areaAccumulator)
	functionNodes := cpg.NodesByKind(graph.NodeFunction)
	ctx.Graph.Functions = len(functionNodes)
	ctx.Graph.CallSites = cpg.KindCount(graph.NodeCallSite)
	ctx.Graph.HTTPEndpoints = cpg.KindCount(graph.NodeHTTPEndpoint)
	ctx.Graph.StorageOperations = cpg.KindCount(graph.NodeDBOperation)
	ctx.Graph.ExternalCalls = cpg.KindCount(graph.NodeExternalCall)
	ctx.Graph.TaintFlows = countEdges(cpg.Edges(), graph.EdgeTaint)
	ctx.Graph.TestFunctions = 0

	controllerNames := make(map[string]bool)
	if arch != nil {
		for _, sequence := range arch.ReconcileSequences {
			controllerNames[sequence.Controller] = true
		}
		for _, watch := range arch.ControllerWatch {
			if watch.Controller != "" {
				controllerNames[watch.Controller] = true
			}
		}
	}

	for _, node := range nodes {
		nodeByID[node.ID] = node
		if node.File != "" {
			sourceFiles[node.File] = true
			language := languageName(node.Language)
			if languageFiles[language] == nil {
				languageFiles[language] = make(map[string]bool)
			}
			languageFiles[language][node.File] = true
			areaPath := filepath.ToSlash(filepath.Dir(node.File))
			area := areas[areaPath]
			if area == nil {
				area = &areaAccumulator{files: make(map[string]bool), languages: make(map[string]bool)}
				areas[areaPath] = area
			}
			area.files[node.File] = true
			area.languages[language] = true
		}
	}

	for _, node := range functionNodes {
		language := languageName(node.Language)
		languageFunctions[language]++
		isTest := isTestFunction(node)
		if isTest {
			ctx.Graph.TestFunctions++
			languageTests[language]++
		}
		if node.File == "" {
			continue
		}
		areaPath := filepath.ToSlash(filepath.Dir(node.File))
		area := areas[areaPath]
		area.functions++
		if isTest {
			area.tests++
		}
		callers := countIncomingCalls(cpg, node.ID)
		area.symbols = append(area.symbols, rankedSymbol{
			symbol: Symbol{
				Name:        node.Name,
				Language:    node.Language,
				Source:      sourceLocation(node.File, node.Line),
				Complexity:  node.Complexity,
				Callers:     callers,
				Annotations: trueAnnotations(node.Annotations),
			},
			root: node.Name == "main" || node.Annotations["reconciler"] || controllerNames[node.Name],
		})
	}

	ctx.Graph.SourceFiles = len(sourceFiles)
	ctx.Graph.ResolvedCalls = countEdges(cpg.Edges(), graph.EdgeCalls)
	ctx.Languages = make([]LanguageSummary, 0, len(languageFiles))
	for language, files := range languageFiles {
		ctx.Languages = append(ctx.Languages, LanguageSummary{
			Name:          language,
			SourceFiles:   len(files),
			Functions:     languageFunctions[language],
			TestFunctions: languageTests[language],
		})
	}
	sort.Slice(ctx.Languages, func(i, j int) bool { return ctx.Languages[i].Name < ctx.Languages[j].Name })

	areaPaths := make([]string, 0, len(areas))
	for path := range areas {
		areaPaths = append(areaPaths, path)
	}
	sort.Slice(areaPaths, func(i, j int) bool {
		a, b := areas[areaPaths[i]], areas[areaPaths[j]]
		if a.functions != b.functions {
			return a.functions > b.functions
		}
		return areaPaths[i] < areaPaths[j]
	})
	if len(areaPaths) > maxSourceAreas {
		ctx.SourceAreasTruncated = len(areaPaths) - maxSourceAreas
		areaPaths = areaPaths[:maxSourceAreas]
	}
	for _, path := range areaPaths {
		area := areas[path]
		languages := make([]string, 0, len(area.languages))
		for language := range area.languages {
			languages = append(languages, language)
		}
		sort.Strings(languages)
		sort.Slice(area.symbols, func(i, j int) bool {
			if area.symbols[i].root != area.symbols[j].root {
				return area.symbols[i].root
			}
			if area.symbols[i].symbol.Callers != area.symbols[j].symbol.Callers {
				return area.symbols[i].symbol.Callers > area.symbols[j].symbol.Callers
			}
			if area.symbols[i].symbol.Complexity != area.symbols[j].symbol.Complexity {
				return area.symbols[i].symbol.Complexity > area.symbols[j].symbol.Complexity
			}
			if area.symbols[i].symbol.Name != area.symbols[j].symbol.Name {
				return area.symbols[i].symbol.Name < area.symbols[j].symbol.Name
			}
			return area.symbols[i].symbol.Source < area.symbols[j].symbol.Source
		})
		symbolCount := len(area.symbols)
		if symbolCount > maxSymbolsPerArea {
			symbolCount = maxSymbolsPerArea
		}
		symbols := make([]Symbol, 0, symbolCount)
		for _, item := range area.symbols[:symbolCount] {
			symbols = append(symbols, item.symbol)
		}
		ctx.SourceAreas = append(ctx.SourceAreas, SourceArea{
			Path:             path,
			SourceFiles:      len(area.files),
			Functions:        area.functions,
			TestFunctions:    area.tests,
			Languages:        languages,
			Symbols:          symbols,
			SymbolsTruncated: len(area.symbols) - symbolCount,
		})
	}

	ctx.EntryPoints = collectEntryPoints(cpg)
	if len(ctx.EntryPoints) > maxEntryPoints {
		ctx.EntryPointsTruncated = len(ctx.EntryPoints) - maxEntryPoints
		ctx.EntryPoints = ctx.EntryPoints[:maxEntryPoints]
	}

	callRows, testRows := collectCalls(cpg, nodeByID, controllerNames)
	if len(callRows) > maxCallRelationships {
		ctx.CallRelationshipsTruncated = len(callRows) - maxCallRelationships
		callRows = callRows[:maxCallRelationships]
	}
	ctx.CallRelationships = make([]CallRelationship, 0, len(callRows))
	for _, item := range callRows {
		ctx.CallRelationships = append(ctx.CallRelationships, item.relation)
	}
	if len(testRows) > maxTestRelationships {
		ctx.TestRelationshipsTruncated = len(testRows) - maxTestRelationships
		testRows = testRows[:maxTestRelationships]
	}
	ctx.TestRelationships = testRows

	ctx.Operations = collectOperations(cpg)
	if len(ctx.Operations) > maxOperations {
		ctx.OperationsTruncated = len(ctx.Operations) - maxOperations
		ctx.Operations = ctx.Operations[:maxOperations]
	}
	ctx.TaintFlows = collectTaintFlows(cpg, nodeByID)
	if len(ctx.TaintFlows) > maxTaintFlows {
		ctx.TaintFlowsTruncated = len(ctx.TaintFlows) - maxTaintFlows
		ctx.TaintFlows = ctx.TaintFlows[:maxTaintFlows]
	}
	return ctx
}

func readmeDescription(repoPath string) (string, string, bool) {
	root, rootErr := os.OpenRoot(repoPath)
	if rootErr != nil {
		return "", "", false
	}
	defer root.Close()
	for _, name := range []string{"README.md", "README.rst", "README.txt", "README"} {
		info, err := root.Stat(name)
		if err != nil || !info.Mode().IsRegular() || info.Size() > 2*1024*1024 {
			continue
		}
		content, err := root.ReadFile(name)
		if err != nil {
			continue
		}
		lines := strings.Split(string(content), "\n")
		var paragraph []string
		startLine := 0
		inFence := false
		for i, raw := range lines {
			line := strings.TrimSpace(raw)
			if strings.HasPrefix(line, "```") || strings.HasPrefix(line, "~~~") {
				inFence = !inFence
				continue
			}
			if inFence {
				continue
			}
			if line == "" {
				if len(paragraph) > 0 {
					break
				}
				continue
			}
			if len(paragraph) > 0 && isSetextUnderline(line) {
				paragraph = nil
				startLine = 0
				continue
			}
			if len(paragraph) == 0 && isReadmeDecoration(line) {
				continue
			}
			if len(paragraph) > 0 && (strings.HasPrefix(line, "#") || strings.HasPrefix(line, "-") || strings.HasPrefix(line, "*") || strings.HasPrefix(line, "|")) {
				break
			}
			if startLine == 0 {
				startLine = i + 1
			}
			paragraph = append(paragraph, line)
		}
		if len(paragraph) == 0 {
			continue
		}
		text := cleanReadmeText(strings.Join(paragraph, " "))
		if text == "" {
			continue
		}
		runes := []rune(text)
		truncated := len(runes) > maxReadmeDescription
		if truncated {
			text = string(runes[:maxReadmeDescription])
		}
		return text, sourceLocation(name, startLine), truncated
	}
	return "", "", false
}

func isReadmeDecoration(line string) bool {
	return strings.HasPrefix(line, "#") || strings.HasPrefix(line, "[![") || strings.HasPrefix(line, "![") ||
		strings.HasPrefix(line, "<!--") || strings.HasPrefix(line, "<img") || strings.HasPrefix(line, "---") ||
		strings.HasPrefix(line, "```") || strings.HasPrefix(line, "~~~")
}

// isSetextUnderline reports whether line is a Markdown setext heading
// underline: a run of '=' or '-' characters only. A paragraph directly
// followed by such a line is a heading, not description prose.
func isSetextUnderline(line string) bool {
	if line == "" {
		return false
	}
	char := line[0]
	if char != '=' && char != '-' {
		return false
	}
	for i := 0; i < len(line); i++ {
		if line[i] != char {
			return false
		}
	}
	return true
}

func cleanReadmeText(text string) string {
	text = readmeImageRE.ReplaceAllString(text, "")
	text = readmeLinkRE.ReplaceAllString(text, "$1")
	text = readmeCodeRE.ReplaceAllString(text, "$1")
	text = readmeHTMLRE.ReplaceAllString(text, "")
	text = strings.TrimSpace(strings.TrimLeft(text, "-*> "))
	return strings.Join(strings.Fields(text), " ")
}

func collectEntryPoints(cpg *graph.CPG) []EntryPoint {
	var result []EntryPoint
	seen := make(map[string]bool)
	for _, node := range cpg.NodesByKind(graph.NodeFunction) {
		if node.Name != "main" || node.File == "" {
			continue
		}
		entry := EntryPoint{Kind: "executable", Name: node.Name, Source: sourceLocation(node.File, node.Line)}
		if !seen[entry.Kind+"|"+entry.Source] {
			seen[entry.Kind+"|"+entry.Source] = true
			result = append(result, entry)
		}
	}
	for _, node := range cpg.NodesByKind(graph.NodeHTTPEndpoint) {
		if node.File == "" {
			continue
		}
		entry := EntryPoint{
			Kind:       "http",
			Name:       node.Name,
			Method:     node.HTTPMethod,
			Path:       node.Route,
			TrustLevel: string(node.TrustLevel),
			Source:     sourceLocation(node.File, node.Line),
		}
		key := strings.Join([]string{entry.Kind, entry.Name, entry.Method, entry.Path, entry.Source}, "|")
		if !seen[key] {
			seen[key] = true
			result = append(result, entry)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Kind != result[j].Kind {
			return result[i].Kind < result[j].Kind
		}
		if result[i].Path != result[j].Path {
			return result[i].Path < result[j].Path
		}
		if result[i].Name != result[j].Name {
			return result[i].Name < result[j].Name
		}
		return result[i].Source < result[j].Source
	})
	return result
}

func collectControllers(arch *extractor.ComponentArchitecture) []ControllerFlow {
	if arch == nil {
		return nil
	}
	byName := make(map[string]*ControllerFlow)
	for _, sequence := range arch.ReconcileSequences {
		flow := byName[sequence.Controller]
		if flow == nil {
			flow = &ControllerFlow{Name: sequence.Controller, Source: sequence.Source}
			byName[sequence.Controller] = flow
		}
		for _, step := range sequence.Steps {
			flow.Steps = append(flow.Steps, ReconcileStep{
				Order: step.Order, Method: step.Method, Component: step.Component,
				Condition: step.Condition, Source: step.Source,
			})
		}
	}
	for _, watch := range arch.ControllerWatch {
		name := watch.Controller
		if name == "" {
			name = "Unknown controller"
		}
		flow := byName[name]
		if flow == nil {
			flow = &ControllerFlow{Name: name}
			byName[name] = flow
		}
		item := ControllerWatch{Type: watch.Type, GVK: watch.GVK, Source: watch.Source}
		for _, op := range watch.ResourceOps {
			item.ResourceOps = append(item.ResourceOps, ResourceOperation{
				Kind: op.Kind, Group: op.Group, Verb: op.Verb, Condition: op.Condition,
				Source: sourceLocation(op.Source.File, op.Source.Line),
			})
		}
		flow.Watches = append(flow.Watches, item)
	}
	result := make([]ControllerFlow, 0, len(byName))
	for _, flow := range byName {
		sort.Slice(flow.Steps, func(i, j int) bool {
			if flow.Steps[i].Order != flow.Steps[j].Order {
				return flow.Steps[i].Order < flow.Steps[j].Order
			}
			return flow.Steps[i].Method < flow.Steps[j].Method
		})
		sort.Slice(flow.Watches, func(i, j int) bool {
			if flow.Watches[i].Type != flow.Watches[j].Type {
				return flow.Watches[i].Type < flow.Watches[j].Type
			}
			return flow.Watches[i].GVK < flow.Watches[j].GVK
		})
		result = append(result, *flow)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Name < result[j].Name })
	return result
}

func collectCalls(cpg *graph.CPG, nodeByID map[string]*graph.Node, controllerNames map[string]bool) ([]rankedCall, []TestRelationship) {
	parentByCallSite := make(map[string]*graph.Node)
	for _, edge := range cpg.Edges() {
		if edge.Kind != graph.EdgeDataFlow || edge.Label != "contains_call" {
			continue
		}
		parent := nodeByID[edge.From]
		callSite := nodeByID[edge.To]
		if parent != nil && parent.Kind == graph.NodeFunction && callSite != nil && callSite.Kind == graph.NodeCallSite {
			parentByCallSite[edge.To] = parent
		}
	}
	callersByFunction := make(map[string]int)
	for _, edge := range cpg.Edges() {
		if edge.Kind == graph.EdgeCalls {
			callersByFunction[edge.To]++
		}
	}
	var calls []rankedCall
	var tests []TestRelationship
	seenCalls := make(map[string]bool)
	seenTests := make(map[string]bool)
	for _, edge := range cpg.Edges() {
		if edge.Kind != graph.EdgeCalls {
			continue
		}
		callSite := nodeByID[edge.From]
		caller := parentByCallSite[edge.From]
		callee := nodeByID[edge.To]
		if callSite == nil || caller == nil || callee == nil || callee.Kind != graph.NodeFunction {
			continue
		}
		relation := CallRelationship{
			Caller:         caller.Name,
			CallerSource:   sourceLocation(caller.File, caller.Line),
			Callee:         callee.Name,
			CalleeSource:   sourceLocation(callee.File, callee.Line),
			CallSiteSource: sourceLocation(callSite.File, callSite.Line),
			Confidence:     string(edge.Confidence),
		}
		key := strings.Join([]string{caller.ID, callee.ID, callSite.ID}, "|")
		if seenCalls[key] {
			continue
		}
		seenCalls[key] = true
		calls = append(calls, rankedCall{
			relation: relation,
			root:     caller.Name == "main" || caller.Annotations["reconciler"] || controllerNames[caller.Name],
			callers:  callersByFunction[callee.ID],
		})
		if isTestFunction(caller) && !isTestFunction(callee) {
			testKey := strings.Join([]string{caller.ID, callee.ID}, "|")
			if !seenTests[testKey] {
				seenTests[testKey] = true
				tests = append(tests, TestRelationship{
					Test: caller.Name, TestSource: sourceLocation(caller.File, caller.Line),
					Target: callee.Name, TargetSource: sourceLocation(callee.File, callee.Line),
					Confidence: string(edge.Confidence),
				})
			}
		}
	}
	sort.Slice(calls, func(i, j int) bool {
		if calls[i].root != calls[j].root {
			return calls[i].root
		}
		if calls[i].callers != calls[j].callers {
			return calls[i].callers > calls[j].callers
		}
		if calls[i].relation.CallerSource != calls[j].relation.CallerSource {
			return calls[i].relation.CallerSource < calls[j].relation.CallerSource
		}
		if calls[i].relation.Callee != calls[j].relation.Callee {
			return calls[i].relation.Callee < calls[j].relation.Callee
		}
		return calls[i].relation.CallSiteSource < calls[j].relation.CallSiteSource
	})
	sort.Slice(tests, func(i, j int) bool {
		if tests[i].TestSource != tests[j].TestSource {
			return tests[i].TestSource < tests[j].TestSource
		}
		if tests[i].Target != tests[j].Target {
			return tests[i].Target < tests[j].Target
		}
		return tests[i].TargetSource < tests[j].TargetSource
	})
	return calls, tests
}

func collectOperations(cpg *graph.CPG) []Operation {
	var result []Operation
	for _, kind := range []graph.NodeKind{graph.NodeDBOperation, graph.NodeExternalCall} {
		for _, node := range cpg.NodesByKind(kind) {
			if node.File == "" {
				continue
			}
			target := node.Table
			if target == "" {
				target = node.TypeName
			}
			result = append(result, Operation{
				Kind: string(node.Kind), Name: node.Name, Operation: node.Operation,
				Target: target, Source: sourceLocation(node.File, node.Line),
			})
		}
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Kind != result[j].Kind {
			return result[i].Kind < result[j].Kind
		}
		if result[i].Source != result[j].Source {
			return result[i].Source < result[j].Source
		}
		return result[i].Name < result[j].Name
	})
	return result
}

func collectTaintFlows(cpg *graph.CPG, nodeByID map[string]*graph.Node) []TaintFlow {
	var result []TaintFlow
	seen := make(map[string]bool)
	for _, edge := range cpg.Edges() {
		if edge.Kind != graph.EdgeTaint {
			continue
		}
		source, sink := nodeByID[edge.From], nodeByID[edge.To]
		if source == nil || sink == nil {
			continue
		}
		flow := TaintFlow{
			Label: edge.Label, Confidence: string(edge.Confidence),
			Source: sourceLocation(source.File, source.Line),
			Sink:   sourceLocation(sink.File, sink.Line),
		}
		key := strings.Join([]string{flow.Label, flow.Source, flow.Sink}, "|")
		if !seen[key] {
			seen[key] = true
			result = append(result, flow)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Label != result[j].Label {
			return result[i].Label < result[j].Label
		}
		if result[i].Source != result[j].Source {
			return result[i].Source < result[j].Source
		}
		return result[i].Sink < result[j].Sink
	})
	return result
}

func countEdges(edges []*graph.Edge, kind graph.EdgeKind) int {
	count := 0
	for _, edge := range edges {
		if edge.Kind == kind {
			count++
		}
	}
	return count
}

func countIncomingCalls(cpg *graph.CPG, nodeID string) int {
	count := 0
	for _, edge := range cpg.InEdges(nodeID) {
		if edge.Kind == graph.EdgeCalls {
			count++
		}
	}
	return count
}

func isTestFunction(node *graph.Node) bool {
	return node.IsTest || node.Annotations[testingdomain.AnnotIsTestFunc]
}

func languageName(language string) string {
	if language == "" {
		return "unknown"
	}
	return language
}

func sourceLocation(file string, line int) string {
	if file == "" {
		return ""
	}
	if line > 0 {
		return file + ":" + strconv.Itoa(line)
	}
	return file
}

func trueAnnotations(annotations map[string]bool) []string {
	var result []string
	for name, enabled := range annotations {
		if enabled {
			result = append(result, name)
		}
	}
	sort.Strings(result)
	return result
}
