package security

import (
	"strings"

	"github.com/ugiordan/architecture-analyzer/pkg/domains"
	"github.com/ugiordan/architecture-analyzer/pkg/graph"
)

type GoAnnotator struct{}

func (a *GoAnnotator) Annotate(g *graph.CPG, archData *domains.ArchitectureData) error {
	// First pass: annotate individual nodes
	for _, cs := range g.NodesByKind(graph.NodeCallSite) {
		if cs.Language != "go" {
			continue
		}
		a.annotateCallSite(g, cs)
	}
	for _, sl := range g.NodesByKind(graph.NodeStructLiteral) {
		if sl.Language != "go" {
			continue
		}
		a.annotateStructLiteral(g, sl)
	}
	// Second pass: annotate functions based on contained nodes
	for _, fn := range g.NodesByKind(graph.NodeFunction) {
		if fn.Language != "go" {
			continue
		}
		a.annotateFunction(g, fn)
	}
	// Third pass: classify trust levels
	a.classifyTrust(g)
	// Cross-language annotations (auth, storage, external calls).
	// These were previously in a separate legacy SecurityAnnotator.
	a.annotateAuthAndStorage(g)
	a.annotateExternalCalls(g)
	return nil
}

func (a *GoAnnotator) classifyTrust(g *graph.CPG) {
	// Go HTTP endpoints default to untrusted
	for _, ep := range g.NodesByKind(graph.NodeHTTPEndpoint) {
		if ep.Language != "go" {
			continue
		}
		g.SetTrustLevel(ep.ID, graph.TrustUntrusted)
	}

	for _, fn := range g.NodesByKind(graph.NodeFunction) {
		if fn.Language != "go" {
			continue
		}
		paramTypes := strings.Join(fn.ParamTypes, ",")

		// Admission webhook handlers are semi-trusted (authenticated by API server)
		if strings.Contains(paramTypes, "admission.Request") || strings.Contains(paramTypes, "AdmissionReview") {
			g.SetTrustLevel(fn.ID, graph.TrustSemiTrusted)
			continue
		}

		// Controller Reconcile functions are trusted (internal loop)
		if fn.Name == "Reconcile" && (strings.Contains(paramTypes, "ctrl.Request") || strings.Contains(paramTypes, "reconcile.Request")) {
			g.SetTrustLevel(fn.ID, graph.TrustTrusted)
			continue
		}

		// Init/setup functions are trusted
		if fn.Name == "init" || fn.Name == "main" || fn.Name == "SetupWithManager" {
			g.SetTrustLevel(fn.ID, graph.TrustTrusted)
			continue
		}
	}
}

func (a *GoAnnotator) annotateFunction(g *graph.CPG, fn *graph.Node) {
	// sec:handles_admission: function with admission.Request parameter type
	paramTypes := strings.Join(fn.ParamTypes, ",")
	if strings.Contains(paramTypes, "admission.Request") {
		g.SetAnnotation(fn.ID, AnnotHandlesAdmission, true)
	}

	// Propagate annotations from contained call sites and struct literals
	for _, edge := range g.OutEdges(fn.ID) {
		if edge.Kind != graph.EdgeDataFlow {
			continue
		}
		target := g.GetNode(edge.To)
		if target == nil {
			continue
		}

		if target.Kind == graph.NodeCallSite {
			if target.Annotations[AnnotCreatesRBAC] {
				g.SetAnnotation(fn.ID, AnnotCreatesRBAC, true)
			}
			if target.Annotations[AnnotAccessesSecret] {
				g.SetAnnotation(fn.ID, AnnotAccessesSecret, true)
			}
			if target.Annotations[AnnotConfiguresCache] {
				g.SetAnnotation(fn.ID, AnnotConfiguresCache, true)
			}
			if target.Annotations[AnnotWritesPlaintextSecret] {
				g.SetAnnotation(fn.ID, AnnotWritesPlaintextSecret, true)
			}
		}
		if target.Kind == graph.NodeStructLiteral {
			if target.Annotations[AnnotGeneratesCert] {
				g.SetAnnotation(fn.ID, AnnotGeneratesCert, true)
			}
			if target.Annotations[AnnotConfiguresCache] {
				g.SetAnnotation(fn.ID, AnnotConfiguresCache, true)
			}
		}
	}

	// sec:binds_subject: check for system:authenticated strings in RBAC functions
	if fn.Annotations[AnnotCreatesRBAC] {
		a.checkBindsSubject(g, fn)
	}
}

func (a *GoAnnotator) annotateCallSite(g *graph.CPG, cs *graph.Node) {
	name := cs.Name
	argTypes := cs.Properties["arg_types"]
	stringArgs := cs.Properties["string_args"]

	// sec:creates_rbac
	if isClientMutation(name) && containsRBACType(argTypes) {
		g.SetAnnotation(cs.ID, AnnotCreatesRBAC, true)
	}

	// sec:accesses_secret
	if isClientAccess(name) && strings.Contains(argTypes, "Secret") {
		g.SetAnnotation(cs.ID, AnnotAccessesSecret, true)
	}

	// sec:configures_cache
	if isCacheConfig(name) {
		g.SetAnnotation(cs.ID, AnnotConfiguresCache, true)
	}

	// sec:writes_plaintext_secret
	if isFileWrite(name) && hasSecretArg(stringArgs, argTypes) {
		g.SetAnnotation(cs.ID, AnnotWritesPlaintextSecret, true)
	}
}

func (a *GoAnnotator) annotateStructLiteral(g *graph.CPG, sl *graph.Node) {
	typeName := sl.StructType
	fields := strings.Join(sl.FieldNames, ",")

	// sec:generates_cert
	if strings.Contains(typeName, "Certificate") {
		if strings.Contains(fields, "IsCA") || strings.Contains(fields, "KeyUsage") || strings.Contains(fields, "SerialNumber") {
			g.SetAnnotation(sl.ID, AnnotGeneratesCert, true)
		}
	}

	// sec:configures_cache
	if strings.Contains(typeName, "ByObject") {
		g.SetAnnotation(sl.ID, AnnotConfiguresCache, true)
	}
}

func (a *GoAnnotator) checkBindsSubject(g *graph.CPG, fn *graph.Node) {
	for _, edge := range g.OutEdges(fn.ID) {
		if edge.Kind != graph.EdgeDataFlow {
			continue
		}
		target := g.GetNode(edge.To)
		if target == nil {
			continue
		}
		if target.Kind == graph.NodeStructLiteral {
			sv := target.Properties["string_values"]
			if containsSubjectString(sv) {
				g.SetAnnotation(fn.ID, AnnotBindsSubject, true)
				return
			}
		}
		if target.Kind == graph.NodeCallSite {
			sa := target.Properties["string_args"]
			if containsSubjectString(sa) {
				g.SetAnnotation(fn.ID, AnnotBindsSubject, true)
				return
			}
		}
	}
}

func isClientMutation(name string) bool {
	for _, s := range []string{".Create", ".Update", ".Patch"} {
		if strings.HasSuffix(name, s) {
			return true
		}
	}
	return false
}

func isClientAccess(name string) bool {
	for _, s := range []string{".Get", ".List", ".Create", ".Update"} {
		if strings.HasSuffix(name, s) {
			return true
		}
	}
	return false
}

func containsRBACType(argTypes string) bool {
	for _, rt := range []string{"Role", "ClusterRole", "RoleBinding", "ClusterRoleBinding"} {
		if strings.Contains(argTypes, rt) {
			return true
		}
	}
	return false
}

func isCacheConfig(name string) bool {
	for _, p := range []string{"cache.New", "ctrl.NewManager", "NewCache"} {
		if strings.Contains(name, p) {
			return true
		}
	}
	return false
}

func isFileWrite(name string) bool {
	for _, p := range []string{"WriteFile", "ReplaceStringsInFile"} {
		if strings.HasSuffix(name, p) {
			return true
		}
	}
	return false
}

func hasSecretArg(stringArgs, argTypes string) bool {
	combined := strings.ToLower(stringArgs + " " + argTypes)
	// Use word-boundary-aware matching to avoid false positives
	// ("key" matching "keyboard", "monkey", "MapKeyValue", etc.)
	for _, w := range []string{"secret", "password", "apikey", "api_key", "token", "credential", "passw"} {
		if strings.Contains(combined, w) {
			return true
		}
	}
	// Check for standalone "key" only when preceded by common secret prefixes.
	for _, prefix := range []string{"_key", "-key", "secretkey", "privatekey", "sshkey", "authkey"} {
		if strings.Contains(combined, prefix) {
			return true
		}
	}
	return false
}

func containsSubjectString(s string) bool {
	for _, sub := range []string{"system:authenticated", "system:unauthenticated", "system:serviceaccount:"} {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}

// annotateAuthAndStorage detects auth decorators and DB storage operations.
// Runs on ALL languages (decorator/DB patterns are cross-language).
func (a *GoAnnotator) annotateAuthAndStorage(g *graph.CPG) {
	for _, fn := range g.NodesByKind(graph.NodeFunction) {
		for _, dec := range fn.Decorators {
			lower := strings.ToLower(dec)
			if strings.Contains(lower, "auth") || strings.Contains(lower, "login_required") ||
				strings.Contains(lower, "require_admin") || strings.Contains(lower, "authenticated") {
				g.SetAnnotation(fn.ID, "has_auth", true)
			}
			if strings.Contains(lower, "rate_limit") || strings.Contains(lower, "limiter") {
				g.SetAnnotation(fn.ID, "has_rate_limit", true)
			}
		}
		for _, edge := range g.OutEdges(fn.ID) {
			target := g.GetNode(edge.To)
			if target == nil {
				continue
			}
			if target.Kind == graph.NodeDBOperation {
				if target.Operation == "write" {
					g.SetAnnotation(fn.ID, "writes_storage", true)
					g.SetAnnotation(fn.ID, "mutates_state", true)
				} else if target.Operation == "read" {
					g.SetAnnotation(fn.ID, "reads_storage", true)
				}
			}
			if target.Kind == graph.NodeCallSite {
				for _, inner := range g.OutEdges(target.ID) {
					it := g.GetNode(inner.To)
					if it != nil && it.Kind == graph.NodeDBOperation {
						if it.Operation == "write" {
							g.SetAnnotation(fn.ID, "writes_storage", true)
							g.SetAnnotation(fn.ID, "mutates_state", true)
						} else if it.Operation == "read" {
							g.SetAnnotation(fn.ID, "reads_storage", true)
						}
					}
				}
			}
		}
	}
	for _, op := range g.NodesByKind(graph.NodeDBOperation) {
		if op.Operation == "write" {
			g.SetAnnotation(op.ID, "writes_storage", true)
		} else if op.Operation == "read" {
			g.SetAnnotation(op.ID, "reads_storage", true)
		}
	}
}

// annotateExternalCalls detects HTTP client calls and namespace-crossing operations.
func (a *GoAnnotator) annotateExternalCalls(g *graph.CPG) {
	for _, cs := range g.NodesByKind(graph.NodeCallSite) {
		name := strings.ToLower(cs.Name)
		isExternal := strings.HasPrefix(name, "http.") && (strings.Contains(name, "post") ||
			strings.Contains(name, "get") || strings.Contains(name, "do"))
		if strings.Contains(name, "client.do") || strings.Contains(name, "client.post") ||
			strings.Contains(name, "client.get") {
			isExternal = true
		}
		if isExternal {
			g.SetAnnotation(cs.ID, AnnotCallsExternal, true)
			for _, edge := range g.InEdges(cs.ID) {
				src := g.GetNode(edge.From)
				if src != nil && src.Kind == graph.NodeFunction {
					g.SetAnnotation(src.ID, AnnotCallsExternal, true)
				}
			}
		}
		if strings.Contains(name, "namespace") && (strings.Contains(name, "get") ||
			strings.Contains(name, "list") || strings.Contains(name, "client.")) {
			g.SetAnnotation(cs.ID, AnnotCrossesNamespace, true)
			for _, edge := range g.InEdges(cs.ID) {
				src := g.GetNode(edge.From)
				if src != nil && src.Kind == graph.NodeFunction {
					g.SetAnnotation(src.ID, AnnotCrossesNamespace, true)
				}
			}
		}
	}
}
