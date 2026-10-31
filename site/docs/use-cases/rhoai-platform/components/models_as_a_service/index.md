# models-as-a-service

> **Architecture snapshot: 2026-10-31** (2026-10-31)


**Repository:** red-hat-data-services/models-as-a-service  
**Analyzer:** arch-analyzer dev  
**Extracted:** 2026-10-31T04:56:33Z

## Summary

| Metric | Count |
|--------|-------|
| CRDs | 0 |
| Deployments | 5 |
| Services | 5 |
| Secrets | 5 |
| Cluster Roles | 0 |
| Controller Watches | 27 |

## Component Architecture

CRDs, controllers, and owned Kubernetes resources.

```mermaid
graph LR
    %% Component architecture for models-as-a-service

    classDef crd fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef controller fill:#3498db,stroke:#2980b9,color:#fff
    classDef owned fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef dep fill:#f39c12,stroke:#e67e22,color:#fff

    subgraph controller["models-as-a-service Controller"]
        dep_1["maas-api"]
        class dep_1 controller
        dep_2["maas-api"]
        class dep_2 controller
        dep_3["maas-controller"]
        class dep_3 controller
        dep_4["maas-discovery"]
        class dep_4 controller
        dep_5["payload-processing"]
        class dep_5 controller
    end

    watch_6["AITenant"] -->|"Watches"| controller
    class watch_6 external
    watch_7["HTTPRoute"] -->|"Watches"| controller
    class watch_7 external
    watch_8["InferenceObjective"] -->|"Watches"| controller
    class watch_8 external
    watch_9["LLMInferenceService"] -->|"Watches"| controller
    class watch_9 external
    watch_10["MaaSAuthPolicy"] -->|"Watches"| controller
    class watch_10 external
    watch_11["MaaSModelRef"] -->|"Watches"| controller
    class watch_11 external
    watch_12["MaaSSubscription"] -->|"Watches"| controller
    class watch_12 external
    watch_13["MaasTenantConfig"] -->|"Watches"| controller
    class watch_13 external
    watch_14["Namespace"] -->|"Watches"| controller
    class watch_14 external
    watch_15["Tenant"] -->|"Watches"| controller
    class watch_15 external
    controller -.->|"depends on"| odh_16["ai-gateway-controller"]
    class odh_16 dep
    controller -.->|"depends on"| odh_17["ai-gateway-controller"]
    class odh_17 dep
    controller -.->|"depends on"| odh_18["ai-gateway-controller"]
    class odh_18 dep
    controller -.->|"depends on"| odh_19["kserve"]
    class odh_19 dep
    controller -.->|"depends on"| odh_20["models-as-a-service"]
    class odh_20 dep
    controller -.->|"depends on"| odh_21["models-as-a-service"]
    class odh_21 dep
```

### CRDs

No CRDs found in analyzed sources.

## Dependencies

### Internal Platform Dependencies

| Component | Interaction |
|-----------|-------------|
| ai-gateway-controller | Go module dependency: github.com/opendatahub-io/ai-gateway-controller |
| ai-gateway-controller | Go module dependency: github.com/opendatahub-io/ai-gateway-controller |
| ai-gateway-controller | Go module dependency: github.com/opendatahub-io/ai-gateway-controller |
| kserve | Go module dependency: github.com/opendatahub-io/kserve |
| models-as-a-service | Go module dependency: github.com/opendatahub-io/models-as-a-service/maas-common |
| models-as-a-service | Go module dependency: github.com/opendatahub-io/models-as-a-service/maas-common |

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/go-logr/logr | v1.4.4 |
| github.com/go-logr/logr | v1.4.4 |
| github.com/go-logr/zapr | v1.3.0 |
| github.com/prometheus/client_golang | v1.24.1 |
| github.com/prometheus/client_model | v0.6.3 |
| google.golang.org/grpc | v1.84.0 |
| k8s.io/api | v0.37.1 |
| k8s.io/api | v0.35.3 |
| k8s.io/api | v0.36.4 |
| k8s.io/apiextensions-apiserver | v0.36.4 |
| k8s.io/apimachinery | v0.36.4 |
| k8s.io/apimachinery | v0.35.3 |
| k8s.io/apimachinery | v0.35.3 |
| k8s.io/apimachinery | v0.37.1 |
| k8s.io/apiserver | v0.36.4 |
| k8s.io/client-go | v0.37.1 |
| k8s.io/client-go | v0.36.4 |
| k8s.io/client-go | v0.35.3 |
| sigs.k8s.io/controller-runtime | v0.24.1 |
| sigs.k8s.io/controller-runtime | v0.25.2 |

