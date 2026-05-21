package architecture

import (
	"strings"

	"github.com/ugiordan/architecture-analyzer/pkg/domains"
	"github.com/ugiordan/architecture-analyzer/pkg/graph"
)

type PythonAnnotator struct{}

func (a *PythonAnnotator) Annotate(g *graph.CPG, _ *domains.ArchitectureData) error {
	a.annotateClasses(g)
	a.annotateFactoryMethods(g)
	a.annotateSDKClients(g)
	return nil
}

func (a *PythonAnnotator) annotateClasses(g *graph.CPG) {
	for _, cls := range g.NodesByKind(graph.NodeClass) {
		if cls.Language != "python" {
			continue
		}
		if len(cls.BaseClasses) > 0 {
			g.SetAnnotation(cls.ID, AnnotImplements, true)
		}
		if isAbstractBase(cls) {
			g.SetAnnotation(cls.ID, AnnotAbstractBase, true)
		}
	}
}

func isAbstractBase(cls *graph.Node) bool {
	for _, base := range cls.BaseClasses {
		if base == "ABC" || base == "ABCMeta" || strings.HasSuffix(base, ".ABC") {
			return true
		}
	}
	lower := strings.ToLower(cls.Name)
	return strings.HasPrefix(lower, "base") || strings.HasPrefix(lower, "abstract")
}

// annotateFactoryMethods marks functions that return different class
// instantiations based on conditionals. Detected via: function contains
// multiple StructLiteral call sites (class instantiations) connected by
// data flow, suggesting a dispatch pattern.
func (a *PythonAnnotator) annotateFactoryMethods(g *graph.CPG) {
	for _, fn := range g.NodesByKind(graph.NodeFunction) {
		if fn.Language != "python" {
			continue
		}
		instantiations := 0
		var types []string
		for _, edge := range g.OutEdges(fn.ID) {
			if edge.Kind != graph.EdgeDataFlow {
				continue
			}
			target := g.GetNode(edge.To)
			if target == nil {
				continue
			}
			if target.Kind == graph.NodeStructLiteral {
				instantiations++
				types = append(types, target.Name)
			}
			if target.Kind == graph.NodeCallSite {
				for _, inner := range g.OutEdges(target.ID) {
					it := g.GetNode(inner.To)
					if it != nil && it.Kind == graph.NodeStructLiteral {
						instantiations++
						types = append(types, it.Name)
					}
				}
			}
		}
		if instantiations >= 2 {
			g.SetAnnotation(fn.ID, AnnotFactoryMethod, true)
			if fn.Properties == nil {
				fn.Properties = make(map[string]string)
			}
			fn.Properties["factory_types"] = strings.Join(types, ",")
		}
	}
}

var sdkClientPatterns = []string{
	"openai", "anthropic", "llamastackclient",
	"chromadb", "elasticsearch", "weaviate", "pinecone",
	"qdrantclient", "milvusclient",
	"boto3", "requests.session",
	"httpx.client", "httpx.asyncclient",
	"aiohttp.clientsession",
	"mlflow",
}

func (a *PythonAnnotator) annotateSDKClients(g *graph.CPG) {
	for _, cs := range g.NodesByKind(graph.NodeCallSite) {
		if cs.Language != "python" {
			continue
		}
		lower := strings.ToLower(cs.Name)
		for _, pattern := range sdkClientPatterns {
			if strings.Contains(lower, pattern) {
				g.SetAnnotation(cs.ID, AnnotSDKClient, true)
				for _, edge := range g.InEdges(cs.ID) {
					src := g.GetNode(edge.From)
					if src != nil && src.Kind == graph.NodeFunction {
						g.SetAnnotation(src.ID, AnnotSDKClient, true)
					}
				}
				break
			}
		}
	}
}
