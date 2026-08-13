package layers

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ugiordan/architecture-analyzer/pkg/extractor"
	"github.com/ugiordan/architecture-analyzer/pkg/graph"
	"github.com/ugiordan/architecture-analyzer/pkg/query"
	"github.com/ugiordan/architecture-analyzer/pkg/srclang"
)

const (
	annotNetPolSelector    = "netpol:namespace_selector"
	annotNetPolTenantReach = "netpol:tenant_reach"
	annotNetPolNoRestrict  = "netpol:no_restriction"
)

type NetpolicySelector struct {
	repoPath string
	body     *srclang.BodyExtractor
}

func NewNetpolicySelector(repoPath string) *NetpolicySelector {
	return &NetpolicySelector{
		repoPath: repoPath,
		body:     srclang.NewBodyExtractor(),
	}
}


func (s *NetpolicySelector) Select(cpg *graph.CPG, arch *extractor.ComponentArchitecture, findings []query.Finding, _ []extractor.SecurityAnnotation) (*srclang.Layer, []srclang.Warning) {
	layer := &srclang.Layer{Name: "netpolicy"}
	var warnings []srclang.Warning

	layer.Findings = s.convertFindings(findings)

	if arch != nil {
		s.addNetworkPolicyResources(arch, layer)
	}
	s.addNetpolicyFunctions(cpg, layer, &warnings)
	s.addFindingReferencedFunctions(cpg, layer, &warnings)
	s.addRelationships(cpg, layer)

	sortFindings(layer)
	capFunctions(layer)
	capCodeBodies(layer)
	return layer, warnings
}

func (s *NetpolicySelector) addNetworkPolicyResources(arch *extractor.ComponentArchitecture, layer *srclang.Layer) {
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
			labels, _ := flattenLabels(np.PodSelector)
			res.Children = append(res.Children, srclang.ResourceChild{
				XMLContent: fmt.Sprintf(`<pod-selector labels="%s"/>`, xmlEscAttr(labels)),
			})
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

		res.Summary = s.netpolicySummary(np)
		layer.Resources = append(layer.Resources, res)

		layer.Findings = append(layer.Findings, netpolRiskFindings(np)...)
	}
}

func (s *NetpolicySelector) netpolicySummary(np extractor.NetworkPolicy) string {
	var parts []string
	if len(np.PolicyTypes) > 0 {
		parts = append(parts, strings.Join(np.PolicyTypes, "+"))
	}
	parts = append(parts, fmt.Sprintf("%d ingress, %d egress", len(np.IngressRules), len(np.EgressRules)))
	if len(np.Issues) > 0 {
		parts = append(parts, fmt.Sprintf("%d issues", len(np.Issues)))
	}
	return strings.Join(parts, ", ")
}

func (s *NetpolicySelector) addNetpolicyFunctions(cpg *graph.CPG, layer *srclang.Layer, warnings *[]srclang.Warning) {
	fileMap := make(map[string]*srclang.File)

	for _, node := range cpg.Nodes() {
		if !isNetpolicyRelevant(node) {
			continue
		}
		if node.File == "" {
			continue
		}

		sf, ok := fileMap[node.File]
		if !ok {
			sf = &srclang.File{
				Path:     node.File,
				Language: node.Language,
			}
			fileMap[node.File] = sf
		}

		fn := s.buildNetpolicyFunction(node, warnings)
		fn.Metas = s.netpolicyMetas(node)
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

func (s *NetpolicySelector) addFindingReferencedFunctions(cpg *graph.CPG, layer *srclang.Layer, warnings *[]srclang.Warning) {
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
	for i := range layer.Files {
		fileMap[layer.Files[i].Path] = &layer.Files[i]
	}

	for _, node := range cpg.Nodes() {
		if node.Kind != graph.NodeFunction {
			continue
		}
		if node.File == "" || node.Line == 0 {
			continue
		}
		locKey := node.File + ":" + fmt.Sprintf("%d", node.Line)
		if existingFuncs[locKey] {
			continue
		}

		referenced := false
		if _, ok := findingLocations[locKey]; ok {
			referenced = true
		}
		if !referenced && node.EndLine > 0 && node.EndLine >= node.Line {
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

		sf, ok := fileMap[node.File]
		if !ok {
			newFile := &srclang.File{
				Path:     node.File,
				Language: node.Language,
			}
			fileMap[node.File] = newFile
			sf = newFile
		}

		fn := s.buildNetpolicyFunction(node, warnings)
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

func (s *NetpolicySelector) addRelationships(cpg *graph.CPG, layer *srclang.Layer) {
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

func (s *NetpolicySelector) buildNetpolicyFunction(node *graph.Node, warnings *[]srclang.Warning) srclang.Function {
	fn := srclang.Function{
		Name:       node.Name,
		Kind:       functionKind(node),
		SourceLine: node.Line,
	}

	if node.Complexity > 0 {
		fn.Complexity = node.Complexity
	}

	if node.EndLine > 0 && node.Line > 0 && node.EndLine >= node.Line {
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

	return fn
}

func (s *NetpolicySelector) netpolicyMetas(node *graph.Node) []srclang.Meta {
	var metas []srclang.Meta

	if node.Annotations[annotNetPolSelector] {
		metas = append(metas, srclang.Meta{Domain: "netpolicy", Key: "namespace-selector", Value: "true"})
	}

	if node.Annotations[annotNetPolTenantReach] {
		metas = append(metas, srclang.Meta{Domain: "netpolicy", Key: "tenant-reach", Value: "true"})
	}

	if node.Annotations[annotNetPolNoRestrict] {
		metas = append(metas, srclang.Meta{Domain: "netpolicy", Key: "no-restriction", Value: "true"})
	}

	return metas
}

func isNetpolicyRelevant(node *graph.Node) bool {
	if node.Kind != graph.NodeFunction {
		return false
	}
	for _, annot := range []string{
		annotNetPolSelector,
		annotNetPolTenantReach,
		annotNetPolNoRestrict,
	} {
		if node.Annotations[annot] {
			return true
		}
	}
	return false
}

func (s *NetpolicySelector) convertFindings(findings []query.Finding) []srclang.Finding {
	var result []srclang.Finding
	counters := make(map[string]int)
	for _, f := range findings {
		if f.Domain != "netpolicy" {
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
