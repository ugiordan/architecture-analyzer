package renderer

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/ugiordan/architecture-analyzer/pkg/config"
)

// TierReportEntry represents a single row in the tier report, combining CRDs,
// REST APIs, and SDKs into a uniform format matching the KCS article layout.
type TierReportEntry struct {
	APIVersionExample string `json:"api_version_example"`
	APITier           string `json:"api_tier"`
	Component         string `json:"component,omitempty"`
	Category          string `json:"category"` // "CRD", "REST API", "SDK"
}

// TierReportResult holds the full tier report output including coverage gaps.
type TierReportResult struct {
	Version   string            `json:"version"`
	Entries   []TierReportEntry `json:"entries"`
	Gaps      TierReportGaps    `json:"gaps,omitempty"`
}

// TierReportGaps tracks CRDs missing tier assignments and stale mapping entries.
type TierReportGaps struct {
	MissingTier []string `json:"missing_tier,omitempty"`
	StaleEntries []string `json:"stale_entries,omitempty"`
}

// BuildTierReport generates a tier report by joining aggregated platform data
// with an API tiers config. It returns the structured result for rendering in
// any output format.
func BuildTierReport(platformData map[string]interface{}, tierCfg *config.APITiersConfig) *TierReportResult {
	result := &TierReportResult{
		Version: tierCfg.Version,
	}

	crds := getSlice(platformData, "crds")

	// Track which tier config CRD entries were matched for stale detection
	matched := make(map[int]bool)

	for _, crd := range crds {
		group := getStr(crd, "group", "")
		kind := getStr(crd, "kind", "")
		version := getStr(crd, "version", "")
		owner := getStr(crd, "owner", "")

		tier := getStr(crd, "api_tier", "")
		if tier == "" {
			tier, _ = tierCfg.MatchCRDTier(group, kind, version)
		}

		// Mark matching config entries
		for i, entry := range tierCfg.CRDs {
			if entry.Group != group {
				continue
			}
			kindMatch := entry.Kind == kind || entry.Kind == "*"
			versionMatch := entry.Version == version || entry.Version == "*"
			if kindMatch && versionMatch {
				matched[i] = true
			}
		}

		// Build plural.group/version format for the API version example
		plural := strings.ToLower(kind) + "s"
		apiVersion := fmt.Sprintf("%s.%s/%s", plural, group, version)

		entry := TierReportEntry{
			APIVersionExample: apiVersion,
			APITier:           tier,
			Component:         owner,
			Category:          "CRD",
		}
		result.Entries = append(result.Entries, entry)

		if tier == "" {
			result.Gaps.MissingTier = append(result.Gaps.MissingTier, apiVersion)
		}
	}

	// Detect stale entries (tier config CRDs with no extracted match)
	for i, entry := range tierCfg.CRDs {
		if entry.Kind == "*" {
			continue
		}
		if !matched[i] {
			stale := fmt.Sprintf("%s/%s %s", entry.Group, entry.Version, entry.Kind)
			result.Gaps.StaleEntries = append(result.Gaps.StaleEntries, stale)
		}
	}

	// Add REST APIs from tier config
	for _, api := range tierCfg.RESTAPIs {
		result.Entries = append(result.Entries, TierReportEntry{
			APIVersionExample: api.Name,
			APITier:           api.Tier,
			Component:         api.Component,
			Category:          "REST API",
		})
	}

	// Add SDKs from tier config
	for _, sdk := range tierCfg.SDKs {
		result.Entries = append(result.Entries, TierReportEntry{
			APIVersionExample: sdk.Name,
			APITier:           sdk.Tier,
			Component:         sdk.Component,
			Category:          "SDK",
		})
	}

	sort.Slice(result.Entries, func(i, j int) bool {
		if result.Entries[i].APITier != result.Entries[j].APITier {
			return result.Entries[i].APITier < result.Entries[j].APITier
		}
		return result.Entries[i].APIVersionExample < result.Entries[j].APIVersionExample
	})

	return result
}

// RenderTierReportMarkdown renders the tier report as a markdown table.
func RenderTierReportMarkdown(report *TierReportResult) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("# API Tier Report (%s)\n\n", report.Version))
	b.WriteString("| API Version / Name | API Tier | Category | Component |\n")
	b.WriteString("|--------------------|----------|----------|-----------|\n")

	for _, entry := range report.Entries {
		tier := entry.APITier
		if tier == "" {
			tier = "*unassigned*"
		}
		b.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n",
			escapeMdCell(entry.APIVersionExample),
			escapeMdCell(tier),
			escapeMdCell(entry.Category),
			escapeMdCell(entry.Component)))
	}
	b.WriteString("\n")

	if len(report.Gaps.MissingTier) > 0 {
		b.WriteString("## Coverage Gaps\n\n")
		b.WriteString("CRDs found by extraction but missing from api_tiers.yaml:\n\n")
		for _, gap := range report.Gaps.MissingTier {
			b.WriteString(fmt.Sprintf("- `%s`\n", gap))
		}
		b.WriteString("\n")
	}

	if len(report.Gaps.StaleEntries) > 0 {
		b.WriteString("## Stale Entries\n\n")
		b.WriteString("Entries in api_tiers.yaml that match no extracted CRD:\n\n")
		for _, stale := range report.Gaps.StaleEntries {
			b.WriteString(fmt.Sprintf("- `%s`\n", stale))
		}
		b.WriteString("\n")
	}

	return b.String()
}

// RenderTierReportCSV renders the tier report as CSV.
func RenderTierReportCSV(report *TierReportResult) string {
	var b strings.Builder

	b.WriteString("api_version_example,api_tier,category,component\n")
	for _, entry := range report.Entries {
		b.WriteString(fmt.Sprintf("%q,%q,%q,%q\n",
			entry.APIVersionExample, entry.APITier, entry.Category, entry.Component))
	}

	return b.String()
}

// RenderTierReportJSON renders the tier report as indented JSON.
func RenderTierReportJSON(report *TierReportResult) (string, error) {
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshaling tier report: %w", err)
	}
	return string(data) + "\n", nil
}
