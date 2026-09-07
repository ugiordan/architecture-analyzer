# opendatahub-operator: Network

## Service Map

*1 unique services (3 total, duplicates from test fixtures collapsed).*

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    opendatahub_operator["opendatahub-operator"]:::component
    opendatahub_operator --> svc_0["webhook-service\nClusterIP: 443/TCP"]:::svc
    opendatahub_operator -.-> ext_s3[["s3\nobject-storage"]]:::ext
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| webhook-service | ClusterIP | 443/TCP | [`config/rhaii/webhook/service.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/rhaii/webhook/service.yaml) |
| webhook-service | ClusterIP | 443/TCP | [`config/rhoai/webhook/service.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/rhoai/webhook/service.yaml) |
| webhook-service | ClusterIP | 443/TCP | [`config/webhook/service.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/webhook/service.yaml) |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| deny-all-cascade-test | Ingress, Egress | [`cmd/mcp-server/scenarios/cascading-failure-networkpolicy.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/cmd/mcp-server/scenarios/cascading-failure-networkpolicy.yaml) |
| dscInit.Spec.ApplicationsNamespace | Ingress | [`internal/controller/dscinitialization/utils.go:160`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/dscinitialization/utils.go#L160) |

## Network Policy Graph

Visual representation of NetworkPolicy rules. Ingress rules show what traffic is allowed into pods, egress rules show what traffic is allowed out.

```mermaid
graph LR
    classDef policy fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef pod fill:#3498db,stroke:#2980b9,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff

    opendatahub_operator["opendatahub-operator\nPods"]:::pod
    np_0_deny_all_cascade_test{{"deny-all-cascade-test\nIngress, Egress"}}:::policy
    np_0_deny_all_cascade_test --> opendatahub_operator
    np_1_dscInit_Spec_ApplicationsNamespace{{"dscInit.Spec.ApplicationsNamespace\nIngress"}}:::policy
    np_1_dscInit_Spec_ApplicationsNamespace --> opendatahub_operator
```

