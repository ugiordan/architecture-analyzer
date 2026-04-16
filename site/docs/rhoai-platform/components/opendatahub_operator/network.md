# opendatahub-operator: Network

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff

    opendatahub_operator["opendatahub-operator"]:::component
    opendatahub_operator --> svc_0["webhook-service\nClusterIP: 443/TCP"]:::svc
    opendatahub_operator --> svc_1["webhook-service\nClusterIP: 443/TCP"]:::svc
    opendatahub_operator --> svc_2["webhook-service\nClusterIP: 443/TCP"]:::svc
    opendatahub_operator --> svc_3["odh-dashboard\nClusterIP: 8443/TCP"]:::svc
    opendatahub_operator --> svc_4["kserve-controller-manager-service\nClusterIP: 8443/TCP"]:::svc
    opendatahub_operator --> svc_5["kserve-webhook-server-service\nClusterIP: 443/TCP"]:::svc
    opendatahub_operator --> svc_6["odh-model-controller-webhook-service\nClusterIP: 443/TCP"]:::svc
    opendatahub_operator --> svc_7["webhook-service\nClusterIP: 443/TCP"]:::svc
    opendatahub_operator --> svc_8["kuberay-operator\nClusterIP: 8080/TCP"]:::svc
    opendatahub_operator --> svc_9["webhook-service\nClusterIP: 443/TCP"]:::svc
    opendatahub_operator --> svc_10["training-operator\nClusterIP: 8080/TCP, 443/TCP"]:::svc
    opendatahub_operator --> svc_11["service\nClusterIP: 443/TCP"]:::svc
    opendatahub_operator --> svc_12["service\nClusterIP: 8080/TCP"]:::svc
    opendatahub_operator --> svc_13["webhook-service\nClusterIP: 443/TCP"]:::svc
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| webhook-service | ClusterIP | 443/TCP | `config/rhaii/webhook/service.yaml` |
| webhook-service | ClusterIP | 443/TCP | `config/rhoai/webhook/service.yaml` |
| webhook-service | ClusterIP | 443/TCP | `config/webhook/service.yaml` |
| odh-dashboard | ClusterIP | 8443/TCP | `opt/manifests/dashboard/core-bases/base/service.yaml` |
| kserve-controller-manager-service | ClusterIP | 8443/TCP | `opt/manifests/kserve/manager/service.yaml` |
| kserve-webhook-server-service | ClusterIP | 443/TCP | `opt/manifests/kserve/webhook/service.yaml` |
| odh-model-controller-webhook-service | ClusterIP | 443/TCP | `opt/manifests/modelcontroller/webhook/service.yaml` |
| webhook-service | ClusterIP | 443/TCP | `opt/manifests/modelregistry/webhook/service.yaml` |
| kuberay-operator | ClusterIP | 8080/TCP | `opt/manifests/ray/manager/service.yaml` |
| webhook-service | ClusterIP | 443/TCP | `opt/manifests/ray/webhook/service.yaml` |
| training-operator | ClusterIP | 8080/TCP, 443/TCP | `opt/manifests/trainingoperator/base/service.yaml` |
| service | ClusterIP | 443/TCP | `opt/manifests/workbenches/kf-notebook-controller/manager/service.yaml` |
| service | ClusterIP | 8080/TCP | `opt/manifests/workbenches/odh-notebook-controller/manager/service.yaml` |
| webhook-service | ClusterIP | 443/TCP | `opt/manifests/workbenches/odh-notebook-controller/webhook/service.yaml` |

### Ingress / Routing

| Kind | Name | Hosts | Paths | TLS | Source |
|------|------|-------|-------|-----|--------|
| HTTPRoute | odh-dashboard |  | / | no | `opt/manifests/dashboard/core-bases/base/httproute.yaml` |
| Gateway | kserve-ingress-gateway |  |  | no | `opt/manifests/kserve/overlays/test/gateway/ingress_gateway.yaml` |
| Ingress | rayclient-ingress | localhost | / | no | `opt/manifests/ray/samples/ingress-rayclient-tls.yaml` |
| Route | odh-dashboard |  |  | yes | `opt/manifests/dashboard/core-bases/base/routes.yaml` |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| odh-dashboard-allow-ports | Ingress | `opt/manifests/dashboard/modular-architecture/networkpolicy.yaml` |
| kserve-controller-manager |  | `opt/manifests/kserve/overlays/odh/network-policies.yaml` |

