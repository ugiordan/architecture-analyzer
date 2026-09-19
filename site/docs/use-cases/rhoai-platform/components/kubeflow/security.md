# kubeflow: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| kf-notebook-controller-metrics-tls | Opaque | deployment/deployment |
| odh-notebook-controller-metrics-tls | Opaque | deployment/manager |
| odh-notebook-controller-webhook-cert | kubernetes.io/tls | service/webhook-service |
| webhook-server-cert | Opaque | deployment/controller-manager |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| controller-manager | kube-rbac-proxy | ? | ? | ? | [`components/notebook-controller/config/default/manager_auth_proxy_patch.yaml`](https://github.com/red-hat-data-services/kubeflow/blob/c8d3879d766dcf243b9bf1666709e6a4dd074039/components/notebook-controller/config/default/manager_auth_proxy_patch.yaml) |
| controller-manager | manager | ? | ? | ? | [`components/notebook-controller/config/default/manager_auth_proxy_patch.yaml`](https://github.com/red-hat-data-services/kubeflow/blob/c8d3879d766dcf243b9bf1666709e6a4dd074039/components/notebook-controller/config/default/manager_auth_proxy_patch.yaml) |
| controller-manager | manager | ? | ? | ? | [`components/notebook-controller/config/default/manager_image_patch.yaml`](https://github.com/red-hat-data-services/kubeflow/blob/c8d3879d766dcf243b9bf1666709e6a4dd074039/components/notebook-controller/config/default/manager_image_patch.yaml) |
| controller-manager | manager | ? | ? | ? | [`components/notebook-controller/config/default/manager_prometheus_metrics_patch.yaml`](https://github.com/red-hat-data-services/kubeflow/blob/c8d3879d766dcf243b9bf1666709e6a4dd074039/components/notebook-controller/config/default/manager_prometheus_metrics_patch.yaml) |
| controller-manager | manager | ? | ? | ? | [`components/notebook-controller/config/default/manager_webhook_patch.yaml`](https://github.com/red-hat-data-services/kubeflow/blob/c8d3879d766dcf243b9bf1666709e6a4dd074039/components/notebook-controller/config/default/manager_webhook_patch.yaml) |
| deployment | manager | ? | ? | ? | [`components/notebook-controller/config/manager/manager.yaml`](https://github.com/red-hat-data-services/kubeflow/blob/c8d3879d766dcf243b9bf1666709e6a4dd074039/components/notebook-controller/config/manager/manager.yaml) |
| deployment | manager | ? | ? | ? | [`components/notebook-controller/config/overlays/openshift/manager_openshift_patch.yaml`](https://github.com/red-hat-data-services/kubeflow/blob/c8d3879d766dcf243b9bf1666709e6a4dd074039/components/notebook-controller/config/overlays/openshift/manager_openshift_patch.yaml) |
| deployment | manager | true | true | ? | [`components/notebook-controller/config/overlays/openshift/manager_tls_patch.yaml`](https://github.com/red-hat-data-services/kubeflow/blob/c8d3879d766dcf243b9bf1666709e6a4dd074039/components/notebook-controller/config/overlays/openshift/manager_tls_patch.yaml) |
| manager | manager | ? | ? | ? | [`components/odh-notebook-controller/config/manager/manager.yaml`](https://github.com/red-hat-data-services/kubeflow/blob/c8d3879d766dcf243b9bf1666709e6a4dd074039/components/odh-notebook-controller/config/manager/manager.yaml) |
| manager | manager | ? | true | ? | [`components/odh-notebook-controller/config/overlays/openshift/manager_tls_patch.yaml`](https://github.com/red-hat-data-services/kubeflow/blob/c8d3879d766dcf243b9bf1666709e6a4dd074039/components/odh-notebook-controller/config/overlays/openshift/manager_tls_patch.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `components/notebook-controller/Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 1001:0 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `components/notebook-controller/Dockerfile.ci` | gcr.io/distroless/base:debug | 2 |  |  |  |  | No USER directive found (defaults to root) |
| `components/notebook-controller/Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:fac0bde5012de6a8475bb5dc18db968cdc90b12d9dc2ffaddb4a4975d0b8a04f | 2 | 1001:0 |  | multi-arch |  |  |
| `components/odh-notebook-controller/Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 1001:0 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `components/odh-notebook-controller/Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:fac0bde5012de6a8475bb5dc18db968cdc90b12d9dc2ffaddb4a4975d0b8a04f | 2 | 1001:0 |  | multi-arch |  |  |

