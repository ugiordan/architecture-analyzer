package renderer

import (
	"strings"
	"testing"
)

func TestRenderComponentAgentContextPage(t *testing.T) {
	data := map[string]interface{}{
		"component":  "demo",
		"repo":       "example/demo",
		"commit_sha": "abc123",
		"agent_context": map[string]interface{}{
			"component":              "demo",
			"repo":                   "example/demo",
			"commit_sha":             "abc123",
			"graph_status":           "available",
			"summary":                "A demo component summary. It manages demo resources.",
			"repository_description": "Demo project overview",
			"description_source":     "README.md:3",
			"graph": map[string]interface{}{
				"source_files":       2,
				"functions":          3,
				"test_functions":     1,
				"resolved_calls":     1,
				"http_endpoints":     1,
				"storage_operations": 1,
				"external_calls":     0,
			},
			"languages": []interface{}{map[string]interface{}{
				"name": "go", "source_files": 2, "functions": 3, "test_functions": 1,
			}},
			"source_areas": []interface{}{map[string]interface{}{
				"path": "pkg", "source_files": 2, "functions": 3, "test_functions": 1,
				"languages": []interface{}{"go"},
				"symbols": []interface{}{map[string]interface{}{
					"name": "Serve", "source": "pkg/server.go:12", "callers": 1,
				}},
			}},
			"entry_points": []interface{}{map[string]interface{}{
				"kind": "http", "name": "GET /ready", "method": "GET", "path": "/ready",
				"trust_level": "untrusted", "source": "pkg/server.go:10",
			}},
			"call_relationships": []interface{}{map[string]interface{}{
				"caller": "main", "caller_source": "cmd/main.go:2", "callee": "Serve",
				"callee_source": "pkg/server.go:12", "call_site_source": "cmd/main.go:4", "confidence": "CERTAIN",
			}},
			"test_relationships": []interface{}{map[string]interface{}{
				"test": "TestServe", "test_source": "pkg/server_test.go:8", "target": "Serve",
				"target_source": "pkg/server.go:12", "confidence": "CERTAIN",
			}},
			"coverage": map[string]interface{}{"go_ast": "full"},
		},
	}

	content := renderComponentAgentContextPage(data, "demo")
	for _, want := range []string{
		"# demo: Agent Context",
		"[Extracted Data](details.md)",
		"Analyzer Architecture Summary",
		"- A demo component summary.",
		"- It manages demo resources.",
		"Untrusted Repository Text",
		"Source Areas (1 directories)",
		"<details markdown>",
		"`pkg/server.go:12`",
		"https://github.com/example/demo/blob/abc123/pkg/server.go#L12",
		"GET /ready",
		"Call Relationships (1)",
		"CERTAIN",
		"Tests and Targets (1)",
		"go_ast",
	} {
		if !strings.Contains(content, want) {
			t.Errorf("agent context page missing %q", want)
		}
	}
}

func TestRenderComponentDocsIncludesAgentContext(t *testing.T) {
	pages := renderComponentDocs(map[string]interface{}{"component": "demo"}, "components/demo")
	for _, page := range pages {
		if page.Path == "components/demo/agent-context.md" {
			if !strings.Contains(page.Content, "Agent Context Not Available") {
				t.Fatalf("missing-sidecar page should state context is unavailable, got:\n%s", page.Content)
			}
			return
		}
	}
	t.Fatal("agent-context page was not generated")
}
