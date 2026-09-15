# agents-operator: Cache Architecture

Controller-runtime cache configuration controls which Kubernetes resources are cached in-memory. Misconfigured caches (cluster-wide watches on high-cardinality types without filters) are a primary cause of operator OOM kills.

## Cache Architecture

### Manager Configuration

| Property | Value |
|----------|-------|
| Manager file | `kagenti-operator/cmd/main.go` |
| Cache scope | namespace-scoped |
| DefaultTransform | no |
| Memory limit | 128Mi |

### Filtered Types

| Type | Filter Kind | Filter |
|------|-------------|--------|
| corev1.ConfigMap | namespace | namespace-scoped |

### Issues

- No DefaultTransform: managedFields cached for all objects (wasted memory). Add cache.DefaultTransform to strip managedFields and reduce memory footprint
- No GOMEMLIMIT set in deployment (Go GC cannot pressure-tune). Set GOMEMLIMIT to 80-90% of container memory limit for optimal GC behavior
- Type AgentCard is watched but has no cache filter (cluster-wide informer)
- Type AgentRuntime is watched but has no cache filter (cluster-wide informer)
- Type Certificate is watched but has no cache filter (cluster-wide informer)
- Type DataScienceCluster is watched but has no cache filter (cluster-wide informer)
- Type Deployment is watched but has no cache filter (cluster-wide informer)
- Type Kuadrant is watched but has no cache filter (cluster-wide informer)
- Type Namespace is watched but has no cache filter (cluster-wide informer)
- Type NetworkPolicy is watched but has no cache filter (cluster-wide informer)
- Type Role is watched but has no cache filter (cluster-wide informer)
- Type RoleBinding is watched but has no cache filter (cluster-wide informer)
- Type Secret is watched but has no cache filter (cluster-wide informer)
- Type StatefulSet is watched but has no cache filter (cluster-wide informer)
- Type TektonConfig is watched but has no cache filter (cluster-wide informer)

