package conf

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DbWorker struct {
	dsn string
	Db *sql.DB
}

func NewMysql() DbWorker {
	var mysql = MysqlConfig{
		"mysql",
		"192.168.1.186",
		3306,
		"hyperf",
		"root",
		"root",
		"utf8mb4",
	}
	var dbWorker = DbWorker{
		dsn: mysql.GetDsnString(),
		Db:nil,
	}
	dbWorker.Db, _ = sql.Open(mysql.GetDriver(), dbWorker.dsn)
	return dbWorker
}
