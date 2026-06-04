package flow

import (
	"encoding/json"
	"testing"

	"github.com/ugiordan/architecture-analyzer/pkg/renderer"
)

// ---------- helpers ----------

func minimalGraph() renderer.FlowGraph {
	return renderer.FlowGraph{
		Component: "test-component",
		Nodes:     []renderer.FlowNode{},
		Edges:     []renderer.FlowEdge{},
		Paths:     []renderer.FlowPath{},
	}
}

func singleServiceGraph() renderer.FlowGraph {
	return renderer.FlowGraph{
		Component: "kserve",
		Nodes: []renderer.FlowNode{
			{ID: "client", Label: "Client", Type: renderer.FlowNodeIngress, Layer: 0},
			{ID: "svc-api", Label: "api-server", Type: renderer.FlowNodeService, Layer: 3,
				Meta: map[string]string{"type": "ClusterIP", "ports": "8080"}},
			{ID: "dep-ctrl", Label: "kserve-controller", Type: renderer.FlowNodeDeployment, Layer: 4},
		},
		Edges: []renderer.FlowEdge{
			{ID: "e-client-svc-api-route", From: "client", To: "svc-api", Type: "route"},
			{ID: "e-svc-api-dep-ctrl-target", From: "svc-api", To: "dep-ctrl", Type: "target"},
		},
		Paths: []renderer.FlowPath{
			{Name: "Request Flow", Edges: []string{"e-client-svc-api-route", "e-svc-api-dep-ctrl-target"}, Color: "#3498db"},
		},
	}
}

func multiFlowGraph(pathCount int) renderer.FlowGraph {
	g := renderer.FlowGraph{
		Component: "big-app",
		Nodes: []renderer.FlowNode{
			{ID: "a", Label: "A", Type: renderer.FlowNodeIngress, Layer: 0},
			{ID: "b", Label: "B", Type: renderer.FlowNodeService, Layer: 3},
		},
		Edges: []renderer.FlowEdge{
			{ID: "e-a-b", From: "a", To: "b", Type: "route"},
		},
	}
	for i := 0; i < pathCount; i++ {
		name := string(rune('A'+i)) + " Flow"
		g.Paths = append(g.Paths, renderer.FlowPath{
			Name:  name,
			Edges: []string{"e-a-b"},
			Color: "#000",
		})
	}
	return g
}

// ---------- empty graph ----------

func TestConvertDiagram_EmptyGraph(t *testing.T) {
	d := ConvertDiagram(minimalGraph(), nil)

	if d.Meta.Title != "test-component Architecture" {
		t.Errorf("title = %q, want %q", d.Meta.Title, "test-component Architecture")
	}
	if len(d.Nodes) != 0 {
		t.Errorf("expected 0 nodes, got %d", len(d.Nodes))
	}
	if len(d.Flows) != 0 {
		t.Errorf("expected 0 flows, got %d", len(d.Flows))
	}
	if d.Canvas.Width != 1400 || d.Canvas.Height != 900 {
		t.Errorf("canvas = %dx%d, want 1400x900", d.Canvas.Width, d.Canvas.Height)
	}
}

// ---------- node type mapping ----------

func TestConvertDiagram_NodeTypeMapping(t *testing.T) {
	cases := []struct {
		flowType renderer.FlowNodeType
		wantType string
		wantCol  string
	}{
		{renderer.FlowNodeIngress, "icon", "#58a6ff"},
		{renderer.FlowNodeWebhook, "hexagon", "#d29922"},
		{renderer.FlowNodeService, "box", "#009596"},
		{renderer.FlowNodeDeployment, "box", "#3e8635"},
		{renderer.FlowNodeExternal, "barrel", "#8957e5"},
		{renderer.FlowNodeCRD, "box", "#0066cc"},
	}
	for _, tc := range cases {
		t.Run(string(tc.flowType), func(t *testing.T) {
			g := renderer.FlowGraph{
				Component: "test",
				Nodes: []renderer.FlowNode{
					{ID: "n1", Label: "Node", Type: tc.flowType, Layer: 0},
				},
				Edges: []renderer.FlowEdge{},
				Paths: []renderer.FlowPath{},
			}
			d := ConvertDiagram(g, nil)
			n, ok := d.Nodes["n1"]
			if !ok {
				t.Fatal("node n1 not found")
			}
			if n.Type != tc.wantType {
				t.Errorf("type = %q, want %q", n.Type, tc.wantType)
			}
			if n.Color != tc.wantCol {
				t.Errorf("color = %q, want %q", n.Color, tc.wantCol)
			}
		})
	}
}

// ---------- flows from paths ----------

func TestConvertDiagram_FlowFromPaths(t *testing.T) {
	d := ConvertDiagram(singleServiceGraph(), nil)

	flow, ok := d.Flows["request-flow"]
	if !ok {
		t.Fatalf("flow 'request-flow' not found, have: %v", flowKeys(d.Flows))
	}
	if flow.Label != "Request Flow" {
		t.Errorf("label = %q, want %q", flow.Label, "Request Flow")
	}
	if len(flow.Steps) != 2 {
		t.Fatalf("steps = %d, want 2", len(flow.Steps))
	}

	s0 := flow.Steps[0]
	if s0.Mode != "arrow" {
		t.Errorf("step[0].mode = %q, want 'arrow'", s0.Mode)
	}
	if s0.From != "client" || s0.To != "svc-api" {
		t.Errorf("step[0] from/to = %q/%q, want client/svc-api", s0.From, s0.To)
	}
	if s0.Num != 1 {
		t.Errorf("step[0].num = %d, want 1", s0.Num)
	}

	s1 := flow.Steps[1]
	if s1.From != "svc-api" || s1.To != "dep-ctrl" {
		t.Errorf("step[1] from/to = %q/%q, want svc-api/dep-ctrl", s1.From, s1.To)
	}
	if s1.Num != 2 {
		t.Errorf("step[1].num = %d, want 2", s1.Num)
	}
}

// ---------- tooltips from meta ----------

func TestConvertDiagram_TooltipsFromMeta(t *testing.T) {
	d := ConvertDiagram(singleServiceGraph(), nil)

	tip, ok := d.Tooltips["svc-api"]
	if !ok {
		t.Fatal("tooltip for svc-api not found")
	}
	if tip.Title != "api-server" {
		t.Errorf("title = %q, want %q", tip.Title, "api-server")
	}
	if tip.Details["type"] != "ClusterIP" {
		t.Errorf("details[type] = %q, want ClusterIP", tip.Details["type"])
	}
	if tip.Details["ports"] != "8080" {
		t.Errorf("details[ports] = %q, want 8080", tip.Details["ports"])
	}
}

func TestConvertDiagram_NoTooltipForNoMeta(t *testing.T) {
	d := ConvertDiagram(singleServiceGraph(), nil)
	if _, ok := d.Tooltips["dep-ctrl"]; ok {
		t.Error("node without meta should not have a tooltip")
	}
}

// ---------- legend ----------

func TestConvertDiagram_LegendOnlyPresentTypes(t *testing.T) {
	g := renderer.FlowGraph{
		Component: "test",
		Nodes: []renderer.FlowNode{
			{ID: "a", Label: "A", Type: renderer.FlowNodeService, Layer: 3},
			{ID: "b", Label: "B", Type: renderer.FlowNodeExternal, Layer: 5},
		},
		Edges: []renderer.FlowEdge{},
		Paths: []renderer.FlowPath{},
	}
	d := ConvertDiagram(g, nil)

	if len(d.Legend) != 2 {
		t.Fatalf("legend entries = %d, want 2, got %v", len(d.Legend), d.Legend)
	}
	// Service should come before External (layer order).
	if d.Legend[0].Label != "Service" {
		t.Errorf("legend[0].label = %q, want Service", d.Legend[0].Label)
	}
	if d.Legend[1].Label != "External" {
		t.Errorf("legend[1].label = %q, want External", d.Legend[1].Label)
	}
}

func TestConvertDiagram_LegendEmpty(t *testing.T) {
	d := ConvertDiagram(minimalGraph(), nil)
	if len(d.Legend) != 0 {
		t.Errorf("empty graph legend = %v, want empty", d.Legend)
	}
}

// ---------- mode selection ----------

func TestConvertDiagram_ModeLive(t *testing.T) {
	d := ConvertDiagram(multiFlowGraph(5), nil)
	if d.Mode != "live" {
		t.Errorf("mode = %q, want live (5 flows)", d.Mode)
	}
}

func TestConvertDiagram_ModePlay(t *testing.T) {
	d := ConvertDiagram(multiFlowGraph(6), nil)
	if d.Mode != "play" {
		t.Errorf("mode = %q, want play (6 flows)", d.Mode)
	}
}

func TestConvertDiagram_ModeLiveZero(t *testing.T) {
	d := ConvertDiagram(minimalGraph(), nil)
	if d.Mode != "live" {
		t.Errorf("mode = %q, want live (0 flows)", d.Mode)
	}
}

// ---------- flow order ----------

func TestConvertDiagram_FlowOrderMatchesPathOrder(t *testing.T) {
	g := multiFlowGraph(3)
	d := ConvertDiagram(g, nil)

	if len(d.FlowOrder) != 3 {
		t.Fatalf("flowOrder length = %d, want 3", len(d.FlowOrder))
	}
	// Paths: "A Flow", "B Flow", "C Flow" -> slugs: "a-flow", "b-flow", "c-flow"
	want := []string{"a-flow", "b-flow", "c-flow"}
	for i, w := range want {
		if d.FlowOrder[i] != w {
			t.Errorf("flowOrder[%d] = %q, want %q", i, d.FlowOrder[i], w)
		}
	}
}

func TestConvertDiagram_DefaultFlowFirstPath(t *testing.T) {
	d := ConvertDiagram(singleServiceGraph(), nil)
	if d.DefaultFlow != "request-flow" {
		t.Errorf("defaultFlow = %q, want request-flow", d.DefaultFlow)
	}
}

func TestConvertDiagram_DefaultFlowEmptyWhenNoPaths(t *testing.T) {
	d := ConvertDiagram(minimalGraph(), nil)
	if d.DefaultFlow != "" {
		t.Errorf("defaultFlow = %q, want empty", d.DefaultFlow)
	}
}

// ---------- repo from data ----------

func TestConvertDiagram_RepoFromData(t *testing.T) {
	data := map[string]interface{}{
		"repo": "org/my-repo",
	}
	d := ConvertDiagram(minimalGraph(), data)
	if d.Meta.Repo != "org/my-repo" {
		t.Errorf("repo = %q, want org/my-repo", d.Meta.Repo)
	}
}

func TestConvertDiagram_RepoEmptyWhenNilData(t *testing.T) {
	d := ConvertDiagram(minimalGraph(), nil)
	if d.Meta.Repo != "" {
		t.Errorf("repo = %q, want empty", d.Meta.Repo)
	}
}

// ---------- layer field excluded from JSON ----------

func TestDiagramNode_LayerExcludedFromJSON(t *testing.T) {
	g := renderer.FlowGraph{
		Component: "test",
		Nodes: []renderer.FlowNode{
			{ID: "n1", Label: "Node", Type: renderer.FlowNodeService, Layer: 3},
		},
		Edges: []renderer.FlowEdge{},
		Paths: []renderer.FlowPath{},
	}
	d := ConvertDiagram(g, nil)

	b, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}

	// The word "layer" should not appear in the JSON output.
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	nodesRaw := raw["nodes"].(map[string]interface{})
	n1Raw := nodesRaw["n1"].(map[string]interface{})
	if _, exists := n1Raw["layer"]; exists {
		t.Error("layer field should be excluded from JSON (json:\"-\")")
	}
}

func TestDiagramNode_LayerAccessor(t *testing.T) {
	g := renderer.FlowGraph{
		Component: "test",
		Nodes: []renderer.FlowNode{
			{ID: "n1", Label: "Node", Type: renderer.FlowNodeService, Layer: 3},
		},
		Edges: []renderer.FlowEdge{},
		Paths: []renderer.FlowPath{},
	}
	d := ConvertDiagram(g, nil)
	n := d.Nodes["n1"]
	if n.Layer() != 3 {
		t.Errorf("Layer() = %d, want 3", n.Layer())
	}
}

// ---------- edge label preserved ----------

func TestConvertDiagram_EdgeLabelPreserved(t *testing.T) {
	g := renderer.FlowGraph{
		Component: "test",
		Nodes: []renderer.FlowNode{
			{ID: "a", Label: "A", Type: renderer.FlowNodeIngress, Layer: 0},
			{ID: "b", Label: "B", Type: renderer.FlowNodeService, Layer: 3},
		},
		Edges: []renderer.FlowEdge{
			{ID: "e1", From: "a", To: "b", Type: "route", Label: "HTTP GET"},
		},
		Paths: []renderer.FlowPath{
			{Name: "Test", Edges: []string{"e1"}, Color: "#000"},
		},
	}
	d := ConvertDiagram(g, nil)
	flow := d.Flows["test"]
	if len(flow.Steps) != 1 {
		t.Fatalf("steps = %d, want 1", len(flow.Steps))
	}
	if flow.Steps[0].Label != "HTTP GET" {
		t.Errorf("step label = %q, want %q", flow.Steps[0].Label, "HTTP GET")
	}
}

// ---------- missing edge in path is skipped ----------

func TestConvertDiagram_MissingEdgeSkipped(t *testing.T) {
	g := renderer.FlowGraph{
		Component: "test",
		Nodes: []renderer.FlowNode{
			{ID: "a", Label: "A", Type: renderer.FlowNodeIngress, Layer: 0},
		},
		Edges: []renderer.FlowEdge{},
		Paths: []renderer.FlowPath{
			{Name: "Broken", Edges: []string{"nonexistent-edge"}, Color: "#000"},
		},
	}
	d := ConvertDiagram(g, nil)
	flow := d.Flows["broken"]
	if len(flow.Steps) != 0 {
		t.Errorf("steps = %d, want 0 (missing edge should be skipped)", len(flow.Steps))
	}
}

// ---------- slugify ----------

func TestSlugify(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"Request Flow", "request-flow"},
		{"my-controller external calls", "my-controller-external-calls"},
		{"  Spaced  Out  ", "spaced-out"},
		{"UPPER", "upper"},
		{"a/b:c", "a-b-c"},
		{"", ""},
	}
	for _, tc := range cases {
		got := slugify(tc.input)
		if got != tc.want {
			t.Errorf("slugify(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

// ---------- JSON round-trip ----------

func TestConvertDiagram_JSONRoundTrip(t *testing.T) {
	d := ConvertDiagram(singleServiceGraph(), map[string]interface{}{"repo": "org/repo"})
	b, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var d2 Diagram
	if err := json.Unmarshal(b, &d2); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if d2.Meta.Title != d.Meta.Title {
		t.Errorf("round-trip title = %q, want %q", d2.Meta.Title, d.Meta.Title)
	}
	if d2.Meta.Repo != "org/repo" {
		t.Errorf("round-trip repo = %q, want org/repo", d2.Meta.Repo)
	}
	if len(d2.Flows) != len(d.Flows) {
		t.Errorf("round-trip flows = %d, want %d", len(d2.Flows), len(d.Flows))
	}
	if d2.Mode != d.Mode {
		t.Errorf("round-trip mode = %q, want %q", d2.Mode, d.Mode)
	}
}

// ---------- helper ----------

func flowKeys(m map[string]DiagramFlow) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
