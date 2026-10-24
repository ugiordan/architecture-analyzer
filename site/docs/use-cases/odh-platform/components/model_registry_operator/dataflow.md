# model-registry-operator: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

| Type | GVK | Source |
|------|-----|--------|
| For | aihub/v1alpha1/AIHub | [`internal/controller/aihub_controller.go:800`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/aihub_controller.go#L800) |
| For | api/v1beta1/ModelRegistry | [`internal/controller/modelregistry_controller.go:296`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/modelregistry_controller.go#L296) |
| For | catalog/v1alpha1/Catalog | [`internal/controller/catalog_controller.go:1736`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/catalog_controller.go#L1736) |
| Owns | /v1/ConfigMap | [`internal/controller/aihub_controller.go:804`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/aihub_controller.go#L804) |
| Owns | /v1/ConfigMap | [`internal/controller/catalog_controller.go:1740`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/catalog_controller.go#L1740) |
| Owns | /v1/Secret | [`internal/controller/catalog_controller.go:1741`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/catalog_controller.go#L1741) |
| Owns | /v1/Service | [`internal/controller/modelregistry_controller.go:297`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/modelregistry_controller.go#L297) |
| Owns | /v1/Service | [`internal/controller/aihub_controller.go:802`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/aihub_controller.go#L802) |
| Owns | /v1/Service | [`internal/controller/catalog_controller.go:1738`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/catalog_controller.go#L1738) |
| Owns | /v1/ServiceAccount | [`internal/controller/aihub_controller.go:803`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/aihub_controller.go#L803) |
| Owns | /v1/ServiceAccount | [`internal/controller/modelregistry_controller.go:298`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/modelregistry_controller.go#L298) |
| Owns | /v1/ServiceAccount | [`internal/controller/catalog_controller.go:1739`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/catalog_controller.go#L1739) |
| Owns | admissionregistration.k8s.io/v1/MutatingWebhookConfiguration | [`internal/controller/aihub_controller.go:811`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/aihub_controller.go#L811) |
| Owns | admissionregistration.k8s.io/v1/ValidatingWebhookConfiguration | [`internal/controller/aihub_controller.go:810`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/aihub_controller.go#L810) |
| Owns | apps/v1/Deployment | [`internal/controller/catalog_controller.go:1737`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/catalog_controller.go#L1737) |
| Owns | apps/v1/Deployment | [`internal/controller/modelregistry_controller.go:299`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/modelregistry_controller.go#L299) |
| Owns | apps/v1/Deployment | [`internal/controller/aihub_controller.go:801`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/aihub_controller.go#L801) |
| Owns | catalog/v1alpha1/Catalog | [`internal/controller/aihub_controller.go:812`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/aihub_controller.go#L812) |
| Owns | networking.k8s.io/v1/NetworkPolicy | [`internal/controller/modelregistry_controller.go:301`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/modelregistry_controller.go#L301) |
| Owns | networking.k8s.io/v1/NetworkPolicy | [`internal/controller/aihub_controller.go:805`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/aihub_controller.go#L805) |
| Owns | networking.k8s.io/v1/NetworkPolicy | [`internal/controller/catalog_controller.go:1742`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/catalog_controller.go#L1742) |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | [`internal/controller/aihub_controller.go:808`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/aihub_controller.go#L808) |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | [`internal/controller/aihub_controller.go:809`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/aihub_controller.go#L809) |
| Owns | rbac.authorization.k8s.io/v1/Role | [`internal/controller/aihub_controller.go:806`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/aihub_controller.go#L806) |
| Owns | rbac.authorization.k8s.io/v1/Role | [`internal/controller/catalog_controller.go:1743`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/catalog_controller.go#L1743) |
| Owns | rbac.authorization.k8s.io/v1/Role | [`internal/controller/modelregistry_controller.go:300`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/modelregistry_controller.go#L300) |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | [`internal/controller/catalog_controller.go:1744`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/catalog_controller.go#L1744) |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | [`internal/controller/aihub_controller.go:807`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/aihub_controller.go#L807) |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | [`internal/controller/modelregistry_controller.go:307`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/modelregistry_controller.go#L307) |
| Owns | route/v1/Route | [`internal/controller/catalog_controller.go:1753`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/catalog_controller.go#L1753) |
| Watches | /v1/Pod | [`internal/controller/catalog_controller.go:1749`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/catalog_controller.go#L1749) |
| Watches | config/v1/APIServer | [`internal/controller/catalog_controller.go:1746`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/catalog_controller.go#L1746) |
| Watches | config/v1/APIServer | [`internal/controller/modelregistry_controller.go:303`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/modelregistry_controller.go#L303) |
| Watches | image/v1/ImageStream | [`internal/controller/catalog_controller.go:1754`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/controller/catalog_controller.go#L1754) |

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
    KubernetesAPI-->>+aihub_controller_manager: Watch Pod (informer)
    KubernetesAPI-->>+aihub_controller_manager: Watch APIServer (informer)
    KubernetesAPI-->>+aihub_controller_manager: Watch APIServer (informer)
    KubernetesAPI-->>+aihub_controller_manager: Watch ImageStream (informer)

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
| conversion-unknown | conversion | /convert |  | system/webhook-service |  |  | [`config/crd/patches/webhook_in_modelregistries.yaml`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/config/crd/patches/webhook_in_modelregistries.yaml) |
| mmodelregistry.opendatahub.io | mutating | /mutate-modelregistry-opendatahub-io-modelregistry | Fail | opendatahub/model-registry-operator-webhook-service | config/overlays/odh |  | [`config/webhook/manifests.yaml`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/config/webhook/manifests.yaml), [`kustomize:config/overlays/odh (model-registry-operator-mutating-webhook-configuration)`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/kustomize:config/overlays/odh %28model-registry-operator-mutating-webhook-configuration%29), [`internal/webhook/modelregistry_webhook.go`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/webhook/modelregistry_webhook.go) |
| vcatalog.aihub.opendatahub.io | validating | /validate-aihub-opendatahub-io-catalog | Fail | opendatahub/catalog-webhook-service | config/overlays/odh |  | [`config/overlays/catalog/webhook.yaml`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/config/overlays/catalog/webhook.yaml), [`kustomize:config/overlays/odh (catalog-validating-webhook-configuration)`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/kustomize:config/overlays/odh %28catalog-validating-webhook-configuration%29), [`internal/webhook/catalog_webhook.go`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/webhook/catalog_webhook.go) |
| vmodelregistry.opendatahub.io | validating | /validate-modelregistry-opendatahub-io-modelregistry | Fail | opendatahub/model-registry-operator-webhook-service | config/overlays/odh |  | [`config/webhook/manifests.yaml`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/config/webhook/manifests.yaml), [`kustomize:config/overlays/odh (model-registry-operator-validating-webhook-configuration)`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/kustomize:config/overlays/odh %28model-registry-operator-validating-webhook-configuration%29), [`internal/webhook/modelregistry_webhook.go`](https://github.com/opendatahub-io/model-registry-operator/blob/8d79e0e8f076b449e73acee31d8f98a7340cb264/internal/webhook/modelregistry_webhook.go) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

