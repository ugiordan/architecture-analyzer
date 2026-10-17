# eval-hub: Cache Architecture

Controller-runtime cache configuration controls which Kubernetes resources are cached in-memory. Misconfigured caches (cluster-wide watches on high-cardinality types without filters) are a primary cause of operator OOM kills.

## Cache Architecture

### Manager Configuration

| Property | Value |
|----------|-------|
| Manager file | `cmd/eval_runtime_init/main.go` |
| Cache scope | cluster-wide |
| DefaultTransform | no |

