# notebooks: Network

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff

    notebooks["notebooks"]:::component
    notebooks --> svc_0["notebook\nClusterIP: 8888/TCP"]:::svc
    notebooks --> svc_1["notebook\nClusterIP: 8888/TCP"]:::svc
    notebooks --> svc_2["notebook\nClusterIP: 8888/TCP"]:::svc
    notebooks --> svc_3["notebook\nClusterIP: 8888/TCP"]:::svc
    notebooks --> svc_4["notebook\nClusterIP: 8888/TCP"]:::svc
    notebooks --> svc_5["notebook\nClusterIP: 8888/TCP"]:::svc
    notebooks --> svc_6["notebook\nClusterIP: 8888/TCP"]:::svc
    notebooks --> svc_7["notebook\nClusterIP: 8888/TCP"]:::svc
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| notebook | ClusterIP | 8888/TCP | `jupyter/datascience/ubi9-python-3.12/kustomize/base/service.yaml` |
| notebook | ClusterIP | 8888/TCP | `jupyter/minimal/ubi9-python-3.12/kustomize/base/service.yaml` |
| notebook | ClusterIP | 8888/TCP | `jupyter/pytorch/ubi9-python-3.12/kustomize/base/service.yaml` |
| notebook | ClusterIP | 8888/TCP | `jupyter/pytorch+llmcompressor/ubi9-python-3.12/kustomize/base/service.yaml` |
| notebook | ClusterIP | 8888/TCP | `jupyter/rocm/pytorch/ubi9-python-3.12/kustomize/base/service.yaml` |
| notebook | ClusterIP | 8888/TCP | `jupyter/rocm/tensorflow/ubi9-python-3.12/kustomize/base/service.yaml` |
| notebook | ClusterIP | 8888/TCP | `jupyter/tensorflow/ubi9-python-3.12/kustomize/base/service.yaml` |
| notebook | ClusterIP | 8888/TCP | `jupyter/trustyai/ubi9-python-3.12/kustomize/base/service.yaml` |

!!! warning "No Network Policies"
    No NetworkPolicy resources found. All pod-to-pod traffic is allowed by default.

