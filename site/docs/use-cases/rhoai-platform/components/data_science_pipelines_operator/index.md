# data-science-pipelines-operator

> **Architecture snapshot: 2026-09-07** (2026-09-07)


**Repository:** opendatahub-io/data-science-pipelines-operator  
**Analyzer:** arch-analyzer dev  
**Extracted:** 2026-09-07T03:57:07Z

## Summary

| Metric | Count |
|--------|-------|
| CRDs | 4 |
| Deployments | 8 |
| Services | 8 |
| Secrets | 0 |
| Cluster Roles | 4 |
| Controller Watches | 12 |

## Component Architecture

CRDs, controllers, and owned Kubernetes resources.

```mermaid
graph LR
    %% Component architecture for data-science-pipelines-operator

    classDef crd fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef controller fill:#3498db,stroke:#2980b9,color:#fff
    classDef owned fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef dep fill:#f39c12,stroke:#e67e22,color:#fff

    subgraph controller["data-science-pipelines-operator Controller"]
        dep_1["data-science-pipelines-operator-controller-manager"]
        class dep_1 controller
        dep_2["ds-pipeline-workflow-controller-template-value"]
        class dep_2 controller
        dep_3["mariadb-template-value"]
        class dep_3 controller
        dep_4["minio-template-value"]
        class dep_4 controller
        dep_5["template-value"]
        class dep_5 controller
        dep_6["template-value"]
        class dep_6 controller
        dep_7["template-value"]
        class dep_7 controller
        dep_8["template-value"]
        class dep_8 controller
    end

    crd_DataSciencePipelinesApplication{{"DataSciencePipelinesApplication\ndatasciencepipelinesapplications.opendatahub.io/v1"}}
    class crd_DataSciencePipelinesApplication crd
    crd_DataSciencePipelinesApplication -->|"For (reconciles)"| controller
    crd_ScheduledWorkflow{{"ScheduledWorkflow\nkubeflow.org/v1beta1"}}
    class crd_ScheduledWorkflow crd
    crd_Pipeline{{"Pipeline\npipelines.kubeflow.org/v2beta1"}}
    class crd_Pipeline crd
    crd_PipelineVersion{{"PipelineVersion\npipelines.kubeflow.org/v2beta1"}}
    class crd_PipelineVersion crd
    controller -->|"Owns"| owned_9["ConfigMap"]
    class owned_9 owned
    controller -->|"Owns"| owned_10["Deployment"]
    class owned_10 owned
    controller -->|"Owns"| owned_11["NetworkPolicy"]
    class owned_11 owned
    controller -->|"Owns"| owned_12["PersistentVolumeClaim"]
    class owned_12 owned
    controller -->|"Owns"| owned_13["Role"]
    class owned_13 owned
    controller -->|"Owns"| owned_14["RoleBinding"]
    class owned_14 owned
    controller -->|"Owns"| owned_15["Route"]
    class owned_15 owned
    controller -->|"Owns"| owned_16["Secret"]
    class owned_16 owned
    controller -->|"Owns"| owned_17["Service"]
    class owned_17 owned
    controller -->|"Owns"| owned_18["ServiceAccount"]
    class owned_18 owned
    controller -.->|"depends on"| odh_19["mlflow-operator"]
    class odh_19 dep
    controller -.->|"depends on"| odh_20["operator-chaos"]
    class odh_20 dep
```

### CRDs

| Group | Version | Kind | Scope | Fields | Validation Rules | Discovery | Source |
|-------|---------|------|-------|--------|------------------|-----------|--------|
| datasciencepipelinesapplications.opendatahub.io | v1 | DataSciencePipelinesApplication | Namespaced | 208 | 2 | YAML | [`config/crd/bases/datasciencepipelinesapplications.opendatahub.io_datasciencepipelinesapplications.yaml`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/17634f8eb9c92fcc6a71f9f4d5e38cda37cd762c/config/crd/bases/datasciencepipelinesapplications.opendatahub.io_datasciencepipelinesapplications.yaml) |
| kubeflow.org | v1beta1 | ScheduledWorkflow | Namespaced | 5 | 0 | YAML | [`config/crd/bases/scheduledworkflows.yaml`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/17634f8eb9c92fcc6a71f9f4d5e38cda37cd762c/config/crd/bases/scheduledworkflows.yaml) |
| pipelines.kubeflow.org | v2beta1 | Pipeline | Namespaced | 7 | 0 | YAML | [`config/crd/bases/pipelines.kubeflow.org_pipelines.yaml`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/17634f8eb9c92fcc6a71f9f4d5e38cda37cd762c/config/crd/bases/pipelines.kubeflow.org_pipelines.yaml) |
| pipelines.kubeflow.org | v2beta1 | PipelineVersion | Namespaced | 19 | 0 | YAML | [`config/crd/bases/pipelines.kubeflow.org_pipelineversions.yaml`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/17634f8eb9c92fcc6a71f9f4d5e38cda37cd762c/config/crd/bases/pipelines.kubeflow.org_pipelineversions.yaml) |

## Dependencies

### Internal Platform Dependencies

| Component | Interaction |
|-----------|-------------|
| mlflow-operator | Go module dependency: github.com/opendatahub-io/mlflow-operator/api |
| operator-chaos | Go module dependency: github.com/opendatahub-io/operator-chaos |

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/go-logr/logr | v1.4.3 |
| github.com/prometheus/client_golang | v1.23.2 |
| k8s.io/api | v0.35.3 |
| k8s.io/apimachinery | v0.35.3 |
| k8s.io/client-go | v0.35.3 |
| sigs.k8s.io/controller-runtime | v0.23.3 |

