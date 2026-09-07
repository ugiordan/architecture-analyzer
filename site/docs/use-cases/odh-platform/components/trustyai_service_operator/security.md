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
| trustyai-service-operator-controller-manager | manager | true | ? | ? | [`kustomize:config/overlays/odh`](https://github.com/trustyai-explainability/trustyai-service-operator/blob/285d221d65b92c891a5c47baa7c48662ab9c6bd5/kustomize:config/overlays/odh) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi8/ubi-minimal:latest | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi8/ubi-minimal:latest |
| `Dockerfile.driver` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 65532:65532 |  |  |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `Dockerfile.lmes-job` | registry.access.redhat.com/ubi9/python-311@sha256:fccda5088dd13d2a3f2659e4c904beb42fc164a0c909e765f01af31c58affae3 | 1 | 65532:65532 |  |  |  |  |
| `Dockerfile.orchestrator` | ${UBI_MINIMAL_BASE_IMAGE}:${UBI_BASE_IMAGE_TAG} | 6 | orchestr8 |  |  |  | Unpinned base image: rust-builder; Unpinned base image: fms-guardrails-orchestr8-builder; Unpinned base image: fms-guardrails-orchestr8-builder; Unpinned base image: fms-guardrails-orchestr8-builder |
| `trustyai-operator-module/Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 1000 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |

