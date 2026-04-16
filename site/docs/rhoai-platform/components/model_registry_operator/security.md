# model-registry-operator: Security

## Secrets

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| webhook-server-cert | Opaque | deployment/controller-manager |
| controller-manager-metrics-service | Opaque | deployment/controller-manager |

## Deployment Security Controls

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| controller-manager | manager | ? | ? | ? | `config/default/manager_auth_proxy_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `config/default/manager_config_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `config/default/manager_webhook_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `config/manager/manager.yaml` |
| controller-manager | manager | ? | ? | ? | `config/overlays/odh/patches/manager_auth_proxy_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `config/overlays/odh/patches/manager_istio_config_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `config/overlays/odh/patches/manager_migration_env_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `config/overlays/odh/patches/manager_webhook_patch.yaml` |

## Build Security

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |

## Cache Architecture

### Manager Configuration

| Property | Value |
|----------|-------|
| Manager file | `cmd/main.go` |
| Cache scope | cluster-wide |
| DefaultTransform | no |
| Memory limit | 512Mi |

### Filtered Types

| Type | Filter Kind | Filter |
|------|-------------|--------|
| appsv1.Deployment | label | app.kubernetes.io/created-by=model-registry-operator |
| corev1.ConfigMap | label | app.kubernetes.io/created-by=model-registry-operator |
| corev1.PersistentVolumeClaim | label | app.kubernetes.io/created-by=model-registry-operator |
| corev1.ServiceAccount | label | app.kubernetes.io/created-by=model-registry-operator |
| corev1.Service | label | app.kubernetes.io/created-by=model-registry-operator |
| networkingv1.NetworkPolicy | label | app.kubernetes.io/created-by=model-registry-operator |
| rbacv1.ClusterRoleBinding | label | app.kubernetes.io/created-by=model-registry-operator |
| rbacv1.RoleBinding | label | app.kubernetes.io/created-by=model-registry-operator |
| rbacv1.Role | label | app.kubernetes.io/created-by=model-registry-operator |

### Cache-Bypassed Types (DisableFor)

- corev1.Secret

### Issues

- No DefaultTransform: managedFields cached for all objects (wasted memory)
- No GOMEMLIMIT set in deployment (Go GC cannot pressure-tune)
- Type ModelRegistry is watched but has no cache filter (cluster-wide informer)

