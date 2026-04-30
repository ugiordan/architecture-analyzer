package extractor

import (
	"testing"
)

func TestClassifyTier(t *testing.T) {
	tests := []struct {
		name     string
		expected ComponentTier
	}{
		{"opendatahub-operator", TierCore},
		{"odh-dashboard", TierCore},
		{"notebook-controller", TierCore},
		{"kserve", TierML},
		{"modelmesh-serving", TierML},
		{"codeflare-operator", TierML},
		{"ray-operator", TierML},
		{"data-science-pipelines-operator", TierData},
		{"trustyai-service", TierMonitoring},
		{"some-unknown-thing", TierUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tier := classifyTier(tt.name)
			if tier != tt.expected {
				t.Errorf("classifyTier(%q) = %s, want %s", tt.name, tier, tt.expected)
			}
		})
	}
}

func TestClassifyType(t *testing.T) {
	tests := []struct {
		name     string
		comp     KustomizeComponent
		expected ComponentType
	}{
		{
			"operator with CRDs",
			KustomizeComponent{Name: "my-operator", ManagedCRDs: []string{"Foo"}},
			TypeOperator,
		},
		{
			"controller without CRDs",
			KustomizeComponent{Name: "my-controller"},
			TypeController,
		},
		{
			"dashboard",
			KustomizeComponent{Name: "odh-dashboard"},
			TypeUI,
		},
		{
			"api server",
			KustomizeComponent{Name: "model-api-server"},
			TypeServer,
		},
		{
			"has CRDs only",
			KustomizeComponent{Name: "foo", ManagedCRDs: []string{"Bar"}},
			TypeOperator,
		},
		{
			"has images only",
			KustomizeComponent{Name: "foo", ImageParams: []ImageParam{{}}},
			TypeServer,
		},
		{
			"unknown",
			KustomizeComponent{Name: "foo"},
			TypeUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typ := classifyType(tt.comp)
			if typ != tt.expected {
				t.Errorf("classifyType(%q) = %s, want %s", tt.name, typ, tt.expected)
			}
		})
	}
}

func TestBuildComponentMap(t *testing.T) {
	discovery := &PlatformDiscovery{
		Components: []KustomizeComponent{
			{
				Name:        "kserve",
				ImageParams: []ImageParam{{}, {}, {}},
				ManagedCRDs: []string{"InferenceService", "ServingRuntime"},
			},
			{
				Name:         "odh-dashboard",
				ImageParams:  []ImageParam{{}},
				OverlayPaths: []string{"overlays/odh"},
			},
			{
				Name:        "data-science-pipelines-operator",
				ImageParams: []ImageParam{{}, {}},
				ManagedCRDs: []string{"DataSciencePipeline"},
			},
		},
	}

	cm := BuildComponentMap(discovery, "opendatahub-io")

	if len(cm.Components) != 3 {
		t.Fatalf("expected 3 components, got %d", len(cm.Components))
	}

	// Core should come first, then ML, then Data
	if cm.Components[0].Tier != TierCore {
		t.Errorf("expected first component to be core tier, got %s (%s)", cm.Components[0].Tier, cm.Components[0].Name)
	}

	if cm.Summary.TotalComponents != 3 {
		t.Errorf("expected 3 total components, got %d", cm.Summary.TotalComponents)
	}
	if cm.Summary.TotalImages != 6 {
		t.Errorf("expected 6 total images, got %d", cm.Summary.TotalImages)
	}
	if cm.Summary.TotalCRDs != 3 {
		t.Errorf("expected 3 total CRDs, got %d", cm.Summary.TotalCRDs)
	}

	// Check repo URL construction
	if cm.Components[0].Repo == "" {
		t.Error("expected repo URL to be set")
	}
}

func TestBuildComponentMapNil(t *testing.T) {
	cm := BuildComponentMap(nil, "")
	if len(cm.Components) != 0 {
		t.Errorf("expected 0 components for nil discovery, got %d", len(cm.Components))
	}
}

func TestBuildComponentMapEmpty(t *testing.T) {
	cm := BuildComponentMap(&PlatformDiscovery{}, "")
	if len(cm.Components) != 0 {
		t.Errorf("expected 0 components for empty discovery, got %d", len(cm.Components))
	}
}

func TestClassifyTierMLOverridesCore(t *testing.T) {
	// Components with both "operator" and an ML keyword should be ML, not core
	tests := []struct {
		name     string
		expected ComponentTier
	}{
		{"kserve-operator", TierML},
		{"codeflare-operator-controller", TierML},
		{"data-science-pipelines-operator", TierData},
	}
	for _, tt := range tests {
		tier := classifyTier(tt.name)
		if tier != tt.expected {
			t.Errorf("classifyTier(%q) = %s, want %s (specific tier should override core)", tt.name, tier, tt.expected)
		}
	}
}

func TestBuildComponentMapSortsCoreThenML(t *testing.T) {
	discovery := &PlatformDiscovery{
		Components: []KustomizeComponent{
			{Name: "kserve"},       // ml
			{Name: "odh-dashboard"}, // core
			{Name: "trustyai"},     // monitoring
		},
	}
	cm := BuildComponentMap(discovery, "")
	if cm.Components[0].Tier != TierCore {
		t.Errorf("first component should be core, got %s (%s)", cm.Components[0].Tier, cm.Components[0].Name)
	}
	if cm.Components[1].Tier != TierML {
		t.Errorf("second component should be ml, got %s (%s)", cm.Components[1].Tier, cm.Components[1].Name)
	}
	if cm.Components[2].Tier != TierMonitoring {
		t.Errorf("third component should be monitoring, got %s (%s)", cm.Components[2].Tier, cm.Components[2].Name)
	}
}

func TestBuildComponentMapNoOrg(t *testing.T) {
	discovery := &PlatformDiscovery{
		Components: []KustomizeComponent{{Name: "test"}},
	}
	cm := BuildComponentMap(discovery, "")
	if cm.Components[0].Repo != "" {
		t.Errorf("expected empty repo without org, got %q", cm.Components[0].Repo)
	}
}

func TestTierOrdering(t *testing.T) {
	if tierOrder(TierCore) >= tierOrder(TierML) {
		t.Error("core should come before ml")
	}
	if tierOrder(TierML) >= tierOrder(TierData) {
		t.Error("ml should come before data")
	}
	if tierOrder(TierData) >= tierOrder(TierMonitoring) {
		t.Error("data should come before monitoring")
	}
}
