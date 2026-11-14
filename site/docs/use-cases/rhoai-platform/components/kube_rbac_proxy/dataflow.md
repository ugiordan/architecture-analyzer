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
| * | / | [`cmd/kube-rbac-proxy/app/kube-rbac-proxy.go:369`](https://github.com/red-hat-data-services/kube-rbac-proxy/blob/b2c529c74a63a29b830dd28bbed3989766ea9aa4/cmd/kube-rbac-proxy/app/kube-rbac-proxy.go#L369) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

