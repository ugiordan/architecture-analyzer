package extractor

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

type ghaWorkflow struct {
	Name        string             `yaml:"name"`
	On          interface{}        `yaml:"on"`
	Permissions interface{}        `yaml:"permissions"`
	Jobs        map[string]ghaJob  `yaml:"jobs"`
}

type ghaJob struct {
	Steps       []ghaStep   `yaml:"steps"`
	Permissions interface{} `yaml:"permissions"`
}

type ghaStep struct {
	Name string                 `yaml:"name"`
	Uses string                 `yaml:"uses"`
	With map[string]interface{} `yaml:"with"`
}

var shaRefRE = regexp.MustCompile(`^[0-9a-f]{40}([0-9a-f]{24})?$`)

func evalGitHubActionsWorkflows(repoPath string) []SecurityAnnotation {
	files := findWorkflowFiles(repoPath)
	if len(files) == 0 {
		return nil
	}

	var annotations []SecurityAnnotation
	for _, f := range files {
		wf, err := parseGHAWorkflow(f)
		if err != nil {
			continue
		}
		rel, _ := filepath.Rel(repoPath, f)
		if rel == "" {
			rel = f
		}
		annotations = append(annotations, checkMissingPermissions(wf, rel)...)
		annotations = append(annotations, checkPullRequestTarget(wf, rel)...)
		annotations = append(annotations, checkUnpinnedActions(wf, rel)...)
	}
	return annotations
}

func findWorkflowFiles(repoPath string) []string {
	dir := filepath.Join(repoPath, ".github", "workflows")
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	var files []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if strings.HasSuffix(name, ".yml") || strings.HasSuffix(name, ".yaml") {
			files = append(files, filepath.Join(dir, name))
		}
	}
	return files
}

func parseGHAWorkflow(path string) (*ghaWorkflow, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if info.Size() > 10*1024*1024 {
		return nil, fmt.Errorf("workflow file too large: %s", path)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var wf ghaWorkflow
	if err := yaml.Unmarshal(data, &wf); err != nil {
		return nil, err
	}
	return &wf, nil
}

func ghaWorkflowTriggers(on interface{}) map[string]bool {
	triggers := make(map[string]bool)
	switch v := on.(type) {
	case string:
		if v != "" {
			triggers[v] = true
		}
	case []interface{}:
		for _, item := range v {
			if s, ok := item.(string); ok && s != "" {
				triggers[s] = true
			}
		}
	case map[string]interface{}:
		for k := range v {
			triggers[k] = true
		}
	}
	return triggers
}

func isReusableWorkflowOnly(triggers map[string]bool) bool {
	return len(triggers) == 1 && triggers["workflow_call"]
}

func isSHAPin(ref string) bool {
	return shaRefRE.MatchString(ref)
}

func refReferencesForkedCode(ref string) bool {
	lower := strings.ToLower(ref)
	return strings.Contains(lower, "github.event.pull_request.head") ||
		strings.Contains(lower, "github.head_ref") ||
		strings.Contains(lower, "refs/pull/")
}

func checkMissingPermissions(wf *ghaWorkflow, source string) []SecurityAnnotation {
	triggers := ghaWorkflowTriggers(wf.On)
	if len(triggers) == 0 {
		return nil
	}
	if isReusableWorkflowOnly(triggers) {
		return nil
	}
	if wf.Permissions != nil {
		return nil
	}
	if allJobsHavePermissions(wf) {
		return nil
	}
	return []SecurityAnnotation{{
		Type:     "GHA_MISSING_PERMISSIONS",
		Severity: "high",
		Source:   source,
		Description: fmt.Sprintf(
			"Workflow %q has no top-level permissions block. "+
				"Without explicit permissions, the GITHUB_TOKEN gets default "+
				"permissions which may be overly broad (write access to contents, "+
				"packages, etc.).",
			wf.Name,
		),
	}}
}

func allJobsHavePermissions(wf *ghaWorkflow) bool {
	if len(wf.Jobs) == 0 {
		return false
	}
	for _, job := range wf.Jobs {
		if job.Permissions == nil {
			return false
		}
	}
	return true
}

func isCheckoutAction(uses string) bool {
	parts := strings.SplitN(uses, "@", 2)
	return parts[0] == "actions/checkout"
}

func checkPullRequestTarget(wf *ghaWorkflow, source string) []SecurityAnnotation {
	triggers := ghaWorkflowTriggers(wf.On)
	if !triggers["pull_request_target"] {
		return nil
	}

	var annotations []SecurityAnnotation
	for jobName, job := range wf.Jobs {
		for stepIdx, step := range job.Steps {
			if !isCheckoutAction(step.Uses) {
				continue
			}
			refVal, ok := step.With["ref"]
			if !ok {
				continue
			}
			refStr, ok := refVal.(string)
			if !ok || refStr == "" {
				continue
			}
			if !refReferencesForkedCode(refStr) {
				continue
			}
			stepName := step.Name
			if stepName == "" {
				stepName = fmt.Sprintf("step %d", stepIdx+1)
			}
			annotations = append(annotations, SecurityAnnotation{
				Type:     "GHA_PULL_REQUEST_TARGET",
				Severity: "critical",
				Source:   source,
				Description: fmt.Sprintf(
					"Workflow %q uses pull_request_target trigger and job %q, %s "+
						"checks out fork code (ref: %s). This allows arbitrary code "+
						"from a fork PR to run with write access to the base repo's "+
						"secrets and GITHUB_TOKEN.",
					wf.Name, jobName, stepName, refStr,
				),
			})
		}
	}
	return annotations
}

func checkUnpinnedActions(wf *ghaWorkflow, source string) []SecurityAnnotation {
	var annotations []SecurityAnnotation
	for _, job := range wf.Jobs {
		for _, step := range job.Steps {
			if step.Uses == "" {
				continue
			}
			uses := step.Uses
			if strings.HasPrefix(uses, "./") || strings.HasPrefix(uses, "docker://") {
				continue
			}
			atIdx := strings.LastIndex(uses, "@")
			if atIdx < 0 {
				annotations = append(annotations, SecurityAnnotation{
					Type:     "GHA_UNPINNED_ACTION",
					Severity: "medium",
					Source:   source,
					Resource: uses,
					Description: fmt.Sprintf(
						"Action %q has no version pin. Use a SHA-pinned reference "+
							"(action@sha256hash) for supply chain integrity.",
						uses,
					),
				})
				continue
			}
			ref := uses[atIdx+1:]
			if !isSHAPin(ref) {
				annotations = append(annotations, SecurityAnnotation{
					Type:     "GHA_UNPINNED_ACTION",
					Severity: "medium",
					Source:   source,
					Resource: uses,
					Description: fmt.Sprintf(
						"Action %q uses mutable ref %q instead of a SHA pin. "+
							"A compromised action version can inject malicious code "+
							"into the workflow.",
						uses[:atIdx], ref,
					),
				})
			}
		}
	}
	return annotations
}
