package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// APITiersConfig represents the api_tiers.yaml file structure.
type APITiersConfig struct {
	Version  string          `yaml:"version" json:"version"`
	CRDs     []CRDTierEntry  `yaml:"crds" json:"crds"`
	RESTAPIs []APITierEntry  `yaml:"rest_apis" json:"rest_apis"`
	SDKs     []APITierEntry  `yaml:"sdks" json:"sdks"`
}

// CRDTierEntry maps a CRD (by group/kind/version) to a tier.
// Kind and Version support "*" as a wildcard.
type CRDTierEntry struct {
	Group   string `yaml:"group" json:"group"`
	Kind    string `yaml:"kind" json:"kind"`
	Version string `yaml:"version" json:"version"`
	Tier    string `yaml:"tier" json:"tier"`
}

// APITierEntry maps a REST API or SDK to a tier.
type APITierEntry struct {
	Name      string `yaml:"name" json:"name"`
	Component string `yaml:"component" json:"component"`
	Tier      string `yaml:"tier" json:"tier"`
}

// LoadAPITiersConfig reads and parses an api_tiers.yaml file.
func LoadAPITiersConfig(path string) (*APITiersConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading api tiers config: %w", err)
	}

	var cfg APITiersConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing api tiers config: %w", err)
	}

	return &cfg, nil
}

// MatchCRDTier finds the tier for a CRD by group, kind, and version.
// Returns the tier string and true if a match is found.
// Exact matches take priority over wildcard matches.
func (c *APITiersConfig) MatchCRDTier(group, kind, version string) (string, bool) {
	var wildcardMatch string
	for _, entry := range c.CRDs {
		if entry.Group != group {
			continue
		}
		kindMatch := entry.Kind == kind || entry.Kind == "*"
		versionMatch := entry.Version == version || entry.Version == "*"
		if !kindMatch || !versionMatch {
			continue
		}
		if entry.Kind == kind && entry.Version == version {
			return entry.Tier, true
		}
		if wildcardMatch == "" {
			wildcardMatch = entry.Tier
		}
	}
	if wildcardMatch != "" {
		return wildcardMatch, true
	}
	return "", false
}
