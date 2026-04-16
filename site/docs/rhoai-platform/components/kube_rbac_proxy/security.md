# kube-rbac-proxy: Security

## Secrets

## Deployment Security Controls

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| kube-rbac-proxy | kube-rbac-proxy | ? | ? | ? | `examples/non-resource-url/deployment.yaml` |
| kube-rbac-proxy | prometheus-example-app | ? | ? | ? | `examples/non-resource-url/deployment.yaml` |
| kube-rbac-proxy | kube-rbac-proxy | ? | ? | ? | `examples/non-resource-url-token-request/deployment.yaml` |
| kube-rbac-proxy | prometheus-example-app | ? | ? | ? | `examples/non-resource-url-token-request/deployment.yaml` |
| kube-rbac-proxy | kube-rbac-proxy | ? | ? | ? | `examples/oidc/deployment.yaml` |
| kube-rbac-proxy | prometheus-example-app | ? | ? | ? | `examples/oidc/deployment.yaml` |
| kube-rbac-proxy | kube-rbac-proxy | ? | ? | ? | `examples/resource-attributes/deployment.yaml` |
| kube-rbac-proxy | prometheus-example-app | ? | ? | ? | `examples/resource-attributes/deployment.yaml` |
| kube-rbac-proxy | kube-rbac-proxy | ? | ? | ? | `examples/rewrites/deployment.yaml` |
| kube-rbac-proxy | prometheus-example-app | ? | ? | ? | `examples/rewrites/deployment.yaml` |
| kube-rbac-proxy | kube-rbac-proxy | ? | ? | ? | `examples/static-auth/deployment.yaml` |
| kube-rbac-proxy | prometheus-example-app | ? | ? | ? | `examples/static-auth/deployment.yaml` |
| kube-rbac-proxy | kube-rbac-proxy | ? | ? | ? | `test/kubetest/testtemplates/data/deployment.yaml` |
| kube-rbac-proxy | prometheus-example-app | ? | ? | ? | `test/kubetest/testtemplates/data/deployment.yaml` |

## Build Security

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | $BASEIMAGE | 1 | 65532:65532 |  |  |  | Unpinned base image: $BASEIMAGE |
| `Dockerfile.ocp` | registry.ci.openshift.org/ocp/4.22:base-rhel9 | 2 | 65534 |  |  |  |  |
| `examples/example-client/Dockerfile` | alpine:3.12 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `examples/example-client-http2/Dockerfile` |  | 0 |  |  |  |  | No USER directive found (defaults to root) |
| `examples/example-client-urlquery/Dockerfile` | alpine:3.7 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `examples/grpcc/Dockerfile` | node:8.9.4-alpine | 1 |  |  |  |  | No USER directive found (defaults to root) |

