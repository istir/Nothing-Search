package backend

import (
	"fmt"
	"log"
	"os"
)

func LoadFiles(baseDir string) []File {
	// const baseDir = "/Users/lina/Library/Mobile Documents/com~apple~CloudDocs"
	files, err := os.ReadDir(baseDir)
	if err != nil {
		log.Fatal(err)
	}
	var fileT []File
	for _, file := range files {
		fileT = append(fileT, File{Name: file.Name(), IsDir: file.IsDir(), Url: fmt.Sprintf("%s/%s", baseDir, file.Name())})
	}

	return fileT
}
