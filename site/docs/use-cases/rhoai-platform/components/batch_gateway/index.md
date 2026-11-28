# batch-gateway

> **Architecture snapshot: 2026-11-28** (2026-11-28)


[Overview](index.md) · [Agent Context](agent-context.md) · [Extracted Data](details.md) · [Network](network.md) · [RBAC](rbac.md) · [Security](security.md) · [Dataflow](dataflow.md)

**Repository:** red-hat-data-services/batch-gateway  
**Analyzer:** arch-analyzer 38b24267ad379636973f678b4ca92689c2c6fdf3  
**Extracted:** 2026-11-28T07:18:23Z

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
    %% Component architecture for batch-gateway

    classDef crd fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef controller fill:#3498db,stroke:#2980b9,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef owned fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef dep fill:#f39c12,stroke:#e67e22,color:#fff

    component_batch_gateway["batch-gateway"]:::component
```

### CRDs

No CRDs found in analyzed sources.

## Dependencies

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/go-logr/logr | v1.4.4 |
| github.com/prometheus/client_golang | v1.25.0 |
| github.com/prometheus/client_model | v0.6.3 |
| k8s.io/api | v0.37.2 |
| k8s.io/apimachinery | v0.37.2 |
| k8s.io/client-go | v0.37.2 |

