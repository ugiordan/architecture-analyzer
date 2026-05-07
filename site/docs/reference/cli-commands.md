# CLI Commands

## Architecture Commands

### arch-analyzer analyze

Extract architecture data and render diagrams.

```bash
arch-analyzer analyze <repo-path> --output-dir <dir>
```

| Flag | Description |
|------|-------------|
| `--output-dir` | Directory for all output files |

**Output**: `component-architecture.json` + all diagram/report files.

### arch-analyzer extract

Extract architecture data only (no rendering).

```bash
arch-analyzer extract <repo-path> --output <file>
```

| Flag | Description |
|------|-------------|
| `--output` | Output JSON file path |

### arch-analyzer render

Render diagrams from existing JSON.

```bash
arch-analyzer render <json-file> --output-dir <dir> [--formats <list>]
```

| Flag | Description |
|------|-------------|
| `--output-dir` | Directory for diagram files |
| `--formats` | Comma-separated list: `rbac`, `component`, `security`, `dependencies`, `c4`, `dataflow`, `report` |

### arch-analyzer docs

Generate browsable documentation pages from architecture JSON.

```bash
arch-analyzer docs --output-dir <dir> <json-file>
```

| Flag | Description |
|------|-------------|
| `--output-dir` | Directory for generated markdown pages (default: `docs`) |
| `--prefix` | Path prefix for the nav snippet output |

Auto-detects whether the input is a single component or aggregated platform JSON. For platform data, generates per-component deep-dive pages under `components/`.

## Code Graph Commands

### arch-analyzer scan

Build code property graph and run security queries.

```bash
arch-analyzer scan <repo-path> [flags]
```

| Flag | Description |
|------|-------------|
| `--format` | Output format: `json` or `sarif` |
| `--output` | Output file path |
| `--domains` | Comma-separated domain list: `security`, `testing`, `upgrade` |
| `--with-arch` | Enable architecture enrichment (loads or generates architecture data) |
| `--import-sarif` | Comma-separated SARIF files to ingest alongside the scan |

### arch-analyzer graph

Export the code property graph as JSON or DOT.

```bash
arch-analyzer graph <repo-path> [flags]
```

| Flag | Description |
|------|-------------|
| `--output` | Output file path |
| `--format` | Output format: `json` (default) or `dot` |

Exports the raw CPG (nodes, edges, basic blocks) for inspection or custom analysis. JSON output includes `schema_version: 2`.

### arch-analyzer diff

Structural diff between two code-graph.json files.

```bash
arch-analyzer diff <base.json> <head.json> [flags]
```

| Flag | Description |
|------|-------------|
| `--format` | Output format: `json` or `text` |
| `--kind` | Filter by finding kind (comma-separated) |
| `--output` | Output file path |

Detects new/removed functions, changed complexity, new call edges, trust level changes, and new annotations. Useful for PR review automation.

### arch-analyzer ingest

Ingest external scanner SARIF findings and map them to CPG nodes.

```bash
arch-analyzer ingest <sarif-file> [flags]
```

| Flag | Description |
|------|-------------|
| `--graph` | Existing code-graph.json to enrich |
| `--output` | Output file path for enriched graph |

Maps SARIF findings to the tightest-fitting CPG node at each location, adding `sarif:<tool>:<rule_id>` annotations.

### arch-analyzer domains

List all registered analysis domains.

```bash
arch-analyzer domains
```

Output includes domain name, supported languages, dependencies, and query count.

## Contract Validation Commands

### arch-analyzer extract-schema

Extract CRD JSON schemas for contract validation.

```bash
arch-analyzer extract-schema <repo-path> --output-dir <dir>
```

| Flag | Description |
|------|-------------|
| `--output-dir` | Directory for schema JSON files |

### arch-analyzer validate

Validate CRD changes against baseline contracts.

```bash
arch-analyzer validate <repo-path> --contracts-dir <dir>
```

| Flag | Description |
|------|-------------|
| `--contracts-dir` | Directory containing baseline schemas |

Exit code 1 if breaking changes detected.

## Platform Commands

### arch-analyzer aggregate

Merge multiple component analyses into platform view.

```bash
arch-analyzer aggregate <results-dir> --output-dir <dir>
```

| Flag | Description |
|------|-------------|
| `--output-dir` | Directory for platform output |

Discovers all `component-architecture.json` files in the results directory recursively.

### arch-analyzer aggregate-cpg

Merge code graphs into a platform-wide CPG.

```bash
arch-analyzer aggregate-cpg <results-dir>
```

### arch-analyzer discover

Discover platform components from kustomize manifests.

```bash
arch-analyzer discover <operator-repo-path> [flags]
```

| Flag | Description |
|------|-------------|
| `--output` | Output file path |
| `--format` | Output format: `json`, `text`, or `map` |
| `--org` | GitHub organization |
| `--platform` | Platform name |

### arch-analyzer build-config

Extract build metadata (OCP versions, architectures, OLM).

```bash
arch-analyzer build-config <dir>
```

### arch-analyzer konflux

Parse Konflux snapshot image mappings.

```bash
arch-analyzer konflux <snapshot-file-or-dir>
```

### arch-analyzer platforms

List platforms defined in scan config.

```bash
arch-analyzer platforms <scan-config.yaml> [flags]
```

| Flag | Description |
|------|-------------|
| `--platform` | Filter by platform name |
| `--output` | Output file path |

### arch-analyzer version-compat

Check API version compatibility against target OCP/k8s version.

```bash
arch-analyzer version-compat <arch.json> [flags]
```

| Flag | Description |
|------|-------------|
| `--target-version` | Target OCP or Kubernetes version |

## Combined Commands

### arch-analyzer full-analysis

Run architecture extraction + code graph scan.

```bash
arch-analyzer full-analysis <repo-path> --output-dir <dir> [flags]
```

| Flag | Description |
|------|-------------|
| `--output-dir` | Directory for all output files |
| `--domains` | Comma-separated domain list |
| `--import-sarif` | Comma-separated SARIF files to ingest |

### arch-analyzer version

Print version information.

```bash
arch-analyzer version
```
