package parser

import (
	"context"
	"fmt"
	"strings"
	"sync/atomic"
	"unicode"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/python"

	"github.com/ugiordan/architecture-analyzer/pkg/graph"
)

// PythonSkipDirs lists directories that should be skipped when scanning Python projects.
var PythonSkipDirs = []string{"__pycache__", "migrations", ".tox", "venv", "site-packages", ".venv", "env"}

// PythonTestPatterns lists filename patterns that identify test files.
var PythonTestPatterns = []string{"test_*", "*_test.py"}

// pythonBuiltins is the set of PascalCase names that are builtins or standard
// exception types, not user-defined class instantiations.
var pythonBuiltins = map[string]bool{
	"True": true, "False": true, "None": true,
	"Exception": true, "BaseException": true, "ValueError": true,
	"TypeError": true, "KeyError": true, "IndexError": true,
	"AttributeError": true, "RuntimeError": true, "OSError": true,
	"IOError": true, "ImportError": true, "StopIteration": true,
	"NotImplementedError": true, "PermissionError": true,
	"FileNotFoundError": true, "ConnectionError": true,
	"TimeoutError": true, "UnicodeError": true,
}

// httpMethods is the set of decorator method names that indicate HTTP route handlers.
var httpMethods = map[string]bool{
	"route": true, "get": true, "post": true,
	"put": true, "delete": true, "patch": true,
}

// dbCallOps maps Python DB call patterns to their operation type (read/write).
var dbCallOps = map[string]string{
	"session.query":      "read",
	"session.execute":    "write",
	"session.add":        "write",
	"session.commit":     "write",
	"cursor.execute":     "write",
	"db.execute":         "write",
	"connection.execute": "write",
}

// PythonParser extracts code property graph nodes from Python source files using tree-sitter.
// Each goroutine MUST use its own PythonParser instance (tree-sitter parsers are not thread-safe).
// When running multiple instances in parallel, pass a shared counter via NewPythonParserWithSeq.
type PythonParser struct {
	parser *sitter.Parser
	idSeq  *atomic.Int64
}

// NewPythonParser creates a parser for Python source files backed by tree-sitter.
func NewPythonParser() *PythonParser {
	p := sitter.NewParser()
	p.SetLanguage(python.GetLanguage())
	return &PythonParser{parser: p, idSeq: &atomic.Int64{}}
}

// NewPythonParserWithSeq creates a parser that shares an ID counter with other instances.
func NewPythonParserWithSeq(seq *atomic.Int64) *PythonParser {
	p := sitter.NewParser()
	p.SetLanguage(python.GetLanguage())
	return &PythonParser{parser: p, idSeq: seq}
}

func (pp *PythonParser) Language() string     { return "python" }
func (pp *PythonParser) Extensions() []string { return []string{".py"} }

func (pp *PythonParser) nextID(prefix string) string {
	id := pp.idSeq.Add(1)
	return fmt.Sprintf("%s_%d", prefix, id)
}

// ParseFile parses a Python source file and returns extracted nodes and edges.
func (pp *PythonParser) ParseFile(path string, content []byte) (*ParseResult, error) {
	if len(content) > maxFileSize {
		return nil, fmt.Errorf("file too large (%d bytes, max %d)", len(content), maxFileSize)
	}
	tree, err := pp.parser.ParseCtx(context.Background(), nil, content)
	if err != nil {
		return nil, fmt.Errorf("tree-sitter parse failed: %w", err)
	}
	defer tree.Close()

	result := &ParseResult{}
	root := tree.RootNode()
	pp.walk(root, content, path, "", result)
	return result, nil
}

// walk recursively traverses the AST. className tracks the enclosing class for method extraction.
func (pp *PythonParser) walk(node *sitter.Node, src []byte, file, className string, result *ParseResult) {
	switch node.Type() {
	case "function_definition":
		pp.extractFunction(node, src, file, className, nil, result)
		return // children handled inside extractFunction
	case "decorated_definition":
		pp.extractDecorated(node, src, file, className, result)
		return // children handled inside extractDecorated
	case "class_definition":
		pp.extractClass(node, src, file, result)
		return // children handled inside extractClass
	case "call":
		pp.extractCallSite(node, src, file, result)
	}

	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		if child != nil {
			pp.walk(child, src, file, className, result)
		}
	}
}

// extractClass processes a class_definition node, walking its body with the class name set.
func (pp *PythonParser) extractClass(node *sitter.Node, src []byte, file string, result *ParseResult) {
	nameNode := node.ChildByFieldName("name")
	if nameNode == nil {
		return
	}
	clsName := nameNode.Content(src)

	body := node.ChildByFieldName("body")
	if body == nil {
		return
	}
	for i := 0; i < int(body.ChildCount()); i++ {
		child := body.Child(i)
		if child != nil {
			pp.walk(child, src, file, clsName, result)
		}
	}
}

// extractDecorated handles a decorated_definition node: collects decorators, then
// delegates to extractFunction for the inner function_definition.
func (pp *PythonParser) extractDecorated(node *sitter.Node, src []byte, file, className string, result *ParseResult) {
	var decorators []string
	var fnNode *sitter.Node

	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		if child == nil {
			continue
		}
		switch child.Type() {
		case "decorator":
			decorators = append(decorators, child.Content(src))
		case "function_definition":
			fnNode = child
		}
	}

	if fnNode != nil {
		pp.extractFunction(fnNode, src, file, className, decorators, result)
	}
}

// extractFunction creates a Function node and checks decorators for HTTP route patterns.
func (pp *PythonParser) extractFunction(node *sitter.Node, src []byte, file, className string, decorators []string, result *ParseResult) {
	nameNode := node.ChildByFieldName("name")
	if nameNode == nil {
		return
	}
	name := nameNode.Content(src)

	fn := &graph.Node{
		ID:          pp.nextID("fn"),
		Kind:        graph.NodeFunction,
		Name:        name,
		File:        file,
		Line:        int(node.StartPoint().Row) + 1,
		EndLine:     int(node.EndPoint().Row) + 1,
		Language:    "python",
		TypeName:    className,
		Decorators:  decorators,
		Annotations: make(map[string]bool),
		Properties:  make(map[string]string),
	}
	result.Functions = append(result.Functions, fn)

	// Check decorators for HTTP route patterns
	for _, dec := range decorators {
		pp.maybeExtractHTTPHandler(dec, fn, file, result)
	}

	// Walk function body for call sites, etc.
	body := node.ChildByFieldName("body")
	if body != nil {
		for i := 0; i < int(body.ChildCount()); i++ {
			child := body.Child(i)
			if child != nil {
				pp.walk(child, src, file, className, result)
			}
		}
	}
}

// maybeExtractHTTPHandler checks if a decorator string matches an HTTP route pattern
// like @app.route("/path"), @app.get("/path"), @router.post("/path"), etc.
func (pp *PythonParser) maybeExtractHTTPHandler(decorator string, fn *graph.Node, file string, result *ParseResult) {
	// Strip the leading @
	dec := strings.TrimPrefix(decorator, "@")

	// We need: <identifier>.<method>(...) where method is in httpMethods
	// Find the method call part
	parenIdx := strings.Index(dec, "(")
	if parenIdx < 0 {
		return
	}
	callPart := dec[:parenIdx]
	dotIdx := strings.LastIndex(callPart, ".")
	if dotIdx < 0 {
		return
	}
	method := callPart[dotIdx+1:]
	if !httpMethods[method] {
		return
	}

	// Extract route path from first string argument
	argPart := dec[parenIdx+1:]
	route := extractStringArg(argPart)

	handler := &graph.Node{
		ID:          pp.nextID("http"),
		Kind:        graph.NodeHTTPEndpoint,
		Name:        fn.Name,
		File:        file,
		Line:        fn.Line,
		Language:    "python",
		Annotations: make(map[string]bool),
		Properties:  make(map[string]string),
	}
	if route != "" {
		handler.Properties["route"] = route
	}
	handler.Properties["method"] = strings.ToUpper(method)
	result.HTTPHandlers = append(result.HTTPHandlers, handler)
}

// extractStringArg extracts the first quoted string from a decorator argument list.
// Input example: `"/users", methods=["GET"])`
func extractStringArg(s string) string {
	for _, q := range []byte{'"', '\''} {
		start := strings.IndexByte(s, q)
		if start < 0 {
			continue
		}
		end := strings.IndexByte(s[start+1:], q)
		if end < 0 {
			continue
		}
		return s[start+1 : start+1+end]
	}
	return ""
}

// extractCallSite creates a CallSite node from a call expression, and detects
// DB operations and class instantiations (struct literals).
func (pp *PythonParser) extractCallSite(node *sitter.Node, src []byte, file string, result *ParseResult) {
	fnNode := node.ChildByFieldName("function")
	if fnNode == nil {
		return
	}
	callText := fnNode.Content(src)
	line := int(node.StartPoint().Row) + 1

	cs := &graph.Node{
		ID:         pp.nextID("call"),
		Kind:       graph.NodeCallSite,
		Name:       callText,
		File:       file,
		Line:       line,
		Language:   "python",
		Properties: make(map[string]string),
	}
	result.CallSites = append(result.CallSites, cs)

	// Check for DB operations
	if op, ok := dbCallOps[callText]; ok {
		dbOp := &graph.Node{
			ID:         pp.nextID("db"),
			Kind:       graph.NodeDBOperation,
			Name:       callText,
			File:       file,
			Line:       line,
			Language:   "python",
			Properties: map[string]string{"operation": op},
		}
		result.DBOperations = append(result.DBOperations, dbOp)
	}

	// Also check for dotted DB patterns where the receiver varies:
	// e.g., "db.session.execute", "self.session.query"
	if !strings.Contains(callText, ".") {
		// Simple call: check if it's a PascalCase class instantiation
		pp.maybeExtractStructLiteral(callText, node, src, file, line, result)
		return
	}

	// For dotted calls, check the last two segments against dbCallOps
	parts := strings.Split(callText, ".")
	if len(parts) >= 2 {
		suffix := parts[len(parts)-2] + "." + parts[len(parts)-1]
		if op, ok := dbCallOps[suffix]; ok {
			// Avoid duplicate if already matched the full callText
			if _, alreadyMatched := dbCallOps[callText]; !alreadyMatched {
				dbOp := &graph.Node{
					ID:         pp.nextID("db"),
					Kind:       graph.NodeDBOperation,
					Name:       callText,
					File:       file,
					Line:       line,
					Language:   "python",
					Properties: map[string]string{"operation": op},
				}
				result.DBOperations = append(result.DBOperations, dbOp)
			}
		}
	}
}

// maybeExtractStructLiteral checks if a simple (non-dotted) call name looks like
// a PascalCase class instantiation and, if so, creates a StructLiteral node.
func (pp *PythonParser) maybeExtractStructLiteral(name string, node *sitter.Node, src []byte, file string, line int, result *ParseResult) {
	if len(name) == 0 {
		return
	}
	// Must start with an uppercase letter
	if !unicode.IsUpper(rune(name[0])) {
		return
	}
	// Skip Python builtins
	if pythonBuiltins[name] {
		return
	}

	sl := &graph.Node{
		ID:         pp.nextID("struct"),
		Kind:       graph.NodeStructLiteral,
		Name:       name,
		File:       file,
		Line:       line,
		Language:   "python",
		Properties: make(map[string]string),
	}
	result.StructLiterals = append(result.StructLiterals, sl)
}
