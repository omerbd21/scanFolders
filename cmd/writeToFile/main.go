package writeToFile

import "scanFolders/pkg/writeToFile"

func WriteFile(files []string, size string, filename string)  {
	writeToFile.Write(files, size, filename)
}
