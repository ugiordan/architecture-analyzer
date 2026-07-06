package compile

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/ugiordan/architecture-analyzer/pkg/srclang"
)

type platformFile struct {
	Platform       string               `json:"platform"`
	Version        string               `json:"version"`
	ComponentCount int                  `json:"component_count"`
	Components     []string             `json:"components"`
	DependencyGraph []platformEdgeRaw   `json:"dependency_graph"`
	ComponentData  []platformComponent  `json:"component_data"`
}

type platformEdgeRaw struct {
	From string `json:"from"`
	To   string `json:"to"`
	Type string `json:"type"`
}

type platformComponent struct {
	Component   string                `json:"component"`
	Deployments []platformDeployment  `json:"deployments"`
}

type platformDeployment struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

func extractPlatform(path string, component string) (*srclang.Platform, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading platform file: %w", err)
	}

	var pf platformFile
	if err := json.Unmarshal(data, &pf); err != nil {
		return nil, fmt.Errorf("parsing platform file: %w", err)
	}

	platform := &srclang.Platform{
		Name:       pf.Platform,
		Version:    pf.Version,
		Components: pf.ComponentCount,
	}

	seen := make(map[string]bool)
	for _, comp := range pf.ComponentData {
		for _, dep := range comp.Deployments {
			ns := dep.Namespace
			if ns == "" || seen[ns] {
				continue
			}
			seen[ns] = true
			platform.Topology = append(platform.Topology, srclang.Namespace{
				Name: ns,
			})
		}
	}

	for _, edge := range pf.DependencyGraph {
		edgeType, target := splitEdgeType(edge.Type)

		if edge.To == component && edge.From != component {
			platform.Inbound = append(platform.Inbound, srclang.PlatformEdge{
				Peer:   edge.From,
				Type:   edgeType,
				Target: target,
			})
		}
		if edge.From == component && edge.To != component {
			platform.Outbound = append(platform.Outbound, srclang.PlatformEdge{
				Peer:   edge.To,
				Type:   edgeType,
				Target: target,
			})
		}
	}

	return platform, nil
}

func splitEdgeType(raw string) (edgeType, target string) {
	if idx := strings.IndexByte(raw, ':'); idx >= 0 {
		return raw[:idx], raw[idx+1:]
	}
	return raw, ""
}
