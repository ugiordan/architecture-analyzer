package srclang

import (
	"fmt"
	"os"
	"strings"
)

type BodyExtractor struct {
	cache map[string][]string
}

func NewBodyExtractor() *BodyExtractor {
	return &BodyExtractor{cache: make(map[string][]string)}
}

func (be *BodyExtractor) Extract(file string, startLine, endLine int) (string, error) {
	lines, err := be.loadFile(file)
	if err != nil {
		return "", err
	}

	if startLine < 1 {
		startLine = 1
	}
	if endLine > len(lines) {
		endLine = len(lines)
	}
	if startLine > len(lines) {
		return "", fmt.Errorf("start line %d exceeds file length %d", startLine, len(lines))
	}
	if startLine > endLine {
		return "", nil
	}

	selected := lines[startLine-1 : endLine]
	return strings.Join(selected, "\n"), nil
}

func (be *BodyExtractor) loadFile(path string) ([]string, error) {
	if cached, ok := be.cache[path]; ok {
		return cached, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}

	content := strings.TrimRight(string(data), "\n")
	if content == "" {
		be.cache[path] = nil
		return nil, nil
	}
	lines := strings.Split(content, "\n")
	be.cache[path] = lines
	return lines, nil
}
