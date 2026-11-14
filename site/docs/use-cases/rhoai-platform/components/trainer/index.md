# trainer

> **Architecture snapshot: 2026-11-14** (2026-11-14)


**Repository:** red-hat-data-services/trainer  
**Analyzer:** arch-analyzer dev  
**Extracted:** 2026-11-14T05:50:52Z

## Summary

| Metric | Count |
|--------|-------|
| CRDs | 4 |
| Deployments | 4 |
| Services | 0 |
| Secrets | 1 |
| Cluster Roles | 11 |
| Controller Watches | 9 |

## Component Architecture

CRDs, controllers, and owned Kubernetes resources.

```mermaid
graph LR
    %% Component architecture for trainer

    classDef crd fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef controller fill:#3498db,stroke:#2980b9,color:#fff
    classDef owned fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef dep fill:#f39c12,stroke:#e67e22,color:#fff

    subgraph controller["trainer Controller"]
        dep_1["controller-manager"]
        class dep_1 controller
        dep_2["kubeflow-trainer-controller-manager"]
        class dep_2 controller
        dep_3["kubeflow-trainer-controller-manager"]
        class dep_3 controller
        dep_4["kubeflow-trainer-controller-manager"]
        class dep_4 controller
    end

    crd_ClusterTrainingRuntime{{"ClusterTrainingRuntime\ntrainer.kubeflow.org/v1alpha1"}}
    class crd_ClusterTrainingRuntime crd
    crd_OptimizationJob{{"OptimizationJob\ntrainer.kubeflow.org/v1alpha1"}}
    class crd_OptimizationJob crd
    crd_OptimizationJob -->|"For (reconciles)"| controller
    crd_TrainJob{{"TrainJob\ntrainer.kubeflow.org/v1alpha1"}}
    class crd_TrainJob crd
    crd_TrainingRuntime{{"TrainingRuntime\ntrainer.kubeflow.org/v1alpha1"}}
    class crd_TrainingRuntime crd
    controller -->|"Owns"| owned_5["Deployment"]
    class owned_5 owned
    controller -->|"Owns"| owned_6["Service"]
    class owned_6 owned
    controller -->|"Owns"| owned_7["TrainJob"]
    class owned_7 owned
    watch_8["ConfigMap"] -->|"Watches"| controller
    class watch_8 external
    watch_9["Deployment"] -->|"Watches"| controller
    class watch_9 external
    watch_10["Service"] -->|"Watches"| controller
    class watch_10 external
    watch_11["ValidatingWebhookConfiguration"] -->|"Watches"| controller
    class watch_11 external
    controller -.->|"depends on"| odh_12["odh-platform-utilities"]
    class odh_12 dep
    controller -.->|"depends on"| odh_13["odh-platform-utilities"]
    class odh_13 dep
```

### CRDs

| Group | Version | Kind | Scope | Fields | Validation Rules | Discovery | Source |
|-------|---------|------|-------|--------|------------------|-----------|--------|
| trainer.kubeflow.org | v1alpha1 | ClusterTrainingRuntime | Cluster | 1307 | 19 | YAML | [`manifests/base/crds/trainer.kubeflow.org_clustertrainingruntimes.yaml`](https://github.com/red-hat-data-services/trainer/blob/cc8278f05532a988e858a8d7292adc50bab98800/manifests/base/crds/trainer.kubeflow.org_clustertrainingruntimes.yaml) |
| trainer.kubeflow.org | v1alpha1 | OptimizationJob | Namespaced | 729 | 21 | YAML + Go AST | [`manifests/base/crds/trainer.kubeflow.org_optimizationjobs.yaml`](https://github.com/red-hat-data-services/trainer/blob/cc8278f05532a988e858a8d7292adc50bab98800/manifests/base/crds/trainer.kubeflow.org_optimizationjobs.yaml) |
| trainer.kubeflow.org | v1alpha1 | TrainJob | Namespaced | 708 | 14 | YAML | [`manifests/base/crds/trainer.kubeflow.org_trainjobs.yaml`](https://github.com/red-hat-data-services/trainer/blob/cc8278f05532a988e858a8d7292adc50bab98800/manifests/base/crds/trainer.kubeflow.org_trainjobs.yaml) |
| trainer.kubeflow.org | v1alpha1 | TrainingRuntime | Namespaced | 1307 | 19 | YAML | [`manifests/base/crds/trainer.kubeflow.org_trainingruntimes.yaml`](https://github.com/red-hat-data-services/trainer/blob/cc8278f05532a988e858a8d7292adc50bab98800/manifests/base/crds/trainer.kubeflow.org_trainingruntimes.yaml) |

## Dependencies

### Internal Platform Dependencies

| Component | Interaction |
|-----------|-------------|
| odh-platform-utilities | Go module dependency: github.com/opendatahub-io/odh-platform-utilities |
| odh-platform-utilities | Go module dependency: github.com/opendatahub-io/odh-platform-utilities/framework |

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/go-logr/logr | v1.4.4 |
| github.com/go-logr/logr | v1.4.4 |
| github.com/operator-framework/api | v0.42.0 |
| github.com/prometheus/client_golang | v1.25.0 |
| google.golang.org/grpc | v1.84.0 |
| k8s.io/api | v0.37.1 |
| k8s.io/api | v0.37.1 |
| k8s.io/apiextensions-apiserver | v0.37.1 |
| k8s.io/apimachinery | v0.37.1 |
| k8s.io/apimachinery | v0.37.1 |
| k8s.io/apiserver | v0.37.1 |
| k8s.io/client-go | v0.37.1 |
| k8s.io/client-go | v0.37.1 |
| sigs.k8s.io/controller-runtime | v0.25.2 |
| sigs.k8s.io/controller-runtime | v0.25.2 |

