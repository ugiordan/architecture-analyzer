# data-science-pipelines-operator: Security

## Secrets

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| mariadb-certs | Opaque | deployment/mariadb |
| ds-pipeline-db-test | Opaque | deployment/mariadb |
| minio-certs | Opaque | deployment/minio |
| minio | Opaque | deployment/minio |

## Deployment Security Controls

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| mariadb | mariadb | ? | ? | ? | `.github/resources/mariadb/deployment.yaml` |
| minio | minio | ? | ? | ? | `.github/resources/minio/deployment.yaml` |
| controller-manager | manager | ? | ? | ? | `config/manager/manager.yaml` |

## Build Security

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | ${USER}:${USER} |  | multi-arch | yes | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `.github/build/Dockerfile` | ${CI_BASE} | 2 | root |  |  |  | Unpinned base image: ${CI_BASE}; Unpinned base image: ${CI_BASE}; Container runs as root user |
| `.github/scripts/python_package_upload/Dockerfile` | docker.io/python:3.9 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `docs/example_pipelines/iris/Dockerfile` | docker.io/python:3.9.17 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `tests/resources/Dockerfile` | docker.io/python:3.9.17 | 1 |  |  |  |  | No USER directive found (defaults to root) |

## Cache Architecture

### Manager Configuration

| Property | Value |
|----------|-------|
| Manager file | `main.go` |
| Cache scope | cluster-wide |
| DefaultTransform | yes |
| GOMEMLIMIT | 3600MiB |
| Memory limit | 4Gi |

### Filtered Types

| Type | Filter Kind | Filter |
|------|-------------|--------|
| admv1.ValidatingWebhookConfiguration | field | field selector |
| admv1.MutatingWebhookConfiguration | field | field selector |
| corev1.Pod | label | label selector |
| corev1.Service | label | label selector |
| corev1.ServiceAccount | label | label selector |
| corev1.PersistentVolumeClaim | label | label selector |
| rbacv1.Role | label | label selector |
| rbacv1.RoleBinding | label | label selector |
| routev1.Route | label | label selector |

### Transform-Stripped Types

- corev1.ConfigMap
- corev1.Secret

### Issues

- Type ConfigMap is watched but has no cache filter (cluster-wide informer)
- Type DataSciencePipelinesApplication is watched but has no cache filter (cluster-wide informer)
- Type Deployment is watched but has no cache filter (cluster-wide informer)
- Type NetworkPolicy is watched but has no cache filter (cluster-wide informer)
- Type Secret is watched but has no cache filter (cluster-wide informer)

