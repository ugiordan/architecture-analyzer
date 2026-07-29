package compile

import (
	"fmt"
	"path/filepath"
	"strings"
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
	case "architecture":
		sel := layers.NewArchitectureSelector(opts.RepoPath)
		layer, warnings = sel.Select(opts.CPG, opts.Arch, opts.Findings, nil)
	case "testing":
		sel := layers.NewTestingSelector(opts.RepoPath)
		layer, warnings = sel.Select(opts.CPG, opts.Arch, opts.Findings, nil)
	case "upgrade":
		sel := layers.NewUpgradeSelector(opts.RepoPath)
		layer, warnings = sel.Select(opts.CPG, opts.Arch, opts.Findings, nil)
	case "netpolicy":
		sel := layers.NewNetpolicySelector(opts.RepoPath)
		layer, warnings = sel.Select(opts.CPG, opts.Arch, opts.Findings, nil)
	default:
		return nil, fmt.Errorf("unsupported layer %q (supports: security, architecture, testing, upgrade, netpolicy)", opts.Layer)
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

const BundleThreshold = 500_000

// SplitBundle splits a large Document into an index document + shard documents.
// Returns nil if the document is small enough for single-file output.
func SplitBundle(doc *srclang.Document) *srclang.Bundle {
	layer := &doc.Body.Layer

	var shards []srclang.Shard
	shardDocs := make(map[string]*srclang.Document)

	// Shard 1: findings
	if len(layer.Findings) > 0 {
		findingsPath := "findings.srclg"
		shards = append(shards, srclang.Shard{
			Type:  "findings",
			Path:  findingsPath,
			Count: len(layer.Findings),
		})
		shardDocs[findingsPath] = &srclang.Document{
			Version: doc.Version,
			Head: srclang.Head{
				Component:   doc.Head.Component,
				Layer:       doc.Head.Layer,
				ParentIndex: "index.srclg",
			},
			Body: srclang.Body{
				Layer: srclang.Layer{
					Name:     layer.Name,
					Findings: layer.Findings,
				},
			},
		}
	}

	// Shard per source file (functions with code)
	for _, file := range layer.Files {
		hasCode := false
		for _, fn := range file.Functions {
			if fn.Code != "" {
				hasCode = true
				break
			}
		}
		if !hasCode {
			continue
		}

		safeName := sanitizeFilePath(file.Path)
		shardPath := "files/" + safeName + ".srclg"

		// Collect relationships touching this file's functions
		fileFuncs := make(map[string]bool)
		for _, fn := range file.Functions {
			fileFuncs[fn.Name] = true
		}
		var fileRels []srclang.Relationship
		for _, rel := range layer.Relationships {
			if (rel.From.File == file.Path && fileFuncs[rel.From.Function]) ||
				(rel.To.File == file.Path && fileFuncs[rel.To.Function]) {
				fileRels = append(fileRels, rel)
			}
		}

		shards = append(shards, srclang.Shard{
			Type:  "functions",
			Path:  shardPath,
			File:  file.Path,
			Count: len(file.Functions),
		})
		shardDocs[shardPath] = &srclang.Document{
			Version: doc.Version,
			Head: srclang.Head{
				Component:   doc.Head.Component,
				Layer:       doc.Head.Layer,
				ParentIndex: "index.srclg",
			},
			Body: srclang.Body{
				Layer: srclang.Layer{
					Name:          layer.Name,
					Files:         []srclang.File{file},
					Relationships: fileRels,
				},
			},
		}
	}

	// Build index document (compact: no code, no finding descriptions)
	indexDoc := &srclang.Document{
		Version: doc.Version,
		Head:    doc.Head,
		Body:    doc.Body,
	}
	indexDoc.Head.Index = &srclang.Index{Shards: shards}

	return &srclang.Bundle{
		IndexDoc: indexDoc,
		Shards:   shardDocs,
	}
}

func sanitizeFilePath(path string) string {
	s := strings.ReplaceAll(path, "/", "_")
	s = strings.ReplaceAll(s, "\\", "_")
	s = strings.ReplaceAll(s, "..", "_")
	return s
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
