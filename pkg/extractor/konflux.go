package extractor

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// KonfluxSnapshot represents a parsed Konflux/EC snapshot file that maps
// container images to their source repositories and commits.
type KonfluxSnapshot struct {
	Application string              `json:"application,omitempty"`
	Components  []KonfluxComponent  `json:"components"`
	Source      string              `json:"source,omitempty"`
}

// KonfluxComponent is a single component entry in a Konflux snapshot.
type KonfluxComponent struct {
	Name           string `json:"name"`
	ContainerImage string `json:"containerImage"`
	Repository     string `json:"repository,omitempty"`
	Revision       string `json:"revision,omitempty"`
}

// KonfluxImageIndex is a lookup table from container image to source info,
// built from one or more snapshots.
type KonfluxImageIndex struct {
	Images    map[string]KonfluxComponent `json:"images"`
	Snapshots int                         `json:"snapshots_parsed"`
}

// Lookup returns the source info for a container image reference.
// It tries exact match first, then matches by image name (ignoring tag/digest).
func (idx *KonfluxImageIndex) Lookup(image string) (KonfluxComponent, bool) {
	if c, ok := idx.Images[image]; ok {
		return c, true
	}
	// Try without tag/digest
	base := stripTagDigest(image)
	if c, ok := idx.Images[base]; ok {
		return c, true
	}
	// Try matching by base name
	imageName := imageBaseName(image)
	for key, c := range idx.Images {
		if imageBaseName(key) == imageName {
			return c, true
		}
	}
	return KonfluxComponent{}, false
}

// Components returns all components sorted by name.
func (idx *KonfluxImageIndex) Components() []KonfluxComponent {
	seen := make(map[string]bool)
	var out []KonfluxComponent
	for _, c := range idx.Images {
		if !seen[c.Name] {
			seen[c.Name] = true
			out = append(out, c)
		}
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i].Name < out[j].Name
	})
	return out
}

// ParseKonfluxSnapshot reads a single Konflux snapshot file (JSON).
// The expected format is either the full Snapshot CR or a components array.
func ParseKonfluxSnapshot(path string) (*KonfluxSnapshot, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading snapshot: %w", err)
	}

	// Try full Snapshot CR format first
	var cr struct {
		Kind string `json:"kind"`
		Spec struct {
			Application string `json:"application"`
			Components  []struct {
				Name           string `json:"name"`
				ContainerImage string `json:"containerImage"`
				Source         struct {
					Git struct {
						URL      string `json:"url"`
						Revision string `json:"revision"`
					} `json:"git"`
				} `json:"source"`
			} `json:"components"`
		} `json:"spec"`
	}
	if err := json.Unmarshal(data, &cr); err == nil && len(cr.Spec.Components) > 0 {
		snap := &KonfluxSnapshot{
			Application: cr.Spec.Application,
			Source:      path,
		}
		for _, c := range cr.Spec.Components {
			snap.Components = append(snap.Components, KonfluxComponent{
				Name:           c.Name,
				ContainerImage: c.ContainerImage,
				Repository:     c.Source.Git.URL,
				Revision:       c.Source.Git.Revision,
			})
		}
		return snap, nil
	}

	// Try flat components array
	var flat struct {
		Application string `json:"application"`
		Components  []struct {
			Name           string `json:"name"`
			ContainerImage string `json:"containerImage"`
			Repository     string `json:"repository"`
			Revision       string `json:"revision"`
			Source         struct {
				Git struct {
					URL      string `json:"url"`
					Revision string `json:"revision"`
				} `json:"git"`
			} `json:"source"`
		} `json:"components"`
	}
	if err := json.Unmarshal(data, &flat); err == nil && len(flat.Components) > 0 {
		snap := &KonfluxSnapshot{
			Application: flat.Application,
			Source:      path,
		}
		for _, c := range flat.Components {
			repo := c.Repository
			rev := c.Revision
			if repo == "" {
				repo = c.Source.Git.URL
			}
			if rev == "" {
				rev = c.Source.Git.Revision
			}
			snap.Components = append(snap.Components, KonfluxComponent{
				Name:           c.Name,
				ContainerImage: c.ContainerImage,
				Repository:     repo,
				Revision:       rev,
			})
		}
		return snap, nil
	}

	return nil, fmt.Errorf("unrecognized snapshot format in %s", path)
}

// ParseKonfluxDir scans a directory for Konflux snapshot files and builds
// a unified image index.
func ParseKonfluxDir(dir string) (*KonfluxImageIndex, error) {
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return nil, fmt.Errorf("resolving path: %w", err)
	}

	idx := &KonfluxImageIndex{
		Images: make(map[string]KonfluxComponent),
	}

	patterns := []string{"*.json", "snapshot*.yaml", "snapshot*.yml"}
	for _, pattern := range patterns {
		matches, _ := filepath.Glob(filepath.Join(absDir, pattern))
		for _, match := range matches {
			snap, err := ParseKonfluxSnapshot(match)
			if err != nil {
				continue // skip non-snapshot files
			}
			idx.Snapshots++
			for _, c := range snap.Components {
				if c.ContainerImage != "" {
					idx.Images[c.ContainerImage] = c
					// Also index by base (without tag)
					base := stripTagDigest(c.ContainerImage)
					if base != c.ContainerImage {
						idx.Images[base] = c
					}
				}
			}
		}
	}

	return idx, nil
}

func stripTagDigest(image string) string {
	// Remove @sha256:... digest
	if idx := strings.Index(image, "@"); idx > 0 {
		return image[:idx]
	}
	// Remove :tag
	if idx := strings.LastIndex(image, ":"); idx > 0 {
		// Make sure it's not part of a port number (e.g., localhost:5000/image)
		after := image[idx+1:]
		if !strings.Contains(after, "/") {
			return image[:idx]
		}
	}
	return image
}

func imageBaseName(image string) string {
	s := stripTagDigest(image)
	if idx := strings.LastIndex(s, "/"); idx >= 0 {
		return s[idx+1:]
	}
	return s
}
