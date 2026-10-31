# eval-hub: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

No controller watches found in analyzed sources.

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for eval-hub

    participant KubernetesAPI as Kubernetes API
    participant eval_hub as eval-hub
```

### HTTP Endpoints

| Method | Path | Source |
|--------|------|--------|
| * | / | [`internal/evalhub_mcp/server/server.go:246`](https://github.com/red-hat-data-services/eval-hub/blob/df68e322915da44c2805191bfab9b1a474ff5676/internal/evalhub_mcp/server/server.go#L246) |
| * | /health | [`internal/evalhub_mcp/server/server.go:241`](https://github.com/red-hat-data-services/eval-hub/blob/df68e322915da44c2805191bfab9b1a474ff5676/internal/evalhub_mcp/server/server.go#L241) |
| * | /metrics | [`internal/eval_hub/server/metrics_server.go:28`](https://github.com/red-hat-data-services/eval-hub/blob/df68e322915da44c2805191bfab9b1a474ff5676/internal/eval_hub/server/metrics_server.go#L28) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

