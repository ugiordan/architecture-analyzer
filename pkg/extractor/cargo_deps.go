package extractor

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var (
	// cargoSectionRE matches a TOML table header like [dependencies] or
	// [dev-dependencies.foo] (target/platform-specific).
	cargoSectionRE = regexp.MustCompile(`^\[([^\]]+)\]`)
	// cargoEntryRE matches `name = "version"` or `name = { ... }` lines.
	// Keys may contain dots for TOML dotted-key syntax (serde.workspace = true).
	cargoEntryRE = regexp.MustCompile(`^([A-Za-z0-9_.-]+)\s*=\s*(.+)$`)
	// cargoVersionStringRE matches a bare version requirement string value.
	cargoVersionStringRE = regexp.MustCompile(`^"([^"]*)"`)
	// cargoTableFieldRE extracts a `field = "value"` pair inside a table value.
	cargoTableFieldRE = regexp.MustCompile(`(\w+)\s*=\s*"([^"]*)"`)
	// cargoBoolFieldRE extracts a `field = true|false` pair inside a table value.
	cargoBoolFieldRE = regexp.MustCompile(`(\w+)\s*=\s*(true|false)`)
	// cargoRustVersionRE matches rust-version under [package] or [workspace.package].
	cargoRustVersionRE = regexp.MustCompile(`^rust-version\s*=\s*"([^"]*)"`)
	// cargoWorkspaceExcludeRE matches the start of `exclude = [...]` under
	// [workspace], whether or not the array closes on the same line.
	cargoWorkspaceExcludeRE = regexp.MustCompile(`^exclude\s*=\s*(\[.*)?$`)
	// cargoQuotedItemRE extracts a quoted string from an array literal.
	cargoQuotedItemRE = regexp.MustCompile(`"([^"]*)"`)
)

// cargoDepListSections maps Cargo.toml dependency-list table names to the
// crate Category they produce. Order matters: more specific sections first so
// that "workspace.dependencies" is not matched by the plain "dependencies"
// prefix rule below.
var cargoDepListSections = []struct {
	base     string
	category string
}{
	{"workspace.dependencies", "workspace"},
	{"dev-dependencies", "dev"},
	{"build-dependencies", "build"},
	{"dependencies", ""},
}

// cargoSectionInfo classifies a TOML section header. It returns the crate
// category, the dependency name for dependency subtables (e.g.
// [dependencies.awesome] describes one dependency named "awesome"), and
// whether the section holds dependencies at all. Unrelated tables such as
// [package.metadata.dependencies] are rejected.
func cargoSectionInfo(section string) (category, depName string, isDep bool) {
	s := section
	if strings.HasPrefix(s, "target.") {
		rest := s[len("target."):]
		if strings.HasPrefix(rest, "'") {
			// Quoted cfg expression: strip target.'cfg(...)'.
			if idx := strings.Index(rest, "'."); idx >= 0 && idx+2 < len(rest) {
				s = rest[idx+2:]
			} else {
				return "", "", false
			}
		} else {
			// Bare target triple: strip target.x86_64-unknown-linux-gnu.
			if idx := strings.Index(rest, "."); idx >= 0 && idx+1 < len(rest) {
				s = rest[idx+1:]
			} else {
				return "", "", false
			}
		}
	}
	for _, base := range cargoDepListSections {
		if s == base.base {
			return base.category, "", true
		}
		if strings.HasPrefix(s, base.base+".") {
			return base.category, s[len(base.base)+1:], true
		}
	}
	return "", "", false
}

// discoverCargoManifests finds all Cargo.toml manifests in the repository,
// honoring workspace membership: manifests under a [workspace] root's
// exclude entries are skipped so excluded packages cannot contribute (or
// override) dependencies.
func discoverCargoManifests(repoPath string) []string {
	manifests := findFiles(repoPath, []string{"Cargo.toml", "**/Cargo.toml"})

	type wsExcludes struct {
		dir     string
		entries []string
	}
	var excluded []wsExcludes
	for _, m := range manifests {
		entries, ok := parseWorkspaceExclude(m)
		if ok && len(entries) > 0 {
			excluded = append(excluded, wsExcludes{dir: filepath.Dir(m), entries: entries})
		}
	}

	var result []string
	for _, m := range manifests {
		dir := filepath.Dir(m)
		skip := false
		for _, ex := range excluded {
			if dir != ex.dir && !strings.HasPrefix(dir, ex.dir+string(filepath.Separator)) {
				continue
			}
			rel, err := filepath.Rel(ex.dir, dir)
			if err != nil {
				continue
			}
			for _, entry := range ex.entries {
				if rel == entry || strings.HasPrefix(rel, entry+string(filepath.Separator)) {
					skip = true
					break
				}
			}
			if skip {
				break
			}
		}
		if !skip {
			result = append(result, m)
		}
	}
	return result
}

// parseWorkspaceExclude reads a manifest and returns the [workspace] exclude
// entries. The bool reports whether the manifest declares a [workspace] table
// at all. Supports single-line and multi-line arrays.
func parseWorkspaceExclude(fpath string) ([]string, bool) {
	data, err := os.ReadFile(fpath)
	if err != nil {
		return nil, false
	}

	section := ""
	isWorkspace := false
	var entries []string
	accumulating := false
	var arrayBuf strings.Builder

	flush := func() {
		if accumulating {
			for _, m := range cargoQuotedItemRE.FindAllStringSubmatch(arrayBuf.String(), -1) {
				entries = append(entries, m[1])
			}
			arrayBuf.Reset()
			accumulating = false
		}
	}

	for _, rawLine := range strings.Split(string(data), "\n") {
		line := strings.TrimSpace(rawLine)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if m := cargoSectionRE.FindStringSubmatch(line); m != nil {
			flush()
			section = m[1]
			isWorkspace = isWorkspace || section == "workspace"
			continue
		}
		if section != "workspace" {
			continue
		}
		if accumulating {
			arrayBuf.WriteString(" ")
			arrayBuf.WriteString(line)
			if strings.Contains(line, "]") {
				flush()
			}
			continue
		}
		if m := cargoWorkspaceExcludeRE.FindStringSubmatch(line); m != nil {
			arrayBuf.WriteString(m[1])
			accumulating = true
			if strings.Contains(m[1], "]") {
				flush()
			}
		}
	}
	flush()

	return entries, isWorkspace
}

// extractCargoDeps scans a repository for Cargo.toml manifests and extracts
// crate dependencies with version requirements. Dependencies declared with
// `workspace = true` in a member manifest are resolved against the
// workspace-level [workspace.dependencies] entries.
func extractCargoDeps(repoPath string) []CargoCrate {
	manifests := discoverCargoManifests(repoPath)

	parsed := make([][]CargoCrate, len(manifests))
	wsDeps := make(map[string]CargoCrate)
	for i, fpath := range manifests {
		deps, _ := parseCargoToml(fpath, repoPath)
		parsed[i] = deps
		for _, dep := range deps {
			if dep.Category == "workspace" {
				// Index by declared key so renamed entries stay findable by
				// members using the local alias.
				wsDeps[cargoCrateKey(dep)] = dep
			}
		}
	}

	seen := make(map[string]CargoCrate)
	for _, deps := range parsed {
		for _, dep := range deps {
			if dep.workspaceInherit {
				if ws, ok := wsDeps[cargoCrateKey(dep)]; ok {
					if dep.Version == "" {
						dep.Version = ws.Version
					}
					if dep.Kind == "" {
						dep.Kind = ws.Kind
					}
					if dep.Ref == "" {
						dep.Ref = ws.Ref
					}
					// Adopt the resolved display name so the member entry
					// merges with the workspace entry instead of duplicating.
					dep.Name = ws.Name
				}
			}
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

// extractRustVersion returns the rust-version declared in the root manifest
// ([package] rust-version, falling back to an inheritable
// [workspace.package] rust-version), or "" when absent. Only the literal
// root manifest is read: under ExtractAll's active FileIndex, glob patterns
// can match nested manifests (e.g. Cargo/Cargo.toml) whose versions must
// not win over the root's.
func extractRustVersion(repoPath string) string {
	rootManifest := filepath.Join(repoPath, "Cargo.toml")
	info, err := os.Lstat(rootManifest)
	if err != nil || info.IsDir() || info.Mode()&os.ModeSymlink != 0 {
		return ""
	}
	_, rustVersion := parseCargoToml(rootManifest, repoPath)
	return rustVersion
}

// parseCargoToml parses one Cargo.toml manifest, returning its crate
// dependencies and the manifest's rust-version (when present). Dependency
// subtables ([dependencies.awesome]) are parsed as a single named dependency.
func parseCargoToml(fpath, repoPath string) ([]CargoCrate, string) {
	data, err := os.ReadFile(fpath)
	if err != nil {
		log.Printf("warning: skipping %s: %v", fpath, err)
		return nil, ""
	}

	source := relativePath(repoPath, fpath)
	section := ""
	var crates []CargoCrate
	packageRV, workspaceRV := "", ""

	// Dependency subtable state: [dependencies.awesome] fields accumulate
	// until the next section header.
	var subName, subCategory string
	var subFields map[string]string

	// Dotted-key state: `serde.workspace = true` lines accumulate fields
	// per declared key until the next section header.
	type dottedEntry struct {
		category string
		fields   map[string]string
	}
	var dotted map[string]dottedEntry

	flushSubtable := func() {
		if subName == "" {
			return
		}
		crate := CargoCrate{Name: subName, key: subName, Category: subCategory, Source: source}
		applyCrateFields(&crate, subFields)
		crates = append(crates, crate)
		subName, subCategory, subFields = "", "", nil
	}
	flushDotted := func() {
		for name, entry := range dotted {
			crate := CargoCrate{Name: name, key: name, Category: entry.category, Source: source}
			applyCrateFields(&crate, entry.fields)
			crates = append(crates, crate)
		}
		dotted = nil
	}

	for _, rawLine := range strings.Split(string(data), "\n") {
		line := strings.TrimSpace(rawLine)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if m := cargoSectionRE.FindStringSubmatch(line); m != nil {
			flushSubtable()
			flushDotted()
			section = m[1]
			continue
		}

		if section == "package" {
			if m := cargoRustVersionRE.FindStringSubmatch(line); m != nil {
				packageRV = m[1]
			}
			continue
		}
		if section == "workspace.package" {
			if m := cargoRustVersionRE.FindStringSubmatch(line); m != nil {
				workspaceRV = m[1]
			}
			continue
		}

		category, depName, isDep := cargoSectionInfo(section)
		if !isDep {
			continue
		}

		m := cargoEntryRE.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		key, value := m[1], strings.TrimSpace(m[2])

		if depName != "" {
			// Inside a dependency subtable: every key is a field of the
			// single dependency named by the section header.
			if subFields == nil {
				subFields = make(map[string]string)
			}
			if qm := cargoVersionStringRE.FindStringSubmatch(value); qm != nil {
				subFields[key] = qm[1]
			} else if value == "true" || value == "false" {
				subFields[key] = value
			}
			if subName == "" {
				subName, subCategory = depName, category
			}
			continue
		}

		// Dotted-key syntax: `serde.workspace = true` declares a field of
		// the dependency named by the first path segment.
		if dotIdx := strings.Index(key, "."); dotIdx > 0 && dotIdx < len(key)-1 {
			base, field := key[:dotIdx], key[dotIdx+1:]
			var val string
			if qm := cargoVersionStringRE.FindStringSubmatch(value); qm != nil {
				val = qm[1]
			} else if value == "true" || value == "false" {
				val = value
			} else {
				continue
			}
			if dotted == nil {
				dotted = make(map[string]dottedEntry)
			}
			entry := dotted[base]
			if entry.fields == nil {
				entry.category = category
				entry.fields = make(map[string]string)
			}
			entry.fields[field] = val
			dotted[base] = entry
			continue
		}

		crate := CargoCrate{Name: key, key: key, Category: category, Source: source}
		if vm := cargoVersionStringRE.FindStringSubmatch(value); vm != nil {
			crate.Version = vm[1]
		} else if strings.HasPrefix(value, "{") {
			applyCrateFields(&crate, parseCargoTableFields(value))
		} else {
			continue
		}
		crates = append(crates, crate)
	}
	flushSubtable()
	flushDotted()

	rustVersion := packageRV
	if rustVersion == "" {
		rustVersion = workspaceRV
	}
	return crates, rustVersion
}

// parseCargoTableFields extracts `field = "value"` and `field = true|false`
// pairs from an inline table value like { version = "1.0", git = "..." }.
func parseCargoTableFields(value string) map[string]string {
	fields := make(map[string]string)
	for _, fm := range cargoTableFieldRE.FindAllStringSubmatch(value, -1) {
		fields[fm[1]] = fm[2]
	}
	for _, fm := range cargoBoolFieldRE.FindAllStringSubmatch(value, -1) {
		fields[fm[1]] = fm[2]
	}
	return fields
}

// applyCrateFields fills a crate from parsed dependency fields. Source fields
// (git/path) are handled independently of the version requirement so that
// combinations like { version = "1.0", git = "...", branch = "next" } keep
// both the requirement and the source.
func applyCrateFields(crate *CargoCrate, fields map[string]string) {
	if fields["workspace"] == "true" {
		crate.workspaceInherit = true
	}
	if v := fields["version"]; v != "" {
		crate.Version = v
	}
	if g := fields["git"]; g != "" {
		crate.Kind = "git"
		crate.Ref = firstNonEmpty(fields["tag"], fields["branch"], fields["rev"])
	}
	if p := fields["path"]; p != "" {
		crate.Kind = "path"
		crate.Ref = p
	}
	if rename := fields["package"]; rename != "" && rename != crate.Name {
		crate.Name = crate.Name + " (" + rename + ")"
	}
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}

// mergeCrate keeps the first occurrence of a crate, preferring normal
// dependencies over dev/build ones for the same name. An entry without a
// version never replaces an existing versioned entry.
func mergeCrate(seen map[string]CargoCrate, crate CargoCrate) {
	existing, ok := seen[crate.Name]
	if !ok {
		seen[crate.Name] = crate
		return
	}
	if existing.Version != "" && crate.Version == "" {
		return
	}
	if existing.Category != "" && crate.Category == "" {
		seen[crate.Name] = crate
	}
}
