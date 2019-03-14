package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	databaseConn *sql.DB
)

func init() {
	var (
		err error
	)

	databaseConn, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:8806)/?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
}
