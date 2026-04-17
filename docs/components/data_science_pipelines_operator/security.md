# data-science-pipelines-operator: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| mariadb-certs | Opaque | deployment/mariadb |
| ds-pipeline-db-test | Opaque | deployment/mariadb |
| minio-certs | Opaque | deployment/minio |
| minio | Opaque | deployment/minio |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| mariadb | mariadb | ? | ? | ? | `.github/resources/mariadb/deployment.yaml` |
| minio | minio | ? | ? | ? | `.github/resources/minio/deployment.yaml` |
| controller-manager | manager | ? | ? | ? | `config/manager/manager.yaml` |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | ${USER}:${USER} |  | multi-arch | yes | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `.github/build/Dockerfile` | ${CI_BASE} | 2 | root |  |  |  | Unpinned base image: ${CI_BASE}; Unpinned base image: ${CI_BASE}; Container runs as root user |
| `.github/scripts/python_package_upload/Dockerfile` | docker.io/python:3.9 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `docs/example_pipelines/iris/Dockerfile` | docker.io/python:3.9.17 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `tests/resources/Dockerfile` | docker.io/python:3.9.17 | 1 |  |  |  |  | No USER directive found (defaults to root) |

