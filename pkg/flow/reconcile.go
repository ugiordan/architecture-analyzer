package flow

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ugiordan/architecture-analyzer/pkg/renderer"
)

// kindFromGVK extracts the Kind (last segment) from a slash-separated GVK
// string like "serving/v1beta1/InferenceService".
func kindFromGVK(gvk string) string {
	gvk = strings.TrimRight(gvk, "/")
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
	// For multiple deployments, try to match by controller name.
	// Otherwise, create a per-controller fallback node.
	resolveDeploymentID := func(controller string) (string, bool) {
		// Single deployment shortcut.
		if len(deployments) == 1 {
			name := renderer.GetStr(deployments[0], "name", "")
			if name != "" {
				id := "dep-" + renderer.FlowNodeID(name)
				return id, existingNodes[id]
			}
		}

		// Try name matching: controller "FooReconciler" might match deployment "foo-controller".
		controllerLower := strings.ToLower(controller)
		controllerLower = strings.TrimSuffix(controllerLower, "reconciler")
		controllerLower = strings.TrimSuffix(controllerLower, "controller")
		controllerLower = strings.TrimSpace(controllerLower)

		for _, n := range g.Nodes {
			if n.Type != renderer.FlowNodeDeployment {
				continue
			}
			if len(controllerLower) >= 3 && strings.Contains(strings.ToLower(n.Label), controllerLower) {
				return n.ID, true
			}
		}

		// Fallback: create a per-controller node.
		fallbackID := "ctrl-" + renderer.FlowNodeID(controller)
		return fallbackID, existingNodes[fallbackID]
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

	// Process controllers in sorted order for deterministic output.
	ctrlNames := make([]string, 0, len(forByCtrl))
	for ctrl := range forByCtrl {
		ctrlNames = append(ctrlNames, ctrl)
	}
	sort.Strings(ctrlNames)

	for _, ctrl := range ctrlNames {
		forWatches := forByCtrl[ctrl]
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
			crdID := "crd-" + renderer.FlowNodeID(kind)
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
			reconcileEdgeID := fmt.Sprintf("reconcile-%s-%s", renderer.FlowNodeID(kind), depID)
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

				ownedID := "owned-" + renderer.FlowNodeID(ownKind)
				addNode(renderer.FlowNode{
					ID:    ownedID,
					Label: ownKind,
					Type:  renderer.FlowNodeCRD,
					Layer: 6,
				})

				manageEdgeID := fmt.Sprintf("reconcile-%s-%s-%s", depID, renderer.FlowNodeID(ctrl), ownedID)
				addEdge(renderer.FlowEdge{
					ID:    manageEdgeID,
					From:  depID,
					To:    ownedID,
					Type:  "manages",
					Label: "manages",
				})
				pathEdges = append(pathEdges, manageEdgeID)
			}

			// Build reconciliation path (skip if one with this name already exists).
			pathName := kind + " Reconciliation"
			alreadyExists := false
			for _, p := range g.Paths {
				if p.Name == pathName {
					alreadyExists = true
					break
				}
			}
			if !alreadyExists {
				g.Paths = append(g.Paths, renderer.FlowPath{
					Name:  pathName,
					Edges: pathEdges,
					Color: "#0066cc",
				})
			}
		}
	}
}
