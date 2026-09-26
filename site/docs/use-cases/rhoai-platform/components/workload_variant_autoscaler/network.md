# workload-variant-autoscaler: Network

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
| controller-manager-metrics-service | ClusterIP | 8443/TCP | [`legacy/config/base/manager/service.yaml`](https://github.com/red-hat-data-services/workload-variant-autoscaler/blob/af7d918a087f9db38e19120774b60b0378615662/legacy/config/base/manager/service.yaml) |

!!! warning "No Network Policies"
    No NetworkPolicy resources were found in the analyzed sources. Network policies may exist in overlays, Helm values, or cluster-level configurations not captured by static analysis.

