# odh-model-controller: Network

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff

    odh_model_controller["odh-model-controller"]:::component
    odh_model_controller --> svc_0["odh-model-controller-webhook-service\nClusterIP: 443/TCP"]:::svc
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| odh-model-controller-webhook-service | ClusterIP | 443/TCP | `config/webhook/service.yaml` |

!!! warning "No Network Policies"
    No NetworkPolicy resources found. All pod-to-pod traffic is allowed by default.

