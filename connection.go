package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func connectDB() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/golang")

	if err != nil {
		panic(err.Error())
	}
}
