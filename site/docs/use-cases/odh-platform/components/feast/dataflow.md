# feast: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

| Type | GVK | Source |
|------|-----|--------|
| For | api/v1/FeatureStore | [`infra/feast-operator/internal/controller/featurestore_controller.go:414`](https://github.com/feast-dev/feast/blob/6ad332bf6c06499319483b0d3c87d96165c41d53/infra/feast-operator/internal/controller/featurestore_controller.go#L414) |
| Owns | /v1/ConfigMap | [`infra/feast-operator/internal/controller/featurestore_controller.go:415`](https://github.com/feast-dev/feast/blob/6ad332bf6c06499319483b0d3c87d96165c41d53/infra/feast-operator/internal/controller/featurestore_controller.go#L415) |
| Owns | /v1/PersistentVolumeClaim | [`infra/feast-operator/internal/controller/featurestore_controller.go:418`](https://github.com/feast-dev/feast/blob/6ad332bf6c06499319483b0d3c87d96165c41d53/infra/feast-operator/internal/controller/featurestore_controller.go#L418) |
| Owns | /v1/Service | [`infra/feast-operator/internal/controller/featurestore_controller.go:417`](https://github.com/feast-dev/feast/blob/6ad332bf6c06499319483b0d3c87d96165c41d53/infra/feast-operator/internal/controller/featurestore_controller.go#L417) |
| Owns | /v1/ServiceAccount | [`infra/feast-operator/internal/controller/featurestore_controller.go:419`](https://github.com/feast-dev/feast/blob/6ad332bf6c06499319483b0d3c87d96165c41d53/infra/feast-operator/internal/controller/featurestore_controller.go#L419) |
| Owns | apps/v1/Deployment | [`infra/feast-operator/internal/controller/featurestore_controller.go:416`](https://github.com/feast-dev/feast/blob/6ad332bf6c06499319483b0d3c87d96165c41d53/infra/feast-operator/internal/controller/featurestore_controller.go#L416) |
| Owns | autoscaling/v2/HorizontalPodAutoscaler | [`infra/feast-operator/internal/controller/featurestore_controller.go:423`](https://github.com/feast-dev/feast/blob/6ad332bf6c06499319483b0d3c87d96165c41d53/infra/feast-operator/internal/controller/featurestore_controller.go#L423) |
| Owns | batch/v1/CronJob | [`infra/feast-operator/internal/controller/featurestore_controller.go:422`](https://github.com/feast-dev/feast/blob/6ad332bf6c06499319483b0d3c87d96165c41d53/infra/feast-operator/internal/controller/featurestore_controller.go#L422) |
| Owns | policy/v1/PodDisruptionBudget | [`infra/feast-operator/internal/controller/featurestore_controller.go:424`](https://github.com/feast-dev/feast/blob/6ad332bf6c06499319483b0d3c87d96165c41d53/infra/feast-operator/internal/controller/featurestore_controller.go#L424) |
| Owns | rbac.authorization.k8s.io/v1/Role | [`infra/feast-operator/internal/controller/featurestore_controller.go:421`](https://github.com/feast-dev/feast/blob/6ad332bf6c06499319483b0d3c87d96165c41d53/infra/feast-operator/internal/controller/featurestore_controller.go#L421) |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | [`infra/feast-operator/internal/controller/featurestore_controller.go:420`](https://github.com/feast-dev/feast/blob/6ad332bf6c06499319483b0d3c87d96165c41d53/infra/feast-operator/internal/controller/featurestore_controller.go#L420) |
| Owns | route/v1/Route | [`infra/feast-operator/internal/controller/featurestore_controller.go:428`](https://github.com/feast-dev/feast/blob/6ad332bf6c06499319483b0d3c87d96165c41d53/infra/feast-operator/internal/controller/featurestore_controller.go#L428) |
| Watches | api/v1/FeatureStore | [`infra/feast-operator/internal/controller/featurestore_controller.go:425`](https://github.com/feast-dev/feast/blob/6ad332bf6c06499319483b0d3c87d96165c41d53/infra/feast-operator/internal/controller/featurestore_controller.go#L425) |

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for feast

    participant KubernetesAPI as Kubernetes API
    participant controller_manager as controller-manager

    KubernetesAPI->>+controller_manager: Watch FeatureStore (reconcile)
    controller_manager->>KubernetesAPI: Create/Update ConfigMap
    controller_manager->>KubernetesAPI: Create/Update PersistentVolumeClaim
    controller_manager->>KubernetesAPI: Create/Update Service
    controller_manager->>KubernetesAPI: Create/Update ServiceAccount
    controller_manager->>KubernetesAPI: Create/Update Deployment
    controller_manager->>KubernetesAPI: Create/Update HorizontalPodAutoscaler
    controller_manager->>KubernetesAPI: Create/Update CronJob
    controller_manager->>KubernetesAPI: Create/Update PodDisruptionBudget
    controller_manager->>KubernetesAPI: Create/Update Role
    controller_manager->>KubernetesAPI: Create/Update RoleBinding
    controller_manager->>KubernetesAPI: Create/Update Route
    KubernetesAPI-->>+controller_manager: Watch FeatureStore (informer)

    Note over controller_manager: Exposed Services
    Note right of controller_manager: uvicorn-server:6566/TCP []
```

### HTTP Endpoints

| Method | Path | Source |
|--------|------|--------|
| * | /get-online-features | [`go/internal/feast/server/http_server.go:388`](https://github.com/feast-dev/feast/blob/6ad332bf6c06499319483b0d3c87d96165c41d53/go/internal/feast/server/http_server.go#L388) |
| * | /get-online-features | [`go/internal/feast/server/http_server.go:402`](https://github.com/feast-dev/feast/blob/6ad332bf6c06499319483b0d3c87d96165c41d53/go/internal/feast/server/http_server.go#L402) |
| * | /health | [`go/internal/feast/server/http_server.go:389`](https://github.com/feast-dev/feast/blob/6ad332bf6c06499319483b0d3c87d96165c41d53/go/internal/feast/server/http_server.go#L389) |
| * | /health | [`go/internal/feast/server/http_server.go:403`](https://github.com/feast-dev/feast/blob/6ad332bf6c06499319483b0d3c87d96165c41d53/go/internal/feast/server/http_server.go#L403) |
| * | /metrics | [`go/main.go:204`](https://github.com/feast-dev/feast/blob/6ad332bf6c06499319483b0d3c87d96165c41d53/go/main.go#L204) |
| * | /metrics | [`go/main.go:264`](https://github.com/feast-dev/feast/blob/6ad332bf6c06499319483b0d3c87d96165c41d53/go/main.go#L264) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

### Helm

**Chart:** feast v0.66.0

