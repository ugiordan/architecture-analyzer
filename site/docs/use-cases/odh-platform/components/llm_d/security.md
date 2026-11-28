# llm-d: Security

[Overview](index.md) · [Agent Context](agent-context.md) · [Extracted Data](details.md) · [Network](network.md) · [RBAC](rbac.md) · [Security](security.md) · [Dataflow](dataflow.md)

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| llm-d-hf-token | Opaque | deployment/render |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| fma-requester | requester | ? | ? | ? | [`guides/fast-model-actuation/requester/deployment.yaml`](https://github.com/llm-d/llm-d/blob/490eca1567005ebbe07ca43e464f6aee3ae580e8/guides/fast-model-actuation/requester/deployment.yaml) |
| interactive-pod | benchmark-runner | ? | ? | ? | [`helpers/interactive-pod/manifests/deployment.yaml`](https://github.com/llm-d/llm-d/blob/490eca1567005ebbe07ca43e464f6aee3ae580e8/helpers/interactive-pod/manifests/deployment.yaml) |
| llm-d-coordinator | coordinator | ? | ? | ? | [`guides/coord-disaggregation/coordinator/base/deployment.yaml`](https://github.com/llm-d/llm-d/blob/490eca1567005ebbe07ca43e464f6aee3ae580e8/guides/coord-disaggregation/coordinator/base/deployment.yaml) |
| mooncake-master-store | mooncake-master-store | ? | ? | ? | [`helpers/mooncake-master-store/base/deployment.yaml`](https://github.com/llm-d/llm-d/blob/490eca1567005ebbe07ca43e464f6aee3ae580e8/helpers/mooncake-master-store/base/deployment.yaml) |
| render | vllm-render | ? | ? | ? | [`guides/p2p-kv-cache-sharing/render/standalone/deployment.yaml`](https://github.com/llm-d/llm-d/blob/490eca1567005ebbe07ca43e464f6aee3ae580e8/guides/p2p-kv-cache-sharing/render/standalone/deployment.yaml) |
| render | vllm-render | ? | ? | ? | [`guides/precise-prefix-cache-routing/render/standalone/deployment.yaml`](https://github.com/llm-d/llm-d/blob/490eca1567005ebbe07ca43e464f6aee3ae580e8/guides/precise-prefix-cache-routing/render/standalone/deployment.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `docker/Dockerfile.rdma-tools` | nvcr.io/nvidia/cuda:${CUDA_MAJOR}.${CUDA_MINOR}.${CUDA_PATCH}-runtime-ubi9 | 2 |  |  |  |  | No USER directive found (defaults to root) |
| `docker/Dockerfile.xpu` | ${BASE_IMAGE} | 1 |  |  |  |  | Unpinned base image: ${BASE_IMAGE}; No USER directive found (defaults to root) |
| `docker/Dockerfile.xpu-omni` | vllm-omni | 2 |  |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: vllm-omni; No USER directive found (defaults to root) |
| `guides/agent-substrate/hermes/actor/Dockerfile` | python:3.13-slim@sha256:e48f59cdba69a3cbd9ed2cf0f348d90d45a8d350d98134cf1070bc106ee9ed90 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `guides/modelexpress-p2p/image/Dockerfile` | ${BASE_IMAGE} | 1 |  |  |  |  | Unpinned base image: ${BASE_IMAGE}; No USER directive found (defaults to root) |
| `helpers/interactive-pod/build/Dockerfile` | python:3.12-slim | 1 |  |  |  |  | No USER directive found (defaults to root) |

