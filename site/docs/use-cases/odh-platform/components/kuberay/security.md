# kuberay: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| kuberay-operator | kuberay-operator | ? | ? | ? | [`ray-operator/config/manager/manager.yaml`](https://github.com/opendatahub-io/kuberay/blob/e2b85bf7c3476dfd3224fa5cdaa67f235115eb53/ray-operator/config/manager/manager.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `apiserver/Dockerfile` | scratch | 2 | 65532:65532 |  |  |  | Unpinned base image: scratch |
| `proto/Dockerfile` | registry.access.redhat.com/ubi9/go-toolset:1.20.10 | 1 | 65532:65532 |  |  |  |  |
| `ray-operator/Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:9.3 | 2 | 65532:65532 |  |  |  |  |

