# RBAC Surface

52 cluster roles across the platform.

## Permission Scope by Component

```mermaid
graph TD
    classDef wide fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef medium fill:#f39c12,stroke:#e67e22,color:#fff
    classDef narrow fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff

    data_science_pipelines_operator["data-science-pipelines-operator\n4 roles"]:::component
    data_science_pipelines_operator --> role_1["manager-role\n55 resource types"]:::wide
    data_science_pipelines_operator --> role_2["manager-argo-role\n22 resource types"]:::medium
    data_science_pipelines_operator --> role_3["aggregate-dspa-admin-edit\n4 resource types"]:::narrow
    data_science_pipelines_operator --> role_4["1 more roles"]:::narrow
    kserve["kserve\n2 roles"]:::component
    kserve --> role_5["kserve-manager-role\n45 resource types"]:::wide
    kserve --> role_6["kserve-proxy-role\n2 resource types"]:::narrow
    model_registry_operator["model-registry-operator\n6 roles"]:::component
    model_registry_operator --> role_7["manager-role\n27 resource types"]:::medium
    model_registry_operator --> role_8["proxy-role\n2 resource types"]:::narrow
    model_registry_operator --> role_9["modelregistry-admin-role\n2 resource types"]:::narrow
    model_registry_operator --> role_10["3 more roles"]:::narrow
    odh_dashboard["odh-dashboard\n1 roles"]:::component
    odh_dashboard --> role_11["odh-dashboard\n40 resource types"]:::wide
    odh_model_controller["odh-model-controller\n7 roles"]:::component
    odh_model_controller --> role_12["odh-model-controller-role\n41 resource types"]:::wide
    odh_model_controller --> role_13["kserve-prometheus-k8s\n3 resource types"]:::narrow
    odh_model_controller --> role_14["account-editor-role\n2 resource types"]:::narrow
    odh_model_controller --> role_15["4 more roles"]:::narrow
    opendatahub_operator["opendatahub-operator\n23 roles"]:::component
    opendatahub_operator --> role_16["modelregistry-viewer-role\n2 resource types"]:::narrow
    opendatahub_operator --> role_17["kueue-viewer-role\n2 resource types"]:::narrow
    opendatahub_operator --> role_18["dashboard-viewer-role\n2 resource types"]:::narrow
    opendatahub_operator --> role_19["20 more roles"]:::narrow
    trustyai_service_operator["trustyai-service-operator\n9 roles"]:::component
    trustyai_service_operator --> role_20["manager-role\n44 resource types"]:::wide
    trustyai_service_operator --> role_21["evalhub-proxy-role\n4 resource types"]:::narrow
    trustyai_service_operator --> role_22["proxy-role\n2 resource types"]:::narrow
    trustyai_service_operator --> role_23["6 more roles"]:::narrow
```

## Roles by Component

| Owner | Role | Resource Types |
|-------|------|----------------|
| data-science-pipelines-operator | manager-role | 55 |
| data-science-pipelines-operator | manager-argo-role | 22 |
| data-science-pipelines-operator | aggregate-dspa-admin-edit | 4 |
| data-science-pipelines-operator | aggregate-dspa-admin-view | 4 |
| kserve | kserve-manager-role | 45 |
| kserve | kserve-proxy-role | 2 |
| model-registry-operator | manager-role | 27 |
| model-registry-operator | proxy-role | 2 |
| model-registry-operator | modelregistry-admin-role | 2 |
| model-registry-operator | modelregistry-editor-role | 2 |
| model-registry-operator | modelregistry-viewer-role | 2 |
| model-registry-operator | metrics-reader | 0 |
| odh-dashboard | odh-dashboard | 40 |
| odh-model-controller | odh-model-controller-role | 41 |
| odh-model-controller | kserve-prometheus-k8s | 3 |
| odh-model-controller | account-editor-role | 2 |
| odh-model-controller | account-viewer-role | 2 |
| odh-model-controller | proxy-role | 2 |
| odh-model-controller | metrics-auth-role | 2 |
| odh-model-controller | metrics-reader | 0 |
| opendatahub-operator | modelregistry-viewer-role | 2 |
| opendatahub-operator | kueue-viewer-role | 2 |
| opendatahub-operator | dashboard-viewer-role | 2 |
| opendatahub-operator | datasciencepipelines-editor-role | 2 |
| opendatahub-operator | datasciencepipelines-viewer-role | 2 |
| opendatahub-operator | kserve-editor-role | 2 |
| opendatahub-operator | kserve-viewer-role | 2 |
| opendatahub-operator | ray-viewer-role | 2 |
| opendatahub-operator | ray-editor-role | 2 |
| opendatahub-operator | modelregistry-editor-role | 2 |
| opendatahub-operator | dashboard-editor-role | 2 |
| opendatahub-operator | monitoring-viewer-role | 2 |
| opendatahub-operator | kueue-editor-role | 2 |
| opendatahub-operator | trainingoperator-editor-role | 2 |
| opendatahub-operator | trainingoperator-viewer-role | 2 |
| opendatahub-operator | trustyai-editor-role | 2 |
| opendatahub-operator | trustyai-viewer-role | 2 |
| opendatahub-operator | workbenches-editor-role | 2 |
| opendatahub-operator | workbenches-viewer-role | 2 |
| opendatahub-operator | auth-editor-role | 2 |
| opendatahub-operator | auth-viewer-role | 2 |
| opendatahub-operator | monitoring-editor-role | 2 |
| opendatahub-operator | metrics-reader | 0 |
| trustyai-service-operator | manager-role | 44 |
| trustyai-service-operator | evalhub-proxy-role | 4 |
| trustyai-service-operator | proxy-role | 2 |
| trustyai-service-operator | nemoguardrail-editor-role | 2 |
| trustyai-service-operator | nemoguardrail-viewer-role | 2 |
| trustyai-service-operator | lmeval-user-role | 2 |
| trustyai-service-operator | trustyaiservice-editor-role | 2 |
| trustyai-service-operator | trustyaiservice-viewer-role | 2 |
| trustyai-service-operator | metrics-reader | 0 |

