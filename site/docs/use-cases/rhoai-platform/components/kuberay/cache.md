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

### Issues

- No GOMEMLIMIT set in deployment (Go GC cannot pressure-tune). Set GOMEMLIMIT to 80-90% of container memory limit for optimal GC behavior
- No cache configuration: all informers are cluster-wide (OOM risk). See https://book.kubebuilder.io/reference/watching-resources/filtering for cache filtering patterns
- Type Certificate is watched but has no cache filter (cluster-wide informer)
- Type Issuer is watched but has no cache filter (cluster-wide informer)
- Type Job is watched but has no cache filter (cluster-wide informer)
- Type NetworkPolicy is watched but has no cache filter (cluster-wide informer)
- Type PersistentVolumeClaim is watched but has no cache filter (cluster-wide informer)
- Type Pod is watched but has no cache filter (cluster-wide informer)
- Type RayCluster is watched but has no cache filter (cluster-wide informer)
- Type RayCronJob is watched but has no cache filter (cluster-wide informer)
- Type RayJob is watched but has no cache filter (cluster-wide informer)
- Type RayService is watched but has no cache filter (cluster-wide informer)
- Type Route is watched but has no cache filter (cluster-wide informer)
- Type Secret is watched but has no cache filter (cluster-wide informer)
- Type Service is watched but has no cache filter (cluster-wide informer)
- Type ServiceAccount is watched but has no cache filter (cluster-wide informer)

