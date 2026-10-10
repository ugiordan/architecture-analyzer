# rhods-operator: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| opendatahub-operator-controller-webhook-cert | kubernetes.io/tls | deployment/controller-manager, service/webhook-service |
| opendatahub-operator-metrics-tls | Opaque | deployment/controller-manager |
| redhat-ods-operator-controller-webhook-cert | kubernetes.io/tls | deployment/rhods-operator, service/webhook-service |
| rhods-operator-metrics-tls | Opaque | deployment/rhods-operator |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| aws-cloud-manager-operator | manager | ? | ? | ? | [`config/cloudmanager/aws/local/manager_pull_policy_patch.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/934d4c80c14f8b7723383fd0a6b5a572620e34c2/config/cloudmanager/aws/local/manager_pull_policy_patch.yaml) |
| aws-cloud-manager-operator | manager | ? | true | ? | [`config/cloudmanager/aws/manager/manager.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/934d4c80c14f8b7723383fd0a6b5a572620e34c2/config/cloudmanager/aws/manager/manager.yaml) |
| aws-cloud-manager-operator | manager | ? | ? | ? | [`config/cloudmanager/aws/rhoai/manager_rhoai_patch.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/934d4c80c14f8b7723383fd0a6b5a572620e34c2/config/cloudmanager/aws/rhoai/manager_rhoai_patch.yaml) |
| azure-cloud-manager-operator | manager | ? | ? | ? | [`config/cloudmanager/azure/local/manager_pull_policy_patch.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/934d4c80c14f8b7723383fd0a6b5a572620e34c2/config/cloudmanager/azure/local/manager_pull_policy_patch.yaml) |
| azure-cloud-manager-operator | manager | ? | true | ? | [`config/cloudmanager/azure/manager/manager.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/934d4c80c14f8b7723383fd0a6b5a572620e34c2/config/cloudmanager/azure/manager/manager.yaml) |
| azure-cloud-manager-operator | manager | ? | ? | ? | [`config/cloudmanager/azure/rhoai/manager_rhoai_patch.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/934d4c80c14f8b7723383fd0a6b5a572620e34c2/config/cloudmanager/azure/rhoai/manager_rhoai_patch.yaml) |
| controller-manager | manager | ? | ? | ? | [`config/default/manager_webhook_patch.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/934d4c80c14f8b7723383fd0a6b5a572620e34c2/config/default/manager_webhook_patch.yaml) |
| controller-manager | manager | ? | ? | ? | [`config/rhaii/odh-operator/manager_auth_proxy_patch.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/934d4c80c14f8b7723383fd0a6b5a572620e34c2/config/rhaii/odh-operator/manager_auth_proxy_patch.yaml) |
| controller-manager | manager | ? | ? | ? | [`config/rhaii/odh-operator/manager_webhook_patch.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/934d4c80c14f8b7723383fd0a6b5a572620e34c2/config/rhaii/odh-operator/manager_webhook_patch.yaml) |
| controller-manager | manager | ? | ? | ? | [`config/default/manager_auth_proxy_patch.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/934d4c80c14f8b7723383fd0a6b5a572620e34c2/config/default/manager_auth_proxy_patch.yaml) |
| controller-manager | manager | ? | ? | ? | [`config/rhaii/odh-operator/manager_patch.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/934d4c80c14f8b7723383fd0a6b5a572620e34c2/config/rhaii/odh-operator/manager_patch.yaml) |
| controller-manager | manager | ? | true | ? | [`config/manager/manager.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/934d4c80c14f8b7723383fd0a6b5a572620e34c2/config/manager/manager.yaml) |
| controller-manager | manager | ? | ? | ? | [`config/rhaii/odh-local/manager_pull_policy_patch.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/934d4c80c14f8b7723383fd0a6b5a572620e34c2/config/rhaii/odh-local/manager_pull_policy_patch.yaml) |
| coreweave-cloud-manager-operator | manager | ? | ? | ? | [`config/cloudmanager/coreweave/local/manager_pull_policy_patch.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/934d4c80c14f8b7723383fd0a6b5a572620e34c2/config/cloudmanager/coreweave/local/manager_pull_policy_patch.yaml) |
| coreweave-cloud-manager-operator | manager | ? | true | ? | [`config/cloudmanager/coreweave/manager/manager.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/934d4c80c14f8b7723383fd0a6b5a572620e34c2/config/cloudmanager/coreweave/manager/manager.yaml) |
| coreweave-cloud-manager-operator | manager | ? | ? | ? | [`config/cloudmanager/coreweave/rhoai/manager_rhoai_patch.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/934d4c80c14f8b7723383fd0a6b5a572620e34c2/config/cloudmanager/coreweave/rhoai/manager_rhoai_patch.yaml) |
| rhods-operator | rhods-operator | ? | ? | ? | [`config/rhaii/rhoai/operator/manager_auth_proxy_patch.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/934d4c80c14f8b7723383fd0a6b5a572620e34c2/config/rhaii/rhoai/operator/manager_auth_proxy_patch.yaml) |
| rhods-operator | rhods-operator | ? | ? | ? | [`config/rhaii/rhoai/operator/manager_patch.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/934d4c80c14f8b7723383fd0a6b5a572620e34c2/config/rhaii/rhoai/operator/manager_patch.yaml) |
| rhods-operator | rhods-operator | ? | ? | ? | [`config/rhaii/rhoai/operator/manager_webhook_patch.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/934d4c80c14f8b7723383fd0a6b5a572620e34c2/config/rhaii/rhoai/operator/manager_webhook_patch.yaml) |
| rhods-operator | rhods-operator | ? | ? | ? | [`config/rhoai/default/manager_auth_proxy_patch.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/934d4c80c14f8b7723383fd0a6b5a572620e34c2/config/rhoai/default/manager_auth_proxy_patch.yaml) |
| rhods-operator | rhods-operator | ? | ? | ? | [`config/rhoai/default/manager_webhook_patch.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/934d4c80c14f8b7723383fd0a6b5a572620e34c2/config/rhoai/default/manager_webhook_patch.yaml) |
| rhods-operator | rhods-operator | ? | true | ? | [`config/rhoai/manager/manager.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/934d4c80c14f8b7723383fd0a6b5a572620e34c2/config/rhoai/manager/manager.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfiles/Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 4 | 1001 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/toolbox; Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `Dockerfiles/Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:d34666200b64d5a183c867487fd2f48ee8547eb18134e5647099970e26f08af9 | 2 | 1001 |  | multi-arch |  |  |
| `Dockerfiles/build-bundle.Dockerfile` | registry.access.redhat.com/ubi9/go-toolset:$GOLANG_VERSION | 1 | root |  |  |  | Container runs as root user |
| `Dockerfiles/bundle.Dockerfile` | scratch | 2 | root |  |  |  | Unpinned base image: scratch; Container runs as root user |
| `Dockerfiles/catalog.Dockerfile` | quay.io/operator-framework/opm:latest | 2 |  |  |  |  | Unpinned base image: quay.io/operator-framework/opm:latest; Unpinned base image: quay.io/operator-framework/opm:latest; No USER directive found (defaults to root) |
| `Dockerfiles/e2e-tests/e2e-tests.Dockerfile` | registry.access.redhat.com/ubi9/go-toolset:$GOLANG_VERSION | 2 | 1001 |  | multi-arch |  |  |
| `Dockerfiles/rhoai-bundle.Dockerfile` | scratch | 2 | root |  |  |  | Unpinned base image: scratch; Container runs as root user |
| `Dockerfiles/rhoai.Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:9.8@sha256:5ed244b62bbf4095080144d9d35eb8fcd3d39a9801f94aadd63b9d10978a01ae | 4 | 1001 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/toolbox |
| `Dockerfiles/toolbox.Dockerfile` | registry.fedoraproject.org/fedora-toolbox:38 | 1 |  |  |  |  | No USER directive found (defaults to root) |

