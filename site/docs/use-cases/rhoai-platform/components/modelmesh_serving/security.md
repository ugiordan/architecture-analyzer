# modelmesh-serving: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

### Secrets Referenced

| Name | Type | Referenced By |
|------|------|---------------|
| modelmesh-webhook-server-cert | Opaque | deployment/modelmesh-controller |

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| controller-manager | kube-rbac-proxy | ? | ? | ? | [`config/default/manager_auth_proxy_patch.yaml`](https://github.com/red-hat-data-services/modelmesh-serving/blob/457babd26714baa0e8e684ee23181e2645e58396/config/default/manager_auth_proxy_patch.yaml) |
| controller-manager | manager | ? | ? | ? | [`config/default/manager_auth_proxy_patch.yaml`](https://github.com/red-hat-data-services/modelmesh-serving/blob/457babd26714baa0e8e684ee23181e2645e58396/config/default/manager_auth_proxy_patch.yaml) |
| etcd | etcd | ? | ? | ? | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/modelmesh-serving/blob/457babd26714baa0e8e684ee23181e2645e58396/kustomize:config/overlays/odh) |
| modelmesh-controller | manager | ? | ? | ? | [`kustomize:config/overlays/odh`](https://github.com/red-hat-data-services/modelmesh-serving/blob/457babd26714baa0e8e684ee23181e2645e58396/kustomize:config/overlays/odh) |
| template-value-template-value | mm | ? | ? | ? | [`config/internal/base/deployment.yaml.tmpl`](https://github.com/red-hat-data-services/modelmesh-serving/blob/457babd26714baa0e8e684ee23181e2645e58396/config/internal/base/deployment.yaml.tmpl) |
| template-value-template-value | oauth-proxy | ? | ? | ? | [`config/internal/base/deployment.yaml.tmpl`](https://github.com/red-hat-data-services/modelmesh-serving/blob/457babd26714baa0e8e684ee23181e2645e58396/config/internal/base/deployment.yaml.tmpl) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi9/ubi-minimal:9.5 | 2 | ${USER} |  | multi-arch |  | Unpinned base image: ${DEV_IMAGE} |
| `Dockerfile.develop` | registry.access.redhat.com/ubi9/go-toolset:$GOLANG_VERSION | 1 | root |  | multi-arch |  | Container runs as root user |
| `Dockerfile.develop.ci` | registry.access.redhat.com/ubi9/go-toolset:$GOLANG_VERSION | 1 | root |  | multi-arch |  | Container runs as root user |
| `Dockerfile.konflux` | registry.redhat.io/ubi9/ubi-minimal-pqc@sha256:58d09489b1d9123e24d691e47af6c390fe5e08f50c7f4f2967bb64df7ab72aa0 | 2 | ${USER} |  | multi-arch |  |  |

