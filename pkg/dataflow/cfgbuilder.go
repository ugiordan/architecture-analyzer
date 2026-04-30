package dataflow

import (
	"fmt"

	"github.com/ugiordan/architecture-analyzer/pkg/graph"
)

// MaxBlocksPerFunction is the limit on NodeBasicBlock nodes per function.
// When exceeded, CFG construction stops for that function.
const MaxBlocksPerFunction = 200

// MaxASTDepth is the maximum recursion depth for AST walking functions
// (collectNodeIDs, containsPanic). Prevents stack overflow on deeply nested input.
const MaxASTDepth = 1000

// Valid EdgeControlFlow labels used by CFG construction:
//   - "entry"        function -> entry block
//   - "exit"         block -> exit block (return, panic, raise, throw, or function end)
//   - "true_branch"  condition true, or switch/match case
//   - "false_branch" condition false, or switch/match default/no-else
//   - "fallthrough"  sequential flow, or branch merging
//   - "loop_back"    loop body end -> loop header
//   - "loop_exit"    loop header -> after-loop block

// CFGBuilder accumulates basic blocks and control flow edges during
// CFG construction. One instance per function body analysis.
type CFGBuilder struct {
	blocks []*graph.Node
	edges  []*graph.Edge
	funcID string
	file   string
}

// NewCFGBuilder creates a CFGBuilder for the given function.
func NewCFGBuilder(funcID, file string) *CFGBuilder {
	return &CFGBuilder{
		funcID: funcID,
		file:   file,
	}
}

// NewBlock creates a NodeBasicBlock node. The block ID includes the funcID
// and a sequence number to ensure uniqueness across functions and blocks.
func (cb *CFGBuilder) NewBlock(name string, line int) *graph.Node {
	seq := len(cb.blocks)
	blockName := fmt.Sprintf("%s-%d", name, seq)
	block := &graph.Node{
		ID:       graph.NodeID(graph.NodeBasicBlock, cb.funcID+"/"+blockName, cb.file, line, 0),
		Kind:     graph.NodeBasicBlock,
		Name:     blockName,
		File:     cb.file,
		Line:     line,
		ParentID: cb.funcID,
	}
	cb.blocks = append(cb.blocks, block)
	return block
}

// AddMember appends a node ID to a block's Members list.
func (cb *CFGBuilder) AddMember(block *graph.Node, memberID string) {
	block.Members = append(block.Members, memberID)
}

// AddEdge creates an EdgeControlFlow edge between two blocks.
func (cb *CFGBuilder) AddEdge(fromBlockID, toBlockID, label string) {
	cb.edges = append(cb.edges, &graph.Edge{
		From:       fromBlockID,
		To:         toBlockID,
		Kind:       graph.EdgeControlFlow,
		Label:      label,
		Confidence: graph.ConfidenceCertain,
	})
}

// BlockCount returns the number of blocks created so far.
func (cb *CFGBuilder) BlockCount() int {
	return len(cb.blocks)
}

// Result returns all accumulated blocks and edges.
func (cb *CFGBuilder) Result() ([]*graph.Node, []*graph.Edge) {
	return cb.blocks, cb.edges
}
