package layers

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ugiordan/architecture-analyzer/pkg/extractor"
	"github.com/ugiordan/architecture-analyzer/pkg/graph"
	"github.com/ugiordan/architecture-analyzer/pkg/query"
)

func TestNetpolicySelector_IncludesNetworkPolicies(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component: "test",
		NetworkPolicies: []extractor.NetworkPolicy{
			{
				Name:        "allow-dashboard",
				Source:      "manifests/np.yaml",
				PolicyTypes: []string{"Ingress"},
				PodSelector: map[string]interface{}{"matchLabels": map[string]interface{}{"app": "dashboard"}},
				IngressRules: []map[string]interface{}{
					{"from": []interface{}{map[string]interface{}{"namespaceSelector": map[string]interface{}{"matchLabels": map[string]interface{}{"opendatahub.io/dashboard": "true"}}}}},
				},
			},
		},
	}

	sel := NewNetpolicySelector(t.TempDir())
	layer, _ := sel.Select(cpg, arch, nil, nil)

	if layer.Name != "netpolicy" {
		t.Errorf("layer name = %q, want %q", layer.Name, "netpolicy")
	}

	if len(layer.Resources) == 0 {
		t.Fatal("expected at least one NetworkPolicy resource")
	}

	r := layer.Resources[0]
	if r.Kind != "NetworkPolicy" {
		t.Errorf("kind = %q, want NetworkPolicy", r.Kind)
	}
	if r.Name != "allow-dashboard" {
		t.Errorf("name = %q, want allow-dashboard", r.Name)
	}
	if r.Origin != "manifest" {
		t.Errorf("origin = %q, want manifest", r.Origin)
	}
	if !strings.Contains(r.Summary, "Ingress") {
		t.Errorf("summary should contain policy type, got %q", r.Summary)
	}
	if !strings.Contains(r.Summary, "1 ingress") {
		t.Errorf("summary should contain ingress count, got %q", r.Summary)
	}
}

func TestNetpolicySelector_ProgrammaticOrigin(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component: "test",
		NetworkPolicies: []extractor.NetworkPolicy{
			{Name: "dynamic-np", Source: "pkg/controller/netpol.go"},
		},
	}

	sel := NewNetpolicySelector(t.TempDir())
	layer, _ := sel.Select(cpg, arch, nil, nil)

	if len(layer.Resources) == 0 {
		t.Fatal("expected resource")
	}
	if layer.Resources[0].Origin != "programmatic" {
		t.Errorf("origin = %q, want programmatic", layer.Resources[0].Origin)
	}
}

func TestNetpolicySelector_FiltersNetpolicyFindings(t *testing.T) {
	cpg := graph.NewCPG()
	findings := []query.Finding{
		{RuleID: "CGA-N01", Severity: "high", Domain: "netpolicy", Message: "Bare namespaceSelector", File: "np.yaml", Line: 1},
		{RuleID: "CGA-N02", Severity: "medium", Domain: "netpolicy", Message: "Tenant reach", File: "np.yaml", Line: 5},
		{RuleID: "CGA-S01", Severity: "high", Domain: "security", Message: "Security issue", File: "a.go", Line: 1},
		{RuleID: "CGA-U01", Severity: "medium", Domain: "upgrade", Message: "Upgrade issue", File: "b.go", Line: 1},
	}

	sel := NewNetpolicySelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, findings, nil)

	if len(layer.Findings) != 2 {
		t.Fatalf("expected 2 netpolicy findings, got %d", len(layer.Findings))
	}
	for _, f := range layer.Findings {
		if f.Domain != "netpolicy" {
			t.Errorf("unexpected domain %q", f.Domain)
		}
	}
}

func TestNetpolicySelector_FindingsSorted(t *testing.T) {
	cpg := graph.NewCPG()
	findings := []query.Finding{
		{RuleID: "CGA-N02", Severity: "medium", Domain: "netpolicy", Message: "Medium", File: "a.yaml", Line: 1},
		{RuleID: "CGA-N01", Severity: "high", Domain: "netpolicy", Message: "High", File: "b.yaml", Line: 2},
	}

	sel := NewNetpolicySelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, findings, nil)

	if layer.Findings[0].Severity != "high" {
		t.Errorf("first finding severity = %q, want high", layer.Findings[0].Severity)
	}
	if layer.Findings[1].Severity != "medium" {
		t.Errorf("second finding severity = %q, want medium", layer.Findings[1].Severity)
	}
}

func TestNetpolicySelector_SelectsAnnotatedFunctions(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "createNetworkPolicy",
		File: "netpol.go", Line: 10, EndLine: 30,
		Language:    "go",
		Annotations: map[string]bool{annotNetPolSelector: true},
	})
	cpg.AddNode(&graph.Node{
		ID: "fn2", Kind: graph.NodeFunction, Name: "checkTenantReach",
		File: "trust.go", Line: 5,
		Language:    "go",
		Annotations: map[string]bool{annotNetPolTenantReach: true},
	})
	cpg.AddNode(&graph.Node{
		ID: "fn3", Kind: graph.NodeFunction, Name: "wideOpenPolicy",
		File: "open.go", Line: 1,
		Language:    "go",
		Annotations: map[string]bool{annotNetPolNoRestrict: true},
	})

	dir := t.TempDir()
	writeTestFile(t, dir, "netpol.go", 35)

	sel := NewNetpolicySelector(dir)
	layer, _ := sel.Select(cpg, nil, nil, nil)

	names := make(map[string]bool)
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			names[fn.Name] = true
		}
	}

	if !names["createNetworkPolicy"] {
		t.Error("createNetworkPolicy not selected (has namespace_selector annotation)")
	}
	if !names["checkTenantReach"] {
		t.Error("checkTenantReach not selected (has tenant_reach annotation)")
	}
	if !names["wideOpenPolicy"] {
		t.Error("wideOpenPolicy not selected (has no_restriction annotation)")
	}
}

func TestNetpolicySelector_Metas(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "setupNetPol",
		File: "np.go", Line: 1,
		Language: "go",
		Annotations: map[string]bool{
			annotNetPolSelector:    true,
			annotNetPolTenantReach: true,
			annotNetPolNoRestrict:  true,
		},
	})

	sel := NewNetpolicySelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	metaKeys := make(map[string]bool)
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "setupNetPol" {
				for _, m := range fn.Metas {
					metaKeys[m.Key] = true
				}
			}
		}
	}

	for _, expected := range []string{"namespace-selector", "tenant-reach", "no-restriction"} {
		if !metaKeys[expected] {
			t.Errorf("expected meta %q", expected)
		}
	}
}

func TestNetpolicySelector_FindingReferencedFunctions(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "applyLabels",
		File: "labels.go", Line: 10, EndLine: 40,
		Language: "go",
	})

	findings := []query.Finding{
		{RuleID: "CGA-N01", Severity: "high", Domain: "netpolicy", Message: "issue",
			File: "labels.go", Line: 25},
	}

	dir := t.TempDir()
	writeTestFile(t, dir, "labels.go", 50)

	sel := NewNetpolicySelector(dir)
	layer, _ := sel.Select(cpg, nil, findings, nil)

	found := false
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "applyLabels" {
				found = true
			}
		}
	}
	if !found {
		t.Error("applyLabels should be included (finding at line 25 is within 10-40)")
	}
}

func TestNetpolicySelector_CallRelationships(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "createNetPol",
		File: "netpol.go", Line: 10,
		Annotations: map[string]bool{annotNetPolSelector: true},
	})
	cpg.AddNode(&graph.Node{
		ID: "fn2", Kind: graph.NodeFunction, Name: "helper",
		File: "helper.go", Line: 5,
	})
	cpg.AddEdge(&graph.Edge{From: "fn1", To: "fn2", Kind: graph.EdgeCalls})

	sel := NewNetpolicySelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	if len(layer.Relationships) == 0 {
		t.Fatal("expected at least one relationship")
	}
	if layer.Relationships[0].Kind != "calls" {
		t.Errorf("kind = %q, want calls", layer.Relationships[0].Kind)
	}
	if layer.Relationships[0].To.Resolved == nil || *layer.Relationships[0].To.Resolved != false {
		t.Error("expected To.Resolved = false for non-selected function")
	}
}

func TestNetpolicySelector_EmptyInputs(t *testing.T) {
	cpg := graph.NewCPG()
	sel := NewNetpolicySelector(t.TempDir())
	layer, warnings := sel.Select(cpg, nil, nil, nil)

	if layer.Name != "netpolicy" {
		t.Errorf("layer name = %q", layer.Name)
	}
	if len(layer.Files) != 0 {
		t.Errorf("expected 0 files, got %d", len(layer.Files))
	}
	if len(layer.Findings) != 0 {
		t.Errorf("expected 0 findings, got %d", len(layer.Findings))
	}
	if len(layer.Resources) != 0 {
		t.Errorf("expected 0 resources, got %d", len(layer.Resources))
	}
	if len(layer.Relationships) != 0 {
		t.Errorf("expected 0 relationships, got %d", len(layer.Relationships))
	}
	if len(warnings) != 0 {
		t.Errorf("expected 0 warnings, got %d", len(warnings))
	}
}

func TestNetpolicySelector_PathTraversalRejected(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "evilFunc",
		File: "../../../etc/passwd", Line: 1, EndLine: 5,
		Language:    "go",
		Annotations: map[string]bool{annotNetPolSelector: true},
	})

	sel := NewNetpolicySelector(t.TempDir())
	layer, warnings := sel.Select(cpg, nil, nil, nil)

	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Code != "" {
				t.Error("path traversal should not extract code")
			}
		}
	}

	foundWarning := false
	for _, w := range warnings {
		if strings.Contains(w.Message, "path traversal") {
			foundWarning = true
		}
	}
	if !foundWarning {
		t.Error("expected path traversal warning")
	}
}

func TestNetpolicySelector_ExcludesNonNetpolicyFunctions(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "RegularFunc",
		File: "regular.go", Line: 5,
		Language: "go",
	})

	sel := NewNetpolicySelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	if len(layer.Files) != 0 {
		t.Errorf("expected no files, got %d", len(layer.Files))
	}
}

func TestNetpolicySelector_BodyExtractionFailure(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "badFunc",
		File: "short.go", Line: 50, EndLine: 60,
		Language:    "go",
		Annotations: map[string]bool{annotNetPolSelector: true},
	})

	dir := t.TempDir()
	writeTestFile(t, dir, "short.go", 5)

	sel := NewNetpolicySelector(dir)
	layer, warnings := sel.Select(cpg, nil, nil, nil)

	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "badFunc" && fn.Code != "" {
				t.Error("should not have code body")
			}
		}
	}

	foundWarning := false
	for _, w := range warnings {
		if strings.Contains(w.Message, "body extraction failed") {
			foundWarning = true
		}
	}
	if !foundWarning {
		t.Error("expected body extraction failure warning")
	}
}

func TestNetpolicySelector_RelationshipCap(t *testing.T) {
	cpg := graph.NewCPG()

	for i := 0; i < 200; i++ {
		fnID := fmt.Sprintf("fn%d", i)
		helperID := fmt.Sprintf("helper%d", i)

		cpg.AddNode(&graph.Node{
			ID: fnID, Kind: graph.NodeFunction, Name: fmt.Sprintf("NetPol%d", i),
			File: fmt.Sprintf("np_%d.go", i), Line: 1,
			Language:    "go",
			Annotations: map[string]bool{annotNetPolSelector: true},
		})
		cpg.AddNode(&graph.Node{
			ID: helperID, Kind: graph.NodeFunction, Name: fmt.Sprintf("helper%d", i),
			File: fmt.Sprintf("h_%d.go", i), Line: 1,
			Language: "go",
		})
		cpg.AddEdge(&graph.Edge{From: fnID, To: helperID, Kind: graph.EdgeCalls})
	}

	sel := NewNetpolicySelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	if len(layer.Relationships) > maxRelationships {
		t.Errorf("relationships = %d, should be capped at %d", len(layer.Relationships), maxRelationships)
	}
}

func TestNetpolicySelector_NoDuplicateFindingReferencedFunctions(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "createNetPol",
		File: "netpol.go", Line: 10, EndLine: 30,
		Language:    "go",
		Annotations: map[string]bool{annotNetPolSelector: true},
	})

	findings := []query.Finding{
		{RuleID: "CGA-N01", Severity: "high", Domain: "netpolicy", Message: "issue",
			File: "netpol.go", Line: 15},
	}

	dir := t.TempDir()
	writeTestFile(t, dir, "netpol.go", 40)

	sel := NewNetpolicySelector(dir)
	layer, _ := sel.Select(cpg, nil, findings, nil)

	count := 0
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "createNetPol" {
				count++
			}
		}
	}
	if count != 1 {
		t.Errorf("createNetPol should appear exactly once, got %d", count)
	}
}

func TestNetpolicySelector_NetworkPolicyWithIssues(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component: "test",
		NetworkPolicies: []extractor.NetworkPolicy{
			{
				Name:   "problem-np",
				Source: "manifests/np.yaml",
				Issues: []string{"missing egress policy", "overly permissive"},
			},
		},
	}

	sel := NewNetpolicySelector(t.TempDir())
	layer, _ := sel.Select(cpg, arch, nil, nil)

	if len(layer.Resources) == 0 {
		t.Fatal("expected resource")
	}
	if !strings.Contains(layer.Resources[0].Summary, "2 issues") {
		t.Errorf("summary should mention issues, got %q", layer.Resources[0].Summary)
	}

	issueCount := 0
	for _, c := range layer.Resources[0].Children {
		if strings.Contains(c.XMLContent, "<issue>") {
			issueCount++
		}
	}
	if issueCount != 2 {
		t.Errorf("expected 2 issue children, got %d", issueCount)
	}
}

func TestNetpolicySelector_RiskFindingsFromNetworkPolicies(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component: "test",
		NetworkPolicies: []extractor.NetworkPolicy{
			{
				Name:        "missing-egress",
				Source:      "manifests/np.yaml",
				PolicyTypes: []string{"Ingress"},
			},
		},
	}

	sel := NewNetpolicySelector(t.TempDir())
	layer, _ := sel.Select(cpg, arch, nil, nil)

	foundRisk := false
	for _, f := range layer.Findings {
		if f.Rule == "NETPOL_MISSING_EGRESS_POLICY" {
			foundRisk = true
		}
	}
	if !foundRisk {
		t.Error("expected NETPOL_MISSING_EGRESS_POLICY risk finding")
	}
}
