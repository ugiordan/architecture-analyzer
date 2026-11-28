# fms-guardrails-orchestrator: Dataflow

[Overview](index.md) · [Agent Context](agent-context.md) · [Extracted Data](details.md) · [Network](network.md) · [RBAC](rbac.md) · [Security](security.md) · [Dataflow](dataflow.md)

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

No controller watches found in analyzed sources.

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for fms-guardrails-orchestrator

    participant KubernetesAPI as Kubernetes API
    participant fms_guardrails_orchestrator as fms-guardrails-orchestrator
```

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

