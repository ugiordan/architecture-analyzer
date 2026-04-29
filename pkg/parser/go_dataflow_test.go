package parser

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/ugiordan/architecture-analyzer/pkg/graph"
)

func filterEdgesByLabel(edges []*graph.Edge, label string) []*graph.Edge {
	var out []*graph.Edge
	for _, e := range edges {
		if e.Label == label {
			out = append(out, e)
		}
	}
	return out
}

func parseDataFlowSample(t *testing.T) *ParseResult {
	t.Helper()
	content, err := os.ReadFile("../../testdata/dataflow_sample.go")
	if err != nil {
		t.Fatalf("Failed to read test fixture: %v", err)
	}
	p := NewGoParser()
	result, err := p.ParseFile("testdata/dataflow_sample.go", content)
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}
	return result
}

func TestGoDataFlowParameters(t *testing.T) {
	result := parseDataFlowSample(t)

	if len(result.Parameters) == 0 {
		t.Fatal("expected Parameters, got 0")
	}

	names := make(map[string]bool)
	for _, p := range result.Parameters {
		names[p.Name] = true
		if p.Kind != graph.NodeParameter {
			t.Errorf("parameter %q has Kind=%s, want NodeParameter", p.Name, p.Kind)
		}
	}

	// HandleRequest has w and r
	if !names["w"] {
		t.Error("expected parameter 'w' from HandleRequest")
	}
	if !names["r"] {
		t.Error("expected parameter 'r' from HandleRequest")
	}
}

func TestGoDataFlowVariables(t *testing.T) {
	result := parseDataFlowSample(t)

	if len(result.Variables) < 6 {
		t.Errorf("expected at least 6 variables, got %d", len(result.Variables))
		for _, v := range result.Variables {
			t.Logf("  variable: %s", v.Name)
		}
	}

	names := make(map[string]bool)
	for _, v := range result.Variables {
		names[v.Name] = true
		if v.Kind != graph.NodeVariable {
			t.Errorf("variable %q has Kind=%s, want NodeVariable", v.Name, v.Kind)
		}
	}

	for _, expected := range []string{"body", "review", "name", "query", "x", "y"} {
		if !names[expected] {
			t.Errorf("expected variable %q not found", expected)
		}
	}
}

func TestGoDataFlowDeclaresEdges(t *testing.T) {
	result := parseDataFlowSample(t)

	declares := filterEdgesByLabel(result.Edges, "declares")
	if len(declares) == 0 {
		t.Fatal("expected declares edges, got 0")
	}

	for _, e := range declares {
		if e.Kind != graph.EdgeDataFlow {
			t.Errorf("declares edge has Kind=%s, want EdgeDataFlow", e.Kind)
		}
		if e.Confidence != graph.ConfidenceCertain {
			t.Errorf("declares edge has Confidence=%s, want CERTAIN", e.Confidence)
		}
	}
}

func TestGoDataFlowAssignsEdges(t *testing.T) {
	result := parseDataFlowSample(t)

	assigns := filterEdgesByLabel(result.Edges, "assigns")
	if len(assigns) < 2 {
		t.Errorf("expected at least 2 assigns edges, got %d", len(assigns))
		for _, e := range assigns {
			t.Logf("  assigns: %s -> %s", e.From, e.To)
		}
	}
}

func TestGoDataFlowPassesToEdges(t *testing.T) {
	result := parseDataFlowSample(t)

	passesTo := filterEdgesByLabel(result.Edges, "passes_to")
	if len(passesTo) < 1 {
		t.Errorf("expected at least 1 passes_to edge, got %d", len(passesTo))
		for _, e := range result.Edges {
			t.Logf("  edge: %s %s -> %s", e.Label, e.From, e.To)
		}
	}
}

func TestGoDataFlowMutatesEdges(t *testing.T) {
	result := parseDataFlowSample(t)

	mutates := filterEdgesByLabel(result.Edges, "mutates")
	if len(mutates) < 1 {
		t.Errorf("expected at least 1 mutates edge, got %d", len(mutates))
	}

	for _, e := range mutates {
		if e.Confidence != graph.ConfidenceInferred {
			t.Errorf("mutates edge has Confidence=%s, want INFERRED", e.Confidence)
		}
	}
}

func TestGoDataFlowFieldAccessEdges(t *testing.T) {
	result := parseDataFlowSample(t)

	fieldAccess := filterEdgesByLabel(result.Edges, "field_access")
	if len(fieldAccess) < 1 {
		t.Errorf("expected at least 1 field_access edge, got %d", len(fieldAccess))
		for _, e := range result.Edges {
			t.Logf("  edge: %s %s -> %s", e.Label, e.From, e.To)
		}
	}
}

func TestGoDataFlowReturnsEdges(t *testing.T) {
	result := parseDataFlowSample(t)

	returns := filterEdgesByLabel(result.Edges, "returns")
	if len(returns) < 1 {
		t.Errorf("expected at least 1 returns edge, got %d", len(returns))
	}
}

func TestGoDataFlowReadsEdges(t *testing.T) {
	result := parseDataFlowSample(t)

	reads := filterEdgesByLabel(result.Edges, "reads")
	if len(reads) < 1 {
		t.Errorf("expected at least 1 reads edge, got %d", len(reads))
		for _, e := range result.Edges {
			t.Logf("  edge: %s %s -> %s", e.Label, e.From, e.To)
		}
	}
}

func TestGoDataFlowBlankIdentifierSkipped(t *testing.T) {
	result := parseDataFlowSample(t)

	for _, v := range result.Variables {
		if v.Name == "_" {
			t.Error("blank identifier '_' should not create a variable node")
		}
	}
	for _, p := range result.Parameters {
		if p.Name == "_" {
			t.Error("blank identifier '_' should not create a parameter node")
		}
	}
}

func TestGoDataFlowMaxVariablesLimit(t *testing.T) {
	// Generate a Go file with 600 variable assignments in one function
	var sb strings.Builder
	sb.WriteString("package testdata\n\nfunc bigFunc() {\n")
	for i := 0; i < 600; i++ {
		fmt.Fprintf(&sb, "\tv%d := %d\n", i, i)
	}
	sb.WriteString("}\n")

	p := NewGoParser()
	result, err := p.ParseFile("testdata/big.go", []byte(sb.String()))
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}

	// Should be capped at MaxVariablesPerFunction (500)
	if len(result.Variables) > 500 {
		t.Errorf("got %d variables, expected at most 500 (MaxVariablesPerFunction)", len(result.Variables))
	}

	// Function should be annotated as truncated
	var bigFunc *graph.Node
	for _, fn := range result.Functions {
		if fn.Name == "bigFunc" {
			bigFunc = fn
			break
		}
	}
	if bigFunc == nil {
		t.Fatal("bigFunc not found")
	}
	if !bigFunc.Annotations["dataflow:truncated"] {
		t.Error("expected dataflow:truncated annotation on bigFunc")
	}
}
