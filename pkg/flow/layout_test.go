package flow

import (
	"testing"

	"github.com/ugiordan/architecture-analyzer/pkg/renderer"
)

// ---------- helper ----------

func layoutGraph(nodes []renderer.FlowNode) Diagram {
	g := renderer.FlowGraph{
		Component: "layout-test",
		Nodes:     nodes,
		Edges:     []renderer.FlowEdge{},
		Paths:     []renderer.FlowPath{},
	}
	d := ConvertDiagram(g, nil)
	return d
}

// ---------- basic position assignment ----------

func TestApplyLayout_AssignsPositions(t *testing.T) {
	d := layoutGraph([]renderer.FlowNode{
		{ID: "a", Label: "A", Type: renderer.FlowNodeIngress, Layer: 0},
		{ID: "b", Label: "B", Type: renderer.FlowNodeService, Layer: 3},
		{ID: "c", Label: "C", Type: renderer.FlowNodeDeployment, Layer: 3},
	})

	na := d.Nodes["a"]
	nb := d.Nodes["b"]
	nc := d.Nodes["c"]

	// Node in layer 0 should have a different X than nodes in layer 3.
	if na.X == nb.X {
		t.Errorf("nodes in different layers should have different X: a.X=%d, b.X=%d", na.X, nb.X)
	}

	// Nodes in the same layer should have the same X.
	if nb.X != nc.X {
		t.Errorf("nodes in same layer should have same X: b.X=%d, c.X=%d", nb.X, nc.X)
	}

	// All Y positions should be non-zero (centered on a canvas >= 900).
	if na.Y == 0 {
		t.Error("node a Y should be non-zero after layout")
	}
	if nb.Y == 0 {
		t.Error("node b Y should be non-zero after layout")
	}
	if nc.Y == 0 {
		t.Error("node c Y should be non-zero after layout")
	}

	// Nodes in the same layer should have different Y positions.
	if nb.Y == nc.Y {
		t.Errorf("nodes b and c in same layer should have different Y: b.Y=%d, c.Y=%d", nb.Y, nc.Y)
	}
}

// ---------- minimum canvas size ----------

func TestApplyLayout_CanvasMinimumSize(t *testing.T) {
	d := layoutGraph([]renderer.FlowNode{
		{ID: "solo", Label: "Solo", Type: renderer.FlowNodeService, Layer: 0},
	})

	if d.Canvas.Width < 1400 {
		t.Errorf("canvas width = %d, want >= 1400", d.Canvas.Width)
	}
	if d.Canvas.Height < 900 {
		t.Errorf("canvas height = %d, want >= 900", d.Canvas.Height)
	}
}

// ---------- dynamic spacing for many nodes ----------

func TestApplyLayout_DynamicSpacing(t *testing.T) {
	nodes := make([]renderer.FlowNode, 20)
	for i := 0; i < 20; i++ {
		id := "n" + string(rune('a'+i))
		nodes[i] = renderer.FlowNode{
			ID: id, Label: id, Type: renderer.FlowNodeService, Layer: 3,
		}
	}
	d := layoutGraph(nodes)

	// 20 nodes in one layer: verticalSpacing = max(80, 900/21) = 80
	// height = max(20*80+100, 900) = max(1700, 900) = 1700
	if d.Canvas.Height <= 900 {
		t.Errorf("canvas height = %d, want > 900 for 20 nodes in one layer", d.Canvas.Height)
	}
}

// ---------- empty diagram ----------

func TestApplyLayout_NoNodes(t *testing.T) {
	d := layoutGraph(nil)

	// Should not panic, and should have minimum canvas.
	if d.Canvas.Width < 1400 {
		t.Errorf("empty canvas width = %d, want >= 1400", d.Canvas.Width)
	}
	if d.Canvas.Height < 900 {
		t.Errorf("empty canvas height = %d, want >= 900", d.Canvas.Height)
	}
}

// ---------- single node centered ----------

func TestApplyLayout_SingleNode(t *testing.T) {
	d := layoutGraph([]renderer.FlowNode{
		{ID: "only", Label: "Only", Type: renderer.FlowNodeIngress, Layer: 0},
	})

	n := d.Nodes["only"]

	// With 1 node on a 900px canvas:
	// verticalSpacing = max(80, 900/2) = 450
	// totalLayerHeight = 1 * 450 = 450
	// verticalOffset = (900 - 450) / 2 + 450 / 2 = 225 + 225 = 450
	// y = 450 + 0 * 450 = 450
	// The node should be roughly centered vertically.
	if n.Y != 450 {
		t.Errorf("single node Y = %d, want 450 (centered)", n.Y)
	}

	// X should be at the first column.
	if n.X != 80 {
		t.Errorf("single node X = %d, want 80", n.X)
	}
}

// ---------- ConvertDiagram integrates layout ----------

func TestConvertDiagram_IntegratesLayout(t *testing.T) {
	g := renderer.FlowGraph{
		Component: "test",
		Nodes: []renderer.FlowNode{
			{ID: "a", Label: "A", Type: renderer.FlowNodeIngress, Layer: 0},
			{ID: "b", Label: "B", Type: renderer.FlowNodeService, Layer: 3},
		},
		Edges: []renderer.FlowEdge{},
		Paths: []renderer.FlowPath{},
	}
	d := ConvertDiagram(g, nil)

	// After ConvertDiagram, layout should already be applied.
	na := d.Nodes["a"]
	nb := d.Nodes["b"]

	if na.X == 0 && nb.X == 0 {
		t.Error("layout not applied: both nodes have X=0")
	}
	if na.Y == 0 && nb.Y == 0 {
		t.Error("layout not applied: both nodes have Y=0")
	}
}
