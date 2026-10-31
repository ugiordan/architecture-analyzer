# trustyai-service-operator: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

| Type | GVK | Source |
|------|-----|--------|
| For | apis/v1alpha1/TrustyAI | [`trustyai-operator-module/pkg/trustyaimodule/reconciler.go:505`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/5bb9e544b7142cbfff3cc9708298d01b5414b28f/trustyai-operator-module/pkg/trustyaimodule/reconciler.go#L505) |
| For | batch/v1/Job | [`controllers/evalhub/evaluation_job_failure_reconciler.go:210`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/5bb9e544b7142cbfff3cc9708298d01b5414b28f/controllers/evalhub/evaluation_job_failure_reconciler.go#L210) |
| For | evalhub/v1/EvalHub | [`controllers/evalhub/evalhub_controller.go:395`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/5bb9e544b7142cbfff3cc9708298d01b5414b28f/controllers/evalhub/evalhub_controller.go#L395) |
| For | kueue/v1beta1/Workload | [`controllers/evalhub/evaluation_failed_kueue_workloads_reconciler.go:90`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/5bb9e544b7142cbfff3cc9708298d01b5414b28f/controllers/evalhub/evaluation_failed_kueue_workloads_reconciler.go#L90) |
| For | lmes/v1alpha1/LMEvalJob | [`controllers/lmes/lmevaljob_controller.go:351`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/5bb9e544b7142cbfff3cc9708298d01b5414b28f/controllers/lmes/lmevaljob_controller.go#L351) |
| For | nemo_guardrails/v1alpha1/NemoGuardrails | [`controllers/nemo_guardrails/nemoguardrail_controller.go:271`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/5bb9e544b7142cbfff3cc9708298d01b5414b28f/controllers/nemo_guardrails/nemoguardrail_controller.go#L271) |
| For | tas/v1/TrustyAIService | [`controllers/tas/trustyaiservice_controller.go:343`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/5bb9e544b7142cbfff3cc9708298d01b5414b28f/controllers/tas/trustyaiservice_controller.go#L343) |
| Owns | /v1/ConfigMap | [`controllers/evalhub/evalhub_controller.go:399`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/5bb9e544b7142cbfff3cc9708298d01b5414b28f/controllers/evalhub/evalhub_controller.go#L399) |
| Owns | /v1/Service | [`trustyai-operator-module/pkg/trustyaimodule/reconciler.go:507`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/5bb9e544b7142cbfff3cc9708298d01b5414b28f/trustyai-operator-module/pkg/trustyaimodule/reconciler.go#L507) |
| Owns | /v1/Service | [`controllers/evalhub/evalhub_controller.go:398`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/5bb9e544b7142cbfff3cc9708298d01b5414b28f/controllers/evalhub/evalhub_controller.go#L398) |
| Owns | /v1/ServiceAccount | [`trustyai-operator-module/pkg/trustyaimodule/reconciler.go:506`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/5bb9e544b7142cbfff3cc9708298d01b5414b28f/trustyai-operator-module/pkg/trustyaimodule/reconciler.go#L506) |
| Owns | apps/v1/Deployment | [`controllers/tas/trustyaiservice_controller.go:344`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/5bb9e544b7142cbfff3cc9708298d01b5414b28f/controllers/tas/trustyaiservice_controller.go#L344) |
| Owns | apps/v1/Deployment | [`controllers/evalhub/evalhub_controller.go:396`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/5bb9e544b7142cbfff3cc9708298d01b5414b28f/controllers/evalhub/evalhub_controller.go#L396) |
| Owns | apps/v1/Deployment | [`trustyai-operator-module/pkg/trustyaimodule/reconciler.go:508`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/5bb9e544b7142cbfff3cc9708298d01b5414b28f/trustyai-operator-module/pkg/trustyaimodule/reconciler.go#L508) |
| Owns | monitoring/v1/ServiceMonitor | [`controllers/evalhub/evalhub_controller.go:404`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/5bb9e544b7142cbfff3cc9708298d01b5414b28f/controllers/evalhub/evalhub_controller.go#L404) |
| Owns | networking.k8s.io/v1/NetworkPolicy | [`controllers/lmes/lmevaljob_controller.go:357`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/5bb9e544b7142cbfff3cc9708298d01b5414b28f/controllers/lmes/lmevaljob_controller.go#L357) |
| Owns | networking.k8s.io/v1/NetworkPolicy | [`controllers/tas/trustyaiservice_controller.go:345`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/5bb9e544b7142cbfff3cc9708298d01b5414b28f/controllers/tas/trustyaiservice_controller.go#L345) |
| Owns | networking.k8s.io/v1/NetworkPolicy | [`controllers/evalhub/evalhub_controller.go:397`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/5bb9e544b7142cbfff3cc9708298d01b5414b28f/controllers/evalhub/evalhub_controller.go#L397) |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | [`trustyai-operator-module/pkg/trustyaimodule/reconciler.go:509`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/5bb9e544b7142cbfff3cc9708298d01b5414b28f/trustyai-operator-module/pkg/trustyaimodule/reconciler.go#L509) |
| Watches | /v1/ConfigMap | [`controllers/evalhub/evalhub_controller.go:401`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/5bb9e544b7142cbfff3cc9708298d01b5414b28f/controllers/evalhub/evalhub_controller.go#L401) |
| Watches | /v1/Namespace | [`controllers/evalhub/evalhub_controller.go:400`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/5bb9e544b7142cbfff3cc9708298d01b5414b28f/controllers/evalhub/evalhub_controller.go#L400) |
| Watches | /v1/Pod | [`controllers/tas/trustyaiservice_controller.go:350`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/5bb9e544b7142cbfff3cc9708298d01b5414b28f/controllers/tas/trustyaiservice_controller.go#L350) |
| Watches | serving/v1beta1/InferenceService | [`controllers/tas/trustyaiservice_controller.go:356`](https://github.com/red-hat-data-services/trustyai-service-operator/blob/5bb9e544b7142cbfff3cc9708298d01b5414b28f/controllers/tas/trustyaiservice_controller.go#L356) |

### Programmatic Resource Operations

| Verb | Kind | Group | Condition |
|------|------|-------|----------|
| create | ConfigMap |  |  |
| update | ConfigMap |  |  |
| create | Deployment | apps |  |
| update | Deployment | apps |  |
| update | EvalHub | evalhub |  |
| patch | Job | batch |  |
| create | Route | route |  |
| update | Route | route |  |
| create | Service |  |  |
| update | Service |  |  |
| create | ServiceMonitor | monitoring |  |
| update | LMEvalJob | lmes |  |
| patch | Pod |  |  |
| patch | NemoGuardrails | nemo_guardrails |  |
| update | TrustyAIService | tas |  |

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for trustyai-service-operator

    participant KubernetesAPI as Kubernetes API
    participant trustyai_service_operator_controller_manager as trustyai-service-operator-controller-manager

    KubernetesAPI->>+trustyai_service_operator_controller_manager: Watch TrustyAI (reconcile)
    KubernetesAPI->>+trustyai_service_operator_controller_manager: Watch Job (reconcile)
    KubernetesAPI->>+trustyai_service_operator_controller_manager: Watch EvalHub (reconcile)
    KubernetesAPI->>+trustyai_service_operator_controller_manager: Watch Workload (reconcile)
    KubernetesAPI->>+trustyai_service_operator_controller_manager: Watch LMEvalJob (reconcile)
    KubernetesAPI->>+trustyai_service_operator_controller_manager: Watch NemoGuardrails (reconcile)
    KubernetesAPI->>+trustyai_service_operator_controller_manager: Watch TrustyAIService (reconcile)
    trustyai_service_operator_controller_manager->>KubernetesAPI: Create/Update ConfigMap
    trustyai_service_operator_controller_manager->>KubernetesAPI: Create/Update Service
    trustyai_service_operator_controller_manager->>KubernetesAPI: Create/Update Service
    trustyai_service_operator_controller_manager->>KubernetesAPI: Create/Update ServiceAccount
    trustyai_service_operator_controller_manager->>KubernetesAPI: Create/Update Deployment
    trustyai_service_operator_controller_manager->>KubernetesAPI: Create/Update Deployment
    trustyai_service_operator_controller_manager->>KubernetesAPI: Create/Update Deployment
    trustyai_service_operator_controller_manager->>KubernetesAPI: Create/Update ServiceMonitor
    trustyai_service_operator_controller_manager->>KubernetesAPI: Create/Update NetworkPolicy
    trustyai_service_operator_controller_manager->>KubernetesAPI: Create/Update NetworkPolicy
    trustyai_service_operator_controller_manager->>KubernetesAPI: Create/Update NetworkPolicy
    trustyai_service_operator_controller_manager->>KubernetesAPI: Create/Update RoleBinding
    KubernetesAPI-->>+trustyai_service_operator_controller_manager: Watch ConfigMap (informer)
    KubernetesAPI-->>+trustyai_service_operator_controller_manager: Watch Namespace (informer)
    KubernetesAPI-->>+trustyai_service_operator_controller_manager: Watch Pod (informer)
    KubernetesAPI-->>+trustyai_service_operator_controller_manager: Watch InferenceService (informer)

    Note over trustyai_service_operator_controller_manager: Exposed Services
    Note right of trustyai_service_operator_controller_manager: trustyai-service-operator-controller-manager-metrics-service:8443/TCP [https]
    Note right of trustyai_service_operator_controller_manager: trustyai-service-operator-metrics-service:8443/TCP [https]
    Note right of trustyai_service_operator_controller_manager: trustyai-service-operator-webhook-service:443/TCP []

    Note over KubernetesAPI: Defined CRDs
    Note right of KubernetesAPI: EvalHub (trustyai.opendatahub.io/v1)
    Note right of KubernetesAPI: TrustyAIService (trustyai.opendatahub.io/v1)
    Note right of KubernetesAPI: EvalHub (trustyai.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: LMEvalJob (trustyai.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: TrustyAIService (trustyai.opendatahub.io/v1alpha1)
```

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

