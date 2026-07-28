package layers

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ugiordan/architecture-analyzer/pkg/graph"
	"github.com/ugiordan/architecture-analyzer/pkg/query"
)

func TestTestingSelector_SelectsTestFunctions(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "test1", Kind: graph.NodeFunction, Name: "TestReconcile",
		File: "controller_test.go", Line: 10, EndLine: 30,
		Language:    "go",
		Annotations: map[string]bool{annotIsTestFunc: true},
	})
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "Reconcile",
		File: "controller.go", Line: 5, EndLine: 40,
		Language: "go",
	})

	dir := t.TempDir()
	writeTestFile(t, dir, "controller_test.go", 40)
	writeTestFile(t, dir, "controller.go", 50)

	sel := NewTestingSelector(dir)
	layer, _ := sel.Select(cpg, nil, nil, nil)

	if layer.Name != "testing" {
		t.Errorf("layer name = %q, want %q", layer.Name, "testing")
	}

	found := false
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "TestReconcile" {
				found = true
				if fn.Code == "" {
					t.Error("test function should have code body")
				}
			}
		}
	}
	if !found {
		t.Error("TestReconcile not selected")
	}
}

func TestTestingSelector_SelectsTestHelpers(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "helper1", Kind: graph.NodeFunction, Name: "createTestNamespace",
		File: "helpers_test.go", Line: 5, EndLine: 20,
		Language:    "go",
		Annotations: map[string]bool{annotIsTestHelper: true},
	})

	dir := t.TempDir()
	writeTestFile(t, dir, "helpers_test.go", 25)

	sel := NewTestingSelector(dir)
	layer, _ := sel.Select(cpg, nil, nil, nil)

	found := false
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "createTestNamespace" {
				found = true
				hasMeta := false
				for _, m := range fn.Metas {
					if m.Key == "test-kind" && m.Value == "helper" {
						hasMeta = true
					}
				}
				if !hasMeta {
					t.Error("expected test-kind=helper meta")
				}
			}
		}
	}
	if !found {
		t.Error("test helper not selected")
	}
}

func TestTestingSelector_PairsTestWithTarget(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "test1", Kind: graph.NodeFunction, Name: "TestReconcile",
		File: "controller_test.go", Line: 10,
		Language:    "go",
		Annotations: map[string]bool{annotIsTestFunc: true},
	})
	cpg.AddNode(&graph.Node{
		ID: "cs1", Kind: graph.NodeCallSite, Name: "Reconcile",
	})
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "Reconcile",
		File: "controller.go", Line: 5, EndLine: 40,
		Language: "go",
	})

	cpg.AddEdge(&graph.Edge{From: "test1", To: "cs1", Kind: graph.EdgeDataFlow})
	cpg.AddEdge(&graph.Edge{From: "cs1", To: "fn1", Kind: graph.EdgeCalls})

	dir := t.TempDir()
	writeTestFile(t, dir, "controller.go", 50)

	sel := NewTestingSelector(dir)
	layer, _ := sel.Select(cpg, nil, nil, nil)

	foundTarget := false
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "Reconcile" {
				foundTarget = true
				hasTested := false
				for _, m := range fn.Metas {
					if m.Key == "tested" && m.Value == "true" {
						hasTested = true
					}
				}
				if !hasTested {
					t.Error("target function should have tested=true meta")
				}
			}
		}
	}
	if !foundTarget {
		t.Error("target function Reconcile not included")
	}

	if len(layer.Relationships) == 0 {
		t.Fatal("expected at least one 'tests' relationship")
	}
	rel := layer.Relationships[0]
	if rel.Kind != "tests" {
		t.Errorf("relationship kind = %q, want %q", rel.Kind, "tests")
	}
}

func TestTestingSelector_TestKindMetas(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "t1", Kind: graph.NodeFunction, Name: "TestUnit",
		File: "unit_test.go", Line: 1,
		Language:    "go",
		Annotations: map[string]bool{annotIsTestFunc: true},
	})
	cpg.AddNode(&graph.Node{
		ID: "t2", Kind: graph.NodeFunction, Name: "TestIntegration",
		File: "integration_test.go", Line: 1,
		Language:    "go",
		Annotations: map[string]bool{annotIsTestFunc: true, annotUsesEnvtest: true},
	})
	cpg.AddNode(&graph.Node{
		ID: "t3", Kind: graph.NodeFunction, Name: "BenchmarkReconcile",
		File: "bench_test.go", Line: 1,
		Language:    "go",
		Annotations: map[string]bool{annotIsTestFunc: true},
	})

	sel := NewTestingSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	metas := make(map[string]string)
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			for _, m := range fn.Metas {
				if m.Key == "test-kind" {
					metas[fn.Name] = m.Value
				}
			}
		}
	}

	if metas["TestUnit"] != "unit" {
		t.Errorf("TestUnit test-kind = %q, want %q", metas["TestUnit"], "unit")
	}
	if metas["TestIntegration"] != "integration" {
		t.Errorf("TestIntegration test-kind = %q, want %q", metas["TestIntegration"], "integration")
	}
	if metas["BenchmarkReconcile"] != "benchmark" {
		t.Errorf("BenchmarkReconcile test-kind = %q, want %q", metas["BenchmarkReconcile"], "benchmark")
	}
}

func TestTestingSelector_MockTargetMetas(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "t1", Kind: graph.NodeFunction, Name: "TestWithFake",
		File: "fake_test.go", Line: 1,
		Language:    "go",
		Annotations: map[string]bool{annotIsTestFunc: true, annotUsesFakeClient: true},
	})
	cpg.AddNode(&graph.Node{
		ID: "t2", Kind: graph.NodeFunction, Name: "TestWithEnvtest",
		File: "envtest_test.go", Line: 1,
		Language:    "go",
		Annotations: map[string]bool{annotIsTestFunc: true, annotUsesEnvtest: true},
	})

	sel := NewTestingSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	mockTargets := make(map[string][]string)
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			for _, m := range fn.Metas {
				if m.Key == "mock-target" {
					mockTargets[fn.Name] = append(mockTargets[fn.Name], m.Value)
				}
			}
		}
	}

	if len(mockTargets["TestWithFake"]) != 1 {
		t.Fatalf("TestWithFake expected 1 mock-target, got %d", len(mockTargets["TestWithFake"]))
	}
	if mockTargets["TestWithFake"][0] != "fake-client" {
		t.Errorf("TestWithFake mock-target = %q, want %q", mockTargets["TestWithFake"][0], "fake-client")
	}
	if len(mockTargets["TestWithEnvtest"]) != 1 {
		t.Fatalf("TestWithEnvtest expected 1 mock-target, got %d", len(mockTargets["TestWithEnvtest"]))
	}
	if mockTargets["TestWithEnvtest"][0] != "envtest" {
		t.Errorf("TestWithEnvtest mock-target = %q, want %q", mockTargets["TestWithEnvtest"][0], "envtest")
	}
}

func TestTestingSelector_FiltersTestingFindings(t *testing.T) {
	cpg := graph.NewCPG()
	findings := []query.Finding{
		{RuleID: "CGA-T01", Severity: "medium", Domain: "testing", Message: "Untested security func", File: "a.go", Line: 1},
		{RuleID: "CGA-T02", Severity: "low", Domain: "testing", Message: "Fake only integration", File: "b.go", Line: 2},
		{RuleID: "CGA-S01", Severity: "high", Domain: "security", Message: "Security issue", File: "c.go", Line: 3},
		{RuleID: "CGA-A01", Severity: "medium", Domain: "architecture", Message: "Arch issue", File: "d.go", Line: 4},
	}

	sel := NewTestingSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, findings, nil)

	if len(layer.Findings) != 2 {
		t.Fatalf("expected 2 testing findings, got %d", len(layer.Findings))
	}
	for _, f := range layer.Findings {
		if f.Domain != "testing" {
			t.Errorf("unexpected domain %q", f.Domain)
		}
	}
}

func TestTestingSelector_FindingsSorted(t *testing.T) {
	cpg := graph.NewCPG()
	findings := []query.Finding{
		{RuleID: "CGA-T04", Severity: "low", Domain: "testing", Message: "Consolidation", File: "a.go", Line: 1},
		{RuleID: "CGA-T01", Severity: "medium", Domain: "testing", Message: "Untested", File: "b.go", Line: 2},
		{RuleID: "CGA-T03", Severity: "medium", Domain: "testing", Message: "Missing error paths", File: "c.go", Line: 3},
	}

	sel := NewTestingSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, findings, nil)

	if len(layer.Findings) != 3 {
		t.Fatalf("expected 3 findings, got %d", len(layer.Findings))
	}
	if layer.Findings[0].Severity != "medium" {
		t.Errorf("first finding severity = %q, want medium", layer.Findings[0].Severity)
	}
	if layer.Findings[2].Severity != "low" {
		t.Errorf("last finding severity = %q, want low", layer.Findings[2].Severity)
	}
}

func TestTestingSelector_EmptyInputs(t *testing.T) {
	cpg := graph.NewCPG()
	sel := NewTestingSelector(t.TempDir())
	layer, warnings := sel.Select(cpg, nil, nil, nil)

	if layer.Name != "testing" {
		t.Errorf("layer name = %q, want %q", layer.Name, "testing")
	}
	if len(layer.Files) != 0 {
		t.Errorf("expected 0 files, got %d", len(layer.Files))
	}
	if len(layer.Findings) != 0 {
		t.Errorf("expected 0 findings, got %d", len(layer.Findings))
	}
	if len(layer.Relationships) != 0 {
		t.Errorf("expected 0 relationships, got %d", len(layer.Relationships))
	}
	if len(warnings) != 0 {
		t.Errorf("expected 0 warnings, got %d", len(warnings))
	}
}

func TestTestingSelector_PathTraversalRejected(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "t1", Kind: graph.NodeFunction, Name: "TestEvil",
		File: "../../../etc/passwd", Line: 1, EndLine: 5,
		Language:    "go",
		Annotations: map[string]bool{annotIsTestFunc: true},
	})

	sel := NewTestingSelector(t.TempDir())
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

func TestTestingSelector_TableDrivenMeta(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "t1", Kind: graph.NodeFunction, Name: "TestTableDriven",
		File: "table_test.go", Line: 1,
		Language:    "go",
		Annotations: map[string]bool{annotIsTestFunc: true, annotTableDriven: true},
	})

	sel := NewTestingSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	found := false
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "TestTableDriven" {
				for _, m := range fn.Metas {
					if m.Key == "table-driven" && m.Value == "true" {
						found = true
					}
				}
			}
		}
	}
	if !found {
		t.Error("expected table-driven=true meta")
	}
}

func TestTestingSelector_NoDuplicateTargets(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "test1", Kind: graph.NodeFunction, Name: "TestA",
		File: "a_test.go", Line: 1,
		Language:    "go",
		Annotations: map[string]bool{annotIsTestFunc: true},
	})
	cpg.AddNode(&graph.Node{
		ID: "test2", Kind: graph.NodeFunction, Name: "TestB",
		File: "b_test.go", Line: 1,
		Language:    "go",
		Annotations: map[string]bool{annotIsTestFunc: true},
	})
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "SharedFunc",
		File: "shared.go", Line: 5, EndLine: 20,
		Language: "go",
	})

	cpg.AddNode(&graph.Node{ID: "cs1", Kind: graph.NodeCallSite, Name: "SharedFunc"})
	cpg.AddNode(&graph.Node{ID: "cs2", Kind: graph.NodeCallSite, Name: "SharedFunc"})
	cpg.AddEdge(&graph.Edge{From: "test1", To: "cs1", Kind: graph.EdgeDataFlow})
	cpg.AddEdge(&graph.Edge{From: "cs1", To: "fn1", Kind: graph.EdgeCalls})
	cpg.AddEdge(&graph.Edge{From: "test2", To: "cs2", Kind: graph.EdgeDataFlow})
	cpg.AddEdge(&graph.Edge{From: "cs2", To: "fn1", Kind: graph.EdgeCalls})

	dir := t.TempDir()
	writeTestFile(t, dir, "shared.go", 25)

	sel := NewTestingSelector(dir)
	layer, _ := sel.Select(cpg, nil, nil, nil)

	count := 0
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "SharedFunc" {
				count++
			}
		}
	}
	if count != 1 {
		t.Errorf("SharedFunc should appear exactly once, got %d", count)
	}
}

func TestTestingSelector_ExcludesNonTestFunctions(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "RegularFunc",
		File: "regular.go", Line: 5,
		Language: "go",
	})
	cpg.AddNode(&graph.Node{
		ID: "fn2", Kind: graph.NodeFunction, Name: "AnotherFunc",
		File: "another.go", Line: 10,
		Language: "go",
	})

	sel := NewTestingSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	if len(layer.Files) != 0 {
		t.Errorf("expected no files (no test functions), got %d", len(layer.Files))
	}
}

func TestTestingSelector_ErrorPathMeta(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "t1", Kind: graph.NodeFunction, Name: "TestErrorCase",
		File: "error_test.go", Line: 1,
		Language:    "go",
		Annotations: map[string]bool{annotIsTestFunc: true, annotErrorPath: true},
	})

	sel := NewTestingSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	found := false
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "TestErrorCase" {
				for _, m := range fn.Metas {
					if m.Key == "error-path" && m.Value == "true" {
						found = true
					}
				}
			}
		}
	}
	if !found {
		t.Error("expected error-path=true meta")
	}
}

func TestTestingSelector_BenchmarkWithEnvtestPriority(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "t1", Kind: graph.NodeFunction, Name: "BenchmarkReconcile",
		File: "bench_test.go", Line: 1,
		Language:    "go",
		Annotations: map[string]bool{annotIsTestFunc: true, annotUsesEnvtest: true},
	})

	sel := NewTestingSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "BenchmarkReconcile" {
				for _, m := range fn.Metas {
					if m.Key == "test-kind" && m.Value != "integration" {
						t.Errorf("envtest should take priority over benchmark name: test-kind = %q, want integration", m.Value)
					}
				}
			}
		}
	}
}

func TestTestingSelector_RelationshipCap(t *testing.T) {
	cpg := graph.NewCPG()

	for i := 0; i < 200; i++ {
		testID := fmt.Sprintf("test%d", i)
		csID := fmt.Sprintf("cs%d", i)
		fnID := fmt.Sprintf("fn%d", i)
		testFile := fmt.Sprintf("test_%d_test.go", i)
		fnFile := fmt.Sprintf("impl_%d.go", i)

		cpg.AddNode(&graph.Node{
			ID: testID, Kind: graph.NodeFunction, Name: fmt.Sprintf("Test%d", i),
			File: testFile, Line: 1,
			Language:    "go",
			Annotations: map[string]bool{annotIsTestFunc: true},
		})
		cpg.AddNode(&graph.Node{ID: csID, Kind: graph.NodeCallSite, Name: fmt.Sprintf("Func%d", i)})
		cpg.AddNode(&graph.Node{
			ID: fnID, Kind: graph.NodeFunction, Name: fmt.Sprintf("Func%d", i),
			File: fnFile, Line: 1,
			Language: "go",
		})
		cpg.AddEdge(&graph.Edge{From: testID, To: csID, Kind: graph.EdgeDataFlow})
		cpg.AddEdge(&graph.Edge{From: csID, To: fnID, Kind: graph.EdgeCalls})
	}

	sel := NewTestingSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	if len(layer.Relationships) > maxRelationships {
		t.Errorf("relationships = %d, should be capped at %d", len(layer.Relationships), maxRelationships)
	}
}

func TestTestingSelector_SubtestsMeta(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "t1", Kind: graph.NodeFunction, Name: "TestWithSubtests",
		File: "sub_test.go", Line: 1,
		Language:    "go",
		Annotations: map[string]bool{annotIsTestFunc: true, annotSubtests: true},
	})

	sel := NewTestingSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	found := false
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "TestWithSubtests" {
				for _, m := range fn.Metas {
					if m.Key == "subtests" && m.Value == "true" {
						found = true
					}
				}
			}
		}
	}
	if !found {
		t.Error("expected subtests=true meta")
	}
}

func TestTestingSelector_NoTestKindDuplicateForHelperWithTestFunc(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "t1", Kind: graph.NodeFunction, Name: "TestHelper",
		File: "helper_test.go", Line: 1,
		Language:    "go",
		Annotations: map[string]bool{annotIsTestFunc: true, annotIsTestHelper: true},
	})

	sel := NewTestingSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	testKindCount := 0
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "TestHelper" {
				for _, m := range fn.Metas {
					if m.Key == "test-kind" {
						testKindCount++
					}
				}
			}
		}
	}
	if testKindCount != 1 {
		t.Errorf("expected exactly 1 test-kind meta, got %d", testKindCount)
	}
}

func TestTestingSelector_BodyExtractionFailure(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "t1", Kind: graph.NodeFunction, Name: "TestBadRange",
		File: "short.go", Line: 50, EndLine: 60,
		Language:    "go",
		Annotations: map[string]bool{annotIsTestFunc: true},
	})

	dir := t.TempDir()
	writeTestFile(t, dir, "short.go", 5)

	sel := NewTestingSelector(dir)
	layer, warnings := sel.Select(cpg, nil, nil, nil)

	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "TestBadRange" && fn.Code != "" {
				t.Error("should not have code body when lines exceed file length")
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

func TestTestingSelector_SharedFile(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "test1", Kind: graph.NodeFunction, Name: "TestHelper",
		File: "shared_test.go", Line: 20, EndLine: 30,
		Language:    "go",
		Annotations: map[string]bool{annotIsTestFunc: true},
	})
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "helperImpl",
		File: "shared_test.go", Line: 5, EndLine: 15,
		Language: "go",
	})
	cpg.AddNode(&graph.Node{ID: "cs1", Kind: graph.NodeCallSite, Name: "helperImpl"})
	cpg.AddEdge(&graph.Edge{From: "test1", To: "cs1", Kind: graph.EdgeDataFlow})
	cpg.AddEdge(&graph.Edge{From: "cs1", To: "fn1", Kind: graph.EdgeCalls})

	dir := t.TempDir()
	writeTestFile(t, dir, "shared_test.go", 35)

	sel := NewTestingSelector(dir)
	layer, _ := sel.Select(cpg, nil, nil, nil)

	fileCount := 0
	for _, f := range layer.Files {
		if f.Path == "shared_test.go" {
			fileCount++
			if len(f.Functions) != 2 {
				t.Errorf("expected 2 functions in shared file, got %d", len(f.Functions))
			}
		}
	}
	if fileCount != 1 {
		t.Errorf("expected 1 file entry for shared_test.go, got %d", fileCount)
	}
}

func TestTestingSelector_HasTargetMetaOnTestFunc(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "test1", Kind: graph.NodeFunction, Name: "TestReconcile",
		File: "controller_test.go", Line: 10,
		Language:    "go",
		Annotations: map[string]bool{annotIsTestFunc: true},
	})
	cpg.AddNode(&graph.Node{ID: "cs1", Kind: graph.NodeCallSite, Name: "Reconcile"})
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "Reconcile",
		File: "controller.go", Line: 5,
		Language: "go",
	})
	cpg.AddEdge(&graph.Edge{From: "test1", To: "cs1", Kind: graph.EdgeDataFlow})
	cpg.AddEdge(&graph.Edge{From: "cs1", To: "fn1", Kind: graph.EdgeCalls})

	sel := NewTestingSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "TestReconcile" {
				hasTarget := false
				for _, m := range fn.Metas {
					if m.Key == "has-target" && m.Value == "true" {
						hasTarget = true
					}
					if m.Key == "tested" {
						t.Error("test function should use 'has-target' not 'tested'")
					}
				}
				if !hasTarget {
					t.Error("expected has-target=true meta on test function")
				}
			}
		}
	}
}
