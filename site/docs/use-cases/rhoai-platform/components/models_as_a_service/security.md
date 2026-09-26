# models-as-a-service: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| maas-api-metrics-tls | Opaque | deployment/maas-api |
| maas-api-serving-cert | kubernetes.io/tls | deployment/maas-api, service/maas-api |
| maas-controller-metrics-tls | Opaque | deployment/maas-controller |
| maas-controller-webhook-cert | kubernetes.io/tls | deployment/maas-controller, service/maas-controller-webhook-service |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| maas-api | maas-api | true | true | ? | [`deployment/base/maas-api/core/deployment.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/4a9abc9f3969b97de5ce73ca037c19d84126c43d/deployment/base/maas-api/core/deployment.yaml) |
| maas-api | maas-api | ? | ? | ? | [`deployment/base/maas-api/overlays/tls/deployment-patch.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/4a9abc9f3969b97de5ce73ca037c19d84126c43d/deployment/base/maas-api/overlays/tls/deployment-patch.yaml) |
| maas-controller | manager | ? | true | ? | [`deployment/base/maas-controller/manager/manager.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/4a9abc9f3969b97de5ce73ca037c19d84126c43d/deployment/base/maas-controller/manager/manager.yaml) |
| payload-processing | payload-processing | ? | true | ? | [`deployment/base/payload-processing/instance/deployment.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/4a9abc9f3969b97de5ce73ca037c19d84126c43d/deployment/base/payload-processing/instance/deployment.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `maas-api/Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 1001 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `maas-api/Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:22c70d6d7fa920df140c5facdce1f42178642ca7613590da5466f69418ddf57a | 2 | 1001 |  | multi-arch |  |  |
| `maas-controller/Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 1001 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `maas-controller/Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:22c70d6d7fa920df140c5facdce1f42178642ca7613590da5466f69418ddf57a | 2 | 1001 |  | multi-arch |  |  |
| `maas-discovery/Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 1001 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `maas-discovery/Dockerfile.konflux` | registry.access.redhat.com/ubi9/ubi-minimal@sha256:9c6d058e39cdbd97c9e0b83c5c394490869dfe96e8f9b3f0fbde0a08f4220adf | 2 | 1001 |  | multi-arch |  |  |

