package database

import (
	"log"
	"nothingsearch/backend/config"
	"os"
	"path"
)

func createDatabase() {
	dbPath := getDatabasePath()
	_, err := os.Stat(dbPath)
	if err == nil {
		return
	}
	if _, err := os.Create(dbPath); err != nil {
		log.Fatal("Creating database failed", err)
	}
}

func getDatabasePath() string {
	var path = path.Join(config.GetUserConfigPath(), "database.db")
	return path
}
