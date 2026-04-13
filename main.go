package main

import (
	"log/slog"
	"serv-test/database"
	runServer "serv-test/server"
)

type app struct {
	Logger *slog.Logger
}

func main() {

	database.DatabaseConnect()
	mux := RoutHandlers()
	runServer.RunServer(mux)

}
