# ogx-k8s-operator: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| ogx-k8s-operator-controller-manager-metrics-service-tls | Opaque | deployment/controller-manager |
| ogx-k8s-operator-webhook-cert | Opaque | deployment/controller-manager |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| ogx-k8s-operator-controller-manager | manager | ? | true | ? | [`kustomize:config/overlays/odh`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/kustomize:config/overlays/odh) |
| operator | manager | ? | true | ? | [`ogx-module/config/manager/manager.yaml`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/ogx-module/config/manager/manager.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 1001 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:3e009398a8aa8eec621393fbf308c5e622f174900e44e8d5fe224c637920924a | 2 | 1001 |  | multi-arch |  |  |
| `ogx-module/Dockerfile` | registry.access.redhat.com/ubi9/ubi-micro:latest | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-micro:latest |
| `ogx-module/Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:3e009398a8aa8eec621393fbf308c5e622f174900e44e8d5fe224c637920924a | 2 | 1001 |  | multi-arch |  |  |

