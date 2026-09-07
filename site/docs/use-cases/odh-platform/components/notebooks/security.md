# notebooks: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| notebook | notebook | ? | ? | ? | [`jupyter/baseline/ubi9-python-3.12/kustomize/base/statefulset.yaml`](https://github.com/opendatahub-io/notebooks/blob/9ba54dc738b06c12d8844cc62771b4c2cc1e2d42/jupyter/baseline/ubi9-python-3.12/kustomize/base/statefulset.yaml) |
| notebook | notebook | ? | ? | ? | [`jupyter/datascience/ubi9-python-3.12/kustomize/base/statefulset.yaml`](https://github.com/opendatahub-io/notebooks/blob/9ba54dc738b06c12d8844cc62771b4c2cc1e2d42/jupyter/datascience/ubi9-python-3.12/kustomize/base/statefulset.yaml) |
| notebook | notebook | ? | ? | ? | [`jupyter/minimal/ubi9-python-3.12/kustomize/base/statefulset.yaml`](https://github.com/opendatahub-io/notebooks/blob/9ba54dc738b06c12d8844cc62771b4c2cc1e2d42/jupyter/minimal/ubi9-python-3.12/kustomize/base/statefulset.yaml) |
| notebook | notebook | ? | ? | ? | [`jupyter/pytorch/ubi9-python-3.12/kustomize/base/statefulset.yaml`](https://github.com/opendatahub-io/notebooks/blob/9ba54dc738b06c12d8844cc62771b4c2cc1e2d42/jupyter/pytorch/ubi9-python-3.12/kustomize/base/statefulset.yaml) |
| notebook | notebook | ? | ? | ? | [`jupyter/pytorch+llmcompressor/ubi9-python-3.12/kustomize/base/statefulset.yaml`](https://github.com/opendatahub-io/notebooks/blob/9ba54dc738b06c12d8844cc62771b4c2cc1e2d42/jupyter/pytorch+llmcompressor/ubi9-python-3.12/kustomize/base/statefulset.yaml) |
| notebook | notebook | ? | ? | ? | [`jupyter/rocm/pytorch/ubi9-python-3.12/kustomize/base/statefulset.yaml`](https://github.com/opendatahub-io/notebooks/blob/9ba54dc738b06c12d8844cc62771b4c2cc1e2d42/jupyter/rocm/pytorch/ubi9-python-3.12/kustomize/base/statefulset.yaml) |
| notebook | notebook | ? | ? | ? | [`jupyter/rocm/tensorflow/ubi9-python-3.12/kustomize/base/statefulset.yaml`](https://github.com/opendatahub-io/notebooks/blob/9ba54dc738b06c12d8844cc62771b4c2cc1e2d42/jupyter/rocm/tensorflow/ubi9-python-3.12/kustomize/base/statefulset.yaml) |
| notebook | notebook | ? | ? | ? | [`jupyter/tensorflow/ubi9-python-3.12/kustomize/base/statefulset.yaml`](https://github.com/opendatahub-io/notebooks/blob/9ba54dc738b06c12d8844cc62771b4c2cc1e2d42/jupyter/tensorflow/ubi9-python-3.12/kustomize/base/statefulset.yaml) |
| notebook | notebook | ? | ? | ? | [`jupyter/trustyai/ubi9-python-3.12/kustomize/base/statefulset.yaml`](https://github.com/opendatahub-io/notebooks/blob/9ba54dc738b06c12d8844cc62771b4c2cc1e2d42/jupyter/trustyai/ubi9-python-3.12/kustomize/base/statefulset.yaml) |

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `base-images/cpu/c9s-python-3.12/Dockerfile.cpu` | quay.io/centos/centos:stream9@sha256:d323b7623e947245a8eb506fbb0ad0e55eb2ae2d2407b66741a15f372caf9bdc | 2 | ${CNB_USER_ID}:${CNB_GROUP_ID} |  | multi-arch |  |  |
| `base-images/cuda/12.9/c9s-python-3.12/Dockerfile.cuda` | quay.io/opendatahub/odh-midstream-cuda-base-12-9:1.20260810.0 | 2 | 1001 |  | multi-arch |  |  |
| `base-images/cuda/13.0/c9s-python-3.12/Dockerfile.cuda` | quay.io/opendatahub/odh-midstream-cuda-base-13-0:1.20260810.0 | 2 | 1001 |  | multi-arch |  |  |
| `base-images/rocm/7.14/c9s-python-3.12/Dockerfile.rocm` | quay.io/opendatahub/odh-midstream-rocm-base-7-14:1.20260810.0 | 2 | 1001 |  | multi-arch |  |  |
| `codeserver-baseline/ubi9-python-3.12/Dockerfile.konflux.cpu` | codeserver | 5 | 1001 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE}; Unpinned base image: cpu-base; Unpinned base image: codeserver; Unpinned base image: codeserver |
| `codeserver/ubi9-python-3.12/Dockerfile.konflux.cpu` | codeserver | 5 | 1001 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE}; Unpinned base image: cpu-base; Unpinned base image: codeserver; Unpinned base image: codeserver |
| `jupyter/baseline/ubi9-python-3.12/Dockerfile.konflux.cpu` | jupyter-minimal | 3 | 1001 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: cpu-base; Unpinned base image: jupyter-minimal |
| `jupyter/datascience/ubi9-python-3.12/Dockerfile.konflux.cpu` | jupyter-minimal | 4 | 1001 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: cpu-base; Unpinned base image: jupyter-minimal |
| `jupyter/minimal/ubi9-python-3.12/Dockerfile.konflux.cpu` | cpu-base | 2 | 1001 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: cpu-base |
| `jupyter/minimal/ubi9-python-3.12/Dockerfile.konflux.cuda` | cuda-base | 2 | 1001 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: cuda-base |
| `jupyter/minimal/ubi9-python-3.12/Dockerfile.konflux.rocm` | rocm-base | 2 | 1001 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: rocm-base |
| `jupyter/pytorch+llmcompressor/ubi9-python-3.12/Dockerfile.konflux.cuda` | cuda-jupyter-datascience | 5 | 1001 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: cuda-base; Unpinned base image: cuda-jupyter-minimal; Unpinned base image: cuda-jupyter-datascience |
| `jupyter/pytorch/ubi9-python-3.12/Dockerfile.konflux.cuda` | cuda-jupyter-datascience | 5 | 1001 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: cuda-base; Unpinned base image: cuda-jupyter-minimal; Unpinned base image: cuda-jupyter-datascience |
| `jupyter/rocm/pytorch/ubi9-python-3.12/Dockerfile.konflux.rocm` | rocm-jupyter-datascience | 5 | 1001 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: rocm-base; Unpinned base image: rocm-jupyter-minimal; Unpinned base image: rocm-jupyter-datascience |
| `jupyter/rocm/tensorflow/ubi9-python-3.12/Dockerfile.konflux.rocm` | rocm-jupyter-datascience | 5 | 1001 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: rocm-base; Unpinned base image: rocm-jupyter-minimal; Unpinned base image: rocm-jupyter-datascience |
| `jupyter/tensorflow/ubi9-python-3.12/Dockerfile.konflux.cuda` | cuda-jupyter-datascience | 5 | 1001 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: cuda-base; Unpinned base image: cuda-jupyter-minimal; Unpinned base image: cuda-jupyter-datascience |
| `jupyter/trustyai/ubi9-python-3.12/Dockerfile.konflux.cpu` | jupyter-datascience | 5 | 1001 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: cpu-base; Unpinned base image: jupyter-minimal; Unpinned base image: jupyter-datascience |
| `runtimes/baseline/ubi9-python-3.12/Dockerfile.konflux.cpu` | cpu-base | 2 | 1001 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: cpu-base |
| `runtimes/datascience/ubi9-python-3.12/Dockerfile.konflux.cpu` | cpu-base | 2 | 1001 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: cpu-base |
| `runtimes/minimal/ubi9-python-3.12/Dockerfile.konflux.cpu` | cpu-base | 2 | 1001 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: cpu-base |
| `runtimes/pytorch+llmcompressor/ubi9-python-3.12/Dockerfile.konflux.cuda` | cuda-base | 2 | 1001 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: cuda-base |
| `runtimes/pytorch/ubi9-python-3.12/Dockerfile.konflux.cuda` | cuda-base | 2 | 1001 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: cuda-base |
| `runtimes/rocm-pytorch/ubi9-python-3.12/Dockerfile.konflux.rocm` | rocm-base | 2 | 1001 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: rocm-base |
| `runtimes/rocm-tensorflow/ubi9-python-3.12/Dockerfile.konflux.rocm` | rocm-base | 2 | 1001 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: rocm-base |
| `runtimes/tensorflow/ubi9-python-3.12/Dockerfile.konflux.cuda` | cuda-base | 2 | 1001 |  | multi-arch |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: cuda-base |
| `scripts/buildinputs/Dockerfile` | scratch | 2 | 65532:0 |  | multi-arch |  | Unpinned base image: scratch |
| `scripts/check-payload/Dockerfile` | registry.access.redhat.com/ubi9-minimal:latest | 2 | 65532:0 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi9-minimal:latest |
| `scripts/lockfile-generators/Dockerfile.rpm-lockfile` | ${BASE_IMAGE} | 1 | root |  |  |  | Unpinned base image: ${BASE_IMAGE}; Container runs as root user |

