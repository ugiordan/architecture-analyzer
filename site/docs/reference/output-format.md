# Output Format

## component-architecture.json

The core output format. Contains all data extracted by the 22 extractors.

### Top-level structure

```json
{
  "component": "my-operator",
  "repo": "github.com/org/my-operator",
  "extracted_at": "2026-04-14T10:30:00Z",
  "analyzer_version": "0.2.0",
  "crds": [],
  "rbac": {},
  "services": [],
  "deployments": [],
  "network_policies": [],
  "controller_watches": {},
  "dependencies": {},
  "secrets_referenced": [],
  "dockerfiles": [],
  "helm": {},
  "webhooks": [],
  "configmaps": [],
  "http_endpoints": [],
  "ingress_routing": [],
  "external_connections": [],
  "feature_gates": [],
  "cache_config": {},
  "operator_config": [],
  "reconcile_sequences": [],
  "prometheus_metrics": [],
  "status_conditions": [],
  "platform_detection": {},
  "ingress_routing": [],
  "template_files": [],
  "security_annotations": []
}
```

### Key types

#### CRD

```json
{
  "group": "datasciencecluster.opendatahub.io",
  "version": "v1",
  "kind": "DataScienceCluster",
  "scope": "Cluster",
  "field_count": 42,
  "cel_rules": 3,
  "source_file": "config/crd/bases/datasciencecluster.yaml"
}
```

#### RBAC

```json
{
  "cluster_roles": [
    {
      "name": "manager-role",
      "rules": [
        {
          "api_groups": [""],
          "resources": ["secrets"],
          "verbs": ["get", "list", "watch"]
        }
      ],
      "source": "config/rbac/role.yaml"
    }
  ],
  "role_bindings": [],
  "kubebuilder_markers": []
}
```

#### Controller Watches

```json
{
  "controllers": [
    {
      "name": "DSCController",
      "file": "controllers/dsc_controller.go",
      "for": {
        "group": "datasciencecluster.opendatahub.io",
        "version": "v1",
        "kind": "DataScienceCluster"
      },
      "owns": [
        { "group": "apps", "version": "v1", "kind": "Deployment" }
      ],
      "watches": [
        { "group": "", "version": "v1", "kind": "ConfigMap" }
      ]
    }
  ]
}
```

#### Dependencies

```json
{
  "go_version": "1.25",
  "toolchain": "go1.25.0",
  "go_modules": [
    { "module": "sigs.k8s.io/controller-runtime", "version": "v0.23.3" }
  ],
  "replace_directives": [
    {
      "original": "github.com/org/old-module",
      "replacement": "github.com/org/new-module",
      "version": "v1.2.0"
    }
  ],
  "internal_odh": [
    {
      "component": "opendatahub-operator",
      "interaction": "Go module dependency: github.com/opendatahub-io/opendatahub-operator/v2"
    }
  ]
}
```

#### External Connections

```json
[
  {
    "type": "database",
    "service": "postgres",
    "target": "postgres://***@db.example.com:5432/mydb",
    "source": "pkg/storage/db.go:42",
    "function": "NewStore"
  },
  {
    "type": "messaging",
    "service": "kafka",
    "target": "",
    "source": "pkg/events/producer.go:18",
    "function": "InitProducer"
  }
]
```

#### Feature Gates

```json
[
  {
    "name": "PipelineReuse",
    "default": true,
    "pre_release": "Beta",
    "source": "pkg/features/gates.go:15"
  },
  {
    "name": "ExperimentalAPI",
    "default": false,
    "pre_release": "Alpha",
    "source": "pkg/features/gates.go:16"
  },
  {
    "name": "DebugMode",
    "default": true,
    "source": "cmd/main.go:42",
    "runtime_set": true
  }
]
```

#### Cache Config

```json
{
  "scope": "cluster",
  "filtered_types": ["ConfigMap", "Secret"],
  "disabled_types": [],
  "implicit_informers": [
    {
      "type": "Namespace",
      "source": "controllers/dsc_controller.go:145",
      "reason": "client.Get call for unwatched type"
    }
  ],
  "gomemlimit": "512MiB",
  "container_memory_limit": "1Gi",
  "default_transform": false,
  "findings": [
    {
      "severity": "warning",
      "message": "Missing DefaultTransform - managedFields consuming extra memory",
      "recommendation": "Add cache.DefaultTransform to strip managedFields"
    }
  ]
}
```

#### Operator Config

```json
[
  {
    "name": "DefaultDeploymentServiceAccount",
    "value": "ds-pipeline",
    "category": "name_pattern",
    "source": "controllers/dsp_params.go"
  },
  {
    "name": "APIServerImage",
    "value": "quay.io/opendatahub/ds-pipelines-api-server",
    "category": "image",
    "source": "controllers/config/defaults.go"
  }
]
```

#### Reconcile Sequences

```json
[
  {
    "controller": "DSPAReconciler",
    "steps": [
      {
        "method": "ReconcileDatabase",
        "component": "Database",
        "conditional": "p.DatabaseHealthy()",
        "source": "controllers/dspa_controller.go:85"
      },
      {
        "method": "ReconcileStorage",
        "component": "Storage",
        "source": "controllers/dspa_controller.go:92"
      }
    ],
    "source": "controllers/dspa_controller.go"
  }
]
```

#### Prometheus Metrics

```json
[
  {
    "name": "dspo_reconciliation_duration_seconds",
    "type": "histogram",
    "help": "Time taken to reconcile a DSPA resource",
    "labels": ["dspa_name", "dspa_namespace"],
    "namespace": "dspo",
    "source": "controllers/metrics.go"
  }
]
```

#### Status Conditions

```json
[
  {
    "type": "DatabaseAvailable",
    "reasons": ["DatabaseCreated", "DatabaseFailed", "ExternalDBInUse"],
    "source": "controllers/status.go"
  }
]
```

#### Platform Detection

```json
{
  "capabilities": [
    {
      "name": "IsOpenShift",
      "check": "whether the cluster is OpenShift",
      "source": "pkg/config/platform.go"
    }
  ],
  "conditionals": [
    {
      "condition": "p.IsOpenShift",
      "resource_kind": "Route",
      "action": "create",
      "source": "controllers/reconciler.go"
    }
  ]
}
```

#### Ingress Routing

```json
[
  {
    "kind": "Route",
    "name": "my-route",
    "hosts": ["app.example.com"],
    "paths": ["/api"],
    "backend": "my-service",
    "tls": false,
    "source": "internal/controller/config/templates/http-route.yaml.tmpl"
  },
  {
    "kind": "Ingress",
    "name": "my-ingress",
    "hosts": ["app.example.com"],
    "tls": true,
    "source": "config/networking/ingress.yaml"
  }
]
```

Supported kinds: `Gateway`, `HTTPRoute`, `Ingress`, `VirtualService`, `DestinationRule`, `ServiceEntry`, `Route` (OpenShift). Extracted from YAML manifests and `.yaml.tmpl` template files. RBAC-inferred entries include `rbac_verbs` and `note` fields.

#### Template Files

```json
[
  {
    "path": "internal/controller/config/templates/deployment.yaml.tmpl",
    "resource_kinds": ["Deployment"],
    "conditionals": [".Spec.Postgres", ".Spec.MySQL"]
  }
]
```

Go template files (`.yaml.tmpl`) found in the repository. Each entry includes the Kubernetes resource kinds defined in the template and any `{{if}}` conditional guards.

#### Security Annotations

```json
[
  {
    "type": "RBAC_CLUSTER_SCOPE_SENSITIVE",
    "severity": "high",
    "resource": "secrets",
    "verbs": ["create", "update", "patch", "delete"],
    "source": "config/rbac/role.yaml",
    "description": "ClusterRole \"manager-role\" grants cluster-wide create/update/patch/delete secrets"
  },
  {
    "type": "ROUTE_NO_TLS",
    "severity": "medium",
    "resource": "http-route",
    "source": "internal/controller/config/templates/http-route.yaml.tmpl",
    "description": "Route \"http-route\" has no TLS configuration."
  }
]
```

Security evaluation results from `security_eval.go`. These are also surfaced as `SEC-*` findings in `security-findings.json` when running `full-analysis`.

| Type | Rule prefix | What it detects |
|------|-------------|-----------------|
| `RBAC_CLUSTER_SCOPE_SENSITIVE` | `SEC-RBAC` | ClusterRoles granting cluster-wide CRUD on secrets, CRBs, SCCs, nodes, pods/exec |
| `SECRET_IN_CONTAINER_ARGS` | `SEC-SECRET` | Container args/command referencing secrets via `$(VAR)` |
| `CRD_CONFUSED_DEPUTY` | `SEC-CRD` | CRDs with user-settable image fields deployed with operator ServiceAccount |
| `MISSING_AUTH_REQUIREMENT` | `SEC-AUTH` | Optional auth components with mutual exclusion but no "at least one" rule |
| `ROUTE_NO_TLS` | `SEC-ROUTE` | OpenShift Routes with no `spec.tls` block |
| `GHA_UNPINNED_ACTION` | `SEC-GHA` | GitHub Actions using tag references instead of SHA pins |
| `GHA_MISSING_PERMISSIONS` | `SEC-GHA` | Workflows without explicit permissions blocks |

## component-report.md

Generated as part of the `analyze` and `full-analysis` diagram output (in the `diagrams/` directory). A human-readable markdown summary of all extracted architecture data.

Sections: APIs Exposed (CRDs, webhooks, HTTP endpoints), Dependencies (Go modules, internal ODH deps), Network Architecture (services, network policies, ingress), RBAC Surface (cluster roles, role bindings), Deployments (containers, security context, resources, probes), and Operational (metrics, status conditions, feature gates).

## build-config.json

Build metadata extracted from Dockerfiles, OLM bundle, and CI configuration.

```json
{
  "ocp_versions": {"min": "4.14", "max": "4.17"},
  "architectures": ["amd64", "arm64"],
  "go_version": "1.22",
  "base_images": ["registry.access.redhat.com/ubi9/go-toolset:1.22"],
  "fips_enabled": false,
  "olm": {"default_channel": "stable", "channels": ["stable", "alpha"]}
}
```

## quick-index.json

Lightweight function and call graph index from tree-sitter parsing. Produced by `quick-index` command.

```json
{
  "schema_version": 1,
  "repo_path": "/path/to/repo",
  "indexed_at": "2026-07-14T12:00:00Z",
  "stats": {
    "functions": 282,
    "call_sites": 4660,
    "call_edges": 1737,
    "http_endpoints": 0,
    "db_operations": 0,
    "classes": 0,
    "parse_time_ms": 70
  },
  "functions": [
    {
      "name": "Reconcile",
      "file": "controllers/reconciler.go",
      "line": 45,
      "end_line": 120,
      "language": "go",
      "kind": "method",
      "receiver": "MyReconciler",
      "params": ["ctx context.Context", "req ctrl.Request"],
      "return_type": "ctrl.Result, error",
      "complexity": 12
    }
  ],
  "calls": [
    {
      "caller": "r.Reconcile",
      "caller_file": "controllers/reconciler.go",
      "callee": "Create",
      "callee_file": "pkg/client/client.go",
      "line": 67,
      "confidence": "certain"
    }
  ],
  "http_endpoints": [
    {
      "name": "handlePredict",
      "route": "/v1/models/{model}:predict",
      "method": "POST",
      "file": "pkg/handler/predict.go",
      "line": 30
    }
  ]
}
```

## SrcLang context bundle (.srclg)

Structured XML format for LLM agent consumption. Produced by `context-bundle` command. Contains a domain-specialized view of the repository with semantic annotations.

```xml
<?xml version="1.0" encoding="UTF-8"?>
<srclang version="0.0.1" xmlns="https://srclang.dev/ns/core/0">
  <head>
    <producer>arch-analyzer 0.2.2</producer>
    <repository uri="https://github.com/org/my-operator">
      <commit sha="abc1234"/>
    </repository>
    <component>my-operator</component>
    <extracted>2026-07-14T10:00:00Z</extracted>
    <layer name="security"/>
    <languages>
      <language name="go"/>
    </languages>
    <platform name="RHOAI" components="38">
      <inbound>
        <edge from="other-component" type="go-module"/>
      </inbound>
      <outbound>
        <edge to="kserve" type="code-ref"/>
      </outbound>
    </platform>
  </head>
  <body>
    <layer name="security">
      <finding id="SEC-RBAC-001" domain="extraction" severity="high" rule="RBAC_CLUSTER_SCOPE_SENSITIVE">
        <source file="config/rbac/role.yaml"/>
        <title>ClusterRole grants cluster-wide secrets CRUD</title>
        <description>...</description>
      </finding>
      <file path="pkg/webhook/handler.go" language="go">
        <function name="Handle" kind="method" trust="untrusted" taint-role="source">
          <source line="42"/>
          <code><![CDATA[func (h *Handler) Handle(...) { ... }]]></code>
        </function>
      </file>
      <resource kind="ClusterRole" name="manager-role">
        <source file="config/rbac/role.yaml"/>
      </resource>
      <relationship kind="calls">
        <from function="Handle" file="pkg/webhook/handler.go" line="42"/>
        <to function="Create" file="pkg/client/client.go" resolved="false"/>
      </relationship>
    </layer>
  </body>
</srclang>
```

Available layers: `security` (security-relevant functions, taint paths, RBAC, network policies, findings) and `architecture` (CRDs, controller watches, reconcile sequences, external connections, API surface).

## code-graph.json

The code property graph output. Contains all nodes, edges, basic blocks, and optionally taint findings.

### Top-level structure

```json
{
  "schema_version": 2,
  "nodes": [],
  "edges": [],
  "taint_findings": []
}
```

### Node

```json
{
  "id": "pkg/handler/auth.go::HandleLogin",
  "kind": "Function",
  "name": "HandleLogin",
  "file": "pkg/handler/auth.go",
  "line": 42,
  "end_line": 85,
  "language": "go",
  "type_name": "",
  "complexity": 8,
  "param_names": ["w", "r"],
  "param_types": ["http.ResponseWriter", "*http.Request"],
  "return_type": "",
  "trust_level": "untrusted",
  "is_test": false,
  "annotations": {
    "handles_user_input": true,
    "sec:handles_request": true
  }
}
```

Node kinds: `File`, `Function`, `Parameter`, `Call`, `StructLiteral`, `Variable`, `BasicBlock`.

Trust levels: `untrusted` (public HTTP, no auth), `semi_trusted` (webhook, auth middleware), `trusted` (controller Reconcile, init).

### Edge

```json
{
  "source": "pkg/handler/auth.go::HandleLogin",
  "target": "pkg/db/users.go::FindUser",
  "kind": "EdgeCalls",
  "label": "",
  "confidence": "CERTAIN"
}
```

Edge kinds and labels:

- `EdgeCalls`: function-to-function call (confidence: CERTAIN, INFERRED, UNCERTAIN)
- `EdgeContains`: file-to-function, function-to-literal containment
- `EdgeAliases`: type alias relationship
- `EdgeDataFlow`: intraprocedural data flow (labels: `assigns`, `reads`, `passes_to`, `field_access`, `returns`)
- `EdgeControlFlow`: CFG edges (labels: `true_branch`, `false_branch`, `fallthrough`, `loop_back`, `loop_exit`, `exception`, `entry`, `exit`)

### Taint Finding

```json
{
  "rule": "taint-to-sink",
  "source": {
    "id": "pkg/handler/auth.go::HandleLogin::r",
    "file": "pkg/handler/auth.go",
    "line": 42
  },
  "sink": {
    "id": "pkg/db/users.go::FindUser::db.Query",
    "file": "pkg/db/users.go",
    "line": 67
  },
  "path": ["HandleLogin::r", "HandleLogin::username", "FindUser::query", "FindUser::db.Query"],
  "sanitized": false,
  "cross_function": true
}
```

## security-findings.json

Array of findings from CPG domain queries and security evaluation annotations. Produced by `full-analysis` and `scan` commands.

```json
[
  {
    "rule_id": "CGA-S01",
    "severity": "high",
    "message": "Webhook handler accepts untrusted input without validation",
    "file": "pkg/webhook/handler.go",
    "line": 42,
    "domain": "security",
    "architecture_ref": ""
  },
  {
    "rule_id": "SEC-RBAC-001",
    "severity": "high",
    "message": "ClusterRole \"manager-role\" grants cluster-wide create/update/patch/delete secrets",
    "file": "config/rbac/role.yaml",
    "domain": "security",
    "architecture_ref": "security_annotations:RBAC_CLUSTER_SCOPE_SENSITIVE"
  }
]
```

Finding sources:

- **CGA-* rules**: CPG domain query results (security, testing, upgrade, architecture, netpolicy domains)
- **SEC-* rules**: Security evaluation annotations converted to finding format (see [Security Annotations](#security-annotations))
- **External SARIF**: Findings ingested via `--import-sarif` flag (tool name and version preserved)
```

## Security findings (SARIF)

Standard SARIF 2.1.0 format compatible with GitHub Code Scanning:

```json
{
  "$schema": "https://raw.githubusercontent.com/oasis-tcs/sarif-spec/master/Schemata/sarif-schema-2.1.0.json",
  "version": "2.1.0",
  "runs": [
    {
      "tool": {
        "driver": {
          "name": "architecture-analyzer",
          "version": "0.2.0",
          "rules": [...]
        }
      },
      "results": [...]
    }
  ]
}
```

## Platform aggregation output

```json
{
  "components": ["repo-a", "repo-b", "repo-c"],
  "aggregated_at": "2026-04-14T10:30:00Z",
  "crd_ownership": {},
  "cross_dependencies": [],
  "rbac_overlap": [],
  "network_mesh": []
}
```
