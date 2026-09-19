# codeflare-operator: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| webhook-server-cert | Opaque | deployment/manager |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| manager | manager | ? | ? | ? | [`config/default/manager_webhook_patch.yaml`](https://github.com/project-codeflare/codeflare-operator/blob/aaf69895319f7a73f3f7e147a0b60214c7ad72c4/config/default/manager_webhook_patch.yaml) |
| manager | manager | ? | ? | ? | [`config/manager/manager.yaml`](https://github.com/project-codeflare/codeflare-operator/blob/aaf69895319f7a73f3f7e147a0b60214c7ad72c4/config/manager/manager.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:fac0bde5012de6a8475bb5dc18db968cdc90b12d9dc2ffaddb4a4975d0b8a04f | 2 | 65532:65532 |  | multi-arch |  |  |

