# opendatahub-operator: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

| Type | GVK | Source |
|------|-----|--------|
| Owns | /v1/ConfigMap | `internal/controller/components/modelcontroller/modelcontroller_controller.go:53` |
| Owns | /v1/ConfigMap | `internal/controller/components/feastoperator/feastoperator_controller.go:28` |
| Owns | /v1/ConfigMap | `internal/controller/components/sparkoperator/sparkoperator_controller.go:45` |
| Owns | /v1/ConfigMap | `internal/controller/components/trainer/trainer_controller.go:50` |
| Owns | /v1/ConfigMap | `internal/controller/components/kserve/kserve_controller.go:59` |
| Owns | /v1/ConfigMap | `internal/controller/services/monitoring/monitoring_controller.go:96` |
| Owns | /v1/ConfigMap | `internal/controller/components/kueue/kueue_controller.go:59` |
| Owns | /v1/ConfigMap | `internal/controller/components/ray/ray_controller.go:48` |
| Owns | /v1/ConfigMap | `internal/controller/components/modelsasservice/modelsasservice_controller.go:63` |
| Owns | /v1/ConfigMap | `internal/controller/components/trainingoperator/trainingoperator_controller.go:45` |
| Owns | /v1/ConfigMap | `internal/controller/components/dashboard/dashboard_controller.go:56` |
| Owns | /v1/ConfigMap | `internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go:46` |
| Owns | /v1/ConfigMap | `internal/controller/components/llamastackoperator/llamastackoperator_controller.go:29` |
| Owns | /v1/ConfigMap | `internal/controller/components/modelregistry/modelregistry_controller.go:49` |
| Owns | /v1/ConfigMap | `internal/controller/components/mlflowoperator/mlflowoperator_controller.go:48` |
| Owns | /v1/ConfigMap | `internal/controller/components/modelregistry/modelregistry_controller.go:48` |
| Owns | /v1/ConfigMap | `internal/controller/components/trustyai/trustyai_controller.go:63` |
| Owns | /v1/ConfigMap | `internal/controller/components/workbenches/workbenches_controller.go:49` |
| Owns | /v1/Secret | `internal/controller/components/kserve/kserve_controller.go:57` |
| Owns | /v1/Secret | `internal/controller/components/modelcontroller/modelcontroller_controller.go:55` |
| Owns | /v1/Secret | `internal/controller/components/workbenches/workbenches_controller.go:50` |
| Owns | /v1/Secret | `internal/controller/components/modelregistry/modelregistry_controller.go:50` |
| Owns | /v1/Secret | `internal/controller/components/dashboard/dashboard_controller.go:57` |
| Owns | /v1/Secret | `internal/controller/components/kueue/kueue_controller.go:60` |
| Owns | /v1/Secret | `internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go:47` |
| Owns | /v1/Secret | `internal/controller/components/ray/ray_controller.go:49` |
| Owns | /v1/Secret | `internal/controller/services/monitoring/monitoring_controller.go:97` |
| Owns | /v1/Service | `internal/controller/components/feastoperator/feastoperator_controller.go:34` |
| Owns | /v1/Service | `internal/controller/components/kueue/kueue_controller.go:66` |
| Owns | /v1/Service | `internal/controller/components/sparkoperator/sparkoperator_controller.go:51` |
| Owns | /v1/Service | `internal/controller/components/ray/ray_controller.go:55` |
| Owns | /v1/Service | `internal/controller/services/monitoring/monitoring_controller.go:98` |
| Owns | /v1/Service | `internal/controller/components/dashboard/dashboard_controller.go:63` |
| Owns | /v1/Service | `internal/controller/components/kserve/kserve_controller.go:58` |
| Owns | /v1/Service | `internal/controller/components/trainer/trainer_controller.go:55` |
| Owns | /v1/Service | `internal/controller/components/modelsasservice/modelsasservice_controller.go:64` |
| Owns | /v1/Service | `internal/controller/components/modelregistry/modelregistry_controller.go:55` |
| Owns | /v1/Service | `internal/controller/components/trainingoperator/trainingoperator_controller.go:50` |
| Owns | /v1/Service | `internal/controller/components/modelcontroller/modelcontroller_controller.go:62` |
| Owns | /v1/Service | `internal/controller/components/trustyai/trustyai_controller.go:69` |
| Owns | /v1/Service | `internal/controller/components/mlflowoperator/mlflowoperator_controller.go:53` |
| Owns | /v1/Service | `internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go:53` |
| Owns | /v1/Service | `internal/controller/components/llamastackoperator/llamastackoperator_controller.go:35` |
| Owns | /v1/Service | `internal/controller/components/workbenches/workbenches_controller.go:56` |
| Owns | /v1/ServiceAccount | `internal/controller/components/modelregistry/modelregistry_controller.go:56` |
| Owns | /v1/ServiceAccount | `internal/controller/components/kserve/kserve_controller.go:60` |
| Owns | /v1/ServiceAccount | `internal/controller/components/feastoperator/feastoperator_controller.go:33` |
| Owns | /v1/ServiceAccount | `internal/controller/components/llamastackoperator/llamastackoperator_controller.go:34` |
| Owns | /v1/ServiceAccount | `internal/controller/services/monitoring/monitoring_controller.go:99` |
| Owns | /v1/ServiceAccount | `internal/controller/components/sparkoperator/sparkoperator_controller.go:50` |
| Owns | /v1/ServiceAccount | `internal/controller/components/ray/ray_controller.go:54` |
| Owns | /v1/ServiceAccount | `internal/controller/components/trainer/trainer_controller.go:54` |
| Owns | /v1/ServiceAccount | `internal/controller/components/kueue/kueue_controller.go:65` |
| Owns | /v1/ServiceAccount | `internal/controller/components/dashboard/dashboard_controller.go:62` |
| Owns | /v1/ServiceAccount | `internal/controller/components/modelsasservice/modelsasservice_controller.go:65` |
| Owns | /v1/ServiceAccount | `internal/controller/components/workbenches/workbenches_controller.go:55` |
| Owns | /v1/ServiceAccount | `internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go:52` |
| Owns | /v1/ServiceAccount | `internal/controller/components/mlflowoperator/mlflowoperator_controller.go:54` |
| Owns | /v1/ServiceAccount | `internal/controller/components/trainingoperator/trainingoperator_controller.go:49` |
| Owns | /v1/ServiceAccount | `internal/controller/components/trustyai/trustyai_controller.go:64` |
| Owns | /v1/ServiceAccount | `internal/controller/components/modelcontroller/modelcontroller_controller.go:54` |
| Owns | admissionregistration.k8s.io/v1/MutatingWebhookConfiguration | `internal/controller/components/kueue/kueue_controller.go:70` |
| Owns | admissionregistration.k8s.io/v1/MutatingWebhookConfiguration | `internal/controller/components/modelregistry/modelregistry_controller.go:58` |
| Owns | admissionregistration.k8s.io/v1/MutatingWebhookConfiguration | `internal/controller/components/workbenches/workbenches_controller.go:57` |
| Owns | admissionregistration.k8s.io/v1/MutatingWebhookConfiguration | `internal/controller/components/kserve/kserve_controller.go:66` |
| Owns | admissionregistration.k8s.io/v1/MutatingWebhookConfiguration | `internal/controller/components/sparkoperator/sparkoperator_controller.go:53` |
| Owns | admissionregistration.k8s.io/v1/ValidatingAdmissionPolicy | `internal/controller/components/kserve/kserve_controller.go:68` |
| Owns | admissionregistration.k8s.io/v1/ValidatingAdmissionPolicyBinding | `internal/controller/components/kserve/kserve_controller.go:69` |
| Owns | admissionregistration.k8s.io/v1/ValidatingWebhookConfiguration | `internal/controller/components/sparkoperator/sparkoperator_controller.go:54` |
| Owns | admissionregistration.k8s.io/v1/ValidatingWebhookConfiguration | `internal/controller/components/kserve/kserve_controller.go:67` |
| Owns | admissionregistration.k8s.io/v1/ValidatingWebhookConfiguration | `internal/controller/components/modelcontroller/modelcontroller_controller.go:63` |
| Owns | admissionregistration.k8s.io/v1/ValidatingWebhookConfiguration | `internal/controller/components/modelregistry/modelregistry_controller.go:59` |
| Owns | admissionregistration.k8s.io/v1/ValidatingWebhookConfiguration | `internal/controller/components/kueue/kueue_controller.go:71` |
| Owns | apis/v1/Gateway | `internal/controller/components/modelsasservice/modelsasservice_controller.go:74` |
| Owns | apis/v1/HTTPRoute | `internal/controller/components/dashboard/dashboard_controller.go:70` |
| Owns | apis/v1/HTTPRoute | `internal/controller/components/modelsasservice/modelsasservice_controller.go:73` |
| Owns | apps/v1/Deployment | `internal/controller/components/trustyai/trustyai_controller.go:70` |
| Owns | apps/v1/Deployment | `internal/controller/components/kueue/kueue_controller.go:72` |
| Owns | apps/v1/Deployment | `internal/controller/components/mlflowoperator/mlflowoperator_controller.go:57` |
| Owns | apps/v1/Deployment | `internal/controller/components/sparkoperator/sparkoperator_controller.go:55` |
| Owns | apps/v1/Deployment | `internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go:55` |
| Owns | apps/v1/Deployment | `internal/controller/components/ray/ray_controller.go:56` |
| Owns | apps/v1/Deployment | `internal/controller/components/trainer/trainer_controller.go:56` |
| Owns | apps/v1/Deployment | `internal/controller/components/modelsasservice/modelsasservice_controller.go:66` |
| Owns | apps/v1/Deployment | `internal/controller/components/dashboard/dashboard_controller.go:67` |
| Owns | apps/v1/Deployment | `internal/controller/components/feastoperator/feastoperator_controller.go:35` |
| Owns | apps/v1/Deployment | `internal/controller/components/workbenches/workbenches_controller.go:58` |
| Owns | apps/v1/Deployment | `internal/controller/components/trainingoperator/trainingoperator_controller.go:51` |
| Owns | apps/v1/Deployment | `internal/controller/components/modelregistry/modelregistry_controller.go:57` |
| Owns | apps/v1/Deployment | `internal/controller/components/kserve/kserve_controller.go:70` |
| Owns | apps/v1/Deployment | `internal/controller/components/llamastackoperator/llamastackoperator_controller.go:37` |
| Owns | apps/v1/Deployment | `internal/controller/components/modelcontroller/modelcontroller_controller.go:65` |
| Owns | apps/v1/Deployment | `internal/controller/services/monitoring/monitoring_controller.go:94` |
| Owns | batch/v1/Job | `internal/controller/services/monitoring/monitoring_controller.go:95` |
| Owns | components/v1alpha1/Dashboard | `internal/controller/datasciencecluster/datasciencecluster_controller.go:46` |
| Owns | components/v1alpha1/DataSciencePipelines | `internal/controller/datasciencecluster/datasciencecluster_controller.go:54` |
| Owns | components/v1alpha1/FeastOperator | `internal/controller/datasciencecluster/datasciencecluster_controller.go:58` |
| Owns | components/v1alpha1/Kserve | `internal/controller/datasciencecluster/datasciencecluster_controller.go:55` |
| Owns | components/v1alpha1/Kueue | `internal/controller/datasciencecluster/datasciencecluster_controller.go:51` |
| Owns | components/v1alpha1/LlamaStackOperator | `internal/controller/datasciencecluster/datasciencecluster_controller.go:59` |
| Owns | components/v1alpha1/MLflowOperator | `internal/controller/datasciencecluster/datasciencecluster_controller.go:60` |
| Owns | components/v1alpha1/ModelController | `internal/controller/datasciencecluster/datasciencecluster_controller.go:57` |
| Owns | components/v1alpha1/ModelRegistry | `internal/controller/datasciencecluster/datasciencecluster_controller.go:49` |
| Owns | components/v1alpha1/Ray | `internal/controller/datasciencecluster/datasciencecluster_controller.go:48` |
| Owns | components/v1alpha1/SparkOperator | `internal/controller/datasciencecluster/datasciencecluster_controller.go:61` |
| Owns | components/v1alpha1/Trainer | `internal/controller/datasciencecluster/datasciencecluster_controller.go:53` |
| Owns | components/v1alpha1/TrainingOperator | `internal/controller/datasciencecluster/datasciencecluster_controller.go:52` |
| Owns | components/v1alpha1/TrustyAI | `internal/controller/datasciencecluster/datasciencecluster_controller.go:50` |
| Owns | components/v1alpha1/Workbenches | `internal/controller/datasciencecluster/datasciencecluster_controller.go:47` |
| Owns | console/v1/ConsoleLink | `internal/controller/components/dashboard/dashboard_controller.go:71` |
| Owns | console/v1/ConsoleLink | `internal/controller/components/mlflowoperator/mlflowoperator_controller.go:55` |
| Owns | maas/v1alpha1/MaaSTenant | `internal/controller/datasciencecluster/datasciencecluster_controller.go:56` |
| Owns | monitoring/v1/PodMonitor | `internal/controller/components/sparkoperator/sparkoperator_controller.go:52` |
| Owns | monitoring/v1/PodMonitor | `internal/controller/components/trainingoperator/trainingoperator_controller.go:46` |
| Owns | monitoring/v1/PodMonitor | `internal/controller/components/kueue/kueue_controller.go:68` |
| Owns | monitoring/v1/PodMonitor | `internal/controller/components/trainer/trainer_controller.go:51` |
| Owns | monitoring/v1/PrometheusRule | `internal/controller/components/kueue/kueue_controller.go:69` |
| Owns | monitoring/v1/ServiceMonitor | `internal/controller/components/mlflowoperator/mlflowoperator_controller.go:56` |
| Owns | monitoring/v1/ServiceMonitor | `internal/controller/components/modelcontroller/modelcontroller_controller.go:56` |
| Owns | monitoring/v1/ServiceMonitor | `internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go:54` |
| Owns | networking.k8s.io/v1/NetworkPolicy | `internal/controller/components/modelsasservice/modelsasservice_controller.go:71` |
| Owns | networking.k8s.io/v1/NetworkPolicy | `internal/controller/components/modelcontroller/modelcontroller_controller.go:57` |
| Owns | networking.k8s.io/v1/NetworkPolicy | `internal/controller/components/kserve/kserve_controller.go:65` |
| Owns | networking.k8s.io/v1/NetworkPolicy | `internal/controller/components/kueue/kueue_controller.go:67` |
| Owns | networking.k8s.io/v1/NetworkPolicy | `internal/controller/services/monitoring/monitoring_controller.go:93` |
| Owns | networking.k8s.io/v1/NetworkPolicy | `internal/controller/components/dashboard/dashboard_controller.go:73` |
| Owns | policy/v1/PodDisruptionBudget | `internal/controller/components/llamastackoperator/llamastackoperator_controller.go:36` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | `internal/controller/components/modelregistry/modelregistry_controller.go:53` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | `internal/controller/components/sparkoperator/sparkoperator_controller.go:49` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | `internal/controller/components/mlflowoperator/mlflowoperator_controller.go:49` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | `internal/controller/components/kueue/kueue_controller.go:62` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | `internal/controller/components/dashboard/dashboard_controller.go:59` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | `internal/controller/components/workbenches/workbenches_controller.go:52` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | `internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go:49` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | `internal/controller/components/kserve/kserve_controller.go:63` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | `internal/controller/components/ray/ray_controller.go:51` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | `internal/controller/components/llamastackoperator/llamastackoperator_controller.go:33` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | `internal/controller/components/trainer/trainer_controller.go:53` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | `internal/controller/components/feastoperator/feastoperator_controller.go:32` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | `internal/controller/components/trustyai/trustyai_controller.go:66` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | `internal/controller/components/modelsasservice/modelsasservice_controller.go:68` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | `internal/controller/components/modelcontroller/modelcontroller_controller.go:59` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | `internal/controller/services/monitoring/monitoring_controller.go:91` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | `internal/controller/services/auth/auth_controller.go:60` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | `internal/controller/components/trainingoperator/trainingoperator_controller.go:48` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | `internal/controller/components/trainer/trainer_controller.go:52` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | `internal/controller/services/monitoring/monitoring_controller.go:92` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | `internal/controller/components/modelcontroller/modelcontroller_controller.go:61` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | `internal/controller/services/auth/auth_controller.go:59` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | `internal/controller/components/modelregistry/modelregistry_controller.go:54` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | `internal/controller/components/trustyai/trustyai_controller.go:65` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | `internal/controller/components/modelsasservice/modelsasservice_controller.go:69` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | `internal/controller/components/feastoperator/feastoperator_controller.go:31` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | `internal/controller/components/ray/ray_controller.go:50` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | `internal/controller/components/trainingoperator/trainingoperator_controller.go:47` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | `internal/controller/components/kserve/kserve_controller.go:64` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | `internal/controller/components/kueue/kueue_controller.go:61` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | `internal/controller/components/mlflowoperator/mlflowoperator_controller.go:50` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | `internal/controller/components/workbenches/workbenches_controller.go:51` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | `internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go:48` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | `internal/controller/components/dashboard/dashboard_controller.go:58` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | `internal/controller/components/llamastackoperator/llamastackoperator_controller.go:32` |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | `internal/controller/components/sparkoperator/sparkoperator_controller.go:48` |
| Owns | rbac.authorization.k8s.io/v1/Role | `internal/controller/components/modelregistry/modelregistry_controller.go:51` |
| Owns | rbac.authorization.k8s.io/v1/Role | `internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go:50` |
| Owns | rbac.authorization.k8s.io/v1/Role | `internal/controller/components/llamastackoperator/llamastackoperator_controller.go:31` |
| Owns | rbac.authorization.k8s.io/v1/Role | `internal/controller/components/sparkoperator/sparkoperator_controller.go:47` |
| Owns | rbac.authorization.k8s.io/v1/Role | `internal/controller/components/dashboard/dashboard_controller.go:60` |
| Owns | rbac.authorization.k8s.io/v1/Role | `internal/controller/components/kueue/kueue_controller.go:63` |
| Owns | rbac.authorization.k8s.io/v1/Role | `internal/controller/components/workbenches/workbenches_controller.go:53` |
| Owns | rbac.authorization.k8s.io/v1/Role | `internal/controller/components/mlflowoperator/mlflowoperator_controller.go:51` |
| Owns | rbac.authorization.k8s.io/v1/Role | `internal/controller/components/ray/ray_controller.go:52` |
| Owns | rbac.authorization.k8s.io/v1/Role | `internal/controller/components/modelcontroller/modelcontroller_controller.go:58` |
| Owns | rbac.authorization.k8s.io/v1/Role | `internal/controller/services/monitoring/monitoring_controller.go:89` |
| Owns | rbac.authorization.k8s.io/v1/Role | `internal/controller/components/kserve/kserve_controller.go:61` |
| Owns | rbac.authorization.k8s.io/v1/Role | `internal/controller/services/auth/auth_controller.go:61` |
| Owns | rbac.authorization.k8s.io/v1/Role | `internal/controller/components/trustyai/trustyai_controller.go:67` |
| Owns | rbac.authorization.k8s.io/v1/Role | `internal/controller/components/feastoperator/feastoperator_controller.go:30` |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | `internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go:51` |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | `internal/controller/components/llamastackoperator/llamastackoperator_controller.go:30` |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | `internal/controller/components/workbenches/workbenches_controller.go:54` |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | `internal/controller/components/feastoperator/feastoperator_controller.go:29` |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | `internal/controller/components/modelcontroller/modelcontroller_controller.go:60` |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | `internal/controller/components/kueue/kueue_controller.go:64` |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | `internal/controller/components/trustyai/trustyai_controller.go:68` |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | `internal/controller/services/auth/auth_controller.go:62` |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | `internal/controller/components/kserve/kserve_controller.go:62` |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | `internal/controller/services/monitoring/monitoring_controller.go:90` |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | `internal/controller/components/modelregistry/modelregistry_controller.go:52` |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | `internal/controller/components/mlflowoperator/mlflowoperator_controller.go:52` |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | `internal/controller/components/ray/ray_controller.go:53` |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | `internal/controller/components/sparkoperator/sparkoperator_controller.go:46` |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | `internal/controller/components/dashboard/dashboard_controller.go:61` |
| Owns | route/v1/Route | `internal/controller/components/dashboard/dashboard_controller.go:69` |
| Owns | route/v1/Route | `internal/controller/services/monitoring/monitoring_controller.go:101` |
| Owns | security/v1/SecurityContextConstraints | `internal/controller/components/ray/ray_controller.go:57` |
| Owns | security/v1/SecurityContextConstraints | `internal/controller/components/datasciencepipelines/datasciencepipelines_controller.go:56` |
| Owns | template/v1/Template | `internal/controller/components/modelcontroller/modelcontroller_controller.go:64` |
| Watches | /v1/Namespace | `internal/controller/components/kueue/kueue_controller.go:124` |
| Watches | /v1/Namespace | `internal/controller/components/modelregistry/modelregistry_controller.go:67` |
| Watches | /v1/Namespace | `internal/controller/components/workbenches/workbenches_controller.go:66` |
| Watches | apis/v1/HTTPRoute | `internal/controller/components/mlflowoperator/mlflowoperator_controller.go:66` |
| Watches | rbac.authorization.k8s.io/v1/ClusterRole | `internal/controller/components/kueue/kueue_controller.go:118` |
| Watches | services/v1alpha1/Auth | `internal/controller/components/kueue/kueue_controller.go:135` |

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for opendatahub-operator

    participant KubernetesAPI as Kubernetes API
    participant odh_dashboard as odh-dashboard
    participant training_operator as training-operator
    participant azure_cloud_manager_operator as azure-cloud-manager-operator
    participant coreweave_cloud_manager_operator as coreweave-cloud-manager-operator
    participant controller_manager as controller-manager
    participant rhods_operator as rhods-operator
    participant kserve_controller_manager as kserve-controller-manager
    participant kserve_localmodel_controller_manager as kserve-localmodel-controller-manager
    participant odh_model_controller as odh-model-controller
    participant kuberay_operator as kuberay-operator
    participant deployment as deployment
    participant manager as manager

    odh_dashboard->>KubernetesAPI: Create/Update ConfigMap
    odh_dashboard->>KubernetesAPI: Create/Update Secret
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRole
    odh_dashboard->>KubernetesAPI: Create/Update Role
    odh_dashboard->>KubernetesAPI: Create/Update RoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update ServiceAccount
    odh_dashboard->>KubernetesAPI: Create/Update Service
    odh_dashboard->>KubernetesAPI: Create/Update Deployment
    odh_dashboard->>KubernetesAPI: Create/Update Route
    odh_dashboard->>KubernetesAPI: Create/Update HTTPRoute
    odh_dashboard->>KubernetesAPI: Create/Update ConsoleLink
    odh_dashboard->>KubernetesAPI: Create/Update NetworkPolicy
    odh_dashboard->>KubernetesAPI: Create/Update ConfigMap
    odh_dashboard->>KubernetesAPI: Create/Update Secret
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRole
    odh_dashboard->>KubernetesAPI: Create/Update Role
    odh_dashboard->>KubernetesAPI: Create/Update RoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update ServiceAccount
    odh_dashboard->>KubernetesAPI: Create/Update Service
    odh_dashboard->>KubernetesAPI: Create/Update ServiceMonitor
    odh_dashboard->>KubernetesAPI: Create/Update Deployment
    odh_dashboard->>KubernetesAPI: Create/Update SecurityContextConstraints
    odh_dashboard->>KubernetesAPI: Create/Update ConfigMap
    odh_dashboard->>KubernetesAPI: Create/Update RoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update Role
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRole
    odh_dashboard->>KubernetesAPI: Create/Update ServiceAccount
    odh_dashboard->>KubernetesAPI: Create/Update Service
    odh_dashboard->>KubernetesAPI: Create/Update Deployment
    odh_dashboard->>KubernetesAPI: Create/Update Secret
    odh_dashboard->>KubernetesAPI: Create/Update Service
    odh_dashboard->>KubernetesAPI: Create/Update ConfigMap
    odh_dashboard->>KubernetesAPI: Create/Update ServiceAccount
    odh_dashboard->>KubernetesAPI: Create/Update Role
    odh_dashboard->>KubernetesAPI: Create/Update RoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRole
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update NetworkPolicy
    odh_dashboard->>KubernetesAPI: Create/Update MutatingWebhookConfiguration
    odh_dashboard->>KubernetesAPI: Create/Update ValidatingWebhookConfiguration
    odh_dashboard->>KubernetesAPI: Create/Update ValidatingAdmissionPolicy
    odh_dashboard->>KubernetesAPI: Create/Update ValidatingAdmissionPolicyBinding
    odh_dashboard->>KubernetesAPI: Create/Update Deployment
    odh_dashboard->>KubernetesAPI: Create/Update ConfigMap
    odh_dashboard->>KubernetesAPI: Create/Update Secret
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRole
    odh_dashboard->>KubernetesAPI: Create/Update Role
    odh_dashboard->>KubernetesAPI: Create/Update RoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update ServiceAccount
    odh_dashboard->>KubernetesAPI: Create/Update Service
    odh_dashboard->>KubernetesAPI: Create/Update NetworkPolicy
    odh_dashboard->>KubernetesAPI: Create/Update PodMonitor
    odh_dashboard->>KubernetesAPI: Create/Update PrometheusRule
    odh_dashboard->>KubernetesAPI: Create/Update MutatingWebhookConfiguration
    odh_dashboard->>KubernetesAPI: Create/Update ValidatingWebhookConfiguration
    odh_dashboard->>KubernetesAPI: Create/Update Deployment
    KubernetesAPI-->>+odh_dashboard: Watch ClusterRole (informer)
    KubernetesAPI-->>+odh_dashboard: Watch Namespace (informer)
    KubernetesAPI-->>+odh_dashboard: Watch Auth (informer)
    odh_dashboard->>KubernetesAPI: Create/Update ConfigMap
    odh_dashboard->>KubernetesAPI: Create/Update RoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update Role
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRole
    odh_dashboard->>KubernetesAPI: Create/Update ServiceAccount
    odh_dashboard->>KubernetesAPI: Create/Update Service
    odh_dashboard->>KubernetesAPI: Create/Update PodDisruptionBudget
    odh_dashboard->>KubernetesAPI: Create/Update Deployment
    odh_dashboard->>KubernetesAPI: Create/Update ConfigMap
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRole
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update Role
    odh_dashboard->>KubernetesAPI: Create/Update RoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update Service
    odh_dashboard->>KubernetesAPI: Create/Update ServiceAccount
    odh_dashboard->>KubernetesAPI: Create/Update ConsoleLink
    odh_dashboard->>KubernetesAPI: Create/Update ServiceMonitor
    odh_dashboard->>KubernetesAPI: Create/Update Deployment
    KubernetesAPI-->>+odh_dashboard: Watch HTTPRoute (informer)
    odh_dashboard->>KubernetesAPI: Create/Update ConfigMap
    odh_dashboard->>KubernetesAPI: Create/Update ServiceAccount
    odh_dashboard->>KubernetesAPI: Create/Update Secret
    odh_dashboard->>KubernetesAPI: Create/Update ServiceMonitor
    odh_dashboard->>KubernetesAPI: Create/Update NetworkPolicy
    odh_dashboard->>KubernetesAPI: Create/Update Role
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRole
    odh_dashboard->>KubernetesAPI: Create/Update RoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update Service
    odh_dashboard->>KubernetesAPI: Create/Update ValidatingWebhookConfiguration
    odh_dashboard->>KubernetesAPI: Create/Update Template
    odh_dashboard->>KubernetesAPI: Create/Update Deployment
    odh_dashboard->>KubernetesAPI: Create/Update ConfigMap
    odh_dashboard->>KubernetesAPI: Create/Update ConfigMap
    odh_dashboard->>KubernetesAPI: Create/Update Secret
    odh_dashboard->>KubernetesAPI: Create/Update Role
    odh_dashboard->>KubernetesAPI: Create/Update RoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRole
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update Service
    odh_dashboard->>KubernetesAPI: Create/Update ServiceAccount
    odh_dashboard->>KubernetesAPI: Create/Update Deployment
    odh_dashboard->>KubernetesAPI: Create/Update MutatingWebhookConfiguration
    odh_dashboard->>KubernetesAPI: Create/Update ValidatingWebhookConfiguration
    KubernetesAPI-->>+odh_dashboard: Watch Namespace (informer)
    odh_dashboard->>KubernetesAPI: Create/Update ConfigMap
    odh_dashboard->>KubernetesAPI: Create/Update Service
    odh_dashboard->>KubernetesAPI: Create/Update ServiceAccount
    odh_dashboard->>KubernetesAPI: Create/Update Deployment
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRole
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update NetworkPolicy
    odh_dashboard->>KubernetesAPI: Create/Update HTTPRoute
    odh_dashboard->>KubernetesAPI: Create/Update Gateway
    odh_dashboard->>KubernetesAPI: Create/Update ConfigMap
    odh_dashboard->>KubernetesAPI: Create/Update Secret
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRole
    odh_dashboard->>KubernetesAPI: Create/Update Role
    odh_dashboard->>KubernetesAPI: Create/Update RoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update ServiceAccount
    odh_dashboard->>KubernetesAPI: Create/Update Service
    odh_dashboard->>KubernetesAPI: Create/Update Deployment
    odh_dashboard->>KubernetesAPI: Create/Update SecurityContextConstraints
    odh_dashboard->>KubernetesAPI: Create/Update ConfigMap
    odh_dashboard->>KubernetesAPI: Create/Update RoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update Role
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRole
    odh_dashboard->>KubernetesAPI: Create/Update ServiceAccount
    odh_dashboard->>KubernetesAPI: Create/Update Service
    odh_dashboard->>KubernetesAPI: Create/Update PodMonitor
    odh_dashboard->>KubernetesAPI: Create/Update MutatingWebhookConfiguration
    odh_dashboard->>KubernetesAPI: Create/Update ValidatingWebhookConfiguration
    odh_dashboard->>KubernetesAPI: Create/Update Deployment
    odh_dashboard->>KubernetesAPI: Create/Update ConfigMap
    odh_dashboard->>KubernetesAPI: Create/Update PodMonitor
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRole
    odh_dashboard->>KubernetesAPI: Create/Update ServiceAccount
    odh_dashboard->>KubernetesAPI: Create/Update Service
    odh_dashboard->>KubernetesAPI: Create/Update Deployment
    odh_dashboard->>KubernetesAPI: Create/Update ConfigMap
    odh_dashboard->>KubernetesAPI: Create/Update PodMonitor
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRole
    odh_dashboard->>KubernetesAPI: Create/Update ServiceAccount
    odh_dashboard->>KubernetesAPI: Create/Update Service
    odh_dashboard->>KubernetesAPI: Create/Update Deployment
    odh_dashboard->>KubernetesAPI: Create/Update ConfigMap
    odh_dashboard->>KubernetesAPI: Create/Update ServiceAccount
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRole
    odh_dashboard->>KubernetesAPI: Create/Update Role
    odh_dashboard->>KubernetesAPI: Create/Update RoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update Service
    odh_dashboard->>KubernetesAPI: Create/Update Deployment
    odh_dashboard->>KubernetesAPI: Create/Update ConfigMap
    odh_dashboard->>KubernetesAPI: Create/Update Secret
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRole
    odh_dashboard->>KubernetesAPI: Create/Update Role
    odh_dashboard->>KubernetesAPI: Create/Update RoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update ServiceAccount
    odh_dashboard->>KubernetesAPI: Create/Update Service
    odh_dashboard->>KubernetesAPI: Create/Update MutatingWebhookConfiguration
    odh_dashboard->>KubernetesAPI: Create/Update Deployment
    KubernetesAPI-->>+odh_dashboard: Watch Namespace (informer)
    odh_dashboard->>KubernetesAPI: Create/Update Dashboard
    odh_dashboard->>KubernetesAPI: Create/Update Workbenches
    odh_dashboard->>KubernetesAPI: Create/Update Ray
    odh_dashboard->>KubernetesAPI: Create/Update ModelRegistry
    odh_dashboard->>KubernetesAPI: Create/Update TrustyAI
    odh_dashboard->>KubernetesAPI: Create/Update Kueue
    odh_dashboard->>KubernetesAPI: Create/Update TrainingOperator
    odh_dashboard->>KubernetesAPI: Create/Update Trainer
    odh_dashboard->>KubernetesAPI: Create/Update DataSciencePipelines
    odh_dashboard->>KubernetesAPI: Create/Update Kserve
    odh_dashboard->>KubernetesAPI: Create/Update MaaSTenant
    odh_dashboard->>KubernetesAPI: Create/Update ModelController
    odh_dashboard->>KubernetesAPI: Create/Update FeastOperator
    odh_dashboard->>KubernetesAPI: Create/Update LlamaStackOperator
    odh_dashboard->>KubernetesAPI: Create/Update MLflowOperator
    odh_dashboard->>KubernetesAPI: Create/Update SparkOperator
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRole
    odh_dashboard->>KubernetesAPI: Create/Update Role
    odh_dashboard->>KubernetesAPI: Create/Update RoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update Role
    odh_dashboard->>KubernetesAPI: Create/Update RoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRole
    odh_dashboard->>KubernetesAPI: Create/Update ClusterRoleBinding
    odh_dashboard->>KubernetesAPI: Create/Update NetworkPolicy
    odh_dashboard->>KubernetesAPI: Create/Update Deployment
    odh_dashboard->>KubernetesAPI: Create/Update Job
    odh_dashboard->>KubernetesAPI: Create/Update ConfigMap
    odh_dashboard->>KubernetesAPI: Create/Update Secret
    odh_dashboard->>KubernetesAPI: Create/Update Service
    odh_dashboard->>KubernetesAPI: Create/Update ServiceAccount
    odh_dashboard->>KubernetesAPI: Create/Update Route

    Note over odh_dashboard: Exposed Services
    Note right of odh_dashboard: webhook-service:443/TCP []
    Note right of odh_dashboard: webhook-service:443/TCP []
    Note right of odh_dashboard: webhook-service:443/TCP []
    Note right of odh_dashboard: odh-dashboard:8443/TCP [dashboard-ui]
    Note right of odh_dashboard: kserve-controller-manager-service:8443/TCP []
    Note right of odh_dashboard: kserve-webhook-server-service:443/TCP []
    Note right of odh_dashboard: odh-model-controller-webhook-service:443/TCP []
    Note right of odh_dashboard: webhook-service:443/TCP []
    Note right of odh_dashboard: kuberay-operator:8080/TCP [monitoring-port]
    Note right of odh_dashboard: webhook-service:443/TCP []
    Note right of odh_dashboard: training-operator:8080/TCP [monitoring-port]
    Note right of odh_dashboard: training-operator:443/TCP [webhook-server]
    Note right of odh_dashboard: service:443/TCP []
    Note right of odh_dashboard: service:8080/TCP [metrics]
    Note right of odh_dashboard: webhook-service:443/TCP [webhook]

    Note over KubernetesAPI: Defined CRDs
    Note right of KubernetesAPI: Dashboard (components.platform.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: DataSciencePipelines (components.platform.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: FeastOperator (components.platform.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: Kserve (components.platform.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: Kueue (components.platform.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: LlamaStackOperator (components.platform.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: MLflowOperator (components.platform.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: ModelController (components.platform.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: ModelRegistry (components.platform.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: ModelsAsService (components.platform.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: Ray (components.platform.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: SparkOperator (components.platform.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: Trainer (components.platform.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: TrainingOperator (components.platform.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: TrustyAI (components.platform.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: Workbenches (components.platform.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: AzureKubernetesEngine (infrastructure.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: CoreWeaveKubernetesEngine (infrastructure.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: HardwareProfile (infrastructure.opendatahub.io/v1)
    Note right of KubernetesAPI: HardwareProfile (infrastructure.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: Auth (services.platform.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: GatewayConfig (services.platform.opendatahub.io/v1alpha1)
    Note right of KubernetesAPI: Monitoring (services.platform.opendatahub.io/v1alpha1)
```

### Webhooks

| Name | Type | Path | Failure Policy | Service | Source |
|------|------|------|----------------|---------|--------|
| mraycluster.kb.io | mutating | /mutate-ray-io-v1-raycluster | Fail | $(namespace)/kuberay-webhook-service | `opt/manifests/ray/openshift/webhook.yaml` |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

### ConfigMaps

| Name | Data Keys | Source |
|------|-----------|--------|
| blackbox | blackbox.yml | `config/monitoring/blackbox-exporter/external/blackbox-exporter-external-configmap.yaml` |
| blackbox | blackbox.yml | `config/monitoring/blackbox-exporter/internal/blackbox-exporter-internal-configmap.yaml` |
| workflow-controller-configmap |  | `opt/manifests/datasciencepipelines/argo/configmap.workflow-controller-configmap.yaml` |

### Helm

**Chart:** test-chart v0.1.0

