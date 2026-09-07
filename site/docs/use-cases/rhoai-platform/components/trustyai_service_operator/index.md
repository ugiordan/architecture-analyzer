# trustyai-service-operator

> **Architecture snapshot: 2026-09-07** (2026-09-07)


**Repository:** trustyai-explainability/trustyai-service-operator  
**Analyzer:** arch-analyzer dev  
**Extracted:** 2026-09-07T03:59:05Z

## Summary

| Metric | Count |
|--------|-------|
| CRDs | 5 |
| Deployments | 1 |
| Services | 3 |
| Secrets | 1 |
| Cluster Roles | 0 |
| Controller Watches | 18 |

## Component Architecture

CRDs, controllers, and owned Kubernetes resources.

```mermaid
graph LR
    %% Component architecture for trustyai-service-operator

    classDef crd fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef controller fill:#3498db,stroke:#2980b9,color:#fff
    classDef owned fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff
    classDef dep fill:#f39c12,stroke:#e67e22,color:#fff

    subgraph controller["trustyai-service-operator Controller"]
        dep_1["trustyai-service-operator-controller-manager"]
        class dep_1 controller
    end

    crd_EvalHub{{"EvalHub\ntrustyai.opendatahub.io/v1"}}
    class crd_EvalHub crd
    crd_EvalHub -->|"For (reconciles)"| controller
    crd_TrustyAIService{{"TrustyAIService\ntrustyai.opendatahub.io/v1"}}
    class crd_TrustyAIService crd
    crd_TrustyAIService -->|"For (reconciles)"| controller
    crd_EvalHub{{"EvalHub\ntrustyai.opendatahub.io/v1alpha1"}}
    class crd_EvalHub crd
    crd_EvalHub -->|"For (reconciles)"| controller
    crd_LMEvalJob{{"LMEvalJob\ntrustyai.opendatahub.io/v1alpha1"}}
    class crd_LMEvalJob crd
    crd_LMEvalJob -->|"For (reconciles)"| controller
    crd_TrustyAIService{{"TrustyAIService\ntrustyai.opendatahub.io/v1alpha1"}}
    class crd_TrustyAIService crd
    crd_TrustyAIService -->|"For (reconciles)"| controller
    controller -->|"Owns"| owned_2["ConfigMap"]
    class owned_2 owned
    controller -->|"Owns"| owned_3["Deployment"]
    class owned_3 owned
    controller -->|"Owns"| owned_4["Service"]
    class owned_4 owned
    controller -->|"Owns"| owned_5["ServiceMonitor"]
    class owned_5 owned
    watch_6["ConfigMap"] -->|"Watches"| controller
    class watch_6 external
    watch_7["InferenceService"] -->|"Watches"| controller
    class watch_7 external
    watch_8["Namespace"] -->|"Watches"| controller
    class watch_8 external
    controller -.->|"depends on"| odh_9["odh-platform-utilities"]
    class odh_9 dep
```

### CRDs

| Group | Version | Kind | Scope | Fields | Validation Rules | Discovery | Source |
|-------|---------|------|-------|--------|------------------|-----------|--------|
| trustyai.opendatahub.io | v1 | EvalHub | Namespaced | 19 | 0 | Go AST | [`/home/runner/work/_temp/arch-analyzer-repos/trustyai-service-operator/api/evalhub/v1/evalhub_types.go`](https://github.com/trustyai-explainability/trustyai-service-operator/blob/870559ac1034accf95ac655072444bf28d36ca19//home/runner/work/_temp/arch-analyzer-repos/trustyai-service-operator/api/evalhub/v1/evalhub_types.go) |
| trustyai.opendatahub.io | v1 | TrustyAIService | Namespaced | 19 | 0 | Go AST | [`/home/runner/work/_temp/arch-analyzer-repos/trustyai-service-operator/api/tas/v1/trustyaiservice_types.go`](https://github.com/trustyai-explainability/trustyai-service-operator/blob/870559ac1034accf95ac655072444bf28d36ca19//home/runner/work/_temp/arch-analyzer-repos/trustyai-service-operator/api/tas/v1/trustyaiservice_types.go) |
| trustyai.opendatahub.io | v1alpha1 | EvalHub | Namespaced | 19 | 0 | Go AST | [`/home/runner/work/_temp/arch-analyzer-repos/trustyai-service-operator/api/evalhub/v1alpha1/evalhub_types.go`](https://github.com/trustyai-explainability/trustyai-service-operator/blob/870559ac1034accf95ac655072444bf28d36ca19//home/runner/work/_temp/arch-analyzer-repos/trustyai-service-operator/api/evalhub/v1alpha1/evalhub_types.go) |
| trustyai.opendatahub.io | v1alpha1 | LMEvalJob | Namespaced | 19 | 0 | Go AST | [`/home/runner/work/_temp/arch-analyzer-repos/trustyai-service-operator/api/lmes/v1alpha1/lmevaljob_types.go`](https://github.com/trustyai-explainability/trustyai-service-operator/blob/870559ac1034accf95ac655072444bf28d36ca19//home/runner/work/_temp/arch-analyzer-repos/trustyai-service-operator/api/lmes/v1alpha1/lmevaljob_types.go) |
| trustyai.opendatahub.io | v1alpha1 | TrustyAIService | Namespaced | 19 | 0 | Go AST | [`/home/runner/work/_temp/arch-analyzer-repos/trustyai-service-operator/api/tas/v1alpha1/trustyaiservice_types.go`](https://github.com/trustyai-explainability/trustyai-service-operator/blob/870559ac1034accf95ac655072444bf28d36ca19//home/runner/work/_temp/arch-analyzer-repos/trustyai-service-operator/api/tas/v1alpha1/trustyaiservice_types.go) |

## Dependencies

### Internal Platform Dependencies

| Component | Interaction |
|-----------|-------------|
| odh-platform-utilities | Go module dependency: github.com/opendatahub-io/odh-platform-utilities |

### Key External Dependencies

| Module | Version |
|--------|---------|
| github.com/go-logr/logr | v1.4.4 |
| github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring | v0.64.1 |
| github.com/prometheus/client_golang | v1.24.1 |
| k8s.io/api | v0.35.3 |
| k8s.io/api | v0.34.5 |
| k8s.io/apiextensions-apiserver | v0.34.5 |
| k8s.io/apimachinery | v0.35.4 |
| k8s.io/apimachinery | v0.34.5 |
| k8s.io/client-go | v0.34.5 |
| k8s.io/client-go | v0.35.3 |
| sigs.k8s.io/controller-runtime | v0.22.5 |
| sigs.k8s.io/controller-runtime | v0.23.3 |

