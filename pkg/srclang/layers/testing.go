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
	annotIsTestFunc     = "test:is_test_func"
	annotIsTestHelper   = "test:is_test_helper"
	annotUsesFakeClient = "test:uses_fake_client"
	annotUsesEnvtest    = "test:uses_envtest"
	annotTableDriven    = "test:table_driven"
	annotErrorPath      = "test:error_path"
	annotSubtests       = "test:subtests"
)

type TestingSelector struct {
	repoPath string
	body     *srclang.BodyExtractor
}

func NewTestingSelector(repoPath string) *TestingSelector {
	return &TestingSelector{
		repoPath: repoPath,
		body:     srclang.NewBodyExtractor(),
	}
}


func (s *TestingSelector) Select(cpg *graph.CPG, _ *extractor.ComponentArchitecture, findings []query.Finding, _ []extractor.SecurityAnnotation) (*srclang.Layer, []srclang.Warning) {
	layer := &srclang.Layer{Name: "testing"}
	var warnings []srclang.Warning

	layer.Findings = s.convertFindings(findings)

	testTargets := s.buildTestTargetMap(cpg)
	s.addTestFunctions(cpg, testTargets, layer, &warnings)
	s.addTestedFunctions(cpg, testTargets, layer, &warnings)
	s.addTestRelationships(cpg, testTargets, layer)

	sortFindings(layer)
	capFunctions(layer)
	capCodeBodies(layer)
	return layer, warnings
}

type testTarget struct {
	testFuncID   string
	targetFuncID string
}

// buildTestTargetMap finds which test functions call which implementation functions.
// Uses the same edge traversal pattern as the testing domain queries:
// EdgeCalls: CallSite -> Function, EdgeDataFlow: Function -> CallSite.
func (s *TestingSelector) buildTestTargetMap(cpg *graph.CPG) []testTarget {
	testFnIDs := make(map[string]bool)
	for _, fn := range cpg.NodesByKind(graph.NodeFunction) {
		if fn.Annotations[annotIsTestFunc] {
			testFnIDs[fn.ID] = true
		}
	}

	var targets []testTarget
	seen := make(map[string]bool)

	for _, fn := range cpg.NodesByKind(graph.NodeFunction) {
		if testFnIDs[fn.ID] {
			continue
		}
		for _, inEdge := range cpg.InEdges(fn.ID) {
			if inEdge.Kind != graph.EdgeCalls {
				continue
			}
			for _, parentEdge := range cpg.InEdges(inEdge.From) {
				if parentEdge.Kind == graph.EdgeDataFlow && testFnIDs[parentEdge.From] {
					key := parentEdge.From + ":" + fn.ID
					if !seen[key] {
						seen[key] = true
						targets = append(targets, testTarget{
							testFuncID:   parentEdge.From,
							targetFuncID: fn.ID,
						})
					}
				}
			}
		}
	}
	return targets
}

func (s *TestingSelector) addTestFunctions(cpg *graph.CPG, targets []testTarget, layer *srclang.Layer, warnings *[]srclang.Warning) {
	fileMap := make(map[string]*srclang.File)

	testFuncHasTarget := make(map[string]bool)
	for _, t := range targets {
		testFuncHasTarget[t.testFuncID] = true
	}

	for _, fn := range cpg.NodesByKind(graph.NodeFunction) {
		if !fn.Annotations[annotIsTestFunc] && !fn.Annotations[annotIsTestHelper] {
			continue
		}
		if fn.File == "" {
			continue
		}

		sf, ok := fileMap[fn.File]
		if !ok {
			sf = &srclang.File{
				Path:     fn.File,
				Language: fn.Language,
			}
			fileMap[fn.File] = sf
		}

		slFn := s.buildFunction(fn, warnings)
		slFn.Metas = s.testMetas(fn, testFuncHasTarget)
		sf.Functions = append(sf.Functions, slFn)
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

func (s *TestingSelector) addTestedFunctions(cpg *graph.CPG, targets []testTarget, layer *srclang.Layer, warnings *[]srclang.Warning) {
	existingFuncs := make(map[string]bool)
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			existingFuncs[f.Path+":"+fmt.Sprintf("%d", fn.SourceLine)] = true
		}
	}

	targetIDs := make(map[string]bool)
	for _, t := range targets {
		targetIDs[t.targetFuncID] = true
	}

	fileMap := make(map[string]*srclang.File)
	for i := range layer.Files {
		fileMap[layer.Files[i].Path] = &layer.Files[i]
	}

	for _, fn := range cpg.NodesByKind(graph.NodeFunction) {
		if !targetIDs[fn.ID] {
			continue
		}
		if fn.File == "" || fn.Line == 0 {
			continue
		}
		locKey := fn.File + ":" + fmt.Sprintf("%d", fn.Line)
		if existingFuncs[locKey] {
			continue
		}

		sf, ok := fileMap[fn.File]
		if !ok {
			newFile := &srclang.File{
				Path:     fn.File,
				Language: fn.Language,
			}
			fileMap[fn.File] = newFile
			sf = newFile
		}

		slFn := s.buildFunction(fn, warnings)
		slFn.Metas = append(slFn.Metas, srclang.Meta{Domain: "testing", Key: "tested", Value: "true"})
		sf.Functions = append(sf.Functions, slFn)
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

func (s *TestingSelector) addTestRelationships(cpg *graph.CPG, targets []testTarget, layer *srclang.Layer) {
	nodeEndpoint := func(id string) (srclang.Endpoint, bool) {
		node := cpg.GetNode(id)
		if node == nil || node.File == "" {
			return srclang.Endpoint{}, false
		}
		return srclang.Endpoint{
			Function: node.Name,
			File:     node.File,
			Line:     node.Line,
		}, true
	}

	for _, t := range targets {
		from, ok1 := nodeEndpoint(t.testFuncID)
		to, ok2 := nodeEndpoint(t.targetFuncID)
		if !ok1 || !ok2 {
			continue
		}

		layer.Relationships = append(layer.Relationships, srclang.Relationship{
			Kind: "tests",
			From: from,
			To:   to,
		})
	}

	if len(layer.Relationships) > maxRelationships {
		layer.Relationships = layer.Relationships[:maxRelationships]
	}
}

func (s *TestingSelector) buildFunction(node *graph.Node, warnings *[]srclang.Warning) srclang.Function {
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

func (s *TestingSelector) testMetas(node *graph.Node, testFuncHasTarget map[string]bool) []srclang.Meta {
	var metas []srclang.Meta

	if node.Annotations[annotIsTestFunc] {
		kind := "unit"
		if node.Annotations[annotUsesEnvtest] {
			kind = "integration"
		} else if strings.HasPrefix(node.Name, "Benchmark") {
			kind = "benchmark"
		}
		metas = append(metas, srclang.Meta{Domain: "testing", Key: "test-kind", Value: kind})
	} else if node.Annotations[annotIsTestHelper] {
		metas = append(metas, srclang.Meta{Domain: "testing", Key: "test-kind", Value: "helper"})
	}

	if node.Annotations[annotTableDriven] {
		metas = append(metas, srclang.Meta{Domain: "testing", Key: "table-driven", Value: "true"})
	}

	if node.Annotations[annotSubtests] {
		metas = append(metas, srclang.Meta{Domain: "testing", Key: "subtests", Value: "true"})
	}

	if node.Annotations[annotUsesFakeClient] {
		metas = append(metas, srclang.Meta{Domain: "testing", Key: "mock-target", Value: "fake-client"})
	}

	if node.Annotations[annotUsesEnvtest] {
		metas = append(metas, srclang.Meta{Domain: "testing", Key: "mock-target", Value: "envtest"})
	}

	if node.Annotations[annotErrorPath] {
		metas = append(metas, srclang.Meta{Domain: "testing", Key: "error-path", Value: "true"})
	}

	if testFuncHasTarget[node.ID] {
		metas = append(metas, srclang.Meta{Domain: "testing", Key: "has-target", Value: "true"})
	}

	return metas
}

func (s *TestingSelector) convertFindings(findings []query.Finding) []srclang.Finding {
	var result []srclang.Finding
	counters := make(map[string]int)
	for _, f := range findings {
		if f.Domain != "testing" {
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
