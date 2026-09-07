# Network Topology

83 Kubernetes services across the platform.

## Network Topology Graph

Interactive service mesh view of the platform. Drag nodes to rearrange, hover to highlight connections, click for details. Double-click background to fit all.

<div class="topology-toolbar">
  <button data-action="fit" title="Fit to view">Fit</button>
  <button data-action="zoom-in" title="Zoom in">+</button>
  <button data-action="zoom-out" title="Zoom out">&minus;</button>
  <button data-action="relayout" title="Re-run layout">Relayout</button>
  <button data-action="fullscreen" title="Toggle fullscreen">Fullscreen</button>
</div>
<div class="cytoscape-topology">
  <script type="application/json">
  {"components":[{"id":"agents_operator","name":"agents-operator","serviceCount":2,"netpolCount":1,"hasIngress":true},{"id":"ai4rag","name":"ai4rag","serviceCount":0,"netpolCount":0,"hasIngress":false},{"id":"argo_workflows","name":"argo-workflows","serviceCount":0,"netpolCount":0,"hasIngress":false},{"id":"batch_gateway","name":"batch-gateway","serviceCount":0,"netpolCount":0,"hasIngress":false},{"id":"codeflare_operator","name":"codeflare-operator","serviceCount":1,"netpolCount":0,"hasIngress":true},{"id":"data_science_pipelines","name":"data-science-pipelines","serviceCount":1,"netpolCount":1,"hasIngress":true},{"id":"data_science_pipelines_operator","name":"data-science-pipelines-operator","serviceCount":7,"netpolCount":2,"hasIngress":true},{"id":"distributed_workloads","name":"distributed-workloads","serviceCount":0,"netpolCount":0,"hasIngress":false},{"id":"eval_hub","name":"eval-hub","serviceCount":0,"netpolCount":0,"hasIngress":false},{"id":"feast","name":"feast","serviceCount":1,"netpolCount":0,"hasIngress":false},{"id":"fms_guardrails_orchestrator","name":"fms-guardrails-orchestrator","serviceCount":0,"netpolCount":0,"hasIngress":false},{"id":"guardrails_detectors","name":"guardrails-detectors","serviceCount":0,"netpolCount":0,"hasIngress":false},{"id":"kserve","name":"kserve","serviceCount":6,"netpolCount":2,"hasIngress":true},{"id":"kserve_autogluon_server","name":"kserve-autogluon-server","serviceCount":6,"netpolCount":0,"hasIngress":true},{"id":"kube_auth_proxy","name":"kube-auth-proxy","serviceCount":0,"netpolCount":0,"hasIngress":false},{"id":"kube_rbac_proxy","name":"kube-rbac-proxy","serviceCount":0,"netpolCount":0,"hasIngress":false},{"id":"kubeflow","name":"kubeflow","serviceCount":2,"netpolCount":0,"hasIngress":false},{"id":"kuberay","name":"kuberay","serviceCount":2,"netpolCount":0,"hasIngress":false},{"id":"kueue","name":"kueue","serviceCount":2,"netpolCount":0,"hasIngress":false},{"id":"llama_stack","name":"llama-stack","serviceCount":1,"netpolCount":0,"hasIngress":false},{"id":"llama_stack_provider_trustyai_garak","name":"llama-stack-provider-trustyai-garak","serviceCount":0,"netpolCount":0,"hasIngress":false},{"id":"llm_d","name":"llm-d","serviceCount":3,"netpolCount":0,"hasIngress":true},{"id":"llm_d_inference_scheduler","name":"llm-d-inference-scheduler","serviceCount":6,"netpolCount":0,"hasIngress":true},{"id":"llm_d_kv_cache","name":"llm-d-kv-cache","serviceCount":0,"netpolCount":0,"hasIngress":false},{"id":"lm_evaluation_harness","name":"lm-evaluation-harness","serviceCount":0,"netpolCount":0,"hasIngress":false},{"id":"mlflow","name":"mlflow","serviceCount":1,"netpolCount":0,"hasIngress":false},{"id":"mlflow_operator","name":"mlflow-operator","serviceCount":1,"netpolCount":0,"hasIngress":true},{"id":"model_registry","name":"model-registry","serviceCount":1,"netpolCount":0,"hasIngress":false},{"id":"model_registry_operator","name":"model-registry-operator","serviceCount":4,"netpolCount":4,"hasIngress":true},{"id":"modelmesh","name":"modelmesh","serviceCount":1,"netpolCount":1,"hasIngress":false},{"id":"modelmesh_serving","name":"modelmesh-serving","serviceCount":3,"netpolCount":4,"hasIngress":false},{"id":"models_as_a_service","name":"models-as-a-service","serviceCount":3,"netpolCount":10,"hasIngress":true},{"id":"notebooks","name":"notebooks","serviceCount":1,"netpolCount":0,"hasIngress":false},{"id":"odh_dashboard","name":"odh-dashboard","serviceCount":16,"netpolCount":17,"hasIngress":true},{"id":"odh_model_controller","name":"odh-model-controller","serviceCount":3,"netpolCount":1,"hasIngress":true},{"id":"ogx_k8s_operator","name":"ogx-k8s-operator","serviceCount":2,"netpolCount":1,"hasIngress":true},{"id":"opendatahub_operator","name":"opendatahub-operator","serviceCount":1,"netpolCount":2,"hasIngress":false},{"id":"spark_operator","name":"spark-operator","serviceCount":1,"netpolCount":1,"hasIngress":true},{"id":"trainer","name":"trainer","serviceCount":0,"netpolCount":0,"hasIngress":false},{"id":"training_operator","name":"training-operator","serviceCount":1,"netpolCount":0,"hasIngress":false},{"id":"trustyai_service_operator","name":"trustyai-service-operator","serviceCount":3,"netpolCount":0,"hasIngress":false},{"id":"workload_variant_autoscaler","name":"workload-variant-autoscaler","serviceCount":1,"netpolCount":0,"hasIngress":false}],"services":[{"id":"agents_operator_svc_0","name":"bundle-service","parent":"agents_operator","ports":"8080"},{"id":"agents_operator_svc_1","name":"webhook-service","parent":"agents_operator","ports":"443"},{"id":"codeflare_operator_svc_0","name":"webhook-service","parent":"codeflare_operator","ports":"443"},{"id":"data_science_pipelines_svc_0","name":"kubeflow-pipelines-profile-controller","parent":"data_science_pipelines","ports":"80"},{"id":"data_science_pipelines_operator_svc_0","name":"data-science-pipelines-operator-service","parent":"data_science_pipelines_operator","ports":"8080"},{"id":"data_science_pipelines_operator_svc_1","name":"ds-pipeline-workflow-controller-metrics-template-value","parent":"data_science_pipelines_operator","ports":"9090"},{"id":"data_science_pipelines_operator_svc_2","name":"mariadb-template-value","parent":"data_science_pipelines_operator","ports":"3306"},{"id":"data_science_pipelines_operator_svc_3","name":"minio-service","parent":"data_science_pipelines_operator","ports":"9000"},{"id":"data_science_pipelines_operator_svc_4","name":"minio-template-value","parent":"data_science_pipelines_operator","ports":"9000,80"},{"id":"data_science_pipelines_operator_svc_5","name":"ml-pipeline","parent":"data_science_pipelines_operator","ports":"8443,8888,8887"},{"id":"data_science_pipelines_operator_svc_6","name":"template-value","parent":"data_science_pipelines_operator","ports":"8443,8888,8887"},{"id":"feast_svc_0","name":"uvicorn-server","parent":"feast","ports":"6566"},{"id":"kserve_svc_0","name":"kserve-controller-manager-metrics-service","parent":"kserve","ports":"8443"},{"id":"kserve_svc_1","name":"kserve-controller-manager-service","parent":"kserve","ports":"8443"},{"id":"kserve_svc_2","name":"kserve-webhook-server-service","parent":"kserve","ports":"443"},{"id":"kserve_svc_3","name":"llmisvc-controller-manager-service","parent":"kserve","ports":"8443"},{"id":"kserve_svc_4","name":"llmisvc-webhook-server-service","parent":"kserve","ports":"443"},{"id":"kserve_svc_5","name":"localmodel-webhook-server-service","parent":"kserve","ports":"443"},{"id":"kserve_autogluon_server_svc_0","name":"kserve-controller-manager-metrics-service","parent":"kserve_autogluon_server","ports":"8443"},{"id":"kserve_autogluon_server_svc_1","name":"kserve-controller-manager-service","parent":"kserve_autogluon_server","ports":"8443"},{"id":"kserve_autogluon_server_svc_2","name":"kserve-webhook-server-service","parent":"kserve_autogluon_server","ports":"443"},{"id":"kserve_autogluon_server_svc_3","name":"llmisvc-controller-manager-service","parent":"kserve_autogluon_server","ports":"8443"},{"id":"kserve_autogluon_server_svc_4","name":"llmisvc-webhook-server-service","parent":"kserve_autogluon_server","ports":"443"},{"id":"kserve_autogluon_server_svc_5","name":"localmodel-webhook-server-service","parent":"kserve_autogluon_server","ports":"443"},{"id":"kubeflow_svc_0","name":"service","parent":"kubeflow","ports":"443"},{"id":"kubeflow_svc_1","name":"webhook-service","parent":"kubeflow","ports":"443"},{"id":"kuberay_svc_0","name":"kuberay-operator","parent":"kuberay","ports":"8080"},{"id":"kuberay_svc_1","name":"webhook-service","parent":"kuberay","ports":"443"},{"id":"kueue_svc_0","name":"visibility-server","parent":"kueue","ports":"443"},{"id":"kueue_svc_1","name":"webhook-service","parent":"kueue","ports":"443"},{"id":"llama_stack_svc_0","name":"cli-port-default","parent":"llama_stack","ports":"8081"},{"id":"llm_d_svc_0","name":"mooncake-client","parent":"llm_d","ports":"50052"},{"id":"llm_d_svc_1","name":"mooncake-master-store","parent":"llm_d","ports":"50051,8080,9003"},{"id":"llm_d_svc_2","name":"render","parent":"llm_d","ports":"8000"},{"id":"llm_d_inference_scheduler_svc_0","name":"${EPP_NAME}","parent":"llm_d_inference_scheduler","ports":"9002,5557,9090"},{"id":"llm_d_inference_scheduler_svc_1","name":"e2e-epp","parent":"llm_d_inference_scheduler","ports":"9002,5557"},{"id":"llm_d_inference_scheduler_svc_2","name":"e2e-epp-metrics","parent":"llm_d_inference_scheduler","ports":"9090"},{"id":"llm_d_inference_scheduler_svc_3","name":"inference-gateway-istio-nodeport","parent":"llm_d_inference_scheduler","ports":"15021,80"},{"id":"llm_d_inference_scheduler_svc_4","name":"istiod-llm-d-gateway","parent":"llm_d_inference_scheduler","ports":"15010,15012,443,15014"},{"id":"llm_d_inference_scheduler_svc_5","name":"service","parent":"llm_d_inference_scheduler","ports":"8080"},{"id":"mlflow_svc_0","name":"env-port-default","parent":"mlflow","ports":"9137"},{"id":"mlflow_operator_svc_0","name":"mlflow-operator-controller-manager-metrics-service","parent":"mlflow_operator","ports":"8443"},{"id":"model_registry_svc_0","name":"model-catalog","parent":"model_registry","ports":"8080"},{"id":"model_registry_operator_svc_0","name":"catalog-webhook-service","parent":"model_registry_operator","ports":"443"},{"id":"model_registry_operator_svc_1","name":"model-registry-operator-controller-manager-metrics-service","parent":"model_registry_operator","ports":"8443"},{"id":"model_registry_operator_svc_2","name":"model-registry-operator-webhook-service","parent":"model_registry_operator","ports":"443"},{"id":"model_registry_operator_svc_3","name":"template-value-postgres","parent":"model_registry_operator","ports":"5432"},{"id":"modelmesh_svc_0","name":"model-mesh","parent":"modelmesh","ports":"8033"},{"id":"modelmesh_serving_svc_0","name":"etcd","parent":"modelmesh_serving","ports":"2379"},{"id":"modelmesh_serving_svc_1","name":"modelmesh-controller","parent":"modelmesh_serving","ports":"8080"},{"id":"modelmesh_serving_svc_2","name":"modelmesh-webhook-server-service","parent":"modelmesh_serving","ports":"9443"},{"id":"models_as_a_service_svc_0","name":"maas-api","parent":"models_as_a_service","ports":"8080"},{"id":"models_as_a_service_svc_1","name":"maas-controller-webhook-service","parent":"models_as_a_service","ports":"443"},{"id":"models_as_a_service_svc_2","name":"payload-processing","parent":"models_as_a_service","ports":"9004"},{"id":"notebooks_svc_0","name":"notebook","parent":"notebooks","ports":"8888"},{"id":"odh_dashboard_svc_0","name":"odh-dashboard","parent":"odh_dashboard","ports":"8443,8943"},{"id":"odh_dashboard_svc_1","name":"odh-dashboard-agent-ops-ui","parent":"odh_dashboard","ports":"8843"},{"id":"odh_dashboard_svc_2","name":"odh-dashboard-automl-ui","parent":"odh_dashboard","ports":"8643"},{"id":"odh_dashboard_svc_3","name":"odh-dashboard-autorag-ui","parent":"odh_dashboard","ports":"8743"},{"id":"odh_dashboard_svc_4","name":"odh-dashboard-data-registry-ui","parent":"odh_dashboard","ports":"9043"},{"id":"odh_dashboard_svc_5","name":"odh-dashboard-eval-hub-ui","parent":"odh_dashboard","ports":"8543"},{"id":"odh_dashboard_svc_6","name":"odh-dashboard-gen-ai-ui","parent":"odh_dashboard","ports":"8143"},{"id":"odh_dashboard_svc_7","name":"odh-dashboard-maas-ui","parent":"odh_dashboard","ports":"8243"},{"id":"odh_dashboard_svc_8","name":"odh-dashboard-mlflow-ui","parent":"odh_dashboard","ports":"8343"},{"id":"odh_dashboard_svc_9","name":"odh-dashboard-model-registry-ui","parent":"odh_dashboard","ports":"8043"},{"id":"odh_dashboard_svc_10","name":"odh-dashboard-notebooks-ui","parent":"odh_dashboard","ports":"9043"},{"id":"odh_dashboard_svc_11","name":"rhaii-dashboard","parent":"odh_dashboard","ports":"4000"},{"id":"odh_dashboard_svc_12","name":"workspaces-backend","parent":"odh_dashboard","ports":"4000"},{"id":"odh_dashboard_svc_13","name":"workspaces-controller-metrics-service","parent":"odh_dashboard","ports":"8080"},{"id":"odh_dashboard_svc_14","name":"workspaces-frontend","parent":"odh_dashboard","ports":"8080"},{"id":"odh_dashboard_svc_15","name":"workspaces-webhook-service","parent":"odh_dashboard","ports":"443"},{"id":"odh_model_controller_svc_0","name":"model-serving-api","parent":"odh_model_controller","ports":"443,8080"},{"id":"odh_model_controller_svc_1","name":"odh-model-controller-metrics-service","parent":"odh_model_controller","ports":"8443"},{"id":"odh_model_controller_svc_2","name":"odh-model-controller-webhook-service","parent":"odh_model_controller","ports":"443"},{"id":"ogx_k8s_operator_svc_0","name":"ogx-k8s-operator-controller-manager-metrics-service","parent":"ogx_k8s_operator","ports":"8443"},{"id":"ogx_k8s_operator_svc_1","name":"ogx-k8s-operator-webhook-service","parent":"ogx_k8s_operator","ports":"443"},{"id":"opendatahub_operator_svc_0","name":"webhook-service","parent":"opendatahub_operator","ports":"443"},{"id":"spark_operator_svc_0","name":"spark-operator-webhook-svc","parent":"spark_operator","ports":"443"},{"id":"training_operator_svc_0","name":"training-operator","parent":"training_operator","ports":"8080,443"},{"id":"trustyai_service_operator_svc_0","name":"trustyai-service-operator-controller-manager-metrics-service","parent":"trustyai_service_operator","ports":"8443"},{"id":"trustyai_service_operator_svc_1","name":"trustyai-service-operator-metrics-service","parent":"trustyai_service_operator","ports":"8443"},{"id":"trustyai_service_operator_svc_2","name":"trustyai-service-operator-webhook-service","parent":"trustyai_service_operator","ports":"443"},{"id":"workload_variant_autoscaler_svc_0","name":"controller-manager-metrics-service","parent":"workload_variant_autoscaler","ports":"8443"}],"externals":[{"id":"ext_urllib","name":"urllib","type":"api"},{"id":"ext_chromadb","name":"chromadb","type":"api"},{"id":"ext_httpx","name":"httpx","type":"api"},{"id":"ext_milvus","name":"milvus","type":"api"},{"id":"ext_openai","name":"openai","type":"api"},{"id":"ext_grpc","name":"grpc","type":"grpc"},{"id":"ext_azure_blob","name":"azure-blob","type":"object-storage"},{"id":"ext_gcs","name":"gcs","type":"object-storage"},{"id":"ext_redis","name":"redis","type":"database"},{"id":"ext_s3","name":"s3","type":"object-storage"},{"id":"ext_requests","name":"requests","type":"api"},{"id":"ext_minio","name":"minio","type":"object-storage"},{"id":"ext_mysql","name":"mysql","type":"database"},{"id":"ext_elasticsearch","name":"elasticsearch","type":"api"},{"id":"ext_mlflow","name":"mlflow","type":"api"},{"id":"ext_qdrant","name":"qdrant","type":"api"},{"id":"ext_weaviate","name":"weaviate","type":"api"},{"id":"ext_mongodb","name":"mongodb","type":"database"},{"id":"ext_postgres","name":"postgres","type":"database"},{"id":"ext_sqlalchemy","name":"sqlalchemy","type":"database"},{"id":"ext_sqlite","name":"sqlite","type":"database"},{"id":"ext_aiohttp","name":"aiohttp","type":"api"},{"id":"ext_anthropic","name":"anthropic","type":"api"},{"id":"ext_cohere","name":"cohere","type":"api"},{"id":"ext_ogx","name":"ogx","type":"api"},{"id":"ext_wandb","name":"wandb","type":"api"}],"edges":[{"from":"codeflare_operator","to":"opendatahub_operator","type":"module"},{"from":"data_science_pipelines_operator","to":"mlflow_operator","type":"module"},{"from":"data_science_pipelines_operator","to":"operator_chaos","type":"module"},{"from":"kserve","to":"odh_platform_utilities","type":"module"},{"from":"kserve","to":"kserve_autogluon_server","type":"watches"},{"from":"kubeflow","to":"data_science_pipelines_operator","type":"module"},{"from":"kubeflow","to":"operator_chaos","type":"module"},{"from":"llm_d_inference_scheduler","to":"kserve","type":"watches"},{"from":"model_registry_operator","to":"odh_platform_utilities","type":"module"},{"from":"model_registry_operator","to":"operator_chaos","type":"module"},{"from":"model_registry","to":"kserve_autogluon_server","type":"watches"},{"from":"modelmesh_serving","to":"kserve_autogluon_server","type":"watches"},{"from":"odh_dashboard","to":"mlflow_go","type":"module"},{"from":"odh_dashboard","to":"odh_platform_utilities","type":"module"},{"from":"odh_dashboard","to":"ogx_k8s_operator","type":"module"},{"from":"odh_model_controller","to":"kserve","type":"module"},{"from":"odh_model_controller","to":"kserve_autogluon_server","type":"watches"},{"from":"ogx_k8s_operator","to":"odh_platform_utilities","type":"module"},{"from":"opendatahub_operator","to":"models_as_a_service","type":"module"},{"from":"opendatahub_operator","to":"odh_platform_utilities","type":"module"},{"from":"spark_operator","to":"odh_platform_utilities","type":"module"},{"from":"trustyai_service_operator","to":"odh_platform_utilities","type":"module"},{"from":"workload_variant_autoscaler","to":"kserve","type":"watches"},{"from":"kserve_autogluon_server","to":"kserve","type":"module"},{"from":"model_registry_operator","to":"opendatahub_operator","type":"module"},{"from":"modelmesh_serving","to":"opendatahub_operator","type":"module"},{"from":"agents_operator","to":"ext_urllib","type":"external"},{"from":"ai4rag","to":"ext_chromadb","type":"external"},{"from":"ai4rag","to":"ext_httpx","type":"external"},{"from":"ai4rag","to":"ext_milvus","type":"external"},{"from":"ai4rag","to":"ext_openai","type":"external"},{"from":"argo_workflows","to":"ext_grpc","type":"external"},{"from":"argo_workflows","to":"ext_azure_blob","type":"external"},{"from":"argo_workflows","to":"ext_gcs","type":"external"},{"from":"batch_gateway","to":"ext_urllib","type":"external"},{"from":"batch_gateway","to":"ext_redis","type":"external"},{"from":"batch_gateway","to":"ext_s3","type":"external"},{"from":"data_science_pipelines","to":"ext_requests","type":"external"},{"from":"data_science_pipelines","to":"ext_grpc","type":"external"},{"from":"data_science_pipelines","to":"ext_gcs","type":"external"},{"from":"data_science_pipelines","to":"ext_minio","type":"external"},{"from":"data_science_pipelines","to":"ext_s3","type":"external"},{"from":"data_science_pipelines_operator","to":"ext_requests","type":"external"},{"from":"data_science_pipelines_operator","to":"ext_mysql","type":"external"},{"from":"data_science_pipelines_operator","to":"ext_minio","type":"external"},{"from":"eval_hub","to":"ext_requests","type":"external"},{"from":"eval_hub","to":"ext_s3","type":"external"},{"from":"feast","to":"ext_elasticsearch","type":"external"},{"from":"feast","to":"ext_milvus","type":"external"},{"from":"feast","to":"ext_mlflow","type":"external"},{"from":"feast","to":"ext_qdrant","type":"external"},{"from":"feast","to":"ext_requests","type":"external"},{"from":"feast","to":"ext_urllib","type":"external"},{"from":"feast","to":"ext_weaviate","type":"external"},{"from":"feast","to":"ext_mongodb","type":"external"},{"from":"feast","to":"ext_mysql","type":"external"},{"from":"feast","to":"ext_postgres","type":"external"},{"from":"feast","to":"ext_redis","type":"external"},{"from":"feast","to":"ext_sqlalchemy","type":"external"},{"from":"feast","to":"ext_sqlite","type":"external"},{"from":"feast","to":"ext_grpc","type":"external"},{"from":"feast","to":"ext_azure_blob","type":"external"},{"from":"feast","to":"ext_gcs","type":"external"},{"from":"feast","to":"ext_s3","type":"external"},{"from":"kserve","to":"ext_httpx","type":"external"},{"from":"kserve","to":"ext_requests","type":"external"},{"from":"kserve","to":"ext_weaviate","type":"external"},{"from":"kserve","to":"ext_azure_blob","type":"external"},{"from":"kserve","to":"ext_gcs","type":"external"},{"from":"kserve","to":"ext_s3","type":"external"},{"from":"kserve_autogluon_server","to":"ext_httpx","type":"external"},{"from":"kserve_autogluon_server","to":"ext_requests","type":"external"},{"from":"kserve_autogluon_server","to":"ext_weaviate","type":"external"},{"from":"kserve_autogluon_server","to":"ext_azure_blob","type":"external"},{"from":"kserve_autogluon_server","to":"ext_gcs","type":"external"},{"from":"kserve_autogluon_server","to":"ext_s3","type":"external"},{"from":"kube_auth_proxy","to":"ext_redis","type":"external"},{"from":"kubeflow","to":"ext_requests","type":"external"},{"from":"kuberay","to":"ext_requests","type":"external"},{"from":"kuberay","to":"ext_urllib","type":"external"},{"from":"kuberay","to":"ext_grpc","type":"external"},{"from":"kuberay","to":"ext_azure_blob","type":"external"},{"from":"kuberay","to":"ext_gcs","type":"external"},{"from":"llama_stack","to":"ext_aiohttp","type":"external"},{"from":"llama_stack","to":"ext_anthropic","type":"external"},{"from":"llama_stack","to":"ext_chromadb","type":"external"},{"from":"llama_stack","to":"ext_cohere","type":"external"},{"from":"llama_stack","to":"ext_elasticsearch","type":"external"},{"from":"llama_stack","to":"ext_httpx","type":"external"},{"from":"llama_stack","to":"ext_milvus","type":"external"},{"from":"llama_stack","to":"ext_ogx","type":"external"},{"from":"llama_stack","to":"ext_openai","type":"external"},{"from":"llama_stack","to":"ext_qdrant","type":"external"},{"from":"llama_stack","to":"ext_requests","type":"external"},{"from":"llama_stack","to":"ext_urllib","type":"external"},{"from":"llama_stack","to":"ext_mongodb","type":"external"},{"from":"llama_stack","to":"ext_postgres","type":"external"},{"from":"llama_stack","to":"ext_redis","type":"external"},{"from":"llama_stack","to":"ext_sqlalchemy","type":"external"},{"from":"llama_stack","to":"ext_sqlite","type":"external"},{"from":"llama_stack","to":"ext_s3","type":"external"},{"from":"llama_stack_provider_trustyai_garak","to":"ext_httpx","type":"external"},{"from":"llama_stack_provider_trustyai_garak","to":"ext_urllib","type":"external"},{"from":"llama_stack_provider_trustyai_garak","to":"ext_s3","type":"external"},{"from":"llm_d_inference_scheduler","to":"ext_redis","type":"external"},{"from":"llm_d_kv_cache","to":"ext_redis","type":"external"},{"from":"llm_d_kv_cache","to":"ext_sqlalchemy","type":"external"},{"from":"lm_evaluation_harness","to":"ext_anthropic","type":"external"},{"from":"lm_evaluation_harness","to":"ext_httpx","type":"external"},{"from":"lm_evaluation_harness","to":"ext_openai","type":"external"},{"from":"lm_evaluation_harness","to":"ext_requests","type":"external"},{"from":"lm_evaluation_harness","to":"ext_wandb","type":"external"},{"from":"mlflow","to":"ext_aiohttp","type":"external"},{"from":"mlflow","to":"ext_httpx","type":"external"},{"from":"mlflow","to":"ext_mlflow","type":"external"},{"from":"mlflow","to":"ext_openai","type":"external"},{"from":"mlflow","to":"ext_requests","type":"external"},{"from":"mlflow","to":"ext_urllib","type":"external"},{"from":"mlflow","to":"ext_postgres","type":"external"},{"from":"mlflow","to":"ext_redis","type":"external"},{"from":"mlflow","to":"ext_sqlalchemy","type":"external"},{"from":"mlflow","to":"ext_azure_blob","type":"external"},{"from":"mlflow","to":"ext_gcs","type":"external"},{"from":"mlflow","to":"ext_s3","type":"external"},{"from":"mlflow_operator","to":"ext_mlflow","type":"external"},{"from":"model_registry","to":"ext_aiohttp","type":"external"},{"from":"model_registry","to":"ext_requests","type":"external"},{"from":"model_registry","to":"ext_urllib","type":"external"},{"from":"model_registry","to":"ext_s3","type":"external"},{"from":"modelmesh_serving","to":"ext_azure_blob","type":"external"},{"from":"notebooks","to":"ext_requests","type":"external"},{"from":"notebooks","to":"ext_urllib","type":"external"},{"from":"notebooks","to":"ext_minio","type":"external"},{"from":"odh_dashboard","to":"ext_s3","type":"external"},{"from":"odh_model_controller","to":"ext_s3","type":"external"},{"from":"opendatahub_operator","to":"ext_s3","type":"external"}]}
  </script>
</div>
<div class="topology-legend">
  <span><span class="swatch" style="background:#3498db"></span> Component</span>
  <span><span class="swatch" style="background:#27ae60"></span> Has Ingress</span>
  <span><span class="swatch" style="background:#3498db;border:2px solid #f39c12"></span> Has NetworkPolicy</span>
  <span><span class="swatch" style="background:#e74c3c;border-radius:2px;transform:rotate(45deg)"></span> External</span>
  <span><span class="line-swatch" style="background:#e74c3c"></span> CRD Watch</span>
  <span><span class="line-swatch" style="background:#9b59b6"></span> Sidecar</span>
  <span><span class="line-swatch" style="background:#95a5a6;border-top:2px dashed #95a5a6;height:0"></span> Module</span>
  <span><span class="line-swatch" style="background:#e67e22;border-top:2px dotted #e67e22;height:0"></span> External</span>
</div>

## Cross-Component Service References

Services referenced across component boundaries. When component A defines a service that component B also references, it indicates a deployment dependency.

```mermaid
graph LR
    classDef comp fill:#3498db,stroke:#2980b9,color:#fff

    kserve["kserve"]:::comp
    kserve_autogluon_server["kserve-autogluon-server"]:::comp

    kserve_autogluon_server -.->|"kserve-controller-manager-service"| kserve
    kserve_autogluon_server -.->|"llmisvc-controller-manager-service"| kserve
```

## Services by Component

| Component | Services | Webhook (443) | Metrics (8443) | Data |
|-----------|----------|---------------|----------------|------|
| agents-operator | 2 | 1 | 0 | 1 |
| codeflare-operator | 1 | 1 | 0 | 0 |
| data-science-pipelines | 1 | 0 | 0 | 1 |
| data-science-pipelines-operator | 7 | 0 | 2 | 5 |
| feast | 1 | 0 | 0 | 1 |
| kserve | 6 | 3 | 3 | 0 |
| kserve-autogluon-server | 6 | 3 | 3 | 0 |
| kubeflow | 2 | 2 | 0 | 0 |
| kuberay | 2 | 1 | 0 | 1 |
| kueue | 2 | 2 | 0 | 0 |
| llama-stack | 1 | 0 | 0 | 1 |
| llm-d | 3 | 0 | 0 | 3 |
| llm-d-inference-scheduler | 6 | 1 | 0 | 5 |
| mlflow | 1 | 0 | 0 | 1 |
| mlflow-operator | 1 | 0 | 1 | 0 |
| model-registry | 1 | 0 | 0 | 1 |
| model-registry-operator | 4 | 2 | 1 | 1 |
| modelmesh | 1 | 0 | 0 | 1 |
| modelmesh-serving | 3 | 0 | 0 | 3 |
| models-as-a-service | 3 | 1 | 0 | 2 |
| notebooks | 1 | 0 | 0 | 1 |
| odh-dashboard | 16 | 1 | 1 | 14 |
| odh-model-controller | 3 | 2 | 1 | 0 |
| ogx-k8s-operator | 2 | 1 | 1 | 0 |
| opendatahub-operator | 1 | 1 | 0 | 0 |
| spark-operator | 1 | 1 | 0 | 0 |
| training-operator | 1 | 1 | 0 | 0 |
| trustyai-service-operator | 3 | 1 | 2 | 0 |
| workload-variant-autoscaler | 1 | 0 | 1 | 0 |

## Service Detail

Per-component service breakdown with exact port numbers and protocols.

### agents-operator (2 services)

| Service | Type | Ports |
|---------|------|-------|
| bundle-service | ClusterIP | 8080/TCP |
| webhook-service | ClusterIP | 443/TCP |

### codeflare-operator (1 services)

| Service | Type | Ports |
|---------|------|-------|
| webhook-service | ClusterIP | 443/TCP |

### data-science-pipelines (1 services)

| Service | Type | Ports |
|---------|------|-------|
| kubeflow-pipelines-profile-controller | ClusterIP | 80/TCP |

### data-science-pipelines-operator (7 services)

| Service | Type | Ports |
|---------|------|-------|
| data-science-pipelines-operator-service | ClusterIP | 8080/TCP |
| ds-pipeline-workflow-controller-metrics-template-value | ClusterIP | 9090/TCP |
| mariadb-template-value | ClusterIP | 3306/TCP |
| minio-service | ClusterIP | 9000/TCP |
| minio-template-value | ClusterIP | 9000/TCP, 80/TCP |
| ml-pipeline | ClusterIP | 8443/TCP, 8888/TCP, 8887/TCP |
| template-value | ClusterIP | 8443/TCP, 8888/TCP, 8887/TCP |

### feast (1 services)

| Service | Type | Ports |
|---------|------|-------|
| uvicorn-server | python-source | 6566/TCP |

### kserve (6 services)

| Service | Type | Ports |
|---------|------|-------|
| kserve-controller-manager-metrics-service | ClusterIP | 8443/TCP |
| kserve-controller-manager-service | ClusterIP | 8443/TCP |
| kserve-webhook-server-service | ClusterIP | 443/TCP |
| llmisvc-controller-manager-service | ClusterIP | 8443/TCP |
| llmisvc-webhook-server-service | ClusterIP | 443/TCP |
| localmodel-webhook-server-service | ClusterIP | 443/TCP |

### kserve-autogluon-server (6 services)

| Service | Type | Ports |
|---------|------|-------|
| kserve-controller-manager-metrics-service | ClusterIP | 8443/TCP |
| kserve-controller-manager-service | ClusterIP | 8443/TCP |
| kserve-webhook-server-service | ClusterIP | 443/TCP |
| llmisvc-controller-manager-service | ClusterIP | 8443/TCP |
| llmisvc-webhook-server-service | ClusterIP | 443/TCP |
| localmodel-webhook-server-service | ClusterIP | 443/TCP |

### kubeflow (2 services)

| Service | Type | Ports |
|---------|------|-------|
| service | ClusterIP | 443/TCP |
| webhook-service | ClusterIP | 443/TCP |

### kuberay (2 services)

| Service | Type | Ports |
|---------|------|-------|
| kuberay-operator | ClusterIP | 8080/TCP |
| webhook-service | ClusterIP | 443/TCP |

### kueue (2 services)

| Service | Type | Ports |
|---------|------|-------|
| visibility-server | ClusterIP | 443/TCP |
| webhook-service | ClusterIP | 443/TCP |

### llama-stack (1 services)

| Service | Type | Ports |
|---------|------|-------|
| cli-port-default | python-source | 8081/TCP |

### llm-d (3 services)

| Service | Type | Ports |
|---------|------|-------|
| mooncake-client | ClusterIP | 50052/TCP |
| mooncake-master-store | ClusterIP | 50051/TCP, 8080/TCP, 9003/TCP |
| render | ClusterIP | 8000/TCP |

### llm-d-inference-scheduler (6 services)

| Service | Type | Ports |
|---------|------|-------|
| ${EPP_NAME} | ClusterIP | 9002/TCP, 5557/TCP, 9090/TCP |
| e2e-epp | ClusterIP | 9002/TCP, 5557/TCP |
| e2e-epp-metrics | NodePort | 9090/TCP |
| inference-gateway-istio-nodeport | NodePort | 15021/TCP, 80/TCP |
| istiod-llm-d-gateway | ClusterIP | 15010/TCP, 15012/TCP, 443/TCP, 15014/TCP |
| service | ClusterIP | 8080/TCP |

### mlflow (1 services)

| Service | Type | Ports |
|---------|------|-------|
| env-port-default | python-source | 9137/TCP |

### mlflow-operator (1 services)

| Service | Type | Ports |
|---------|------|-------|
| mlflow-operator-controller-manager-metrics-service | ClusterIP | 8443/TCP |

### model-registry (1 services)

| Service | Type | Ports |
|---------|------|-------|
| model-catalog | ClusterIP | 8080/TCP |

### model-registry-operator (4 services)

| Service | Type | Ports |
|---------|------|-------|
| catalog-webhook-service | ClusterIP | 443/TCP |
| model-registry-operator-controller-manager-metrics-service | ClusterIP | 8443/TCP |
| model-registry-operator-webhook-service | ClusterIP | 443/TCP |
| template-value-postgres | ClusterIP | 5432/TCP |

### modelmesh (1 services)

| Service | Type | Ports |
|---------|------|-------|
| model-mesh | ClusterIP | 8033/TCP |

### modelmesh-serving (3 services)

| Service | Type | Ports |
|---------|------|-------|
| etcd | ClusterIP | 2379/TCP |
| modelmesh-controller | ClusterIP | 8080/TCP |
| modelmesh-webhook-server-service | ClusterIP | 9443/TCP |

### models-as-a-service (3 services)

| Service | Type | Ports |
|---------|------|-------|
| maas-api | ClusterIP | 8080/TCP |
| maas-controller-webhook-service | ClusterIP | 443/TCP |
| payload-processing | ClusterIP | 9004/TCP |

### notebooks (1 services)

| Service | Type | Ports |
|---------|------|-------|
| notebook | ClusterIP | 8888/TCP |

### odh-dashboard (16 services)

| Service | Type | Ports |
|---------|------|-------|
| odh-dashboard | ClusterIP | 8443/TCP, 8943/TCP |
| odh-dashboard-agent-ops-ui | ClusterIP | 8843/TCP |
| odh-dashboard-automl-ui | ClusterIP | 8643/TCP |
| odh-dashboard-autorag-ui | ClusterIP | 8743/TCP |
| odh-dashboard-data-registry-ui | ClusterIP | 9043/TCP |
| odh-dashboard-eval-hub-ui | ClusterIP | 8543/TCP |
| odh-dashboard-gen-ai-ui | ClusterIP | 8143/TCP |
| odh-dashboard-maas-ui | ClusterIP | 8243/TCP |
| odh-dashboard-mlflow-ui | ClusterIP | 8343/TCP |
| odh-dashboard-model-registry-ui | ClusterIP | 8043/TCP |
| odh-dashboard-notebooks-ui | ClusterIP | 9043/TCP |
| rhaii-dashboard | ClusterIP | 4000/TCP |
| workspaces-backend | ClusterIP | 4000/TCP |
| workspaces-controller-metrics-service | ClusterIP | 8080/TCP |
| workspaces-frontend | ClusterIP | 8080/TCP |
| workspaces-webhook-service | ClusterIP | 443/TCP |

### odh-model-controller (3 services)

| Service | Type | Ports |
|---------|------|-------|
| model-serving-api | ClusterIP | 443/TCP, 8080/TCP |
| odh-model-controller-metrics-service | ClusterIP | 8443/TCP |
| odh-model-controller-webhook-service | ClusterIP | 443/TCP |

### ogx-k8s-operator (2 services)

| Service | Type | Ports |
|---------|------|-------|
| ogx-k8s-operator-controller-manager-metrics-service | ClusterIP | 8443/TCP |
| ogx-k8s-operator-webhook-service | ClusterIP | 443/TCP |

### opendatahub-operator (1 services)

| Service | Type | Ports |
|---------|------|-------|
| webhook-service | ClusterIP | 443/TCP |

### spark-operator (1 services)

| Service | Type | Ports |
|---------|------|-------|
| spark-operator-webhook-svc | ClusterIP | 443/TCP |

### training-operator (1 services)

| Service | Type | Ports |
|---------|------|-------|
| training-operator | ClusterIP | 8080/TCP, 443/TCP |

### trustyai-service-operator (3 services)

| Service | Type | Ports |
|---------|------|-------|
| trustyai-service-operator-controller-manager-metrics-service | ClusterIP | 8443/TCP |
| trustyai-service-operator-metrics-service | ClusterIP | 8443/TCP |
| trustyai-service-operator-webhook-service | ClusterIP | 443/TCP |

### workload-variant-autoscaler (1 services)

| Service | Type | Ports |
|---------|------|-------|
| controller-manager-metrics-service | ClusterIP | 8443/TCP |

## Port Patterns

- **15010/TCP**: istiod-llm-d-gateway
- **15012/TCP**: istiod-llm-d-gateway
- **15014/TCP**: istiod-llm-d-gateway
- **15021/TCP**: inference-gateway-istio-nodeport
- **2379/TCP**: etcd
- **3306/TCP**: mariadb-template-value
- **4000/TCP**: rhaii-dashboard, workspaces-backend
- **443/TCP**: webhook-service, webhook-service, kserve-webhook-server-service, llmisvc-webhook-server-service, localmodel-webhook-server-service, kserve-webhook-server-service, llmisvc-webhook-server-service, localmodel-webhook-server-service, service, webhook-service, webhook-service, visibility-server, webhook-service, istiod-llm-d-gateway, catalog-webhook-service, model-registry-operator-webhook-service, maas-controller-webhook-service, workspaces-webhook-service, model-serving-api, odh-model-controller-webhook-service, ogx-k8s-operator-webhook-service, webhook-service, spark-operator-webhook-svc, training-operator, trustyai-service-operator-webhook-service
- **50051/TCP**: mooncake-master-store
- **50052/TCP**: mooncake-client
- **5432/TCP**: template-value-postgres
- **5557/TCP**: ${EPP_NAME}, e2e-epp
- **6566/TCP**: uvicorn-server
- **80/TCP**: minio-template-value, kubeflow-pipelines-profile-controller, inference-gateway-istio-nodeport
- **8000/TCP**: render
- **8033/TCP**: model-mesh
- **8043/TCP**: odh-dashboard-model-registry-ui
- **8080/TCP**: mooncake-master-store, bundle-service, data-science-pipelines-operator-service, kuberay-operator, service, model-catalog, modelmesh-controller, maas-api, workspaces-controller-metrics-service, workspaces-frontend, model-serving-api, training-operator
- **8081/TCP**: cli-port-default
- **8143/TCP**: odh-dashboard-gen-ai-ui
- **8243/TCP**: odh-dashboard-maas-ui
- **8343/TCP**: odh-dashboard-mlflow-ui
- **8443/TCP**: ml-pipeline, template-value, kserve-controller-manager-metrics-service, kserve-controller-manager-service, llmisvc-controller-manager-service, kserve-controller-manager-metrics-service, kserve-controller-manager-service, llmisvc-controller-manager-service, mlflow-operator-controller-manager-metrics-service, model-registry-operator-controller-manager-metrics-service, odh-dashboard, odh-model-controller-metrics-service, ogx-k8s-operator-controller-manager-metrics-service, trustyai-service-operator-controller-manager-metrics-service, trustyai-service-operator-metrics-service, controller-manager-metrics-service
- **8543/TCP**: odh-dashboard-eval-hub-ui
- **8643/TCP**: odh-dashboard-automl-ui
- **8743/TCP**: odh-dashboard-autorag-ui
- **8843/TCP**: odh-dashboard-agent-ops-ui
- **8887/TCP**: ml-pipeline, template-value
- **8888/TCP**: ml-pipeline, template-value, notebook
- **8943/TCP**: odh-dashboard
- **9000/TCP**: minio-service, minio-template-value
- **9002/TCP**: ${EPP_NAME}, e2e-epp
- **9003/TCP**: mooncake-master-store
- **9004/TCP**: payload-processing
- **9043/TCP**: odh-dashboard-data-registry-ui, odh-dashboard-notebooks-ui
- **9090/TCP**: ds-pipeline-workflow-controller-metrics-template-value, ${EPP_NAME}, e2e-epp-metrics
- **9137/TCP**: env-port-default
- **9443/TCP**: modelmesh-webhook-server-service

