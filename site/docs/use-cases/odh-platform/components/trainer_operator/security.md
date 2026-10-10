# trainer-operator: Security

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
| controller-manager | manager | ? | true | ? | [`config/manager/manager.yaml`](https://github.com/opendatahub-io/trainer-operator/blob/90322983edf391c7adb195654221cc83d172b2a3/config/manager/manager.yaml) |
| kubeflow-trainer-controller-manager | manager | true | true | ? | [`manifests/trainer/base/manager/manager.yaml`](https://github.com/opendatahub-io/trainer-operator/blob/90322983edf391c7adb195654221cc83d172b2a3/manifests/trainer/base/manager/manager.yaml) |
| kubeflow-trainer-controller-manager | manager | ? | ? | ? | [`manifests/trainer/base/manager/manager_config_patch.yaml`](https://github.com/opendatahub-io/trainer-operator/blob/90322983edf391c7adb195654221cc83d172b2a3/manifests/trainer/base/manager/manager_config_patch.yaml) |
| kubeflow-trainer-controller-manager | manager | ? | ? | ? | [`manifests/trainer/rhoai/manager_config_patch.yaml`](https://github.com/opendatahub-io/trainer-operator/blob/90322983edf391c7adb195654221cc83d172b2a3/manifests/trainer/rhoai/manager_config_patch.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:9.8 | 2 | 65532:65532 |  |  |  |  |

