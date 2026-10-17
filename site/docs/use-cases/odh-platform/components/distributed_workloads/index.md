# distributed-workloads

> **Architecture snapshot: 2026-10-17** (2026-10-17)


**Repository:** opendatahub-io/distributed-workloads  
**Analyzer:** arch-analyzer dev  
**Extracted:** 2026-10-17T03:54:05Z

## Summary

| Metric | Count |
|--------|-------|
| CRDs | 0 |
| Deployments | 0 |
| Services | 0 |
| Secrets | 0 |
| Cluster Roles | 0 |
| Controller Watches | 0 |

## Component Architecture

CRDs, controllers, and owned Kubernetes resources.

```mermaid
graph LR
    %% Component architecture for distributed-workloads

    classDef crd fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef controller fill:#3498db,stroke:#2980b9,color:#fff
    classDef owned fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef dep fill:#f39c12,stroke:#e67e22,color:#fff

    subgraph controller["distributed-workloads Controller"]
        ctrl_1["Controller"]
        class ctrl_1 controller
    end

    controller -.->|"depends on"| odh_2["trainer"]
    class odh_2 dep
```

### CRDs

No CRDs found in analyzed sources.

## Dependencies

### Internal Platform Dependencies

| Component | Interaction |
|-----------|-------------|
| trainer | Go module dependency: github.com/opendatahub-io/trainer/v2 |

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/operator-framework/api | v0.36.0 |
| github.com/operator-framework/operator-lifecycle-manager | v0.38.0 |
| github.com/prometheus/client_golang | v1.25.0 |
| github.com/prometheus/common | v0.72.0 |
| k8s.io/api | v0.37.1 |
| k8s.io/apimachinery | v0.37.1 |
| k8s.io/client-go | v0.37.1 |

