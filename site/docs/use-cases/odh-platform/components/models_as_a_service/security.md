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
| maas-discovery-tls | Opaque | deployment/maas-discovery |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| maas-api | maas-api | true | true | ? | [`deployment/base/maas-api/core/deployment.yaml`](https://github.com/opendatahub-io/models-as-a-service/blob/26e3116f6ddbe50a73fe9c31b61fbd6d3c817411/deployment/base/maas-api/core/deployment.yaml) |
| maas-api | maas-api | ? | ? | ? | [`deployment/base/maas-api/overlays/tls/deployment-patch.yaml`](https://github.com/opendatahub-io/models-as-a-service/blob/26e3116f6ddbe50a73fe9c31b61fbd6d3c817411/deployment/base/maas-api/overlays/tls/deployment-patch.yaml) |
| maas-controller | manager | ? | true | ? | [`deployment/base/maas-controller/manager/manager.yaml`](https://github.com/opendatahub-io/models-as-a-service/blob/26e3116f6ddbe50a73fe9c31b61fbd6d3c817411/deployment/base/maas-controller/manager/manager.yaml) |
| maas-discovery | maas-discovery | true | true | ? | [`deployment/base/maas-discovery/core/deployment.yaml`](https://github.com/opendatahub-io/models-as-a-service/blob/26e3116f6ddbe50a73fe9c31b61fbd6d3c817411/deployment/base/maas-discovery/core/deployment.yaml) |
| payload-processing | payload-processing | ? | true | ? | [`deployment/base/payload-processing/instance/deployment.yaml`](https://github.com/opendatahub-io/models-as-a-service/blob/26e3116f6ddbe50a73fe9c31b61fbd6d3c817411/deployment/base/payload-processing/instance/deployment.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `maas-api/Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 1001 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `maas-api/Dockerfile.konflux` | registry.access.redhat.com/ubi9/ubi-minimal@sha256:80f3902b6dcb47005a90e14140eef9080ccc1bb22df70ee16b27d5891524edb2 | 2 | 1001 |  | multi-arch |  |  |
| `maas-controller/Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 1001 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `maas-controller/Dockerfile.konflux` | registry.access.redhat.com/ubi9/ubi-minimal@sha256:83006d535923fcf1345067873524a3980316f51794f01d8655be55d6e9387183 | 2 | 1001 |  | multi-arch |  |  |
| `maas-discovery/Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 1001 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `maas-discovery/Dockerfile.konflux` | registry.access.redhat.com/ubi9/ubi-minimal@sha256:80f3902b6dcb47005a90e14140eef9080ccc1bb22df70ee16b27d5891524edb2 | 2 | 1001 |  | multi-arch |  |  |

