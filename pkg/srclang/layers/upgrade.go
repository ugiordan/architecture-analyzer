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

const (
	annotVersionConversion = "upgrade:version_conversion"
	annotFeatureGate       = "upgrade:feature_gate"
	annotPreReleaseAPI     = "upgrade:pre_release_api"
	annotMigration         = "upgrade:migration"
	annotVersionCheck      = "upgrade:version_check"
)

type UpgradeSelector struct {
	repoPath string
	body     *srclang.BodyExtractor
}

func NewUpgradeSelector(repoPath string) *UpgradeSelector {
	return &UpgradeSelector{
		repoPath: repoPath,
		body:     srclang.NewBodyExtractor(),
	}
}

func (s *UpgradeSelector) safeJoin(file string) (string, bool) {
	fullPath := filepath.Join(s.repoPath, file)
	absPath, err := filepath.Abs(fullPath)
	if err != nil {
		return "", false
	}
	absRepo, err := filepath.Abs(s.repoPath)
	if err != nil {
		return "", false
	}
	if !strings.HasPrefix(absPath, absRepo+string(filepath.Separator)) && absPath != absRepo {
		return "", false
	}
	return fullPath, true
}

func (s *UpgradeSelector) Select(cpg *graph.CPG, arch *extractor.ComponentArchitecture, findings []query.Finding, _ []extractor.SecurityAnnotation) (*srclang.Layer, []srclang.Warning) {
	layer := &srclang.Layer{Name: "upgrade"}
	var warnings []srclang.Warning

	layer.Findings = s.convertFindings(findings)
	s.addUpgradeFunctions(cpg, layer, &warnings)
	s.addFindingReferencedFunctions(cpg, layer, &warnings)
	if arch != nil {
		s.addFeatureGates(arch, layer)
	}
	s.addRelationships(cpg, layer)

	sortFindings(layer)
	capFunctions(layer)
	capCodeBodies(layer)
	return layer, warnings
}

func (s *UpgradeSelector) addUpgradeFunctions(cpg *graph.CPG, layer *srclang.Layer, warnings *[]srclang.Warning) {
	fileMap := make(map[string]*srclang.File)

	for _, node := range cpg.Nodes() {
		if !isUpgradeRelevant(node) {
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

		fn := s.buildUpgradeFunction(node, warnings)
		fn.Metas = s.upgradeMetas(node)
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

func (s *UpgradeSelector) addFindingReferencedFunctions(cpg *graph.CPG, layer *srclang.Layer, warnings *[]srclang.Warning) {
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

		fn := s.buildUpgradeFunction(node, warnings)
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

func (s *UpgradeSelector) addFeatureGates(arch *extractor.ComponentArchitecture, layer *srclang.Layer) {
	for _, fg := range arch.FeatureGates {
		defaultStr := "disabled"
		if fg.Default {
			defaultStr = "enabled"
		}
		summary := fmt.Sprintf("default=%s", defaultStr)
		if fg.PreRelease != "" {
			summary += fmt.Sprintf(", stage=%s", fg.PreRelease)
		}
		if fg.LockToDefault {
			summary += ", locked"
		}

		res := srclang.Resource{
			Kind:       "FeatureGate",
			Name:       fg.Name,
			SourceFile: fg.Source,
			Summary:    summary,
		}
		layer.Resources = append(layer.Resources, res)
	}
}

func (s *UpgradeSelector) addRelationships(cpg *graph.CPG, layer *srclang.Layer) {
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

func (s *UpgradeSelector) buildUpgradeFunction(node *graph.Node, warnings *[]srclang.Warning) srclang.Function {
	fn := srclang.Function{
		Name:       node.Name,
		Kind:       functionKind(node),
		SourceLine: node.Line,
	}

	if node.Complexity > 0 {
		fn.Complexity = node.Complexity
	}

	if node.EndLine > 0 && node.Line > 0 && node.EndLine >= node.Line {
		fullPath, ok := s.safeJoin(node.File)
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

func (s *UpgradeSelector) upgradeMetas(node *graph.Node) []srclang.Meta {
	var metas []srclang.Meta

	if node.Annotations[annotVersionConversion] {
		metas = append(metas, srclang.Meta{Domain: "upgrade", Key: "conversion", Value: "true"})
	}

	if node.Annotations[annotFeatureGate] {
		metas = append(metas, srclang.Meta{Domain: "upgrade", Key: "feature-gate", Value: "true"})
	}

	if node.Annotations[annotMigration] {
		metas = append(metas, srclang.Meta{Domain: "upgrade", Key: "migration", Value: "true"})
	}

	if node.Annotations[annotPreReleaseAPI] {
		metas = append(metas, srclang.Meta{Domain: "upgrade", Key: "pre-release-api", Value: "true"})
	}

	if node.Annotations[annotVersionCheck] {
		metas = append(metas, srclang.Meta{Domain: "upgrade", Key: "version-check", Value: "true"})
	}

	return metas
}

func isUpgradeRelevant(node *graph.Node) bool {
	if node.Kind != graph.NodeFunction {
		return false
	}
	for _, annot := range []string{
		annotVersionConversion,
		annotFeatureGate,
		annotPreReleaseAPI,
		annotMigration,
		annotVersionCheck,
	} {
		if node.Annotations[annot] {
			return true
		}
	}
	return false
}

func (s *UpgradeSelector) convertFindings(findings []query.Finding) []srclang.Finding {
	var result []srclang.Finding
	counters := make(map[string]int)
	for _, f := range findings {
		if f.Domain != "upgrade" {
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
