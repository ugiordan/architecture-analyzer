package layers

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ugiordan/architecture-analyzer/pkg/extractor"
	"github.com/ugiordan/architecture-analyzer/pkg/graph"
	"github.com/ugiordan/architecture-analyzer/pkg/query"
	"github.com/ugiordan/architecture-analyzer/pkg/srclang"
)

// safeJoin resolves repoPath+file and rejects paths that escape repoPath.
func safeJoin(repoPath, file string) (string, bool) {
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

var securityAnnotations = map[string]bool{
	"webhook_handler":   true,
	"auth_check":        true,
	"input_validator":   true,
	"crypto_operation":  true,
	"deserialization":   true,
	"sql_construction":  true,
	"command_execution": true,
	"file_access":       true,
	"network_policy":    true,
	"secret_access":     true,
	"taint_source":      true,
	"taint_sink":        true,
}

var securityNodeKinds = map[graph.NodeKind]bool{
	graph.NodeHTTPEndpoint: true,
	graph.NodeDBOperation:  true,
}

type SecuritySelector struct {
	repoPath string
	body     *srclang.BodyExtractor
}

func NewSecuritySelector(repoPath string) *SecuritySelector {
	return &SecuritySelector{
		repoPath: repoPath,
		body:     srclang.NewBodyExtractor(),
	}
}


func (s *SecuritySelector) Select(cpg *graph.CPG, arch *extractor.ComponentArchitecture, findings []query.Finding, extractionAnnotations []extractor.SecurityAnnotation) (*srclang.Layer, []srclang.Warning) {
	layer := &srclang.Layer{Name: "security"}
	var warnings []srclang.Warning

	layer.Findings = s.convertFindings(findings)
	layer.Findings = append(layer.Findings, convertSecurityAnnotations(extractionAnnotations)...)
	s.addSecurityFunctions(cpg, layer, &warnings)
	s.addFindingReferencedFunctions(cpg, layer, &warnings)
	if arch != nil {
		s.addRBAC(arch, layer)
		s.addNetworkPolicies(arch, layer)
	}
	s.addRelationships(cpg, layer)

	sortFindings(layer)
	capFunctions(layer)
	capCodeBodies(layer)
	return layer, warnings
}

func (s *SecuritySelector) addSecurityFunctions(cpg *graph.CPG, layer *srclang.Layer, warnings *[]srclang.Warning) {
	fileMap := make(map[string]*srclang.File)

	for _, node := range cpg.Nodes() {
		if !isSecurityRelevant(node) {
			continue
		}

		file := node.File
		if file == "" {
			continue
		}

		sf, ok := fileMap[file]
		if !ok {
			sf = &srclang.File{
				Path:     file,
				Language: node.Language,
			}
			fileMap[file] = sf
		}

		fn := srclang.Function{
			Name:       node.Name,
			Kind:       functionKind(node),
			SourceLine: node.Line,
		}

		if node.Complexity > 0 {
			fn.Complexity = node.Complexity
		}

		if node.EndLine > 0 && node.Line > 0 {
			fullPath, ok := safeJoin(s.repoPath,file)
			if !ok {
				*warnings = append(*warnings, srclang.Warning{
					File:    file,
					Message: fmt.Sprintf("path traversal attempt for %s", node.Name),
				})
			} else {
				code, err := s.body.Extract(fullPath, node.Line, node.EndLine)
				if err != nil {
					*warnings = append(*warnings, srclang.Warning{
						File:    file,
						Message: fmt.Sprintf("body extraction failed for %s: %v", node.Name, err),
					})
				} else {
					fn.Code = code
					fn.BodyLines = node.EndLine - node.Line + 1
				}
			}
		}

		if node.Annotations["taint_source"] {
			fn.TaintRole = "source"
			fn.Trust = "untrusted"
		} else if node.Annotations["taint_sink"] {
			fn.TaintRole = "sink"
		}

		fn.Params = extractParams(node)
		fn.Returns = extractReturns(node)
		if node.TypeName != "" {
			fn.ReceiverType = node.TypeName
		}

		sf.Functions = append(sf.Functions, fn)
	}

	for _, sf := range fileMap {
		sort.Slice(sf.Functions, func(i, j int) bool {
			return sf.Functions[i].SourceLine < sf.Functions[j].SourceLine
		})
		layer.Files = append(layer.Files, *sf)
	}
	sort.Slice(layer.Files, func(i, j int) bool {
		return layer.Files[i].Path < layer.Files[j].Path
	})
}

func (s *SecuritySelector) addFindingReferencedFunctions(cpg *graph.CPG, layer *srclang.Layer, warnings *[]srclang.Warning) {
	existingFuncs := make(map[string]bool)
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			existingFuncs[f.Path+":"+fmt.Sprintf("%d", fn.SourceLine)] = true
		}
	}

	findingLocations := make(map[string]int)
	for _, f := range layer.Findings {
		if f.SourceFile != "" && f.SourceLine > 0 {
			key := f.SourceFile + ":" + fmt.Sprintf("%d", f.SourceLine)
			findingLocations[key] = f.SourceLine
		}
	}

	fileMap := make(map[string]*srclang.File)
	for _, sf := range layer.Files {
		fileMap[sf.Path] = &sf
	}

	for _, node := range cpg.Nodes() {
		if node.Kind != graph.NodeFunction && node.Kind != graph.NodeHTTPEndpoint {
			continue
		}
		if node.File == "" || node.Line == 0 {
			continue
		}
		locKey := node.File + ":" + fmt.Sprintf("%d", node.Line)
		if existingFuncs[locKey] {
			continue
		}
		if _, referenced := findingLocations[locKey]; !referenced {
			if node.EndLine > 0 {
				for fl := node.Line; fl <= node.EndLine; fl++ {
					ck := node.File + ":" + fmt.Sprintf("%d", fl)
					if _, ok := findingLocations[ck]; ok {
						referenced = true
						break
					}
				}
			}
			if !referenced {
				continue
			}
		}

		sf, ok := fileMap[node.File]
		if !ok {
			newFile := &srclang.File{
				Path:     node.File,
				Language: node.Language,
			}
			fileMap[node.File] = newFile
			sf = newFile
		}

		fn := srclang.Function{
			Name:       node.Name,
			Kind:       functionKind(node),
			SourceLine: node.Line,
		}
		if node.Complexity > 0 {
			fn.Complexity = node.Complexity
		}
		if node.EndLine > 0 && node.Line > 0 {
			fullPath, ok := safeJoin(s.repoPath,node.File)
			if !ok {
				*warnings = append(*warnings, srclang.Warning{
					File:    node.File,
					Message: fmt.Sprintf("path traversal attempt for %s", node.Name),
				})
			} else {
				code, err := s.body.Extract(fullPath, node.Line, node.EndLine)
				if err != nil {
					*warnings = append(*warnings, srclang.Warning{
						File:    node.File,
						Message: fmt.Sprintf("body extraction failed for %s: %v", node.Name, err),
					})
				} else {
					fn.Code = code
					fn.BodyLines = node.EndLine - node.Line + 1
				}
			}
		}
		fn.Params = extractParams(node)
		fn.Returns = extractReturns(node)
		if node.TypeName != "" {
			fn.ReceiverType = node.TypeName
		}
		sf.Functions = append(sf.Functions, fn)
		existingFuncs[locKey] = true
	}

	layer.Files = nil
	for _, sf := range fileMap {
		sort.Slice(sf.Functions, func(i, j int) bool {
			return sf.Functions[i].SourceLine < sf.Functions[j].SourceLine
		})
		layer.Files = append(layer.Files, *sf)
	}
	sort.Slice(layer.Files, func(i, j int) bool {
		return layer.Files[i].Path < layer.Files[j].Path
	})
}

func (s *SecuritySelector) addRBAC(arch *extractor.ComponentArchitecture, layer *srclang.Layer) {
	if arch.RBAC == nil {
		return
	}
	bindingsByRole := s.buildBindingIndex(arch.RBAC)

	for _, cr := range arch.RBAC.ClusterRoles {
		res := srclang.Resource{
			Kind:       "ClusterRole",
			Name:       cr.Name,
			SourceFile: cr.Source,
			APIGroup:   "rbac.authorization.k8s.io",
			Scope:      "cluster",
			Summary:    fmt.Sprintf("%d rules", len(cr.Rules)),
		}
		res.Children = append(res.Children, rbacRuleChildren(cr.Rules)...)
		if cr.AggregationRule != nil {
			res.Children = append(res.Children, rbacAggregationChild(cr.AggregationRule))
		}
		res.Children = append(res.Children, rbacBindingChildren(bindingsByRole["ClusterRole/"+cr.Name])...)
		layer.Resources = append(layer.Resources, res)
	}
	for _, r := range arch.RBAC.Roles {
		res := srclang.Resource{
			Kind:       "Role",
			Name:       r.Name,
			SourceFile: r.Source,
			APIGroup:   "rbac.authorization.k8s.io",
			Scope:      "namespaced",
			Summary:    fmt.Sprintf("%d rules", len(r.Rules)),
		}
		res.Children = append(res.Children, rbacRuleChildren(r.Rules)...)
		res.Children = append(res.Children, rbacBindingChildren(bindingsByRole["Role/"+r.Name])...)
		layer.Resources = append(layer.Resources, res)
	}
}

func (s *SecuritySelector) buildBindingIndex(rbac *extractor.RBACData) map[string][]extractor.RBACBinding {
	index := make(map[string][]extractor.RBACBinding)
	for _, b := range rbac.ClusterRoleBindings {
		key := "ClusterRole/" + b.RoleRef
		index[key] = append(index[key], b)
	}
	for _, b := range rbac.RoleBindings {
		key := "Role/" + b.RoleRef
		index[key] = append(index[key], b)
		key2 := "ClusterRole/" + b.RoleRef
		index[key2] = append(index[key2], b)
	}
	return index
}

func rbacRuleChildren(rules []extractor.RBACRule) []srclang.ResourceChild {
	var children []srclang.ResourceChild
	for _, r := range rules {
		attrs := fmt.Sprintf(`apiGroups="%s" resources="%s" verbs="%s"`,
			xmlEscAttr(strings.Join(r.APIGroups, ",")),
			xmlEscAttr(strings.Join(r.Resources, ",")),
			xmlEscAttr(strings.Join(r.Verbs, ",")),
		)
		if len(r.ResourceNames) > 0 {
			attrs += fmt.Sprintf(` resourceNames="%s"`, xmlEscAttr(strings.Join(r.ResourceNames, ",")))
		}
		children = append(children, srclang.ResourceChild{
			XMLContent: fmt.Sprintf("<rule %s/>", attrs),
		})
	}
	return children
}

func rbacAggregationChild(labels map[string]string) srclang.ResourceChild {
	var parts []string
	for k, v := range labels {
		parts = append(parts, fmt.Sprintf("%s=%s", k, v))
	}
	sort.Strings(parts)
	return srclang.ResourceChild{
		XMLContent: fmt.Sprintf(`<aggregation-rule labels="%s"/>`, xmlEscAttr(strings.Join(parts, ","))),
	}
}

func rbacBindingChildren(bindings []extractor.RBACBinding) []srclang.ResourceChild {
	var children []srclang.ResourceChild
	for _, b := range bindings {
		for _, subj := range b.Subjects {
			subjStr := subj.Kind + ":" + subj.Name
			if subj.Namespace != "" {
				subjStr = subj.Kind + ":" + subj.Namespace + "/" + subj.Name
			}
			children = append(children, srclang.ResourceChild{
				XMLContent: fmt.Sprintf(`<binding name="%s" subject="%s" source="%s"/>`,
					xmlEscAttr(b.Name), xmlEscAttr(subjStr), xmlEscAttr(b.Source)),
			})
		}
	}
	return children
}

func xmlEscAttr(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, "\"", "&quot;")
	return s
}

func (s *SecuritySelector) addNetworkPolicies(arch *extractor.ComponentArchitecture, layer *srclang.Layer) {
	for _, np := range arch.NetworkPolicies {
		origin := "manifest"
		if strings.HasSuffix(np.Source, ".go") {
			origin = "programmatic"
		}
		res := srclang.Resource{
			Kind:       "NetworkPolicy",
			Name:       np.Name,
			SourceFile: np.Source,
			Origin:     origin,
		}
		if len(np.PolicyTypes) > 0 {
			res.Children = append(res.Children, srclang.ResourceChild{
				XMLContent: fmt.Sprintf(`<policy-types>%s</policy-types>`, xmlEscAttr(strings.Join(np.PolicyTypes, ","))),
			})
		}
		if np.PodSelector != nil {
			res.Children = append(res.Children, netpolSelectorChild("pod-selector", np.PodSelector))
		}
		for _, rule := range np.IngressRules {
			res.Children = append(res.Children, netpolRuleChild("ingress-rule", rule))
		}
		for _, rule := range np.EgressRules {
			res.Children = append(res.Children, netpolRuleChild("egress-rule", rule))
		}
		for _, issue := range np.Issues {
			res.Children = append(res.Children, srclang.ResourceChild{
				XMLContent: fmt.Sprintf(`<issue>%s</issue>`, xmlEscAttr(issue)),
			})
		}
		layer.Resources = append(layer.Resources, res)

		layer.Findings = append(layer.Findings, netpolRiskFindings(np)...)
	}
}

func netpolRiskFindings(np extractor.NetworkPolicy) []srclang.Finding {
	var findings []srclang.Finding
	counters := make(map[string]int)

	emit := func(rule, severity, title string) {
		counters[rule]++
		findings = append(findings, srclang.Finding{
			ID:          fmt.Sprintf("ext-%s-%03d", rule, counters[rule]),
			Domain:      "extraction",
			Severity:    severity,
			Rule:        rule,
			SourceFile:  np.Source,
			Title:       title,
			Description: title,
		})
	}

	for _, rule := range np.EgressRules {
		peers := extractPeers(rule, "to")
		for _, peer := range peers {
			if ipBlock, ok := peer["ipBlock"].(map[string]interface{}); ok {
				cidrStr := fmt.Sprintf("%v", ipBlock["cidr"])
				if cidrStr == "0.0.0.0/0" || cidrStr == "::/0" {
					ports := flattenPorts(rule["ports"])
					emit("NETPOL_UNRESTRICTED_EGRESS", "medium",
						fmt.Sprintf("NetworkPolicy %q allows egress to %s on ports %s", np.Name, cidrStr, ports))
				}
			}
		}
	}

	for _, rule := range np.IngressRules {
		peers := extractPeers(rule, "from")
		for _, peer := range peers {
			if ns, ok := peer["namespaceSelector"]; ok {
				nsMap, _ := ns.(map[string]interface{})
				if nsMap != nil {
					_, hasLabels := nsMap["matchLabels"]
					_, hasExpressions := nsMap["matchExpressions"]
					if !hasLabels && !hasExpressions {
						emit("NETPOL_CROSS_NAMESPACE_INGRESS", "medium",
							fmt.Sprintf("NetworkPolicy %q allows ingress from any namespace (empty namespaceSelector)", np.Name))
					}
				}
			}
		}
	}

	hasEgress := false
	for _, pt := range np.PolicyTypes {
		if strings.EqualFold(pt, "Egress") {
			hasEgress = true
		}
	}
	if len(np.PolicyTypes) > 0 && !hasEgress {
		emit("NETPOL_MISSING_EGRESS_POLICY", "low",
			fmt.Sprintf("NetworkPolicy %q does not include Egress in policyTypes (all egress allowed by default)", np.Name))
	}

	return findings
}

func extractPeers(rule map[string]interface{}, key string) []map[string]interface{} {
	v, ok := rule[key]
	if !ok {
		return nil
	}
	peers, ok := v.([]interface{})
	if !ok {
		return nil
	}
	var result []map[string]interface{}
	for _, p := range peers {
		if pm, ok := p.(map[string]interface{}); ok {
			result = append(result, pm)
		}
	}
	return result
}

func netpolSelectorChild(tag string, selector map[string]interface{}) srclang.ResourceChild {
	labels, _ := flattenLabels(selector)
	return srclang.ResourceChild{
		XMLContent: fmt.Sprintf(`<%s labels="%s"/>`, tag, xmlEscAttr(labels)),
	}
}

func netpolRuleChild(tag string, rule map[string]interface{}) srclang.ResourceChild {
	var attrs []string
	if from, ok := rule["from"]; ok {
		attrs = append(attrs, fmt.Sprintf(`from="%s"`, xmlEscAttr(flattenNetpolPeers(from))))
	}
	if to, ok := rule["to"]; ok {
		attrs = append(attrs, fmt.Sprintf(`to="%s"`, xmlEscAttr(flattenNetpolPeers(to))))
	}
	if ports, ok := rule["ports"]; ok {
		attrs = append(attrs, fmt.Sprintf(`ports="%s"`, xmlEscAttr(flattenPorts(ports))))
	}
	return srclang.ResourceChild{
		XMLContent: fmt.Sprintf(`<%s %s/>`, tag, strings.Join(attrs, " ")),
	}
}

func flattenLabels(selector map[string]interface{}) (string, bool) {
	ml, ok := selector["matchLabels"]
	if !ok {
		return "{}", false
	}
	labels, ok := ml.(map[string]interface{})
	if !ok {
		return "{}", false
	}
	var parts []string
	for k, v := range labels {
		parts = append(parts, fmt.Sprintf("%s=%v", k, v))
	}
	sort.Strings(parts)
	return strings.Join(parts, ","), true
}

func flattenNetpolPeers(v interface{}) string {
	peers, ok := v.([]interface{})
	if !ok {
		return fmt.Sprintf("%v", v)
	}
	var parts []string
	for _, p := range peers {
		pm, ok := p.(map[string]interface{})
		if !ok {
			continue
		}
		if ns, ok := pm["namespaceSelector"]; ok {
			nsMap, _ := ns.(map[string]interface{})
			labels, _ := flattenLabels(nsMap)
			if labels == "{}" {
				parts = append(parts, "namespace:*")
			} else {
				parts = append(parts, "namespace:"+labels)
			}
		}
		if pod, ok := pm["podSelector"]; ok {
			podMap, _ := pod.(map[string]interface{})
			labels, _ := flattenLabels(podMap)
			if labels == "{}" {
				parts = append(parts, "pod:*")
			} else {
				parts = append(parts, "pod:"+labels)
			}
		}
		if cidr, ok := pm["ipBlock"]; ok {
			cb, _ := cidr.(map[string]interface{})
			if c, ok := cb["cidr"]; ok {
				parts = append(parts, fmt.Sprintf("cidr:%v", c))
			}
		}
	}
	if len(parts) == 0 {
		return "any"
	}
	return strings.Join(parts, ";")
}

func flattenPorts(v interface{}) string {
	var items []interface{}
	switch vt := v.(type) {
	case []interface{}:
		items = vt
	default:
		b, err := json.Marshal(v)
		if err != nil {
			return fmt.Sprintf("%v", v)
		}
		if err := json.Unmarshal(b, &items); err != nil {
			return string(b)
		}
	}
	var parts []string
	for _, p := range items {
		pm, ok := p.(map[string]interface{})
		if !ok {
			continue
		}
		port := fmt.Sprintf("%v", pm["port"])
		if f, ok := pm["port"].(float64); ok {
			port = fmt.Sprintf("%d", int(f))
		}
		if proto, ok := pm["protocol"]; ok {
			port += "/" + fmt.Sprintf("%v", proto)
		}
		parts = append(parts, port)
	}
	if len(parts) == 0 {
		return fmt.Sprintf("%v", v)
	}
	return strings.Join(parts, ",")
}

const maxRelationships = 150
const maxFunctions = 200
const maxCodeBodies = 60
const maxCodeBytesTotal = 200_000

func capFunctions(layer *srclang.Layer) {
	total := 0
	for _, f := range layer.Files {
		total += len(f.Functions)
	}
	if total <= maxFunctions {
		return
	}

	type funcEntry struct {
		fileIdx int
		funcIdx int
		score   int
	}
	var entries []funcEntry
	for fi, f := range layer.Files {
		for fni, fn := range f.Functions {
			score := fn.Complexity
			if fn.TaintRole != "" {
				score += 20
			}
			if fn.Trust == "untrusted" {
				score += 10
			}
			if fn.Code != "" {
				score += 5
			}
			if score == 0 {
				score = 1
			}
			entries = append(entries, funcEntry{fileIdx: fi, funcIdx: fni, score: score})
		}
	}
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].score != entries[j].score {
			return entries[i].score > entries[j].score
		}
		if entries[i].fileIdx != entries[j].fileIdx {
			return entries[i].fileIdx < entries[j].fileIdx
		}
		return entries[i].funcIdx < entries[j].funcIdx
	})

	keep := make(map[string]bool)
	for i, e := range entries {
		if i >= maxFunctions {
			break
		}
		key := fmt.Sprintf("%d:%d", e.fileIdx, e.funcIdx)
		keep[key] = true
	}

	for fi := range layer.Files {
		var kept []srclang.Function
		for fni, fn := range layer.Files[fi].Functions {
			key := fmt.Sprintf("%d:%d", fi, fni)
			if keep[key] {
				kept = append(kept, fn)
			}
		}
		layer.Files[fi].Functions = kept
	}

	// Remove empty files
	var files []srclang.File
	for _, f := range layer.Files {
		if len(f.Functions) > 0 {
			files = append(files, f)
		}
	}
	layer.Files = files
}

func capCodeBodies(layer *srclang.Layer) {
	type funcRef struct {
		fileIdx int
		funcIdx int
		score   int
		codeLen int
	}

	var refs []funcRef
	for fi, f := range layer.Files {
		for fni, fn := range f.Functions {
			if fn.Code == "" {
				continue
			}
			score := fn.Complexity
			if fn.TaintRole == "source" || fn.TaintRole == "sink" {
				score += 20
			}
			if fn.Trust == "untrusted" {
				score += 10
			}
			if score == 0 {
				score = 1
			}
			refs = append(refs, funcRef{
				fileIdx: fi,
				funcIdx: fni,
				score:   score,
				codeLen: len(fn.Code),
			})
		}
	}

	if len(refs) <= maxCodeBodies {
		totalBytes := 0
		for _, r := range refs {
			totalBytes += r.codeLen
		}
		if totalBytes <= maxCodeBytesTotal {
			return
		}
	}

	sort.Slice(refs, func(i, j int) bool {
		if refs[i].score != refs[j].score {
			return refs[i].score > refs[j].score
		}
		if refs[i].fileIdx != refs[j].fileIdx {
			return refs[i].fileIdx < refs[j].fileIdx
		}
		return refs[i].funcIdx < refs[j].funcIdx
	})

	totalBytes := 0
	for i, r := range refs {
		if i >= maxCodeBodies || totalBytes+r.codeLen > maxCodeBytesTotal {
			layer.Files[r.fileIdx].Functions[r.funcIdx].Code = ""
			layer.Files[r.fileIdx].Functions[r.funcIdx].BodyLines = 0
		} else {
			totalBytes += r.codeLen
		}
	}
}

func (s *SecuritySelector) addRelationships(cpg *graph.CPG, layer *srclang.Layer) {
	selectedFuncs := make(map[string]bool)
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			key := f.Path + ":" + fn.Name + ":" + fmt.Sprintf("%d", fn.SourceLine)
			selectedFuncs[key] = true
		}
	}

	var bothSelected []srclang.Relationship
	var oneSelected []srclang.Relationship

	for _, edge := range cpg.Edges() {
		if edge.Kind != graph.EdgeCalls {
			continue
		}
		fromNode := cpg.GetNode(edge.From)
		toNode := cpg.GetNode(edge.To)
		if fromNode == nil || toNode == nil {
			continue
		}
		fromKey := fromNode.File + ":" + fromNode.Name + ":" + fmt.Sprintf("%d", fromNode.Line)
		toKey := toNode.File + ":" + toNode.Name + ":" + fmt.Sprintf("%d", toNode.Line)

		fromSelected := selectedFuncs[fromKey]
		toSelected := selectedFuncs[toKey]
		if !fromSelected && !toSelected {
			continue
		}

		rel := srclang.Relationship{
			Kind: "calls",
			From: srclang.Endpoint{
				Function: fromNode.Name,
				File:     fromNode.File,
				Line:     fromNode.Line,
			},
			To: srclang.Endpoint{
				Function: toNode.Name,
				File:     toNode.File,
				Line:     toNode.Line,
			},
		}

		if !fromSelected {
			resolved := false
			rel.From.Resolved = &resolved
		}
		if !toSelected {
			resolved := false
			rel.To.Resolved = &resolved
		}

		if fromSelected && toSelected {
			bothSelected = append(bothSelected, rel)
		} else {
			oneSelected = append(oneSelected, rel)
		}
	}

	if len(bothSelected) > maxRelationships {
		bothSelected = bothSelected[:maxRelationships]
	}
	layer.Relationships = append(layer.Relationships, bothSelected...)
	remaining := maxRelationships - len(bothSelected)
	if remaining > 0 && len(oneSelected) > 0 {
		if len(oneSelected) > remaining {
			oneSelected = oneSelected[:remaining]
		}
		layer.Relationships = append(layer.Relationships, oneSelected...)
	}
}

func (s *SecuritySelector) convertFindings(findings []query.Finding) []srclang.Finding {
	var result []srclang.Finding
	counters := make(map[string]int)
	for _, f := range findings {
		if f.Domain != "security" && f.Domain != "netpolicy" && !strings.HasPrefix(f.Domain, "external/") {
			continue
		}
		counters[f.RuleID]++
		result = append(result, srclang.Finding{
			ID:          fmt.Sprintf("%s-%03d", f.RuleID, counters[f.RuleID]),
			Domain:      f.Domain,
			Severity:    f.Severity,
			Rule:        f.RuleID,
			SourceFile:  f.File,
			SourceLine:  f.Line,
			Title:       f.Message,
			Description: f.Message,
		})
	}
	return result
}

func convertSecurityAnnotations(annotations []extractor.SecurityAnnotation) []srclang.Finding {
	var result []srclang.Finding
	counters := make(map[string]int)
	for _, a := range annotations {
		counters[a.Type]++
		result = append(result, srclang.Finding{
			ID:          fmt.Sprintf("ext-%s-%03d", a.Type, counters[a.Type]),
			Domain:      "extraction",
			Severity:    a.Severity,
			Rule:        a.Type,
			SourceFile:  a.Source,
			Title:       a.Description,
			Description: a.Description,
		})
	}
	return result
}

func isSecurityRelevant(node *graph.Node) bool {
	if securityNodeKinds[node.Kind] {
		return true
	}
	if node.Kind != graph.NodeFunction {
		return false
	}
	for ann := range node.Annotations {
		if securityAnnotations[ann] {
			return true
		}
	}
	return false
}

func functionKind(node *graph.Node) string {
	if node.Kind == graph.NodeHTTPEndpoint {
		return "handler"
	}
	if node.TypeName != "" {
		return "method"
	}
	return "function"
}

func extractParams(node *graph.Node) []srclang.Param {
	var params []srclang.Param
	for i, name := range node.ParamNames {
		p := srclang.Param{Name: name}
		if i < len(node.ParamTypes) {
			p.Type = node.ParamTypes[i]
		}
		params = append(params, p)
	}
	return params
}

func extractReturns(node *graph.Node) []srclang.Return {
	if node.ReturnType == "" {
		return nil
	}
	return []srclang.Return{{Type: node.ReturnType}}
}

const maxFindings = 200

func sortFindings(layer *srclang.Layer) {
	severityOrder := map[string]int{
		"critical": 0, "high": 1, "medium": 2, "low": 3, "info": 4,
	}
	sort.SliceStable(layer.Findings, func(i, j int) bool {
		si, ok := severityOrder[layer.Findings[i].Severity]
		if !ok {
			si = 99
		}
		sj, ok2 := severityOrder[layer.Findings[j].Severity]
		if !ok2 {
			sj = 99
		}
		return si < sj
	})
	if len(layer.Findings) > maxFindings {
		layer.Findings = layer.Findings[:maxFindings]
	}
}
