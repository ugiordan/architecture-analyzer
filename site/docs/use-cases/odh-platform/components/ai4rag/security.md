# ai4rag: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `ai4rag/assets_generator/starter_kit_templates/agentic_rag/Containerfile.openshell` | ${BASE_IMAGE} | 1 | 1001 |  |  |  | Unpinned base image: ${BASE_IMAGE} |

