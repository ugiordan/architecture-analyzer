# trustyai-service-operator: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| webhook-server-cert | kubernetes.io/tls | deployment/controller-manager, service/webhook-service |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| trustyai-service-operator-controller-manager | manager | true | ? | ? | [`kustomize:config/overlays/odh`](https://github.com/trustyai-explainability/trustyai-service-operator/blob/870559ac1034accf95ac655072444bf28d36ca19/kustomize:config/overlays/odh) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi8/ubi-minimal:latest | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi8/ubi-minimal:latest |
| `Dockerfile.driver` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 65532:65532 |  |  |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:8a842ac769de709143e4edeace516f2008dfdc431b64670ad3353fa323b44736 | 2 | 65532:65532 |  | multi-arch |  |  |
| `Dockerfile.konflux.driver` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:8a842ac769de709143e4edeace516f2008dfdc431b64670ad3353fa323b44736 | 2 | 65532:65532 |  |  |  |  |
| `Dockerfile.lmes-job` | registry.access.redhat.com/ubi9/python-311@sha256:608902aba04dd1210be6db6c3b4e288d8478f978737b677f28d7d48452611506 | 1 | 65532:65532 |  |  |  |  |
| `Dockerfile.orchestrator` | ${UBI_MINIMAL_BASE_IMAGE}:${UBI_BASE_IMAGE_TAG} | 6 | orchestr8 |  |  |  | Unpinned base image: rust-builder; Unpinned base image: fms-guardrails-orchestr8-builder; Unpinned base image: fms-guardrails-orchestr8-builder; Unpinned base image: fms-guardrails-orchestr8-builder |
| `trustyai-operator-module/Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 1000 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |

