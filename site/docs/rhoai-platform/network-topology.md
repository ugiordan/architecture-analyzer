# Network Topology

29 Kubernetes services across the platform.

## Service Map

Services grouped by owning component. Dashed arrows show cross-component references (component A deploys or references a service defined by component B).

```mermaid
graph LR
    classDef webhook fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef metrics fill:#f39c12,stroke:#e67e22,color:#fff
    classDef data fill:#2ecc71,stroke:#27ae60,color:#fff

    subgraph data_science_pipelines_operator_sub["data-science-pipelines-operator"]
        svc_1["mariadb\n3306"]:::data
        svc_2["minio\n9000,9001"]:::data
        svc_3["pypi-server\n8080"]:::data
    end
    subgraph kserve_sub["kserve"]
        svc_4["llmisvc-controller-manager-service\n8443"]:::metrics
        svc_5["kserve-controller-manager-service\n8443"]:::metrics
        svc_6["llmisvc-webhook-server-service\n443"]:::webhook
        svc_7["localmodel-webhook-server-service\n443"]:::webhook
        svc_8["kserve-webhook-server-service\n443"]:::webhook
        svc_9["webhook-service\n443"]:::webhook
    end
    subgraph kube_auth_proxy_sub["kube-auth-proxy"]
        svc_10["kube-rbac-proxy\n8443"]:::metrics
    end
    subgraph kube_rbac_proxy_sub["kube-rbac-proxy"]
        svc_11["kube-rbac-proxy\n8443"]:::metrics
    end
    subgraph kuberay_sub["kuberay"]
        svc_12["kuberay-operator\n8080"]:::data
        svc_13["webhook-service\n443"]:::webhook
    end
    subgraph model_registry_operator_sub["model-registry-operator"]
        svc_14["webhook-service\n443"]:::webhook
    end
    subgraph notebooks_sub["notebooks"]
        svc_15["notebook\n8888"]:::data
    end
    subgraph odh_dashboard_sub["odh-dashboard"]
        svc_16["odh-dashboard\n8443"]:::metrics
        svc_17["workspaces-backend\n4000"]:::data
        svc_18["workspaces-webhook-service\n443"]:::webhook
        svc_19["workspaces-controller-metrics-service\n8080"]:::data
        svc_20["workspaces-frontend\n8080"]:::data
    end
    subgraph odh_model_controller_sub["odh-model-controller"]
        svc_21["odh-model-controller-webhook-service\n443"]:::webhook
    end
    subgraph opendatahub_operator_sub["opendatahub-operator"]
        svc_22["webhook-service\n443"]:::webhook
        svc_23["odh-dashboard\n8443"]:::metrics
        svc_24["kserve-controller-manager-service\n8443"]:::metrics
        svc_25["kserve-webhook-server-service\n443"]:::webhook
        svc_26["odh-model-controller-webhook-service\n443"]:::webhook
        svc_27["kuberay-operator\n8080"]:::data
        svc_28["training-operator\n8080,443"]:::webhook
        svc_29["service\n443"]:::webhook
    end

    %% Cross-component service references
    svc_24 -.->|"refs kserve-controller-manager-service"| svc_5
    svc_23 -.->|"refs odh-dashboard"| svc_16
    svc_26 -.->|"refs odh-model-controller-webhook-service"| svc_21
    svc_25 -.->|"refs kserve-webhook-server-service"| svc_8
    svc_27 -.->|"refs kuberay-operator"| svc_12
```

## Services by Component

| Owner | Service | Ports |
|-------|---------|-------|
| data-science-pipelines-operator | mariadb | 3306/TCP |
| data-science-pipelines-operator | minio | 9000/TCP, 9001/TCP |
| data-science-pipelines-operator | pypi-server | 8080/TCP |
| kserve | llmisvc-controller-manager-service | 8443/TCP |
| kserve | kserve-controller-manager-service | 8443/TCP |
| kserve | llmisvc-webhook-server-service | 443/TCP |
| kserve | localmodel-webhook-server-service | 443/TCP |
| kserve | kserve-webhook-server-service | 443/TCP |
| kserve | webhook-service | 443/TCP |
| kube-auth-proxy | kube-rbac-proxy | 8443/TCP |
| kube-rbac-proxy | kube-rbac-proxy | 8443/TCP |
| kuberay | kuberay-operator | 8080/TCP |
| kuberay | webhook-service | 443/TCP |
| model-registry-operator | webhook-service | 443/TCP |
| notebooks | notebook | 8888/TCP |
| odh-dashboard | odh-dashboard | 8443/TCP |
| odh-dashboard | workspaces-backend | 4000/TCP |
| odh-dashboard | workspaces-webhook-service | 443/TCP |
| odh-dashboard | workspaces-controller-metrics-service | 8080/TCP |
| odh-dashboard | workspaces-frontend | 8080/TCP |
| odh-model-controller | odh-model-controller-webhook-service | 443/TCP |
| opendatahub-operator | webhook-service | 443/TCP |
| opendatahub-operator | odh-dashboard | 8443/TCP |
| opendatahub-operator | kserve-controller-manager-service | 8443/TCP |
| opendatahub-operator | kserve-webhook-server-service | 443/TCP |
| opendatahub-operator | odh-model-controller-webhook-service | 443/TCP |
| opendatahub-operator | kuberay-operator | 8080/TCP |
| opendatahub-operator | training-operator | 8080/TCP, 443/TCP |
| opendatahub-operator | service | 443/TCP |

## Port Patterns

- **3306/TCP**: mariadb
- **4000/TCP**: workspaces-backend
- **443/TCP**: llmisvc-webhook-server-service, localmodel-webhook-server-service, kserve-webhook-server-service, webhook-service, webhook-service, webhook-service, workspaces-webhook-service, odh-model-controller-webhook-service, webhook-service, kserve-webhook-server-service, odh-model-controller-webhook-service, training-operator, service
- **8080/TCP**: pypi-server, kuberay-operator, workspaces-controller-metrics-service, workspaces-frontend, kuberay-operator, training-operator
- **8443/TCP**: llmisvc-controller-manager-service, kserve-controller-manager-service, kube-rbac-proxy, kube-rbac-proxy, odh-dashboard, odh-dashboard, kserve-controller-manager-service
- **8888/TCP**: notebook
- **9000/TCP**: minio
- **9001/TCP**: minio

