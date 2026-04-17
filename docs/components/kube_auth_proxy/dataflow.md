# kube-auth-proxy: Dataflow

## Controller Watches

Kubernetes resources this controller monitors for changes. Each watch triggers reconciliation when the watched resource is created, updated, or deleted.

No controller watches found.

## Reconciliation Flow

How the controller interacts with the Kubernetes API during reconciliation.

```mermaid
sequenceDiagram
    %% Static dataflow for kube-auth-proxy

    participant KubernetesAPI as Kubernetes API
    participant kube_auth_proxy as kube-auth-proxy
    participant example_app as example-app
    participant kube_rbac_proxy as kube-rbac-proxy
    participant kube_rbac_proxy_verb_override as kube-rbac-proxy-verb-override


    Note over kube_auth_proxy: Exposed Services
    Note right of kube_auth_proxy: kube-rbac-proxy:8443/TCP [https]
    Note right of kube_auth_proxy: kube-rbac-proxy:8443/TCP [https]
    Note right of kube_auth_proxy: kube-rbac-proxy:8443/TCP [https]
    Note right of kube_auth_proxy: kube-rbac-proxy:8443/TCP [https]
    Note right of kube_auth_proxy: kube-rbac-proxy:8443/TCP [https]
    Note right of kube_auth_proxy: kube-rbac-proxy:8443/TCP [https]
    Note right of kube_auth_proxy: kube-rbac-proxy:8443/TCP [https]
    Note right of kube_auth_proxy: kube-rbac-proxy:8443/TCP [https]
    Note right of kube_auth_proxy: kube-rbac-proxy:8443/TCP [https]
    Note right of kube_auth_proxy: kube-rbac-proxy:8443/TCP [https]
    Note right of kube_auth_proxy: kube-rbac-proxy:8443/TCP [https]
```

### HTTP Endpoints

| Method | Path | Source |
|--------|------|--------|
| * | / | `kube-rbac-proxy/cmd/kube-rbac-proxy/app/kube-rbac-proxy.go:324` |
| * | /.well-known/oauth-authorization-server | `test/integration/testutil/mock_openshift_oauth.go:41` |
| * | /oauth/authorize | `test/integration/testutil/mock_openshift_oauth.go:42` |
| * | /oauth/token | `test/integration/testutil/mock_openshift_oauth.go:43` |
| * | /apis/user.openshift.io/v1/users/~ | `test/integration/testutil/mock_openshift_oauth.go:44` |

## Configuration

ConfigMaps and Helm values that control this component's runtime behavior.

### ConfigMaps

| Name | Data Keys | Source |
|------|-----------|--------|
| kube-rbac-proxy | config-file.yaml | `kube-rbac-proxy/test/e2e/static-auth/configmap-non-resource.yaml` |
| kube-rbac-proxy | config-file.yaml | `kube-rbac-proxy/test/e2e/static-auth/configmap-resource.yaml` |

### Helm

**Chart:** kubernetes v7.14.1

