# kubeflow

> **Architecture snapshot: 2026-09-15** (2026-09-15)


**Repository:** opendatahub-io/kubeflow  
**Analyzer:** arch-analyzer dev  
**Extracted:** 2026-09-15T04:21:39Z

## Summary

| Metric | Count |
|--------|-------|
| CRDs | 0 |
| Deployments | 9 |
| Services | 3 |
| Secrets | 4 |
| Cluster Roles | 0 |
| Controller Watches | 14 |

## Component Architecture

CRDs, controllers, and owned Kubernetes resources.

```mermaid
graph LR
    %% Component architecture for kubeflow

    classDef crd fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef controller fill:#3498db,stroke:#2980b9,color:#fff
    classDef owned fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef dep fill:#f39c12,stroke:#e67e22,color:#fff

    subgraph controller["kubeflow Controller"]
        dep_1["controller-manager"]
        class dep_1 controller
        dep_2["controller-manager"]
        class dep_2 controller
        dep_3["controller-manager"]
        class dep_3 controller
        dep_4["controller-manager"]
        class dep_4 controller
        dep_5["deployment"]
        class dep_5 controller
        dep_6["deployment"]
        class dep_6 controller
        dep_7["deployment"]
        class dep_7 controller
        dep_8["manager"]
        class dep_8 controller
        dep_9["manager"]
        class dep_9 controller
    end

    controller -->|"Owns"| owned_10["ConfigMap"]
    class owned_10 owned
    controller -->|"Owns"| owned_11["NetworkPolicy"]
    class owned_11 owned
    controller -->|"Owns"| owned_12["RoleBinding"]
    class owned_12 owned
    controller -->|"Owns"| owned_13["Secret"]
    class owned_13 owned
    controller -->|"Owns"| owned_14["Service"]
    class owned_14 owned
    controller -->|"Owns"| owned_15["ServiceAccount"]
    class owned_15 owned
    controller -->|"Owns"| owned_16["StatefulSet"]
    class owned_16 owned
    watch_17["ConfigMap"] -->|"Watches"| controller
    class watch_17 external
    watch_18["HTTPRoute"] -->|"Watches"| controller
    class watch_18 external
    watch_19["ReferenceGrant"] -->|"Watches"| controller
    class watch_19 external
    controller -.->|"depends on"| odh_20["data-science-pipelines-operator"]
    class odh_20 dep
    controller -.->|"depends on"| odh_21["operator-chaos"]
    class odh_21 dep
    controller -.->|"depends on"| odh_22["operator-chaos"]
    class odh_22 dep
```

### CRDs

No CRDs found in analyzed sources.

## Dependencies

### Internal Platform Dependencies

| Component | Interaction |
|-----------|-------------|
| data-science-pipelines-operator | Go module dependency: github.com/opendatahub-io/data-science-pipelines-operator |
| operator-chaos | Go module dependency: github.com/opendatahub-io/operator-chaos |
| operator-chaos | Go module dependency: github.com/opendatahub-io/operator-chaos |

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/go-logr/logr | v1.4.3 |
| github.com/go-logr/logr | v1.4.3 |
| github.com/prometheus/client_golang | v1.23.2 |
| k8s.io/api | v0.35.2 |
| k8s.io/api | v0.35.2 |
| k8s.io/apiextensions-apiserver | v0.35.1 |
| k8s.io/apimachinery | v0.35.2 |
| k8s.io/apimachinery | v0.35.2 |
| k8s.io/client-go | v0.35.2 |
| k8s.io/client-go | v0.35.2 |
| sigs.k8s.io/controller-runtime | v0.23.3 |
| sigs.k8s.io/controller-runtime | v0.23.3 |

