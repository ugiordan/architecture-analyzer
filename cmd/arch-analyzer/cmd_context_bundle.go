package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ugiordan/architecture-analyzer/pkg/extractor"
	"github.com/ugiordan/architecture-analyzer/pkg/srclang"
	"github.com/ugiordan/architecture-analyzer/pkg/srclang/compile"
)

func cmdContextBundle(args []string) error {
	fs := flag.NewFlagSet("context-bundle", flag.ExitOnError)
	layer := fs.String("layer", "security", "Domain layer (security)")
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

	// Extract architecture data
	arch, err := extractor.ExtractAll(repoPath, nil)
	if err != nil {
		return fmt.Errorf("extraction failed: %w", err)
	}

	// Build CPG for function extraction (needed by SecuritySelector)
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

	// Optionally run security scan for findings and taint analysis
	if *withScan {
		archData := prepareArchData(repoPath, cpg, "")
		findings, scanErr := runSecurityScan(cpg, *domainList, archData)
		if scanErr != nil {
			return fmt.Errorf("security scan failed: %w", scanErr)
		}
		opts.Findings = findings
	}

	// Ingest external SARIF findings (kube-chainsaw, tekton-guard, helm-guard, etc.)
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

	f, err := os.Create(*output)
	if err != nil {
		return fmt.Errorf("creating output file: %w", err)
	}
	defer f.Close()

	if err := srclang.WriteDocument(f, doc); err != nil {
		return fmt.Errorf("writing srclang document: %w", err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("closing output file: %w", err)
	}

	fmt.Printf("SrcLang document written to %s (%s layer)\n", *output, *layer)
	return nil
}
