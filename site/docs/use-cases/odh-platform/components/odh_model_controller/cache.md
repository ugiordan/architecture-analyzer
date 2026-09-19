# odh-model-controller: Cache Architecture

Controller-runtime cache configuration controls which Kubernetes resources are cached in-memory. Misconfigured caches (cluster-wide watches on high-cardinality types without filters) are a primary cause of operator OOM kills.

## Cache Architecture

### Manager Configuration

| Property | Value |
|----------|-------|
| Manager file | `cmd/main.go` |
| Cache scope | cluster-wide |
| DefaultTransform | yes |
| GOMEMLIMIT | 1800MiB |
| Memory limit | 2Gi |

### Filtered Types

| Type | Filter Kind | Filter |
|------|-------------|--------|
| corev1.Pod | label | component=predictor |
| corev1.Secret | label | opendatahub.io/managed=true |

### Cache-Bypassed Types (DisableFor)

- corev1.ConfigMap

### Issues

- Cache bypass (DisableFor) configured for corev1.ConfigMap. This is a common fix for OOM caused by informer cache flooding from high-cardinality types (e.g., opendatahub-io/model-registry-operator#457)
- Type Account is watched but has no cache filter (cluster-wide informer)
- Type AuthPolicy is watched but has no cache filter (cluster-wide informer)
- Type Authorino is watched but has no cache filter (cluster-wide informer)
- Type ClusterRoleBinding is watched but has no cache filter (cluster-wide informer)
- Type EnvoyFilter is watched but has no cache filter (cluster-wide informer)
- Type Gateway is watched but has no cache filter (cluster-wide informer)
- Type InferenceGraph is watched but has no cache filter (cluster-wide informer)
- Type InferenceService is watched but has no cache filter (cluster-wide informer)
- Type Kuadrant is watched but has no cache filter (cluster-wide informer)
- Type LLMInferenceService is watched but has no cache filter (cluster-wide informer)
- Type LLMInferenceServiceConfig is watched but has no cache filter (cluster-wide informer)
- Type Namespace is watched but has no cache filter (cluster-wide informer)
- Type NetworkPolicy is watched but has no cache filter (cluster-wide informer)
- Type PodMonitor is watched but has no cache filter (cluster-wide informer)
- Type Role is watched but has no cache filter (cluster-wide informer)
- Type RoleBinding is watched but has no cache filter (cluster-wide informer)
- Type Route is watched but has no cache filter (cluster-wide informer)
- Type Service is watched but has no cache filter (cluster-wide informer)
- Type ServiceAccount is watched but has no cache filter (cluster-wide informer)
- Type ServiceMonitor is watched but has no cache filter (cluster-wide informer)
- Type ServingRuntime is watched but has no cache filter (cluster-wide informer)
- Type Template is watched but has no cache filter (cluster-wide informer)
- Type TriggerAuthentication is watched but has no cache filter (cluster-wide informer)

