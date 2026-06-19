package extractor

import (
	"sort"
	"testing"
)

func TestResolveExtractorGroups_NilReturnsAll(t *testing.T) {
	active, err := resolveExtractorGroups(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(active) != len(allGroups) {
		t.Errorf("got %d groups, want %d", len(active), len(allGroups))
	}
	for _, g := range allGroups {
		if !active[g.name] {
			t.Errorf("group %q missing from active set", g.name)
		}
	}
}

func TestResolveExtractorGroups_EmptyReturnsAll(t *testing.T) {
	active, err := resolveExtractorGroups([]string{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(active) != len(allGroups) {
		t.Errorf("got %d groups, want %d", len(active), len(allGroups))
	}
}

func TestResolveExtractorGroups_SingleGroup(t *testing.T) {
	active, err := resolveExtractorGroups([]string{"docker"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !active["docker"] {
		t.Error("docker should be active")
	}
	if active["helm"] {
		t.Error("helm should not be active")
	}
}

func TestResolveExtractorGroups_DependencyExpansion(t *testing.T) {
	active, err := resolveExtractorGroups([]string{"controllers"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !active["controllers"] {
		t.Error("controllers should be active")
	}
	if !active["k8s_core"] {
		t.Error("k8s_core should be auto-included as dependency of controllers")
	}
	if active["docker"] {
		t.Error("docker should not be active")
	}
}

func TestResolveExtractorGroups_OperatorDependency(t *testing.T) {
	active, err := resolveExtractorGroups([]string{"operator"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !active["operator"] {
		t.Error("operator should be active")
	}
	if !active["observability"] {
		t.Error("observability should be auto-included as dependency of operator")
	}
}

func TestResolveExtractorGroups_UnknownGroupError(t *testing.T) {
	_, err := resolveExtractorGroups([]string{"nonexistent"})
	if err == nil {
		t.Fatal("expected error for unknown group")
	}
}

func TestResolveExtractorGroups_MultipleGroups(t *testing.T) {
	active, err := resolveExtractorGroups([]string{"docker", "kustomize"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !active["docker"] || !active["kustomize"] {
		t.Error("both docker and kustomize should be active")
	}
	if len(active) != 2 {
		t.Errorf("got %d active groups, want 2", len(active))
	}
}

func TestResolveExtractorGroups_TrimsWhitespace(t *testing.T) {
	active, err := resolveExtractorGroups([]string{" docker ", " helm"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !active["docker"] || !active["helm"] {
		t.Error("whitespace-trimmed groups should be active")
	}
}

func TestShouldRun(t *testing.T) {
	active := map[string]bool{"docker": true, "helm": true}
	if !shouldRun(active, "docker") {
		t.Error("shouldRun should return true for active group")
	}
	if shouldRun(active, "crds") {
		t.Error("shouldRun should return false for inactive group")
	}
}

func TestExtractorGroupNames_Sorted(t *testing.T) {
	names := ExtractorGroupNames()
	if len(names) != len(allGroups) {
		t.Errorf("got %d names, want %d", len(names), len(allGroups))
	}
	if !sort.StringsAreSorted(names) {
		t.Error("names should be sorted")
	}
}

func TestNeedsGoPackages(t *testing.T) {
	tests := []struct {
		active map[string]bool
		want   bool
	}{
		{map[string]bool{"docker": true}, false},
		{map[string]bool{"crds": true}, true},
		{map[string]bool{"webhooks": true}, true},
		{map[string]bool{"controllers": true}, true},
		{map[string]bool{"docker": true, "operator": true}, true},
		{map[string]bool{"helm": true, "kustomize": true}, false},
	}
	for _, tt := range tests {
		got := needsGoPackages(tt.active)
		if got != tt.want {
			t.Errorf("needsGoPackages(%v) = %v, want %v", tt.active, got, tt.want)
		}
	}
}
