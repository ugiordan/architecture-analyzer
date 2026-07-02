package srclang

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBodyExtractor_Extract(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "main.go")
	content := `package main

import "fmt"

func hello() {
	fmt.Println("hello")
}

func world() {
	fmt.Println("world")
}
`
	if err := os.WriteFile(src, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	be := NewBodyExtractor()

	tests := []struct {
		name      string
		file      string
		startLine int
		endLine   int
		want      string
	}{
		{
			name:      "extract hello function",
			file:      src,
			startLine: 5,
			endLine:   7,
			want:      "func hello() {\n\tfmt.Println(\"hello\")\n}",
		},
		{
			name:      "extract world function",
			file:      src,
			startLine: 9,
			endLine:   11,
			want:      "func world() {\n\tfmt.Println(\"world\")\n}",
		},
		{
			name:      "single line",
			file:      src,
			startLine: 1,
			endLine:   1,
			want:      "package main",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := be.Extract(tt.file, tt.startLine, tt.endLine)
			if err != nil {
				t.Fatalf("Extract() error: %v", err)
			}
			if got != tt.want {
				t.Errorf("Extract() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestBodyExtractor_CachesFileContent(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "main.go")
	if err := os.WriteFile(src, []byte("line1\nline2\nline3\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	be := NewBodyExtractor()

	got1, _ := be.Extract(src, 1, 1)
	if got1 != "line1" {
		t.Fatalf("first call: got %q, want %q", got1, "line1")
	}

	got2, _ := be.Extract(src, 2, 3)
	if got2 != "line2\nline3" {
		t.Fatalf("second call: got %q, want %q", got2, "line2\nline3")
	}

	if len(be.cache) != 1 {
		t.Errorf("cache should have 1 entry, got %d", len(be.cache))
	}
}

func TestBodyExtractor_MissingFile(t *testing.T) {
	be := NewBodyExtractor()
	_, err := be.Extract("/nonexistent/file.go", 1, 5)
	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestBodyExtractor_OutOfRange(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "main.go")
	if err := os.WriteFile(src, []byte("line1\nline2\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	be := NewBodyExtractor()
	got, err := be.Extract(src, 1, 100)
	if err != nil {
		t.Fatalf("Extract() error: %v", err)
	}
	if got != "line1\nline2" {
		t.Errorf("got %q, want clamped to available lines", got)
	}
}

func TestBodyExtractor_InvertedRange(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "main.go")
	if err := os.WriteFile(src, []byte("line1\nline2\nline3\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	be := NewBodyExtractor()
	got, err := be.Extract(src, 3, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "" {
		t.Errorf("expected empty string for inverted range, got %q", got)
	}
}
