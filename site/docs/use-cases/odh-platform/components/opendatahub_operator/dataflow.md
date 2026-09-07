# opendatahub-operator: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

| Type | GVK | Source |
|------|-----|--------|
| Owns | /v1/ConfigMap | [`internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go:48`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go#L48) |
| Owns | /v1/ConfigMap | [`internal/controller/components/kueue/kueue_controller.go:60`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/kueue/kueue_controller.go#L60) |
| Owns | /v1/ConfigMap | [`internal/controller/components/ray/ray_controller.go:50`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/ray/ray_controller.go#L50) |
| Owns | /v1/ConfigMap | [`internal/controller/components/trustyai/trustyai_controller.go:50`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/trustyai/trustyai_controller.go#L50) |
| Owns | /v1/Secret | [`internal/controller/components/ray/ray_controller.go:51`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/ray/ray_controller.go#L51) |
| Owns | /v1/Secret | [`internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go:49`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go#L49) |
| Owns | /v1/Secret | [`internal/controller/components/kueue/kueue_controller.go:61`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/kueue/kueue_controller.go#L61) |
| Owns | /v1/Service | [`internal/controller/components/trustyai/trustyai_controller.go:56`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/trustyai/trustyai_controller.go#L56) |
| Owns | /v1/Service | [`internal/controller/components/ray/ray_controller.go:57`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/ray/ray_controller.go#L57) |
| Owns | /v1/Service | [`internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go:55`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go#L55) |
| Owns | /v1/Service | [`internal/controller/components/kueue/kueue_controller.go:67`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/kueue/kueue_controller.go#L67) |
| Owns | /v1/ServiceAccount | [`internal/controller/components/ray/ray_controller.go:56`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/ray/ray_controller.go#L56) |
| Owns | /v1/ServiceAccount | [`internal/controller/components/trustyai/trustyai_controller.go:51`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/trustyai/trustyai_controller.go#L51) |
| Owns | /v1/ServiceAccount | [`internal/controller/components/kueue/kueue_controller.go:66`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/kueue/kueue_controller.go#L66) |
| Owns | /v1/ServiceAccount | [`internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go:54`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go#L54) |
| Owns | admissionregistration.k8s.io/v1/MutatingWebhookConfiguration | [`internal/controller/components/kueue/kueue_controller.go:71`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/kueue/kueue_controller.go#L71) |
| Owns | admissionregistration.k8s.io/v1/ValidatingWebhookConfiguration | [`internal/controller/components/kueue/kueue_controller.go:72`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/kueue/kueue_controller.go#L72) |
| Owns | apps/v1/Deployment | [`internal/controller/components/kueue/kueue_controller.go:73`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/kueue/kueue_controller.go#L73) |
| Owns | apps/v1/Deployment | [`internal/controller/components/trustyai/trustyai_controller.go:57`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/trustyai/trustyai_controller.go#L57) |
| Owns | apps/v1/Deployment | [`internal/controller/components/ray/ray_controller.go:58`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/ray/ray_controller.go#L58) |
| Owns | apps/v1/Deployment | [`internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go:57`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go#L57) |
| Owns | monitoring/v1/PodMonitor | [`internal/controller/components/kueue/kueue_controller.go:69`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/kueue/kueue_controller.go#L69) |
| Owns | monitoring/v1/PrometheusRule | [`internal/controller/components/kueue/kueue_controller.go:70`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/kueue/kueue_controller.go#L70) |
| Owns | monitoring/v1/ServiceMonitor | [`internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go:56`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go#L56) |
| Owns | networking.k8s.io/v1/NetworkPolicy | [`internal/controller/components/kueue/kueue_controller.go:68`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/kueue/kueue_controller.go#L68) |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | [`internal/controller/services/auth/auth_controller.go:60`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/services/auth/auth_controller.go#L60) |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | [`internal/controller/components/trustyai/trustyai_controller.go:53`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/trustyai/trustyai_controller.go#L53) |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | [`internal/controller/components/ray/ray_controller.go:53`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/ray/ray_controller.go#L53) |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | [`internal/controller/components/kueue/kueue_controller.go:63`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/kueue/kueue_controller.go#L63) |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | [`internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go:51`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go#L51) |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | [`internal/controller/services/auth/auth_controller.go:59`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/services/auth/auth_controller.go#L59) |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | [`internal/controller/components/kueue/kueue_controller.go:62`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/kueue/kueue_controller.go#L62) |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | [`internal/controller/components/trustyai/trustyai_controller.go:52`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/trustyai/trustyai_controller.go#L52) |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | [`internal/controller/components/ray/ray_controller.go:52`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/ray/ray_controller.go#L52) |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | [`internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go:50`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go#L50) |
| Owns | rbac.authorization.k8s.io/v1/Role | [`internal/controller/components/trustyai/trustyai_controller.go:54`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/trustyai/trustyai_controller.go#L54) |
| Owns | rbac.authorization.k8s.io/v1/Role | [`internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go:52`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go#L52) |
| Owns | rbac.authorization.k8s.io/v1/Role | [`internal/controller/services/auth/auth_controller.go:61`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/services/auth/auth_controller.go#L61) |
| Owns | rbac.authorization.k8s.io/v1/Role | [`internal/controller/components/kueue/kueue_controller.go:64`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/kueue/kueue_controller.go#L64) |
| Owns | rbac.authorization.k8s.io/v1/Role | [`internal/controller/components/ray/ray_controller.go:54`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/ray/ray_controller.go#L54) |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | [`internal/controller/components/kueue/kueue_controller.go:65`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/kueue/kueue_controller.go#L65) |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | [`internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go:53`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go#L53) |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | [`internal/controller/components/trustyai/trustyai_controller.go:55`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/trustyai/trustyai_controller.go#L55) |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | [`internal/controller/components/ray/ray_controller.go:55`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/ray/ray_controller.go#L55) |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | [`internal/controller/services/auth/auth_controller.go:62`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/services/auth/auth_controller.go#L62) |
| Owns | security/v1/SecurityContextConstraints | [`internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go:58`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go#L58) |
| Owns | security/v1/SecurityContextConstraints | [`internal/controller/components/ray/ray_controller.go:59`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/ray/ray_controller.go#L59) |
| Watches | /v1/Namespace | [`internal/controller/components/kueue/kueue_controller.go:140`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/kueue/kueue_controller.go#L140) |
| Watches | rbac.authorization.k8s.io/v1/ClusterRole | [`internal/controller/components/kueue/kueue_controller.go:134`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/kueue/kueue_controller.go#L134) |
| Watches | services/v1alpha1/Auth | [`internal/controller/components/kueue/kueue_controller.go:151`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/internal/controller/components/kueue/kueue_controller.go#L151) |

### Programmatic Resource Operations

| Verb | Kind | Group | Condition |
|------|------|-------|----------|
| update | DSCInitialization | dscinitialization |  |
| delete | FeatureTracker | features |  |

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for opendatahub-operator

    participant KubernetesAPI as Kubernetes API
    participant aws_cloud_manager_operator as aws-cloud-manager-operator
    participant azure_cloud_manager_operator as azure-cloud-manager-operator
    participant controller_manager as controller-manager
    participant coreweave_cloud_manager_operator as coreweave-cloud-manager-operator
    participant rhods_operator as rhods-operator

    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ConfigMap
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ConfigMap
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ConfigMap
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ConfigMap
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update Secret
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update Secret
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update Secret
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update Service
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update Service
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update Service
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update Service
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ServiceAccount
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ServiceAccount
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ServiceAccount
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ServiceAccount
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update MutatingWebhookConfiguration
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ValidatingWebhookConfiguration
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update Deployment
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update Deployment
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update Deployment
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update Deployment
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update PodMonitor
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update PrometheusRule
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ServiceMonitor
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update NetworkPolicy
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ClusterRole
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ClusterRole
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ClusterRole
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ClusterRole
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ClusterRole
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ClusterRoleBinding
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ClusterRoleBinding
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ClusterRoleBinding
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ClusterRoleBinding
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update ClusterRoleBinding
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update Role
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update Role
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update Role
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update Role
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update Role
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update RoleBinding
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update RoleBinding
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update RoleBinding
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update RoleBinding
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update RoleBinding
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update SecurityContextConstraints
    aws_cloud_manager_operator->>KubernetesAPI: Create/Update SecurityContextConstraints
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
| Defaulter-webhook | mutating | /mutate-datasciencecluster-v2 |  |  |  |  |  |
| Defaulter-webhook | mutating | /mutate-datasciencecluster-v1 |  |  |  |  |  |
| Injector-webhook | mutating | /mutate-prometheus-monitors |  |  |  |  |  |
| Validator-webhook | validating | /validate-datasciencecluster-v2 |  |  |  |  |  |
| Validator-webhook | validating | /validate-dscinitialization-v1 |  |  |  |  |  |
| Validator-webhook | validating | /validate-dscinitialization-v2 |  |  |  |  |  |
| Validator-webhook | validating | /validate-datasciencecluster-v1 |  |  |  |  |  |
| conversion-unknown | conversion | /convert |  | system/webhook-service |  | .ManagementState == operatorv1.Managed &amp;&amp; .ManagementState == "" | [`config/crd/patches/webhook_in_datasciencecluster_datascienceclusters.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/crd/patches/webhook_in_datasciencecluster_datascienceclusters.yaml) |
| conversion-unknown | conversion | /convert |  | system/webhook-service |  | .ManagementState == operatorv1.Managed &amp;&amp; .ManagementState == "" | [`config/crd/patches/webhook_in_dscinitialization_dscinitializations.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/crd/patches/webhook_in_dscinitialization_dscinitializations.yaml) |
| conversion-unknown | conversion | /convert |  | system/webhook-service |  | .ManagementState == operatorv1.Managed &amp;&amp; .ManagementState == "" | [`config/crd/patches/webhook_in_services_auths.yaml`](https://github.com/opendatahub-io/opendatahub-operator/blob/a51f8abe6d56585efe2082a3bcc90351f9b4eefc/config/crd/patches/webhook_in_services_auths.yaml) |

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

