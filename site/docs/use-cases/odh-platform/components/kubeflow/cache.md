# kubeflow: Cache Architecture

Controller-runtime cache configuration controls which Kubernetes resources are cached in-memory. Misconfigured caches (cluster-wide watches on high-cardinality types without filters) are a primary cause of operator OOM kills.

## Cache Architecture

### Manager Configuration

| Property | Value |
|----------|-------|
| Manager file | `components/odh-notebook-controller/main.go` |
| Cache scope | cluster-wide |
| DefaultTransform | no |
| Memory limit | 4Gi |

### Issues

- No GOMEMLIMIT set in deployment (Go GC cannot pressure-tune). Set GOMEMLIMIT to 80-90% of container memory limit for optimal GC behavior
- No cache configuration: all informers are cluster-wide (OOM risk). See https://book.kubebuilder.io/reference/watching-resources/filtering for cache filtering patterns
- Type ConfigMap is watched but has no cache filter (cluster-wide informer)
- Type HTTPRoute is watched but has no cache filter (cluster-wide informer)
- Type Namespace is watched but has no cache filter (cluster-wide informer)
- Type NetworkPolicy is watched but has no cache filter (cluster-wide informer)
- Type Notebook is watched but has no cache filter (cluster-wide informer)
- Type ReferenceGrant is watched but has no cache filter (cluster-wide informer)
- Type RoleBinding is watched but has no cache filter (cluster-wide informer)
- Type Secret is watched but has no cache filter (cluster-wide informer)
- Type Service is watched but has no cache filter (cluster-wide informer)
- Type ServiceAccount is watched but has no cache filter (cluster-wide informer)
- Type StatefulSet is watched but has no cache filter (cluster-wide informer)

