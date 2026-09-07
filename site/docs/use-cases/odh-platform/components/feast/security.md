# feast: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| controller-manager | manager | ? | ? | ? | [`infra/feast-operator/config/default/manager_config_patch.yaml`](https://github.com/feast-dev/feast/blob/6ad332bf6c06499319483b0d3c87d96165c41d53/infra/feast-operator/config/default/manager_config_patch.yaml) |
| controller-manager | manager | ? | ? | ? | [`infra/feast-operator/config/manager/manager.yaml`](https://github.com/feast-dev/feast/blob/6ad332bf6c06499319483b0d3c87d96165c41d53/infra/feast-operator/config/manager/manager.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `.gitpod.Dockerfile` | gitpod/workspace-base | 1 |  |  |  |  | Unpinned base image: gitpod/workspace-base; No USER directive found (defaults to root) |
| `go/infra/docker/feature-server/Dockerfile` | golang:1.25 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `infra/feast-operator/Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:9.8 | 2 | 65532:65532 |  | multi-arch |  |  |
| `infra/feast-operator/bundle.Dockerfile` | scratch | 1 |  |  |  |  | Unpinned base image: scratch; No USER directive found (defaults to root) |
| `java/infra/docker/feature-server/Dockerfile` | amazoncorretto:11 | 2 |  |  |  |  | No USER directive found (defaults to root) |
| `java/infra/docker/feature-server/Dockerfile.dev` | openjdk:11-jre | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `sdk/python/feast/infra/compute_engines/aws_lambda/Dockerfile` | public.ecr.aws/lambda/python:3.9 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `sdk/python/feast/infra/compute_engines/kubernetes/Dockerfile` | debian:11-slim | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `sdk/python/feast/infra/compute_engines/spark_application/Dockerfile` | apache/spark:4.0.1 | 1 | spark |  |  |  |  |
| `sdk/python/feast/infra/feature_servers/multicloud/Dockerfile` | registry.access.redhat.com/ubi9/python-312-minimal:latest | 1 | 1001 |  |  |  | Unpinned base image: registry.access.redhat.com/ubi9/python-312-minimal:latest |
| `sdk/python/feast/infra/feature_servers/multicloud/Dockerfile.dev` | registry.access.redhat.com/ubi9/python-312-minimal:latest | 1 | 1001 |  |  |  | Unpinned base image: registry.access.redhat.com/ubi9/python-312-minimal:latest |
| `sdk/python/feast/infra/feature_servers/multicloud/offline/Dockerfile.binary` | yarn-builder:latest | 1 |  |  |  |  | Unpinned base image: yarn-builder:latest; No USER directive found (defaults to root) |
| `sdk/python/feast/infra/feature_servers/multicloud/offline/Dockerfile.binary.release` | registry.access.redhat.com/ubi9/python-312-minimal:latest | 1 |  |  |  |  | Unpinned base image: registry.access.redhat.com/ubi9/python-312-minimal:latest; No USER directive found (defaults to root) |
| `sdk/python/feast/infra/feature_servers/multicloud/offline/Dockerfile.builder.yarn` | registry.access.redhat.com/ubi9/python-312-minimal:latest | 1 | 1001 |  |  |  | Unpinned base image: registry.access.redhat.com/ubi9/python-312-minimal:latest |
| `sdk/python/feast/infra/feature_servers/multicloud/offline/Dockerfile.builder.yum` | registry.access.redhat.com/ubi9/python-312-minimal:latest | 1 | 1001 |  |  |  | Unpinned base image: registry.access.redhat.com/ubi9/python-312-minimal:latest |
| `sdk/python/feast/infra/feature_servers/multicloud/offline/Dockerfile.sdist` | yarn-builder:latest | 1 | 1001 |  |  |  | Unpinned base image: yarn-builder:latest |
| `sdk/python/feast/infra/feature_servers/multicloud/offline/Dockerfile.sdist.release` | registry.access.redhat.com/ubi9/python-312-minimal:latest | 1 | 1001 |  |  |  | Unpinned base image: registry.access.redhat.com/ubi9/python-312-minimal:latest |
| `sdk/python/feast/infra/transformation_servers/Dockerfile` | python:3.11-slim | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `ui/docker/Dockerfile` | node:17.9.0-slim | 1 |  |  |  |  | No USER directive found (defaults to root) |

