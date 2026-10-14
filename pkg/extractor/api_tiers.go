package extractor

import (
	"github.com/ugiordan/architecture-analyzer/pkg/config"
)

// ApplyAPITiers enriches extracted CRDs with tier assignments from an
// APITiersConfig. Each CRD is matched by group/kind/version, with wildcard
// support. CRDs that already have a tier set are not overwritten.
func ApplyAPITiers(arch *ComponentArchitecture, tierCfg *config.APITiersConfig) {
	if tierCfg == nil {
		return
	}
	for i := range arch.CRDs {
		if arch.CRDs[i].APITier != "" {
			continue
		}
		tier, ok := tierCfg.MatchCRDTier(arch.CRDs[i].Group, arch.CRDs[i].Kind, arch.CRDs[i].Version)
		if ok {
			arch.CRDs[i].APITier = tier
		}
	}
}
