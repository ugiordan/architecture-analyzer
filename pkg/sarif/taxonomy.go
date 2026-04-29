package sarif

// NormalizedSeverity maps tool-specific severity levels to a unified taxonomy.
// SARIF defines: error, warning, note, none. Tools also use custom levels.
// We normalize to: critical, high, medium, low, informational.
func NormalizedSeverity(level string) string {
	switch level {
	case "error":
		return "high"
	case "warning":
		return "medium"
	case "note":
		return "low"
	case "none", "":
		return "informational"
	// Pass through already-normalized values
	case "critical", "high", "medium", "low", "informational":
		return level
	default:
		return "medium"
	}
}

// CWEToCategory maps CWE identifiers to broad security categories for
// cross-referencing with domain analyzers.
var CWEToCategory = map[string]string{
	// Injection
	"CWE-78":  "injection",
	"CWE-79":  "injection",
	"CWE-89":  "injection",
	"CWE-94":  "injection",
	"CWE-917": "injection",

	// Authentication / Authorization
	"CWE-287": "auth",
	"CWE-306": "auth",
	"CWE-862": "auth",
	"CWE-863": "auth",

	// Cryptography
	"CWE-326": "crypto",
	"CWE-327": "crypto",
	"CWE-328": "crypto",
	"CWE-330": "crypto",
	"CWE-338": "crypto",

	// Data Exposure
	"CWE-200": "data-exposure",
	"CWE-209": "data-exposure",
	"CWE-532": "data-exposure",
	"CWE-312": "data-exposure",
	"CWE-319": "data-exposure",

	// Path Traversal / File
	"CWE-22":  "path-traversal",
	"CWE-23":  "path-traversal",
	"CWE-73":  "path-traversal",
	"CWE-434": "path-traversal",

	// SSRF / Request Forgery
	"CWE-918": "ssrf",
	"CWE-352": "csrf",

	// Deserialization
	"CWE-502": "deserialization",

	// Hardcoded Credentials
	"CWE-798": "hardcoded-credentials",
	"CWE-259": "hardcoded-credentials",

	// K8s / Container specific
	"CWE-250": "privilege-escalation",
	"CWE-269": "privilege-escalation",
}

// categoryPriority defines preference order when multiple CWE categories match.
// Higher value = higher priority. More specific categories are preferred over broad ones.
var categoryPriority = map[string]int{
	"injection":             1,
	"auth":                  2,
	"crypto":                3,
	"data-exposure":         4,
	"path-traversal":        5,
	"ssrf":                  6,
	"csrf":                  6,
	"deserialization":       7,
	"hardcoded-credentials": 8,
	"privilege-escalation":  9,
}

// CategoryForCWEs returns the highest-priority security category for a set of CWEs.
// When multiple CWEs map to different categories, the most specific category wins
// (determined by categoryPriority). Returns empty string if no CWEs match.
func CategoryForCWEs(cwes []string) string {
	bestCat := ""
	bestPri := -1
	for _, cwe := range cwes {
		if cat, ok := CWEToCategory[cwe]; ok {
			pri := categoryPriority[cat]
			if pri > bestPri {
				bestPri = pri
				bestCat = cat
			}
		}
	}
	return bestCat
}
