package extractor

import (
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

var (
	// cargoSectionRE matches a TOML table header like [dependencies] or
	// [dev-dependencies.foo] (target/platform-specific).
	cargoSectionRE = regexp.MustCompile(`^\[([^\]]+)\]`)
	// cargoEntryRE matches `name = "version"` or `name = { ... }` lines.
	cargoEntryRE = regexp.MustCompile(`^([A-Za-z0-9_-]+)\s*=\s*(.+)$`)
	// cargoVersionStringRE matches a bare version requirement string value.
	cargoVersionStringRE = regexp.MustCompile(`^"([^"]*)"`)
	// cargoTableFieldRE extracts a `field = "value"` pair inside a table value.
	cargoTableFieldRE = regexp.MustCompile(`(\w+)\s*=\s*"([^"]*)"`)
	// cargoRustVersionRE matches rust-version under [package].
	cargoRustVersionRE = regexp.MustCompile(`^rust-version\s*=\s*"([^"]*)"`)
)

// cargoDepSections maps TOML section names (and prefixes, for target-specific
// tables like dev-dependencies.bar) to the crate Category they produce.
var cargoDepSections = []struct {
	section  string
	category string
}{
	// Order matters: more specific sections first so that
	// "workspace.dependencies" is not matched by the plain "dependencies"
	// suffix rule below.
	{"workspace.dependencies", "workspace"},
	{"dev-dependencies", "dev"},
	{"build-dependencies", "build"},
	{"dependencies", ""},
}

// extractCargoDeps scans a repository for Cargo.toml manifests and extracts
// crate dependencies with version requirements.
func extractCargoDeps(repoPath string) []CargoCrate {
	manifestFiles := findFiles(repoPath, []string{"Cargo.toml", "**/Cargo.toml"})

	seen := make(map[string]CargoCrate)
	for _, fpath := range manifestFiles {
		deps, _ := parseCargoToml(fpath, repoPath)
		for _, dep := range deps {
			mergeCrate(seen, dep)
		}
	}

	result := make([]CargoCrate, 0, len(seen))
	for _, crate := range seen {
		result = append(result, crate)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result
}

// extractRustVersion returns the rust-version declared in the root manifest,
// or "" when absent.
func extractRustVersion(repoPath string) string {
	rootPath := findFiles(repoPath, []string{"Cargo.toml"})
	for _, fpath := range rootPath {
		if _, rustVersion := parseCargoToml(fpath, repoPath); rustVersion != "" {
			return rustVersion
		}
	}
	return ""
}

// parseCargoToml parses one Cargo.toml manifest, returning its crate
// dependencies and the manifest's rust-version (when present).
func parseCargoToml(fpath, repoPath string) ([]CargoCrate, string) {
	data, err := os.ReadFile(fpath)
	if err != nil {
		log.Printf("warning: skipping %s: %v", fpath, err)
		return nil, ""
	}

	source := relativePath(repoPath, fpath)
	section := ""
	var crates []CargoCrate
	rustVersion := ""

	for _, rawLine := range strings.Split(string(data), "\n") {
		line := strings.TrimSpace(rawLine)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if m := cargoSectionRE.FindStringSubmatch(line); m != nil {
			section = m[1]
			continue
		}

		if section == "package" {
			if m := cargoRustVersionRE.FindStringSubmatch(line); m != nil {
				rustVersion = m[1]
			}
			continue
		}

		category, ok := cargoCategoryForSection(section)
		if !ok {
			continue
		}

		m := cargoEntryRE.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		name, value := m[1], strings.TrimSpace(m[2])
		crate := CargoCrate{Name: name, Category: category, Source: source}

		if vm := cargoVersionStringRE.FindStringSubmatch(value); vm != nil {
			crate.Version = vm[1]
		} else if strings.HasPrefix(value, "{") {
			fields := make(map[string]string)
			for _, fm := range cargoTableFieldRE.FindAllStringSubmatch(value, -1) {
				fields[fm[1]] = fm[2]
			}
			switch {
			case fields["version"] != "":
				crate.Version = fields["version"]
			case fields["git"] != "":
				crate.Kind = "git"
				crate.Version = fields["tag"]
				if crate.Version == "" {
					crate.Version = fields["branch"]
				}
				if crate.Version == "" {
					crate.Version = fields["rev"]
				}
				if crate.Version == "" {
					crate.Version = "git"
				}
			case fields["path"] != "":
				crate.Kind = "path"
				crate.Version = fields["path"]
			}
			if rename := fields["package"]; rename != "" {
				crate.Name = name + " (" + rename + ")"
			}
		} else {
			continue
		}

		crates = append(crates, crate)
	}

	return crates, rustVersion
}

// cargoCategoryForSection maps a TOML section name (possibly target-specific,
// e.g. "dev-dependencies.foo" or "target.'cfg(unix)'.dependencies") to a crate
// category. Returns false for non-dependency sections.
func cargoCategoryForSection(section string) (string, bool) {
	for _, s := range cargoDepSections {
		if section == s.section || strings.HasPrefix(section, s.section+".") ||
			strings.HasSuffix(section, "."+s.section) {
			return s.category, true
		}
	}
	return "", false
}

// mergeCrate keeps the first occurrence of a crate, preferring normal
// dependencies over dev/build ones for the same name.
func mergeCrate(seen map[string]CargoCrate, crate CargoCrate) {
	existing, ok := seen[crate.Name]
	if !ok {
		seen[crate.Name] = crate
		return
	}
	if existing.Category != "" && crate.Category == "" {
		seen[crate.Name] = crate
	}
}
