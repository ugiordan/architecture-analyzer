package layers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ugiordan/architecture-analyzer/pkg/extractor"
	"github.com/ugiordan/architecture-analyzer/pkg/graph"
	"github.com/ugiordan/architecture-analyzer/pkg/query"
)

func writeTestFileWithDir(t *testing.T, dir, name string, lines int) {
	t.Helper()
	full := filepath.Join(dir, name)
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		t.Fatal(err)
	}
	var sb []byte
	for i := 1; i <= lines; i++ {
		sb = append(sb, []byte("// line "+string(rune('0'+i%10))+"\n")...)
	}
	if err := os.WriteFile(full, sb, 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestUpgradeSelector_SelectsConversionFunctions(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "ConvertTo",
		File: "api/v1/convert.go", Line: 10, EndLine: 30,
		Language: "go", TypeName: "Widget",
		Annotations: map[string]bool{annotVersionConversion: true},
	})
	cpg.AddNode(&graph.Node{
		ID: "fn2", Kind: graph.NodeFunction, Name: "ConvertFrom",
		File: "api/v1/convert.go", Line: 35, EndLine: 55,
		Language: "go", TypeName: "Widget",
		Annotations: map[string]bool{annotVersionConversion: true},
	})

	dir := t.TempDir()
	writeTestFileWithDir(t, dir, "api/v1/convert.go", 60)

	sel := NewUpgradeSelector(dir)
	layer, _ := sel.Select(cpg, nil, nil, nil)

	if layer.Name != "upgrade" {
		t.Errorf("layer name = %q, want %q", layer.Name, "upgrade")
	}

	names := make(map[string]bool)
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			names[fn.Name] = true
			if fn.ReceiverType != "Widget" {
				t.Errorf("%s ReceiverType = %q, want Widget", fn.Name, fn.ReceiverType)
			}
		}
	}
	if !names["ConvertTo"] {
		t.Error("ConvertTo not selected")
	}
	if !names["ConvertFrom"] {
		t.Error("ConvertFrom not selected")
	}
}

func TestUpgradeSelector_SelectsMigrationFunctions(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "MigrateDatabase",
		File: "migrate.go", Line: 5, EndLine: 25,
		Language:    "go",
		Annotations: map[string]bool{annotMigration: true},
	})

	dir := t.TempDir()
	writeTestFile(t, dir, "migrate.go", 30)

	sel := NewUpgradeSelector(dir)
	layer, _ := sel.Select(cpg, nil, nil, nil)

	found := false
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "MigrateDatabase" {
				found = true
				hasMeta := false
				for _, m := range fn.Metas {
					if m.Key == "migration" && m.Value == "true" {
						hasMeta = true
					}
				}
				if !hasMeta {
					t.Error("expected migration=true meta")
				}
			}
		}
	}
	if !found {
		t.Error("MigrateDatabase not selected")
	}
}

func TestUpgradeSelector_SelectsFeatureGateFunctions(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "enableNewFeature",
		File: "feature.go", Line: 10, EndLine: 20,
		Language:    "go",
		Annotations: map[string]bool{annotFeatureGate: true},
	})

	dir := t.TempDir()
	writeTestFile(t, dir, "feature.go", 25)

	sel := NewUpgradeSelector(dir)
	layer, _ := sel.Select(cpg, nil, nil, nil)

	found := false
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "enableNewFeature" {
				found = true
				hasMeta := false
				for _, m := range fn.Metas {
					if m.Key == "feature-gate" && m.Value == "true" {
						hasMeta = true
					}
				}
				if !hasMeta {
					t.Error("expected feature-gate=true meta")
				}
			}
		}
	}
	if !found {
		t.Error("enableNewFeature not selected")
	}
}

func TestUpgradeSelector_SelectsVersionCheckFunctions(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "checkClusterVersion",
		File: "version.go", Line: 5,
		Language:    "go",
		Annotations: map[string]bool{annotVersionCheck: true},
	})

	sel := NewUpgradeSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	found := false
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "checkClusterVersion" {
				found = true
				hasMeta := false
				for _, m := range fn.Metas {
					if m.Key == "version-check" && m.Value == "true" {
						hasMeta = true
					}
				}
				if !hasMeta {
					t.Error("expected version-check=true meta")
				}
			}
		}
	}
	if !found {
		t.Error("checkClusterVersion not selected")
	}
}

func TestUpgradeSelector_SelectsPreReleaseAPIFunctions(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "useAlphaAPI",
		File: "alpha.go", Line: 5,
		Language:    "go",
		Annotations: map[string]bool{annotPreReleaseAPI: true},
	})

	sel := NewUpgradeSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	found := false
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "useAlphaAPI" {
				found = true
				hasMeta := false
				for _, m := range fn.Metas {
					if m.Key == "pre-release-api" && m.Value == "true" {
						hasMeta = true
					}
				}
				if !hasMeta {
					t.Error("expected pre-release-api=true meta")
				}
			}
		}
	}
	if !found {
		t.Error("useAlphaAPI not selected")
	}
}

func TestUpgradeSelector_FiltersUpgradeFindings(t *testing.T) {
	cpg := graph.NewCPG()
	findings := []query.Finding{
		{RuleID: "CGA-U01", Severity: "medium", Domain: "upgrade", Message: "Unconverted CRD", File: "a.go", Line: 1},
		{RuleID: "CGA-U02", Severity: "low", Domain: "upgrade", Message: "Pre-release API", File: "b.go", Line: 2},
		{RuleID: "CGA-S01", Severity: "high", Domain: "security", Message: "Security issue", File: "c.go", Line: 3},
		{RuleID: "CGA-T01", Severity: "medium", Domain: "testing", Message: "Testing issue", File: "d.go", Line: 4},
	}

	sel := NewUpgradeSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, findings, nil)

	if len(layer.Findings) != 2 {
		t.Fatalf("expected 2 upgrade findings, got %d", len(layer.Findings))
	}
	for _, f := range layer.Findings {
		if f.Domain != "upgrade" {
			t.Errorf("unexpected domain %q", f.Domain)
		}
	}
}

func TestUpgradeSelector_FindingsSorted(t *testing.T) {
	cpg := graph.NewCPG()
	findings := []query.Finding{
		{RuleID: "CGA-U02", Severity: "low", Domain: "upgrade", Message: "Low", File: "a.go", Line: 1},
		{RuleID: "CGA-U01", Severity: "medium", Domain: "upgrade", Message: "Medium", File: "b.go", Line: 2},
	}

	sel := NewUpgradeSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, findings, nil)

	if len(layer.Findings) != 2 {
		t.Fatalf("expected 2 findings, got %d", len(layer.Findings))
	}
	if layer.Findings[0].Severity != "medium" {
		t.Errorf("first finding severity = %q, want medium", layer.Findings[0].Severity)
	}
	if layer.Findings[1].Severity != "low" {
		t.Errorf("second finding severity = %q, want low", layer.Findings[1].Severity)
	}
}

func TestUpgradeSelector_FeatureGateResources(t *testing.T) {
	cpg := graph.NewCPG()
	arch := &extractor.ComponentArchitecture{
		Component: "test",
		FeatureGates: []extractor.FeatureGate{
			{Name: "NewInference", Default: true, PreRelease: "Beta", Source: "pkg/features/gates.go:15"},
			{Name: "ExperimentalCache", Default: false, PreRelease: "Alpha", LockToDefault: true, Source: "pkg/features/gates.go:20"},
		},
	}

	sel := NewUpgradeSelector(t.TempDir())
	layer, _ := sel.Select(cpg, arch, nil, nil)

	if len(layer.Resources) != 2 {
		t.Fatalf("expected 2 FeatureGate resources, got %d", len(layer.Resources))
	}

	for _, r := range layer.Resources {
		if r.Kind != "FeatureGate" {
			t.Errorf("resource kind = %q, want FeatureGate", r.Kind)
		}
	}

	found := false
	for _, r := range layer.Resources {
		if r.Name == "ExperimentalCache" {
			found = true
			if !strings.Contains(r.Summary, "disabled") {
				t.Errorf("summary should contain 'disabled', got %q", r.Summary)
			}
			if !strings.Contains(r.Summary, "locked") {
				t.Errorf("summary should contain 'locked', got %q", r.Summary)
			}
			if !strings.Contains(r.Summary, "Alpha") {
				t.Errorf("summary should contain 'Alpha', got %q", r.Summary)
			}
		}
	}
	if !found {
		t.Error("ExperimentalCache resource not found")
	}
}

func TestUpgradeSelector_FindingReferencedFunctions(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "handleRequest",
		File: "handler.go", Line: 10, EndLine: 50,
		Language: "go",
	})

	findings := []query.Finding{
		{RuleID: "CGA-U02", Severity: "low", Domain: "upgrade", Message: "Pre-release API",
			File: "handler.go", Line: 25},
	}

	dir := t.TempDir()
	writeTestFile(t, dir, "handler.go", 60)

	sel := NewUpgradeSelector(dir)
	layer, _ := sel.Select(cpg, nil, findings, nil)

	found := false
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "handleRequest" {
				found = true
			}
		}
	}
	if !found {
		t.Error("handleRequest should be included (finding at line 25 is within 10-50)")
	}
}

func TestUpgradeSelector_CallRelationships(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "ConvertTo",
		File: "convert.go", Line: 10,
		Annotations: map[string]bool{annotVersionConversion: true},
	})
	cpg.AddNode(&graph.Node{
		ID: "fn2", Kind: graph.NodeFunction, Name: "helper",
		File: "helper.go", Line: 5,
	})
	cpg.AddEdge(&graph.Edge{
		From: "fn1", To: "fn2", Kind: graph.EdgeCalls,
	})

	sel := NewUpgradeSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	if len(layer.Relationships) == 0 {
		t.Fatal("expected at least one relationship")
	}

	rel := layer.Relationships[0]
	if rel.Kind != "calls" {
		t.Errorf("relationship kind = %q, want calls", rel.Kind)
	}
	if rel.From.Function != "ConvertTo" {
		t.Errorf("from = %q, want ConvertTo", rel.From.Function)
	}
	if rel.To.Resolved == nil || *rel.To.Resolved != false {
		t.Error("expected To.Resolved = false for non-selected function")
	}
}

func TestUpgradeSelector_EmptyInputs(t *testing.T) {
	cpg := graph.NewCPG()
	sel := NewUpgradeSelector(t.TempDir())
	layer, warnings := sel.Select(cpg, nil, nil, nil)

	if layer.Name != "upgrade" {
		t.Errorf("layer name = %q, want %q", layer.Name, "upgrade")
	}
	if len(layer.Files) != 0 {
		t.Errorf("expected 0 files, got %d", len(layer.Files))
	}
	if len(layer.Findings) != 0 {
		t.Errorf("expected 0 findings, got %d", len(layer.Findings))
	}
	if len(layer.Resources) != 0 {
		t.Errorf("expected 0 resources, got %d", len(layer.Resources))
	}
	if len(layer.Relationships) != 0 {
		t.Errorf("expected 0 relationships, got %d", len(layer.Relationships))
	}
	if len(warnings) != 0 {
		t.Errorf("expected 0 warnings, got %d", len(warnings))
	}
}

func TestUpgradeSelector_PathTraversalRejected(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "ConvertTo",
		File: "../../../etc/passwd", Line: 1, EndLine: 5,
		Language:    "go",
		Annotations: map[string]bool{annotVersionConversion: true},
	})

	sel := NewUpgradeSelector(t.TempDir())
	layer, warnings := sel.Select(cpg, nil, nil, nil)

	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Code != "" {
				t.Error("path traversal should not extract code")
			}
		}
	}

	foundWarning := false
	for _, w := range warnings {
		if strings.Contains(w.Message, "path traversal") {
			foundWarning = true
		}
	}
	if !foundWarning {
		t.Error("expected path traversal warning")
	}
}

func TestUpgradeSelector_ExcludesNonUpgradeFunctions(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "RegularFunc",
		File: "regular.go", Line: 5,
		Language: "go",
	})

	sel := NewUpgradeSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	if len(layer.Files) != 0 {
		t.Errorf("expected no files (no upgrade functions), got %d", len(layer.Files))
	}
}

func TestUpgradeSelector_BodyExtractionFailure(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "MigrateOld",
		File: "short.go", Line: 50, EndLine: 60,
		Language:    "go",
		Annotations: map[string]bool{annotMigration: true},
	})

	dir := t.TempDir()
	writeTestFile(t, dir, "short.go", 5)

	sel := NewUpgradeSelector(dir)
	layer, warnings := sel.Select(cpg, nil, nil, nil)

	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "MigrateOld" && fn.Code != "" {
				t.Error("should not have code body when lines exceed file length")
			}
		}
	}

	foundWarning := false
	for _, w := range warnings {
		if strings.Contains(w.Message, "body extraction failed") {
			foundWarning = true
		}
	}
	if !foundWarning {
		t.Error("expected body extraction failure warning")
	}
}

func TestUpgradeSelector_MultipleMetas(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "MigrateAndConvert",
		File: "dual.go", Line: 5,
		Language: "go",
		Annotations: map[string]bool{
			annotVersionConversion: true,
			annotMigration:         true,
		},
	})

	sel := NewUpgradeSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	metaKeys := make(map[string]bool)
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "MigrateAndConvert" {
				for _, m := range fn.Metas {
					metaKeys[m.Key] = true
				}
			}
		}
	}

	if !metaKeys["conversion"] {
		t.Error("expected conversion meta")
	}
	if !metaKeys["migration"] {
		t.Error("expected migration meta")
	}
}

func TestUpgradeSelector_NoDuplicateFindingReferencedFunctions(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "ConvertTo",
		File: "convert.go", Line: 10, EndLine: 30,
		Language:    "go",
		Annotations: map[string]bool{annotVersionConversion: true},
	})

	findings := []query.Finding{
		{RuleID: "CGA-U01", Severity: "medium", Domain: "upgrade", Message: "issue",
			File: "convert.go", Line: 15},
	}

	dir := t.TempDir()
	writeTestFile(t, dir, "convert.go", 40)

	sel := NewUpgradeSelector(dir)
	layer, _ := sel.Select(cpg, nil, findings, nil)

	count := 0
	for _, f := range layer.Files {
		for _, fn := range f.Functions {
			if fn.Name == "ConvertTo" {
				count++
			}
		}
	}
	if count != 1 {
		t.Errorf("ConvertTo should appear exactly once, got %d", count)
	}
}

func TestUpgradeSelector_RelationshipCap(t *testing.T) {
	cpg := graph.NewCPG()

	for i := 0; i < 200; i++ {
		fnID := fmt.Sprintf("fn%d", i)
		helperID := fmt.Sprintf("helper%d", i)

		cpg.AddNode(&graph.Node{
			ID: fnID, Kind: graph.NodeFunction, Name: fmt.Sprintf("Convert%d", i),
			File: fmt.Sprintf("convert_%d.go", i), Line: 1,
			Language:    "go",
			Annotations: map[string]bool{annotVersionConversion: true},
		})
		cpg.AddNode(&graph.Node{
			ID: helperID, Kind: graph.NodeFunction, Name: fmt.Sprintf("helper%d", i),
			File: fmt.Sprintf("helper_%d.go", i), Line: 1,
			Language: "go",
		})
		cpg.AddEdge(&graph.Edge{From: fnID, To: helperID, Kind: graph.EdgeCalls})
	}

	sel := NewUpgradeSelector(t.TempDir())
	layer, _ := sel.Select(cpg, nil, nil, nil)

	if len(layer.Relationships) > maxRelationships {
		t.Errorf("relationships = %d, should be capped at %d", len(layer.Relationships), maxRelationships)
	}
}

func TestUpgradeSelector_SharedFile(t *testing.T) {
	cpg := graph.NewCPG()
	cpg.AddNode(&graph.Node{
		ID: "fn1", Kind: graph.NodeFunction, Name: "ConvertTo",
		File: "api/convert.go", Line: 10, EndLine: 30,
		Language:    "go",
		Annotations: map[string]bool{annotVersionConversion: true},
	})
	cpg.AddNode(&graph.Node{
		ID: "fn2", Kind: graph.NodeFunction, Name: "Hub",
		File: "api/convert.go", Line: 35, EndLine: 40,
		Language:    "go",
		Annotations: map[string]bool{annotVersionConversion: true},
	})

	dir := t.TempDir()
	writeTestFileWithDir(t, dir, "api/convert.go", 50)

	sel := NewUpgradeSelector(dir)
	layer, _ := sel.Select(cpg, nil, nil, nil)

	fileCount := 0
	for _, f := range layer.Files {
		if f.Path == "api/convert.go" {
			fileCount++
			if len(f.Functions) != 2 {
				t.Errorf("expected 2 functions in shared file, got %d", len(f.Functions))
			}
		}
	}
	if fileCount != 1 {
		t.Errorf("expected 1 file entry for api/convert.go, got %d", fileCount)
	}
}
