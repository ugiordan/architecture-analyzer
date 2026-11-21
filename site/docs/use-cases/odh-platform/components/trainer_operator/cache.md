# trainer-operator: Cache Architecture

Controller-runtime cache configuration controls which Kubernetes resources are cached in-memory. Misconfigured caches (cluster-wide watches on high-cardinality types without filters) are a primary cause of operator OOM kills.

## Cache Architecture

### Manager Configuration

| Property | Value |
|----------|-------|
| Manager file | `cmd/main.go` |
| Cache scope | namespace-scoped |
| DefaultTransform | yes |
| Memory limit | 512Mi |

### Filtered Types

| Type | Filter Kind | Filter |
|------|-------------|--------|
| admissionv1.ValidatingWebhookConfiguration | label | platformlabels.PlatformPartOf=trainerPartOf (constants, resolved at runtime) |
| appsv1.Deployment | label | platformlabels.PlatformPartOf=trainerPartOf (constants, resolved at runtime) |
| corev1.Service | label | platformlabels.PlatformPartOf=trainerPartOf (constants, resolved at runtime) |

### Cache-Bypassed Types (DisableFor)

- corev1.Namespace

### Issues

- No GOMEMLIMIT set in deployment (Go GC cannot pressure-tune). Set GOMEMLIMIT to 80-90% of container memory limit for optimal GC behavior
- Type ConfigMap is watched but has no cache filter (cluster-wide informer)

