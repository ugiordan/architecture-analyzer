# llm-d-inference-scheduler: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile.builder` | golang:1.27.4 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `Dockerfile.coordinator` | ${BASE_IMAGE} | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE} |
| `Dockerfile.epp` | ${BASE_IMAGE} | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE} |
| `Dockerfile.konflux.coordinator` | registry.redhat.io/ubi9/ubi-minimal-pqc:9.8-1793092152@sha256:72187ec58a94bd958ff48e41881d7c5d8bc0b49586cc1661da5ba4d92374c206 | 2 | 65532:65532 |  | multi-arch |  |  |
| `Dockerfile.konflux.epp` | registry.redhat.io/ubi9/ubi-minimal-pqc:9.8-1793092152@sha256:72187ec58a94bd958ff48e41881d7c5d8bc0b49586cc1661da5ba4d92374c206 | 2 | 65532:65532 |  | multi-arch |  |  |
| `Dockerfile.konflux.sidecar` | registry.redhat.io/ubi9/ubi-minimal-pqc:9.8-1793092152@sha256:72187ec58a94bd958ff48e41881d7c5d8bc0b49586cc1661da5ba4d92374c206 | 2 | 65532:65532 |  | multi-arch |  |  |
| `Dockerfile.sidecar` | ${BASE_IMAGE} | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE} |

