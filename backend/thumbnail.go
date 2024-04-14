package backend

import (
	"fmt"
	"image"
	"log"
	"os"

	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"

	"github.com/disintegration/imaging"
)

func CreateThumbnail(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening image file:", err)
		return
	}
	defer file.Close()
	img, err := getImageFromFile(file)
	if err != nil {
		return
	}
	if img == nil {
		fmt.Printf("Error creating thumb for file %s", filePath)
		return
	}
	thumbnail := createThumbnail(img, 256)
	if thumbnail == nil {
		println("error creating thumb")
		return
	}
	saveThumbnailToDisk(filePath, thumbnail)
	// GetThumbnailPath
}

func saveThumbnailToDisk(filePath string, img image.Image) {
	thumbPath := GetThumbnailPath(filePath)
	file, err := os.Create(thumbPath)
	if err != nil {
		log.Fatal("Error creating thumb dir!", err)
	}
	defer file.Close()

	err = jpeg.Encode(file, img, &jpeg.Options{Quality: 50})
	if err != nil {
		log.Fatal("Error encoding", err)
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

func createThumbnail(image image.Image, size int) image.Image {
	thumbnail := imaging.Fit(image, size, size, imaging.Lanczos)
	return thumbnail
	// buf := new(bytes.Buffer)
	// err := imaging.Encode(buf, thumbnail, imaging.JPEG)
	// if err != nil {
	// 	// handle error
	// 	return []byte{}
	// }
	// return buf.Bytes()
}
