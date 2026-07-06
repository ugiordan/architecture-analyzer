package srclang

import (
	"bytes"
	"strings"
	"testing"
)

func TestWriteDocument_MinimalSecurity(t *testing.T) {
	doc := &Document{
		Version: "0.0.1",
		Head: Head{
			Producer:  "arch-analyzer 0.2.0",
			Component: "test-operator",
			Layer:     "security",
			Repository: &Repository{
				URI:    "https://github.com/example/test-operator",
				Commit: "abc1234def5678",
				Branch: "main",
			},
			Extracted: "2026-07-01T10:00:00Z",
			Languages: []Language{{Name: "go", Version: "1.22"}},
		},
		Body: Body{
			Layer: Layer{Name: "security"},
		},
	}

	var buf bytes.Buffer
	err := WriteDocument(&buf, doc)
	if err != nil {
		t.Fatalf("WriteDocument() error: %v", err)
	}
	out := buf.String()

	checks := []string{
		`<?xml version="1.0" encoding="UTF-8"?>`,
		`<srclang version="0.0.1" xmlns="https://srclang.dev/ns/core/0">`,
		`<producer>arch-analyzer 0.2.0</producer>`,
		`<component>test-operator</component>`,
		`<layer name="security"/>`,
		`<language name="go" version="1.22"/>`,
		`<layer name="security">`,
		`</srclang>`,
	}
	for _, c := range checks {
		if !strings.Contains(out, c) {
			t.Errorf("output missing %q", c)
		}
	}
}

func TestWriteDocument_FunctionWithCode(t *testing.T) {
	doc := &Document{
		Version: "0.0.1",
		Head: Head{
			Component: "test",
			Layer:     "security",
		},
		Body: Body{
			Layer: Layer{
				Name: "security",
				Files: []File{{
					Path:     "main.go",
					Language: "go",
					Lines:    100,
					Functions: []Function{{
						Name:       "handleRequest",
						Kind:       "function",
						SourceLine: 10,
						Complexity: 5,
						BodyLines:  20,
						Code:       "func handleRequest(w http.ResponseWriter, r *http.Request) {\n\t// body\n}",
						Trust:      "untrusted",
						TaintRole:  "source",
					}},
				}},
			},
		},
	}

	var buf bytes.Buffer
	if err := WriteDocument(&buf, doc); err != nil {
		t.Fatalf("WriteDocument() error: %v", err)
	}
	out := buf.String()

	checks := []string{
		`<file path="main.go" language="go" lines="100">`,
		`<function name="handleRequest" kind="function" trust="untrusted" taint-role="source">`,
		`<source line="10"/>`,
		`<metrics complexity="5" lines="20"/>`,
		`<![CDATA[func handleRequest`,
		`</function>`,
		`</file>`,
	}
	for _, c := range checks {
		if !strings.Contains(out, c) {
			t.Errorf("output missing %q", c)
		}
	}
}

func TestWriteDocument_CDATASplitting(t *testing.T) {
	doc := &Document{
		Version: "0.0.1",
		Head:    Head{Component: "test", Layer: "security"},
		Body: Body{
			Layer: Layer{
				Name: "security",
				Files: []File{{
					Path: "main.go", Language: "go",
					Functions: []Function{{
						Name:       "f",
						Kind:       "function",
						SourceLine: 1,
						Code:       `if a ]]> b { return }`,
					}},
				}},
			},
		},
	}

	var buf bytes.Buffer
	if err := WriteDocument(&buf, doc); err != nil {
		t.Fatalf("WriteDocument() error: %v", err)
	}
	out := buf.String()

	if !strings.Contains(out, "]]]]><![CDATA[>") {
		t.Error("CDATA section not properly split for ]]> in source")
	}
}

func TestWriteDocument_EmptyAndZeroValues(t *testing.T) {
	doc := &Document{
		Version: "0.0.1",
		Head:    Head{Component: "test", Layer: "security"},
		Body: Body{
			Layer: Layer{
				Name: "security",
				Files: []File{{
					Path:     "main.go",
					Language: "go",
					Functions: []Function{{
						Name:       "empty",
						Kind:       "function",
						SourceLine: 1,
						Code:       "",
						Complexity: 0,
						BodyLines:  0,
					}},
				}},
			},
		},
	}

	var buf bytes.Buffer
	if err := WriteDocument(&buf, doc); err != nil {
		t.Fatalf("WriteDocument() error: %v", err)
	}
	out := buf.String()

	if strings.Contains(out, "<code>") {
		t.Error("empty code should not produce <code> element")
	}
	if strings.Contains(out, "<metrics") {
		t.Error("zero metrics should not produce <metrics> element")
	}
}

func TestWriteDocument_Relationship(t *testing.T) {
	doc := &Document{
		Version: "0.0.1",
		Head:    Head{Component: "test", Layer: "security"},
		Body: Body{
			Layer: Layer{
				Name: "security",
				Relationships: []Relationship{{
					Kind: "calls",
					From: Endpoint{Function: "Reconcile", File: "controller.go", Line: 45},
					To:   Endpoint{Function: "Create", File: "client.go", Resolved: boolPtr(false)},
				}},
			},
		},
	}

	var buf bytes.Buffer
	if err := WriteDocument(&buf, doc); err != nil {
		t.Fatalf("WriteDocument() error: %v", err)
	}
	out := buf.String()

	if !strings.Contains(out, `resolved="false"`) {
		t.Error("missing resolved=false on unresolved endpoint")
	}
}

func TestWriteDocument_Finding(t *testing.T) {
	doc := &Document{
		Version: "0.0.1",
		Head:    Head{Component: "test", Layer: "security"},
		Body: Body{
			Layer: Layer{
				Name: "security",
				Findings: []Finding{{
					ID:          "CGA-N01-001",
					Domain:      "netpolicy",
					Severity:    "high",
					Rule:        "CGA-N01",
					SourceFile:  "utils.go",
					SourceLine:  160,
					Title:       "Bare namespaceSelector",
					Description: "NetworkPolicy allows unrestricted access",
					Evidence: []Ref{
						{Type: "function", Name: "createPeer", File: "utils.go", Line: 200},
					},
				}},
			},
		},
	}

	var buf bytes.Buffer
	if err := WriteDocument(&buf, doc); err != nil {
		t.Fatalf("WriteDocument() error: %v", err)
	}
	out := buf.String()

	checks := []string{
		`<finding id="CGA-N01-001" domain="netpolicy" severity="high" rule="CGA-N01">`,
		`<title>Bare namespaceSelector</title>`,
		`<ref type="function" name="createPeer"`,
	}
	for _, c := range checks {
		if !strings.Contains(out, c) {
			t.Errorf("output missing %q", c)
		}
	}
}

func TestWriteDocument_AttributeEscaping(t *testing.T) {
	doc := &Document{
		Version: "0.0.1",
		Head:    Head{Component: "test", Layer: "security\nwith\nnewlines"},
		Body:    Body{Layer: Layer{Name: "security"}},
	}
	var buf bytes.Buffer
	if err := WriteDocument(&buf, doc); err != nil {
		t.Fatalf("WriteDocument() error: %v", err)
	}
	out := buf.String()
	if strings.Contains(out, `name="security`+"\n") {
		t.Error("raw newlines should be escaped in attributes")
	}
	if !strings.Contains(out, "&#10;") {
		t.Error("newlines should be escaped as &#10;")
	}
}

func TestWriteDocument_Platform(t *testing.T) {
	doc := &Document{
		Version: "0.0.1",
		Head: Head{
			Component: "kserve",
			Layer:     "security",
			Platform: &Platform{
				Name:       "RHOAI",
				Version:    "rhoai.next",
				Components: 69,
				Topology: []Namespace{
					{Name: "redhat-ods-applications", Zone: "control-plane"},
					{Name: "rhods-notebooks", Zone: "tenant"},
				},
				Inbound: []PlatformEdge{
					{Peer: "odh-model-controller", Type: "watches-crd", Target: "InferenceService"},
				},
				Outbound: []PlatformEdge{
					{Peer: "knative-serving", Type: "go-module", Target: "knative.dev/serving"},
				},
			},
		},
		Body: Body{Layer: Layer{Name: "security"}},
	}

	var buf bytes.Buffer
	if err := WriteDocument(&buf, doc); err != nil {
		t.Fatalf("WriteDocument() error: %v", err)
	}
	out := buf.String()

	checks := []string{
		`<platform name="RHOAI" version="rhoai.next" components="69">`,
		`<namespace name="redhat-ods-applications" zone="control-plane"/>`,
		`<namespace name="rhods-notebooks" zone="tenant"/>`,
		`<edge from="odh-model-controller" type="watches-crd" target="InferenceService"/>`,
		`<edge to="knative-serving" type="go-module" target="knative.dev/serving"/>`,
		`<inbound>`,
		`<outbound>`,
		`<topology>`,
		`</platform>`,
	}
	for _, c := range checks {
		if !strings.Contains(out, c) {
			t.Errorf("output missing %q", c)
		}
	}
}

func TestWriteDocument_PlatformEmpty(t *testing.T) {
	doc := &Document{
		Version: "0.0.1",
		Head: Head{
			Component: "test",
			Layer:     "security",
			Platform: &Platform{
				Name:       "TestPlatform",
				Components: 5,
			},
		},
		Body: Body{Layer: Layer{Name: "security"}},
	}

	var buf bytes.Buffer
	if err := WriteDocument(&buf, doc); err != nil {
		t.Fatalf("WriteDocument() error: %v", err)
	}
	out := buf.String()

	if !strings.Contains(out, `<platform name="TestPlatform" components="5">`) {
		t.Error("missing platform element")
	}
	if strings.Contains(out, "<topology>") {
		t.Error("empty topology should not produce <topology> element")
	}
	if strings.Contains(out, "<inbound>") {
		t.Error("empty inbound should not produce <inbound> element")
	}
}

func TestWriteDocument_NilPlatform(t *testing.T) {
	doc := &Document{
		Version: "0.0.1",
		Head:    Head{Component: "test", Layer: "security"},
		Body:    Body{Layer: Layer{Name: "security"}},
	}

	var buf bytes.Buffer
	if err := WriteDocument(&buf, doc); err != nil {
		t.Fatalf("WriteDocument() error: %v", err)
	}
	out := buf.String()

	if strings.Contains(out, "<platform") {
		t.Error("nil platform should not produce <platform> element")
	}
}

func boolPtr(b bool) *bool { return &b }
