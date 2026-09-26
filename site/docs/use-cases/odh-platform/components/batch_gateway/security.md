# batch-gateway: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `docker/Dockerfile.apiserver` | gcr.io/distroless/static:nonroot | 2 | 65532:65532 |  | multi-arch |  |  |
| `docker/Dockerfile.apiserver.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:b280df203f9ec5d03439ae79cc3f26a01dc04f01c9259853398be438a334ade2 | 2 | 1001:1001 |  | multi-arch |  |  |
| `docker/Dockerfile.gc` | gcr.io/distroless/static:nonroot | 2 | 65532:65532 |  | multi-arch |  |  |
| `docker/Dockerfile.gc.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:b280df203f9ec5d03439ae79cc3f26a01dc04f01c9259853398be438a334ade2 | 2 | 1001:1001 |  | multi-arch |  |  |
| `docker/Dockerfile.processor` | gcr.io/distroless/static:nonroot | 2 | 65532:65532 |  | multi-arch |  |  |
| `docker/Dockerfile.processor.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:b280df203f9ec5d03439ae79cc3f26a01dc04f01c9259853398be438a334ade2 | 2 | 1001:1001 |  | multi-arch |  |  |

