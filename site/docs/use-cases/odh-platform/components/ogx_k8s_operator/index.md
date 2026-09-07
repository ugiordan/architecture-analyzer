# ogx-k8s-operator

> **Architecture snapshot: 2026-09-07** (2026-09-07)


**Repository:** ogx-ai/ogx-k8s-operator  
**Analyzer:** arch-analyzer dev  
**Extracted:** 2026-09-07T03:59:02Z

## Summary

| Metric | Count |
|--------|-------|
| CRDs | 2 |
| Deployments | 3 |
| Services | 2 |
| Secrets | 2 |
| Cluster Roles | 5 |
| Controller Watches | 22 |

## Component Architecture

CRDs, controllers, and owned Kubernetes resources.

```mermaid
graph LR
    %% Component architecture for ogx-k8s-operator

    classDef crd fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef controller fill:#3498db,stroke:#2980b9,color:#fff
    classDef owned fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef dep fill:#f39c12,stroke:#e67e22,color:#fff

    subgraph controller["ogx-k8s-operator Controller"]
        dep_1["deployment"]
        class dep_1 controller
        dep_2["ogx-k8s-operator-controller-manager"]
        class dep_2 controller
        dep_3["operator"]
        class dep_3 controller
    end

    crd_LlamaStackDistribution{{"LlamaStackDistribution\nllamastack.io/v1alpha1"}}
    class crd_LlamaStackDistribution crd
    crd_OGXServer{{"OGXServer\nogx.io/v1beta1"}}
    class crd_OGXServer crd
    crd_OGXServer -->|"For (reconciles)"| controller
    controller -->|"Owns"| owned_4["ClusterRole"]
    class owned_4 owned
    controller -->|"Owns"| owned_5["ClusterRoleBinding"]
    class owned_5 owned
    controller -->|"Owns"| owned_6["ConfigMap"]
    class owned_6 owned
    controller -->|"Owns"| owned_7["CustomResourceDefinition"]
    class owned_7 owned
    controller -->|"Owns"| owned_8["Deployment"]
    class owned_8 owned
    controller -->|"Owns"| owned_9["HorizontalPodAutoscaler"]
    class owned_9 owned
    controller -->|"Owns"| owned_10["Ingress"]
    class owned_10 owned
    controller -->|"Owns"| owned_11["NetworkPolicy"]
    class owned_11 owned
    controller -->|"Owns"| owned_12["PersistentVolumeClaim"]
    class owned_12 owned
    controller -->|"Owns"| owned_13["PodDisruptionBudget"]
    class owned_13 owned
    controller -->|"Owns"| owned_14["Role"]
    class owned_14 owned
    controller -->|"Owns"| owned_15["RoleBinding"]
    class owned_15 owned
    controller -->|"Owns"| owned_16["Service"]
    class owned_16 owned
    controller -->|"Owns"| owned_17["ServiceAccount"]
    class owned_17 owned
    controller -->|"Owns"| owned_18["ValidatingWebhookConfiguration"]
    class owned_18 owned
    controller -.->|"depends on"| odh_19["odh-platform-utilities"]
    class odh_19 dep
```

### CRDs

| Group | Version | Kind | Scope | Fields | Validation Rules | Discovery | Source |
|-------|---------|------|-------|--------|------------------|-----------|--------|
| llamastack.io | v1alpha1 | LlamaStackDistribution | Namespaced | 372 | 1 | YAML | [`config/crd/bases/llamastack.io_llamastackdistributions.yaml`](https://github.com/ogx-ai/ogx-k8s-operator/blob/e70c51de8847c4f108e84e6fb1a8aa7ab2bad7b4/config/crd/bases/llamastack.io_llamastackdistributions.yaml) |
| ogx.io | v1beta1 | OGXServer | Namespaced | 942 | 125 | YAML | [`config/crd/bases/ogx.io_ogxservers.yaml`](https://github.com/ogx-ai/ogx-k8s-operator/blob/e70c51de8847c4f108e84e6fb1a8aa7ab2bad7b4/config/crd/bases/ogx.io_ogxservers.yaml) |

## Dependencies

### Internal Platform Dependencies

| Component | Interaction |
|-----------|-------------|
| odh-platform-utilities | Go module dependency: github.com/opendatahub-io/odh-platform-utilities |

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/go-logr/logr | v1.4.4 |
| github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring | v0.82.0 |
| k8s.io/api | v0.35.2 |
| k8s.io/api | v0.35.2 |
| k8s.io/apiextensions-apiserver | v0.35.1 |
| k8s.io/apiextensions-apiserver | v0.35.2 |
| k8s.io/apimachinery | v0.35.2 |
| k8s.io/apimachinery | v0.35.4 |
| k8s.io/client-go | v0.35.2 |
| k8s.io/client-go | v0.35.2 |
| sigs.k8s.io/controller-runtime | v0.23.3 |
| sigs.k8s.io/controller-runtime | v0.23.3 |

