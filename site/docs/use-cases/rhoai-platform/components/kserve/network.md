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
| kserve-controller-manager-metrics-service | ClusterIP | 8443/TCP | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kustomize:config/overlays/odh) |
| kserve-controller-manager-service | ClusterIP | 8443/TCP | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kustomize:config/overlays/odh) |
| kserve-webhook-server-service | ClusterIP | 443/TCP | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kustomize:config/overlays/odh) |
| llmisvc-controller-manager-service | ClusterIP | 8443/TCP | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kustomize:config/overlays/odh) |
| llmisvc-webhook-server-service | ClusterIP | 443/TCP | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kustomize:config/overlays/odh) |
| localmodel-webhook-server-service | ClusterIP | 443/TCP | [`config/webhook/localmodel/service.yaml`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/config/webhook/localmodel/service.yaml) |
| model-serving-api | ClusterIP | 443/TCP, 8080/TCP | [`kserve-module/prefetched-manifests-rhoai/modelcontroller/server/service.yaml`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/prefetched-manifests-rhoai/modelcontroller/server/service.yaml) |
| modelexpress-operator-metrics | ClusterIP | 8080/TCP | [`kserve-module/prefetched-manifests-rhoai/modelexpress/manager/service.yaml`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/prefetched-manifests-rhoai/modelexpress/manager/service.yaml) |
| modelexpress-operator-metrics | ClusterIP | 8443/TCP | [`kserve-module/prefetched-manifests-rhoai/modelexpress/components/openshift/service-patch.yaml`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/prefetched-manifests-rhoai/modelexpress/components/openshift/service-patch.yaml) |
| odh-model-controller-webhook-service | ClusterIP | 443/TCP | [`kserve-module/prefetched-manifests-rhoai/modelcontroller/webhook/service.yaml`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/prefetched-manifests-rhoai/modelcontroller/webhook/service.yaml) |

### Ingress / Routing

| Kind | Name | Hosts | Paths | TLS | Source |
|------|------|-------|-------|-----|--------|
| HTTPRoute | rbac-inferred |  |  | no | [`rbac/kserve-manager-role`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/rbac/kserve-manager-role) |
| Ingress | rbac-inferred |  |  | no | [`rbac/kserve-manager-role`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/rbac/kserve-manager-role) |
| VirtualService | rbac-inferred |  |  | no | [`rbac/kserve-manager-role`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/rbac/kserve-manager-role) |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| kserve-controller-manager |  | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kustomize:config/overlays/odh) |
| llmisvc-controller-manager |  | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kustomize:config/overlays/odh) |
| modelexpress-operator |  | [`kserve-module/prefetched-manifests-rhoai/modelexpress/components/openshift/networkpolicy-patch.yaml`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/prefetched-manifests-rhoai/modelexpress/components/openshift/networkpolicy-patch.yaml) |
| modelexpress-operator | Ingress, Egress | [`kserve-module/prefetched-manifests-rhoai/modelexpress/manager/networkpolicy.yaml`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/prefetched-manifests-rhoai/modelexpress/manager/networkpolicy.yaml) |
| odh-model-controller-webhook-allow-ingress |  | [`kserve-module/prefetched-manifests-rhoai/modelcontroller/webhook/networkpolicy.yaml`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/prefetched-manifests-rhoai/modelcontroller/webhook/networkpolicy.yaml) |

## Network Policy Graph

Visual representation of NetworkPolicy rules. Ingress rules show what traffic is allowed into pods, egress rules show what traffic is allowed out.

```mermaid
graph LR
    classDef policy fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef pod fill:#3498db,stroke:#2980b9,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff

    kserve["kserve\nPods"]:::pod
    np_0_kserve_controller_manager{{"kserve-controller-manager\nIngress"}}:::policy
    np_0_kserve_controller_manager --> kserve
    np_1_llmisvc_controller_manager{{"llmisvc-controller-manager\nIngress"}}:::policy
    np_1_llmisvc_controller_manager --> kserve
    np_2_modelexpress_operator{{"modelexpress-operator\nIngress"}}:::policy
    np_2_modelexpress_operator --> kserve
    np_3_modelexpress_operator{{"modelexpress-operator\nIngress, Egress"}}:::policy
    np_3_modelexpress_operator --> kserve
    np_4_odh_model_controller_webhook_allow_ingress{{"odh-model-controller-webhook-allow-ingress\nIngress"}}:::policy
    np_4_odh_model_controller_webhook_allow_ingress --> kserve
```

