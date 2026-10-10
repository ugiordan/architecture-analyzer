# trainer: Cache Architecture

Controller-runtime cache configuration controls which Kubernetes resources are cached in-memory. Misconfigured caches (cluster-wide watches on high-cardinality types without filters) are a primary cause of operator OOM kills.

## Cache Architecture

### Manager Configuration

| Property | Value |
|----------|-------|
| Manager file | `cmd/trainer-controller-manager/main.go` |
| Cache scope | cluster-wide |
| DefaultTransform | no |

### Implicit Informers (OOM Risk)

| Type | Source | Risk |
|------|--------|------|
| corev1.Secret | `pkg/runtime/framework/plugins/mpi/mpi.go:264` | **HIGH** |

### Issues

- Implicit informer for corev1.Secret via client.Get at pkg/runtime/framework/plugins/mpi/mpi.go:264 (cluster-wide, OOM risk). This bypasses cache filters and creates a full cluster-wide watch
- No cache configuration: all informers are cluster-wide (OOM risk). See https://book.kubebuilder.io/reference/watching-resources/filtering for cache filtering patterns
- Type Deployment is watched but has no cache filter (cluster-wide informer)
- Type OptimizationJob is watched but has no cache filter (cluster-wide informer)
- Type Service is watched but has no cache filter (cluster-wide informer)
- Type TrainJob is watched but has no cache filter (cluster-wide informer)

