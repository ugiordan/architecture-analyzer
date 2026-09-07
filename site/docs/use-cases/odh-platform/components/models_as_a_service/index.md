# models-as-a-service

> **Architecture snapshot: 2026-09-07** (2026-09-07)


**Repository:** opendatahub-io/models-as-a-service  
**Analyzer:** arch-analyzer dev  
**Extracted:** 2026-09-07T03:56:41Z

## Summary

| Metric | Count |
|--------|-------|
| CRDs | 0 |
| Deployments | 4 |
| Services | 4 |
| Secrets | 4 |
| Cluster Roles | 0 |
| Controller Watches | 21 |

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
        dep_4["payload-processing"]
        class dep_4 controller
    end

    watch_5["AITenant"] -->|"Watches"| controller
    class watch_5 external
    watch_6["HTTPRoute"] -->|"Watches"| controller
    class watch_6 external
    watch_7["LLMInferenceService"] -->|"Watches"| controller
    class watch_7 external
    watch_8["MaaSAuthPolicy"] -->|"Watches"| controller
    class watch_8 external
    watch_9["MaaSModelRef"] -->|"Watches"| controller
    class watch_9 external
    watch_10["MaaSSubscription"] -->|"Watches"| controller
    class watch_10 external
    watch_11["Namespace"] -->|"Watches"| controller
    class watch_11 external
```

### CRDs

No CRDs found in analyzed sources.

## Dependencies

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/go-logr/logr | v1.4.4 |
| github.com/go-logr/logr | v1.4.3 |
| github.com/prometheus/client_golang | v1.24.1 |
| github.com/prometheus/client_model | v0.6.3 |
| google.golang.org/grpc | v1.83.2 |
| k8s.io/api | v0.35.3 |
| k8s.io/api | v0.37.0 |
| k8s.io/apiextensions-apiserver | v0.35.3 |
| k8s.io/apimachinery | v0.35.3 |
| k8s.io/apimachinery | v0.37.0 |
| k8s.io/client-go | v0.37.0 |
| k8s.io/client-go | v0.35.3 |
| sigs.k8s.io/controller-runtime | v0.25.0 |
| sigs.k8s.io/controller-runtime | v0.22.5 |

