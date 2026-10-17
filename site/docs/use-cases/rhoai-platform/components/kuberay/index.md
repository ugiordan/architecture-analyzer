# kuberay

> **Architecture snapshot: 2026-10-17** (2026-10-17)


**Repository:** ray-project/kuberay  
**Analyzer:** arch-analyzer dev  
**Extracted:** 2026-10-17T03:54:53Z

## Summary

| Metric | Count |
|--------|-------|
| CRDs | 0 |
| Deployments | 4 |
| Services | 2 |
| Secrets | 1 |
| Cluster Roles | 0 |
| Controller Watches | 23 |

## Component Architecture

CRDs, controllers, and owned Kubernetes resources.

```mermaid
graph LR
    %% Component architecture for kuberay

    classDef crd fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef controller fill:#3498db,stroke:#2980b9,color:#fff
    classDef owned fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef dep fill:#f39c12,stroke:#e67e22,color:#fff

    subgraph controller["kuberay Controller"]
        dep_1["kuberay-operator"]
        class dep_1 controller
        dep_2["kuberay-operator"]
        class dep_2 controller
        dep_3["kuberay-operator"]
        class dep_3 controller
        dep_4["kuberay-operator"]
        class dep_4 controller
    end

    controller -->|"Owns"| owned_5["Certificate"]
    class owned_5 owned
    controller -->|"Owns"| owned_6["Issuer"]
    class owned_6 owned
    controller -->|"Owns"| owned_7["Job"]
    class owned_7 owned
    controller -->|"Owns"| owned_8["NetworkPolicy"]
    class owned_8 owned
    controller -->|"Owns"| owned_9["PersistentVolumeClaim"]
    class owned_9 owned
    controller -->|"Owns"| owned_10["Pod"]
    class owned_10 owned
    controller -->|"Owns"| owned_11["RayCluster"]
    class owned_11 owned
    controller -->|"Owns"| owned_12["Route"]
    class owned_12 owned
    controller -->|"Owns"| owned_13["Secret"]
    class owned_13 owned
    controller -->|"Owns"| owned_14["Service"]
    class owned_14 owned
    controller -->|"Owns"| owned_15["ServiceAccount"]
    class owned_15 owned
```

### CRDs

No CRDs found in analyzed sources.

## Dependencies

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/go-logr/logr | v1.4.3 |
| github.com/go-logr/logr | v1.4.3 |
| github.com/go-logr/zapr | v1.3.0 |
| github.com/go-logr/zerologr | v1.2.3 |
| github.com/prometheus/client_golang | v1.23.2 |
| github.com/prometheus/client_golang | v1.23.2 |
| google.golang.org/grpc | v1.83.1 |
| k8s.io/api | v0.36.0 |
| k8s.io/api | v0.35.0 |
| k8s.io/api | v0.35.0 |
| k8s.io/api | v0.36.0 |
| k8s.io/apimachinery | v0.36.0 |
| k8s.io/apimachinery | v0.36.0 |
| k8s.io/apimachinery | v0.35.0 |
| k8s.io/apimachinery | v0.35.0 |
| k8s.io/apiserver | v0.36.0 |
| k8s.io/client-go | v0.35.0 |
| k8s.io/client-go | v0.35.0 |
| k8s.io/client-go | v0.36.0 |
| k8s.io/client-go | v0.36.0 |
| sigs.k8s.io/controller-runtime | v0.24.1 |
| sigs.k8s.io/controller-runtime | v0.22.4 |
| sigs.k8s.io/controller-runtime | v0.24.1 |

