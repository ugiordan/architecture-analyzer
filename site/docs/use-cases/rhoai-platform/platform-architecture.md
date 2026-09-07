# Platform Architecture

## CRD Ownership Map

The platform defines 70 CRDs. Each CRD is owned by the component that declares it.

| Owner | CRDs | Count |
|-------|------|-------|
| **ai-gateway-payload-processing** | ExternalModel, ExternalProvider | 2 |
| **codeflare-operator** | AppWrapper | 1 |
| **data-science-pipelines** | CompositeController, ControllerRevision, DecoratorController | 3 |
| **data-science-pipelines-operator** | DataSciencePipelinesApplication, Pipeline, PipelineVersion, ScheduledWorkflow | 4 |
| **kserve** | ClusterServingRuntime, ClusterStorageContainer, InferenceGraph, InferencePool, InferenceService, LLMInferenceService, LLMInferenceServiceConfig, LocalModelCache, LocalModelNamespaceCache, LocalModelNode, LocalModelNodeGroup, ServingRuntime, TrainedModel | 13 |
| **kueue** | ClusterQueue, LocalQueue | 2 |
| **llm-d-inference-scheduler** | InferenceModelRewrite, InferenceObjective | 2 |
| **mlflow-operator** | MLflow, MLflowConfig, MLflowOperator | 3 |
| **model-registry-operator** | ModelRegistry | 1 |
| **modelmesh-serving** | ClusterServingRuntime, InferenceService, Predictor, ServingRuntime | 4 |
| **odh-model-controller** | Account | 1 |
| **ogx-k8s-operator** | LlamaStackDistribution, OGXServer | 2 |
| **rhods-operator** | FeatureTracker | 1 |
| **spark-operator** | ScheduledSparkApplication, SparkApplication, SparkConnect | 3 |
| **trainer** | ClusterTrainingRuntime, TrainJob, TrainingRuntime | 3 |
| **training-operator** | JAXJob, MPIJob, PaddleJob, PyTorchJob, TFJob, XGBoostJob | 6 |
| **trustyai-service-operator** | EvalHub, LMEvalJob, TrustyAIService | 3 |

## Cross-Component Dependencies

Relationships detected through Go module imports and CRD watch patterns.

| From | To | Type |
|------|----|------|
| codeflare-operator | opendatahub-operator | go-module |
| data-science-pipelines-operator | mlflow-operator | go-module |
| data-science-pipelines-operator | operator-chaos | go-module |
| kserve | odh-platform-utilities | go-module |
| kubeflow | data-science-pipelines-operator | go-module |
| kubeflow | operator-chaos | go-module |
| llm-d-inference-scheduler | kserve | watches-crd:InferencePool |
| mlflow-operator | mlflow-operator | go-module |
| model-registry-operator | odh-platform-utilities | go-module |
| model-registry-operator | operator-chaos | go-module |
| model-registry | kserve | watches-crd:InferenceService |
| modelmesh-serving | kserve | watches-crd:ServingRuntime |
| models-as-a-service | ai-gateway-payload-processing | watches-crd:ExternalModel |
| odh-cli | opendatahub-operator | go-module |
| odh-dashboard | mlflow-go | go-module |
| odh-dashboard | odh-dashboard | go-module |
| odh-dashboard | odh-platform-utilities | go-module |
| odh-dashboard | ogx-k8s-operator | go-module |
| odh-model-controller | kserve | go-module |
| odh-model-controller | kserve | watches-crd:InferenceGraph |
| odh-model-controller | kserve | watches-crd:ServingRuntime |
| odh-model-controller | kserve | watches-crd:LLMInferenceService |
| odh-model-controller | kserve | watches-crd:InferenceService |
| ogx-k8s-operator | odh-platform-utilities | go-module |
| rhods-operator | models-as-a-service | go-module |
| rhods-operator | odh-platform-utilities | go-module |
| rhods-operator | rhods-operator | go-module |
| spark-operator | odh-platform-utilities | go-module |
| trustyai-service-operator | odh-platform-utilities | go-module |
| workload-variant-autoscaler | kserve | watches-crd:InferencePool |
| codeflare-operator | rhods-operator | webhook-ref |
| model-registry-operator | rhods-operator | webhook-ref |
| modelmesh-serving | rhods-operator | webhook-ref |

**Tightest coupling:** `odh-model-controller -> kserve` (5 dependency edges).

## Webhooks

**Total webhooks across platform**: 95

| Component | Webhooks |
|-----------|----------|
| agents-operator | 2 |
| codeflare-operator | 4 |
| data-science-pipelines-operator | 1 |
| kserve | 21 |
| kubeflow | 2 |
| kuberay | 4 |
| kueue | 20 |
| llm-d-inference-scheduler | 5 |
| model-registry-operator | 4 |
| modelmesh-serving | 2 |
| models-as-a-service | 4 |
| odh-model-controller | 8 |
| ogx-k8s-operator | 1 |
| rhods-operator | 4 |
| spark-operator | 8 |
| trainer | 4 |
| training-operator | 1 |

### Cross-Component Webhooks

Webhooks whose service reference points to a different component:

| Webhook | Type | Owner | Target Component | Target Type | Path |
|---------|------|-------|------------------|-------------|------|
| mappwrapper.kb.io | mutating | codeflare-operator | rhods-operator |  | /mutate-workload-codeflare-dev-v1beta2-appwrapper |
| mraycluster.ray.openshift.ai | mutating | codeflare-operator | rhods-operator | rayClusterWebhook | /mutate-ray-io-v1-raycluster |
| vappwrapper.kb.io | validating | codeflare-operator | rhods-operator |  | /validate-workload-codeflare-dev-v1beta2-appwrapper |
| vraycluster.ray.openshift.ai | validating | codeflare-operator | rhods-operator | rayClusterWebhook | /validate-ray-io-v1-raycluster |
| conversion-unknown | conversion | model-registry-operator | rhods-operator |  | /convert |
| conversion-unknown | conversion | modelmesh-serving | rhods-operator |  | /convert |

#### Webhook Behavioral Analysis

Field-level operations extracted from Go AST analysis of webhook handlers:

| Webhook | Owner | Field | Operation | Condition |
|---------|-------|-------|-----------|----------|
| mraycluster.ray.openshift.ai | codeflare-operator | spec.headGroupSpec.template.spec.containers | set | ptr.Deref(...) |
| mraycluster.ray.openshift.ai | codeflare-operator | spec.headGroupSpec.template.spec.volumes | set | ptr.Deref(...) |
| mraycluster.ray.openshift.ai | codeflare-operator | spec.headGroupSpec.template.spec.serviceAccountName | set | ptr.Deref(...) |
| mraycluster.ray.openshift.ai | codeflare-operator | spec.headGroupSpec.template.spec.initContainers | set | ptr.Deref(...) |
| mraycluster.ray.openshift.ai | codeflare-operator | template.spec.volumes | set | ptr.Deref(...) |
| mraycluster.ray.openshift.ai | codeflare-operator | template.spec.initContainers | set | ptr.Deref(...) |

### External Webhooks

| Webhook | Type | Owner | Target Type | Path | Failure Policy |
|---------|------|-------|-------------|------|----------------|
| mraycluster.kb.io | mutating | kuberay |  | /mutate-ray-io-v1-raycluster | Fail |

