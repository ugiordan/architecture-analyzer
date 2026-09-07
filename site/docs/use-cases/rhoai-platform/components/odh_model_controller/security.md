# odh-model-controller: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| model-serving-api-tls | kubernetes.io/tls | service/model-serving-api |
| odh-model-controller-webhook-cert | kubernetes.io/tls | deployment/odh-model-controller, service/odh-model-controller-webhook-service |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| model-serving-api | server | ? | true | ? | [`kustomize:config/overlays/odh`](https://github.com/opendatahub-io/odh-model-controller/blob/5fe45b407f4f72f32b2d14497c9311d331bc8772/kustomize:config/overlays/odh) |
| odh-model-controller | manager | ? | ? | ? | [`kustomize:config/overlays/odh`](https://github.com/opendatahub-io/odh-model-controller/blob/5fe45b407f4f72f32b2d14497c9311d331bc8772/kustomize:config/overlays/odh) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Containerfile` | registry.access.redhat.com/ubi9/ubi-minimal:9.5 | 2 | 65532:65532 |  | multi-arch |  |  |
| `Containerfile.server` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 1000:1000 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `Containerfile.server.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:8a842ac769de709143e4edeace516f2008dfdc431b64670ad3353fa323b44736 | 2 | ${USER} |  | multi-arch |  |  |
| `Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:3e009398a8aa8eec621393fbf308c5e622f174900e44e8d5fe224c637920924a | 2 | ${USER} |  | multi-arch |  |  |
| `dev_tools/Containerfile.devspace` | registry.access.redhat.com/ubi9/go-toolset:1.25.8 | 1 | root |  |  |  | Container runs as root user |

