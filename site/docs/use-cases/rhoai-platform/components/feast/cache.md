# feast: Cache Architecture

Controller-runtime cache configuration controls which Kubernetes resources are cached in-memory. Misconfigured caches (cluster-wide watches on high-cardinality types without filters) are a primary cause of operator OOM kills.

## Cache Architecture

### Manager Configuration

| Property | Value |
|----------|-------|
| Manager file | `infra/feast-operator/cmd/main.go` |
| Cache scope | cluster-wide |
| DefaultTransform | yes |
| GOMEMLIMIT | 230MiB |
| Memory limit | 256Mi |

### Filtered Types

| Type | Filter Kind | Filter |
|------|-------------|--------|
| appsv1.Deployment | label | label selector |
| autoscalingv2.HorizontalPodAutoscaler | label | label selector |
| batchv1.CronJob | label | label selector |
| corev1.ConfigMap | label | label selector |
| corev1.PersistentVolumeClaim | label | label selector |
| corev1.Service | label | label selector |
| corev1.ServiceAccount | label | label selector |
| policyv1.PodDisruptionBudget | label | label selector |
| rbacv1.Role | label | label selector |
| rbacv1.RoleBinding | label | label selector |

### Cache-Bypassed Types (DisableFor)

- appsv1.Deployment
- autoscalingv2.HorizontalPodAutoscaler
- batchv1.CronJob
- corev1.ConfigMap
- corev1.PersistentVolumeClaim
- corev1.Secret
- corev1.Service
- corev1.ServiceAccount
- policyv1.PodDisruptionBudget
- rbacv1.Role
- rbacv1.RoleBinding

### Issues

- Cache bypass (DisableFor) configured for corev1.ConfigMap. This is a common fix for OOM caused by informer cache flooding from high-cardinality types (e.g., opendatahub-io/model-registry-operator#457)
- Cache bypass (DisableFor) configured for corev1.Secret. This is a common fix for OOM caused by informer cache flooding from high-cardinality types (e.g., opendatahub-io/model-registry-operator#457)
- Type FeatureStore is watched but has no cache filter (cluster-wide informer)
- Type Route is watched but has no cache filter (cluster-wide informer)

