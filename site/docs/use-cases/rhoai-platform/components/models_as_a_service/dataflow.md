# models-as-a-service: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

| Type | GVK | Source |
|------|-----|--------|
| For | apps/v1/Deployment | [`maas-controller/pkg/controller/maas/self_deployment_controller.go:1080`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/self_deployment_controller.go#L1080) |
| For | maas/v1alpha1/AITenant | [`maas-controller/pkg/controller/maas/aitenant_controller.go:290`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/aitenant_controller.go#L290) |
| For | maas/v1alpha1/ExternalModel | [`maas-controller/pkg/reconciler/externalmodel/reconciler.go:439`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/reconciler/externalmodel/reconciler.go#L439) |
| For | maas/v1alpha1/MaaSAuthPolicy | [`maas-controller/pkg/controller/maas/maasauthpolicy_controller.go:2128`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/maasauthpolicy_controller.go#L2128) |
| For | maas/v1alpha1/MaaSModelRef | [`maas-controller/pkg/controller/maas/maasmodelref_controller.go:710`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/maasmodelref_controller.go#L710) |
| For | maas/v1alpha1/MaaSSubscription | [`maas-controller/pkg/controller/maas/maassubscription_controller.go:1414`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/maassubscription_controller.go#L1414) |
| For | maas/v1alpha1/MaasTenantConfig | [`maas-controller/pkg/controller/maas/tenant_controller.go:320`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/tenant_controller.go#L320) |
| Watches | /v1/Namespace | [`maas-controller/pkg/controller/maas/maasauthpolicy_controller.go:2172`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/maasauthpolicy_controller.go#L2172) |
| Watches | /v1/Namespace | [`maas-controller/pkg/controller/maas/maassubscription_controller.go:1493`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/maassubscription_controller.go#L1493) |
| Watches | apis/v1/HTTPRoute | [`maas-controller/pkg/controller/maas/maasmodelref_controller.go:716`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/maasmodelref_controller.go#L716) |
| Watches | apis/v1/HTTPRoute | [`maas-controller/pkg/controller/maas/maassubscription_controller.go:1427`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/maassubscription_controller.go#L1427) |
| Watches | apis/v1/HTTPRoute | [`maas-controller/pkg/controller/maas/maasauthpolicy_controller.go:2134`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/maasauthpolicy_controller.go#L2134) |
| Watches | apix/v1alpha2/InferenceObjective | [`maas-controller/pkg/controller/maas/maassubscription_controller.go:1465`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/maassubscription_controller.go#L1465) |
| Watches | maas/v1alpha1/AITenant | [`maas-controller/pkg/controller/maas/maasmodelref_controller.go:783`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/maasmodelref_controller.go#L783) |
| Watches | maas/v1alpha1/AITenant | [`maas-controller/pkg/controller/maas/maassubscription_controller.go:1436`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/maassubscription_controller.go#L1436) |
| Watches | maas/v1alpha1/AITenant | [`maas-controller/pkg/controller/maas/maasauthpolicy_controller.go:2148`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/maasauthpolicy_controller.go#L2148) |
| Watches | maas/v1alpha1/MaaSAuthPolicy | [`maas-controller/pkg/controller/maas/maasmodelref_controller.go:777`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/maasmodelref_controller.go#L777) |
| Watches | maas/v1alpha1/MaaSModelRef | [`maas-controller/pkg/controller/maas/maassubscription_controller.go:1431`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/maassubscription_controller.go#L1431) |
| Watches | maas/v1alpha1/MaaSModelRef | [`maas-controller/pkg/controller/maas/maasmodelref_controller.go:726`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/maasmodelref_controller.go#L726) |
| Watches | maas/v1alpha1/MaaSModelRef | [`maas-controller/pkg/controller/maas/maasauthpolicy_controller.go:2138`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/maasauthpolicy_controller.go#L2138) |
| Watches | maas/v1alpha1/MaaSSubscription | [`maas-controller/pkg/controller/maas/maasmodelref_controller.go:773`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/maasmodelref_controller.go#L773) |
| Watches | maas/v1alpha1/MaasTenantConfig | [`maas-controller/pkg/controller/maas/maassubscription_controller.go:1441`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/maassubscription_controller.go#L1441) |
| Watches | maas/v1alpha1/Tenant | [`maas-controller/pkg/controller/maas/maassubscription_controller.go:1444`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/maassubscription_controller.go#L1444) |
| Watches | serving/v1alpha2/LLMInferenceService | [`maas-controller/pkg/controller/maas/maassubscription_controller.go:1453`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/maassubscription_controller.go#L1453) |
| Watches | serving/v1alpha2/LLMInferenceService | [`maas-controller/pkg/controller/maas/maasmodelref_controller.go:752`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-controller/pkg/controller/maas/maasmodelref_controller.go#L752) |

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for models-as-a-service

    participant KubernetesAPI as Kubernetes API
    participant maas_api as maas-api
    participant maas_controller as maas-controller
    participant maas_discovery as maas-discovery
    participant payload_processing as payload-processing

    KubernetesAPI->>+maas_api: Watch Deployment (reconcile)
    KubernetesAPI->>+maas_api: Watch AITenant (reconcile)
    KubernetesAPI->>+maas_api: Watch ExternalModel (reconcile)
    KubernetesAPI->>+maas_api: Watch MaaSAuthPolicy (reconcile)
    KubernetesAPI->>+maas_api: Watch MaaSModelRef (reconcile)
    KubernetesAPI->>+maas_api: Watch MaaSSubscription (reconcile)
    KubernetesAPI->>+maas_api: Watch MaasTenantConfig (reconcile)
    KubernetesAPI-->>+maas_api: Watch Namespace (informer)
    KubernetesAPI-->>+maas_api: Watch Namespace (informer)
    KubernetesAPI-->>+maas_api: Watch HTTPRoute (informer)
    KubernetesAPI-->>+maas_api: Watch HTTPRoute (informer)
    KubernetesAPI-->>+maas_api: Watch HTTPRoute (informer)
    KubernetesAPI-->>+maas_api: Watch InferenceObjective (informer)
    KubernetesAPI-->>+maas_api: Watch AITenant (informer)
    KubernetesAPI-->>+maas_api: Watch AITenant (informer)
    KubernetesAPI-->>+maas_api: Watch AITenant (informer)
    KubernetesAPI-->>+maas_api: Watch MaaSAuthPolicy (informer)
    KubernetesAPI-->>+maas_api: Watch MaaSModelRef (informer)
    KubernetesAPI-->>+maas_api: Watch MaaSModelRef (informer)
    KubernetesAPI-->>+maas_api: Watch MaaSModelRef (informer)
    KubernetesAPI-->>+maas_api: Watch MaaSSubscription (informer)
    KubernetesAPI-->>+maas_api: Watch MaasTenantConfig (informer)
    KubernetesAPI-->>+maas_api: Watch Tenant (informer)
    KubernetesAPI-->>+maas_api: Watch LLMInferenceService (informer)
    KubernetesAPI-->>+maas_api: Watch LLMInferenceService (informer)

    Note over maas_api: Exposed Services
    Note right of maas_api: maas-api:8080/TCP [http]
    Note right of maas_api: maas-api:0/TCP []
    Note right of maas_api: maas-api:8443/TCP [https]
    Note right of maas_api: maas-controller-webhook-service:443/TCP []
    Note right of maas_api: maas-discovery:8443/TCP [https]
    Note right of maas_api: payload-processing:9004/TCP []
```

### Webhooks

| Name | Type | Path | Failure Policy | Service | Overlays | Enable Condition | Sources |
|------|------|------|----------------|---------|----------|------------------|----------|
| vaitenant.kb.io | validating | /validate-maas-opendatahub-io-v1alpha1-aitenant | Fail | system/maas-controller-webhook-service |  |  | [`deployment/base/maas-controller/webhook/validating_webhook_configuration.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/deployment/base/maas-controller/webhook/validating_webhook_configuration.yaml) |
| vexternalmodel.kb.io | validating | /validate-maas-opendatahub-io-v1alpha1-externalmodel | Fail | system/maas-controller-webhook-service |  |  | [`deployment/base/maas-controller/webhook/validating_webhook_configuration.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/deployment/base/maas-controller/webhook/validating_webhook_configuration.yaml) |
| vmaasauthpolicy.kb.io | validating | /validate-maas-opendatahub-io-v1alpha1-maasauthpolicy | Fail | system/maas-controller-webhook-service |  |  | [`deployment/base/maas-controller/webhook/validating_webhook_configuration.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/deployment/base/maas-controller/webhook/validating_webhook_configuration.yaml) |
| vmaasmodelref.kb.io | validating | /validate-maas-opendatahub-io-v1alpha1-maasmodelref | Fail | system/maas-controller-webhook-service |  |  | [`deployment/base/maas-controller/webhook/validating_webhook_configuration.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/deployment/base/maas-controller/webhook/validating_webhook_configuration.yaml) |
| vmaassubscription.kb.io | validating | /validate-maas-opendatahub-io-v1alpha1-maassubscription | Fail | system/maas-controller-webhook-service |  |  | [`deployment/base/maas-controller/webhook/validating_webhook_configuration.yaml`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/deployment/base/maas-controller/webhook/validating_webhook_configuration.yaml) |

### HTTP Endpoints

| Method | Path | Source |
|--------|------|--------|
| OPTIONS | /*path | [`maas-api/cmd/main.go:171`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-api/cmd/main.go#L171) |
| DELETE | /:id | [`maas-api/cmd/main.go:349`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-api/cmd/main.go#L349) |
| GET | /:id | [`maas-api/cmd/main.go:348`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-api/cmd/main.go#L348) |
| * | /api-keys | [`maas-api/cmd/main.go:344`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-api/cmd/main.go#L344) |
| POST | /api-keys/cleanup | [`maas-api/cmd/main.go:367`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-api/cmd/main.go#L367) |
| POST | /api-keys/search | [`maas-api/cmd/main.go:352`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-api/cmd/main.go#L352) |
| POST | /api-keys/validate | [`maas-api/cmd/main.go:365`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-api/cmd/main.go#L365) |
| POST | /bulk-revoke | [`maas-api/cmd/main.go:347`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-api/cmd/main.go#L347) |
| GET | /config | [`maas-api/cmd/main.go:345`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-api/cmd/main.go#L345) |
| GET | /health | [`maas-api/cmd/main.go:250`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-api/cmd/main.go#L250) |
| GET | /healthz | [`maas-discovery/internal/handler/handler.go:84`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-discovery/internal/handler/handler.go#L84) |
| * | /internal/v1 | [`maas-api/cmd/main.go:364`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-api/cmd/main.go#L364) |
| GET | /me | [`maas-api/cmd/main.go:319`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-api/cmd/main.go#L319) |
| * | /metrics | [`maas-api/internal/metrics/server.go:86`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-api/internal/metrics/server.go#L86) |
| GET | /model/:model-id/subscriptions | [`maas-api/cmd/main.go:333`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-api/cmd/main.go#L333) |
| GET | /models | [`maas-api/cmd/main.go:315`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-api/cmd/main.go#L315) |
| GET | /readyz | [`maas-discovery/internal/handler/handler.go:85`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-discovery/internal/handler/handler.go#L85) |
| GET | /subscriptions | [`maas-api/cmd/main.go:332`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-api/cmd/main.go#L332) |
| POST | /subscriptions/select | [`maas-api/cmd/main.go:370`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-api/cmd/main.go#L370) |
| GET | /tenants | [`maas-api/cmd/main.go:356`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-api/cmd/main.go#L356) |
| DELETE | /tenants/:tenant/api-keys | [`maas-api/cmd/main.go:368`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-api/cmd/main.go#L368) |
| DELETE | /tenants/:tenant/subscriptions/:subscription/api-keys | [`maas-api/cmd/main.go:369`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-api/cmd/main.go#L369) |
| * | /v1 | [`maas-api/cmd/main.go:258`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-api/cmd/main.go#L258) |
| GET | /v1/tenants | [`maas-discovery/internal/handler/handler.go:83`](https://github.com/red-hat-data-services/models-as-a-service/blob/d7a54d5d7b3a0baaa2283ceafc8088ac934ec5a7/maas-discovery/internal/handler/handler.go#L83) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

