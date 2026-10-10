# models-as-a-service: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

| Type | GVK | Source |
|------|-----|--------|
| For | apps/v1/Deployment | [`maas-controller/pkg/controller/maas/self_deployment_controller.go:1015`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/self_deployment_controller.go#L1015) |
| For | maas/v1alpha1/AITenant | [`maas-controller/pkg/controller/maas/aitenant_controller.go:290`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/aitenant_controller.go#L290) |
| For | maas/v1alpha1/ExternalModel | [`maas-controller/pkg/reconciler/externalmodel/reconciler.go:439`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/reconciler/externalmodel/reconciler.go#L439) |
| For | maas/v1alpha1/MaaSAuthPolicy | [`maas-controller/pkg/controller/maas/maasauthpolicy_controller.go:2061`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/maasauthpolicy_controller.go#L2061) |
| For | maas/v1alpha1/MaaSModelRef | [`maas-controller/pkg/controller/maas/maasmodelref_controller.go:671`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/maasmodelref_controller.go#L671) |
| For | maas/v1alpha1/MaaSSubscription | [`maas-controller/pkg/controller/maas/maassubscription_controller.go:1246`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/maassubscription_controller.go#L1246) |
| For | maas/v1alpha1/MaasTenantConfig | [`maas-controller/pkg/controller/maas/tenant_controller.go:313`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/tenant_controller.go#L313) |
| Watches | /v1/Namespace | [`maas-controller/pkg/controller/maas/maasauthpolicy_controller.go:2105`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/maasauthpolicy_controller.go#L2105) |
| Watches | /v1/Namespace | [`maas-controller/pkg/controller/maas/maassubscription_controller.go:1325`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/maassubscription_controller.go#L1325) |
| Watches | apis/v1/HTTPRoute | [`maas-controller/pkg/controller/maas/maasmodelref_controller.go:677`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/maasmodelref_controller.go#L677) |
| Watches | apis/v1/HTTPRoute | [`maas-controller/pkg/controller/maas/maassubscription_controller.go:1259`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/maassubscription_controller.go#L1259) |
| Watches | apis/v1/HTTPRoute | [`maas-controller/pkg/controller/maas/maasauthpolicy_controller.go:2067`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/maasauthpolicy_controller.go#L2067) |
| Watches | apix/v1alpha2/InferenceObjective | [`maas-controller/pkg/controller/maas/maassubscription_controller.go:1297`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/maassubscription_controller.go#L1297) |
| Watches | maas/v1alpha1/AITenant | [`maas-controller/pkg/controller/maas/maasmodelref_controller.go:744`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/maasmodelref_controller.go#L744) |
| Watches | maas/v1alpha1/AITenant | [`maas-controller/pkg/controller/maas/maassubscription_controller.go:1268`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/maassubscription_controller.go#L1268) |
| Watches | maas/v1alpha1/AITenant | [`maas-controller/pkg/controller/maas/maasauthpolicy_controller.go:2081`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/maasauthpolicy_controller.go#L2081) |
| Watches | maas/v1alpha1/MaaSAuthPolicy | [`maas-controller/pkg/controller/maas/maasmodelref_controller.go:738`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/maasmodelref_controller.go#L738) |
| Watches | maas/v1alpha1/MaaSModelRef | [`maas-controller/pkg/controller/maas/maassubscription_controller.go:1263`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/maassubscription_controller.go#L1263) |
| Watches | maas/v1alpha1/MaaSModelRef | [`maas-controller/pkg/controller/maas/maasmodelref_controller.go:687`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/maasmodelref_controller.go#L687) |
| Watches | maas/v1alpha1/MaaSModelRef | [`maas-controller/pkg/controller/maas/maasauthpolicy_controller.go:2071`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/maasauthpolicy_controller.go#L2071) |
| Watches | maas/v1alpha1/MaaSSubscription | [`maas-controller/pkg/controller/maas/maasmodelref_controller.go:734`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/maasmodelref_controller.go#L734) |
| Watches | maas/v1alpha1/MaasTenantConfig | [`maas-controller/pkg/controller/maas/maassubscription_controller.go:1273`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/maassubscription_controller.go#L1273) |
| Watches | maas/v1alpha1/Tenant | [`maas-controller/pkg/controller/maas/maassubscription_controller.go:1276`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/maassubscription_controller.go#L1276) |
| Watches | serving/v1alpha2/LLMInferenceService | [`maas-controller/pkg/controller/maas/maassubscription_controller.go:1285`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/maassubscription_controller.go#L1285) |
| Watches | serving/v1alpha2/LLMInferenceService | [`maas-controller/pkg/controller/maas/maasmodelref_controller.go:713`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-controller/pkg/controller/maas/maasmodelref_controller.go#L713) |

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
| vaitenant.kb.io | validating | /validate-maas-opendatahub-io-v1alpha1-aitenant | Fail | system/maas-controller-webhook-service |  |  | [`deployment/base/maas-controller/webhook/validating_webhook_configuration.yaml`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/deployment/base/maas-controller/webhook/validating_webhook_configuration.yaml) |
| vmaasauthpolicy.kb.io | validating | /validate-maas-opendatahub-io-v1alpha1-maasauthpolicy | Fail | system/maas-controller-webhook-service |  |  | [`deployment/base/maas-controller/webhook/validating_webhook_configuration.yaml`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/deployment/base/maas-controller/webhook/validating_webhook_configuration.yaml) |
| vmaasmodelref.kb.io | validating | /validate-maas-opendatahub-io-v1alpha1-maasmodelref | Fail | system/maas-controller-webhook-service |  |  | [`deployment/base/maas-controller/webhook/validating_webhook_configuration.yaml`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/deployment/base/maas-controller/webhook/validating_webhook_configuration.yaml) |
| vmaassubscription.kb.io | validating | /validate-maas-opendatahub-io-v1alpha1-maassubscription | Fail | system/maas-controller-webhook-service |  |  | [`deployment/base/maas-controller/webhook/validating_webhook_configuration.yaml`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/deployment/base/maas-controller/webhook/validating_webhook_configuration.yaml) |

### HTTP Endpoints

| Method | Path | Source |
|--------|------|--------|
| OPTIONS | /*path | [`maas-api/cmd/main.go:171`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-api/cmd/main.go#L171) |
| DELETE | /:id | [`maas-api/cmd/main.go:315`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-api/cmd/main.go#L315) |
| GET | /:id | [`maas-api/cmd/main.go:314`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-api/cmd/main.go#L314) |
| * | /api-keys | [`maas-api/cmd/main.go:310`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-api/cmd/main.go#L310) |
| POST | /api-keys/cleanup | [`maas-api/cmd/main.go:332`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-api/cmd/main.go#L332) |
| POST | /api-keys/search | [`maas-api/cmd/main.go:318`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-api/cmd/main.go#L318) |
| POST | /api-keys/validate | [`maas-api/cmd/main.go:330`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-api/cmd/main.go#L330) |
| POST | /bulk-revoke | [`maas-api/cmd/main.go:313`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-api/cmd/main.go#L313) |
| GET | /config | [`maas-api/cmd/main.go:311`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-api/cmd/main.go#L311) |
| GET | /health | [`maas-api/cmd/main.go:250`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-api/cmd/main.go#L250) |
| GET | /healthz | [`maas-discovery/internal/handler/handler.go:50`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-discovery/internal/handler/handler.go#L50) |
| * | /internal/v1 | [`maas-api/cmd/main.go:329`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-api/cmd/main.go#L329) |
| * | /metrics | [`maas-api/internal/metrics/server.go:86`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-api/internal/metrics/server.go#L86) |
| GET | /model/:model-id/subscriptions | [`maas-api/cmd/main.go:303`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-api/cmd/main.go#L303) |
| GET | /models | [`maas-api/cmd/main.go:298`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-api/cmd/main.go#L298) |
| GET | /readyz | [`maas-discovery/internal/handler/handler.go:51`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-discovery/internal/handler/handler.go#L51) |
| GET | /subscriptions | [`maas-api/cmd/main.go:302`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-api/cmd/main.go#L302) |
| POST | /subscriptions/select | [`maas-api/cmd/main.go:335`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-api/cmd/main.go#L335) |
| GET | /tenants | [`maas-api/cmd/main.go:322`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-api/cmd/main.go#L322) |
| DELETE | /tenants/:tenant/api-keys | [`maas-api/cmd/main.go:333`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-api/cmd/main.go#L333) |
| DELETE | /tenants/:tenant/subscriptions/:subscription/api-keys | [`maas-api/cmd/main.go:334`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-api/cmd/main.go#L334) |
| * | /v1 | [`maas-api/cmd/main.go:258`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-api/cmd/main.go#L258) |
| GET | /v1/tenants | [`maas-discovery/internal/handler/handler.go:49`](https://github.com/opendatahub-io/models-as-a-service/blob/b2c10457e65e8c20cd8aaddddc4a3ef7a9f0454d/maas-discovery/internal/handler/handler.go#L49) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

