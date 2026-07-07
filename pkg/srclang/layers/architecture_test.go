package layers

import (
	"strings"
	"testing"

	"github.com/ugiordan/architecture-analyzer/pkg/extractor"
	"github.com/ugiordan/architecture-analyzer/pkg/graph"
	"github.com/ugiordan/architecture-analyzer/pkg/query"
)

func TestArchitectureSelector_IncludesCRDs(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component: "test",
		CRDs: []extractor.CRD{
			{Group: "serving.kserve.io", Version: "v1beta1", Kind: "InferenceService", Scope: "Namespaced", FieldsCount: 142, Source: "config/crd/inferenceservice.yaml"},
		},
	}

	sel := NewArchitectureSelector(t.TempDir())
	layer, _ := sel.Select(cpg, arch, nil, nil)

	if layer.Name != "architecture" {
		t.Errorf("layer name = %q, want %q", layer.Name, "architecture")
	}
	found := false
	for _, r := range layer.Resources {
		if r.Kind == "CustomResourceDefinition" && r.Name == "InferenceService" {
			found = true
			if r.APIGroup != "serving.kserve.io" {
				t.Errorf("APIGroup = %q", r.APIGroup)
			}
			if r.FieldCount != 142 {
				t.Errorf("FieldCount = %d", r.FieldCount)
			}
		}
	}
	if !found {
		t.Error("CRD not found in resources")
	}
}

func TestArchitectureSelector_IncludesControllerWatches(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component: "test",
		ControllerWatch: []extractor.ControllerWatch{
			{Type: "For", GVK: "apis/v1beta1/InferenceService", Controller: "ISReconciler", Source: "pkg/reconciler/setup.go"},
			{Type: "Owns", GVK: "apps/v1/Deployment", Controller: "ISReconciler", Source: "pkg/reconciler/setup.go"},
		},
	}

	sel := NewArchitectureSelector(t.TempDir())
	layer, _ := sel.Select(cpg, arch, nil, nil)

	if len(layer.Relationships) < 2 {
		t.Fatalf("expected at least 2 relationships, got %d", len(layer.Relationships))
	}
	foundFor := false
	foundOwns := false
	for _, r := range layer.Relationships {
		if r.Kind == "For" {
			foundFor = true
		}
		if r.Kind == "Owns" {
			foundOwns = true
		}
	}
	if !foundFor {
		t.Error("missing For relationship")
	}
	if !foundOwns {
		t.Error("missing Owns relationship")
	}
}

func TestArchitectureSelector_IncludesServices(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component: "test",
		Services: []extractor.Service{
			{Name: "kserve-webhook", Source: "config/default/service.yaml", TargetDeployment: "kserve-controller-manager", Ports: []extractor.ServicePort{{Name: "https", Port: 443}}},
		},
	}

	sel := NewArchitectureSelector(t.TempDir())
	layer, _ := sel.Select(cpg, arch, nil, nil)

	found := false
	for _, r := range layer.Resources {
		if r.Kind == "Service" && r.Name == "kserve-webhook" {
			found = true
			if r.Summary == "" {
				t.Error("service summary should include target and ports")
			}
		}
	}
	if !found {
		t.Error("Service not found")
	}
}

func TestArchitectureSelector_IncludesDeployments(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component: "test",
		Deployments: []extractor.Deployment{
			{Name: "controller-manager", Kind: "Deployment", Source: "config/manager/manager.yaml", ServiceAccount: "controller-sa",
				Containers: []extractor.Container{{Name: "manager"}, {Name: "kube-rbac-proxy"}}},
		},
	}

	sel := NewArchitectureSelector(t.TempDir())
	layer, _ := sel.Select(cpg, arch, nil, nil)

	found := false
	for _, r := range layer.Resources {
		if r.Kind == "Deployment" && r.Name == "controller-manager" {
			found = true
			if r.Summary == "" {
				t.Error("deployment summary should include containers and SA")
			}
		}
	}
	if !found {
		t.Error("Deployment not found")
	}
}

func TestArchitectureSelector_IncludesExternalConnections(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component: "test",
		ExternalConnections: []extractor.ExternalConnection{
			{Type: "database", Service: "postgres", Source: "pkg/db/connect.go:15", Function: "NewDBClient"},
		},
	}

	sel := NewArchitectureSelector(t.TempDir())
	layer, _ := sel.Select(cpg, arch, nil, nil)

	if len(layer.Relationships) == 0 {
		t.Fatal("expected external connection relationship")
	}
	if layer.Relationships[0].Kind != "external-database" {
		t.Errorf("kind = %q, want %q", layer.Relationships[0].Kind, "external-database")
	}
}

func TestArchitectureSelector_IncludesWebhooks(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component: "test",
		Webhooks: []extractor.WebhookConfig{
			{Name: "validate-inferenceservice", Type: "validating", Path: "/validate-inferenceservice",
				Sources: []extractor.SourceRef{{Type: "go_handler", File: "pkg/webhook/validate.go"}}},
		},
	}

	sel := NewArchitectureSelector(t.TempDir())
	layer, _ := sel.Select(cpg, arch, nil, nil)

	found := false
	for _, r := range layer.Resources {
		if r.Kind == "Webhook" && r.Name == "validate-inferenceservice" {
			found = true
		}
	}
	if !found {
		t.Error("Webhook not found")
	}
}

func TestArchitectureSelector_IncludesReconcileSequences(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component: "test",
		ReconcileSequences: []extractor.ReconcileSequence{
			{Controller: "ISReconciler", Source: "pkg/reconciler/reconciler.go",
				Steps: []extractor.ReconcileStep{
					{Order: 1, Method: "reconcileDeployment", Source: "pkg/reconciler/deployment.go"},
					{Order: 2, Method: "reconcileService", Source: "pkg/reconciler/service.go"},
				}},
		},
	}

	sel := NewArchitectureSelector(t.TempDir())
	layer, _ := sel.Select(cpg, arch, nil, nil)

	stepCount := 0
	for _, r := range layer.Relationships {
		if r.Kind == "reconcile-step" {
			stepCount++
		}
	}
	if stepCount != 2 {
		t.Errorf("expected 2 reconcile-step relationships, got %d", stepCount)
	}
}

func TestArchitectureSelector_IncludesReconcileFunctions(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "Reconcile",
		File: "controller.go", Line: 10, EndLine: 50,
		Language: "go", TypeName: "ISReconciler",
	})
	cpg.AddNode(&graph.Node{
		ID: "fn2", Kind: graph.NodeFunction, Name: "helperFunc",
		File: "utils.go", Line: 5, EndLine: 15,
		Language: "go",
	})

	dir := t.TempDir()
	writeTestFile(t, dir, "controller.go", 60)

	sel := NewArchitectureSelector(dir)
	layer, _ := sel.Select(cpg, nil, nil, nil)

	found := false
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "Reconcile" {
				found = true
			}
			if fn.Name == "helperFunc" {
				t.Error("helperFunc should not be selected (not a reconciler)")
			}
		}
	}
	if !found {
		t.Error("Reconcile function not found")
	}
}

func TestArchitectureSelector_IncludesArchitectureFindings(t *testing.T) {
	cpg := graph.NewCPG()
	findings := []query.Finding{
		{RuleID: "CGA-A01", Severity: "medium", Message: "abstraction layer", File: "pkg/store.go", Line: 10, Domain: "architecture"},
		{RuleID: "CGA-S01", Severity: "high", Message: "security finding", File: "pkg/auth.go", Line: 5, Domain: "security"},
	}

	sel := NewArchitectureSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, findings, nil)

	if len(layer.Findings) != 1 {
		t.Fatalf("expected 1 architecture finding (not security), got %d", len(layer.Findings))
	}
	if layer.Findings[0].Domain != "architecture" {
		t.Errorf("domain = %q, want %q", layer.Findings[0].Domain, "architecture")
	}
}

func TestArchitectureSelector_EmptyInputs(t *testing.T) {
	cpg := graph.NewCPG()
	sel := NewArchitectureSelector(t.TempDir())
	layer, warnings := sel.Select(cpg, nil, nil, nil)

	if layer.Name != "architecture" {
		t.Errorf("name = %q", layer.Name)
	}
	if len(warnings) != 0 {
		t.Errorf("expected 0 warnings, got %d", len(warnings))
	}
}

func TestArchitectureSelector_CompileIntegration(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "ReconcileDeployment",
		File: "controller.go", Line: 10, EndLine: 30,
		Language: "go",
	})

	arch := &extractor.ComponentArchitecture{
		Component: "test-operator",
		CRDs:      []extractor.CRD{{Kind: "Foo", Group: "test.io", Version: "v1"}},
		Services:  []extractor.Service{{Name: "foo-svc", Source: "svc.yaml"}},
	}

	dir := t.TempDir()
	writeTestFile(t, dir, "controller.go", 40)

	sel := NewArchitectureSelector(dir)
	layer, _ := sel.Select(cpg, arch, nil, nil)

	if len(layer.Resources) < 2 {
		t.Errorf("expected at least 2 resources (CRD + Service), got %d", len(layer.Resources))
	}
	if len(layer.Files) == 0 {
		t.Error("expected reconcile function in files")
	}
}

func TestArchitectureSelector_PathTraversalRejected(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "Reconcile",
		File: "../../../etc/passwd", Line: 1, EndLine: 5,
		Language: "go",
	})
	sel := NewArchitectureSelector(t.TempDir())
	layer, warnings := sel.Select(cpg, nil, nil, nil)
	if len(layer.Files) > 0 {
		for _, f := range layer.Files {
			if len(f.Functions) > 0 && f.Functions[0].Code != "" {
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
