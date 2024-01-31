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
		fmt.Println(fmt.Sprintf("file://%s/%s", baseDir, file.Name()))
		fileT = append(fileT, File{Name: file.Name(), IsDir: file.IsDir(), url: fmt.Sprintf("file://%s/%s", baseDir, file.Name())})
	}

	return fileT
}
