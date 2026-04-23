package builder

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ugiordan/architecture-analyzer/pkg/graph"
)

func TestBuilderMultiLanguage(t *testing.T) {
	tmpDir := t.TempDir()

	fixtures := map[string]string{
		"simple_http_server.go": "../../testdata/simple_http_server.go",
		"flask_app.py":          "../../testdata/flask_app.py",
		"express_server.ts":     "../../testdata/express_server.ts",
		"actix_handler.rs":      "../../testdata/actix_handler.rs",
	}

	for dest, src := range fixtures {
		content, err := os.ReadFile(src)
		if err != nil {
			t.Fatalf("Failed to read %s: %v", src, err)
		}
		if err := os.WriteFile(filepath.Join(tmpDir, dest), content, 0644); err != nil {
			t.Fatalf("Failed to write %s: %v", dest, err)
		}
	}

	b := NewBuilder()
	cpg, err := b.BuildFromDir(tmpDir)
	if err != nil {
		t.Fatalf("BuildFromDir failed: %v", err)
	}

	// Check functions from all 4 languages
	languages := make(map[string]int)
	for _, fn := range cpg.NodesByKind(graph.NodeFunction) {
		languages[fn.Language]++
	}
	for _, lang := range []string{"go", "python", "typescript", "rust"} {
		if languages[lang] == 0 {
			t.Errorf("expected functions from language %q, got 0", lang)
		}
	}
	t.Logf("Functions by language: %v", languages)

	// Check HTTP handlers across languages
	handlers := cpg.NodesByKind(graph.NodeHTTPEndpoint)
	if len(handlers) < 5 {
		t.Errorf("expected at least 5 HTTP handlers across all languages, got %d", len(handlers))
	}
	handlerLangs := make(map[string]int)
	for _, h := range handlers {
		handlerLangs[h.Language]++
	}
	t.Logf("HTTP handlers by language: %v", handlerLangs)

	// Check DB operations
	dbOps := cpg.NodesByKind(graph.NodeDBOperation)
	if len(dbOps) < 2 {
		t.Errorf("expected at least 2 DB operations, got %d", len(dbOps))
	}

	// Check no node ID collisions
	ids := make(map[string]bool)
	for _, n := range cpg.Nodes() {
		if ids[n.ID] {
			t.Errorf("node ID collision: %s", n.ID)
		}
		ids[n.ID] = true
	}
	t.Logf("Total unique nodes: %d", len(ids))

	// Check containment edges exist
	containment := 0
	calls := 0
	for _, e := range cpg.Edges() {
		if e.Kind == graph.EdgeDataFlow && e.Label == "contains_call" {
			containment++
		}
		if e.Kind == graph.EdgeCalls {
			calls++
		}
	}
	if containment == 0 {
		t.Error("expected containment edges (EdgeDataFlow contains_call), got 0")
	}
	t.Logf("Containment edges: %d, Call edges: %d", containment, calls)
}
