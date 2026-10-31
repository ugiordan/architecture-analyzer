# llm-d-kv-cache: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| 0 | cmd | ? | ? | ? | [`deploy/common/statefulset.yaml`](https://github.com/red-hat-data-services/llm-d-kv-cache/blob/fb80f035daf5408b7be78adaca31cc963a9b770c/deploy/common/statefulset.yaml) |
| llm-d-kv-cache-pvc-evictor | pvc-evictor | true | ? | ? | [`kv_connectors/pvc_evictor/deploy/deployment.yaml`](https://github.com/red-hat-data-services/llm-d-kv-cache/blob/fb80f035daf5408b7be78adaca31cc963a9b770c/kv_connectors/pvc_evictor/deploy/deployment.yaml) |

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
| `kv_connectors/pvc_evictor/go/Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc:9.8-1793278488@sha256:58d09489b1d9123e24d691e47af6c390fe5e08f50c7f4f2967bb64df7ab72aa0 | 2 | 65532:65532 |  | multi-arch |  |  |
| `services/uds_tokenizer/Dockerfile` | python:3.12-slim | 2 | 65532:65532 |  | multi-arch |  |  |
| `services/uds_tokenizer/Dockerfile.konflux` | runtime | 4 | 65532:65532 |  | multi-arch |  | Unpinned base image: runtime; Unpinned base image: runtime |

