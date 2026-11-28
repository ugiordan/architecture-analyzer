# modelmesh-serving: Network

[Overview](index.md) · [Agent Context](agent-context.md) · [Extracted Data](details.md) · [Network](network.md) · [RBAC](rbac.md) · [Security](security.md) · [Cache](cache.md) · [Dataflow](dataflow.md)

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    modelmesh_serving["modelmesh-serving"]:::component
    modelmesh_serving --> svc_0["etcd\nClusterIP: 2379/TCP"]:::svc
    modelmesh_serving --> svc_1["modelmesh-controller\nClusterIP: 8080/TCP"]:::svc
    modelmesh_serving --> svc_2["modelmesh-webhook-server-service\nClusterIP: 9443/TCP"]:::svc
    modelmesh_serving -.-> ext_azure_blob[["azure-blob\nobject-storage"]]:::ext
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| etcd | ClusterIP | 2379/TCP | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/modelmesh-serving/blob/2751b227f532385b73e407e5a6f09ee61920406a/kustomize:config/overlays/odh) |
| modelmesh-controller | ClusterIP | 8080/TCP | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/modelmesh-serving/blob/2751b227f532385b73e407e5a6f09ee61920406a/kustomize:config/overlays/odh) |
| modelmesh-webhook-server-service | ClusterIP | 9443/TCP | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/modelmesh-serving/blob/2751b227f532385b73e407e5a6f09ee61920406a/kustomize:config/overlays/odh) |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| etcd | Ingress | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/modelmesh-serving/blob/2751b227f532385b73e407e5a6f09ee61920406a/kustomize:config/overlays/odh) |
| modelmesh-controller | Ingress | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/modelmesh-serving/blob/2751b227f532385b73e407e5a6f09ee61920406a/kustomize:config/overlays/odh) |
| modelmesh-runtimes | Ingress | [`config/rbac/common/networkpolicy-runtimes.yaml`](https://github.com/red-hat-data-services/modelmesh-serving/blob/2751b227f532385b73e407e5a6f09ee61920406a/config/rbac/common/networkpolicy-runtimes.yaml) |
| modelmesh-webhook | Ingress | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/modelmesh-serving/blob/2751b227f532385b73e407e5a6f09ee61920406a/kustomize:config/overlays/odh) |

## Network Policy Graph

Visual representation of NetworkPolicy rules. Ingress rules show what traffic is allowed into pods, egress rules show what traffic is allowed out.

```mermaid
graph LR
    classDef policy fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef pod fill:#3498db,stroke:#2980b9,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff

    modelmesh_serving["modelmesh-serving\nPods"]:::pod
    np_0_etcd{{"etcd\nIngress"}}:::policy
    src_0_0_0["ns: modelmesh-enabled=true"]:::external
    src_0_0_0 -->|ingress| np_0_etcd
    src_0_0_1["all pods"]:::external
    src_0_0_1 -->|ingress| np_0_etcd
    np_0_etcd --> modelmesh_serving
    np_1_modelmesh_controller{{"modelmesh-controller\nIngress"}}:::policy
    src_1_0_all["all pods"]:::external
    src_1_0_all -->|ingress| np_1_modelmesh_controller
    np_1_modelmesh_controller --> modelmesh_serving
    np_2_modelmesh_runtimes{{"modelmesh-runtimes\nIngress"}}:::policy
    src_2_0_0["pods: app.kubernetes.io/managed-by=modelmesh-controller"]:::external
    src_2_0_0 -->|ingress| np_2_modelmesh_runtimes
    src_2_1_all["all pods"]:::external
    src_2_1_all -->|ingress| np_2_modelmesh_runtimes
    src_2_2_all["all pods"]:::external
    src_2_2_all -->|ingress| np_2_modelmesh_runtimes
    np_2_modelmesh_runtimes --> modelmesh_serving
    np_3_modelmesh_webhook{{"modelmesh-webhook\nIngress"}}:::policy
    src_3_0_all["all pods"]:::external
    src_3_0_all -->|ingress| np_3_modelmesh_webhook
    np_3_modelmesh_webhook --> modelmesh_serving
```

