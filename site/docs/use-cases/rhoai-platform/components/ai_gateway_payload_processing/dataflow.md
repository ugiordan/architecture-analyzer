# ai-gateway-payload-processing: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

| Type | GVK | Source |
|------|-----|--------|
| For | inference/v1alpha1/ExternalModel | [`pkg/controller/externalmodel/reconciler.go:318`](https://github.com/red-hat-data-services/ai-gateway-payload-processing/blob/0680ab42966f1cb0c61b73ee9530bc02b413a719/pkg/controller/externalmodel/reconciler.go#L318) |
| For | inference/v1alpha1/ExternalProvider | [`pkg/controller/externalprovider/reconciler.go:405`](https://github.com/red-hat-data-services/ai-gateway-payload-processing/blob/0680ab42966f1cb0c61b73ee9530bc02b413a719/pkg/controller/externalprovider/reconciler.go#L405) |
| Owns | /v1/Service | [`pkg/controller/externalprovider/reconciler.go:406`](https://github.com/red-hat-data-services/ai-gateway-payload-processing/blob/0680ab42966f1cb0c61b73ee9530bc02b413a719/pkg/controller/externalprovider/reconciler.go#L406) |
| Owns | apis/v1/HTTPRoute | [`pkg/controller/externalmodel/reconciler.go:319`](https://github.com/red-hat-data-services/ai-gateway-payload-processing/blob/0680ab42966f1cb0c61b73ee9530bc02b413a719/pkg/controller/externalmodel/reconciler.go#L319) |
| Watches | inference/v1alpha1/ExternalModel | [`pkg/controller/externalprovider/reconciler.go:409`](https://github.com/red-hat-data-services/ai-gateway-payload-processing/blob/0680ab42966f1cb0c61b73ee9530bc02b413a719/pkg/controller/externalprovider/reconciler.go#L409) |
| Watches | inference/v1alpha1/ExternalProvider | [`pkg/controller/externalmodel/reconciler.go:320`](https://github.com/red-hat-data-services/ai-gateway-payload-processing/blob/0680ab42966f1cb0c61b73ee9530bc02b413a719/pkg/controller/externalmodel/reconciler.go#L320) |

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for ai-gateway-payload-processing

    participant KubernetesAPI as Kubernetes API
    participant ai_gateway_payload_processing as ai-gateway-payload-processing

    KubernetesAPI->>+ai_gateway_payload_processing: Watch ExternalModel (reconcile)
    KubernetesAPI->>+ai_gateway_payload_processing: Watch ExternalProvider (reconcile)
    ai_gateway_payload_processing->>KubernetesAPI: Create/Update Service
    ai_gateway_payload_processing->>KubernetesAPI: Create/Update HTTPRoute
    KubernetesAPI-->>+ai_gateway_payload_processing: Watch ExternalModel (informer)
    KubernetesAPI-->>+ai_gateway_payload_processing: Watch ExternalProvider (informer)

    Note over KubernetesAPI: Defined CRDs
    Note right of KubernetesAPI: ExternalModel (inference.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: ExternalProvider (inference.opendatahub.io/v1alpha1)
```

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

### Helm

**Chart:** payload-processing v0.1.0

