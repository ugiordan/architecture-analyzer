# batch-gateway: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

No controller watches found in analyzed sources.

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for batch-gateway

    participant KubernetesAPI as Kubernetes API
    participant batch_gateway as batch-gateway
```

### HTTP Endpoints

| Method | Path | Source |
|--------|------|--------|
| * | / | [`internal/apiserver/common/rest.go:74`](https://github.com/llm-d/batch-gateway/blob/80c35a41ad4bdf92ef3fa53f25bb8d6ac74e49d5/internal/apiserver/common/rest.go#L74) |
| * | /health | [`cmd/batch-gc/main.go:164`](https://github.com/llm-d/batch-gateway/blob/80c35a41ad4bdf92ef3fa53f25bb8d6ac74e49d5/cmd/batch-gc/main.go#L164) |
| * | /metrics | [`cmd/batch-gc/main.go:163`](https://github.com/llm-d/batch-gateway/blob/80c35a41ad4bdf92ef3fa53f25bb8d6ac74e49d5/cmd/batch-gc/main.go#L163) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

### Helm

**Chart:** batch-gateway v0.1.0

