// Package flow converts renderer.FlowGraph into the flowlens diagram.json
// format for browser-based visualization.
package flow

import (
	"strings"

	"github.com/ugiordan/architecture-analyzer/pkg/renderer"
)

// ---------- flowlens diagram types ----------

// Diagram is the top-level structure written as diagram.json.
type Diagram struct {
	Meta        DiagramMeta            `json:"meta"`
	Canvas      DiagramCanvas          `json:"canvas"`
	Nodes       map[string]DiagramNode `json:"nodes"`
	Flows       map[string]DiagramFlow `json:"flows"`
	Tooltips    map[string]TooltipDef  `json:"tooltips,omitempty"`
	Legend      []LegendEntry          `json:"legend,omitempty"`
	FlowOrder   []string               `json:"flowOrder,omitempty"`
	DefaultFlow string                 `json:"defaultFlow,omitempty"`
	Mode        string                 `json:"mode,omitempty"`
}

// DiagramMeta holds title, subtitle, and optional repository link.
type DiagramMeta struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle,omitempty"`
	Repo     string `json:"repo,omitempty"`
}

// DiagramCanvas sets the SVG viewport dimensions.
type DiagramCanvas struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// DiagramNode is a positioned visual element.
// The unexported layer field is used by auto-layout (Task 3) and excluded
// from JSON output.
type DiagramNode struct {
	X        int    `json:"x"`
	Y        int    `json:"y"`
	W        int    `json:"w,omitempty"`
	H        int    `json:"h,omitempty"`
	Type     string `json:"type"`
	Label    string `json:"label"`
	Sublabel string `json:"sublabel,omitempty"`
	Color    string `json:"color,omitempty"`
	layer    int
}

// Layer returns the layout layer for this node. Used by auto-layout.
func (n DiagramNode) Layer() int { return n.layer }

// DiagramFlow groups ordered arrow steps under a label.
type DiagramFlow struct {
	Label string      `json:"label"`
	Steps []ArrowStep `json:"steps"`
}

// ArrowStep draws a numbered arrow between two nodes.
type ArrowStep struct {
	Mode  string `json:"mode"`
	From  string `json:"from"`
	To    string `json:"to"`
	Num   int    `json:"num"`
	Label string `json:"label,omitempty"`
}

// TooltipDef provides hover content for a node.
type TooltipDef struct {
	Title       string            `json:"title,omitempty"`
	Description string            `json:"description,omitempty"`
	Details     map[string]string `json:"details,omitempty"`
}

// LegendEntry maps a human label to a color swatch.
type LegendEntry struct {
	Label string `json:"label"`
	Color string `json:"color"`
}

// ---------- node type/color mapping ----------

type nodeMapping struct {
	diagramType string
	color       string
	legendLabel string
}

var nodeTypeMap = map[renderer.FlowNodeType]nodeMapping{
	renderer.FlowNodeIngress:    {diagramType: "icon", color: "#58a6ff", legendLabel: "Client / Ingress"},
	renderer.FlowNodeWebhook:    {diagramType: "hexagon", color: "#d29922", legendLabel: "Webhook"},
	renderer.FlowNodeService:    {diagramType: "box", color: "#009596", legendLabel: "Service"},
	renderer.FlowNodeDeployment: {diagramType: "box", color: "#3e8635", legendLabel: "Deployment"},
	renderer.FlowNodeExternal:   {diagramType: "barrel", color: "#8957e5", legendLabel: "External"},
	renderer.FlowNodeCRD:        {diagramType: "box", color: "#0066cc", legendLabel: "CRD"},
}

// ---------- conversion ----------

// ConvertDiagram transforms a renderer.FlowGraph into a flowlens Diagram.
// The optional data map (raw component-architecture.json) is used to extract
// the repository field. ApplyLayout is called automatically to assign
// positions based on each node's layer.
func ConvertDiagram(g renderer.FlowGraph, data map[string]interface{}) Diagram {
	d := Diagram{
		Meta: DiagramMeta{
			Title: g.Component + " Architecture",
		},
		Canvas: DiagramCanvas{Width: 1400, Height: 900},
		Nodes:  make(map[string]DiagramNode),
		Flows:  make(map[string]DiagramFlow),
	}

	// Extract repo from raw data if available.
	if data != nil {
		if repo, ok := data["repo"].(string); ok && repo != "" {
			d.Meta.Repo = repo
		}
	}

	// Track which node types are present for the legend.
	seenTypes := map[renderer.FlowNodeType]bool{}

	// Map FlowNodes to DiagramNodes.
	for _, n := range g.Nodes {
		mapping, ok := nodeTypeMap[n.Type]
		if !ok {
			mapping = nodeMapping{diagramType: "box", color: "#999999", legendLabel: string(n.Type)}
		}
		seenTypes[n.Type] = true

		dn := DiagramNode{
			Type:  mapping.diagramType,
			Label: n.Label,
			Color: mapping.color,
			layer: n.Layer,
		}
		d.Nodes[n.ID] = dn
	}

	// Build edge lookup for path resolution.
	edgeByID := make(map[string]renderer.FlowEdge, len(g.Edges))
	for _, e := range g.Edges {
		edgeByID[e.ID] = e
	}

	// Convert FlowPaths to DiagramFlows.
	for _, p := range g.Paths {
		flowID := slugify(p.Name)
		flow := DiagramFlow{
			Label: p.Name,
			Steps: make([]ArrowStep, 0, len(p.Edges)),
		}
		for i, eid := range p.Edges {
			e, ok := edgeByID[eid]
			if !ok {
				continue
			}
			flow.Steps = append(flow.Steps, ArrowStep{
				Mode:  "arrow",
				From:  e.From,
				To:    e.To,
				Num:   i + 1,
				Label: e.Label,
			})
		}
		d.Flows[flowID] = flow
		d.FlowOrder = append(d.FlowOrder, flowID)
	}

	// Set default flow to the first one if any exist.
	if len(d.FlowOrder) > 0 {
		d.DefaultFlow = d.FlowOrder[0]
	}

	// Generate tooltips from node Meta.
	tooltips := make(map[string]TooltipDef)
	for _, n := range g.Nodes {
		if len(n.Meta) == 0 {
			continue
		}
		tooltips[n.ID] = TooltipDef{
			Title:   n.Label,
			Details: copyMap(n.Meta),
		}
	}
	if len(tooltips) > 0 {
		d.Tooltips = tooltips
	}

	// Build legend from present node types, preserving layer order.
	typeOrder := []renderer.FlowNodeType{
		renderer.FlowNodeIngress,
		renderer.FlowNodeWebhook,
		renderer.FlowNodeService,
		renderer.FlowNodeDeployment,
		renderer.FlowNodeExternal,
		renderer.FlowNodeCRD,
	}
	for _, nt := range typeOrder {
		if !seenTypes[nt] {
			continue
		}
		m := nodeTypeMap[nt]
		d.Legend = append(d.Legend, LegendEntry{
			Label: m.legendLabel,
			Color: m.color,
		})
	}

	// Mode: live for small flow counts, play for larger ones.
	if len(d.Flows) > 5 {
		d.Mode = "play"
	} else {
		d.Mode = "live"
	}

	ApplyLayout(&d)

	return d
}

// ---------- helpers ----------

// slugify lowercases and replaces non-alphanumeric runs with hyphens.
func slugify(s string) string {
	s = strings.ToLower(s)
	var b strings.Builder
	prevDash := false
	for _, ch := range s {
		if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
			b.WriteRune(ch)
			prevDash = false
		} else if !prevDash {
			b.WriteByte('-')
			prevDash = true
		}
	}
	return strings.Trim(b.String(), "-")
}

// copyMap returns a shallow copy of a string map.
func copyMap(m map[string]string) map[string]string {
	out := make(map[string]string, len(m))
	for k, v := range m {
		out[k] = v
	}
	return out
}
