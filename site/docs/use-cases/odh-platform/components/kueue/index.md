# kueue

> **Architecture snapshot: 2026-09-07** (2026-09-07)


**Repository:** opendatahub-io/kueue  
**Analyzer:** arch-analyzer dev  
**Extracted:** 2026-09-07T03:59:40Z

## Summary

| Metric | Count |
|--------|-------|
| CRDs | 2 |
| Deployments | 11 |
| Services | 2 |
| Secrets | 1 |
| Cluster Roles | 0 |
| Controller Watches | 22 |

## Component Architecture

CRDs, controllers, and owned Kubernetes resources.

```mermaid
graph LR
    %% Component architecture for kueue

    classDef crd fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef controller fill:#3498db,stroke:#2980b9,color:#fff
    classDef owned fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef dep fill:#f39c12,stroke:#e67e22,color:#fff

    subgraph controller["kueue Controller"]
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
        dep_6["controller-manager"]
        class dep_6 controller
        dep_7["controller-manager"]
        class dep_7 controller
        dep_8["controller-manager"]
        class dep_8 controller
        dep_9["controller-manager"]
        class dep_9 controller
        dep_10["controller-manager"]
        class dep_10 controller
        dep_11["controller-manager"]
        class dep_11 controller
    end

    crd_ClusterQueue{{"ClusterQueue\nvisibility.kueue.x-k8s.io/v1beta1"}}
    class crd_ClusterQueue crd
    crd_LocalQueue{{"LocalQueue\nvisibility.kueue.x-k8s.io/v1beta1"}}
    class crd_LocalQueue crd
    controller -->|"Owns"| owned_12["ProvisioningRequest"]
    class owned_12 owned
    controller -->|"Owns"| owned_13["Workload"]
    class owned_13 owned
    watch_14["AdmissionCheck"] -->|"Watches"| controller
    class watch_14 external
    watch_15["ClusterQueue"] -->|"Watches"| controller
    class watch_15 external
    watch_16["LimitRange"] -->|"Watches"| controller
    class watch_16 external
    watch_17["LocalQueue"] -->|"Watches"| controller
    class watch_17 external
    watch_18["Namespace"] -->|"Watches"| controller
    class watch_18 external
    watch_19["Pod"] -->|"Watches"| controller
    class watch_19 external
    watch_20["ProvisioningRequestConfig"] -->|"Watches"| controller
    class watch_20 external
    watch_21["ResourceFlavor"] -->|"Watches"| controller
    class watch_21 external
    watch_22["RuntimeClass"] -->|"Watches"| controller
    class watch_22 external
    watch_23["Workload"] -->|"Watches"| controller
    class watch_23 external
```

### CRDs

| Group | Version | Kind | Scope | Fields | Validation Rules | Discovery | Source |
|-------|---------|------|-------|--------|------------------|-----------|--------|
| visibility.kueue.x-k8s.io | v1beta1 | ClusterQueue | Namespaced | 18 | 0 | Go AST | [`/home/runner/work/_temp/arch-analyzer-repos/kueue/apis/visibility/v1beta1/types.go`](https://github.com/opendatahub-io/kueue/blob/97024bd289d2cc5c9369b40d9f3483ab1483143d//home/runner/work/_temp/arch-analyzer-repos/kueue/apis/visibility/v1beta1/types.go) |
| visibility.kueue.x-k8s.io | v1beta1 | LocalQueue | Namespaced | 18 | 0 | Go AST | [`/home/runner/work/_temp/arch-analyzer-repos/kueue/apis/visibility/v1beta1/types.go`](https://github.com/opendatahub-io/kueue/blob/97024bd289d2cc5c9369b40d9f3483ab1483143d//home/runner/work/_temp/arch-analyzer-repos/kueue/apis/visibility/v1beta1/types.go) |

## Dependencies

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/go-logr/logr | v1.4.2 |
| github.com/prometheus/client_golang | v1.21.1 |
| github.com/prometheus/client_model | v0.6.1 |
| k8s.io/api | v0.32.3 |
| k8s.io/api | v0.32.3 |
| k8s.io/apimachinery | v0.32.3 |
| k8s.io/apimachinery | v0.32.3 |
| k8s.io/apiserver | v0.32.3 |
| k8s.io/client-go | v0.32.3 |
| k8s.io/client-go | v0.32.3 |
| sigs.k8s.io/controller-runtime | v0.19.4 |

