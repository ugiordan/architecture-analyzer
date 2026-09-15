# rest-proxy: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi9/ubi-micro:9.5 | 3 | ${USER} |  | multi-arch |  | Unpinned base image: develop |
| `Dockerfile.konflux` | registry.access.redhat.com/ubi9/ubi-minimal@sha256:186a94b76e386782f576c9c49813b16dceb2ba63102af5a28405dcefec2806d0 | 2 | ${USER} |  | multi-arch |  |  |

