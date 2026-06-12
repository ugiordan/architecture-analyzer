package netpolicy

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ugiordan/architecture-analyzer/pkg/graph"
	"github.com/ugiordan/architecture-analyzer/pkg/query"
)

func netpolicyQueries() []query.Rule {
	return []query.Rule{
		{ID: "CGA-N01", Name: "netpol-bare-namespace-selector", Domain: "netpolicy", Severity: "high", Run: queryBareNamespaceSelector},
		{ID: "CGA-N02", Name: "netpol-tenant-reach", Domain: "netpolicy", Severity: "high", Run: queryTenantReach},
	}
}

// tenantNamespacePatterns identifies namespace names/patterns that run tenant workloads.
var tenantNamespacePatterns = []string{
	"notebook", "workbench", "project", "sandbox", "user", "tenant",
	"pipeline", "serving", "inference", "training",
}

// queryBareNamespaceSelector finds NetworkPolicies that use namespaceSelector
// without podSelector or port restrictions. A bare namespaceSelector grants
// full access from any pod in matching namespaces to all ports on target pods.
func queryBareNamespaceSelector(g *graph.CPG) []query.Finding {
	if g.ArchData == nil {
		return nil
	}

	var findings []query.Finding

	for _, np := range g.ArchData.NetworkPolicies {
		for _, rule := range np.IngressRules {
			selectors := extractNamespaceSelectors(rule)
			if len(selectors) == 0 {
				continue
			}

			hasPodSelector := hasPodSelectorInRule(rule)
			hasPorts := hasPortsInRule(rule)

			if !hasPodSelector && !hasPorts {
				for _, sel := range selectors {
					findings = append(findings, query.Finding{
						RuleID:   "CGA-N01",
						Severity: "high",
						Message: fmt.Sprintf(
							"NetworkPolicy %q allows ingress from namespaces matching %s=%s without podSelector or port restrictions. Any pod in matching namespaces gets full access.",
							np.Name, sel.key, sel.value,
						),
						File:   np.Source,
						Domain: "netpolicy",
					})
				}
			}
		}
	}

	return findings
}

// queryTenantReach traces whether namespaceSelector labels are applied to
// namespaces that run tenant/user workloads (notebooks, pipelines, model serving).
// If so, tenant code can reach control plane services.
func queryTenantReach(g *graph.CPG) []query.Finding {
	if g.ArchData == nil {
		return nil
	}

	var findings []query.Finding

	for _, np := range g.ArchData.NetworkPolicies {
		for _, rule := range np.IngressRules {
			selectors := extractNamespaceSelectors(rule)
			for _, sel := range selectors {
				// Check if this label key appears in contexts related to tenant namespaces
				labelSites := findLabelApplicationSites(g, sel.key)
				for _, site := range labelSites {
					if site.isTenant {
						hasPodSelector := hasPodSelectorInRule(rule)
						hasPorts := hasPortsInRule(rule)
						restriction := "none"
						if hasPodSelector {
							restriction = "podSelector"
						}
						if hasPorts {
							if restriction != "none" {
								restriction += " + ports"
							} else {
								restriction = "ports"
							}
						}

						severity := "high"
						if hasPodSelector && hasPorts {
							severity = "low"
						} else if hasPodSelector || hasPorts {
							severity = "medium"
						}

						findings = append(findings, query.Finding{
							RuleID:   "CGA-N02",
							Severity: severity,
							Message: fmt.Sprintf(
								"NetworkPolicy %q allows ingress from tenant namespace (%s: %s). Restriction: %s. Tenant workloads (%s) can reach control plane services.",
								np.Name, site.context, sel.key, restriction, site.tenantType,
							),
							File:   np.Source,
							Domain: "netpolicy",
						})
					}
				}
			}
		}
	}

	return findings
}

type labelSelector struct {
	key   string
	value string
}

type labelSite struct {
	file       string
	line       int
	context    string
	isTenant   bool
	tenantType string
}

func extractNamespaceSelectors(rule json.RawMessage) []labelSelector {
	var parsed map[string]interface{}
	if err := json.Unmarshal(rule, &parsed); err != nil {
		return nil
	}

	var selectors []labelSelector

	fromList, ok := parsed["from"].([]interface{})
	if !ok {
		return nil
	}

	for _, from := range fromList {
		fromMap, ok := from.(map[string]interface{})
		if !ok {
			continue
		}
		nsSel, ok := fromMap["namespaceSelector"].(map[string]interface{})
		if !ok {
			continue
		}
		matchLabels, ok := nsSel["matchLabels"].(map[string]interface{})
		if ok {
			for k, v := range matchLabels {
				selectors = append(selectors, labelSelector{key: k, value: fmt.Sprintf("%v", v)})
			}
		} else {
			// Empty namespaceSelector {} matches ALL namespaces
			selectors = append(selectors, labelSelector{key: "*", value: "(all namespaces)"})
		}
	}

	return selectors
}

func hasPodSelectorInRule(rule json.RawMessage) bool {
	var parsed map[string]interface{}
	if err := json.Unmarshal(rule, &parsed); err != nil {
		return false
	}

	fromList, ok := parsed["from"].([]interface{})
	if !ok {
		return false
	}

	for _, from := range fromList {
		fromMap, ok := from.(map[string]interface{})
		if !ok {
			continue
		}
		if _, ok := fromMap["podSelector"]; ok {
			return true
		}
	}
	return false
}

func hasPortsInRule(rule json.RawMessage) bool {
	var parsed map[string]interface{}
	if err := json.Unmarshal(rule, &parsed); err != nil {
		return false
	}

	ports, ok := parsed["ports"].([]interface{})
	return ok && len(ports) > 0
}

// findLabelApplicationSites searches the CPG for call sites and string literals
// that reference a given label key. It checks surrounding context for tenant
// namespace indicators.
func findLabelApplicationSites(g *graph.CPG, labelKey string) []labelSite {
	var sites []labelSite

	// Search all call sites and struct literals for the label key
	for _, cs := range g.NodesByKind(graph.NodeCallSite) {
		stringArgs := cs.Properties["string_args"]
		if !strings.Contains(stringArgs, labelKey) {
			continue
		}

		isTenant, tenantType := classifyTenantContext(cs.File, cs.Name, stringArgs)
		sites = append(sites, labelSite{
			file:       cs.File,
			line:       cs.Line,
			context:    cs.Name,
			isTenant:   isTenant,
			tenantType: tenantType,
		})
	}

	// Also check struct literals (label maps in Go)
	for _, sl := range g.NodesByKind(graph.NodeStructLiteral) {
		fieldNames := strings.Join(sl.FieldNames, ",")
		if !strings.Contains(fieldNames, labelKey) {
			continue
		}

		isTenant, tenantType := classifyTenantContext(sl.File, sl.Name, fieldNames)
		sites = append(sites, labelSite{
			file:       sl.File,
			line:       sl.Line,
			context:    sl.Name,
			isTenant:   isTenant,
			tenantType: tenantType,
		})
	}

	return sites
}

func classifyTenantContext(file, funcName, content string) (bool, string) {
	combined := strings.ToLower(file + " " + funcName + " " + content)
	for _, pattern := range tenantNamespacePatterns {
		if strings.Contains(combined, pattern) {
			return true, inferTenantType(combined)
		}
	}
	return false, ""
}

func inferTenantType(context string) string {
	switch {
	case strings.Contains(context, "notebook") || strings.Contains(context, "workbench"):
		return "jupyter notebooks"
	case strings.Contains(context, "pipeline"):
		return "data science pipelines"
	case strings.Contains(context, "serving") || strings.Contains(context, "inference"):
		return "model serving"
	case strings.Contains(context, "training"):
		return "training jobs"
	case strings.Contains(context, "sandbox") || strings.Contains(context, "user"):
		return "user workloads"
	default:
		return "tenant workloads"
	}
}
