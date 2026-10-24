# data-science-pipelines-operator: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| data-science-pipelines-operator-controller-manager | manager | ? | true | ? | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/2b0cb79c44205212347ffb26bafc298b91bee952/kustomize:config/overlays/odh) |
| ds-pipeline-workflow-controller-template-value | ds-pipeline-workflow-controller | true | true | ? | [`config/internal/workflow-controller/deployment.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/2b0cb79c44205212347ffb26bafc298b91bee952/config/internal/workflow-controller/deployment.yaml.tmpl) |
| mariadb-template-value | mariadb | ? | ? | ? | [`config/internal/mariadb/default/deployment.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/2b0cb79c44205212347ffb26bafc298b91bee952/config/internal/mariadb/default/deployment.yaml.tmpl) |
| minio-template-value | minio | ? | ? | ? | [`config/internal/minio/default/deployment.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/2b0cb79c44205212347ffb26bafc298b91bee952/config/internal/minio/default/deployment.yaml.tmpl) |
| template-value | ds-pipeline-api-server | ? | ? | ? | [`config/internal/apiserver/default/deployment.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/2b0cb79c44205212347ffb26bafc298b91bee952/config/internal/apiserver/default/deployment.yaml.tmpl) |
| template-value | kube-rbac-proxy | ? | ? | ? | [`config/internal/apiserver/default/deployment.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/2b0cb79c44205212347ffb26bafc298b91bee952/config/internal/apiserver/default/deployment.yaml.tmpl) |
| template-value | kube-rbac-proxy-metrics | ? | ? | ? | [`config/internal/apiserver/default/deployment.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/2b0cb79c44205212347ffb26bafc298b91bee952/config/internal/apiserver/default/deployment.yaml.tmpl) |
| template-value | ds-pipeline-persistenceagent | ? | ? | ? | [`config/internal/persistence-agent/deployment.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/2b0cb79c44205212347ffb26bafc298b91bee952/config/internal/persistence-agent/deployment.yaml.tmpl) |
| template-value | ds-pipeline-scheduledworkflow | ? | ? | ? | [`config/internal/scheduled-workflow/deployment.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/2b0cb79c44205212347ffb26bafc298b91bee952/config/internal/scheduled-workflow/deployment.yaml.tmpl) |
| template-value | ds-pipeline-webhook | ? | ? | ? | [`config/internal/webhook/deployment.yaml.tmpl`](https://github.com/red-hat-data-services/data-science-pipelines-operator/blob/2b0cb79c44205212347ffb26bafc298b91bee952/config/internal/webhook/deployment.yaml.tmpl) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 4 | ${USER}:${USER} |  | multi-arch |  | Unpinned base image: go-toolset-${BUILDER_ARCH}; Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:d0fdfb1401ee754872f60c2527563671c16112080e5254139133908178700911 | 2 | ${USER}:${USER} |  | multi-arch |  |  |

