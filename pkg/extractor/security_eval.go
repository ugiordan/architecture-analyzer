package extractor

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

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
	annotations = append(annotations, evalCRDConfusedDeputy(arch, repoPath)...)
	annotations = append(annotations, evalMissingMutualExclusion(arch, repoPath)...)
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
		allContainers := append(dep.Containers, dep.InitContainers...)
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
			allStrings := append(c.Args, c.Command...)
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
	fullPath := filepath.Join(repoPath, crd.Source)
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
		fullPath := filepath.Join(repoPath, crd.Source)
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
								nullable, _ := ps["nullable"].(bool)
								xNullable, _ := ps["x-kubernetes-preserve-unknown-fields"].(bool)
								_, hasPointerMarker := ps["x-kubernetes-int-or-string"]
								if nullable || xNullable || hasPointerMarker || isOptionalPointer(ps) {
									foundAuthFields = append(foundAuthFields, authName)
								}
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

func isOptionalPointer(schema map[string]interface{}) bool {
	_, hasDefault := schema["default"]
	return !hasDefault
}

func containsAny(s string, substrs []string) bool {
	for _, sub := range substrs {
		if strings.Contains(s, strings.ToLower(sub)) {
			return true
		}
	}
	return false
}
