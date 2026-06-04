package flow

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/ugiordan/architecture-analyzer/pkg/renderer"
)

// reconcileNodeID produces a safe HTML element ID from a label.
// Letters, digits, hyphens, and underscores are kept; everything else
// becomes a hyphen.
func reconcileNodeID(s string) string {
	var b strings.Builder
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '-' || ch == '_' {
			b.WriteRune(ch)
		} else {
			b.WriteByte('-')
		}
	}
	result := b.String()
	if result == "" {
		return "node"
	}
	return result
}

// kindFromGVK extracts the Kind (last segment) from a slash-separated GVK
// string like "serving/v1beta1/InferenceService".
func kindFromGVK(gvk string) string {
	if idx := strings.LastIndex(gvk, "/"); idx >= 0 && idx+1 < len(gvk) {
		return gvk[idx+1:]
	}
	return gvk
}

// AddReconcileFlows analyzes controller_watches and crds from component
// architecture data and appends reconciliation nodes, edges, and paths
// to the existing FlowGraph. Each "For" watch that targets a CRD kind
// produces a CRD node, an edge to the controller's deployment, and
// edges for any "Owns" resources on the same controller.
func AddReconcileFlows(g *renderer.FlowGraph, data map[string]interface{}) {
	watches := renderer.GetSlice(data, "controller_watches")
	crds := renderer.GetSlice(data, "crds")
	deployments := renderer.GetSlice(data, "deployments")

	if len(watches) == 0 {
		return
	}

	// Build CRD kind set.
	crdKinds := map[string]bool{}
	for _, crd := range crds {
		kind := renderer.GetStr(crd, "kind", "")
		if kind != "" {
			crdKinds[kind] = true
		}
	}

	// Index existing node IDs so we don't add duplicates.
	existingNodes := map[string]bool{}
	for _, n := range g.Nodes {
		existingNodes[n.ID] = true
	}

	// Index existing edge keys for dedup.
	existingEdges := map[string]bool{}
	for _, e := range g.Edges {
		existingEdges[e.ID] = true
	}

	// Group watches by controller name.
	forByCtrl := map[string][]map[string]interface{}{}
	ownsByCtrl := map[string][]map[string]interface{}{}
	for _, w := range watches {
		wType := renderer.GetStr(w, "type", "")
		ctrl := renderer.GetStr(w, "controller", "")
		if ctrl == "" {
			continue
		}
		switch wType {
		case "For":
			forByCtrl[ctrl] = append(forByCtrl[ctrl], w)
		case "Owns":
			ownsByCtrl[ctrl] = append(ownsByCtrl[ctrl], w)
		}
	}

	// Resolve the deployment target for controllers.
	// If only one deployment exists, all controllers map to it.
	// Otherwise try to find a matching deployment node in the graph.
	resolveDeploymentID := func(controller string) (string, bool) {
		// Single deployment shortcut.
		if len(deployments) == 1 {
			name := renderer.GetStr(deployments[0], "name", "")
			if name != "" {
				id := "dep-" + reconcileNodeID(name)
				return id, existingNodes[id]
			}
		}

		// Try to find a deployment node that already exists in the graph.
		for _, n := range g.Nodes {
			if n.Type == renderer.FlowNodeDeployment {
				return n.ID, true
			}
		}

		// Fallback: create a generic controller node.
		return "ctrl-" + reconcileNodeID(controller), false
	}

	addNode := func(n renderer.FlowNode) {
		if existingNodes[n.ID] {
			return
		}
		existingNodes[n.ID] = true
		g.Nodes = append(g.Nodes, n)
	}

	addEdge := func(e renderer.FlowEdge) {
		if existingEdges[e.ID] {
			return
		}
		existingEdges[e.ID] = true
		g.Edges = append(g.Edges, e)
	}

	// Process each controller's For watches.
	for ctrl, forWatches := range forByCtrl {
		for _, fw := range forWatches {
			gvk := renderer.GetStr(fw, "gvk", "")
			kind := kindFromGVK(gvk)
			if kind == "" {
				continue
			}

			// Skip kinds that don't match any CRD.
			if !crdKinds[kind] {
				continue
			}

			// CRD node.
			crdID := "crd-" + reconcileNodeID(kind)
			addNode(renderer.FlowNode{
				ID:    crdID,
				Label: kind,
				Type:  renderer.FlowNodeCRD,
				Layer: 5,
				Meta:  map[string]string{"gvk": gvk, "controller": ctrl},
			})

			// Resolve deployment target.
			depID, depExists := resolveDeploymentID(ctrl)
			if !depExists {
				// Create a deployment-like node for the controller.
				addNode(renderer.FlowNode{
					ID:    depID,
					Label: ctrl,
					Type:  renderer.FlowNodeDeployment,
					Layer: 4,
				})
			}

			// Edge: CRD -> controller deployment (reconcile).
			reconcileEdgeID := fmt.Sprintf("reconcile-%s-%s", reconcileNodeID(kind), depID)
			addEdge(renderer.FlowEdge{
				ID:    reconcileEdgeID,
				From:  crdID,
				To:    depID,
				Type:  "reconcile",
				Label: "reconciles",
			})

			// Collect path edges starting with the reconcile edge.
			pathEdges := []string{reconcileEdgeID}

			// Process Owns entries for this controller.
			for _, ow := range ownsByCtrl[ctrl] {
				ownGVK := renderer.GetStr(ow, "gvk", "")
				ownKind := kindFromGVK(ownGVK)
				if ownKind == "" {
					continue
				}

				ownedID := "owned-" + reconcileNodeID(ownKind)
				addNode(renderer.FlowNode{
					ID:    ownedID,
					Label: ownKind,
					Type:  renderer.FlowNodeCRD,
					Layer: 6,
				})

				manageEdgeID := fmt.Sprintf("reconcile-%s-%s", depID, ownedID)
				addEdge(renderer.FlowEdge{
					ID:    manageEdgeID,
					From:  depID,
					To:    ownedID,
					Type:  "manages",
					Label: "manages",
				})
				pathEdges = append(pathEdges, manageEdgeID)
			}

			// Build reconciliation path.
			g.Paths = append(g.Paths, renderer.FlowPath{
				Name:  kind + " Reconciliation",
				Edges: pathEdges,
				Color: "#0066cc",
			})
		}
	}
}
