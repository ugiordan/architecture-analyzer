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

type CodegenSelector struct {
	repoPath string
	body     *srclang.BodyExtractor
}

func NewCodegenSelector(repoPath string) *CodegenSelector {
	return &CodegenSelector{
		repoPath: repoPath,
		body:     srclang.NewBodyExtractor(),
	}
}


func (s *CodegenSelector) Select(cpg *graph.CPG, arch *extractor.ComponentArchitecture, findings []query.Finding, annotations []extractor.SecurityAnnotation) (*srclang.Layer, []srclang.Warning) {
	layer := &srclang.Layer{Name: "codegen"}
	var warnings []srclang.Warning

	layer.Findings = s.convertFindings(findings)
	layer.Findings = append(layer.Findings, convertSecurityAnnotations(annotations)...)

	s.addAllFunctions(cpg, layer, &warnings)
	s.addHTTPEndpoints(cpg, layer)
	if arch != nil {
		s.addImports(arch, layer)
	}
	s.addRelationships(cpg, layer)

	sortFindings(layer)
	capFunctions(layer)
	capCodeBodies(layer)
	return layer, warnings
}

func (s *CodegenSelector) addAllFunctions(cpg *graph.CPG, layer *srclang.Layer, warnings *[]srclang.Warning) {
	fileMap := make(map[string]*srclang.File)

	for _, node := range cpg.NodesByKind(graph.NodeFunction) {
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

		fn := s.buildFunction(node, warnings)
		fn.Metas = s.codegenMetas(node)
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

func (s *CodegenSelector) addHTTPEndpoints(cpg *graph.CPG, layer *srclang.Layer) {
	existingFuncs := make(map[string]bool)
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			existingFuncs[f.Path+":"+fmt.Sprintf("%d", fn.SourceLine)] = true
		}
	}

	fileMap := make(map[string]*srclang.File)
	for i := range layer.Files {
		fileMap[layer.Files[i].Path] = &layer.Files[i]
	}

	for _, node := range cpg.NodesByKind(graph.NodeHTTPEndpoint) {
		if node.File == "" {
			continue
		}
		locKey := node.File + ":" + fmt.Sprintf("%d", node.Line)
		if existingFuncs[locKey] {
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

		sf.Functions = append(sf.Functions, srclang.Function{
			Name:       node.Name,
			Kind:       "handler",
			SourceLine: node.Line,
		})
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

func (s *CodegenSelector) addImports(arch *extractor.ComponentArchitecture, layer *srclang.Layer) {
	if arch.Dependencies == nil {
		return
	}

	for _, mod := range arch.Dependencies.GoModules {
		imp := srclang.Import{
			Module:  mod.Module,
			Kind:    "go-module",
			Version: mod.Version,
		}
		layer.Imports = append(layer.Imports, imp)
	}

	for _, pkg := range arch.Dependencies.PythonPackages {
		imp := srclang.Import{
			Module:  pkg.Name,
			Kind:    "python-package",
			Version: pkg.Version,
		}
		if pkg.Source != "" {
			imp.Path = pkg.Source
		}
		layer.Imports = append(layer.Imports, imp)
	}
}

func (s *CodegenSelector) addRelationships(cpg *graph.CPG, layer *srclang.Layer) {
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

func (s *CodegenSelector) buildFunction(node *graph.Node, warnings *[]srclang.Warning) srclang.Function {
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

	if node.TrustLevel != "" {
		fn.Trust = string(node.TrustLevel)
	}

	if node.Annotations["taint_source"] {
		fn.TaintRole = "source"
		if fn.Trust == "" {
			fn.Trust = "untrusted"
		}
	} else if node.Annotations["taint_sink"] {
		fn.TaintRole = "sink"
	}

	return fn
}

func (s *CodegenSelector) codegenMetas(node *graph.Node) []srclang.Meta {
	var metas []srclang.Meta

	role := classifyCodeRole(node)
	metas = append(metas, srclang.Meta{Domain: "codegen", Key: "code-role", Value: role})

	if isTestCode(node) {
		metas = append(metas, srclang.Meta{Domain: "codegen", Key: "test-only", Value: "true"})
	}

	if isGeneratedCode(node) {
		metas = append(metas, srclang.Meta{Domain: "codegen", Key: "generated", Value: "true"})
	}

	if node.Complexity > 10 {
		metas = append(metas, srclang.Meta{Domain: "codegen", Key: "change-risk", Value: "high"})
	} else if node.Complexity > 5 {
		metas = append(metas, srclang.Meta{Domain: "codegen", Key: "change-risk", Value: "medium"})
	}

	return metas
}

func isTestCode(node *graph.Node) bool {
	return node.IsTest ||
		strings.HasSuffix(node.File, "_test.go") ||
		strings.HasPrefix(filepath.Base(node.File), "test_")
}

func classifyCodeRole(node *graph.Node) string {
	if isTestCode(node) {
		return "test-only"
	}

	if isGeneratedCode(node) {
		return "generated"
	}

	name := node.Name
	if len(name) > 0 {
		first := name[0]
		if first >= 'A' && first <= 'Z' {
			return "public-api"
		}
	}

	return "internal"
}

func isGeneratedCode(node *graph.Node) bool {
	file := strings.ToLower(node.File)
	return strings.Contains(file, "zz_generated") ||
		strings.Contains(file, "generated.go") ||
		strings.Contains(file, "deepcopy") ||
		strings.Contains(file, "_gen.go") ||
		strings.Contains(file, "/generated/")
}

func (s *CodegenSelector) convertFindings(findings []query.Finding) []srclang.Finding {
	var result []srclang.Finding
	counters := make(map[string]int)
	for _, f := range findings {
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
