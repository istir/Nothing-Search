package backend

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "image/jpeg"
	_ "image/png"
)

type FileLoader struct {
	http.Handler
}

func (h *FileLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var err error
	fileName := req.URL.Path
	// var thumbnailSize int
	// size, err := strconv.Atoi(req.URL.Query().Get("size"))
	// if err != nil {
	// 	thumbnailSize = 0
	// } else {
	// 	thumbnailSize = size
	// }
	thumbPath := GetThumbnailPath(fileName)
	if _, err := os.Stat(thumbPath); err != nil {
		// create thumbnail first
		fmt.Println("Creating thumbnail", thumbPath, fileName)
		CreateThumbnail(fileName)
	}
	file, err := os.Open(thumbPath)
	if err != nil {
		fmt.Println("Error opening image file:", err)
		return
	}
	defer file.Close()

	// if thumbnailSize > 0 {
	// 	img, err := getImageFromFile(file)
	// 	if err != nil {
	// 		return
	// 	}
	// 	if img == nil {
	// 		res.Write(convertFileToBytes(file))
	// 		return
	// 	}
	// 	thumbnail := createThumbnail(img, thumbnailSize)
	// 	if thumbnail != nil {
	// 		res.Write(thumbnail)
	// 		return
	// 	} else {
	// 		println("error creating thumb")
	// 	}

	// 	res.Write(convertFileToBytes(file))
	// 	defer file.Close()
	// }
	if file != nil {
		b, err := convertFileToBytes(file)
		if err != nil {
			return
		}
		status, err := res.Write(b)
		if err != nil {
			log.Fatal("ERROR", err, status)
		}
	}

}

// func getImageFromFile(file *os.File) (image.Image, error) {

// 	img, _, err := image.Decode(file)
// 	if err != nil {
// 		fmt.Println("Error decoding image file:", err)
// 		return nil, err
// 	}

// 	return img, nil
// }

func convertFileToBytes(file *os.File) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(file)
	if err != nil {
		// fmt.Println("Error converting file to bytes:", err)
		log.Fatal("Error converting file to bytes", err)
		return nil, err
	}
	bufferBytes := buf.Bytes()
	if len(bufferBytes) <= 0 {
		return nil, errors.New("No bytes for file")
	}
	return bufferBytes, nil
}

// func createThumbnail(image image.Image, size int) []byte {
// 	thumbnail := imaging.Fit(image, size, size, imaging.Lanczos)
// 	buf := new(bytes.Buffer)
// 	err := imaging.Encode(buf, thumbnail, imaging.PNG)
// 	if err != nil {
// 		// handle error
// 		return []byte{}
// 	}
// 	return buf.Bytes()
// }

func NewFileLoader() *FileLoader {
	return &FileLoader{}
}
