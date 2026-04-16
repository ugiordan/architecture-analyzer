# trustyai-service-operator: Security

## Secrets

## Deployment Security Controls

### Container Security Contexts

| Deployment | Container | RunAsNonRoot | ReadOnlyFS | Privileged | Source |
|------------|-----------|--------------|------------|------------|--------|
| controller-manager | manager | true | ? | ? | `config/manager/manager.yaml` |

## Build Security

| Path | Base Image | Stages | User | Ports | Architectures | FIPS | Issues |
|------|------------|--------|------|-------|---------------|------|--------|
| `Dockerfile` | registry.access.redhat.com/ubi8/ubi-minimal:latest | 2 | 65532:65532 |  | multi-arch |  | Unpinned base image: registry.access.redhat.com/ubi8/ubi-minimal:latest |
| `Dockerfile.driver` | registry.access.redhat.com/ubi9/ubi-minimal:latest | 2 | 65532:65532 |  |  |  | Unpinned base image: registry.access.redhat.com/ubi9/ubi-minimal:latest |
| `Dockerfile.lmes-job` | registry.access.redhat.com/ubi9/python-311@sha256:fccda5088dd13d2a3f2659e4c904beb42fc164a0c909e765f01af31c58affae3 | 1 | 65532:65532 |  |  |  |  |
| `Dockerfile.orchestrator` | ${UBI_MINIMAL_BASE_IMAGE}:${UBI_BASE_IMAGE_TAG} | 6 | orchestr8 |  |  |  | Unpinned base image: rust-builder; Unpinned base image: fms-guardrails-orchestr8-builder; Unpinned base image: fms-guardrails-orchestr8-builder; Unpinned base image: fms-guardrails-orchestr8-builder |
| `tests/Dockerfile` | registry.access.redhat.com/ubi8:8.10-1020 | 1 |  |  |  |  | No USER directive found (defaults to root) |

## Cache Architecture

### Manager Configuration

| Property | Value |
|----------|-------|
| Manager file | `cmd/main.go` |
| Cache scope | cluster-wide |
| DefaultTransform | no |
| Memory limit | 700Mi |

### Issues

- No GOMEMLIMIT set in deployment (Go GC cannot pressure-tune)
- No cache configuration: all informers are cluster-wide (OOM risk)
- Type ConfigMap is watched but has no cache filter (cluster-wide informer)
- Type Deployment is watched but has no cache filter (cluster-wide informer)
- Type EvalHub is watched but has no cache filter (cluster-wide informer)
- Type GuardrailsOrchestrator is watched but has no cache filter (cluster-wide informer)
- Type InferenceService is watched but has no cache filter (cluster-wide informer)
- Type LMEvalJob is watched but has no cache filter (cluster-wide informer)
- Type NemoGuardrails is watched but has no cache filter (cluster-wide informer)
- Type Service is watched but has no cache filter (cluster-wide informer)
- Type TrustyAIService is watched but has no cache filter (cluster-wide informer)

