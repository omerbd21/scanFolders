package size

import (
	"fmt"
	"scanFolders/pkg/size"
)

func Size(path string)  {
	folderSize, _ := size.FolderSize(path)
	fmt.Println("Directory", path, "Size (KB)", folderSize)
}
