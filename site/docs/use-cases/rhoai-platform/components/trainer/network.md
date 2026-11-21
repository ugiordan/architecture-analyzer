# trainer: Network

### Services

No services found in analyzed sources.

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| kubeflow-trainer-controller-manager | Ingress | [`manifests/rhoai/networkpolicy-ingress.yaml`](https://github.com/red-hat-data-services/trainer/blob/2f3298cdf9a308731456dd30918a63aeff8fb5e8/manifests/rhoai/networkpolicy-ingress.yaml) |
| kubeflow-trainer-controller-manager-egress | Egress | [`manifests/rhoai/networkpolicy-egress.yaml`](https://github.com/red-hat-data-services/trainer/blob/2f3298cdf9a308731456dd30918a63aeff8fb5e8/manifests/rhoai/networkpolicy-egress.yaml) |

## Network Policy Graph

Visual representation of NetworkPolicy rules. Ingress rules show what traffic is allowed into pods, egress rules show what traffic is allowed out.

```mermaid
graph LR
    classDef policy fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef pod fill:#3498db,stroke:#2980b9,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff

    trainer["trainer\nPods"]:::pod
    np_0_kubeflow_trainer_controller_manager{{"kubeflow-trainer-controller-manager\nIngress"}}:::policy
    src_0_0_all["all pods"]:::external
    src_0_0_all -->|ingress| np_0_kubeflow_trainer_controller_manager
    src_0_1_0["all namespaces"]:::external
    src_0_1_0 -->|ingress| np_0_kubeflow_trainer_controller_manager
    src_0_2_0["ns: kubernetes.io/metadata.name=openshift-monitoring"]:::external
    src_0_2_0 -->|ingress| np_0_kubeflow_trainer_controller_manager
    np_0_kubeflow_trainer_controller_manager --> trainer
    np_1_kubeflow_trainer_controller_manager_egress{{"kubeflow-trainer-controller-manager-egress\nEgress"}}:::policy
    np_1_kubeflow_trainer_controller_manager_egress --> trainer
    dst_1_0_all["all destinations"]:::external
    trainer -->|egress| dst_1_0_all
```

