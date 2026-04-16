# odh-dashboard: Dataflow

## Controller Watches

| Type | GVK | Source |
|------|-----|--------|
| For | api/v1beta1/Workspace | `packages/notebooks/upstream/workspaces/controller/internal/controller/workspace_controller.go:469` |
| For | api/v1beta1/WorkspaceKind | `packages/notebooks/upstream/workspaces/controller/internal/controller/workspacekind_controller.go:175` |
| Owns | /v1/Service | `packages/notebooks/upstream/workspaces/controller/internal/controller/workspace_controller.go:471` |
| Owns | apps/v1/StatefulSet | `packages/notebooks/upstream/workspaces/controller/internal/controller/workspace_controller.go:470` |
| Owns | networking/v1/VirtualService | `packages/notebooks/upstream/workspaces/controller/internal/controller/workspace_controller.go:475` |

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for odh-dashboard

    participant KubernetesAPI as Kubernetes API
    participant odh_dashboard as odh-dashboard
    participant workspaces_backend as workspaces-backend
    participant workspaces_frontend as workspaces-frontend
    participant workspaces_controller as workspaces-controller

    KubernetesAPI->>+odh_dashboard: Watch Workspace (reconcile)
    odh_dashboard->>KubernetesAPI: Create/Update StatefulSet
    odh_dashboard->>KubernetesAPI: Create/Update Service
    odh_dashboard->>KubernetesAPI: Create/Update VirtualService
    KubernetesAPI->>+odh_dashboard: Watch WorkspaceKind (reconcile)

    Note over odh_dashboard: Exposed Services
    Note right of odh_dashboard: odh-dashboard:8443/TCP [dashboard-ui]
    Note right of odh_dashboard: workspaces-backend:4000/TCP [http-api]
    Note right of odh_dashboard: workspaces-webhook-service:443/TCP [https-webhook]
    Note right of odh_dashboard: workspaces-controller-metrics-service:8080/TCP [metrics]
    Note right of odh_dashboard: workspaces-frontend:8080/TCP [http-ui]
```

## Configuration

### ConfigMaps

| Name | Data Keys | Source |
|------|-----------|--------|
| model-registry-ui-config | images-jobs-async-upload | `manifests/common/model-registry/configmap.yaml` |
| federation-config | module-federation-config.json | `manifests/modular-architecture/federation-configmap.yaml` |
| federation-config | module-federation-config.json | `manifests/rhoai/shared/base/federation-configmap.yaml` |

