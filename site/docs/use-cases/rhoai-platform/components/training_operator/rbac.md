# training-operator: RBAC

ServiceAccount bindings, roles, and resource permissions.

## RBAC Overview

This component defines a large RBAC surface (183 diagram lines). The graph below groups roles by permission scope.

```mermaid
graph LR
    classDef wide fill:#e74c3c,stroke:#c0392b,color:#fff
    classDef medium fill:#f39c12,stroke:#d68910,color:#fff
    classDef narrow fill:#2ecc71,stroke:#27ae60,color:#fff
    classDef subject fill:#3498db,stroke:#2980b9,color:#fff

    subgraph wide["Wide Scope (>30 resources)"]
    training_operator["training-operator\n32 resources"]:::wide
    end
    subgraph med["Medium Scope (10-30)"]
    kubeflow_training_edit["kubeflow-training-edit\n15 resources"]:::medium
    kubeflow_training_view["kubeflow-training-view\n13 resources"]:::medium
    training_edit["training-edit\n12 resources"]:::medium
    training_view["training-view\n12 resources"]:::medium
    end
    subgraph nar["Narrow Scope (<10)"]
    kubeflow_training_admin["kubeflow-training-admin"]:::narrow
    training_operator_webhook["training-operator-webhook\n1 resources"]:::narrow
    end

    subj_training_operator["training-operator\nServiceAccount"]:::subject
    subj_training_operator -->|binds| training_operator
    subj_training_operator -->|binds| training_operator_webhook
```

## Bindings

Subject-to-role mappings defining who has access to what.

| Binding | Type | Role | Subject |
|---------|------|------|---------|
| training-operator | ClusterRoleBinding | training-operator | ServiceAccount/training-operator |
| training-operator-webhook | RoleBinding | training-operator-webhook | ServiceAccount/training-operator |

## Role Details

Per-rule breakdown of API groups, resources, and verbs for each role.

| Role | Kind | API Groups | Resources | Verbs |
|------|------|------------|-----------|-------|
| kubeflow-training-edit | ClusterRole |  | mpijobs, tfjobs, pytorchjobs, xgboostjobs, paddlejobs, jaxjobs | create, delete, get, list, patch, update, watch |
| kubeflow-training-edit | ClusterRole |  | mpijobs/status, tfjobs/status, pytorchjobs/status, xgboostjobs/status, paddlejobs/status, jaxjobs/status | get |
| kubeflow-training-edit | ClusterRole |  | localqueues | get, list |
| kubeflow-training-edit | ClusterRole |  | persistentvolumeclaims | create, delete, get, list, watch |
| kubeflow-training-edit | ClusterRole |  | events | get, list, watch |
| kubeflow-training-view | ClusterRole |  | mpijobs, tfjobs, pytorchjobs, xgboostjobs, paddlejobs, jaxjobs | get, list, watch |
| kubeflow-training-view | ClusterRole |  | mpijobs/status, tfjobs/status, pytorchjobs/status, xgboostjobs/status, paddlejobs/status, jaxjobs/status | get |
| kubeflow-training-view | ClusterRole |  | localqueues | get, list |
| training-edit | ClusterRole |  | mpijobs, tfjobs, pytorchjobs, mxjobs, xgboostjobs, paddlejobs | create, delete, get, list, patch, update, watch |
| training-edit | ClusterRole |  | mpijobs/status, tfjobs/status, pytorchjobs/status, mxjobs/status, xgboostjobs/status, paddlejobs/status | get |
| training-operator | ClusterRole |  | configmaps | create, list, update, watch |
| training-operator | ClusterRole |  | events | create, delete, get, list, patch, update, watch |
| training-operator | ClusterRole |  | pods | create, delete, get, list, patch, update, watch |
| training-operator | ClusterRole |  | pods/exec | create |
| training-operator | ClusterRole |  | serviceaccounts | create, get, list, watch |
| training-operator | ClusterRole |  | services | create, delete, get, list, watch |
| training-operator | ClusterRole |  | validatingwebhookconfigurations | get, list, update, watch |
| training-operator | ClusterRole |  | horizontalpodautoscalers | create, delete, get, list, patch, update, watch |
| training-operator | ClusterRole |  | jaxjobs | create, delete, get, list, patch, update, watch |
| training-operator | ClusterRole |  | jaxjobs/finalizers | update |
| training-operator | ClusterRole |  | jaxjobs/status | get, patch, update |
| training-operator | ClusterRole |  | mpijobs | create, delete, get, list, patch, update, watch |
| training-operator | ClusterRole |  | mpijobs/finalizers | update |
| training-operator | ClusterRole |  | mpijobs/status | get, patch, update |
| training-operator | ClusterRole |  | paddlejobs | create, delete, get, list, patch, update, watch |
| training-operator | ClusterRole |  | paddlejobs/finalizers | update |
| training-operator | ClusterRole |  | paddlejobs/status | get, patch, update |
| training-operator | ClusterRole |  | pytorchjobs | create, delete, get, list, patch, update, watch |
| training-operator | ClusterRole |  | pytorchjobs/finalizers | update |
| training-operator | ClusterRole |  | pytorchjobs/status | get, patch, update |
| training-operator | ClusterRole |  | tfjobs | create, delete, get, list, patch, update, watch |
| training-operator | ClusterRole |  | tfjobs/finalizers | update |
| training-operator | ClusterRole |  | tfjobs/status | get, patch, update |
| training-operator | ClusterRole |  | xgboostjobs | create, delete, get, list, patch, update, watch |
| training-operator | ClusterRole |  | xgboostjobs/finalizers | update |
| training-operator | ClusterRole |  | xgboostjobs/status | get, patch, update |
| training-operator | ClusterRole |  | rolebindings | create, list, update, watch |
| training-operator | ClusterRole |  | roles | create, list, update, watch |
| training-operator | ClusterRole |  | podgroups | create, delete, get, list, patch, update, watch |
| training-operator | ClusterRole |  | podgroups | create, delete, get, list, patch, update, watch |
| training-operator | ClusterRole |  | apiservers | get |
| training-operator | ClusterRole |  | networkpolicies | create, delete, get, list, patch, update, watch |
| training-view | ClusterRole |  | mpijobs, tfjobs, pytorchjobs, mxjobs, xgboostjobs, paddlejobs | get, list, watch |
| training-view | ClusterRole |  | mpijobs/status, tfjobs/status, pytorchjobs/status, mxjobs/status, xgboostjobs/status, paddlejobs/status | get |
| training-operator-webhook | Role |  | secrets | get, list, update, watch |

