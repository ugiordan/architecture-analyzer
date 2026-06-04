package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ugiordan/architecture-analyzer/pkg/flow"
	"github.com/ugiordan/architecture-analyzer/pkg/renderer"
)

func cmdFlow(args []string) error {
	fs := flag.NewFlagSet("flow", flag.ExitOnError)
	output := fs.String("o", "flow.html", "Output file path")
	diagramOnly := fs.Bool("diagram", false, "Output flowlens diagram JSON only, no HTML")
	title := fs.String("title", "", "Override diagram title")
	fs.Parse(args)

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: arch-analyzer flow <component-architecture.json> [-o flow.html] [-diagram] [-title title]")
	}

	data, err := loadJSON(fs.Arg(0))
	if err != nil {
		return fmt.Errorf("loading %s: %w", fs.Arg(0), err)
	}

	// Build request-path flow graph
	g := renderer.BuildFlowGraph(data)

	// Add reconciliation flows
	flow.AddReconcileFlows(&g, data)

	// Convert to flowlens diagram
	d := flow.ConvertDiagram(g, data)

	if *title != "" {
		d.Meta.Title = *title
	}

	if *diagramOnly {
		out, err := json.MarshalIndent(d, "", "  ")
		if err != nil {
			return fmt.Errorf("marshaling diagram: %w", err)
		}
		if *output != "flow.html" {
			if err := os.MkdirAll(filepath.Dir(*output), 0o755); err != nil {
				return err
			}
			return os.WriteFile(*output, out, 0o644)
		}
		fmt.Println(string(out))
		return nil
	}

	html, err := flow.GenerateHTML(d)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(*output), 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(*output, []byte(html), 0o644); err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "flowlens diagram written to %s\n", *output)
	return nil
}
