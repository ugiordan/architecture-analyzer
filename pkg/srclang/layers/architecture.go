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

type ArchitectureSelector struct {
	repoPath string
	body     *srclang.BodyExtractor
}

func NewArchitectureSelector(repoPath string) *ArchitectureSelector {
	return &ArchitectureSelector{
		repoPath: repoPath,
		body:     srclang.NewBodyExtractor(),
	}
}

func (s *ArchitectureSelector) Select(cpg *graph.CPG, arch *extractor.ComponentArchitecture, findings []query.Finding, _ []extractor.SecurityAnnotation) (*srclang.Layer, []srclang.Warning) {
	layer := &srclang.Layer{Name: "architecture"}
	var warnings []srclang.Warning

	if arch != nil {
		s.addCRDs(arch, layer)
		s.addControllerWatches(arch, layer)
		s.addServices(arch, layer)
		s.addDeployments(arch, layer)
		s.addExternalConnections(arch, layer)
		s.addWebhooks(arch, layer)
		s.addReconcileSequences(arch, layer)
	}

	s.addArchitectureFindings(findings, layer)
	s.addReconcileFunctions(cpg, layer, &warnings)

	return layer, warnings
}

func (s *ArchitectureSelector) addCRDs(arch *extractor.ComponentArchitecture, layer *srclang.Layer) {
	for _, crd := range arch.CRDs {
		r := srclang.Resource{
			Kind:       "CustomResourceDefinition",
			Name:       crd.Kind,
			SourceFile: crd.Source,
			APIGroup:   crd.Group,
			APIVersion: crd.Version,
			Scope:      crd.Scope,
			FieldCount: crd.FieldsCount,
		}
		if crd.GoSource != "" {
			r.Summary = fmt.Sprintf("Go AST enriched, hub=%s", crd.HubVersion)
		}
		layer.Resources = append(layer.Resources, r)
	}
}

func (s *ArchitectureSelector) addControllerWatches(arch *extractor.ComponentArchitecture, layer *srclang.Layer) {
	controllers := make(map[string][]extractor.ControllerWatch)
	for _, cw := range arch.ControllerWatch {
		controllers[cw.Controller] = append(controllers[cw.Controller], cw)
	}

	for controller, watches := range controllers {
		for _, w := range watches {
			layer.Relationships = append(layer.Relationships, srclang.Relationship{
				Kind: w.Type,
				From: srclang.Endpoint{
					Function: controller,
					File:     w.Source,
				},
				To: srclang.Endpoint{
					Resource: w.GVK,
				},
			})
		}
	}
}

func (s *ArchitectureSelector) addServices(arch *extractor.ComponentArchitecture, layer *srclang.Layer) {
	for _, svc := range arch.Services {
		r := srclang.Resource{
			Kind:       "Service",
			Name:       svc.Name,
			SourceFile: svc.Source,
		}
		if svc.TargetDeployment != "" {
			r.Summary = fmt.Sprintf("targets %s", svc.TargetDeployment)
		}
		if len(svc.Ports) > 0 {
			var portStrs []string
			for _, p := range svc.Ports {
				portStrs = append(portStrs, fmt.Sprintf("%s/%d", p.Name, p.Port))
			}
			if r.Summary != "" {
				r.Summary += ", "
			}
			r.Summary += "ports: " + strings.Join(portStrs, ", ")
		}
		layer.Resources = append(layer.Resources, r)
	}
}

func (s *ArchitectureSelector) addDeployments(arch *extractor.ComponentArchitecture, layer *srclang.Layer) {
	for _, dep := range arch.Deployments {
		r := srclang.Resource{
			Kind:       dep.Kind,
			Name:       dep.Name,
			SourceFile: dep.Source,
		}
		if r.Kind == "" {
			r.Kind = "Deployment"
		}
		var containerNames []string
		for _, c := range dep.Containers {
			containerNames = append(containerNames, c.Name)
		}
		if len(containerNames) > 0 {
			r.Summary = "containers: " + strings.Join(containerNames, ", ")
		}
		if dep.ServiceAccount != "" {
			if r.Summary != "" {
				r.Summary += ", "
			}
			r.Summary += "sa=" + dep.ServiceAccount
		}
		layer.Resources = append(layer.Resources, r)
	}
}

func (s *ArchitectureSelector) addExternalConnections(arch *extractor.ComponentArchitecture, layer *srclang.Layer) {
	for _, ec := range arch.ExternalConnections {
		layer.Relationships = append(layer.Relationships, srclang.Relationship{
			Kind: "external",
			From: srclang.Endpoint{
				Function: ec.Function,
				File:     ec.Source,
			},
			To: srclang.Endpoint{
				Resource: ec.Service,
			},
		})
	}
}

func (s *ArchitectureSelector) addWebhooks(arch *extractor.ComponentArchitecture, layer *srclang.Layer) {
	for _, wh := range arch.Webhooks {
		r := srclang.Resource{
			Kind: "Webhook",
			Name: wh.Name,
		}
		if len(wh.Sources) > 0 {
			r.SourceFile = wh.Sources[0].File
		}
		r.Summary = wh.Type
		if wh.Path != "" {
			r.Summary += " path=" + wh.Path
		}
		layer.Resources = append(layer.Resources, r)
	}
}

func (s *ArchitectureSelector) addReconcileSequences(arch *extractor.ComponentArchitecture, layer *srclang.Layer) {
	for _, rs := range arch.ReconcileSequences {
		for _, step := range rs.Steps {
			layer.Relationships = append(layer.Relationships, srclang.Relationship{
				Kind: "reconcile-step",
				From: srclang.Endpoint{
					Function: rs.Controller,
					File:     rs.Source,
				},
				To: srclang.Endpoint{
					Function: step.Method,
					File:     step.Source,
				},
			})
		}
	}
}

func (s *ArchitectureSelector) addArchitectureFindings(findings []query.Finding, layer *srclang.Layer) {
	counters := make(map[string]int)
	for _, f := range findings {
		if f.Domain != "architecture" {
			continue
		}
		counters[f.RuleID]++
		layer.Findings = append(layer.Findings, srclang.Finding{
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
}

func (s *ArchitectureSelector) addReconcileFunctions(cpg *graph.CPG, layer *srclang.Layer, warnings *[]srclang.Warning) {
	fileMap := make(map[string]*srclang.File)

	for _, node := range cpg.Nodes() {
		if node.Kind != graph.NodeFunction {
			continue
		}
		if node.File == "" || node.Line == 0 {
			continue
		}
		isReconcile := strings.HasPrefix(node.Name, "Reconcile") ||
			strings.HasPrefix(node.Name, "reconcile") ||
			strings.Contains(node.Name, "Reconciler") ||
			node.Annotations["reconciler"]
		if !isReconcile {
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

		fn := srclang.Function{
			Name:       node.Name,
			Kind:       "method",
			SourceLine: node.Line,
		}
		if node.Complexity > 0 {
			fn.Complexity = node.Complexity
		}
		if node.EndLine > 0 {
			fullPath, ok := safeJoinPath(s.repoPath, node.File)
			if ok {
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

func safeJoinPath(repoPath, file string) (string, bool) {
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
