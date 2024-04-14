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

func ConvertDbFileToFile(file db.File) File {
	return File{
		Name: file.Name.String,
		Url:  file.Path.String,
		// LastModified: file.DateModified.(int8),
		IsDir: false,
	}
}

func GetPathsFromFiles(files []File) []string {
	var paths []string
	for _, file := range files {
		paths = append(paths, file.Url)
	}
	return paths
}

func GetFilesByBasePath(path string) []File {
	response, err := dbQueries.GetFilesByDirectory(ctx, sql.NullString{String: path + "%", Valid: true})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	var items []File
	for _, item := range response {
		items = append(items, ConvertDbFileToFile(item))
	}
	return items
}

func GetFilesByPaths(paths []string) []File {
	p := make([]sql.NullString, len(paths))
	for i, v := range paths {
		p[i] = sql.NullString{
			String: v,
			Valid:  true,
		}
	}
	response, err := dbQueries.GetFilesByPaths(ctx, p)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	var items []File
	for _, item := range response {
		items = append(items, ConvertDbFileToFile(item))
	}
	return items
}

func GetFileByPath(path string) File {
	arr := make([]string, 1)
	arr = append(arr, path)
	return GetFilesByPaths(arr)[0]
}

func createAndInsertThumbnail(filePath string, qtx *db.Queries) {
	go CreateThumbnail(filePath)
	qtx.SetThumbnail(ctx, db.SetThumbnailParams{
		ThumbnailExists: sql.NullInt64{
			Int64: 1,
			Valid: true,
		},
		Path: sql.NullString{
			String: filePath,
			Valid:  true,
		},
	})
}

func AddThumbnailsForFiles(filePaths []string) {

	// tx, err := database.Begin()
	// if err != nil {
	// 	log.Fatal("balls", err)
	// 	return
	// }
	// qtx := dbQueries.WithTx(tx)
	// for _, filePath := range filePaths {
	// 	// createAndInsertThumbnail(filePath, qtx)
	// }
	// tx.Commit()
}

func AddFilesToDb(files []File) {
	// f := GetFiles(GetPathsFromFiles(files))
	// fmt.Printf("Found %d files", len(f))
	// var pathsToCreateThumbnails []string
	tx, err := database.Begin()
	if err != nil {
		log.Fatal("balls", err)
		return
	}
	defer tx.Rollback()

	qtx := dbQueries.WithTx(tx)
	for _, file := range files {
		// foundFile, err := qtx.GetFile(ctx, sql.NullString{
		// 	String: file.Url,
		// 	Valid:  true,
		// })
		// if err == nil {
		// 	log.Printf("File %s doesn't exist, creating thumbnail", file.Url)
		// 	pathsToCreateThumbnails = append(pathsToCreateThumbnails, file.Url)
		// 	// createAndInsertThumbnail(file.Url, qtx)
		// } else if foundFile.ThumbnailExists.Int64 == 0 {
		// 	// log.Printf("File %s exists without thumbnail, creating", file.Url)
		// 	pathsToCreateThumbnails = append(pathsToCreateThumbnails, file.Url)
		// 	// createAndInsertThumbnail(file.Url, qtx)
		// }
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
	// go AddThumbnailsForFiles(pathsToCreateThumbnails)
}
