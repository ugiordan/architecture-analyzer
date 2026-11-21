# mlflow: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile.konflux` | importlib.metadata | 4 | 1001 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: importlib.metadata |
| `dev/Dockerfile.protos` | python:3.10-slim | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `dev/Dockerfile.protos.dockerignore` |  | 0 |  |  |  |  | No USER directive found (defaults to root) |
| `docker/Dockerfile` | python:3.11-slim-trixie@sha256:90744cff8f32887f075c47d747a173ff333e9e98801667af93c357fa9f5e28ff | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `docker/Dockerfile.full` | python:3.11-slim-trixie@sha256:90744cff8f32887f075c47d747a173ff333e9e98801667af93c357fa9f5e28ff | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `docker/Dockerfile.full.dev` | python:3.11-slim-trixie@sha256:90744cff8f32887f075c47d747a173ff333e9e98801667af93c357fa9f5e28ff | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `mlflow/R/mlflow/Dockerfile.dev` | rocker/r-ver:4.4.2 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `mlflow/R/mlflow/Dockerfile.r-devel` | rocker/r-ver:devel | 1 |  |  |  |  | No USER directive found (defaults to root) |

