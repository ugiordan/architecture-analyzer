package extractor

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractComponentName(t *testing.T) {
	content := `
func (d *Dashboard) GetComponentName() string {
	return "dashboard"
}
`
	name := extractComponentName(content)
	if name != "dashboard" {
		t.Errorf("expected 'dashboard', got %q", name)
	}
}

func TestExtractComponentNameNotFound(t *testing.T) {
	content := `func (d *Dashboard) SomeOtherMethod() string { return "foo" }`
	name := extractComponentName(content)
	if name != "" {
		t.Errorf("expected empty string, got %q", name)
	}
}

func TestExtractImageParams(t *testing.T) {
	content := `
var imageParamMap = map[string]string{
	"odh-dashboard-image": "RELATED_IMAGE_ODH_DASHBOARD",
	"oauth-proxy":         "RELATED_IMAGE_OAUTH_PROXY",
}
`
	params := extractImageParams(content)
	if len(params) != 2 {
		t.Fatalf("expected 2 params, got %d", len(params))
	}
	// Check one specific entry
	found := false
	for _, p := range params {
		if p.ParamsKey == "odh-dashboard-image" && p.EnvVar == "RELATED_IMAGE_ODH_DASHBOARD" {
			found = true
		}
	}
	if !found {
		t.Error("expected to find 'odh-dashboard-image' -> 'RELATED_IMAGE_ODH_DASHBOARD'")
	}
}

func TestExtractImageParamsDedup(t *testing.T) {
	content := `
	"my-image": "RELATED_IMAGE_FOO",
	"my-image": "RELATED_IMAGE_FOO",
`
	params := extractImageParams(content)
	if len(params) != 1 {
		t.Errorf("expected 1 deduped param, got %d", len(params))
	}
}

func TestExtractOverlayPaths(t *testing.T) {
	content := `
	overlayDir := "opt/manifests/dashboard/overlays/odh"
	rhoaiPath := "opt/manifests/dashboard/overlays/rhoai"
`
	paths := extractOverlayPaths(content)
	if len(paths) != 2 {
		t.Fatalf("expected 2 overlay paths, got %d", len(paths))
	}
}

func TestExtractFeatureFlags(t *testing.T) {
	content := `
if features.ModelServing {
	doSomething()
}
if features.KServeEnabled {
	doSomethingElse()
}
`
	flags := extractFeatureFlags(content)
	if len(flags) != 2 {
		t.Fatalf("expected 2 feature flags, got %d", len(flags))
	}
}

func TestExtractFeatureFlagsSkipsFalsePositives(t *testing.T) {
	content := `
var x features.Feature
x = features.FeatureGate
y := features.String()
`
	flags := extractFeatureFlags(content)
	if len(flags) != 0 {
		t.Errorf("expected 0 feature flags (all false positives), got %d: %v", len(flags), flags)
	}
}

func TestParseParamsEnv(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "params.env")
	content := `# Comment
odh-dashboard-image=quay.io/opendatahub/dashboard:latest
oauth-proxy=quay.io/openshift/oauth-proxy:v4.15

# Another comment
empty=
`
	os.WriteFile(path, []byte(content), 0o644)

	pe, err := parseParamsEnv(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(pe.Params) != 3 {
		t.Fatalf("expected 3 params, got %d", len(pe.Params))
	}
	if pe.Params["odh-dashboard-image"] != "quay.io/opendatahub/dashboard:latest" {
		t.Errorf("unexpected value: %q", pe.Params["odh-dashboard-image"])
	}
	if pe.Params["empty"] != "" {
		t.Errorf("expected empty value, got %q", pe.Params["empty"])
	}
}

func TestResolveImageDefaults(t *testing.T) {
	components := []KustomizeComponent{
		{
			Name: "dashboard",
			ImageParams: []ImageParam{
				{EnvVar: "RELATED_IMAGE_DASH", ParamsKey: "dash-image"},
				{EnvVar: "RELATED_IMAGE_OAUTH", ParamsKey: "oauth-image"},
			},
		},
	}
	params := &ParamsEnv{
		Params: map[string]string{
			"dash-image": "quay.io/opendatahub/dashboard:latest",
		},
	}

	resolveImageDefaults(components, params)

	if components[0].ImageParams[0].DefaultImage != "quay.io/opendatahub/dashboard:latest" {
		t.Errorf("expected resolved image, got %q", components[0].ImageParams[0].DefaultImage)
	}
	if components[0].ImageParams[1].DefaultImage != "" {
		t.Errorf("expected empty default for unresolved key, got %q", components[0].ImageParams[1].DefaultImage)
	}
}

func TestResolveImageDefaultsNilParams(t *testing.T) {
	components := []KustomizeComponent{
		{Name: "test", ImageParams: []ImageParam{{EnvVar: "X", ParamsKey: "y"}}},
	}
	resolveImageDefaults(components, nil)
	if components[0].ImageParams[0].DefaultImage != "" {
		t.Error("expected empty default with nil params")
	}
}

func TestDiscoverPlatformComponents(t *testing.T) {
	dir := t.TempDir()

	// Create a support file
	supportDir := filepath.Join(dir, "components", "dashboard")
	os.MkdirAll(supportDir, 0o755)

	supportContent := `package components

func (d *Dashboard) GetComponentName() string {
	return "dashboard"
}

var imageParamMap = map[string]string{
	"dashboard-image": "RELATED_IMAGE_DASHBOARD",
}
`
	os.WriteFile(filepath.Join(supportDir, "dashboard_support.go"), []byte(supportContent), 0o644)

	// Create a params.env
	configDir := filepath.Join(dir, "config")
	os.MkdirAll(configDir, 0o755)
	os.WriteFile(filepath.Join(configDir, "params.env"), []byte("dashboard-image=quay.io/test:v1\n"), 0o644)

	pd, err := DiscoverPlatformComponents(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(pd.Components) != 1 {
		t.Fatalf("expected 1 component, got %d", len(pd.Components))
	}
	if pd.Components[0].Name != "dashboard" {
		t.Errorf("expected name 'dashboard', got %q", pd.Components[0].Name)
	}
	if len(pd.Components[0].ImageParams) != 1 {
		t.Fatalf("expected 1 image param, got %d", len(pd.Components[0].ImageParams))
	}
	if pd.Components[0].ImageParams[0].DefaultImage != "quay.io/test:v1" {
		t.Errorf("expected resolved default image, got %q", pd.Components[0].ImageParams[0].DefaultImage)
	}
}

func TestDiscoverPlatformComponentsNoSupportFiles(t *testing.T) {
	dir := t.TempDir()
	pd, err := DiscoverPlatformComponents(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(pd.Components) != 0 {
		t.Errorf("expected 0 components, got %d", len(pd.Components))
	}
}

func TestExtractComponentNameFallbackFromFilename(t *testing.T) {
	dir := t.TempDir()
	supportDir := filepath.Join(dir, "components")
	os.MkdirAll(supportDir, 0o755)

	// No GetComponentName method, name derived from filename
	content := `package components
var imageParamMap = map[string]string{
	"img": "RELATED_IMAGE_FOO",
}
`
	os.WriteFile(filepath.Join(supportDir, "kserve_support.go"), []byte(content), 0o644)

	pd, err := DiscoverPlatformComponents(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(pd.Components) != 1 {
		t.Fatalf("expected 1 component, got %d", len(pd.Components))
	}
	if pd.Components[0].Name != "kserve" {
		t.Errorf("expected fallback name 'kserve', got %q", pd.Components[0].Name)
	}
}

func TestExtractComponentNameFallbackComponentSuffix(t *testing.T) {
	dir := t.TempDir()
	supportDir := filepath.Join(dir, "pkg")
	os.MkdirAll(supportDir, 0o755)

	content := `package pkg
var x = 1
`
	os.WriteFile(filepath.Join(supportDir, "ray_component.go"), []byte(content), 0o644)

	pd, err := DiscoverPlatformComponents(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(pd.Components) != 1 {
		t.Fatalf("expected 1 component, got %d", len(pd.Components))
	}
	if pd.Components[0].Name != "ray" {
		t.Errorf("expected fallback name 'ray', got %q", pd.Components[0].Name)
	}
}

func TestDiscoverSkipsVendorAndGitDirs(t *testing.T) {
	dir := t.TempDir()

	// Create support files in vendor and .git (should be skipped)
	for _, subdir := range []string{"vendor/pkg", ".git/hooks"} {
		d := filepath.Join(dir, subdir)
		os.MkdirAll(d, 0o755)
		os.WriteFile(filepath.Join(d, "fake_support.go"), []byte(`package x`), 0o644)
	}

	pd, err := DiscoverPlatformComponents(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(pd.Components) != 0 {
		t.Errorf("expected 0 components (vendor/git skipped), got %d", len(pd.Components))
	}
}

func TestExtractImageParamsNoMatches(t *testing.T) {
	content := `var x = map[string]int{"foo": 42}`
	params := extractImageParams(content)
	if len(params) != 0 {
		t.Errorf("expected 0 params for non-RELATED_IMAGE content, got %d", len(params))
	}
}

func TestExtractOverlayPathsDedup(t *testing.T) {
	content := `
	p1 := "config/overlays/odh"
	p2 := "config/overlays/odh"
`
	paths := extractOverlayPaths(content)
	if len(paths) != 1 {
		t.Errorf("expected 1 deduped path, got %d", len(paths))
	}
}

func TestParseParamsEnvWithEqualsInValue(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "params.env")
	// Values can contain = (e.g. base64 encoded strings)
	content := "my-key=value=with=equals\n"
	os.WriteFile(path, []byte(content), 0o644)

	pe, err := parseParamsEnv(path)
	if err != nil {
		t.Fatal(err)
	}
	if pe.Params["my-key"] != "value=with=equals" {
		t.Errorf("expected 'value=with=equals', got %q", pe.Params["my-key"])
	}
}

func TestParseParamsEnvMissingFile(t *testing.T) {
	_, err := parseParamsEnv("/nonexistent/params.env")
	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestDiscoverMultipleSupportFiles(t *testing.T) {
	dir := t.TempDir()

	for _, comp := range []struct{ dir, name string }{
		{"components/dashboard", "dashboard"},
		{"components/kserve", "kserve"},
	} {
		d := filepath.Join(dir, comp.dir)
		os.MkdirAll(d, 0o755)
		content := `package components
func (c *C) GetComponentName() string { return "` + comp.name + `" }
`
		os.WriteFile(filepath.Join(d, comp.name+"_support.go"), []byte(content), 0o644)
	}

	pd, err := DiscoverPlatformComponents(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(pd.Components) != 2 {
		t.Fatalf("expected 2 components, got %d", len(pd.Components))
	}
	// Should be sorted alphabetically
	if pd.Components[0].Name != "dashboard" {
		t.Errorf("expected 'dashboard' first (sorted), got %q", pd.Components[0].Name)
	}
	if pd.Components[1].Name != "kserve" {
		t.Errorf("expected 'kserve' second, got %q", pd.Components[1].Name)
	}
}

func TestFormatSummaryEmpty(t *testing.T) {
	pd := &PlatformDiscovery{}
	summary := pd.FormatSummary()
	if summary != "No kustomize components discovered." {
		t.Errorf("unexpected empty summary: %q", summary)
	}
}

func TestPlatformDiscoverySummary(t *testing.T) {
	pd := &PlatformDiscovery{
		Components: []KustomizeComponent{
			{Name: "dashboard", ImageParams: []ImageParam{{}, {}}},
			{Name: "kserve", ImageParams: []ImageParam{{}}},
		},
	}
	names := pd.ComponentNames()
	if len(names) != 2 {
		t.Errorf("expected 2 names, got %d", len(names))
	}
	if pd.TotalImageParams() != 3 {
		t.Errorf("expected 3 total image params, got %d", pd.TotalImageParams())
	}
	summary := pd.FormatSummary()
	if summary == "" {
		t.Error("expected non-empty summary")
	}
}
