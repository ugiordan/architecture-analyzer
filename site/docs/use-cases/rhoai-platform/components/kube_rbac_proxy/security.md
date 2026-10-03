# kube-rbac-proxy: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | $BASEIMAGE | 1 | 65532:65532 |  |  |  | Unpinned base image: $BASEIMAGE |
| `Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:e8673bd338b3f093501a9a9088f043d8959df120edd9947583e2d9bb3aacc3a8 | 2 | 65534 |  | multi-arch |  |  |
| `Dockerfile.ocp` | registry.ci.openshift.org/ocp/4.22:base-rhel9 | 2 | 65534 |  |  |  |  |
| `Dockerfile.redhat` | registry.access.redhat.com/ubi9/ubi-minimal:9.8@sha256:8ebe2ad8fdf3cab3e5a53c1edc69194c98209cfadab24b884f4ad9ebcf7bbbfc | 2 | 65534 |  | multi-arch |  |  |

