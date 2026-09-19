# ai-gateway-payload-processing: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:9.8@sha256:463cae32c6f6f5594b11a5c22de275016bd8545ce58a6373388e8b24f13fc15c | 2 | 1001 |  | multi-arch |  |  |
| `Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:fac0bde5012de6a8475bb5dc18db968cdc90b12d9dc2ffaddb4a4975d0b8a04f | 2 | 1001 |  | multi-arch |  |  |
| `Dockerfile.konflux.e2e` | registry.access.redhat.com/ubi9/ubi-minimal@sha256:7b8e25a1b56ca4d00219198f3b5b51a3e1693a5c4f5369c5e190d7d6cb3f980e | 3 | 1001 |  | multi-arch |  | Unpinned base image: registry.redhat.io/openshift4/ose-cli:latest |

