package backend

import (
	"bytes"
	"fmt"
	"image"
	"net/http"
	"os"
	"strconv"

	"github.com/disintegration/imaging"

	_ "image/jpeg"
	_ "image/png"
)

type FileLoader struct {
	http.Handler
}

func (h *FileLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var err error
	// image.RegisterFormat("png", "PNG", png.Decode, png.DecodeConfig)
	// image.RegisterFormat("jpeg", "JPEG", jpeg.Decode, jpeg.DecodeConfig)
	// requestedFilename := strings.TrimPrefix(req.URL.Path, "/")
	fileName := req.URL.Path
	var thumbnailSize int
	size, err := strconv.Atoi(req.URL.Query().Get("size"))
	if err != nil {
		thumbnailSize = 0
	} else {
		thumbnailSize = size
	}
	// fileName, thumbnailSize := getThumbnailDataFromThumbnail(requestedFilename)
	println("Requesting file:", fileName, size, thumbnailSize)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening image file:", err)
		return
	}
	defer file.Close()

	// return
	// file, err := os.Open(fileName)
	// fileData := file.
	// if err != nil {
	// 	res.WriteHeader(http.StatusBadRequest)
	// 	res.Write([]byte(fmt.Sprintf("Could not load file %s", fileName)))
	// }
	println("thumbnail size %s", thumbnailSize)
	if thumbnailSize > 0 {
		img, err := getImageFromFile(file)
		if err != nil {
			return
		}
		if img == nil {
			res.Write(convertFileToBytes(file))
			return
		}
		println("thumbnail exists %s", thumbnailSize)
		thumbnail := createThumbnail(img, thumbnailSize)
		if thumbnail != nil {
			res.Write(thumbnail)
			return
		} else {
			println("error creating thumb")
		}

		res.Write(convertFileToBytes(file))
		defer file.Close()
	}
}

func getImageFromFile(file *os.File) (image.Image, error) {

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image file:", err)
		return nil, err
	}

	return img, nil
}

func convertFileToBytes(file *os.File) []byte {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(file)
	if err != nil {
		fmt.Println("Error converting file to bytes:", err)
		return nil
	}
	return buf.Bytes()
}

func createThumbnail(image image.Image, size int) []byte {
	thumbnail := imaging.Fit(image, size, size, imaging.Lanczos)
	buf := new(bytes.Buffer)
	err := imaging.Encode(buf, thumbnail, imaging.PNG)
	if err != nil {
		// handle error
		return []byte{}
	}
	return buf.Bytes()
}

func NewFileLoader() *FileLoader {
	return &FileLoader{}
}
