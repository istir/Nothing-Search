package main

import (
	"context"
	"fmt"
	"nothingsearch/backend"
	"nothingsearch/backend/database"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	database.ConnectToDb()
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) LoadAllFilesToDb(baseDir string) {
	go backend.LoadAllFilesToDb(baseDir)
}

func (a *App) LoadFiles(baseDir string, lastEvaluatedKey string) backend.FileResponse {
	return backend.LoadFiles(baseDir, lastEvaluatedKey)
}
