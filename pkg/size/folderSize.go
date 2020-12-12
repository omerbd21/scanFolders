package size

import (
	"os"
	"path/filepath"
)
// Gets the size of a folder and returns it in KB
func FolderSize(path string) (int64, error){
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir(){
			size += info.Size()
		}
		return err
	})
	return size /1024, err
}
