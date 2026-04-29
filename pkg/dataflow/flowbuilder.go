package dataflow

import "github.com/ugiordan/architecture-analyzer/pkg/graph"

// MaxVariablesPerFunction is the limit on NodeVariable nodes per function.
// When exceeded, analyzeFunctionBody() stops creating new variables.
const MaxVariablesPerFunction = 500

// FlowBuilder accumulates data flow nodes and edges during function body analysis.
type FlowBuilder struct {
	nodes    []*graph.Node
	edges    []*graph.Edge
	varCount int
}

func NewFlowBuilder() *FlowBuilder {
	return &FlowBuilder{}
}

func (fb *FlowBuilder) AddVariable(node *graph.Node) {
	fb.nodes = append(fb.nodes, node)
	fb.varCount++
}

func (fb *FlowBuilder) AddParameter(node *graph.Node) {
	fb.nodes = append(fb.nodes, node)
}

func (fb *FlowBuilder) VariableCount() int {
	return fb.varCount
}

func (fb *FlowBuilder) addEdge(from, to, label string, confidence graph.EdgeConfidence) {
	fb.edges = append(fb.edges, &graph.Edge{
		From:       from,
		To:         to,
		Kind:       graph.EdgeDataFlow,
		Label:      label,
		Confidence: confidence,
	})
}

func (fb *FlowBuilder) AddDeclares(funcID, paramID string) {
	fb.addEdge(funcID, paramID, "declares", graph.ConfidenceCertain)
}

func (fb *FlowBuilder) AddAssign(fromID, toVarID string) {
	fb.addEdge(fromID, toVarID, "assigns", graph.ConfidenceCertain)
}

func (fb *FlowBuilder) AddRead(varID, toExprID string) {
	fb.addEdge(varID, toExprID, "reads", graph.ConfidenceCertain)
}

func (fb *FlowBuilder) AddPassesTo(varID, callSiteID string) {
	fb.addEdge(varID, callSiteID, "passes_to", graph.ConfidenceCertain)
}

func (fb *FlowBuilder) AddMutates(callSiteID, varID string) {
	fb.addEdge(callSiteID, varID, "mutates", graph.ConfidenceInferred)
}

func (fb *FlowBuilder) AddFieldAccess(objVarID, fieldVarID string) {
	fb.addEdge(objVarID, fieldVarID, "field_access", graph.ConfidenceCertain)
}

func (fb *FlowBuilder) AddReturn(varID, funcID string) {
	fb.addEdge(varID, funcID, "returns", graph.ConfidenceCertain)
}

func (fb *FlowBuilder) Result() ([]*graph.Node, []*graph.Edge) {
	return fb.nodes, fb.edges
}
