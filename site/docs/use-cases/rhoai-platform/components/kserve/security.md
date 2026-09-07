# kserve: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| epp-metrics-token | Opaque | deployment/controller-manager |
| kserve-webhook-server-cert | Opaque | deployment/kserve-controller-manager |
| llmisvc-webhook-server-cert | Opaque | deployment/llmisvc-controller-manager |
| localmodel-webhook-server-cert | Opaque | deployment/kserve-localmodel-controller-manager |
| metrics-server-cert | Opaque | deployment/controller-manager |
| model-serving-api-tls | kubernetes.io/tls | service/model-serving-api |
| odh-model-controller-webhook-cert | kubernetes.io/tls | deployment/odh-model-controller, service/odh-model-controller-webhook-service |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| kserve-controller-manager | manager | true | true | false | [`kustomize:config/overlays/odh`](https://github.com/kserve/kserve/blob/731e4c5ee14219321cd0568cc7dade12d9d3d393/kustomize:config/overlays/odh) |
| kserve-localmodel-controller-manager | manager | true | true | false | [`config/localmodels/manager.yaml`](https://github.com/kserve/kserve/blob/731e4c5ee14219321cd0568cc7dade12d9d3d393/config/localmodels/manager.yaml) |
| kserve-module-controller-manager | manager | true | true | false | [`kserve-module/config/manager.yaml`](https://github.com/kserve/kserve/blob/731e4c5ee14219321cd0568cc7dade12d9d3d393/kserve-module/config/manager.yaml) |
| llmisvc-controller-manager | manager | true | true | false | [`kustomize:config/overlays/odh`](https://github.com/kserve/kserve/blob/731e4c5ee14219321cd0568cc7dade12d9d3d393/kustomize:config/overlays/odh) |
| odh-model-controller | manager | ? | ? | ? | [`kserve-module/prefetched-manifests-rhoai/modelcontroller/default/manager_webhook_patch.yaml`](https://github.com/kserve/kserve/blob/731e4c5ee14219321cd0568cc7dade12d9d3d393/kserve-module/prefetched-manifests-rhoai/modelcontroller/default/manager_webhook_patch.yaml) |
| odh-model-controller | manager | ? | ? | ? | [`kserve-module/prefetched-manifests-rhoai/modelcontroller/manager/manager.yaml`](https://github.com/kserve/kserve/blob/731e4c5ee14219321cd0568cc7dade12d9d3d393/kserve-module/prefetched-manifests-rhoai/modelcontroller/manager/manager.yaml) |
| odh-model-controller | manager | ? | ? | ? | [`kserve-module/prefetched-manifests-rhoai/modelcontroller/overlays/xks/manager_xks_env_patch.yaml`](https://github.com/kserve/kserve/blob/731e4c5ee14219321cd0568cc7dade12d9d3d393/kserve-module/prefetched-manifests-rhoai/modelcontroller/overlays/xks/manager_xks_env_patch.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 4 | 1000:1000 |  |  |  | Unpinned base image: deps; Unpinned base image: deps; Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `Dockerfiles/Dockerfile.konflux.kserve-module-controller` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:3e009398a8aa8eec621393fbf308c5e622f174900e44e8d5fe224c637920924a | 2 | 1000:1000 |  | multi-arch |  |  |
| `agent.Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 4 | 1000:1000 |  |  |  | Unpinned base image: deps; Unpinned base image: deps; Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `dev_tools/kserve-devspace-debug.Containerfile` | golang:1.24-bookworm | 1 | vscode |  |  |  |  |
| `dev_tools/kserve-devspace-dev.Containerfile` | golang:1.24-bookworm | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `kserve-module-controller.Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 1000:1000 |  |  |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `llmisvc-controller.Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 4 | 0 |  |  |  | Unpinned base image: deps; Unpinned base image: deps; Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest; Container runs as root user |
| `localmodel-agent.Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 4 | 1000:1000 |  |  |  | Unpinned base image: deps; Unpinned base image: deps; Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `localmodel.Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 4 | 1000:1000 |  |  |  | Unpinned base image: deps; Unpinned base image: deps; Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `python/aiffairness.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/artexplainer.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/autogluon.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/custom_model.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/custom_model_grpc.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/custom_tokenizer.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/custom_transformer.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/custom_transformer_grpc.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/error_404_isvc.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/huggingface_server.Dockerfile` | cuda-runtime | 6 | 1000 |  |  |  | Unpinned base image: cuda-base; Unpinned base image: cuda-runtime; Unpinned base image: cuda-devel; Unpinned base image: base; Unpinned base image: cuda-runtime |
| `python/huggingface_server_cpu.Dockerfile` | base | 3 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: base; Unpinned base image: base |
| `python/lgb.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/paddle.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/pmml.Dockerfile` | ${BASE_IMAGE} | 2 | 1001 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/predictiveserver.Dockerfile` | ${BASE_IMAGE} | 2 | 1001 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/sklearn.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/storage-initializer.Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 1000 |  |  |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest; Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `python/success_200_isvc.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/xgb.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `qpext/qpext.Dockerfile` | gcr.io/distroless/static:nonroot | 4 |  |  |  |  | Unpinned base image: deps; Unpinned base image: deps; No USER directive found (defaults to root) |
| `router.Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 4 | 1000:1000 |  |  |  | Unpinned base image: deps; Unpinned base image: deps; Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `tools/tf2openapi/Dockerfile` | gcr.io/distroless/static:nonroot | 2 |  |  |  |  | No USER directive found (defaults to root) |

