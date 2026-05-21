package architecture

import (
	"fmt"
	"strings"

	"github.com/ugiordan/architecture-analyzer/pkg/graph"
	"github.com/ugiordan/architecture-analyzer/pkg/query"
)

func architectureQueries() []query.Rule {
	return []query.Rule{
		{ID: "CGA-A01", Name: "abstraction-layers", Domain: "architecture", Severity: "informational", Run: queryAbstractionLayers},
		{ID: "CGA-A02", Name: "external-api-surface", Domain: "architecture", Severity: "informational", Run: queryExternalAPISurface},
		{ID: "CGA-A03", Name: "factory-dispatch", Domain: "architecture", Severity: "informational", Run: queryFactoryDispatch},
		{ID: "CGA-A04", Name: "unimplemented-interface", Domain: "architecture", Severity: "low", Run: queryUnimplementedInterface},
	}
}

// CGA-A01: Surfaces class hierarchies with abstract bases and their implementations.
func queryAbstractionLayers(g *graph.CPG) []query.Finding {
	var findings []query.Finding

	// Build base → implementations map
	baseToImpls := make(map[string][]string)
	for _, cls := range g.NodesByKind(graph.NodeClass) {
		for _, base := range cls.BaseClasses {
			baseToImpls[base] = append(baseToImpls[base], cls.Name)
		}
	}

	for _, cls := range g.NodesByKind(graph.NodeClass) {
		if !cls.Annotations[AnnotAbstractBase] {
			continue
		}
		impls := baseToImpls[cls.Name]
		if len(impls) == 0 {
			continue
		}
		findings = append(findings, query.Finding{
			RuleID:   "CGA-A01",
			Severity: "informational",
			Message:  fmt.Sprintf("Abstraction layer: %s has %d implementation(s): %s", cls.Name, len(impls), strings.Join(impls, ", ")),
			File:     cls.File,
			Line:     cls.Line,
			NodeID:   cls.ID,
		})
	}
	return findings
}

// CGA-A02: Surfaces functions that use external SDK clients.
func queryExternalAPISurface(g *graph.CPG) []query.Finding {
	var findings []query.Finding
	seen := make(map[string]bool)

	for _, fn := range g.NodesByKind(graph.NodeFunction) {
		if !fn.Annotations[AnnotSDKClient] {
			continue
		}
		if seen[fn.ID] {
			continue
		}
		seen[fn.ID] = true

		// Collect which SDK clients this function uses
		var clients []string
		for _, edge := range g.OutEdges(fn.ID) {
			target := g.GetNode(edge.To)
			if target != nil && target.Annotations[AnnotSDKClient] {
				clients = append(clients, target.Name)
			}
		}

		msg := fmt.Sprintf("Function %s uses external SDK client(s)", fn.Name)
		if len(clients) > 0 {
			msg = fmt.Sprintf("Function %s calls external SDK: %s", fn.Name, strings.Join(clients, ", "))
		}

		findings = append(findings, query.Finding{
			RuleID:   "CGA-A02",
			Severity: "informational",
			Message:  msg,
			File:     fn.File,
			Line:     fn.Line,
			NodeID:   fn.ID,
		})
	}
	return findings
}

// CGA-A03: Surfaces factory functions that dispatch to multiple implementations.
func queryFactoryDispatch(g *graph.CPG) []query.Finding {
	var findings []query.Finding

	for _, fn := range g.NodesByKind(graph.NodeFunction) {
		if !fn.Annotations[AnnotFactoryMethod] {
			continue
		}
		types := fn.Properties["factory_types"]
		findings = append(findings, query.Finding{
			RuleID:   "CGA-A03",
			Severity: "informational",
			Message:  fmt.Sprintf("Factory function %s dispatches to: %s", fn.Name, types),
			File:     fn.File,
			Line:     fn.Line,
			NodeID:   fn.ID,
		})
	}
	return findings
}

// CGA-A04: Abstract bases with no implementations found in the analyzed codebase.
func queryUnimplementedInterface(g *graph.CPG) []query.Finding {
	var findings []query.Finding

	implemented := make(map[string]bool)
	for _, cls := range g.NodesByKind(graph.NodeClass) {
		for _, base := range cls.BaseClasses {
			implemented[base] = true
		}
	}

	for _, cls := range g.NodesByKind(graph.NodeClass) {
		if !cls.Annotations[AnnotAbstractBase] {
			continue
		}
		if implemented[cls.Name] {
			continue
		}
		findings = append(findings, query.Finding{
			RuleID:   "CGA-A04",
			Severity: "low",
			Message:  fmt.Sprintf("Abstract base %s has no implementations found in analyzed sources", cls.Name),
			File:     cls.File,
			Line:     cls.Line,
			NodeID:   cls.ID,
		})
	}
	return findings
}
