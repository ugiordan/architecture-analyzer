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
| `docker/Dockerfile.apiserver.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:b3a0b012de0c578082e1bc5b01c8bc27461b4eeb571a0b1a2b21a07dbd845a62 | 2 | 1001:1001 |  | multi-arch |  |  |
| `docker/Dockerfile.gc` | gcr.io/distroless/static:nonroot | 2 | 65532:65532 |  | multi-arch |  |  |
| `docker/Dockerfile.gc.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:b3a0b012de0c578082e1bc5b01c8bc27461b4eeb571a0b1a2b21a07dbd845a62 | 2 | 1001:1001 |  | multi-arch |  |  |
| `docker/Dockerfile.processor` | gcr.io/distroless/static:nonroot | 2 | 65532:65532 |  | multi-arch |  |  |
| `docker/Dockerfile.processor.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:b3a0b012de0c578082e1bc5b01c8bc27461b4eeb571a0b1a2b21a07dbd845a62 | 2 | 1001:1001 |  | multi-arch |  |  |

