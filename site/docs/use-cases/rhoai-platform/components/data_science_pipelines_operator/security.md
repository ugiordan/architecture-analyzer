# data-science-pipelines-operator: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| data-science-pipelines-operator-controller-manager | manager | ? | true | ? | [`kustomize:config/overlays/odh`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/a7254f1bf9f155a6e91359737335bcc6d4d18e1f/kustomize:config/overlays/odh) |
| ds-pipeline-workflow-controller-template-value | ds-pipeline-workflow-controller | true | true | ? | [`config/internal/workflow-controller/deployment.yaml.tmpl`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/a7254f1bf9f155a6e91359737335bcc6d4d18e1f/config/internal/workflow-controller/deployment.yaml.tmpl) |
| mariadb-template-value | mariadb | ? | ? | ? | [`config/internal/mariadb/default/deployment.yaml.tmpl`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/a7254f1bf9f155a6e91359737335bcc6d4d18e1f/config/internal/mariadb/default/deployment.yaml.tmpl) |
| minio-template-value | minio | ? | ? | ? | [`config/internal/minio/default/deployment.yaml.tmpl`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/a7254f1bf9f155a6e91359737335bcc6d4d18e1f/config/internal/minio/default/deployment.yaml.tmpl) |
| template-value | ds-pipeline-api-server | ? | ? | ? | [`config/internal/apiserver/default/deployment.yaml.tmpl`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/a7254f1bf9f155a6e91359737335bcc6d4d18e1f/config/internal/apiserver/default/deployment.yaml.tmpl) |
| template-value | kube-rbac-proxy | ? | ? | ? | [`config/internal/apiserver/default/deployment.yaml.tmpl`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/a7254f1bf9f155a6e91359737335bcc6d4d18e1f/config/internal/apiserver/default/deployment.yaml.tmpl) |
| template-value | ds-pipeline-persistenceagent | ? | ? | ? | [`config/internal/persistence-agent/deployment.yaml.tmpl`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/a7254f1bf9f155a6e91359737335bcc6d4d18e1f/config/internal/persistence-agent/deployment.yaml.tmpl) |
| template-value | ds-pipeline-scheduledworkflow | ? | ? | ? | [`config/internal/scheduled-workflow/deployment.yaml.tmpl`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/a7254f1bf9f155a6e91359737335bcc6d4d18e1f/config/internal/scheduled-workflow/deployment.yaml.tmpl) |
| template-value | ds-pipeline-webhook | ? | ? | ? | [`config/internal/webhook/deployment.yaml.tmpl`](https://github.com/opendatahub-io/data-science-pipelines-operator/blob/a7254f1bf9f155a6e91359737335bcc6d4d18e1f/config/internal/webhook/deployment.yaml.tmpl) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 4 | ${USER}:${USER} |  | multi-arch |  | Unpinned base image: go-toolset-${BUILDER_ARCH}; Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:fac0bde5012de6a8475bb5dc18db968cdc90b12d9dc2ffaddb4a4975d0b8a04f | 2 | ${USER}:${USER} |  | multi-arch |  |  |

