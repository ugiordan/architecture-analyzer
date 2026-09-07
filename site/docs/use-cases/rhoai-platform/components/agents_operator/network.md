# agents-operator: Network

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    agents_operator["agents-operator"]:::component
    agents_operator --> svc_0["bundle-service\nClusterIP: 8080/TCP"]:::svc
    agents_operator --> svc_1["webhook-service\nClusterIP: 443/TCP"]:::svc
    agents_operator -.-> ext_urllib[["urllib\napi"]]:::ext
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| bundle-service | ClusterIP | 8080/TCP | [`kagenti-operator/config/bundleservice/service.yaml`](https://github.com/red-hat-data-services/agents-operator/blob/cea595935e348e2f2c4bffc5481350fd8509d2d1/kagenti-operator/config/bundleservice/service.yaml) |
| webhook-service | ClusterIP | 443/TCP | [`kagenti-operator/config/webhook/service.yaml`](https://github.com/red-hat-data-services/agents-operator/blob/cea595935e348e2f2c4bffc5481350fd8509d2d1/kagenti-operator/config/webhook/service.yaml) |

### Ingress / Routing

| Kind | Name | Hosts | Paths | TLS | Source |
|------|------|-------|-------|-----|--------|
| HTTPRoute | token-broker-oauth-callback | token-broker.localtest.me | /oauth/callback | no | [`token-broker/deploy/04-httproute.yaml`](https://github.com/red-hat-data-services/agents-operator/blob/cea595935e348e2f2c4bffc5481350fd8509d2d1/token-broker/deploy/04-httproute.yaml) |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| bundle-service | Ingress, Egress | [`kagenti-operator/config/bundleservice/networkpolicy.yaml`](https://github.com/red-hat-data-services/agents-operator/blob/cea595935e348e2f2c4bffc5481350fd8509d2d1/kagenti-operator/config/bundleservice/networkpolicy.yaml) |

## Network Policy Graph

Visual representation of NetworkPolicy rules. Ingress rules show what traffic is allowed into pods, egress rules show what traffic is allowed out.

```mermaid
graph LR
    classDef policy fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef pod fill:#3498db,stroke:#2980b9,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff

    agents_operator["agents-operator\nPods"]:::pod
    np_0_bundle_service{{"bundle-service\nIngress, Egress"}}:::policy
    np_0_bundle_service --> agents_operator
```

