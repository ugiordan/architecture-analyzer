# Installation

## Download pre-built binary (recommended)

Download the latest release from [GitHub Releases](https://github.com/ugiordan/architecture-analyzer/releases):

```bash
# Linux (amd64)
curl -L https://github.com/ugiordan/architecture-analyzer/releases/latest/download/arch-analyzer-linux-amd64 -o arch-analyzer
chmod +x arch-analyzer

# macOS (Apple Silicon)
curl -L https://github.com/ugiordan/architecture-analyzer/releases/latest/download/arch-analyzer-darwin-arm64 -o arch-analyzer
chmod +x arch-analyzer
```

## Build from source

Requires Go 1.25+ and git.

```bash
git clone https://github.com/ugiordan/architecture-analyzer.git
cd architecture-analyzer
go build -o arch-analyzer ./cmd/arch-analyzer/
```

## Verify installation

```bash
./arch-analyzer version
```

Expected output:

```
arch-analyzer 0.1.0
```

## Optional: Add to PATH

```bash
# Move to a directory in your PATH
sudo mv arch-analyzer /usr/local/bin/

# Or add the project directory to PATH
export PATH="$PATH:$(pwd)"
```

## Tree-sitter dependency

The code property graph and security scanning features use tree-sitter for Go source parsing. Tree-sitter is included as a Go dependency and compiled automatically during `go build`. No additional installation needed.

## GitHub Actions

For CI/CD usage, the analyzer builds in the workflow and runs against repos. See [CI Integration](../guides/ci-integration.md) for workflow examples.
