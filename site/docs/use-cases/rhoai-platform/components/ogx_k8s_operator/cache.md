# ogx-k8s-operator: Cache Architecture

Controller-runtime cache configuration controls which Kubernetes resources are cached in-memory. Misconfigured caches (cluster-wide watches on high-cardinality types without filters) are a primary cause of operator OOM kills.

## Cache Architecture

### Manager Configuration

| Property | Value |
|----------|-------|
| Manager file | `ogx-module/cmd/ogx-module/main.go` |
| Cache scope | namespace-scoped |
| DefaultTransform | yes |
| GOMEMLIMIT | 800MiB |
| Memory limit | 1Gi |

### Filtered Types

| Type | Filter Kind | Filter |
|------|-------------|--------|
| appsv1.Deployment | label | label selector |
| autoscalingv2.HorizontalPodAutoscaler | label | label selector |
| corev1.ConfigMap | label | controllers.WatchLabelKey=controllers.WatchLabelValue (constants, resolved at runtime) |
| corev1.PersistentVolumeClaim | label | label selector |
| corev1.Secret | label | controllers.WatchLabelKey=controllers.WatchLabelValue (constants, resolved at runtime) |
| corev1.Service | label | label selector |
| networkingv1.Ingress | label | label selector |
| networkingv1.NetworkPolicy | label | label selector |
| policyv1.PodDisruptionBudget | label | label selector |

### Issues

- GOMEMLIMIT ratio 78.1% is below recommended 80% minimum (GC cannot pressure-tune effectively)
- Type ClusterRole is watched but has no cache filter (cluster-wide informer)
- Type ClusterRoleBinding is watched but has no cache filter (cluster-wide informer)
- Type CustomResourceDefinition is watched but has no cache filter (cluster-wide informer)
- Type OGX is watched but has no cache filter (cluster-wide informer)
- Type OGXServer is watched but has no cache filter (cluster-wide informer)
- Type Role is watched but has no cache filter (cluster-wide informer)
- Type RoleBinding is watched but has no cache filter (cluster-wide informer)
- Type ServiceAccount is watched but has no cache filter (cluster-wide informer)
- Type ValidatingWebhookConfiguration is watched but has no cache filter (cluster-wide informer)

