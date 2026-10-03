# praxis-extproc: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| payload-processing | payload-processing | true | true | ? | [`deploy/base/instance/deployment.yaml`](https://github.com/opendatahub-io/praxis-extproc/blob/028e3dec3e6627420048e6446995b1089bdb1c7e/deploy/base/instance/deployment.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Containerfile` | registry.access.redhat.com/ubi9/ubi-minimal@${UBI9_MINIMAL_DIGEST} | 4 | 1001 |  |  |  | Unpinned base image: toolchain; Unpinned base image: builder |

