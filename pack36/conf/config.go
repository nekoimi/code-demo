package conf

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
)

type DbWorker struct {
	dsn string
	Db *sql.DB
}

func NewMysql() DbWorker {
	var makeDsn = func(conf MysqlConfig) string {
		// 支持下面几种DSN写法，具体看mysql服务端配置，常见为第1种
		// user:password@tcp(localhost:5555)/dbname?charset=utf8
		// user@unix(/path/to/socket)/dbname?charset=utf8
		// user:password@/dbname
		// user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
		var dsnstr strings.Builder
		dsnstr.WriteString(GetUserName())
		dsnstr.WriteString(":")
		dsnstr.WriteString(GetPassWord())
		dsnstr.WriteString("@tcp(")
		dsnstr.WriteString(GetHost())
		dsnstr.WriteString(":")
		dsnstr.WriteString(strconv.Itoa(GetPort()))
		dsnstr.WriteString(")/")
		dsnstr.WriteString(GetDbName())
		dsnstr.WriteString("?charset=")
		dsnstr.WriteString(GetCharset())
		return dsnstr.String()
	}
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
		dsn:makeDsn(mysql),
		Db:nil,
	}
	dbWorker.Db, _ = sql.Open(GetDriver(), dbWorker.dsn)
	return dbWorker
}
