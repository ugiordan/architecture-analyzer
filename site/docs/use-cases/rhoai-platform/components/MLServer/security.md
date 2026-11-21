# MLServer: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | mlserver.settings | 4 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE}; Unpinned base image: pathlib; Unpinned base image: mlserver.settings |
| `Dockerfile.cuda` | mlserver.settings | 4 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE}; Unpinned base image: pathlib; Unpinned base image: mlserver.settings |
| `Dockerfile.konflux` | mlserver.settings | 3 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: pathlib; Unpinned base image: mlserver.settings |
| `Dockerfile.konflux.cuda` | mlserver.settings | 3 | 1000 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: pathlib; Unpinned base image: mlserver.settings |

