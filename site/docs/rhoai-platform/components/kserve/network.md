# kserve: Network

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff

    kserve["kserve"]:::component
    kserve --> svc_0["llmisvc-controller-manager-service\nClusterIP: 8443/TCP"]:::svc
    kserve --> svc_1["kserve-controller-manager-service\nClusterIP: 8443/TCP"]:::svc
    kserve --> svc_2["llmisvc-webhook-server-service\nClusterIP: 443/TCP"]:::svc
    kserve --> svc_3["localmodel-webhook-server-service\nClusterIP: 443/TCP"]:::svc
    kserve --> svc_4["kserve-webhook-server-service\nClusterIP: 443/TCP"]:::svc
    kserve --> svc_5["webhook-service\nClusterIP: 443/TCP"]:::svc
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| llmisvc-controller-manager-service | ClusterIP | 8443/TCP | `config/llmisvc/service.yaml` |
| kserve-controller-manager-service | ClusterIP | 8443/TCP | `config/manager/service.yaml` |
| llmisvc-webhook-server-service | ClusterIP | 443/TCP | `config/webhook/llmisvc/service.yaml` |
| localmodel-webhook-server-service | ClusterIP | 443/TCP | `config/webhook/localmodel/service.yaml` |
| kserve-webhook-server-service | ClusterIP | 443/TCP | `config/webhook/service.yaml` |
| webhook-service | ClusterIP | 443/TCP | `test/webhooks/service.yaml` |

### Ingress / Routing

| Kind | Name | Hosts | Paths | TLS | Source |
|------|------|-------|-------|-----|--------|
| Gateway | knative-ingress-gateway |  |  | no | `docs/openshift/serverless/gateways.yaml` |
| Gateway | knative-local-gateway |  |  | no | `docs/openshift/serverless/gateways.yaml` |
| Gateway | ai-gateway |  |  | no | `docs/samples/llmisvc/e2e-gpt-oss/gateway.yaml` |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| kserve-controller-manager |  | `config/overlays/odh/network-policies.yaml` |

