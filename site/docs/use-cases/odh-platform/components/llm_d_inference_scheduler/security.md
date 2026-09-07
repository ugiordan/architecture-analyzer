# llm-d-inference-scheduler: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| cacerts | Opaque | deployment/istiod-llm-d-gateway |
| istio-kubeconfig | Opaque | deployment/istiod-llm-d-gateway |
| istiod-tls | Opaque | deployment/istiod-llm-d-gateway |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| ${EPP_NAME} | epp | ? | ? | ? | [`deploy/components/inference-gateway/deployment.yaml`](https://github.com/llm-d/llm-d-inference-scheduler/blob/d9a2167b9d89bf54b5fc5d76c3318e3c08f82084/deploy/components/inference-gateway/deployment.yaml) |
| 0 | cmd | ? | ? | ? | [`deploy/environments/kubernetes-base/common/statefulset.yaml`](https://github.com/llm-d/llm-d-inference-scheduler/blob/d9a2167b9d89bf54b5fc5d76c3318e3c08f82084/deploy/environments/kubernetes-base/common/statefulset.yaml) |
| istiod-llm-d-gateway | discovery | true | true | ? | [`deploy/components/istio-control-plane/deployment.yaml`](https://github.com/llm-d/llm-d-inference-scheduler/blob/d9a2167b9d89bf54b5fc5d76c3318e3c08f82084/deploy/components/istio-control-plane/deployment.yaml) |
| llm-d-coordinator | coordinator | ? | ? | ? | [`deploy/coordinator/deployment.yaml`](https://github.com/llm-d/llm-d-inference-scheduler/blob/d9a2167b9d89bf54b5fc5d76c3318e3c08f82084/deploy/coordinator/deployment.yaml) |
| vllm-d | vllm | ? | ? | ? | [`deploy/components/vllm-decode/deployment.yaml`](https://github.com/llm-d/llm-d-inference-scheduler/blob/d9a2167b9d89bf54b5fc5d76c3318e3c08f82084/deploy/components/vllm-decode/deployment.yaml) |
| vllm-e | vllm | ? | ? | ? | [`deploy/components/vllm-encode/deployment.yaml`](https://github.com/llm-d/llm-d-inference-scheduler/blob/d9a2167b9d89bf54b5fc5d76c3318e3c08f82084/deploy/components/vllm-encode/deployment.yaml) |
| vllm-p | vllm | ? | ? | ? | [`deploy/components/vllm-prefill/deployment.yaml`](https://github.com/llm-d/llm-d-inference-scheduler/blob/d9a2167b9d89bf54b5fc5d76c3318e3c08f82084/deploy/components/vllm-prefill/deployment.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile.builder` | golang:1.26.7 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `Dockerfile.coordinator` | ${BASE_IMAGE} | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE} |
| `Dockerfile.epp` | ${BASE_IMAGE} | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE} |
| `Dockerfile.konflux.epp` | registry.access.redhat.com/ubi9/ubi-minimal:9.7@sha256:d91be7cea9f03a757d69ad7fcdfcd7849dba820110e7980d5e2a1f46ed06ea3b | 2 | 65532:65532 |  | multi-arch |  |  |
| `Dockerfile.konflux.sidecar` | registry.access.redhat.com/ubi9/ubi-minimal:9.7@sha256:d91be7cea9f03a757d69ad7fcdfcd7849dba820110e7980d5e2a1f46ed06ea3b | 2 | 65532:65532 |  | multi-arch |  |  |
| `Dockerfile.sidecar` | ${BASE_IMAGE} | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE} |

