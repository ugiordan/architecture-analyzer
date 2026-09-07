# models-as-a-service: Network

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    models_as_a_service["models-as-a-service"]:::component
    models_as_a_service --> svc_0["maas-api\nClusterIP: 8080/TCP,9090/TCP"]:::svc
    models_as_a_service --> svc_1["maas-api\nClusterIP: 0/TCP,8443/TCP"]:::svc
    models_as_a_service --> svc_2["maas-controller-webhook-service\nClusterIP: 443/TCP"]:::svc
    models_as_a_service --> svc_3["payload-processing\nClusterIP: 9004/TCP"]:::svc
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| maas-api | ClusterIP | 8080/TCP, 9090/TCP | [`deployment/base/maas-api/core/service.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/c5d5cab6fed7f11ea101e101993b8c9b80de60ef/deployment/base/maas-api/core/service.yaml) |
| maas-api | ClusterIP | 0/TCP, 8443/TCP | [`deployment/base/maas-api/overlays/tls/service-patch.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/c5d5cab6fed7f11ea101e101993b8c9b80de60ef/deployment/base/maas-api/overlays/tls/service-patch.yaml) |
| maas-controller-webhook-service | ClusterIP | 443/TCP | [`deployment/base/maas-controller/webhook/service.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/c5d5cab6fed7f11ea101e101993b8c9b80de60ef/deployment/base/maas-controller/webhook/service.yaml) |
| payload-processing | ClusterIP | 9004/TCP | [`deployment/base/payload-processing/instance/service.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/c5d5cab6fed7f11ea101e101993b8c9b80de60ef/deployment/base/payload-processing/instance/service.yaml) |

### Ingress / Routing

| Kind | Name | Hosts | Paths | TLS | Source |
|------|------|-------|-------|-----|--------|
| DestinationRule | maas-api-backend-tls |  |  | no | [`deployment/base/maas-api/overlays/tls/destinationrule.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/c5d5cab6fed7f11ea101e101993b8c9b80de60ef/deployment/base/maas-api/overlays/tls/destinationrule.yaml) |
| HTTPRoute | maas-api-route |  | /v1/models, /v1/subscriptions, /v1/api-keys, /maas-api | no | [`deployment/base/maas-api/networking/httproute.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/c5d5cab6fed7f11ea101e101993b8c9b80de60ef/deployment/base/maas-api/networking/httproute.yaml) |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| maas-api-allow-cleanup | Ingress | [`deployment/base/maas-api/core/networkpolicy-cleanup-ingress.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/c5d5cab6fed7f11ea101e101993b8c9b80de60ef/deployment/base/maas-api/core/networkpolicy-cleanup-ingress.yaml) |
| maas-api-allow-deployment-ns | Ingress | [`deployment/base/maas-api/networking/maas-deployment-ns-networkpolicy.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/c5d5cab6fed7f11ea101e101993b8c9b80de60ef/deployment/base/maas-api/networking/maas-deployment-ns-networkpolicy.yaml) |
| maas-api-allow-gateway | Ingress | [`deployment/base/maas-api/networking/maas-gateway-networkpolicy.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/c5d5cab6fed7f11ea101e101993b8c9b80de60ef/deployment/base/maas-api/networking/maas-gateway-networkpolicy.yaml) |
| maas-api-allow-monitoring | Ingress | [`deployment/base/maas-api/monitoring/networkpolicy.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/c5d5cab6fed7f11ea101e101993b8c9b80de60ef/deployment/base/maas-api/monitoring/networkpolicy.yaml) |
| maas-api-cleanup-restrict | Egress, Ingress | [`deployment/base/maas-api/core/networkpolicy-cleanup.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/c5d5cab6fed7f11ea101e101993b8c9b80de60ef/deployment/base/maas-api/core/networkpolicy-cleanup.yaml) |
| maas-api-cleanup-restrict |  | [`deployment/base/maas-api/overlays/tls/networkpolicy-cleanup-patch.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/c5d5cab6fed7f11ea101e101993b8c9b80de60ef/deployment/base/maas-api/overlays/tls/networkpolicy-cleanup-patch.yaml) |
| maas-authorino-allow | Ingress | [`deployment/base/maas-api/networking/maas-authorino-networkpolicy.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/c5d5cab6fed7f11ea101e101993b8c9b80de60ef/deployment/base/maas-api/networking/maas-authorino-networkpolicy.yaml) |
| maas-authorino-allow | Ingress | [`scripts/data/maas-authorino-networkpolicy.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/c5d5cab6fed7f11ea101e101993b8c9b80de60ef/scripts/data/maas-authorino-networkpolicy.yaml) |
| maas-controller-allow-monitoring | Ingress | [`deployment/base/maas-controller/monitoring/networkpolicy.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/c5d5cab6fed7f11ea101e101993b8c9b80de60ef/deployment/base/maas-controller/monitoring/networkpolicy.yaml) |
| payload-processing | Ingress, Egress | [`deployment/base/payload-processing/networking/networkpolicy.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/c5d5cab6fed7f11ea101e101993b8c9b80de60ef/deployment/base/payload-processing/networking/networkpolicy.yaml) |

## Network Policy Graph

Visual representation of NetworkPolicy rules. Ingress rules show what traffic is allowed into pods, egress rules show what traffic is allowed out.

```mermaid
graph LR
    classDef policy fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef pod fill:#3498db,stroke:#2980b9,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff

    models_as_a_service["models-as-a-service\nPods"]:::pod
    np_0_maas_api_allow_cleanup{{"maas-api-allow-cleanup\nIngress"}}:::policy
    np_0_maas_api_allow_cleanup --> models_as_a_service
    np_1_maas_api_allow_deployment_ns{{"maas-api-allow-deployment-ns\nIngress"}}:::policy
    np_1_maas_api_allow_deployment_ns --> models_as_a_service
    np_2_maas_api_allow_gateway{{"maas-api-allow-gateway\nIngress"}}:::policy
    np_2_maas_api_allow_gateway --> models_as_a_service
    np_3_maas_api_allow_monitoring{{"maas-api-allow-monitoring\nIngress"}}:::policy
    np_3_maas_api_allow_monitoring --> models_as_a_service
    np_4_maas_api_cleanup_restrict{{"maas-api-cleanup-restrict\nEgress, Ingress"}}:::policy
    np_4_maas_api_cleanup_restrict --> models_as_a_service
    np_5_maas_api_cleanup_restrict{{"maas-api-cleanup-restrict\nIngress"}}:::policy
    np_5_maas_api_cleanup_restrict --> models_as_a_service
    np_6_maas_authorino_allow{{"maas-authorino-allow\nIngress"}}:::policy
    np_6_maas_authorino_allow --> models_as_a_service
    np_7_maas_authorino_allow{{"maas-authorino-allow\nIngress"}}:::policy
    np_7_maas_authorino_allow --> models_as_a_service
    np_8_maas_controller_allow_monitoring{{"maas-controller-allow-monitoring\nIngress"}}:::policy
    np_8_maas_controller_allow_monitoring --> models_as_a_service
    np_9_payload_processing{{"payload-processing\nIngress, Egress"}}:::policy
    np_9_payload_processing --> models_as_a_service
```

