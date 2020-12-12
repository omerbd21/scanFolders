package main

import (
	"fmt"
	"os"
	"scanFolders/cmd/scan"
	"scanFolders/cmd/size"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter a folder path.")
	}
	folderRoot := os.Args[1]
	scan.Scan(folderRoot)
	size.Size(folderRoot)
}
