# workload-variant-autoscaler: Network

[Overview](index.md) · [Agent Context](agent-context.md) · [Extracted Data](details.md) · [Network](network.md) · [RBAC](rbac.md) · [Security](security.md) · [Cache](cache.md) · [Dataflow](dataflow.md)

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    workload_variant_autoscaler["workload-variant-autoscaler"]:::component
    workload_variant_autoscaler --> svc_0["controller-manager-metrics-service\nClusterIP: 8443/TCP"]:::svc
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| controller-manager-metrics-service | ClusterIP | 8443/TCP | [`legacy/config/base/manager/service.yaml`](https://github.com/opendatahub-io/workload-variant-autoscaler/blob/32fa863c73641afbd8ff9d6eb8244ec68b67c04b/legacy/config/base/manager/service.yaml) |

!!! warning "No Network Policies"
    No NetworkPolicy resources were found in the analyzed sources. Network policies may exist in overlays, Helm values, or cluster-level configurations not captured by static analysis.

