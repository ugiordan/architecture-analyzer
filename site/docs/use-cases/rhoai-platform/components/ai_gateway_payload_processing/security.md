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
| `Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc:9.8@sha256:d638f3b49b79575329aacdae2739f954effe2467c12147711e73445c4fc97819 | 2 | 1001 |  | multi-arch |  |  |
| `Dockerfile.konflux.e2e` | registry.access.redhat.com/ubi9/ubi-minimal:9.8@sha256:f26ef64172ea61f5ac5ba4be8c5c738bf8760b22b3e5148ce7b40f7308ab32db | 3 | 1001 |  | multi-arch |  | Unpinned base image: registry.redhat.io/openshift4/ose-cli:latest |

