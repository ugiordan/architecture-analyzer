# kubeflow: Network

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    kubeflow["kubeflow"]:::component
    kubeflow --> svc_0["service\nClusterIP: 443/TCP"]:::svc
    kubeflow --> svc_1["service\nClusterIP: 8443/TCP"]:::svc
    kubeflow --> svc_2["webhook-service\nClusterIP: 443/TCP"]:::svc
    kubeflow -.-> ext_requests[["requests\napi"]]:::ext
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| service | ClusterIP | 443/TCP | [`components/notebook-controller/config/manager/service.yaml`](https://github.com/opendatahub-io/kubeflow/blob/1832b8a731ac724c32855fef4d5dc69910f46ebc/components/notebook-controller/config/manager/service.yaml) |
| service | ClusterIP | 8443/TCP | [`components/odh-notebook-controller/config/manager/service.yaml`](https://github.com/opendatahub-io/kubeflow/blob/1832b8a731ac724c32855fef4d5dc69910f46ebc/components/odh-notebook-controller/config/manager/service.yaml) |
| webhook-service | ClusterIP | 443/TCP | [`components/odh-notebook-controller/config/webhook/service.yaml`](https://github.com/opendatahub-io/kubeflow/blob/1832b8a731ac724c32855fef4d5dc69910f46ebc/components/odh-notebook-controller/config/webhook/service.yaml) |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| notebook-controller-allow-ingress | Ingress | [`components/notebook-controller/config/overlays/openshift/network-policies.yaml`](https://github.com/opendatahub-io/kubeflow/blob/1832b8a731ac724c32855fef4d5dc69910f46ebc/components/notebook-controller/config/overlays/openshift/network-policies.yaml) |
| odh-notebook-controller-allow-ingress | Ingress | [`components/odh-notebook-controller/config/overlays/openshift/network-policies.yaml`](https://github.com/opendatahub-io/kubeflow/blob/1832b8a731ac724c32855fef4d5dc69910f46ebc/components/odh-notebook-controller/config/overlays/openshift/network-policies.yaml) |

## Network Policy Graph

Visual representation of NetworkPolicy rules. Ingress rules show what traffic is allowed into pods, egress rules show what traffic is allowed out.

```mermaid
graph LR
    classDef policy fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef pod fill:#3498db,stroke:#2980b9,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff

    kubeflow["kubeflow\nPods"]:::pod
    np_0_notebook_controller_allow_ingress{{"notebook-controller-allow-ingress\nIngress"}}:::policy
    np_0_notebook_controller_allow_ingress --> kubeflow
    np_1_odh_notebook_controller_allow_ingress{{"odh-notebook-controller-allow-ingress\nIngress"}}:::policy
    np_1_odh_notebook_controller_allow_ingress --> kubeflow
```

