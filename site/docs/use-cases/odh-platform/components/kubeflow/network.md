# kubeflow: Network

[Overview](index.md) · [Agent Context](agent-context.md) · [Extracted Data](details.md) · [Network](network.md) · [RBAC](rbac.md) · [Security](security.md) · [Cache](cache.md) · [Dataflow](dataflow.md)

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
| service | ClusterIP | 443/TCP | [`components/notebook-controller/config/manager/service.yaml`](https://github.com/opendatahub-io/kubeflow/blob/6b078dbe8bb560f7af91b933e8abb3fc26ea1bf1/components/notebook-controller/config/manager/service.yaml) |
| service | ClusterIP | 8443/TCP | [`components/odh-notebook-controller/config/manager/service.yaml`](https://github.com/opendatahub-io/kubeflow/blob/6b078dbe8bb560f7af91b933e8abb3fc26ea1bf1/components/odh-notebook-controller/config/manager/service.yaml) |
| webhook-service | ClusterIP | 443/TCP | [`components/odh-notebook-controller/config/webhook/service.yaml`](https://github.com/opendatahub-io/kubeflow/blob/6b078dbe8bb560f7af91b933e8abb3fc26ea1bf1/components/odh-notebook-controller/config/webhook/service.yaml) |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| notebook-controller-allow-ports | Ingress, Egress | [`components/notebook-controller/config/overlays/openshift/network-policies.yaml`](https://github.com/opendatahub-io/kubeflow/blob/6b078dbe8bb560f7af91b933e8abb3fc26ea1bf1/components/notebook-controller/config/overlays/openshift/network-policies.yaml) |
| odh-notebook-controller-allow-ports | Ingress, Egress | [`components/odh-notebook-controller/config/overlays/openshift/network-policies.yaml`](https://github.com/opendatahub-io/kubeflow/blob/6b078dbe8bb560f7af91b933e8abb3fc26ea1bf1/components/odh-notebook-controller/config/overlays/openshift/network-policies.yaml) |

## Network Policy Graph

Visual representation of NetworkPolicy rules. Ingress rules show what traffic is allowed into pods, egress rules show what traffic is allowed out.

```mermaid
graph LR
    classDef policy fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef pod fill:#3498db,stroke:#2980b9,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff

    kubeflow["kubeflow\nPods"]:::pod
    np_0_notebook_controller_allow_ports{{"notebook-controller-allow-ports\nIngress, Egress"}}:::policy
    src_0_0_0["ns: kubernetes.io/metadata.name=openshift-monitoring"]:::external
    src_0_0_0 -->|ingress| np_0_notebook_controller_allow_ports
    src_0_0_1["ns: kubernetes.io/metadata.name=openshift-user-workload-monitoring"]:::external
    src_0_0_1 -->|ingress| np_0_notebook_controller_allow_ports
    src_0_0_2["ns: kubernetes.io/metadata.name=redhat-ods-monitoring"]:::external
    src_0_0_2 -->|ingress| np_0_notebook_controller_allow_ports
    src_0_0_3["ns: kubernetes.io/metadata.name=opendatahub-monitoring"]:::external
    src_0_0_3 -->|ingress| np_0_notebook_controller_allow_ports
    src_0_1_all["all pods"]:::external
    src_0_1_all -->|ingress| np_0_notebook_controller_allow_ports
    np_0_notebook_controller_allow_ports --> kubeflow
    dst_0_0_all["all destinations"]:::external
    kubeflow -->|egress| dst_0_0_all
    np_1_odh_notebook_controller_allow_ports{{"odh-notebook-controller-allow-ports\nIngress, Egress"}}:::policy
    src_1_0_all["all pods"]:::external
    src_1_0_all -->|ingress| np_1_odh_notebook_controller_allow_ports
    src_1_1_0["ns: kubernetes.io/metadata.name=openshift-monitoring"]:::external
    src_1_1_0 -->|ingress| np_1_odh_notebook_controller_allow_ports
    src_1_1_1["ns: kubernetes.io/metadata.name=openshift-user-workload-monitoring"]:::external
    src_1_1_1 -->|ingress| np_1_odh_notebook_controller_allow_ports
    src_1_1_2["ns: kubernetes.io/metadata.name=redhat-ods-monitoring"]:::external
    src_1_1_2 -->|ingress| np_1_odh_notebook_controller_allow_ports
    src_1_1_3["ns: kubernetes.io/metadata.name=opendatahub-monitoring"]:::external
    src_1_1_3 -->|ingress| np_1_odh_notebook_controller_allow_ports
    src_1_2_all["all pods"]:::external
    src_1_2_all -->|ingress| np_1_odh_notebook_controller_allow_ports
    np_1_odh_notebook_controller_allow_ports --> kubeflow
    dst_1_0_all["all destinations"]:::external
    kubeflow -->|egress| dst_1_0_all
```

