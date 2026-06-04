package flow

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"testing"
)

// testDiagram returns a minimal Diagram suitable for HTML generation tests.
func testDiagram() Diagram {
	return Diagram{
		Meta:   DiagramMeta{Title: "Test Diagram"},
		Canvas: DiagramCanvas{Width: 800, Height: 600},
		Nodes: map[string]DiagramNode{
			"a": {X: 100, Y: 100, Type: "box", Label: "Alpha"},
		},
		Flows: map[string]DiagramFlow{
			"f1": {Label: "flow-1", Steps: []ArrowStep{
				{Mode: "arrow", From: "a", To: "a", Num: 1},
			}},
		},
	}
}

func TestGenerateHTML_ContainsCSP(t *testing.T) {
	html, err := GenerateHTML(testDiagram())
	if err != nil {
		t.Fatalf("GenerateHTML: %v", err)
	}
	if !strings.Contains(html, "Content-Security-Policy") {
		t.Error("HTML missing Content-Security-Policy meta tag")
	}
	if !strings.Contains(html, "sha256-") {
		t.Error("CSP does not contain sha256 hashes")
	}
}

func TestGenerateHTML_ContainsFlowlensBundle(t *testing.T) {
	html, err := GenerateHTML(testDiagram())
	if err != nil {
		t.Fatalf("GenerateHTML: %v", err)
	}
	if !strings.Contains(html, "FlowLens") {
		t.Error("HTML does not contain the FlowLens bundle")
	}
}

func TestGenerateHTML_ContainsDiagramJSON(t *testing.T) {
	html, err := GenerateHTML(testDiagram())
	if err != nil {
		t.Fatalf("GenerateHTML: %v", err)
	}
	if !strings.Contains(html, "flowlens-data") {
		t.Error("HTML missing flowlens-data script element")
	}
	if !strings.Contains(html, "Test Diagram") {
		t.Error("HTML missing diagram title in JSON")
	}
}

func TestGenerateHTML_EscapesJSON(t *testing.T) {
	d := testDiagram()
	d.Meta.Title = "</script>"
	html, err := GenerateHTML(d)
	if err != nil {
		t.Fatalf("GenerateHTML: %v", err)
	}
	// The raw "</script>" must not appear as a literal closing tag.
	// json.Marshal escapes < and > as < and > in string values,
	// and escapeForHTML handles any remaining </ sequences in the full JSON.
	dataStart := strings.Index(html, `id="flowlens-data">`)
	if dataStart == -1 {
		t.Fatal("flowlens-data script tag not found")
	}
	dataSection := html[dataStart:]
	dataEnd := strings.Index(dataSection, "</script>")
	if dataEnd == -1 {
		t.Fatal("could not find closing </script> for data section")
	}
	jsonContent := dataSection[len(`id="flowlens-data">`):dataEnd]
	if strings.Contains(jsonContent, "</script>") {
		t.Error("JSON data contains literal </script> tag, XSS vector")
	}
}

func TestGenerateHTML_CSPHashCount(t *testing.T) {
	html, err := GenerateHTML(testDiagram())
	if err != nil {
		t.Fatalf("GenerateHTML: %v", err)
	}
	count := strings.Count(html, "'sha256-")
	if count != 2 {
		t.Errorf("expected 2 CSP sha256 hashes, got %d", count)
	}
}

func TestEmbeddedBundleNotEmpty(t *testing.T) {
	if len(flowlensBundle) == 0 {
		t.Fatal("flowlensBundle is empty; go:embed likely failed")
	}
}

func TestEscapeForHTML(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "closing script tag",
			input: `{"v":"</script>"}`,
			want:  "{\"v\":\"<\\/script>\"}",
		},
		{
			name:  "HTML comment",
			input: `{"v":"<!--comment-->"}`,
			want:  "{\"v\":\"<\\!--comment-->\"}",
		},
		{
			name:  "no special chars",
			input: `{"key":"value"}`,
			want:  `{"key":"value"}`,
		},
		{
			name:  "multiple occurrences",
			input: "</</",
			want:  "<\\/<\\/",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := escapeForHTML(tc.input)
			if got != tc.want {
				t.Errorf("escapeForHTML(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestSHA256Base64(t *testing.T) {
	// SHA-256 of "hello" is well-known.
	h := sha256.Sum256([]byte("hello"))
	want := base64.StdEncoding.EncodeToString(h[:])
	got := sha256Base64("hello")
	if got != want {
		t.Errorf("sha256Base64(\"hello\") = %q, want %q", got, want)
	}
	// Verify against the known hex digest converted to base64.
	// SHA-256("hello") = 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
	const knownBase64 = "LPJNul+wow4m6DsqxbninhsWHlwfp0JecwQzYpOLmCQ="
	if got != knownBase64 {
		t.Errorf("sha256Base64(\"hello\") = %q, want known %q", got, knownBase64)
	}
}
