# praxis-extproc

> **Architecture snapshot: 2026-11-21** (2026-11-21)


**Repository:** opendatahub-io/praxis-extproc  
**Analyzer:** arch-analyzer f62b7632f89631498789b54d78e8c552c27100c7  
**Extracted:** 2026-11-21T06:38:32Z

See [Agent Context](agent-context.md) for repository structure and code relationships, and [Extracted Data](details.md) for dependency and configuration details.

## Summary

| Metric | Count |
|--------|-------|
| CRDs | 0 |
| Deployments | 1 |
| Services | 1 |
| Secrets | 1 |
| Cluster Roles | 0 |
| Controller Watches | 0 |

## Component Architecture

CRDs, controllers, and owned Kubernetes resources.

```mermaid
graph LR
    %% Component architecture for praxis-extproc

    classDef crd fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef controller fill:#3498db,stroke:#2980b9,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef owned fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef dep fill:#f39c12,stroke:#e67e22,color:#fff

    subgraph controller["praxis-extproc Controller"]
        dep_1["payload-processing"]
        class dep_1 controller
    end
```

### CRDs

No CRDs found in analyzed sources.

## Dependencies

