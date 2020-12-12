package scan

import (
	"os"
	"path/filepath"
	"strings"
)
// Scans a folder and returns list of files
func ScanFolder(rootFolder string) []string {
	var files []string

	err := filepath.Walk(rootFolder, func(path string, info os.FileInfo, err error) error {
		pathWithoutPrefix := strings.TrimPrefix(path, rootFolder)
		files = append(files, pathWithoutPrefix)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}
