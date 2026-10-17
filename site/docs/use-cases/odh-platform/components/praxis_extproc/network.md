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
| payload-processing | ClusterIP | 9004/TCP | [`deploy/base/instance/service.yaml`](https://github.com/opendatahub-io/praxis-extproc/blob/06db0b9e5fece6b523332a5990346abd8148c5e9/deploy/base/instance/service.yaml) |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| payload-processing | Ingress, Egress | [`deploy/overlays/odh/networking/networkpolicy.yaml`](https://github.com/opendatahub-io/praxis-extproc/blob/06db0b9e5fece6b523332a5990346abd8148c5e9/deploy/overlays/odh/networking/networkpolicy.yaml) |

## Network Policy Graph

Visual representation of NetworkPolicy rules. Ingress rules show what traffic is allowed into pods, egress rules show what traffic is allowed out.

```mermaid
graph LR
    classDef policy fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef pod fill:#3498db,stroke:#2980b9,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff

    praxis_extproc["praxis-extproc\nPods"]:::pod
    np_0_payload_processing{{"payload-processing\nIngress, Egress"}}:::policy
    np_0_payload_processing --> praxis_extproc
```

