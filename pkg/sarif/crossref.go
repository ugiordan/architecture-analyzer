package sarif

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ugiordan/architecture-analyzer/pkg/graph"
	"github.com/ugiordan/architecture-analyzer/pkg/query"
)

// CorrelatedFinding represents an internal finding that has corroborating
// evidence from an external scanner.
type CorrelatedFinding struct {
	InternalFinding query.Finding
	ExternalNodes   []*graph.Node
	Category        string
}

// CrossReference finds internal findings (from domain analyzers) that are
// corroborated by external scanner findings linked to the same CPG node.
// This correlation strengthens confidence: if both CGA-008 (plaintext secrets)
// and semgrep's hardcoded-credential rule flag the same function, the finding
// is higher confidence than either alone.
func CrossReference(cpg *graph.CPG, internalFindings []query.Finding) []CorrelatedFinding {
	// Build index: nodeID -> external findings linked to that node via REPORTED_BY
	nodeToExternal := make(map[string][]*graph.Node)
	for _, ef := range cpg.SARIFFindings() {
		// Find which CPG nodes have REPORTED_BY edges pointing to this external finding
		edges := cpg.InEdges(ef.ID)
		for _, e := range edges {
			if e.Kind == graph.EdgeReportedBy {
				nodeToExternal[e.From] = append(nodeToExternal[e.From], ef)
			}
		}
	}

	var correlated []CorrelatedFinding
	for _, f := range internalFindings {
		if f.NodeID == "" {
			continue
		}
		if externals, ok := nodeToExternal[f.NodeID]; ok {
			cat := ""
			for _, ext := range externals {
				if c := CategoryForCWEs(ext.CWEs); c != "" {
					cat = c
					break
				}
			}
			correlated = append(correlated, CorrelatedFinding{
				InternalFinding: f,
				ExternalNodes:   externals,
				Category:        cat,
			})
		}
	}
	return correlated
}

// FormatCorrelations returns a human-readable summary of cross-referenced findings.
func FormatCorrelations(correlations []CorrelatedFinding) string {
	if len(correlations) == 0 {
		return ""
	}
	var b strings.Builder
	fmt.Fprintf(&b, "\n%d cross-referenced finding(s) (internal + external corroboration):\n", len(correlations))
	for _, c := range correlations {
		tools := make(map[string]bool)
		for _, ext := range c.ExternalNodes {
			tools[ext.ToolName] = true
		}
		var toolNames []string
		for t := range tools {
			toolNames = append(toolNames, t)
		}
		sort.Strings(toolNames)
		fmt.Fprintf(&b, "  %s %s (%s:%d) corroborated by %s",
			c.InternalFinding.RuleID,
			c.InternalFinding.Message,
			c.InternalFinding.File,
			c.InternalFinding.Line,
			strings.Join(toolNames, ", "))
		if c.Category != "" {
			fmt.Fprintf(&b, " [%s]", c.Category)
		}
		b.WriteString("\n")
	}
	return b.String()
}
