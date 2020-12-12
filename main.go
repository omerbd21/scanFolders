package main

import (
	"fmt"
	"os"
	"scanFolders/cmd/scan"
	"scanFolders/cmd/size"
	"scanFolders/cmd/writeToFile"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(`Usage: main.exe [folder] [filename]
folder: The folder you want to scan
filename: The file to write the output to.
Omer Ben David TM`)
	}
	folderRoot := os.Args[1]
	filename := os.Args[2]
	writeToFile.WriteFile(scan.Scan(folderRoot), size.Size(folderRoot), filename)
}
