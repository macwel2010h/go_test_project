package main

import (
	"fmt"
	"serv-test/database"
	runServer "serv-test/server"
)

func main() {
	fmt.Println("starting the web server...")

	database.DatabaseConnect()
	mux := RoutHandlers()
	runServer.RunServer(mux)

}
