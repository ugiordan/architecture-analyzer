# kubeflow: Cache Architecture

Controller-runtime cache configuration controls which Kubernetes resources are cached in-memory. Misconfigured caches (cluster-wide watches on high-cardinality types without filters) are a primary cause of operator OOM kills.

## Cache Architecture

### Manager Configuration

| Property | Value |
|----------|-------|
| Manager file | `components/odh-notebook-controller/main.go` |
| Cache scope | cluster-wide |
| DefaultTransform | yes |
| Memory limit | 4Gi |

### Cache-Bypassed Types (DisableFor)

- corev1.ConfigMap
- corev1.Secret

### Issues

- Cache bypass (DisableFor) configured for corev1.ConfigMap. This is a common fix for OOM caused by informer cache flooding from high-cardinality types (e.g., opendatahub-io/model-registry-operator#457)
- Cache bypass (DisableFor) configured for corev1.Secret. This is a common fix for OOM caused by informer cache flooding from high-cardinality types (e.g., opendatahub-io/model-registry-operator#457)
- No GOMEMLIMIT set in deployment (Go GC cannot pressure-tune). Set GOMEMLIMIT to 80-90% of container memory limit for optimal GC behavior
- Type HTTPRoute is watched but has no cache filter (cluster-wide informer)
- Type NetworkPolicy is watched but has no cache filter (cluster-wide informer)
- Type Notebook is watched but has no cache filter (cluster-wide informer)
- Type ReferenceGrant is watched but has no cache filter (cluster-wide informer)
- Type RoleBinding is watched but has no cache filter (cluster-wide informer)
- Type Service is watched but has no cache filter (cluster-wide informer)
- Type ServiceAccount is watched but has no cache filter (cluster-wide informer)
- Type StatefulSet is watched but has no cache filter (cluster-wide informer)

