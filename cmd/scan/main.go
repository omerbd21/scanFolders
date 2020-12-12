package scan

import (
	"scanFolders/pkg/scan"
)

func Scan(rootFolder string) []string {
	files := scan.ScanFolder(rootFolder)
	return files
}
