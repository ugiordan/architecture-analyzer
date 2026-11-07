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
| controller-manager | manager | ? | ? | ? | [`legacy/config/base/manager/deployment.yaml`](https://github.com/red-hat-data-services/workload-variant-autoscaler/blob/d3170fc4122c86060775671560891305c67886d4/legacy/config/base/manager/deployment.yaml) |
| controller-manager | manager | ? | ? | ? | [`legacy/config/components/openshift/deployment-patch.yaml`](https://github.com/red-hat-data-services/workload-variant-autoscaler/blob/d3170fc4122c86060775671560891305c67886d4/legacy/config/components/openshift/deployment-patch.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `legacy/Dockerfile` | gcr.io/distroless/static:nonroot@sha256:f7f8f729987ad0fdf6b05eeeae94b26e6a0f613bdf46feea7fc40f7bd72953e6 | 2 | 65532:65532 |  | multi-arch |  |  |
| `legacy/Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:c7fdcf2617a608d07f95fa01d8936fc8bf898096a60e5ebc432aae8ead50ea6c | 2 | 65532:65532 |  | multi-arch |  |  |

