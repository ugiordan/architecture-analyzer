# data-science-pipelines: Network

[Overview](index.md) · [Agent Context](agent-context.md) · [Extracted Data](details.md) · [Network](network.md) · [RBAC](rbac.md) · [Security](security.md) · [Cache](cache.md) · [Dataflow](dataflow.md)

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    data_science_pipelines["data-science-pipelines"]:::component
    data_science_pipelines --> svc_0["kubeflow-pipelines-profile-controller\nClusterIP: 80/TCP"]:::svc
    data_science_pipelines -.-> ext_requests[["requests\napi"]]:::ext
    data_science_pipelines -.-> ext_grpc[["grpc\ngrpc"]]:::ext
    data_science_pipelines -.-> ext_gcs[["gcs\nobject-storage"]]:::ext
    data_science_pipelines -.-> ext_minio[["minio\nobject-storage"]]:::ext
    data_science_pipelines -.-> ext_s3[["s3\nobject-storage"]]:::ext
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| kubeflow-pipelines-profile-controller | ClusterIP | 80/TCP | [`manifests/kustomize/base/installs/multi-user/pipelines-profile-controller/service.yaml`](https://github.com/opendatahub-io/data-science-pipelines/blob/99d5da66f758d22eb96f8bda2ae46161100e75a6/manifests/kustomize/base/installs/multi-user/pipelines-profile-controller/service.yaml) |

### Ingress / Routing

| Kind | Name | Hosts | Paths | TLS | Source |
|------|------|-------|-------|-----|--------|
| DestinationRule | rbac-inferred |  |  | no | [`rbac/kubeflow-metacontroller`](https://github.com/opendatahub-io/data-science-pipelines/blob/99d5da66f758d22eb96f8bda2ae46161100e75a6/rbac/kubeflow-metacontroller) |
| Route | ml-pipeline-ui |  |  | yes | [`manifests/kustomize/env/openshift/base/route.yaml`](https://github.com/opendatahub-io/data-science-pipelines/blob/99d5da66f758d22eb96f8bda2ae46161100e75a6/manifests/kustomize/env/openshift/base/route.yaml) |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| seaweedfs | Ingress | [`manifests/kustomize/third-party/seaweedfs/base/seaweedfs/seaweedfs-networkpolicy.yaml`](https://github.com/opendatahub-io/data-science-pipelines/blob/99d5da66f758d22eb96f8bda2ae46161100e75a6/manifests/kustomize/third-party/seaweedfs/base/seaweedfs/seaweedfs-networkpolicy.yaml) |

## Network Policy Graph

Visual representation of NetworkPolicy rules. Ingress rules show what traffic is allowed into pods, egress rules show what traffic is allowed out.

```mermaid
graph LR
    classDef policy fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef pod fill:#3498db,stroke:#2980b9,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff

    data_science_pipelines["data-science-pipelines\nPods"]:::pod
    np_0_seaweedfs{{"seaweedfs\nIngress"}}:::policy
    src_0_0_0["all namespaces"]:::external
    src_0_0_0 -->|ingress| np_0_seaweedfs
    src_0_1_0["all namespaces"]:::external
    src_0_1_0 -->|ingress| np_0_seaweedfs
    src_0_2_0["all pods"]:::external
    src_0_2_0 -->|ingress| np_0_seaweedfs
    np_0_seaweedfs --> data_science_pipelines
```

