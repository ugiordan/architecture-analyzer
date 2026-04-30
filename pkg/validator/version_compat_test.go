package validator

import (
	"testing"
)

func TestParseKubeVersion(t *testing.T) {
	tests := []struct {
		input string
		major int
		minor int
	}{
		{"1.27", 1, 27},
		{"v1.25.3", 1, 25},
		{"4.14", 4, 14},
	}
	for _, tt := range tests {
		v, err := ParseKubeVersion(tt.input)
		if err != nil {
			t.Errorf("ParseKubeVersion(%q) error: %v", tt.input, err)
			continue
		}
		if v.Major != tt.major || v.Minor != tt.minor {
			t.Errorf("ParseKubeVersion(%q) = %d.%d, want %d.%d", tt.input, v.Major, v.Minor, tt.major, tt.minor)
		}
	}
}

func TestParseKubeVersionInvalid(t *testing.T) {
	invalids := []string{"", "1", "abc", "1.abc"}
	for _, s := range invalids {
		_, err := ParseKubeVersion(s)
		if err == nil {
			t.Errorf("expected error for %q", s)
		}
	}
}

func TestOCPToKubeVersion(t *testing.T) {
	tests := []struct {
		ocp          string
		expectedKube string
	}{
		{"4.14", "1.27"},
		{"4.15", "1.28"},
		{"4.16", "1.29"},
		{"4.12", "1.25"},
	}
	for _, tt := range tests {
		v, err := OCPToKubeVersion(tt.ocp)
		if err != nil {
			t.Errorf("OCPToKubeVersion(%q) error: %v", tt.ocp, err)
			continue
		}
		if v.String() != tt.expectedKube {
			t.Errorf("OCPToKubeVersion(%q) = %s, want %s", tt.ocp, v.String(), tt.expectedKube)
		}
	}
}

func TestKubeVersionAtLeast(t *testing.T) {
	v127 := KubeVersion{1, 27}
	v125 := KubeVersion{1, 25}
	v200 := KubeVersion{2, 0}

	if !v127.AtLeast(v125) {
		t.Error("1.27 should be >= 1.25")
	}
	if v125.AtLeast(v127) {
		t.Error("1.25 should not be >= 1.27")
	}
	if !v127.AtLeast(v127) {
		t.Error("1.27 should be >= 1.27")
	}
	if !v200.AtLeast(v127) {
		t.Error("2.0 should be >= 1.27")
	}
}

func TestCheckVersionCompatCRDv1beta1(t *testing.T) {
	arch := map[string]interface{}{
		"crds": []interface{}{
			map[string]interface{}{
				"group":   "apiextensions.k8s.io",
				"version": "v1beta1",
				"kind":    "CustomResourceDefinition",
				"source":  "config/crd/bases/foo.yaml",
			},
		},
	}

	// Target OCP 4.14 (kube 1.27): v1beta1 was removed in 1.22
	result, err := CheckVersionCompat(arch, "4.14")
	if err != nil {
		t.Fatal(err)
	}
	if result.Compatible {
		t.Error("expected incompatible: CRD v1beta1 removed in 1.22, target is 1.27")
	}
	if len(result.Issues) == 0 {
		t.Fatal("expected at least one issue")
	}
	if result.Issues[0].Severity != "error" {
		t.Errorf("expected severity 'error', got %q", result.Issues[0].Severity)
	}
}

func TestCheckVersionCompatCRDv1(t *testing.T) {
	arch := map[string]interface{}{
		"crds": []interface{}{
			map[string]interface{}{
				"group":   "apiextensions.k8s.io",
				"version": "v1",
				"kind":    "CustomResourceDefinition",
				"source":  "config/crd/bases/foo.yaml",
			},
		},
	}

	result, err := CheckVersionCompat(arch, "4.14")
	if err != nil {
		t.Fatal(err)
	}
	if !result.Compatible {
		t.Error("expected compatible: CRD v1 is fine")
	}
}

func TestCheckVersionCompatIngressV1beta1(t *testing.T) {
	arch := map[string]interface{}{
		"ingress_routing": []interface{}{
			map[string]interface{}{
				"kind":   "Ingress",
				"name":   "my-ingress",
				"source": "config/networking.k8s.io/v1beta1/ingress.yaml",
			},
		},
	}

	result, err := CheckVersionCompat(arch, "4.14")
	if err != nil {
		t.Fatal(err)
	}
	if result.Compatible {
		t.Error("expected incompatible: Ingress v1beta1 removed in 1.22")
	}
}

func TestCheckVersionCompatEmpty(t *testing.T) {
	result, err := CheckVersionCompat(map[string]interface{}{}, "4.14")
	if err != nil {
		t.Fatal(err)
	}
	if !result.Compatible {
		t.Error("expected compatible for empty arch data")
	}
	if result.KubeVersion != "1.27" {
		t.Errorf("expected kube version 1.27, got %q", result.KubeVersion)
	}
}

func TestCheckVersionCompatKubeVersionDirect(t *testing.T) {
	// Use kube version directly instead of OCP
	result, err := CheckVersionCompat(map[string]interface{}{}, "1.27")
	if err != nil {
		t.Fatal(err)
	}
	if result.KubeVersion != "1.27" {
		t.Errorf("expected kube version 1.27, got %q", result.KubeVersion)
	}
}

func TestCheckVersionCompatHPAv2beta2(t *testing.T) {
	arch := map[string]interface{}{
		"crds": []interface{}{
			map[string]interface{}{
				"group":   "autoscaling",
				"version": "v2beta2",
				"kind":    "HorizontalPodAutoscaler",
				"source":  "config/hpa.yaml",
			},
		},
	}

	// OCP 4.13 -> kube 1.26: v2beta2 removed in 1.26
	result, err := CheckVersionCompat(arch, "4.13")
	if err != nil {
		t.Fatal(err)
	}
	if result.Compatible {
		t.Error("expected incompatible: HPA v2beta2 removed in 1.26")
	}
}

func TestCheckVersionCompatDeprecatedButNotRemoved(t *testing.T) {
	// Test with a version where API is deprecated but not yet removed
	arch := map[string]interface{}{
		"crds": []interface{}{
			map[string]interface{}{
				"group":   "apiextensions.k8s.io",
				"version": "v1beta1",
				"kind":    "CustomResourceDefinition",
				"source":  "test.yaml",
			},
		},
	}

	// Kube 1.17: v1beta1 deprecated in 1.16 but not removed until 1.22
	result, err := CheckVersionCompat(arch, "1.17")
	if err != nil {
		t.Fatal(err)
	}
	// Should have warning but still be compatible
	if !result.Compatible {
		t.Error("expected compatible (deprecated but not removed)")
	}
	if len(result.Issues) == 0 {
		t.Error("expected deprecation warning")
	}
	if result.Issues[0].Severity != "warning" {
		t.Errorf("expected severity 'warning', got %q", result.Issues[0].Severity)
	}
}

func TestCheckVersionCompatBadVersion(t *testing.T) {
	_, err := CheckVersionCompat(map[string]interface{}{}, "invalid")
	if err == nil {
		t.Error("expected error for invalid version")
	}
}
