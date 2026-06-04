package flow

import (
	"testing"

	"github.com/ugiordan/architecture-analyzer/pkg/renderer"
)

// ---------- test helpers ----------

func emptyGraph() renderer.FlowGraph {
	return renderer.FlowGraph{
		Component: "test",
		Nodes:     []renderer.FlowNode{},
		Edges:     []renderer.FlowEdge{},
		Paths:     []renderer.FlowPath{},
	}
}

func graphWithDeployment(depName string) renderer.FlowGraph {
	return renderer.FlowGraph{
		Component: "test",
		Nodes: []renderer.FlowNode{
			{ID: "dep-" + reconcileNodeID(depName), Label: depName, Type: renderer.FlowNodeDeployment, Layer: 4},
		},
		Edges: []renderer.FlowEdge{},
		Paths: []renderer.FlowPath{},
	}
}

func findNode(g renderer.FlowGraph, id string) *renderer.FlowNode {
	for i := range g.Nodes {
		if g.Nodes[i].ID == id {
			return &g.Nodes[i]
		}
	}
	return nil
}

func findEdge(g renderer.FlowGraph, id string) *renderer.FlowEdge {
	for i := range g.Edges {
		if g.Edges[i].ID == id {
			return &g.Edges[i]
		}
	}
	return nil
}

func findPath(g renderer.FlowGraph, name string) *renderer.FlowPath {
	for i := range g.Paths {
		if g.Paths[i].Name == name {
			return &g.Paths[i]
		}
	}
	return nil
}

// ---------- matching CRD ----------

func TestAddReconcileFlows_MatchingCRD(t *testing.T) {
	g := graphWithDeployment("my-controller")
	data := map[string]interface{}{
		"controller_watches": []interface{}{
			map[string]interface{}{
				"type":       "For",
				"gvk":        "serving/v1beta1/InferenceService",
				"controller": "ISReconciler",
				"source":     "pkg/ctrl/is.go:42",
			},
		},
		"crds": []interface{}{
			map[string]interface{}{
				"kind":    "InferenceService",
				"group":   "serving.kserve.io",
				"version": "v1beta1",
			},
		},
		"deployments": []interface{}{
			map[string]interface{}{"name": "my-controller"},
		},
	}

	AddReconcileFlows(&g, data)

	// CRD node should be created.
	crdNode := findNode(g, "crd-InferenceService")
	if crdNode == nil {
		t.Fatal("expected crd-InferenceService node")
	}
	if crdNode.Type != renderer.FlowNodeCRD {
		t.Errorf("crd node type = %q, want crd", crdNode.Type)
	}
	if crdNode.Layer != 5 {
		t.Errorf("crd node layer = %d, want 5", crdNode.Layer)
	}
	if crdNode.Label != "InferenceService" {
		t.Errorf("crd node label = %q, want InferenceService", crdNode.Label)
	}

	// Edge: CRD -> deployment.
	edge := findEdge(g, "reconcile-InferenceService-dep-my-controller")
	if edge == nil {
		t.Fatal("expected reconcile edge from CRD to deployment")
	}
	if edge.From != "crd-InferenceService" {
		t.Errorf("edge.From = %q, want crd-InferenceService", edge.From)
	}
	if edge.To != "dep-my-controller" {
		t.Errorf("edge.To = %q, want dep-my-controller", edge.To)
	}
	if edge.Label != "reconciles" {
		t.Errorf("edge.Label = %q, want reconciles", edge.Label)
	}

	// Path should exist.
	path := findPath(g, "InferenceService Reconciliation")
	if path == nil {
		t.Fatal("expected InferenceService Reconciliation path")
	}
	if path.Color != "#0066cc" {
		t.Errorf("path color = %q, want #0066cc", path.Color)
	}
	if len(path.Edges) < 1 {
		t.Errorf("path edges = %d, want >= 1", len(path.Edges))
	}
}

// ---------- skips core types ----------

func TestAddReconcileFlows_SkipsCoreTypes(t *testing.T) {
	g := emptyGraph()
	data := map[string]interface{}{
		"controller_watches": []interface{}{
			map[string]interface{}{
				"type":       "For",
				"gvk":        "/v1/ConfigMap",
				"controller": "ConfigMapReconciler",
				"source":     "pkg/ctrl/cm.go:10",
			},
			map[string]interface{}{
				"type":       "For",
				"gvk":        "/v1/Pod",
				"controller": "PodReconciler",
				"source":     "pkg/ctrl/pod.go:10",
			},
			map[string]interface{}{
				"type":       "For",
				"gvk":        "/v1/Secret",
				"controller": "SecretReconciler",
				"source":     "pkg/ctrl/secret.go:10",
			},
		},
		"crds":        []interface{}{},
		"deployments": []interface{}{},
	}

	AddReconcileFlows(&g, data)

	if len(g.Nodes) != 0 {
		t.Errorf("expected 0 nodes (core types skipped), got %d", len(g.Nodes))
	}
	if len(g.Edges) != 0 {
		t.Errorf("expected 0 edges (core types skipped), got %d", len(g.Edges))
	}
	if len(g.Paths) != 0 {
		t.Errorf("expected 0 paths (core types skipped), got %d", len(g.Paths))
	}
}

// ---------- single deployment ----------

func TestAddReconcileFlows_SingleDeployment(t *testing.T) {
	g := graphWithDeployment("operator")
	data := map[string]interface{}{
		"controller_watches": []interface{}{
			map[string]interface{}{
				"type":       "For",
				"gvk":        "api/v1/Widget",
				"controller": "WidgetReconciler",
				"source":     "ctrl/widget.go:1",
			},
			map[string]interface{}{
				"type":       "For",
				"gvk":        "api/v1/Gadget",
				"controller": "GadgetReconciler",
				"source":     "ctrl/gadget.go:1",
			},
		},
		"crds": []interface{}{
			map[string]interface{}{"kind": "Widget", "group": "example.io", "version": "v1"},
			map[string]interface{}{"kind": "Gadget", "group": "example.io", "version": "v1"},
		},
		"deployments": []interface{}{
			map[string]interface{}{"name": "operator"},
		},
	}

	AddReconcileFlows(&g, data)

	// Both CRDs should point to the same deployment.
	widgetEdge := findEdge(g, "reconcile-Widget-dep-operator")
	gadgetEdge := findEdge(g, "reconcile-Gadget-dep-operator")
	if widgetEdge == nil {
		t.Error("missing reconcile edge for Widget")
	}
	if gadgetEdge == nil {
		t.Error("missing reconcile edge for Gadget")
	}
	if widgetEdge != nil && widgetEdge.To != "dep-operator" {
		t.Errorf("Widget edge target = %q, want dep-operator", widgetEdge.To)
	}
	if gadgetEdge != nil && gadgetEdge.To != "dep-operator" {
		t.Errorf("Gadget edge target = %q, want dep-operator", gadgetEdge.To)
	}

	// The deployment node should NOT be duplicated.
	depCount := 0
	for _, n := range g.Nodes {
		if n.Type == renderer.FlowNodeDeployment {
			depCount++
		}
	}
	if depCount != 1 {
		t.Errorf("deployment node count = %d, want 1 (no duplicates)", depCount)
	}
}

// ---------- owns entries ----------

func TestAddReconcileFlows_OwnsEntries(t *testing.T) {
	g := graphWithDeployment("dspa-operator")
	data := map[string]interface{}{
		"controller_watches": []interface{}{
			map[string]interface{}{
				"type":       "For",
				"gvk":        "api/v1/DSPA",
				"controller": "DSPAReconciler",
				"source":     "ctrl/dspa.go:1",
			},
			map[string]interface{}{
				"type":       "Owns",
				"gvk":        "/v1/ConfigMap",
				"controller": "DSPAReconciler",
				"source":     "ctrl/dspa.go:2",
			},
			map[string]interface{}{
				"type":       "Owns",
				"gvk":        "apps/v1/Deployment",
				"controller": "DSPAReconciler",
				"source":     "ctrl/dspa.go:3",
			},
		},
		"crds": []interface{}{
			map[string]interface{}{"kind": "DSPA", "group": "api.example.io", "version": "v1"},
		},
		"deployments": []interface{}{
			map[string]interface{}{"name": "dspa-operator"},
		},
	}

	AddReconcileFlows(&g, data)

	// Owned resource nodes.
	cmNode := findNode(g, "owned-ConfigMap")
	if cmNode == nil {
		t.Fatal("expected owned-ConfigMap node")
	}
	if cmNode.Layer != 6 {
		t.Errorf("owned node layer = %d, want 6", cmNode.Layer)
	}

	depOwnedNode := findNode(g, "owned-Deployment")
	if depOwnedNode == nil {
		t.Fatal("expected owned-Deployment node")
	}

	// Manages edges.
	cmEdge := findEdge(g, "reconcile-dep-dspa-operator-owned-ConfigMap")
	if cmEdge == nil {
		t.Fatal("expected manages edge to owned-ConfigMap")
	}
	if cmEdge.Label != "manages" {
		t.Errorf("manages edge label = %q, want manages", cmEdge.Label)
	}
	if cmEdge.Type != "manages" {
		t.Errorf("manages edge type = %q, want manages", cmEdge.Type)
	}
	if cmEdge.From != "dep-dspa-operator" {
		t.Errorf("manages edge from = %q, want dep-dspa-operator", cmEdge.From)
	}

	// Path should include reconcile + manages edges.
	path := findPath(g, "DSPA Reconciliation")
	if path == nil {
		t.Fatal("expected DSPA Reconciliation path")
	}
	// At least 3 edges: reconcile + 2 manages.
	if len(path.Edges) != 3 {
		t.Errorf("path edges = %d, want 3", len(path.Edges))
	}
}

// ---------- no watches ----------

func TestAddReconcileFlows_NoWatches(t *testing.T) {
	g := emptyGraph()
	data := map[string]interface{}{}

	AddReconcileFlows(&g, data)

	if len(g.Nodes) != 0 {
		t.Errorf("expected 0 nodes, got %d", len(g.Nodes))
	}
	if len(g.Edges) != 0 {
		t.Errorf("expected 0 edges, got %d", len(g.Edges))
	}
	if len(g.Paths) != 0 {
		t.Errorf("expected 0 paths, got %d", len(g.Paths))
	}
}

// ---------- multiple CRDs ----------

func TestAddReconcileFlows_MultipleCRDs(t *testing.T) {
	g := graphWithDeployment("operator")
	data := map[string]interface{}{
		"controller_watches": []interface{}{
			map[string]interface{}{
				"type":       "For",
				"gvk":        "api/v1/Alpha",
				"controller": "AlphaReconciler",
				"source":     "ctrl/alpha.go:1",
			},
			map[string]interface{}{
				"type":       "For",
				"gvk":        "api/v1/Beta",
				"controller": "BetaReconciler",
				"source":     "ctrl/beta.go:1",
			},
		},
		"crds": []interface{}{
			map[string]interface{}{"kind": "Alpha", "group": "example.io", "version": "v1"},
			map[string]interface{}{"kind": "Beta", "group": "example.io", "version": "v1"},
		},
		"deployments": []interface{}{
			map[string]interface{}{"name": "operator"},
		},
	}

	AddReconcileFlows(&g, data)

	alphaPath := findPath(g, "Alpha Reconciliation")
	betaPath := findPath(g, "Beta Reconciliation")

	if alphaPath == nil {
		t.Error("expected Alpha Reconciliation path")
	}
	if betaPath == nil {
		t.Error("expected Beta Reconciliation path")
	}

	// Should have 2 CRD nodes + 1 existing deployment = 3 total.
	if len(g.Nodes) != 3 {
		t.Errorf("node count = %d, want 3 (2 CRDs + 1 deployment)", len(g.Nodes))
	}
}

// ---------- no matching CRD ----------

func TestAddReconcileFlows_NoMatchingCRD(t *testing.T) {
	g := graphWithDeployment("operator")
	data := map[string]interface{}{
		"controller_watches": []interface{}{
			map[string]interface{}{
				"type":       "For",
				"gvk":        "api/v1/UnknownThing",
				"controller": "UnknownReconciler",
				"source":     "ctrl/unknown.go:1",
			},
		},
		"crds": []interface{}{
			map[string]interface{}{"kind": "SomethingElse", "group": "example.io", "version": "v1"},
		},
		"deployments": []interface{}{
			map[string]interface{}{"name": "operator"},
		},
	}

	initialNodeCount := len(g.Nodes)
	AddReconcileFlows(&g, data)

	if len(g.Nodes) != initialNodeCount {
		t.Errorf("node count changed from %d to %d, want no change (no matching CRD)", initialNodeCount, len(g.Nodes))
	}
	if len(g.Edges) != 0 {
		t.Errorf("expected 0 edges, got %d", len(g.Edges))
	}
	if len(g.Paths) != 0 {
		t.Errorf("expected 0 paths, got %d", len(g.Paths))
	}
}

// ---------- controller fallback when no deployment exists ----------

func TestAddReconcileFlows_FallbackControllerNode(t *testing.T) {
	g := emptyGraph() // no deployment nodes
	data := map[string]interface{}{
		"controller_watches": []interface{}{
			map[string]interface{}{
				"type":       "For",
				"gvk":        "api/v1/Widget",
				"controller": "WidgetReconciler",
				"source":     "ctrl/widget.go:1",
			},
		},
		"crds": []interface{}{
			map[string]interface{}{"kind": "Widget", "group": "example.io", "version": "v1"},
		},
		"deployments": []interface{}{},
	}

	AddReconcileFlows(&g, data)

	// A fallback controller node should be created.
	ctrlNode := findNode(g, "ctrl-WidgetReconciler")
	if ctrlNode == nil {
		t.Fatal("expected fallback controller node ctrl-WidgetReconciler")
	}
	if ctrlNode.Type != renderer.FlowNodeDeployment {
		t.Errorf("fallback node type = %q, want deployment", ctrlNode.Type)
	}

	// Edge should point to the fallback node.
	edge := findEdge(g, "reconcile-Widget-ctrl-WidgetReconciler")
	if edge == nil {
		t.Fatal("expected reconcile edge to fallback controller")
	}
	if edge.To != "ctrl-WidgetReconciler" {
		t.Errorf("edge.To = %q, want ctrl-WidgetReconciler", edge.To)
	}
}

// ---------- kindFromGVK ----------

func TestKindFromGVK(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"serving/v1beta1/InferenceService", "InferenceService"},
		{"/v1/ConfigMap", "ConfigMap"},
		{"apps/v1/Deployment", "Deployment"},
		{"InferenceService", "InferenceService"},
		{"", ""},
		{"a/b/c/d", "d"},
	}
	for _, tc := range cases {
		got := kindFromGVK(tc.input)
		if got != tc.want {
			t.Errorf("kindFromGVK(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

// ---------- reconcileNodeID ----------

func TestReconcileNodeID(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"InferenceService", "InferenceService"},
		{"my-controller", "my-controller"},
		{"foo.bar/baz:123", "foo-bar-baz-123"},
		{"", "node"},
		{"a b c", "a-b-c"},
	}
	for _, tc := range cases {
		got := reconcileNodeID(tc.input)
		if got != tc.want {
			t.Errorf("reconcileNodeID(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

// ---------- duplicate nodes are not added ----------

func TestAddReconcileFlows_NoDuplicateNodes(t *testing.T) {
	g := graphWithDeployment("operator")
	data := map[string]interface{}{
		"controller_watches": []interface{}{
			map[string]interface{}{
				"type":       "For",
				"gvk":        "api/v1/Widget",
				"controller": "WidgetReconciler",
				"source":     "ctrl/widget.go:1",
			},
		},
		"crds": []interface{}{
			map[string]interface{}{"kind": "Widget", "group": "example.io", "version": "v1"},
		},
		"deployments": []interface{}{
			map[string]interface{}{"name": "operator"},
		},
	}

	// Call twice to verify idempotency.
	AddReconcileFlows(&g, data)
	nodeCountAfterFirst := len(g.Nodes)
	edgeCountAfterFirst := len(g.Edges)

	AddReconcileFlows(&g, data)

	if len(g.Nodes) != nodeCountAfterFirst {
		t.Errorf("second call added nodes: %d -> %d", nodeCountAfterFirst, len(g.Nodes))
	}
	if len(g.Edges) != edgeCountAfterFirst {
		t.Errorf("second call added edges: %d -> %d", edgeCountAfterFirst, len(g.Edges))
	}
}
