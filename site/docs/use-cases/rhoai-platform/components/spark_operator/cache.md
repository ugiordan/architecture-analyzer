# spark-operator: Cache Architecture

Controller-runtime cache configuration controls which Kubernetes resources are cached in-memory. Misconfigured caches (cluster-wide watches on high-cardinality types without filters) are a primary cause of operator OOM kills.

## Cache Architecture

### Manager Configuration

| Property | Value |
|----------|-------|
| Manager file | `spark-operator-module/cmd/spark-operator-module/main.go` |
| Cache scope | cluster-wide |
| DefaultTransform | no |
| Memory limit | 512Mi |

### Implicit Informers (OOM Risk)

| Type | Source | Risk |
|------|--------|------|
| corev1.Pod | `internal/controller/sparkapplication/controller.go:1330` | **HIGH** |
| networkingv1.Ingress | `internal/controller/sparkapplication/controller.go:1345` | medium |

### Issues

- Implicit informer for corev1.Pod via client.Get at internal/controller/sparkapplication/controller.go:1330 (cluster-wide, OOM risk). This bypasses cache filters and creates a full cluster-wide watch
- No GOMEMLIMIT set in deployment (Go GC cannot pressure-tune). Set GOMEMLIMIT to 80-90% of container memory limit for optimal GC behavior
- No cache configuration: all informers are cluster-wide (OOM risk). See https://book.kubebuilder.io/reference/watching-resources/filtering for cache filtering patterns
- Type ClusterRole is watched but has no cache filter (cluster-wide informer)
- Type ClusterRoleBinding is watched but has no cache filter (cluster-wide informer)
- Type ConfigMap is watched but has no cache filter (cluster-wide informer)
- Type CustomResourceDefinition is watched but has no cache filter (cluster-wide informer)
- Type Deployment is watched but has no cache filter (cluster-wide informer)
- Type MutatingWebhookConfiguration is watched but has no cache filter (cluster-wide informer)
- Type NetworkPolicy is watched but has no cache filter (cluster-wide informer)
- Type Role is watched but has no cache filter (cluster-wide informer)
- Type RoleBinding is watched but has no cache filter (cluster-wide informer)
- Type Service is watched but has no cache filter (cluster-wide informer)
- Type ServiceAccount is watched but has no cache filter (cluster-wide informer)
- Type SparkConnect is watched but has no cache filter (cluster-wide informer)
- Type SparkOperator is watched but has no cache filter (cluster-wide informer)
- Type ValidatingWebhookConfiguration is watched but has no cache filter (cluster-wide informer)

