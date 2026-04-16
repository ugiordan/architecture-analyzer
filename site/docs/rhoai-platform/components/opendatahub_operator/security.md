# opendatahub-operator: Security

## Secrets

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

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfiles/Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 3 | 1001 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/toolbox; Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `vendor/github.com/itchyny/gojq/Dockerfile` | gcr.io/distroless/static:debug | 2 |  |  |  |  | No USER directive found (defaults to root) |
| `vendor/github.com/pelletier/go-toml/v2/Dockerfile` | scratch | 1 |  |  |  |  | Unpinned base image: scratch; No USER directive found (defaults to root) |

## Cache Architecture

### Manager Configuration

| Property | Value |
|----------|-------|
| Manager file | `cmd/main.go` |
| Cache scope | namespace-scoped |
| DefaultTransform | yes |
| Memory limit | 4Gi |

### Filtered Types

| Type | Filter Kind | Filter |
|------|-------------|--------|
| corev1.Secret | namespace | namespace-scoped |
| corev1.ConfigMap | namespace | namespace-scoped |
| appsv1.Deployment | namespace | namespace-scoped |
| networkingv1.NetworkPolicy | namespace | namespace-scoped |
| rbacv1.Role | namespace | namespace-scoped |
| rbacv1.RoleBinding | namespace | namespace-scoped |
| rbacv1.ClusterRole | label | label selector |
| rbacv1.ClusterRoleBinding | label | label selector |

### Cache-Bypassed Types (DisableFor)

- ofapiv1alpha1.Subscription
- authorizationv1.SelfSubjectRulesReview
- corev1.Pod
- userv1.Group
- ofapiv1alpha1.CatalogSource

### Implicit Informers (OOM Risk)

| Type | Source | Risk |
|------|--------|------|
| client.ListOptions | `internal/controller/components/modelsasservice/modelsasservice_controller_prerequisites.go:163` | medium |

### Issues

- No GOMEMLIMIT set in deployment (Go GC cannot pressure-tune)
- Type Auth is watched but has no cache filter (cluster-wide informer)
- Type ConsoleLink is watched but has no cache filter (cluster-wide informer)
- Type Dashboard is watched but has no cache filter (cluster-wide informer)
- Type DataSciencePipelines is watched but has no cache filter (cluster-wide informer)
- Type FeastOperator is watched but has no cache filter (cluster-wide informer)
- Type Gateway is watched but has no cache filter (cluster-wide informer)
- Type HTTPRoute is watched but has no cache filter (cluster-wide informer)
- Type Job is watched but has no cache filter (cluster-wide informer)
- Type Kserve is watched but has no cache filter (cluster-wide informer)
- Type Kueue is watched but has no cache filter (cluster-wide informer)
- Type LlamaStackOperator is watched but has no cache filter (cluster-wide informer)
- Type MLflowOperator is watched but has no cache filter (cluster-wide informer)
- Type MaaSTenant is watched but has no cache filter (cluster-wide informer)
- Type ModelController is watched but has no cache filter (cluster-wide informer)
- Type ModelRegistry is watched but has no cache filter (cluster-wide informer)
- Type MutatingWebhookConfiguration is watched but has no cache filter (cluster-wide informer)
- Type Namespace is watched but has no cache filter (cluster-wide informer)
- Type PodDisruptionBudget is watched but has no cache filter (cluster-wide informer)
- Type PodMonitor is watched but has no cache filter (cluster-wide informer)
- Type PrometheusRule is watched but has no cache filter (cluster-wide informer)
- Type Ray is watched but has no cache filter (cluster-wide informer)
- Type Route is watched but has no cache filter (cluster-wide informer)
- Type SecurityContextConstraints is watched but has no cache filter (cluster-wide informer)
- Type Service is watched but has no cache filter (cluster-wide informer)
- Type ServiceAccount is watched but has no cache filter (cluster-wide informer)
- Type ServiceMonitor is watched but has no cache filter (cluster-wide informer)
- Type SparkOperator is watched but has no cache filter (cluster-wide informer)
- Type Template is watched but has no cache filter (cluster-wide informer)
- Type Trainer is watched but has no cache filter (cluster-wide informer)
- Type TrainingOperator is watched but has no cache filter (cluster-wide informer)
- Type TrustyAI is watched but has no cache filter (cluster-wide informer)
- Type ValidatingAdmissionPolicy is watched but has no cache filter (cluster-wide informer)
- Type ValidatingAdmissionPolicyBinding is watched but has no cache filter (cluster-wide informer)
- Type ValidatingWebhookConfiguration is watched but has no cache filter (cluster-wide informer)
- Type Workbenches is watched but has no cache filter (cluster-wide informer)

