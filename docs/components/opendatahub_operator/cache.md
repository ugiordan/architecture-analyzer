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
| corev1.Secret | namespace | namespace-scoped |
| corev1.ConfigMap | namespace | namespace-scoped |
| appsv1.Deployment | namespace | namespace-scoped |
| networkingv1.NetworkPolicy | namespace | namespace-scoped |
| rbacv1.Role | namespace | namespace-scoped |
| rbacv1.RoleBinding | namespace | namespace-scoped |
| rbacv1.ClusterRole | label | label selector |
| rbacv1.ClusterRoleBinding | label | label selector |

### Cache-Bypassed Types (DisableFor)

- ofapiv1alpha1.Subscription
- authorizationv1.SelfSubjectRulesReview
- corev1.Pod
- userv1.Group
- ofapiv1alpha1.CatalogSource

### Implicit Informers (OOM Risk)

| Type | Source | Risk |
|------|--------|------|
| client.ListOptions | `internal/controller/components/modelsasservice/modelsasservice_controller_prerequisites.go:163` | medium |

### Issues

- No GOMEMLIMIT set in deployment (Go GC cannot pressure-tune)
- Type Auth is watched but has no cache filter (cluster-wide informer)
- Type ConsoleLink is watched but has no cache filter (cluster-wide informer)
- Type Dashboard is watched but has no cache filter (cluster-wide informer)
- Type DataSciencePipelines is watched but has no cache filter (cluster-wide informer)
- Type FeastOperator is watched but has no cache filter (cluster-wide informer)
- Type Gateway is watched but has no cache filter (cluster-wide informer)
- Type HTTPRoute is watched but has no cache filter (cluster-wide informer)
- Type Job is watched but has no cache filter (cluster-wide informer)
- Type Kserve is watched but has no cache filter (cluster-wide informer)
- Type Kueue is watched but has no cache filter (cluster-wide informer)
- Type LlamaStackOperator is watched but has no cache filter (cluster-wide informer)
- Type MLflowOperator is watched but has no cache filter (cluster-wide informer)
- Type MaaSTenant is watched but has no cache filter (cluster-wide informer)
- Type ModelController is watched but has no cache filter (cluster-wide informer)
- Type ModelRegistry is watched but has no cache filter (cluster-wide informer)
- Type MutatingWebhookConfiguration is watched but has no cache filter (cluster-wide informer)
- Type Namespace is watched but has no cache filter (cluster-wide informer)
- Type PodDisruptionBudget is watched but has no cache filter (cluster-wide informer)
- Type PodMonitor is watched but has no cache filter (cluster-wide informer)
- Type PrometheusRule is watched but has no cache filter (cluster-wide informer)
- Type Ray is watched but has no cache filter (cluster-wide informer)
- Type Route is watched but has no cache filter (cluster-wide informer)
- Type SecurityContextConstraints is watched but has no cache filter (cluster-wide informer)
- Type Service is watched but has no cache filter (cluster-wide informer)
- Type ServiceAccount is watched but has no cache filter (cluster-wide informer)
- Type ServiceMonitor is watched but has no cache filter (cluster-wide informer)
- Type SparkOperator is watched but has no cache filter (cluster-wide informer)
- Type Template is watched but has no cache filter (cluster-wide informer)
- Type Trainer is watched but has no cache filter (cluster-wide informer)
- Type TrainingOperator is watched but has no cache filter (cluster-wide informer)
- Type TrustyAI is watched but has no cache filter (cluster-wide informer)
- Type ValidatingAdmissionPolicy is watched but has no cache filter (cluster-wide informer)
- Type ValidatingAdmissionPolicyBinding is watched but has no cache filter (cluster-wide informer)
- Type ValidatingWebhookConfiguration is watched but has no cache filter (cluster-wide informer)
- Type Workbenches is watched but has no cache filter (cluster-wide informer)

