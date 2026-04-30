package extractor

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestParseKonfluxSnapshotCR(t *testing.T) {
	dir := t.TempDir()
	cr := map[string]interface{}{
		"kind":       "Snapshot",
		"apiVersion": "appstudio.redhat.com/v1alpha1",
		"spec": map[string]interface{}{
			"application": "rhoai",
			"components": []interface{}{
				map[string]interface{}{
					"name":           "odh-dashboard",
					"containerImage": "quay.io/rhoai/odh-dashboard@sha256:abc123",
					"source": map[string]interface{}{
						"git": map[string]interface{}{
							"url":      "https://github.com/red-hat-data-services/odh-dashboard",
							"revision": "abc123def456",
						},
					},
				},
				map[string]interface{}{
					"name":           "kserve-controller",
					"containerImage": "quay.io/rhoai/kserve-controller:v0.12",
					"source": map[string]interface{}{
						"git": map[string]interface{}{
							"url":      "https://github.com/red-hat-data-services/kserve",
							"revision": "def789",
						},
					},
				},
			},
		},
	}
	data, _ := json.MarshalIndent(cr, "", "  ")
	path := filepath.Join(dir, "snapshot.json")
	os.WriteFile(path, data, 0o644)

	snap, err := ParseKonfluxSnapshot(path)
	if err != nil {
		t.Fatal(err)
	}
	if snap.Application != "rhoai" {
		t.Errorf("expected application 'rhoai', got %q", snap.Application)
	}
	if len(snap.Components) != 2 {
		t.Fatalf("expected 2 components, got %d", len(snap.Components))
	}
	if snap.Components[0].Name != "odh-dashboard" {
		t.Errorf("expected name 'odh-dashboard', got %q", snap.Components[0].Name)
	}
	if snap.Components[0].Repository != "https://github.com/red-hat-data-services/odh-dashboard" {
		t.Errorf("expected repo URL, got %q", snap.Components[0].Repository)
	}
	if snap.Components[0].Revision != "abc123def456" {
		t.Errorf("expected revision 'abc123def456', got %q", snap.Components[0].Revision)
	}
}

func TestParseKonfluxSnapshotFlat(t *testing.T) {
	dir := t.TempDir()
	flat := map[string]interface{}{
		"application": "odh",
		"components": []interface{}{
			map[string]interface{}{
				"name":           "dashboard",
				"containerImage": "quay.io/odh/dashboard:latest",
				"repository":     "https://github.com/opendatahub-io/odh-dashboard",
				"revision":       "main",
			},
		},
	}
	data, _ := json.MarshalIndent(flat, "", "  ")
	path := filepath.Join(dir, "snap.json")
	os.WriteFile(path, data, 0o644)

	snap, err := ParseKonfluxSnapshot(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(snap.Components) != 1 {
		t.Fatalf("expected 1 component, got %d", len(snap.Components))
	}
	if snap.Components[0].Repository != "https://github.com/opendatahub-io/odh-dashboard" {
		t.Errorf("unexpected repo: %q", snap.Components[0].Repository)
	}
}

func TestParseKonfluxSnapshotBadFormat(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "bad.json"), []byte(`{"foo": "bar"}`), 0o644)
	_, err := ParseKonfluxSnapshot(filepath.Join(dir, "bad.json"))
	if err == nil {
		t.Error("expected error for unrecognized format")
	}
}

func TestParseKonfluxSnapshotMissing(t *testing.T) {
	_, err := ParseKonfluxSnapshot("/nonexistent/snapshot.json")
	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestParseKonfluxDir(t *testing.T) {
	dir := t.TempDir()

	// Write two snapshot files
	for i, app := range []string{"app1", "app2"} {
		snap := map[string]interface{}{
			"application": app,
			"components": []interface{}{
				map[string]interface{}{
					"name":           "comp-" + app,
					"containerImage": "quay.io/test/" + app + ":latest",
					"repository":     "https://github.com/test/" + app,
					"revision":       "abc",
				},
			},
		}
		data, _ := json.MarshalIndent(snap, "", "  ")
		os.WriteFile(filepath.Join(dir, "snapshot"+string(rune('0'+i))+".json"), data, 0o644)
	}

	idx, err := ParseKonfluxDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if idx.Snapshots != 2 {
		t.Errorf("expected 2 snapshots parsed, got %d", idx.Snapshots)
	}
	if len(idx.Components()) != 2 {
		t.Errorf("expected 2 components, got %d", len(idx.Components()))
	}
}

func TestKonfluxImageIndexLookup(t *testing.T) {
	idx := &KonfluxImageIndex{
		Images: map[string]KonfluxComponent{
			"quay.io/test/dashboard:v1.0": {
				Name:           "dashboard",
				ContainerImage: "quay.io/test/dashboard:v1.0",
				Repository:     "https://github.com/test/dashboard",
			},
			"quay.io/test/dashboard": {
				Name:           "dashboard",
				ContainerImage: "quay.io/test/dashboard:v1.0",
				Repository:     "https://github.com/test/dashboard",
			},
		},
	}

	// Exact match
	c, ok := idx.Lookup("quay.io/test/dashboard:v1.0")
	if !ok {
		t.Error("expected exact match")
	}
	if c.Name != "dashboard" {
		t.Errorf("expected name 'dashboard', got %q", c.Name)
	}

	// Base match (without tag)
	c, ok = idx.Lookup("quay.io/test/dashboard:v2.0")
	if !ok {
		t.Error("expected base match")
	}

	// Name-only match
	c, ok = idx.Lookup("registry.example.com/other/dashboard:latest")
	if !ok {
		t.Error("expected name-only match")
	}

	// No match
	_, ok = idx.Lookup("quay.io/test/unknown:v1.0")
	if ok {
		t.Error("expected no match for unknown image")
	}
}

func TestStripTagDigest(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"quay.io/test/img:v1.0", "quay.io/test/img"},
		{"quay.io/test/img@sha256:abc123", "quay.io/test/img"},
		{"quay.io/test/img", "quay.io/test/img"},
		{"localhost:5000/img:tag", "localhost:5000/img"},
	}
	for _, tt := range tests {
		got := stripTagDigest(tt.input)
		if got != tt.expected {
			t.Errorf("stripTagDigest(%q) = %q, want %q", tt.input, got, tt.expected)
		}
	}
}

func TestImageBaseName(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"quay.io/test/dashboard:v1.0", "dashboard"},
		{"registry.io/org/my-app@sha256:abc", "my-app"},
		{"simple-image", "simple-image"},
	}
	for _, tt := range tests {
		got := imageBaseName(tt.input)
		if got != tt.expected {
			t.Errorf("imageBaseName(%q) = %q, want %q", tt.input, got, tt.expected)
		}
	}
}

func TestParseKonfluxDirEmpty(t *testing.T) {
	dir := t.TempDir()
	idx, err := ParseKonfluxDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if idx.Snapshots != 0 {
		t.Errorf("expected 0 snapshots for empty dir, got %d", idx.Snapshots)
	}
}
