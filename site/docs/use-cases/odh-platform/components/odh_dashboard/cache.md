# odh-dashboard: Cache Architecture

[Overview](index.md) · [Agent Context](agent-context.md) · [Extracted Data](details.md) · [Network](network.md) · [RBAC](rbac.md) · [Security](security.md) · [Cache](cache.md) · [Dataflow](dataflow.md)

Controller-runtime cache configuration controls which Kubernetes resources are cached in-memory. Misconfigured caches (cluster-wide watches on high-cardinality types without filters) are a primary cause of operator OOM kills.

## Cache Architecture

### Manager Configuration

| Property | Value |
|----------|-------|
| Manager file | `packages/notebooks/upstream/workspaces/controller/cmd/main.go` |
| Cache scope | cluster-wide |
| DefaultTransform | no |
| Memory limit | 256Mi |

### Issues

- No GOMEMLIMIT set in deployment (Go GC cannot pressure-tune). Set GOMEMLIMIT to 80-90% of container memory limit for optimal GC behavior

