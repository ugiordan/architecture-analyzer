# ogx-k8s-operator

> **Architecture snapshot: 2026-10-24** (2026-10-24)


**Repository:** opendatahub-io/ogx-k8s-operator  
**Analyzer:** arch-analyzer dev  
**Extracted:** 2026-10-24T04:06:52Z

## Summary

| Metric | Count |
|--------|-------|
| CRDs | 2 |
| Deployments | 2 |
| Services | 2 |
| Secrets | 1 |
| Cluster Roles | 5 |
| Controller Watches | 10 |

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
    end

    crd_LlamaStackDistribution{{"LlamaStackDistribution\nllamastack.io/v1alpha1"}}
    class crd_LlamaStackDistribution crd
    crd_OGXServer{{"OGXServer\nogx.io/v1beta1"}}
    class crd_OGXServer crd
    crd_OGXServer -->|"For (reconciles)"| controller
    controller -->|"Owns"| owned_3["ConfigMap"]
    class owned_3 owned
    controller -->|"Owns"| owned_4["Deployment"]
    class owned_4 owned
    controller -->|"Owns"| owned_5["HorizontalPodAutoscaler"]
    class owned_5 owned
    controller -->|"Owns"| owned_6["Ingress"]
    class owned_6 owned
    controller -->|"Owns"| owned_7["NetworkPolicy"]
    class owned_7 owned
    controller -->|"Owns"| owned_8["PersistentVolumeClaim"]
    class owned_8 owned
    controller -->|"Owns"| owned_9["PodDisruptionBudget"]
    class owned_9 owned
    controller -->|"Owns"| owned_10["Service"]
    class owned_10 owned
```

### CRDs

| Group | Version | Kind | Scope | Fields | Validation Rules | Discovery | Source |
|-------|---------|------|-------|--------|------------------|-----------|--------|
| llamastack.io | v1alpha1 | LlamaStackDistribution | Namespaced | 371 | 1 | YAML | [`config/crd/bases/llamastack.io_llamastackdistributions.yaml`](https://github.com/opendatahub-io/ogx-k8s-operator/blob/e2e94610d6d884caae50b8615b3b67d3df9f6518/config/crd/bases/llamastack.io_llamastackdistributions.yaml) |
| ogx.io | v1beta1 | OGXServer | Namespaced | 892 | 118 | YAML | [`config/crd/bases/ogx.io_ogxservers.yaml`](https://github.com/opendatahub-io/ogx-k8s-operator/blob/e2e94610d6d884caae50b8615b3b67d3df9f6518/config/crd/bases/ogx.io_ogxservers.yaml) |

## Dependencies

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/go-logr/logr | v1.4.3 |
| k8s.io/api | v0.34.3 |
| k8s.io/apiextensions-apiserver | v0.34.3 |
| k8s.io/apimachinery | v0.34.3 |
| k8s.io/client-go | v0.34.3 |
| sigs.k8s.io/controller-runtime | v0.22.4 |

