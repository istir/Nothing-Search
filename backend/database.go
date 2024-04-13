package backend

import (
	"context"
	"database/sql"
	"log"
	"nothingsearch/db"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB
var dbQueries *db.Queries

var ctx = context.Background()

func initDb() {
	createDatabase()
	createdDb, err := sql.Open("sqlite3", getDatabasePath())

	if err != nil {
		log.Fatal("Couldn't connect to db", err)
	}
	dbQueries = db.New(createdDb)
	database = createdDb
	// createTable("files")
}

func ConnectToDb() {
	if database != nil {
		return
	}
	initDb()
}

func getDatabasePath() string {
	var path = path.Join(GetUserConfigPath(), "database.db")
	return path
}

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

// func createTable(name string) {

// 	var createTableQuery = fmt.Sprintf(`
// 		CREATE TABLE IF NOT EXISTS %s (
// 			"id"	INTEGER UNIQUE,
// 			"name"	TEXT,
// 			"path"	INTEGER UNIQUE,
// 			"thumbnail_exists"	INTEGER DEFAULT 0,
// 			"date_modified"	INTEGER,
// 			"date_created"	INTEGER,
// 			PRIMARY KEY("id" AUTOINCREMENT)
// )
// 	`, name)
// 	database.Exec(createTableQuery, name)

// }

func AddFilesToDb(files []File) {
	// var payload
	// for index, file := range files {

	// }
	//
	tx, err := database.Begin()
	if err != nil {
		log.Fatal("balls", err)
		return
	}
	defer tx.Rollback()

	qtx := dbQueries.WithTx(tx)
	for _, file := range files {

		qtx.CreateFile(ctx, db.CreateFileParams{
			Name: sql.NullString{
				String: file.Name,
				Valid:  true,
			},
			Path: sql.NullString{
				String: file.Url,
				Valid:  true,
			},
			DateModified: sql.NullInt64{
				Int64: 0,
				Valid: false,
			},
			DateCreated: sql.NullInt64{
				Int64: 0,
				Valid: false,
			},
		})
	}
	tx.Commit()
}
