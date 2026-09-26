# llama-stack: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

No controller watches found in analyzed sources.

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for llama-stack

    participant KubernetesAPI as Kubernetes API
    participant llama_stack as llama-stack


    Note over llama_stack: Exposed Services
    Note right of llama_stack: cli-port-default:8081/TCP []
```

### HTTP Endpoints

| Method | Path | Source |
|--------|------|--------|
| GET | /v1/admin/tools | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/batches | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/batches | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/batches/{batch_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/batches/{batch_id}/cancel | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/chat/completions | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/chat/completions | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/chat/completions/{completion_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/chat/completions/{completion_id}/messages | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/completions | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/conversations | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| DELETE | /v1/conversations/{conversation_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/conversations/{conversation_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/conversations/{conversation_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/conversations/{conversation_id}/items | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/conversations/{conversation_id}/items | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| DELETE | /v1/conversations/{conversation_id}/items/{item_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/conversations/{conversation_id}/items/{item_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/embeddings | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/files | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/files | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| DELETE | /v1/files/{file_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/files/{file_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/files/{file_id}/content | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/health | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/inspect/routes | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/messages | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/messages/batches | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/messages/batches | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/messages/batches/{message_batch_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/messages/batches/{message_batch_id}/cancel | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/messages/batches/{message_batch_id}/results | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/messages/count_tokens | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/models | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/models/{model_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/prompts | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/prompts | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| DELETE | /v1/prompts/{prompt_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/prompts/{prompt_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| PUT | /v1/prompts/{prompt_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| PUT | /v1/prompts/{prompt_id}/set-default-version | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/prompts/{prompt_id}/versions | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/providers | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/providers/{provider_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/responses | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/responses | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/responses/compact | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| DELETE | /v1/responses/{response_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/responses/{response_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/responses/{response_id}/cancel | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/responses/{response_id}/input_items | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/vector-io/insert | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/vector-io/query | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/vector_stores | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/vector_stores | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| DELETE | /v1/vector_stores/{vector_store_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/vector_stores/{vector_store_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/vector_stores/{vector_store_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/vector_stores/{vector_store_id}/file_batches | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/vector_stores/{vector_store_id}/file_batches/{batch_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/vector_stores/{vector_store_id}/file_batches/{batch_id}/cancel | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/vector_stores/{vector_store_id}/file_batches/{batch_id}/files | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/vector_stores/{vector_store_id}/files | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/vector_stores/{vector_store_id}/files | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| DELETE | /v1/vector_stores/{vector_store_id}/files/{file_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/vector_stores/{vector_store_id}/files/{file_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/vector_stores/{vector_store_id}/files/{file_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/vector_stores/{vector_store_id}/files/{file_id}/content | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1/vector_stores/{vector_store_id}/search | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1/version | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1alpha/admin/connectors | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1alpha/admin/connectors/{connector_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1alpha/admin/connectors/{connector_id}/tools | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1alpha/admin/connectors/{connector_id}/tools/{tool_name} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1alpha/admin/health | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1alpha/admin/inspect/routes | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1alpha/admin/providers | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1alpha/admin/providers/{provider_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1alpha/admin/version | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1alpha/file-processors/jobs | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1alpha/file-processors/jobs | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1alpha/file-processors/jobs/{job_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1alpha/file-processors/jobs/{job_id}/cancel | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1alpha/file-processors/process | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1alpha/inference/rerank | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1alpha/interactions | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1alpha/skills | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1alpha/skills | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| DELETE | /v1alpha/skills/{skill_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1alpha/skills/{skill_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1alpha/skills/{skill_id} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1alpha/skills/{skill_id}/content | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1alpha/skills/{skill_id}/versions | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| POST | /v1alpha/skills/{skill_id}/versions | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| DELETE | /v1alpha/skills/{skill_id}/versions/{version} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1alpha/skills/{skill_id}/versions/{version} | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |
| GET | /v1alpha/skills/{skill_id}/versions/{version}/content | [`client-sdks/spec/openapi.yml`](https://github.com/opendatahub-io/llama-stack/blob/3459711a3502d7b2ac7d8d3c8802164dfd51c582/client-sdks/spec/openapi.yml) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

