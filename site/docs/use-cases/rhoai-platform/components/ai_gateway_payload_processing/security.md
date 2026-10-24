# ai-gateway-payload-processing: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal@sha256:8ebe2ad8fdf3cab3e5a53c1edc69194c98209cfadab24b884f4ad9ebcf7bbbfc | 2 | 1001 |  | multi-arch |  |  |
| `Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:d0fdfb1401ee754872f60c2527563671c16112080e5254139133908178700911 | 2 | 1001 |  | multi-arch |  |  |
| `Dockerfile.konflux.e2e` | registry.access.redhat.com/ubi9/ubi-minimal@sha256:2d4c4551af6831dcd26bbf86108c4385de5183ec8fe6b5868f6f9640fe47f4ed | 3 | 1001 |  | multi-arch |  | Unpinned base image: registry.redhat.io/openshift4/ose-cli:latest |

