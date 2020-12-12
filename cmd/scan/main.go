package scan

import (
	"fmt"
	"scanFolders/pkg/scan"
)

func Scan(rootFolder string) {
	files := scan.ScanFolder(rootFolder)
	for _, file := range files {
		fmt.Println(file)
	}
}
