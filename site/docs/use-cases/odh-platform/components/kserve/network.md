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
| kserve-controller-manager-metrics-service | ClusterIP | 8443/TCP | [`kustomize:config/overlays/odh`](https://github.com/opendatahub-io/kserve/blob/e0c6c8559bd78fa9d7bb8692bb737c8e7eb4c931/kustomize:config/overlays/odh) |
| kserve-controller-manager-service | ClusterIP | 8443/TCP | [`kustomize:config/overlays/odh`](https://github.com/opendatahub-io/kserve/blob/e0c6c8559bd78fa9d7bb8692bb737c8e7eb4c931/kustomize:config/overlays/odh) |
| kserve-webhook-server-service | ClusterIP | 443/TCP | [`kustomize:config/overlays/odh`](https://github.com/opendatahub-io/kserve/blob/e0c6c8559bd78fa9d7bb8692bb737c8e7eb4c931/kustomize:config/overlays/odh) |
| llmisvc-controller-manager-service | ClusterIP | 8443/TCP | [`kustomize:config/overlays/odh`](https://github.com/opendatahub-io/kserve/blob/e0c6c8559bd78fa9d7bb8692bb737c8e7eb4c931/kustomize:config/overlays/odh) |
| llmisvc-webhook-server-service | ClusterIP | 443/TCP | [`kustomize:config/overlays/odh`](https://github.com/opendatahub-io/kserve/blob/e0c6c8559bd78fa9d7bb8692bb737c8e7eb4c931/kustomize:config/overlays/odh) |
| localmodel-webhook-server-service | ClusterIP | 443/TCP | [`config/webhook/localmodel/service.yaml`](https://github.com/opendatahub-io/kserve/blob/e0c6c8559bd78fa9d7bb8692bb737c8e7eb4c931/config/webhook/localmodel/service.yaml) |

### Ingress / Routing

| Kind | Name | Hosts | Paths | TLS | Source |
|------|------|-------|-------|-----|--------|
| HTTPRoute | rbac-inferred |  |  | no | [`rbac/kserve-manager-role`](https://github.com/opendatahub-io/kserve/blob/e0c6c8559bd78fa9d7bb8692bb737c8e7eb4c931/rbac/kserve-manager-role) |
| Ingress | rbac-inferred |  |  | no | [`rbac/kserve-manager-role`](https://github.com/opendatahub-io/kserve/blob/e0c6c8559bd78fa9d7bb8692bb737c8e7eb4c931/rbac/kserve-manager-role) |
| VirtualService | rbac-inferred |  |  | no | [`rbac/kserve-manager-role`](https://github.com/opendatahub-io/kserve/blob/e0c6c8559bd78fa9d7bb8692bb737c8e7eb4c931/rbac/kserve-manager-role) |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| kserve-controller-manager | Ingress, Egress | [`kustomize:config/overlays/odh`](https://github.com/opendatahub-io/kserve/blob/e0c6c8559bd78fa9d7bb8692bb737c8e7eb4c931/kustomize:config/overlays/odh) |
| kserve-localmodel-controller-manager | Ingress, Egress | [`config/overlays/odh-modelcache/network-policies.yaml`](https://github.com/opendatahub-io/kserve/blob/e0c6c8559bd78fa9d7bb8692bb737c8e7eb4c931/config/overlays/odh-modelcache/network-policies.yaml) |
| kserve-localmodel-download | Egress | [`config/overlays/odh-modelcache/network-policies.yaml`](https://github.com/opendatahub-io/kserve/blob/e0c6c8559bd78fa9d7bb8692bb737c8e7eb4c931/config/overlays/odh-modelcache/network-policies.yaml) |
| kserve-localmodelnode-agent | Ingress, Egress | [`config/overlays/odh-modelcache/network-policies.yaml`](https://github.com/opendatahub-io/kserve/blob/e0c6c8559bd78fa9d7bb8692bb737c8e7eb4c931/config/overlays/odh-modelcache/network-policies.yaml) |
| llmisvc-controller-manager | Ingress, Egress | [`kustomize:config/overlays/odh`](https://github.com/opendatahub-io/kserve/blob/e0c6c8559bd78fa9d7bb8692bb737c8e7eb4c931/kustomize:config/overlays/odh) |

## Network Policy Graph

Visual representation of NetworkPolicy rules. Ingress rules show what traffic is allowed into pods, egress rules show what traffic is allowed out.

```mermaid
graph LR
    classDef policy fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef pod fill:#3498db,stroke:#2980b9,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff

    kserve["kserve\nPods"]:::pod
    np_0_kserve_controller_manager{{"kserve-controller-manager\nIngress, Egress"}}:::policy
    np_0_kserve_controller_manager --> kserve
    np_1_kserve_localmodel_controller_manager{{"kserve-localmodel-controller-manager\nIngress, Egress"}}:::policy
    np_1_kserve_localmodel_controller_manager --> kserve
    np_2_kserve_localmodel_download{{"kserve-localmodel-download\nEgress"}}:::policy
    np_2_kserve_localmodel_download --> kserve
    np_3_kserve_localmodelnode_agent{{"kserve-localmodelnode-agent\nIngress, Egress"}}:::policy
    np_3_kserve_localmodelnode_agent --> kserve
    np_4_llmisvc_controller_manager{{"llmisvc-controller-manager\nIngress, Egress"}}:::policy
    np_4_llmisvc_controller_manager --> kserve
```

