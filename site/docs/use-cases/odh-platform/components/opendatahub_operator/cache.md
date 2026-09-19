# opendatahub-operator: Cache Architecture

Controller-runtime cache configuration controls which Kubernetes resources are cached in-memory. Misconfigured caches (cluster-wide watches on high-cardinality types without filters) are a primary cause of operator OOM kills.

## Cache Architecture

### Manager Configuration

| Property | Value |
|----------|-------|
| Manager file | `cmd/main.go` |
| Cache scope | namespace-scoped |
| DefaultTransform | yes |
| Memory limit | 4Gi |

### Filtered Types

| Type | Filter Kind | Filter |
|------|-------------|--------|
| appsv1.Deployment | namespace | namespace-scoped |
| corev1.ConfigMap | namespace | namespace-scoped |
| corev1.Secret | namespace | namespace-scoped |
| extv1.CustomResourceDefinition | label | label selector |
| networkingv1.NetworkPolicy | namespace | namespace-scoped |
| rbacv1.ClusterRole | label | label selector |
| rbacv1.ClusterRoleBinding | label | label selector |
| rbacv1.Role | namespace | namespace-scoped |
| rbacv1.RoleBinding | namespace | namespace-scoped |

### Cache-Bypassed Types (DisableFor)

- authorizationv1.SelfSubjectRulesReview
- corev1.Pod
- ofapiv1alpha1.CatalogSource
- ofapiv1alpha1.Subscription
- userv1.Group

### Issues

- No GOMEMLIMIT set in deployment (Go GC cannot pressure-tune). Set GOMEMLIMIT to 80-90% of container memory limit for optimal GC behavior
- Type Auth is watched but has no cache filter (cluster-wide informer)
- Type MutatingWebhookConfiguration is watched but has no cache filter (cluster-wide informer)
- Type Namespace is watched but has no cache filter (cluster-wide informer)
- Type PodMonitor is watched but has no cache filter (cluster-wide informer)
- Type PrometheusRule is watched but has no cache filter (cluster-wide informer)
- Type SecurityContextConstraints is watched but has no cache filter (cluster-wide informer)
- Type Service is watched but has no cache filter (cluster-wide informer)
- Type ServiceAccount is watched but has no cache filter (cluster-wide informer)
- Type ServiceMonitor is watched but has no cache filter (cluster-wide informer)
- Type ValidatingWebhookConfiguration is watched but has no cache filter (cluster-wide informer)

