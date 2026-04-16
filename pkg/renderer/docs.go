package renderer

import (
	"fmt"
	"strings"
)

// DocsPage represents a single documentation page.
type DocsPage struct {
	Path    string // relative path, e.g. "index.md" or "components/kserve/index.md"
	Content string
}

// RenderDocs generates a complete set of documentation pages from architecture
// data. It auto-detects whether the input is a single component or aggregated
// platform data (by checking for the "platform" key). Returns a slice of pages
// with relative paths suitable for embedding in a docs site.
func RenderDocs(data map[string]interface{}) []DocsPage {
	if _, ok := data["platform"]; ok {
		return renderPlatformDocs(data)
	}
	return renderComponentDocs(data, "")
}

// NavSnippet generates a YAML navigation snippet for mkdocs.yml from the
// rendered docs pages.
func NavSnippet(pages []DocsPage, prefix string) string {
	var b strings.Builder
	for _, p := range pages {
		path := p.Path
		if prefix != "" {
			path = prefix + "/" + path
		}
		// Extract title from first H1 line
		title := extractTitle(p.Content)
		b.WriteString(fmt.Sprintf("    - %s: %s\n", title, path))
	}
	return b.String()
}

func extractTitle(content string) string {
	for _, line := range strings.SplitN(content, "\n", 5) {
		if strings.HasPrefix(line, "# ") {
			return strings.TrimPrefix(line, "# ")
		}
	}
	return "Untitled"
}

// renderComponentDocs generates docs pages for a single component.
func renderComponentDocs(data map[string]interface{}, pathPrefix string) []DocsPage {
	component := getStr(data, "component", "unknown")
	var pages []DocsPage

	prefix := pathPrefix
	if prefix != "" && !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}

	pages = append(pages, DocsPage{
		Path:    prefix + "index.md",
		Content: renderComponentIndexPage(data, component),
	})
	pages = append(pages, DocsPage{
		Path:    prefix + "network.md",
		Content: renderComponentNetworkPage(data, component),
	})
	pages = append(pages, DocsPage{
		Path:    prefix + "rbac.md",
		Content: renderComponentRBACPage(data, component),
	})
	pages = append(pages, DocsPage{
		Path:    prefix + "security.md",
		Content: renderComponentSecurityPage(data, component),
	})
	pages = append(pages, DocsPage{
		Path:    prefix + "dataflow.md",
		Content: renderComponentDataflowPage(data, component),
	})

	return pages
}

// --- Component index page ---

func renderComponentIndexPage(data map[string]interface{}, component string) string {
	var b strings.Builder

	repo := getStr(data, "repo", "")
	version := getStr(data, "analyzer_version", "")
	extractedAt := getStr(data, "extracted_at", "")

	b.WriteString(fmt.Sprintf("# %s\n\n", component))
	if repo != "" {
		b.WriteString(fmt.Sprintf("**Repository:** %s  \n", repo))
	}
	b.WriteString(fmt.Sprintf("**Analyzer:** rhoai-analyzer %s  \n", version))
	b.WriteString(fmt.Sprintf("**Extracted:** %s\n\n", extractedAt))

	// Summary table
	crds := getSlice(data, "crds")
	services := getSlice(data, "services")
	secrets := getSlice(data, "secrets_referenced")
	rbac := getMap(data, "rbac")
	deployments := getSlice(data, "deployments")
	watches := getSlice(data, "controller_watches")

	clusterRoleCount := 0
	if rbac != nil {
		clusterRoleCount = len(getSlice(rbac, "cluster_roles"))
	}

	b.WriteString("## Summary\n\n")
	b.WriteString("| Metric | Count |\n")
	b.WriteString("|--------|-------|\n")
	b.WriteString(fmt.Sprintf("| CRDs | %d |\n", len(crds)))
	b.WriteString(fmt.Sprintf("| Deployments | %d |\n", len(deployments)))
	b.WriteString(fmt.Sprintf("| Services | %d |\n", len(services)))
	b.WriteString(fmt.Sprintf("| Secrets | %d |\n", len(secrets)))
	b.WriteString(fmt.Sprintf("| Cluster Roles | %d |\n", clusterRoleCount))
	b.WriteString(fmt.Sprintf("| Controller Watches | %d |\n", len(watches)))
	b.WriteString("\n")

	// Component architecture diagram (inline mermaid)
	b.WriteString("## Component Architecture\n\n")
	b.WriteString("CRDs, controllers, and owned Kubernetes resources.\n\n")
	componentDiagram := (&ComponentRenderer{}).Render(data)
	b.WriteString("```mermaid\n")
	b.WriteString(componentDiagram)
	b.WriteString("\n```\n\n")

	// CRD table
	renderCRDTable(&b, data)

	// Dependencies
	b.WriteString("## Dependencies\n\n")
	renderDependencySection(&b, data)

	return b.String()
}

// --- Component network page ---

func renderComponentNetworkPage(data map[string]interface{}, component string) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("# %s: Network\n\n", component))

	services := getSlice(data, "services")
	netpols := getSlice(data, "network_policies")

	if len(services) > 0 {
		// Dedup services by (name, type, ports) to avoid test fixture duplicates
		type svcKey struct {
			name, svcType, ports string
		}
		type dedupedSvc struct {
			svc    map[string]interface{}
			count  int
			isTest bool
		}
		seen := make(map[svcKey]*dedupedSvc)
		var order []svcKey
		for _, svc := range services {
			name := getStr(svc, "name", "")
			svcType := getStr(svc, "type", "ClusterIP")
			ports := getSlice(svc, "ports")
			var portParts []string
			for _, p := range ports {
				portParts = append(portParts, fmt.Sprintf("%v/%s", p["port"], getStr(p, "protocol", "TCP")))
			}
			key := svcKey{name, svcType, strings.Join(portParts, ",")}
			source := getStr(svc, "source", "")
			isTest := strings.Contains(source, "/test/") || strings.Contains(source, "/e2e/") || strings.Contains(source, "/testdata/")
			if existing, ok := seen[key]; ok {
				existing.count++
				if isTest {
					existing.isTest = true
				}
			} else {
				seen[key] = &dedupedSvc{svc: svc, count: 1, isTest: isTest}
				order = append(order, key)
			}
		}

		// Only show unique services in the diagram
		b.WriteString("## Service Map\n\n")
		if len(order) < len(services) {
			b.WriteString(fmt.Sprintf("*%d unique services (%d total, duplicates from test fixtures collapsed).*\n\n", len(order), len(services)))
		}
		b.WriteString("```mermaid\n")
		b.WriteString("graph LR\n")
		b.WriteString("    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff\n")
		b.WriteString("    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff\n")
		b.WriteString("    classDef component fill:#3498db,stroke:#2980b9,color:#fff\n\n")
		compID := sanitizeID(component)
		b.WriteString(fmt.Sprintf("    %s[\"%s\"]:::%s\n", compID, escapeLabel(component), "component"))
		for i, key := range order {
			ds := seen[key]
			svcID := fmt.Sprintf("svc_%d", i)
			cls := "svc"
			if ds.isTest {
				cls = "test"
			}
			b.WriteString(fmt.Sprintf("    %s --> %s[\"%s\\n%s: %s\"]:::%s\n",
				compID, svcID, escapeLabel(key.name), key.svcType, key.ports, cls))
		}
		b.WriteString("```\n\n")
	}

	renderServiceTable(&b, data)
	renderIngressTable(&b, data)
	renderNetworkPolicyTable(&b, data)

	if len(netpols) == 0 {
		b.WriteString("!!! warning \"No Network Policies\"\n")
		b.WriteString("    No NetworkPolicy resources found. All pod-to-pod traffic is allowed by default.\n\n")
	}

	return b.String()
}

// --- Component RBAC page ---

func renderComponentRBACPage(data map[string]interface{}, component string) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("# %s: RBAC\n\n", component))

	// Check if RBAC data has actual content
	rbac := getMap(data, "rbac")
	hasBindings := rbac != nil && (len(getSlice(rbac, "cluster_role_bindings")) > 0 || len(getSlice(rbac, "role_bindings")) > 0)
	hasRoles := rbac != nil && (len(getSlice(rbac, "cluster_roles")) > 0 || len(getSlice(rbac, "roles")) > 0)

	if !hasBindings && !hasRoles {
		b.WriteString("!!! info \"No RBAC Resources\"\n")
		b.WriteString("    This component does not define any ClusterRoles, Roles, or RoleBindings.\n")
		b.WriteString("    It may operate as a sidecar or library used by other components that define their own RBAC.\n\n")
		return b.String()
	}

	// Count how many nodes the RBAC diagram would produce
	rbacDiagram := (&RBACRenderer{}).Render(data)
	lineCount := strings.Count(rbacDiagram, "\n")

	if lineCount > 60 {
		// Too large for a diagram: show a summary instead
		b.WriteString("## RBAC Summary\n\n")
		b.WriteString(fmt.Sprintf("This component defines a large RBAC surface (%d rules). The table below summarizes permissions by role.\n\n", lineCount))

		// Summary table: role -> resource count, verb breadth
		if hasRoles {
			type roleSummary struct {
				name      string
				kind      string
				resources int
				hasAdmin  bool
			}
			var summaries []roleSummary
			for _, role := range getSlice(rbac, "cluster_roles") {
				name := getStr(role, "name", "")
				rules := getSlice(role, "rules")
				resCount := 0
				admin := false
				for _, rule := range rules {
					resources := getStringSlice(rule, "resources")
					resCount += len(resources)
					verbs := getStringSlice(rule, "verbs")
					for _, v := range verbs {
						if v == "*" {
							admin = true
						}
					}
				}
				summaries = append(summaries, roleSummary{name, "ClusterRole", resCount, admin})
			}
			for _, role := range getSlice(rbac, "roles") {
				name := getStr(role, "name", "")
				rules := getSlice(role, "rules")
				resCount := 0
				for _, rule := range rules {
					resCount += len(getStringSlice(rule, "resources"))
				}
				summaries = append(summaries, roleSummary{name, "Role", resCount, false})
			}

			b.WriteString("| Role | Kind | Resources | Wildcard |\n")
			b.WriteString("|------|------|-----------|----------|\n")
			for _, s := range summaries {
				wildcard := ""
				if s.hasAdmin {
					wildcard = "yes"
				}
				b.WriteString(fmt.Sprintf("| %s | %s | %d | %s |\n", s.name, s.kind, s.resources, wildcard))
			}
			b.WriteString("\n")
		}

		if hasBindings {
			b.WriteString("### Bindings\n\n")
			b.WriteString("| Binding | Type | Role | Subject |\n")
			b.WriteString("|---------|------|------|---------|\n")
			for _, binding := range getSlice(rbac, "cluster_role_bindings") {
				name := getStr(binding, "name", "")
				roleRef := getStr(binding, "role_ref", "")
				for _, subj := range getSlice(binding, "subjects") {
					b.WriteString(fmt.Sprintf("| %s | ClusterRoleBinding | %s | %s/%s |\n",
						name, roleRef, getStr(subj, "kind", ""), getStr(subj, "name", "")))
				}
			}
			for _, binding := range getSlice(rbac, "role_bindings") {
				name := getStr(binding, "name", "")
				roleRef := getStr(binding, "role_ref", "")
				for _, subj := range getSlice(binding, "subjects") {
					b.WriteString(fmt.Sprintf("| %s | RoleBinding | %s | %s/%s |\n",
						name, roleRef, getStr(subj, "kind", ""), getStr(subj, "name", "")))
				}
			}
			b.WriteString("\n")
		}

		// Full details in expandable section
		b.WriteString("<details>\n<summary>Full RBAC hierarchy diagram</summary>\n\n")
		b.WriteString("```mermaid\n")
		b.WriteString(rbacDiagram)
		b.WriteString("\n```\n\n</details>\n\n")
	} else {
		// Small enough for inline diagram
		b.WriteString("## RBAC Hierarchy\n\n")
		b.WriteString("ServiceAccount bindings, roles, and resource permissions.\n\n")
		b.WriteString("```mermaid\n")
		b.WriteString(rbacDiagram)
		b.WriteString("\n```\n\n")
	}

	renderRBACSection(&b, data)
	return b.String()
}

// --- Component security page ---

func renderComponentSecurityPage(data map[string]interface{}, component string) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("# %s: Security\n\n", component))

	// Secrets
	b.WriteString("## Secrets\n\n")
	renderSecretTable(&b, data)

	// Deployment security
	b.WriteString("## Deployment Security Controls\n\n")
	renderSecurityContextSection(&b, data)

	// Dockerfiles
	b.WriteString("## Build Security\n\n")
	renderDockerfileTable(&b, data)

	// Cache architecture (OOM risks)
	renderCacheSection(&b, data)

	return b.String()
}

// --- Component dataflow page ---

func renderComponentDataflowPage(data map[string]interface{}, component string) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("# %s: Dataflow\n\n", component))

	// Controller watch table
	b.WriteString("## Controller Watches\n\n")
	renderControllerWatchTable(&b, data)

	// Dataflow diagram
	dataflowDiagram := (&DataflowRenderer{}).Render(data)
	if dataflowDiagram != "" {
		b.WriteString("## Reconciliation Flow\n\n")
		b.WriteString("How the controller interacts with the Kubernetes API during reconciliation.\n\n")
		b.WriteString("```mermaid\n")
		b.WriteString(dataflowDiagram)
		b.WriteString("\n```\n\n")
	}

	// Webhooks
	renderWebhookTable(&b, data)
	renderHTTPEndpointTable(&b, data)

	// Configuration
	b.WriteString("## Configuration\n\n")
	renderConfigMapTable(&b, data)
	renderHelmSection(&b, data)

	return b.String()
}
