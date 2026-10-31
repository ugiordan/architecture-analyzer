# data-science-pipelines: Cache Architecture

Controller-runtime cache configuration controls which Kubernetes resources are cached in-memory. Misconfigured caches (cluster-wide watches on high-cardinality types without filters) are a primary cause of operator OOM kills.

## Cache Architecture

### Manager Configuration

| Property | Value |
|----------|-------|
| Manager file | `backend/src/v2/cmd/launcher-v2/main.go` |
| Cache scope | namespace-scoped |
| DefaultTransform | no |

