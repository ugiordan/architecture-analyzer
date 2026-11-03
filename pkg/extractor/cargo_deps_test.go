package extractor

import (
	"os"
	"path/filepath"
	"strings"
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
		{Name: "praxis-ai-filters", Version: "", Kind: "git", Ref: "v0.9.0"},
		{Name: "praxis-core (praxis-proxy-core)", Version: "0.7.4"},
		{Name: "praxis-proto (praxis-proxy-proto)", Version: "", Kind: "path", Ref: "proto"},
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
		if got.Ref != want.Ref {
			t.Errorf("%s ref: got %q, want %q", want.Name, got.Ref, want.Ref)
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

func TestExtractRustVersionWorkspacePackage(t *testing.T) {
	dir := t.TempDir()
	manifest := `[workspace]
members = ["crate-a"]

[workspace.package]
rust-version = "1.85"
edition = "2021"
`
	if err := os.WriteFile(filepath.Join(dir, "Cargo.toml"), []byte(manifest), 0644); err != nil {
		t.Fatal(err)
	}
	if v := extractRustVersion(dir); v != "1.85" {
		t.Errorf("rust version: got %q, want 1.85", v)
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

func TestExtractCargoDepsWorkspaceExclude(t *testing.T) {
	dir := t.TempDir()
	root := `[workspace]
members = ["app"]
exclude = ["vendored-thing"]
`
	if err := os.WriteFile(filepath.Join(dir, "Cargo.toml"), []byte(root), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(dir, "app"), 0755); err != nil {
		t.Fatal(err)
	}
	app := `[package]
name = "app"

[dependencies]
serde = "1.0.228"
`
	if err := os.WriteFile(filepath.Join(dir, "app", "Cargo.toml"), []byte(app), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(dir, "vendored-thing", "nested"), 0755); err != nil {
		t.Fatal(err)
	}
	excluded := `[package]
name = "vendored-thing"

[dependencies]
serde = "0.9.0"
sneaky = "1.0.0"
`
	if err := os.WriteFile(filepath.Join(dir, "vendored-thing", "Cargo.toml"), []byte(excluded), 0644); err != nil {
		t.Fatal(err)
	}

	crates := extractCargoDeps(dir)
	for _, c := range crates {
		if c.Name == "sneaky" {
			t.Errorf("excluded package dependency leaked into results: %+v", c)
		}
		if c.Name == "serde" && c.Version == "0.9.0" {
			t.Errorf("excluded package overrode member dependency: %+v", c)
		}
	}
	byName := make(map[string]CargoCrate, len(crates))
	for _, c := range crates {
		byName[c.Name] = c
	}
	if got, ok := byName["serde"]; !ok || got.Version != "1.0.228" || got.Source != filepath.Join("app", "Cargo.toml") {
		t.Errorf("member dependency serde: got %+v, ok=%v", got, ok)
	}
}

func TestExtractCargoDepsWorkspaceInherit(t *testing.T) {
	dir := t.TempDir()
	root := `[workspace]
members = ["app"]

[workspace.dependencies]
regex = "1.10.3"
`
	if err := os.WriteFile(filepath.Join(dir, "Cargo.toml"), []byte(root), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(dir, "app"), 0755); err != nil {
		t.Fatal(err)
	}
	member := `[package]
name = "app"

[dependencies]
regex = { workspace = true }
`
	if err := os.WriteFile(filepath.Join(dir, "app", "Cargo.toml"), []byte(member), 0644); err != nil {
		t.Fatal(err)
	}

	crates := extractCargoDeps(dir)
	byName := make(map[string]CargoCrate, len(crates))
	for _, c := range crates {
		byName[c.Name] = c
	}
	got, ok := byName["regex"]
	if !ok {
		t.Fatalf("missing regex dependency; got %+v", crates)
	}
	if got.Version != "1.10.3" {
		t.Errorf("inherited version: got %q, want 1.10.3", got.Version)
	}
	if got.Category != "" {
		t.Errorf("inherited dependency category: got %q, want \"\" (member-level)", got.Category)
	}
	if got.Source != filepath.Join("app", "Cargo.toml") {
		t.Errorf("inherited dependency source: got %q, want member manifest", got.Source)
	}
}

func TestExtractCargoDepsVersionAndGitCombo(t *testing.T) {
	dir := t.TempDir()
	manifest := `[package]
name = "combo"

[dependencies]
foo = { version = "1.10.3", git = "https://github.com/example/foo.git", branch = "next" }
bar = { version = "2.0.0", path = "../bar" }
`
	if err := os.WriteFile(filepath.Join(dir, "Cargo.toml"), []byte(manifest), 0644); err != nil {
		t.Fatal(err)
	}

	crates := extractCargoDeps(dir)
	byName := make(map[string]CargoCrate, len(crates))
	for _, c := range crates {
		byName[c.Name] = c
	}

	if got, ok := byName["foo"]; !ok {
		t.Fatalf("missing foo; got %+v", crates)
	} else {
		if got.Version != "1.10.3" {
			t.Errorf("foo version: got %q, want 1.10.3", got.Version)
		}
		if got.Kind != "git" {
			t.Errorf("foo kind: got %q, want git", got.Kind)
		}
		if got.Ref != "next" {
			t.Errorf("foo ref: got %q, want next", got.Ref)
		}
	}
	if got, ok := byName["bar"]; !ok {
		t.Fatalf("missing bar; got %+v", crates)
	} else {
		if got.Version != "2.0.0" {
			t.Errorf("bar version: got %q, want 2.0.0", got.Version)
		}
		if got.Kind != "path" {
			t.Errorf("bar kind: got %q, want path", got.Kind)
		}
		if got.Ref != "../bar" {
			t.Errorf("bar ref: got %q, want ../bar", got.Ref)
		}
	}
}

func TestExtractCargoDepsSubtables(t *testing.T) {
	dir := t.TempDir()
	manifest := `[package]
name = "subtables"

[dependencies.awesome]
version = "1.3.5"

[dev-dependencies.old]
git = "https://github.com/example/old.git"
tag = "v1"

[package.metadata.dependencies]
something = "not-a-crate"

[target.'cfg(unix)'.dependencies.unix-only]
version = "0.1"
`
	if err := os.WriteFile(filepath.Join(dir, "Cargo.toml"), []byte(manifest), 0644); err != nil {
		t.Fatal(err)
	}

	crates := extractCargoDeps(dir)
	byName := make(map[string]CargoCrate, len(crates))
	for _, c := range crates {
		byName[c.Name] = c
	}

	if got, ok := byName["awesome"]; !ok || got.Version != "1.3.5" || got.Category != "" {
		t.Errorf("dependency subtable awesome: got %+v, ok=%v", got, ok)
	}
	if got, ok := byName["old"]; !ok || got.Category != "dev" || got.Kind != "git" || got.Ref != "v1" {
		t.Errorf("dev dependency subtable old: got %+v, ok=%v", got, ok)
	}
	if got, ok := byName["unix-only"]; !ok || got.Version != "0.1" {
		t.Errorf("target dependency subtable unix-only: got %+v, ok=%v", got, ok)
	}
	for _, name := range []string{"version", "git", "tag", "something"} {
		if _, ok := byName[name]; ok {
			t.Errorf("subtable field %q leaked as a crate", name)
		}
	}
}

func TestExtractCargoDepsMultilineExclude(t *testing.T) {
	dir := t.TempDir()
	root := `[workspace]
members = ["app"]
exclude = [
    "vendored-thing",
    "also-excluded",
]
`
	if err := os.WriteFile(filepath.Join(dir, "Cargo.toml"), []byte(root), 0644); err != nil {
		t.Fatal(err)
	}
	for _, excluded := range []string{"vendored-thing", "also-excluded"} {
		if err := os.MkdirAll(filepath.Join(dir, excluded), 0755); err != nil {
			t.Fatal(err)
		}
		manifest := "[package]\nname = \"x\"\n\n[dependencies]\nsneaky = \"1.0.0\"\n"
		if err := os.WriteFile(filepath.Join(dir, excluded, "Cargo.toml"), []byte(manifest), 0644); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.MkdirAll(filepath.Join(dir, "app"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "app", "Cargo.toml"), []byte("[package]\nname = \"app\"\n"), 0644); err != nil {
		t.Fatal(err)
	}

	crates := extractCargoDeps(dir)
	if len(crates) != 0 {
		t.Errorf("excluded package dependencies leaked into results: %+v", crates)
	}
}

func TestExtractCargoDepsBareTargetSection(t *testing.T) {
	dir := t.TempDir()
	manifest := `[package]
name = "bare-target"

[target.x86_64-unknown-linux-gnu.dependencies]
foo = "1.0.0"

[target.'cfg(windows)'.dev-dependencies]
win-only = "0.1"
`
	if err := os.WriteFile(filepath.Join(dir, "Cargo.toml"), []byte(manifest), 0644); err != nil {
		t.Fatal(err)
	}

	crates := extractCargoDeps(dir)
	byName := make(map[string]CargoCrate, len(crates))
	for _, c := range crates {
		byName[c.Name] = c
	}
	if got, ok := byName["foo"]; !ok || got.Version != "1.0.0" {
		t.Errorf("bare target dependency foo: got %+v, ok=%v", got, ok)
	}
	if got, ok := byName["win-only"]; !ok || got.Category != "dev" {
		t.Errorf("quoted target dev-dependency win-only: got %+v, ok=%v", got, ok)
	}
}

func TestExtractCargoDepsWorkspaceInheritRenamed(t *testing.T) {
	dir := t.TempDir()
	root := `[workspace]
members = ["app"]

[workspace.dependencies]
alias = { package = "actual", version = "1.2.3" }
`
	if err := os.WriteFile(filepath.Join(dir, "Cargo.toml"), []byte(root), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(dir, "app"), 0755); err != nil {
		t.Fatal(err)
	}
	member := `[package]
name = "app"

[dependencies]
alias = { workspace = true }
`
	if err := os.WriteFile(filepath.Join(dir, "app", "Cargo.toml"), []byte(member), 0644); err != nil {
		t.Fatal(err)
	}

	crates := extractCargoDeps(dir)
	var renamed []CargoCrate
	for _, c := range crates {
		if strings.Contains(c.Name, "alias") || strings.Contains(c.Name, "actual") {
			renamed = append(renamed, c)
		}
	}
	if len(renamed) != 1 {
		t.Fatalf("expected a single resolved entry, got %+v", renamed)
	}
	got := renamed[0]
	if got.Name != "alias (actual)" {
		t.Errorf("resolved name: got %q, want %q", got.Name, "alias (actual)")
	}
	if got.Version != "1.2.3" {
		t.Errorf("resolved version: got %q, want 1.2.3", got.Version)
	}
	if got.Category != "" {
		t.Errorf("resolved category: got %q, want \"\" (member-level)", got.Category)
	}
}

func TestExtractRustVersionRootManifestOnly(t *testing.T) {
	dir := t.TempDir()
	root := "[package]\nname = \"root\"\nrust-version = \"1.90\"\n"
	if err := os.WriteFile(filepath.Join(dir, "Cargo.toml"), []byte(root), 0644); err != nil {
		t.Fatal(err)
	}
	// A nested manifest whose basename is Cargo.toml (e.g. inside a directory
	// named Cargo) must not override the root's rust-version.
	if err := os.MkdirAll(filepath.Join(dir, "Cargo"), 0755); err != nil {
		t.Fatal(err)
	}
	nested := "[package]\nname = \"nested\"\nrust-version = \"1.50\"\n"
	if err := os.WriteFile(filepath.Join(dir, "Cargo", "Cargo.toml"), []byte(nested), 0644); err != nil {
		t.Fatal(err)
	}
	if v := extractRustVersion(dir); v != "1.90" {
		t.Errorf("rust version: got %q, want 1.90 (root manifest)", v)
	}
}

func TestExtractCargoDepsDottedKeys(t *testing.T) {
	dir := t.TempDir()
	root := `[workspace]
members = ["app"]

[workspace.dependencies]
serde = "1.0.228"
`
	if err := os.WriteFile(filepath.Join(dir, "Cargo.toml"), []byte(root), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(dir, "app"), 0755); err != nil {
		t.Fatal(err)
	}
	member := `[package]
name = "app"

[dependencies]
serde.workspace = true
foo.version = "1.2"
`
	if err := os.WriteFile(filepath.Join(dir, "app", "Cargo.toml"), []byte(member), 0644); err != nil {
		t.Fatal(err)
	}

	crates := extractCargoDeps(dir)
	byName := make(map[string]CargoCrate, len(crates))
	for _, c := range crates {
		byName[c.Name] = c
	}
	if got, ok := byName["serde"]; !ok || got.Version != "1.0.228" || got.Category != "" {
		t.Errorf("dotted workspace inheritance serde: got %+v, ok=%v", got, ok)
	}
	if got, ok := byName["foo"]; !ok || got.Version != "1.2" {
		t.Errorf("dotted version declaration foo: got %+v, ok=%v", got, ok)
	}
	// Dotted field names must not leak as crates.
	for _, name := range []string{"workspace", "version"} {
		if _, ok := byName[name]; ok {
			t.Errorf("dotted field %q leaked as a crate", name)
		}
	}
}

func TestExtractCargoDepsSubtableWorkspaceInherit(t *testing.T) {
	dir := t.TempDir()
	root := `[workspace]
members = ["app"]

[workspace.dependencies]
serde = "1.0.228"
`
	if err := os.WriteFile(filepath.Join(dir, "Cargo.toml"), []byte(root), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(dir, "app"), 0755); err != nil {
		t.Fatal(err)
	}
	member := `[package]
name = "app"

[dependencies.serde]
workspace = true
`
	if err := os.WriteFile(filepath.Join(dir, "app", "Cargo.toml"), []byte(member), 0644); err != nil {
		t.Fatal(err)
	}

	crates := extractCargoDeps(dir)
	byName := make(map[string]CargoCrate, len(crates))
	for _, c := range crates {
		byName[c.Name] = c
	}
	got, ok := byName["serde"]
	if !ok {
		t.Fatalf("missing serde dependency; got %+v", crates)
	}
	if got.Version != "1.0.228" {
		t.Errorf("inherited version: got %q, want 1.0.228", got.Version)
	}
	if got.Category != "" {
		t.Errorf("inherited dependency category: got %q, want \"\" (member-level)", got.Category)
	}
	if got.Source != filepath.Join("app", "Cargo.toml") {
		t.Errorf("inherited dependency source: got %q, want member manifest", got.Source)
	}
}

func TestMergeCrateEmptyVersionDoesNotReplace(t *testing.T) {
	seen := map[string]CargoCrate{
		"regex": {Name: "regex", Version: "1.10.3", Category: "workspace"},
	}
	mergeCrate(seen, CargoCrate{Name: "regex", Version: "", Category: ""})
	if seen["regex"].Version != "1.10.3" {
		t.Errorf("empty version replaced versioned entry: %+v", seen["regex"])
	}

	// A versioned normal dependency still takes precedence over the
	// workspace-level entry.
	mergeCrate(seen, CargoCrate{Name: "regex", Version: "1.10.3", Category: ""})
	if seen["regex"].Category != "" {
		t.Errorf("versioned normal dependency did not take precedence: %+v", seen["regex"])
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
