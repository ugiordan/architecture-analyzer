# llm-d: Network

## Service Map

*3 unique services (7 total, duplicates from test fixtures collapsed).*

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    llm_d["llm-d"]:::component
    llm_d --> svc_0["mooncake-client\nClusterIP: 50052/TCP"]:::svc
    llm_d --> svc_1["mooncake-master-store\nClusterIP: 50051/TCP,8080/TCP,9003/TCP"]:::svc
    llm_d --> svc_2["render\nClusterIP: 8000/TCP"]:::svc
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| mooncake-client | ClusterIP | 50052/TCP | [`helpers/mooncake-client/base/service.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/helpers/mooncake-client/base/service.yaml) |
| mooncake-master-store | ClusterIP | 50051/TCP, 8080/TCP, 9003/TCP | [`helpers/mooncake-master-store/base/service.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/helpers/mooncake-master-store/base/service.yaml) |
| render | ClusterIP | 8000/TCP | [`guides/p2p-kv-cache-sharing/render/service.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/guides/p2p-kv-cache-sharing/render/service.yaml) |
| render | ClusterIP | 8000/TCP | [`guides/p2p-kv-cache-sharing/render/standalone/service.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/guides/p2p-kv-cache-sharing/render/standalone/service.yaml) |
| render | ClusterIP | 8000/TCP | [`guides/precise-prefix-cache-routing/render/service.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/guides/precise-prefix-cache-routing/render/service.yaml) |
| render | ClusterIP | 8000/TCP | [`guides/precise-prefix-cache-routing/render/standalone/service.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/guides/precise-prefix-cache-routing/render/standalone/service.yaml) |
| render | ClusterIP | 8000/TCP | [`guides/wide-ep-lws/render/service.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/guides/wide-ep-lws/render/service.yaml) |

### Ingress / Routing

| Kind | Name | Hosts | Paths | TLS | Source |
|------|------|-------|-------|-----|--------|
| Gateway | llm-d-inference-gateway |  |  | no | [`guides/recipes/gateway/agentgateway/gateway.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/guides/recipes/gateway/agentgateway/gateway.yaml) |
| Gateway | llm-d-inference-gateway |  |  | no | [`guides/recipes/gateway/base/gateway.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/guides/recipes/gateway/base/gateway.yaml) |
| Gateway | llm-d-inference-gateway |  |  | no | [`guides/recipes/gateway/envoy-ai-gateway/gateway.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/guides/recipes/gateway/envoy-ai-gateway/gateway.yaml) |
| Gateway | llm-d-inference-gateway |  |  | no | [`guides/recipes/gateway/istio/gateway.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/guides/recipes/gateway/istio/gateway.yaml) |
| Gateway | llm-d-inference-gateway |  |  | no | [`guides/recipes/gateway/kgateway/gateway.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/guides/recipes/gateway/kgateway/gateway.yaml) |
| HTTPRoute | ${GUIDE_NAME} |  | /v1/completions, /v1/chat/completions, /inference/v1/generate, /v1/completions, /v1/chat/completions, /inference/v1/generate, /v1/completions, /v1/chat/completions, /inference/v1/generate, / | no | [`guides/coord-disaggregation/router/httproute-3-epp.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/guides/coord-disaggregation/router/httproute-3-epp.yaml) |
| HTTPRoute | ${GUIDE_NAME} |  | /v1/completions, /v1/completions, /v1/completions, /v1/chat/completions, /v1/chat/completions, /v1/chat/completions, /inference/v1/generate, /inference/v1/generate, /inference/v1/generate, / | no | [`guides/coord-disaggregation/router/httproute.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/guides/coord-disaggregation/router/httproute.yaml) |
| HTTPRoute | coordinator |  | /v1/completions, /v1/chat/completions, /inference/v1/generate | no | [`guides/coord-disaggregation/coordinator/httproute.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/guides/coord-disaggregation/coordinator/httproute.yaml) |
| HTTPRoute | deepseek-route |  |  | no | [`guides/multi-model-routing/manifests/httproutes.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/guides/multi-model-routing/manifests/httproutes.yaml) |
| HTTPRoute | qwen-route |  |  | no | [`guides/multi-model-routing/manifests/httproutes.yaml`](https://github.com/llm-d/llm-d/blob/080c14d957755da4cc363745638619fc748558f1/guides/multi-model-routing/manifests/httproutes.yaml) |

!!! warning "No Network Policies"
    No NetworkPolicy resources were found in the analyzed sources. Network policies may exist in overlays, Helm values, or cluster-level configurations not captured by static analysis.

