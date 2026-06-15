package extractor

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractCopyAddInstructions(t *testing.T) {
	dockerfile := `FROM golang:1.22 AS builder
WORKDIR /workspace
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o manager cmd/main.go

FROM registry.access.redhat.com/ubi9/ubi-minimal:9.4
WORKDIR /
COPY --from=builder /workspace/manager .
COPY config/manifests/ /manifests/
COPY deploy/crds/my-crd.yaml /crds/
ADD https://example.com/file.tar.gz /tmp/
USER 65532:65532
EXPOSE 8080
`

	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "Dockerfile"), []byte(dockerfile), 0644); err != nil {
		t.Fatal(err)
	}

	result := extractDockerfiles(dir)
	if len(result) != 1 {
		t.Fatalf("expected 1 dockerfile, got %d", len(result))
	}

	df := result[0]
	if len(df.CopyInstructions) != 4 {
		t.Fatalf("expected 4 final-stage copy instructions, got %d", len(df.CopyInstructions))
	}

	tests := []struct {
		name            string
		idx             int
		sources         []string
		destination     string
		fromStage       string
		originalSources []string
		isURL           bool
	}{
		{
			name:            "COPY --from=builder",
			idx:             0,
			sources:         []string{"/workspace/manager"},
			destination:     ".",
			fromStage:       "builder",
			originalSources: []string{"go.mod", "go.sum", "."},
		},
		{
			name:        "COPY config/manifests/",
			idx:         1,
			sources:     []string{"config/manifests/"},
			destination: "/manifests/",
		},
		{
			name:        "COPY deploy/crds/",
			idx:         2,
			sources:     []string{"deploy/crds/my-crd.yaml"},
			destination: "/crds/",
		},
		{
			name:        "ADD URL",
			idx:         3,
			sources:     []string{"https://example.com/file.tar.gz"},
			destination: "/tmp/",
			isURL:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ci := df.CopyInstructions[tt.idx]
			if ci.FromStage != tt.fromStage {
				t.Errorf("from_stage: got %q, want %q", ci.FromStage, tt.fromStage)
			}
			if ci.IsURL != tt.isURL {
				t.Errorf("is_url: got %v, want %v", ci.IsURL, tt.isURL)
			}
			if ci.Destination != tt.destination {
				t.Errorf("destination: got %q, want %q", ci.Destination, tt.destination)
			}
			assertStringSlice(t, "sources", ci.Sources, tt.sources)
			assertStringSlice(t, "original_sources", ci.OriginalSources, tt.originalSources)
		})
	}
}

func TestExtractCopyAddThreeStages(t *testing.T) {
	dockerfile := `FROM golang:1.22 AS deps
COPY go.mod go.sum ./
RUN go mod download

FROM deps AS builder
COPY . .
RUN go build -o app

FROM scratch
COPY --from=builder /app /app
COPY config/crd/bases/ /crds/
`

	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "Dockerfile"), []byte(dockerfile), 0644); err != nil {
		t.Fatal(err)
	}

	result := extractDockerfiles(dir)
	df := result[0]

	if len(df.CopyInstructions) != 2 {
		t.Fatalf("expected 2 final-stage copy instructions, got %d", len(df.CopyInstructions))
	}

	// COPY --from=builder traces through deps→builder chain
	ci0 := df.CopyInstructions[0]
	if ci0.FromStage != "builder" {
		t.Errorf("from_stage: got %q, want %q", ci0.FromStage, "builder")
	}
	assertStringSlice(t, "original_sources", ci0.OriginalSources, []string{".", "go.mod", "go.sum"})

	// Direct host copy
	ci1 := df.CopyInstructions[1]
	if ci1.FromStage != "" {
		t.Errorf("from_stage: got %q, want empty", ci1.FromStage)
	}
	if ci1.OriginalSources != nil {
		t.Errorf("original_sources should be nil for host copy, got %v", ci1.OriginalSources)
	}
}

func TestExtractCopyAddJSONForm(t *testing.T) {
	dockerfile := `FROM alpine
COPY ["src/app.yaml", "config/base.yml", "/app/"]
`

	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "Dockerfile"), []byte(dockerfile), 0644); err != nil {
		t.Fatal(err)
	}

	result := extractDockerfiles(dir)
	ci := result[0].CopyInstructions[0]

	if len(ci.Sources) != 2 {
		t.Fatalf("expected 2 sources, got %d", len(ci.Sources))
	}
	if ci.Sources[0] != "src/app.yaml" || ci.Sources[1] != "config/base.yml" {
		t.Errorf("sources: got %v", ci.Sources)
	}
	if ci.Destination != "/app/" {
		t.Errorf("destination: got %q", ci.Destination)
	}
}

func TestExtractCopyAddBuildArgPath(t *testing.T) {
	dockerfile := `FROM alpine
COPY ${DIR}/ /app/
`

	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "Dockerfile"), []byte(dockerfile), 0644); err != nil {
		t.Fatal(err)
	}

	result := extractDockerfiles(dir)
	ci := result[0].CopyInstructions[0]

	if ci.Sources[0] != "${DIR}/" {
		t.Errorf("expected literal ${DIR}/, got %q", ci.Sources[0])
	}
}

func TestExtractCopyAddOriginalSources(t *testing.T) {
	dockerfile := `FROM alpine AS fetcher
COPY config/manifests/ /data/
RUN process /data/

FROM scratch
COPY --from=fetcher /data/output /output
`

	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "Dockerfile"), []byte(dockerfile), 0644); err != nil {
		t.Fatal(err)
	}

	result := extractDockerfiles(dir)
	ci := result[0].CopyInstructions[0]

	if ci.FromStage != "fetcher" {
		t.Errorf("from_stage: got %q, want %q", ci.FromStage, "fetcher")
	}
	assertStringSlice(t, "original_sources", ci.OriginalSources, []string{"config/manifests/"})
}

func TestSplitCopyArgs(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{"src dest", []string{"src", "dest"}},
		{"a.yaml b.yaml /dest/", []string{"a.yaml", "b.yaml", "/dest/"}},
		{`["src", "dest"]`, []string{"src", "dest"}},
		{`["a.yaml", "b.yaml", "/app/"]`, []string{"a.yaml", "b.yaml", "/app/"}},
		{`["a,b.txt", "dest/"]`, []string{"a,b.txt", "dest/"}},
	}

	for _, tt := range tests {
		got := splitCopyArgs(tt.input)
		if len(got) != len(tt.want) {
			t.Errorf("splitCopyArgs(%q) = %v, want %v", tt.input, got, tt.want)
			continue
		}
		for i := range got {
			if got[i] != tt.want[i] {
				t.Errorf("splitCopyArgs(%q)[%d] = %q, want %q", tt.input, i, got[i], tt.want[i])
			}
		}
	}
}

func TestCollectStageHostSources(t *testing.T) {
	hostSources := map[int][]string{
		0: {"go.mod", "go.sum"},
		1: {"."},
	}
	parents := map[int]int{1: 0}

	// Stage 1 inherits from stage 0
	got := collectStageHostSources(1, hostSources, parents)
	want := []string{".", "go.mod", "go.sum"}
	assertStringSlice(t, "collected", got, want)

	// Stage 0 standalone
	got0 := collectStageHostSources(0, hostSources, parents)
	assertStringSlice(t, "stage0", got0, []string{"go.mod", "go.sum"})

	// Stage with no sources
	got2 := collectStageHostSources(2, hostSources, parents)
	if len(got2) != 0 {
		t.Errorf("expected empty, got %v", got2)
	}

	// Circular parent chain must not loop forever
	circParents := map[int]int{0: 1, 1: 0}
	circSources := map[int][]string{0: {"a"}, 1: {"b"}}
	gotCirc := collectStageHostSources(0, circSources, circParents)
	assertStringSlice(t, "circular", gotCirc, []string{"a", "b"})
}

func TestMultilineCopyWithBackslash(t *testing.T) {
	dockerfile := `FROM golang:1.22 as builder
COPY go.mod go.sum ./

FROM registry.access.redhat.com/ubi9/ubi-minimal:9.4
COPY --from=builder \
     /workspace/manager \
     .
ADD config/manifests/ /manifests/
`

	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "Dockerfile"), []byte(dockerfile), 0644); err != nil {
		t.Fatal(err)
	}

	result := extractDockerfiles(dir)
	if len(result) != 1 {
		t.Fatalf("expected 1 dockerfile, got %d", len(result))
	}

	df := result[0]
	if len(df.CopyInstructions) != 2 {
		t.Fatalf("expected 2 final-stage copy instructions, got %d: %+v", len(df.CopyInstructions), df.CopyInstructions)
	}

	found := false
	for _, c := range df.CopyInstructions {
		if c.FromStage == "builder" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected FromStage=builder in copy instructions, got %+v", df.CopyInstructions)
	}
}

func TestLowercaseDirectives(t *testing.T) {
	dockerfile := `from golang:1.22 as builder
copy go.mod go.sum ./
add https://example.com/file.tar.gz /tmp/

FROM registry.access.redhat.com/ubi9/ubi-minimal:9.4
Copy --from=builder /workspace/manager .
ADD config/manifests/ /manifests/
`

	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "Dockerfile"), []byte(dockerfile), 0644); err != nil {
		t.Fatal(err)
	}

	result := extractDockerfiles(dir)
	if len(result) != 1 {
		t.Fatalf("expected 1 dockerfile, got %d", len(result))
	}

	df := result[0]

	// Should detect 2 FROM instructions (case-insensitive)
	if df.Stages != 2 {
		t.Errorf("expected 2 stages, got %d", df.Stages)
	}

	// Should detect 2 final-stage COPY/ADD instructions (case-insensitive)
	if len(df.CopyInstructions) != 2 {
		t.Fatalf("expected 2 final-stage copy instructions, got %d", len(df.CopyInstructions))
	}
}

func assertStringSlice(t *testing.T, label string, got, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("%s: got %v (len %d), want %v (len %d)", label, got, len(got), want, len(want))
		return
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("%s[%d]: got %q, want %q", label, i, got[i], want[i])
		}
	}
}
