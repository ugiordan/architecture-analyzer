package architecture

import (
	"github.com/ugiordan/architecture-analyzer/pkg/domains"
	"github.com/ugiordan/architecture-analyzer/pkg/graph"
	"github.com/ugiordan/architecture-analyzer/pkg/query"
)

type Analyzer struct {
	annotators map[string]domains.Annotator
}

func New() *Analyzer {
	return &Analyzer{
		annotators: map[string]domains.Annotator{
			"python": &PythonAnnotator{},
		},
	}
}

func (a *Analyzer) Name() string                { return "architecture" }
func (a *Analyzer) SupportedLanguages() []string { return []string{"python"} }
func (a *Analyzer) Dependencies() []string       { return nil }

func (a *Analyzer) Annotate(g *graph.CPG, lang string, archData *domains.ArchitectureData) error {
	ann, ok := a.annotators[lang]
	if !ok {
		return nil
	}
	return ann.Annotate(g, archData)
}

func (a *Analyzer) Queries() []query.Rule {
	return architectureQueries()
}
