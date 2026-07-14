package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ugiordan/architecture-analyzer/pkg/builder"
	"github.com/ugiordan/architecture-analyzer/pkg/graph"
)

type quickIndexOutput struct {
	SchemaVersion int                `json:"schema_version"`
	RepoPath      string             `json:"repo_path"`
	IndexedAt     string             `json:"indexed_at"`
	Stats         quickIndexStats    `json:"stats"`
	Functions     []quickIndexFunc   `json:"functions"`
	Calls         []quickIndexCall   `json:"calls"`
	HTTPEndpoints []quickIndexHTTP   `json:"http_endpoints,omitempty"`
	DBOperations  []quickIndexDB     `json:"db_operations,omitempty"`
	Classes       []quickIndexClass  `json:"classes,omitempty"`
}

type quickIndexStats struct {
	Functions     int     `json:"functions"`
	CallSites     int     `json:"call_sites"`
	CallEdges     int     `json:"call_edges"`
	HTTPEndpoints int     `json:"http_endpoints"`
	DBOperations  int     `json:"db_operations"`
	Classes       int     `json:"classes"`
	ParseTimeMS   int64   `json:"parse_time_ms"`
}

type quickIndexFunc struct {
	Name       string   `json:"name"`
	File       string   `json:"file"`
	Line       int      `json:"line"`
	EndLine    int      `json:"end_line,omitempty"`
	Language   string   `json:"language,omitempty"`
	Kind       string   `json:"kind"`
	Receiver   string   `json:"receiver,omitempty"`
	Params     []string `json:"params,omitempty"`
	ReturnType string   `json:"return_type,omitempty"`
	Complexity int      `json:"complexity,omitempty"`
	IsTest     bool     `json:"is_test,omitempty"`
}

type quickIndexCall struct {
	Caller     string `json:"caller"`
	CallerFile string `json:"caller_file"`
	Callee     string `json:"callee"`
	CalleeFile string `json:"callee_file"`
	Line       int    `json:"line"`
	Confidence string `json:"confidence"`
}

type quickIndexHTTP struct {
	Name       string `json:"name"`
	Route      string `json:"route"`
	Method     string `json:"method,omitempty"`
	File       string `json:"file"`
	Line       int    `json:"line"`
}

type quickIndexDB struct {
	Name string `json:"name"`
	File string `json:"file"`
	Line int    `json:"line"`
}

type quickIndexClass struct {
	Name    string   `json:"name"`
	File    string   `json:"file"`
	Line    int      `json:"line"`
	Methods []string `json:"methods,omitempty"`
}

func cmdQuickIndex(args []string) error {
	fs := flag.NewFlagSet("quick-index", flag.ExitOnError)
	output := fs.String("output", "quick-index.json", "Output JSON file")
	fs.Parse(reorderArgs(fs, args))

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: arch-analyzer quick-index [--output file] <repo-path>")
	}
	repoPath := fs.Arg(0)

	start := time.Now()

	b := builder.NewBuilder()
	cpg, err := b.BuildFromDir(repoPath)
	if err != nil {
		return fmt.Errorf("building index: %w", err)
	}

	elapsed := time.Since(start)

	idx := buildQuickIndex(cpg, repoPath, elapsed)

	data, err := json.MarshalIndent(idx, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling index: %w", err)
	}

	if *output == "-" {
		_, err = os.Stdout.Write(data)
		return err
	}

	if err := os.WriteFile(*output, data, 0o644); err != nil {
		return fmt.Errorf("writing output: %w", err)
	}

	fmt.Printf("Quick index: %d functions, %d call edges, %d HTTP endpoints in %dms\n",
		idx.Stats.Functions, idx.Stats.CallEdges, idx.Stats.HTTPEndpoints, idx.Stats.ParseTimeMS)
	fmt.Printf("Written to: %s\n", *output)
	return nil
}

func buildQuickIndex(cpg *graph.CPG, repoPath string, elapsed time.Duration) quickIndexOutput {
	functions := cpg.NodesByKind(graph.NodeFunction)
	callSites := cpg.NodesByKind(graph.NodeCallSite)
	httpEndpoints := cpg.NodesByKind(graph.NodeHTTPEndpoint)
	dbOps := cpg.NodesByKind(graph.NodeDBOperation)
	classes := cpg.NodesByKind(graph.NodeClass)

	// Build function index
	var funcs []quickIndexFunc
	for _, fn := range functions {
		kind := "function"
		if fn.TypeName != "" {
			kind = "method"
		}
		if fn.IsTest {
			kind = "test"
		}

		f := quickIndexFunc{
			Name:       fn.Name,
			File:       fn.File,
			Line:       fn.Line,
			EndLine:    fn.EndLine,
			Language:   fn.Language,
			Kind:       kind,
			Receiver:   fn.TypeName,
			ReturnType: fn.ReturnType,
			Complexity: fn.Complexity,
			IsTest:     fn.IsTest,
		}
		if len(fn.ParamNames) > 0 {
			for i, name := range fn.ParamNames {
				p := name
				if i < len(fn.ParamTypes) && fn.ParamTypes[i] != "" {
					p += " " + fn.ParamTypes[i]
				}
				f.Params = append(f.Params, p)
			}
		}
		funcs = append(funcs, f)
	}

	// Build call edges (only resolved EdgeCalls, not dataflow)
	funcByID := make(map[string]*graph.Node)
	for _, fn := range functions {
		funcByID[fn.ID] = fn
	}
	callSiteByID := make(map[string]*graph.Node)
	for _, cs := range callSites {
		callSiteByID[cs.ID] = cs
	}

	var calls []quickIndexCall
	for _, edge := range cpg.Edges() {
		if edge.Kind != graph.EdgeCalls {
			continue
		}
		from := callSiteByID[edge.From]
		to := funcByID[edge.To]
		if from == nil || to == nil {
			continue
		}

		conf := "certain"
		switch edge.Confidence {
		case graph.ConfidenceInferred:
			conf = "inferred"
		case graph.ConfidenceUncertain:
			conf = "uncertain"
		}

		calls = append(calls, quickIndexCall{
			Caller:     from.Name,
			CallerFile: from.File,
			Callee:     to.Name,
			CalleeFile: to.File,
			Line:       from.Line,
			Confidence: conf,
		})
	}

	// HTTP endpoints
	var httpList []quickIndexHTTP
	for _, ep := range httpEndpoints {
		httpList = append(httpList, quickIndexHTTP{
			Name:   ep.Name,
			Route:  ep.Route,
			Method: ep.HTTPMethod,
			File:   ep.File,
			Line:   ep.Line,
		})
	}

	// DB operations
	var dbList []quickIndexDB
	for _, op := range dbOps {
		dbList = append(dbList, quickIndexDB{
			Name: op.Name,
			File: op.File,
			Line: op.Line,
		})
	}

	// Classes with methods
	var classList []quickIndexClass
	for _, cls := range classes {
		c := quickIndexClass{
			Name: cls.Name,
			File: cls.File,
			Line: cls.Line,
		}
		for _, fn := range functions {
			if fn.TypeName == cls.Name && fn.File == cls.File {
				c.Methods = append(c.Methods, fn.Name)
			}
		}
		classList = append(classList, c)
	}

	return quickIndexOutput{
		SchemaVersion: 1,
		RepoPath:      repoPath,
		IndexedAt:     time.Now().UTC().Format(time.RFC3339),
		Stats: quickIndexStats{
			Functions:     len(functions),
			CallSites:     len(callSites),
			CallEdges:     len(calls),
			HTTPEndpoints: len(httpEndpoints),
			DBOperations:  len(dbOps),
			Classes:       len(classes),
			ParseTimeMS:   elapsed.Milliseconds(),
		},
		Functions:     funcs,
		Calls:         calls,
		HTTPEndpoints: httpList,
		DBOperations:  dbList,
		Classes:       classList,
	}
}
