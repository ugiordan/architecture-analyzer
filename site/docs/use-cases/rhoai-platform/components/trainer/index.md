# trainer

> **Architecture snapshot: 2026-09-19** (2026-09-19)


**Repository:** kubeflow/trainer  
**Analyzer:** arch-analyzer dev  
**Extracted:** 2026-09-19T04:03:52Z

## Summary

| Metric | Count |
|--------|-------|
| CRDs | 4 |
| Deployments | 4 |
| Services | 0 |
| Secrets | 1 |
| Cluster Roles | 10 |
| Controller Watches | 1 |

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
        dep_1["kubeflow-trainer-controller-manager"]
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
    crd_TrainJob{{"TrainJob\ntrainer.kubeflow.org/v1alpha1"}}
    class crd_TrainJob crd
    crd_TrainingRuntime{{"TrainingRuntime\ntrainer.kubeflow.org/v1alpha1"}}
    class crd_TrainingRuntime crd
```

### CRDs

| Group | Version | Kind | Scope | Fields | Validation Rules | Discovery | Source |
|-------|---------|------|-------|--------|------------------|-----------|--------|
| trainer.kubeflow.org | v1alpha1 | ClusterTrainingRuntime | Cluster | 1307 | 18 | YAML | [`manifests/base/crds/trainer.kubeflow.org_clustertrainingruntimes.yaml`](https://github.com/kubeflow/trainer/blob/28e71a937606ddca86844c8ea33451e4686f4c75/manifests/base/crds/trainer.kubeflow.org_clustertrainingruntimes.yaml) |
| trainer.kubeflow.org | v1alpha1 | OptimizationJob | Namespaced | 709 | 17 | YAML + Go AST | [`manifests/base/crds/trainer.kubeflow.org_optimizationjobs.yaml`](https://github.com/kubeflow/trainer/blob/28e71a937606ddca86844c8ea33451e4686f4c75/manifests/base/crds/trainer.kubeflow.org_optimizationjobs.yaml) |
| trainer.kubeflow.org | v1alpha1 | TrainJob | Namespaced | 688 | 11 | YAML | [`manifests/base/crds/trainer.kubeflow.org_trainjobs.yaml`](https://github.com/kubeflow/trainer/blob/28e71a937606ddca86844c8ea33451e4686f4c75/manifests/base/crds/trainer.kubeflow.org_trainjobs.yaml) |
| trainer.kubeflow.org | v1alpha1 | TrainingRuntime | Namespaced | 1307 | 18 | YAML | [`manifests/base/crds/trainer.kubeflow.org_trainingruntimes.yaml`](https://github.com/kubeflow/trainer/blob/28e71a937606ddca86844c8ea33451e4686f4c75/manifests/base/crds/trainer.kubeflow.org_trainingruntimes.yaml) |

## Dependencies

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/go-logr/logr | v1.4.4 |
| k8s.io/api | v0.37.0 |
| k8s.io/apiextensions-apiserver | v0.37.0 |
| k8s.io/apimachinery | v0.37.0 |
| k8s.io/apiserver | v0.37.0 |
| k8s.io/client-go | v0.37.0 |
| sigs.k8s.io/controller-runtime | v0.24.1 |

