# llm-d-inference-scheduler: Security

[Overview](index.md) · [Agent Context](agent-context.md) · [Extracted Data](details.md) · [Network](network.md) · [RBAC](rbac.md) · [Security](security.md) · [Dataflow](dataflow.md)

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile.builder` | golang:1.27.4 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `Dockerfile.coordinator` | ${BASE_IMAGE} | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE} |
| `Dockerfile.epp` | ${BASE_IMAGE} | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE} |
| `Dockerfile.konflux.coordinator` | registry.access.redhat.com/ubi10/ubi-minimal:10.2@sha256:f0c7b9ae9914fcf183e7c36fc46864903c2d2caf3af327421913ae54b3d29e94 | 2 | 65532:65532 |  | multi-arch |  |  |
| `Dockerfile.konflux.epp` | registry.access.redhat.com/ubi10/ubi-minimal:10.2@sha256:f0c7b9ae9914fcf183e7c36fc46864903c2d2caf3af327421913ae54b3d29e94 | 2 | 65532:65532 |  | multi-arch |  |  |
| `Dockerfile.konflux.sidecar` | registry.access.redhat.com/ubi10/ubi-minimal:10.2@sha256:f0c7b9ae9914fcf183e7c36fc46864903c2d2caf3af327421913ae54b3d29e94 | 2 | 65532:65532 |  | multi-arch |  |  |
| `Dockerfile.sidecar` | ${BASE_IMAGE} | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE} |

