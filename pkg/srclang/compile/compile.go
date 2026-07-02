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
	RepoPath string
	Layer    string
	CPG      *graph.CPG
	Arch     *extractor.ComponentArchitecture
	Findings []query.Finding
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
		layer, warnings = sel.Select(opts.CPG, opts.Arch, opts.Findings)
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
	if len(arch.PythonK8sCalls) > 0 || len(arch.ExternalConnections) > 0 {
		langs = append(langs, srclang.Language{Name: "python"})
	}
	return langs
}
