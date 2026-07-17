package extractor

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func safeRepoJoin(repoPath, file string) (string, bool) {
	fullPath := filepath.Join(repoPath, file)
	absPath, err := filepath.Abs(fullPath)
	if err != nil {
		return "", false
	}
	absRepo, err := filepath.Abs(repoPath)
	if err != nil {
		return "", false
	}
	if !strings.HasPrefix(absPath, absRepo+string(filepath.Separator)) && absPath != absRepo {
		return "", false
	}
	return fullPath, true
}

var sensitiveRBACResources = map[string]bool{
	"secrets":                    true,
	"configmaps":                true,
	"clusterrolebindings":       true,
	"securitycontextconstraints": true,
	"roles":                     true,
	"clusterroles":              true,
	"rolebindings":              true,
	"serviceaccounts":           true,
	"pods/exec":                 true,
	"nodes":                     true,
}

var mutatingVerbs = map[string]bool{
	"create": true,
	"update": true,
	"patch":  true,
	"delete": true,
	"*":      true,
}

var envVarRefRE = regexp.MustCompile(`\$\(([A-Za-z_][A-Za-z0-9_]*)\)`)

var imageFieldRE = regexp.MustCompile(`(?i)(image|container[_-]?image|runtime[_-]?image|sidecar[_-]?image)`)

func EvaluateSecurityAnnotations(arch *ComponentArchitecture, repoPath string) []SecurityAnnotation {
	var annotations []SecurityAnnotation
	annotations = append(annotations, evalRBACScope(arch)...)
	annotations = append(annotations, evalSecretExposureInArgs(arch)...)
	annotations = append(annotations, evalRouteNoTLS(arch)...)
	annotations = append(annotations, evalCRDConfusedDeputy(arch, repoPath)...)
	annotations = append(annotations, evalMissingMutualExclusion(arch, repoPath)...)
	annotations = append(annotations, evalHardcodedSecretValues(arch)...)
	annotations = append(annotations, evalPermissivePasswordEnv(arch)...)
	annotations = append(annotations, evalAuthBypassArgs(arch)...)
	annotations = append(annotations, evalDebugEndpoints(arch, repoPath)...)
	annotations = append(annotations, evalKustomizeSecurityDeletion(arch, repoPath)...)
	annotations = append(annotations, evalSecretInURL(arch, repoPath)...)
	annotations = append(annotations, evalGitHubActionsWorkflows(repoPath)...)
	annotations = append(annotations, evalScaffoldingNameCollisions(arch)...)
	return annotations
}

func evalRBACScope(arch *ComponentArchitecture) []SecurityAnnotation {
	if arch.RBAC == nil {
		return nil
	}
	var annotations []SecurityAnnotation
	for _, cr := range arch.RBAC.ClusterRoles {
		for _, rule := range cr.Rules {
			for _, resource := range rule.Resources {
				if !sensitiveRBACResources[strings.ToLower(resource)] {
					continue
				}
				var dangerousVerbs []string
				for _, verb := range rule.Verbs {
					if mutatingVerbs[strings.ToLower(verb)] {
						dangerousVerbs = append(dangerousVerbs, verb)
					}
				}
				if len(dangerousVerbs) == 0 {
					continue
				}
				severity := "high"
				if strings.ToLower(resource) == "configmaps" {
					severity = "medium"
				}
				annotations = append(annotations, SecurityAnnotation{
					Type:        "RBAC_CLUSTER_SCOPE_SENSITIVE",
					Severity:    severity,
					Resource:    resource,
					Verbs:       dangerousVerbs,
					Source:      cr.Source,
					Description: fmt.Sprintf("ClusterRole %q grants cluster-wide %s %s", cr.Name, strings.Join(dangerousVerbs, "/"), resource),
				})
			}
		}
	}
	return annotations
}

func evalSecretExposureInArgs(arch *ComponentArchitecture) []SecurityAnnotation {
	var annotations []SecurityAnnotation
	for _, dep := range arch.Deployments {
		allContainers := make([]Container, 0, len(dep.Containers)+len(dep.InitContainers))
		allContainers = append(allContainers, dep.Containers...)
		allContainers = append(allContainers, dep.InitContainers...)
		for _, c := range allContainers {
			secretEnvNames := make(map[string]string)
			for _, ref := range c.EnvVarRefs {
				if ref.SecretName != "" {
					secretEnvNames[ref.Name] = ref.SecretName
				}
			}
			if len(secretEnvNames) == 0 {
				continue
			}
			allStrings := make([]string, 0, len(c.Args)+len(c.Command))
			allStrings = append(allStrings, c.Args...)
			allStrings = append(allStrings, c.Command...)
			for _, arg := range allStrings {
				matches := envVarRefRE.FindAllStringSubmatch(arg, -1)
				for _, match := range matches {
					envName := match[1]
					if secretName, ok := secretEnvNames[envName]; ok {
						annotations = append(annotations, SecurityAnnotation{
							Type:        "SECRET_IN_CONTAINER_ARGS",
							Severity:    "medium",
							Source:      dep.Source,
							Container:   c.Name,
							EnvVar:      envName,
							Description: fmt.Sprintf("Container %q in %s uses $(%s) in args/command, which references secret %q. Kubelet expands this into /proc/1/cmdline, exposing the secret value to any process that can read procfs.", c.Name, dep.Name, envName, secretName),
						})
					}
				}
			}
		}
	}
	return annotations
}

func evalRouteNoTLS(arch *ComponentArchitecture) []SecurityAnnotation {
	var annotations []SecurityAnnotation
	for _, res := range arch.IngressRouting {
		if res.Kind != "Route" {
			continue
		}
		if res.TLS {
			continue
		}
		if len(res.RBACVerbs) > 0 {
			continue
		}
		annotations = append(annotations, SecurityAnnotation{
			Type:        "ROUTE_NO_TLS",
			Severity:    "medium",
			Source:      res.Source,
			Resource:    res.Name,
			Description: fmt.Sprintf("Route %q has no TLS configuration. Traffic between the router and backend service is unencrypted.", res.Name),
		})
	}
	return annotations
}

func evalCRDConfusedDeputy(arch *ComponentArchitecture, repoPath string) []SecurityAnnotation {
	var annotations []SecurityAnnotation
	for _, crd := range arch.CRDs {
		if crd.Source == "" {
			continue
		}
		imageFields := findImageFieldsInCRD(crd, repoPath)
		if len(imageFields) == 0 {
			continue
		}
		hasServiceAccount := false
		for _, dep := range arch.Deployments {
			if dep.ServiceAccount != "" {
				hasServiceAccount = true
				break
			}
		}
		if !hasServiceAccount {
			continue
		}
		for _, field := range imageFields {
			annotations = append(annotations, SecurityAnnotation{
				Type:        "CRD_CONFUSED_DEPUTY",
				Severity:    "high",
				Source:      crd.Source,
				CRDKind:     crd.Kind,
				Field:       field,
				Description: fmt.Sprintf("CRD %s has user-settable image field %q. The operator deploys containers using this value with its own ServiceAccount. A user-controlled image runs with the operator's permissions.", crd.Kind, field),
			})
		}
	}
	return annotations
}

func findImageFieldsInCRD(crd CRD, repoPath string) []string {
	if crd.Source == "" {
		return nil
	}
	fullPath, ok := safeRepoJoin(repoPath, crd.Source)
	if !ok {
		return nil
	}
	docs := parseYAMLSafe(fullPath)
	if len(docs) == 0 {
		return nil
	}
	for _, doc := range docs {
		kind, _ := doc["kind"].(string)
		if kind != "CustomResourceDefinition" {
			continue
		}
		spec, ok := doc["spec"].(map[string]interface{})
		if !ok {
			continue
		}
		versions := toSliceOfMaps(spec["versions"])
		for _, ver := range versions {
			schema, ok := ver["schema"].(map[string]interface{})
			if !ok {
				continue
			}
			openAPI, ok := schema["openAPIV3Schema"].(map[string]interface{})
			if !ok {
				continue
			}
			specProps := navigateToSpec(openAPI)
			if specProps == nil {
				continue
			}
			return findImageFields(specProps, "spec", 0)
		}
	}
	return nil
}

func navigateToSpec(schema map[string]interface{}) map[string]interface{} {
	props, ok := schema["properties"].(map[string]interface{})
	if !ok {
		return nil
	}
	specProp, ok := props["spec"].(map[string]interface{})
	if !ok {
		return nil
	}
	specProps, ok := specProp["properties"].(map[string]interface{})
	if !ok {
		return nil
	}
	return specProps
}

func findImageFields(props map[string]interface{}, prefix string, depth int) []string {
	if depth > 10 {
		return nil
	}
	var fields []string
	for name, propSchema := range props {
		fullPath := prefix + "." + name
		if imageFieldRE.MatchString(name) {
			ps, ok := propSchema.(map[string]interface{})
			if ok {
				pType, _ := ps["type"].(string)
				if pType == "string" || pType == "" {
					fields = append(fields, fullPath)
				}
			}
		}
		if ps, ok := propSchema.(map[string]interface{}); ok {
			if subProps, ok := ps["properties"].(map[string]interface{}); ok {
				fields = append(fields, findImageFields(subProps, fullPath, depth+1)...)
			}
		}
	}
	return fields
}

func evalMissingMutualExclusion(arch *ComponentArchitecture, repoPath string) []SecurityAnnotation {
	authComponentNames := []string{
		"kubeRBACProxy", "kube_rbac_proxy", "kubeRbacProxy",
		"oauthProxy", "oauth_proxy", "oauthproxy",
		"authorino", "authConfig", "auth_config",
	}

	var annotations []SecurityAnnotation
	for _, crd := range arch.CRDs {
		if crd.Source == "" {
			continue
		}
		fullPath, ok := safeRepoJoin(repoPath, crd.Source)
		if !ok {
			continue
		}
		docs := parseYAMLSafe(fullPath)
		for _, doc := range docs {
			kind, _ := doc["kind"].(string)
			if kind != "CustomResourceDefinition" {
				continue
			}
			spec, ok := doc["spec"].(map[string]interface{})
			if !ok {
				continue
			}
			versions := toSliceOfMaps(spec["versions"])
			for _, ver := range versions {
				schema, ok := ver["schema"].(map[string]interface{})
				if !ok {
					continue
				}
				openAPI, ok := schema["openAPIV3Schema"].(map[string]interface{})
				if !ok {
					continue
				}
				specProps := navigateToSpec(openAPI)
				if specProps == nil {
					continue
				}
				var foundAuthFields []string
				for _, authName := range authComponentNames {
					if _, exists := specProps[authName]; exists {
						ps, ok := specProps[authName].(map[string]interface{})
						if ok {
							pType, _ := ps["type"].(string)
							if pType == "object" || pType == "" {
								foundAuthFields = append(foundAuthFields, authName)
							}
						}
					}
				}
				if len(foundAuthFields) < 2 {
					continue
				}
				hasExclusion := false
				hasRequirement := false
				for _, rule := range crd.ValidationRules {
					lowerRule := strings.ToLower(rule)
					if containsAny(lowerRule, foundAuthFields) {
						hasExclusion = true
					}
					if strings.Contains(lowerRule, "has(") && containsAny(lowerRule, foundAuthFields) {
						if strings.Contains(lowerRule, "||") {
							hasRequirement = true
						}
					}
				}
				if hasExclusion && !hasRequirement {
					annotations = append(annotations, SecurityAnnotation{
						Type:        "MISSING_AUTH_REQUIREMENT",
						Severity:    "high",
						Source:      crd.Source,
						CRDKind:     crd.Kind,
						Field:       strings.Join(foundAuthFields, ", "),
						Description: fmt.Sprintf("CRD %s has optional auth component fields (%s) with mutual exclusion (XValidation prevents both) but no rule requiring at least one. The REST API can be exposed unauthenticated when all auth components are omitted.", crd.Kind, strings.Join(foundAuthFields, ", ")),
					})
				}
			}
		}
	}
	return annotations
}

func containsAny(s string, substrs []string) bool {
	for _, sub := range substrs {
		if strings.Contains(s, strings.ToLower(sub)) {
			return true
		}
	}
	return false
}

var knownPlaceholderSecrets = []string{
	"secret", "changeme", "password", "password123", "admin",
	"default", "placeholder", "fixme", "todo", "replace_me",
}

func evalHardcodedSecretValues(arch *ComponentArchitecture) []SecurityAnnotation {
	var annotations []SecurityAnnotation
	secretEnvPatterns := []string{
		"secret", "password", "token", "key", "credential",
	}
	for _, dep := range arch.Deployments {
		for _, c := range append(dep.Containers, dep.InitContainers...) {
			for name, value := range c.EnvVars {
				nameLower := strings.ToLower(name)
				isSecretEnv := false
				for _, p := range secretEnvPatterns {
					if strings.Contains(nameLower, p) {
						isSecretEnv = true
						break
					}
				}
				if !isSecretEnv {
					continue
				}
				valueLower := strings.ToLower(strings.TrimSpace(value))
				for _, placeholder := range knownPlaceholderSecrets {
					if valueLower == placeholder {
						annotations = append(annotations, SecurityAnnotation{
							Type:        "HARDCODED_SECRET_VALUE",
							Severity:    "high",
							Source:      dep.Source,
							Container:   c.Name,
							EnvVar:      name,
							Description: fmt.Sprintf("Container %q in %s has env var %s=%q which is a known placeholder secret value", c.Name, dep.Name, name, value),
						})
						break
					}
				}
			}
		}
	}
	return annotations
}

var passwordPermissiveRE = regexp.MustCompile(`(?i)(ALLOW_EMPTY_PASSWORD|NO_PASSWORD|SKIP_PASSWORD|DISABLE_AUTH)`)

func evalPermissivePasswordEnv(arch *ComponentArchitecture) []SecurityAnnotation {
	var annotations []SecurityAnnotation
	for _, dep := range arch.Deployments {
		for _, c := range append(dep.Containers, dep.InitContainers...) {
			for name, value := range c.EnvVars {
				if !passwordPermissiveRE.MatchString(name) {
					continue
				}
				valueLower := strings.ToLower(strings.TrimSpace(value))
				if valueLower == "true" || valueLower == "1" || valueLower == "yes" {
					annotations = append(annotations, SecurityAnnotation{
						Type:        "PERMISSIVE_PASSWORD_ENV",
						Severity:    "medium",
						Source:      dep.Source,
						Container:   c.Name,
						EnvVar:      name,
						Description: fmt.Sprintf("Container %q in %s sets %s=%s, disabling password authentication", c.Name, dep.Name, name, value),
					})
				}
			}
		}
	}
	return annotations
}

var authBypassPatterns = []string{
	"--skip-auth-regex",
	"--ignore-paths",
	"--insecure-skip-tls-verify",
	"--skip-auth-preflight",
}

func evalAuthBypassArgs(arch *ComponentArchitecture) []SecurityAnnotation {
	var annotations []SecurityAnnotation
	for _, dep := range arch.Deployments {
		for _, c := range append(dep.Containers, dep.InitContainers...) {
			allArgs := append(c.Command, c.Args...)
			for _, arg := range allArgs {
				argLower := strings.ToLower(arg)
				for _, pattern := range authBypassPatterns {
					if strings.HasPrefix(argLower, pattern) {
						annotations = append(annotations, SecurityAnnotation{
							Type:        "AUTH_BYPASS_ARG",
							Severity:    "medium",
							Source:      dep.Source,
							Container:   c.Name,
							Description: fmt.Sprintf("Container %q in %s uses %s which bypasses authentication or TLS verification", c.Name, dep.Name, arg),
						})
					}
				}
			}
		}
	}
	return annotations
}

var pprofImportRE = regexp.MustCompile(`"net/http/pprof"`)
var pprofRegisterRE = regexp.MustCompile(`pprof\.(Register|Handler)`)

func evalDebugEndpoints(arch *ComponentArchitecture, repoPath string) []SecurityAnnotation {
	var annotations []SecurityAnnotation
	if arch.GoASTMode == "" {
		return nil
	}
	goFiles := findGoFiles(repoPath)
	for _, f := range goFiles {
		if strings.Contains(f, "vendor/") || strings.Contains(f, "test") {
			continue
		}
		fullPath, ok := safeRepoJoin(repoPath, f)
		if !ok {
			continue
		}
		content := readFileSafe(fullPath)
		if content == "" {
			continue
		}
		if pprofImportRE.MatchString(content) || pprofRegisterRE.MatchString(content) {
			annotations = append(annotations, SecurityAnnotation{
				Type:        "DEBUG_ENDPOINT_PPROF",
				Severity:    "medium",
				Source:      f,
				Description: fmt.Sprintf("File %s imports or registers pprof debug endpoint. Pprof exposes heap dumps, goroutine stacks, and CPU profiles which may leak sensitive data.", f),
			})
		}
	}
	return annotations
}

func evalKustomizeSecurityDeletion(arch *ComponentArchitecture, repoPath string) []SecurityAnnotation {
	var annotations []SecurityAnnotation
	kustomizeFiles := findFilesByName(repoPath, "kustomization.yaml", "kustomization.yml")
	securityResourceKinds := []string{"networkpolicy", "role", "clusterrole", "securitycontextconstraints"}
	for _, f := range kustomizeFiles {
		fullPath, ok := safeRepoJoin(repoPath, f)
		if !ok {
			continue
		}
		docs := parseYAMLSafe(fullPath)
		for _, doc := range docs {
			patches := toSliceOfMaps(doc["patches"])
			for _, patch := range patches {
				target, ok := patch["target"].(map[string]interface{})
				if !ok {
					continue
				}
				kind, _ := target["kind"].(string)
				kindLower := strings.ToLower(kind)
				isSecurity := false
				for _, sk := range securityResourceKinds {
					if kindLower == sk {
						isSecurity = true
						break
					}
				}
				if !isSecurity {
					continue
				}
				patchStr, _ := patch["patch"].(string)
				if strings.Contains(patchStr, "$patch: delete") || strings.Contains(patchStr, "op: remove") {
					annotations = append(annotations, SecurityAnnotation{
						Type:        "KUSTOMIZE_SECURITY_DELETION",
						Severity:    "high",
						Source:      f,
						Resource:    kind,
						Description: fmt.Sprintf("Kustomize overlay in %s deletes %s resource, removing security controls", f, kind),
					})
				}
			}
		}
	}
	return annotations
}

var secretInURLRE = regexp.MustCompile(`(?i)(api[_-]?key|token|secret|password|credential)[=/]`)

func evalSecretInURL(arch *ComponentArchitecture, repoPath string) []SecurityAnnotation {
	var annotations []SecurityAnnotation
	for _, dep := range arch.Deployments {
		for _, c := range append(dep.Containers, dep.InitContainers...) {
			allArgs := append(c.Command, c.Args...)
			for _, arg := range allArgs {
				if !strings.Contains(arg, "http") {
					continue
				}
				if secretInURLRE.MatchString(arg) {
					annotations = append(annotations, SecurityAnnotation{
						Type:        "SECRET_IN_URL",
						Severity:    "medium",
						Source:      dep.Source,
						Container:   c.Name,
						Description: fmt.Sprintf("Container %q in %s passes a secret/key/token in URL: %s", c.Name, dep.Name, truncate(arg, 120)),
					})
				}
			}
		}
	}
	return annotations
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

func findFilesByName(repoPath string, names ...string) []string {
	nameSet := make(map[string]bool)
	for _, n := range names {
		nameSet[n] = true
	}
	var files []string
	filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			base := filepath.Base(path)
			if base == "vendor" || base == ".git" || base == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}
		if nameSet[filepath.Base(path)] {
			rel, err := filepath.Rel(repoPath, path)
			if err == nil {
				files = append(files, rel)
			}
		}
		return nil
	})
	return files
}

func readFileSafe(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	if len(data) > 512*1024 {
		return ""
	}
	return string(data)
}
