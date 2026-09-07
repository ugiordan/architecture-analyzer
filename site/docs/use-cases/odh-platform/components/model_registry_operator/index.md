# model-registry-operator

> **Architecture snapshot: 2026-09-07** (2026-09-07)


**Repository:** opendatahub-io/model-registry-operator  
**Analyzer:** arch-analyzer dev  
**Extracted:** 2026-09-07T03:56:48Z

## Summary

| Metric | Count |
|--------|-------|
| CRDs | 1 |
| Deployments | 4 |
| Services | 5 |
| Secrets | 4 |
| Cluster Roles | 6 |
| Controller Watches | 30 |

## Component Architecture

CRDs, controllers, and owned Kubernetes resources.

```mermaid
graph LR
    %% Component architecture for model-registry-operator

    classDef crd fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef controller fill:#3498db,stroke:#2980b9,color:#fff
    classDef owned fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef dep fill:#f39c12,stroke:#e67e22,color:#fff

    subgraph controller["model-registry-operator Controller"]
        dep_1["aihub-controller-manager"]
        class dep_1 controller
        dep_2["catalog-controller-manager"]
        class dep_2 controller
        dep_3["model-registry-operator-controller-manager"]
        class dep_3 controller
        dep_4["template-value"]
        class dep_4 controller
    end

    crd_ModelRegistry{{"ModelRegistry\nmodelregistry.opendatahub.io/v1beta1"}}
    class crd_ModelRegistry crd
    crd_ModelRegistry -->|"For (reconciles)"| controller
    controller -->|"Owns"| owned_5["Catalog"]
    class owned_5 owned
    controller -->|"Owns"| owned_6["ClusterRole"]
    class owned_6 owned
    controller -->|"Owns"| owned_7["ClusterRoleBinding"]
    class owned_7 owned
    controller -->|"Owns"| owned_8["ConfigMap"]
    class owned_8 owned
    controller -->|"Owns"| owned_9["Deployment"]
    class owned_9 owned
    controller -->|"Owns"| owned_10["MutatingWebhookConfiguration"]
    class owned_10 owned
    controller -->|"Owns"| owned_11["NetworkPolicy"]
    class owned_11 owned
    controller -->|"Owns"| owned_12["Role"]
    class owned_12 owned
    controller -->|"Owns"| owned_13["RoleBinding"]
    class owned_13 owned
    controller -->|"Owns"| owned_14["Route"]
    class owned_14 owned
    controller -->|"Owns"| owned_15["Secret"]
    class owned_15 owned
    controller -->|"Owns"| owned_16["Service"]
    class owned_16 owned
    controller -->|"Owns"| owned_17["ServiceAccount"]
    class owned_17 owned
    controller -->|"Owns"| owned_18["ValidatingWebhookConfiguration"]
    class owned_18 owned
    controller -.->|"depends on"| odh_19["odh-platform-utilities"]
    class odh_19 dep
    controller -.->|"depends on"| odh_20["operator-chaos"]
    class odh_20 dep
```

### CRDs

| Group | Version | Kind | Scope | Fields | Validation Rules | Discovery | Source |
|-------|---------|------|-------|--------|------------------|-----------|--------|
| modelregistry.opendatahub.io | v1beta1 | ModelRegistry | Namespaced | 113 | 6 | YAML | [`config/crd/bases/modelregistry.opendatahub.io_modelregistries.yaml`](https://github.com/opendatahub-io/model-registry-operator/blob/655e3f2fe563582cb56b8a603c852325f70c51c5/config/crd/bases/modelregistry.opendatahub.io_modelregistries.yaml) |

## Dependencies

### Internal Platform Dependencies

| Component | Interaction |
|-----------|-------------|
| odh-platform-utilities | Go module dependency: github.com/opendatahub-io/odh-platform-utilities |
| operator-chaos | Go module dependency: github.com/opendatahub-io/operator-chaos |

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/go-logr/logr | v1.4.4 |
| k8s.io/api | v0.37.0 |
| k8s.io/apiextensions-apiserver | v0.37.0 |
| k8s.io/apimachinery | v0.37.0 |
| k8s.io/client-go | v0.37.0 |
| sigs.k8s.io/controller-runtime | v0.24.1 |

