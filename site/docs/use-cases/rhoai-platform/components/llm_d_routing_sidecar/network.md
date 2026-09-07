# llm-d-routing-sidecar: Network

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    llm_d_routing_sidecar["llm-d-routing-sidecar"]:::component
    llm_d_routing_sidecar --> svc_0["service\nClusterIP: 8080/TCP"]:::svc
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| service | ClusterIP | 8080/TCP | [`deploy/common/service.yaml`](https://github.com/llm-d/llm-d-routing-sidecar/blob/214ed72b3bcd2ea0d66ae2f15d82e0037a726c06/deploy/common/service.yaml) |

### Ingress / Routing

| Kind | Name | Hosts | Paths | TLS | Source |
|------|------|-------|-------|-----|--------|
| Route | route |  |  | yes | [`deploy/openshift/route.yaml`](https://github.com/llm-d/llm-d-routing-sidecar/blob/214ed72b3bcd2ea0d66ae2f15d82e0037a726c06/deploy/openshift/route.yaml) |
| Route | route |  |  | no | [`deploy/openshift/patch-route.yaml`](https://github.com/llm-d/llm-d-routing-sidecar/blob/214ed72b3bcd2ea0d66ae2f15d82e0037a726c06/deploy/openshift/patch-route.yaml) |

!!! warning "No Network Policies"
    No NetworkPolicy resources were found in the analyzed sources. Network policies may exist in overlays, Helm values, or cluster-level configurations not captured by static analysis.

