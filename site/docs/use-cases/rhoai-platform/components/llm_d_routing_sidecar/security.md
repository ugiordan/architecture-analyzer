# llm-d-routing-sidecar: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| 0 | cmd | ? | ? | ? | [`deploy/common/statefulset.yaml`](https://github.com/llm-d/llm-d-routing-sidecar/blob/214ed72b3bcd2ea0d66ae2f15d82e0037a726c06/deploy/common/statefulset.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi9/ubi-micro:latest | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-micro:latest |
| `Dockerfile.konflux` | registry.access.redhat.com/ubi9/ubi:latest | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi:latest |

