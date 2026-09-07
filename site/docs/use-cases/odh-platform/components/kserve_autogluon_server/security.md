# kserve-autogluon-server: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| kserve-webhook-server-cert | Opaque | deployment/kserve-controller-manager |
| llmisvc-webhook-server-cert | Opaque | deployment/llmisvc-controller-manager |
| localmodel-webhook-server-cert | Opaque | deployment/kserve-localmodel-controller-manager |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| kserve-controller-manager | manager | true | true | false | [`kustomize:config/overlays/all`](https://github.com/kserve/kserve-autogluon-server/blob/aad6612fe382c04130a0675500087545cb6edc3f/kustomize:config/overlays/all) |
| kserve-controller-manager | kube-rbac-proxy | true | true | false | [`kustomize:config/overlays/all`](https://github.com/kserve/kserve-autogluon-server/blob/aad6612fe382c04130a0675500087545cb6edc3f/kustomize:config/overlays/all) |
| kserve-localmodel-controller-manager | manager | true | true | false | [`kustomize:config/overlays/all`](https://github.com/kserve/kserve-autogluon-server/blob/aad6612fe382c04130a0675500087545cb6edc3f/kustomize:config/overlays/all) |
| llmisvc-controller-manager | manager | true | true | false | [`kustomize:config/overlays/all`](https://github.com/kserve/kserve-autogluon-server/blob/aad6612fe382c04130a0675500087545cb6edc3f/kustomize:config/overlays/all) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | gcr.io/distroless/static:nonroot | 4 |  |  |  |  | Unpinned base image: deps; Unpinned base image: deps; No USER directive found (defaults to root) |
| `agent.Dockerfile` | gcr.io/distroless/static:nonroot | 4 |  |  |  |  | Unpinned base image: deps; Unpinned base image: deps; No USER directive found (defaults to root) |
| `llmisvc-controller.Dockerfile` | gcr.io/distroless/static:nonroot | 4 |  |  |  |  | Unpinned base image: deps; Unpinned base image: deps; No USER directive found (defaults to root) |
| `localmodel-agent.Dockerfile` | gcr.io/distroless/static:nonroot | 4 |  |  |  |  | Unpinned base image: deps; Unpinned base image: deps; No USER directive found (defaults to root) |
| `localmodel.Dockerfile` | gcr.io/distroless/static:nonroot | 4 |  |  |  |  | Unpinned base image: deps; Unpinned base image: deps; No USER directive found (defaults to root) |
| `python/aiffairness.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/artexplainer.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/autogluon.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/custom_model.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/custom_model_grpc.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/custom_tokenizer.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/custom_transformer.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/custom_transformer_grpc.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/error_404_isvc.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/huggingface_server.Dockerfile` | cuda-runtime | 5 | 1000 |  |  |  | Unpinned base image: cuda-runtime; Unpinned base image: cuda-devel; Unpinned base image: base; Unpinned base image: cuda-runtime |
| `python/huggingface_server_cpu.Dockerfile` | base | 3 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: base; Unpinned base image: base |
| `python/lgb.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/paddle.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/pmml.Dockerfile` | ${BASE_IMAGE} | 2 | 1001 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/predictiveserver.Dockerfile` | ${BASE_IMAGE} | 2 | 1001 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/sklearn.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/storage-initializer.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/success_200_isvc.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `python/xgb.Dockerfile` | ${BASE_IMAGE} | 2 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `qpext/qpext.Dockerfile` | gcr.io/distroless/static:nonroot | 4 |  |  |  |  | Unpinned base image: deps; Unpinned base image: deps; No USER directive found (defaults to root) |
| `router.Dockerfile` | gcr.io/distroless/static:nonroot | 4 |  |  |  |  | Unpinned base image: deps; Unpinned base image: deps; No USER directive found (defaults to root) |
| `tools/tf2openapi/Dockerfile` | gcr.io/distroless/static:nonroot | 2 |  |  |  |  | No USER directive found (defaults to root) |

