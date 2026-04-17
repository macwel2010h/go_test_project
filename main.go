package main

import (
	"database/sql"
	"log/slog"
	handlers "serv-test/api"
	"serv-test/database"
	"serv-test/internal/models"
	runServer "serv-test/server"
)

type App struct {
	Logger *slog.Logger
	DB     *sql.DB
}

var app = &App{}

func main() {
	app.DB = database.DatabaseConnect()
	handlers.SetUserModel(&models.UserModel{DB: app.DB})

	mux := RoutHandlers()
	runServer.RunServer(mux)
}
