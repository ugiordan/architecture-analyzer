# kubeflow: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

| Type | GVK | Source |
|------|-----|--------|
| For | api/v1/Notebook | [`components/odh-notebook-controller/controllers/notebook_controller.go:761`](https://github.com/opendatahub-io/kubeflow/blob/9ae2b2af88080c54381fe22652403fc61291e134/components/odh-notebook-controller/controllers/notebook_controller.go#L761) |
| For | api/v1beta1/Notebook | [`components/notebook-controller/controllers/culling_controller.go:580`](https://github.com/opendatahub-io/kubeflow/blob/9ae2b2af88080c54381fe22652403fc61291e134/components/notebook-controller/controllers/culling_controller.go#L580) |
| For | api/v1beta1/Notebook | [`components/notebook-controller/controllers/notebook_controller.go:801`](https://github.com/opendatahub-io/kubeflow/blob/9ae2b2af88080c54381fe22652403fc61291e134/components/notebook-controller/controllers/notebook_controller.go#L801) |
| Owns | /v1/ConfigMap | [`components/odh-notebook-controller/controllers/notebook_controller.go:765`](https://github.com/opendatahub-io/kubeflow/blob/9ae2b2af88080c54381fe22652403fc61291e134/components/odh-notebook-controller/controllers/notebook_controller.go#L765) |
| Owns | /v1/Secret | [`components/odh-notebook-controller/controllers/notebook_controller.go:764`](https://github.com/opendatahub-io/kubeflow/blob/9ae2b2af88080c54381fe22652403fc61291e134/components/odh-notebook-controller/controllers/notebook_controller.go#L764) |
| Owns | /v1/Service | [`components/odh-notebook-controller/controllers/notebook_controller.go:763`](https://github.com/opendatahub-io/kubeflow/blob/9ae2b2af88080c54381fe22652403fc61291e134/components/odh-notebook-controller/controllers/notebook_controller.go#L763) |
| Owns | /v1/Service | [`components/notebook-controller/controllers/notebook_controller.go:803`](https://github.com/opendatahub-io/kubeflow/blob/9ae2b2af88080c54381fe22652403fc61291e134/components/notebook-controller/controllers/notebook_controller.go#L803) |
| Owns | /v1/ServiceAccount | [`components/odh-notebook-controller/controllers/notebook_controller.go:762`](https://github.com/opendatahub-io/kubeflow/blob/9ae2b2af88080c54381fe22652403fc61291e134/components/odh-notebook-controller/controllers/notebook_controller.go#L762) |
| Owns | apps/v1/StatefulSet | [`components/notebook-controller/controllers/notebook_controller.go:802`](https://github.com/opendatahub-io/kubeflow/blob/9ae2b2af88080c54381fe22652403fc61291e134/components/notebook-controller/controllers/notebook_controller.go#L802) |
| Owns | networking.k8s.io/v1/NetworkPolicy | [`components/odh-notebook-controller/controllers/notebook_controller.go:766`](https://github.com/opendatahub-io/kubeflow/blob/9ae2b2af88080c54381fe22652403fc61291e134/components/odh-notebook-controller/controllers/notebook_controller.go#L766) |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | [`components/odh-notebook-controller/controllers/notebook_controller.go:767`](https://github.com/opendatahub-io/kubeflow/blob/9ae2b2af88080c54381fe22652403fc61291e134/components/odh-notebook-controller/controllers/notebook_controller.go#L767) |
| Watches | /v1/ConfigMap | [`components/odh-notebook-controller/controllers/notebook_controller.go:840`](https://github.com/opendatahub-io/kubeflow/blob/9ae2b2af88080c54381fe22652403fc61291e134/components/odh-notebook-controller/controllers/notebook_controller.go#L840) |
| Watches | /v1/Namespace | [`components/odh-notebook-controller/controllers/notebook_controller.go:841`](https://github.com/opendatahub-io/kubeflow/blob/9ae2b2af88080c54381fe22652403fc61291e134/components/odh-notebook-controller/controllers/notebook_controller.go#L841) |
| Watches | apis/v1/HTTPRoute | [`components/odh-notebook-controller/controllers/notebook_controller.go:771`](https://github.com/opendatahub-io/kubeflow/blob/9ae2b2af88080c54381fe22652403fc61291e134/components/odh-notebook-controller/controllers/notebook_controller.go#L771) |
| Watches | apis/v1beta1/ReferenceGrant | [`components/odh-notebook-controller/controllers/notebook_controller.go:801`](https://github.com/opendatahub-io/kubeflow/blob/9ae2b2af88080c54381fe22652403fc61291e134/components/odh-notebook-controller/controllers/notebook_controller.go#L801) |

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for kubeflow

    participant KubernetesAPI as Kubernetes API
    participant controller_manager as controller-manager
    participant deployment as deployment
    participant manager as manager

    KubernetesAPI->>+controller_manager: Watch Notebook (reconcile)
    KubernetesAPI->>+controller_manager: Watch Notebook (reconcile)
    KubernetesAPI->>+controller_manager: Watch Notebook (reconcile)
    controller_manager->>KubernetesAPI: Create/Update ConfigMap
    controller_manager->>KubernetesAPI: Create/Update Secret
    controller_manager->>KubernetesAPI: Create/Update Service
    controller_manager->>KubernetesAPI: Create/Update Service
    controller_manager->>KubernetesAPI: Create/Update ServiceAccount
    controller_manager->>KubernetesAPI: Create/Update StatefulSet
    controller_manager->>KubernetesAPI: Create/Update NetworkPolicy
    controller_manager->>KubernetesAPI: Create/Update RoleBinding
    KubernetesAPI-->>+controller_manager: Watch ConfigMap (informer)
    KubernetesAPI-->>+controller_manager: Watch Namespace (informer)
    KubernetesAPI-->>+controller_manager: Watch HTTPRoute (informer)
    KubernetesAPI-->>+controller_manager: Watch ReferenceGrant (informer)

    Note over controller_manager: Exposed Services
    Note right of controller_manager: service:443/TCP []
    Note right of controller_manager: service:8443/TCP [metrics]
    Note right of controller_manager: webhook-service:443/TCP [webhook]
```

### Webhooks

| Name | Type | Path | Failure Policy | Service | Overlays | Enable Condition | Sources |
|------|------|------|----------------|---------|----------|------------------|----------|
| notebooks-validation.opendatahub.io | validating | /validate-notebook-v1 | fail |  |  |  | [`components/odh-notebook-controller/controllers/notebook_validating_webhook.go`](https://github.com/opendatahub-io/kubeflow/blob/9ae2b2af88080c54381fe22652403fc61291e134/components/odh-notebook-controller/controllers/notebook_validating_webhook.go), [`components/odh-notebook-controller/controllers/notebook_validating_webhook.go`](https://github.com/opendatahub-io/kubeflow/blob/9ae2b2af88080c54381fe22652403fc61291e134/components/odh-notebook-controller/controllers/notebook_validating_webhook.go) |
| notebooks.opendatahub.io | mutating | /mutate-notebook-v1 | fail |  |  |  | [`components/odh-notebook-controller/controllers/notebook_mutating_webhook.go`](https://github.com/opendatahub-io/kubeflow/blob/9ae2b2af88080c54381fe22652403fc61291e134/components/odh-notebook-controller/controllers/notebook_mutating_webhook.go), [`components/odh-notebook-controller/controllers/notebook_mutating_webhook.go`](https://github.com/opendatahub-io/kubeflow/blob/9ae2b2af88080c54381fe22652403fc61291e134/components/odh-notebook-controller/controllers/notebook_mutating_webhook.go) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

