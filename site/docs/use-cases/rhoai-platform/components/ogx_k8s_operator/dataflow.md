# ogx-k8s-operator: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

| Type | GVK | Source |
|------|-----|--------|
| For | api/v1beta1/OGXServer | [`controllers/ogxserver_controller.go:819`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/controllers/ogxserver_controller.go#L819) |
| For | apis/v1alpha1/OGX | [`ogx-module/internal/controller/ogx/reconciler.go:140`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/ogx-module/internal/controller/ogx/reconciler.go#L140) |
| Owns | /v1/ConfigMap | [`controllers/ogxserver_controller.go:826`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/controllers/ogxserver_controller.go#L826) |
| Owns | /v1/ConfigMap | [`ogx-module/internal/controller/ogx/reconciler.go:142`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/ogx-module/internal/controller/ogx/reconciler.go#L142) |
| Owns | /v1/PersistentVolumeClaim | [`controllers/ogxserver_controller.go:839`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/controllers/ogxserver_controller.go#L839) |
| Owns | /v1/Service | [`controllers/ogxserver_controller.go:825`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/controllers/ogxserver_controller.go#L825) |
| Owns | /v1/Service | [`ogx-module/internal/controller/ogx/reconciler.go:143`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/ogx-module/internal/controller/ogx/reconciler.go#L143) |
| Owns | /v1/ServiceAccount | [`ogx-module/internal/controller/ogx/reconciler.go:144`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/ogx-module/internal/controller/ogx/reconciler.go#L144) |
| Owns | admissionregistration.k8s.io/v1/ValidatingWebhookConfiguration | [`ogx-module/internal/controller/ogx/reconciler.go:150`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/ogx-module/internal/controller/ogx/reconciler.go#L150) |
| Owns | apiextensions.k8s.io/v1/CustomResourceDefinition | [`ogx-module/internal/controller/ogx/reconciler.go:151`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/ogx-module/internal/controller/ogx/reconciler.go#L151) |
| Owns | apps/v1/Deployment | [`ogx-module/internal/controller/ogx/reconciler.go:141`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/ogx-module/internal/controller/ogx/reconciler.go#L141) |
| Owns | apps/v1/Deployment | [`controllers/ogxserver_controller.go:822`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/controllers/ogxserver_controller.go#L822) |
| Owns | autoscaling/v2/HorizontalPodAutoscaler | [`controllers/ogxserver_controller.go:824`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/controllers/ogxserver_controller.go#L824) |
| Owns | networking.k8s.io/v1/Ingress | [`controllers/ogxserver_controller.go:838`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/controllers/ogxserver_controller.go#L838) |
| Owns | networking.k8s.io/v1/NetworkPolicy | [`controllers/ogxserver_controller.go:837`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/controllers/ogxserver_controller.go#L837) |
| Owns | policy/v1/PodDisruptionBudget | [`ogx-module/internal/controller/ogx/reconciler.go:145`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/ogx-module/internal/controller/ogx/reconciler.go#L145) |
| Owns | policy/v1/PodDisruptionBudget | [`controllers/ogxserver_controller.go:823`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/controllers/ogxserver_controller.go#L823) |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | [`ogx-module/internal/controller/ogx/reconciler.go:148`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/ogx-module/internal/controller/ogx/reconciler.go#L148) |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | [`ogx-module/internal/controller/ogx/reconciler.go:149`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/ogx-module/internal/controller/ogx/reconciler.go#L149) |
| Owns | rbac.authorization.k8s.io/v1/Role | [`ogx-module/internal/controller/ogx/reconciler.go:146`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/ogx-module/internal/controller/ogx/reconciler.go#L146) |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | [`ogx-module/internal/controller/ogx/reconciler.go:147`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/ogx-module/internal/controller/ogx/reconciler.go#L147) |

### Programmatic Resource Operations

| Verb | Kind | Group | Condition |
|------|------|-------|----------|
| delete | ConfigMap |  |  |
| create | ConfigMap |  |  |
| patch | ConfigMap |  |  |

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for ogx-k8s-operator

    participant KubernetesAPI as Kubernetes API
    participant deployment as deployment
    participant ogx_k8s_operator_controller_manager as ogx-k8s-operator-controller-manager
    participant operator as operator

    KubernetesAPI->>+deployment: Watch OGXServer (reconcile)
    KubernetesAPI->>+deployment: Watch OGX (reconcile)
    deployment->>KubernetesAPI: Create/Update ConfigMap
    deployment->>KubernetesAPI: Create/Update ConfigMap
    deployment->>KubernetesAPI: Create/Update PersistentVolumeClaim
    deployment->>KubernetesAPI: Create/Update Service
    deployment->>KubernetesAPI: Create/Update Service
    deployment->>KubernetesAPI: Create/Update ServiceAccount
    deployment->>KubernetesAPI: Create/Update ValidatingWebhookConfiguration
    deployment->>KubernetesAPI: Create/Update CustomResourceDefinition
    deployment->>KubernetesAPI: Create/Update Deployment
    deployment->>KubernetesAPI: Create/Update Deployment
    deployment->>KubernetesAPI: Create/Update HorizontalPodAutoscaler
    deployment->>KubernetesAPI: Create/Update Ingress
    deployment->>KubernetesAPI: Create/Update NetworkPolicy
    deployment->>KubernetesAPI: Create/Update PodDisruptionBudget
    deployment->>KubernetesAPI: Create/Update PodDisruptionBudget
    deployment->>KubernetesAPI: Create/Update ClusterRole
    deployment->>KubernetesAPI: Create/Update ClusterRoleBinding
    deployment->>KubernetesAPI: Create/Update Role
    deployment->>KubernetesAPI: Create/Update RoleBinding

    Note over deployment: Exposed Services
    Note right of deployment: ogx-k8s-operator-controller-manager-metrics-service:8443/TCP [https]
    Note right of deployment: ogx-k8s-operator-webhook-service:443/TCP []

    Note over KubernetesAPI: Defined CRDs
    Note right of KubernetesAPI: LlamaStackDistribution (llamastack.io/v1alpha1)
    Note right of KubernetesAPI: OGXServer (ogx.io/v1beta1)
```

### Webhooks

| Name | Type | Path | Failure Policy | Service | Overlays | Enable Condition | Sources |
|------|------|------|----------------|---------|----------|------------------|----------|
| vogxserver.kb.io | validating | /validate-ogx-io-v1beta1-ogxserver | Fail | opendatahub/ogx-k8s-operator-webhook-service | config/overlays/odh |  | [`config/webhook/manifests.yaml`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/config/webhook/manifests.yaml), [`kustomize:config/overlays/odh (ogx-k8s-operator-validating-webhook-configuration)`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/kustomize:config/overlays/odh %28ogx-k8s-operator-validating-webhook-configuration%29) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

### ConfigMaps

| Name | Data Keys | Source |
|------|-----------|--------|
| config | platform-name | [`ogx-module/config/manager/configmap.yaml`](https://github.com/ogx-ai/ogx-k8s-operator/blob/bc1c6449e36dad7d00eea7bb5324772ca76dc794/ogx-module/config/manager/configmap.yaml) |

