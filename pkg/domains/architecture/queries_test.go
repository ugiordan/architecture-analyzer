package architecture

import (
	"testing"

	"github.com/ugiordan/architecture-analyzer/pkg/graph"
)

func TestQueryAbstractionLayers(t *testing.T) {
	cpg := graph.NewCPG()

	base := &graph.Node{
		ID: "cls_base", Kind: graph.NodeClass, Name: "BaseStore",
		File: "store.py", Line: 1, Language: "python",
		BaseClasses: []string{"ABC"},
	}
	impl1 := &graph.Node{
		ID: "cls_impl1", Kind: graph.NodeClass, Name: "RedisStore",
		File: "redis.py", Line: 1, Language: "python",
		BaseClasses: []string{"BaseStore"},
	}
	impl2 := &graph.Node{
		ID: "cls_impl2", Kind: graph.NodeClass, Name: "SQLStore",
		File: "sql.py", Line: 1, Language: "python",
		BaseClasses: []string{"BaseStore"},
	}

	cpg.AddNode(base)
	cpg.AddNode(impl1)
	cpg.AddNode(impl2)

	ann := &PythonAnnotator{}
	ann.annotateClasses(cpg)

	findings := queryAbstractionLayers(cpg)
	if len(findings) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(findings))
	}
	if findings[0].RuleID != "CGA-A01" {
		t.Errorf("rule ID = %s, want CGA-A01", findings[0].RuleID)
	}
}

func TestQueryExternalAPISurface(t *testing.T) {
	cpg := graph.NewCPG()

	fn := &graph.Node{
		ID: "fn_1", Kind: graph.NodeFunction, Name: "get_embeddings",
		File: "embed.py", Line: 5, Language: "python",
	}
	cs := &graph.Node{
		ID: "call_1", Kind: graph.NodeCallSite, Name: "OpenAI",
		File: "embed.py", Line: 6, Language: "python",
	}

	cpg.AddNode(fn)
	cpg.AddNode(cs)
	cpg.AddEdge(&graph.Edge{From: fn.ID, To: cs.ID, Kind: graph.EdgeDataFlow})
	cpg.AddEdge(&graph.Edge{From: fn.ID, To: cs.ID, Kind: graph.EdgeCalls})

	ann := &PythonAnnotator{}
	ann.annotateSDKClients(cpg)

	if !fn.Annotations[AnnotSDKClient] {
		t.Error("expected fn to have arch:sdk_client annotation")
	}

	findings := queryExternalAPISurface(cpg)
	if len(findings) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(findings))
	}
	if findings[0].RuleID != "CGA-A02" {
		t.Errorf("rule ID = %s, want CGA-A02", findings[0].RuleID)
	}
}

func TestQueryFactoryDispatch(t *testing.T) {
	cpg := graph.NewCPG()

	fn := &graph.Node{
		ID: "fn_factory", Kind: graph.NodeFunction, Name: "get_vector_store",
		File: "factory.py", Line: 10, Language: "python",
	}
	sl1 := &graph.Node{
		ID: "struct_1", Kind: graph.NodeStructLiteral, Name: "ChromaStore",
		File: "factory.py", Line: 12, Language: "python",
	}
	sl2 := &graph.Node{
		ID: "struct_2", Kind: graph.NodeStructLiteral, Name: "LSStore",
		File: "factory.py", Line: 14, Language: "python",
	}

	cpg.AddNode(fn)
	cpg.AddNode(sl1)
	cpg.AddNode(sl2)
	cpg.AddEdge(&graph.Edge{From: fn.ID, To: sl1.ID, Kind: graph.EdgeDataFlow})
	cpg.AddEdge(&graph.Edge{From: fn.ID, To: sl2.ID, Kind: graph.EdgeDataFlow})

	ann := &PythonAnnotator{}
	ann.annotateFactoryMethods(cpg)

	if !fn.Annotations[AnnotFactoryMethod] {
		t.Error("expected fn to have arch:factory_method annotation")
	}

	findings := queryFactoryDispatch(cpg)
	if len(findings) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(findings))
	}
	if findings[0].RuleID != "CGA-A03" {
		t.Errorf("rule ID = %s, want CGA-A03", findings[0].RuleID)
	}
}

func TestQueryUnimplementedInterface(t *testing.T) {
	cpg := graph.NewCPG()

	base := &graph.Node{
		ID: "cls_orphan", Kind: graph.NodeClass, Name: "AbstractProcessor",
		File: "proc.py", Line: 1, Language: "python",
		BaseClasses: []string{"ABC"},
	}
	cpg.AddNode(base)

	ann := &PythonAnnotator{}
	ann.annotateClasses(cpg)

	findings := queryUnimplementedInterface(cpg)
	if len(findings) != 1 {
		t.Fatalf("expected 1 finding for orphan abstract base, got %d", len(findings))
	}
	if findings[0].RuleID != "CGA-A04" {
		t.Errorf("rule ID = %s, want CGA-A04", findings[0].RuleID)
	}

	// Add an implementation, should clear the finding
	impl := &graph.Node{
		ID: "cls_concrete", Kind: graph.NodeClass, Name: "ConcreteProcessor",
		File: "concrete.py", Line: 1, Language: "python",
		BaseClasses: []string{"AbstractProcessor"},
	}
	cpg.AddNode(impl)

	findings = queryUnimplementedInterface(cpg)
	if len(findings) != 0 {
		t.Errorf("expected 0 findings after adding implementation, got %d", len(findings))
	}
}
