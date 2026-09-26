# rhods-operator: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

| Type | GVK | Source |
|------|-----|--------|
| Owns | /v1/ConfigMap | [`internal/controller/components/kueue/kueue_controller.go:63`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/internal/controller/components/kueue/kueue_controller.go#L63) |
| Owns | /v1/Secret | [`internal/controller/components/kueue/kueue_controller.go:64`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/internal/controller/components/kueue/kueue_controller.go#L64) |
| Owns | /v1/Service | [`internal/controller/components/kueue/kueue_controller.go:70`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/internal/controller/components/kueue/kueue_controller.go#L70) |
| Owns | /v1/ServiceAccount | [`internal/controller/components/kueue/kueue_controller.go:69`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/internal/controller/components/kueue/kueue_controller.go#L69) |
| Owns | admissionregistration.k8s.io/v1/MutatingWebhookConfiguration | [`internal/controller/components/kueue/kueue_controller.go:74`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/internal/controller/components/kueue/kueue_controller.go#L74) |
| Owns | admissionregistration.k8s.io/v1/ValidatingWebhookConfiguration | [`internal/controller/components/kueue/kueue_controller.go:75`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/internal/controller/components/kueue/kueue_controller.go#L75) |
| Owns | apps/v1/Deployment | [`internal/controller/components/kueue/kueue_controller.go:76`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/internal/controller/components/kueue/kueue_controller.go#L76) |
| Owns | monitoring/v1/PodMonitor | [`internal/controller/components/kueue/kueue_controller.go:72`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/internal/controller/components/kueue/kueue_controller.go#L72) |
| Owns | monitoring/v1/PrometheusRule | [`internal/controller/components/kueue/kueue_controller.go:73`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/internal/controller/components/kueue/kueue_controller.go#L73) |
| Owns | networking.k8s.io/v1/NetworkPolicy | [`internal/controller/components/kueue/kueue_controller.go:71`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/internal/controller/components/kueue/kueue_controller.go#L71) |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | [`internal/controller/components/kueue/kueue_controller.go:66`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/internal/controller/components/kueue/kueue_controller.go#L66) |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | [`internal/controller/services/auth/auth_controller.go:60`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/internal/controller/services/auth/auth_controller.go#L60) |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | [`internal/controller/services/auth/auth_controller.go:59`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/internal/controller/services/auth/auth_controller.go#L59) |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | [`internal/controller/components/kueue/kueue_controller.go:65`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/internal/controller/components/kueue/kueue_controller.go#L65) |
| Owns | rbac.authorization.k8s.io/v1/Role | [`internal/controller/components/kueue/kueue_controller.go:67`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/internal/controller/components/kueue/kueue_controller.go#L67) |
| Owns | rbac.authorization.k8s.io/v1/Role | [`internal/controller/services/auth/auth_controller.go:61`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/internal/controller/services/auth/auth_controller.go#L61) |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | [`internal/controller/components/kueue/kueue_controller.go:68`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/internal/controller/components/kueue/kueue_controller.go#L68) |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | [`internal/controller/services/auth/auth_controller.go:62`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/internal/controller/services/auth/auth_controller.go#L62) |
| Watches | /v1/Namespace | [`internal/controller/components/kueue/kueue_controller.go:150`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/internal/controller/components/kueue/kueue_controller.go#L150) |
| Watches | rbac.authorization.k8s.io/v1/ClusterRole | [`internal/controller/components/kueue/kueue_controller.go:144`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/internal/controller/components/kueue/kueue_controller.go#L144) |
| Watches | services/v1alpha1/Auth | [`internal/controller/components/kueue/kueue_controller.go:161`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/internal/controller/components/kueue/kueue_controller.go#L161) |

### Programmatic Resource Operations

| Verb | Kind | Group | Condition |
|------|------|-------|----------|
| update | DSCInitialization | dscinitialization |  |
| delete | FeatureTracker | features |  |

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for rhods-operator

    participant KubernetesAPI as Kubernetes API
    participant aws_cloud_manager_operator as aws-cloud-manager-operator
    participant azure_cloud_manager_operator as azure-cloud-manager-operator
    participant controller_manager as controller-manager
    participant coreweave_cloud_manager_operator as coreweave-cloud-manager-operator
    participant rhods_operator as rhods-operator

    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ConfigMap
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update Secret
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update Service
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ServiceAccount
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update MutatingWebhookConfiguration
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ValidatingWebhookConfiguration
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update Deployment
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update PodMonitor
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update PrometheusRule
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update NetworkPolicy
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ClusterRole
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ClusterRole
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ClusterRoleBinding
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ClusterRoleBinding
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update Role
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update Role
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update RoleBinding
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update RoleBinding
    KubernetesAPI-->>+aws_cloud_manager_operator: Watch Namespace (informer)
    KubernetesAPI-->>+aws_cloud_manager_operator: Watch ClusterRole (informer)
    KubernetesAPI-->>+aws_cloud_manager_operator: Watch Auth (informer)

    Note over aws_cloud_manager_operator: Exposed Services
    Note right of aws_cloud_manager_operator: webhook-service:443/TCP []
    Note right of aws_cloud_manager_operator: webhook-service:443/TCP []
    Note right of aws_cloud_manager_operator: webhook-service:443/TCP []

    Note over KubernetesAPI: Defined CRDs
    Note right of KubernetesAPI: FeatureTracker (features.opendatahub.io/v1)
```

### Webhooks

| Name | Type | Path | Failure Policy | Service | Overlays | Enable Condition | Sources |
|------|------|------|----------------|---------|----------|------------------|----------|
| Defaulter-webhook | mutating | /mutate-datasciencecluster-v1 |  |  |  |  |  |
| Defaulter-webhook | mutating | /mutate-datasciencecluster-v2 |  |  |  |  |  |
| Injector-webhook | mutating | /mutate-prometheus-monitors |  |  |  |  |  |
| Validator-webhook | validating | /validate-datasciencecluster-v1 |  |  |  |  |  |
| Validator-webhook | validating | /validate-datasciencecluster-v2 |  |  |  |  |  |
| Validator-webhook | validating | /validate-dscinitialization-v1 |  |  |  |  |  |
| Validator-webhook | validating | /validate-dscinitialization-v2 |  |  |  |  |  |
| conversion-unknown | conversion | /convert |  | system/webhook-service |  | .ManagementState == operatorv1.Managed &amp;&amp; .ManagementState == "" | [`config/crd/patches/webhook_in_datasciencecluster_datascienceclusters.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/config/crd/patches/webhook_in_datasciencecluster_datascienceclusters.yaml) |
| conversion-unknown | conversion | /convert |  | system/webhook-service |  | .ManagementState == operatorv1.Managed &amp;&amp; .ManagementState == "" | [`config/crd/patches/webhook_in_dscinitialization_dscinitializations.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/config/crd/patches/webhook_in_dscinitialization_dscinitializations.yaml) |
| conversion-unknown | conversion | /convert |  | system/webhook-service |  | .ManagementState == operatorv1.Managed &amp;&amp; .ManagementState == "" | [`config/crd/patches/webhook_in_services_auths.yaml`](https://github.com/opendatahub-io/rhods-operator/blob/c797c263e2b046b249b9b6f501b5bec8bc6f8527/config/crd/patches/webhook_in_services_auths.yaml) |

#### Defaulter-webhook Behavior

| Field | Operation | Condition |
|-------|-----------|----------|
| registriesNamespace | set | modelRegistry.ManagementState == operatorv1.Managed &amp;&amp; modelRegistry.RegistriesNamespace == "" |
| nim.managementState | set | kserve.ManagementState == operatorv1.Managed &amp;&amp; kserve.NIM.ManagementState == "" |

#### Defaulter-webhook Behavior

| Field | Operation | Condition |
|-------|-----------|----------|
| registriesNamespace | set | modelRegistry.ManagementState == operatorv1.Managed &amp;&amp; modelRegistry.RegistriesNamespace == "" |
| nim.managementState | set | kserve.ManagementState == operatorv1.Managed &amp;&amp; kserve.NIM.ManagementState == "" |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

