# agents-operator

> **Architecture snapshot: 2026-09-07** (2026-09-07)


**Repository:** red-hat-data-services/agents-operator  
**Analyzer:** arch-analyzer dev  
**Extracted:** 2026-09-07T03:57:32Z

## Summary

| Metric | Count |
|--------|-------|
| CRDs | 0 |
| Deployments | 3 |
| Services | 2 |
| Secrets | 1 |
| Cluster Roles | 0 |
| Controller Watches | 20 |

## Component Architecture

CRDs, controllers, and owned Kubernetes resources.

```mermaid
graph LR
    %% Component architecture for agents-operator

    classDef crd fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef controller fill:#3498db,stroke:#2980b9,color:#fff
    classDef owned fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef dep fill:#f39c12,stroke:#e67e22,color:#fff

    subgraph controller["agents-operator Controller"]
        dep_1["bundle-service"]
        class dep_1 controller
        dep_2["controller-manager"]
        class dep_2 controller
        dep_3["controller-manager"]
        class dep_3 controller
    end

    controller -->|"Owns"| owned_4["Certificate"]
    class owned_4 owned
    controller -->|"Owns"| owned_5["ConfigMap"]
    class owned_5 owned
    controller -->|"Owns"| owned_6["NetworkPolicy"]
    class owned_6 owned
    controller -->|"Owns"| owned_7["Role"]
    class owned_7 owned
    controller -->|"Owns"| owned_8["RoleBinding"]
    class owned_8 owned
    watch_9["Certificate"] -->|"Watches"| controller
    class watch_9 external
    watch_10["Deployment"] -->|"Watches"| controller
    class watch_10 external
    watch_11["Secret"] -->|"Watches"| controller
    class watch_11 external
```

### CRDs

No CRDs found in analyzed sources.

## Dependencies

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/go-logr/logr | v1.4.3 |
| github.com/prometheus/client_golang | v1.23.2 |
| github.com/prometheus/client_model | v0.6.2 |
| google.golang.org/grpc | v1.81.1 |
| google.golang.org/grpc | v1.81.1 |
| k8s.io/api | v0.36.2 |
| k8s.io/apiextensions-apiserver | v0.36.2 |
| k8s.io/apimachinery | v0.36.2 |
| k8s.io/client-go | v0.36.2 |
| sigs.k8s.io/controller-runtime | v0.24.1 |

