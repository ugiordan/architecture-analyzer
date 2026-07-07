package compile

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/ugiordan/architecture-analyzer/pkg/extractor"
	"github.com/ugiordan/architecture-analyzer/pkg/graph"
	"github.com/ugiordan/architecture-analyzer/pkg/query"
	"github.com/ugiordan/architecture-analyzer/pkg/srclang"
	"github.com/ugiordan/architecture-analyzer/pkg/srclang/layers"
)

type Options struct {
	RepoPath            string
	Layer               string
	CPG                 *graph.CPG
	Arch                *extractor.ComponentArchitecture
	Findings            []query.Finding
	SecurityAnnotations []extractor.SecurityAnnotation
	PlatformFile        string
}

func Compile(opts Options) (*srclang.Document, error) {
	if opts.CPG == nil {
		return nil, fmt.Errorf("CPG is required for layer %q", opts.Layer)
	}

	var layer *srclang.Layer
	var warnings []srclang.Warning

	switch opts.Layer {
	case "security":
		sel := layers.NewSecuritySelector(opts.RepoPath)
		layer, warnings = sel.Select(opts.CPG, opts.Arch, opts.Findings, opts.SecurityAnnotations)
	default:
		return nil, fmt.Errorf("unsupported layer %q (v0.0.1 supports: security)", opts.Layer)
	}

	doc := &srclang.Document{
		Version: "0.0.1",
		Head: srclang.Head{
			Layer:     layer.Name,
			Extracted: time.Now().UTC().Format(time.RFC3339),
		},
		Body: srclang.Body{Layer: *layer},
	}

	if opts.Arch != nil {
		doc.Head.Component = opts.Arch.Component
		doc.Head.Producer = "arch-analyzer " + opts.Arch.AnalyzerVersion
		if opts.Arch.Repo != "" {
			doc.Head.Repository = &srclang.Repository{
				URI:    opts.Arch.Repo,
				Commit: opts.Arch.CommitSHA,
			}
		}
		doc.Head.Languages = detectLanguages(opts.Arch)
	} else if opts.RepoPath != "" {
		doc.Head.Component = filepath.Base(opts.RepoPath)
	}

	if opts.PlatformFile != "" {
		p, err := extractPlatform(opts.PlatformFile, doc.Head.Component)
		if err != nil {
			warnings = append(warnings, srclang.Warning{
				Message: fmt.Sprintf("platform extraction failed: %v", err),
			})
		} else {
			doc.Head.Platform = p
		}
	}

	if len(warnings) > 0 {
		doc.Head.Diagnostics = warnings
	}

	return doc, nil
}

func detectLanguages(arch *extractor.ComponentArchitecture) []srclang.Language {
	var langs []srclang.Language
	if arch.GoASTMode != "" {
		langs = append(langs, srclang.Language{Name: "go"})
	}
	hasPython := len(arch.PythonK8sCalls) > 0
	if !hasPython && arch.Dependencies != nil {
		hasPython = len(arch.Dependencies.PythonPackages) > 0
	}
	if hasPython {
		langs = append(langs, srclang.Language{Name: "python"})
	}
	return langs
}
