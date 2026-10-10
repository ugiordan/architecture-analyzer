# mcp-lifecycle-module-operator: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

| Type | GVK | Source |
|------|-----|--------|
| For | api/v1alpha1/MCPLifecycleOperator | [`internal/controller/mcplifecycleoperator_reconciler.go:709`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/bfacb7cefc37fc46f22209f5accd93bed36b3718/internal/controller/mcplifecycleoperator_reconciler.go#L709) |
| Watches | /v1/ConfigMap | [`internal/controller/mcplifecycleoperator_reconciler.go:710`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/bfacb7cefc37fc46f22209f5accd93bed36b3718/internal/controller/mcplifecycleoperator_reconciler.go#L710) |
| Watches | /v1/Service | [`internal/controller/mcplifecycleoperator_reconciler.go:713`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/bfacb7cefc37fc46f22209f5accd93bed36b3718/internal/controller/mcplifecycleoperator_reconciler.go#L713) |
| Watches | /v1/ServiceAccount | [`internal/controller/mcplifecycleoperator_reconciler.go:712`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/bfacb7cefc37fc46f22209f5accd93bed36b3718/internal/controller/mcplifecycleoperator_reconciler.go#L712) |
| Watches | apiextensions.k8s.io/v1/CustomResourceDefinition | [`internal/controller/mcplifecycleoperator_reconciler.go:718`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/bfacb7cefc37fc46f22209f5accd93bed36b3718/internal/controller/mcplifecycleoperator_reconciler.go#L718) |
| Watches | apps/v1/Deployment | [`internal/controller/mcplifecycleoperator_reconciler.go:711`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/bfacb7cefc37fc46f22209f5accd93bed36b3718/internal/controller/mcplifecycleoperator_reconciler.go#L711) |
| Watches | config/v1/APIServer | [`internal/controller/mcplifecycleoperator_reconciler.go:721`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/bfacb7cefc37fc46f22209f5accd93bed36b3718/internal/controller/mcplifecycleoperator_reconciler.go#L721) |
| Watches | rbac.authorization.k8s.io/v1/ClusterRole | [`internal/controller/mcplifecycleoperator_reconciler.go:714`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/bfacb7cefc37fc46f22209f5accd93bed36b3718/internal/controller/mcplifecycleoperator_reconciler.go#L714) |
| Watches | rbac.authorization.k8s.io/v1/ClusterRoleBinding | [`internal/controller/mcplifecycleoperator_reconciler.go:715`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/bfacb7cefc37fc46f22209f5accd93bed36b3718/internal/controller/mcplifecycleoperator_reconciler.go#L715) |
| Watches | rbac.authorization.k8s.io/v1/Role | [`internal/controller/mcplifecycleoperator_reconciler.go:716`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/bfacb7cefc37fc46f22209f5accd93bed36b3718/internal/controller/mcplifecycleoperator_reconciler.go#L716) |
| Watches | rbac.authorization.k8s.io/v1/RoleBinding | [`internal/controller/mcplifecycleoperator_reconciler.go:717`](https://github.com/opendatahub-io/mcp-lifecycle-module-operator/blob/bfacb7cefc37fc46f22209f5accd93bed36b3718/internal/controller/mcplifecycleoperator_reconciler.go#L717) |

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for mcp-lifecycle-module-operator

    participant KubernetesAPI as Kubernetes API
    participant controller_manager as controller-manager

    KubernetesAPI->>+controller_manager: Watch MCPLifecycleOperator (reconcile)
    KubernetesAPI-->>+controller_manager: Watch ConfigMap (informer)
    KubernetesAPI-->>+controller_manager: Watch Service (informer)
    KubernetesAPI-->>+controller_manager: Watch ServiceAccount (informer)
    KubernetesAPI-->>+controller_manager: Watch CustomResourceDefinition (informer)
    KubernetesAPI-->>+controller_manager: Watch Deployment (informer)
    KubernetesAPI-->>+controller_manager: Watch APIServer (informer)
    KubernetesAPI-->>+controller_manager: Watch ClusterRole (informer)
    KubernetesAPI-->>+controller_manager: Watch ClusterRoleBinding (informer)
    KubernetesAPI-->>+controller_manager: Watch Role (informer)
    KubernetesAPI-->>+controller_manager: Watch RoleBinding (informer)

    Note over controller_manager: Exposed Services
    Note right of controller_manager: webhook-service:443/TCP []

    Note over KubernetesAPI: Defined CRDs
    Note right of KubernetesAPI: MCPLifecycleOperator (components.platform.opendatahub.io/v1alpha1)
```

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

