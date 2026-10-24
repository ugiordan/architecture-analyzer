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
| * | / | [`internal/evalhub_mcp/server/server.go:241`](https://github.com/opendatahub-io/eval-hub/blob/79066e0fe85dd4bdb8c9b699ae29dbd22e1591d6/internal/evalhub_mcp/server/server.go#L241) |
| * | /health | [`internal/evalhub_mcp/server/server.go:236`](https://github.com/opendatahub-io/eval-hub/blob/79066e0fe85dd4bdb8c9b699ae29dbd22e1591d6/internal/evalhub_mcp/server/server.go#L236) |
| * | /metrics | [`internal/eval_hub/server/metrics_server.go:28`](https://github.com/opendatahub-io/eval-hub/blob/79066e0fe85dd4bdb8c9b699ae29dbd22e1591d6/internal/eval_hub/server/metrics_server.go#L28) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

