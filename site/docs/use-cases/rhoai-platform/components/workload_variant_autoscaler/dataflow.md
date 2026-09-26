# workload-variant-autoscaler: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

| Type | GVK | Source |
|------|-----|--------|
| For | /v1/ConfigMap | [`legacy/internal/controller/configmap_reconciler.go:109`](https://github.com/red-hat-data-services/workload-variant-autoscaler/blob/af7d918a087f9db38e19120774b60b0378615662/legacy/internal/controller/configmap_reconciler.go#L109) |
| For | api/v1/InferencePool | [`legacy/internal/controller/inferencepool_reconciler.go:113`](https://github.com/red-hat-data-services/workload-variant-autoscaler/blob/af7d918a087f9db38e19120774b60b0378615662/legacy/internal/controller/inferencepool_reconciler.go#L113) |
| For | apix/v1alpha2/InferencePool | [`legacy/internal/controller/inferencepool_reconciler.go:109`](https://github.com/red-hat-data-services/workload-variant-autoscaler/blob/af7d918a087f9db38e19120774b60b0378615662/legacy/internal/controller/inferencepool_reconciler.go#L109) |
| For | autoscaling/v2/HorizontalPodAutoscaler | [`legacy/internal/controller/hpa_reconciler.go:67`](https://github.com/red-hat-data-services/workload-variant-autoscaler/blob/af7d918a087f9db38e19120774b60b0378615662/legacy/internal/controller/hpa_reconciler.go#L67) |
| For | keda/v1alpha1/ScaledObject | [`legacy/internal/controller/scaledobject_reconciler.go:68`](https://github.com/red-hat-data-services/workload-variant-autoscaler/blob/af7d918a087f9db38e19120774b60b0378615662/legacy/internal/controller/scaledobject_reconciler.go#L68) |

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for workload-variant-autoscaler

    participant KubernetesAPI as Kubernetes API
    participant controller_manager as controller-manager

    KubernetesAPI->>+controller_manager: Watch ConfigMap (reconcile)
    KubernetesAPI->>+controller_manager: Watch InferencePool (reconcile)
    KubernetesAPI->>+controller_manager: Watch InferencePool (reconcile)
    KubernetesAPI->>+controller_manager: Watch HorizontalPodAutoscaler (reconcile)
    KubernetesAPI->>+controller_manager: Watch ScaledObject (reconcile)

    Note over controller_manager: Exposed Services
    Note right of controller_manager: controller-manager-metrics-service:8443/TCP [https]
```

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

### ConfigMaps

| Name | Data Keys | Source |
|------|-----------|--------|
| manager-config | config.yaml | [`legacy/config/components/openshift/configmap-patch.yaml`](https://github.com/red-hat-data-services/workload-variant-autoscaler/blob/af7d918a087f9db38e19120774b60b0378615662/legacy/config/components/openshift/configmap-patch.yaml) |
| service-classes-config | freemium.yaml, premium.yaml | [`legacy/deploy/configmap-serviceclass.yaml`](https://github.com/red-hat-data-services/workload-variant-autoscaler/blob/af7d918a087f9db38e19120774b60b0378615662/legacy/deploy/configmap-serviceclass.yaml) |
| wva-queueing-model-config | default | [`legacy/deploy/configmap-queueing-model.yaml`](https://github.com/red-hat-data-services/workload-variant-autoscaler/blob/af7d918a087f9db38e19120774b60b0378615662/legacy/deploy/configmap-queueing-model.yaml) |
| wva-saturation-scaling-config | default | [`legacy/deploy/configmap-saturation-scaling.yaml`](https://github.com/red-hat-data-services/workload-variant-autoscaler/blob/af7d918a087f9db38e19120774b60b0378615662/legacy/deploy/configmap-saturation-scaling.yaml) |

