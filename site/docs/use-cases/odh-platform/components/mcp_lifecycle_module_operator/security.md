# mcp-lifecycle-module-operator: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| mcplo-webhook-server-cert | Opaque | deployment/controller-manager |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| controller-manager | manager | ? | true | ? | [`config/manager/manager.yaml`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/6ed00689e045b97ffdd1027ed4d0471cf41b098c/config/manager/manager.yaml) |
| controller-manager | manager | ? | ? | ? | [`config/manifests/mcp-lifecycle-operator/coverage/manager_coverage_patch.yaml`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/6ed00689e045b97ffdd1027ed4d0471cf41b098c/config/manifests/mcp-lifecycle-operator/coverage/manager_coverage_patch.yaml) |
| controller-manager | manager | ? | ? | ? | [`config/manifests/mcp-lifecycle-operator/default/manager_webhook_patch.yaml`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/6ed00689e045b97ffdd1027ed4d0471cf41b098c/config/manifests/mcp-lifecycle-operator/default/manager_webhook_patch.yaml) |
| controller-manager | manager | ? | ? | ? | [`config/manifests/mcp-lifecycle-operator/default-with-webhook/manager_webhook_patch.yaml`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/6ed00689e045b97ffdd1027ed4d0471cf41b098c/config/manifests/mcp-lifecycle-operator/default-with-webhook/manager_webhook_patch.yaml) |
| controller-manager | manager | true | true | ? | [`config/manifests/mcp-lifecycle-operator/manager/manager.yaml`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/6ed00689e045b97ffdd1027ed4d0471cf41b098c/config/manifests/mcp-lifecycle-operator/manager/manager.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile.konflux` | registry.access.redhat.com/ubi9/ubi-minimal@sha256:4b79c93106c14ad1c234907f83441ff0e4f8d24e78426a36e159098e9f7e6fa5 | 3 | 65532:65532 |  | multi-arch |  |  |

