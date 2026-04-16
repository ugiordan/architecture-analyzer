# kube-auth-proxy: Network

## Service Map

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff

    kube_auth_proxy["kube-auth-proxy"]:::component
    kube_auth_proxy --> svc_0["kube-rbac-proxy\nClusterIP: 8443/TCP"]:::svc
    kube_auth_proxy --> svc_1["kube-rbac-proxy\nClusterIP: 8443/TCP"]:::svc
    kube_auth_proxy --> svc_2["kube-rbac-proxy\nClusterIP: 8443/TCP"]:::svc
    kube_auth_proxy --> svc_3["kube-rbac-proxy\nClusterIP: 8443/TCP"]:::svc
    kube_auth_proxy --> svc_4["kube-rbac-proxy\nClusterIP: 8443/TCP"]:::svc
    kube_auth_proxy --> svc_5["kube-rbac-proxy\nClusterIP: 8443/TCP"]:::svc
    kube_auth_proxy --> svc_6["kube-rbac-proxy\nClusterIP: 8443/TCP"]:::svc
    kube_auth_proxy --> svc_7["kube-rbac-proxy\nClusterIP: 8443/TCP"]:::svc
    kube_auth_proxy --> svc_8["kube-rbac-proxy\nClusterIP: 8443/TCP"]:::svc
    kube_auth_proxy --> svc_9["kube-rbac-proxy\nClusterIP: 8443/TCP"]:::svc
    kube_auth_proxy --> svc_10["kube-rbac-proxy\nClusterIP: 8443/TCP"]:::svc
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| kube-rbac-proxy | ClusterIP | 8443/TCP | `kube-rbac-proxy/test/e2e/allowpaths/service.yaml` |
| kube-rbac-proxy | ClusterIP | 8443/TCP | `kube-rbac-proxy/test/e2e/basics/service.yaml` |
| kube-rbac-proxy | ClusterIP | 8443/TCP | `kube-rbac-proxy/test/e2e/clientcertificates/service.yaml` |
| kube-rbac-proxy | ClusterIP | 8443/TCP | `kube-rbac-proxy/test/e2e/flags/service.yaml` |
| kube-rbac-proxy | ClusterIP | 8443/TCP | `kube-rbac-proxy/test/e2e/h2c-upstream/service.yaml` |
| kube-rbac-proxy | ClusterIP | 8443/TCP | `kube-rbac-proxy/test/e2e/hardcoded_authorizer/service.yaml` |
| kube-rbac-proxy | ClusterIP | 8443/TCP | `kube-rbac-proxy/test/e2e/http2/service.yaml` |
| kube-rbac-proxy | ClusterIP | 8443/TCP | `kube-rbac-proxy/test/e2e/ignorepaths/service.yaml` |
| kube-rbac-proxy | ClusterIP | 8443/TCP | `kube-rbac-proxy/test/e2e/static-auth/service.yaml` |
| kube-rbac-proxy | ClusterIP | 8443/TCP | `kube-rbac-proxy/test/e2e/tokenmasking/service.yaml` |
| kube-rbac-proxy | ClusterIP | 8443/TCP | `kube-rbac-proxy/test/e2e/tokenrequest/service.yaml` |

!!! warning "No Network Policies"
    No NetworkPolicy resources found. All pod-to-pod traffic is allowed by default.

