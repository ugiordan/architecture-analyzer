# llm-d-kv-cache: Security

[Overview](index.md) · [Agent Context](agent-context.md) · [Extracted Data](details.md) · [Network](network.md) · [RBAC](rbac.md) · [Security](security.md) · [Dataflow](dataflow.md)

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| 0 | cmd | ? | ? | ? | [`deploy/common/statefulset.yaml`](https://github.com/opendatahub-io/llm-d-kv-cache/blob/67d00aa3d0e0b9e41e51f12d06b651570039f79f/deploy/common/statefulset.yaml) |
| llm-d-kv-cache-pvc-evictor | pvc-evictor | true | ? | ? | [`kv_connectors/pvc_evictor/deploy/deployment.yaml`](https://github.com/opendatahub-io/llm-d-kv-cache/blob/67d00aa3d0e0b9e41e51f12d06b651570039f79f/kv_connectors/pvc_evictor/deploy/deployment.yaml) |

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
| `kv_connectors/pvc_evictor/go/Dockerfile.konflux` | registry.access.redhat.com/ubi10/ubi-minimal:10.2@sha256:ab67cd02caf8706e34a695a1432ebe965934f6128ec352a6c3f4048b43a4fb5e | 2 | 65532:65532 |  | multi-arch |  |  |
| `services/uds_tokenizer/Dockerfile` | python:3.12-slim | 2 | 65532:65532 |  | multi-arch |  |  |
| `services/uds_tokenizer/Dockerfile.konflux` | registry.access.redhat.com/ubi9/python-312:9.7 | 2 | 65532:65532 |  | multi-arch |  |  |

