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
| GET | / | [`fern/openapi.yml`](https://github.com/red-hat-data-services/NeMo-Guardrails/blob/0a26145b25bf7984e7f87feb328c6b7bfe1660ce/fern/openapi.yml) |
| GET | /healthz | [`fern/openapi.yml`](https://github.com/red-hat-data-services/NeMo-Guardrails/blob/0a26145b25bf7984e7f87feb328c6b7bfe1660ce/fern/openapi.yml) |
| GET | /v1/challenges | [`fern/openapi.yml`](https://github.com/red-hat-data-services/NeMo-Guardrails/blob/0a26145b25bf7984e7f87feb328c6b7bfe1660ce/fern/openapi.yml) |
| POST | /v1/chat/completions | [`fern/openapi.yml`](https://github.com/red-hat-data-services/NeMo-Guardrails/blob/0a26145b25bf7984e7f87feb328c6b7bfe1660ce/fern/openapi.yml) |
| GET | /v1/health | [`fern/openapi.yml`](https://github.com/red-hat-data-services/NeMo-Guardrails/blob/0a26145b25bf7984e7f87feb328c6b7bfe1660ce/fern/openapi.yml) |
| GET | /v1/models | [`fern/openapi.yml`](https://github.com/red-hat-data-services/NeMo-Guardrails/blob/0a26145b25bf7984e7f87feb328c6b7bfe1660ce/fern/openapi.yml) |
| GET | /v1/rails/configs | [`fern/openapi.yml`](https://github.com/red-hat-data-services/NeMo-Guardrails/blob/0a26145b25bf7984e7f87feb328c6b7bfe1660ce/fern/openapi.yml) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

