# trainer: Network

### Services

No services found in analyzed sources.

### Network Policies

| Name | Policy Types | Source |
|------|-------------|--------|
| kubeflow-trainer-controller-manager | Ingress | [`manifests/rhoai/networkpolicy-ingress.yaml`](https://github.com/kubeflow/trainer/blob/24c87a6591d04a247b7408d19cddda6bd434c842/manifests/rhoai/networkpolicy-ingress.yaml) |
| kubeflow-trainer-controller-manager-egress | Egress | [`manifests/rhoai/networkpolicy-egress.yaml`](https://github.com/kubeflow/trainer/blob/24c87a6591d04a247b7408d19cddda6bd434c842/manifests/rhoai/networkpolicy-egress.yaml) |

## Network Policy Graph

Visual representation of NetworkPolicy rules. Ingress rules show what traffic is allowed into pods, egress rules show what traffic is allowed out.

```mermaid
graph LR
    classDef policy fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef pod fill:#3498db,stroke:#2980b9,color:#fff
    classDef external fill:#95a5a6,stroke:#7f8c8d,color:#fff

    trainer["trainer\nPods"]:::pod
    np_0_kubeflow_trainer_controller_manager{{"kubeflow-trainer-controller-manager\nIngress"}}:::policy
    np_0_kubeflow_trainer_controller_manager --> trainer
    np_1_kubeflow_trainer_controller_manager_egress{{"kubeflow-trainer-controller-manager-egress\nEgress"}}:::policy
    np_1_kubeflow_trainer_controller_manager_egress --> trainer
```

