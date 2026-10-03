# odh-cli: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi9/ubi:9.8@sha256:094ea2ecfd3225af8f93807b99daa9ff33710fc705ebdf6e8466f46ed605585c | 2 | root |  | multi-arch |  | Container runs as root user |
| `Dockerfile.konflux` | registry.redhat.io/openshift4/ose-cli-rhel9@sha256:b36a4c332f05bd71c5315b6cfe38987809bae09c2865217e1c7240fa6bfd813b | 2 | root |  | multi-arch |  | Container runs as root user |

