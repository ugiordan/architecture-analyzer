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
| controller-manager | manager | ? | ? | ? | [`config/base/manager/deployment.yaml`](https://github.com/llm-d/workload-variant-autoscaler/blob/33e8dd112ae0f98b8b1f895aa0f70a5e973de851/config/base/manager/deployment.yaml) |
| controller-manager | manager | ? | ? | ? | [`config/components/openshift/deployment-patch.yaml`](https://github.com/llm-d/workload-variant-autoscaler/blob/33e8dd112ae0f98b8b1f895aa0f70a5e973de851/config/components/openshift/deployment-patch.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | gcr.io/distroless/static:nonroot@sha256:f7f8f729987ad0fdf6b05eeeae94b26e6a0f613bdf46feea7fc40f7bd72953e6 | 2 | 65532:65532 |  | multi-arch |  |  |
| `Dockerfile.konflux` | registry.access.redhat.com/ubi9/ubi-minimal:9.8-1786380870@sha256:7c372902c8d211db2d25c8277ba534a73b92742a334874dced829a63b0f21221 | 2 | 65532:65532 |  | multi-arch |  |  |

