package flow

import "sort"

// ApplyLayout assigns x/y positions to all nodes in the diagram based on
// their layer field, using a BFS-style layered layout. It also computes and
// sets the canvas dimensions.
func ApplyLayout(d *Diagram) {
	if len(d.Nodes) == 0 {
		d.Canvas = DiagramCanvas{
			Width:  max(1400, 0),
			Height: max(900, 0),
		}
		return
	}

	// Group node IDs by layer.
	layers := map[int][]string{}
	for id, n := range d.Nodes {
		layers[n.Layer()] = append(layers[n.Layer()], id)
	}

	// Sort IDs within each layer for deterministic output.
	for _, ids := range layers {
		sort.Strings(ids)
	}

	// Collect and sort the distinct layer indices.
	layerIndices := make([]int, 0, len(layers))
	for l := range layers {
		layerIndices = append(layerIndices, l)
	}
	sort.Ints(layerIndices)

	layerCount := len(layerIndices)

	// Find the maximum number of nodes in any single layer.
	maxNodesInLayer := 0
	for _, ids := range layers {
		if len(ids) > maxNodesInLayer {
			maxNodesInLayer = len(ids)
		}
	}

	// Compute canvas dimensions.
	canvasWidth := max(layerCount*240+160, 1400)
	verticalSpacing := max(80, 900/(maxNodesInLayer+1))
	canvasHeight := max(maxNodesInLayer*verticalSpacing+100, 900)

	// Position each node within its layer.
	for i, layerIdx := range layerIndices {
		ids := layers[layerIdx]
		x := i*240 + 80

		totalLayerHeight := len(ids) * verticalSpacing
		verticalOffset := (canvasHeight-totalLayerHeight)/2 + verticalSpacing/2

		for j, id := range ids {
			n := d.Nodes[id]
			n.X = x
			n.Y = verticalOffset + j*verticalSpacing
			d.Nodes[id] = n
		}
	}

	d.Canvas = DiagramCanvas{Width: canvasWidth, Height: canvasHeight}
}
