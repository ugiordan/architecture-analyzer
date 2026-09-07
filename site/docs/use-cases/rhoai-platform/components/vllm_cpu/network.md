# vllm-cpu: Network

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    vllm_cpu["vllm-cpu"]:::component
    vllm_cpu --> svc_0["cli-port-default\npython-source: 8000/TCP"]:::svc
    vllm_cpu -.-> ext_aiohttp[["aiohttp\napi"]]:::ext
    vllm_cpu -.-> ext_httpx[["httpx\napi"]]:::ext
    vllm_cpu -.-> ext_openai[["openai\napi"]]:::ext
    vllm_cpu -.-> ext_requests[["requests\napi"]]:::ext
    vllm_cpu -.-> ext_urllib[["urllib\napi"]]:::ext
    vllm_cpu -.-> ext_sqlalchemy[["sqlalchemy\ndatabase"]]:::ext
    vllm_cpu -.-> ext_s3[["s3\nobject-storage"]]:::ext
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| cli-port-default | python-source | 8000/TCP | [`benchmarks/benchmark_serving_structured_output.py:905`](https://github.com/red-hat-data-services/vllm-cpu/blob/a7f683ee6b25b07450044e5dd324a61163da3a9a/benchmarks/benchmark_serving_structured_output.py#L905) |

!!! warning "No Network Policies"
    No NetworkPolicy resources were found in the analyzed sources. Network policies may exist in overlays, Helm values, or cluster-level configurations not captured by static analysis.

