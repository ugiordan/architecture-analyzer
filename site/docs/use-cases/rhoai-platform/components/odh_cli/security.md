# odh-cli: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi9/ubi:latest | 2 | root |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi:latest; Container runs as root user |
| `Dockerfile.konflux` | registry.redhat.io/openshift4/ose-cli-rhel9@sha256:b4e48cec3040a0f18fe5697e128d8deb1c9409c28b6b3249c789f7beaba10fea | 2 | root |  | multi-arch |  | Container runs as root user |

