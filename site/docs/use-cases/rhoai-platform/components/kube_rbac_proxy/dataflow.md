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
| * | / | [`cmd/kube-rbac-proxy/app/kube-rbac-proxy.go:341`](https://github.com/brancz/kube-rbac-proxy/blob/a20020af4f1d890f8f12dfea8a23f03054ddccc7/cmd/kube-rbac-proxy/app/kube-rbac-proxy.go#L341) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

