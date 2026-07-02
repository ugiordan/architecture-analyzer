package layers

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ugiordan/architecture-analyzer/pkg/extractor"
	"github.com/ugiordan/architecture-analyzer/pkg/graph"
	"github.com/ugiordan/architecture-analyzer/pkg/query"
	"github.com/ugiordan/architecture-analyzer/pkg/srclang"
)

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

func (s *SecuritySelector) Select(cpg *graph.CPG, arch *extractor.ComponentArchitecture, findings []query.Finding) (*srclang.Layer, []srclang.Warning) {
	layer := &srclang.Layer{Name: "security"}
	var warnings []srclang.Warning

	layer.Findings = s.convertFindings(findings)
	s.addSecurityFunctions(cpg, layer, &warnings)
	if arch != nil {
		s.addRBAC(arch, layer)
		s.addNetworkPolicies(arch, layer)
	}
	s.addRelationships(cpg, layer)

	sortFindings(layer)
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
			fullPath := filepath.Join(s.repoPath, file)
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

func (s *SecuritySelector) addRBAC(arch *extractor.ComponentArchitecture, layer *srclang.Layer) {
	if arch.RBAC == nil {
		return
	}
	for _, cr := range arch.RBAC.ClusterRoles {
		layer.Resources = append(layer.Resources, srclang.Resource{
			Kind:       "ClusterRole",
			Name:       cr.Name,
			SourceFile: cr.Source,
			Summary:    fmt.Sprintf("%d rules", len(cr.Rules)),
		})
	}
	for _, r := range arch.RBAC.Roles {
		layer.Resources = append(layer.Resources, srclang.Resource{
			Kind:       "Role",
			Name:       r.Name,
			SourceFile: r.Source,
			Summary:    fmt.Sprintf("%d rules", len(r.Rules)),
		})
	}
}

func (s *SecuritySelector) addNetworkPolicies(arch *extractor.ComponentArchitecture, layer *srclang.Layer) {
	for _, np := range arch.NetworkPolicies {
		origin := "manifest"
		if strings.HasSuffix(np.Source, ".go") {
			origin = "programmatic"
		}
		layer.Resources = append(layer.Resources, srclang.Resource{
			Kind:       "NetworkPolicy",
			Name:       np.Name,
			SourceFile: np.Source,
			Origin:     origin,
		})
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
		if !selectedFuncs[fromKey] && !selectedFuncs[toKey] {
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

		if !selectedFuncs[fromKey] || !selectedFuncs[toKey] {
			resolved := false
			if !selectedFuncs[fromKey] {
				rel.From.Resolved = &resolved
			}
			if !selectedFuncs[toKey] {
				rel.To.Resolved = &resolved
			}
		}

		layer.Relationships = append(layer.Relationships, rel)
	}
}

func (s *SecuritySelector) convertFindings(findings []query.Finding) []srclang.Finding {
	var result []srclang.Finding
	counter := 0
	for _, f := range findings {
		if f.Domain != "security" && f.Domain != "netpolicy" {
			continue
		}
		counter++
		result = append(result, srclang.Finding{
			ID:          fmt.Sprintf("%s-%03d", f.RuleID, counter),
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

func sortFindings(layer *srclang.Layer) {
	severityOrder := map[string]int{
		"critical": 0, "high": 1, "medium": 2, "low": 3, "info": 4,
	}
	sort.Slice(layer.Findings, func(i, j int) bool {
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
}
