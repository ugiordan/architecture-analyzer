package flow

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/ugiordan/architecture-analyzer/pkg/renderer"
)

func TestIntegration_RealComponentArchitecture(t *testing.T) {
	testFile := "../../results/kserve/component-architecture.json"
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Skip("test fixture not found: " + testFile)
	}

	raw, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("reading test fixture: %v", err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(raw, &data); err != nil {
		t.Fatalf("parsing JSON: %v", err)
	}

	// Build full pipeline
	g := renderer.BuildFlowGraph(data)
	AddReconcileFlows(&g, data)
	d := ConvertDiagram(g, data)

	// Meta
	if d.Meta.Title == "" {
		t.Error("diagram should have a title")
	}
	t.Logf("Title: %s", d.Meta.Title)

	// Nodes
	if len(d.Nodes) == 0 {
		t.Error("diagram should have nodes")
	}
	t.Logf("Nodes: %d", len(d.Nodes))

	// Flows
	if len(d.Flows) == 0 {
		t.Error("diagram should have at least one flow")
	}
	t.Logf("Flows: %d", len(d.Flows))
	for fid, f := range d.Flows {
		t.Logf("  Flow %q: %d steps", fid, len(f.Steps))
	}

	// All nodes should have positions (from layout)
	for id, n := range d.Nodes {
		if n.X == 0 && n.Y == 0 && n.Type != "container" {
			t.Errorf("node %q (%s) has no position", id, n.Label)
		}
	}

	// All flow steps should reference valid nodes
	for fid, f := range d.Flows {
		for i, step := range f.Steps {
			if _, ok := d.Nodes[step.From]; !ok {
				t.Errorf("flow %q step %d references missing from node %q", fid, i, step.From)
			}
			if _, ok := d.Nodes[step.To]; !ok {
				t.Errorf("flow %q step %d references missing to node %q", fid, i, step.To)
			}
		}
	}

	// Canvas should be reasonably sized
	if d.Canvas.Width < 1400 {
		t.Errorf("canvas width %d < 1400", d.Canvas.Width)
	}
	if d.Canvas.Height < 900 {
		t.Errorf("canvas height %d < 900", d.Canvas.Height)
	}

	// Legend should exist
	if len(d.Legend) == 0 {
		t.Error("diagram should have legend entries")
	}

	// Mode should be set
	if d.Mode != "live" && d.Mode != "play" {
		t.Errorf("mode should be 'live' or 'play', got %q", d.Mode)
	}
	t.Logf("Mode: %s", d.Mode)

	// FlowOrder should match flows
	if len(d.FlowOrder) != len(d.Flows) {
		t.Errorf("flowOrder length %d != flows length %d", len(d.FlowOrder), len(d.Flows))
	}

	// Generate HTML
	html, err := GenerateHTML(d)
	if err != nil {
		t.Fatalf("GenerateHTML: %v", err)
	}
	if !strings.Contains(html, "FlowLens") {
		t.Error("HTML should contain FlowLens bundle")
	}
	if !strings.Contains(html, "Content-Security-Policy") {
		t.Error("HTML should contain CSP")
	}
	if !strings.Contains(html, d.Meta.Title) {
		t.Error("HTML should contain diagram title")
	}
	t.Logf("HTML size: %d bytes", len(html))
}

func TestIntegration_MultipleOperators(t *testing.T) {
	fixtures := []string{
		"../../results/kserve/component-architecture.json",
		"../../results/codeflare-operator/component-architecture.json",
		"../../results/odh-model-controller/component-architecture.json",
	}

	for _, testFile := range fixtures {
		name := strings.TrimPrefix(testFile, "../../results/")
		name = strings.TrimSuffix(name, "/component-architecture.json")

		t.Run(name, func(t *testing.T) {
			if _, err := os.Stat(testFile); os.IsNotExist(err) {
				t.Skip("fixture not found: " + testFile)
			}

			raw, err := os.ReadFile(testFile)
			if err != nil {
				t.Fatalf("reading: %v", err)
			}

			var data map[string]interface{}
			if err := json.Unmarshal(raw, &data); err != nil {
				t.Fatalf("parsing: %v", err)
			}

			g := renderer.BuildFlowGraph(data)
			AddReconcileFlows(&g, data)
			d := ConvertDiagram(g, data)

			// Should not panic and should produce valid diagram
			if d.Meta.Title == "" {
				t.Error("missing title")
			}
			if len(d.Nodes) == 0 {
				t.Error("no nodes")
			}

			// All step references should be valid
			for fid, f := range d.Flows {
				for i, step := range f.Steps {
					if _, ok := d.Nodes[step.From]; !ok {
						t.Errorf("flow %q step %d: missing from %q", fid, i, step.From)
					}
					if _, ok := d.Nodes[step.To]; !ok {
						t.Errorf("flow %q step %d: missing to %q", fid, i, step.To)
					}
				}
			}

			// HTML generation should succeed
			html, err := GenerateHTML(d)
			if err != nil {
				t.Fatalf("GenerateHTML: %v", err)
			}
			if len(html) < 1000 {
				t.Error("HTML too small")
			}

			t.Logf("%s: %d nodes, %d flows, %d bytes HTML", name, len(d.Nodes), len(d.Flows), len(html))
		})
	}
}
