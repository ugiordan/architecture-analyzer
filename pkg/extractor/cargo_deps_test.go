package extractor

import (
	"os"
	"path/filepath"
	"testing"
)

const sampleCargoToml = `[package]
name = "praxis-extproc"
version = "0.1.0"
edition = "2024"
rust-version = "1.96"

[dependencies]
async-trait = "0.1.89"
clap = { version = "4.6.1", features = ["derive"] }
praxis-ai-filters = { git = "https://github.com/praxis-proxy/ai.git", tag = "v0.9.0", default-features = false }
praxis-core = { version = "0.7.4", package = "praxis-proxy-core" }
praxis-proto = { path = "proto", package = "praxis-proxy-proto" }

[dev-dependencies]
serde_json = "1.0.151"
tower = { version = "0.5.3", features = ["util"] }

[build-dependencies]
prost-build = "0.14"

[features]
default = ["aws-sigv4"]
aws-sigv4 = ["praxis-ai-filters/aws-sigv4-filter"]

[target.'cfg(unix)'.dependencies]
openssl = "0.10.81"

[lints]
rust = { dead_code = "deny" }

[profile.release]
lto = true
`

func TestExtractCargoDeps(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "Cargo.toml"), []byte(sampleCargoToml), 0644); err != nil {
		t.Fatal(err)
	}

	crates := extractCargoDeps(dir)
	byName := make(map[string]CargoCrate, len(crates))
	for _, c := range crates {
		byName[c.Name] = c
	}

	expect := []CargoCrate{
		{Name: "async-trait", Version: "0.1.89"},
		{Name: "clap", Version: "4.6.1"},
		{Name: "praxis-ai-filters", Version: "v0.9.0", Kind: "git"},
		{Name: "praxis-core (praxis-proxy-core)", Version: "0.7.4"},
		{Name: "praxis-proto (praxis-proxy-proto)", Version: "proto", Kind: "path"},
		{Name: "openssl", Version: "0.10.81"},
	}
	for _, want := range expect {
		got, ok := byName[want.Name]
		if !ok {
			t.Errorf("missing crate %q; got %+v", want.Name, crates)
			continue
		}
		if got.Version != want.Version {
			t.Errorf("%s version: got %q, want %q", want.Name, got.Version, want.Version)
		}
		if got.Kind != want.Kind {
			t.Errorf("%s kind: got %q, want %q", want.Name, got.Kind, want.Kind)
		}
	}

	if got, ok := byName["serde_json"]; !ok || got.Category != "dev" {
		t.Errorf("dev-dependency serde_json: got %+v, ok=%v", got, ok)
	}
	if got, ok := byName["prost-build"]; !ok || got.Category != "build" {
		t.Errorf("build-dependency prost-build: got %+v, ok=%v", got, ok)
	}

	// Non-dependency sections must not produce crates.
	for _, name := range []string{"default", "aws-sigv4", "dead_code", "lto", "praxis-extproc"} {
		if _, ok := byName[name]; ok {
			t.Errorf("unexpected crate %q from non-dependency section", name)
		}
	}

	// Sorted output.
	for i := 1; i < len(crates); i++ {
		if crates[i-1].Name > crates[i].Name {
			t.Errorf("crates not sorted: %q before %q", crates[i-1].Name, crates[i].Name)
			break
		}
	}
}

func TestExtractRustVersion(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "Cargo.toml"), []byte(sampleCargoToml), 0644); err != nil {
		t.Fatal(err)
	}
	if v := extractRustVersion(dir); v != "1.96" {
		t.Errorf("rust version: got %q, want 1.96", v)
	}
}

func TestExtractCargoDepsWorkspaceMembers(t *testing.T) {
	dir := t.TempDir()
	root := `[workspace]
members = ["proto"]

[workspace.dependencies]
tokio = "1.52.2"

[dependencies]
serde = "1.0.228"
`
	if err := os.WriteFile(filepath.Join(dir, "Cargo.toml"), []byte(root), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(dir, "proto"), 0755); err != nil {
		t.Fatal(err)
	}
	member := `[package]
name = "praxis-proxy-proto"

[dependencies]
prost = "0.14.3"
`
	if err := os.WriteFile(filepath.Join(dir, "proto", "Cargo.toml"), []byte(member), 0644); err != nil {
		t.Fatal(err)
	}

	crates := extractCargoDeps(dir)
	byName := make(map[string]CargoCrate, len(crates))
	for _, c := range crates {
		byName[c.Name] = c
	}

	if got, ok := byName["tokio"]; !ok || got.Category != "workspace" {
		t.Errorf("workspace dependency tokio: got %+v, ok=%v", got, ok)
	}
	if got, ok := byName["prost"]; !ok || got.Source != filepath.Join("proto", "Cargo.toml") {
		t.Errorf("member manifest dependency prost: got %+v, ok=%v", got, ok)
	}
	// The workspace table itself must not be misfiled as a normal dependency.
	if got, ok := byName["members"]; ok {
		t.Errorf("unexpected crate from [workspace] table: %+v", got)
	}
}

func TestExtractCargoDepsNoManifest(t *testing.T) {
	dir := t.TempDir()
	if crates := extractCargoDeps(dir); len(crates) != 0 {
		t.Errorf("expected no crates without Cargo.toml, got %+v", crates)
	}
	if v := extractRustVersion(dir); v != "" {
		t.Errorf("expected empty rust version, got %q", v)
	}
}
