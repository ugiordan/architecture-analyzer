# kuberay: Network

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    kuberay["kuberay"]:::component
    kuberay --> svc_0["kuberay-operator\nClusterIP: 8080/TCP"]:::svc
    kuberay --> svc_1["webhook-service\nClusterIP: 443/TCP"]:::svc
    kuberay -.-> ext_requests[["requests\napi"]]:::ext
    kuberay -.-> ext_urllib[["urllib\napi"]]:::ext
    kuberay -.-> ext_grpc[["grpc\ngrpc"]]:::ext
    kuberay -.-> ext_azure_blob[["azure-blob\nobject-storage"]]:::ext
    kuberay -.-> ext_gcs[["gcs\nobject-storage"]]:::ext
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| kuberay-operator | ClusterIP | 8080/TCP | [`ray-operator/config/manager/service.yaml`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/config/manager/service.yaml) |
| webhook-service | ClusterIP | 443/TCP | [`ray-operator/config/webhook/service.yaml`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/config/webhook/service.yaml) |

!!! warning "No Network Policies"
    No NetworkPolicy resources were found in the analyzed sources. Network policies may exist in overlays, Helm values, or cluster-level configurations not captured by static analysis.

