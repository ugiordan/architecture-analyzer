package extractor

import "fmt"

var scaffoldingNames = map[string]bool{
	"controller-manager":                true,
	"controller-manager-metrics-monitor": true,
	"controller-manager-metrics-service": true,
	"leader-election-role":              true,
	"leader-election-rolebinding":       true,
	"manager-role":                      true,
	"manager-rolebinding":               true,
	"proxy-role":                        true,
	"proxy-rolebinding":                 true,
	"metrics-reader":                    true,
	"metrics-auth-role":                 true,
	"metrics-auth-rolebinding":          true,
}

func evalScaffoldingNameCollisions(arch *ComponentArchitecture) []SecurityAnnotation {
	if arch == nil {
		return nil
	}
	var annotations []SecurityAnnotation

	if arch.RBAC != nil {
		for _, cr := range arch.RBAC.ClusterRoles {
			if scaffoldingNames[cr.Name] {
				annotations = append(annotations, SecurityAnnotation{
					Type:     "SCAFFOLDING_NAME_COLLISION",
					Severity: "medium",
					Source:   cr.Source,
					Resource: cr.Name,
					Description: fmt.Sprintf(
						"ClusterRole uses default kubebuilder name %q. "+
							"When multiple operators are deployed, these generic "+
							"names collide. Prefix with the operator name.",
						cr.Name,
					),
				})
			}
		}
		for _, r := range arch.RBAC.Roles {
			if scaffoldingNames[r.Name] {
				annotations = append(annotations, SecurityAnnotation{
					Type:     "SCAFFOLDING_NAME_COLLISION",
					Severity: "medium",
					Source:   r.Source,
					Resource: r.Name,
					Description: fmt.Sprintf(
						"Role uses default kubebuilder name %q. "+
							"When multiple operators are deployed, these generic "+
							"names collide. Prefix with the operator name.",
						r.Name,
					),
				})
			}
		}
		for _, crb := range arch.RBAC.ClusterRoleBindings {
			if scaffoldingNames[crb.Name] {
				annotations = append(annotations, SecurityAnnotation{
					Type:     "SCAFFOLDING_NAME_COLLISION",
					Severity: "medium",
					Source:   crb.Source,
					Resource: crb.Name,
					Description: fmt.Sprintf(
						"ClusterRoleBinding uses default kubebuilder name %q. "+
							"Prefix with the operator name to avoid collisions.",
						crb.Name,
					),
				})
			}
		}
		for _, rb := range arch.RBAC.RoleBindings {
			if scaffoldingNames[rb.Name] {
				annotations = append(annotations, SecurityAnnotation{
					Type:     "SCAFFOLDING_NAME_COLLISION",
					Severity: "medium",
					Source:   rb.Source,
					Resource: rb.Name,
					Description: fmt.Sprintf(
						"RoleBinding uses default kubebuilder name %q. "+
							"Prefix with the operator name to avoid collisions.",
						rb.Name,
					),
				})
			}
		}
	}

	for _, dep := range arch.Deployments {
		if scaffoldingNames[dep.Name] {
			annotations = append(annotations, SecurityAnnotation{
				Type:     "SCAFFOLDING_NAME_COLLISION",
				Severity: "medium",
				Source:   dep.Source,
				Resource: dep.Name,
				Description: fmt.Sprintf(
					"Deployment uses default kubebuilder name %q. "+
						"Prefix with the operator name to avoid collisions.",
					dep.Name,
				),
			})
		}
	}

	for _, svc := range arch.Services {
		if scaffoldingNames[svc.Name] {
			annotations = append(annotations, SecurityAnnotation{
				Type:     "SCAFFOLDING_NAME_COLLISION",
				Severity: "medium",
				Source:   svc.Source,
				Resource: svc.Name,
				Description: fmt.Sprintf(
					"Service uses default kubebuilder name %q. "+
						"Prefix with the operator name to avoid collisions.",
					svc.Name,
				),
			})
		}
	}

	return annotations
}
