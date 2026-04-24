package main

import (
	handlers "serv-test/api"
	"serv-test/database"
	"serv-test/internal/models"
	runServer "serv-test/server"
	"time"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
)

func main() {

	db := database.DatabaseConnect()
	handlers.SetUserModel(&models.UserModel{DB: app.DB})

	SessionManager := scs.New()
	SessionManager.Store = mysqlstore.New(app.DB)
	SessionManager.Lifetime = 12 * time.Hour

	mux := RoutHandlers()
	runServer.RunServer(mux)

}
