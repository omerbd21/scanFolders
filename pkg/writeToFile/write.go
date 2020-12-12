package writeToFile

import (
	"bufio"
	"fmt"
	"log"
	"os"
)
// Writes a list of files and the size to a file specified in filename
func Write(files []string, size string, filename string){
	filename +=".txt"
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	writer:= bufio.NewWriter(f)
	for _, file := range files{
		fmt.Fprintln(writer, file)
	}
	writer.Flush()
	fmt.Fprintln(f, "-------------------")
	f.Write([]byte(size))
	defer f.Close()
}
