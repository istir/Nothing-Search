package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DbClient *gorm.DB

func initDb() {
	createDatabase()

	gormDb, err := gorm.Open(sqlite.Open(getDatabasePath()), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		panic("failed to connect database")
	}

	DbClient = gormDb
	migrate()
}

func ConnectToDb() {
	if DbClient != nil {
		return
	}
	initDb()
}
