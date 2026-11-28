# model-registry: Network

[Overview](index.md) · [Agent Context](agent-context.md) · [Extracted Data](details.md) · [Network](network.md) · [RBAC](rbac.md) · [Security](security.md) · [Cache](cache.md) · [Dataflow](dataflow.md)

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    model_registry["model-registry"]:::component
    model_registry --> svc_0["model-catalog\nClusterIP: 8080/TCP"]:::svc
    model_registry -.-> ext_aiohttp[["aiohttp\napi"]]:::ext
    model_registry -.-> ext_requests[["requests\napi"]]:::ext
    model_registry -.-> ext_urllib[["urllib\napi"]]:::ext
    model_registry -.-> ext_s3[["s3\nobject-storage"]]:::ext
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| model-catalog | ClusterIP | 8080/TCP | [`manifests/kustomize/options/catalog/base/service.yaml`](https://github.com/opendatahub-io/model-registry/blob/51d3c2fd69c5d1b191e8c4a4a8c29b644ef6555a/manifests/kustomize/options/catalog/base/service.yaml) |

!!! warning "No Network Policies"
    No NetworkPolicy resources were found in the analyzed sources. Network policies may exist in overlays, Helm values, or cluster-level configurations not captured by static analysis.

