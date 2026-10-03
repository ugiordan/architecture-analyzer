# mcp-lifecycle-module-operator: Cache Architecture

Controller-runtime cache configuration controls which Kubernetes resources are cached in-memory. Misconfigured caches (cluster-wide watches on high-cardinality types without filters) are a primary cause of operator OOM kills.

## Cache Architecture

### Manager Configuration

| Property | Value |
|----------|-------|
| Manager file | `cmd/main.go` |
| Cache scope | namespace-scoped |
| DefaultTransform | yes |
| GOMEMLIMIT | 482344960 |
| Memory limit | 512Mi |

### Filtered Types

| Type | Filter Kind | Filter |
|------|-------------|--------|
| extv1.CustomResourceDefinition | label | label selector |
| rbacv1.ClusterRole | label | label selector |
| rbacv1.ClusterRoleBinding | label | label selector |

### Issues

- Type APIServer is watched but has no cache filter (cluster-wide informer)
- Type ConfigMap is watched but has no cache filter (cluster-wide informer)
- Type Deployment is watched but has no cache filter (cluster-wide informer)
- Type MCPLifecycleOperator is watched but has no cache filter (cluster-wide informer)
- Type Role is watched but has no cache filter (cluster-wide informer)
- Type RoleBinding is watched but has no cache filter (cluster-wide informer)
- Type Service is watched but has no cache filter (cluster-wide informer)
- Type ServiceAccount is watched but has no cache filter (cluster-wide informer)

