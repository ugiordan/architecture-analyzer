# model-registry: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

| Type | GVK | Source |
|------|-----|--------|
| For | serving/v1beta1/InferenceService | [`cmd/controller/internal/controllers/inferenceservice_controller.go:44`](https://github.com/kubeflow/model-registry/blob/9b5a0587297f1dd94d08338cf4241df725a29592/cmd/controller/internal/controllers/inferenceservice_controller.go#L44) |
| For | serving/v1beta1/InferenceService | [`pkg/inferenceservice-controller/controller.go:251`](https://github.com/kubeflow/model-registry/blob/9b5a0587297f1dd94d08338cf4241df725a29592/pkg/inferenceservice-controller/controller.go#L251) |

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for model-registry

    participant KubernetesAPI as Kubernetes API
    participant controller_manager as controller-manager
    participant minio as minio
    participant model_catalog_server as model-catalog-server

    KubernetesAPI->>+controller_manager: Watch InferenceService (reconcile)
    KubernetesAPI->>+controller_manager: Watch InferenceService (reconcile)

    Note over controller_manager: Exposed Services
    Note right of controller_manager: model-catalog:8080/TCP [http-api]
```

### HTTP Endpoints

| Method | Path | Source |
|--------|------|--------|
| GET | /healthcheck | [`clients/ui/bff/openapi/swagger.yaml`](https://github.com/kubeflow/model-registry/blob/9b5a0587297f1dd94d08338cf4241df725a29592/clients/ui/bff/openapi/swagger.yaml) |
| GET | /healthcheck | [`clients/ui/bff/openapi/swagger.json`](https://github.com/kubeflow/model-registry/blob/9b5a0587297f1dd94d08338cf4241df725a29592/clients/ui/bff/openapi/swagger.json) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

### ConfigMaps

| Name | Data Keys | Source |
|------|-----------|--------|
| auth-proxy-config | nginx.conf | [`manifests/kustomize/options/ui/overlays/standalone/auth-proxy-configmap.yaml`](https://github.com/kubeflow/model-registry/blob/9b5a0587297f1dd94d08338cf4241df725a29592/manifests/kustomize/options/ui/overlays/standalone/auth-proxy-configmap.yaml) |
| model-registry-configmap | MODEL_REGISTRY_DATA_STORE_TYPE, MODEL_REGISTRY_REST_SERVICE_HOST, MODEL_REGISTRY_REST_SERVICE_PORT | [`manifests/kustomize/base/model-registry-configmap.yaml`](https://github.com/kubeflow/model-registry/blob/9b5a0587297f1dd94d08338cf4241df725a29592/manifests/kustomize/base/model-registry-configmap.yaml) |

