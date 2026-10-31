# kube-rbac-proxy: Dataflow

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
| * | / | [`cmd/kube-rbac-proxy/app/kube-rbac-proxy.go:361`](https://github.com/red-hat-data-services/kube-rbac-proxy/blob/62baa2d5220ae25e5f7bab005fe3d5a442513292/cmd/kube-rbac-proxy/app/kube-rbac-proxy.go#L361) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

