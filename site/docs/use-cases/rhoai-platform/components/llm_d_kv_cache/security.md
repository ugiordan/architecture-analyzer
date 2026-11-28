# llm-d-kv-cache: Security

[Overview](index.md) · [Agent Context](agent-context.md) · [Extracted Data](details.md) · [Network](network.md) · [RBAC](rbac.md) · [Security](security.md) · [Dataflow](dataflow.md)

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| 0 | cmd | ? | ? | ? | [`deploy/common/statefulset.yaml`](https://github.com/red-hat-data-services/llm-d-kv-cache/blob/d6085b1e6c74a5eb545e57deb1eab24921675145/deploy/common/statefulset.yaml) |
| llm-d-kv-cache-pvc-evictor | pvc-evictor | true | ? | ? | [`kv_connectors/pvc_evictor/deploy/deployment.yaml`](https://github.com/red-hat-data-services/llm-d-kv-cache/blob/d6085b1e6c74a5eb545e57deb1eab24921675145/kv_connectors/pvc_evictor/deploy/deployment.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi9/ubi:latest | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi:latest |
| `kv_connectors/llmd_fs_backend/Dockerfile.dev` | ${VLLM_IMAGE} | 1 |  |  |  |  | Unpinned base image: ${VLLM_IMAGE}; No USER directive found (defaults to root) |
| `kv_connectors/llmd_fs_backend/Dockerfile.wheel` | nvcr.io/nvidia/cuda:${CUDA_BASE_TAG} | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `kv_connectors/pvc_evictor/Dockerfile` | python:3.12-slim | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `kv_connectors/pvc_evictor/Dockerfile.konflux` | runtime | 3 | 65532:65532 |  |  |  | Unpinned base image: runtime |
| `kv_connectors/pvc_evictor/go/Dockerfile` | ${BASE_IMAGE} | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE} |
| `kv_connectors/pvc_evictor/go/Dockerfile.konflux` | registry.access.redhat.com/ubi10/ubi-minimal:10.3-1795196920@sha256:d63890a635f8f19177bce1775f74f587a47f4456844f1392456c4e3673a48753 | 2 | 65532:65532 |  | multi-arch |  |  |
| `services/uds_tokenizer/Dockerfile` | python:3.12-slim | 2 | 65532:65532 |  | multi-arch |  |  |
| `services/uds_tokenizer/Dockerfile.konflux` | runtime | 4 | 65532:65532 |  | multi-arch |  | Unpinned base image: runtime; Unpinned base image: runtime |

