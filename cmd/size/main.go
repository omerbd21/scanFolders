package size

import (
	"fmt"
	"scanFolders/pkg/size"
)

func Size(path string) string  {
	folderSize, _ := size.FolderSize(path)
	return fmt.Sprintln("Directory", path, "Size (KB)", folderSize)
}
