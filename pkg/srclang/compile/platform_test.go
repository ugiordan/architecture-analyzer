package compile

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractPlatform_BasicEdges(t *testing.T) {
	platformJSON := `{
		"platform": "RHOAI",
		"component_count": 38,
		"components": ["kserve", "odh-model-controller", "knative-serving"],
		"dependency_graph": [
			{"from": "odh-model-controller", "to": "kserve", "type": "go-module"},
			{"from": "kserve", "to": "knative-serving", "type": "go-module"},
			{"from": "kserve", "to": "kserve-serving-runtime", "type": "watches-crd:InferenceService"}
		],
		"component_data": []
	}`

	path := writeTempJSON(t, platformJSON)
	p, err := extractPlatform(path, "kserve")
	if err != nil {
		t.Fatalf("extractPlatform() error: %v", err)
	}

	if p.Name != "RHOAI" {
		t.Errorf("name = %q, want %q", p.Name, "RHOAI")
	}
	if p.Components != 38 {
		t.Errorf("components = %d, want %d", p.Components, 38)
	}
	if len(p.Inbound) != 1 {
		t.Fatalf("expected 1 inbound edge, got %d", len(p.Inbound))
	}
	if p.Inbound[0].Peer != "odh-model-controller" {
		t.Errorf("inbound[0].Peer = %q, want %q", p.Inbound[0].Peer, "odh-model-controller")
	}
	if p.Inbound[0].Type != "go-module" {
		t.Errorf("inbound[0].Type = %q, want %q", p.Inbound[0].Type, "go-module")
	}
	if len(p.Outbound) != 2 {
		t.Fatalf("expected 2 outbound edges, got %d", len(p.Outbound))
	}
	if p.Outbound[0].Peer != "knative-serving" {
		t.Errorf("outbound[0].Peer = %q, want %q", p.Outbound[0].Peer, "knative-serving")
	}
}

func TestExtractPlatform_EdgeTypeSplitting(t *testing.T) {
	platformJSON := `{
		"platform": "TestPlatform",
		"component_count": 5,
		"components": ["comp-a", "comp-b"],
		"dependency_graph": [
			{"from": "comp-a", "to": "comp-b", "type": "watches-crd:InferenceService"},
			{"from": "comp-b", "to": "comp-a", "type": "code-ref"}
		],
		"component_data": []
	}`

	path := writeTempJSON(t, platformJSON)
	p, err := extractPlatform(path, "comp-a")
	if err != nil {
		t.Fatalf("extractPlatform() error: %v", err)
	}

	if len(p.Outbound) != 1 {
		t.Fatalf("expected 1 outbound, got %d", len(p.Outbound))
	}
	if p.Outbound[0].Type != "watches-crd" {
		t.Errorf("outbound type = %q, want %q", p.Outbound[0].Type, "watches-crd")
	}
	if p.Outbound[0].Target != "InferenceService" {
		t.Errorf("outbound target = %q, want %q", p.Outbound[0].Target, "InferenceService")
	}

	if len(p.Inbound) != 1 {
		t.Fatalf("expected 1 inbound, got %d", len(p.Inbound))
	}
	if p.Inbound[0].Type != "code-ref" {
		t.Errorf("inbound type = %q, want %q", p.Inbound[0].Type, "code-ref")
	}
	if p.Inbound[0].Target != "" {
		t.Errorf("inbound target = %q, want empty", p.Inbound[0].Target)
	}
}

func TestExtractPlatform_NoEdgesForComponent(t *testing.T) {
	platformJSON := `{
		"platform": "TestPlatform",
		"component_count": 3,
		"components": ["comp-a", "comp-b", "comp-c"],
		"dependency_graph": [
			{"from": "comp-a", "to": "comp-b", "type": "go-module"}
		],
		"component_data": []
	}`

	path := writeTempJSON(t, platformJSON)
	p, err := extractPlatform(path, "comp-c")
	if err != nil {
		t.Fatalf("extractPlatform() error: %v", err)
	}

	if len(p.Inbound) != 0 {
		t.Errorf("expected 0 inbound edges, got %d", len(p.Inbound))
	}
	if len(p.Outbound) != 0 {
		t.Errorf("expected 0 outbound edges, got %d", len(p.Outbound))
	}
}

func TestExtractPlatform_SelfEdgesExcluded(t *testing.T) {
	platformJSON := `{
		"platform": "TestPlatform",
		"component_count": 2,
		"components": ["comp-a"],
		"dependency_graph": [
			{"from": "comp-a", "to": "comp-a", "type": "code-ref"}
		],
		"component_data": []
	}`

	path := writeTempJSON(t, platformJSON)
	p, err := extractPlatform(path, "comp-a")
	if err != nil {
		t.Fatalf("extractPlatform() error: %v", err)
	}

	if len(p.Inbound) != 0 {
		t.Errorf("self-edges should not appear in inbound, got %d", len(p.Inbound))
	}
	if len(p.Outbound) != 0 {
		t.Errorf("self-edges should not appear in outbound, got %d", len(p.Outbound))
	}
}

func TestExtractPlatform_FileNotFound(t *testing.T) {
	_, err := extractPlatform("/nonexistent/platform.json", "test")
	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestExtractPlatform_InvalidJSON(t *testing.T) {
	path := writeTempJSON(t, `{invalid json`)
	_, err := extractPlatform(path, "test")
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}

func TestSplitEdgeType(t *testing.T) {
	tests := []struct {
		raw        string
		wantType   string
		wantTarget string
	}{
		{"go-module", "go-module", ""},
		{"watches-crd:InferenceService", "watches-crd", "InferenceService"},
		{"code-ref", "code-ref", ""},
		{"watches-crd:Local:Model", "watches-crd", "Local:Model"},
	}

	for _, tt := range tests {
		typ, target := splitEdgeType(tt.raw)
		if typ != tt.wantType {
			t.Errorf("splitEdgeType(%q) type = %q, want %q", tt.raw, typ, tt.wantType)
		}
		if target != tt.wantTarget {
			t.Errorf("splitEdgeType(%q) target = %q, want %q", tt.raw, target, tt.wantTarget)
		}
	}
}

func TestExtractPlatform_TopologyFromComponentData(t *testing.T) {
	platformJSON := `{
		"platform": "RHOAI",
		"component_count": 3,
		"components": ["kserve", "odh-model-controller", "notebooks"],
		"dependency_graph": [],
		"component_data": [
			{"component": "kserve", "deployments": [
				{"name": "kserve-controller", "namespace": "redhat-ods-applications"},
				{"name": "kserve-webhook", "namespace": "redhat-ods-applications"}
			]},
			{"component": "notebooks", "deployments": [
				{"name": "notebook-controller", "namespace": "rhods-notebooks"}
			]},
			{"component": "odh-model-controller", "deployments": [
				{"name": "omc", "namespace": "redhat-ods-applications"}
			]}
		]
	}`
	path := writeTempJSON(t, platformJSON)

	p, err := extractPlatform(path, "kserve")
	if err != nil {
		t.Fatalf("extractPlatform() error: %v", err)
	}

	if len(p.Topology) != 2 {
		t.Fatalf("expected 2 namespaces (deduped), got %d", len(p.Topology))
	}
	nsNames := make(map[string]bool)
	for _, ns := range p.Topology {
		nsNames[ns.Name] = true
	}
	if !nsNames["redhat-ods-applications"] {
		t.Error("missing namespace redhat-ods-applications")
	}
	if !nsNames["rhods-notebooks"] {
		t.Error("missing namespace rhods-notebooks")
	}
}

func writeTempJSON(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "platform-architecture.json")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}
