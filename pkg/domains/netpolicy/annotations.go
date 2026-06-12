// Package netpolicy provides NetworkPolicy trust chain analysis for code property graphs.
// It traces namespaceSelector labels to their application sites, identifies tenant
// workload namespaces, and flags trust boundary violations where tenant code can
// reach control plane services without podSelector or port restrictions.
package netpolicy

const (
	AnnotNetPolSelector    = "netpol:namespace_selector"
	AnnotNetPolTenantReach = "netpol:tenant_reach"
	AnnotNetPolNoRestrict  = "netpol:no_restriction"
)
