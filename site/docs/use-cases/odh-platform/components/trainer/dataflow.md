# trainer: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

| Type | GVK | Source |
|------|-----|--------|
| For | trainer/v1alpha1/OptimizationJob | [`pkg/controller/optimizationjob_controller.go:447`](https://github.com/opendatahub-io/trainer/blob/dc34bddaeca85ad47d5dd7f26bd73eec70042821/pkg/controller/optimizationjob_controller.go#L447) |
| Owns | /v1/Service | [`pkg/controller/optimizationjob_controller.go:450`](https://github.com/opendatahub-io/trainer/blob/dc34bddaeca85ad47d5dd7f26bd73eec70042821/pkg/controller/optimizationjob_controller.go#L450) |
| Owns | apps/v1/Deployment | [`pkg/controller/optimizationjob_controller.go:449`](https://github.com/opendatahub-io/trainer/blob/dc34bddaeca85ad47d5dd7f26bd73eec70042821/pkg/controller/optimizationjob_controller.go#L449) |
| Owns | trainer/v1alpha1/TrainJob | [`pkg/controller/optimizationjob_controller.go:448`](https://github.com/opendatahub-io/trainer/blob/dc34bddaeca85ad47d5dd7f26bd73eec70042821/pkg/controller/optimizationjob_controller.go#L448) |
| Watches | /v1/ConfigMap | [`rhoai/operator/internal/controller/trainer_controller.go:213`](https://github.com/opendatahub-io/trainer/blob/dc34bddaeca85ad47d5dd7f26bd73eec70042821/rhoai/operator/internal/controller/trainer_controller.go#L213) |
| Watches | /v1/Service | [`rhoai/operator/internal/controller/trainer_controller.go:212`](https://github.com/opendatahub-io/trainer/blob/dc34bddaeca85ad47d5dd7f26bd73eec70042821/rhoai/operator/internal/controller/trainer_controller.go#L212) |
| Watches | admissionregistration.k8s.io/v1/ValidatingWebhookConfiguration | [`rhoai/operator/internal/controller/trainer_controller.go:214`](https://github.com/opendatahub-io/trainer/blob/dc34bddaeca85ad47d5dd7f26bd73eec70042821/rhoai/operator/internal/controller/trainer_controller.go#L214) |
| Watches | apps/v1/Deployment | [`rhoai/operator/internal/controller/trainer_controller.go:210`](https://github.com/opendatahub-io/trainer/blob/dc34bddaeca85ad47d5dd7f26bd73eec70042821/rhoai/operator/internal/controller/trainer_controller.go#L210) |

### Programmatic Resource Operations

| Verb | Kind | Group | Condition |
|------|------|-------|----------|
| patch | ClusterTrainingRuntime | trainer |  |
| patch | OptimizationJob | trainer |  |
| patch | TrainingRuntime | trainer |  |
| patch | TrainJob | trainer |  |
| delete | JobSet | jobset |  |

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for trainer

    participant KubernetesAPI as Kubernetes API
    participant controller_manager as controller-manager
    participant kubeflow_trainer_controller_manager as kubeflow-trainer-controller-manager

    KubernetesAPI->>+controller_manager: Watch OptimizationJob (reconcile)
    controller_manager->>KubernetesAPI: Create/Update Service
    controller_manager->>KubernetesAPI: Create/Update Deployment
    controller_manager->>KubernetesAPI: Create/Update TrainJob
    KubernetesAPI-->>+controller_manager: Watch ConfigMap (informer)
    KubernetesAPI-->>+controller_manager: Watch Service (informer)
    KubernetesAPI-->>+controller_manager: Watch ValidatingWebhookConfiguration (informer)
    KubernetesAPI-->>+controller_manager: Watch Deployment (informer)

    Note over KubernetesAPI: Defined CRDs
    Note right of KubernetesAPI: ClusterTrainingRuntime (trainer.kubeflow.org/v1alpha1)
    Note right of KubernetesAPI: OptimizationJob (trainer.kubeflow.org/v1alpha1)
    Note right of KubernetesAPI: TrainJob (trainer.kubeflow.org/v1alpha1)
    Note right of KubernetesAPI: TrainingRuntime (trainer.kubeflow.org/v1alpha1)
```

### Webhooks

| Name | Type | Path | Failure Policy | Service | Overlays | Enable Condition | Sources |
|------|------|------|----------------|---------|----------|------------------|----------|
| ClusterTrainingRuntimeValidator-webhook | validating | /validate-trainer-kubeflow-org-v1alpha1-clustertrainingruntime |  |  |  |  |  |
| OptimizationJobValidator-webhook | validating | /validate-trainer-kubeflow-org-v1alpha1-optimizationjob |  |  |  |  |  |
| TrainJobDefaulter-webhook | mutating | /mutate-trainer-kubeflow-org-v1alpha1-trainjob |  |  |  |  |  |
| TrainJobValidator-webhook | validating | /validate-trainer-kubeflow-org-v1alpha1-trainjob |  |  |  |  |  |
| TrainingRuntimeValidator-webhook | validating | /validate-trainer-kubeflow-org-v1alpha1-trainingruntime |  |  |  |  |  |

#### OptimizationJobValidator-webhook Behavior

| Field | Operation | Condition |
|-------|-----------|----------|
| spec | forbidden |  |

#### TrainJobDefaulter-webhook Behavior

| Field | Operation | Condition |
|-------|-----------|----------|
| time | set | ok |

### HTTP Endpoints

| Method | Path | Source |
|--------|------|--------|
| * | / | [`pkg/statusserver/server.go:89`](https://github.com/opendatahub-io/trainer/blob/dc34bddaeca85ad47d5dd7f26bd73eec70042821/pkg/statusserver/server.go#L89) |
| * | POST  | [`pkg/statusserver/server.go:88`](https://github.com/opendatahub-io/trainer/blob/dc34bddaeca85ad47d5dd7f26bd73eec70042821/pkg/statusserver/server.go#L88) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

### Helm

**Chart:** kubeflow-trainer v2.3.0

