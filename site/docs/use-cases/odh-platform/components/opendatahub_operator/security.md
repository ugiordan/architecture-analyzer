# opendatahub-operator: Security

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
| aws-cloud-manager-operator | manager | ? | ? | ? | [`config/cloudmanager/aws/local/manager_pull_policy_patch.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/cloudmanager/aws/local/manager_pull_policy_patch.yaml) |
| aws-cloud-manager-operator | manager | ? | true | ? | [`config/cloudmanager/aws/manager/manager.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/cloudmanager/aws/manager/manager.yaml) |
| aws-cloud-manager-operator | manager | ? | ? | ? | [`config/cloudmanager/aws/rhoai/manager_rhoai_patch.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/cloudmanager/aws/rhoai/manager_rhoai_patch.yaml) |
| azure-cloud-manager-operator | manager | ? | ? | ? | [`config/cloudmanager/azure/local/manager_pull_policy_patch.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/cloudmanager/azure/local/manager_pull_policy_patch.yaml) |
| azure-cloud-manager-operator | manager | ? | true | ? | [`config/cloudmanager/azure/manager/manager.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/cloudmanager/azure/manager/manager.yaml) |
| azure-cloud-manager-operator | manager | ? | ? | ? | [`config/cloudmanager/azure/rhoai/manager_rhoai_patch.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/cloudmanager/azure/rhoai/manager_rhoai_patch.yaml) |
| controller-manager | manager | ? | ? | ? | [`config/default/manager_webhook_patch.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/default/manager_webhook_patch.yaml) |
| controller-manager | manager | ? | ? | ? | [`config/rhaii/odh-operator/manager_auth_proxy_patch.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/rhaii/odh-operator/manager_auth_proxy_patch.yaml) |
| controller-manager | manager | ? | ? | ? | [`config/rhaii/odh-operator/manager_webhook_patch.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/rhaii/odh-operator/manager_webhook_patch.yaml) |
| controller-manager | manager | ? | ? | ? | [`config/default/manager_auth_proxy_patch.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/default/manager_auth_proxy_patch.yaml) |
| controller-manager | manager | ? | ? | ? | [`config/rhaii/odh-operator/manager_patch.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/rhaii/odh-operator/manager_patch.yaml) |
| controller-manager | manager | ? | true | ? | [`config/manager/manager.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/manager/manager.yaml) |
| controller-manager | manager | ? | ? | ? | [`config/rhaii/odh-local/manager_pull_policy_patch.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/rhaii/odh-local/manager_pull_policy_patch.yaml) |
| coreweave-cloud-manager-operator | manager | ? | ? | ? | [`config/cloudmanager/coreweave/local/manager_pull_policy_patch.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/cloudmanager/coreweave/local/manager_pull_policy_patch.yaml) |
| coreweave-cloud-manager-operator | manager | ? | true | ? | [`config/cloudmanager/coreweave/manager/manager.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/cloudmanager/coreweave/manager/manager.yaml) |
| coreweave-cloud-manager-operator | manager | ? | ? | ? | [`config/cloudmanager/coreweave/rhoai/manager_rhoai_patch.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/cloudmanager/coreweave/rhoai/manager_rhoai_patch.yaml) |
| rhods-operator | rhods-operator | ? | ? | ? | [`config/rhaii/rhoai/operator/manager_auth_proxy_patch.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/rhaii/rhoai/operator/manager_auth_proxy_patch.yaml) |
| rhods-operator | rhods-operator | ? | ? | ? | [`config/rhaii/rhoai/operator/manager_patch.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/rhaii/rhoai/operator/manager_patch.yaml) |
| rhods-operator | rhods-operator | ? | ? | ? | [`config/rhaii/rhoai/operator/manager_webhook_patch.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/rhaii/rhoai/operator/manager_webhook_patch.yaml) |
| rhods-operator | rhods-operator | ? | ? | ? | [`config/rhoai/default/manager_auth_proxy_patch.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/rhoai/default/manager_auth_proxy_patch.yaml) |
| rhods-operator | rhods-operator | ? | ? | ? | [`config/rhoai/default/manager_webhook_patch.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/rhoai/default/manager_webhook_patch.yaml) |
| rhods-operator | rhods-operator | ? | true | ? | [`config/rhoai/manager/manager.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/rhoai/manager/manager.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfiles/Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 4 | 1001 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/toolbox; Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `Dockerfiles/build-bundle.Dockerfile` | registry.access.redhat.com/ubi9/go-toolset:$GOLANG_VERSION | 1 | root |  |  |  | Container runs as root user |
| `Dockerfiles/bundle.Dockerfile` | scratch | 2 | root |  |  |  | Unpinned base image: scratch; Container runs as root user |
| `Dockerfiles/catalog.Dockerfile` | quay.io/operator-framework/opm:latest | 2 |  |  |  |  | Unpinned base image: quay.io/operator-framework/opm:latest; Unpinned base image: quay.io/operator-framework/opm:latest; No USER directive found (defaults to root) |
| `Dockerfiles/e2e-tests/e2e-tests.Dockerfile` | golang:$GOLANG_VERSION | 2 | root |  | multi-arch |  | Container runs as root user |
| `Dockerfiles/rhoai-bundle.Dockerfile` | scratch | 2 | root |  |  |  | Unpinned base image: scratch; Container runs as root user |
| `Dockerfiles/rhoai.Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 4 | 1001 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/toolbox; Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `Dockerfiles/toolbox.Dockerfile` | registry.fedoraproject.org/fedora-toolbox:38 | 1 |  |  |  |  | No USER directive found (defaults to root) |

