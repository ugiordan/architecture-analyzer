# trainer: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| kubeflow-trainer-webhook-cert | Opaque | deployment/kubeflow-trainer-controller-manager |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| kubeflow-trainer-controller-manager | manager | true | true | ? | [`manifests/base/manager/manager.yaml`](https://github.com/kubeflow/trainer/blob/6d6b623b4f54315b09d597b438e2daf911e20fb2/manifests/base/manager/manager.yaml) |
| kubeflow-trainer-controller-manager | manager | ? | ? | ? | [`manifests/base/manager/manager_config_patch.yaml`](https://github.com/kubeflow/trainer/blob/6d6b623b4f54315b09d597b438e2daf911e20fb2/manifests/base/manager/manager_config_patch.yaml) |
| kubeflow-trainer-controller-manager | manager | ? | ? | ? | [`manifests/rhoai/manager_config_patch.yaml`](https://github.com/kubeflow/trainer/blob/6d6b623b4f54315b09d597b438e2daf911e20fb2/manifests/rhoai/manager_config_patch.yaml) |
| kubeflow-trainer-controller-manager | manager | ? | ? | ? | [`manifests/rhoai/manager_metrics_patch.yaml`](https://github.com/kubeflow/trainer/blob/6d6b623b4f54315b09d597b438e2daf911e20fb2/manifests/rhoai/manager_metrics_patch.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `cmd/data_cache/Dockerfile` | debian:bookworm-slim | 2 | cache_user |  |  |  |  |
| `cmd/initializers/dataset/Dockerfile` | python:3.14-slim-bookworm | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `cmd/initializers/model/Dockerfile` | python:3.14-slim-bookworm | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `cmd/runtimes/deepspeed/Dockerfile` | nvidia/cuda:13.3.0-devel-ubuntu22.04 | 2 |  |  |  |  | No USER directive found (defaults to root) |
| `cmd/runtimes/mlx/Dockerfile` | nvidia/cuda:13.3.0-devel-ubuntu22.04 | 2 |  |  |  |  | No USER directive found (defaults to root) |
| `cmd/runtimes/xgboost/Dockerfile` | nvidia/cuda:12.4.0-runtime-ubuntu22.04 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `cmd/trainer-controller-manager/Dockerfile` | gcr.io/distroless/static:nonroot | 2 |  |  |  |  | No USER directive found (defaults to root) |
| `cmd/trainer-controller-manager/Dockerfile.odh` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 65532:65532 |  |  |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `cmd/trainer-controller-manager/Dockerfile.rhoai.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:3e009398a8aa8eec621393fbf308c5e622f174900e44e8d5fe224c637920924a | 2 | 65532:65532 |  |  |  |  |
| `cmd/trainers/torchtune/Dockerfile` | pytorch/pytorch:2.9.1-cuda12.8-cudnn9-runtime | 1 |  |  |  |  | No USER directive found (defaults to root) |

