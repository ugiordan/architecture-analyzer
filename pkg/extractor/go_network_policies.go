package extractor

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	// Matches networkingv1.NetworkPolicy{ or &networkingv1.NetworkPolicy{
	netpolStructRE = regexp.MustCompile(`(?:=|:=)\s*&?networkingv1\.NetworkPolicy\{`)

	// Matches NamespaceSelector with MatchLabels inline
	namespaceSelectorRE = regexp.MustCompile(`NamespaceSelector:\s*&metav1\.LabelSelector\{`)
	matchLabelsInlineRE = regexp.MustCompile(`MatchLabels:\s*map\[string\]string\{([^}]+)\}`)

	// Matches createNetworkPolicyPeer("key", "value") or similar helper patterns
	netpolPeerHelperRE = regexp.MustCompile(`(?:createNetworkPolicyPeer|NetworkPolicyPeer)\(\s*(?:(\w+(?:\.\w+)*)|"([^"]+)")\s*,\s*(?:(\w+(?:\.\w+)*)|"([^"]+)")`)

	// Matches PodSelector in NetworkPolicySpec
	podSelectorRE = regexp.MustCompile(`PodSelector:\s*metav1\.LabelSelector\{`)

	// Matches Ports in NetworkPolicyIngressRule
	netpolPortsRE = regexp.MustCompile(`Ports:\s*\[\]networkingv1\.NetworkPolicyPort\{`)

	// Matches ObjectMeta Name field
	netpolNameRE = regexp.MustCompile(`Name:\s*(?:(\w+(?:\.\w+)*)|"([^"]+)")`)
)

// extractGoNetworkPolicies scans Go source files for programmatic NetworkPolicy
// creation and returns them as NetworkPolicy entries.
func extractGoNetworkPolicies(repoPath string) []NetworkPolicy {
	goFiles := findFiles(repoPath, []string{"**/*.go"})

	var policies []NetworkPolicy

	for _, fpath := range goFiles {
		if strings.Contains(fpath, "_test.go") || strings.Contains(fpath, "/vendor/") {
			continue
		}

		info, err := os.Lstat(fpath)
		if err != nil || info.Mode()&os.ModeSymlink != 0 {
			continue
		}
		if info.Size() > maxFileSize {
			continue
		}
		data, err := os.ReadFile(fpath)
		if err != nil {
			continue
		}

		content := string(data)
		if !strings.Contains(content, "networkingv1.NetworkPolicy{") {
			continue
		}

		source := relativePath(repoPath, fpath)
		filePolicies := parseGoNetworkPolicies(content, source)
		policies = append(policies, filePolicies...)
	}

	return policies
}

func parseGoNetworkPolicies(content, source string) []NetworkPolicy {
	var policies []NetworkPolicy
	lines := strings.Split(content, "\n")

	// Find all NetworkPolicy struct literal starts
	for i, line := range lines {
		if !netpolStructRE.MatchString(line) {
			continue
		}
		// Skip if this is just Owns(&networkingv1.NetworkPolicy{}) or cache config
		trimmed := strings.TrimSpace(line)
		if strings.Contains(trimmed, "Owns(") || strings.Contains(trimmed, "Watches(") ||
			strings.Contains(trimmed, "For(") || strings.Contains(trimmed, "InformerFor") ||
			strings.Contains(trimmed, "func()") {
			continue
		}

		// Extract the block from this line to the matching closing brace
		block := extractGoBlock(lines, i)
		if block == "" {
			continue
		}

		np := NetworkPolicy{
			Source:      source + ":" + strconv.Itoa(i+1),
			PolicyTypes: []string{"Ingress"},
		}

		// Extract name
		if m := netpolNameRE.FindStringSubmatch(block); m != nil {
			if m[2] != "" {
				np.Name = m[2]
			} else if m[1] != "" {
				np.Name = m[1]
			}
		}
		if np.Name == "" {
			np.Name = "programmatic-netpol"
		}

		// Extract namespace selectors from helper calls
		var ingressRules []map[string]interface{}
		for _, m := range netpolPeerHelperRE.FindAllStringSubmatch(block, -1) {
			key := m[2]
			if key == "" {
				key = m[1]
			}
			value := m[4]
			if value == "" {
				value = m[3]
			}

			hasPodSelector := podSelectorRE.MatchString(block)
			hasPorts := netpolPortsRE.MatchString(block)

			rule := map[string]interface{}{
				"from": []interface{}{
					map[string]interface{}{
						"namespaceSelector": map[string]interface{}{
							"matchLabels": map[string]interface{}{key: value},
						},
					},
				},
			}
			if hasPorts {
				rule["ports"] = []interface{}{}
			}
			_ = hasPodSelector
			ingressRules = append(ingressRules, rule)
		}

		// Also check for inline MatchLabels (direct struct literal, not helper)
		for _, m := range matchLabelsInlineRE.FindAllStringSubmatch(block, -1) {
			pairs := parseInlineMap(m[1])
			for k, v := range pairs {
				rule := map[string]interface{}{
					"from": []interface{}{
						map[string]interface{}{
							"namespaceSelector": map[string]interface{}{
								"matchLabels": map[string]interface{}{k: v},
							},
						},
					},
				}
				ingressRules = append(ingressRules, rule)
			}
		}

		if len(ingressRules) > 0 {
			np.IngressRules = ingressRules
			np.Issues = append(np.Issues, "programmatic NetworkPolicy created in Go source")
			policies = append(policies, np)
		}
	}

	return policies
}

// extractGoBlock gets the Go code block starting from a line with an opening brace,
// up to the matching closing brace. Returns empty if no match within 50 lines.
func extractGoBlock(lines []string, startLine int) string {
	depth := 0
	var b strings.Builder
	for i := startLine; i < len(lines) && i < startLine+50; i++ {
		line := lines[i]
		b.WriteString(line)
		b.WriteByte('\n')
		depth += strings.Count(line, "{") - strings.Count(line, "}")
		if depth <= 0 && i > startLine {
			break
		}
	}
	return b.String()
}

// parseInlineMap extracts key-value pairs from a Go map literal string like:
// "key1": "value1", "key2": "value2"
func parseInlineMap(s string) map[string]string {
	result := make(map[string]string)
	re := regexp.MustCompile(`"([^"]+)":\s*"([^"]+)"`)
	for _, m := range re.FindAllStringSubmatch(s, -1) {
		result[m[1]] = m[2]
	}
	return result
}
