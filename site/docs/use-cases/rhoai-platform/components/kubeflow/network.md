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
| service | ClusterIP | 443/TCP | [`components/notebook-controller/config/manager/service.yaml`](https://github.com/red-hat-data-services/kubeflow/blob/5b7ff478aecf00c6954ce097bf77741e3fe40715/components/notebook-controller/config/manager/service.yaml) |
| service | ClusterIP | 8443/TCP | [`components/odh-notebook-controller/config/manager/service.yaml`](https://github.com/red-hat-data-services/kubeflow/blob/5b7ff478aecf00c6954ce097bf77741e3fe40715/components/odh-notebook-controller/config/manager/service.yaml) |
| webhook-service | ClusterIP | 443/TCP | [`components/odh-notebook-controller/config/webhook/service.yaml`](https://github.com/red-hat-data-services/kubeflow/blob/5b7ff478aecf00c6954ce097bf77741e3fe40715/components/odh-notebook-controller/config/webhook/service.yaml) |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| notebook-controller-allow-ports | Ingress, Egress | [`components/notebook-controller/config/overlays/openshift/network-policies.yaml`](https://github.com/red-hat-data-services/kubeflow/blob/5b7ff478aecf00c6954ce097bf77741e3fe40715/components/notebook-controller/config/overlays/openshift/network-policies.yaml) |
| odh-notebook-controller-allow-ports | Ingress, Egress | [`components/odh-notebook-controller/config/overlays/openshift/network-policies.yaml`](https://github.com/red-hat-data-services/kubeflow/blob/5b7ff478aecf00c6954ce097bf77741e3fe40715/components/odh-notebook-controller/config/overlays/openshift/network-policies.yaml) |

## Network Policy Graph

Visual representation of NetworkPolicy rules. Ingress rules show what traffic is allowed into pods, egress rules show what traffic is allowed out.

```mermaid
graph LR
    classDef policy fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef pod fill:#3498db,stroke:#2980b9,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff

    kubeflow["kubeflow\nPods"]:::pod
    np_0_notebook_controller_allow_ports{{"notebook-controller-allow-ports\nIngress, Egress"}}:::policy
    np_0_notebook_controller_allow_ports --> kubeflow
    np_1_odh_notebook_controller_allow_ports{{"odh-notebook-controller-allow-ports\nIngress, Egress"}}:::policy
    np_1_odh_notebook_controller_allow_ports --> kubeflow
```

