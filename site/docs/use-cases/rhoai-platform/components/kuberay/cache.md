# kuberay: Cache Architecture

Controller-runtime cache configuration controls which Kubernetes resources are cached in-memory. Misconfigured caches (cluster-wide watches on high-cardinality types without filters) are a primary cause of operator OOM kills.

## Cache Architecture

### Manager Configuration

| Property | Value |
|----------|-------|
| Manager file | `ray-operator/main.go` |
| Cache scope | namespace-scoped |
| DefaultTransform | no |
| Memory limit | 512Mi |

### Filtered Types

| Type | Filter Kind | Filter |
|------|-------------|--------|
| batchv1.Job | label | label selector |

### Issues

- No DefaultTransform: managedFields cached for all objects (wasted memory). Add cache.DefaultTransform to strip managedFields and reduce memory footprint
- No GOMEMLIMIT set in deployment (Go GC cannot pressure-tune). Set GOMEMLIMIT to 80-90% of container memory limit for optimal GC behavior
- Type NetworkPolicy is watched but has no cache filter (cluster-wide informer)
- Type Pod is watched but has no cache filter (cluster-wide informer)
- Type RayCluster is watched but has no cache filter (cluster-wide informer)
- Type RayCronJob is watched but has no cache filter (cluster-wide informer)
- Type RayJob is watched but has no cache filter (cluster-wide informer)
- Type RayService is watched but has no cache filter (cluster-wide informer)
- Type Route is watched but has no cache filter (cluster-wide informer)
- Type Secret is watched but has no cache filter (cluster-wide informer)
- Type Service is watched but has no cache filter (cluster-wide informer)
- Type ServiceAccount is watched but has no cache filter (cluster-wide informer)

