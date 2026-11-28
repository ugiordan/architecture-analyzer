# modelmesh-runtime-adapter: Dataflow

[Overview](index.md) · [Agent Context](agent-context.md) · [Extracted Data](details.md) · [Network](network.md) · [RBAC](rbac.md) · [Security](security.md) · [Dataflow](dataflow.md)

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

No controller watches found in analyzed sources.

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for modelmesh-runtime-adapter

    participant KubernetesAPI as Kubernetes API
    participant modelmesh_runtime_adapter as modelmesh-runtime-adapter
```

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

