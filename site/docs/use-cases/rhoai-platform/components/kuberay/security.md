# kuberay: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| webhook-server-cert | Opaque | deployment/kuberay-operator |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| kuberay-operator | kuberay-operator | ? | ? | ? | [`ray-operator/config/overlays/kubernetes-was-v1alpha2/deployment-override.yaml`](https://github.com/red-hat-data-services/kuberay/blob/4cb669708e722b664a9fc91635a074da2752abad/ray-operator/config/overlays/kubernetes-was-v1alpha2/deployment-override.yaml) |
| kuberay-operator | kuberay-operator | ? | ? | ? | [`ray-operator/config/overlays/test-overrides/deployment-override.yaml`](https://github.com/red-hat-data-services/kuberay/blob/4cb669708e722b664a9fc91635a074da2752abad/ray-operator/config/overlays/test-overrides/deployment-override.yaml) |
| kuberay-operator | kuberay-operator | ? | ? | ? | [`ray-operator/config/default-with-webhooks/manager_webhook_patch.yaml`](https://github.com/red-hat-data-services/kuberay/blob/4cb669708e722b664a9fc91635a074da2752abad/ray-operator/config/default-with-webhooks/manager_webhook_patch.yaml) |
| kuberay-operator | kuberay-operator | ? | true | ? | [`ray-operator/config/manager/manager.yaml`](https://github.com/red-hat-data-services/kuberay/blob/4cb669708e722b664a9fc91635a074da2752abad/ray-operator/config/manager/manager.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `apiserver/Dockerfile` | scratch | 2 | 65532:65532 |  |  |  | Unpinned base image: scratch |
| `apiserver/Dockerfile.buildx` | scratch | 1 | 65532:65532 |  | multi-arch |  | Unpinned base image: scratch |
| `benchmark/perf-tests/images/ray-pytorch/Dockerfile` | rayproject/ray:2.46.0 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `dashboard/Dockerfile` | base | 4 | nextjs |  |  |  | Unpinned base image: base; Unpinned base image: base; Unpinned base image: base |
| `historyserver/Dockerfile.collector` | gcr.io/distroless/base-debian12:debug-nonroot | 2 | 1000:1000 |  |  |  |  |
| `historyserver/Dockerfile.collector.buildx` | gcr.io/distroless/base-debian12:nonroot | 1 | 1000:1000 |  | multi-arch |  |  |
| `historyserver/Dockerfile.historyserver` | gcr.io/distroless/base-debian12:nonroot | 2 | 65532:65532 |  |  |  |  |
| `historyserver/Dockerfile.historyserver.buildx` | gcr.io/distroless/base-debian12:nonroot | 1 | 65532:65532 |  | multi-arch |  |  |
| `historyserver/cmd/collector/Dockerfile` |  | 0 |  |  |  |  | No USER directive found (defaults to root) |
| `historyserver/cmd/historyserver/Dockerfile` |  | 0 |  |  |  |  | No USER directive found (defaults to root) |
| `proto/Dockerfile` | golang:1.26-bookworm | 1 | 65532:65532 |  |  |  |  |
| `ray-operator/Dockerfile` | gcr.io/distroless/base-debian12:nonroot | 3 | 65532:65532 |  |  |  | Unpinned base image: scratch |
| `ray-operator/Dockerfile.buildx` | gcr.io/distroless/base-debian12:nonroot | 1 | 65532:65532 |  | multi-arch |  |  |
| `ray-operator/Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:58d09489b1d9123e24d691e47af6c390fe5e08f50c7f4f2967bb64df7ab72aa0 | 2 | 65532:65532 |  | multi-arch |  |  |
| `ray-operator/Dockerfile.rhoai` | registry.access.redhat.com/ubi9/ubi:latest | 2 | 65532:65532 |  |  |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi:latest |
| `ray-operator/Dockerfile.submitter.buildx` | scratch | 1 | 65532:65532 |  | multi-arch |  | Unpinned base image: scratch |

