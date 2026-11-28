# opendatahub-operator: Network

[Overview](index.md) · [Agent Context](agent-context.md) · [Extracted Data](details.md) · [Network](network.md) · [RBAC](rbac.md) · [Security](security.md) · [Cache](cache.md) · [Dataflow](dataflow.md)

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
| webhook-service | ClusterIP | 443/TCP | [`config/rhaii/webhook/service.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/aa0297559a280eb7030cc513cdba70d4243a44c2/config/rhaii/webhook/service.yaml) |
| webhook-service | ClusterIP | 443/TCP | [`config/rhoai/webhook/service.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/aa0297559a280eb7030cc513cdba70d4243a44c2/config/rhoai/webhook/service.yaml) |
| webhook-service | ClusterIP | 443/TCP | [`config/webhook/service.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/aa0297559a280eb7030cc513cdba70d4243a44c2/config/webhook/service.yaml) |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| deny-all-cascade-test | Ingress, Egress | [`cmd/mcp-server/scenarios/cascading-failure-networkpolicy.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/aa0297559a280eb7030cc513cdba70d4243a44c2/cmd/mcp-server/scenarios/cascading-failure-networkpolicy.yaml) |
| dscInit.Spec.ApplicationsNamespace | Ingress | [`internal/controller/dscinitialization/utils.go:164`](https://github.com/opendatahub-io/opendatahub-operator/blob/aa0297559a280eb7030cc513cdba70d4243a44c2/internal/controller/dscinitialization/utils.go#L164) |

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
    src_1_0_0["ns: labels.ODH.OwnedNamespace=labels.True"]:::external
    src_1_0_0 -->|ingress| np_1_dscInit_Spec_ApplicationsNamespace
    src_1_1_0["ns: labels.CustomizedAppNamespace=labels.True"]:::external
    src_1_1_0 -->|ingress| np_1_dscInit_Spec_ApplicationsNamespace
    src_1_2_0["ns: network.openshift.io/policy-group=ingress"]:::external
    src_1_2_0 -->|ingress| np_1_dscInit_Spec_ApplicationsNamespace
    src_1_3_0["ns: kubernetes.io/metadata.name=openshift-host-network"]:::external
    src_1_3_0 -->|ingress| np_1_dscInit_Spec_ApplicationsNamespace
    src_1_4_0["ns: kubernetes.io/metadata.name=openshift-monitoring"]:::external
    src_1_4_0 -->|ingress| np_1_dscInit_Spec_ApplicationsNamespace
    src_1_5_0["ns: kubernetes.io/metadata.name=openshift-cluster-observability-operator"]:::external
    src_1_5_0 -->|ingress| np_1_dscInit_Spec_ApplicationsNamespace
    np_1_dscInit_Spec_ApplicationsNamespace --> opendatahub_operator
```

