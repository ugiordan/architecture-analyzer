package dataflow

import "testing"

func TestSymbolTableDefineAndResolve(t *testing.T) {
	st := NewSymbolTable()
	st.Define("x", "var_abc123")

	id, ok := st.Resolve("x")
	if !ok {
		t.Fatal("expected x to resolve")
	}
	if id != "var_abc123" {
		t.Errorf("got %q, want %q", id, "var_abc123")
	}
}

func TestSymbolTableResolveUndefined(t *testing.T) {
	st := NewSymbolTable()

	_, ok := st.Resolve("x")
	if ok {
		t.Error("expected x to not resolve")
	}
}

func TestSymbolTableLastWriteWins(t *testing.T) {
	st := NewSymbolTable()
	st.Define("x", "var_first")
	st.Define("x", "var_second")

	id, ok := st.Resolve("x")
	if !ok {
		t.Fatal("expected x to resolve")
	}
	if id != "var_second" {
		t.Errorf("got %q, want %q", id, "var_second")
	}
}

func TestSymbolTableMultipleVariables(t *testing.T) {
	st := NewSymbolTable()
	st.Define("a", "var_a")
	st.Define("b", "var_b")
	st.Define("c", "var_c")

	for _, tc := range []struct{ name, wantID string }{
		{"a", "var_a"},
		{"b", "var_b"},
		{"c", "var_c"},
	} {
		id, ok := st.Resolve(tc.name)
		if !ok {
			t.Errorf("%s: expected to resolve", tc.name)
		}
		if id != tc.wantID {
			t.Errorf("%s: got %q, want %q", tc.name, id, tc.wantID)
		}
	}
}
