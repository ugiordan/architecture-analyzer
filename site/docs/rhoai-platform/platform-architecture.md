# Platform Architecture

## CRD Ownership Map

The platform defines 50 CRDs. Each CRD is owned by exactly one component.

| Owner | CRDs |
|-------|------|
| **data-science-pipelines-operator** | DataSciencePipelinesApplication, Pipeline, PipelineVersion, ScheduledWorkflow |
| **kserve** | ClusterServingRuntime, ClusterStorageContainer, InferenceGraph, InferenceService, LLMInferenceService, LLMInferenceServiceConfig, LocalModelCache, LocalModelNamespaceCache, LocalModelNode, LocalModelNodeGroup, ServingRuntime, TrainedModel |
| **odh-model-controller** | Account |
| **opendatahub-operator** | Auth, AzureKubernetesEngine, CoreWeaveKubernetesEngine, Dashboard, DataSciencePipelines, FeastOperator, GatewayConfig, HardwareProfile, Kserve, Kueue, LlamaStackOperator, MLflowOperator, ModelController, ModelRegistry, ModelsAsService, Monitoring, Ray, SparkOperator, Trainer, TrainingOperator, TrustyAI, Workbenches |
| **trustyai-service-operator** | EvalHub, GuardrailsOrchestrator, LMEvalJob, NemoGuardrails, TrustyAIService |

## Cross-Component Dependencies

Relationships detected through Go module imports and CRD watch patterns.

| From | To | Type |
|------|----|------|
| odh-dashboard | llama-stack-k8s-operator | go-module |
| odh-dashboard | mlflow-go | go-module |
| odh-model-controller | kserve | go-module |
| odh-model-controller | kserve | watches-crd:InferenceGraph |
| odh-model-controller | kserve | watches-crd:InferenceService |
| odh-model-controller | kserve | watches-crd:LLMInferenceService |
| odh-model-controller | kserve | watches-crd:ServingRuntime |
| opendatahub-operator | models-as-a-service | go-module |
| opendatahub-operator | opendatahub-operator | go-module |

**Tightest coupling:** `odh-model-controller -> kserve` (5 dependency edges).

