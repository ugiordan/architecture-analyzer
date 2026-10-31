# kserve: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

| Type | GVK | Source |
|------|-----|--------|
| For | apis/v1alpha1/Kserve | [`kserve-module/pkg/kservemodule/setup.go:172`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/pkg/kservemodule/setup.go#L172) |
| For | serving/v1alpha1/InferenceGraph | [`pkg/controller/v1alpha1/inferencegraph/controller.go:385`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/inferencegraph/controller.go#L385) |
| For | serving/v1alpha1/KernelCacheNodeGroup | [`pkg/controller/v1alpha1/kernelcache/reconcilers/node_reconciler.go:160`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/kernelcache/reconcilers/node_reconciler.go#L160) |
| For | serving/v1alpha1/LocalModelCache | [`pkg/controller/v1alpha1/localmodel/reconcilers/localmodelcache_reconciler.go:340`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/localmodel/reconcilers/localmodelcache_reconciler.go#L340) |
| For | serving/v1alpha1/LocalModelNamespaceCache | [`pkg/controller/v1alpha1/localmodel/reconcilers/localmodelnamespacecache_reconciler.go:431`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/localmodel/reconcilers/localmodelnamespacecache_reconciler.go#L431) |
| For | serving/v1alpha1/LocalModelNode | [`pkg/controller/v1alpha1/localmodelnode/controller.go:646`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/localmodelnode/controller.go#L646) |
| For | serving/v1alpha1/TrainedModel | [`pkg/controller/v1alpha1/trainedmodel/controller.go:306`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/trainedmodel/controller.go#L306) |
| For | serving/v1alpha2/LLMInferenceService | [`pkg/controller/v1alpha2/llmisvc/controller.go:456`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L456) |
| For | serving/v1alpha2/LLMInferenceServiceConfig | [`pkg/controller/v1alpha2/llmisvc/llmisvcconfig_controller.go:241`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/llmisvcconfig_controller.go#L241) |
| For | serving/v1beta1/InferenceService | [`pkg/controller/v1alpha1/kernelcache/reconcilers/capture_reconciler.go:307`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/kernelcache/reconcilers/capture_reconciler.go#L307) |
| For | serving/v1beta1/InferenceService | [`pkg/controller/v1beta1/inferenceservice/controller.go:685`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1beta1/inferenceservice/controller.go#L685) |
| For | serving/v1beta1/InferenceService | [`pkg/controller/v1alpha1/kernelcache/reconcilers/capture_controller.go:562`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/kernelcache/reconcilers/capture_controller.go#L562) |
| Owns | /v1/ConfigMap | [`kserve-module/pkg/kservemodule/setup.go:173`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/pkg/kservemodule/setup.go#L173) |
| Owns | /v1/PersistentVolume | [`pkg/controller/v1alpha1/localmodel/reconcilers/localmodelcache_reconciler.go:341`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/localmodel/reconcilers/localmodelcache_reconciler.go#L341) |
| Owns | /v1/PersistentVolume | [`kserve-module/pkg/kservemodule/setup.go:177`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/pkg/kservemodule/setup.go#L177) |
| Owns | /v1/PersistentVolumeClaim | [`pkg/controller/v1alpha1/localmodel/reconcilers/localmodelcache_reconciler.go:342`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/localmodel/reconcilers/localmodelcache_reconciler.go#L342) |
| Owns | /v1/PersistentVolumeClaim | [`pkg/controller/v1alpha1/localmodel/reconcilers/localmodelnamespacecache_reconciler.go:432`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/localmodel/reconcilers/localmodelnamespacecache_reconciler.go#L432) |
| Owns | /v1/PersistentVolumeClaim | [`kserve-module/pkg/kservemodule/setup.go:178`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/pkg/kservemodule/setup.go#L178) |
| Owns | /v1/Secret | [`kserve-module/pkg/kservemodule/setup.go:174`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/pkg/kservemodule/setup.go#L174) |
| Owns | /v1/Secret | [`pkg/controller/v1alpha2/llmisvc/controller.go:460`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L460) |
| Owns | /v1/Service | [`pkg/controller/v1alpha2/llmisvc/controller.go:461`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L461) |
| Owns | /v1/Service | [`kserve-module/pkg/kservemodule/setup.go:175`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/pkg/kservemodule/setup.go#L175) |
| Owns | /v1/Service | [`pkg/controller/v1beta1/inferenceservice/controller.go:687`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1beta1/inferenceservice/controller.go#L687) |
| Owns | /v1/ServiceAccount | [`kserve-module/pkg/kservemodule/setup.go:176`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/pkg/kservemodule/setup.go#L176) |
| Owns | admissionregistration.k8s.io/v1/MutatingWebhookConfiguration | [`kserve-module/pkg/kservemodule/setup.go:186`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/pkg/kservemodule/setup.go#L186) |
| Owns | admissionregistration.k8s.io/v1/ValidatingWebhookConfiguration | [`kserve-module/pkg/kservemodule/setup.go:187`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/pkg/kservemodule/setup.go#L187) |
| Owns | api/v1/InferencePool | [`pkg/controller/v1alpha2/llmisvc/controller.go:491`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L491) |
| Owns | apis/v1/HTTPRoute | [`pkg/controller/v1beta1/inferenceservice/controller.go:731`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1beta1/inferenceservice/controller.go#L731) |
| Owns | apis/v1/HTTPRoute | [`pkg/controller/v1alpha2/llmisvc/controller.go:483`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L483) |
| Owns | apis/v1beta1/OpenTelemetryCollector | [`pkg/controller/v1beta1/inferenceservice/controller.go:713`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1beta1/inferenceservice/controller.go#L713) |
| Owns | apps/v1/DaemonSet | [`kserve-module/pkg/kservemodule/setup.go:180`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/pkg/kservemodule/setup.go#L180) |
| Owns | apps/v1/Deployment | [`pkg/controller/v1alpha1/inferencegraph/controller.go:386`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/inferencegraph/controller.go#L386) |
| Owns | apps/v1/Deployment | [`pkg/controller/v1alpha2/llmisvc/controller.go:459`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L459) |
| Owns | apps/v1/Deployment | [`kserve-module/pkg/kservemodule/setup.go:179`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/pkg/kservemodule/setup.go#L179) |
| Owns | apps/v1/Deployment | [`pkg/controller/v1beta1/inferenceservice/controller.go:686`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1beta1/inferenceservice/controller.go#L686) |
| Owns | autoscaling/v2/HorizontalPodAutoscaler | [`pkg/controller/v1alpha2/llmisvc/controller.go:462`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L462) |
| Owns | batch/v1/Job | [`pkg/controller/v1alpha1/localmodel/reconcilers/localmodelnamespacecache_reconciler.go:434`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/localmodel/reconcilers/localmodelnamespacecache_reconciler.go#L434) |
| Owns | batch/v1/Job | [`pkg/controller/v1alpha1/localmodelnode/controller.go:647`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/localmodelnode/controller.go#L647) |
| Owns | disaggregatedset/v1/DisaggregatedSet | [`pkg/controller/v1alpha2/llmisvc/controller.go:510`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L510) |
| Owns | gie/v1alpha2pool/InferencePool | [`pkg/controller/v1alpha2/llmisvc/controller.go:496`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L496) |
| Owns | keda/v1alpha1/ScaledObject | [`pkg/controller/v1alpha2/llmisvc/controller.go:501`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L501) |
| Owns | keda/v1alpha1/ScaledObject | [`pkg/controller/v1beta1/inferenceservice/controller.go:696`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1beta1/inferenceservice/controller.go#L696) |
| Owns | leaderworkerset/v1/LeaderWorkerSet | [`pkg/controller/v1alpha2/llmisvc/controller.go:505`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L505) |
| Owns | networking.k8s.io/v1/Ingress | [`pkg/controller/v1beta1/inferenceservice/controller.go:737`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1beta1/inferenceservice/controller.go#L737) |
| Owns | networking.k8s.io/v1/Ingress | [`pkg/controller/v1alpha2/llmisvc/controller.go:458`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L458) |
| Owns | networking.k8s.io/v1/NetworkPolicy | [`kserve-module/pkg/kservemodule/setup.go:181`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/pkg/kservemodule/setup.go#L181) |
| Owns | networking/v1beta1/VirtualService | [`pkg/controller/v1beta1/inferenceservice/controller.go:719`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1beta1/inferenceservice/controller.go#L719) |
| Owns | rbac.authorization.k8s.io/v1/ClusterRole | [`kserve-module/pkg/kservemodule/setup.go:184`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/pkg/kservemodule/setup.go#L184) |
| Owns | rbac.authorization.k8s.io/v1/ClusterRoleBinding | [`kserve-module/pkg/kservemodule/setup.go:185`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/pkg/kservemodule/setup.go#L185) |
| Owns | rbac.authorization.k8s.io/v1/Role | [`kserve-module/pkg/kservemodule/setup.go:182`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/pkg/kservemodule/setup.go#L182) |
| Owns | rbac.authorization.k8s.io/v1/RoleBinding | [`kserve-module/pkg/kservemodule/setup.go:183`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/pkg/kservemodule/setup.go#L183) |
| Owns | resource/v1/ResourceClaimTemplate | [`pkg/controller/v1alpha2/llmisvc/controller.go:521`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L521) |
| Owns | security/v1/SecurityContextConstraints | [`kserve-module/pkg/kservemodule/setup.go:233`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/pkg/kservemodule/setup.go#L233) |
| Owns | serving/v1/Service | [`pkg/controller/v1beta1/inferenceservice/controller.go:690`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1beta1/inferenceservice/controller.go#L690) |
| Owns | serving/v1/Service | [`pkg/controller/v1alpha1/inferencegraph/controller.go:393`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/inferencegraph/controller.go#L393) |
| Watches | /v1/ConfigMap | [`pkg/controller/v1alpha2/llmisvc/controller.go:464`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L464) |
| Watches | /v1/ConfigMap | [`pkg/controller/v1alpha2/llmisvc/controller.go:463`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L463) |
| Watches | /v1/ConfigMap | [`pkg/controller/v1alpha1/kernelcache/reconcilers/capture_controller.go:563`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/kernelcache/reconcilers/capture_controller.go#L563) |
| Watches | /v1/ConfigMap | [`pkg/controller/v1alpha1/kernelcache/reconcilers/node_reconciler.go:162`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/kernelcache/reconcilers/node_reconciler.go#L162) |
| Watches | /v1/ConfigMap | [`kserve-module/pkg/kservemodule/setup.go:193`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/pkg/kservemodule/setup.go#L193) |
| Watches | /v1/Node | [`pkg/controller/v1alpha1/kernelcache/reconcilers/node_reconciler.go:161`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/kernelcache/reconcilers/node_reconciler.go#L161) |
| Watches | /v1/Node | [`pkg/controller/v1alpha1/localmodel/reconcilers/localmodelcache_reconciler.go:371`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/localmodel/reconcilers/localmodelcache_reconciler.go#L371) |
| Watches | /v1/Node | [`pkg/controller/v1alpha1/localmodel/reconcilers/localmodelnamespacecache_reconciler.go:464`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/localmodel/reconcilers/localmodelnamespacecache_reconciler.go#L464) |
| Watches | /v1/Node | [`kserve-module/pkg/kservemodule/setup.go:203`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/pkg/kservemodule/setup.go#L203) |
| Watches | /v1/PersistentVolumeClaim | [`pkg/controller/v1alpha1/localmodel/reconcilers/localmodelnamespacecache_reconciler.go:466`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/localmodel/reconcilers/localmodelnamespacecache_reconciler.go#L466) |
| Watches | /v1/Pod | [`pkg/controller/v1alpha1/kernelcache/reconcilers/capture_controller.go:574`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/kernelcache/reconcilers/capture_controller.go#L574) |
| Watches | /v1/Pod | [`pkg/controller/v1alpha2/llmisvc/controller.go:465`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L465) |
| Watches | /v1/Pod | [`pkg/controller/v1beta1/inferenceservice/controller.go:741`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1beta1/inferenceservice/controller.go#L741) |
| Watches | /v1/Pod | [`pkg/controller/v1alpha1/kernelcache/reconcilers/capture_reconciler.go:309`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/kernelcache/reconcilers/capture_reconciler.go#L309) |
| Watches | /v1/ServiceAccount | [`pkg/controller/v1alpha2/llmisvc/controller.go:468`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L468) |
| Watches | api/v1/InferencePool | [`pkg/controller/v1alpha2/llmisvc/controller.go:492`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L492) |
| Watches | apis/v1/Gateway | [`pkg/controller/v1alpha2/llmisvc/controller.go:487`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L487) |
| Watches | apis/v1/HTTPRoute | [`pkg/controller/v1alpha2/llmisvc/controller.go:484`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L484) |
| Watches | disaggregatedset/v1/DisaggregatedSetRoleScaler | [`pkg/controller/v1alpha2/llmisvc/controller.go:515`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L515) |
| Watches | gie/v1alpha2pool/InferencePool | [`pkg/controller/v1alpha2/llmisvc/controller.go:497`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L497) |
| Watches | node/v1/RuntimeClass | [`kserve-module/pkg/kservemodule/setup.go:210`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/pkg/kservemodule/setup.go#L210) |
| Watches | rbac.authorization.k8s.io/v1/RoleBinding | [`pkg/controller/v1alpha1/kernelcache/reconcilers/capture_controller.go:564`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/kernelcache/reconcilers/capture_controller.go#L564) |
| Watches | serving/v1alpha1/ClusterServingRuntime | [`pkg/controller/v1beta1/inferenceservice/controller.go:748`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1beta1/inferenceservice/controller.go#L748) |
| Watches | serving/v1alpha1/ClusterServingRuntime | [`pkg/controller/v1alpha2/llmisvc/controller.go:531`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L531) |
| Watches | serving/v1alpha1/KernelCacheCapture | [`pkg/controller/v1alpha1/kernelcache/reconcilers/capture_reconciler.go:308`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/kernelcache/reconcilers/capture_reconciler.go#L308) |
| Watches | serving/v1alpha1/LocalModelNamespaceCache | [`pkg/controller/v1alpha1/localmodel/reconcilers/localmodelnamespacecache_reconciler.go:463`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/localmodel/reconcilers/localmodelnamespacecache_reconciler.go#L463) |
| Watches | serving/v1alpha1/LocalModelNode | [`pkg/controller/v1alpha1/localmodel/reconcilers/localmodelcache_reconciler.go:373`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/localmodel/reconcilers/localmodelcache_reconciler.go#L373) |
| Watches | serving/v1alpha1/LocalModelNode | [`pkg/controller/v1alpha1/localmodel/reconcilers/localmodelnamespacecache_reconciler.go:465`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/localmodel/reconcilers/localmodelnamespacecache_reconciler.go#L465) |
| Watches | serving/v1alpha1/ServingRuntime | [`pkg/controller/v1alpha2/llmisvc/controller.go:536`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L536) |
| Watches | serving/v1alpha1/ServingRuntime | [`pkg/controller/v1beta1/inferenceservice/controller.go:740`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1beta1/inferenceservice/controller.go#L740) |
| Watches | serving/v1alpha2/LLMInferenceService | [`pkg/controller/v1alpha1/localmodel/reconcilers/localmodelcache_reconciler.go:366`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/localmodel/reconcilers/localmodelcache_reconciler.go#L366) |
| Watches | serving/v1alpha2/LLMInferenceService | [`pkg/controller/v1alpha2/llmisvc/llmisvcconfig_controller.go:242`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/llmisvcconfig_controller.go#L242) |
| Watches | serving/v1alpha2/LLMInferenceService | [`pkg/controller/v1alpha1/kernelcache/reconcilers/capture_controller.go:582`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/kernelcache/reconcilers/capture_controller.go#L582) |
| Watches | serving/v1alpha2/LLMInferenceService | [`pkg/controller/v1alpha1/localmodel/reconcilers/localmodelnamespacecache_reconciler.go:458`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/localmodel/reconcilers/localmodelnamespacecache_reconciler.go#L458) |
| Watches | serving/v1alpha2/LLMInferenceServiceConfig | [`pkg/controller/v1alpha2/llmisvc/controller.go:457`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/controller.go#L457) |
| Watches | serving/v1beta1/InferenceService | [`pkg/controller/v1alpha1/localmodel/reconcilers/localmodelcache_reconciler.go:364`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/localmodel/reconcilers/localmodelcache_reconciler.go#L364) |
| Watches | serving/v1beta1/InferenceService | [`pkg/controller/v1alpha1/localmodel/reconcilers/localmodelnamespacecache_reconciler.go:456`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha1/localmodel/reconcilers/localmodelnamespacecache_reconciler.go#L456) |

### Programmatic Resource Operations

| Verb | Kind | Group | Condition |
|------|------|-------|----------|
| update | KernelCacheNode | serving |  |
| create | ConfigMap |  |  |
| update | ConfigMap |  |  |
| update | LLMInferenceService | serving |  |
| create | LLMInferenceService | serving |  |
| update | LLMInferenceServiceConfig | serving |  |
| patch | HTTPRoute | apis |  |
| create | KernelCache | serving |  |
| delete | RoleBinding | rbac.authorization.k8s.io |  |
| delete | KernelCacheCapture | serving |  |
| update | KernelCacheCapture | serving |  |
| create | KernelCacheNode | serving |  |
| delete | KernelCacheNode | serving |  |
| patch | DaemonSet | apps |  |
| create | RoleBinding | rbac.authorization.k8s.io |  |
| patch | RoleBinding | rbac.authorization.k8s.io |  |
| update | KernelCache | serving |  |
| patch | LocalModelCache | serving |  |
| patch | LocalModelNamespaceCache | serving |  |
| create | Deployment | apps |  |
| patch | Deployment | apps |  |
| delete | Deployment | apps |  |
| delete | HTTPRoute | apis |  |
| create | HTTPRoute | apis |  |
| update | HTTPRoute | apis |  |
| create | VirtualService | networking |  |
| update | VirtualService | networking |  |
| delete | VirtualService | networking |  |
| create | Service |  |  |
| update | Service |  |  |
| delete | Service |  |  |
| delete | Ingress | networking.k8s.io |  |
| create | Ingress | networking.k8s.io |  |
| update | Ingress | networking.k8s.io |  |
| patch | Service |  |  |
| create | HorizontalPodAutoscaler | autoscaling |  |
| update | HorizontalPodAutoscaler | autoscaling |  |
| delete | HorizontalPodAutoscaler | autoscaling |  |
| delete | ScaledObject | keda |  |
| create | ScaledObject | keda |  |
| update | ScaledObject | keda |  |
| delete | OpenTelemetryCollector | apis |  |
| create | OpenTelemetryCollector | apis |  |
| update | OpenTelemetryCollector | apis |  |
| delete | Service | serving |  |
| update | Service | serving |  |
| create | Service | serving |  |
| delete | TrainedModel | serving |  |
| update | TrainedModel | serving |  |
| patch | InferenceService | serving |  |

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for kserve

    participant KubernetesAPI as Kubernetes API
    participant kserve_controller_manager as kserve-controller-manager
    participant kserve_localmodel_controller_manager as kserve-localmodel-controller-manager
    participant kserve_module_controller_manager as kserve-module-controller-manager
    participant llmisvc_controller_manager as llmisvc-controller-manager
    participant modelexpress_operator as modelexpress-operator
    participant odh_model_controller as odh-model-controller

    KubernetesAPI->>+kserve_controller_manager: Watch Kserve (reconcile)
    KubernetesAPI->>+kserve_controller_manager: Watch InferenceGraph (reconcile)
    KubernetesAPI->>+kserve_controller_manager: Watch KernelCacheNodeGroup (reconcile)
    KubernetesAPI->>+kserve_controller_manager: Watch LocalModelCache (reconcile)
    KubernetesAPI->>+kserve_controller_manager: Watch LocalModelNamespaceCache (reconcile)
    KubernetesAPI->>+kserve_controller_manager: Watch LocalModelNode (reconcile)
    KubernetesAPI->>+kserve_controller_manager: Watch TrainedModel (reconcile)
    KubernetesAPI->>+kserve_controller_manager: Watch LLMInferenceService (reconcile)
    KubernetesAPI->>+kserve_controller_manager: Watch LLMInferenceServiceConfig (reconcile)
    KubernetesAPI->>+kserve_controller_manager: Watch InferenceService (reconcile)
    KubernetesAPI->>+kserve_controller_manager: Watch InferenceService (reconcile)
    KubernetesAPI->>+kserve_controller_manager: Watch InferenceService (reconcile)
    kserve_controller_manager->>KubernetesAPI: Create/Update ConfigMap
    kserve_controller_manager->>KubernetesAPI: Create/Update PersistentVolume
    kserve_controller_manager->>KubernetesAPI: Create/Update PersistentVolume
    kserve_controller_manager->>KubernetesAPI: Create/Update PersistentVolumeClaim
    kserve_controller_manager->>KubernetesAPI: Create/Update PersistentVolumeClaim
    kserve_controller_manager->>KubernetesAPI: Create/Update PersistentVolumeClaim
    kserve_controller_manager->>KubernetesAPI: Create/Update Secret
    kserve_controller_manager->>KubernetesAPI: Create/Update Secret
    kserve_controller_manager->>KubernetesAPI: Create/Update Service
    kserve_controller_manager->>KubernetesAPI: Create/Update Service
    kserve_controller_manager->>KubernetesAPI: Create/Update Service
    kserve_controller_manager->>KubernetesAPI: Create/Update ServiceAccount
    kserve_controller_manager->>KubernetesAPI: Create/Update MutatingWebhookConfiguration
    kserve_controller_manager->>KubernetesAPI: Create/Update ValidatingWebhookConfiguration
    kserve_controller_manager->>KubernetesAPI: Create/Update InferencePool
    kserve_controller_manager->>KubernetesAPI: Create/Update HTTPRoute
    kserve_controller_manager->>KubernetesAPI: Create/Update HTTPRoute
    kserve_controller_manager->>KubernetesAPI: Create/Update OpenTelemetryCollector
    kserve_controller_manager->>KubernetesAPI: Create/Update DaemonSet
    kserve_controller_manager->>KubernetesAPI: Create/Update Deployment
    kserve_controller_manager->>KubernetesAPI: Create/Update Deployment
    kserve_controller_manager->>KubernetesAPI: Create/Update Deployment
    kserve_controller_manager->>KubernetesAPI: Create/Update Deployment
    kserve_controller_manager->>KubernetesAPI: Create/Update HorizontalPodAutoscaler
    kserve_controller_manager->>KubernetesAPI: Create/Update Job
    kserve_controller_manager->>KubernetesAPI: Create/Update Job
    kserve_controller_manager->>KubernetesAPI: Create/Update DisaggregatedSet
    kserve_controller_manager->>KubernetesAPI: Create/Update InferencePool
    kserve_controller_manager->>KubernetesAPI: Create/Update ScaledObject
    kserve_controller_manager->>KubernetesAPI: Create/Update ScaledObject
    kserve_controller_manager->>KubernetesAPI: Create/Update LeaderWorkerSet
    kserve_controller_manager->>KubernetesAPI: Create/Update Ingress
    kserve_controller_manager->>KubernetesAPI: Create/Update Ingress
    kserve_controller_manager->>KubernetesAPI: Create/Update NetworkPolicy
    kserve_controller_manager->>KubernetesAPI: Create/Update VirtualService
    kserve_controller_manager->>KubernetesAPI: Create/Update ClusterRole
    kserve_controller_manager->>KubernetesAPI: Create/Update ClusterRoleBinding
    kserve_controller_manager->>KubernetesAPI: Create/Update Role
    kserve_controller_manager->>KubernetesAPI: Create/Update RoleBinding
    kserve_controller_manager->>KubernetesAPI: Create/Update ResourceClaimTemplate
    kserve_controller_manager->>KubernetesAPI: Create/Update SecurityContextConstraints
    kserve_controller_manager->>KubernetesAPI: Create/Update Service
    kserve_controller_manager->>KubernetesAPI: Create/Update Service
    KubernetesAPI-->>+kserve_controller_manager: Watch ConfigMap (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch ConfigMap (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch ConfigMap (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch ConfigMap (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch ConfigMap (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch Node (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch Node (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch Node (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch Node (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch PersistentVolumeClaim (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch Pod (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch Pod (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch Pod (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch Pod (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch ServiceAccount (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch InferencePool (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch Gateway (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch HTTPRoute (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch DisaggregatedSetRoleScaler (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch InferencePool (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch RuntimeClass (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch RoleBinding (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch ClusterServingRuntime (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch ClusterServingRuntime (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch KernelCacheCapture (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch LocalModelNamespaceCache (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch LocalModelNode (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch LocalModelNode (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch ServingRuntime (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch ServingRuntime (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch LLMInferenceService (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch LLMInferenceService (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch LLMInferenceService (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch LLMInferenceService (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch LLMInferenceServiceConfig (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch InferenceService (informer)
    KubernetesAPI-->>+kserve_controller_manager: Watch InferenceService (informer)

    Note over kserve_controller_manager: Exposed Services
    Note right of kserve_controller_manager: kserve-controller-manager-metrics-service:8443/TCP [https]
    Note right of kserve_controller_manager: kserve-controller-manager-service:8443/TCP []
    Note right of kserve_controller_manager: kserve-webhook-server-service:443/TCP []
    Note right of kserve_controller_manager: llmisvc-controller-manager-service:8443/TCP [https]
    Note right of kserve_controller_manager: llmisvc-webhook-server-service:443/TCP [https]
    Note right of kserve_controller_manager: localmodel-webhook-server-service:443/TCP []
    Note right of kserve_controller_manager: model-serving-api:443/TCP [https]
    Note right of kserve_controller_manager: model-serving-api:8080/TCP [metrics]
    Note right of kserve_controller_manager: modelexpress-operator-metrics:8080/TCP [http-metrics]
    Note right of kserve_controller_manager: modelexpress-operator-metrics:8443/TCP [https-metrics]
    Note right of kserve_controller_manager: odh-model-controller-webhook-service:443/TCP []

    Note over KubernetesAPI: Defined CRDs
    Note right of KubernetesAPI: ClusterServingRuntime (/v1alpha1)
    Note right of KubernetesAPI: ClusterStorageContainer (/v1alpha1)
    Note right of KubernetesAPI: InferenceGraph (/v1alpha1)
    Note right of KubernetesAPI: KernelCache (/v1alpha1)
    Note right of KubernetesAPI: KernelCacheCapture (/v1alpha1)
    Note right of KubernetesAPI: KernelCacheNode (/v1alpha1)
    Note right of KubernetesAPI: KernelCacheNodeGroup (/v1alpha1)
    Note right of KubernetesAPI: LLMInferenceService (/v1alpha1)
    Note right of KubernetesAPI: LLMInferenceServiceConfig (/v1alpha1)
    Note right of KubernetesAPI: LocalModelCache (/v1alpha1)
    Note right of KubernetesAPI: LocalModelNamespaceCache (/v1alpha1)
    Note right of KubernetesAPI: LocalModelNode (/v1alpha1)
    Note right of KubernetesAPI: LocalModelNodeGroup (/v1alpha1)
    Note right of KubernetesAPI: ServingRuntime (/v1alpha1)
    Note right of KubernetesAPI: TrainedModel (/v1alpha1)
    Note right of KubernetesAPI: LLMInferenceService (/v1alpha2)
    Note right of KubernetesAPI: LLMInferenceServiceConfig (/v1alpha2)
    Note right of KubernetesAPI: InferenceService (/v1beta1)
    Note right of KubernetesAPI: InferencePool (inference.networking.x-k8s.io/v1alpha2pool)
    Note right of KubernetesAPI: ClusterServingRuntime (serving.kserve.io/v1alpha1)
    Note right of KubernetesAPI: ClusterStorageContainer (serving.kserve.io/v1alpha1)
    Note right of KubernetesAPI: InferenceGraph (serving.kserve.io/v1alpha1)
    Note right of KubernetesAPI: KernelCache (serving.kserve.io/v1alpha1)
    Note right of KubernetesAPI: KernelCacheCapture (serving.kserve.io/v1alpha1)
    Note right of KubernetesAPI: KernelCacheNode (serving.kserve.io/v1alpha1)
    Note right of KubernetesAPI: KernelCacheNodeGroup (serving.kserve.io/v1alpha1)
    Note right of KubernetesAPI: LocalModelCache (serving.kserve.io/v1alpha1)
    Note right of KubernetesAPI: LocalModelNamespaceCache (serving.kserve.io/v1alpha1)
    Note right of KubernetesAPI: LocalModelNode (serving.kserve.io/v1alpha1)
    Note right of KubernetesAPI: LocalModelNodeGroup (serving.kserve.io/v1alpha1)
    Note right of KubernetesAPI: ServingRuntime (serving.kserve.io/v1alpha1)
    Note right of KubernetesAPI: TrainedModel (serving.kserve.io/v1alpha1)
    Note right of KubernetesAPI: LLMInferenceService (serving.kserve.io/v1alpha2)
    Note right of KubernetesAPI: LLMInferenceServiceConfig (serving.kserve.io/v1alpha2)
    Note right of KubernetesAPI: InferenceService (serving.kserve.io/v1beta1)
```

### Webhooks

| Name | Type | Path | Failure Policy | Service | Overlays | Enable Condition | Sources |
|------|------|------|----------------|---------|----------|------------------|----------|
| InferenceGraphValidator-webhook | validating | /validate-inferencegraph |  |  |  |  |  |
| InferenceServiceDefaulter-webhook | mutating | /mutate-inferenceservices |  |  |  |  |  |
| InferenceServiceValidator-webhook | validating | /validate-inferenceservices |  |  |  |  |  |
| LocalModelCacheValidator-webhook | validating | /validate-localmodelcaches |  |  |  |  |  |
| LocalModelNamespaceCacheValidator-webhook | validating | /validate-localmodelnamespacecaches |  |  |  |  |  |
| PodMutator-webhook | mutating | /mutate-kernelcache-pods |  |  |  |  |  |
| TrainedModelValidator-webhook | validating | /validate-trainedmodel |  |  |  |  |  |
| clusterservingruntime.kserve-webhook-server.validator | validating | /validate-serving-kserve-io-v1alpha1-clusterservingruntime | Fail | opendatahub/kserve-webhook-server-service | config/overlays/odh |  | [`config/webhook/manifests.yaml`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/config/webhook/manifests.yaml), [`kustomize:config/overlays/odh (clusterservingruntime.serving.kserve.io)`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kustomize:config/overlays/odh %28clusterservingruntime.serving.kserve.io%29) |
| conversion-unknown | conversion | /convert |  | kserve/llmisvc-webhook-server-service |  |  | [`config/crd/full/llmisvc/llmisvc_conversion_webhook_patch.yaml`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/config/crd/full/llmisvc/llmisvc_conversion_webhook_patch.yaml) |
| conversion-unknown | conversion | /convert |  | kserve/llmisvc-webhook-server-service |  |  | [`config/crd/full/llmisvc/llmisvcconfig_conversion_webhook_patch.yaml`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/config/crd/full/llmisvc/llmisvcconfig_conversion_webhook_patch.yaml) |
| conversion-unknown | conversion | /convert |  | kserve/llmisvc-webhook-server-service |  |  | [`config/crd/minimal/llmisvc/llmisvc_conversion_webhook_patch.yaml`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/config/crd/minimal/llmisvc/llmisvc_conversion_webhook_patch.yaml) |
| conversion-unknown | conversion | /convert |  | kserve/llmisvc-webhook-server-service |  |  | [`config/crd/minimal/llmisvc/llmisvcconfig_conversion_webhook_patch.yaml`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/config/crd/minimal/llmisvc/llmisvcconfig_conversion_webhook_patch.yaml) |
| inferencegraph.kserve-webhook-server.validator | validating | /validate-serving-kserve-io-v1alpha1-inferencegraph | Fail | opendatahub/kserve-webhook-server-service | config/overlays/odh |  | [`config/webhook/manifests.yaml`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/config/webhook/manifests.yaml), [`kustomize:config/overlays/odh (inferencegraph.serving.kserve.io)`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kustomize:config/overlays/odh %28inferencegraph.serving.kserve.io%29) |
| inferenceservice.kserve-webhook-server.defaulter | mutating | /mutate-serving-kserve-io-v1beta1-inferenceservice | Fail | opendatahub/kserve-webhook-server-service | config/overlays/odh |  | [`config/webhook/manifests.yaml`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/config/webhook/manifests.yaml), [`kustomize:config/overlays/odh (inferenceservice.serving.kserve.io)`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kustomize:config/overlays/odh %28inferenceservice.serving.kserve.io%29) |
| inferenceservice.kserve-webhook-server.pod-mutator | mutating | /mutate-pods | Fail | opendatahub/kserve-webhook-server-service | config/overlays/odh |  | [`config/webhook/manifests.yaml`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/config/webhook/manifests.yaml), [`kustomize:config/overlays/odh (inferenceservice.serving.kserve.io)`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kustomize:config/overlays/odh %28inferenceservice.serving.kserve.io%29) |
| inferenceservice.kserve-webhook-server.validator | validating | /validate-serving-kserve-io-v1beta1-inferenceservice | Fail | opendatahub/kserve-webhook-server-service | config/overlays/odh |  | [`config/webhook/manifests.yaml`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/config/webhook/manifests.yaml), [`kustomize:config/overlays/odh (inferenceservice.serving.kserve.io)`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kustomize:config/overlays/odh %28inferenceservice.serving.kserve.io%29) |
| llminferenceservice.kserve-webhook-server.v1alpha1.defaulter | mutating | /mutate-serving-kserve-io-v1alpha1-llminferenceservice | Fail | opendatahub/llmisvc-webhook-server-service | config/overlays/odh |  | [`kustomize:config/overlays/odh (llminferenceservice.serving.kserve.io)`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kustomize:config/overlays/odh %28llminferenceservice.serving.kserve.io%29) |
| llminferenceservice.kserve-webhook-server.v1alpha1.validator | validating | /validate-serving-kserve-io-v1alpha1-llminferenceservice | Fail | opendatahub/llmisvc-webhook-server-service | config/overlays/odh |  | [`kustomize:config/overlays/odh (llminferenceservice.serving.kserve.io)`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kustomize:config/overlays/odh %28llminferenceservice.serving.kserve.io%29) |
| llminferenceservice.kserve-webhook-server.v1alpha2.defaulter | mutating | /mutate-serving-kserve-io-v1alpha2-llminferenceservice | Fail | opendatahub/llmisvc-webhook-server-service | config/overlays/odh |  | [`kustomize:config/overlays/odh (llminferenceservice.serving.kserve.io)`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kustomize:config/overlays/odh %28llminferenceservice.serving.kserve.io%29) |
| llminferenceservice.kserve-webhook-server.v1alpha2.validator | validating | /validate-serving-kserve-io-v1alpha2-llminferenceservice | Fail | opendatahub/llmisvc-webhook-server-service | config/overlays/odh |  | [`kustomize:config/overlays/odh (llminferenceservice.serving.kserve.io)`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kustomize:config/overlays/odh %28llminferenceservice.serving.kserve.io%29) |
| llminferenceserviceconfig.kserve-webhook-server.v1alpha1.validator | validating | /validate-serving-kserve-io-v1alpha1-llminferenceserviceconfig | Fail | opendatahub/llmisvc-webhook-server-service | config/overlays/odh |  | [`kustomize:config/overlays/odh (llminferenceserviceconfig.serving.kserve.io)`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kustomize:config/overlays/odh %28llminferenceserviceconfig.serving.kserve.io%29) |
| llminferenceserviceconfig.kserve-webhook-server.v1alpha2.validator | validating | /validate-serving-kserve-io-v1alpha2-llminferenceserviceconfig | Fail | opendatahub/llmisvc-webhook-server-service | config/overlays/odh |  | [`kustomize:config/overlays/odh (llminferenceserviceconfig.serving.kserve.io)`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kustomize:config/overlays/odh %28llminferenceserviceconfig.serving.kserve.io%29) |
| localmodelcache.kserve-webhook-server.validator | validating |  |  |  |  |  | [`config/localmodels/webhook_cainjection_patch.yaml`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/config/localmodels/webhook_cainjection_patch.yaml) |
| servingruntime.kserve-webhook-server.validator | validating | /validate-serving-kserve-io-v1alpha1-servingruntime | Fail | opendatahub/kserve-webhook-server-service | config/overlays/odh |  | [`config/webhook/manifests.yaml`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/config/webhook/manifests.yaml), [`kustomize:config/overlays/odh (servingruntime.serving.kserve.io)`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kustomize:config/overlays/odh %28servingruntime.serving.kserve.io%29) |
| trainedmodel.kserve-webhook-server.validator | validating | /validate-serving-kserve-io-v1alpha1-trainedmodel | Fail | opendatahub/kserve-webhook-server-service | config/overlays/odh |  | [`config/webhook/manifests.yaml`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/config/webhook/manifests.yaml), [`kustomize:config/overlays/odh (trainedmodel.serving.kserve.io)`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kustomize:config/overlays/odh %28trainedmodel.serving.kserve.io%29) |

#### llminferenceservice.kserve-webhook-server.v1alpha1.validator Behavior

| Field | Operation | Condition |
|-------|-----------|----------|
| spec | invalid |  |
| group | invalid |  |
| group | required |  |
| weight | required |  |
| worker | invalid |  |
| dataLocal | invalid |  |
| data | invalid |  |
| pipeline | invalid |  |
| replicas | invalid |  |
| inline | invalid |  |
| ref.name | invalid |  |
| scaling.wva | forbidden |  |

#### llminferenceservice.kserve-webhook-server.v1alpha2.validator Behavior

| Field | Operation | Condition |
|-------|-----------|----------|
| spec | invalid |  |
| group | invalid |  |
| group | required |  |
| weight | required |  |
| worker | invalid |  |
| dataLocal | invalid |  |
| data | invalid |  |
| pipeline | invalid |  |
| replicas | invalid |  |
| inline | invalid |  |
| ref.name | invalid |  |
| scaling.wva | forbidden |  |

#### llminferenceserviceconfig.kserve-webhook-server.v1alpha1.validator Behavior

| Field | Operation | Condition |
|-------|-----------|----------|
| spec.baseRefs | forbidden |  |
| replicas | invalid |  |

#### llminferenceserviceconfig.kserve-webhook-server.v1alpha2.validator Behavior

| Field | Operation | Condition |
|-------|-----------|----------|
| spec.baseRefs | forbidden |  |
| replicas | invalid |  |

### HTTP Endpoints

| Method | Path | Source |
|--------|------|--------|
| * | / | [`cmd/router/main.go:675`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/cmd/router/main.go#L675) |
| * | gateway.networking.k8s.io | [`pkg/controller/v1alpha2/llmisvc/config_merge.go:1127`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/config_merge.go#L1127) |
| * | gateway.networking.k8s.io | [`pkg/controller/v1alpha2/llmisvc/fixture/gwapi_builders.go:213`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/fixture/gwapi_builders.go#L213) |
| * | gateway.networking.k8s.io | [`pkg/controller/v1alpha2/llmisvc/fixture/gwapi_builders.go:231`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/fixture/gwapi_builders.go#L231) |
| * | gateway.networking.k8s.io | [`pkg/controller/v1alpha2/llmisvc/fixture/gwapi_builders.go:433`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/fixture/gwapi_builders.go#L433) |
| * | gateway.networking.k8s.io | [`pkg/controller/v1alpha2/llmisvc/fixture/gwapi_builders.go:740`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/fixture/gwapi_builders.go#L740) |
| * | inference.networking.k8s.io | [`pkg/controller/v1alpha2/llmisvc/fixture/gwapi_builders.go:293`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/fixture/gwapi_builders.go#L293) |
| * | inference.networking.k8s.io | [`pkg/controller/v1alpha2/llmisvc/router.go:816`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/router.go#L816) |
| * | inference.networking.x-k8s.io | [`pkg/controller/v1alpha2/llmisvc/fixture/gwapi_builders.go:307`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/pkg/controller/v1alpha2/llmisvc/fixture/gwapi_builders.go#L307) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

### ConfigMaps

| Name | Data Keys | Source |
|------|-----------|--------|
| inferenceservice-config | _example, agent, autoscaler, batcher, credentials, deploy, explainers, inferenceService, ingress, kernelcache, localModel, logger, metricsAggregator, opentelemetryCollector, router, security, storageInitializer | [`charts/kserve-llmisvc-resources/files/common/configmap.yaml`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/charts/kserve-llmisvc-resources/files/common/configmap.yaml) |
| inferenceservice-config | _example, agent, autoscaler, batcher, credentials, deploy, explainers, inferenceService, ingress, kernelcache, localModel, logger, metricsAggregator, opentelemetryCollector, router, security, storageInitializer | [`charts/kserve-resources/files/common/configmap.yaml`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/charts/kserve-resources/files/common/configmap.yaml) |
| manager-config | config.yaml | [`kserve-module/prefetched-manifests-rhoai/wva/components/openshift/configmap-patch.yaml`](https://github.com/red-hat-data-services/kserve/blob/4db69f8e7e04d9babc0ac88d3b691bdadacdb860/kserve-module/prefetched-manifests-rhoai/wva/components/openshift/configmap-patch.yaml) |

### Helm

**Chart:** kserve-crd vv0.21.0

