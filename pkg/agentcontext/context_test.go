package agentcontext

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	testingdomain "github.com/ugiordan/architecture-analyzer/pkg/domains/testing"
	"github.com/ugiordan/architecture-analyzer/pkg/extractor"
	"github.com/ugiordan/architecture-analyzer/pkg/graph"
)

func TestBuildIncludesRepositoryAndGraphContext(t *testing.T) {
	repoPath := t.TempDir()
	readme := "# Demo\n\nDemo repo overview with [docs](https://example.test).\n\n## More\n"
	if err := os.WriteFile(filepath.Join(repoPath, "README.md"), []byte(readme), 0o644); err != nil {
		t.Fatal(err)
	}

	cpg := graph.NewCPG()
	nodes := []*graph.Node{
		{ID: "main", Kind: graph.NodeFunction, Name: "main", File: "cmd/main.go", Line: 2, Language: "go"},
		{ID: "main-call", Kind: graph.NodeCallSite, Name: "serve", File: "cmd/main.go", Line: 4, Language: "go"},
		{ID: "test", Kind: graph.NodeFunction, Name: "TestServe", File: "pkg/serve_test.go", Line: 8, Language: "go", Annotations: map[string]bool{testingdomain.AnnotIsTestFunc: true}},
		{ID: "test-call", Kind: graph.NodeCallSite, Name: "serve", File: "pkg/serve_test.go", Line: 9, Language: "go"},
		{ID: "serve", Kind: graph.NodeFunction, Name: "serve", File: "pkg/serve.go", Line: 12, Language: "go", Complexity: 3},
		{ID: "endpoint", Kind: graph.NodeHTTPEndpoint, Name: "GET /ready", File: "pkg/serve.go", Line: 10, Language: "go", HTTPMethod: "GET", Route: "/ready", TrustLevel: graph.TrustUntrusted},
		{ID: "db", Kind: graph.NodeDBOperation, Name: "records", File: "pkg/store.go", Line: 14, Language: "go", Operation: "read", Table: "records"},
		{ID: "external", Kind: graph.NodeExternalCall, Name: "client.Do", File: "pkg/client.go", Line: 21, Language: "go", TypeName: "http"},
	}
	for _, node := range nodes {
		if err := cpg.AddNode(node); err != nil {
			t.Fatal(err)
		}
	}
	cpg.AddEdge(&graph.Edge{From: "main", To: "main-call", Kind: graph.EdgeDataFlow, Label: "contains_call"})
	cpg.AddEdge(&graph.Edge{From: "main-call", To: "serve", Kind: graph.EdgeCalls, Confidence: graph.ConfidenceCertain})
	cpg.AddEdge(&graph.Edge{From: "test", To: "test-call", Kind: graph.EdgeDataFlow, Label: "contains_call"})
	cpg.AddEdge(&graph.Edge{From: "test-call", To: "serve", Kind: graph.EdgeCalls, Confidence: graph.ConfidenceInferred})
	cpg.AddEdge(&graph.Edge{From: "endpoint", To: "external", Kind: graph.EdgeTaint, Label: "request->http", Confidence: graph.ConfidenceCertain})

	arch := &extractor.ComponentArchitecture{
		Component:       "demo",
		Repo:            "example/demo",
		CommitSHA:       "abc123",
		AnalyzerVersion: "test",
		Summary:         "demo component summary",
		DataCoverage:    map[string]string{"status": "full"},
		ReconcileSequences: []extractor.ReconcileSequence{{
			Controller: "DemoReconciler",
			Source:     "controllers/demo.go:4",
			Steps: []extractor.ReconcileStep{{
				Order: 1, Method: "reconcileService", Source: "controllers/service.go:10",
			}},
		}},
		ControllerWatch: []extractor.ControllerWatch{{
			Type: "For", GVK: "demo.io/v1/Demo", Controller: "DemoReconciler", Source: "controllers/setup.go:8",
		}},
	}

	got := Build(repoPath, arch, cpg)
	if got.GraphStatus != "available" || got.Graph.SourceFiles != 5 || got.Graph.Functions != 3 || got.Graph.TestFunctions != 1 || got.Graph.ResolvedCalls != 2 {
		t.Fatalf("unexpected graph summary: %+v", got.Graph)
	}
	if got.RepositoryDescription != "Demo repo overview with docs." || got.DescriptionSource != "README.md:3" {
		t.Fatalf("unexpected README description: %q from %q", got.RepositoryDescription, got.DescriptionSource)
	}
	hasHTTPEntrypoint := false
	for _, entry := range got.EntryPoints {
		if entry.Kind == "http" && entry.TrustLevel == string(graph.TrustUntrusted) {
			hasHTTPEntrypoint = true
		}
	}
	if len(got.EntryPoints) != 2 || !hasHTTPEntrypoint {
		t.Fatalf("unexpected entry points: %+v", got.EntryPoints)
	}
	if len(got.Controllers) != 1 || len(got.Controllers[0].Steps) != 1 || len(got.Controllers[0].Watches) != 1 {
		t.Fatalf("unexpected controller context: %+v", got.Controllers)
	}
	if len(got.CallRelationships) != 2 || got.CallRelationships[0].Caller != "main" {
		t.Fatalf("unexpected call relationships: %+v", got.CallRelationships)
	}
	if len(got.TestRelationships) != 1 || got.TestRelationships[0].Target != "serve" || got.TestRelationships[0].Confidence != string(graph.ConfidenceInferred) {
		t.Fatalf("unexpected test relationships: %+v", got.TestRelationships)
	}
	if len(got.Operations) != 2 || len(got.TaintFlows) != 1 || got.TaintFlows[0].Label != "request->http" {
		t.Fatalf("unexpected operation or taint context: %+v %+v", got.Operations, got.TaintFlows)
	}
	if len(got.SourceAreas) == 0 || got.SourceAreas[0].Path != "pkg" {
		t.Fatalf("unexpected source areas: %+v", got.SourceAreas)
	}

	if again := Build(repoPath, arch, cpg); !reflect.DeepEqual(got, again) {
		t.Fatal("context output was not deterministic")
	}
}

func TestBuildWithoutGraphRetainsArchitectureContext(t *testing.T) {
	arch := &extractor.ComponentArchitecture{
		Component: "demo",
		Summary:   "architecture-only",
		ControllerWatch: []extractor.ControllerWatch{{
			Type: "For", GVK: "demo.io/v1/Demo", Controller: "DemoReconciler", Source: "controllers/setup.go:8",
		}},
	}
	got := Build(t.TempDir(), arch, nil)
	if got.GraphStatus != "unavailable" || got.Component != "demo" || got.Summary != "architecture-only" {
		t.Fatalf("unexpected partial context: %+v", got)
	}
	if len(got.Controllers) != 1 || got.Controllers[0].Name != "DemoReconciler" {
		t.Fatalf("architecture-only context should retain controllers, got: %+v", got.Controllers)
	}
}

func TestReadmeDescriptionSkipsSetextHeading(t *testing.T) {
	repo := t.TempDir()
	readme := "Demo\n====\n\nRuns the demo service.\n"
	if err := os.WriteFile(filepath.Join(repo, "README.md"), []byte(readme), 0o644); err != nil {
		t.Fatal(err)
	}
	description, source, truncated := readmeDescription(repo)
	if description != "Runs the demo service." || source != "README.md:4" || truncated {
		t.Fatalf("setext heading should be skipped, got %q from %q (truncated=%v)", description, source, truncated)
	}
}

func TestReadmeDescriptionRejectsSymlinkEscape(t *testing.T) {
	parent := t.TempDir()
	repo := filepath.Join(parent, "repo")
	if err := os.MkdirAll(repo, 0o755); err != nil {
		t.Fatal(err)
	}
	outside := filepath.Join(parent, "outside-secret.txt")
	if err := os.WriteFile(outside, []byte("# Stolen\n\nsecret contents outside the repo\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(outside, filepath.Join(repo, "README.md")); err != nil {
		t.Fatal(err)
	}

	description, source, truncated := readmeDescription(repo)
	if description != "" || source != "" || truncated {
		t.Fatalf("escaping README symlink must be rejected, got %q from %q (truncated=%v)", description, source, truncated)
	}

	ctx := Build(repo, nil, nil)
	if ctx.RepositoryDescription != "" {
		t.Fatalf("escaping README content must not reach the context, got %q", ctx.RepositoryDescription)
	}
}
