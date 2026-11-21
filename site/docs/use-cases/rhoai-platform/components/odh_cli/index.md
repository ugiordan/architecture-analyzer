# odh-cli

> **Architecture snapshot: 2026-11-21** (2026-11-21)


**Repository:** red-hat-data-services/odh-cli  
**Analyzer:** arch-analyzer f62b7632f89631498789b54d78e8c552c27100c7  
**Extracted:** 2026-11-21T06:39:20Z

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
    %% Component architecture for odh-cli

    classDef crd fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef controller fill:#3498db,stroke:#2980b9,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef owned fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef dep fill:#f39c12,stroke:#e67e22,color:#fff

    component_odh_cli["odh-cli"]:::component

    component_odh_cli -.->|"depends on"| odh_1["odh-platform-utilities"]
    class odh_1 dep
    component_odh_cli -.->|"depends on"| odh_2["opendatahub-operator"]
    class odh_2 dep
    component_odh_cli -.->|"depends on"| odh_3["opendatahub-operator"]
    class odh_3 dep
    component_odh_cli -.->|"depends on"| odh_4["opendatahub-operator"]
    class odh_4 dep
```

### CRDs

No CRDs found in analyzed sources.

## Dependencies

### Internal Platform Dependencies

| Component | Interaction |
|-----------|-------------|
| odh-platform-utilities | Go module dependency: github.com/opendatahub-io/odh-platform-utilities |
| opendatahub-operator | Go module dependency: github.com/opendatahub-io/opendatahub-operator/pkg/clusterhealth |
| opendatahub-operator | Go module dependency: github.com/opendatahub-io/opendatahub-operator/pkg/failureclassifier |
| opendatahub-operator | Go module dependency: github.com/opendatahub-io/opendatahub-operator/pkg/mcptools |

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/operator-framework/api | v0.39.0 |
| github.com/operator-framework/operator-lifecycle-manager | v0.40.0 |
| k8s.io/api | v0.35.3 |
| k8s.io/apiextensions-apiserver | v0.35.3 |
| k8s.io/apimachinery | v0.35.4 |
| k8s.io/client-go | v0.35.3 |
| sigs.k8s.io/controller-runtime | v0.23.3 |

