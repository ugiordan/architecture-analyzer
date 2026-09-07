# odh-dashboard: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| agent-ops-proxy-tls | kubernetes.io/tls | deployment/agent-ops-ui, service/odh-dashboard-agent-ops-ui |
| automl-proxy-tls | kubernetes.io/tls | deployment/automl-ui, service/odh-dashboard-automl-ui |
| autorag-proxy-tls | kubernetes.io/tls | deployment/autorag-ui, service/odh-dashboard-autorag-ui |
| dashboard-operator-metrics-tls | Opaque | deployment/dashboard-operator |
| dashboard-operator-webhook-tls | Opaque | deployment/dashboard-operator |
| dashboard-proxy-tls | kubernetes.io/tls | deployment/odh-dashboard, service/odh-dashboard |
| data-registry-proxy-tls | kubernetes.io/tls | deployment/data-registry-ui, service/odh-dashboard-data-registry-ui |
| eval-hub-proxy-tls | kubernetes.io/tls | deployment/eval-hub-ui, service/odh-dashboard-eval-hub-ui |
| gen-ai-proxy-tls | kubernetes.io/tls | deployment/gen-ai-ui, service/odh-dashboard-gen-ai-ui |
| maas-proxy-tls | kubernetes.io/tls | deployment/maas-ui, service/odh-dashboard-maas-ui |
| mlflow-proxy-tls | kubernetes.io/tls | deployment/mlflow-ui, service/odh-dashboard-mlflow-ui |
| model-registry-proxy-tls | kubernetes.io/tls | deployment/model-registry-ui, service/odh-dashboard-model-registry-ui |
| notebooks-proxy-tls | kubernetes.io/tls | deployment/notebooks-ui, service/odh-dashboard-notebooks-ui |
| webhook-server-cert | Opaque | deployment/workspaces-controller |
| workspaces-controller-webhook-cert | Opaque | deployment/workspaces-controller |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| agent-ops-ui | agent-ops-ui | true | true | ? | [`manifests/modules/agent-ops/deployment.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/manifests/modules/agent-ops/deployment.yaml) |
| automl-ui | automl-ui | true | true | ? | [`manifests/modules/automl/deployment.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/manifests/modules/automl/deployment.yaml) |
| autorag-ui | autorag-ui | true | true | ? | [`manifests/modules/autorag/deployment.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/manifests/modules/autorag/deployment.yaml) |
| dashboard-operator | manager | ? | ? | ? | [`dashboard-operator/config/manager/manager.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/dashboard-operator/config/manager/manager.yaml) |
| data-registry-ui | data-registry-ui | true | true | ? | [`manifests/modules/data-registry/deployment.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/manifests/modules/data-registry/deployment.yaml) |
| eval-hub-ui | eval-hub-ui | true | true | ? | [`manifests/modules/eval-hub/deployment.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/manifests/modules/eval-hub/deployment.yaml) |
| gen-ai-ui | gen-ai-ui | true | true | ? | [`manifests/modules/gen-ai/deployment.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/manifests/modules/gen-ai/deployment.yaml) |
| maas-ui | maas-ui | true | true | ? | [`manifests/modules/maas/deployment.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/manifests/modules/maas/deployment.yaml) |
| mlflow-ui | mlflow-ui | true | true | ? | [`manifests/modules/mlflow/deployment.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/manifests/modules/mlflow/deployment.yaml) |
| model-registry-ui | model-registry-ui | true | true | ? | [`manifests/modules/model-registry/deployment.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/manifests/modules/model-registry/deployment.yaml) |
| notebooks-ui | notebooks-ui | true | true | ? | [`manifests/modules/notebooks/deployment.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/manifests/modules/notebooks/deployment.yaml) |
| odh-dashboard | odh-dashboard | ? | ? | ? | [`manifests/base/deployment.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/manifests/base/deployment.yaml) |
| odh-dashboard | kube-rbac-proxy | ? | ? | ? | [`manifests/base/deployment.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/manifests/base/deployment.yaml) |
| odh-dashboard | core-bff | true | true | ? | [`manifests/base/deployment.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/manifests/base/deployment.yaml) |
| rhaii-dashboard | rhaii-dashboard | true | true | ? | [`distributions/core-bff/manifests/base/deployment.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/distributions/core-bff/manifests/base/deployment.yaml) |
| workspaces-backend | workspaces-backend | ? | ? | ? | [`packages/notebooks/upstream/workspaces/backend/manifests/kustomize/base/deployment.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/packages/notebooks/upstream/workspaces/backend/manifests/kustomize/base/deployment.yaml) |
| workspaces-backend | workspaces-backend | ? | ? | ? | [`packages/notebooks/upstream/workspaces/backend/manifests/kustomize/components/istio/deployment_patch.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/packages/notebooks/upstream/workspaces/backend/manifests/kustomize/components/istio/deployment_patch.yaml) |
| workspaces-controller | manager | ? | ? | ? | [`packages/notebooks/upstream/workspaces/controller/manifests/kustomize/components/certmanager/deployment_patch.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/packages/notebooks/upstream/workspaces/controller/manifests/kustomize/components/certmanager/deployment_patch.yaml) |
| workspaces-controller | manager | ? | ? | ? | [`packages/notebooks/upstream/workspaces/controller/manifests/kustomize/components/istio/deployment_patch.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/packages/notebooks/upstream/workspaces/controller/manifests/kustomize/components/istio/deployment_patch.yaml) |
| workspaces-controller | manager | ? | ? | ? | [`packages/notebooks/upstream/workspaces/controller/manifests/kustomize/components/openshift-webhook/deployment_patch.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/packages/notebooks/upstream/workspaces/controller/manifests/kustomize/components/openshift-webhook/deployment_patch.yaml) |
| workspaces-controller | manager | ? | ? | ? | [`packages/notebooks/upstream/workspaces/controller/manifests/kustomize/components/prometheus/deployment_patch.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/packages/notebooks/upstream/workspaces/controller/manifests/kustomize/components/prometheus/deployment_patch.yaml) |
| workspaces-controller | manager | true | ? | ? | [`packages/notebooks/upstream/workspaces/controller/manifests/kustomize/base/manager/manager.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/packages/notebooks/upstream/workspaces/controller/manifests/kustomize/base/manager/manager.yaml) |
| workspaces-controller | manager | ? | ? | ? | [`packages/notebooks/upstream/workspaces/controller/manifests/kustomize/components/gateway/manager_gateway_patch.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/packages/notebooks/upstream/workspaces/controller/manifests/kustomize/components/gateway/manager_gateway_patch.yaml) |
| workspaces-frontend | workspaces-frontend | ? | ? | ? | [`packages/notebooks/upstream/workspaces/frontend/manifests/kustomize/base/deployment.yaml`](https://github.com/opendatahub-io/odh-dashboard/blob/2d6770547c827f86b1f99d466b78aca5a899a6cc/packages/notebooks/upstream/workspaces/frontend/manifests/kustomize/base/deployment.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | ${MINIMAL_IMAGE} | 2 | 1001:0 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${MINIMAL_IMAGE} |
| `dashboard-operator/Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:9.3 | 2 | 65532:65532 |  | multi-arch |  |  |
| `dashboard-operator/Dockerfile.dockerignore` |  | 0 |  |  |  |  | No USER directive found (defaults to root) |
| `distributions/core-bff/Dockerfile` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `distributions/core-bff/Dockerfile.workspace` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `distributions/rhaii/Dockerfile` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/agent-ops/Dockerfile` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/agent-ops/Dockerfile.workspace` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/automl/Dockerfile` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/automl/Dockerfile.workspace` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/autorag/Dockerfile` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/autorag/Dockerfile.workspace` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/data-registry/Dockerfile` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/data-registry/Dockerfile.workspace` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/eval-hub/Dockerfile` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/eval-hub/Dockerfile.workspace` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/gen-ai/Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:9.3 | 3 | 1001 |  |  |  | Unpinned base image: registry.access.redhat.com/ubi9/nodejs-20 |
| `packages/gen-ai/Dockerfile.workspace` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/gen-ai/bff/Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:9.3 | 2 | 1001 |  |  |  |  |
| `packages/maas/Dockerfile` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/maas/Dockerfile.workspace` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/mlflow/Dockerfile` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/mlflow/Dockerfile.workspace` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/model-registry/Dockerfile.workspace` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/model-registry/upstream/Dockerfile` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/model-registry/upstream/Dockerfile.standalone` | release | 4 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE}; Unpinned base image: release |
| `packages/notebooks/Dockerfile.workspace` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `packages/notebooks/upstream/workspaces/backend/Dockerfile` | gcr.io/distroless/static:nonroot | 2 | 65532:65532 |  | multi-arch |  |  |
| `packages/notebooks/upstream/workspaces/controller/Dockerfile` | gcr.io/distroless/static:nonroot | 2 | 65532:65532 |  | multi-arch |  |  |
| `packages/notebooks/upstream/workspaces/frontend/Dockerfile` | nginx:alpine | 2 | 101:101 |  |  |  |  |
| `packages/notebooks/upstream/workspaces/frontend/Dockerfile.dev` | node:20-slim | 1 | 1000:1000 |  |  |  |  |
| `packages/plugin-template/Dockerfile.workspace` | ${DISTROLESS_BASE_IMAGE} | 3 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${NODE_BASE_IMAGE}; Unpinned base image: ${GOLANG_BASE_IMAGE}; Unpinned base image: ${DISTROLESS_BASE_IMAGE} |
| `scripts/ci/Dockerfile` | quay.io/fedora/fedora:43 | 1 | $USER |  |  |  |  |

