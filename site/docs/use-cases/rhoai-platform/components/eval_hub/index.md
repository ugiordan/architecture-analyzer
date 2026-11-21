# eval-hub

> **Architecture snapshot: 2026-11-21** (2026-11-21)


**Repository:** red-hat-data-services/eval-hub  
**Analyzer:** arch-analyzer f62b7632f89631498789b54d78e8c552c27100c7  
**Extracted:** 2026-11-21T06:38:44Z

See [Agent Context](agent-context.md) for repository structure and code relationships, and [Extracted Data](details.md) for dependency and configuration details.

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
    %% Component architecture for eval-hub

    classDef crd fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef controller fill:#3498db,stroke:#2980b9,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef owned fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef dep fill:#f39c12,stroke:#e67e22,color:#fff

    component_eval_hub["eval-hub"]:::component

    component_eval_hub -.->|"depends on"| odh_1["mlflow-go"]
    class odh_1 dep
```

### CRDs

No CRDs found in analyzed sources.

## Dependencies

### Internal Platform Dependencies

| Component | Interaction |
|-----------|-------------|
| mlflow-go | Go module dependency: github.com/opendatahub-io/mlflow-go |

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/go-logr/logr | v1.4.4 |
| github.com/prometheus/client_golang | v1.25.0 |
| github.com/prometheus/client_model | v0.6.3 |
| google.golang.org/grpc | v1.85.0 |
| k8s.io/api | v0.37.2 |
| k8s.io/apimachinery | v0.37.2 |
| k8s.io/client-go | v0.37.2 |

