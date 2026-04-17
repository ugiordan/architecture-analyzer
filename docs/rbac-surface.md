# RBAC Surface

52 cluster roles across the platform.

## Permission Scope by Component

Each bar shows the widest role (by resource type count). Scope: 🔴 wide (>30), 🟠 medium (10-30), 🟢 narrow (<10).

```mermaid
%%{init: {'theme': 'base', 'themeVariables': {'fontSize': '12px'}, 'xyChart': {'width': 700, 'height': 380}}}%%
xychart-beta horizontal
    title "Widest Role Scope (resource types)"
    x-axis ["data-science-pipelines-operator", "kserve", "model-registry-operator", "odh-dashboard", "odh-model-controller", "opendatahub-operator", "trustyai-service-operator"]
    y-axis "Resource types" 0 --> 60
    bar [55, 45, 27, 40, 41, 2, 44]
```

## RBAC Binding Graph

Subject-to-role bindings across all platform components. Edge direction shows who has access to what.

```mermaid
graph LR
    classDef role fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef subject fill:#3498db,stroke:#2980b9,color:#fff

    sa_controller_manager["controller-manager\nServiceAccount"]:::subject
    role_manager_argo_role["manager-argo-role"]:::role
    sa_controller_manager -->|data-science-pipelines-operator| role_manager_argo_role
    role_manager_role["manager-role"]:::role
    sa_controller_manager -->|data-science-pipelines-operator| role_manager_role
    sa_kserve_controller_manager["kserve-controller-manager\nServiceAccount"]:::subject
    role_kserve_proxy_role["kserve-proxy-role"]:::role
    sa_kserve_controller_manager -->|kserve| role_kserve_proxy_role
    role_kserve_manager_role["kserve-manager-role"]:::role
    sa_kserve_controller_manager -->|kserve| role_kserve_manager_role
    role_proxy_role["proxy-role"]:::role
    sa_controller_manager -->|model-registry-operator| role_proxy_role
    sa_controller_manager -->|model-registry-operator| role_manager_role
    sa_odh_dashboard["odh-dashboard\nServiceAccount"]:::subject
    role_system_auth_delegator["system:auth-delegator"]:::role
    sa_odh_dashboard -->|odh-dashboard| role_system_auth_delegator
    role_cluster_monitoring_view["cluster-monitoring-view"]:::role
    sa_odh_dashboard -->|odh-dashboard| role_cluster_monitoring_view
    role_odh_dashboard["odh-dashboard"]:::role
    sa_odh_dashboard -->|odh-dashboard| role_odh_dashboard
    sa_odh_model_controller["odh-model-controller\nServiceAccount"]:::subject
    sa_odh_model_controller -->|odh-model-controller| role_proxy_role
    role_metrics_auth_role["metrics-auth-role"]:::role
    sa_controller_manager -->|odh-model-controller| role_metrics_auth_role
    role_odh_model_controller_role["odh-model-controller-role"]:::role
    sa_odh_model_controller -->|odh-model-controller| role_odh_model_controller_role
    role_controller_manager_role["controller-manager-role"]:::role
    sa_controller_manager -->|opendatahub-operator| role_controller_manager_role
    sa_controller_manager -->|trustyai-service-operator| role_proxy_role
    sa_system_authenticated["system:authenticated\nGroup"]:::subject
    role_lmeval_user_role["lmeval-user-role"]:::role
    sa_system_authenticated -->|trustyai-service-operator| role_lmeval_user_role
    sa_controller_manager -->|trustyai-service-operator| role_manager_role
    sa_controller_manager -->|trustyai-service-operator| role_system_auth_delegator
```

## Roles by Component

| Component | Roles | Widest Role | Resources | Scope |
|-----------|-------|-------------|-----------|-------|
| data-science-pipelines-operator | 4 | manager-role | 55 | **wide** |
| kserve | 2 | kserve-manager-role | 45 | **wide** |
| model-registry-operator | 6 | manager-role | 27 | medium |
| odh-dashboard | 1 | odh-dashboard | 40 | **wide** |
| odh-model-controller | 7 | odh-model-controller-role | 41 | **wide** |
| opendatahub-operator | 23 | modelregistry-viewer-role | 2 | narrow |
| trustyai-service-operator | 9 | manager-role | 44 | **wide** |

