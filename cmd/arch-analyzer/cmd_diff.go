package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ugiordan/architecture-analyzer/pkg/diff"
	"github.com/ugiordan/architecture-analyzer/pkg/graph"
)

func cmdDiff(args []string) error {
	fs := flag.NewFlagSet("diff", flag.ExitOnError)
	outputFile := fs.String("output", "", "Output file (default: stdout)")
	format := fs.String("format", "json", "Output format: json, text")
	kindFilter := fs.String("kind", "", "Filter by node kind (comma-separated, e.g. Function,HTTPEndpoint)")
	fs.Parse(args)

	if fs.NArg() < 2 {
		return fmt.Errorf("usage: arch-analyzer diff <base.json> <head.json> [--format json|text] [--kind kinds] [--output file]")
	}

	basePath := fs.Arg(0)
	headPath := fs.Arg(1)

	base, err := loadSnapshot(basePath)
	if err != nil {
		return fmt.Errorf("loading base: %w", err)
	}
	head, err := loadSnapshot(headPath)
	if err != nil {
		return fmt.Errorf("loading head: %w", err)
	}

	result, err := diff.Compare(base, head)
	if err != nil {
		return err
	}

	// Apply kind filter if specified
	if *kindFilter != "" {
		kinds := make(map[string]bool)
		for _, k := range strings.Split(*kindFilter, ",") {
			kinds[strings.TrimSpace(k)] = true
		}
		filterByKind(result, kinds)
	}

	var content []byte
	switch *format {
	case "json":
		content, err = json.MarshalIndent(result, "", "  ")
		if err != nil {
			return err
		}
		content = append(content, '\n')
	case "text":
		content = []byte(formatDiffText(result))
	default:
		return fmt.Errorf("unknown format: %s", *format)
	}

	if *outputFile != "" {
		if wErr := os.WriteFile(*outputFile, content, 0o644); wErr != nil {
			return wErr
		}
	} else {
		os.Stdout.Write(content)
	}

	if result.HasDifferences() {
		return errDiffFound
	}
	return nil
}

func loadSnapshot(path string) (diff.GraphSnapshot, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return diff.GraphSnapshot{}, err
	}
	var snap diff.GraphSnapshot
	if err := json.Unmarshal(data, &snap); err != nil {
		return diff.GraphSnapshot{}, fmt.Errorf("parsing %s: %w", path, err)
	}
	return snap, nil
}

func filterByKind(d *diff.GraphDiff, kinds map[string]bool) {
	d.Nodes.Added = filterNodes(d.Nodes.Added, kinds)
	d.Nodes.Removed = filterNodes(d.Nodes.Removed, kinds)

	filtered := make([]diff.NodeChange, 0)
	for _, mc := range d.Nodes.Modified {
		if kinds[string(mc.After.Kind)] {
			filtered = append(filtered, mc)
		}
	}
	d.Nodes.Modified = filtered

	// Filter edges to only include those connecting filtered nodes.
	nodeIDs := make(map[string]bool)
	for _, n := range d.Nodes.Added {
		nodeIDs[n.ID] = true
	}
	for _, n := range d.Nodes.Removed {
		nodeIDs[n.ID] = true
	}
	for _, mc := range d.Nodes.Modified {
		nodeIDs[mc.After.ID] = true
	}
	d.Edges.Added = filterEdgesForNodes(d.Edges.Added, nodeIDs)
	d.Edges.Removed = filterEdgesForNodes(d.Edges.Removed, nodeIDs)

	// Recompute summary
	d.Summary.NodesAdded = len(d.Nodes.Added)
	d.Summary.NodesRemoved = len(d.Nodes.Removed)
	d.Summary.NodesModified = len(d.Nodes.Modified)
	d.Summary.EdgesAdded = len(d.Edges.Added)
	d.Summary.EdgesRemoved = len(d.Edges.Removed)
	d.Summary.ByKind = make(map[string]diff.KindCounts)
	d.Summary.ByLanguage = make(map[string]diff.KindCounts)
	for _, n := range d.Nodes.Added {
		incFilterCount(d.Summary.ByKind, string(n.Kind), "added")
		incFilterCount(d.Summary.ByLanguage, n.Language, "added")
	}
	for _, n := range d.Nodes.Removed {
		incFilterCount(d.Summary.ByKind, string(n.Kind), "removed")
		incFilterCount(d.Summary.ByLanguage, n.Language, "removed")
	}
	for _, mc := range d.Nodes.Modified {
		incFilterCount(d.Summary.ByKind, string(mc.After.Kind), "modified")
		incFilterCount(d.Summary.ByLanguage, mc.After.Language, "modified")
	}
}

func filterEdgesForNodes(edges []graph.Edge, nodeIDs map[string]bool) []graph.Edge {
	var result []graph.Edge
	for _, e := range edges {
		if nodeIDs[e.From] || nodeIDs[e.To] {
			result = append(result, e)
		}
	}
	return result
}

func incFilterCount(m map[string]diff.KindCounts, key, op string) {
	if key == "" {
		return
	}
	c := m[key]
	switch op {
	case "added":
		c.Added++
	case "removed":
		c.Removed++
	case "modified":
		c.Modified++
	}
	m[key] = c
}

func filterNodes(nodes []graph.Node, kinds map[string]bool) []graph.Node {
	filtered := make([]graph.Node, 0)
	for _, n := range nodes {
		if kinds[string(n.Kind)] {
			filtered = append(filtered, n)
		}
	}
	return filtered
}

func formatDiffText(d *diff.GraphDiff) string {
	var b strings.Builder

	if d.BaseVersion != "" || d.HeadVersion != "" {
		fmt.Fprintf(&b, "Graph Diff: %s -> %s\n\n", d.BaseVersion, d.HeadVersion)
	}

	fmt.Fprintf(&b, "Nodes: +%d added, -%d removed, ~%d modified\n",
		d.Summary.NodesAdded, d.Summary.NodesRemoved, d.Summary.NodesModified)
	fmt.Fprintf(&b, "Edges: +%d added, -%d removed\n",
		d.Summary.EdgesAdded, d.Summary.EdgesRemoved)

	if len(d.Nodes.Modified) > 0 {
		b.WriteString("\nModified:\n")
		for _, mc := range d.Nodes.Modified {
			fmt.Fprintf(&b, "  %s:%d %s %s\n", mc.After.File, mc.After.Line, mc.After.Kind, mc.After.Name)
			for _, fc := range mc.Changes {
				fmt.Fprintf(&b, "    %s: %v -> %v\n", fc.Field, fc.OldValue, fc.NewValue)
			}
		}
	}

	if len(d.Nodes.Added) > 0 {
		b.WriteString("\nAdded:\n")
		for _, n := range d.Nodes.Added {
			fmt.Fprintf(&b, "  %s:%d %s %s\n", n.File, n.Line, n.Kind, n.Name)
		}
	}

	if len(d.Nodes.Removed) > 0 {
		b.WriteString("\nRemoved:\n")
		for _, n := range d.Nodes.Removed {
			fmt.Fprintf(&b, "  %s:%d %s %s\n", n.File, n.Line, n.Kind, n.Name)
		}
	}

	return b.String()
}

// cmdDiscover discovers platform components from an operator repo's kustomize manifests.
