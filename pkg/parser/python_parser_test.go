package parser

import (
	"os"
	"testing"

	"github.com/ugiordan/architecture-analyzer/pkg/graph"
)

func TestPythonParserLanguageAndExtensions(t *testing.T) {
	p := NewPythonParser()
	if p.Language() != "python" {
		t.Errorf("expected language 'python', got %q", p.Language())
	}
	exts := p.Extensions()
	if len(exts) != 1 || exts[0] != ".py" {
		t.Errorf("expected extensions [.py], got %v", exts)
	}
}

func TestPythonParserFlaskApp(t *testing.T) {
	content, err := os.ReadFile("../../testdata/flask_app.py")
	if err != nil {
		t.Fatalf("Failed to read test fixture: %v", err)
	}

	p := NewPythonParser()
	result, err := p.ParseFile("testdata/flask_app.py", content)
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}

	// --- Functions ---
	t.Run("functions", func(t *testing.T) {
		if len(result.Functions) < 5 {
			t.Errorf("expected at least 5 functions, got %d", len(result.Functions))
			for _, fn := range result.Functions {
				t.Logf("  function: %s (type=%s)", fn.Name, fn.TypeName)
			}
		}

		fnMap := make(map[string]*graph.Node)
		for _, fn := range result.Functions {
			fnMap[fn.Name] = fn
		}

		for _, expected := range []string{"get_users", "create_user", "delete_user", "run_backup", "run_migration"} {
			if fnMap[expected] == nil {
				t.Errorf("expected function %q not found", expected)
			}
		}

		// All functions should have language "python"
		for _, fn := range result.Functions {
			if fn.Language != "python" {
				t.Errorf("function %q has language %q, expected 'python'", fn.Name, fn.Language)
			}
		}
	})

	// --- Methods with TypeName ---
	t.Run("methods", func(t *testing.T) {
		fnMap := make(map[string]*graph.Node)
		for _, fn := range result.Functions {
			fnMap[fn.Name] = fn
		}

		for _, method := range []string{"get_all", "create"} {
			fn := fnMap[method]
			if fn == nil {
				t.Errorf("expected method %q not found", method)
				continue
			}
			if fn.TypeName != "UserService" {
				t.Errorf("method %q TypeName = %q, want 'UserService'", method, fn.TypeName)
			}
		}
	})

	// --- HTTP handlers ---
	t.Run("http_handlers", func(t *testing.T) {
		if len(result.HTTPHandlers) < 3 {
			t.Errorf("expected at least 3 HTTP handlers, got %d", len(result.HTTPHandlers))
			for _, h := range result.HTTPHandlers {
				t.Logf("  handler: %s route=%s", h.Name, h.Route)
			}
		}

		hasUsersRoute := false
		for _, h := range result.HTTPHandlers {
			if h.Route == "/users" {
				hasUsersRoute = true
				break
			}
		}
		if !hasUsersRoute {
			t.Error("expected an HTTP handler with route '/users'")
		}
	})

	// --- Call sites ---
	t.Run("call_sites", func(t *testing.T) {
		if len(result.CallSites) == 0 {
			t.Error("expected call sites, got 0")
		}

		hasSubprocessRun := false
		for _, cs := range result.CallSites {
			if cs.Name == "subprocess.run" {
				hasSubprocessRun = true
				break
			}
		}
		if !hasSubprocessRun {
			t.Error("expected call site 'subprocess.run'")
			for _, cs := range result.CallSites {
				t.Logf("  call: %s", cs.Name)
			}
		}
	})

	// --- DB operations ---
	t.Run("db_operations", func(t *testing.T) {
		if len(result.DBOperations) < 2 {
			t.Errorf("expected at least 2 DB operations, got %d", len(result.DBOperations))
			for _, op := range result.DBOperations {
				t.Logf("  db op: %s (op=%s)", op.Name, op.Operation)
			}
		}

		hasRead, hasWrite := false, false
		for _, op := range result.DBOperations {
			switch op.Operation {
			case "read":
				hasRead = true
			case "write":
				hasWrite = true
			}
		}
		if !hasRead {
			t.Error("expected a DB read operation")
		}
		if !hasWrite {
			t.Error("expected a DB write operation")
		}
	})

	// --- Struct literals (class instantiations) ---
	t.Run("struct_literals", func(t *testing.T) {
		if len(result.StructLiterals) == 0 {
			t.Error("expected struct literals (class instantiations), got 0")
		}

		names := make(map[string]bool)
		for _, sl := range result.StructLiterals {
			names[sl.Name] = true
		}

		if !names["UserService"] && !names["User"] {
			t.Error("expected UserService or User class instantiation")
			for _, sl := range result.StructLiterals {
				t.Logf("  struct literal: %s", sl.Name)
			}
		}
	})

	// --- Decorators ---
	t.Run("decorators", func(t *testing.T) {
		var getUsersFn *graph.Node
		for _, fn := range result.Functions {
			if fn.Name == "get_users" {
				getUsersFn = fn
				break
			}
		}
		if getUsersFn == nil {
			t.Fatal("expected to find get_users function")
		}
		if len(getUsersFn.Decorators) == 0 {
			t.Error("expected get_users to have decorators")
		}
	})
}

func TestPythonParserComputesComplexity(t *testing.T) {
	content, err := os.ReadFile("../../testdata/complexity_sample.py")
	if err != nil {
		t.Fatalf("Failed to read test fixture: %v", err)
	}

	p := NewPythonParser()
	result, err := p.ParseFile("testdata/complexity_sample.py", content)
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}

	expected := map[string]int{
		"simple_func":        1,
		"complex_func":       6, // if + and + elif + for + if + base (or not counted)
		"loop_func":          4, // while + except + for + base
		"comprehension_func": 3, // if + comprehension-if + base
	}

	for _, fn := range result.Functions {
		if want, ok := expected[fn.Name]; ok {
			if fn.Complexity != want {
				t.Errorf("function %s: complexity = %d, want %d", fn.Name, fn.Complexity, want)
			}
			delete(expected, fn.Name)
		}
	}
	for name := range expected {
		t.Errorf("function %s not found in parse result", name)
	}
}

func TestPythonParserFastAPIApp(t *testing.T) {
	content, err := os.ReadFile("../../testdata/fastapi_app.py")
	if err != nil {
		t.Fatalf("Failed to read test fixture: %v", err)
	}

	p := NewPythonParser()
	result, err := p.ParseFile("testdata/fastapi_app.py", content)
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}

	// --- HTTP handlers ---
	t.Run("http_handlers", func(t *testing.T) {
		if len(result.HTTPHandlers) < 2 {
			t.Errorf("expected at least 2 HTTP handlers, got %d", len(result.HTTPHandlers))
			for _, h := range result.HTTPHandlers {
				t.Logf("  handler: %s route=%s", h.Name, h.Route)
			}
		}
	})

	// --- Call sites include pickle.loads ---
	t.Run("pickle_loads", func(t *testing.T) {
		found := false
		for _, cs := range result.CallSites {
			if cs.Name == "pickle.loads" {
				found = true
				break
			}
		}
		if !found {
			t.Error("expected call site 'pickle.loads'")
			for _, cs := range result.CallSites {
				t.Logf("  call: %s", cs.Name)
			}
		}
	})

	// --- Struct literal: DataProcessor ---
	t.Run("data_processor_instantiation", func(t *testing.T) {
		found := false
		for _, sl := range result.StructLiterals {
			if sl.Name == "DataProcessor" {
				found = true
				break
			}
		}
		if !found {
			t.Error("expected DataProcessor class instantiation")
			for _, sl := range result.StructLiterals {
				t.Logf("  struct literal: %s", sl.Name)
			}
		}
	})
}

func TestPythonParserClassHierarchy(t *testing.T) {
	content := []byte(`
from abc import ABC, abstractmethod

class BaseVectorStore(ABC):
    @abstractmethod
    def insert(self, data):
        pass

    @abstractmethod
    def query(self, q):
        pass

class ChromaVectorStore(BaseVectorStore):
    def insert(self, data):
        self.client.add(data)

    def query(self, q):
        return self.client.query(q)

class LSVectorStore(BaseVectorStore):
    def insert(self, data):
        self.ls_client.vector_io.insert(data)

    def query(self, q):
        return self.ls_client.vector_io.query(q)

class SimpleClass:
    def method(self):
        pass
`)

	p := NewPythonParser()
	result, err := p.ParseFile("vector_store.py", content)
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}

	t.Run("class_nodes_created", func(t *testing.T) {
		if len(result.Classes) != 4 {
			t.Errorf("expected 4 class nodes, got %d", len(result.Classes))
			for _, c := range result.Classes {
				t.Logf("  class: %s bases=%v", c.Name, c.BaseClasses)
			}
		}
	})

	t.Run("base_classes_extracted", func(t *testing.T) {
		clsMap := make(map[string]*graph.Node)
		for _, c := range result.Classes {
			clsMap[c.Name] = c
		}

		// BaseVectorStore inherits from ABC
		if base := clsMap["BaseVectorStore"]; base == nil {
			t.Error("BaseVectorStore class not found")
		} else if len(base.BaseClasses) != 1 || base.BaseClasses[0] != "ABC" {
			t.Errorf("BaseVectorStore.BaseClasses = %v, want [ABC]", base.BaseClasses)
		}

		// ChromaVectorStore inherits from BaseVectorStore
		if chroma := clsMap["ChromaVectorStore"]; chroma == nil {
			t.Error("ChromaVectorStore class not found")
		} else if len(chroma.BaseClasses) != 1 || chroma.BaseClasses[0] != "BaseVectorStore" {
			t.Errorf("ChromaVectorStore.BaseClasses = %v, want [BaseVectorStore]", chroma.BaseClasses)
		}

		// LSVectorStore inherits from BaseVectorStore
		if ls := clsMap["LSVectorStore"]; ls == nil {
			t.Error("LSVectorStore class not found")
		} else if len(ls.BaseClasses) != 1 || ls.BaseClasses[0] != "BaseVectorStore" {
			t.Errorf("LSVectorStore.BaseClasses = %v, want [BaseVectorStore]", ls.BaseClasses)
		}

		// SimpleClass has no base classes
		if simple := clsMap["SimpleClass"]; simple == nil {
			t.Error("SimpleClass class not found")
		} else if len(simple.BaseClasses) != 0 {
			t.Errorf("SimpleClass.BaseClasses = %v, want []", simple.BaseClasses)
		}
	})

	t.Run("class_nodes_have_correct_kind", func(t *testing.T) {
		for _, c := range result.Classes {
			if c.Kind != graph.NodeClass {
				t.Errorf("class %s has kind %s, want Class", c.Name, c.Kind)
			}
			if c.Language != "python" {
				t.Errorf("class %s has language %s, want python", c.Name, c.Language)
			}
		}
	})

	t.Run("methods_have_type_name", func(t *testing.T) {
		methods := map[string]string{}
		for _, fn := range result.Functions {
			methods[fn.Name] = fn.TypeName
		}
		if methods["insert"] != "LSVectorStore" && methods["insert"] != "ChromaVectorStore" {
			t.Errorf("insert method TypeName = %q, want a VectorStore class", methods["insert"])
		}
	})
}
