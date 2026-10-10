# llm-d: Network

## Service Map

*4 unique services (9 total, duplicates from test fixtures collapsed).*

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    llm_d["llm-d"]:::component
    llm_d --> svc_0["decode\nClusterIP: 6379/TCP,8000/TCP"]:::svc
    llm_d --> svc_1["mooncake-client\nClusterIP: 50052/TCP"]:::svc
    llm_d --> svc_2["mooncake-master-store\nClusterIP: 50051/TCP,8080/TCP,9003/TCP"]:::svc
    llm_d --> svc_3["render\nClusterIP: 8000/TCP"]:::svc
    llm_d -.-> ext_urllib[["urllib\napi"]]:::ext
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| decode | ClusterIP | 8000/TCP, 6379/TCP | [`guides/tiered-prefix-cache/modelserver/tpu/v7/vllm/native/cpu/multi-host/service.yaml`](https://github.com/llm-d/llm-d/blob/c690e9acdab3f1c426e3e03ad89de99ebc56bc50/guides/tiered-prefix-cache/modelserver/tpu/v7/vllm/native/cpu/multi-host/service.yaml) |
| mooncake-client | ClusterIP | 50052/TCP | [`helpers/mooncake-client/base/service.yaml`](https://github.com/llm-d/llm-d/blob/c690e9acdab3f1c426e3e03ad89de99ebc56bc50/helpers/mooncake-client/base/service.yaml) |
| mooncake-master-store | ClusterIP | 50051/TCP, 8080/TCP, 9003/TCP | [`helpers/mooncake-master-store/base/service.yaml`](https://github.com/llm-d/llm-d/blob/c690e9acdab3f1c426e3e03ad89de99ebc56bc50/helpers/mooncake-master-store/base/service.yaml) |
| render | ClusterIP | 8000/TCP | [`guides/coord-disaggregation/render/service.yaml`](https://github.com/llm-d/llm-d/blob/c690e9acdab3f1c426e3e03ad89de99ebc56bc50/guides/coord-disaggregation/render/service.yaml) |
| render | ClusterIP | 8000/TCP | [`guides/models/glm-5-2/render/service.yaml`](https://github.com/llm-d/llm-d/blob/c690e9acdab3f1c426e3e03ad89de99ebc56bc50/guides/models/glm-5-2/render/service.yaml) |
| render | ClusterIP | 8000/TCP | [`guides/p2p-kv-cache-sharing/render/service.yaml`](https://github.com/llm-d/llm-d/blob/c690e9acdab3f1c426e3e03ad89de99ebc56bc50/guides/p2p-kv-cache-sharing/render/service.yaml) |
| render | ClusterIP | 8000/TCP | [`guides/p2p-kv-cache-sharing/render/standalone/service.yaml`](https://github.com/llm-d/llm-d/blob/c690e9acdab3f1c426e3e03ad89de99ebc56bc50/guides/p2p-kv-cache-sharing/render/standalone/service.yaml) |
| render | ClusterIP | 8000/TCP | [`guides/precise-prefix-cache-routing/render/service.yaml`](https://github.com/llm-d/llm-d/blob/c690e9acdab3f1c426e3e03ad89de99ebc56bc50/guides/precise-prefix-cache-routing/render/service.yaml) |
| render | ClusterIP | 8000/TCP | [`guides/precise-prefix-cache-routing/render/standalone/service.yaml`](https://github.com/llm-d/llm-d/blob/c690e9acdab3f1c426e3e03ad89de99ebc56bc50/guides/precise-prefix-cache-routing/render/standalone/service.yaml) |

### Ingress / Routing

| Kind | Name | Hosts | Paths | TLS | Source |
|------|------|-------|-------|-----|--------|
| Gateway | llm-d-inference-gateway |  |  | no | [`guides/recipes/gateway/agentgateway/gateway.yaml`](https://github.com/llm-d/llm-d/blob/c690e9acdab3f1c426e3e03ad89de99ebc56bc50/guides/recipes/gateway/agentgateway/gateway.yaml) |
| Gateway | llm-d-inference-gateway |  |  | no | [`guides/recipes/gateway/base/gateway.yaml`](https://github.com/llm-d/llm-d/blob/c690e9acdab3f1c426e3e03ad89de99ebc56bc50/guides/recipes/gateway/base/gateway.yaml) |
| Gateway | llm-d-inference-gateway |  |  | no | [`guides/recipes/gateway/envoy-ai-gateway/gateway.yaml`](https://github.com/llm-d/llm-d/blob/c690e9acdab3f1c426e3e03ad89de99ebc56bc50/guides/recipes/gateway/envoy-ai-gateway/gateway.yaml) |
| Gateway | llm-d-inference-gateway |  |  | no | [`guides/recipes/gateway/istio/gateway.yaml`](https://github.com/llm-d/llm-d/blob/c690e9acdab3f1c426e3e03ad89de99ebc56bc50/guides/recipes/gateway/istio/gateway.yaml) |
| Gateway | llm-d-inference-gateway |  |  | no | [`guides/recipes/gateway/kgateway/gateway.yaml`](https://github.com/llm-d/llm-d/blob/c690e9acdab3f1c426e3e03ad89de99ebc56bc50/guides/recipes/gateway/kgateway/gateway.yaml) |
| HTTPRoute | ${GUIDE_NAME} |  | /v1/completions, /v1/completions, /v1/completions, /v1/chat/completions, /v1/chat/completions, /v1/chat/completions, /inference/v1/generate, /inference/v1/generate, /inference/v1/generate, / | no | [`guides/coord-disaggregation/router/httproute/base/httproute.yaml`](https://github.com/llm-d/llm-d/blob/c690e9acdab3f1c426e3e03ad89de99ebc56bc50/guides/coord-disaggregation/router/httproute/base/httproute.yaml) |
| HTTPRoute | ${GUIDE_NAME} |  | /v1/completions, /v1/chat/completions, /inference/v1/generate, /v1/completions, /v1/chat/completions, /inference/v1/generate, /v1/completions, /v1/chat/completions, /inference/v1/generate, / | no | [`guides/coord-disaggregation/router/httproute-3-epp/base/httproute.yaml`](https://github.com/llm-d/llm-d/blob/c690e9acdab3f1c426e3e03ad89de99ebc56bc50/guides/coord-disaggregation/router/httproute-3-epp/base/httproute.yaml) |
| HTTPRoute | agentic-api |  | /v1/responses, /v1/conversations, /v1/messages, /v1/models | no | [`guides/agentic-api/manifests/gateway/routes.yaml`](https://github.com/llm-d/llm-d/blob/c690e9acdab3f1c426e3e03ad89de99ebc56bc50/guides/agentic-api/manifests/gateway/routes.yaml) |
| HTTPRoute | agentic-api-internal | epp.gateway.internal | / | no | [`guides/agentic-api/manifests/gateway/routes.yaml`](https://github.com/llm-d/llm-d/blob/c690e9acdab3f1c426e3e03ad89de99ebc56bc50/guides/agentic-api/manifests/gateway/routes.yaml) |
| HTTPRoute | coordinator |  | /v1/completions, /v1/chat/completions, /inference/v1/generate | no | [`guides/coord-disaggregation/coordinator/base/httproute.yaml`](https://github.com/llm-d/llm-d/blob/c690e9acdab3f1c426e3e03ad89de99ebc56bc50/guides/coord-disaggregation/coordinator/base/httproute.yaml) |
| HTTPRoute | deepseek-route |  |  | no | [`guides/multi-model-routing/manifests/httproutes.yaml`](https://github.com/llm-d/llm-d/blob/c690e9acdab3f1c426e3e03ad89de99ebc56bc50/guides/multi-model-routing/manifests/httproutes.yaml) |
| HTTPRoute | omni-serving-endpoints |  | /v1/audio/speech, /v1/audio/transcriptions, /v1/images/generations | no | [`guides/omni-serving/httproutes.yaml`](https://github.com/llm-d/llm-d/blob/c690e9acdab3f1c426e3e03ad89de99ebc56bc50/guides/omni-serving/httproutes.yaml) |
| HTTPRoute | qwen-route |  |  | no | [`guides/multi-model-routing/manifests/httproutes.yaml`](https://github.com/llm-d/llm-d/blob/c690e9acdab3f1c426e3e03ad89de99ebc56bc50/guides/multi-model-routing/manifests/httproutes.yaml) |

!!! warning "No Network Policies"
    No NetworkPolicy resources were found in the analyzed sources. Network policies may exist in overlays, Helm values, or cluster-level configurations not captured by static analysis.

