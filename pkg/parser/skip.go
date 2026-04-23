package parser

import "strings"

// ShouldSkipDir returns true if a directory should be excluded from parsing.
// Each language parser has its own skip list; this provides common directories
// that are always skipped regardless of language.
func ShouldSkipDir(dirName string, languageSkips []string) bool {
	for _, skip := range languageSkips {
		if dirName == skip {
			return true
		}
	}
	return false
}

// IsTestFile checks if a filename matches common test file patterns for a language.
func IsTestFile(filename string, patterns []string) bool {
	for _, pat := range patterns {
		if strings.HasPrefix(pat, "*") {
			// suffix match: *_test.py -> check if filename ends with _test.py
			suffix := pat[1:]
			if strings.HasSuffix(filename, suffix) {
				return true
			}
		} else if strings.HasSuffix(pat, "*") {
			// prefix match: test_* -> check if filename starts with test_
			prefix := pat[:len(pat)-1]
			if strings.HasPrefix(filename, prefix) {
				return true
			}
		} else if strings.Contains(pat, "*") {
			// wildcard in middle: *.spec.ts -> split and check
			parts := strings.SplitN(pat, "*", 2)
			if strings.HasPrefix(filename, parts[0]) && strings.HasSuffix(filename, parts[1]) {
				return true
			}
		}
	}
	return false
}
