package config

import (
	"database/sql"
	"log/slog"
	logger "serv-test/log"

	"github.com/alexedwards/scs/v2"
)

type Application struct {
	Logger         *slog.Logger
	DB             *sql.DB
	SessionManager *scs.SessionManager
}

var app = &Application{
	Logger: logger.Logger,
}
