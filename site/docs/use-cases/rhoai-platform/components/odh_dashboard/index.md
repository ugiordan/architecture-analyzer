# odh-dashboard

> **Architecture snapshot: 2026-09-26** (2026-09-26)


**Repository:** red-hat-data-services/odh-dashboard  
**Analyzer:** arch-analyzer dev  
**Extracted:** 2026-09-26T04:28:03Z

## Summary

| Metric | Count |
|--------|-------|
| CRDs | 0 |
| Deployments | 24 |
| Services | 18 |
| Secrets | 18 |
| Cluster Roles | 13 |
| Controller Watches | 17 |

## Component Architecture

CRDs, controllers, and owned Kubernetes resources.

```mermaid
graph LR
    %% Component architecture for odh-dashboard

    classDef crd fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef controller fill:#3498db,stroke:#2980b9,color:#fff
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
        dep_9["maas-consumer-portal"]
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
        dep_16["workspaces-backend"]
        class dep_16 controller
        dep_17["workspaces-backend"]
        class dep_17 controller
        dep_18["workspaces-controller"]
        class dep_18 controller
        dep_19["workspaces-controller"]
        class dep_19 controller
        dep_20["workspaces-controller"]
        class dep_20 controller
        dep_21["workspaces-controller"]
        class dep_21 controller
        dep_22["workspaces-controller"]
        class dep_22 controller
        dep_23["workspaces-controller"]
        class dep_23 controller
        dep_24["workspaces-frontend"]
        class dep_24 controller
    end

    controller -->|"Owns"| owned_25["ClusterRole"]
    class owned_25 owned
    controller -->|"Owns"| owned_26["ClusterRoleBinding"]
    class owned_26 owned
    controller -->|"Owns"| owned_27["ConfigMap"]
    class owned_27 owned
    controller -->|"Owns"| owned_28["Deployment"]
    class owned_28 owned
    controller -->|"Owns"| owned_29["NetworkPolicy"]
    class owned_29 owned
    controller -->|"Owns"| owned_30["PodDisruptionBudget"]
    class owned_30 owned
    controller -->|"Owns"| owned_31["RoleBinding"]
    class owned_31 owned
    controller -->|"Owns"| owned_32["Secret"]
    class owned_32 owned
    controller -->|"Owns"| owned_33["Service"]
    class owned_33 owned
    controller -->|"Owns"| owned_34["ServiceAccount"]
    class owned_34 owned
    controller -->|"Owns"| owned_35["StatefulSet"]
    class owned_35 owned
    controller -->|"Owns"| owned_36["VirtualService"]
    class owned_36 owned
    controller -.->|"depends on"| odh_37["mlflow-go"]
    class odh_37 dep
    controller -.->|"depends on"| odh_38["mlflow-go"]
    class odh_38 dep
    controller -.->|"depends on"| odh_39["odh-dashboard"]
    class odh_39 dep
    controller -.->|"depends on"| odh_40["odh-dashboard"]
    class odh_40 dep
    controller -.->|"depends on"| odh_41["odh-dashboard"]
    class odh_41 dep
    controller -.->|"depends on"| odh_42["odh-dashboard"]
    class odh_42 dep
    controller -.->|"depends on"| odh_43["odh-dashboard"]
    class odh_43 dep
    controller -.->|"depends on"| odh_44["odh-dashboard"]
    class odh_44 dep
    controller -.->|"depends on"| odh_45["odh-dashboard"]
    class odh_45 dep
    controller -.->|"depends on"| odh_46["odh-dashboard"]
    class odh_46 dep
    controller -.->|"depends on"| odh_47["odh-dashboard"]
    class odh_47 dep
    controller -.->|"depends on"| odh_48["odh-dashboard"]
    class odh_48 dep
    controller -.->|"depends on"| odh_49["odh-dashboard"]
    class odh_49 dep
    controller -.->|"depends on"| odh_50["odh-dashboard"]
    class odh_50 dep
    controller -.->|"depends on"| odh_51["odh-platform-utilities"]
    class odh_51 dep
    controller -.->|"depends on"| odh_52["odh-platform-utilities"]
    class odh_52 dep
    controller -.->|"depends on"| odh_53["ogx-k8s-operator"]
    class odh_53 dep
    controller -.->|"depends on"| odh_54["operator-chaos"]
    class odh_54 dep
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
| odh-dashboard | Go module dependency: github.com/opendatahub-io/odh-dashboard/packages/autox-core/services |
| odh-dashboard | Go module dependency: github.com/opendatahub-io/odh-dashboard/pkg/tls |
| odh-dashboard | Go module dependency: github.com/opendatahub-io/odh-dashboard/packages/autox-core/services |
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
| github.com/go-logr/logr | v1.4.3 |
| k8s.io/api | v0.34.3 |
| k8s.io/api | v0.34.3 |
| k8s.io/api | v0.34.3 |
| k8s.io/api | v0.34.3 |
| k8s.io/api | v0.34.1 |
| k8s.io/api | v0.37.0 |
| k8s.io/api | v0.34.3 |
| k8s.io/api | v0.34.3 |
| k8s.io/api | v0.36.2 |
| k8s.io/api | v0.34.1 |
| k8s.io/api | v0.36.3 |
| k8s.io/api | v0.34.3 |
| k8s.io/api | v0.34.3 |
| k8s.io/api | v0.34.3 |
| k8s.io/api | v0.34.1 |
| k8s.io/apiextensions-apiserver | v0.34.1 |
| k8s.io/apiextensions-apiserver | v0.34.3 |
| k8s.io/apiextensions-apiserver | v0.36.1 |
| k8s.io/apimachinery | v0.34.3 |
| k8s.io/apimachinery | v0.36.2 |
| k8s.io/apimachinery | v0.34.3 |
| k8s.io/apimachinery | v0.34.3 |
| k8s.io/apimachinery | v0.34.3 |
| k8s.io/apimachinery | v0.34.1 |
| k8s.io/apimachinery | v0.34.3 |
| k8s.io/apimachinery | v0.34.3 |
| k8s.io/apimachinery | v0.34.3 |
| k8s.io/apimachinery | v0.36.3 |
| k8s.io/apimachinery | v0.37.0 |
| k8s.io/apimachinery | v0.34.3 |
| k8s.io/apimachinery | v0.34.1 |
| k8s.io/apimachinery | v0.34.3 |
| k8s.io/apimachinery | v0.34.3 |
| k8s.io/apimachinery | v0.34.1 |
| k8s.io/apiserver | v0.34.1 |
| k8s.io/client-go | v0.34.3 |
| k8s.io/client-go | v0.34.3 |
| k8s.io/client-go | v0.34.3 |
| k8s.io/client-go | v0.34.3 |
| k8s.io/client-go | v0.34.3 |
| k8s.io/client-go | v0.34.3 |
| k8s.io/client-go | v0.37.0 |
| k8s.io/client-go | v0.36.1 |
| k8s.io/client-go | v0.34.1 |
| k8s.io/client-go | v0.34.3 |
| k8s.io/client-go | v0.34.3 |
| k8s.io/client-go | v0.34.3 |
| k8s.io/client-go | v0.34.1 |
| k8s.io/client-go | v0.34.1 |
| k8s.io/client-go | v0.36.3 |
| k8s.io/client-go | v0.34.3 |
| sigs.k8s.io/controller-runtime | v0.22.3 |
| sigs.k8s.io/controller-runtime | v0.22.4 |
| sigs.k8s.io/controller-runtime | v0.22.3 |
| sigs.k8s.io/controller-runtime | v0.22.1 |
| sigs.k8s.io/controller-runtime | v0.22.3 |
| sigs.k8s.io/controller-runtime | v0.25.0 |
| sigs.k8s.io/controller-runtime | v0.22.3 |
| sigs.k8s.io/controller-runtime | v0.24.1 |
| sigs.k8s.io/controller-runtime | v0.24.1 |
| sigs.k8s.io/controller-runtime | v0.22.3 |
| sigs.k8s.io/controller-runtime | v0.22.3 |
| sigs.k8s.io/controller-runtime | v0.22.1 |

