# mlflow-operator: Network

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    mlflow_operator["mlflow-operator"]:::component
    mlflow_operator --> svc_0["mlflow-operator-controller-manager-metrics-service\nClusterIP: 8443/TCP"]:::svc
    mlflow_operator -.-> ext_mlflow[["mlflow\napi"]]:::ext
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| mlflow-operator-controller-manager-metrics-service | ClusterIP | 8443/TCP | [`kustomize:config/overlays/odh`](https://github.com/opendatahub-io/mlflow-operator/blob/6df31b94f07284e6e0465e174e3b0c039d219726/kustomize:config/overlays/odh) |

### Ingress / Routing

| Kind | Name | Hosts | Paths | TLS | Source |
|------|------|-------|-------|-----|--------|
| HTTPRoute | rbac-inferred |  |  | no | [`rbac/manager-role`](https://github.com/opendatahub-io/mlflow-operator/blob/6df31b94f07284e6e0465e174e3b0c039d219726/rbac/manager-role) |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| mlflow-operator-controller-manager | Ingress | [`kustomize:config/overlays/odh`](https://github.com/opendatahub-io/mlflow-operator/blob/6df31b94f07284e6e0465e174e3b0c039d219726/kustomize:config/overlays/odh) |

## Network Policy Graph

Visual representation of NetworkPolicy rules. Ingress rules show what traffic is allowed into pods, egress rules show what traffic is allowed out.

```mermaid
graph LR
    classDef policy fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef pod fill:#3498db,stroke:#2980b9,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff

    mlflow_operator["mlflow-operator\nPods"]:::pod
    np_0_mlflow_operator_controller_manager{{"mlflow-operator-controller-manager\nIngress"}}:::policy
    np_0_mlflow_operator_controller_manager --> mlflow_operator
```

