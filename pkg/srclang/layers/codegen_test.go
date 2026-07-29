package layers

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ugiordan/architecture-analyzer/pkg/extractor"
	"github.com/ugiordan/architecture-analyzer/pkg/graph"
	"github.com/ugiordan/architecture-analyzer/pkg/query"
)

func TestCodegenSelector_SelectsAllFunctions(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "Reconcile",
		File: "controller.go", Line: 10, EndLine: 50,
		Language: "go", TypeName: "Reconciler",
	})
	cpg.AddNode(&graph.Node{
		ID: "fn2", Kind: graph.NodeFunction, Name: "helper",
		File: "utils.go", Line: 5, EndLine: 15,
		Language: "go",
	})

	dir := t.TempDir()
	writeTestFile(t, dir, "controller.go", 60)
	writeTestFile(t, dir, "utils.go", 20)

	sel := NewCodegenSelector(dir)
	layer, _ := sel.Select(cpg, nil, nil, nil)

	if layer.Name != "codegen" {
		t.Errorf("layer name = %q, want codegen", layer.Name)
	}

	names := make(map[string]bool)
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			names[fn.Name] = true
		}
	}

	if !names["Reconcile"] {
		t.Error("Reconcile not selected")
	}
	if !names["helper"] {
		t.Error("helper not selected (codegen selects ALL functions)")
	}
}

func TestCodegenSelector_CodeRolePublicAPI(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "ExportedFunc",
		File: "api.go", Line: 5, Language: "go",
	})

	sel := NewCodegenSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "ExportedFunc" {
				for _, m := range fn.Metas {
					if m.Key == "code-role" && m.Value != "public-api" {
						t.Errorf("ExportedFunc code-role = %q, want public-api", m.Value)
					}
				}
			}
		}
	}
}

func TestCodegenSelector_CodeRoleInternal(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "privateFunc",
		File: "internal.go", Line: 5, Language: "go",
	})

	sel := NewCodegenSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "privateFunc" {
				for _, m := range fn.Metas {
					if m.Key == "code-role" && m.Value != "internal" {
						t.Errorf("privateFunc code-role = %q, want internal", m.Value)
					}
				}
			}
		}
	}
}

func TestCodegenSelector_CodeRoleTestOnly(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "TestSomething",
		File: "controller_test.go", Line: 5, Language: "go", IsTest: true,
	})

	sel := NewCodegenSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	foundRole := ""
	foundTestOnly := false
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "TestSomething" {
				for _, m := range fn.Metas {
					if m.Key == "code-role" {
						foundRole = m.Value
					}
					if m.Key == "test-only" && m.Value == "true" {
						foundTestOnly = true
					}
				}
			}
		}
	}
	if foundRole != "test-only" {
		t.Errorf("code-role = %q, want test-only", foundRole)
	}
	if !foundTestOnly {
		t.Error("expected test-only=true meta")
	}
}

func TestCodegenSelector_CodeRoleGenerated(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "DeepCopyInto",
		File: "zz_generated_deepcopy.go", Line: 10, Language: "go",
	})

	sel := NewCodegenSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	foundRole := ""
	foundGenerated := false
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "DeepCopyInto" {
				for _, m := range fn.Metas {
					if m.Key == "code-role" {
						foundRole = m.Value
					}
					if m.Key == "generated" && m.Value == "true" {
						foundGenerated = true
					}
				}
			}
		}
	}
	if foundRole != "generated" {
		t.Errorf("code-role = %q, want generated", foundRole)
	}
	if !foundGenerated {
		t.Error("expected generated=true meta")
	}
}

func TestCodegenSelector_ChangeRiskMetas(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "SimpleFunc",
		File: "simple.go", Line: 1, Language: "go", Complexity: 3,
	})
	cpg.AddNode(&graph.Node{
		ID: "fn2", Kind: graph.NodeFunction, Name: "MediumFunc",
		File: "medium.go", Line: 1, Language: "go", Complexity: 8,
	})
	cpg.AddNode(&graph.Node{
		ID: "fn3", Kind: graph.NodeFunction, Name: "ComplexFunc",
		File: "complex.go", Line: 1, Language: "go", Complexity: 15,
	})

	sel := NewCodegenSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	risks := make(map[string]string)
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			for _, m := range fn.Metas {
				if m.Key == "change-risk" {
					risks[fn.Name] = m.Value
				}
			}
		}
	}

	if _, ok := risks["SimpleFunc"]; ok {
		t.Error("SimpleFunc (complexity=3) should not have change-risk meta")
	}
	if risks["MediumFunc"] != "medium" {
		t.Errorf("MediumFunc change-risk = %q, want medium", risks["MediumFunc"])
	}
	if risks["ComplexFunc"] != "high" {
		t.Errorf("ComplexFunc change-risk = %q, want high", risks["ComplexFunc"])
	}
}

func TestCodegenSelector_IncludesAllFindings(t *testing.T) {
	cpg := graph.NewCPG()
	findings := []query.Finding{
		{RuleID: "CGA-S01", Severity: "high", Domain: "security", Message: "Security", File: "a.go", Line: 1},
		{RuleID: "CGA-T01", Severity: "medium", Domain: "testing", Message: "Testing", File: "b.go", Line: 2},
		{RuleID: "CGA-U01", Severity: "medium", Domain: "upgrade", Message: "Upgrade", File: "c.go", Line: 3},
		{RuleID: "CGA-N01", Severity: "high", Domain: "netpolicy", Message: "Netpolicy", File: "d.go", Line: 4},
	}

	sel := NewCodegenSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, findings, nil)

	if len(layer.Findings) != 4 {
		t.Fatalf("expected 4 findings (all domains), got %d", len(layer.Findings))
	}

	domains := make(map[string]bool)
	for _, f := range layer.Findings {
		domains[f.Domain] = true
	}
	for _, d := range []string{"security", "testing", "upgrade", "netpolicy"} {
		if !domains[d] {
			t.Errorf("missing domain %q", d)
		}
	}
}

func TestCodegenSelector_IncludesSecurityAnnotations(t *testing.T) {
	cpg := graph.NewCPG()
	annotations := []extractor.SecurityAnnotation{
		{Type: "RBAC_CLUSTER_SCOPE", Severity: "high", Source: "rbac.yaml", Description: "Wide RBAC"},
	}

	sel := NewCodegenSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, annotations)

	if len(layer.Findings) != 1 {
		t.Fatalf("expected 1 extraction finding, got %d", len(layer.Findings))
	}
	if layer.Findings[0].Domain != "extraction" {
		t.Errorf("domain = %q, want extraction", layer.Findings[0].Domain)
	}
}

func TestCodegenSelector_GoModuleImports(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component: "test",
		Dependencies: &extractor.DependencyData{
			GoModules: []extractor.GoModule{
				{Module: "sigs.k8s.io/controller-runtime", Version: "v0.17.0"},
				{Module: "k8s.io/api", Version: "v0.29.0"},
			},
		},
	}

	sel := NewCodegenSelector(t.TempDir())
	layer, _ := sel.Select(cpg, arch, nil, nil)

	if len(layer.Imports) != 2 {
		t.Fatalf("expected 2 imports, got %d", len(layer.Imports))
	}
	if layer.Imports[0].Kind != "go-module" {
		t.Errorf("import kind = %q, want go-module", layer.Imports[0].Kind)
	}
}

func TestCodegenSelector_PythonPackageImports(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component: "test",
		Dependencies: &extractor.DependencyData{
			PythonPackages: []extractor.PythonPackage{
				{Name: "kfp", Version: "2.5.0", Source: "requirements.txt"},
			},
		},
	}

	sel := NewCodegenSelector(t.TempDir())
	layer, _ := sel.Select(cpg, arch, nil, nil)

	if len(layer.Imports) != 1 {
		t.Fatalf("expected 1 import, got %d", len(layer.Imports))
	}
	if layer.Imports[0].Kind != "python-package" {
		t.Errorf("kind = %q, want python-package", layer.Imports[0].Kind)
	}
	if layer.Imports[0].Path != "requirements.txt" {
		t.Errorf("path = %q, want requirements.txt", layer.Imports[0].Path)
	}
}

func TestCodegenSelector_NilDependencies(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{Component: "test"}

	sel := NewCodegenSelector(t.TempDir())
	layer, _ := sel.Select(cpg, arch, nil, nil)

	if len(layer.Imports) != 0 {
		t.Errorf("expected 0 imports with nil dependencies, got %d", len(layer.Imports))
	}
}

func TestCodegenSelector_HTTPEndpoints(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "ep1", Kind: graph.NodeHTTPEndpoint, Name: "/api/v1/models",
		File: "routes.go", Line: 100,
	})

	sel := NewCodegenSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	found := false
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "/api/v1/models" && fn.Kind == "handler" {
				found = true
			}
		}
	}
	if !found {
		t.Error("expected HTTP endpoint in codegen layer")
	}
}

func TestCodegenSelector_TaintAnnotations(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "HandleInput",
		File: "handler.go", Line: 10,
		Language:    "go",
		Annotations: map[string]bool{"taint_source": true},
	})
	cpg.AddNode(&graph.Node{
		ID: "fn2", Kind: graph.NodeFunction, Name: "RunQuery",
		File: "db.go", Line: 5,
		Language:    "go",
		Annotations: map[string]bool{"taint_sink": true},
	})

	sel := NewCodegenSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "HandleInput" {
				if fn.TaintRole != "source" {
					t.Errorf("HandleInput TaintRole = %q, want source", fn.TaintRole)
				}
				if fn.Trust != "untrusted" {
					t.Errorf("HandleInput Trust = %q, want untrusted", fn.Trust)
				}
			}
			if fn.Name == "RunQuery" {
				if fn.TaintRole != "sink" {
					t.Errorf("RunQuery TaintRole = %q, want sink", fn.TaintRole)
				}
			}
		}
	}
}

func TestCodegenSelector_CallRelationships(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "Caller",
		File: "caller.go", Line: 10, Language: "go",
	})
	cpg.AddNode(&graph.Node{
		ID: "fn2", Kind: graph.NodeFunction, Name: "Callee",
		File: "callee.go", Line: 5, Language: "go",
	})
	cpg.AddEdge(&graph.Edge{From: "fn1", To: "fn2", Kind: graph.EdgeCalls})

	sel := NewCodegenSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	if len(layer.Relationships) == 0 {
		t.Fatal("expected at least one relationship")
	}
	if layer.Relationships[0].Kind != "calls" {
		t.Errorf("kind = %q, want calls", layer.Relationships[0].Kind)
	}
}

func TestCodegenSelector_EmptyInputs(t *testing.T) {
	cpg := graph.NewCPG()
	sel := NewCodegenSelector(t.TempDir())
	layer, warnings := sel.Select(cpg, nil, nil, nil)

	if layer.Name != "codegen" {
		t.Errorf("layer name = %q", layer.Name)
	}
	if len(layer.Files) != 0 {
		t.Errorf("expected 0 files, got %d", len(layer.Files))
	}
	if len(layer.Findings) != 0 {
		t.Errorf("expected 0 findings, got %d", len(layer.Findings))
	}
	if len(layer.Imports) != 0 {
		t.Errorf("expected 0 imports, got %d", len(layer.Imports))
	}
	if len(layer.Relationships) != 0 {
		t.Errorf("expected 0 relationships, got %d", len(layer.Relationships))
	}
	if len(warnings) != 0 {
		t.Errorf("expected 0 warnings, got %d", len(warnings))
	}
}

func TestCodegenSelector_PathTraversalRejected(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "Evil",
		File: "../../../etc/passwd", Line: 1, EndLine: 5,
		Language: "go",
	})

	sel := NewCodegenSelector(t.TempDir())
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

func TestCodegenSelector_BodyExtractionFailure(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "BadFunc",
		File: "short.go", Line: 50, EndLine: 60,
		Language: "go",
	})

	dir := t.TempDir()
	writeTestFile(t, dir, "short.go", 5)

	sel := NewCodegenSelector(dir)
	layer, warnings := sel.Select(cpg, nil, nil, nil)

	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "BadFunc" && fn.Code != "" {
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

func TestCodegenSelector_RelationshipCap(t *testing.T) {
	cpg := graph.NewCPG()

	for i := 0; i < 200; i++ {
		cpg.AddNode(&graph.Node{
			ID: fmt.Sprintf("fn%d", i), Kind: graph.NodeFunction,
			Name: fmt.Sprintf("Func%d", i),
			File: fmt.Sprintf("f_%d.go", i), Line: 1, Language: "go",
		})
		cpg.AddNode(&graph.Node{
			ID: fmt.Sprintf("helper%d", i), Kind: graph.NodeFunction,
			Name: fmt.Sprintf("helper%d", i),
			File: fmt.Sprintf("h_%d.go", i), Line: 1, Language: "go",
		})
		cpg.AddEdge(&graph.Edge{From: fmt.Sprintf("fn%d", i), To: fmt.Sprintf("helper%d", i), Kind: graph.EdgeCalls})
	}

	sel := NewCodegenSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	if len(layer.Relationships) > maxRelationships {
		t.Errorf("relationships = %d, should be capped at %d", len(layer.Relationships), maxRelationships)
	}
}

func TestCodegenSelector_TestFileWithoutIsTestFlag(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "setup_fixtures",
		File: "test_utils.py", Line: 5, Language: "python",
	})

	sel := NewCodegenSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	foundRole := ""
	foundTestOnly := false
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "setup_fixtures" {
				for _, m := range fn.Metas {
					if m.Key == "code-role" {
						foundRole = m.Value
					}
					if m.Key == "test-only" && m.Value == "true" {
						foundTestOnly = true
					}
				}
			}
		}
	}
	if foundRole != "test-only" {
		t.Errorf("code-role = %q, want test-only", foundRole)
	}
	if !foundTestOnly {
		t.Error("expected test-only=true meta even when IsTest is false (file prefix match)")
	}
}

func TestCodegenSelector_GeneratedFilePatterns(t *testing.T) {
	patterns := []struct {
		file string
		want string
	}{
		{"zz_generated_deepcopy.go", "generated"},
		{"types_generated.go", "generated"},
		{"api/v1/deepcopy.go", "generated"},
		{"pkg/client/generated/clientset.go", "generated"},
		{"pkg/api/types_gen.go", "generated"},
		{"pkg/controller/reconciler.go", "public-api"},
		{"pkg/attestation/attest_config.go", "public-api"},
		{"test_utils.py", "test-only"},
	}

	for _, tt := range patterns {
		cpg := graph.NewCPG()
		cpg.AddNode(&graph.Node{
			ID: "fn1", Kind: graph.NodeFunction, Name: "SomeFunc",
			File: tt.file, Line: 1, Language: "go",
		})

		sel := NewCodegenSelector(t.TempDir())
		layer, _ := sel.Select(cpg, nil, nil, nil)

		for _, f := range layer.Files {
			for _, fn := range f.Functions {
				for _, m := range fn.Metas {
					if m.Key == "code-role" && m.Value != tt.want {
						t.Errorf("file %q: code-role = %q, want %q", tt.file, m.Value, tt.want)
					}
				}
			}
		}
	}
}
