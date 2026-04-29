package sarif

import "testing"

func TestNormalizedSeverity(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"error", "high"},
		{"warning", "medium"},
		{"note", "low"},
		{"none", "informational"},
		{"", "informational"},
		{"critical", "critical"},
		{"high", "high"},
		{"medium", "medium"},
		{"low", "low"},
		{"informational", "informational"},
		{"unknown-level", "medium"},
	}
	for _, tt := range tests {
		got := NormalizedSeverity(tt.input)
		if got != tt.want {
			t.Errorf("NormalizedSeverity(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestCategoryForCWEs(t *testing.T) {
	tests := []struct {
		name string
		cwes []string
		want string
	}{
		{"sql injection", []string{"CWE-89"}, "injection"},
		{"xss", []string{"CWE-79"}, "injection"},
		{"auth bypass", []string{"CWE-287"}, "auth"},
		{"multiple CWEs highest priority wins", []string{"CWE-798", "CWE-89"}, "hardcoded-credentials"},
		{"unknown CWE", []string{"CWE-999"}, ""},
		{"empty", nil, ""},
		{"mixed known and unknown", []string{"CWE-999", "CWE-327"}, "crypto"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CategoryForCWEs(tt.cwes)
			if got != tt.want {
				t.Errorf("CategoryForCWEs(%v) = %q, want %q", tt.cwes, got, tt.want)
			}
		})
	}
}
