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
| manager | manager | ? | ? | ? | [`config/default/manager_webhook_patch.yaml`](https://github.com/red-hat-data-services/codeflare-operator/blob/6b74024f8f7835d0a470a1ebc4fc4184dc36f3a3/config/default/manager_webhook_patch.yaml) |
| manager | manager | ? | ? | ? | [`config/manager/manager.yaml`](https://github.com/red-hat-data-services/codeflare-operator/blob/6b74024f8f7835d0a470a1ebc4fc4184dc36f3a3/config/manager/manager.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:d0fdfb1401ee754872f60c2527563671c16112080e5254139133908178700911 | 2 | 65532:65532 |  | multi-arch |  |  |

