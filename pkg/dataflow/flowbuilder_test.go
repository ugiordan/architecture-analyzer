package dataflow

import (
	"testing"

	"github.com/ugiordan/architecture-analyzer/pkg/graph"
)

func TestFlowBuilderAddVariableAndParameter(t *testing.T) {
	fb := NewFlowBuilder()

	varNode := &graph.Node{ID: "var_abc", Kind: graph.NodeVariable, Name: "x"}
	paramNode := &graph.Node{ID: "param_def", Kind: graph.NodeParameter, Name: "r"}

	fb.AddVariable(varNode)
	fb.AddParameter(paramNode)

	nodes, edges := fb.Result()
	if len(nodes) != 2 {
		t.Fatalf("got %d nodes, want 2", len(nodes))
	}
	if len(edges) != 0 {
		t.Fatalf("got %d edges, want 0", len(edges))
	}
	if nodes[0].ID != "var_abc" {
		t.Errorf("first node ID = %q, want var_abc", nodes[0].ID)
	}
	if nodes[1].ID != "param_def" {
		t.Errorf("second node ID = %q, want param_def", nodes[1].ID)
	}
}

func TestFlowBuilderAddDeclares(t *testing.T) {
	fb := NewFlowBuilder()
	fb.AddDeclares("fn_123", "param_456")

	_, edges := fb.Result()
	if len(edges) != 1 {
		t.Fatalf("got %d edges, want 1", len(edges))
	}
	e := edges[0]
	if e.From != "fn_123" || e.To != "param_456" {
		t.Errorf("edge = %s -> %s, want fn_123 -> param_456", e.From, e.To)
	}
	if e.Kind != graph.EdgeDataFlow {
		t.Errorf("kind = %s, want DATA_FLOW", e.Kind)
	}
	if e.Label != "declares" {
		t.Errorf("label = %q, want declares", e.Label)
	}
	if e.Confidence != graph.ConfidenceCertain {
		t.Errorf("confidence = %s, want CERTAIN", e.Confidence)
	}
}

func TestFlowBuilderAddAssign(t *testing.T) {
	fb := NewFlowBuilder()
	fb.AddAssign("call_read", "var_body")

	_, edges := fb.Result()
	if len(edges) != 1 {
		t.Fatalf("got %d edges, want 1", len(edges))
	}
	e := edges[0]
	if e.From != "call_read" || e.To != "var_body" {
		t.Errorf("edge = %s -> %s, want call_read -> var_body", e.From, e.To)
	}
	if e.Label != "assigns" {
		t.Errorf("label = %q, want assigns", e.Label)
	}
	if e.Confidence != graph.ConfidenceCertain {
		t.Errorf("confidence = %s, want CERTAIN", e.Confidence)
	}
}

func TestFlowBuilderAddRead(t *testing.T) {
	fb := NewFlowBuilder()
	fb.AddRead("var_name", "var_query")

	_, edges := fb.Result()
	if len(edges) != 1 {
		t.Fatalf("got %d edges, want 1", len(edges))
	}
	e := edges[0]
	if e.From != "var_name" || e.To != "var_query" {
		t.Errorf("edge = %s -> %s, want var_name -> var_query", e.From, e.To)
	}
	if e.Label != "reads" {
		t.Errorf("label = %q, want reads", e.Label)
	}
}

func TestFlowBuilderAddPassesTo(t *testing.T) {
	fb := NewFlowBuilder()
	fb.AddPassesTo("var_body", "call_unmarshal")

	_, edges := fb.Result()
	if len(edges) != 1 {
		t.Fatalf("got %d edges, want 1", len(edges))
	}
	e := edges[0]
	if e.From != "var_body" || e.To != "call_unmarshal" {
		t.Errorf("edge = %s -> %s, want var_body -> call_unmarshal", e.From, e.To)
	}
	if e.Label != "passes_to" {
		t.Errorf("label = %q, want passes_to", e.Label)
	}
}

func TestFlowBuilderAddMutates(t *testing.T) {
	fb := NewFlowBuilder()
	fb.AddMutates("call_unmarshal", "var_review")

	_, edges := fb.Result()
	if len(edges) != 1 {
		t.Fatalf("got %d edges, want 1", len(edges))
	}
	e := edges[0]
	if e.From != "call_unmarshal" || e.To != "var_review" {
		t.Errorf("edge = %s -> %s, want call_unmarshal -> var_review", e.From, e.To)
	}
	if e.Label != "mutates" {
		t.Errorf("label = %q, want mutates", e.Label)
	}
	if e.Confidence != graph.ConfidenceInferred {
		t.Errorf("confidence = %s, want INFERRED", e.Confidence)
	}
}

func TestFlowBuilderAddFieldAccess(t *testing.T) {
	fb := NewFlowBuilder()
	fb.AddFieldAccess("var_review", "var_review.Request")

	_, edges := fb.Result()
	if len(edges) != 1 {
		t.Fatalf("got %d edges, want 1", len(edges))
	}
	e := edges[0]
	if e.From != "var_review" || e.To != "var_review.Request" {
		t.Errorf("edge = %s -> %s", e.From, e.To)
	}
	if e.Label != "field_access" {
		t.Errorf("label = %q, want field_access", e.Label)
	}
}

func TestFlowBuilderAddReturn(t *testing.T) {
	fb := NewFlowBuilder()
	fb.AddReturn("var_result", "fn_handle")

	_, edges := fb.Result()
	if len(edges) != 1 {
		t.Fatalf("got %d edges, want 1", len(edges))
	}
	e := edges[0]
	if e.From != "var_result" || e.To != "fn_handle" {
		t.Errorf("edge = %s -> %s, want var_result -> fn_handle", e.From, e.To)
	}
	if e.Label != "returns" {
		t.Errorf("label = %q, want returns", e.Label)
	}
}

func TestFlowBuilderVariableCount(t *testing.T) {
	fb := NewFlowBuilder()
	if fb.VariableCount() != 0 {
		t.Errorf("initial count = %d, want 0", fb.VariableCount())
	}
	fb.AddVariable(&graph.Node{ID: "var_1", Kind: graph.NodeVariable})
	fb.AddVariable(&graph.Node{ID: "var_2", Kind: graph.NodeVariable})
	if fb.VariableCount() != 2 {
		t.Errorf("count = %d, want 2", fb.VariableCount())
	}
	// Parameters don't count toward variable limit
	fb.AddParameter(&graph.Node{ID: "param_1", Kind: graph.NodeParameter})
	if fb.VariableCount() != 2 {
		t.Errorf("count after param = %d, want 2", fb.VariableCount())
	}
}
