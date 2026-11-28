# odh-dashboard

> **Architecture snapshot: 2026-11-28** (2026-11-28)


[Overview](index.md) · [Agent Context](agent-context.md) · [Extracted Data](details.md) · [Network](network.md) · [RBAC](rbac.md) · [Security](security.md) · [Cache](cache.md) · [Dataflow](dataflow.md)

**Repository:** opendatahub-io/odh-dashboard  
**Analyzer:** arch-analyzer 38b24267ad379636973f678b4ca92689c2c6fdf3  
**Extracted:** 2026-11-28T07:12:47Z

## Summary

| Metric | Count |
|--------|-------|
| CRDs | 0 |
| Deployments | 16 |
| Services | 15 |
| Secrets | 14 |
| Cluster Roles | 14 |
| Controller Watches | 0 |

## Component Architecture

CRDs, controllers, and owned Kubernetes resources.

```mermaid
graph LR
    %% Component architecture for odh-dashboard

    classDef crd fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef controller fill:#3498db,stroke:#2980b9,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef owned fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef dep fill:#f39c12,stroke:#e67e22,color:#fff

    subgraph controller["odh-dashboard Controller"]
        dep_1["agent-ops-ui"]
        class dep_1 controller
        dep_2["automl-ui"]
        class dep_2 controller
        dep_3["autorag-ui"]
        class dep_3 controller
        dep_4["dashboard-operator"]
        class dep_4 controller
        dep_5["data-connect-hub-ui"]
        class dep_5 controller
        dep_6["data-registry-ui"]
        class dep_6 controller
        dep_7["eval-hub-ui"]
        class dep_7 controller
        dep_8["gen-ai-ui"]
        class dep_8 controller
        dep_9["maas-portal"]
        class dep_9 controller
        dep_10["maas-ui"]
        class dep_10 controller
        dep_11["mlflow-ui"]
        class dep_11 controller
        dep_12["model-registry-ui"]
        class dep_12 controller
        dep_13["notebooks-ui"]
        class dep_13 controller
        dep_14["odh-dashboard"]
        class dep_14 controller
        dep_15["rhaii-dashboard"]
        class dep_15 controller
        dep_16["rhaii-dashboard"]
        class dep_16 controller
    end

    controller -.->|"depends on"| odh_17["mlflow-go"]
    class odh_17 dep
    controller -.->|"depends on"| odh_18["mlflow-go"]
    class odh_18 dep
    controller -.->|"depends on"| odh_19["odh-dashboard"]
    class odh_19 dep
    controller -.->|"depends on"| odh_20["odh-dashboard"]
    class odh_20 dep
    controller -.->|"depends on"| odh_21["odh-dashboard"]
    class odh_21 dep
    controller -.->|"depends on"| odh_22["odh-dashboard"]
    class odh_22 dep
    controller -.->|"depends on"| odh_23["odh-dashboard"]
    class odh_23 dep
    controller -.->|"depends on"| odh_24["odh-dashboard"]
    class odh_24 dep
    controller -.->|"depends on"| odh_25["odh-dashboard"]
    class odh_25 dep
    controller -.->|"depends on"| odh_26["odh-dashboard"]
    class odh_26 dep
    controller -.->|"depends on"| odh_27["odh-dashboard"]
    class odh_27 dep
    controller -.->|"depends on"| odh_28["odh-dashboard"]
    class odh_28 dep
    controller -.->|"depends on"| odh_29["odh-dashboard"]
    class odh_29 dep
    controller -.->|"depends on"| odh_30["odh-dashboard"]
    class odh_30 dep
    controller -.->|"depends on"| odh_31["odh-dashboard"]
    class odh_31 dep
    controller -.->|"depends on"| odh_32["odh-dashboard"]
    class odh_32 dep
    controller -.->|"depends on"| odh_33["odh-platform-utilities"]
    class odh_33 dep
    controller -.->|"depends on"| odh_34["odh-platform-utilities"]
    class odh_34 dep
    controller -.->|"depends on"| odh_35["ogx-k8s-operator"]
    class odh_35 dep
    controller -.->|"depends on"| odh_36["operator-chaos"]
    class odh_36 dep
```

### CRDs

No CRDs found in analyzed sources.

## Dependencies

### Internal Platform Dependencies

| Component | Interaction |
|-----------|-------------|
| mlflow-go | Go module dependency: github.com/opendatahub-io/mlflow-go |
| mlflow-go | Go module dependency: github.com/opendatahub-io/mlflow-go |
| odh-dashboard | Go module dependency: github.com/opendatahub-io/odh-dashboard/pkg/tls |
| odh-dashboard | Go module dependency: github.com/opendatahub-io/odh-dashboard/pkg/tls |
| odh-dashboard | Go module dependency: github.com/opendatahub-io/odh-dashboard/pkg/tls |
| odh-dashboard | Go module dependency: github.com/opendatahub-io/odh-dashboard/pkg/tls |
| odh-dashboard | Go module dependency: github.com/opendatahub-io/odh-dashboard/packages/autox-core/services |
| odh-dashboard | Go module dependency: github.com/opendatahub-io/odh-dashboard/pkg/tls |
| odh-dashboard | Go module dependency: github.com/opendatahub-io/odh-dashboard/packages/autox-core/services |
| odh-dashboard | Go module dependency: github.com/opendatahub-io/odh-dashboard/pkg/tls |
| odh-dashboard | Go module dependency: github.com/opendatahub-io/odh-dashboard/pkg/tls |
| odh-dashboard | Go module dependency: github.com/opendatahub-io/odh-dashboard/pkg/tls |
| odh-dashboard | Go module dependency: github.com/opendatahub-io/odh-dashboard/pkg/tls |
| odh-dashboard | Go module dependency: github.com/opendatahub-io/odh-dashboard/pkg/tls |
| odh-dashboard | Go module dependency: github.com/opendatahub-io/odh-dashboard/pkg/tls |
| odh-dashboard | Go module dependency: github.com/opendatahub-io/odh-dashboard/pkg/tls |
| odh-platform-utilities | Go module dependency: github.com/opendatahub-io/odh-platform-utilities/framework |
| odh-platform-utilities | Go module dependency: github.com/opendatahub-io/odh-platform-utilities |
| ogx-k8s-operator | Go module dependency: github.com/opendatahub-io/ogx-k8s-operator |
| operator-chaos | Go module dependency: github.com/opendatahub-io/operator-chaos |

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/go-logr/logr | v1.4.4 |
| google.golang.org/grpc | v1.84.0 |
| k8s.io/api | v0.34.3 |
| k8s.io/api | v0.34.3 |
| k8s.io/api | v0.36.2 |
| k8s.io/api | v0.34.3 |
| k8s.io/api | v0.34.3 |
| k8s.io/api | v0.34.3 |
| k8s.io/api | v0.36.4 |
| k8s.io/api | v0.37.0 |
| k8s.io/api | v0.36.3 |
| k8s.io/api | v0.34.3 |
| k8s.io/api | v0.34.3 |
| k8s.io/api | v0.34.3 |
| k8s.io/api | v0.34.3 |
| k8s.io/api | v0.36.4 |
| k8s.io/api | v0.34.3 |
| k8s.io/api | v0.34.3 |
| k8s.io/apiextensions-apiserver | v0.34.1 |
| k8s.io/apiextensions-apiserver | v0.34.3 |
| k8s.io/apiextensions-apiserver | v0.36.1 |
| k8s.io/apimachinery | v0.36.3 |
| k8s.io/apimachinery | v0.34.3 |
| k8s.io/apimachinery | v0.36.4 |
| k8s.io/apimachinery | v0.34.3 |
| k8s.io/apimachinery | v0.34.3 |
| k8s.io/apimachinery | v0.36.2 |
| k8s.io/apimachinery | v0.34.3 |
| k8s.io/apimachinery | v0.34.3 |
| k8s.io/apimachinery | v0.34.3 |
| k8s.io/apimachinery | v0.34.3 |
| k8s.io/apimachinery | v0.34.3 |
| k8s.io/apimachinery | v0.34.3 |
| k8s.io/apimachinery | v0.34.3 |
| k8s.io/apimachinery | v0.34.3 |
| k8s.io/apimachinery | v0.37.1 |
| k8s.io/apimachinery | v0.34.3 |
| k8s.io/apimachinery | v0.36.4 |
| k8s.io/apiserver | v0.36.4 |
| k8s.io/client-go | v0.34.3 |
| k8s.io/client-go | v0.34.3 |
| k8s.io/client-go | v0.36.3 |
| k8s.io/client-go | v0.34.3 |
| k8s.io/client-go | v0.34.3 |
| k8s.io/client-go | v0.34.3 |
| k8s.io/client-go | v0.36.4 |
| k8s.io/client-go | v0.34.3 |
| k8s.io/client-go | v0.34.3 |
| k8s.io/client-go | v0.34.3 |
| k8s.io/client-go | v0.36.4 |
| k8s.io/client-go | v0.37.0 |
| k8s.io/client-go | v0.34.3 |
| k8s.io/client-go | v0.34.3 |
| k8s.io/client-go | v0.34.3 |
| k8s.io/client-go | v0.36.1 |
| k8s.io/client-go | v0.34.3 |
| sigs.k8s.io/controller-runtime | v0.25.1 |
| sigs.k8s.io/controller-runtime | v0.22.3 |
| sigs.k8s.io/controller-runtime | v0.24.1 |
| sigs.k8s.io/controller-runtime | v0.22.3 |
| sigs.k8s.io/controller-runtime | v0.22.3 |
| sigs.k8s.io/controller-runtime | v0.24.1 |
| sigs.k8s.io/controller-runtime | v0.22.3 |
| sigs.k8s.io/controller-runtime | v0.24.1 |
| sigs.k8s.io/controller-runtime | v0.22.4 |
| sigs.k8s.io/controller-runtime | v0.24.1 |
| sigs.k8s.io/controller-runtime | v0.22.3 |
| sigs.k8s.io/controller-runtime | v0.22.3 |
| sigs.k8s.io/controller-runtime | v0.22.3 |

