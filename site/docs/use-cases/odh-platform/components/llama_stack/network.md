# llama-stack: Network

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    llama_stack["llama-stack"]:::component
    llama_stack --> svc_0["cli-port-default\npython-source: 8081/TCP"]:::svc
    llama_stack -.-> ext_aiohttp[["aiohttp\napi"]]:::ext
    llama_stack -.-> ext_anthropic[["anthropic\napi"]]:::ext
    llama_stack -.-> ext_chromadb[["chromadb\napi"]]:::ext
    llama_stack -.-> ext_cohere[["cohere\napi"]]:::ext
    llama_stack -.-> ext_elasticsearch[["elasticsearch\napi"]]:::ext
    llama_stack -.-> ext_httpx[["httpx\napi"]]:::ext
    llama_stack -.-> ext_milvus[["milvus\napi"]]:::ext
    llama_stack -.-> ext_ogx[["ogx\napi"]]:::ext
    llama_stack -.-> ext_openai[["openai\napi"]]:::ext
    llama_stack -.-> ext_qdrant[["qdrant\napi"]]:::ext
    llama_stack -.-> ext_requests[["requests\napi"]]:::ext
    llama_stack -.-> ext_urllib[["urllib\napi"]]:::ext
    llama_stack -.-> ext_mongodb[["mongodb\ndatabase"]]:::ext
    llama_stack -.-> ext_postgres[["postgres\ndatabase"]]:::ext
    llama_stack -.-> ext_redis[["redis\ndatabase"]]:::ext
    llama_stack -.-> ext_sqlalchemy[["sqlalchemy\ndatabase"]]:::ext
    llama_stack -.-> ext_sqlite[["sqlite\ndatabase"]]:::ext
    llama_stack -.-> ext_s3[["s3\nobject-storage"]]:::ext
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| cli-port-default | python-source | 8081/TCP | [`benchmarking/k8s-benchmark/openai-mock-server.py:191`](https://github.com/opendatahub-io/llama-stack/blob/1acb60086c391f316a50cdb62f812abd2e924403/benchmarking/k8s-benchmark/openai-mock-server.py#L191) |

!!! warning "No Network Policies"
    No NetworkPolicy resources were found in the analyzed sources. Network policies may exist in overlays, Helm values, or cluster-level configurations not captured by static analysis.

