# llm-d: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| llm-d-hf-token | Opaque | deployment/llm-d-coordinator, deployment/render |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| fma-requester | requester | ? | ? | ? | [`guides/fast-model-actuation/requester/deployment.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/guides/fast-model-actuation/requester/deployment.yaml) |
| interactive-pod | benchmark-runner | ? | ? | ? | [`helpers/interactive-pod/manifests/deployment.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/helpers/interactive-pod/manifests/deployment.yaml) |
| llm-d-coordinator | coordinator | ? | ? | ? | [`guides/coord-disaggregation/coordinator/deployment.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/guides/coord-disaggregation/coordinator/deployment.yaml) |
| llm-d-coordinator | vllm-render | ? | ? | ? | [`guides/coord-disaggregation/coordinator/deployment.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/guides/coord-disaggregation/coordinator/deployment.yaml) |
| mooncake-master-store | mooncake-master-store | ? | ? | ? | [`helpers/mooncake-master-store/base/deployment.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/helpers/mooncake-master-store/base/deployment.yaml) |
| render | vllm-render | ? | ? | ? | [`guides/p2p-kv-cache-sharing/render/standalone/deployment.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/guides/p2p-kv-cache-sharing/render/standalone/deployment.yaml) |
| render | vllm-render | ? | ? | ? | [`guides/precise-prefix-cache-routing/render/standalone/deployment.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/guides/precise-prefix-cache-routing/render/standalone/deployment.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `docker/Dockerfile.cpu` | vllm-build | 5 |  |  | multi-arch |  | Unpinned base image: base-common; Unpinned base image: base-common; Unpinned base image: base-${TARGETARCH}; Unpinned base image: vllm-build; No USER directive found (defaults to root) |
| `docker/Dockerfile.cuda` | nvcr.io/nvidia/cuda:${CUDA_MAJOR}.${CUDA_MINOR}.${CUDA_PATCH}-runtime-${FINAL_BASE_IMAGE_SUFFIX} | 2 | 2000 |  | multi-arch |  |  |
| `docker/Dockerfile.rdma-tools` | nvcr.io/nvidia/cuda:${CUDA_MAJOR}.${CUDA_MINOR}.${CUDA_PATCH}-runtime-ubi9 | 2 |  |  |  |  | No USER directive found (defaults to root) |
| `docker/Dockerfile.rocm` | base | 3 |  |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: base; Unpinned base image: base; No USER directive found (defaults to root) |
| `docker/Dockerfile.xpu` | ${BASE_IMAGE} | 1 |  |  |  |  | Unpinned base image: ${BASE_IMAGE}; No USER directive found (defaults to root) |
| `guides/modelexpress-p2p/image/Dockerfile` | ${BASE_IMAGE} | 1 |  |  |  |  | Unpinned base image: ${BASE_IMAGE}; No USER directive found (defaults to root) |
| `helpers/interactive-pod/build/Dockerfile` | python:3.12-slim | 1 |  |  |  |  | No USER directive found (defaults to root) |

