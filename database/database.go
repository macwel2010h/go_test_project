package database

import (
	"context"
	"database/sql"
	logger "serv-test/log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func DatabaseConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:abcd@1234@tcp(localhost:3306)/user_db")
	if err != nil {
		panic(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		panic(err)
	}
	logger.Logger.Info("Database connected successfully.")
	return db
}

func InsertUser() {

}
