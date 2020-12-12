package writeToFile

import (
	"log"
	"os"
	"strings"
)

func Write(files []string, size string){
	f, err := os.Create("files.txt")
	if err != nil {
		log.Fatal(err)
	}
	filesInBytes := "\x00" + strings.Join(files, "\x20\x00")
	data := []byte(filesInBytes)
	f.Write(data)
	f.Write([]byte(size))
	defer f.Close()
}
