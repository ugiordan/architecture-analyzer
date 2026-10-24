# trustyai-service-operator: Network

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    trustyai_service_operator["trustyai-service-operator"]:::component
    trustyai_service_operator --> svc_0["trustyai-service-operator-controller-manager-metrics-service\nClusterIP: 8443/TCP"]:::svc
    trustyai_service_operator --> svc_1["trustyai-service-operator-metrics-service\nClusterIP: 8443/TCP"]:::svc
    trustyai_service_operator --> svc_2["trustyai-service-operator-webhook-service\nClusterIP: 443/TCP"]:::svc
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| trustyai-service-operator-controller-manager-metrics-service | ClusterIP | 8443/TCP | [`kustomize:config/overlays/odh`](https://github.com/opendatahub-io/trustyai-service-operator/blob/81f45a36ebad1f0f800d230321b589adf7bd404e/kustomize:config/overlays/odh) |
| trustyai-service-operator-metrics-service | ClusterIP | 8443/TCP | [`kustomize:config/overlays/odh`](https://github.com/opendatahub-io/trustyai-service-operator/blob/81f45a36ebad1f0f800d230321b589adf7bd404e/kustomize:config/overlays/odh) |
| trustyai-service-operator-webhook-service | ClusterIP | 443/TCP | [`kustomize:config/overlays/odh`](https://github.com/opendatahub-io/trustyai-service-operator/blob/81f45a36ebad1f0f800d230321b589adf7bd404e/kustomize:config/overlays/odh) |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| controller-manager-allow-traffic | Ingress, Egress | [`trustyai-operator-module/config/manager/networkpolicy.yaml`](https://github.com/opendatahub-io/trustyai-service-operator/blob/81f45a36ebad1f0f800d230321b589adf7bd404e/trustyai-operator-module/config/manager/networkpolicy.yaml) |
| trustyai-service-operator-controller-manager-allow-egress | Egress | [`kustomize:config/overlays/odh`](https://github.com/opendatahub-io/trustyai-service-operator/blob/81f45a36ebad1f0f800d230321b589adf7bd404e/kustomize:config/overlays/odh) |
| trustyai-service-operator-controller-manager-allow-ports | Ingress | [`kustomize:config/overlays/odh`](https://github.com/opendatahub-io/trustyai-service-operator/blob/81f45a36ebad1f0f800d230321b589adf7bd404e/kustomize:config/overlays/odh) |

## Network Policy Graph

Visual representation of NetworkPolicy rules. Ingress rules show what traffic is allowed into pods, egress rules show what traffic is allowed out.

```mermaid
graph LR
    classDef policy fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef pod fill:#3498db,stroke:#2980b9,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff

    trustyai_service_operator["trustyai-service-operator\nPods"]:::pod
    np_0_controller_manager_allow_traffic{{"controller-manager-allow-traffic\nIngress, Egress"}}:::policy
    np_0_controller_manager_allow_traffic --> trustyai_service_operator
    np_1_trustyai_service_operator_controller_manager_allow_egress{{"trustyai-service-operator-controller-manager-allow-egress\nEgress"}}:::policy
    np_1_trustyai_service_operator_controller_manager_allow_egress --> trustyai_service_operator
    np_2_trustyai_service_operator_controller_manager_allow_ports{{"trustyai-service-operator-controller-manager-allow-ports\nIngress"}}:::policy
    np_2_trustyai_service_operator_controller_manager_allow_ports --> trustyai_service_operator
```

