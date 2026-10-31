# praxis-extproc: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| payload-processing | payload-processing | true | true | ? | [`deploy/base/instance/deployment.yaml`](https://github.com/opendatahub-io/praxis-extproc/blob/24f2d003ed303a950b6bd8ab042f0d62a5ab4278/deploy/base/instance/deployment.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Containerfile` | registry.access.redhat.com/ubi9/ubi-minimal@${UBI9_MINIMAL_DIGEST} | 4 | 1001 |  |  |  | Unpinned base image: toolchain; Unpinned base image: builder |
| `Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@${UBI9_MINIMAL_PQC_DIGEST} | 2 | 1001 |  |  |  |  |

