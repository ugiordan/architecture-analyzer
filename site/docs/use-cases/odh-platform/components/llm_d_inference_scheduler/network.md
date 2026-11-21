# llm-d-inference-scheduler: Network

### Services

No services found in analyzed sources.

### Ingress / Routing

| Kind | Name | Hosts | Paths | TLS | Source |
|------|------|-------|-------|-----|--------|
| Gateway | inference-gateway |  |  | no | [`config/manifests/gateway/agentgateway/gateway.yaml`](https://github.com/opendatahub-io/llm-d-inference-scheduler/blob/30fb1d0cff419b5cba63e71c1b521983afc921d1/config/manifests/gateway/agentgateway/gateway.yaml) |
| Gateway | inference-gateway |  |  | no | [`config/manifests/gateway/envoyaigateway/gateway.yaml`](https://github.com/opendatahub-io/llm-d-inference-scheduler/blob/30fb1d0cff419b5cba63e71c1b521983afc921d1/config/manifests/gateway/envoyaigateway/gateway.yaml) |
| Gateway | inference-gateway |  |  | no | [`config/manifests/gateway/gke/gateway.yaml`](https://github.com/opendatahub-io/llm-d-inference-scheduler/blob/30fb1d0cff419b5cba63e71c1b521983afc921d1/config/manifests/gateway/gke/gateway.yaml) |
| Gateway | inference-gateway |  |  | no | [`config/manifests/gateway/istio/gateway.yaml`](https://github.com/opendatahub-io/llm-d-inference-scheduler/blob/30fb1d0cff419b5cba63e71c1b521983afc921d1/config/manifests/gateway/istio/gateway.yaml) |
| Gateway | inference-gateway |  |  | no | [`config/manifests/gateway/nginxgatewayfabric/gateway.yaml`](https://github.com/opendatahub-io/llm-d-inference-scheduler/blob/30fb1d0cff419b5cba63e71c1b521983afc921d1/config/manifests/gateway/nginxgatewayfabric/gateway.yaml) |

!!! warning "No Network Policies"
    No NetworkPolicy resources were found in the analyzed sources. Network policies may exist in overlays, Helm values, or cluster-level configurations not captured by static analysis.

