package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ugiordan/architecture-analyzer/pkg/aggregator"
	"github.com/ugiordan/architecture-analyzer/pkg/config"
	"bufio"
	"github.com/ugiordan/architecture-analyzer/pkg/extractor"
	"github.com/ugiordan/architecture-analyzer/pkg/validator"
)

func cmdDiscover(args []string) error {
	fs := flag.NewFlagSet("discover", flag.ExitOnError)
	output := fs.String("output", "", "Output file (default: stdout)")
	format := fs.String("format", "text", "Output format: text, json, map")
	org := fs.String("org", "", "GitHub organization for repo URLs")
	platform := fs.String("platform", "", "Platform name (e.g. rhoai, odh)")
	fs.Parse(args)

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: arch-analyzer discover <operator-repo-path> [--output file] [--format text|json|map] [--org org] [--platform name]")
	}

	repoPath := fs.Arg(0)
	discovery, err := extractor.DiscoverPlatformComponents(repoPath)
	if err != nil {
		return err
	}

	orgName := *org
	if orgName == "" {
		// Auto-detect from the repo
		orgName = detectOrgFromRepo(repoPath)
	}

	switch *format {
	case "text":
		fmt.Print(discovery.FormatSummary())
		return nil
	case "json":
		return outputJSON(*output, discovery)
	case "map":
		cm := extractor.BuildComponentMap(discovery, orgName)
		if *platform != "" {
			cm.Platform = *platform
		}
		return outputJSON(*output, cm)
	default:
		return fmt.Errorf("unknown format: %s", *format)
	}
}

// detectOrgFromRepo tries to detect the GitHub org from a repo path.
func detectOrgFromRepo(repoPath string) string {
	absPath, err := filepath.Abs(repoPath)
	if err != nil {
		return ""
	}
	opts := &extractor.ExtractOptions{}
	// Use the same detection logic as extract
	_ = opts
	// Simple: check go.mod
	goModPath := filepath.Join(absPath, "go.mod")
	if f, err := os.Open(goModPath); err == nil {
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if strings.HasPrefix(line, "module ") {
				module := strings.TrimPrefix(line, "module ")
				parts := strings.Split(strings.TrimSpace(module), "/")
				if len(parts) >= 2 && parts[0] == "github.com" {
					return parts[1]
				}
			}
		}
	}
	return ""
}


func cmdBuildConfig(args []string) error {
	fs := flag.NewFlagSet("build-config", flag.ExitOnError)
	output := fs.String("output", "", "Output file (default: stdout)")
	fs.Parse(args)

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: arch-analyzer build-config <dir> [--output file]")
	}

	bc, err := extractor.ParseBuildConfig(fs.Arg(0))
	if err != nil {
		return err
	}

	return outputJSON(*output, bc)
}

// cmdKonflux parses Konflux snapshot files.
func cmdKonflux(args []string) error {
	fs := flag.NewFlagSet("konflux", flag.ExitOnError)
	output := fs.String("output", "", "Output file (default: stdout)")
	format := fs.String("format", "json", "Output format: json, text")
	fs.Parse(args)

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: arch-analyzer konflux <snapshot-file-or-dir> [--output file] [--format json|text]")
	}

	path := fs.Arg(0)
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("path not found: %w", err)
	}

	if info.IsDir() {
		idx, err := extractor.ParseKonfluxDir(path)
		if err != nil {
			return err
		}
		switch *format {
		case "text":
			fmt.Printf("Parsed %d snapshot(s), %d unique images\n", idx.Snapshots, len(idx.Components()))
			for _, c := range idx.Components() {
				fmt.Printf("  %-30s %s\n", c.Name, c.ContainerImage)
				if c.Repository != "" {
					fmt.Printf("  %-30s -> %s@%s\n", "", c.Repository, c.Revision)
				}
			}
			return nil
		default:
			return outputJSON(*output, idx)
		}
	}

	snap, err := extractor.ParseKonfluxSnapshot(path)
	if err != nil {
		return err
	}
	switch *format {
	case "text":
		fmt.Printf("Application: %s (%d components)\n", snap.Application, len(snap.Components))
		for _, c := range snap.Components {
			fmt.Printf("  %-30s %s\n", c.Name, c.ContainerImage)
			if c.Repository != "" {
				fmt.Printf("  %-30s -> %s@%s\n", "", c.Repository, c.Revision)
			}
		}
		return nil
	default:
		return outputJSON(*output, snap)
	}
}

// cmdAggregateCPG merges code graphs from multiple components.
func cmdAggregateCPG(args []string) error {
	fs := flag.NewFlagSet("aggregate-cpg", flag.ExitOnError)
	output := fs.String("output", "", "Output file (default: stdout)")
	fs.Parse(args)

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: arch-analyzer aggregate-cpg <results-dir> [--output file]")
	}

	platform, err := aggregator.AggregateCPGs(fs.Arg(0))
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Merged CPG: %d components, %d nodes, %d edges, %d cross-component links\n",
		platform.ComponentCount, platform.TotalNodes, platform.TotalEdges, platform.CrossEdges)

	return outputJSON(*output, platform)
}

// cmdVersionCompat checks API version compatibility against a target OCP/Kubernetes version.
func cmdVersionCompat(args []string) error {
	fs := flag.NewFlagSet("version-compat", flag.ExitOnError)
	targetVersion := fs.String("target-version", "4.14", "Target OCP or Kubernetes version (e.g. 4.14, 1.27)")
	output := fs.String("output", "", "Output file (default: stdout)")
	format := fs.String("format", "text", "Output format: text, json")
	fs.Parse(args)

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: arch-analyzer version-compat <arch.json> [--target-version ver] [--output file] [--format text|json]")
	}

	data, err := loadJSON(fs.Arg(0))
	if err != nil {
		return err
	}

	result, err := validator.CheckVersionCompat(data, *targetVersion)
	if err != nil {
		return err
	}

	switch *format {
	case "text":
		fmt.Printf("Version Compatibility Check: target %s (Kubernetes %s)\n\n", result.TargetVersion, result.KubeVersion)
		if len(result.Issues) == 0 {
			fmt.Println("No compatibility issues found.")
		} else {
			for _, issue := range result.Issues {
				icon := "WARNING"
				if issue.Severity == "error" {
					icon = "ERROR"
				}
				fmt.Printf("  [%s] %s\n", icon, issue.Message)
				fmt.Printf("          Source: %s\n", issue.Source)
				if issue.Replacement != "" {
					fmt.Printf("          Replace with: %s\n", issue.Replacement)
				}
			}
		}
		if result.Compatible {
			fmt.Println("\nResult: COMPATIBLE")
		} else {
			fmt.Println("\nResult: INCOMPATIBLE")
		}
		return nil
	default:
		return outputJSON(*output, result)
	}
}

// cmdPlatforms lists or queries platform definitions from scan-config.yaml.
func cmdPlatforms(args []string) error {
	fs := flag.NewFlagSet("platforms", flag.ExitOnError)
	platform := fs.String("platform", "", "Show repos for a specific platform")
	output := fs.String("output", "", "Output file (default: stdout)")
	fs.Parse(args)

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: arch-analyzer platforms <scan-config.yaml> [--platform name] [--output file]")
	}

	configPath := fs.Arg(0)

	if *platform == "" {
		names, err := config.ListPlatforms(configPath)
		if err != nil {
			return err
		}
		fmt.Printf("Platforms defined in %s:\n", configPath)
		for _, name := range names {
			fmt.Printf("  - %s\n", name)
		}
		return nil
	}

	specs, platCfg, err := config.LoadPlatformConfig(configPath, *platform)
	if err != nil {
		return err
	}

	result := map[string]interface{}{
		"platform":    *platform,
		"name":        platCfg.Name,
		"description": platCfg.Description,
		"repo_count":  len(specs),
		"repos":       specs,
	}
	if platCfg.OCPVersions != nil {
		result["ocp_versions"] = platCfg.OCPVersions
	}

	return outputJSON(*output, result)
}

// collectRepoSHAs scans a results directory for component-architecture.json files
// and extracts repo -> commit_sha pairs for snapshot metadata.
// Supports both flat (results/<repo>/) and org-namespaced (results/<org>/<repo>/) layouts.
func collectRepoSHAs(resultsDir string) (map[string]string, error) {
	repos := make(map[string]string)
	err := filepath.Walk(resultsDir, func(path string, fi os.FileInfo, walkErr error) error {
		if walkErr != nil || fi.IsDir() || fi.Name() != "component-architecture.json" {
			return nil
		}
		data, loadErr := loadJSON(path)
		if loadErr != nil {
			fmt.Printf("WARN: failed to load %s: %v\n", path, loadErr)
			return nil
		}
		repo, _ := data["repo"].(string)
		sha, _ := data["commit_sha"].(string)
		if repo != "" {
			repos[repo] = sha
		}
		return nil
	})
	if err != nil {
		return repos, fmt.Errorf("scanning results directory %s: %w", resultsDir, err)
	}
	return repos, nil
}

// runSecurityScan runs domain annotations, taint propagation, legacy queries,
// and domain queries against a CPG, returning all findings.
// If domainList is non-empty, only those domains run.
