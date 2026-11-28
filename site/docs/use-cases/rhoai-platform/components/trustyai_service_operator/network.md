# trustyai-service-operator: Network

[Overview](index.md) · [Agent Context](agent-context.md) · [Extracted Data](details.md) · [Network](network.md) · [RBAC](rbac.md) · [Security](security.md) · [Cache](cache.md) · [Dataflow](dataflow.md)

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
| trustyai-service-operator-controller-manager-metrics-service | ClusterIP | 8443/TCP | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/136ae01feaa1c5edd794af980240064097432e93/kustomize:config/overlays/odh) |
| trustyai-service-operator-metrics-service | ClusterIP | 8443/TCP | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/136ae01feaa1c5edd794af980240064097432e93/kustomize:config/overlays/odh) |
| trustyai-service-operator-webhook-service | ClusterIP | 443/TCP | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/136ae01feaa1c5edd794af980240064097432e93/kustomize:config/overlays/odh) |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| controller-manager-allow-traffic | Ingress, Egress | [`trustyai-operator-module/config/manager/networkpolicy.yaml`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/136ae01feaa1c5edd794af980240064097432e93/trustyai-operator-module/config/manager/networkpolicy.yaml) |
| trustyai-service-operator-controller-manager-allow-egress | Egress | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/136ae01feaa1c5edd794af980240064097432e93/kustomize:config/overlays/odh) |
| trustyai-service-operator-controller-manager-allow-ports | Ingress | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/136ae01feaa1c5edd794af980240064097432e93/kustomize:config/overlays/odh) |

## Network Policy Graph

Visual representation of NetworkPolicy rules. Ingress rules show what traffic is allowed into pods, egress rules show what traffic is allowed out.

```mermaid
graph LR
    classDef policy fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef pod fill:#3498db,stroke:#2980b9,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff

    trustyai_service_operator["trustyai-service-operator\nPods"]:::pod
    np_0_controller_manager_allow_traffic{{"controller-manager-allow-traffic\nIngress, Egress"}}:::policy
    src_0_0_all["all pods"]:::external
    src_0_0_all -->|ingress| np_0_controller_manager_allow_traffic
    np_0_controller_manager_allow_traffic --> trustyai_service_operator
    dst_0_0_all["all destinations"]:::external
    trustyai_service_operator -->|egress| dst_0_0_all
    np_1_trustyai_service_operator_controller_manager_allow_egress{{"trustyai-service-operator-controller-manager-allow-egress\nEgress"}}:::policy
    np_1_trustyai_service_operator_controller_manager_allow_egress --> trustyai_service_operator
    dst_1_0_all["all destinations"]:::external
    trustyai_service_operator -->|egress| dst_1_0_all
    np_2_trustyai_service_operator_controller_manager_allow_ports{{"trustyai-service-operator-controller-manager-allow-ports\nIngress"}}:::policy
    src_2_0_0["ns: kubernetes.io/metadata.name=openshift-monitoring"]:::external
    src_2_0_0 -->|ingress| np_2_trustyai_service_operator_controller_manager_allow_ports
    src_2_0_1["ns: kubernetes.io/metadata.name=openshift-user-workload-monitoring"]:::external
    src_2_0_1 -->|ingress| np_2_trustyai_service_operator_controller_manager_allow_ports
    src_2_1_all["all pods"]:::external
    src_2_1_all -->|ingress| np_2_trustyai_service_operator_controller_manager_allow_ports
    src_2_2_all["all pods"]:::external
    src_2_2_all -->|ingress| np_2_trustyai_service_operator_controller_manager_allow_ports
    np_2_trustyai_service_operator_controller_manager_allow_ports --> trustyai_service_operator
```

