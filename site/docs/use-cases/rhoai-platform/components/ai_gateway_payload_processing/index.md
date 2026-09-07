# ai-gateway-payload-processing

> **Architecture snapshot: 2026-09-07** (2026-09-07)


**Repository:** opendatahub-io/ai-gateway-payload-processing  
**Analyzer:** arch-analyzer dev  
**Extracted:** 2026-09-07T04:03:21Z

## Summary

| Metric | Count |
|--------|-------|
| CRDs | 2 |
| Deployments | 0 |
| Services | 0 |
| Secrets | 0 |
| Cluster Roles | 0 |
| Controller Watches | 5 |

## Component Architecture

CRDs, controllers, and owned Kubernetes resources.

```mermaid
graph LR
    %% Component architecture for ai-gateway-payload-processing

    classDef crd fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef controller fill:#3498db,stroke:#2980b9,color:#fff
    classDef owned fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef dep fill:#f39c12,stroke:#e67e22,color:#fff

    subgraph controller["ai-gateway-payload-processing Controller"]
        ctrl_1["Controller"]
        class ctrl_1 controller
    end

    crd_ExternalModel{{"ExternalModel\ninference.opendatahub.io/v1alpha1"}}
    class crd_ExternalModel crd
    crd_ExternalModel -->|"For (reconciles)"| controller
    crd_ExternalProvider{{"ExternalProvider\ninference.opendatahub.io/v1alpha1"}}
    class crd_ExternalProvider crd
    crd_ExternalProvider -->|"For (reconciles)"| controller
    controller -->|"Owns"| owned_2["HTTPRoute"]
    class owned_2 owned
    controller -->|"Owns"| owned_3["Service"]
    class owned_3 owned
    watch_4["ExternalProvider"] -->|"Watches"| controller
    class watch_4 external
```

### CRDs

| Group | Version | Kind | Scope | Fields | Validation Rules | Discovery | Source |
|-------|---------|------|-------|--------|------------------|-----------|--------|
| inference.opendatahub.io | v1alpha1 | ExternalModel | Namespaced | 27 | 0 | YAML | [`config/crd/bases/inference.opendatahub.io_externalmodels.yaml`](https://github.com/opendatahub-io/ai-gateway-payload-processing/blob/582c80e2b18ecda2ede4a40f597b7efc8da0dcc5/config/crd/bases/inference.opendatahub.io_externalmodels.yaml) |
| inference.opendatahub.io | v1alpha1 | ExternalProvider | Namespaced | 20 | 0 | YAML | [`config/crd/bases/inference.opendatahub.io_externalproviders.yaml`](https://github.com/opendatahub-io/ai-gateway-payload-processing/blob/582c80e2b18ecda2ede4a40f597b7efc8da0dcc5/config/crd/bases/inference.opendatahub.io_externalproviders.yaml) |

## Dependencies

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/go-logr/logr | v1.4.4 |
| github.com/prometheus/client_golang | v1.24.1 |
| google.golang.org/grpc | v1.83.0 |
| k8s.io/api | v0.35.7 |
| k8s.io/apimachinery | v0.35.7 |
| k8s.io/client-go | v0.35.7 |
| sigs.k8s.io/controller-runtime | v0.23.3 |

