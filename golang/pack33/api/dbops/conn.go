package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db_driver = "mysql"
	db_connection_dsn = "root:root@tcp(localhost:3306)/go_demo?charset=utf8"

	Connection *sql.DB
	err error
)

func init() {
	Connection, err = sql.Open(db_driver,
		db_connection_dsn)
	if err != nil {
		panic(err.Error())
	}
}
