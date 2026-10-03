# mcp-lifecycle-module-operator: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| webhook-server-cert | Opaque | deployment/controller-manager |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| controller-manager | manager | ? | true | ? | [`config/manager/manager.yaml`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/ed1187c1c5382db14bac17ee53f5a8a1dd5dbd7d/config/manager/manager.yaml) |
| controller-manager | manager | ? | ? | ? | [`config/manifests/mcp-lifecycle-operator/coverage/manager_coverage_patch.yaml`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/ed1187c1c5382db14bac17ee53f5a8a1dd5dbd7d/config/manifests/mcp-lifecycle-operator/coverage/manager_coverage_patch.yaml) |
| controller-manager | manager | ? | ? | ? | [`config/manifests/mcp-lifecycle-operator/default/manager_webhook_patch.yaml`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/ed1187c1c5382db14bac17ee53f5a8a1dd5dbd7d/config/manifests/mcp-lifecycle-operator/default/manager_webhook_patch.yaml) |
| controller-manager | manager | ? | ? | ? | [`config/manifests/mcp-lifecycle-operator/default-with-webhook/manager_webhook_patch.yaml`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/ed1187c1c5382db14bac17ee53f5a8a1dd5dbd7d/config/manifests/mcp-lifecycle-operator/default-with-webhook/manager_webhook_patch.yaml) |
| controller-manager | manager | true | true | ? | [`config/manifests/mcp-lifecycle-operator/manager/manager.yaml`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/ed1187c1c5382db14bac17ee53f5a8a1dd5dbd7d/config/manifests/mcp-lifecycle-operator/manager/manager.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile.konflux` | registry.access.redhat.com/ubi9/ubi-minimal@sha256:7c372902c8d211db2d25c8277ba534a73b92742a334874dced829a63b0f21221 | 3 | 65532:65532 |  | multi-arch |  |  |

