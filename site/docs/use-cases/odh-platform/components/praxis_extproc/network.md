# praxis-extproc: Network

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    praxis_extproc["praxis-extproc"]:::component
    praxis_extproc --> svc_0["payload-processing\nClusterIP: 9004/TCP"]:::svc
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| payload-processing | ClusterIP | 9004/TCP | [`deploy/base/instance/service.yaml`](https://github.com/opendatahub-io/praxis-extproc/blob/814a177260f1d3ce5f2e62f5cdfd8c7c532ea710/deploy/base/instance/service.yaml) |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| payload-processing | Ingress, Egress | [`deploy/overlays/odh/networking/networkpolicy.yaml`](https://github.com/opendatahub-io/praxis-extproc/blob/814a177260f1d3ce5f2e62f5cdfd8c7c532ea710/deploy/overlays/odh/networking/networkpolicy.yaml) |

## Network Policy Graph

Visual representation of NetworkPolicy rules. Ingress rules show what traffic is allowed into pods, egress rules show what traffic is allowed out.

```mermaid
graph LR
    classDef policy fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef pod fill:#3498db,stroke:#2980b9,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff

    praxis_extproc["praxis-extproc\nPods"]:::pod
    np_0_payload_processing{{"payload-processing\nIngress, Egress"}}:::policy
    src_0_0_0["ns: kubernetes.io/metadata.name=openshift-ingress"]:::external
    src_0_0_0 -->|ingress| np_0_payload_processing
    src_0_1_0["ns: kubernetes.io/metadata.name=openshift-monitoring"]:::external
    src_0_1_0 -->|ingress| np_0_payload_processing
    src_0_1_1["ns: kubernetes.io/metadata.name=openshift-user-workload-monitoring"]:::external
    src_0_1_1 -->|ingress| np_0_payload_processing
    src_0_1_2["ns: kubernetes.io/metadata.name=redhat-ods-monitoring"]:::external
    src_0_1_2 -->|ingress| np_0_payload_processing
    np_0_payload_processing --> praxis_extproc
    dst_0_0_all["all destinations"]:::external
    praxis_extproc -->|egress| dst_0_0_all
    dst_0_1_all["all destinations"]:::external
    praxis_extproc -->|egress| dst_0_1_all
```

