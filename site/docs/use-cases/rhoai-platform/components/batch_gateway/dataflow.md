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
| * | / | [`internal/apiserver/common/rest.go:74`](https://github.com/red-hat-data-services/batch-gateway/blob/105c8c1a7bd61e414444133a5437e12825475d6a/internal/apiserver/common/rest.go#L74) |
| * | /health | [`cmd/batch-gc/main.go:178`](https://github.com/red-hat-data-services/batch-gateway/blob/105c8c1a7bd61e414444133a5437e12825475d6a/cmd/batch-gc/main.go#L178) |
| * | /metrics | [`cmd/batch-gc/main.go:177`](https://github.com/red-hat-data-services/batch-gateway/blob/105c8c1a7bd61e414444133a5437e12825475d6a/cmd/batch-gc/main.go#L177) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

### Helm

**Chart:** batch-gateway v0.1.0

