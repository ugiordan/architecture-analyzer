# agents-operator: Security

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
| bundle-service | bundle-service | ? | true | ? | [`kagenti-operator/config/bundleservice/deployment.yaml`](https://github.com/red-hat-data-services/agents-operator/blob/a12017a8ae6ef5b079699207f7d9f5b2fde40e0b/kagenti-operator/config/bundleservice/deployment.yaml) |
| controller-manager | manager | ? | ? | ? | [`kagenti-operator/config/default/manager_webhook_patch.yaml`](https://github.com/red-hat-data-services/agents-operator/blob/a12017a8ae6ef5b079699207f7d9f5b2fde40e0b/kagenti-operator/config/default/manager_webhook_patch.yaml) |
| controller-manager | manager | ? | ? | ? | [`kagenti-operator/config/manager/manager.yaml`](https://github.com/red-hat-data-services/agents-operator/blob/a12017a8ae6ef5b079699207f7d9f5b2fde40e0b/kagenti-operator/config/manager/manager.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `authbridge/cmd/authbridge-envoy/Dockerfile` | registry.access.redhat.com/ubi9/ubi-micro@sha256:2173487b3b72b1a7b11edc908e9bbf1726f9df46a4f78fd6d19a2bab0a701f38 | 3 | 1337 |  |  |  |  |
| `authbridge/cmd/authbridge-lite/Dockerfile` | alpine:3.20@sha256:beefdbd8a1da6d2915566fde36db9db0b524eb737fc57cd1367effd16dc0d06d | 2 | 1001 |  |  |  |  |
| `authbridge/cmd/authbridge-proxy/Dockerfile` | alpine:3.20@sha256:beefdbd8a1da6d2915566fde36db9db0b524eb737fc57cd1367effd16dc0d06d | 2 | 1001 |  |  |  |  |
| `authbridge/cmd/authbridge-proxy/Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal:9.8@sha256:f26ef64172ea61f5ac5ba4be8c5c738bf8760b22b3e5148ce7b40f7308ab32db | 2 | 1001 |  | multi-arch |  |  |
| `authbridge/demos/echo/agent/Dockerfile` | gcr.io/distroless/static-debian12:nonroot | 2 |  |  |  |  | No USER directive found (defaults to root) |
| `authbridge/demos/echo/upstream/Dockerfile` | gcr.io/distroless/static-debian12:nonroot | 2 |  |  |  |  | No USER directive found (defaults to root) |
| `authbridge/demos/finance-sparc/finance-agent/Dockerfile` | gcr.io/distroless/static-debian12:nonroot | 2 |  |  |  |  | No USER directive found (defaults to root) |
| `authbridge/demos/finance-sparc/finance-mcp/Dockerfile` | gcr.io/distroless/static-debian12:nonroot | 2 |  |  |  |  | No USER directive found (defaults to root) |
| `authbridge/demos/ibac/agent/Dockerfile` | gcr.io/distroless/static-debian12:nonroot | 2 |  |  |  |  | No USER directive found (defaults to root) |
| `authbridge/demos/ibac/email-server/Dockerfile` | gcr.io/distroless/static-debian12:nonroot | 2 |  |  |  |  | No USER directive found (defaults to root) |
| `authbridge/demos/ibac/evil-server/Dockerfile` | gcr.io/distroless/static-debian12:nonroot | 2 |  |  |  |  | No USER directive found (defaults to root) |
| `authbridge/proxy-init/Dockerfile.init` | docker.io/library/alpine:3.23 | 1 | root |  |  |  | Container runs as root user |
| `authbridge/sparc-service/Dockerfile` | python:3.12-slim | 1 | sparc |  |  |  |  |
| `kagenti-operator/Dockerfile` | gcr.io/distroless/static:nonroot | 2 | 65532:65532 |  | multi-arch |  |  |
| `kagenti-operator/Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal:9.8@sha256:f26ef64172ea61f5ac5ba4be8c5c738bf8760b22b3e5148ce7b40f7308ab32db | 2 | 65532:65532 |  | multi-arch |  |  |
| `kagenti-operator/cmd/agentcard-signer/Dockerfile` | gcr.io/distroless/static:nonroot | 2 | 65532:65532 |  | multi-arch |  |  |
| `kagenti-operator/cmd/bundle-service/Dockerfile` | gcr.io/distroless/static:nonroot | 2 | 65532:65532 |  | multi-arch |  |  |
| `kagenti-operator/cmd/test-tls-agent/Dockerfile` | gcr.io/distroless/static:nonroot | 2 | 65532:65532 |  | multi-arch |  |  |
| `token-broker/Dockerfile` | gcr.io/distroless/static:nonroot | 2 | 65532:65532 |  | multi-arch |  |  |

