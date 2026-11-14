# NeMo-Guardrails: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

No controller watches found in analyzed sources.

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for NeMo-Guardrails

    participant KubernetesAPI as Kubernetes API
    participant NeMo_Guardrails as NeMo-Guardrails
```

### HTTP Endpoints

| Method | Path | Source |
|--------|------|--------|
| GET | / | [`fern/openapi.yml`](https://github.com/red-hat-data-services/NeMo-Guardrails/blob/04e2d18ebb076712d21692357f5f7b71ea47e296/fern/openapi.yml) |
| GET | /healthz | [`fern/openapi.yml`](https://github.com/red-hat-data-services/NeMo-Guardrails/blob/04e2d18ebb076712d21692357f5f7b71ea47e296/fern/openapi.yml) |
| GET | /v1/challenges | [`fern/openapi.yml`](https://github.com/red-hat-data-services/NeMo-Guardrails/blob/04e2d18ebb076712d21692357f5f7b71ea47e296/fern/openapi.yml) |
| POST | /v1/chat/completions | [`fern/openapi.yml`](https://github.com/red-hat-data-services/NeMo-Guardrails/blob/04e2d18ebb076712d21692357f5f7b71ea47e296/fern/openapi.yml) |
| GET | /v1/health | [`fern/openapi.yml`](https://github.com/red-hat-data-services/NeMo-Guardrails/blob/04e2d18ebb076712d21692357f5f7b71ea47e296/fern/openapi.yml) |
| GET | /v1/models | [`fern/openapi.yml`](https://github.com/red-hat-data-services/NeMo-Guardrails/blob/04e2d18ebb076712d21692357f5f7b71ea47e296/fern/openapi.yml) |
| GET | /v1/rails/configs | [`fern/openapi.yml`](https://github.com/red-hat-data-services/NeMo-Guardrails/blob/04e2d18ebb076712d21692357f5f7b71ea47e296/fern/openapi.yml) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

