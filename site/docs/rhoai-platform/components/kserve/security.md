# kserve: Security

## Secrets

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| llmisvc-webhook-server-cert | Opaque | deployment/llmisvc-controller-manager |
| localmodel-webhook-server-cert | Opaque | deployment/kserve-localmodel-controller-manager |
| kserve-webhook-server-cert | Opaque | deployment/kserve-controller-manager |

## Deployment Security Controls

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| spark-pmml-iris | kfserving-container | ? | ? | ? | `docs/samples/v1beta1/spark/deployment.yaml` |
| kserve-controller-manager | kube-rbac-proxy | true | true | false | `config/default/manager_auth_proxy_patch.yaml` |
| kserve-controller-manager | manager | ? | ? | ? | `config/default/manager_auth_proxy_patch.yaml` |
| kserve-controller-manager | manager | ? | ? | ? | `config/default/manager_image_patch.yaml` |
| kserve-controller-manager | manager | ? | ? | ? | `config/default/manager_prometheus_metrics_patch.yaml` |
| kserve-controller-manager | manager | ? | ? | ? | `config/default/manager_resources_patch.yaml` |
| llmisvc-controller-manager | manager | true | true | false | `config/llmisvc/manager.yaml` |
| kserve-localmodel-controller-manager | manager | true | true | false | `config/localmodels/manager.yaml` |
| kserve-controller-manager | manager | true | true | false | `config/manager/manager.yaml` |
| kserve-controller-manager | manager | ? | ? | ? | `config/overlays/test/manager_image_patch.yaml` |
| kserve-controller-manager | manager | ? | ? | ? | `config/overlays/version-template/manager_image_patch.yaml` |

## Build Security

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 4 | 1000:1000 |  |  |  | Unpinned base image: deps; Unpinned base image: deps; Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `docs/apis/Dockerfile` | registry.access.redhat.com/ubi9/go-toolset:1.25 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `docs/kfp/Dockerfile` | python:3.9-slim-bullseye | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `docs/samples/explanation/aif/germancredit/server/Dockerfile` | python:3.10-slim | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `docs/samples/graph/bgtest/Dockerfile` | scratch | 2 |  |  |  |  | Unpinned base image: scratch; No USER directive found (defaults to root) |
| `docs/samples/v1beta1/custom/paddleserving/Dockerfile` | registry.baidubce.com/paddlepaddle/serving:0.5.0-devel | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `docs/samples/v1beta1/custom/torchserve/torchserve-image/Dockerfile` | ${BASE_IMAGE} | 2 | model-server |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `docs/samples/v1beta1/torchserve/model-archiver/model-archiver-image/Dockerfile` | ubuntu:18.04 | 1 | model-server |  |  |  |  |
| `docs/samples/v1beta1/triton/fastertransformer/transformer/Dockerfile` | python:3.9-slim-bullseye | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `hack/kserve_migration/Dockerfile` | alpine:latest | 1 |  |  |  |  | Unpinned base image: alpine:latest; No USER directive found (defaults to root) |
| `tools/tf2openapi/Dockerfile` | gcr.io/distroless/static:nonroot | 2 |  |  |  |  | No USER directive found (defaults to root) |

## Cache Architecture

### Manager Configuration

| Property | Value |
|----------|-------|
| Manager file | `cmd/manager/main.go` |
| Cache scope | cluster-wide |
| DefaultTransform | no |
| Memory limit | 300Mi |

### Filtered Types

| Type | Filter Kind | Filter |
|------|-------------|--------|
| corev1.Secret | label | label selector |
| corev1.ConfigMap | label | label selector |
| appsv1.Deployment | label | label selector |
| corev1.Pod | label | label selector |
| autoscalingv2.HorizontalPodAutoscaler | label | label selector |

### Issues

- No DefaultTransform: managedFields cached for all objects (wasted memory)
- No GOMEMLIMIT set in deployment (Go GC cannot pressure-tune)
- Type ClusterServingRuntime is watched but has no cache filter (cluster-wide informer)
- Type Gateway is watched but has no cache filter (cluster-wide informer)
- Type HTTPRoute is watched but has no cache filter (cluster-wide informer)
- Type InferenceGraph is watched but has no cache filter (cluster-wide informer)
- Type InferencePool is watched but has no cache filter (cluster-wide informer)
- Type InferenceService is watched but has no cache filter (cluster-wide informer)
- Type Ingress is watched but has no cache filter (cluster-wide informer)
- Type Job is watched but has no cache filter (cluster-wide informer)
- Type LLMInferenceService is watched but has no cache filter (cluster-wide informer)
- Type LLMInferenceServiceConfig is watched but has no cache filter (cluster-wide informer)
- Type LeaderWorkerSet is watched but has no cache filter (cluster-wide informer)
- Type LocalModelCache is watched but has no cache filter (cluster-wide informer)
- Type LocalModelNamespaceCache is watched but has no cache filter (cluster-wide informer)
- Type LocalModelNode is watched but has no cache filter (cluster-wide informer)
- Type Node is watched but has no cache filter (cluster-wide informer)
- Type OpenTelemetryCollector is watched but has no cache filter (cluster-wide informer)
- Type PersistentVolume is watched but has no cache filter (cluster-wide informer)
- Type PersistentVolumeClaim is watched but has no cache filter (cluster-wide informer)
- Type PodMonitor is watched but has no cache filter (cluster-wide informer)
- Type Route is watched but has no cache filter (cluster-wide informer)
- Type ScaledObject is watched but has no cache filter (cluster-wide informer)
- Type Service is watched but has no cache filter (cluster-wide informer)
- Type ServiceMonitor is watched but has no cache filter (cluster-wide informer)
- Type ServingRuntime is watched but has no cache filter (cluster-wide informer)
- Type TrainedModel is watched but has no cache filter (cluster-wide informer)
- Type VariantAutoscaling is watched but has no cache filter (cluster-wide informer)
- Type VirtualService is watched but has no cache filter (cluster-wide informer)

