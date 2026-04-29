package sarif_test

import (
	"testing"

	"github.com/ugiordan/architecture-analyzer/pkg/graph"
	"github.com/ugiordan/architecture-analyzer/pkg/query"
	"github.com/ugiordan/architecture-analyzer/pkg/sarif"
)

func TestCrossReferenceFindsCorroboration(t *testing.T) {
	fnID := graph.NodeID(graph.NodeFunction, "handleLogin", "auth.go", 10, 0)
	fn := &graph.Node{
		ID: fnID, Kind: graph.NodeFunction, Name: "handleLogin",
		File: "auth.go", Line: 10, EndLine: 50,
	}
	cpg := graph.NewCPG()
	if err := cpg.AddNode(fn); err != nil {
		t.Fatal(err)
	}

	// Ingest external finding linked to same function
	report := &sarif.Report{
		Version: "2.1.0",
		Runs: []sarif.Run{{
			Tool: sarif.Tool{Driver: sarif.ToolComponent{
				Name: "semgrep",
				Rules: []sarif.Rule{{
					ID:         "hardcoded-cred",
					Properties: sarif.RuleProperties{Tags: []string{"CWE-798"}},
				}},
			}},
			Results: []sarif.Result{{
				RuleID:  "hardcoded-cred",
				Level:   "error",
				Message: sarif.Message{Text: "hardcoded credential"},
				Locations: []sarif.Location{{
					PhysicalLocation: sarif.PhysicalLocation{
						ArtifactLocation: sarif.ArtifactLocation{URI: "auth.go"},
						Region:           sarif.Region{StartLine: 10, StartColumn: 1},
					},
				}},
			}},
		}},
	}
	sarif.Ingest(cpg, report, "")

	// Internal finding at the same node
	internalFindings := []query.Finding{{
		RuleID:  "CGA-008",
		Message: "Plaintext secret in handleLogin",
		File:    "auth.go",
		Line:    10,
		NodeID:  fnID,
		Domain:  "security",
	}}

	correlated := sarif.CrossReference(cpg, internalFindings)
	if len(correlated) != 1 {
		t.Fatalf("correlated = %d, want 1", len(correlated))
	}
	if correlated[0].InternalFinding.RuleID != "CGA-008" {
		t.Errorf("RuleID = %q", correlated[0].InternalFinding.RuleID)
	}
	if correlated[0].Category != "hardcoded-credentials" {
		t.Errorf("Category = %q, want hardcoded-credentials", correlated[0].Category)
	}
	if len(correlated[0].ExternalNodes) != 1 {
		t.Errorf("ExternalNodes = %d, want 1", len(correlated[0].ExternalNodes))
	}
}

func TestCrossReferenceNoCorroboration(t *testing.T) {
	cpg := graph.NewCPG()

	findings := []query.Finding{{
		RuleID: "CGA-001",
		NodeID: "some-node",
		File:   "a.go",
		Line:   1,
	}}

	correlated := sarif.CrossReference(cpg, findings)
	if len(correlated) != 0 {
		t.Errorf("correlated = %d, want 0", len(correlated))
	}
}

func TestCrossReferenceSkipsEmptyNodeID(t *testing.T) {
	cpg := graph.NewCPG()

	findings := []query.Finding{{
		RuleID: "CGA-001",
		NodeID: "",
		File:   "a.go",
		Line:   1,
	}}

	correlated := sarif.CrossReference(cpg, findings)
	if len(correlated) != 0 {
		t.Errorf("correlated = %d, want 0 (empty NodeID skipped)", len(correlated))
	}
}

func TestFormatCorrelationsEmpty(t *testing.T) {
	result := sarif.FormatCorrelations(nil)
	if result != "" {
		t.Errorf("expected empty string for nil correlations")
	}
}
