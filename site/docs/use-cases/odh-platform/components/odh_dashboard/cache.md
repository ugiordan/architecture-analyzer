# odh-dashboard: Cache Architecture

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
- No cache configuration: all informers are cluster-wide (OOM risk). See https://book.kubebuilder.io/reference/watching-resources/filtering for cache filtering patterns
- Type ConfigMap is watched but has no cache filter (cluster-wide informer)
- Type Dashboard is watched but has no cache filter (cluster-wide informer)
- Type Deployment is watched but has no cache filter (cluster-wide informer)
- Type PodDisruptionBudget is watched but has no cache filter (cluster-wide informer)
- Type Service is watched but has no cache filter (cluster-wide informer)
- Type StatefulSet is watched but has no cache filter (cluster-wide informer)
- Type VirtualService is watched but has no cache filter (cluster-wide informer)
- Type Workspace is watched but has no cache filter (cluster-wide informer)
- Type WorkspaceKind is watched but has no cache filter (cluster-wide informer)

