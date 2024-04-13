package backend

import (
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
