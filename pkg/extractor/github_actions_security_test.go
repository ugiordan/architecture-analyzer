package extractor

import (
	"os"
	"path/filepath"
	"testing"
)

func writeWorkflow(t *testing.T, dir, name, content string) {
	t.Helper()
	wfDir := filepath.Join(dir, ".github", "workflows")
	if err := os.MkdirAll(wfDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(wfDir, name), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestCheckMissingPermissions_NoPermissions(t *testing.T) {
	wf := &ghaWorkflow{
		Name: "CI",
		On:   "push",
	}
	anns := checkMissingPermissions(wf, "ci.yml")
	if len(anns) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(anns))
	}
	if anns[0].Type != "GHA_MISSING_PERMISSIONS" {
		t.Errorf("type = %q, want GHA_MISSING_PERMISSIONS", anns[0].Type)
	}
	if anns[0].Severity != "high" {
		t.Errorf("severity = %q, want high", anns[0].Severity)
	}
}

func TestCheckMissingPermissions_HasPermissions(t *testing.T) {
	wf := &ghaWorkflow{
		Name:        "CI",
		On:          "push",
		Permissions: map[string]interface{}{"contents": "read"},
	}
	anns := checkMissingPermissions(wf, "ci.yml")
	if len(anns) != 0 {
		t.Fatalf("expected 0 findings, got %d", len(anns))
	}
}

func TestCheckMissingPermissions_EmptyPermissions(t *testing.T) {
	wf := &ghaWorkflow{
		Name:        "CI",
		On:          "push",
		Permissions: map[string]interface{}{},
	}
	anns := checkMissingPermissions(wf, "ci.yml")
	if len(anns) != 0 {
		t.Fatalf("expected 0 findings (empty permissions = no permissions), got %d", len(anns))
	}
}

func TestCheckMissingPermissions_ReusableWorkflowOnly(t *testing.T) {
	wf := &ghaWorkflow{
		Name: "Reusable",
		On:   "workflow_call",
	}
	anns := checkMissingPermissions(wf, "reusable.yml")
	if len(anns) != 0 {
		t.Fatalf("expected 0 findings for reusable-only workflow, got %d", len(anns))
	}
}

func TestCheckMissingPermissions_MixedTriggersWithWorkflowCall(t *testing.T) {
	wf := &ghaWorkflow{
		Name: "Mixed",
		On:   []interface{}{"push", "workflow_call"},
	}
	anns := checkMissingPermissions(wf, "mixed.yml")
	if len(anns) != 1 {
		t.Fatalf("expected 1 finding for mixed triggers, got %d", len(anns))
	}
}

func TestCheckMissingPermissions_MapTrigger(t *testing.T) {
	wf := &ghaWorkflow{
		Name: "CI",
		On:   map[string]interface{}{"push": map[string]interface{}{"branches": []interface{}{"main"}}},
	}
	anns := checkMissingPermissions(wf, "ci.yml")
	if len(anns) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(anns))
	}
}

func TestCheckPullRequestTarget_Detected(t *testing.T) {
	wf := &ghaWorkflow{
		Name: "PR Target",
		On:   "pull_request_target",
		Jobs: map[string]ghaJob{
			"build": {
				Steps: []ghaStep{
					{
						Name: "Checkout PR",
						Uses: "actions/checkout@v4",
						With: map[string]interface{}{
							"ref": "${{ github.event.pull_request.head.sha }}",
						},
					},
				},
			},
		},
	}
	anns := checkPullRequestTarget(wf, "pr.yml")
	if len(anns) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(anns))
	}
	if anns[0].Severity != "critical" {
		t.Errorf("severity = %q, want critical", anns[0].Severity)
	}
}

func TestCheckPullRequestTarget_SafeCheckout(t *testing.T) {
	wf := &ghaWorkflow{
		Name: "PR Target Safe",
		On:   "pull_request_target",
		Jobs: map[string]ghaJob{
			"build": {
				Steps: []ghaStep{
					{Uses: "actions/checkout@v4"},
				},
			},
		},
	}
	anns := checkPullRequestTarget(wf, "pr.yml")
	if len(anns) != 0 {
		t.Fatalf("expected 0 findings for default checkout, got %d", len(anns))
	}
}

func TestCheckPullRequestTarget_NoPRT(t *testing.T) {
	wf := &ghaWorkflow{
		Name: "Regular PR",
		On:   "pull_request",
		Jobs: map[string]ghaJob{
			"build": {
				Steps: []ghaStep{
					{
						Uses: "actions/checkout@v4",
						With: map[string]interface{}{
							"ref": "${{ github.event.pull_request.head.sha }}",
						},
					},
				},
			},
		},
	}
	anns := checkPullRequestTarget(wf, "pr.yml")
	if len(anns) != 0 {
		t.Fatalf("expected 0 findings for pull_request (not target), got %d", len(anns))
	}
}

func TestCheckPullRequestTarget_RefsPullPattern(t *testing.T) {
	wf := &ghaWorkflow{
		Name: "PR Target",
		On:   "pull_request_target",
		Jobs: map[string]ghaJob{
			"build": {
				Steps: []ghaStep{
					{
						Uses: "actions/checkout@v4",
						With: map[string]interface{}{
							"ref": "refs/pull/${{ github.event.number }}/head",
						},
					},
				},
			},
		},
	}
	anns := checkPullRequestTarget(wf, "pr.yml")
	if len(anns) != 1 {
		t.Fatalf("expected 1 finding for refs/pull/ pattern, got %d", len(anns))
	}
}

func TestCheckPullRequestTarget_HeadRefPattern(t *testing.T) {
	wf := &ghaWorkflow{
		Name: "PR Target",
		On:   "pull_request_target",
		Jobs: map[string]ghaJob{
			"build": {
				Steps: []ghaStep{
					{
						Uses: "actions/checkout@v4",
						With: map[string]interface{}{
							"ref": "${{ github.head_ref }}",
						},
					},
				},
			},
		},
	}
	anns := checkPullRequestTarget(wf, "pr.yml")
	if len(anns) != 1 {
		t.Fatalf("expected 1 finding for github.head_ref, got %d", len(anns))
	}
}

func TestCheckUnpinnedActions_TagRef(t *testing.T) {
	wf := &ghaWorkflow{
		Name: "CI",
		Jobs: map[string]ghaJob{
			"build": {
				Steps: []ghaStep{
					{Uses: "actions/checkout@v4"},
				},
			},
		},
	}
	anns := checkUnpinnedActions(wf, "ci.yml")
	if len(anns) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(anns))
	}
	if anns[0].Resource != "actions/checkout@v4" {
		t.Errorf("resource = %q, want actions/checkout@v4", anns[0].Resource)
	}
}

func TestCheckUnpinnedActions_SHAPin(t *testing.T) {
	wf := &ghaWorkflow{
		Name: "CI",
		Jobs: map[string]ghaJob{
			"build": {
				Steps: []ghaStep{
					{Uses: "actions/checkout@a5ac7e51b41094c92402da3b24376905380afc29"},
				},
			},
		},
	}
	anns := checkUnpinnedActions(wf, "ci.yml")
	if len(anns) != 0 {
		t.Fatalf("expected 0 findings for SHA-pinned action, got %d", len(anns))
	}
}

func TestCheckUnpinnedActions_LocalAction(t *testing.T) {
	wf := &ghaWorkflow{
		Name: "CI",
		Jobs: map[string]ghaJob{
			"build": {
				Steps: []ghaStep{
					{Uses: "./local-action"},
				},
			},
		},
	}
	anns := checkUnpinnedActions(wf, "ci.yml")
	if len(anns) != 0 {
		t.Fatalf("expected 0 findings for local action, got %d", len(anns))
	}
}

func TestCheckUnpinnedActions_DockerAction(t *testing.T) {
	wf := &ghaWorkflow{
		Name: "CI",
		Jobs: map[string]ghaJob{
			"build": {
				Steps: []ghaStep{
					{Uses: "docker://alpine:3.18"},
				},
			},
		},
	}
	anns := checkUnpinnedActions(wf, "ci.yml")
	if len(anns) != 0 {
		t.Fatalf("expected 0 findings for docker action, got %d", len(anns))
	}
}

func TestCheckUnpinnedActions_ReusableWorkflow(t *testing.T) {
	wf := &ghaWorkflow{
		Name: "CI",
		Jobs: map[string]ghaJob{
			"build": {
				Steps: []ghaStep{
					{Uses: "./.github/workflows/reusable.yml"},
				},
			},
		},
	}
	anns := checkUnpinnedActions(wf, "ci.yml")
	if len(anns) != 0 {
		t.Fatalf("expected 0 findings for reusable workflow call, got %d", len(anns))
	}
}

func TestCheckUnpinnedActions_MultipleUnpinned(t *testing.T) {
	wf := &ghaWorkflow{
		Name: "CI",
		Jobs: map[string]ghaJob{
			"build": {
				Steps: []ghaStep{
					{Uses: "actions/checkout@v4"},
					{Uses: "actions/setup-go@v5"},
					{Uses: "actions/cache@a5ac7e51b41094c92402da3b24376905380afc29"},
				},
			},
		},
	}
	anns := checkUnpinnedActions(wf, "ci.yml")
	if len(anns) != 2 {
		t.Fatalf("expected 2 findings (checkout + setup-go), got %d", len(anns))
	}
}

func TestEvalGitHubActionsWorkflows_EndToEnd(t *testing.T) {
	dir := t.TempDir()
	writeWorkflow(t, dir, "ci.yml", `
name: CI
on: push
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
`)
	writeWorkflow(t, dir, "pr.yml", `
name: PR Target
on: pull_request_target
permissions:
  contents: read
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@a5ac7e51b41094c92402da3b24376905380afc29
        with:
          ref: ${{ github.event.pull_request.head.sha }}
`)
	anns := evalGitHubActionsWorkflows(dir)

	typeCount := map[string]int{}
	for _, a := range anns {
		typeCount[a.Type]++
	}

	if typeCount["GHA_MISSING_PERMISSIONS"] != 1 {
		t.Errorf("GHA_MISSING_PERMISSIONS = %d, want 1 (ci.yml has no permissions)", typeCount["GHA_MISSING_PERMISSIONS"])
	}
	if typeCount["GHA_PULL_REQUEST_TARGET"] != 1 {
		t.Errorf("GHA_PULL_REQUEST_TARGET = %d, want 1 (pr.yml checks out fork)", typeCount["GHA_PULL_REQUEST_TARGET"])
	}
	if typeCount["GHA_UNPINNED_ACTION"] != 2 {
		t.Errorf("GHA_UNPINNED_ACTION = %d, want 2 (checkout@v4 + setup-go@v5 in ci.yml)", typeCount["GHA_UNPINNED_ACTION"])
	}
}

func TestEvalGitHubActionsWorkflows_NoWorkflowDir(t *testing.T) {
	dir := t.TempDir()
	anns := evalGitHubActionsWorkflows(dir)
	if len(anns) != 0 {
		t.Fatalf("expected 0 findings for empty dir, got %d", len(anns))
	}
}

func TestGHAWorkflowTriggers_AllForms(t *testing.T) {
	tests := []struct {
		name     string
		on       interface{}
		expected []string
	}{
		{"string", "push", []string{"push"}},
		{"list", []interface{}{"push", "pull_request"}, []string{"push", "pull_request"}},
		{"map", map[string]interface{}{"push": nil, "schedule": nil}, []string{"push", "schedule"}},
		{"nil", nil, nil},
		{"bool", true, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			triggers := ghaWorkflowTriggers(tt.on)
			if tt.expected == nil {
				if len(triggers) != 0 {
					t.Errorf("expected empty triggers, got %v", triggers)
				}
				return
			}
			for _, exp := range tt.expected {
				if !triggers[exp] {
					t.Errorf("missing trigger %q in %v", exp, triggers)
				}
			}
			if len(triggers) != len(tt.expected) {
				t.Errorf("trigger count = %d, want %d", len(triggers), len(tt.expected))
			}
		})
	}
}

func TestCheckMissingPermissions_JobLevelPermissions(t *testing.T) {
	wf := &ghaWorkflow{
		Name: "CI",
		On:   "push",
		Jobs: map[string]ghaJob{
			"build": {
				Permissions: map[string]interface{}{"contents": "read"},
				Steps:       []ghaStep{{Uses: "actions/checkout@v4"}},
			},
			"test": {
				Permissions: map[string]interface{}{"contents": "read"},
				Steps:       []ghaStep{{Uses: "actions/setup-go@v5"}},
			},
		},
	}
	anns := checkMissingPermissions(wf, "ci.yml")
	if len(anns) != 0 {
		t.Fatalf("expected 0 findings when all jobs have permissions, got %d", len(anns))
	}
}

func TestCheckMissingPermissions_PartialJobPermissions(t *testing.T) {
	wf := &ghaWorkflow{
		Name: "CI",
		On:   "push",
		Jobs: map[string]ghaJob{
			"build": {
				Permissions: map[string]interface{}{"contents": "read"},
			},
			"deploy": {
				Steps: []ghaStep{{Uses: "actions/checkout@v4"}},
			},
		},
	}
	anns := checkMissingPermissions(wf, "ci.yml")
	if len(anns) != 1 {
		t.Fatalf("expected 1 finding when not all jobs have permissions, got %d", len(anns))
	}
}

func TestCheckMissingPermissions_StringPermissions(t *testing.T) {
	wf := &ghaWorkflow{
		Name:        "CI",
		On:          "push",
		Permissions: "read-all",
	}
	anns := checkMissingPermissions(wf, "ci.yml")
	if len(anns) != 0 {
		t.Fatalf("expected 0 findings for permissions: read-all, got %d", len(anns))
	}
}

func TestCheckPullRequestTarget_MapTrigger(t *testing.T) {
	wf := &ghaWorkflow{
		Name: "PR Target Map",
		On:   map[string]interface{}{"pull_request_target": map[string]interface{}{"types": []interface{}{"opened"}}},
		Jobs: map[string]ghaJob{
			"build": {
				Steps: []ghaStep{{
					Uses: "actions/checkout@v4",
					With: map[string]interface{}{"ref": "${{ github.event.pull_request.head.sha }}"},
				}},
			},
		},
	}
	anns := checkPullRequestTarget(wf, "pr.yml")
	if len(anns) != 1 {
		t.Fatalf("expected 1 finding for map-form PRT trigger, got %d", len(anns))
	}
}

func TestCheckUnpinnedActions_RemoteReusableWorkflow(t *testing.T) {
	wf := &ghaWorkflow{
		Name: "CI",
		Jobs: map[string]ghaJob{
			"build": {
				Steps: []ghaStep{
					{Uses: "org/repo/.github/workflows/build.yml@main"},
				},
			},
		},
	}
	anns := checkUnpinnedActions(wf, "ci.yml")
	if len(anns) != 1 {
		t.Fatalf("expected 1 finding for remote reusable workflow with mutable ref, got %d", len(anns))
	}
}

func TestCheckUnpinnedActions_NoAtSeparator(t *testing.T) {
	wf := &ghaWorkflow{
		Name: "CI",
		Jobs: map[string]ghaJob{
			"build": {
				Steps: []ghaStep{
					{Uses: "actions/checkout"},
				},
			},
		},
	}
	anns := checkUnpinnedActions(wf, "ci.yml")
	if len(anns) != 1 {
		t.Fatalf("expected 1 finding for action with no @ separator, got %d", len(anns))
	}
}

func TestIsSHAPin(t *testing.T) {
	tests := []struct {
		ref  string
		want bool
	}{
		{"a5ac7e51b41094c92402da3b24376905380afc29", true},
		{"a5ac7e51b41094c92402da3b24376905380afc29abcdef1234567890abcdef12", true}, // 64-char SHA-256
		{"v4", false},
		{"main", false},
		{"a5ac7e51b41094c92402da3b24376905380afc2", false},  // 39 chars
		{"A5AC7E51B41094C92402DA3B24376905380AFC29", false}, // uppercase
		{"", false},
	}
	for _, tt := range tests {
		if got := isSHAPin(tt.ref); got != tt.want {
			t.Errorf("isSHAPin(%q) = %v, want %v", tt.ref, got, tt.want)
		}
	}
}
