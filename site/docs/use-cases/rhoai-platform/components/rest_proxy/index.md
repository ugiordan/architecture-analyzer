# rest-proxy

> **Architecture snapshot: 2026-09-07** (2026-09-07)


**Repository:** kserve/rest-proxy  
**Analyzer:** arch-analyzer dev  
**Extracted:** 2026-09-07T04:03:13Z

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
    %% Component architecture for rest-proxy

    classDef crd fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef controller fill:#3498db,stroke:#2980b9,color:#fff
    classDef owned fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef dep fill:#f39c12,stroke:#e67e22,color:#fff

    subgraph controller["rest-proxy Controller"]
        ctrl_1["Controller"]
        class ctrl_1 controller
    end
```

### CRDs

No CRDs found in analyzed sources.

## Dependencies

### Key External Dependencies

| Module | Version |
|--------|---------|
| google.golang.org/grpc | v1.56.3 |
| sigs.k8s.io/controller-runtime | v0.14.1 |

