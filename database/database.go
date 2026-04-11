package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func DatabaseConnect() {
	db, err := sql.Open("mysql", "root:abcd@1234>@tcp(127.0.0.1:3306)/helllo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		panic(err)
	}
	fmt.Println("database connected successfully")
}
