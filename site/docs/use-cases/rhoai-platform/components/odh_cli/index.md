# odh-cli

> **Architecture snapshot: 2026-09-07** (2026-09-07)


**Repository:** opendatahub-io/odh-cli  
**Analyzer:** arch-analyzer dev  
**Extracted:** 2026-09-07T04:02:44Z

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
    classDef owned fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef dep fill:#f39c12,stroke:#e67e22,color:#fff

    subgraph controller["odh-cli Controller"]
        ctrl_1["Controller"]
        class ctrl_1 controller
    end

    controller -.->|"depends on"| odh_2["opendatahub-operator"]
    class odh_2 dep
    controller -.->|"depends on"| odh_3["opendatahub-operator"]
    class odh_3 dep
    controller -.->|"depends on"| odh_4["opendatahub-operator"]
    class odh_4 dep
```

### CRDs

No CRDs found in analyzed sources.

## Dependencies

### Internal Platform Dependencies

| Component | Interaction |
|-----------|-------------|
| opendatahub-operator | Go module dependency: github.com/opendatahub-io/opendatahub-operator/pkg/clusterhealth |
| opendatahub-operator | Go module dependency: github.com/opendatahub-io/opendatahub-operator/pkg/failureclassifier |
| opendatahub-operator | Go module dependency: github.com/opendatahub-io/opendatahub-operator/pkg/mcptools |

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/operator-framework/api | v0.39.0 |
| github.com/operator-framework/operator-lifecycle-manager | v0.40.0 |
| k8s.io/api | v0.35.2 |
| k8s.io/apiextensions-apiserver | v0.35.2 |
| k8s.io/apimachinery | v0.35.2 |
| k8s.io/client-go | v0.35.2 |
| sigs.k8s.io/controller-runtime | v0.23.3 |

