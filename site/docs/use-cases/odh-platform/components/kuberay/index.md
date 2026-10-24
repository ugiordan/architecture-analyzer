# kuberay

> **Architecture snapshot: 2026-10-24** (2026-10-24)


**Repository:** opendatahub-io/kuberay  
**Analyzer:** arch-analyzer dev  
**Extracted:** 2026-10-24T04:10:15Z

## Summary

| Metric | Count |
|--------|-------|
| CRDs | 0 |
| Deployments | 1 |
| Services | 1 |
| Secrets | 0 |
| Cluster Roles | 0 |
| Controller Watches | 11 |

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
    end

    controller -->|"Owns"| owned_2["Ingress"]
    class owned_2 owned
    controller -->|"Owns"| owned_3["Job"]
    class owned_3 owned
    controller -->|"Owns"| owned_4["Pod"]
    class owned_4 owned
    controller -->|"Owns"| owned_5["RayCluster"]
    class owned_5 owned
    controller -->|"Owns"| owned_6["Service"]
    class owned_6 owned
```

### CRDs

No CRDs found in analyzed sources.

## Dependencies

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/go-logr/logr | v1.2.4 |
| github.com/go-logr/zapr | v1.2.4 |
| github.com/prometheus/client_golang | v1.16.0 |
| github.com/prometheus/client_golang | v1.16.0 |
| google.golang.org/grpc | v1.59.0 |
| google.golang.org/grpc | v1.54.0 |
| google.golang.org/grpc | v1.59.0 |
| k8s.io/api | v0.28.4 |
| k8s.io/api | v0.28.4 |
| k8s.io/apiextensions-apiserver | v0.28.4 |
| k8s.io/apimachinery | v0.28.4 |
| k8s.io/apimachinery | v0.28.4 |
| k8s.io/apiserver | v0.28.4 |
| k8s.io/client-go | v0.28.4 |
| k8s.io/client-go | v0.28.4 |
| sigs.k8s.io/controller-runtime | v0.16.3 |
| sigs.k8s.io/controller-runtime | v0.16.3 |

