package extractor

import (
	"testing"
)

func TestEvalRBACScope_ClusterRoleSecretsWrite(t *testing.T) {
	arch := &ComponentArchitecture{
		RBAC: &RBACData{
			ClusterRoles: []RBACRole{{
				Name:   "manager-role",
				Source: "config/rbac/role.yaml",
				Rules: []RBACRule{{
					APIGroups: []string{""},
					Resources: []string{"secrets"},
					Verbs:     []string{"get", "list", "create", "update", "delete"},
				}},
			}},
		},
	}

	annotations := evalRBACScope(arch)
	if len(annotations) == 0 {
		t.Fatal("expected RBAC annotation for secrets CRUD")
	}
	found := false
	for _, a := range annotations {
		if a.Type == "RBAC_CLUSTER_SCOPE_SENSITIVE" && a.Resource == "secrets" {
			found = true
			if a.Severity != "high" {
				t.Errorf("severity = %q, want %q", a.Severity, "high")
			}
			if len(a.Verbs) != 3 {
				t.Errorf("verbs count = %d, want 3 (create, update, delete)", len(a.Verbs))
			}
		}
	}
	if !found {
		t.Error("missing RBAC_CLUSTER_SCOPE_SENSITIVE annotation for secrets")
	}
}

func TestEvalRBACScope_ReadOnlyNotFlagged(t *testing.T) {
	arch := &ComponentArchitecture{
		RBAC: &RBACData{
			ClusterRoles: []RBACRole{{
				Name:   "viewer-role",
				Source: "config/rbac/role.yaml",
				Rules: []RBACRule{{
					APIGroups: []string{""},
					Resources: []string{"secrets"},
					Verbs:     []string{"get", "list", "watch"},
				}},
			}},
		},
	}

	annotations := evalRBACScope(arch)
	if len(annotations) != 0 {
		t.Errorf("read-only verbs should not be flagged, got %d annotations", len(annotations))
	}
}

func TestEvalRBACScope_WildcardVerb(t *testing.T) {
	arch := &ComponentArchitecture{
		RBAC: &RBACData{
			ClusterRoles: []RBACRole{{
				Name:   "admin-role",
				Source: "config/rbac/role.yaml",
				Rules: []RBACRule{{
					APIGroups: []string{""},
					Resources: []string{"secrets"},
					Verbs:     []string{"*"},
				}},
			}},
		},
	}

	annotations := evalRBACScope(arch)
	if len(annotations) == 0 {
		t.Fatal("wildcard verb on secrets should be flagged")
	}
}

func TestEvalRBACScope_ConfigmapsMediumSeverity(t *testing.T) {
	arch := &ComponentArchitecture{
		RBAC: &RBACData{
			ClusterRoles: []RBACRole{{
				Name:   "manager-role",
				Source: "config/rbac/role.yaml",
				Rules: []RBACRule{{
					APIGroups: []string{""},
					Resources: []string{"configmaps"},
					Verbs:     []string{"create", "update"},
				}},
			}},
		},
	}

	annotations := evalRBACScope(arch)
	if len(annotations) == 0 {
		t.Fatal("expected annotation for configmaps")
	}
	if annotations[0].Severity != "medium" {
		t.Errorf("configmaps severity = %q, want %q", annotations[0].Severity, "medium")
	}
}

func TestEvalRBACScope_NilRBAC(t *testing.T) {
	arch := &ComponentArchitecture{}
	annotations := evalRBACScope(arch)
	if len(annotations) != 0 {
		t.Errorf("nil RBAC should produce no annotations, got %d", len(annotations))
	}
}

func TestEvalRBACScope_RoleNotFlagged(t *testing.T) {
	arch := &ComponentArchitecture{
		RBAC: &RBACData{
			Roles: []RBACRole{{
				Name:   "namespace-role",
				Source: "config/rbac/role.yaml",
				Rules: []RBACRule{{
					APIGroups: []string{""},
					Resources: []string{"secrets"},
					Verbs:     []string{"create", "update", "delete"},
				}},
			}},
		},
	}

	annotations := evalRBACScope(arch)
	if len(annotations) != 0 {
		t.Errorf("namespace-scoped Role should not be flagged, got %d annotations", len(annotations))
	}
}

func TestEvalRBACScope_MultipleResources(t *testing.T) {
	arch := &ComponentArchitecture{
		RBAC: &RBACData{
			ClusterRoles: []RBACRole{{
				Name:   "manager-role",
				Source: "config/rbac/role.yaml",
				Rules: []RBACRule{
					{Resources: []string{"secrets"}, Verbs: []string{"create"}},
					{Resources: []string{"clusterrolebindings"}, Verbs: []string{"create"}},
					{Resources: []string{"pods"}, Verbs: []string{"create"}},
				},
			}},
		},
	}

	annotations := evalRBACScope(arch)
	if len(annotations) != 2 {
		t.Errorf("expected 2 annotations (secrets + clusterrolebindings), got %d", len(annotations))
	}
}

func TestEvalSecretExposureInArgs(t *testing.T) {
	arch := &ComponentArchitecture{
		Deployments: []Deployment{{
			Name:   "controller-manager",
			Source: "config/manager/manager.yaml",
			Containers: []Container{{
				Name: "manager",
				Args: []string{"--database-password=$(DATABASE_PASSWORD)", "--port=8080"},
				EnvVarRefs: []EnvVarRef{{
					Name:       "DATABASE_PASSWORD",
					SecretName: "db-credentials",
					SecretKey:  "password",
				}},
			}},
		}},
	}

	annotations := evalSecretExposureInArgs(arch)
	if len(annotations) == 0 {
		t.Fatal("expected SECRET_IN_CONTAINER_ARGS annotation")
	}
	a := annotations[0]
	if a.Type != "SECRET_IN_CONTAINER_ARGS" {
		t.Errorf("type = %q", a.Type)
	}
	if a.Container != "manager" {
		t.Errorf("container = %q", a.Container)
	}
	if a.EnvVar != "DATABASE_PASSWORD" {
		t.Errorf("env_var = %q", a.EnvVar)
	}
}

func TestEvalSecretExposureInArgs_NoSecretRef(t *testing.T) {
	arch := &ComponentArchitecture{
		Deployments: []Deployment{{
			Name: "controller",
			Containers: []Container{{
				Name: "manager",
				Args: []string{"--port=$(PORT)"},
				EnvVarRefs: []EnvVarRef{{
					Name:          "PORT",
					ConfigMapName: "config",
					ConfigMapKey:  "port",
				}},
			}},
		}},
	}

	annotations := evalSecretExposureInArgs(arch)
	if len(annotations) != 0 {
		t.Errorf("configmap ref in args should not be flagged, got %d", len(annotations))
	}
}

func TestEvalSecretExposureInArgs_SecretInCommand(t *testing.T) {
	arch := &ComponentArchitecture{
		Deployments: []Deployment{{
			Name: "controller",
			Containers: []Container{{
				Name:    "manager",
				Command: []string{"/bin/sh", "-c", "echo $(SECRET_TOKEN)"},
				EnvVarRefs: []EnvVarRef{{
					Name:       "SECRET_TOKEN",
					SecretName: "api-token",
					SecretKey:  "token",
				}},
			}},
		}},
	}

	annotations := evalSecretExposureInArgs(arch)
	if len(annotations) == 0 {
		t.Fatal("expected annotation for secret in command")
	}
}

func TestEvalSecretExposureInArgs_InitContainer(t *testing.T) {
	arch := &ComponentArchitecture{
		Deployments: []Deployment{{
			Name: "controller",
			InitContainers: []Container{{
				Name: "init-db",
				Args: []string{"--password=$(DB_PASS)"},
				EnvVarRefs: []EnvVarRef{{
					Name:       "DB_PASS",
					SecretName: "db-secret",
					SecretKey:  "password",
				}},
			}},
		}},
	}

	annotations := evalSecretExposureInArgs(arch)
	if len(annotations) == 0 {
		t.Fatal("expected annotation for init container secret exposure")
	}
}

func TestEvalSecretExposureInArgs_NoEnvVarRef(t *testing.T) {
	arch := &ComponentArchitecture{
		Deployments: []Deployment{{
			Name: "controller",
			Containers: []Container{{
				Name: "manager",
				Args: []string{"--flag=$(SOME_VAR)"},
			}},
		}},
	}

	annotations := evalSecretExposureInArgs(arch)
	if len(annotations) != 0 {
		t.Errorf("no env var refs should produce no annotations, got %d", len(annotations))
	}
}

func TestEvalMissingMutualExclusion_NoAnnotations(t *testing.T) {
	arch := &ComponentArchitecture{
		CRDs: []CRD{{
			Kind:   "SimpleResource",
			Source: "config/crd/simple.yaml",
		}},
	}

	annotations := evalMissingMutualExclusion(arch, "")
	if len(annotations) != 0 {
		t.Errorf("simple CRD should not be flagged, got %d", len(annotations))
	}
}

func TestEvaluateSecurityAnnotations_Integration(t *testing.T) {
	arch := &ComponentArchitecture{
		RBAC: &RBACData{
			ClusterRoles: []RBACRole{{
				Name:   "manager-role",
				Source: "config/rbac/role.yaml",
				Rules: []RBACRule{{
					Resources: []string{"secrets"},
					Verbs:     []string{"create", "update", "patch", "delete"},
				}},
			}},
		},
		Deployments: []Deployment{{
			Name:           "controller",
			Source:         "config/manager/manager.yaml",
			ServiceAccount: "controller-manager",
			Containers: []Container{{
				Name: "manager",
				Args: []string{"--db-pass=$(DB_PASSWORD)"},
				EnvVarRefs: []EnvVarRef{{
					Name:       "DB_PASSWORD",
					SecretName: "db-creds",
					SecretKey:  "password",
				}},
			}},
		}},
	}

	arch.IngressRouting = []IngressResource{{
		Kind:   "Route",
		Name:   "http-route",
		Source: "templates/http-route.yaml.tmpl",
		TLS:    false,
	}}

	annotations := EvaluateSecurityAnnotations(arch, "")
	types := make(map[string]bool)
	for _, a := range annotations {
		types[a.Type] = true
	}
	if !types["RBAC_CLUSTER_SCOPE_SENSITIVE"] {
		t.Error("missing RBAC_CLUSTER_SCOPE_SENSITIVE")
	}
	if !types["SECRET_IN_CONTAINER_ARGS"] {
		t.Error("missing SECRET_IN_CONTAINER_ARGS")
	}
	if !types["ROUTE_NO_TLS"] {
		t.Error("missing ROUTE_NO_TLS")
	}
}

func TestEvalRouteNoTLS_Detected(t *testing.T) {
	arch := &ComponentArchitecture{
		IngressRouting: []IngressResource{{
			Kind:   "Route",
			Name:   "http-route",
			Source: "internal/controller/config/templates/http-route.yaml.tmpl",
			TLS:    false,
		}},
	}
	annotations := evalRouteNoTLS(arch)
	if len(annotations) != 1 {
		t.Fatalf("expected 1 annotation, got %d", len(annotations))
	}
	a := annotations[0]
	if a.Type != "ROUTE_NO_TLS" {
		t.Errorf("type = %q, want ROUTE_NO_TLS", a.Type)
	}
	if a.Severity != "medium" {
		t.Errorf("severity = %q, want medium", a.Severity)
	}
	if a.Resource != "http-route" {
		t.Errorf("resource = %q, want http-route", a.Resource)
	}
}

func TestEvalRouteNoTLS_WithTLSNotFlagged(t *testing.T) {
	arch := &ComponentArchitecture{
		IngressRouting: []IngressResource{{
			Kind: "Route",
			Name: "secure-route",
			TLS:  true,
		}},
	}
	annotations := evalRouteNoTLS(arch)
	if len(annotations) != 0 {
		t.Errorf("Route with TLS should not be flagged, got %d", len(annotations))
	}
}

func TestEvalRouteNoTLS_NonRouteNotFlagged(t *testing.T) {
	arch := &ComponentArchitecture{
		IngressRouting: []IngressResource{{
			Kind: "Ingress",
			Name: "my-ingress",
			TLS:  false,
		}},
	}
	annotations := evalRouteNoTLS(arch)
	if len(annotations) != 0 {
		t.Errorf("non-Route kind should not be flagged, got %d", len(annotations))
	}
}

func TestEvalRouteNoTLS_RBACInferredSkipped(t *testing.T) {
	arch := &ComponentArchitecture{
		IngressRouting: []IngressResource{{
			Kind:      "Route",
			Name:      "inferred-route",
			TLS:       false,
			RBACVerbs: []string{"create", "update"},
		}},
	}
	annotations := evalRouteNoTLS(arch)
	if len(annotations) != 0 {
		t.Errorf("RBAC-inferred routes should not be flagged, got %d", len(annotations))
	}
}

func TestEvalRBACScope_PodsExec(t *testing.T) {
	arch := &ComponentArchitecture{
		RBAC: &RBACData{
			ClusterRoles: []RBACRole{{
				Name:   "debug-role",
				Source: "config/rbac/role.yaml",
				Rules: []RBACRule{{
					Resources: []string{"pods/exec"},
					Verbs:     []string{"create"},
				}},
			}},
		},
	}

	annotations := evalRBACScope(arch)
	if len(annotations) == 0 {
		t.Fatal("pods/exec create should be flagged")
	}
}

func TestEvalRBACScope_NonSensitiveResource(t *testing.T) {
	arch := &ComponentArchitecture{
		RBAC: &RBACData{
			ClusterRoles: []RBACRole{{
				Name:   "manager-role",
				Source: "config/rbac/role.yaml",
				Rules: []RBACRule{{
					Resources: []string{"deployments"},
					Verbs:     []string{"create", "update", "delete"},
				}},
			}},
		},
	}

	annotations := evalRBACScope(arch)
	if len(annotations) != 0 {
		t.Errorf("deployments is not sensitive, got %d annotations", len(annotations))
	}
}
