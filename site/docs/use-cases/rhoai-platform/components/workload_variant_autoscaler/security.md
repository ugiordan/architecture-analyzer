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
| controller-manager | manager | ? | ? | ? | [`config/base/manager/deployment.yaml`](https://github.com/llm-d/workload-variant-autoscaler/blob/d4520bf1a85dcd8e7f0c3ce682db83e41d77e4a9/config/base/manager/deployment.yaml) |
| controller-manager | manager | ? | ? | ? | [`config/components/openshift/deployment-patch.yaml`](https://github.com/llm-d/workload-variant-autoscaler/blob/d4520bf1a85dcd8e7f0c3ce682db83e41d77e4a9/config/components/openshift/deployment-patch.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | gcr.io/distroless/static:nonroot@sha256:f7f8f729987ad0fdf6b05eeeae94b26e6a0f613bdf46feea7fc40f7bd72953e6 | 2 | 65532:65532 |  | multi-arch |  |  |
| `Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:a3f7ca6346ee5ca1867d1dcc79b011c5c711bf7a3e43c7cece3323d6d1ef5ac0 | 2 | 65532:65532 |  | multi-arch |  |  |

