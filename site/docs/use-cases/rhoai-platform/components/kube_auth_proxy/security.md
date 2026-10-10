# kube-auth-proxy: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | ${RUNTIME_IMAGE} | 2 |  |  | multi-arch |  | Unpinned base image: ${BUILD_IMAGE}; Unpinned base image: ${RUNTIME_IMAGE}; No USER directive found (defaults to root) |
| `Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:d34666200b64d5a183c867487fd2f48ee8547eb18134e5647099970e26f08af9 | 2 | 1001 |  | multi-arch |  |  |
| `Dockerfile.redhat` | registry.access.redhat.com/ubi9/ubi-minimal:9.8@sha256:beeada7dd17903dfb69fd5f6916c054720bf28a52daaa2f7a1910a1394244bd2 | 2 | 1001 |  | multi-arch |  |  |

