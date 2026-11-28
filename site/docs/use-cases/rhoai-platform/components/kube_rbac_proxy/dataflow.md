# kube-rbac-proxy: Dataflow

[Overview](index.md) · [Agent Context](agent-context.md) · [Extracted Data](details.md) · [Network](network.md) · [RBAC](rbac.md) · [Security](security.md) · [Dataflow](dataflow.md)

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

No controller watches found in analyzed sources.

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for kube-rbac-proxy

    participant KubernetesAPI as Kubernetes API
    participant kube_rbac_proxy as kube-rbac-proxy
```

### HTTP Endpoints

| Method | Path | Source |
|--------|------|--------|
| * | / | [`cmd/kube-rbac-proxy/app/kube-rbac-proxy.go:420`](https://github.com/red-hat-data-services/kube-rbac-proxy/blob/6026bed65dea3d35af239d0f7177832679b44f4b/cmd/kube-rbac-proxy/app/kube-rbac-proxy.go#L420) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

