# Secrets Inventory

25 secrets referenced across the platform. No secret values are extracted, only names, types, and which component references them.

## Secret Distribution

```mermaid
graph TD
    classDef tls fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef opaque fill:#f39c12,stroke:#e67e22,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff

    data_science_pipelines_operator["data-science-pipelines-operator\n4 secrets"]:::component
    data_science_pipelines_operator --> sec_1["mariadb-certs\nds-pipeline-db-test\nminio-certs\nminio"]:::opaque
    kserve["kserve\n3 secrets"]:::component
    kserve --> sec_2["llmisvc-webhook-server-cert\nlocalmodel-webhook-server-cert\nkserve-webhook-server-cert"]:::opaque
    kube_auth_proxy["kube-auth-proxy\n2 secrets"]:::component
    kube_auth_proxy --> sec_3["kube-auth-proxy-secret\nkube-rbac-proxy-client-certificates"]:::opaque
    kuberay["kuberay\n1 secrets"]:::component
    kuberay --> sec_4["webhook-server-cert"]:::opaque
    model_registry_operator["model-registry-operator\n2 secrets"]:::component
    model_registry_operator --> sec_5["webhook-server-cert\ncontroller-manager-metrics-service"]:::opaque
    odh_dashboard["odh-dashboard\n2 secrets"]:::component
    odh_dashboard --> sec_6["dashboard-proxy-tls"]:::tls
    odh_dashboard --> sec_7["webhook-server-cert"]:::opaque
    odh_model_controller["odh-model-controller\n1 secrets"]:::component
    odh_model_controller --> sec_8["odh-model-controller-webhook-cert"]:::tls
    opendatahub_operator["opendatahub-operator\n10 secrets"]:::component
    opendatahub_operator --> sec_9["odh-model-controller-webhook-cert\nodh-notebook-controller-webhook-cert\nopendatahub-operator-controller-webhook-cert\nredhat-ods-operator-controller-webhook-cert\ndashboard-proxy-tls"]:::tls
    opendatahub_operator --> sec_10["webhook-server-cert\ncontroller-manager-metrics-service\nkserve-webhook-server-cert\nkubeflow-training-operator-webhook-cert\ntraining-operator-webhook-cert"]:::opaque
```

## Secrets by Component

| Owner | Secret | Type |
|-------|--------|------|
| data-science-pipelines-operator | mariadb-certs | Opaque |
| data-science-pipelines-operator | ds-pipeline-db-test | Opaque |
| data-science-pipelines-operator | minio-certs | Opaque |
| data-science-pipelines-operator | minio | Opaque |
| kserve | llmisvc-webhook-server-cert | Opaque |
| kserve | localmodel-webhook-server-cert | Opaque |
| kserve | kserve-webhook-server-cert | Opaque |
| kube-auth-proxy | kube-auth-proxy-secret | Opaque |
| kube-auth-proxy | kube-rbac-proxy-client-certificates | Opaque |
| kuberay | webhook-server-cert | Opaque |
| model-registry-operator | webhook-server-cert | Opaque |
| model-registry-operator | controller-manager-metrics-service | Opaque |
| odh-dashboard | dashboard-proxy-tls | kubernetes.io/tls |
| odh-dashboard | webhook-server-cert | Opaque |
| odh-model-controller | odh-model-controller-webhook-cert | kubernetes.io/tls |
| opendatahub-operator | odh-model-controller-webhook-cert | kubernetes.io/tls |
| opendatahub-operator | webhook-server-cert | Opaque |
| opendatahub-operator | controller-manager-metrics-service | Opaque |
| opendatahub-operator | odh-notebook-controller-webhook-cert | kubernetes.io/tls |
| opendatahub-operator | opendatahub-operator-controller-webhook-cert | kubernetes.io/tls |
| opendatahub-operator | redhat-ods-operator-controller-webhook-cert | kubernetes.io/tls |
| opendatahub-operator | kserve-webhook-server-cert | Opaque |
| opendatahub-operator | kubeflow-training-operator-webhook-cert | Opaque |
| opendatahub-operator | dashboard-proxy-tls | kubernetes.io/tls |
| opendatahub-operator | training-operator-webhook-cert | Opaque |

## Patterns

- **Webhook certs** are the dominant secret type (18 of 25 secrets).
- **kubernetes.io/tls** secrets (7) are used for TLS-terminated services.

