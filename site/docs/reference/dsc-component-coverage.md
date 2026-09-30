# DataScienceCluster Component Coverage

Source: [DataScienceCluster v2 component definitions](https://github.com/opendatahub-io/opendatahub-operator/blob/main/api/datasciencecluster/v2/datasciencecluster_types.go)

This table compares active `spec.components` keys with the repositories in
`scan-config.yaml`.

| DataScienceCluster key | Repository coverage | Status |
|---|---|---|
| `dashboard` | `opendatahub-io/odh-dashboard`, `red-hat-data-services/odh-dashboard` | Covered |
| `workbenches` | `opendatahub-io/notebooks`, `red-hat-data-services/notebooks` | Covered |
| `aipipelines` | `data-science-pipelines`, `data-science-pipelines-operator` | Covered |
| `kserve` | `kserve` | Covered |
| `kueue` | `kueue` | Covered |
| `ray` | `kuberay` | Covered |
| `trustyai` | `trustyai-service-operator` | Covered |
| `modelregistry` | `model-registry`, `model-registry-operator` | Covered |
| `trainingoperator` | Historical `training-operator` entry remains only for downstream RHOAI | Deprecated in upstream v2 |
| `feastoperator` | `feast` | Covered |
| `llamastackoperator` | `llama-stack` plus `ogx-k8s-operator` | Deprecated, replaced by OGX |
| `ogx` | `ogx-k8s-operator` | Covered |
| `mlflowoperator` | `mlflow-operator` | Covered |
| `trainer` | `opendatahub-io/trainer-operator`; downstream replacement still needs confirmation | ODH covered, RHOAI pending |
| `sparkoperator` | `spark-operator` | Covered |
| `aigateway` | `ai-gateway-payload-processing` | Covered |
| `mcplifecycleoperator` | `opendatahub-io/mcp-lifecycle-module-operator` | Added to ODH scan |

The following status-only fields are submodules, not additional top-level
`spec.components` keys:

- `maasConsumerPortal`: submodule of Dashboard
- `workbenchesV2`: submodule of Workbenches
- `modelsAsAService`: submodule of AIGateway, covered by `models-as-a-service`
- `batchGateway`: submodule of AIGateway, covered by `batch-gateway`

The upstream v2 source explicitly marks `trainingoperator` and
`llamastackoperator` as backward-compatibility fields. They should not be
treated as missing active components.
