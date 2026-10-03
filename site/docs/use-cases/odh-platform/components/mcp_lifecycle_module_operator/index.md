# mcp-lifecycle-module-operator

> **Architecture snapshot: 2026-10-03** (2026-10-03)


**Repository:** opendatahub-io/mcp-lifecycle-module-operator  
**Analyzer:** arch-analyzer dev  
**Extracted:** 2026-10-03T04:45:00Z

## Summary

| Metric | Count |
|--------|-------|
| CRDs | 1 |
| Deployments | 5 |
| Services | 1 |
| Secrets | 1 |
| Cluster Roles | 1 |
| Controller Watches | 11 |

## Component Architecture

CRDs, controllers, and owned Kubernetes resources.

```mermaid
graph LR
    %% Component architecture for mcp-lifecycle-module-operator

    classDef crd fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef controller fill:#3498db,stroke:#2980b9,color:#fff
    classDef owned fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef dep fill:#f39c12,stroke:#e67e22,color:#fff

    subgraph controller["mcp-lifecycle-module-operator Controller"]
        dep_1["controller-manager"]
        class dep_1 controller
        dep_2["controller-manager"]
        class dep_2 controller
        dep_3["controller-manager"]
        class dep_3 controller
        dep_4["controller-manager"]
        class dep_4 controller
        dep_5["controller-manager"]
        class dep_5 controller
    end

    crd_MCPLifecycleOperator{{"MCPLifecycleOperator\ncomponents.platform.opendatahub.io/v1alpha1"}}
    class crd_MCPLifecycleOperator crd
    crd_MCPLifecycleOperator -->|"For (reconciles)"| controller
    watch_6["APIServer"] -->|"Watches"| controller
    class watch_6 external
    watch_7["ClusterRole"] -->|"Watches"| controller
    class watch_7 external
    watch_8["ClusterRoleBinding"] -->|"Watches"| controller
    class watch_8 external
    watch_9["ConfigMap"] -->|"Watches"| controller
    class watch_9 external
    watch_10["CustomResourceDefinition"] -->|"Watches"| controller
    class watch_10 external
    watch_11["Deployment"] -->|"Watches"| controller
    class watch_11 external
    watch_12["Role"] -->|"Watches"| controller
    class watch_12 external
    watch_13["RoleBinding"] -->|"Watches"| controller
    class watch_13 external
    watch_14["Service"] -->|"Watches"| controller
    class watch_14 external
    watch_15["ServiceAccount"] -->|"Watches"| controller
    class watch_15 external
    controller -.->|"depends on"| odh_16["odh-platform-utilities"]
    class odh_16 dep
```

### CRDs

| Group | Version | Kind | Scope | Fields | Validation Rules | Discovery | Source |
|-------|---------|------|-------|--------|------------------|-----------|--------|
| components.platform.opendatahub.io | v1alpha1 | MCPLifecycleOperator | Cluster | 24 | 1 | YAML | [`config/crd/bases/components.platform.opendatahub.io_mcplifecycleoperators.yaml`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/ed1187c1c5382db14bac17ee53f5a8a1dd5dbd7d/config/crd/bases/components.platform.opendatahub.io_mcplifecycleoperators.yaml) |

## Dependencies

### Internal Platform Dependencies

| Component | Interaction |
|-----------|-------------|
| odh-platform-utilities | Go module dependency: github.com/opendatahub-io/odh-platform-utilities |

### Key External Dependencies

| Module | Version |
|--------|---------|
| k8s.io/api | v0.36.2 |
| k8s.io/apiextensions-apiserver | v0.36.2 |
| k8s.io/apimachinery | v0.36.2 |
| k8s.io/client-go | v0.36.2 |
| sigs.k8s.io/controller-runtime | v0.24.1 |

