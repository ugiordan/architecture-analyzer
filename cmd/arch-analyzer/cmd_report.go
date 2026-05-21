package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ugiordan/architecture-analyzer/pkg/report"
	"github.com/ugiordan/architecture-analyzer/pkg/sbom"
)

func cmdSBOM(args []string) error {
	fs := flag.NewFlagSet("sbom", flag.ExitOnError)
	output := fs.String("output", "", "Output file (default: stdout)")
	fs.Parse(args)

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: arch-analyzer sbom <component-architecture.json> [--output file.json]")
	}

	data, err := loadJSON(fs.Arg(0))
	if err != nil {
		return err
	}

	bomJSON, err := sbom.GenerateJSON(data)
	if err != nil {
		return fmt.Errorf("generating SBOM: %w", err)
	}

	if *output != "" {
		if err := os.MkdirAll(filepath.Dir(*output), 0o755); err != nil {
			return err
		}
		return os.WriteFile(*output, bomJSON, 0o644)
	}

	_, err = os.Stdout.Write(bomJSON)
	fmt.Fprintln(os.Stdout)
	return err
}

func cmdReport(args []string) error {
	fs := flag.NewFlagSet("report", flag.ExitOnError)
	output := fs.String("output", "", "Output file (default: stdout)")
	fs.Parse(args)

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: arch-analyzer report <json-file>... [--output file.md]\n  Accepts one or more component-architecture.json files for cross-component analysis")
	}

	components := make(map[string]map[string]interface{})
	for _, jsonPath := range fs.Args() {
		data, err := loadJSON(jsonPath)
		if err != nil {
			return fmt.Errorf("loading %s: %w", jsonPath, err)
		}
		name := ""
		if c, ok := data["component"].(string); ok && c != "" {
			name = c
		} else {
			name = filepath.Base(filepath.Dir(jsonPath))
		}
		components[name] = data
	}

	md := report.GenerateImageReport(components)

	if *output != "" {
		if err := os.MkdirAll(filepath.Dir(*output), 0o755); err != nil {
			return err
		}
		return os.WriteFile(*output, []byte(md), 0o644)
	}

	fmt.Print(md)
	return nil
}

// cmdFullAnalysis runs architecture extraction + code graph scan.
