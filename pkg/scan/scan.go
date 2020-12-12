package scan

import (
	"os"
	"path/filepath"
)

func ScanFolder(rootFolder string) []string {
	var files []string

	err := filepath.Walk(rootFolder, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}
