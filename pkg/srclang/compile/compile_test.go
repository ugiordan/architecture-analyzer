package compile

import (
	"strings"
	"testing"
	"time"

	"github.com/ugiordan/architecture-analyzer/pkg/extractor"
	"github.com/ugiordan/architecture-analyzer/pkg/graph"
	"github.com/ugiordan/architecture-analyzer/pkg/query"
)

func TestCompile_ProducesValidDocument(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "ValidateWebhook",
		File: "webhook.go", Line: 10, EndLine: 30,
		Language: "go",
		Annotations: map[string]bool{"webhook_handler": true},
	})

	arch := &extractor.ComponentArchitecture{
		Component:       "test-operator",
		Repo:            "https://github.com/example/test-operator",
		CommitSHA:       "abc123",
		AnalyzerVersion: "0.2.0",
	}

	findings := []query.Finding{
		{RuleID: "CGA-S01", Severity: "medium", Message: "test finding",
			File: "webhook.go", Line: 15, Domain: "security"},
	}

	opts := Options{
		RepoPath: t.TempDir(),
		Layer:    "security",
		CPG:      cpg,
		Arch:     arch,
		Findings: findings,
	}

	doc, err := Compile(opts)
	if err != nil {
		t.Fatalf("Compile() error: %v", err)
	}

	if doc.Version != "0.0.1" {
		t.Errorf("version = %q, want %q", doc.Version, "0.0.1")
	}
	if doc.Head.Layer != "security" {
		t.Errorf("layer = %q, want %q", doc.Head.Layer, "security")
	}
	if doc.Head.Component != "test-operator" {
		t.Errorf("component = %q, want %q", doc.Head.Component, "test-operator")
	}
	if doc.Body.Layer.Name != "security" {
		t.Errorf("body layer = %q, want %q", doc.Body.Layer.Name, "security")
	}
}

func TestCompile_UnsupportedLayer(t *testing.T) {
	opts := Options{Layer: "unknown-layer", CPG: graph.NewCPG()}
	_, err := Compile(opts)
	if err == nil {
		t.Error("expected error for unsupported layer")
	}
	if err != nil && !strings.Contains(err.Error(), "unsupported layer") {
		t.Errorf("expected unsupported layer error, got: %v", err)
	}
}

func TestCompile_DetectsGoLanguage(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component:       "test-operator",
		GoASTMode:       "full",
		AnalyzerVersion: "0.2.0",
	}

	opts := Options{
		RepoPath: t.TempDir(),
		Layer:    "security",
		CPG:      cpg,
		Arch:     arch,
	}

	doc, err := Compile(opts)
	if err != nil {
		t.Fatalf("Compile() error: %v", err)
	}

	if len(doc.Head.Languages) == 0 {
		t.Error("expected at least one language")
	}
	if doc.Head.Languages[0].Name != "go" {
		t.Errorf("language = %q, want %q", doc.Head.Languages[0].Name, "go")
	}
}

func TestCompile_DetectsPythonLanguage(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component:       "test-operator",
		PythonK8sCalls:  []extractor.PythonK8sCall{{Operation: "create"}},
		AnalyzerVersion: "0.2.0",
	}

	opts := Options{
		RepoPath: t.TempDir(),
		Layer:    "security",
		CPG:      cpg,
		Arch:     arch,
	}

	doc, err := Compile(opts)
	if err != nil {
		t.Fatalf("Compile() error: %v", err)
	}

	found := false
	for _, lang := range doc.Head.Languages {
		if lang.Name == "python" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected python language to be detected")
	}
}

func TestCompile_GoExternalConnectionsNotPython(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component:       "go-operator",
		GoASTMode:       "full",
		AnalyzerVersion: "0.2.0",
		ExternalConnections: []extractor.ExternalConnection{
			{Type: "database", Service: "postgres"},
		},
	}

	opts := Options{
		RepoPath: t.TempDir(),
		Layer:    "security",
		CPG:      cpg,
		Arch:     arch,
	}

	doc, err := Compile(opts)
	if err != nil {
		t.Fatalf("Compile() error: %v", err)
	}

	for _, lang := range doc.Head.Languages {
		if lang.Name == "python" {
			t.Error("Go-only repo with external connections should not report python")
		}
	}
	if len(doc.Head.Languages) != 1 || doc.Head.Languages[0].Name != "go" {
		t.Errorf("expected [go], got %v", doc.Head.Languages)
	}
}

func TestCompile_IncludesRepositoryInfo(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component:       "test-operator",
		Repo:            "https://github.com/example/test-operator",
		CommitSHA:       "abc123def456",
		AnalyzerVersion: "0.2.0",
	}

	opts := Options{
		RepoPath: t.TempDir(),
		Layer:    "security",
		CPG:      cpg,
		Arch:     arch,
	}

	doc, err := Compile(opts)
	if err != nil {
		t.Fatalf("Compile() error: %v", err)
	}

	if doc.Head.Repository == nil {
		t.Fatal("expected repository info to be set")
	}
	if doc.Head.Repository.URI != "https://github.com/example/test-operator" {
		t.Errorf("repo URI = %q", doc.Head.Repository.URI)
	}
	if doc.Head.Repository.Commit != "abc123def456" {
		t.Errorf("commit = %q", doc.Head.Repository.Commit)
	}
}

func TestCompile_ProducerIncludesVersion(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component:       "test-operator",
		AnalyzerVersion: "0.2.0",
	}

	opts := Options{
		RepoPath: t.TempDir(),
		Layer:    "security",
		CPG:      cpg,
		Arch:     arch,
	}

	doc, err := Compile(opts)
	if err != nil {
		t.Fatalf("Compile() error: %v", err)
	}

	if doc.Head.Producer != "arch-analyzer 0.2.0" {
		t.Errorf("producer = %q, want %q", doc.Head.Producer, "arch-analyzer 0.2.0")
	}
}

func TestCompile_WithSecurityAnnotations(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component:       "test-operator",
		AnalyzerVersion: "0.2.0",
	}
	annotations := []extractor.SecurityAnnotation{
		{
			Type:        "RBAC_CLUSTER_SCOPE_SENSITIVE",
			Severity:    "high",
			Source:      "config/rbac/role.yaml",
			Description: "ClusterRole grants cluster-wide secrets CRUD",
		},
	}

	opts := Options{
		RepoPath:            t.TempDir(),
		Layer:               "security",
		CPG:                 cpg,
		Arch:                arch,
		SecurityAnnotations: annotations,
	}

	doc, err := Compile(opts)
	if err != nil {
		t.Fatalf("Compile() error: %v", err)
	}

	if len(doc.Body.Layer.Findings) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(doc.Body.Layer.Findings))
	}
	f := doc.Body.Layer.Findings[0]
	if f.Domain != "extraction" {
		t.Errorf("finding domain = %q, want %q", f.Domain, "extraction")
	}
	if f.Rule != "RBAC_CLUSTER_SCOPE_SENSITIVE" {
		t.Errorf("finding rule = %q, want %q", f.Rule, "RBAC_CLUSTER_SCOPE_SENSITIVE")
	}
}

func TestCompile_WithPlatformFile(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component:       "kserve",
		AnalyzerVersion: "0.2.0",
	}

	platformJSON := `{
		"platform": "RHOAI",
		"component_count": 38,
		"components": ["kserve", "odh-model-controller"],
		"dependency_graph": [
			{"from": "odh-model-controller", "to": "kserve", "type": "go-module"}
		],
		"component_data": []
	}`
	path := writeTempJSON(t, platformJSON)

	opts := Options{
		RepoPath:     t.TempDir(),
		Layer:        "security",
		CPG:          cpg,
		Arch:         arch,
		PlatformFile: path,
	}

	doc, err := Compile(opts)
	if err != nil {
		t.Fatalf("Compile() error: %v", err)
	}

	if doc.Head.Platform == nil {
		t.Fatal("expected Platform to be set")
	}
	if doc.Head.Platform.Name != "RHOAI" {
		t.Errorf("platform name = %q, want %q", doc.Head.Platform.Name, "RHOAI")
	}
	if len(doc.Head.Platform.Inbound) != 1 {
		t.Fatalf("expected 1 inbound edge, got %d", len(doc.Head.Platform.Inbound))
	}
}

func TestCompile_WithPlatformFile_InvalidFile(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component:       "test-operator",
		AnalyzerVersion: "0.2.0",
	}

	opts := Options{
		RepoPath:     t.TempDir(),
		Layer:        "security",
		CPG:          cpg,
		Arch:         arch,
		PlatformFile: "/nonexistent/platform.json",
	}

	doc, err := Compile(opts)
	if err != nil {
		t.Fatalf("expected no error (graceful degradation), got: %v", err)
	}

	if doc.Head.Platform != nil {
		t.Error("expected Platform to be nil on failure")
	}
	if len(doc.Head.Diagnostics) == 0 {
		t.Error("expected diagnostic warning for failed platform extraction")
	}
}

func TestCompile_NewLayerRoutes(t *testing.T) {
	layers := []string{"testing", "upgrade", "netpolicy", "codegen"}
	for _, layer := range layers {
		t.Run(layer, func(t *testing.T) {
			opts := Options{
				RepoPath: t.TempDir(),
				Layer:    layer,
				CPG:      graph.NewCPG(),
				Arch:     &extractor.ComponentArchitecture{Component: "smoke-test", AnalyzerVersion: "0.0.0"},
			}
			doc, err := Compile(opts)
			if err != nil {
				t.Fatalf("Compile(%q) error: %v", layer, err)
			}
			if doc.Head.Layer != layer {
				t.Errorf("Head.Layer = %q, want %q", doc.Head.Layer, layer)
			}
			if doc.Body.Layer.Name != layer {
				t.Errorf("Body.Layer.Name = %q, want %q", doc.Body.Layer.Name, layer)
			}
		})
	}
}

func TestCompile_ExtractedTimestamp(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component:       "test-operator",
		AnalyzerVersion: "0.2.0",
	}

	opts := Options{
		RepoPath: t.TempDir(),
		Layer:    "security",
		CPG:      cpg,
		Arch:     arch,
	}

	doc, err := Compile(opts)
	if err != nil {
		t.Fatalf("Compile() error: %v", err)
	}

	if doc.Head.Extracted == "" {
		t.Error("expected extracted timestamp")
	}
	if _, err := time.Parse(time.RFC3339, doc.Head.Extracted); err != nil {
		t.Errorf("not valid RFC3339: %v", err)
	}
}
