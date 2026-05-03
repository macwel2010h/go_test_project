package main

import (
	"serv-test/database"
	runServer "serv-test/server"
	"time"

	"serv-test/config"

	"github.com/alexedwards/scs/mysqlstore"
)

func main() {

	database.DatabaseConnect()

	config.App.SessionManager.Store = mysqlstore.New(config.App.DB)
	config.App.SessionManager.Lifetime = 12 * time.Hour

	mux := RoutHandlers()
	runServer.RunServer(mux)

}

// End
// 1
