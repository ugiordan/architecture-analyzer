# odh-model-controller: Network

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    odh_model_controller["odh-model-controller"]:::component
    odh_model_controller --> svc_0["model-serving-api\nClusterIP: 443/TCP,8080/TCP"]:::svc
    odh_model_controller --> svc_1["odh-model-controller-metrics-service\nClusterIP: 8443/TCP"]:::svc
    odh_model_controller --> svc_2["odh-model-controller-webhook-service\nClusterIP: 443/TCP"]:::svc
    odh_model_controller -.-> ext_s3[["s3\nobject-storage"]]:::ext
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| model-serving-api | ClusterIP | 443/TCP, 8080/TCP | [`kustomize:config/overlays/odh`](https://github.com/opendatahub-io/odh-model-controller/blob/43ebf90dd5af130a94cb3b1857f6ebd45ecd28cf/kustomize:config/overlays/odh) |
| odh-model-controller-metrics-service | ClusterIP | 8443/TCP | [`kustomize:config/overlays/odh`](https://github.com/opendatahub-io/odh-model-controller/blob/43ebf90dd5af130a94cb3b1857f6ebd45ecd28cf/kustomize:config/overlays/odh) |
| odh-model-controller-webhook-service | ClusterIP | 443/TCP | [`kustomize:config/overlays/odh`](https://github.com/opendatahub-io/odh-model-controller/blob/43ebf90dd5af130a94cb3b1857f6ebd45ecd28cf/kustomize:config/overlays/odh) |

### Ingress / Routing

| Kind | Name | Hosts | Paths | TLS | Source |
|------|------|-------|-------|-----|--------|
| Gateway | rbac-inferred |  |  | no | [`rbac/odh-model-controller-role`](https://github.com/opendatahub-io/odh-model-controller/blob/43ebf90dd5af130a94cb3b1857f6ebd45ecd28cf/rbac/odh-model-controller-role) |
| HTTPRoute | rbac-inferred |  |  | no | [`rbac/odh-model-controller-role`](https://github.com/opendatahub-io/odh-model-controller/blob/43ebf90dd5af130a94cb3b1857f6ebd45ecd28cf/rbac/odh-model-controller-role) |
| Ingress | rbac-inferred |  |  | no | [`rbac/odh-model-controller-role`](https://github.com/opendatahub-io/odh-model-controller/blob/43ebf90dd5af130a94cb3b1857f6ebd45ecd28cf/rbac/odh-model-controller-role) |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| odh-model-controller-webhook-allow-ingress |  | [`kustomize:config/overlays/odh`](https://github.com/opendatahub-io/odh-model-controller/blob/43ebf90dd5af130a94cb3b1857f6ebd45ecd28cf/kustomize:config/overlays/odh) |

## Network Policy Graph

Visual representation of NetworkPolicy rules. Ingress rules show what traffic is allowed into pods, egress rules show what traffic is allowed out.

```mermaid
graph LR
    classDef policy fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef pod fill:#3498db,stroke:#2980b9,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff

    odh_model_controller["odh-model-controller\nPods"]:::pod
    np_0_odh_model_controller_webhook_allow_ingress{{"odh-model-controller-webhook-allow-ingress\nIngress"}}:::policy
    np_0_odh_model_controller_webhook_allow_ingress --> odh_model_controller
```

