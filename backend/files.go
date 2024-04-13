package backend

import (
	"fmt"
	"log"
	"os"
)

const LIMIT = 500

func LoadFiles(baseDir string) FileResponse {
	// const baseDir = "/Users/lina/Library/Mobile Documents/com~apple~CloudDocs"
	files, err := os.ReadDir(baseDir)
	if err != nil {
		log.Fatal(err)
	}
	var fileT []File
	for iter, file := range files {
		if iter > LIMIT {
			break
		}
		fileT = append(fileT, File{Name: file.Name(), IsDir: file.IsDir(), Url: fmt.Sprintf("%s/%s", baseDir, file.Name())})
	}
	AddFilesToDb(fileT)
	var response FileResponse
	response.Files = fileT
	return response
}
