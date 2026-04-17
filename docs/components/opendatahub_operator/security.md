# opendatahub-operator: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| odh-model-controller-webhook-cert | kubernetes.io/tls | deployment/odh-model-controller, service/odh-model-controller-webhook-service |
| webhook-server-cert | Opaque | deployment/controller-manager, deployment/kuberay-operator |
| controller-manager-metrics-service | Opaque | deployment/controller-manager |
| odh-notebook-controller-webhook-cert | kubernetes.io/tls | service/webhook-service |
| opendatahub-operator-controller-webhook-cert | kubernetes.io/tls | deployment/controller-manager, service/webhook-service |
| redhat-ods-operator-controller-webhook-cert | kubernetes.io/tls | deployment/rhods-operator, service/webhook-service |
| kserve-webhook-server-cert | Opaque | deployment/kserve-controller-manager |
| kubeflow-training-operator-webhook-cert | Opaque | deployment/training-operator |
| dashboard-proxy-tls | kubernetes.io/tls | deployment/odh-dashboard, service/odh-dashboard |
| training-operator-webhook-cert | Opaque | deployment/training-operator |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| odh-dashboard | odh-dashboard | ? | ? | ? | `opt/manifests/dashboard/core-bases/base/deployment.yaml` |
| odh-dashboard | kube-rbac-proxy | ? | ? | ? | `opt/manifests/dashboard/core-bases/base/deployment.yaml` |
| training-operator | training-operator | ? | ? | ? | `opt/manifests/trainingoperator/base/deployment.yaml` |
| azure-cloud-manager-operator | manager | ? | ? | ? | `config/cloudmanager/azure/local/manager_pull_policy_patch.yaml` |
| azure-cloud-manager-operator | manager | ? | true | ? | `config/cloudmanager/azure/manager/manager.yaml` |
| azure-cloud-manager-operator | manager | ? | ? | ? | `config/cloudmanager/azure/rhoai/manager_rhoai_patch.yaml` |
| coreweave-cloud-manager-operator | manager | ? | ? | ? | `config/cloudmanager/coreweave/local/manager_pull_policy_patch.yaml` |
| coreweave-cloud-manager-operator | manager | ? | true | ? | `config/cloudmanager/coreweave/manager/manager.yaml` |
| coreweave-cloud-manager-operator | manager | ? | ? | ? | `config/cloudmanager/coreweave/rhoai/manager_rhoai_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `config/default/manager_auth_proxy_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `config/default/manager_webhook_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `config/manager/manager.yaml` |
| controller-manager | manager | ? | ? | ? | `config/rhaii/odh-local/manager_pull_policy_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `config/rhaii/odh-operator/manager_auth_proxy_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `config/rhaii/odh-operator/manager_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `config/rhaii/odh-operator/manager_webhook_patch.yaml` |
| rhods-operator | rhods-operator | ? | ? | ? | `config/rhaii/rhoai/operator/manager_auth_proxy_patch.yaml` |
| rhods-operator | rhods-operator | ? | ? | ? | `config/rhaii/rhoai/operator/manager_patch.yaml` |
| rhods-operator | rhods-operator | ? | ? | ? | `config/rhaii/rhoai/operator/manager_webhook_patch.yaml` |
| rhods-operator | rhods-operator | ? | ? | ? | `config/rhoai/default/manager_auth_proxy_patch.yaml` |
| rhods-operator | rhods-operator | ? | ? | ? | `config/rhoai/default/manager_webhook_patch.yaml` |
| rhods-operator | rhods-operator | ? | ? | ? | `config/rhoai/manager/manager.yaml` |
| controller-manager | manager | ? | ? | ? | `opt/manifests/datasciencepipelines/manager/manager.yaml` |
| controller-manager | manager | ? | ? | ? | `opt/manifests/feastoperator/default/manager_config_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `opt/manifests/feastoperator/manager/manager.yaml` |
| kserve-controller-manager | kube-rbac-proxy | true | true | false | `opt/manifests/kserve/default/manager_auth_proxy_patch.yaml` |
| kserve-controller-manager | manager | ? | ? | ? | `opt/manifests/kserve/default/manager_auth_proxy_patch.yaml` |
| kserve-controller-manager | manager | ? | ? | ? | `opt/manifests/kserve/default/manager_image_patch.yaml` |
| kserve-controller-manager | manager | ? | ? | ? | `opt/manifests/kserve/default/manager_prometheus_metrics_patch.yaml` |
| kserve-controller-manager | manager | ? | ? | ? | `opt/manifests/kserve/default/manager_resources_patch.yaml` |
| kserve-localmodel-controller-manager | manager | true | true | false | `opt/manifests/kserve/localmodels/manager.yaml` |
| kserve-controller-manager | manager | true | true | false | `opt/manifests/kserve/manager/manager.yaml` |
| kserve-controller-manager | manager | ? | ? | ? | `opt/manifests/kserve/overlays/odh-test/manager_image_patch.yaml` |
| kserve-controller-manager | manager | ? | ? | ? | `opt/manifests/kserve/overlays/test/manager_image_patch.yaml` |
| controller-manager | kube-rbac-proxy | ? | ? | ? | `opt/manifests/llamastackoperator/default/manager_auth_proxy_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `opt/manifests/llamastackoperator/default/manager_auth_proxy_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `opt/manifests/llamastackoperator/default/manager_config_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `opt/manifests/llamastackoperator/manager/manager.yaml` |
| odh-model-controller | manager | ? | ? | ? | `opt/manifests/modelcontroller/default/manager_webhook_patch.yaml` |
| odh-model-controller | manager | ? | ? | ? | `opt/manifests/modelcontroller/manager/manager.yaml` |
| controller-manager | manager | ? | ? | ? | `opt/manifests/modelregistry/default/manager_auth_proxy_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `opt/manifests/modelregistry/default/manager_config_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `opt/manifests/modelregistry/default/manager_webhook_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `opt/manifests/modelregistry/manager/manager.yaml` |
| controller-manager | manager | ? | ? | ? | `opt/manifests/modelregistry/overlays/odh/patches/manager_auth_proxy_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `opt/manifests/modelregistry/overlays/odh/patches/manager_istio_config_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `opt/manifests/modelregistry/overlays/odh/patches/manager_migration_env_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `opt/manifests/modelregistry/overlays/odh/patches/manager_webhook_patch.yaml` |
| kuberay-operator | kuberay-operator | ? | ? | ? | `opt/manifests/ray/default-with-webhooks/manager_webhook_patch.yaml` |
| kuberay-operator | kuberay-operator | ? | ? | ? | `opt/manifests/ray/manager/manager.yaml` |
| training-operator | training-operator | ? | ? | ? | `opt/manifests/trainingoperator/rhoai/manager_config_patch.yaml` |
| training-operator | training-operator | ? | ? | ? | `opt/manifests/trainingoperator/rhoai/manager_metrics_patch.yaml` |
| controller-manager | manager | true | ? | ? | `opt/manifests/trustyai/manager/manager.yaml` |
| controller-manager | kube-rbac-proxy | ? | ? | ? | `opt/manifests/workbenches/kf-notebook-controller/default/manager_auth_proxy_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `opt/manifests/workbenches/kf-notebook-controller/default/manager_auth_proxy_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `opt/manifests/workbenches/kf-notebook-controller/default/manager_image_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `opt/manifests/workbenches/kf-notebook-controller/default/manager_prometheus_metrics_patch.yaml` |
| controller-manager | manager | ? | ? | ? | `opt/manifests/workbenches/kf-notebook-controller/default/manager_webhook_patch.yaml` |
| deployment | manager | ? | ? | ? | `opt/manifests/workbenches/kf-notebook-controller/manager/manager.yaml` |
| deployment | manager | ? | ? | ? | `opt/manifests/workbenches/kf-notebook-controller/overlays/openshift/manager_openshift_patch.yaml` |
| manager | manager | ? | ? | ? | `opt/manifests/workbenches/odh-notebook-controller/manager/manager.yaml` |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfiles/Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 3 | 1001 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/toolbox; Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `vendor/github.com/itchyny/gojq/Dockerfile` | gcr.io/distroless/static:debug | 2 |  |  |  |  | No USER directive found (defaults to root) |
| `vendor/github.com/pelletier/go-toml/v2/Dockerfile` | scratch | 1 |  |  |  |  | Unpinned base image: scratch; No USER directive found (defaults to root) |

