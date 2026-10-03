# mcp-lifecycle-module-operator: Network

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    mcp_lifecycle_module_operator["mcp-lifecycle-module-operator"]:::component
    mcp_lifecycle_module_operator --> svc_0["webhook-service\nClusterIP: 443/TCP"]:::svc
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| webhook-service | ClusterIP | 443/TCP | [`config/manifests/mcp-lifecycle-operator/webhook/service.yaml`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/ed1187c1c5382db14bac17ee53f5a8a1dd5dbd7d/config/manifests/mcp-lifecycle-operator/webhook/service.yaml) |

### Ingress / Routing

| Kind | Name | Hosts | Paths | TLS | Source |
|------|------|-------|-------|-----|--------|
| Gateway | rbac-inferred |  |  | no | [`rbac/manager-role`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/ed1187c1c5382db14bac17ee53f5a8a1dd5dbd7d/rbac/manager-role) |
| HTTPRoute | rbac-inferred |  |  | no | [`rbac/manager-role`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/ed1187c1c5382db14bac17ee53f5a8a1dd5dbd7d/rbac/manager-role) |

!!! warning "No Network Policies"
    No NetworkPolicy resources were found in the analyzed sources. Network policies may exist in overlays, Helm values, or cluster-level configurations not captured by static analysis.

