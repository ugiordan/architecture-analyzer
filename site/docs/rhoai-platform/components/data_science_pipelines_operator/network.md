# data-science-pipelines-operator: Network

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff

    data_science_pipelines_operator["data-science-pipelines-operator"]:::component
    data_science_pipelines_operator --> svc_0["mariadb\nClusterIP: 3306/TCP"]:::svc
    data_science_pipelines_operator --> svc_1["minio\nClusterIP: 9000/TCP, 9001/TCP"]:::svc
    data_science_pipelines_operator --> svc_2["pypi-server\nClusterIP: 8080/TCP"]:::svc
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| mariadb | ClusterIP | 3306/TCP | `.github/resources/mariadb/service.yaml` |
| minio | ClusterIP | 9000/TCP, 9001/TCP | `.github/resources/minio/service.yaml` |
| pypi-server | ClusterIP | 8080/TCP | `.github/resources/pypiserver/base/service.yaml` |

!!! warning "No Network Policies"
    No NetworkPolicy resources found. All pod-to-pod traffic is allowed by default.

