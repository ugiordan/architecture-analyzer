# kuberay: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

| Type | GVK | Source |
|------|-----|--------|
| For | ray/v1/RayCluster | [`ray-operator/controllers/ray/authentication_controller.go:1050`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/controllers/ray/authentication_controller.go#L1050) |
| For | ray/v1/RayCluster | [`ray-operator/controllers/ray/networkpolicy_controller.go:427`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/controllers/ray/networkpolicy_controller.go#L427) |
| For | ray/v1/RayCluster | [`ray-operator/controllers/ray/raycluster_controller.go:2009`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/controllers/ray/raycluster_controller.go#L2009) |
| For | ray/v1/RayCluster | [`ray-operator/controllers/ray/raycluster_mtls_controller.go:834`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/controllers/ray/raycluster_mtls_controller.go#L834) |
| For | ray/v1/RayCronJob | [`ray-operator/controllers/ray/raycronjob_controller.go:182`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/controllers/ray/raycronjob_controller.go#L182) |
| For | ray/v1/RayJob | [`ray-operator/controllers/ray/rayjob_controller.go:888`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/controllers/ray/rayjob_controller.go#L888) |
| For | ray/v1/RayService | [`ray-operator/controllers/ray/rayservice_controller.go:603`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/controllers/ray/rayservice_controller.go#L603) |
| Owns | /v1/Pod | [`ray-operator/controllers/ray/raycluster_controller.go:2014`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/controllers/ray/raycluster_controller.go#L2014) |
| Owns | /v1/Secret | [`ray-operator/controllers/ray/raycluster_controller.go:2016`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/controllers/ray/raycluster_controller.go#L2016) |
| Owns | /v1/Service | [`ray-operator/controllers/ray/rayjob_controller.go:890`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/controllers/ray/rayjob_controller.go#L890) |
| Owns | /v1/Service | [`ray-operator/controllers/ray/raycluster_controller.go:2015`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/controllers/ray/raycluster_controller.go#L2015) |
| Owns | /v1/Service | [`ray-operator/controllers/ray/authentication_controller.go:1053`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/controllers/ray/authentication_controller.go#L1053) |
| Owns | /v1/Service | [`ray-operator/controllers/ray/rayservice_controller.go:609`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/controllers/ray/rayservice_controller.go#L609) |
| Owns | /v1/ServiceAccount | [`ray-operator/controllers/ray/authentication_controller.go:1052`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/controllers/ray/authentication_controller.go#L1052) |
| Owns | batch/v1/Job | [`ray-operator/controllers/ray/rayjob_controller.go:891`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/controllers/ray/rayjob_controller.go#L891) |
| Owns | networking.k8s.io/v1/NetworkPolicy | [`ray-operator/controllers/ray/networkpolicy_controller.go:428`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/controllers/ray/networkpolicy_controller.go#L428) |
| Owns | ray/v1/RayCluster | [`ray-operator/controllers/ray/rayjob_controller.go:889`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/controllers/ray/rayjob_controller.go#L889) |
| Owns | ray/v1/RayCluster | [`ray-operator/controllers/ray/rayservice_controller.go:608`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/controllers/ray/rayservice_controller.go#L608) |
| Owns | ray/v1/RayJob | [`ray-operator/controllers/ray/raycronjob_controller.go:183`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/controllers/ray/raycronjob_controller.go#L183) |
| Owns | route/v1/Route | [`ray-operator/controllers/ray/authentication_controller.go:1055`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/controllers/ray/authentication_controller.go#L1055) |

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for kuberay

    participant KubernetesAPI as Kubernetes API
    participant kuberay_operator as kuberay-operator

    KubernetesAPI->>+kuberay_operator: Watch RayCluster (reconcile)
    KubernetesAPI->>+kuberay_operator: Watch RayCluster (reconcile)
    KubernetesAPI->>+kuberay_operator: Watch RayCluster (reconcile)
    KubernetesAPI->>+kuberay_operator: Watch RayCluster (reconcile)
    KubernetesAPI->>+kuberay_operator: Watch RayCronJob (reconcile)
    KubernetesAPI->>+kuberay_operator: Watch RayJob (reconcile)
    KubernetesAPI->>+kuberay_operator: Watch RayService (reconcile)
    kuberay_operator->>KubernetesAPI: Create/Update Pod
    kuberay_operator->>KubernetesAPI: Create/Update Secret
    kuberay_operator->>KubernetesAPI: Create/Update Service
    kuberay_operator->>KubernetesAPI: Create/Update Service
    kuberay_operator->>KubernetesAPI: Create/Update Service
    kuberay_operator->>KubernetesAPI: Create/Update Service
    kuberay_operator->>KubernetesAPI: Create/Update ServiceAccount
    kuberay_operator->>KubernetesAPI: Create/Update Job
    kuberay_operator->>KubernetesAPI: Create/Update NetworkPolicy
    kuberay_operator->>KubernetesAPI: Create/Update RayCluster
    kuberay_operator->>KubernetesAPI: Create/Update RayCluster
    kuberay_operator->>KubernetesAPI: Create/Update RayJob
    kuberay_operator->>KubernetesAPI: Create/Update Route

    Note over kuberay_operator: Exposed Services
    Note right of kuberay_operator: kuberay-operator:8080/TCP [monitoring-port]
    Note right of kuberay_operator: webhook-service:443/TCP []
```

### Webhooks

| Name | Type | Path | Failure Policy | Service | Overlays | Enable Condition | Sources |
|------|------|------|----------------|---------|----------|------------------|----------|
| mraycluster.kb.io | mutating | /mutate-ray-io-v1-raycluster | Fail | $(namespace)/kuberay-webhook-service |  |  | [`ray-operator/config/openshift/webhook.yaml`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/config/openshift/webhook.yaml) |
| mraycluster.kb.io | mutating | /mutate-ray-io-v1-raycluster | fail |  |  |  | [`ray-operator/pkg/webhooks/v1/raycluster_mutating_webhook.go`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/pkg/webhooks/v1/raycluster_mutating_webhook.go), [`ray-operator/pkg/webhooks/v1/raycluster_mutating_webhook.go`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/pkg/webhooks/v1/raycluster_mutating_webhook.go) |
| vraycluster.kb.io | validating | /validate-ray-io-v1-raycluster | fail |  |  |  | [`ray-operator/pkg/webhooks/v1/raycluster_validating_webhook.go`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/pkg/webhooks/v1/raycluster_validating_webhook.go), [`ray-operator/pkg/webhooks/v1/raycluster_validating_webhook.go`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/pkg/webhooks/v1/raycluster_validating_webhook.go) |
| vrayjob.kb.io | validating | /validate-ray-io-v1-rayjob | fail |  |  |  | [`ray-operator/pkg/webhooks/v1/rayjob_webhook.go`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/pkg/webhooks/v1/rayjob_webhook.go), [`ray-operator/pkg/webhooks/v1/rayjob_webhook.go`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/pkg/webhooks/v1/rayjob_webhook.go) |
| vrayservice.kb.io | validating | /validate-ray-io-v1-rayservice | fail |  |  |  | [`ray-operator/pkg/webhooks/v1/rayservice_webhook.go`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/pkg/webhooks/v1/rayservice_webhook.go), [`ray-operator/pkg/webhooks/v1/rayservice_webhook.go`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/pkg/webhooks/v1/rayservice_webhook.go) |

### HTTP Endpoints

| Method | Path | Source |
|--------|------|--------|
| * | / | [`experimental/cmd/main.go:111`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/experimental/cmd/main.go#L111) |
| GET | / | [`historyserver/pkg/historyserver/router.go:102`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/historyserver/pkg/historyserver/router.go#L102) |
| GET | / | [`historyserver/pkg/historyserver/router.go:64`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/historyserver/pkg/historyserver/router.go#L64) |
| GET | / | [`historyserver/pkg/historyserver/router.go:55`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/historyserver/pkg/historyserver/router.go#L55) |
| GET | /actors | [`historyserver/pkg/historyserver/router.go:277`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/historyserver/pkg/historyserver/router.go#L277) |
| GET | /actors/{single_actor:*} | [`historyserver/pkg/historyserver/router.go:285`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/historyserver/pkg/historyserver/router.go#L285) |
| * | /api/v1/namespaces/{namespace}/services/{service}/proxy | [`apiserversdk/proxy.go:64`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/apiserversdk/proxy.go#L64) |
| * | /api/v1/namespaces/{namespace}/services/{service}/proxy/ | [`apiserversdk/proxy.go:65`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/apiserversdk/proxy.go#L65) |
| * | /apis/ray.io/v1/ | [`apiserversdk/proxy.go:46`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/apiserversdk/proxy.go#L46) |
| GET | /cluster_status | [`historyserver/pkg/historyserver/router.go:111`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/historyserver/pkg/historyserver/router.go#L111) |
| POST | /events | [`historyserver/pkg/collector/eventserver/eventserver.go:90`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/historyserver/pkg/collector/eventserver/eventserver.go#L90) |
| GET | /grafana_health | [`historyserver/pkg/historyserver/router.go:114`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/historyserver/pkg/historyserver/router.go#L114) |
| GET | /jobs/ | [`historyserver/pkg/historyserver/router.go:121`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/historyserver/pkg/historyserver/router.go#L121) |
| GET | /jobs/{job_id} | [`historyserver/pkg/historyserver/router.go:125`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/historyserver/pkg/historyserver/router.go#L125) |
| * | /livez | [`historyserver/pkg/historyserver/router.go:262`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/historyserver/pkg/historyserver/router.go#L262) |
| GET | /prometheus_health | [`historyserver/pkg/historyserver/router.go:117`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/historyserver/pkg/historyserver/router.go#L117) |
| * | /readz | [`historyserver/pkg/historyserver/router.go:256`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/historyserver/pkg/historyserver/router.go#L256) |
| GET | /v0/cluster_metadata | [`historyserver/pkg/historyserver/router.go:130`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/historyserver/pkg/historyserver/router.go#L130) |
| GET | /v0/logs | [`historyserver/pkg/historyserver/router.go:134`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/historyserver/pkg/historyserver/router.go#L134) |
| GET | /v0/logs/{media_type} | [`historyserver/pkg/historyserver/router.go:140`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/historyserver/pkg/historyserver/router.go#L140) |
| GET | /v0/tasks | [`historyserver/pkg/historyserver/router.go:163`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/historyserver/pkg/historyserver/router.go#L163) |
| GET | /v0/tasks/summarize | [`historyserver/pkg/historyserver/router.go:174`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/historyserver/pkg/historyserver/router.go#L174) |
| GET | /v0/tasks/timeline | [`historyserver/pkg/historyserver/router.go:182`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/historyserver/pkg/historyserver/router.go#L182) |
| GET | /{namespace}/{name}/{session} | [`historyserver/pkg/historyserver/router.go:297`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/historyserver/pkg/historyserver/router.go#L297) |
| GET | /{node_id} | [`historyserver/pkg/historyserver/router.go:91`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/historyserver/pkg/historyserver/router.go#L91) |
| GET | /{subpath:*} | [`historyserver/pkg/historyserver/router.go:191`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/historyserver/pkg/historyserver/router.go#L191) |
| * | GET /api/v1/namespaces/{namespace}/events | [`apiserversdk/proxy.go:47`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/apiserversdk/proxy.go#L47) |
| * | POST /apis/ray.io/v1/namespaces/{namespace}/rayclusters | [`apiserversdk/proxy.go:53`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/apiserversdk/proxy.go#L53) |
| * | POST /apis/ray.io/v1/namespaces/{namespace}/rayjobs | [`apiserversdk/proxy.go:55`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/apiserversdk/proxy.go#L55) |
| * | POST /apis/ray.io/v1/namespaces/{namespace}/rayservices | [`apiserversdk/proxy.go:57`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/apiserversdk/proxy.go#L57) |
| * | PUT /apis/ray.io/v1/namespaces/{namespace}/rayclusters/{name} | [`apiserversdk/proxy.go:54`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/apiserversdk/proxy.go#L54) |
| * | PUT /apis/ray.io/v1/namespaces/{namespace}/rayjobs/{name} | [`apiserversdk/proxy.go:56`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/apiserversdk/proxy.go#L56) |
| * | PUT /apis/ray.io/v1/namespaces/{namespace}/rayservices/{name} | [`apiserversdk/proxy.go:58`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/apiserversdk/proxy.go#L58) |
| * | gateway.networking.k8s.io | [`ray-operator/controllers/ray/authentication_controller.go:455`](https://github.com/ray-project/kuberay/blob/92bd061228ae1b70c7d34858ac1efaf6f271f0db/ray-operator/controllers/ray/authentication_controller.go#L455) |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

### Helm

**Chart:** kuberay-apiserver v1.6.2

