# mlflow: Network

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    mlflow["mlflow"]:::component
    mlflow --> svc_0["env-port-default\npython-source: 9137/TCP"]:::svc
    mlflow -.-> ext_aiohttp[["aiohttp\napi"]]:::ext
    mlflow -.-> ext_httpx[["httpx\napi"]]:::ext
    mlflow -.-> ext_mlflow[["mlflow\napi"]]:::ext
    mlflow -.-> ext_openai[["openai\napi"]]:::ext
    mlflow -.-> ext_requests[["requests\napi"]]:::ext
    mlflow -.-> ext_urllib[["urllib\napi"]]:::ext
    mlflow -.-> ext_postgres[["postgres\ndatabase"]]:::ext
    mlflow -.-> ext_redis[["redis\ndatabase"]]:::ext
    mlflow -.-> ext_sqlalchemy[["sqlalchemy\ndatabase"]]:::ext
    mlflow -.-> ext_azure_blob[["azure-blob\nobject-storage"]]:::ext
    mlflow -.-> ext_gcs[["gcs\nobject-storage"]]:::ext
    mlflow -.-> ext_s3[["s3\nobject-storage"]]:::ext
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| env-port-default | python-source | 9137/TCP | [`dev/benchmarks/gateway/fake_server.py:66`](https://github.com/opendatahub-io/mlflow/blob/218c92d73cd0dc5d06cc6604f97ed92d22fc9591/dev/benchmarks/gateway/fake_server.py#L66) |

!!! warning "No Network Policies"
    No NetworkPolicy resources were found in the analyzed sources. Network policies may exist in overlays, Helm values, or cluster-level configurations not captured by static analysis.

