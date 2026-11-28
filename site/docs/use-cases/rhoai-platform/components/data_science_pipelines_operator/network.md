# data-science-pipelines-operator: Network

[Overview](index.md) · [Agent Context](agent-context.md) · [Extracted Data](details.md) · [Network](network.md) · [RBAC](rbac.md) · [Security](security.md) · [Cache](cache.md) · [Dataflow](dataflow.md)

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    data_science_pipelines_operator["data-science-pipelines-operator"]:::component
    data_science_pipelines_operator --> svc_0["data-science-pipelines-operator-service\nClusterIP: 8443/TCP"]:::svc
    data_science_pipelines_operator --> svc_1["ds-pipeline-workflow-controller-metrics-template-value\nClusterIP: 9090/TCP"]:::svc
    data_science_pipelines_operator --> svc_2["mariadb-template-value\nClusterIP: 3306/TCP"]:::svc
    data_science_pipelines_operator --> svc_3["minio-service\nClusterIP: 9000/TCP"]:::svc
    data_science_pipelines_operator --> svc_4["minio-template-value\nClusterIP: 80/TCP,9000/TCP"]:::svc
    data_science_pipelines_operator --> svc_5["ml-pipeline\nClusterIP: 8443/TCP,8887/TCP,8888/TCP"]:::svc
    data_science_pipelines_operator --> svc_6["template-value\nClusterIP: 8443/TCP,8445/TCP,8887/TCP,8888/TCP"]:::svc
    data_science_pipelines_operator --> svc_7["template-value\nClusterIP: 8443/TCP"]:::svc
    data_science_pipelines_operator -.-> ext_requests[["requests\napi"]]:::ext
    data_science_pipelines_operator -.-> ext_minio[["minio\nobject-storage"]]:::ext
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| data-science-pipelines-operator-service | ClusterIP | 8443/TCP | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/18ea33cf3b3ba58dbc92e80c673c55517144f36e/kustomize:config/overlays/odh) |
| ds-pipeline-workflow-controller-metrics-template-value | ClusterIP | 9090/TCP | [`config/internal/workflow-controller/service.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/18ea33cf3b3ba58dbc92e80c673c55517144f36e/config/internal/workflow-controller/service.yaml.tmpl) |
| mariadb-template-value | ClusterIP | 3306/TCP | [`config/internal/mariadb/default/service.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/18ea33cf3b3ba58dbc92e80c673c55517144f36e/config/internal/mariadb/default/service.yaml.tmpl) |
| minio-service | ClusterIP | 9000/TCP | [`config/internal/minio/default/service.minioservice.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/18ea33cf3b3ba58dbc92e80c673c55517144f36e/config/internal/minio/default/service.minioservice.yaml.tmpl) |
| minio-template-value | ClusterIP | 9000/TCP, 80/TCP | [`config/internal/minio/default/service.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/18ea33cf3b3ba58dbc92e80c673c55517144f36e/config/internal/minio/default/service.yaml.tmpl) |
| ml-pipeline | ClusterIP | 8443/TCP, 8888/TCP, 8887/TCP | [`config/internal/apiserver/default/service.ml-pipeline.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/18ea33cf3b3ba58dbc92e80c673c55517144f36e/config/internal/apiserver/default/service.ml-pipeline.yaml.tmpl) |
| template-value | ClusterIP | 8443/TCP, 8445/TCP, 8888/TCP, 8887/TCP | [`config/internal/apiserver/default/service.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/18ea33cf3b3ba58dbc92e80c673c55517144f36e/config/internal/apiserver/default/service.yaml.tmpl) |
| template-value | ClusterIP | 8443/TCP | [`config/internal/webhook/service.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/18ea33cf3b3ba58dbc92e80c673c55517144f36e/config/internal/webhook/service.yaml.tmpl) |

### Ingress / Routing

| Kind | Name | Hosts | Paths | TLS | Source |
|------|------|-------|-------|-----|--------|
| Ingress | rbac-inferred |  |  | no | [`rbac/manager-role`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/18ea33cf3b3ba58dbc92e80c673c55517144f36e/rbac/manager-role) |
| Route | ds-pipeline-md-template-value |  |  | yes | [`config/internal/ml-metadata/route/metadata-envoy.route.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/18ea33cf3b3ba58dbc92e80c673c55517144f36e/config/internal/ml-metadata/route/metadata-envoy.route.yaml.tmpl) |
| Route | minio-template-value |  |  | yes | [`config/internal/minio/route.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/18ea33cf3b3ba58dbc92e80c673c55517144f36e/config/internal/minio/route.yaml.tmpl) |
| Route | template-value |  |  | yes | [`config/internal/apiserver/route/route.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/18ea33cf3b3ba58dbc92e80c673c55517144f36e/config/internal/apiserver/route/route.yaml.tmpl) |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| controller-manager | Ingress | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/18ea33cf3b3ba58dbc92e80c673c55517144f36e/kustomize:config/overlays/odh) |
| ds-pipeline-drivers-template-value | Ingress, Egress | [`config/internal/common/default/driver-networkpolicy.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/18ea33cf3b3ba58dbc92e80c673c55517144f36e/config/internal/common/default/driver-networkpolicy.yaml.tmpl) |
| ds-pipeline-metadata-grpc-template-value | Ingress, Egress | [`config/internal/ml-metadata/metadata-grpc.networkpolicy.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/18ea33cf3b3ba58dbc92e80c673c55517144f36e/config/internal/ml-metadata/metadata-grpc.networkpolicy.yaml.tmpl) |
| mariadb-template-value | Ingress, Egress | [`config/internal/mariadb/default/networkpolicy.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/18ea33cf3b3ba58dbc92e80c673c55517144f36e/config/internal/mariadb/default/networkpolicy.yaml.tmpl) |
| minio-template-value | Ingress, Egress | [`config/internal/minio/default/networkpolicy.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/18ea33cf3b3ba58dbc92e80c673c55517144f36e/config/internal/minio/default/networkpolicy.yaml.tmpl) |
| template-value | Ingress, Egress | [`config/internal/persistence-agent/networkpolicy.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/18ea33cf3b3ba58dbc92e80c673c55517144f36e/config/internal/persistence-agent/networkpolicy.yaml.tmpl) |
| template-value | Ingress, Egress | [`config/internal/scheduled-workflow/networkpolicy.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/18ea33cf3b3ba58dbc92e80c673c55517144f36e/config/internal/scheduled-workflow/networkpolicy.yaml.tmpl) |
| template-value | Ingress | [`config/internal/webhook/networkpolicy.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/18ea33cf3b3ba58dbc92e80c673c55517144f36e/config/internal/webhook/networkpolicy.yaml.tmpl) |

## Network Policy Graph

Visual representation of NetworkPolicy rules. Ingress rules show what traffic is allowed into pods, egress rules show what traffic is allowed out.

```mermaid
graph LR
    classDef policy fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef pod fill:#3498db,stroke:#2980b9,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff

    data_science_pipelines_operator["data-science-pipelines-operator\nPods"]:::pod
    np_0_controller_manager{{"controller-manager\nIngress"}}:::policy
    src_0_0_all["all pods"]:::external
    src_0_0_all -->|ingress| np_0_controller_manager
    src_0_1_all["all pods"]:::external
    src_0_1_all -->|ingress| np_0_controller_manager
    src_0_2_0["ns: kubernetes.io/metadata.name=openshift-monitoring"]:::external
    src_0_2_0 -->|ingress| np_0_controller_manager
    src_0_2_1["ns: kubernetes.io/metadata.name=openshift-user-workload-monitoring"]:::external
    src_0_2_1 -->|ingress| np_0_controller_manager
    src_0_2_2["ns: kubernetes.io/metadata.name=redhat-ods-monitoring"]:::external
    src_0_2_2 -->|ingress| np_0_controller_manager
    src_0_2_3["ns: kubernetes.io/metadata.name=opendatahub-monitoring"]:::external
    src_0_2_3 -->|ingress| np_0_controller_manager
    np_0_controller_manager --> data_science_pipelines_operator
    np_1_ds_pipeline_drivers_template_value{{"ds-pipeline-drivers-template-value\nIngress, Egress"}}:::policy
    np_1_ds_pipeline_drivers_template_value --> data_science_pipelines_operator
    dst_1_0_all["all destinations"]:::external
    data_science_pipelines_operator -->|egress| dst_1_0_all
    dst_1_1_all["all destinations"]:::external
    data_science_pipelines_operator -->|egress| dst_1_1_all
    dst_1_2_0["pods: app=template-value, component=data-science-pipelines"]:::external
    data_science_pipelines_operator -->|egress| dst_1_2_0
    dst_1_3_0["pods: app=ds-pipeline-metadata-grpc-template-value, component=data-science-pipelines"]:::external
    data_science_pipelines_operator -->|egress| dst_1_3_0
    dst_1_4_0["ns: kubernetes.io/metadata.name=template-value"]:::external
    data_science_pipelines_operator -->|egress| dst_1_4_0
    np_2_ds_pipeline_metadata_grpc_template_value{{"ds-pipeline-metadata-grpc-template-value\nIngress, Egress"}}:::policy
    src_2_0_0["pods: app=template-value, component=data-science-pipelines"]:::external
    src_2_0_0 -->|ingress| np_2_ds_pipeline_metadata_grpc_template_value
    src_2_0_1["pods: app=ds-pipeline-metadata-envoy-template-value, component=data-science-pipelines"]:::external
    src_2_0_1 -->|ingress| np_2_ds_pipeline_metadata_grpc_template_value
    src_2_0_2["pods: pipelines.kubeflow.org/v2_component=true"]:::external
    src_2_0_2 -->|ingress| np_2_ds_pipeline_metadata_grpc_template_value
    np_2_ds_pipeline_metadata_grpc_template_value --> data_science_pipelines_operator
    dst_2_0_all["all destinations"]:::external
    data_science_pipelines_operator -->|egress| dst_2_0_all
    dst_2_1_0["pods: app=mariadb-template-value, component=data-science-pipelines"]:::external
    data_science_pipelines_operator -->|egress| dst_2_1_0
    dst_2_2_all["all destinations"]:::external
    data_science_pipelines_operator -->|egress| dst_2_2_all
    np_3_mariadb_template_value{{"mariadb-template-value\nIngress, Egress"}}:::policy
    src_3_0_0["ns: kubernetes.io/metadata.name=template-value"]:::external
    src_3_0_0 -->|ingress| np_3_mariadb_template_value
    src_3_0_1["pods: app=template-value, component=data-science-pipelines"]:::external
    src_3_0_1 -->|ingress| np_3_mariadb_template_value
    src_3_0_2["pods: app=ds-pipeline-metadata-grpc-template-value, component=data-science-pipelines"]:::external
    src_3_0_2 -->|ingress| np_3_mariadb_template_value
    np_3_mariadb_template_value --> data_science_pipelines_operator
    np_4_minio_template_value{{"minio-template-value\nIngress, Egress"}}:::policy
    src_4_0_0["pods: app=template-value, component=data-science-pipelines"]:::external
    src_4_0_0 -->|ingress| np_4_minio_template_value
    src_4_0_1["pods: pipelines.kubeflow.org/v2_component=true"]:::external
    src_4_0_1 -->|ingress| np_4_minio_template_value
    src_4_0_2["ns: kubernetes.io/metadata.name=template-value"]:::external
    src_4_0_2 -->|ingress| np_4_minio_template_value
    src_4_0_3["ns: policy-group.network.openshift.io/ingress="]:::external
    src_4_0_3 -->|ingress| np_4_minio_template_value
    np_4_minio_template_value --> data_science_pipelines_operator
    np_5_template_value{{"template-value\nIngress, Egress"}}:::policy
    np_5_template_value --> data_science_pipelines_operator
    dst_5_0_all["all destinations"]:::external
    data_science_pipelines_operator -->|egress| dst_5_0_all
    dst_5_1_all["all destinations"]:::external
    data_science_pipelines_operator -->|egress| dst_5_1_all
    dst_5_2_all["all destinations"]:::external
    data_science_pipelines_operator -->|egress| dst_5_2_all
    dst_5_3_0["pods: app=template-value, component=data-science-pipelines"]:::external
    data_science_pipelines_operator -->|egress| dst_5_3_0
    np_6_template_value{{"template-value\nIngress, Egress"}}:::policy
    np_6_template_value --> data_science_pipelines_operator
    dst_6_0_all["all destinations"]:::external
    data_science_pipelines_operator -->|egress| dst_6_0_all
    dst_6_1_all["all destinations"]:::external
    data_science_pipelines_operator -->|egress| dst_6_1_all
    dst_6_2_all["all destinations"]:::external
    data_science_pipelines_operator -->|egress| dst_6_2_all
    dst_6_3_0["pods: app=template-value, component=data-science-pipelines"]:::external
    data_science_pipelines_operator -->|egress| dst_6_3_0
    np_7_template_value{{"template-value\nIngress"}}:::policy
    src_7_0_all["all pods"]:::external
    src_7_0_all -->|ingress| np_7_template_value
    np_7_template_value --> data_science_pipelines_operator
```

