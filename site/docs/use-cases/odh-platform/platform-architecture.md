# Platform Architecture

## CRD Ownership Map

The platform defines 94 CRDs. Each CRD is owned by the component that declares it.

| Owner | CRDs | Count |
|-------|------|-------|
| **codeflare-operator** | AppWrapper | 1 |
| **data-science-pipelines** | CompositeController, ControllerRevision, DecoratorController | 3 |
| **data-science-pipelines-operator** | DataSciencePipelinesApplication, Pipeline, PipelineVersion, ScheduledWorkflow | 4 |
| **kserve** | ClusterServingRuntime, ClusterStorageContainer, InferenceGraph, InferencePool, InferenceService, LLMInferenceService, LLMInferenceServiceConfig, LocalModelCache, LocalModelNamespaceCache, LocalModelNode, LocalModelNodeGroup, ServingRuntime, TrainedModel | 13 |
| **kserve-autogluon-server** | ClusterServingRuntime, ClusterStorageContainer, InferenceGraph, InferenceService, LLMInferenceService, LLMInferenceServiceConfig, LocalModelCache, LocalModelNamespaceCache, LocalModelNode, LocalModelNodeGroup, ServingRuntime, TrainedModel | 12 |
| **kueue** | ClusterQueue, LocalQueue | 2 |
| **llm-d-inference-scheduler** | InferenceModelRewrite, InferenceObjective | 2 |
| **mlflow-operator** | MLflow, MLflowConfig, MLflowOperator | 3 |
| **model-registry-operator** | ModelRegistry | 1 |
| **modelmesh-serving** | ClusterServingRuntime, InferenceService, Predictor, ServingRuntime | 4 |
| **odh-model-controller** | Account | 1 |
| **ogx-k8s-operator** | LlamaStackDistribution, OGXServer | 2 |
| **opendatahub-operator** | FeatureTracker | 1 |
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
| kserve | kserve-autogluon-server | watches-crd:InferenceGraph |
| kserve | kserve-autogluon-server | watches-crd:LocalModelCache |
| kserve | kserve-autogluon-server | watches-crd:LocalModelNamespaceCache |
| kserve | kserve-autogluon-server | watches-crd:LocalModelNode |
| kserve | kserve-autogluon-server | watches-crd:TrainedModel |
| kserve | kserve-autogluon-server | watches-crd:LLMInferenceService |
| kserve | kserve-autogluon-server | watches-crd:LLMInferenceServiceConfig |
| kserve | kserve-autogluon-server | watches-crd:InferenceService |
| kubeflow | data-science-pipelines-operator | go-module |
| kubeflow | operator-chaos | go-module |
| llm-d-inference-scheduler | kserve | watches-crd:InferencePool |
| mlflow-operator | mlflow-operator | go-module |
| model-registry-operator | odh-platform-utilities | go-module |
| model-registry-operator | operator-chaos | go-module |
| model-registry | kserve-autogluon-server | watches-crd:InferenceService |
| modelmesh-serving | kserve-autogluon-server | watches-crd:ServingRuntime |
| odh-dashboard | mlflow-go | go-module |
| odh-dashboard | odh-dashboard | go-module |
| odh-dashboard | odh-platform-utilities | go-module |
| odh-dashboard | ogx-k8s-operator | go-module |
| odh-model-controller | kserve | go-module |
| odh-model-controller | kserve-autogluon-server | watches-crd:InferenceGraph |
| odh-model-controller | kserve-autogluon-server | watches-crd:ServingRuntime |
| odh-model-controller | kserve-autogluon-server | watches-crd:LLMInferenceService |
| odh-model-controller | kserve-autogluon-server | watches-crd:InferenceService |
| ogx-k8s-operator | odh-platform-utilities | go-module |
| opendatahub-operator | models-as-a-service | go-module |
| opendatahub-operator | odh-platform-utilities | go-module |
| opendatahub-operator | opendatahub-operator | go-module |
| spark-operator | odh-platform-utilities | go-module |
| trustyai-service-operator | odh-platform-utilities | go-module |
| workload-variant-autoscaler | kserve | watches-crd:InferencePool |
| codeflare-operator | opendatahub-operator | webhook-ref |
| kserve-autogluon-server | kserve | webhook-ref |
| model-registry-operator | opendatahub-operator | webhook-ref |
| modelmesh-serving | opendatahub-operator | webhook-ref |

**Tightest coupling:** `kserve -> kserve-autogluon-server` (8 dependency edges).

## Webhooks

**Total webhooks across platform**: 116

| Component | Webhooks |
|-----------|----------|
| agents-operator | 2 |
| codeflare-operator | 4 |
| data-science-pipelines-operator | 1 |
| kserve | 21 |
| kserve-autogluon-server | 21 |
| kubeflow | 2 |
| kuberay | 4 |
| kueue | 20 |
| llm-d-inference-scheduler | 5 |
| model-registry-operator | 4 |
| modelmesh-serving | 2 |
| models-as-a-service | 4 |
| odh-model-controller | 8 |
| ogx-k8s-operator | 1 |
| opendatahub-operator | 4 |
| spark-operator | 8 |
| trainer | 4 |
| training-operator | 1 |

### Cross-Component Webhooks

Webhooks whose service reference points to a different component:

| Webhook | Type | Owner | Target Component | Target Type | Path |
|---------|------|-------|------------------|-------------|------|
| mappwrapper.kb.io | mutating | codeflare-operator | opendatahub-operator |  | /mutate-workload-codeflare-dev-v1beta2-appwrapper |
| mraycluster.ray.openshift.ai | mutating | codeflare-operator | opendatahub-operator | rayClusterWebhook | /mutate-ray-io-v1-raycluster |
| vappwrapper.kb.io | validating | codeflare-operator | opendatahub-operator |  | /validate-workload-codeflare-dev-v1beta2-appwrapper |
| vraycluster.ray.openshift.ai | validating | codeflare-operator | opendatahub-operator | rayClusterWebhook | /validate-ray-io-v1-raycluster |
| clusterservingruntime.kserve-webhook-server.validator | validating | kserve-autogluon-server | kserve | ServingRuntimeValidator | /validate-serving-kserve-io-v1alpha1-clusterservingruntime |
| conversion-unknown | conversion | kserve-autogluon-server | kserve |  | /convert |
| inferencegraph.kserve-webhook-server.validator | validating | kserve-autogluon-server | kserve |  | /validate-serving-kserve-io-v1alpha1-inferencegraph |
| inferenceservice.kserve-webhook-server.defaulter | mutating | kserve-autogluon-server | kserve |  | /mutate-serving-kserve-io-v1beta1-inferenceservice |
| inferenceservice.kserve-webhook-server.pod-mutator | mutating | kserve-autogluon-server | kserve | Mutator | /mutate-pods |
| inferenceservice.kserve-webhook-server.validator | validating | kserve-autogluon-server | kserve |  | /validate-serving-kserve-io-v1beta1-inferenceservice |
| llminferenceservice.kserve-webhook-server.v1alpha1.defaulter | mutating | kserve-autogluon-server | kserve | LLMInferenceServiceDefaulterV1Alpha1 | /mutate-serving-kserve-io-v1alpha1-llminferenceservice |
| llminferenceservice.kserve-webhook-server.v1alpha1.validator | validating | kserve-autogluon-server | kserve | LLMInferenceServiceValidator | /validate-serving-kserve-io-v1alpha1-llminferenceservice |
| llminferenceservice.kserve-webhook-server.v1alpha2.defaulter | mutating | kserve-autogluon-server | kserve | LLMInferenceServiceDefaulterV1Alpha2 | /mutate-serving-kserve-io-v1alpha2-llminferenceservice |
| llminferenceservice.kserve-webhook-server.v1alpha2.validator | validating | kserve-autogluon-server | kserve | LLMInferenceServiceValidator | /validate-serving-kserve-io-v1alpha2-llminferenceservice |
| llminferenceserviceconfig.kserve-webhook-server.v1alpha1.validator | validating | kserve-autogluon-server | kserve | LLMInferenceServiceConfigValidator | /validate-serving-kserve-io-v1alpha1-llminferenceserviceconfig |
| llminferenceserviceconfig.kserve-webhook-server.v1alpha2.validator | validating | kserve-autogluon-server | kserve | LLMInferenceServiceConfigValidator | /validate-serving-kserve-io-v1alpha2-llminferenceserviceconfig |
| localmodelcache.kserve-webhook-server.validator | validating | kserve-autogluon-server | kserve |  | /validate-serving-kserve-io-v1alpha1-localmodelcache |
| servingruntime.kserve-webhook-server.validator | validating | kserve-autogluon-server | kserve | ServingRuntimeValidator | /validate-serving-kserve-io-v1alpha1-servingruntime |
| trainedmodel.kserve-webhook-server.validator | validating | kserve-autogluon-server | kserve |  | /validate-serving-kserve-io-v1alpha1-trainedmodel |
| conversion-unknown | conversion | model-registry-operator | opendatahub-operator |  | /convert |
| conversion-unknown | conversion | modelmesh-serving | opendatahub-operator |  | /convert |

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
| llminferenceservice.kserve-webhook-server.v1alpha1.validator | kserve-autogluon-server | spec | invalid |  |
| llminferenceservice.kserve-webhook-server.v1alpha1.validator | kserve-autogluon-server | worker | invalid |  |
| llminferenceservice.kserve-webhook-server.v1alpha1.validator | kserve-autogluon-server | dataLocal | invalid |  |
| llminferenceservice.kserve-webhook-server.v1alpha1.validator | kserve-autogluon-server | data | invalid |  |
| llminferenceservice.kserve-webhook-server.v1alpha1.validator | kserve-autogluon-server | pipeline | invalid |  |
| llminferenceservice.kserve-webhook-server.v1alpha1.validator | kserve-autogluon-server | replicas | invalid |  |
| llminferenceservice.kserve-webhook-server.v1alpha1.validator | kserve-autogluon-server | inline | invalid |  |
| llminferenceservice.kserve-webhook-server.v1alpha1.validator | kserve-autogluon-server | ref.name | invalid |  |
| llminferenceservice.kserve-webhook-server.v1alpha2.validator | kserve-autogluon-server | spec | invalid |  |
| llminferenceservice.kserve-webhook-server.v1alpha2.validator | kserve-autogluon-server | worker | invalid |  |
| llminferenceservice.kserve-webhook-server.v1alpha2.validator | kserve-autogluon-server | dataLocal | invalid |  |
| llminferenceservice.kserve-webhook-server.v1alpha2.validator | kserve-autogluon-server | data | invalid |  |
| llminferenceservice.kserve-webhook-server.v1alpha2.validator | kserve-autogluon-server | pipeline | invalid |  |
| llminferenceservice.kserve-webhook-server.v1alpha2.validator | kserve-autogluon-server | replicas | invalid |  |
| llminferenceservice.kserve-webhook-server.v1alpha2.validator | kserve-autogluon-server | inline | invalid |  |
| llminferenceservice.kserve-webhook-server.v1alpha2.validator | kserve-autogluon-server | ref.name | invalid |  |
| llminferenceservice.kserve-webhook-server.v1alpha2.validator | kserve-autogluon-server | maxRank | invalid |  |
| llminferenceservice.kserve-webhook-server.v1alpha2.validator | kserve-autogluon-server | maxAdapters | invalid |  |
| llminferenceservice.kserve-webhook-server.v1alpha2.validator | kserve-autogluon-server | maxCpuAdapters | invalid |  |
| llminferenceserviceconfig.kserve-webhook-server.v1alpha1.validator | kserve-autogluon-server | spec.baseRefs | forbidden |  |
| llminferenceserviceconfig.kserve-webhook-server.v1alpha1.validator | kserve-autogluon-server | replicas | invalid |  |
| llminferenceserviceconfig.kserve-webhook-server.v1alpha2.validator | kserve-autogluon-server | spec.baseRefs | forbidden |  |
| llminferenceserviceconfig.kserve-webhook-server.v1alpha2.validator | kserve-autogluon-server | replicas | invalid |  |

### External Webhooks

| Webhook | Type | Owner | Target Type | Path | Failure Policy |
|---------|------|-------|-------------|------|----------------|
| mraycluster.kb.io | mutating | kuberay |  | /mutate-ray-io-v1-raycluster | Fail |

