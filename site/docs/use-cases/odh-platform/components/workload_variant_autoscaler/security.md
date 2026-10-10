# workload-variant-autoscaler: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| epp-metrics-token | Opaque | deployment/controller-manager |
| metrics-server-cert | Opaque | deployment/controller-manager |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| controller-manager | manager | ? | ? | ? | [`legacy/config/base/manager/deployment.yaml`](https://github.com/opendatahub-io/workload-variant-autoscaler/blob/0b7d784a8681526ed641bbc4255e8bafcd8872fa/legacy/config/base/manager/deployment.yaml) |
| controller-manager | manager | ? | ? | ? | [`legacy/config/components/openshift/deployment-patch.yaml`](https://github.com/opendatahub-io/workload-variant-autoscaler/blob/0b7d784a8681526ed641bbc4255e8bafcd8872fa/legacy/config/components/openshift/deployment-patch.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `legacy/Dockerfile` | gcr.io/distroless/static:nonroot@sha256:f7f8f729987ad0fdf6b05eeeae94b26e6a0f613bdf46feea7fc40f7bd72953e6 | 2 | 65532:65532 |  | multi-arch |  |  |
| `legacy/Dockerfile.konflux` | registry.access.redhat.com/ubi9/ubi-minimal:9.8-1786380870@sha256:7c372902c8d211db2d25c8277ba534a73b92742a334874dced829a63b0f21221 | 2 | 65532:65532 |  | multi-arch |  |  |

