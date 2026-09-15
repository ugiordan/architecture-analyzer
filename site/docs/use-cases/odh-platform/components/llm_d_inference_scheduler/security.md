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
| ${EPP_NAME} | epp | ? | ? | ? | [`deploy/components/inference-gateway/deployment.yaml`](https://github.com/llm-d/llm-d-inference-scheduler/blob/3bc2c129cb53a005361a7d3da424b5d081e24647/deploy/components/inference-gateway/deployment.yaml) |
| 0 | cmd | ? | ? | ? | [`deploy/environments/kubernetes-base/common/statefulset.yaml`](https://github.com/llm-d/llm-d-inference-scheduler/blob/3bc2c129cb53a005361a7d3da424b5d081e24647/deploy/environments/kubernetes-base/common/statefulset.yaml) |
| istiod-llm-d-gateway | discovery | true | true | ? | [`deploy/components/istio-control-plane/deployment.yaml`](https://github.com/llm-d/llm-d-inference-scheduler/blob/3bc2c129cb53a005361a7d3da424b5d081e24647/deploy/components/istio-control-plane/deployment.yaml) |
| llm-d-coordinator | coordinator | ? | ? | ? | [`deploy/coordinator/deployment.yaml`](https://github.com/llm-d/llm-d-inference-scheduler/blob/3bc2c129cb53a005361a7d3da424b5d081e24647/deploy/coordinator/deployment.yaml) |
| vllm-d | vllm | ? | ? | ? | [`deploy/components/vllm-decode/deployment.yaml`](https://github.com/llm-d/llm-d-inference-scheduler/blob/3bc2c129cb53a005361a7d3da424b5d081e24647/deploy/components/vllm-decode/deployment.yaml) |
| vllm-e | vllm | ? | ? | ? | [`deploy/components/vllm-encode/deployment.yaml`](https://github.com/llm-d/llm-d-inference-scheduler/blob/3bc2c129cb53a005361a7d3da424b5d081e24647/deploy/components/vllm-encode/deployment.yaml) |
| vllm-p | vllm | ? | ? | ? | [`deploy/components/vllm-prefill/deployment.yaml`](https://github.com/llm-d/llm-d-inference-scheduler/blob/3bc2c129cb53a005361a7d3da424b5d081e24647/deploy/components/vllm-prefill/deployment.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile.builder` | golang:1.26.8 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `Dockerfile.coordinator` | ${BASE_IMAGE} | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE} |
| `Dockerfile.epp` | ${BASE_IMAGE} | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE} |
| `Dockerfile.konflux.epp` | registry.redhat.io/ubi9/ubi-minimal-pqc:9.8-1788217333@sha256:3e009398a8aa8eec621393fbf308c5e622f174900e44e8d5fe224c637920924a | 2 | 65532:65532 |  | multi-arch |  |  |
| `Dockerfile.konflux.sidecar` | registry.redhat.io/ubi9/ubi-minimal-pqc:9.8-1788217333@sha256:3e009398a8aa8eec621393fbf308c5e622f174900e44e8d5fe224c637920924a | 2 | 65532:65532 |  | multi-arch |  |  |
| `Dockerfile.sidecar` | ${BASE_IMAGE} | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE} |

