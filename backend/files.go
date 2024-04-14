package backend

import (
	"fmt"
	"log"
	"os"
)

const LIMIT = 500

func verifyLastEvaluatedKey(baseDir string, lastEvaluatedKey *string) bool {
	return true
}

func LoadAllFilesToDb(baseDir string) bool {
	log.Printf("LOADING ALL FILES IN: %s", baseDir)
	files, err := os.ReadDir(baseDir)

	// I'll need to get a lil recursion going is file is a dir
	if err != nil {
		log.Fatal(err)
	}
	var filesResponse []File
	for _, file := range files {
		filePath := fmt.Sprintf("%s/%s", baseDir, file.Name())
		if file.IsDir() {
			go LoadAllFilesToDb(filePath)
			continue
		}
		filesResponse = append(filesResponse, File{Name: file.Name(), IsDir: file.IsDir(), Url: filePath})
	}
	go AddFilesToDb(filesResponse)
	return true
}

func getNextParam(files []File) string {
	var next string = files[len(files)-1].Url
	return next
}

func LoadFiles(baseDir string, lastEvaluatedKey string) FileResponse {
	// this should first get data from db actually, but in bg still check for files via old method?
	go LoadAllFilesToDb(baseDir)
	dbFiles := GetFilesByBasePath(baseDir)
	var response FileResponse
	response.Files = dbFiles
	response.Next = getNextParam(dbFiles)
	return response
	//
	// if !verifyLastEvaluatedKey(baseDir, lastEvaluatedKey) {
	// 	// return nil
	// }
	// const baseDir = "/Users/lina/Library/Mobile Documents/com~apple~CloudDocs"
	// files, err := os.ReadDir(baseDir)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var fileT []File
	// for iter, file := range files {
	// 	if iter > LIMIT {
	// 		break
	// 	}
	// 	fileT = append(fileT, File{Name: file.Name(), IsDir: file.IsDir(), Url: fmt.Sprintf("%s/%s", baseDir, file.Name())})
	// }
	// AddFilesToDb(fileT)
	// response.Files = fileT
	// response.Next = getNextParam(fileT)
	// return response
}
