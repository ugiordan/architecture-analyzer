package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ugiordan/architecture-analyzer/pkg/config"
	"github.com/ugiordan/architecture-analyzer/pkg/renderer"
)

func cmdTierReport(args []string) error {
	fs := flag.NewFlagSet("tier-report", flag.ExitOnError)
	tiersFile := fs.String("tiers", "api_tiers.yaml", "Path to api_tiers.yaml tier mapping file")
	output := fs.String("output", "", "Output file (default: stdout)")
	format := fs.String("format", "markdown", "Output format: markdown, csv, json")
	fs.Parse(reorderArgs(fs, args))

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: arch-analyzer tier-report <platform-json-or-dir> --tiers api_tiers.yaml [--format markdown|csv|json] [--output file]")
	}

	tierCfg, err := config.LoadAPITiersConfig(*tiersFile)
	if err != nil {
		return fmt.Errorf("loading tier config: %w", err)
	}

	inputPath := fs.Arg(0)

	// Detect whether input is a JSON file or a directory containing platform-architecture.json
	info, err := os.Stat(inputPath)
	if err != nil {
		return fmt.Errorf("input path: %w", err)
	}
	jsonPath := inputPath
	if info.IsDir() {
		jsonPath = inputPath + "/platform-architecture.json"
		if _, err := os.Stat(jsonPath); err != nil {
			return fmt.Errorf("no platform-architecture.json found in %s", inputPath)
		}
	}

	platformData, err := loadJSON(jsonPath)
	if err != nil {
		return fmt.Errorf("loading platform data: %w", err)
	}

	report := renderer.BuildTierReport(platformData, tierCfg)

	var content string
	switch *format {
	case "markdown", "md":
		content = renderer.RenderTierReportMarkdown(report)
	case "csv":
		content = renderer.RenderTierReportCSV(report)
	case "json":
		content, err = renderer.RenderTierReportJSON(report)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown format %q (supported: markdown, csv, json)", *format)
	}

	if *output != "" {
		if err := os.WriteFile(*output, []byte(content), 0o644); err != nil {
			return fmt.Errorf("writing output: %w", err)
		}
		fmt.Fprintf(os.Stderr, "Tier report written to: %s\n", *output)
	} else {
		fmt.Print(content)
	}

	// Print coverage summary to stderr
	gapCount := len(report.Gaps.MissingTier)
	staleCount := len(report.Gaps.StaleEntries)
	totalCRDs := 0
	for _, e := range report.Entries {
		if e.Category == "CRD" {
			totalCRDs++
		}
	}
	covered := totalCRDs - gapCount
	fmt.Fprintf(os.Stderr, "\nCoverage: %d/%d CRDs have tier assignments", covered, totalCRDs)
	if gapCount > 0 {
		fmt.Fprintf(os.Stderr, " (%d missing)", gapCount)
	}
	if staleCount > 0 {
		fmt.Fprintf(os.Stderr, ", %d stale entries in tier config", staleCount)
	}
	fmt.Fprintln(os.Stderr)

	return nil
}
