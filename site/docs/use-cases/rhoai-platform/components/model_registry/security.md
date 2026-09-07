# model-registry: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| minio-secret | Opaque | deployment/minio |
| model-catalog-hf-api-key | Opaque | deployment/model-catalog-server |
| model-catalog-postgres | Opaque | deployment/model-catalog-server |
| skill-catalog-git-credentials | Opaque | deployment/model-catalog-server |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| controller-manager | manager | ? | ? | ? | [`manifests/kustomize/options/controller/manager/manager.yaml`](https://github.com/kubeflow/model-registry/blob/75376c10211d5678fad2b267b226a92d8bdba6bf/manifests/kustomize/options/controller/manager/manager.yaml) |
| minio | minio | ? | ? | ? | [`scripts/manifests/minio/deployment.yaml`](https://github.com/kubeflow/model-registry/blob/75376c10211d5678fad2b267b226a92d8bdba6bf/scripts/manifests/minio/deployment.yaml) |
| model-catalog-server | catalog | ? | ? | ? | [`manifests/kustomize/options/catalog/base/deployment.yaml`](https://github.com/kubeflow/model-registry/blob/75376c10211d5678fad2b267b226a92d8bdba6bf/manifests/kustomize/options/catalog/base/deployment.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: common; Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:3e009398a8aa8eec621393fbf308c5e622f174900e44e8d5fe224c637920924a | 2 | 1001 |  |  |  |  |
| `Dockerfile.odh` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 65532:65532 |  |  |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `Dockerfile.testops` | registry.access.redhat.com/ubi9/python-312 | 1 | odh |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/python-312 |
| `clients/ui/Dockerfile` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `clients/ui/Dockerfile.standalone` | release | 4 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE}; Unpinned base image: release |
| `cmd/controller/Dockerfile.controller` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `cmd/csi/Dockerfile.csi` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: common; Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `jobs/async-upload/Dockerfile` | registry.access.redhat.com/ubi9/python-312-minimal | 2 | 1000 |  |  |  | Unpinned base image: registry.access.redhat.com/ubi9/python-312-minimal |
| `jobs/async-upload/Dockerfile.konflux` | quay.io/aipcc/base-images/cpu:3.5@sha256:1459233907bdcfe4fdb4f9a8d39de861bacd868d7a1e5cee083ddd46f93c306d | 2 | 1000 |  | multi-arch |  |  |

