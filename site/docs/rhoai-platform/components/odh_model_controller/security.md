# odh-model-controller: Security

## Secrets

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| odh-model-controller-webhook-cert | kubernetes.io/tls | deployment/odh-model-controller, service/odh-model-controller-webhook-service |

## Deployment Security Controls

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| odh-model-controller | manager | ? | ? | ? | `config/default/manager_webhook_patch.yaml` |
| odh-model-controller | manager | ? | ? | ? | `config/manager/manager.yaml` |

## Build Security

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Containerfile` | registry.access.redhat.com/ubi9/ubi-minimal:9.5 | 2 | 65532:65532 |  | multi-arch |  |  |

## Cache Architecture

### Manager Configuration

| Property | Value |
|----------|-------|
| Manager file | `cmd/main.go` |
| Cache scope | cluster-wide |
| DefaultTransform | no |
| Memory limit | 2Gi |

### Filtered Types

| Type | Filter Kind | Filter |
|------|-------------|--------|
| corev1.Secret | label | opendatahub.io/managed=true |
| corev1.Pod | label | component=predictor |

### Issues

- No DefaultTransform: managedFields cached for all objects (wasted memory)
- No GOMEMLIMIT set in deployment (Go GC cannot pressure-tune)
- Type Account is watched but has no cache filter (cluster-wide informer)
- Type AuthPolicy is watched but has no cache filter (cluster-wide informer)
- Type Authorino is watched but has no cache filter (cluster-wide informer)
- Type ClusterRoleBinding is watched but has no cache filter (cluster-wide informer)
- Type ConfigMap is watched but has no cache filter (cluster-wide informer)
- Type EnvoyFilter is watched but has no cache filter (cluster-wide informer)
- Type Gateway is watched but has no cache filter (cluster-wide informer)
- Type InferenceGraph is watched but has no cache filter (cluster-wide informer)
- Type InferenceService is watched but has no cache filter (cluster-wide informer)
- Type Kuadrant is watched but has no cache filter (cluster-wide informer)
- Type LLMInferenceService is watched but has no cache filter (cluster-wide informer)
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

