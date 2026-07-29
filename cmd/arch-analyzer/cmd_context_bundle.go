package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ugiordan/architecture-analyzer/pkg/extractor"
	"github.com/ugiordan/architecture-analyzer/pkg/srclang"
	"github.com/ugiordan/architecture-analyzer/pkg/srclang/compile"
)

func cmdContextBundle(args []string) error {
	fs := flag.NewFlagSet("context-bundle", flag.ExitOnError)
	layer := fs.String("layer", "security", "Domain layer (security, architecture, testing, upgrade, netpolicy, codegen)")
	output := fs.String("output", "context.srclg", "Output .srclg file")
	platformFile := fs.String("platform", "", "Path to platform-architecture.json for platform context")
	withScan := fs.Bool("with-scan", true, "Run CPG scan for findings and taint analysis")
	domainList := fs.String("domains", "", "Comma-separated domain list for scan (default: all)")
	importSARIF := fs.String("sarif", "", "Comma-separated SARIF files to include as findings (e.g., kube-chainsaw.sarif,tekton-guard.sarif)")
	fs.Parse(reorderArgs(fs, args))

	if fs.NArg() < 1 {
		return fmt.Errorf("usage: arch-analyzer context-bundle [flags] <repo-path>")
	}
	repoPath := fs.Arg(0)

	arch, err := extractor.ExtractAll(repoPath, nil)
	if err != nil {
		return fmt.Errorf("extraction failed: %w", err)
	}

	cpg, err := buildCPG(repoPath)
	if err != nil {
		return fmt.Errorf("CPG build failed: %w", err)
	}

	opts := compile.Options{
		RepoPath:            repoPath,
		Layer:               *layer,
		Arch:                arch,
		CPG:                 cpg,
		SecurityAnnotations: arch.SecurityAnnotations,
		PlatformFile:        *platformFile,
	}

	if *withScan {
		archData := prepareArchData(repoPath, cpg, "")
		findings, scanErr := runSecurityScan(cpg, *domainList, archData)
		if scanErr != nil {
			return fmt.Errorf("security scan failed: %w", scanErr)
		}
		opts.Findings = findings
	}

	if *importSARIF != "" {
		externalFindings, sarifErr := ingestSARIFFiles(cpg, *importSARIF)
		if sarifErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: SARIF ingestion: %v\n", sarifErr)
		}
		opts.Findings = append(opts.Findings, externalFindings...)
	}

	doc, err := compile.Compile(opts)
	if err != nil {
		return fmt.Errorf("compile failed: %w", err)
	}

	// Check if the document needs splitting
	var buf bytes.Buffer
	if err := srclang.WriteDocument(&buf, doc); err != nil {
		return fmt.Errorf("serializing document: %w", err)
	}

	if buf.Len() <= compile.BundleThreshold {
		// Small enough for single file
		if err := os.WriteFile(*output, buf.Bytes(), 0o644); err != nil {
			return fmt.Errorf("writing output: %w", err)
		}
		fmt.Printf("SrcLang document written to %s (%s layer, %dKB)\n", *output, *layer, buf.Len()/1024)
		return nil
	}

	// Split into directory bundle
	bundle := compile.SplitBundle(doc)
	bundleDir := *output + ".d"
	if err := os.MkdirAll(filepath.Join(bundleDir, "files"), 0o755); err != nil {
		return fmt.Errorf("creating bundle directory: %w", err)
	}

	// Write index (compact: no code bodies, no finding descriptions)
	indexPath := filepath.Join(bundleDir, "index.srclg")
	indexFile, err := os.Create(indexPath)
	if err != nil {
		return fmt.Errorf("creating index: %w", err)
	}
	if err := srclang.WriteIndexDocument(indexFile, bundle.IndexDoc); err != nil {
		indexFile.Close()
		return fmt.Errorf("writing index: %w", err)
	}
	indexFile.Close()
	indexInfo, _ := os.Stat(indexPath)

	// Write shards
	for shardPath, shardDoc := range bundle.Shards {
		fullPath := filepath.Join(bundleDir, shardPath)
		if dir := filepath.Dir(fullPath); dir != bundleDir {
			os.MkdirAll(dir, 0o755)
		}
		sf, err := os.Create(fullPath)
		if err != nil {
			return fmt.Errorf("creating shard %s: %w", shardPath, err)
		}
		if err := srclang.WriteDocument(sf, shardDoc); err != nil {
			sf.Close()
			return fmt.Errorf("writing shard %s: %w", shardPath, err)
		}
		sf.Close()
	}

	fmt.Printf("SrcLang bundle written to %s/ (%s layer, index %dKB, %d shards)\n",
		bundleDir, *layer, indexInfo.Size()/1024, len(bundle.Shards))

	// Also write single-file version (with budget caps applied) for backward compat
	if err := os.WriteFile(*output, buf.Bytes(), 0o644); err != nil {
		return fmt.Errorf("writing single-file output: %w", err)
	}
	fmt.Printf("SrcLang single-file also written to %s (%dKB, capped)\n", *output, buf.Len()/1024)

	return nil
}
