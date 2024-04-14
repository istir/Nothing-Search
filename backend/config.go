package backend

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path"
)

func GetUserConfigPath() string {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("Couldn't get user config directory", err)
	}
	var dir = path.Join(userConfigDir, "Nothing Search")
	createUserConfigPath(dir)
	createUserConfigPath(path.Join(dir, "thumbnails"))
	return dir
}

func createUserConfigPath(dir string) {
	_, err := os.Stat(dir)
	if err == nil {
		return
	}

	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		log.Fatal("Creating user dir failed", err)
	}
}

func GetThumbnailPath(filePath string) string {
	configPath := GetUserConfigPath()
	// extension := filepath.Ext(filePath)
	encodedName := calculateSHA1(filePath)
	name := fmt.Sprintf("%s.%s", encodedName, "jpg")
	return path.Join(configPath, "thumbnails", name)
}

func calculateSHA1(input string) string {
	hasher := sha1.New()
	hasher.Write([]byte(input))
	sha1Hash := hex.EncodeToString(hasher.Sum(nil))
	return sha1Hash
}
