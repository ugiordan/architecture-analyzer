# training-operator: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

| Type | GVK | Source |
|------|-----|--------|

### Programmatic Resource Operations

| Verb | Kind | Group | Condition |
|------|------|-------|----------|
| create | Pod |  |  |
| create | HorizontalPodAutoscaler | autoscaling |  |
| delete | HorizontalPodAutoscaler | autoscaling |  |
| update | HorizontalPodAutoscaler | autoscaling |  |

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for training-operator

    participant KubernetesAPI as Kubernetes API
    participant training_operator as training-operator


    Note over training_operator: Exposed Services
    Note right of training_operator: training-operator:8080/TCP [monitoring-port]
    Note right of training_operator: training-operator:443/TCP [webhook-server]

    Note over KubernetesAPI: Defined CRDs
    Note right of KubernetesAPI: JAXJob (kubeflow.org/v1)
    Note right of KubernetesAPI: MPIJob (kubeflow.org/v1)
    Note right of KubernetesAPI: PaddleJob (kubeflow.org/v1)
    Note right of KubernetesAPI: PyTorchJob (kubeflow.org/v1)
    Note right of KubernetesAPI: TFJob (kubeflow.org/v1)
    Note right of KubernetesAPI: XGBoostJob (kubeflow.org/v1)
```

### Webhooks

| Name | Type | Path | Failure Policy | Service | Overlays | Enable Condition | Sources |
|------|------|------|----------------|---------|----------|------------------|----------|
| Webhook-webhook | validating | /validate-kubeflow-org-v1-pytorchjob |  |  |  |  |  |
| Webhook-webhook | validating | /validate-kubeflow-org-v1-tfjob |  |  |  |  |  |
| Webhook-webhook | validating | /validate-kubeflow-org-v1-xgboostjob |  |  |  |  |  |
| Webhook-webhook | validating | /validate-kubeflow-org-v1-jaxjob |  |  |  |  |  |
| Webhook-webhook | validating | /validate-kubeflow-org-v1-paddlejob |  |  |  |  |  |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

