package extractor

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestParseBuildConfigYAML(t *testing.T) {
	dir := t.TempDir()
	content := `
ocp_version_min: "4.14"
ocp_version_max: "4.16"
PLATFORMS: linux/amd64 linux/arm64 linux/ppc64le linux/s390x
`
	os.WriteFile(filepath.Join(dir, "config.yaml"), []byte(content), 0o644)

	bc, err := ParseBuildConfig(dir)
	if err != nil {
		t.Fatal(err)
	}
	if bc.OCPVersions.Min != "4.14" {
		t.Errorf("expected min 4.14, got %q", bc.OCPVersions.Min)
	}
	if bc.OCPVersions.Max != "4.16" {
		t.Errorf("expected max 4.16, got %q", bc.OCPVersions.Max)
	}
	if len(bc.Architectures) < 2 {
		t.Errorf("expected at least 2 architectures, got %d: %v", len(bc.Architectures), bc.Architectures)
	}
}

func TestParseBuildConfigJSON(t *testing.T) {
	dir := t.TempDir()
	cfg := map[string]interface{}{
		"ocp_versions": map[string]string{
			"min": "4.13",
			"max": "4.15",
		},
		"architectures": []string{"amd64", "arm64"},
		"olm_features":  []string{"AllNamespaces", "OwnNamespace"},
	}
	data, _ := json.MarshalIndent(cfg, "", "  ")
	os.WriteFile(filepath.Join(dir, "config.json"), data, 0o644)

	bc, err := ParseBuildConfig(dir)
	if err != nil {
		t.Fatal(err)
	}
	if bc.OCPVersions.Min != "4.13" {
		t.Errorf("expected min 4.13, got %q", bc.OCPVersions.Min)
	}
	if len(bc.Architectures) != 2 {
		t.Errorf("expected 2 architectures, got %d", len(bc.Architectures))
	}
	if len(bc.OLMFeatures) != 2 {
		t.Errorf("expected 2 OLM features, got %d", len(bc.OLMFeatures))
	}
}

func TestParseBuildConfigMakefile(t *testing.T) {
	dir := t.TempDir()
	makefile := `
PLATFORMS ?= linux/amd64,linux/arm64,linux/s390x
GOARCH ?= amd64

.PHONY: build
build:
	go build ./...
`
	os.WriteFile(filepath.Join(dir, "Makefile"), []byte(makefile), 0o644)

	bc, err := ParseBuildConfig(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(bc.Architectures) == 0 {
		t.Error("expected architectures from Makefile")
	}
	found := false
	for _, a := range bc.Architectures {
		if a == "amd64" {
			found = true
		}
	}
	if !found {
		t.Errorf("expected amd64 in architectures: %v", bc.Architectures)
	}
}

func TestParseBuildConfigCSVTemplate(t *testing.T) {
	dir := t.TempDir()
	manifestsDir := filepath.Join(dir, "manifests")
	os.MkdirAll(manifestsDir, 0o755)

	csv := `
apiVersion: operators.coreos.com/v1alpha1
kind: ClusterServiceVersion
metadata:
  name: my-operator.v1.0.0
spec:
  minKubeVersion: "1.25.0"
  installModes:
    - type: OwnNamespace
      supported: true
    - type: AllNamespaces
      supported: true
`
	os.WriteFile(filepath.Join(manifestsDir, "my-operator.clusterserviceversion.yaml"), []byte(csv), 0o644)

	bc, err := ParseBuildConfig(dir)
	if err != nil {
		t.Fatal(err)
	}
	if bc.OperatorConfig == nil {
		t.Fatal("expected operator config from CSV")
	}
	if bc.OperatorConfig.MinKubeVersion != "1.25.0" {
		t.Errorf("expected minKubeVersion 1.25.0, got %q", bc.OperatorConfig.MinKubeVersion)
	}
}

func TestParseBuildConfigImageReferences(t *testing.T) {
	dir := t.TempDir()
	content := `# Image references
dashboard=quay.io/opendatahub/dashboard:latest
oauth-proxy=quay.io/openshift/oauth-proxy:v4.15
`
	os.WriteFile(filepath.Join(dir, "image-references"), []byte(content), 0o644)

	bc, err := ParseBuildConfig(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(bc.ImageMappings) != 2 {
		t.Fatalf("expected 2 image mappings, got %d", len(bc.ImageMappings))
	}
	if bc.ImageMappings[0].Component != "dashboard" {
		t.Errorf("expected component 'dashboard', got %q", bc.ImageMappings[0].Component)
	}
}

func TestParseBuildConfigEmpty(t *testing.T) {
	dir := t.TempDir()
	bc, err := ParseBuildConfig(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(bc.Architectures) != 0 {
		t.Errorf("expected no architectures for empty dir, got %v", bc.Architectures)
	}
}

func TestParseBuildConfigBadDir(t *testing.T) {
	_, err := ParseBuildConfig("/nonexistent/path")
	if err == nil {
		t.Error("expected error for nonexistent path")
	}
}

func TestNormalizeArch(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"x86_64", "amd64"},
		{"amd64", "amd64"},
		{"aarch64", "arm64"},
		{"arm64", "arm64"},
		{"ppc64le", "ppc64le"},
		{"s390x", "s390x"},
	}
	for _, tt := range tests {
		got := normalizeArch(tt.input)
		if got != tt.expected {
			t.Errorf("normalizeArch(%q) = %q, want %q", tt.input, got, tt.expected)
		}
	}
}

func TestDedupStrings(t *testing.T) {
	got := dedupStrings([]string{"a", "b", "a", "c", "b"})
	if len(got) != 3 {
		t.Errorf("expected 3 deduped strings, got %d: %v", len(got), got)
	}
}

func TestDedupStringsNil(t *testing.T) {
	got := dedupStrings(nil)
	if got != nil {
		t.Errorf("expected nil for nil input, got %v", got)
	}
}
