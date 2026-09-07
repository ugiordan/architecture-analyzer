# spark-operator: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| webhook-server-cert | Opaque | deployment/controller-manager |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| controller-manager | manager | ? | ? | ? | [`config/default/manager_webhook_patch.yaml`](https://github.com/kubeflow/spark-operator/blob/31bb82c09e1f7240d193625af0e87b3a042fb5fa/config/default/manager_webhook_patch.yaml) |
| spark-operator-controller | controller | true | true | false | [`kustomize:config/overlays/odh`](https://github.com/kubeflow/spark-operator/blob/31bb82c09e1f7240d193625af0e87b3a042fb5fa/kustomize:config/overlays/odh) |
| spark-operator-module-controller-manager | manager | true | true | false | [`spark-operator-module/config/manager/manager.yaml`](https://github.com/kubeflow/spark-operator/blob/31bb82c09e1f7240d193625af0e87b3a042fb5fa/spark-operator-module/config/manager/manager.yaml) |
| spark-operator-webhook | webhook | true | true | false | [`kustomize:config/overlays/odh`](https://github.com/kubeflow/spark-operator/blob/31bb82c09e1f7240d193625af0e87b3a042fb5fa/kustomize:config/overlays/odh) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | ${SPARK_IMAGE} | 2 | ${SPARK_UID}:${SPARK_GID} |  | multi-arch |  | Unpinned base image: ${SPARK_IMAGE} |
| `Dockerfile.konflux` | ${BASE_IMAGE} | 2 | ${SPARK_UID}:${SPARK_GID} |  | multi-arch |  | Unpinned base image: ${GO_BUILDER_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `Dockerfile.konflux.module-controller` | ${BASE_IMAGE} | 2 | 1000:1000 |  | multi-arch |  | Unpinned base image: ${GO_BUILDER_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `docker/Dockerfile.kubectl` | ${BASE_IMAGE} | 2 |  |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE}; No USER directive found (defaults to root) |
| `spark-docker/Dockerfile` | ${SPARK_IMAGE} | 1 | ${spark_uid} |  |  |  | Unpinned base image: ${SPARK_IMAGE} |
| `spark-operator-module-controller.Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:9.6@sha256:8201445bebcb5bd4fe23fcc2a76cd5fec029ab401d270926a1563c03b36f0137 | 2 | 1000:1000 |  |  |  |  |

