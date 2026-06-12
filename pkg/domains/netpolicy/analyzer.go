package netpolicy

import (
	"github.com/ugiordan/architecture-analyzer/pkg/domains"
	"github.com/ugiordan/architecture-analyzer/pkg/graph"
	"github.com/ugiordan/architecture-analyzer/pkg/query"
)

type Analyzer struct{}

func New() *Analyzer {
	return &Analyzer{}
}

func (a *Analyzer) Name() string                { return "netpolicy" }
func (a *Analyzer) SupportedLanguages() []string { return []string{"go", "python"} }
func (a *Analyzer) Dependencies() []string       { return nil }

func (a *Analyzer) Annotate(g *graph.CPG, _ string, _ *domains.ArchitectureData) error {
	return nil
}

func (a *Analyzer) Queries() []query.Rule {
	return netpolicyQueries()
}
