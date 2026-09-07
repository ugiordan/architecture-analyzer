# NeMo-Guardrails: Security

## Secrets

Kubernetes secrets referenced by this component. Only names and types are shown, not values.

## Deployment Security Controls

SecurityContext settings on pod and container specs. These control privilege escalation, filesystem access, and user identity.

## Build Security

Dockerfile patterns and base image analysis. Covers supply chain security: base images, build stages, runtime user, FIPS compliance.

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | python:3.12-slim | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `Dockerfile.konflux` | ${BASE} | 6 | 1001 |  |  |  | Unpinned base image: ${MODELCAR_MINILM}; Unpinned base image: ${MODELCAR_SNOWFLAKE}; Unpinned base image: ${MODELCAR_SPACY}; Unpinned base image: ${MODELCAR_NLTK}; Unpinned base image: ${BASE}; Unpinned base image: ${BASE} |
| `Dockerfile.server` | registry.access.redhat.com/ubi9/python-312:latest | 2 | 1001 |  |  |  | Unpinned base image: registry.access.redhat.com/ubi9/python-312:latest; Unpinned base image: registry.access.redhat.com/ubi9/python-312:latest |
| `nemoguardrails/library/factchecking/align_score/Dockerfile` | python:3.10 | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `nemoguardrails/library/jailbreak_detection/Dockerfile` | python:3.11-slim | 1 |  |  |  |  | No USER directive found (defaults to root) |
| `qa/Dockerfile.qa` | python:3.10 | 1 |  |  |  |  | No USER directive found (defaults to root) |

