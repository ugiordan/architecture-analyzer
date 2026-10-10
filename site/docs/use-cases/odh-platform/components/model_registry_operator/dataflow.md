# model-registry-operator: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

| Type | GVK | Source |
|------|-----|--------|
| For | aihub/v1alpha1/AIHub | [`internal/controller/aihub_controller.go:768`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/aihub_controller.go#L768) |
| For | api/v1beta1/ModelRegistry | [`internal/controller/modelregistry_controller.go:282`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/modelregistry_controller.go#L282) |
| For | catalog/v1alpha1/Catalog | [`internal/controller/catalog_controller.go:1519`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/catalog_controller.go#L1519) |
| Owns | /v1/ConfigMap | [`internal/controller/aihub_controller.go:772`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/aihub_controller.go#L772) |
| Owns | /v1/ConfigMap | [`internal/controller/catalog_controller.go:1523`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/catalog_controller.go#L1523) |
| Owns | /v1/Secret | [`internal/controller/catalog_controller.go:1524`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/catalog_controller.go#L1524) |
| Owns | /v1/Service | [`internal/controller/catalog_controller.go:1521`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/catalog_controller.go#L1521) |
| Owns | /v1/Service | [`internal/controller/modelregistry_controller.go:283`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/modelregistry_controller.go#L283) |
| Owns | /v1/Service | [`internal/controller/aihub_controller.go:770`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/aihub_controller.go#L770) |
| Owns | /v1/ServiceAccount | [`internal/controller/modelregistry_controller.go:284`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/modelregistry_controller.go#L284) |
| Owns | /v1/ServiceAccount | [`internal/controller/aihub_controller.go:771`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/aihub_controller.go#L771) |
| Owns | /v1/ServiceAccount | [`internal/controller/catalog_controller.go:1522`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/catalog_controller.go#L1522) |
| Owns | admissionregistration.k8s.io/v1/MutatingWebhookConfiguration | [`internal/controller/aihub_controller.go:779`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/aihub_controller.go#L779) |
| Owns | admissionregistration.k8s.io/v1/ValidatingWebhookConfiguration | [`internal/controller/aihub_controller.go:778`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/aihub_controller.go#L778) |
| Owns | apps/v1/Deployment | [`internal/controller/catalog_controller.go:1520`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/catalog_controller.go#L1520) |
| Owns | apps/v1/Deployment | [`internal/controller/aihub_controller.go:769`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/aihub_controller.go#L769) |
| Owns | apps/v1/Deployment | [`internal/controller/modelregistry_controller.go:285`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/modelregistry_controller.go#L285) |
| Owns | catalog/v1alpha1/Catalog | [`internal/controller/aihub_controller.go:780`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/aihub_controller.go#L780) |
| Owns | networking.k8s.io/v1/NetworkPolicy | [`internal/controller/catalog_controller.go:1525`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/catalog_controller.go#L1525) |
| Owns | networking.k8s.io/v1/NetworkPolicy | [`internal/controller/modelregistry_controller.go:287`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/modelregistry_controller.go#L287) |
| Owns | networking.k8s.io/v1/NetworkPolicy | [`internal/controller/aihub_controller.go:773`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/aihub_controller.go#L773) |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | [`internal/controller/aihub_controller.go:776`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/aihub_controller.go#L776) |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | [`internal/controller/aihub_controller.go:777`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/aihub_controller.go#L777) |
| Owns | rbac.authorization.k8s.io/v1/Role | [`internal/controller/catalog_controller.go:1526`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/catalog_controller.go#L1526) |
| Owns | rbac.authorization.k8s.io/v1/Role | [`internal/controller/modelregistry_controller.go:286`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/modelregistry_controller.go#L286) |
| Owns | rbac.authorization.k8s.io/v1/Role | [`internal/controller/aihub_controller.go:774`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/aihub_controller.go#L774) |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | [`internal/controller/aihub_controller.go:775`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/aihub_controller.go#L775) |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | [`internal/controller/catalog_controller.go:1527`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/catalog_controller.go#L1527) |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | [`internal/controller/modelregistry_controller.go:289`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/modelregistry_controller.go#L289) |
| Owns | route/v1/Route | [`internal/controller/catalog_controller.go:1530`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/controller/catalog_controller.go#L1530) |

### Programmatic Resource Operations

| Verb | Kind | Group | Condition |
|------|------|-------|----------|
| update | AIHub | aihub |  |
| delete | Deployment | apps |  |
| update | Catalog | catalog |  |
| update | ModelRegistry | api |  |

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for model-registry-operator

    participant KubernetesAPI as Kubernetes API
    participant aihub_controller_manager as aihub-controller-manager
    participant catalog_controller_manager as catalog-controller-manager
    participant model_registry_operator_controller_manager as model-registry-operator-controller-manager
    participant template_value as template-value

    KubernetesAPI->>+aihub_controller_manager: Watch AIHub (reconcile)
    KubernetesAPI->>+aihub_controller_manager: Watch ModelRegistry (reconcile)
    KubernetesAPI->>+aihub_controller_manager: Watch Catalog (reconcile)
    aihub_controller_manager->>KubernetesAPI: Create/Update ConfigMap
    aihub_controller_manager->>KubernetesAPI: Create/Update ConfigMap
    aihub_controller_manager->>KubernetesAPI: Create/Update Secret
    aihub_controller_manager->>KubernetesAPI: Create/Update Service
    aihub_controller_manager->>KubernetesAPI: Create/Update Service
    aihub_controller_manager->>KubernetesAPI: Create/Update Service
    aihub_controller_manager->>KubernetesAPI: Create/Update ServiceAccount
    aihub_controller_manager->>KubernetesAPI: Create/Update ServiceAccount
    aihub_controller_manager->>KubernetesAPI: Create/Update ServiceAccount
    aihub_controller_manager->>KubernetesAPI: Create/Update MutatingWebhookConfiguration
    aihub_controller_manager->>KubernetesAPI: Create/Update ValidatingWebhookConfiguration
    aihub_controller_manager->>KubernetesAPI: Create/Update Deployment
    aihub_controller_manager->>KubernetesAPI: Create/Update Deployment
    aihub_controller_manager->>KubernetesAPI: Create/Update Deployment
    aihub_controller_manager->>KubernetesAPI: Create/Update Catalog
    aihub_controller_manager->>KubernetesAPI: Create/Update NetworkPolicy
    aihub_controller_manager->>KubernetesAPI: Create/Update NetworkPolicy
    aihub_controller_manager->>KubernetesAPI: Create/Update NetworkPolicy
    aihub_controller_manager->>KubernetesAPI: Create/Update ClusterRole
    aihub_controller_manager->>KubernetesAPI: Create/Update ClusterRoleBinding
    aihub_controller_manager->>KubernetesAPI: Create/Update Role
    aihub_controller_manager->>KubernetesAPI: Create/Update Role
    aihub_controller_manager->>KubernetesAPI: Create/Update Role
    aihub_controller_manager->>KubernetesAPI: Create/Update RoleBinding
    aihub_controller_manager->>KubernetesAPI: Create/Update RoleBinding
    aihub_controller_manager->>KubernetesAPI: Create/Update RoleBinding
    aihub_controller_manager->>KubernetesAPI: Create/Update Route

    Note over aihub_controller_manager: Exposed Services
    Note right of aihub_controller_manager: catalog-webhook-service:443/TCP []
    Note right of aihub_controller_manager: model-registry-operator-controller-manager-metrics-service:8443/TCP [https]
    Note right of aihub_controller_manager: model-registry-operator-webhook-service:443/TCP []
    Note right of aihub_controller_manager: template-value-postgres:5432/TCP [postgresql]
    Note right of aihub_controller_manager: template-value-postgres:5432/TCP []

    Note over KubernetesAPI: Defined CRDs
    Note right of KubernetesAPI: ModelRegistry (modelregistry.opendatahub.io/v1beta1)
```

### Webhooks

| Name | Type | Path | Failure Policy | Service | Overlays | Enable Condition | Sources |
|------|------|------|----------------|---------|----------|------------------|----------|
| conversion-unknown | conversion | /convert |  | system/webhook-service |  |  | [`config/crd/patches/webhook_in_modelregistries.yaml`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/config/crd/patches/webhook_in_modelregistries.yaml) |
| mmodelregistry.opendatahub.io | mutating | /mutate-modelregistry-opendatahub-io-modelregistry | Fail | opendatahub/model-registry-operator-webhook-service | config/overlays/odh |  | [`config/webhook/manifests.yaml`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/config/webhook/manifests.yaml), [`kustomize:config/overlays/odh (model-registry-operator-mutating-webhook-configuration)`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/kustomize:config/overlays/odh %28model-registry-operator-mutating-webhook-configuration%29), [`internal/webhook/modelregistry_webhook.go`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/webhook/modelregistry_webhook.go) |
| vcatalog.aihub.opendatahub.io | validating | /validate-aihub-opendatahub-io-catalog | Fail | opendatahub/catalog-webhook-service | config/overlays/odh |  | [`config/overlays/catalog/webhook.yaml`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/config/overlays/catalog/webhook.yaml), [`kustomize:config/overlays/odh (catalog-validating-webhook-configuration)`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/kustomize:config/overlays/odh %28catalog-validating-webhook-configuration%29), [`internal/webhook/catalog_webhook.go`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/webhook/catalog_webhook.go) |
| vmodelregistry.opendatahub.io | validating | /validate-modelregistry-opendatahub-io-modelregistry | Fail | opendatahub/model-registry-operator-webhook-service | config/overlays/odh |  | [`config/webhook/manifests.yaml`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/config/webhook/manifests.yaml), [`kustomize:config/overlays/odh (model-registry-operator-validating-webhook-configuration)`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/kustomize:config/overlays/odh %28model-registry-operator-validating-webhook-configuration%29), [`internal/webhook/modelregistry_webhook.go`](https://github.com/opendatahub-io/model-registry-operator/blob/0493841d3ecd2685731bff495be33b66938a36e4/internal/webhook/modelregistry_webhook.go) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

