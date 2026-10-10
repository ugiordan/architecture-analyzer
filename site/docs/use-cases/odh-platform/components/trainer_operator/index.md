# trainer-operator

> **Architecture snapshot: 2026-10-10** (2026-10-10)


**Repository:** opendatahub-io/trainer-operator  
**Analyzer:** arch-analyzer dev  
**Extracted:** 2026-10-10T05:16:56Z

## Summary

| Metric | Count |
|--------|-------|
| CRDs | 1 |
| Deployments | 4 |
| Services | 0 |
| Secrets | 1 |
| Cluster Roles | 15 |
| Controller Watches | 4 |

## Component Architecture

CRDs, controllers, and owned Kubernetes resources.

```mermaid
graph LR
    %% Component architecture for trainer-operator

    classDef crd fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef controller fill:#3498db,stroke:#2980b9,color:#fff
    classDef owned fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef dep fill:#f39c12,stroke:#e67e22,color:#fff

    subgraph controller["trainer-operator Controller"]
        dep_1["controller-manager"]
        class dep_1 controller
        dep_2["kubeflow-trainer-controller-manager"]
        class dep_2 controller
        dep_3["kubeflow-trainer-controller-manager"]
        class dep_3 controller
        dep_4["kubeflow-trainer-controller-manager"]
        class dep_4 controller
    end

    crd_Trainer{{"Trainer\ncomponents.platform.opendatahub.io/v1alpha1"}}
    class crd_Trainer crd
    watch_5["ConfigMap"] -->|"Watches"| controller
    class watch_5 external
    watch_6["Deployment"] -->|"Watches"| controller
    class watch_6 external
    watch_7["Service"] -->|"Watches"| controller
    class watch_7 external
    watch_8["ValidatingWebhookConfiguration"] -->|"Watches"| controller
    class watch_8 external
    controller -.->|"depends on"| odh_9["odh-platform-utilities"]
    class odh_9 dep
    controller -.->|"depends on"| odh_10["odh-platform-utilities"]
    class odh_10 dep
```

### CRDs

| Group | Version | Kind | Scope | Fields | Validation Rules | Discovery | Source |
|-------|---------|------|-------|--------|------------------|-----------|--------|
| components.platform.opendatahub.io | v1alpha1 | Trainer | Cluster | 30 | 1 | YAML | [`config/crd/bases/components.platform.opendatahub.io_trainers.yaml`](https://github.com/opendatahub-io/trainer-operator/blob/90322983edf391c7adb195654221cc83d172b2a3/config/crd/bases/components.platform.opendatahub.io_trainers.yaml) |

## Dependencies

### Internal Platform Dependencies

| Component | Interaction |
|-----------|-------------|
| odh-platform-utilities | Go module dependency: github.com/opendatahub-io/odh-platform-utilities |
| odh-platform-utilities | Go module dependency: github.com/opendatahub-io/odh-platform-utilities/framework |

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/go-logr/logr | v1.4.3 |
| github.com/operator-framework/api | v0.42.0 |
| k8s.io/api | v0.36.2 |
| k8s.io/apiextensions-apiserver | v0.36.2 |
| k8s.io/apimachinery | v0.36.2 |
| k8s.io/client-go | v0.36.2 |
| sigs.k8s.io/controller-runtime | v0.24.1 |

