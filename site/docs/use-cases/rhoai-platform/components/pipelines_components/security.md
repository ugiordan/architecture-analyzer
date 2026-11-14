# pipelines-components: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.redhat.io/ubi9/python-311@sha256:d7620b96616955d78425518143affdc9463fb1e71d00aa2b7dc2785c54621a0b | 1 | 1001 |  |  |  |  |
| `Dockerfile.konflux.automl` | ${BASE_IMAGE} | 2 | default |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `Dockerfile.konflux.autorag` | ${BASE_IMAGE} | 4 | default |  |  |  | Unpinned base image: ${DOCLING_OGX_MODELCAR}; Unpinned base image: ${WHISPER_MODELCAR}; Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `Dockerfile.konflux.autorag-inference` | ${BASE_IMAGE} | 1 | default |  |  |  | Unpinned base image: ${BASE_IMAGE} |
| `Dockerfile.konflux.pipelines-components` | registry.redhat.io/ubi9/python-312:9.8@sha256:f72227ffc4f3faa3772a6ae91ffc03bcf3c507d6b7a525d2bfe90b3579dcb8af | 1 | 1001 |  |  |  |  |
| `components/training/autorag/shared/assets_generator/starter_kit_templates/agentic_rag/Containerfile.openshell` | ${BASE_IMAGE} | 3 | 1001 |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${PYTHON_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `pipelines/training/automl/autogluon_tabular_training_pipeline/Containerfile` | ${BASE_IMAGE} | 2 |  |  |  |  | Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE}; No USER directive found (defaults to root) |
| `pipelines/training/autorag/documents_rag_optimization_pipeline/Containerfile` | ${BASE_IMAGE} | 4 | default |  |  |  | Unpinned base image: ${DOCLING_OGX_MODELCAR}; Unpinned base image: ${WHISPER_MODELCAR}; Unpinned base image: ${BASE_IMAGE}; Unpinned base image: ${BASE_IMAGE} |
| `pipelines/training/autorag/documents_rag_optimization_pipeline/inference/Containerfile` | ${BASE_IMAGE} | 1 |  |  |  |  | Unpinned base image: ${BASE_IMAGE}; No USER directive found (defaults to root) |

