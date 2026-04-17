# model-registry-operator: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

| Type | GVK | Source |
|------|-----|--------|
| For | api/v1beta1/ModelRegistry | `internal/controller/modelregistry_controller.go:258` |
| Owns | /v1/Service | `internal/controller/modelregistry_controller.go:259` |
| Owns | /v1/ServiceAccount | `internal/controller/modelregistry_controller.go:260` |
| Owns | apps/v1/Deployment | `internal/controller/modelregistry_controller.go:261` |
| Owns | networking.k8s.io/v1/NetworkPolicy | `internal/controller/modelregistry_controller.go:263` |
| Owns | rbac.authorization.k8s.io/v1/Role | `internal/controller/modelregistry_controller.go:262` |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | `internal/controller/modelregistry_controller.go:265` |
| Watches | /v1/ConfigMap | `internal/controller/modelcatalog_controller.go:1258` |
| Watches | /v1/PersistentVolumeClaim | `internal/controller/modelcatalog_controller.go:1262` |
| Watches | /v1/Secret | `internal/controller/modelcatalog_controller.go:1259` |
| Watches | /v1/Service | `internal/controller/modelcatalog_controller.go:1261` |
| Watches | /v1/ServiceAccount | `internal/controller/modelcatalog_controller.go:1260` |
| Watches | apps/v1/Deployment | `internal/controller/modelcatalog_controller.go:1257` |
| Watches | networking.k8s.io/v1/NetworkPolicy | `internal/controller/modelcatalog_controller.go:1266` |
| Watches | rbac.authorization.k8s.io/v1/ClusterRoleBinding | `internal/controller/modelcatalog_controller.go:1263` |
| Watches | rbac.authorization.k8s.io/v1/Role | `internal/controller/modelcatalog_controller.go:1264` |
| Watches | rbac.authorization.k8s.io/v1/RoleBinding | `internal/controller/modelcatalog_controller.go:1265` |

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for model-registry-operator

    participant KubernetesAPI as Kubernetes API
    participant controller_manager as controller-manager

    KubernetesAPI-->>+controller_manager: Watch Deployment (informer)
    KubernetesAPI-->>+controller_manager: Watch ConfigMap (informer)
    KubernetesAPI-->>+controller_manager: Watch Secret (informer)
    KubernetesAPI-->>+controller_manager: Watch ServiceAccount (informer)
    KubernetesAPI-->>+controller_manager: Watch Service (informer)
    KubernetesAPI-->>+controller_manager: Watch PersistentVolumeClaim (informer)
    KubernetesAPI-->>+controller_manager: Watch ClusterRoleBinding (informer)
    KubernetesAPI-->>+controller_manager: Watch Role (informer)
    KubernetesAPI-->>+controller_manager: Watch RoleBinding (informer)
    KubernetesAPI-->>+controller_manager: Watch NetworkPolicy (informer)
    KubernetesAPI->>+controller_manager: Watch ModelRegistry (reconcile)
    controller_manager->>KubernetesAPI: Create/Update Service
    controller_manager->>KubernetesAPI: Create/Update ServiceAccount
    controller_manager->>KubernetesAPI: Create/Update Deployment
    controller_manager->>KubernetesAPI: Create/Update Role
    controller_manager->>KubernetesAPI: Create/Update NetworkPolicy
    controller_manager->>KubernetesAPI: Create/Update RoleBinding

    Note over controller_manager: Exposed Services
    Note right of controller_manager: webhook-service:443/TCP []

    Note over KubernetesAPI: Defined CRDs
    Note right of KubernetesAPI: ModelRegistry (modelregistry.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: ModelRegistry (modelregistry.opendatahub.io/v1beta1)
```

### Webhooks

| Name | Type | Path | Failure Policy | Service | Source |
|------|------|------|----------------|---------|--------|
| mmodelregistry.opendatahub.io | mutating | /mutate-modelregistry-opendatahub-io-modelregistry | Fail | system/webhook-service | `config/webhook/manifests.yaml` |
| vmodelregistry.opendatahub.io | validating | /validate-modelregistry-opendatahub-io-modelregistry | Fail | system/webhook-service | `config/webhook/manifests.yaml` |
| mmodelregistry.opendatahub.io | mutating | /mutate-modelregistry-opendatahub-io-modelregistry | fail |  | `internal/webhook/modelregistry_webhook.go` |
| vmodelregistry.opendatahub.io | validating | /validate-modelregistry-opendatahub-io-modelregistry | fail |  | `internal/webhook/modelregistry_webhook.go` |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

