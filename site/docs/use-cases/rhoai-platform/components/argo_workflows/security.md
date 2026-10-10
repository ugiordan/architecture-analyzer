# argo-workflows: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | gcr.io/distroless/static-debian13:latest@sha256:28efbe90d0b2f2a3ee465cc5b44f3f2cf5533514cf4d51447a977a5dc8e526d0 | 10 | 8737 |  |  |  | Unpinned base image: builder; Unpinned base image: builder; Unpinned base image: builder; Unpinned base image: argoexec-base; Unpinned base image: argoexec-base |
| `Dockerfile.windows` | argoexec-base | 4 | Administrator |  |  |  | Unpinned base image: builder; Unpinned base image: argoexec-base |
| `argo-argoexec/Dockerfile.ODH` | registry.redhat.io/ubi9/ubi-minimal:9.5 | 2 | 2000 |  | multi-arch |  |  |
| `argo-argoexec/Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:d34666200b64d5a183c867487fd2f48ee8547eb18134e5647099970e26f08af9 | 2 | 2000 |  |  |  |  |
| `argo-workflowcontroller/Dockerfile.ODH` | registry.redhat.io/ubi9/ubi-minimal:9.5 | 2 | 8737 |  | multi-arch |  |  |
| `argo-workflowcontroller/Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:d34666200b64d5a183c867487fd2f48ee8547eb18134e5647099970e26f08af9 | 2 | 8737 |  |  |  |  |
| `rhoai/Dockerfile.argoexec` | registry.redhat.io/ubi8/ubi-minimal:latest | 2 | 2000 |  |  |  | Unpinned base image: registry.redhat.io/ubi8/ubi-minimal:latest |
| `rhoai/Dockerfile.workflowcontroller` | registry.redhat.io/ubi8/ubi-minimal:latest | 2 | 8737 |  |  |  | Unpinned base image: registry.redhat.io/ubi8/ubi-minimal:latest |

