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

func TestExtractOverlayPathsSourcePathLiteral(t *testing.T) {
	content := `
	return types.ManifestInfo{
		Path:       basePath,
		ContextDir: ComponentName,
		SourcePath: "overlays/odh",
	}
`
	paths := extractOverlayPaths(content)
	if len(paths) != 1 {
		t.Fatalf("expected 1 overlay path, got %d: %v", len(paths), paths)
	}
	if paths[0] != "overlays/odh" {
		t.Errorf("expected 'overlays/odh', got %q", paths[0])
	}
}

func TestExtractOverlayPathsFromMap(t *testing.T) {
	content := `
	overlaysSourcePaths = map[common.Platform]string{
		cluster.SelfManagedRhoai: "/rhoai",
		cluster.ManagedRhoai:     "/not-supported",
		cluster.OpenDataHub:      "/odh",
	}
`
	paths := extractOverlayPaths(content)
	if len(paths) != 2 {
		t.Fatalf("expected 2 paths (not-supported filtered), got %d: %v", len(paths), paths)
	}
}

func TestExtractOverlayPathsFromConstant(t *testing.T) {
	content := `
	kserveManifestSourcePath           = "overlays/odh"
	kserveManifestSourcePathXKS        = "overlays/odh-xks"
	kserveManifestSourcePathModelCache = "overlays/odh-modelcache"
`
	paths := extractOverlayPaths(content)
	if len(paths) != 3 {
		t.Fatalf("expected 3 paths from constants, got %d: %v", len(paths), paths)
	}
}

func TestExtractOverlayPathsFromManifestsSourcePathMap(t *testing.T) {
	content := `
	ManifestsSourcePath = map[common.Platform]string{
		cluster.SelfManagedRhoai: "overlays/rhoai",
		cluster.ManagedRhoai:     "overlays/rhoai",
		cluster.OpenDataHub:      "overlays/odh",
	}
`
	paths := extractOverlayPaths(content)
	if len(paths) != 2 {
		t.Fatalf("expected 2 deduped paths, got %d: %v", len(paths), paths)
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

	// Create a support file under a components/ directory
	supportDir := filepath.Join(dir, "internal", "controller", "components", "dashboard")
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
	supportDir := filepath.Join(dir, "internal", "controller", "components", "kserve")
	os.MkdirAll(supportDir, 0o755)

	// No GetComponentName method, name derived from filename
	content := `package kserve
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
	supportDir := filepath.Join(dir, "internal", "controller", "components", "ray")
	os.MkdirAll(supportDir, 0o755)

	content := `package ray
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

func TestExtractOverlayPathsRejectsGoImports(t *testing.T) {
	content := `
import (
	"sigs.k8s.io/kustomize/kyaml/yaml"
	"sigs.k8s.io/kustomize/api/builtins"
)

	SourcePath: "overlays/odh",
`
	paths := extractOverlayPaths(content)
	if len(paths) != 1 {
		t.Fatalf("expected 1 path (imports filtered), got %d: %v", len(paths), paths)
	}
	if paths[0] != "overlays/odh" {
		t.Errorf("expected 'overlays/odh', got %q", paths[0])
	}
}

func TestExtractOverlayPathsSkipsComments(t *testing.T) {
	content := `
	// SourcePath: "commented/out"
	/* SourcePath: "block/comment" */
	SourcePath: "real/path",
`
	paths := extractOverlayPaths(content)
	if len(paths) != 1 {
		t.Fatalf("expected 1 path (comments stripped), got %d: %v", len(paths), paths)
	}
	if paths[0] != "real/path" {
		t.Errorf("expected 'real/path', got %q", paths[0])
	}
}

func TestExtractOverlayPathsDedup(t *testing.T) {
	content := `
	SourcePath: "overlays/odh",
	SourcePath: "overlays/odh",
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
		{"internal/controller/components/dashboard", "dashboard"},
		{"internal/controller/components/kserve", "kserve"},
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

func TestFindSupportFilesSkipsNonComponents(t *testing.T) {
	dir := t.TempDir()

	// Create support files in non-component locations (should be skipped)
	for _, subdir := range []string{
		"pkg/controller/actions/deploy",
		"pkg/controller/conditions",
		"internal/controller/services/auth",
		"pkg/manifests/kustomize",
	} {
		d := filepath.Join(dir, subdir)
		os.MkdirAll(d, 0o755)
		os.WriteFile(filepath.Join(d, "action_deploy_support.go"), []byte(`package x`), 0o644)
	}

	// Create one in a valid components path (should be found)
	compDir := filepath.Join(dir, "internal", "controller", "components", "dashboard")
	os.MkdirAll(compDir, 0o755)
	os.WriteFile(filepath.Join(compDir, "dashboard_support.go"), []byte(`package dashboard
func (d *D) GetComponentName() string { return "dashboard" }
`), 0o644)

	pd, err := DiscoverPlatformComponents(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(pd.Components) != 1 {
		t.Fatalf("expected 1 component (non-components filtered), got %d", len(pd.Components))
	}
	if pd.Components[0].Name != "dashboard" {
		t.Errorf("expected 'dashboard', got %q", pd.Components[0].Name)
	}
}

func TestExtractContextDirsLiteral(t *testing.T) {
	content := `
	return types.ManifestInfo{
		Path:       basePath,
		ContextDir: "wva",
		SourcePath: "openshift",
	}
`
	dirs := extractContextDirs(content, "modelcontroller")
	if len(dirs) != 1 {
		t.Fatalf("expected 1 context dir, got %d: %v", len(dirs), dirs)
	}
	if dirs[0] != "wva" {
		t.Errorf("expected 'wva', got %q", dirs[0])
	}
}

func TestExtractContextDirsComponentName(t *testing.T) {
	content := `
	return types.ManifestInfo{
		Path:       basePath,
		ContextDir: ComponentName,
		SourcePath: "overlays/odh",
	}
`
	dirs := extractContextDirs(content, "dashboard")
	if len(dirs) != 1 {
		t.Fatalf("expected 1 context dir, got %d: %v", len(dirs), dirs)
	}
	if dirs[0] != "dashboard" {
		t.Errorf("expected 'dashboard', got %q", dirs[0])
	}
}

func TestExtractContextDirsMixed(t *testing.T) {
	content := `
	return types.ManifestInfo{
		ContextDir: ComponentName,
		SourcePath: "base",
	}
	return types.ManifestInfo{
		ContextDir: "wva",
		SourcePath: "openshift",
	}
`
	dirs := extractContextDirs(content, "modelcontroller")
	if len(dirs) != 2 {
		t.Fatalf("expected 2 context dirs, got %d: %v", len(dirs), dirs)
	}
}

func TestScansSiblingGoFiles(t *testing.T) {
	dir := t.TempDir()
	compDir := filepath.Join(dir, "internal", "controller", "components", "kserve")
	os.MkdirAll(compDir, 0o755)

	// Support file has component name and image params
	supportContent := `package kserve
func (k *K) GetComponentName() string { return "kserve" }
var imageParamMap = map[string]string{
	"kserve-image": "RELATED_IMAGE_KSERVE",
}
`
	os.WriteFile(filepath.Join(compDir, "kserve_support.go"), []byte(supportContent), 0o644)

	// Main file has source path constants
	mainContent := `package kserve
	kserveManifestSourcePath           = "overlays/odh"
	kserveManifestSourcePathXKS        = "overlays/odh-xks"
`
	os.WriteFile(filepath.Join(compDir, "kserve.go"), []byte(mainContent), 0o644)

	// Test file should be ignored
	testContent := `package kserve
	testSourcePath = "test/path"
`
	os.WriteFile(filepath.Join(compDir, "kserve_test.go"), []byte(testContent), 0o644)

	pd, err := DiscoverPlatformComponents(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(pd.Components) != 1 {
		t.Fatalf("expected 1 component, got %d", len(pd.Components))
	}
	comp := pd.Components[0]
	if len(comp.OverlayPaths) != 2 {
		t.Fatalf("expected 2 overlay paths from sibling file, got %d: %v", len(comp.OverlayPaths), comp.OverlayPaths)
	}
}

func TestFullComponentExtraction(t *testing.T) {
	dir := t.TempDir()

	type testComponent struct {
		name             string
		supportContent   string
		siblingContent   string
		wantOverlays     int
		wantContextDirs  int
		wantImageParams  int
	}

	components := []testComponent{
		{
			name: "dashboard",
			supportContent: `package dashboard
func (d *D) GetComponentName() string { return "dashboard" }

var imageParamMap = map[string]string{
	"dashboard-image": "RELATED_IMAGE_DASHBOARD",
}

overlaysSourcePaths = map[common.Platform]string{
	cluster.SelfManagedRhoai: "/rhoai",
	cluster.OpenDataHub:      "/odh",
}

func defaultManifestInfo(basePath string, p common.Platform) types.ManifestInfo {
	return types.ManifestInfo{
		Path:       basePath,
		ContextDir: ComponentName,
		SourcePath: overlaysSourcePaths[p],
	}
}
`,
			wantOverlays:    2,
			wantContextDirs: 1,
			wantImageParams: 1,
		},
		{
			name: "kserve",
			supportContent: `package kserve
import (
	"sigs.k8s.io/kustomize/kyaml/yaml"
)
func kserveManifestInfo(basePath string, sourcePath string) odhtypes.ManifestInfo {
	return odhtypes.ManifestInfo{
		Path:       basePath,
		ContextDir: componentName,
		SourcePath: sourcePath,
	}
}
`,
			siblingContent: `package kserve
	componentName = componentApi.KserveComponentName
	kserveManifestSourcePath           = "overlays/odh"
	kserveManifestSourcePathXKS        = "overlays/odh-xks"
	kserveManifestSourcePathModelCache = "overlays/odh-modelcache"
`,
			wantOverlays:    3,
			wantContextDirs: 1,
			wantImageParams: 0,
		},
		{
			name: "modelcontroller",
			supportContent: `package modelcontroller
func (m *M) GetComponentName() string { return "modelcontroller" }
func mainManifestInfo(basePath string) types.ManifestInfo {
	return types.ManifestInfo{
		ContextDir: ComponentName,
		SourcePath: "base",
	}
}
func wvaManifestInfo(basePath string) types.ManifestInfo {
	return types.ManifestInfo{
		ContextDir: "wva",
		SourcePath: "openshift",
	}
}
`,
			wantOverlays:    2,
			wantContextDirs: 2,
			wantImageParams: 0,
		},
		{
			name: "kueue",
			supportContent: `package kueue
func (k *K) GetComponentName() string { return "kueue" }
`,
			wantOverlays:    0,
			wantContextDirs: 0,
			wantImageParams: 0,
		},
		{
			name: "workbenches",
			supportContent: `package workbenches
func (w *W) GetComponentName() string { return "workbenches" }

const (
	notebooksPath = "notebooks"
	notebookControllerPath = "odh-notebook-controller"
	kfNotebookControllerPath = "kf-notebook-controller"
)

var (
	notebookControllerContextDir   = path.Join(ComponentName, notebookControllerPath)
	kfNotebookControllerContextDir = path.Join(ComponentName, kfNotebookControllerPath)
	notebookContextDir             = path.Join(ComponentName, notebooksPath)

	notebookImagesManifestSourcePath = map[common.Platform]string{
		cluster.SelfManagedRhoai: "rhoai/overlays/additional",
		cluster.OpenDataHub:      "odh/overlays/additional",
	}
)

func notebookControllerManifestInfo(basePath string, sourcePath string) odhtypes.ManifestInfo {
	return odhtypes.ManifestInfo{
		Path:       basePath,
		ContextDir: notebookControllerContextDir,
		SourcePath: sourcePath,
	}
}
`,
			siblingContent: `package workbenches
if err := odhdeploy.ApplyParams(nbcManifestInfo.String(), "params.env", map[string]string{
	"odh-notebook-controller-image": "RELATED_IMAGE_ODH_NOTEBOOK_CONTROLLER_IMAGE",
	"kube-rbac-proxy":               "RELATED_IMAGE_ODH_KUBE_RBAC_PROXY_IMAGE",
}); err != nil {
	return err
}
if err := odhdeploy.ApplyParams(kfnbcManifestInfo.String(), "params.env", map[string]string{
	"odh-kf-notebook-controller-image": "RELATED_IMAGE_ODH_KF_NOTEBOOK_CONTROLLER_IMAGE",
}); err != nil {
	return err
}
`,
			wantOverlays:    2,
			wantContextDirs: 3,
			wantImageParams: 3,
		},
	}

	// Also create non-component files that should be filtered out
	for _, nonComp := range []string{
		"pkg/controller/actions/deploy",
		"internal/controller/services/auth",
	} {
		d := filepath.Join(dir, nonComp)
		os.MkdirAll(d, 0o755)
		os.WriteFile(filepath.Join(d, "action_support.go"), []byte(`package x
func (a *A) GetComponentName() string { return "should-not-appear" }
`), 0o644)
	}

	for _, comp := range components {
		compDir := filepath.Join(dir, "internal", "controller", "components", comp.name)
		os.MkdirAll(compDir, 0o755)
		os.WriteFile(filepath.Join(compDir, comp.name+"_support.go"), []byte(comp.supportContent), 0o644)
		if comp.siblingContent != "" {
			os.WriteFile(filepath.Join(compDir, comp.name+".go"), []byte(comp.siblingContent), 0o644)
		}
	}

	pd, err := DiscoverPlatformComponents(dir)
	if err != nil {
		t.Fatal(err)
	}

	// Verify correct component count (no non-components)
	if len(pd.Components) != len(components) {
		t.Fatalf("expected %d components, got %d", len(components), len(pd.Components))
	}

	// Build a map for easy lookup
	compMap := make(map[string]KustomizeComponent)
	for _, c := range pd.Components {
		compMap[c.Name] = c
	}

	for _, tc := range components {
		c, ok := compMap[tc.name]
		if !ok {
			t.Errorf("component %q not found", tc.name)
			continue
		}
		if len(c.OverlayPaths) != tc.wantOverlays {
			t.Errorf("%s: expected %d overlay paths, got %d: %v",
				tc.name, tc.wantOverlays, len(c.OverlayPaths), c.OverlayPaths)
		}
		if len(c.ContextDirs) != tc.wantContextDirs {
			t.Errorf("%s: expected %d context dirs, got %d: %v",
				tc.name, tc.wantContextDirs, len(c.ContextDirs), c.ContextDirs)
		}
		if len(c.ImageParams) != tc.wantImageParams {
			t.Errorf("%s: expected %d image params, got %d: %v",
				tc.name, tc.wantImageParams, len(c.ImageParams), c.ImageParams)
		}
	}

	// Verify non-components were filtered
	if _, ok := compMap["should-not-appear"]; ok {
		t.Error("non-component 'should-not-appear' should have been filtered out")
	}
}

func TestContextDirPathJoin(t *testing.T) {
	content := `
	notebooksPath = "notebooks"
	notebookControllerPath = "odh-notebook-controller"
	kfNotebookControllerPath = "kf-notebook-controller"

	notebookControllerContextDir   = path.Join(ComponentName, notebookControllerPath)
	kfNotebookControllerContextDir = path.Join(ComponentName, kfNotebookControllerPath)
	notebookContextDir             = path.Join(ComponentName, notebooksPath)
`
	dirs := extractContextDirs(content, "workbenches")
	if len(dirs) != 3 {
		t.Fatalf("expected 3 context dirs from path.Join, got %d: %v", len(dirs), dirs)
	}
	expected := map[string]bool{
		"workbenches/odh-notebook-controller": true,
		"workbenches/kf-notebook-controller":  true,
		"workbenches/notebooks":               true,
	}
	for _, d := range dirs {
		if !expected[d] {
			t.Errorf("unexpected context dir: %q", d)
		}
	}
}

func TestImageParamsFromSiblingFiles(t *testing.T) {
	dir := t.TempDir()
	compDir := filepath.Join(dir, "internal", "controller", "components", "workbenches")
	os.MkdirAll(compDir, 0o755)

	supportContent := `package workbenches
func (w *W) GetComponentName() string { return "workbenches" }
`
	os.WriteFile(filepath.Join(compDir, "workbenches_support.go"), []byte(supportContent), 0o644)

	mainContent := `package workbenches
if err := odhdeploy.ApplyParams(info.String(), "params.env", map[string]string{
	"img-a": "RELATED_IMAGE_A",
	"img-b": "RELATED_IMAGE_B",
	"img-c": "RELATED_IMAGE_C",
}); err != nil {
	return err
}
`
	os.WriteFile(filepath.Join(compDir, "workbenches.go"), []byte(mainContent), 0o644)

	pd, err := DiscoverPlatformComponents(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(pd.Components) != 1 {
		t.Fatalf("expected 1 component, got %d", len(pd.Components))
	}
	if len(pd.Components[0].ImageParams) != 3 {
		t.Errorf("expected 3 image params from sibling file, got %d: %v",
			len(pd.Components[0].ImageParams), pd.Components[0].ImageParams)
	}
}

func TestIsValidOverlayPath(t *testing.T) {
	tests := []struct {
		path string
		want bool
	}{
		{"overlays/odh", true},
		{"/rhoai", true},
		{"base", true},
		{"sigs.k8s.io/kustomize/api/types", false},
		{"github.com/foo/bar", false},
		{"example.org/path", false},
		{"not-supported", false},
		{"/not-supported", false},
	}
	for _, tt := range tests {
		got := isValidOverlayPath(tt.path)
		if got != tt.want {
			t.Errorf("isValidOverlayPath(%q) = %v, want %v", tt.path, got, tt.want)
		}
	}
}

func TestIsComponentPath(t *testing.T) {
	tests := []struct {
		path string
		want bool
	}{
		{"internal/controller/components/dashboard/dashboard_support.go", true},
		{"pkg/controller/actions/deploy/action_deploy_support.go", false},
		{"internal/controller/services/auth/auth_support.go", false},
		{"pkg/manifests/kustomize/kustomize_support.go", false},
	}
	for _, tt := range tests {
		got := isComponentPath(tt.path)
		if got != tt.want {
			t.Errorf("isComponentPath(%q) = %v, want %v", tt.path, got, tt.want)
		}
	}
}
