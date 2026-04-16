# odh-dashboard: Network

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff

    odh_dashboard["odh-dashboard"]:::component
    odh_dashboard --> svc_0["odh-dashboard\nClusterIP: 8443/TCP"]:::svc
    odh_dashboard --> svc_1["workspaces-backend\nClusterIP: 4000/TCP"]:::svc
    odh_dashboard --> svc_2["workspaces-webhook-service\nClusterIP: 443/TCP"]:::svc
    odh_dashboard --> svc_3["workspaces-controller-metrics-service\nClusterIP: 8080/TCP"]:::svc
    odh_dashboard --> svc_4["workspaces-frontend\nClusterIP: 8080/TCP"]:::svc
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| odh-dashboard | ClusterIP | 8443/TCP | `manifests/core-bases/base/service.yaml` |
| workspaces-backend | ClusterIP | 4000/TCP | `packages/notebooks/upstream/workspaces/backend/manifests/kustomize/base/service.yaml` |
| workspaces-webhook-service | ClusterIP | 443/TCP | `packages/notebooks/upstream/workspaces/controller/manifests/kustomize/base/webhook/service.yaml` |
| workspaces-controller-metrics-service | ClusterIP | 8080/TCP | `packages/notebooks/upstream/workspaces/controller/manifests/kustomize/components/prometheus/service.yaml` |
| workspaces-frontend | ClusterIP | 8080/TCP | `packages/notebooks/upstream/workspaces/frontend/manifests/kustomize/base/service.yaml` |

### Ingress / Routing

| Kind | Name | Hosts | Paths | TLS | Source |
|------|------|-------|-------|-----|--------|
| Gateway | kubeflow-gateway |  |  | no | `packages/notebooks/upstream/developing/manifests/istio-gateway/gateway.yaml` |
| HTTPRoute | odh-dashboard |  | / | no | `manifests/core-bases/base/httproute.yaml` |
| Route | odh-dashboard |  |  | yes | `manifests/core-bases/base/routes.yaml` |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| odh-dashboard-allow-ports | Ingress | `manifests/modular-architecture/networkpolicy.yaml` |
| dashboard-perses-access | Ingress | `manifests/observability/odh/network-policy.yaml` |
| dashboard-perses-access | Ingress | `manifests/observability/rhoai/network-policy.yaml` |
| workspaces-controller | Ingress | `packages/notebooks/upstream/workspaces/controller/manifests/kustomize/components/istio/network-policy.yaml` |
| allow-perses-operator-access | Ingress | `packages/observability/setup/network-policy-perses-operator-access.yaml` |

