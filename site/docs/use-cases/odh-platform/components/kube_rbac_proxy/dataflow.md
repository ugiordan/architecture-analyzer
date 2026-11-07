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
| * | / | [`cmd/kube-rbac-proxy/app/kube-rbac-proxy.go:361`](https://github.com/opendatahub-io/kube-rbac-proxy/blob/8fe9016316d294aa3f4ff56d60b33c72cd9c4871/cmd/kube-rbac-proxy/app/kube-rbac-proxy.go#L361) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

