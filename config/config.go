package config

import (
	"database/sql"
	"log/slog"
	"serv-test/database"
	logger "serv-test/log"

	"github.com/alexedwards/scs/v2"
)

type Application struct {
	Logger         *slog.Logger
	DB             *sql.DB
	SessionManager *scs.SessionManager
}

var App = &Application{
	Logger:         logger.Logger,
	DB:             database.DB,
	SessionManager: scs.New(),
}
