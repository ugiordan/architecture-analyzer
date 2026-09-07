# kserve-autogluon-server: Network

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    kserve_autogluon_server["kserve-autogluon-server"]:::component
    kserve_autogluon_server --> svc_0["kserve-controller-manager-metrics-service\nClusterIP: 8443/TCP"]:::svc
    kserve_autogluon_server --> svc_1["kserve-controller-manager-service\nClusterIP: 8443/TCP"]:::svc
    kserve_autogluon_server --> svc_2["kserve-webhook-server-service\nClusterIP: 443/TCP"]:::svc
    kserve_autogluon_server --> svc_3["llmisvc-controller-manager-service\nClusterIP: 8443/TCP"]:::svc
    kserve_autogluon_server --> svc_4["llmisvc-webhook-server-service\nClusterIP: 443/TCP"]:::svc
    kserve_autogluon_server --> svc_5["localmodel-webhook-server-service\nClusterIP: 443/TCP"]:::svc
    kserve_autogluon_server -.-> ext_httpx[["httpx\napi"]]:::ext
    kserve_autogluon_server -.-> ext_requests[["requests\napi"]]:::ext
    kserve_autogluon_server -.-> ext_weaviate[["weaviate\napi"]]:::ext
    kserve_autogluon_server -.-> ext_azure_blob[["azure-blob\nobject-storage"]]:::ext
    kserve_autogluon_server -.-> ext_gcs[["gcs\nobject-storage"]]:::ext
    kserve_autogluon_server -.-> ext_s3[["s3\nobject-storage"]]:::ext
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| kserve-controller-manager-metrics-service | ClusterIP | 8443/TCP | [`kustomize:config/overlays/all`](https://github.com/kserve/kserve-autogluon-server/blob/aad6612fe382c04130a0675500087545cb6edc3f/kustomize:config/overlays/all) |
| kserve-controller-manager-service | ClusterIP | 8443/TCP | [`kustomize:config/overlays/all`](https://github.com/kserve/kserve-autogluon-server/blob/aad6612fe382c04130a0675500087545cb6edc3f/kustomize:config/overlays/all) |
| kserve-webhook-server-service | ClusterIP | 443/TCP | [`kustomize:config/overlays/all`](https://github.com/kserve/kserve-autogluon-server/blob/aad6612fe382c04130a0675500087545cb6edc3f/kustomize:config/overlays/all) |
| llmisvc-controller-manager-service | ClusterIP | 8443/TCP | [`kustomize:config/overlays/all`](https://github.com/kserve/kserve-autogluon-server/blob/aad6612fe382c04130a0675500087545cb6edc3f/kustomize:config/overlays/all) |
| llmisvc-webhook-server-service | ClusterIP | 443/TCP | [`kustomize:config/overlays/all`](https://github.com/kserve/kserve-autogluon-server/blob/aad6612fe382c04130a0675500087545cb6edc3f/kustomize:config/overlays/all) |
| localmodel-webhook-server-service | ClusterIP | 443/TCP | [`kustomize:config/overlays/all`](https://github.com/kserve/kserve-autogluon-server/blob/aad6612fe382c04130a0675500087545cb6edc3f/kustomize:config/overlays/all) |

### Ingress / Routing

| Kind | Name | Hosts | Paths | TLS | Source |
|------|------|-------|-------|-----|--------|
| HTTPRoute | rbac-inferred |  |  | no | [`rbac/kserve-manager-role`](https://github.com/kserve/kserve-autogluon-server/blob/aad6612fe382c04130a0675500087545cb6edc3f/rbac/kserve-manager-role) |
| Ingress | rbac-inferred |  |  | no | [`rbac/kserve-manager-role`](https://github.com/kserve/kserve-autogluon-server/blob/aad6612fe382c04130a0675500087545cb6edc3f/rbac/kserve-manager-role) |
| VirtualService | rbac-inferred |  |  | no | [`rbac/kserve-manager-role`](https://github.com/kserve/kserve-autogluon-server/blob/aad6612fe382c04130a0675500087545cb6edc3f/rbac/kserve-manager-role) |

!!! warning "No Network Policies"
    No NetworkPolicy resources were found in the analyzed sources. Network policies may exist in overlays, Helm values, or cluster-level configurations not captured by static analysis.

