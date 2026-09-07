# data-science-pipelines-operator: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

| Type | GVK | Source |
|------|-----|--------|
| For | api/v1/DataSciencePipelinesApplication | [`controllers/dspipeline_controller.go:893`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/17634f8eb9c92fcc6a71f9f4d5e38cda37cd762c/controllers/dspipeline_controller.go#L893) |
| Owns | /v1/ConfigMap | [`controllers/dspipeline_controller.go:896`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/17634f8eb9c92fcc6a71f9f4d5e38cda37cd762c/controllers/dspipeline_controller.go#L896) |
| Owns | /v1/PersistentVolumeClaim | [`controllers/dspipeline_controller.go:899`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/17634f8eb9c92fcc6a71f9f4d5e38cda37cd762c/controllers/dspipeline_controller.go#L899) |
| Owns | /v1/Secret | [`controllers/dspipeline_controller.go:895`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/17634f8eb9c92fcc6a71f9f4d5e38cda37cd762c/controllers/dspipeline_controller.go#L895) |
| Owns | /v1/Service | [`controllers/dspipeline_controller.go:897`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/17634f8eb9c92fcc6a71f9f4d5e38cda37cd762c/controllers/dspipeline_controller.go#L897) |
| Owns | /v1/ServiceAccount | [`controllers/dspipeline_controller.go:898`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/17634f8eb9c92fcc6a71f9f4d5e38cda37cd762c/controllers/dspipeline_controller.go#L898) |
| Owns | apps/v1/Deployment | [`controllers/dspipeline_controller.go:894`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/17634f8eb9c92fcc6a71f9f4d5e38cda37cd762c/controllers/dspipeline_controller.go#L894) |
| Owns | networking.k8s.io/v1/NetworkPolicy | [`controllers/dspipeline_controller.go:900`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/17634f8eb9c92fcc6a71f9f4d5e38cda37cd762c/controllers/dspipeline_controller.go#L900) |
| Owns | rbac.authorization.k8s.io/v1/Role | [`controllers/dspipeline_controller.go:901`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/17634f8eb9c92fcc6a71f9f4d5e38cda37cd762c/controllers/dspipeline_controller.go#L901) |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | [`controllers/dspipeline_controller.go:902`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/17634f8eb9c92fcc6a71f9f4d5e38cda37cd762c/controllers/dspipeline_controller.go#L902) |
| Owns | route/v1/Route | [`controllers/dspipeline_controller.go:903`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/17634f8eb9c92fcc6a71f9f4d5e38cda37cd762c/controllers/dspipeline_controller.go#L903) |

### Programmatic Resource Operations

| Verb | Kind | Group | Condition |
|------|------|-------|----------|
| update | DataSciencePipelinesApplication | api |  |

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for data-science-pipelines-operator

    participant KubernetesAPI as Kubernetes API
    participant data_science_pipelines_operator_controller_manager as data-science-pipelines-operator-controller-manager
    participant ds_pipeline_workflow_controller_template_value as ds-pipeline-workflow-controller-template-value
    participant mariadb_template_value as mariadb-template-value
    participant minio_template_value as minio-template-value
    participant template_value as template-value

    KubernetesAPI->>+data_science_pipelines_operator_controller_manager: Watch DataSciencePipelinesApplication (reconcile)
    data_science_pipelines_operator_controller_manager->>KubernetesAPI: Create/Update ConfigMap
    data_science_pipelines_operator_controller_manager->>KubernetesAPI: Create/Update PersistentVolumeClaim
    data_science_pipelines_operator_controller_manager->>KubernetesAPI: Create/Update Secret
    data_science_pipelines_operator_controller_manager->>KubernetesAPI: Create/Update Service
    data_science_pipelines_operator_controller_manager->>KubernetesAPI: Create/Update ServiceAccount
    data_science_pipelines_operator_controller_manager->>KubernetesAPI: Create/Update Deployment
    data_science_pipelines_operator_controller_manager->>KubernetesAPI: Create/Update NetworkPolicy
    data_science_pipelines_operator_controller_manager->>KubernetesAPI: Create/Update Role
    data_science_pipelines_operator_controller_manager->>KubernetesAPI: Create/Update RoleBinding
    data_science_pipelines_operator_controller_manager->>KubernetesAPI: Create/Update Route

    Note over data_science_pipelines_operator_controller_manager: Exposed Services
    Note right of data_science_pipelines_operator_controller_manager: data-science-pipelines-operator-service:8080/TCP [metrics]
    Note right of data_science_pipelines_operator_controller_manager: ds-pipeline-workflow-controller-metrics-template-value:9090/TCP [metrics]
    Note right of data_science_pipelines_operator_controller_manager: mariadb-template-value:3306/TCP []
    Note right of data_science_pipelines_operator_controller_manager: minio-service:9000/TCP [http]
    Note right of data_science_pipelines_operator_controller_manager: minio-template-value:9000/TCP [http]
    Note right of data_science_pipelines_operator_controller_manager: minio-template-value:80/TCP [kfp-ui-http]
    Note right of data_science_pipelines_operator_controller_manager: ml-pipeline:8443/TCP [proxy]
    Note right of data_science_pipelines_operator_controller_manager: ml-pipeline:8888/TCP [http]
    Note right of data_science_pipelines_operator_controller_manager: ml-pipeline:8887/TCP [grpc]
    Note right of data_science_pipelines_operator_controller_manager: template-value:8443/TCP [proxy]
    Note right of data_science_pipelines_operator_controller_manager: template-value:8888/TCP [http]
    Note right of data_science_pipelines_operator_controller_manager: template-value:8887/TCP [grpc]
    Note right of data_science_pipelines_operator_controller_manager: template-value:8443/TCP [webhook]

    Note over KubernetesAPI: Defined CRDs
    Note right of KubernetesAPI: DataSciencePipelinesApplication (datasciencepipelinesapplications.opendatahub.io/v1)
    Note right of KubernetesAPI: ScheduledWorkflow (kubeflow.org/v1beta1)
    Note right of KubernetesAPI: Pipeline (pipelines.kubeflow.org/v2beta1)
    Note right of KubernetesAPI: PipelineVersion (pipelines.kubeflow.org/v2beta1)
```

### Webhooks

| Name | Type | Path | Failure Policy | Service | Overlays | Enable Condition | Sources |
|------|------|------|----------------|---------|----------|------------------|----------|
| pipelineversions.pipelines.kubeflow.org | mutating | /webhooks/mutate-pipelineversion | Fail | template-value/template-value |  |  | [`config/internal/webhook/mutating_webhook.yaml.tmpl`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/17634f8eb9c92fcc6a71f9f4d5e38cda37cd762c/config/internal/webhook/mutating_webhook.yaml.tmpl) |
| pipelineversions.pipelines.kubeflow.org | validating | /webhooks/validate-pipelineversion | Fail | template-value/template-value |  |  | [`config/internal/webhook/validating_webhook.yaml.tmpl`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/17634f8eb9c92fcc6a71f9f4d5e38cda37cd762c/config/internal/webhook/validating_webhook.yaml.tmpl) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

### ConfigMaps

| Name | Data Keys | Source |
|------|-----------|--------|
| workflow-controller-configmap |  | [`config/argo/configmap.workflow-controller-configmap.yaml`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/17634f8eb9c92fcc6a71f9f4d5e38cda37cd762c/config/argo/configmap.workflow-controller-configmap.yaml) |

