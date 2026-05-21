package extractor

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// FileIndex pre-walks a repository once and provides fast file lookups
// by glob pattern, avoiding redundant filesystem walks across 30+ extractors.
//
// The index stores paths relative to the repo root AND absolute paths.
// findFiles returns paths in the same format as filepath.Join(root, pattern)
// (typically absolute when root is absolute), so the index must reproduce that.
type FileIndex struct {
	root     string
	absRoot  string
	allFiles []indexedFile
	byExt    map[string][]int // ".go" → indices into allFiles
	once     sync.Once
}

type indexedFile struct {
	abs  string // absolute path
	rel  string // relative to absRoot
	base string // filepath.Base
	ext  string // filepath.Ext
}

// NewFileIndex creates a file index for the given repo root.
// The actual walk is deferred until the first query.
func NewFileIndex(root string) *FileIndex {
	abs, _ := filepath.Abs(root)
	return &FileIndex{
		root:    root,
		absRoot: abs,
		byExt:   make(map[string][]int),
	}
}

func (fi *FileIndex) build() {
	fi.once.Do(func() {
		filepath.WalkDir(fi.absRoot, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			if d.IsDir() {
				name := d.Name()
				if strings.HasPrefix(name, ".") || DefaultExcludedDirs[name] {
					return fs.SkipDir
				}
				return nil
			}
			info, infoErr := d.Info()
			if infoErr != nil {
				return nil
			}
			if info.Mode()&os.ModeSymlink != 0 {
				return nil
			}
			rel, _ := filepath.Rel(fi.absRoot, path)
			base := filepath.Base(path)
			ext := filepath.Ext(base)
			idx := len(fi.allFiles)
			fi.allFiles = append(fi.allFiles, indexedFile{
				abs:  path,
				rel:  rel,
				base: base,
				ext:  ext,
			})
			if ext != "" {
				fi.byExt[ext] = append(fi.byExt[ext], idx)
			}
			return nil
		})
	})
}

// FindByPatterns returns files matching any of the given glob patterns.
// Drop-in replacement for the walk-based findFiles() that uses the pre-built index.
// Returns paths in the same format as findFiles (filepath.Join(root, pattern)).
func (fi *FileIndex) FindByPatterns(patterns []string) []string {
	fi.build()
	seen := make(map[string]bool)
	var result []string

	for _, pattern := range patterns {
		candidates := fi.candidatesForPattern(pattern)
		for _, idx := range candidates {
			f := fi.allFiles[idx]
			if fi.matchPattern(pattern, f.rel, f.base) {
				if !seen[f.abs] {
					seen[f.abs] = true
					result = append(result, f.abs)
				}
			}
		}
	}
	return result
}

// candidatesForPattern narrows the search space using the extension index.
// For patterns with a clear extension (e.g. "**/*.go", "config/crd/*.yaml"),
// only files with that extension are checked. Otherwise all files are checked.
func (fi *FileIndex) candidatesForPattern(pattern string) []int {
	ext := filepath.Ext(pattern)
	// Template extensions like ".tmpl" need the compound extension
	if ext == ".tmpl" {
		// For "*.yaml.tmpl", we still want ".tmpl" files
		if indices, ok := fi.byExt[ext]; ok {
			return indices
		}
	}
	if ext != "" && !strings.Contains(ext, "*") {
		if indices, ok := fi.byExt[ext]; ok {
			return indices
		}
		return nil
	}
	// No extension filter: return all indices
	all := make([]int, len(fi.allFiles))
	for i := range all {
		all[i] = i
	}
	return all
}

// matchPattern checks if a file matches a glob-style pattern.
// Handles recursive "**" globs by delegating to the existing matching functions.
func (fi *FileIndex) matchPattern(pattern, relPath, baseName string) bool {
	if strings.Contains(pattern, "**") {
		prefix, suffix := parseGlobPattern(pattern)
		if prefix != "" {
			prefix = strings.TrimSuffix(prefix, "/")
			prefix = strings.TrimSuffix(prefix, string(filepath.Separator))
			if !strings.HasPrefix(relPath, prefix+"/") && !strings.HasPrefix(relPath, prefix+string(filepath.Separator)) {
				return false
			}
			relPath = relPath[len(prefix)+1:]
		}
		if suffix == "" {
			return true
		}
		return matchGlobSuffix(suffix, relPath, baseName)
	}
	// Simple pattern: match against relative path
	matched, _ := filepath.Match(pattern, relPath)
	if matched {
		return true
	}
	// Also try just the basename for simple filename patterns
	matched, _ = filepath.Match(pattern, baseName)
	return matched
}

// activeFileIndex is the package-level index set by ExtractAll for the duration
// of an extraction run. When non-nil, findFiles uses this instead of walking
// the filesystem per call.
var activeFileIndex *FileIndex
