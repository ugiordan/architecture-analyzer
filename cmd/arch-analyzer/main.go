// Package main implements the arch-analyzer CLI. As the tool grows, consider
// splitting subcommands into dedicated packages under internal/cmd/.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/ugiordan/architecture-analyzer/pkg/aggregator"
	"github.com/ugiordan/architecture-analyzer/pkg/arch"
	"github.com/ugiordan/architecture-analyzer/pkg/config"
	"github.com/ugiordan/architecture-analyzer/pkg/domains"
	"github.com/ugiordan/architecture-analyzer/pkg/domains/architecture"
	"github.com/ugiordan/architecture-analyzer/pkg/domains/netpolicy"
	"github.com/ugiordan/architecture-analyzer/pkg/domains/security"
	testingdomain "github.com/ugiordan/architecture-analyzer/pkg/domains/testing"
	"github.com/ugiordan/architecture-analyzer/pkg/domains/upgrade"
	"github.com/ugiordan/architecture-analyzer/pkg/extractor"
	"github.com/ugiordan/architecture-analyzer/pkg/graph"
	"github.com/ugiordan/architecture-analyzer/pkg/query"
	"github.com/ugiordan/architecture-analyzer/pkg/renderer"
	"github.com/ugiordan/architecture-analyzer/pkg/validator"
)

var version = "dev"

func init() {
	extractor.AnalyzerVersion = version
	domains.Register(security.New())
	domains.Register(testingdomain.New())
	domains.Register(upgrade.New())
	domains.Register(architecture.New())
	domains.Register(netpolicy.New())
}

// versionLabelRe validates version labels for snapshot output directories.
var versionLabelRe = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9._-]{0,63}$`)

var errDiffFound = fmt.Errorf("differences found")

// validateVersionLabel checks that a version label is safe for use in file paths and git tags.
func validateVersionLabel(label string) error {
	if !versionLabelRe.MatchString(label) {
		return fmt.Errorf("invalid version label %q: must match %s", label, versionLabelRe.String())
	}
	return nil
}

// applyVersion adjusts outputDir to include a version subdirectory when version is non-empty.
// e.g. applyVersion("output", "v2.15.0") returns "output/v2.15.0".
func applyVersion(outputDir, ver string) string {
	if ver == "" {
		return outputDir
	}
	return filepath.Join(outputDir, ver)
}

// snapshotMetadata holds version and provenance information written alongside output.
type snapshotMetadata struct {
	Version         string            `json:"version"`
	Timestamp       string            `json:"timestamp"`
	AnalyzerVersion string            `json:"analyzer_version"`
	ReposAnalyzed   map[string]string `json:"repos_analyzed"`
	Platform        string            `json:"platform,omitempty"`
}

// writeSnapshotMetadata writes snapshot-metadata.json to the given directory.
func writeSnapshotMetadata(dir, ver string, repos map[string]string, platform string) error {
	meta := snapshotMetadata{
		Version:         ver,
		Timestamp:       time.Now().UTC().Format(time.RFC3339),
		AnalyzerVersion: version,
		ReposAnalyzed:   repos,
		Platform:        platform,
	}
	path := filepath.Join(dir, "snapshot-metadata.json")
	return writeJSON(path, meta)
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	var err error
	switch cmd {
	case "extract":
		err = cmdExtract(args)
	case "render":
		err = cmdRender(args)
	case "analyze":
		err = cmdAnalyze(args)
	case "aggregate":
		err = cmdAggregate(args)
	case "aggregate-cpg":
		err = cmdAggregateCPG(args)
	case "extract-schema":
		err = cmdExtractSchema(args)
	case "validate":
		err = cmdValidate(args)
	case "scan":
		err = cmdScan(args)
	case "context-bundle":
		err = cmdContextBundle(args)
	case "graph":
		err = cmdGraph(args)
	case "diff":
		err = cmdDiff(args)
	case "flow":
		err = cmdFlow(args)
	case "ingest":
		err = cmdIngest(args)
	case "domains":
		err = cmdDomains()
	case "docs":
		err = cmdDocs(args)
	case "discover":
		err = cmdDiscover(args)
	case "build-config":
		err = cmdBuildConfig(args)
	case "konflux":
		err = cmdKonflux(args)
	case "version-compat":
		err = cmdVersionCompat(args)
	case "platforms":
		err = cmdPlatforms(args)
	case "full-analysis":
		err = cmdFullAnalysis(args)
	case "quick-index":
		err = cmdQuickIndex(args)
	case "sbom":
		err = cmdSBOM(args)
	case "report":
		err = cmdReport(args)
	case "version":
		fmt.Printf("arch-analyzer %s\n", version)
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", cmd)
		printUsage()
		os.Exit(1)
	}

	if err != nil {
		if err == errDiffFound {
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		if cmd == "diff" {
			os.Exit(2)
		}
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`arch-analyzer: Architecture Analyzer and Code Graph Security Scanner

Usage: arch-analyzer <command> [options]

Architecture commands:
  extract <repo-path>                  Extract architecture data from a repository
  render <json-file>                   Render diagrams from architecture JSON
  analyze <repo-path>                  Extract + render in one step
  aggregate <results-dir>              Aggregate multiple component JSONs into platform view
  docs <json-file>                     Generate browsable documentation site from architecture JSON
  flow <json-file>                     Generate interactive flow diagram (flowlens HTML)

Contract validation commands:
  extract-schema <repo-path>           Extract CRD JSON schemas from a repository
  validate <repo-path>                 Validate CRD changes against stored contracts

Code graph commands:
  scan <repo-path>                     Build code graph and run security queries
                                       [--domains d1,d2] [--with-arch] [--import-sarif files]
  context-bundle <repo-path>           Generate SrcLang context bundle for LLM consumption
                                       [--layer security] [--output context.srclg]
                                       [--with-scan] [--domains d1,d2]
  quick-index <repo-path>               Fast function/call index (tree-sitter only, no taint/domains)
                                       [--output file]
  graph <repo-path>                    Export code property graph (JSON or DOT)
  diff <base.json> <head.json>       Structural diff between two code-graph.json files
                                     [--format json|text] [--kind f1,f2] [--output file]
  ingest <sarif-file>                  Ingest external scanner SARIF findings
                                       [--graph code-graph.json] [--output file]
  domains                              List registered analysis domains

Platform commands:
  discover <operator-repo-path>        Discover platform components from kustomize manifests
                                       [--output file] [--format json|text|map]
                                       [--org org] [--platform name]
  build-config <dir>                   Extract build metadata (OCP versions, arches, OLM)
  konflux <snapshot-file-or-dir>       Parse Konflux snapshot image mappings
  platforms <scan-config.yaml>         List platforms defined in scan config
                                       [--platform name] [--output file]
  aggregate-cpg <results-dir>          Merge code graphs into platform-wide CPG
  version-compat <arch.json>           Check API version compatibility against target OCP/k8s
                                       [--target-version ver]

Combined:
  full-analysis <repo-path>            Run architecture extraction + code graph scan
                                       [--domains d1,d2] [--import-sarif files]

Other:
  version                              Print version
  help                                 Show this help`)
}

// cmdExtract extracts architecture data from a repo and writes JSON.
func cmdExtract(args []string) error {
	fs := flag.NewFlagSet("extract", flag.ExitOnError)
	output := fs.String("output", "component-architecture.json", "Output JSON file")
	org := fs.String("org", "", "GitHub organization (auto-detected from go.mod if empty)")
	ver := fs.String("version", "", "Version label for snapshot output (e.g. v2.15.0)")
	aliases := fs.String("aliases", "", "Comma-separated component aliases (e.g. rhods-operator,RHODS)")
	withDeps := fs.Bool("with-deps", false, "Also extract dependencies detected via component_refs")
	scanConfig := fs.String("scan-config", "", "Path to scan-config.yaml for resolving dependency repos")
	extractors := fs.String("extractors", "", "Comma-separated extractor groups to run (default: all). Available: "+strings.Join(extractor.ExtractorGroupNames(), ", "))
	fs.Parse(reorderArgs(fs, args))

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: arch-analyzer extract <repo-path> [--output file.json] [--org org] [--version label] [--aliases list] [--with-deps] [--scan-config config.yaml] [--extractors groups]")
	}

	if *ver != "" {
		if err := validateVersionLabel(*ver); err != nil {
			return err
		}
	}

	opts := &extractor.ExtractOptions{Org: *org, Aliases: parseAliases(*aliases), Extractors: parseAliases(*extractors)}
	arch, err := extractor.ExtractAll(fs.Arg(0), opts)
	if err != nil {
		return err
	}

	outPath := *output
	if *ver != "" {
		dir := filepath.Dir(outPath)
		base := filepath.Base(outPath)
		outPath = filepath.Join(dir, *ver, base)
	}
	if err := writeJSON(outPath, arch); err != nil {
		return err
	}

	if *withDeps && len(arch.ComponentRefs) > 0 {
		return extractDeps(arch, outPath, *org, *scanConfig)
	}
	return nil
}

// extractDeps clones and extracts dependencies found in component_refs.
func extractDeps(primary *extractor.ComponentArchitecture, primaryOutput, org, scanConfigPath string) error {
	repoMap, err := buildRepoMap(scanConfigPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: could not load scan-config for dep resolution: %v\n", err)
		return nil
	}

	outDir := filepath.Dir(primaryOutput)
	scanned := map[string]bool{primary.Component: true}

	for _, ref := range primary.ComponentRefs {
		target := ref.Target
		if scanned[target] {
			continue
		}
		scanned[target] = true

		repoURL, ok := repoMap[target]
		if !ok {
			fmt.Fprintf(os.Stderr, "  Dep %s: no repo mapping found, skipping\n", target)
			continue
		}

		tmpDir, err := os.MkdirTemp("", "arch-dep-"+target+"-")
		if err != nil {
			fmt.Fprintf(os.Stderr, "  Dep %s: failed to create temp dir: %v\n", target, err)
			continue
		}
		defer os.RemoveAll(tmpDir)

		fmt.Fprintf(os.Stderr, "  Dep %s: cloning %s\n", target, repoURL)
		cloneCmd := exec.Command("git", "clone", "--depth", "1", repoURL, tmpDir)
		cloneCmd.Stderr = os.Stderr
		if err := cloneCmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "  Dep %s: clone failed: %v\n", target, err)
			continue
		}

		depOpts := &extractor.ExtractOptions{Org: org}
		depArch, err := extractor.ExtractAll(tmpDir, depOpts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "  Dep %s: extraction failed: %v\n", target, err)
			continue
		}

		depOutput := filepath.Join(outDir, target+"-architecture.json")
		if err := writeJSON(depOutput, depArch); err != nil {
			fmt.Fprintf(os.Stderr, "  Dep %s: write failed: %v\n", target, err)
			continue
		}
		fmt.Fprintf(os.Stderr, "  Dep %s: extracted to %s\n", target, depOutput)
	}
	return nil
}

// buildRepoMap reads scan-config.yaml and builds a component-name → clone-URL map.
func buildRepoMap(scanConfigPath string) (map[string]string, error) {
	if scanConfigPath == "" {
		// Try default location
		scanConfigPath = "scan-config.yaml"
	}
	data, err := os.ReadFile(scanConfigPath)
	if err != nil {
		return nil, err
	}

	// Simple YAML parsing: extract org/repo pairs from the nested structure.
	// Uses line-by-line parsing to avoid a YAML dependency.
	repoMap := make(map[string]string)
	var currentOrg string
	for _, line := range strings.Split(string(data), "\n") {
		trimmed := strings.TrimSpace(line)
		// Detect org keys (lines ending with ":" that are indented under "orgs:")
		if strings.HasSuffix(trimmed, ":") && !strings.Contains(trimmed, " ") && !strings.HasPrefix(trimmed, "#") && !strings.HasPrefix(trimmed, "-") {
			candidate := strings.TrimSuffix(trimmed, ":")
			if strings.Contains(candidate, "-") || strings.Contains(candidate, ".") {
				currentOrg = candidate
			}
		}
		// Detect repo entries (lines starting with "- ")
		if strings.HasPrefix(trimmed, "- ") && currentOrg != "" {
			repo := strings.TrimPrefix(trimmed, "- ")
			repo = strings.TrimSpace(repo)
			if repo != "" && !strings.HasPrefix(repo, "#") && !strings.Contains(repo, ":") {
				url := fmt.Sprintf("https://github.com/%s/%s.git", currentOrg, repo)
				repoMap[repo] = url
			}
		}
	}
	return repoMap, nil
}

// cmdRender renders diagrams from architecture JSON.
func cmdRender(args []string) error {
	fs := flag.NewFlagSet("render", flag.ExitOnError)
	outputDir := fs.String("output-dir", "", "Output directory (default: <json-dir>/diagrams)")
	formats := fs.String("formats", "", "Comma-separated formats: rbac,component,security_network,dependencies,c4,dataflow,report (default: all)")
	fs.Parse(args)

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: arch-analyzer render <json-file> [--output-dir dir] [--formats fmt1,fmt2]")
	}

	jsonPath := fs.Arg(0)
	data, err := loadJSON(jsonPath)
	if err != nil {
		return err
	}

	outDir := *outputDir
	if outDir == "" {
		outDir = filepath.Join(filepath.Dir(jsonPath), "diagrams")
	}

	var fmts []string
	if *formats != "" {
		fmts = strings.Split(*formats, ",")
	}

	diagrams := renderer.RenderAll(data, fmts)
	return writeDiagrams(outDir, diagrams)
}

// cmdAnalyze runs extract + render in one step.
func cmdAnalyze(args []string) error {
	fs := flag.NewFlagSet("analyze", flag.ExitOnError)
	outputDir := fs.String("output-dir", "output", "Output directory")
	org := fs.String("org", "", "GitHub organization (auto-detected from go.mod if empty)")
	ver := fs.String("version", "", "Version label for snapshot output (e.g. v2.15.0)")
	aliases := fs.String("aliases", "", "Comma-separated component aliases (e.g. rhods-operator,RHODS)")
	fs.Parse(args)

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: arch-analyzer analyze <repo-path> [--output-dir dir] [--org org] [--version label] [--aliases list]")
	}

	if *ver != "" {
		if err := validateVersionLabel(*ver); err != nil {
			return err
		}
	}

	repoPath := fs.Arg(0)
	opts := &extractor.ExtractOptions{Org: *org, Aliases: parseAliases(*aliases)}
	arch, err := extractor.ExtractAll(repoPath, opts)
	if err != nil {
		return err
	}

	outDir := applyVersion(*outputDir, *ver)
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("creating output dir: %w", err)
	}

	jsonPath := filepath.Join(outDir, "component-architecture.json")
	if err := writeJSON(jsonPath, arch); err != nil {
		return err
	}
	fmt.Printf("Extracted architecture to: %s\n", jsonPath)

	data, err := loadJSON(jsonPath)
	if err != nil {
		return err
	}

	diagramsDir := filepath.Join(outDir, "diagrams")
	diagrams := renderer.RenderAll(data, nil)
	if err := writeDiagrams(diagramsDir, diagrams); err != nil {
		return err
	}
	fmt.Printf("Rendered %d diagram(s) to: %s\n", len(diagrams), diagramsDir)

	if *ver != "" {
		repos := map[string]string{arch.Repo: arch.CommitSHA}
		if err := writeSnapshotMetadata(outDir, *ver, repos, ""); err != nil {
			return fmt.Errorf("writing snapshot metadata: %w", err)
		}
		fmt.Printf("Snapshot metadata written to: %s/snapshot-metadata.json\n", outDir)
	}

	return nil
}

// cmdAggregate merges multiple component JSONs into platform view.
func cmdAggregate(args []string) error {
	fs := flag.NewFlagSet("aggregate", flag.ExitOnError)
	outputDir := fs.String("output-dir", "platform-output", "Output directory")
	ver := fs.String("version", "", "Version label for snapshot output (e.g. v2.15.0)")
	platform := fs.String("platform", "", "Platform name for snapshot metadata (e.g. rhoai, odh)")
	scanConfig := fs.String("scan-config", "", "Path to scan-config.yaml for platform-aware features (OCP version-compat, overrides)")
	fs.Parse(args)

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: arch-analyzer aggregate <results-dir> [--output-dir dir] [--version label] [--platform name] [--scan-config path]")
	}

	if *ver != "" {
		if err := validateVersionLabel(*ver); err != nil {
			return err
		}
	}

	resultsDir := fs.Arg(0)

	platformData, err := aggregator.Aggregate(resultsDir)
	if err != nil {
		return err
	}

	outDir := applyVersion(*outputDir, *ver)
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("creating output dir: %w", err)
	}

	jsonPath := filepath.Join(outDir, "platform-architecture.json")
	if err := writeJSON(jsonPath, platformData); err != nil {
		return err
	}
	fmt.Printf("Aggregated platform architecture to: %s\n", jsonPath)

	diagramsDir := filepath.Join(outDir, "diagrams")
	diagrams := renderer.RenderPlatformAll(platformData)
	if err := writeDiagrams(diagramsDir, diagrams); err != nil {
		return err
	}
	fmt.Printf("Rendered %d platform diagram(s) to: %s\n", len(diagrams), diagramsDir)

	// CPG aggregation: merge code-graph.json files into platform-wide CPG
	platformCPG, cpgErr := aggregator.AggregateCPGs(resultsDir)
	if cpgErr != nil {
		fmt.Fprintf(os.Stderr, "Warning: CPG aggregation failed: %v\n", cpgErr)
	} else if platformCPG.ComponentCount > 0 {
		// Write summary (without full node/edge arrays) to keep file size manageable
		cpgPath := filepath.Join(outDir, "platform-cpg.json")
		if wErr := writeJSON(cpgPath, platformCPG.Summary()); wErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to write platform CPG: %v\n", wErr)
		} else {
			fmt.Printf("Platform CPG: %d components, %d nodes, %d edges, %d cross-component links\n",
				platformCPG.ComponentCount, platformCPG.TotalNodes, platformCPG.TotalEdges, platformCPG.CrossEdges)
		}
	}

	// Generate per-component markdown docs + INDEX.md + interactions.md
	if err := writeMarkdownDocs(outDir, platformData); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: markdown docs generation failed: %v\n", err)
	}

	// Version compatibility check using platform OCP versions from scan-config
	if *scanConfig != "" && *platform != "" {
		aggregateVersionCompat(resultsDir, outDir, *scanConfig, *platform)
	}

	if *ver != "" {
		repos, err := collectRepoSHAs(resultsDir)
		if err != nil {
			return fmt.Errorf("collecting repo SHAs: %w", err)
		}
		if err := writeSnapshotMetadata(outDir, *ver, repos, *platform); err != nil {
			return fmt.Errorf("writing snapshot metadata: %w", err)
		}
		fmt.Printf("Snapshot metadata written to: %s/snapshot-metadata.json\n", outDir)
	}

	return nil
}

// writeMarkdownDocs generates per-component markdown files, INDEX.md, and
// cross-component/interactions.md. Uses atomic writes: writes to a temp
// directory first, then renames on success.
func writeMarkdownDocs(outDir string, platformData map[string]interface{}) error {
	componentsDir := filepath.Join(outDir, "components")
	crossDir := filepath.Join(outDir, "cross-component")

	// Atomic write: use temp directories, rename on success
	tmpComponentsDir := componentsDir + "-tmp"
	tmpCrossDir := crossDir + "-tmp"

	// Clean up any leftover tmp dirs
	os.RemoveAll(tmpComponentsDir)
	os.RemoveAll(tmpCrossDir)

	if err := os.MkdirAll(tmpComponentsDir, 0o755); err != nil {
		return fmt.Errorf("creating temp components dir: %w", err)
	}
	if err := os.MkdirAll(tmpCrossDir, 0o755); err != nil {
		os.RemoveAll(tmpComponentsDir)
		return fmt.Errorf("creating temp cross-component dir: %w", err)
	}

	// Track filenames for collision detection
	usedFilenames := make(map[string]string) // filename -> component name

	// Render per-component markdown
	componentData := renderer.GetSlice(platformData, "component_data")
	for _, cd := range componentData {
		compName := renderer.GetStr(cd, "component", "unknown")
		filename := renderer.SanitizeFilename(compName) + ".md"

		// Collision detection
		if existingComp, exists := usedFilenames[filename]; exists {
			os.RemoveAll(tmpComponentsDir)
			os.RemoveAll(tmpCrossDir)
			return fmt.Errorf("filename collision: %q and %q both map to %q", existingComp, compName, filename)
		}
		usedFilenames[filename] = compName

		content := renderer.RenderComponentMarkdown(cd)
		filePath := filepath.Join(tmpComponentsDir, filename)
		if err := os.WriteFile(filePath, []byte(content), 0o644); err != nil {
			os.RemoveAll(tmpComponentsDir)
			os.RemoveAll(tmpCrossDir)
			return fmt.Errorf("writing %s: %w", filePath, err)
		}
	}

	// Render INDEX.md to temp file (written alongside atomic swap)
	indexContent := renderer.RenderNavIndex(platformData)
	tmpIndexPath := filepath.Join(outDir, "INDEX.md.tmp")
	if err := os.WriteFile(tmpIndexPath, []byte(indexContent), 0o644); err != nil {
		os.RemoveAll(tmpComponentsDir)
		os.RemoveAll(tmpCrossDir)
		return fmt.Errorf("writing INDEX.md: %w", err)
	}

	// Render interactions.md
	interactionsContent := renderer.RenderInteractions(platformData)
	if err := os.WriteFile(filepath.Join(tmpCrossDir, "interactions.md"), []byte(interactionsContent), 0o644); err != nil {
		os.RemoveAll(tmpComponentsDir)
		os.RemoveAll(tmpCrossDir)
		os.Remove(tmpIndexPath)
		return fmt.Errorf("writing interactions.md: %w", err)
	}

	// Atomic swap: remove old dirs, rename tmp dirs, then rename INDEX.md
	os.RemoveAll(componentsDir)
	os.RemoveAll(crossDir)
	if err := os.Rename(tmpComponentsDir, componentsDir); err != nil {
		os.Remove(tmpIndexPath)
		return fmt.Errorf("renaming components dir: %w", err)
	}
	if err := os.Rename(tmpCrossDir, crossDir); err != nil {
		os.Remove(tmpIndexPath)
		return fmt.Errorf("renaming cross-component dir: %w", err)
	}
	indexPath := filepath.Join(outDir, "INDEX.md")
	if err := os.Rename(tmpIndexPath, indexPath); err != nil {
		return fmt.Errorf("renaming INDEX.md: %w", err)
	}

	fmt.Printf("Generated %d component markdown docs + INDEX.md + interactions.md\n", len(usedFilenames))
	return nil
}

// aggregateVersionCompat runs version compatibility checks against platform OCP versions.
func aggregateVersionCompat(resultsDir, outDir, scanConfigPath, platformName string) {
	_, platCfg, err := config.LoadPlatformConfig(scanConfigPath, platformName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to load platform config for version-compat: %v\n", err)
		return
	}
	if platCfg.OCPVersions == nil || platCfg.OCPVersions.Min == "" {
		fmt.Println("No OCP version constraints defined for platform, skipping version-compat")
		return
	}

	targetVersion := platCfg.OCPVersions.Min
	fmt.Printf("\n=== Version Compatibility Check (target OCP %s) ===\n", targetVersion)

	type componentCompat struct {
		Component string                         `json:"component"`
		Result    *validator.VersionCompatResult `json:"result"`
	}

	var results []componentCompat
	totalIssues := 0

	// Walk results dir for component architecture files
	walkErr := filepath.Walk(resultsDir, func(path string, fi os.FileInfo, walkErr error) error {
		if walkErr != nil || fi.IsDir() || fi.Name() != "component-architecture.json" {
			return nil
		}

		data, err := loadJSON(path)
		if err != nil {
			return nil
		}

		component := filepath.Base(filepath.Dir(path))
		result, err := validator.CheckVersionCompat(data, targetVersion)
		if err != nil {
			fmt.Fprintf(os.Stderr, "  Warning: version-compat for %s: %v\n", component, err)
			return nil
		}

		results = append(results, componentCompat{Component: component, Result: result})
		if len(result.Issues) > 0 {
			totalIssues += len(result.Issues)
			for _, issue := range result.Issues {
				icon := "WARNING"
				if issue.Severity == "error" {
					icon = "ERROR"
				}
				fmt.Printf("  [%s] %s: %s\n", icon, component, issue.Message)
			}
		}
		return nil
	})
	if walkErr != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to walk results dir for version-compat: %v\n", walkErr)
		return
	}

	if totalIssues == 0 {
		fmt.Printf("All %d components compatible with OCP %s\n", len(results), targetVersion)
	} else {
		fmt.Printf("%d issue(s) found across %d components\n", totalIssues, len(results))
	}

	compatPath := filepath.Join(outDir, "version-compat.json")
	output := map[string]interface{}{
		"target_ocp_version": targetVersion,
		"platform":           platformName,
		"components":         results,
		"total_issues":       totalIssues,
	}
	if wErr := writeJSON(compatPath, output); wErr != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to write version-compat: %v\n", wErr)
	} else {
		fmt.Printf("Version compatibility report: %s\n", compatPath)
	}
}

// cmdDocs generates a browsable documentation site from architecture JSON.
func cmdDocs(args []string) error {
	fs := flag.NewFlagSet("docs", flag.ExitOnError)
	outputDir := fs.String("output-dir", "docs", "Output directory for generated docs")
	prefix := fs.String("prefix", "", "Path prefix for nav snippet (e.g. 'rhoai-platform')")
	ver := fs.String("version", "", "Version label for snapshot output (e.g. v2.15.0)")
	fs.Parse(args)

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: arch-analyzer docs <json-file> [--output-dir dir] [--prefix path] [--version label]")
	}

	if *ver != "" {
		if err := validateVersionLabel(*ver); err != nil {
			return err
		}
	}

	data, err := loadJSON(fs.Arg(0))
	if err != nil {
		return err
	}

	pages := renderer.RenderDocs(data)

	// Look for snapshot metadata next to the input JSON file
	var banner string
	metaPath := filepath.Join(filepath.Dir(fs.Arg(0)), "snapshot-metadata.json")
	if metaData, err := loadJSON(metaPath); err == nil {
		snapVer, _ := metaData["version"].(string)
		snapTS, _ := metaData["timestamp"].(string)
		if snapTS != "" {
			// Truncate to date only for display
			if len(snapTS) >= 10 {
				snapTS = snapTS[:10]
			}
		}
		if snapVer != "" {
			banner = fmt.Sprintf("> **Architecture snapshot: %s** (%s)\n\n", snapVer, snapTS)
		}
	}

	outDir := applyVersion(*outputDir, *ver)
	for _, page := range pages {
		content := page.Content
		if banner != "" && strings.HasSuffix(page.Path, "index.md") {
			// Inject version banner after the first heading
			if idx := strings.Index(content, "\n"); idx > 0 {
				content = content[:idx+1] + "\n" + banner + content[idx+1:]
			}
		}
		outPath := filepath.Join(outDir, page.Path)
		if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
			return fmt.Errorf("creating directory for %s: %w", outPath, err)
		}
		if err := os.WriteFile(outPath, []byte(content), 0o644); err != nil {
			return fmt.Errorf("writing %s: %w", outPath, err)
		}
	}

	fmt.Printf("Generated %d documentation pages to: %s\n", len(pages), outDir)

	// Print nav snippet
	navSnippet := renderer.NavSnippet(pages, *prefix)
	if navSnippet != "" {
		fmt.Println("\nmkdocs.yml nav snippet:")
		fmt.Println(navSnippet)
	}

	return nil
}

// cmdExtractSchema extracts CRD JSON schemas from a repo.
func cmdExtractSchema(args []string) error {
	fs := flag.NewFlagSet("extract-schema", flag.ExitOnError)
	outputDir := fs.String("output-dir", "contracts/schemas", "Output directory for schemas")
	repoName := fs.String("repo-name", "", "Repository name (default: directory name)")
	fs.Parse(args)

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: arch-analyzer extract-schema <repo-path> [--output-dir dir] [--repo-name name]")
	}

	repoPath := fs.Arg(0)
	schemas, err := validator.ExtractSchemasFromDir(repoPath)
	if err != nil {
		return err
	}

	if len(schemas) == 0 {
		fmt.Printf("No CRD schemas found in %s\n", repoPath)
		return nil
	}

	name := *repoName
	if name == "" {
		name = filepath.Base(repoPath)
	}

	schemaDir := filepath.Join(*outputDir, name)
	if err := os.MkdirAll(schemaDir, 0o755); err != nil {
		return fmt.Errorf("creating schema dir: %w", err)
	}

	for _, s := range schemas {
		outPath := filepath.Join(schemaDir, s.ResourceKey+".json")
		if err := writeJSON(outPath, s.Schema); err != nil {
			return fmt.Errorf("writing schema %s: %w", s.ResourceKey, err)
		}
		fmt.Printf("Extracted: %s -> %s\n", s.ResourceKey, outPath)
	}

	fmt.Printf("Extracted %d schema(s) to %s\n", len(schemas), schemaDir)
	return nil
}

// cmdValidate validates CRD changes against stored contracts.
func cmdValidate(args []string) error {
	fs := flag.NewFlagSet("validate", flag.ExitOnError)
	contractsDir := fs.String("contracts-dir", "contracts", "Path to contracts directory")
	repoName := fs.String("repo-name", "", "Repository name (default: directory name)")
	fs.Parse(args)

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: arch-analyzer validate <repo-path> [--contracts-dir dir] [--repo-name name]")
	}

	repoPath := fs.Arg(0)
	schemas, err := validator.ExtractSchemasFromDir(repoPath)
	if err != nil {
		return err
	}

	if len(schemas) == 0 {
		fmt.Printf("No CRD schemas found in %s, nothing to validate\n", repoPath)
		return nil
	}

	name := *repoName
	if name == "" {
		name = filepath.Base(repoPath)
	}

	result, err := validator.CheckContract(name, schemas, *contractsDir)
	if err != nil {
		return err
	}

	fmt.Printf("\n%s\n", strings.Repeat("=", 60))
	fmt.Printf("Contract Validation: %s\n", name)
	fmt.Printf("%s\n", strings.Repeat("=", 60))

	for _, check := range result.Checks {
		symbol := "v"
		status := "PASS"
		if !check.IsCompatible {
			symbol = "X"
			status = "FAIL"
		}
		fmt.Printf("  [%s] %s: %s\n", symbol, check.Resource, status)
		for _, d := range check.Details {
			fmt.Printf("      - %s\n", d.Description)
		}
		if len(check.Consumers) > 0 {
			fmt.Printf("      Consumers: %s\n", strings.Join(check.Consumers, ", "))
		}
	}

	if len(result.AffectedConsumers) > 0 {
		fmt.Printf("\nAFFECTED CONSUMERS:\n")
		for _, c := range result.AffectedConsumers {
			fmt.Printf("  - %s: %s\n", c.Repo, c.Usage)
			for _, bc := range c.BreakingChanges {
				fmt.Printf("      Breaking: %s\n", bc)
			}
		}
	}

	if result.IsCompatible {
		fmt.Printf("\nResult: COMPATIBLE\n")
		return nil
	}
	fmt.Printf("\nResult: BREAKING CHANGES DETECTED\n")
	return fmt.Errorf("breaking changes detected")
}

// cmdScan builds a code property graph and runs security queries.

func cmdFullAnalysis(args []string) error {
	fs := flag.NewFlagSet("full-analysis", flag.ExitOnError)
	outputDir := fs.String("output-dir", "output", "Output directory")
	org := fs.String("org", "", "GitHub organization (auto-detected from go.mod if empty)")
	domainList := fs.String("domains", "", "Comma-separated domains (default: all)")
	ver := fs.String("version", "", "Version label for snapshot output (e.g. v2.15.0)")
	importSARIF := fs.String("import-sarif", "", "Comma-separated SARIF files to ingest after building graph")
	aliases := fs.String("aliases", "", "Comma-separated component aliases (e.g. rhods-operator,RHODS)")
	extractorsList := fs.String("extractors", "", "Comma-separated extractor groups to run (default: all). Available: "+strings.Join(extractor.ExtractorGroupNames(), ", "))
	fs.Parse(reorderArgs(fs, args))

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: arch-analyzer full-analysis <repo-path> [--output-dir dir] [--org org] [--domains sec,test] [--extractors groups] [--version label] [--aliases list]")
	}

	if *ver != "" {
		if err := validateVersionLabel(*ver); err != nil {
			return err
		}
	}

	repoPath := fs.Arg(0)
	outDir := applyVersion(*outputDir, *ver)
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("creating output dir: %w", err)
	}

	// Architecture extraction
	fmt.Println("=== Architecture Extraction ===")
	extractOpts := &extractor.ExtractOptions{Org: *org, Aliases: parseAliases(*aliases), Extractors: parseAliases(*extractorsList)}
	archResult, err := extractor.ExtractAll(repoPath, extractOpts)
	var archData *domains.ArchitectureData
	var parsedArch *arch.Data
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: architecture extraction failed: %v\n", err)
	} else {
		jsonPath := filepath.Join(outDir, "component-architecture.json")
		if err := writeJSON(jsonPath, archResult); err != nil {
			return err
		}
		fmt.Printf("Extracted architecture to: %s\n", jsonPath)

		// Prepare arch data for domain analyzers
		raw, mErr := json.Marshal(archResult)
		if mErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to marshal architecture data: %v\n", mErr)
		}
		var data map[string]interface{}
		if err := json.Unmarshal(raw, &data); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to unmarshal architecture data: %v\n", err)
		}
		archData = &domains.ArchitectureData{}

		parsed, parseErr := arch.Parse(data)
		if parseErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: architecture data parsing failed: %v\n", parseErr)
		} else {
			parsedArch = parsed
		}

		data2, loadErr := loadJSON(jsonPath)
		if loadErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to load architecture JSON: %v\n", loadErr)
		} else if data2 != nil {
			diagramsDir := filepath.Join(outDir, "diagrams")
			diagrams := renderer.RenderAll(data2, nil)
			if wErr := writeDiagrams(diagramsDir, diagrams); wErr != nil {
				fmt.Fprintf(os.Stderr, "Warning: failed to write diagrams: %v\n", wErr)
			} else {
				fmt.Printf("Rendered %d diagram(s) to: %s\n", len(diagrams), diagramsDir)
			}
		}
	}

	// Build config extraction
	fmt.Println("\n=== Build Config Extraction ===")
	bc, bcErr := extractor.ParseBuildConfig(repoPath)
	if bcErr != nil {
		fmt.Fprintf(os.Stderr, "Warning: build config extraction failed: %v\n", bcErr)
	} else {
		bcPath := filepath.Join(outDir, "build-config.json")
		if wErr := writeJSON(bcPath, bc); wErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to write build config: %v\n", wErr)
		} else {
			fmt.Printf("Build config written to: %s\n", bcPath)
			if bc.OCPVersions.Min != "" || bc.OCPVersions.Max != "" {
				fmt.Printf("  OCP versions: %s - %s\n", bc.OCPVersions.Min, bc.OCPVersions.Max)
			}
			if len(bc.Architectures) > 0 {
				fmt.Printf("  Architectures: %s\n", strings.Join(bc.Architectures, ", "))
			}
		}
	}

	// Code graph scan with domains
	fmt.Println("\n=== Code Graph Security Scan ===")
	cpg, err := buildCPG(repoPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: code graph build failed: %v\n", err)
	} else {
		if parsedArch != nil {
			cpg.ArchData = parsedArch
		}

		findings, scanErr := runSecurityScan(cpg, *domainList, archData)
		if scanErr != nil {
			return scanErr
		}

		if *importSARIF != "" {
			externalFindings, sarifErr := ingestSARIFFiles(cpg, *importSARIF)
			if sarifErr != nil {
				fmt.Fprintf(os.Stderr, "Warning: SARIF ingestion: %v\n", sarifErr)
			}
			findings = append(findings, externalFindings...)
		}

		if archResult != nil {
			findings = append(findings, securityAnnotationsToFindings(archResult.SecurityAnnotations)...)
		}

		printFindings(cpg, findings)

		findingsPath := filepath.Join(outDir, "security-findings.json")
		if wErr := outputJSON(findingsPath, findings); wErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to write findings: %v\n", wErr)
		} else {
			fmt.Printf("Findings written to: %s\n", findingsPath)
		}

		graphPath := filepath.Join(outDir, "code-graph.json")
		graphData := map[string]interface{}{
			"schema_version": graph.SchemaVersion,
			"nodes":          cpg.Nodes(),
			"edges":          cpg.Edges(),
		}
		if wErr := writeJSON(graphPath, graphData); wErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to write code graph: %v\n", wErr)
		} else {
			fmt.Printf("Code graph written to: %s\n", graphPath)
		}
	}

	// Schema extraction
	fmt.Println("\n=== CRD Schema Extraction ===")
	schemas, err := validator.ExtractSchemasFromDir(repoPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: schema extraction failed: %v\n", err)
	} else if len(schemas) > 0 {
		schemaDir := filepath.Join(outDir, "schemas")
		if mkErr := os.MkdirAll(schemaDir, 0o755); mkErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to create schema dir: %v\n", mkErr)
		} else {
			for _, s := range schemas {
				outPath := filepath.Join(schemaDir, s.ResourceKey+".json")
				if wErr := writeJSON(outPath, s.Schema); wErr != nil {
					fmt.Fprintf(os.Stderr, "Warning: failed to write schema %s: %v\n", s.ResourceKey, wErr)
				}
			}
			fmt.Printf("Extracted %d CRD schema(s) to: %s\n", len(schemas), schemaDir)
		}
	} else {
		fmt.Println("No CRD schemas found")
	}

	if *ver != "" && archResult != nil {
		repos := map[string]string{archResult.Repo: archResult.CommitSHA}
		if err := writeSnapshotMetadata(outDir, *ver, repos, ""); err != nil {
			return fmt.Errorf("writing snapshot metadata: %w", err)
		}
		fmt.Printf("Snapshot metadata written to: %s/snapshot-metadata.json\n", outDir)
	}

	return nil
}

// reorderArgs moves positional arguments after flags so that flag.FlagSet.Parse
// works regardless of argument order. Without this, flags placed after the first
// positional argument are silently ignored by Go's flag package.
func reorderArgs(fs *flag.FlagSet, args []string) []string {
	boolFlags := map[string]bool{}
	fs.VisitAll(func(f *flag.Flag) {
		if bf, ok := f.Value.(interface{ IsBoolFlag() bool }); ok && bf.IsBoolFlag() {
			boolFlags[f.Name] = true
		}
	})

	var flags, positional []string
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if !strings.HasPrefix(arg, "-") {
			positional = append(positional, arg)
			continue
		}
		name := strings.TrimLeft(arg, "-")
		if eq := strings.Index(name, "="); eq >= 0 {
			name = name[:eq]
		}
		flags = append(flags, arg)
		if !boolFlags[name] && !strings.Contains(arg, "=") && i+1 < len(args) {
			i++
			flags = append(flags, args[i])
		}
	}
	return append(flags, positional...)
}

var secAnnotationRulePrefix = map[string]string{
	"RBAC_CLUSTER_SCOPE_SENSITIVE": "SEC-RBAC",
	"ROUTE_NO_TLS":                "SEC-ROUTE",
	"SECRET_IN_CONTAINER_ARGS":    "SEC-SECRET",
	"CRD_CONFUSED_DEPUTY":         "SEC-CRD",
	"MISSING_AUTH_REQUIREMENT":     "SEC-AUTH",
}

func securityAnnotationsToFindings(annotations []extractor.SecurityAnnotation) []query.Finding {
	counters := make(map[string]int)
	var findings []query.Finding
	for _, a := range annotations {
		prefix := secAnnotationRulePrefix[a.Type]
		if prefix == "" {
			prefix = "SEC-EVAL"
		}
		counters[prefix]++
		findings = append(findings, query.Finding{
			RuleID:          fmt.Sprintf("%s-%03d", prefix, counters[prefix]),
			Severity:        a.Severity,
			Message:         a.Description,
			File:            a.Source,
			Domain:          "security",
			ArchitectureRef: "security_annotations:" + a.Type,
		})
	}
	return findings
}
