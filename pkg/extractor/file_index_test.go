package extractor

import (
	"os"
	"path/filepath"
	"sort"
	"testing"
)

func TestFileIndex_FindByPatterns(t *testing.T) {
	dir := t.TempDir()

	files := map[string]string{
		"config/crd/bases/foo.yaml":     "apiVersion: v1",
		"config/crd/bases/bar.yaml":     "apiVersion: v1",
		"config/webhook/hook.yaml":      "kind: MutatingWebhookConfiguration",
		"pkg/controller/reconciler.go":  "package controller",
		"pkg/controller/suite_test.go":  "package controller",
		"cmd/main.go":                   "package main",
		"go.mod":                        "module test",
		"internal/cache/setup.go":       "package cache",
	}
	for rel, content := range files {
		abs := filepath.Join(dir, rel)
		os.MkdirAll(filepath.Dir(abs), 0o755)
		os.WriteFile(abs, []byte(content), 0o644)
	}

	fi := NewFileIndex(dir)

	tests := []struct {
		name     string
		patterns []string
		want     int
	}{
		{"recursive go", []string{"**/*.go"}, 4},
		{"recursive yaml", []string{"**/*.yaml"}, 3},
		{"crd bases", []string{"config/crd/bases/*.yaml"}, 2},
		{"specific file", []string{"go.mod"}, 1},
		{"multiple patterns", []string{"cmd/main.go", "go.mod"}, 2},
		{"no match", []string{"**/*.py"}, 0},
		{"prefix glob", []string{"pkg/**/*.go"}, 2},
		{"webhook yaml", []string{"config/webhook/*.yaml"}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fi.FindByPatterns(tt.patterns)
			if len(got) != tt.want {
				t.Errorf("FindByPatterns(%v) returned %d files, want %d\n  files: %v", tt.patterns, len(got), tt.want, got)
			}
		})
	}
}

func TestFileIndex_SkipsExcludedDirs(t *testing.T) {
	dir := t.TempDir()

	files := map[string]string{
		"pkg/main.go":          "package pkg",
		"vendor/dep/dep.go":    "package dep",
		"testdata/fixture.go":  "package testdata",
		".git/HEAD":            "ref: refs/heads/main",
		"test/e2e_test.go":     "package test",
		"docs/guide.md":        "# Guide",
	}
	for rel, content := range files {
		abs := filepath.Join(dir, rel)
		os.MkdirAll(filepath.Dir(abs), 0o755)
		os.WriteFile(abs, []byte(content), 0o644)
	}

	fi := NewFileIndex(dir)
	got := fi.FindByPatterns([]string{"**/*.go"})

	if len(got) != 1 {
		t.Errorf("expected 1 file (pkg/main.go), got %d: %v", len(got), got)
	}
}

func TestFileIndex_SkipsSymlinks(t *testing.T) {
	dir := t.TempDir()

	realFile := filepath.Join(dir, "real.go")
	os.WriteFile(realFile, []byte("package real"), 0o644)

	linkFile := filepath.Join(dir, "linked.go")
	os.Symlink(realFile, linkFile)

	fi := NewFileIndex(dir)
	got := fi.FindByPatterns([]string{"**/*.go"})

	if len(got) != 1 {
		t.Errorf("expected 1 file (real.go only, not symlink), got %d: %v", len(got), got)
	}
}

func TestFileIndex_Deduplicates(t *testing.T) {
	dir := t.TempDir()

	abs := filepath.Join(dir, "config", "crd", "bases", "crd.yaml")
	os.MkdirAll(filepath.Dir(abs), 0o755)
	os.WriteFile(abs, []byte("apiVersion: v1"), 0o644)

	fi := NewFileIndex(dir)
	got := fi.FindByPatterns([]string{
		"config/crd/bases/*.yaml",
		"**/*.yaml",
		"config/crd/**/*.yaml",
	})

	if len(got) != 1 {
		t.Errorf("expected 1 deduplicated result, got %d: %v", len(got), got)
	}
}

func TestFileIndex_MatchesOriginalFindFiles(t *testing.T) {
	dir := t.TempDir()

	files := map[string]string{
		"config/crd/bases/foo.yaml":    "apiVersion: v1",
		"config/crd/bases/bar.yml":     "apiVersion: v1",
		"config/webhook/hook.yaml":     "kind: Webhook",
		"pkg/controller/reconciler.go": "package controller",
		"cmd/main.go":                  "package main",
		"cmd/worker/main.go":           "package main",
		"internal/cache/cache.go":      "package cache",
	}
	for rel, content := range files {
		abs := filepath.Join(dir, rel)
		os.MkdirAll(filepath.Dir(abs), 0o755)
		os.WriteFile(abs, []byte(content), 0o644)
	}

	patterns := [][]string{
		{"config/crd/bases/*.yaml", "config/crd/bases/*.yml"},
		{"**/*.go"},
		{"cmd/main.go", "cmd/*/main.go"},
		{"config/webhook/*.yaml"},
		{"**/*cache*.go"},
	}

	fi := NewFileIndex(dir)

	for _, pats := range patterns {
		indexed := fi.FindByPatterns(pats)
		original := findFilesSlow(dir, pats)

		// Normalize to absolute paths for comparison
		normalizeAbs := func(paths []string) []string {
			var out []string
			for _, p := range paths {
				abs, _ := filepath.Abs(p)
				out = append(out, abs)
			}
			sort.Strings(out)
			return out
		}

		idxSorted := normalizeAbs(indexed)
		origSorted := normalizeAbs(original)

		if len(idxSorted) != len(origSorted) {
			t.Errorf("patterns %v: indexed=%d files, original=%d files\n  indexed:  %v\n  original: %v",
				pats, len(idxSorted), len(origSorted), idxSorted, origSorted)
			continue
		}

		for i := range idxSorted {
			if idxSorted[i] != origSorted[i] {
				t.Errorf("patterns %v: mismatch at index %d\n  indexed:  %s\n  original: %s",
					pats, i, idxSorted[i], origSorted[i])
			}
		}
	}
}
