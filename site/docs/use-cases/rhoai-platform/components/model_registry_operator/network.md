# model-registry-operator: Network

## Service Map

*4 unique services (5 total, duplicates from test fixtures collapsed).*

```mermaid
graph LR
    classDef svc fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef test fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef component fill:#3498db,stroke:#2980b9,color:#fff
    classDef ext fill:#e74c3c,stroke:#c0392b,color:#fff

    model_registry_operator["model-registry-operator"]:::component
    model_registry_operator --> svc_0["catalog-webhook-service\nClusterIP: 443/TCP"]:::svc
    model_registry_operator --> svc_1["model-registry-operator-controller-manager-metrics-service\nClusterIP: 8443/TCP"]:::svc
    model_registry_operator --> svc_2["model-registry-operator-webhook-service\nClusterIP: 443/TCP"]:::svc
    model_registry_operator --> svc_3["template-value-postgres\nClusterIP: 5432/TCP"]:::svc
```

### Services

| Name | Type | Ports | Source |
|------|------|-------|--------|
| catalog-webhook-service | ClusterIP | 443/TCP | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/model-registry-operator/blob/3c3a420bd1b39f7d97013b600af7ca77d2ab55dc/kustomize:config/overlays/odh) |
| model-registry-operator-controller-manager-metrics-service | ClusterIP | 8443/TCP | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/model-registry-operator/blob/3c3a420bd1b39f7d97013b600af7ca77d2ab55dc/kustomize:config/overlays/odh) |
| model-registry-operator-webhook-service | ClusterIP | 443/TCP | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/model-registry-operator/blob/3c3a420bd1b39f7d97013b600af7ca77d2ab55dc/kustomize:config/overlays/odh) |
| template-value-postgres | ClusterIP | 5432/TCP | [`internal/controller/config/templates/catalog/catalog-postgres-service.yaml.tmpl`](https://github.com/red-hat-data-services/model-registry-operator/blob/3c3a420bd1b39f7d97013b600af7ca77d2ab55dc/internal/controller/config/templates/catalog/catalog-postgres-service.yaml.tmpl) |
| template-value-postgres | ClusterIP | 5432/TCP | [`internal/controller/config/templates/postgres-service.yaml.tmpl`](https://github.com/red-hat-data-services/model-registry-operator/blob/3c3a420bd1b39f7d97013b600af7ca77d2ab55dc/internal/controller/config/templates/postgres-service.yaml.tmpl) |

### Ingress / Routing

| Kind | Name | Hosts | Paths | TLS | Source |
|------|------|-------|-------|-----|--------|
| HTTPRoute | model-catalog |  | /catalog/ | no | [`internal/controller/config/templates/catalog/catalog-gateway-httproute.yaml.tmpl`](https://github.com/red-hat-data-services/model-registry-operator/blob/3c3a420bd1b39f7d97013b600af7ca77d2ab55dc/internal/controller/config/templates/catalog/catalog-gateway-httproute.yaml.tmpl) |
| HTTPRoute | model-registry-template-value |  | /model-registry/template-value/ | no | [`internal/controller/config/templates/gateway/gateway-httproute.yaml.tmpl`](https://github.com/red-hat-data-services/model-registry-operator/blob/3c3a420bd1b39f7d97013b600af7ca77d2ab55dc/internal/controller/config/templates/gateway/gateway-httproute.yaml.tmpl) |
| Route | template-value-http |  |  | no | [`internal/controller/config/templates/http-route.yaml.tmpl`](https://github.com/red-hat-data-services/model-registry-operator/blob/3c3a420bd1b39f7d97013b600af7ca77d2ab55dc/internal/controller/config/templates/http-route.yaml.tmpl) |
| Route | template-value-https | template-value.template-value |  | yes | [`internal/controller/config/templates/catalog/catalog-kube-rbac-proxy-https-route.yaml.tmpl`](https://github.com/red-hat-data-services/model-registry-operator/blob/3c3a420bd1b39f7d97013b600af7ca77d2ab55dc/internal/controller/config/templates/catalog/catalog-kube-rbac-proxy-https-route.yaml.tmpl) |
| Route | template-value-https | template-value-rest.template-value |  | yes | [`internal/controller/config/templates/kube-rbac-proxy/kube-rbac-proxy-https-route.yaml.tmpl`](https://github.com/red-hat-data-services/model-registry-operator/blob/3c3a420bd1b39f7d97013b600af7ca77d2ab55dc/internal/controller/config/templates/kube-rbac-proxy/kube-rbac-proxy-https-route.yaml.tmpl) |

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| catalog-controller-manager | Ingress, Egress | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/model-registry-operator/blob/3c3a420bd1b39f7d97013b600af7ca77d2ab55dc/kustomize:config/overlays/odh) |
| model-registry-operator-controller-manager | Ingress, Egress | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/model-registry-operator/blob/3c3a420bd1b39f7d97013b600af7ca77d2ab55dc/kustomize:config/overlays/odh) |
| template-value-egress | Egress | [`internal/controller/config/templates/catalog/catalog-egress-network-policy.yaml.tmpl`](https://github.com/red-hat-data-services/model-registry-operator/blob/3c3a420bd1b39f7d97013b600af7ca77d2ab55dc/internal/controller/config/templates/catalog/catalog-egress-network-policy.yaml.tmpl) |
| template-value-egress | Egress | [`internal/controller/config/templates/egress-network-policy.yaml.tmpl`](https://github.com/red-hat-data-services/model-registry-operator/blob/3c3a420bd1b39f7d97013b600af7ca77d2ab55dc/internal/controller/config/templates/egress-network-policy.yaml.tmpl) |
| template-value-https-route | Ingress | [`internal/controller/config/templates/catalog/catalog-kube-rbac-proxy-network-policy.yaml.tmpl`](https://github.com/red-hat-data-services/model-registry-operator/blob/3c3a420bd1b39f7d97013b600af7ca77d2ab55dc/internal/controller/config/templates/catalog/catalog-kube-rbac-proxy-network-policy.yaml.tmpl) |
| template-value-https-route | Ingress | [`internal/controller/config/templates/ingress-network-policy.yaml.tmpl`](https://github.com/red-hat-data-services/model-registry-operator/blob/3c3a420bd1b39f7d97013b600af7ca77d2ab55dc/internal/controller/config/templates/ingress-network-policy.yaml.tmpl) |
| template-value-postgres | Ingress | [`internal/controller/config/templates/catalog/catalog-postgres-network-policy.yaml.tmpl`](https://github.com/red-hat-data-services/model-registry-operator/blob/3c3a420bd1b39f7d97013b600af7ca77d2ab55dc/internal/controller/config/templates/catalog/catalog-postgres-network-policy.yaml.tmpl) |
| template-value-postgres | Ingress | [`internal/controller/config/templates/postgres-network-policy.yaml.tmpl`](https://github.com/red-hat-data-services/model-registry-operator/blob/3c3a420bd1b39f7d97013b600af7ca77d2ab55dc/internal/controller/config/templates/postgres-network-policy.yaml.tmpl) |
| template-value-postgres-egress | Egress | [`internal/controller/config/templates/catalog/catalog-postgres-egress-network-policy.yaml.tmpl`](https://github.com/red-hat-data-services/model-registry-operator/blob/3c3a420bd1b39f7d97013b600af7ca77d2ab55dc/internal/controller/config/templates/catalog/catalog-postgres-egress-network-policy.yaml.tmpl) |
| template-value-postgres-egress | Egress | [`internal/controller/config/templates/postgres-egress-network-policy.yaml.tmpl`](https://github.com/red-hat-data-services/model-registry-operator/blob/3c3a420bd1b39f7d97013b600af7ca77d2ab55dc/internal/controller/config/templates/postgres-egress-network-policy.yaml.tmpl) |

## Network Policy Graph

Visual representation of NetworkPolicy rules. Ingress rules show what traffic is allowed into pods, egress rules show what traffic is allowed out.

```mermaid
graph LR
    classDef policy fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef pod fill:#3498db,stroke:#2980b9,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff

    model_registry_operator["model-registry-operator\nPods"]:::pod
    np_0_catalog_controller_manager{{"catalog-controller-manager\nIngress, Egress"}}:::policy
    src_0_0_0["ns: policy-group.network.openshift.io/host-network="]:::external
    src_0_0_0 -->|ingress| np_0_catalog_controller_manager
    np_0_catalog_controller_manager --> model_registry_operator
    dst_0_0_all["all destinations"]:::external
    model_registry_operator -->|egress| dst_0_0_all
    np_1_model_registry_operator_controller_manager{{"model-registry-operator-controller-manager\nIngress, Egress"}}:::policy
    src_1_0_0["ns: kubernetes.io/metadata.name=openshift-monitoring"]:::external
    src_1_0_0 -->|ingress| np_1_model_registry_operator_controller_manager
    src_1_0_1["ns: kubernetes.io/metadata.name=openshift-user-workload-monitoring"]:::external
    src_1_0_1 -->|ingress| np_1_model_registry_operator_controller_manager
    src_1_0_2["ns: kubernetes.io/metadata.name=redhat-ods-monitoring"]:::external
    src_1_0_2 -->|ingress| np_1_model_registry_operator_controller_manager
    src_1_0_3["ns: kubernetes.io/metadata.name=opendatahub-monitoring"]:::external
    src_1_0_3 -->|ingress| np_1_model_registry_operator_controller_manager
    src_1_1_0["ns: policy-group.network.openshift.io/host-network="]:::external
    src_1_1_0 -->|ingress| np_1_model_registry_operator_controller_manager
    np_1_model_registry_operator_controller_manager --> model_registry_operator
    dst_1_0_all["all destinations"]:::external
    model_registry_operator -->|egress| dst_1_0_all
    np_2_template_value_egress{{"template-value-egress\nEgress"}}:::policy
    np_2_template_value_egress --> model_registry_operator
    dst_2_0_all["all destinations"]:::external
    model_registry_operator -->|egress| dst_2_0_all
    np_3_template_value_egress{{"template-value-egress\nEgress"}}:::policy
    np_3_template_value_egress --> model_registry_operator
    dst_3_0_all["all destinations"]:::external
    model_registry_operator -->|egress| dst_3_0_all
    np_4_template_value_https_route{{"template-value-https-route\nIngress"}}:::policy
    src_4_0_0["ns: network.openshift.io/policy-group=ingress"]:::external
    src_4_0_0 -->|ingress| np_4_template_value_https_route
    src_4_0_1["ns: kubernetes.io/metadata.name=template-value"]:::external
    src_4_0_1 -->|ingress| np_4_template_value_https_route
    np_4_template_value_https_route --> model_registry_operator
    np_5_template_value_https_route{{"template-value-https-route\nIngress"}}:::policy
    src_5_0_0["ns: network.openshift.io/policy-group=ingress"]:::external
    src_5_0_0 -->|ingress| np_5_template_value_https_route
    src_5_0_1["ns: opendatahub.io/generated-namespace=true"]:::external
    src_5_0_1 -->|ingress| np_5_template_value_https_route
    np_5_template_value_https_route --> model_registry_operator
    np_6_template_value_postgres{{"template-value-postgres\nIngress"}}:::policy
    src_6_0_0["pods: app.kubernetes.io/name=template-value, app=template-value, component=model-catalog"]:::external
    src_6_0_0 -->|ingress| np_6_template_value_postgres
    np_6_template_value_postgres --> model_registry_operator
    np_7_template_value_postgres{{"template-value-postgres\nIngress"}}:::policy
    src_7_0_0["pods: app.kubernetes.io/name=template-value, app=template-value, component=model-registry"]:::external
    src_7_0_0 -->|ingress| np_7_template_value_postgres
    np_7_template_value_postgres --> model_registry_operator
    np_8_template_value_postgres_egress{{"template-value-postgres-egress\nEgress"}}:::policy
    np_8_template_value_postgres_egress --> model_registry_operator
    np_9_template_value_postgres_egress{{"template-value-postgres-egress\nEgress"}}:::policy
    np_9_template_value_postgres_egress --> model_registry_operator
```

