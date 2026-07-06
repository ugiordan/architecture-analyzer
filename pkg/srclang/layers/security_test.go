package layers

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ugiordan/architecture-analyzer/pkg/extractor"
	"github.com/ugiordan/architecture-analyzer/pkg/graph"
	"github.com/ugiordan/architecture-analyzer/pkg/query"
)

func TestSecuritySelector_SelectsWebhookHandlers(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "Reconcile",
		File: "controller.go", Line: 10, EndLine: 50,
		Language: "go",
		Annotations: map[string]bool{"webhook_handler": true},
	})
	cpg.AddNode(&graph.Node{
		ID: "fn2", Kind: graph.NodeFunction, Name: "helperFunc",
		File: "utils.go", Line: 5, EndLine: 15,
		Language: "go",
	})

	dir := t.TempDir()
	writeTestFile(t, dir, "controller.go", 60)
	writeTestFile(t, dir, "utils.go", 20)

	arch := &extractor.ComponentArchitecture{
		Component: "test-operator",
		Webhooks: []extractor.WebhookConfig{
			{Name: "validate-foo", Type: "validating"},
		},
	}

	sel := NewSecuritySelector(dir)
	layer, _ := sel.Select(cpg, arch, nil, nil)

	if layer.Name != "security" {
		t.Errorf("layer name = %q, want %q", layer.Name, "security")
	}

	funcCount := 0
	for _, f := range layer.Files {
		funcCount += len(f.Functions)
	}
	if funcCount == 0 {
		t.Error("expected at least one security-relevant function")
	}

	found := false
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "Reconcile" {
				found = true
				if fn.Code == "" {
					t.Error("Reconcile should have code body extracted")
				}
			}
		}
	}
	if !found {
		t.Error("Reconcile function not selected by security layer")
	}
}

func TestSecuritySelector_IncludesFindings(t *testing.T) {
	cpg := graph.NewCPG()
	findings := []query.Finding{
		{
			RuleID:   "CGA-N01",
			Severity: "high",
			Message:  "Bare namespaceSelector",
			File:     "utils.go",
			Line:     160,
			Domain:   "netpolicy",
		},
	}

	dir := t.TempDir()
	arch := &extractor.ComponentArchitecture{Component: "test"}

	sel := NewSecuritySelector(dir)
	layer, _ := sel.Select(cpg, arch, findings, nil)

	if len(layer.Findings) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(layer.Findings))
	}
	if layer.Findings[0].Severity != "high" {
		t.Errorf("severity = %q, want %q", layer.Findings[0].Severity, "high")
	}
}

func TestSecuritySelector_IncludesNetworkPolicies(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component: "test",
		NetworkPolicies: []extractor.NetworkPolicy{
			{Name: "default-np", Source: "manifests/np.yaml"},
		},
	}

	dir := t.TempDir()
	sel := NewSecuritySelector(dir)
	layer, _ := sel.Select(cpg, arch, nil, nil)

	if len(layer.Resources) == 0 {
		t.Error("expected NetworkPolicy resource in output")
	}
}

func TestSecuritySelector_IncludesRBAC(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component: "test",
		RBAC: &extractor.RBACData{
			ClusterRoles: []extractor.RBACRole{
				{Name: "manager-role", Rules: make([]extractor.RBACRule, 45), Source: "config/rbac/role.yaml"},
			},
		},
	}

	dir := t.TempDir()
	sel := NewSecuritySelector(dir)
	layer, _ := sel.Select(cpg, arch, nil, nil)

	foundRBAC := false
	for _, r := range layer.Resources {
		if r.Kind == "ClusterRole" {
			foundRBAC = true
		}
	}
	if !foundRBAC {
		t.Error("expected ClusterRole resource")
	}
}

func TestSecuritySelector_TaintAnnotations(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "HandleRequest",
		File: "handler.go", Line: 10, EndLine: 20,
		Language:    "go",
		Annotations: map[string]bool{"taint_source": true},
	})
	cpg.AddNode(&graph.Node{
		ID: "fn2", Kind: graph.NodeFunction, Name: "ExecuteSQL",
		File: "db.go", Line: 5, EndLine: 15,
		Language:    "go",
		Annotations: map[string]bool{"taint_sink": true},
	})

	dir := t.TempDir()
	writeTestFile(t, dir, "handler.go", 25)
	writeTestFile(t, dir, "db.go", 20)

	arch := &extractor.ComponentArchitecture{Component: "test"}
	sel := NewSecuritySelector(dir)
	layer, _ := sel.Select(cpg, arch, nil, nil)

	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "HandleRequest" {
				if fn.TaintRole != "source" {
					t.Errorf("HandleRequest TaintRole = %q, want %q", fn.TaintRole, "source")
				}
				if fn.Trust != "untrusted" {
					t.Errorf("HandleRequest Trust = %q, want %q", fn.Trust, "untrusted")
				}
			}
			if fn.Name == "ExecuteSQL" {
				if fn.TaintRole != "sink" {
					t.Errorf("ExecuteSQL TaintRole = %q, want %q", fn.TaintRole, "sink")
				}
			}
		}
	}
}

func TestSecuritySelector_CallRelationships(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "Validate",
		File:        "validator.go",
		Line:        10,
		Annotations: map[string]bool{"input_validator": true},
	})
	cpg.AddNode(&graph.Node{
		ID:   "fn2",
		Kind: graph.NodeFunction,
		Name: "helper",
		File: "helper.go",
		Line: 5,
	})
	cpg.AddEdge(&graph.Edge{
		From: "fn1",
		To:   "fn2",
		Kind: graph.EdgeCalls,
	})

	dir := t.TempDir()
	arch := &extractor.ComponentArchitecture{Component: "test"}
	sel := NewSecuritySelector(dir)
	layer, _ := sel.Select(cpg, arch, nil, nil)

	if len(layer.Relationships) == 0 {
		t.Fatal("expected at least one relationship")
	}

	rel := layer.Relationships[0]
	if rel.Kind != "calls" {
		t.Errorf("relationship kind = %q, want %q", rel.Kind, "calls")
	}
	if rel.From.Function != "Validate" {
		t.Errorf("from function = %q, want %q", rel.From.Function, "Validate")
	}
	if rel.To.Function != "helper" {
		t.Errorf("to function = %q, want %q", rel.To.Function, "helper")
	}
	if rel.To.Resolved == nil || *rel.To.Resolved != false {
		t.Error("expected To.Resolved = false for non-selected function")
	}
}

func TestSecuritySelector_HTTPEndpoints(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID:   "ep1",
		Kind: graph.NodeHTTPEndpoint,
		Name: "/api/v1/models",
		File: "routes.go",
		Line: 100,
	})

	dir := t.TempDir()
	arch := &extractor.ComponentArchitecture{Component: "test"}
	sel := NewSecuritySelector(dir)
	layer, _ := sel.Select(cpg, arch, nil, nil)

	found := false
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "/api/v1/models" && fn.Kind == "handler" {
				found = true
			}
		}
	}
	if !found {
		t.Error("expected HTTP endpoint to be selected")
	}
}

func TestSecuritySelector_FindingsSorting(t *testing.T) {
	cpg := graph.NewCPG()
	findings := []query.Finding{
		{RuleID: "CGA-001", Severity: "low", Domain: "security", Message: "Low severity", File: "a.go", Line: 1},
		{RuleID: "CGA-002", Severity: "critical", Domain: "security", Message: "Critical severity", File: "b.go", Line: 2},
		{RuleID: "CGA-003", Severity: "medium", Domain: "security", Message: "Medium severity", File: "c.go", Line: 3},
		{RuleID: "CGA-004", Severity: "high", Domain: "security", Message: "High severity", File: "d.go", Line: 4},
	}

	dir := t.TempDir()
	arch := &extractor.ComponentArchitecture{Component: "test"}
	sel := NewSecuritySelector(dir)
	layer, _ := sel.Select(cpg, arch, findings, nil)

	if len(layer.Findings) != 4 {
		t.Fatalf("expected 4 findings, got %d", len(layer.Findings))
	}

	expected := []string{"critical", "high", "medium", "low"}
	for i, f := range layer.Findings {
		if f.Severity != expected[i] {
			t.Errorf("finding[%d] severity = %q, want %q", i, f.Severity, expected[i])
		}
	}
}

func TestSecuritySelector_FiltersNonSecurityFindings(t *testing.T) {
	cpg := graph.NewCPG()
	findings := []query.Finding{
		{RuleID: "CGA-001", Severity: "high", Domain: "security", Message: "Security issue"},
		{RuleID: "CGA-002", Severity: "high", Domain: "performance", Message: "Performance issue"},
		{RuleID: "CGA-003", Severity: "high", Domain: "netpolicy", Message: "Network policy issue"},
		{RuleID: "CGA-004", Severity: "high", Domain: "other", Message: "Other issue"},
	}

	dir := t.TempDir()
	arch := &extractor.ComponentArchitecture{Component: "test"}
	sel := NewSecuritySelector(dir)
	layer, _ := sel.Select(cpg, arch, findings, nil)

	if len(layer.Findings) != 2 {
		t.Errorf("expected 2 findings (security + netpolicy), got %d", len(layer.Findings))
	}
}

func TestSecuritySelector_EmptyInputs(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{Component: "test"}
	dir := t.TempDir()

	sel := NewSecuritySelector(dir)
	layer, warnings := sel.Select(cpg, arch, nil, nil)

	if layer.Name != "security" {
		t.Errorf("layer name = %q, want %q", layer.Name, "security")
	}
	if len(layer.Files) != 0 {
		t.Errorf("expected no files for empty CPG, got %d", len(layer.Files))
	}
	if len(layer.Resources) != 0 {
		t.Errorf("expected no resources for empty arch, got %d", len(layer.Resources))
	}
	if len(layer.Findings) != 0 {
		t.Errorf("expected no findings for nil findings, got %d", len(layer.Findings))
	}
	if len(warnings) != 0 {
		t.Errorf("expected no warnings, got %d", len(warnings))
	}
}

func TestSecuritySelector_DBOperations(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID:   "db1",
		Kind: graph.NodeDBOperation,
		Name: "UpdateUser",
		File: "database.go",
		Line: 42,
	})

	dir := t.TempDir()
	arch := &extractor.ComponentArchitecture{Component: "test"}
	sel := NewSecuritySelector(dir)
	layer, _ := sel.Select(cpg, arch, nil, nil)

	found := false
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "UpdateUser" {
				found = true
			}
		}
	}
	if !found {
		t.Error("expected DBOperation node to be selected")
	}
}

func TestSecuritySelector_IncludesExtractionAnnotations(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{Component: "test"}
	annotations := []extractor.SecurityAnnotation{
		{
			Type:        "RBAC_CLUSTER_SCOPE_SENSITIVE",
			Severity:    "high",
			Source:      "config/rbac/role.yaml",
			Description: "ClusterRole grants cluster-wide secrets CRUD",
		},
		{
			Type:        "SECRET_IN_CONTAINER_ARGS",
			Severity:    "medium",
			Source:      "manifests/deployment.yaml",
			Description: "Container uses $(SECRET_KEY) in args",
		},
	}

	dir := t.TempDir()
	sel := NewSecuritySelector(dir)
	layer, _ := sel.Select(cpg, arch, nil, annotations)

	if len(layer.Findings) != 2 {
		t.Fatalf("expected 2 extraction findings, got %d", len(layer.Findings))
	}

	// Sorted by severity: high before medium
	if layer.Findings[0].Domain != "extraction" {
		t.Errorf("finding[0] domain = %q, want %q", layer.Findings[0].Domain, "extraction")
	}
	if layer.Findings[0].Rule != "RBAC_CLUSTER_SCOPE_SENSITIVE" {
		t.Errorf("finding[0] rule = %q, want %q", layer.Findings[0].Rule, "RBAC_CLUSTER_SCOPE_SENSITIVE")
	}
	if layer.Findings[0].Severity != "high" {
		t.Errorf("finding[0] severity = %q, want %q", layer.Findings[0].Severity, "high")
	}
	if layer.Findings[1].Rule != "SECRET_IN_CONTAINER_ARGS" {
		t.Errorf("finding[1] rule = %q, want %q", layer.Findings[1].Rule, "SECRET_IN_CONTAINER_ARGS")
	}
	if layer.Findings[1].Severity != "medium" {
		t.Errorf("finding[1] severity = %q, want %q", layer.Findings[1].Severity, "medium")
	}
}

func TestSecuritySelector_MergesCPGAndExtractionFindings(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{Component: "test"}

	cpgFindings := []query.Finding{
		{RuleID: "CGA-N01", Severity: "high", Domain: "netpolicy", Message: "Bare namespaceSelector", File: "utils.go", Line: 160},
	}
	annotations := []extractor.SecurityAnnotation{
		{Type: "CRD_CONFUSED_DEPUTY", Severity: "high", Source: "config/crd/test.yaml", Description: "CRD has user-settable image field"},
	}

	dir := t.TempDir()
	sel := NewSecuritySelector(dir)
	layer, _ := sel.Select(cpg, arch, cpgFindings, annotations)

	if len(layer.Findings) != 2 {
		t.Fatalf("expected 2 merged findings (1 CPG + 1 extraction), got %d", len(layer.Findings))
	}

	domains := map[string]bool{}
	for _, f := range layer.Findings {
		domains[f.Domain] = true
	}
	if !domains["netpolicy"] {
		t.Error("expected netpolicy domain finding")
	}
	if !domains["extraction"] {
		t.Error("expected extraction domain finding")
	}
}

func writeTestFile(t *testing.T, dir, name string, lines int) {
	t.Helper()
	var sb []byte
	for i := 1; i <= lines; i++ {
		sb = append(sb, []byte("// line "+string(rune('0'+i%10))+"\n")...)
	}
	if err := os.WriteFile(filepath.Join(dir, name), sb, 0o644); err != nil {
		t.Fatal(err)
	}
}
