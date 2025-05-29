package libx

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// ManifestEntry represents a single entry in the Vite manifest
type ManifestEntry struct {
	File           string   `json:"file,omitempty"`
	Src            string   `json:"src,omitempty"`
	IsEntry        bool     `json:"isEntry,omitempty"`
	IsDynamicEntry bool     `json:"isDynamicEntry,omitempty"`
	Name           string   `json:"name,omitempty"`
	Imports        []string `json:"imports,omitempty"`
	DynamicImports []string `json:"dynamicImports,omitempty"`
	CSS            []string `json:"css,omitempty"`
}

// ViteManifest represents the entire Vite manifest JSON structure
type ViteManifest map[string]ManifestEntry

// ParseManifestFile reads a Vite manifest JSON file from the given path and parses it
func ParseManifestFile(manifestPath string) (ViteManifest, error) {
	// Read the file
	data, err := os.ReadFile(manifestPath)
	if err != nil {
		return nil, err
	}

	// Parse the JSON
	var manifest ViteManifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, err
	}

	return manifest, nil
}

// ParseManifestString parses a Vite manifest from a JSON string
func ParseManifestString(manifestJSON string) (ViteManifest, error) {
	var manifest ViteManifest
	if err := json.Unmarshal([]byte(manifestJSON), &manifest); err != nil {
		return nil, err
	}
	return manifest, nil
}

// GetEntryPoints returns all the entry points in the manifest
func (m ViteManifest) GetEntryPoints() []string {
	var entryPoints []string
	for key, entry := range m {
		if entry.IsEntry {
			entryPoints = append(entryPoints, key)
		}
	}
	return entryPoints
}

// GetDynamicEntries returns all dynamic entries in the manifest
func (m ViteManifest) GetDynamicEntries() []string {
	var dynamicEntries []string
	for key, entry := range m {
		if entry.IsDynamicEntry {
			dynamicEntries = append(dynamicEntries, key)
		}
	}
	return dynamicEntries
}

// GetAssetPath returns the file path for a given module ID
func (m ViteManifest) GetAssetPath(moduleID string) string {
	if entry, exists := m[moduleID]; exists {
		return entry.File
	}
	return ""
}

// GetCSS returns all CSS files associated with a module ID
func (m ViteManifest) GetCSS(moduleID string) []string {
	if entry, exists := m[moduleID]; exists {
		return entry.CSS
	}
	return nil
}

// GetImports returns all imports of a module
func (m ViteManifest) GetImports(moduleID string) []string {
	if entry, exists := m[moduleID]; exists {
		return entry.Imports
	}
	return nil
}

// IsStaticFileExists checks if a static file path (like .css or .js) exists in the manifest
func (m ViteManifest) IsStaticFileExists(filePath string) bool {
	// Clean the path to handle different formats consistently
	filePath = strings.TrimPrefix(filePath, "/")

	// Check if this path directly matches any file in the manifest
	for _, entry := range m {
		if entry.File == filePath {
			return true
		}

		// Check CSS files as well
		for _, css := range entry.CSS {
			if css == filePath {
				return true
			}
		}
	}

	return false
}

// FindModuleByAssetPath returns the module ID that has the given asset path
func (m ViteManifest) FindModuleByAssetPath(assetPath string) (string, error) {
	for moduleID, entry := range m {
		if entry.File == assetPath {
			return moduleID, nil
		}

		// Check CSS files as well
		for _, css := range entry.CSS {
			if css == assetPath {
				return moduleID, nil
			}
		}
	}

	return "", errors.New("asset path not found in manifest")
}

// IsStaticFileExistsOnDisk checks if the static file actually exists on disk
func (m ViteManifest) IsStaticFileExistsOnDisk(filePath string, basePath string) bool {
	// If the file isn't in the manifest, it doesn't exist in our app
	if !m.IsStaticFileExists(filePath) {
		return false
	}

	// Check if the file exists on disk
	fullPath := filepath.Join(basePath, filePath)
	_, err := os.Stat(fullPath)
	return err == nil
}

// GetAllStaticFiles returns all static files (.js and .css) in the manifest
func (m ViteManifest) GetAllStaticFiles() []string {
	var files []string
	seen := make(map[string]bool)

	for _, entry := range m {
		if entry.File != "" && !seen[entry.File] {
			files = append(files, entry.File)
			seen[entry.File] = true
		}

		for _, css := range entry.CSS {
			if !seen[css] {
				files = append(files, css)
				seen[css] = true
			}
		}
	}

	return files
}

// GetAllStaticFilesByType returns all static files of a specific type (.js or .css)
func (m ViteManifest) GetAllStaticFilesByType(fileType string) []string {
	var files []string
	seen := make(map[string]bool)

	// Ensure fileType starts with a dot
	if !strings.HasPrefix(fileType, ".") {
		fileType = "." + fileType
	}

	for _, entry := range m {
		if entry.File != "" && !seen[entry.File] && strings.HasSuffix(entry.File, fileType) {
			files = append(files, entry.File)
			seen[entry.File] = true
		}

		// For CSS files, only include them if the fileType is .css
		if fileType == ".css" {
			for _, css := range entry.CSS {
				if !seen[css] {
					files = append(files, css)
					seen[css] = true
				}
			}
		}
	}

	return files
}

// GetDependencyTree returns a map of all modules that depend on the given module ID
func (m ViteManifest) GetDependencyTree(moduleID string) map[string]bool {
	result := make(map[string]bool)

	for id, entry := range m {
		for _, imp := range entry.Imports {
			if imp == moduleID {
				result[id] = true
				// Recursively find dependencies
				for depID := range m.GetDependencyTree(id) {
					result[depID] = true
				}
			}
		}

		for _, dynImp := range entry.DynamicImports {
			if dynImp == moduleID {
				result[id] = true
				// Recursively find dependencies
				for depID := range m.GetDependencyTree(id) {
					result[depID] = true
				}
			}
		}
	}

	return result
}

// SaveToFile saves the manifest back to a JSON file
func (m ViteManifest) SaveToFile(filePath string) error {
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}

	// Ensure the directory exists
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}
