# kube-rbac-proxy: Dataflow

## Controller Watches

No controller watches found.

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for kube-rbac-proxy

    participant KubernetesAPI as Kubernetes API
    participant kube_rbac_proxy as kube-rbac-proxy


    Note over kube_rbac_proxy: Exposed Services
    Note right of kube_rbac_proxy: kube-rbac-proxy:8443/TCP [https]
```

### HTTP Endpoints

| Method | Path | Source |
|--------|------|--------|
| * | / | `cmd/kube-rbac-proxy/app/kube-rbac-proxy.go:322` |

## Configuration

