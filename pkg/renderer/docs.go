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
		// Build a focused service diagram
		b.WriteString("## Service Map\n\n")
		b.WriteString("```mermaid\n")
		b.WriteString("graph LR\n")
		b.WriteString("    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff\n")
		b.WriteString("    classDef component fill:#3498db,stroke:#2980b9,color:#fff\n\n")
		compID := sanitizeID(component)
		b.WriteString(fmt.Sprintf("    %s[\"%s\"]:::%s\n", compID, escapeLabel(component), "component"))
		for i, svc := range services {
			name := getStr(svc, "name", "")
			svcType := getStr(svc, "type", "ClusterIP")
			ports := getSlice(svc, "ports")
			var portParts []string
			for _, p := range ports {
				portParts = append(portParts, fmt.Sprintf("%v/%s",
					p["port"], getStr(p, "protocol", "TCP")))
			}
			svcID := fmt.Sprintf("svc_%d", i)
			b.WriteString(fmt.Sprintf("    %s --> %s[\"%s\\n%s: %s\"]:::svc\n",
				compID, svcID, escapeLabel(name), svcType, strings.Join(portParts, ", ")))
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

	// Inline RBAC diagram
	rbacDiagram := (&RBACRenderer{}).Render(data)
	if rbacDiagram != "" && !strings.Contains(rbacDiagram, "No RBAC") {
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
