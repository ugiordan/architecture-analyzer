# feast: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

| Type | GVK | Source |
|------|-----|--------|
| For | api/v1/FeatureStore | [`infra/feast-operator/internal/controller/featurestore_controller.go:439`](https://github.com/feast-dev/feast/blob/ff274a6b65df387ef53ffc14b09fad6f07ac27fc/infra/feast-operator/internal/controller/featurestore_controller.go#L439) |
| Owns | /v1/ConfigMap | [`infra/feast-operator/internal/controller/featurestore_controller.go:440`](https://github.com/feast-dev/feast/blob/ff274a6b65df387ef53ffc14b09fad6f07ac27fc/infra/feast-operator/internal/controller/featurestore_controller.go#L440) |
| Owns | /v1/PersistentVolumeClaim | [`infra/feast-operator/internal/controller/featurestore_controller.go:443`](https://github.com/feast-dev/feast/blob/ff274a6b65df387ef53ffc14b09fad6f07ac27fc/infra/feast-operator/internal/controller/featurestore_controller.go#L443) |
| Owns | /v1/Service | [`infra/feast-operator/internal/controller/featurestore_controller.go:442`](https://github.com/feast-dev/feast/blob/ff274a6b65df387ef53ffc14b09fad6f07ac27fc/infra/feast-operator/internal/controller/featurestore_controller.go#L442) |
| Owns | /v1/ServiceAccount | [`infra/feast-operator/internal/controller/featurestore_controller.go:444`](https://github.com/feast-dev/feast/blob/ff274a6b65df387ef53ffc14b09fad6f07ac27fc/infra/feast-operator/internal/controller/featurestore_controller.go#L444) |
| Owns | apps/v1/Deployment | [`infra/feast-operator/internal/controller/featurestore_controller.go:441`](https://github.com/feast-dev/feast/blob/ff274a6b65df387ef53ffc14b09fad6f07ac27fc/infra/feast-operator/internal/controller/featurestore_controller.go#L441) |
| Owns | autoscaling/v2/HorizontalPodAutoscaler | [`infra/feast-operator/internal/controller/featurestore_controller.go:448`](https://github.com/feast-dev/feast/blob/ff274a6b65df387ef53ffc14b09fad6f07ac27fc/infra/feast-operator/internal/controller/featurestore_controller.go#L448) |
| Owns | batch/v1/CronJob | [`infra/feast-operator/internal/controller/featurestore_controller.go:447`](https://github.com/feast-dev/feast/blob/ff274a6b65df387ef53ffc14b09fad6f07ac27fc/infra/feast-operator/internal/controller/featurestore_controller.go#L447) |
| Owns | policy/v1/PodDisruptionBudget | [`infra/feast-operator/internal/controller/featurestore_controller.go:449`](https://github.com/feast-dev/feast/blob/ff274a6b65df387ef53ffc14b09fad6f07ac27fc/infra/feast-operator/internal/controller/featurestore_controller.go#L449) |
| Owns | rbac.authorization.k8s.io/v1/Role | [`infra/feast-operator/internal/controller/featurestore_controller.go:446`](https://github.com/feast-dev/feast/blob/ff274a6b65df387ef53ffc14b09fad6f07ac27fc/infra/feast-operator/internal/controller/featurestore_controller.go#L446) |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | [`infra/feast-operator/internal/controller/featurestore_controller.go:445`](https://github.com/feast-dev/feast/blob/ff274a6b65df387ef53ffc14b09fad6f07ac27fc/infra/feast-operator/internal/controller/featurestore_controller.go#L445) |
| Owns | route/v1/Route | [`infra/feast-operator/internal/controller/featurestore_controller.go:453`](https://github.com/feast-dev/feast/blob/ff274a6b65df387ef53ffc14b09fad6f07ac27fc/infra/feast-operator/internal/controller/featurestore_controller.go#L453) |
| Watches | api/v1/FeatureStore | [`infra/feast-operator/internal/controller/featurestore_controller.go:450`](https://github.com/feast-dev/feast/blob/ff274a6b65df387ef53ffc14b09fad6f07ac27fc/infra/feast-operator/internal/controller/featurestore_controller.go#L450) |

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
| * | /get-online-features | [`go/internal/feast/server/http_server.go:388`](https://github.com/feast-dev/feast/blob/ff274a6b65df387ef53ffc14b09fad6f07ac27fc/go/internal/feast/server/http_server.go#L388) |
| * | /get-online-features | [`go/internal/feast/server/http_server.go:402`](https://github.com/feast-dev/feast/blob/ff274a6b65df387ef53ffc14b09fad6f07ac27fc/go/internal/feast/server/http_server.go#L402) |
| * | /health | [`go/internal/feast/server/http_server.go:389`](https://github.com/feast-dev/feast/blob/ff274a6b65df387ef53ffc14b09fad6f07ac27fc/go/internal/feast/server/http_server.go#L389) |
| * | /health | [`go/internal/feast/server/http_server.go:403`](https://github.com/feast-dev/feast/blob/ff274a6b65df387ef53ffc14b09fad6f07ac27fc/go/internal/feast/server/http_server.go#L403) |
| * | /metrics | [`go/main.go:204`](https://github.com/feast-dev/feast/blob/ff274a6b65df387ef53ffc14b09fad6f07ac27fc/go/main.go#L204) |
| * | /metrics | [`go/main.go:264`](https://github.com/feast-dev/feast/blob/ff274a6b65df387ef53ffc14b09fad6f07ac27fc/go/main.go#L264) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

### Helm

**Chart:** feast v0.66.0

