# models-as-a-service: Network

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    models_as_a_service["models-as-a-service"]:::component
    models_as_a_service --> svc_0["maas-api\nClusterIP: 8080/TCP"]:::svc
    models_as_a_service --> svc_1["maas-api\nClusterIP: 0/TCP,8443/TCP"]:::svc
    models_as_a_service --> svc_2["maas-controller-webhook-service\nClusterIP: 443/TCP"]:::svc
    models_as_a_service --> svc_3["maas-discovery\nClusterIP: 8443/TCP"]:::svc
    models_as_a_service --> svc_4["payload-processing\nClusterIP: 9004/TCP"]:::svc
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| maas-api | ClusterIP | 8080/TCP | [`deployment/base/maas-api/core/service.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/47df205d72a68342f51cc53e91ff67dbe3ea36fa/deployment/base/maas-api/core/service.yaml) |
| maas-api | ClusterIP | 0/TCP, 8443/TCP | [`deployment/base/maas-api/overlays/tls/service-patch.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/47df205d72a68342f51cc53e91ff67dbe3ea36fa/deployment/base/maas-api/overlays/tls/service-patch.yaml) |
| maas-controller-webhook-service | ClusterIP | 443/TCP | [`deployment/base/maas-controller/webhook/service.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/47df205d72a68342f51cc53e91ff67dbe3ea36fa/deployment/base/maas-controller/webhook/service.yaml) |
| maas-discovery | ClusterIP | 8443/TCP | [`deployment/base/maas-discovery/core/service.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/47df205d72a68342f51cc53e91ff67dbe3ea36fa/deployment/base/maas-discovery/core/service.yaml) |
| payload-processing | ClusterIP | 9004/TCP | [`deployment/base/payload-processing/instance/service.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/47df205d72a68342f51cc53e91ff67dbe3ea36fa/deployment/base/payload-processing/instance/service.yaml) |

### Ingress / Routing

| Kind | Name | Hosts | Paths | TLS | Source |
|------|------|-------|-------|-----|--------|
| DestinationRule | maas-api-backend-tls |  |  | no | [`deployment/base/maas-api/overlays/tls/destinationrule.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/47df205d72a68342f51cc53e91ff67dbe3ea36fa/deployment/base/maas-api/overlays/tls/destinationrule.yaml) |
| HTTPRoute | maas-api-route |  | /v1/models, /v1/subscriptions, /v1/api-keys, /maas-api/v1, /maas-api/health, /v1/me | no | [`deployment/base/maas-api/networking/httproute.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/47df205d72a68342f51cc53e91ff67dbe3ea36fa/deployment/base/maas-api/networking/httproute.yaml) |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| maas-api-allow-cleanup | Ingress | [`deployment/base/maas-api/core/networkpolicy-cleanup-ingress.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/47df205d72a68342f51cc53e91ff67dbe3ea36fa/deployment/base/maas-api/core/networkpolicy-cleanup-ingress.yaml) |
| maas-api-allow-deployment-ns | Ingress | [`deployment/base/maas-api/overlays/tls/maas-deployment-ns-networkpolicy.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/47df205d72a68342f51cc53e91ff67dbe3ea36fa/deployment/base/maas-api/overlays/tls/maas-deployment-ns-networkpolicy.yaml) |
| maas-api-allow-gateway | Ingress | [`deployment/base/maas-api/overlays/tls/maas-gateway-networkpolicy.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/47df205d72a68342f51cc53e91ff67dbe3ea36fa/deployment/base/maas-api/overlays/tls/maas-gateway-networkpolicy.yaml) |
| maas-api-allow-monitoring | Ingress | [`deployment/base/maas-api/monitoring/networkpolicy.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/47df205d72a68342f51cc53e91ff67dbe3ea36fa/deployment/base/maas-api/monitoring/networkpolicy.yaml) |
| maas-api-cleanup-restrict | Egress, Ingress | [`deployment/base/maas-api/core/networkpolicy-cleanup.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/47df205d72a68342f51cc53e91ff67dbe3ea36fa/deployment/base/maas-api/core/networkpolicy-cleanup.yaml) |
| maas-api-cleanup-restrict |  | [`deployment/base/maas-api/overlays/tls/networkpolicy-cleanup-patch.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/47df205d72a68342f51cc53e91ff67dbe3ea36fa/deployment/base/maas-api/overlays/tls/networkpolicy-cleanup-patch.yaml) |
| maas-api-egress-restrict | Egress | [`deployment/base/maas-api/networking/maas-api-egress-networkpolicy.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/47df205d72a68342f51cc53e91ff67dbe3ea36fa/deployment/base/maas-api/networking/maas-api-egress-networkpolicy.yaml) |
| maas-authorino-allow | Ingress | [`deployment/base/maas-api/networking/maas-authorino-networkpolicy.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/47df205d72a68342f51cc53e91ff67dbe3ea36fa/deployment/base/maas-api/networking/maas-authorino-networkpolicy.yaml) |
| maas-authorino-allow | Ingress | [`scripts/data/maas-authorino-networkpolicy.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/47df205d72a68342f51cc53e91ff67dbe3ea36fa/scripts/data/maas-authorino-networkpolicy.yaml) |
| maas-controller-allow-monitoring | Ingress | [`deployment/base/maas-controller/monitoring/networkpolicy.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/47df205d72a68342f51cc53e91ff67dbe3ea36fa/deployment/base/maas-controller/monitoring/networkpolicy.yaml) |
| payload-processing | Ingress, Egress | [`deployment/base/payload-processing/networking/networkpolicy.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/47df205d72a68342f51cc53e91ff67dbe3ea36fa/deployment/base/payload-processing/networking/networkpolicy.yaml) |
| usage-logs-collector-egress-restrict | Egress | [`deployment/components/observability/usage-logs/usage-logs-collector-egress-networkpolicy.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/47df205d72a68342f51cc53e91ff67dbe3ea36fa/deployment/components/observability/usage-logs/usage-logs-collector-egress-networkpolicy.yaml) |
| usage-tenancy-proxy-allow-perses | Ingress | [`deployment/components/observability/usage-logs/dsci-managed/usage-tenancy-proxy-networkpolicy.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/47df205d72a68342f51cc53e91ff67dbe3ea36fa/deployment/components/observability/usage-logs/dsci-managed/usage-tenancy-proxy-networkpolicy.yaml) |
| usage-tenancy-proxy-egress-restrict | Egress | [`deployment/components/observability/usage-logs/dsci-managed/usage-tenancy-proxy-networkpolicy.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/47df205d72a68342f51cc53e91ff67dbe3ea36fa/deployment/components/observability/usage-logs/dsci-managed/usage-tenancy-proxy-networkpolicy.yaml) |

## Network Policy Graph

Visual representation of NetworkPolicy rules. Ingress rules show what traffic is allowed into pods, egress rules show what traffic is allowed out.

```mermaid
graph LR
    classDef policy fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef pod fill:#3498db,stroke:#2980b9,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff

    models_as_a_service["models-as-a-service\nPods"]:::pod
    np_0_maas_api_allow_cleanup{{"maas-api-allow-cleanup\nIngress"}}:::policy
    src_0_0_0["pods: app.kubernetes.io/component=api, app.kubernetes.io/name=maas-api, app.kubernetes.io/part-of=models-as-a-service, app=maas-api-cleanup"]:::external
    src_0_0_0 -->|ingress| np_0_maas_api_allow_cleanup
    np_0_maas_api_allow_cleanup --> models_as_a_service
    np_1_maas_api_allow_deployment_ns{{"maas-api-allow-deployment-ns\nIngress"}}:::policy
    src_1_0_0["ns: kubernetes.io/metadata.name=REPLACE_CONTROLLER_NAMESPACE"]:::external
    src_1_0_0 -->|ingress| np_1_maas_api_allow_deployment_ns
    src_1_0_1["ns: kubernetes.io/metadata.name=REPLACE_CONTROLLER_NAMESPACE"]:::external
    src_1_0_1 -->|ingress| np_1_maas_api_allow_deployment_ns
    src_1_0_2["ns: kubernetes.io/metadata.name=REPLACE_CONTROLLER_NAMESPACE"]:::external
    src_1_0_2 -->|ingress| np_1_maas_api_allow_deployment_ns
    np_1_maas_api_allow_deployment_ns --> models_as_a_service
    np_2_maas_api_allow_gateway{{"maas-api-allow-gateway\nIngress"}}:::policy
    src_2_0_0["ns: kubernetes.io/metadata.name=openshift-ingress"]:::external
    src_2_0_0 -->|ingress| np_2_maas_api_allow_gateway
    np_2_maas_api_allow_gateway --> models_as_a_service
    np_3_maas_api_allow_monitoring{{"maas-api-allow-monitoring\nIngress"}}:::policy
    src_3_0_0["ns: kubernetes.io/metadata.name=redhat-ods-monitoring"]:::external
    src_3_0_0 -->|ingress| np_3_maas_api_allow_monitoring
    np_3_maas_api_allow_monitoring --> models_as_a_service
    np_4_maas_api_cleanup_restrict{{"maas-api-cleanup-restrict\nEgress, Ingress"}}:::policy
    np_4_maas_api_cleanup_restrict --> models_as_a_service
    dst_4_0_0["pods: app.kubernetes.io/name=maas-api"]:::external
    models_as_a_service -->|egress| dst_4_0_0
    dst_4_1_all["all destinations"]:::external
    models_as_a_service -->|egress| dst_4_1_all
    np_5_maas_api_cleanup_restrict{{"maas-api-cleanup-restrict\nIngress"}}:::policy
    np_5_maas_api_cleanup_restrict --> models_as_a_service
    dst_5_0_0["pods: app.kubernetes.io/component=api, app.kubernetes.io/name=maas-api, app.kubernetes.io/part-of=models-as-a-service"]:::external
    models_as_a_service -->|egress| dst_5_0_0
    dst_5_1_all["all destinations"]:::external
    models_as_a_service -->|egress| dst_5_1_all
    np_6_maas_api_egress_restrict{{"maas-api-egress-restrict\nEgress"}}:::policy
    np_6_maas_api_egress_restrict --> models_as_a_service
    dst_6_0_0["ns: kubernetes.io/metadata.name=openshift-dns"]:::external
    models_as_a_service -->|egress| dst_6_0_0
    dst_6_1_all["all destinations"]:::external
    models_as_a_service -->|egress| dst_6_1_all
    np_7_maas_authorino_allow{{"maas-authorino-allow\nIngress"}}:::policy
    src_7_0_0["all namespaces"]:::external
    src_7_0_0 -->|ingress| np_7_maas_authorino_allow
    np_7_maas_authorino_allow --> models_as_a_service
    np_8_maas_authorino_allow{{"maas-authorino-allow\nIngress"}}:::policy
    src_8_0_0["all namespaces"]:::external
    src_8_0_0 -->|ingress| np_8_maas_authorino_allow
    np_8_maas_authorino_allow --> models_as_a_service
    np_9_maas_controller_allow_monitoring{{"maas-controller-allow-monitoring\nIngress"}}:::policy
    src_9_0_0["ns: kubernetes.io/metadata.name=redhat-ods-monitoring"]:::external
    src_9_0_0 -->|ingress| np_9_maas_controller_allow_monitoring
    np_9_maas_controller_allow_monitoring --> models_as_a_service
    np_10_payload_processing{{"payload-processing\nIngress, Egress"}}:::policy
    src_10_0_0["ns: kubernetes.io/metadata.name=openshift-ingress"]:::external
    src_10_0_0 -->|ingress| np_10_payload_processing
    src_10_1_0["ns: kubernetes.io/metadata.name=openshift-monitoring"]:::external
    src_10_1_0 -->|ingress| np_10_payload_processing
    src_10_1_1["ns: kubernetes.io/metadata.name=openshift-user-workload-monitoring"]:::external
    src_10_1_1 -->|ingress| np_10_payload_processing
    np_10_payload_processing --> models_as_a_service
    dst_10_0_all["all destinations"]:::external
    models_as_a_service -->|egress| dst_10_0_all
    dst_10_1_all["all destinations"]:::external
    models_as_a_service -->|egress| dst_10_1_all
    dst_10_2_0["ns: kubernetes.io/metadata.name=opendatahub"]:::external
    models_as_a_service -->|egress| dst_10_2_0
    np_11_usage_logs_collector_egress_restrict{{"usage-logs-collector-egress-restrict\nEgress"}}:::policy
    np_11_usage_logs_collector_egress_restrict --> models_as_a_service
    dst_11_0_0["ns: kubernetes.io/metadata.name=openshift-dns"]:::external
    models_as_a_service -->|egress| dst_11_0_0
    dst_11_1_0["pods: app.kubernetes.io/component=lokistack-gateway, app.kubernetes.io/instance=usage, app.kubernetes.io/name=lokistack"]:::external
    models_as_a_service -->|egress| dst_11_1_0
    np_12_usage_tenancy_proxy_allow_perses{{"usage-tenancy-proxy-allow-perses\nIngress"}}:::policy
    src_12_0_0["pods: app.kubernetes.io/managed-by=perses-operator"]:::external
    src_12_0_0 -->|ingress| np_12_usage_tenancy_proxy_allow_perses
    src_12_1_0["ns: network.openshift.io/policy-group=host-network"]:::external
    src_12_1_0 -->|ingress| np_12_usage_tenancy_proxy_allow_perses
    np_12_usage_tenancy_proxy_allow_perses --> models_as_a_service
    np_13_usage_tenancy_proxy_egress_restrict{{"usage-tenancy-proxy-egress-restrict\nEgress"}}:::policy
    np_13_usage_tenancy_proxy_egress_restrict --> models_as_a_service
    dst_13_0_0["ns: kubernetes.io/metadata.name=openshift-dns"]:::external
    models_as_a_service -->|egress| dst_13_0_0
    dst_13_1_all["all destinations"]:::external
    models_as_a_service -->|egress| dst_13_1_all
    dst_13_2_0["pods: app.kubernetes.io/component=lokistack-gateway, app.kubernetes.io/instance=data-science-usage-logs-lokistack, app.kubernetes.io/name=lokistack"]:::external
    models_as_a_service -->|egress| dst_13_2_0
```

