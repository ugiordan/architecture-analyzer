# spark-operator

> **Architecture snapshot: 2026-09-07** (2026-09-07)


**Repository:** kubeflow/spark-operator  
**Analyzer:** arch-analyzer dev  
**Extracted:** 2026-09-07T03:57:01Z

## Summary

| Metric | Count |
|--------|-------|
| CRDs | 3 |
| Deployments | 4 |
| Services | 1 |
| Secrets | 1 |
| Cluster Roles | 6 |
| Controller Watches | 15 |

## Component Architecture

CRDs, controllers, and owned Kubernetes resources.

```mermaid
graph LR
    %% Component architecture for spark-operator

    classDef crd fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef controller fill:#3498db,stroke:#2980b9,color:#fff
    classDef owned fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef dep fill:#f39c12,stroke:#e67e22,color:#fff

    subgraph controller["spark-operator Controller"]
        dep_1["controller-manager"]
        class dep_1 controller
        dep_2["spark-operator-controller"]
        class dep_2 controller
        dep_3["spark-operator-module-controller-manager"]
        class dep_3 controller
        dep_4["spark-operator-webhook"]
        class dep_4 controller
    end

    crd_SparkConnect{{"SparkConnect\nsparkoperator.k8s.io/v1alpha1"}}
    class crd_SparkConnect crd
    crd_SparkConnect -->|"For (reconciles)"| controller
    crd_ScheduledSparkApplication{{"ScheduledSparkApplication\nsparkoperator.k8s.io/v1beta2"}}
    class crd_ScheduledSparkApplication crd
    crd_SparkApplication{{"SparkApplication\nsparkoperator.k8s.io/v1beta2"}}
    class crd_SparkApplication crd
    controller -->|"Owns"| owned_5["ClusterRole"]
    class owned_5 owned
    controller -->|"Owns"| owned_6["ClusterRoleBinding"]
    class owned_6 owned
    controller -->|"Owns"| owned_7["ConfigMap"]
    class owned_7 owned
    controller -->|"Owns"| owned_8["CustomResourceDefinition"]
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
    controller -->|"Owns"| owned_14["Service"]
    class owned_14 owned
    controller -->|"Owns"| owned_15["ServiceAccount"]
    class owned_15 owned
    controller -->|"Owns"| owned_16["ValidatingWebhookConfiguration"]
    class owned_16 owned
    controller -.->|"depends on"| odh_17["odh-platform-utilities"]
    class odh_17 dep
```

### CRDs

| Group | Version | Kind | Scope | Fields | Validation Rules | Discovery | Source |
|-------|---------|------|-------|--------|------------------|-----------|--------|
| sparkoperator.k8s.io | v1alpha1 | SparkConnect | Namespaced | 95 | 0 | YAML | [`config/crd/bases/sparkoperator.k8s.io_sparkconnects.yaml`](https://github.com/kubeflow/spark-operator/blob/ba4f90ed3a4296379cb602b3bcab96524cb92690/config/crd/bases/sparkoperator.k8s.io_sparkconnects.yaml) |
| sparkoperator.k8s.io | v1beta2 | ScheduledSparkApplication | Namespaced | 1741 | 0 | YAML | [`config/crd/bases/sparkoperator.k8s.io_scheduledsparkapplications.yaml`](https://github.com/kubeflow/spark-operator/blob/ba4f90ed3a4296379cb602b3bcab96524cb92690/config/crd/bases/sparkoperator.k8s.io_scheduledsparkapplications.yaml) |
| sparkoperator.k8s.io | v1beta2 | SparkApplication | Namespaced | 1744 | 0 | YAML | [`config/crd/bases/sparkoperator.k8s.io_sparkapplications.yaml`](https://github.com/kubeflow/spark-operator/blob/ba4f90ed3a4296379cb602b3bcab96524cb92690/config/crd/bases/sparkoperator.k8s.io_sparkapplications.yaml) |

## Dependencies

### Internal Platform Dependencies

| Component | Interaction |
|-----------|-------------|
| odh-platform-utilities | Go module dependency: github.com/opendatahub-io/odh-platform-utilities |

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/go-logr/logr | v1.4.3 |
| github.com/prometheus/client_golang | v1.23.2 |
| k8s.io/api | v0.35.4 |
| k8s.io/api | v0.35.2 |
| k8s.io/apiextensions-apiserver | v0.35.2 |
| k8s.io/apiextensions-apiserver | v0.35.4 |
| k8s.io/apimachinery | v0.35.4 |
| k8s.io/apimachinery | v0.35.4 |
| k8s.io/apiserver | v0.35.4 |
| k8s.io/client-go | v0.35.2 |
| k8s.io/client-go | v0.35.4 |
| sigs.k8s.io/controller-runtime | v0.23.3 |
| sigs.k8s.io/controller-runtime | v0.23.3 |

