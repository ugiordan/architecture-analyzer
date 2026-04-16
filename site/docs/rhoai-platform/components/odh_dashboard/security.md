# odh-dashboard: Security

## Secrets

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| dashboard-proxy-tls | kubernetes.io/tls | deployment/odh-dashboard, service/odh-dashboard |
| webhook-server-cert | Opaque | deployment/workspaces-controller |

## Deployment Security Controls

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| odh-dashboard | odh-dashboard | ? | ? | ? | `manifests/core-bases/base/deployment.yaml` |
| odh-dashboard | kube-rbac-proxy | ? | ? | ? | `manifests/core-bases/base/deployment.yaml` |
| workspaces-backend | workspaces-backend | ? | ? | ? | `packages/notebooks/upstream/workspaces/backend/manifests/kustomize/base/deployment.yaml` |
| workspaces-frontend | workspaces-frontend | ? | ? | ? | `packages/notebooks/upstream/workspaces/frontend/manifests/kustomize/base/deployment.yaml` |
| workspaces-controller | manager | ? | ? | ? | `packages/notebooks/upstream/workspaces/controller/manifests/kustomize/components/certmanager/deployment_patch.yaml` |
| workspaces-controller | manager | ? | ? | ? | `packages/notebooks/upstream/workspaces/controller/manifests/kustomize/components/istio/deployment_patch.yaml` |
| workspaces-controller | manager | ? | ? | ? | `packages/notebooks/upstream/workspaces/controller/manifests/kustomize/components/prometheus/deployment_patch.yaml` |
| workspaces-controller | manager | true | ? | ? | `packages/notebooks/upstream/workspaces/controller/manifests/kustomize/base/manager/manager.yaml` |

## Build Security

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | ${BASE_IMAGE} | 2 | 1001:0 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `node_modules/getos/Dockerfile` | debian:7.4 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `node_modules/getos/tests/alpine/3.3/Dockerfile` | alpine:3.3 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `node_modules/getos/tests/debian/7.3/Dockerfile` | debian:7.3 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `node_modules/getos/tests/debian/7.4/Dockerfile` | debian:7.4 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `node_modules/getos/tests/debian/7.5/Dockerfile` | debian:7.5 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `node_modules/getos/tests/debian/7.6/Dockerfile` | debian:7.6 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `node_modules/getos/tests/fedora/20/Dockerfile` | fedora:20 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `node_modules/getos/tests/ubuntu/13.10/Dockerfile` | ubuntu:13.10 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `node_modules/getos/tests/ubuntu/14.04/Dockerfile` | ubuntu:14.04 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `packages/automl/Dockerfile` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/autorag/Dockerfile` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/eval-hub/Dockerfile` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/gen-ai/Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:9.3 | 3 | 1001 |  |  |  | Unpinned base image: registry.access.redhat.com/ubi9/nodejs-20 |
| `packages/gen-ai/bff/Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:9.3 | 2 | 1001 |  |  |  |  |
| `packages/gen-ai/frontend/node_modules/getos/Dockerfile` | debian:7.4 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `packages/gen-ai/frontend/node_modules/getos/tests/alpine/3.3/Dockerfile` | alpine:3.3 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `packages/gen-ai/frontend/node_modules/getos/tests/debian/7.3/Dockerfile` | debian:7.3 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `packages/gen-ai/frontend/node_modules/getos/tests/debian/7.4/Dockerfile` | debian:7.4 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `packages/gen-ai/frontend/node_modules/getos/tests/debian/7.5/Dockerfile` | debian:7.5 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `packages/gen-ai/frontend/node_modules/getos/tests/debian/7.6/Dockerfile` | debian:7.6 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `packages/gen-ai/frontend/node_modules/getos/tests/fedora/20/Dockerfile` | fedora:20 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `packages/gen-ai/frontend/node_modules/getos/tests/ubuntu/13.10/Dockerfile` | ubuntu:13.10 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `packages/gen-ai/frontend/node_modules/getos/tests/ubuntu/14.04/Dockerfile` | ubuntu:14.04 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `packages/maas/Dockerfile` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/mlflow/Dockerfile` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/model-registry/upstream/Dockerfile` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/model-registry/upstream/frontend/node_modules/getos/Dockerfile` | debian:7.4 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `packages/model-registry/upstream/frontend/node_modules/getos/tests/alpine/3.3/Dockerfile` | alpine:3.3 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `packages/model-registry/upstream/frontend/node_modules/getos/tests/debian/7.3/Dockerfile` | debian:7.3 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `packages/model-registry/upstream/frontend/node_modules/getos/tests/debian/7.4/Dockerfile` | debian:7.4 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `packages/model-registry/upstream/frontend/node_modules/getos/tests/debian/7.5/Dockerfile` | debian:7.5 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `packages/model-registry/upstream/frontend/node_modules/getos/tests/debian/7.6/Dockerfile` | debian:7.6 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `packages/model-registry/upstream/frontend/node_modules/getos/tests/fedora/20/Dockerfile` | fedora:20 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `packages/model-registry/upstream/frontend/node_modules/getos/tests/ubuntu/13.10/Dockerfile` | ubuntu:13.10 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `packages/model-registry/upstream/frontend/node_modules/getos/tests/ubuntu/14.04/Dockerfile` | ubuntu:14.04 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `packages/notebooks/upstream/workspaces/backend/Dockerfile` | gcr.io/distroless/static:nonroot | 2 | 65532:65532 |  | multi-arch |  |  |
| `packages/notebooks/upstream/workspaces/controller/Dockerfile` | gcr.io/distroless/static:nonroot | 2 | 65532:65532 |  | multi-arch |  |  |
| `packages/notebooks/upstream/workspaces/frontend/Dockerfile` | nginx:alpine | 2 | 101:101 |  |  |  |  |
| `scripts/ci/Dockerfile` | quay.io/fedora/fedora:43 | 1 | $USER |  |  |  |  |
| `packages/automl/Dockerfile.workspace` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/autorag/Dockerfile.workspace` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/eval-hub/Dockerfile.workspace` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/gen-ai/Dockerfile.workspace` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/maas/Dockerfile.workspace` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/mlflow/Dockerfile.workspace` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/model-registry/Dockerfile.workspace` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/model-registry/upstream/Dockerfile.standalone` | release | 4 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE}; Unpinned base image: release |
| `packages/notebooks/Dockerfile.workspace` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/notebooks/upstream/workspaces/frontend/Dockerfile.dev` | node:20-slim | 1 | 1000:1000 |  |  |  |  |
| `packages/plugin-template/Dockerfile.workspace` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |

