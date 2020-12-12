package scan

import (
	"fmt"
	"os"
	"path/filepath"
)

func Scan() {
	for _, file := range files {
		fmt.Println(file)
	}
}
