# feast

> **Architecture snapshot: 2026-10-31** (2026-10-31)


**Repository:** opendatahub-io/feast  
**Analyzer:** arch-analyzer dev  
**Extracted:** 2026-10-31T04:59:37Z

## Summary

| Metric | Count |
|--------|-------|
| CRDs | 0 |
| Deployments | 3 |
| Services | 1 |
| Secrets | 1 |
| Cluster Roles | 0 |
| Controller Watches | 14 |

## Component Architecture

CRDs, controllers, and owned Kubernetes resources.

```mermaid
graph LR
    %% Component architecture for feast

    classDef crd fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef controller fill:#3498db,stroke:#2980b9,color:#fff
    classDef owned fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef dep fill:#f39c12,stroke:#e67e22,color:#fff

    subgraph controller["feast Controller"]
        dep_1["controller-manager"]
        class dep_1 controller
        dep_2["controller-manager"]
        class dep_2 controller
        dep_3["controller-manager"]
        class dep_3 controller
    end

    controller -->|"Owns"| owned_4["ConfigMap"]
    class owned_4 owned
    controller -->|"Owns"| owned_5["CronJob"]
    class owned_5 owned
    controller -->|"Owns"| owned_6["Deployment"]
    class owned_6 owned
    controller -->|"Owns"| owned_7["HorizontalPodAutoscaler"]
    class owned_7 owned
    controller -->|"Owns"| owned_8["NetworkPolicy"]
    class owned_8 owned
    controller -->|"Owns"| owned_9["PersistentVolumeClaim"]
    class owned_9 owned
    controller -->|"Owns"| owned_10["PodDisruptionBudget"]
    class owned_10 owned
    controller -->|"Owns"| owned_11["Role"]
    class owned_11 owned
    controller -->|"Owns"| owned_12["RoleBinding"]
    class owned_12 owned
    controller -->|"Owns"| owned_13["Route"]
    class owned_13 owned
    controller -->|"Owns"| owned_14["Service"]
    class owned_14 owned
    controller -->|"Owns"| owned_15["ServiceAccount"]
    class owned_15 owned
    watch_16["FeatureStore"] -->|"Watches"| controller
    class watch_16 external
```

### CRDs

No CRDs found in analyzed sources.

## Dependencies

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/prometheus-operator/prometheus-operator/pkg/client | v0.83.0 |
| github.com/prometheus/client_golang | v1.23.2 |
| github.com/prometheus/client_golang | v1.24.1 |
| github.com/prometheus/client_model | v0.6.2 |
| google.golang.org/grpc | v1.84.0 |
| k8s.io/api | v0.35.2 |
| k8s.io/apimachinery | v0.35.2 |
| k8s.io/client-go | v0.35.2 |
| sigs.k8s.io/controller-runtime | v0.23.3 |

