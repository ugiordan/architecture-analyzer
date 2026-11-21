# kserve: Network

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    kserve["kserve"]:::component
    kserve --> svc_0["kserve-controller-manager-metrics-service\nClusterIP: 8443/TCP"]:::svc
    kserve --> svc_1["kserve-controller-manager-service\nClusterIP: 8443/TCP"]:::svc
    kserve --> svc_2["kserve-webhook-server-service\nClusterIP: 443/TCP"]:::svc
    kserve --> svc_3["llmisvc-controller-manager-service\nClusterIP: 8443/TCP"]:::svc
    kserve --> svc_4["llmisvc-webhook-server-service\nClusterIP: 443/TCP"]:::svc
    kserve --> svc_5["localmodel-webhook-server-service\nClusterIP: 443/TCP"]:::svc
    kserve --> svc_6["model-serving-api\nClusterIP: 443/TCP,8080/TCP"]:::svc
    kserve --> svc_7["modelexpress-operator-metrics\nClusterIP: 8080/TCP"]:::svc
    kserve --> svc_8["modelexpress-operator-metrics\nClusterIP: 8443/TCP"]:::svc
    kserve --> svc_9["odh-model-controller-webhook-service\nClusterIP: 443/TCP"]:::svc
    kserve -.-> ext_httpx[["httpx\napi"]]:::ext
    kserve -.-> ext_requests[["requests\napi"]]:::ext
    kserve -.-> ext_urllib[["urllib\napi"]]:::ext
    kserve -.-> ext_weaviate[["weaviate\napi"]]:::ext
    kserve -.-> ext_azure_blob[["azure-blob\nobject-storage"]]:::ext
    kserve -.-> ext_gcs[["gcs\nobject-storage"]]:::ext
    kserve -.-> ext_s3[["s3\nobject-storage"]]:::ext
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| kserve-controller-manager-metrics-service | ClusterIP | 8443/TCP | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/kserve/blob/d0ff8ca1aa8a854d7e3be2a28d3023db2649626b/kustomize:config/overlays/odh) |
| kserve-controller-manager-service | ClusterIP | 8443/TCP | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/kserve/blob/d0ff8ca1aa8a854d7e3be2a28d3023db2649626b/kustomize:config/overlays/odh) |
| kserve-webhook-server-service | ClusterIP | 443/TCP | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/kserve/blob/d0ff8ca1aa8a854d7e3be2a28d3023db2649626b/kustomize:config/overlays/odh) |
| llmisvc-controller-manager-service | ClusterIP | 8443/TCP | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/kserve/blob/d0ff8ca1aa8a854d7e3be2a28d3023db2649626b/kustomize:config/overlays/odh) |
| llmisvc-webhook-server-service | ClusterIP | 443/TCP | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/kserve/blob/d0ff8ca1aa8a854d7e3be2a28d3023db2649626b/kustomize:config/overlays/odh) |
| localmodel-webhook-server-service | ClusterIP | 443/TCP | [`config/webhook/localmodel/service.yaml`](https://github.com/red-hat-data-services/kserve/blob/d0ff8ca1aa8a854d7e3be2a28d3023db2649626b/config/webhook/localmodel/service.yaml) |
| model-serving-api | ClusterIP | 443/TCP, 8080/TCP | [`kserve-module/prefetched-manifests-rhoai/modelcontroller/server/service.yaml`](https://github.com/red-hat-data-services/kserve/blob/d0ff8ca1aa8a854d7e3be2a28d3023db2649626b/kserve-module/prefetched-manifests-rhoai/modelcontroller/server/service.yaml) |
| modelexpress-operator-metrics | ClusterIP | 8080/TCP | [`kserve-module/prefetched-manifests-rhoai/modelexpress/manager/service.yaml`](https://github.com/red-hat-data-services/kserve/blob/d0ff8ca1aa8a854d7e3be2a28d3023db2649626b/kserve-module/prefetched-manifests-rhoai/modelexpress/manager/service.yaml) |
| modelexpress-operator-metrics | ClusterIP | 8443/TCP | [`kserve-module/prefetched-manifests-rhoai/modelexpress/components/openshift/service-patch.yaml`](https://github.com/red-hat-data-services/kserve/blob/d0ff8ca1aa8a854d7e3be2a28d3023db2649626b/kserve-module/prefetched-manifests-rhoai/modelexpress/components/openshift/service-patch.yaml) |
| odh-model-controller-webhook-service | ClusterIP | 443/TCP | [`kserve-module/prefetched-manifests-rhoai/modelcontroller/webhook/service.yaml`](https://github.com/red-hat-data-services/kserve/blob/d0ff8ca1aa8a854d7e3be2a28d3023db2649626b/kserve-module/prefetched-manifests-rhoai/modelcontroller/webhook/service.yaml) |

### Ingress / Routing

| Kind | Name | Hosts | Paths | TLS | Source |
|------|------|-------|-------|-----|--------|
| HTTPRoute | rbac-inferred |  |  | no | [`rbac/kserve-manager-role`](https://github.com/red-hat-data-services/kserve/blob/d0ff8ca1aa8a854d7e3be2a28d3023db2649626b/rbac/kserve-manager-role) |
| Ingress | rbac-inferred |  |  | no | [`rbac/kserve-manager-role`](https://github.com/red-hat-data-services/kserve/blob/d0ff8ca1aa8a854d7e3be2a28d3023db2649626b/rbac/kserve-manager-role) |
| VirtualService | rbac-inferred |  |  | no | [`rbac/kserve-manager-role`](https://github.com/red-hat-data-services/kserve/blob/d0ff8ca1aa8a854d7e3be2a28d3023db2649626b/rbac/kserve-manager-role) |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| kserve-controller-manager | Ingress, Egress | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/kserve/blob/d0ff8ca1aa8a854d7e3be2a28d3023db2649626b/kustomize:config/overlays/odh) |
| kserve-localmodel-controller-manager | Ingress, Egress | [`config/overlays/odh-modelcache/network-policies.yaml`](https://github.com/red-hat-data-services/kserve/blob/d0ff8ca1aa8a854d7e3be2a28d3023db2649626b/config/overlays/odh-modelcache/network-policies.yaml) |
| kserve-localmodel-download | Egress | [`config/overlays/odh-modelcache/network-policies.yaml`](https://github.com/red-hat-data-services/kserve/blob/d0ff8ca1aa8a854d7e3be2a28d3023db2649626b/config/overlays/odh-modelcache/network-policies.yaml) |
| kserve-localmodelnode-agent | Ingress, Egress | [`config/overlays/odh-modelcache/network-policies.yaml`](https://github.com/red-hat-data-services/kserve/blob/d0ff8ca1aa8a854d7e3be2a28d3023db2649626b/config/overlays/odh-modelcache/network-policies.yaml) |
| llmisvc-controller-manager | Ingress, Egress | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/kserve/blob/d0ff8ca1aa8a854d7e3be2a28d3023db2649626b/kustomize:config/overlays/odh) |
| modelexpress-operator |  | [`kserve-module/prefetched-manifests-rhoai/modelexpress/components/openshift/networkpolicy-patch.yaml`](https://github.com/red-hat-data-services/kserve/blob/d0ff8ca1aa8a854d7e3be2a28d3023db2649626b/kserve-module/prefetched-manifests-rhoai/modelexpress/components/openshift/networkpolicy-patch.yaml) |
| modelexpress-operator | Ingress, Egress | [`kserve-module/prefetched-manifests-rhoai/modelexpress/manager/networkpolicy.yaml`](https://github.com/red-hat-data-services/kserve/blob/d0ff8ca1aa8a854d7e3be2a28d3023db2649626b/kserve-module/prefetched-manifests-rhoai/modelexpress/manager/networkpolicy.yaml) |
| odh-model-controller-webhook-allow-ingress |  | [`kserve-module/prefetched-manifests-rhoai/modelcontroller/webhook/networkpolicy.yaml`](https://github.com/red-hat-data-services/kserve/blob/d0ff8ca1aa8a854d7e3be2a28d3023db2649626b/kserve-module/prefetched-manifests-rhoai/modelcontroller/webhook/networkpolicy.yaml) |

## Network Policy Graph

Visual representation of NetworkPolicy rules. Ingress rules show what traffic is allowed into pods, egress rules show what traffic is allowed out.

```mermaid
graph LR
    classDef policy fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef pod fill:#3498db,stroke:#2980b9,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff

    kserve["kserve\nPods"]:::pod
    np_0_kserve_controller_manager{{"kserve-controller-manager\nIngress, Egress"}}:::policy
    src_0_0_0["ns: policy-group.network.openshift.io/host-network="]:::external
    src_0_0_0 -->|ingress| np_0_kserve_controller_manager
    src_0_1_0["ns: network.openshift.io/policy-group=monitoring"]:::external
    src_0_1_0 -->|ingress| np_0_kserve_controller_manager
    np_0_kserve_controller_manager --> kserve
    dst_0_0_all["all destinations"]:::external
    kserve -->|egress| dst_0_0_all
    np_1_kserve_localmodel_controller_manager{{"kserve-localmodel-controller-manager\nIngress, Egress"}}:::policy
    src_1_0_0["ns: policy-group.network.openshift.io/host-network="]:::external
    src_1_0_0 -->|ingress| np_1_kserve_localmodel_controller_manager
    src_1_1_0["ns: network.openshift.io/policy-group=monitoring"]:::external
    src_1_1_0 -->|ingress| np_1_kserve_localmodel_controller_manager
    np_1_kserve_localmodel_controller_manager --> kserve
    dst_1_0_all["all destinations"]:::external
    kserve -->|egress| dst_1_0_all
    np_2_kserve_localmodel_download{{"kserve-localmodel-download\nEgress"}}:::policy
    np_2_kserve_localmodel_download --> kserve
    dst_2_0_all["all destinations"]:::external
    kserve -->|egress| dst_2_0_all
    np_3_kserve_localmodelnode_agent{{"kserve-localmodelnode-agent\nIngress, Egress"}}:::policy
    src_3_0_0["ns: network.openshift.io/policy-group=monitoring"]:::external
    src_3_0_0 -->|ingress| np_3_kserve_localmodelnode_agent
    np_3_kserve_localmodelnode_agent --> kserve
    dst_3_0_all["all destinations"]:::external
    kserve -->|egress| dst_3_0_all
    np_4_llmisvc_controller_manager{{"llmisvc-controller-manager\nIngress, Egress"}}:::policy
    src_4_0_0["ns: policy-group.network.openshift.io/host-network="]:::external
    src_4_0_0 -->|ingress| np_4_llmisvc_controller_manager
    src_4_1_0["ns: network.openshift.io/policy-group=monitoring"]:::external
    src_4_1_0 -->|ingress| np_4_llmisvc_controller_manager
    np_4_llmisvc_controller_manager --> kserve
    dst_4_0_all["all destinations"]:::external
    kserve -->|egress| dst_4_0_all
    np_5_modelexpress_operator{{"modelexpress-operator\nIngress"}}:::policy
    src_5_0_0["all namespaces"]:::external
    src_5_0_0 -->|ingress| np_5_modelexpress_operator
    np_5_modelexpress_operator --> kserve
    np_6_modelexpress_operator{{"modelexpress-operator\nIngress, Egress"}}:::policy
    src_6_0_0["all namespaces"]:::external
    src_6_0_0 -->|ingress| np_6_modelexpress_operator
    np_6_modelexpress_operator --> kserve
    dst_6_0_all["all destinations"]:::external
    kserve -->|egress| dst_6_0_all
    np_7_odh_model_controller_webhook_allow_ingress{{"odh-model-controller-webhook-allow-ingress\nIngress"}}:::policy
    src_7_0_all["all pods"]:::external
    src_7_0_all -->|ingress| np_7_odh_model_controller_webhook_allow_ingress
    np_7_odh_model_controller_webhook_allow_ingress --> kserve
```

