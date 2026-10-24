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
| * | / | [`cmd/kube-rbac-proxy/app/kube-rbac-proxy.go:347`](https://github.com/red-hat-data-services/kube-rbac-proxy/blob/443f6e28aae8b16dfdf500a282dc58cd256bdefc/cmd/kube-rbac-proxy/app/kube-rbac-proxy.go#L347) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

