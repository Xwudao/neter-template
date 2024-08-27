package libx

import (
	"os"
	"path/filepath"
)

type StaticFile struct {
}

func NewStaticFile() *StaticFile {
	return &StaticFile{}
}

// WriteFile writes the contents of the file to the specified path.
func (s *StaticFile) WriteFile(filename string, data []byte) error {
	wd, _ := os.Getwd()
	fp := filepath.Join(wd, "web/public", filename)
	return os.WriteFile(fp, data, 0644)
}
