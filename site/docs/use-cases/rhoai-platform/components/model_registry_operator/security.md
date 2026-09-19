# model-registry-operator: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| aihub-controller-manager-metrics-tls | Opaque | deployment/aihub-controller-manager |
| catalog-webhook-service | Opaque | deployment/catalog-controller-manager |
| controller-manager-metrics-service | Opaque | deployment/controller-manager |
| webhook-server-cert | Opaque | deployment/controller-manager |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| aihub-controller-manager | manager | ? | true | ? | [`config/overlays/aihub/manager.yaml`](https://github.com/opendatahub-io/model-registry-operator/blob/6ca32433140dc135d3b5cbc168007f346bbce959/config/overlays/aihub/manager.yaml) |
| catalog-controller-manager | manager | ? | true | ? | [`kustomize:config/overlays/odh`](https://github.com/opendatahub-io/model-registry-operator/blob/6ca32433140dc135d3b5cbc168007f346bbce959/kustomize:config/overlays/odh) |
| model-registry-operator-controller-manager | manager | ? | ? | ? | [`kustomize:config/overlays/odh`](https://github.com/opendatahub-io/model-registry-operator/blob/6ca32433140dc135d3b5cbc168007f346bbce959/kustomize:config/overlays/odh) |
| template-value | rest-container | ? | ? | ? | [`internal/controller/config/templates/deployment.yaml.tmpl`](https://github.com/opendatahub-io/model-registry-operator/blob/6ca32433140dc135d3b5cbc168007f346bbce959/internal/controller/config/templates/deployment.yaml.tmpl) |
| template-value | kube-rbac-proxy | ? | ? | ? | [`internal/controller/config/templates/deployment.yaml.tmpl`](https://github.com/opendatahub-io/model-registry-operator/blob/6ca32433140dc135d3b5cbc168007f346bbce959/internal/controller/config/templates/deployment.yaml.tmpl) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:fac0bde5012de6a8475bb5dc18db968cdc90b12d9dc2ffaddb4a4975d0b8a04f | 2 | 65532:65532 |  |  |  |  |

