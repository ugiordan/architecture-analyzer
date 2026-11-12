package main

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ugiordan/architecture-analyzer/pkg/query"
)

func TestNormalizeFindingsWritesEmptyArray(t *testing.T) {
	encoded, err := json.Marshal(normalizeFindings(nil))
	if err != nil {
		t.Fatal(err)
	}
	if string(encoded) != "[]" {
		t.Fatalf("normalizeFindings(nil) = %s, want []", encoded)
	}

	findings := []query.Finding{{}}
	if got := normalizeFindings(findings); !reflect.DeepEqual(got, findings) {
		t.Fatalf("normalizeFindings(non-nil) = %v, want %v", got, findings)
	}
}

func TestCollectRepoSHAsIncludesAnalyzerVersions(t *testing.T) {
	resultsDir := t.TempDir()
	componentDir := filepath.Join(resultsDir, "org", "repo")
	if err := os.MkdirAll(componentDir, 0o755); err != nil {
		t.Fatal(err)
	}
	component := `{"repo":"org/repo","commit_sha":"abc123","analyzer_version":"v2"}`
	if err := os.WriteFile(filepath.Join(componentDir, "component-architecture.json"), []byte(component), 0o644); err != nil {
		t.Fatal(err)
	}

	repos, analyzerVersions, err := collectRepoSHAs(resultsDir)
	if err != nil {
		t.Fatal(err)
	}
	if repos["org/repo"] != "abc123" {
		t.Fatalf("repo SHA = %q, want abc123", repos["org/repo"])
	}
	if analyzerVersions["org/repo"] != "v2" {
		t.Fatalf("repo analyzer version = %q, want v2", analyzerVersions["org/repo"])
	}
}

func TestWriteSnapshotMetadataIncludesRepoAnalyzerVersions(t *testing.T) {
	dir := t.TempDir()
	want := map[string]string{"org/repo": "v1"}
	if err := writeSnapshotMetadata(
		dir,
		"2026-11-12",
		map[string]string{"org/repo": "abc123"},
		want,
		"odh",
	); err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "snapshot-metadata.json"))
	if err != nil {
		t.Fatal(err)
	}
	var metadata snapshotMetadata
	if err := json.Unmarshal(data, &metadata); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(metadata.RepoAnalyzerVersions, want) {
		t.Fatalf("repo analyzer versions = %v, want %v", metadata.RepoAnalyzerVersions, want)
	}
}

func TestReorderArgs(t *testing.T) {
	newFS := func() *flag.FlagSet {
		fs := flag.NewFlagSet("test", flag.ContinueOnError)
		fs.String("output", "", "")
		fs.String("extractors", "", "")
		fs.String("org", "", "")
		return fs
	}

	tests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "flags before positional",
			args: []string{"--output", "out.json", "--extractors", "docker", "repo"},
			want: []string{"--output", "out.json", "--extractors", "docker", "repo"},
		},
		{
			name: "positional before flags",
			args: []string{"repo", "--output", "out.json", "--extractors", "docker"},
			want: []string{"--output", "out.json", "--extractors", "docker", "repo"},
		},
		{
			name: "mixed order",
			args: []string{"--output", "out.json", "repo", "--extractors", "docker"},
			want: []string{"--output", "out.json", "--extractors", "docker", "repo"},
		},
		{
			name: "no flags",
			args: []string{"repo"},
			want: []string{"repo"},
		},
		{
			name: "flag with equals",
			args: []string{"repo", "--output=out.json"},
			want: []string{"--output=out.json", "repo"},
		},
		{
			name: "empty args",
			args: []string{},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := newFS()
			got := reorderArgs(fs, tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reorderArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
