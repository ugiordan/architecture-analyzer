package dataflow

// SymbolTable is a flat, scope-aware variable tracker for intraprocedural analysis.
// One instance per function body. Variables are tracked by name with last-write-wins
// semantics (no block scoping).
type SymbolTable struct {
	vars map[string]string // variable name -> most recent NodeVariable/NodeParameter ID
}

// NewSymbolTable creates an empty symbol table.
func NewSymbolTable() *SymbolTable {
	return &SymbolTable{vars: make(map[string]string)}
}

// Define registers a variable name with its node ID. Overwrites any previous binding.
func (s *SymbolTable) Define(name, nodeID string) {
	s.vars[name] = nodeID
}

// Resolve looks up the most recent node ID for the given variable name.
func (s *SymbolTable) Resolve(name string) (nodeID string, ok bool) {
	nodeID, ok = s.vars[name]
	return
}
