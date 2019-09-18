package conf

import (
	"strconv"
	"strings"
)

type Databaser interface {
	GetDriver() string
	GetHost() string
	GetPort() int
	GetDbName() string
	GetUserName() string
	GetPassWord() string
	GetCharset() string
	GetDsnString() string
}

// mysql config
type MysqlConfig struct {
	driver   string
	host     string
	port     int
	dbname   string
	username string
	password string
	charset string
}

func (config *MysqlConfig) GetDriver() string {
	return config.driver
}

func (config *MysqlConfig) GetHost() string {
	return config.host
}

func (config *MysqlConfig) GetPort() int {
	return config.port
}

func (config *MysqlConfig) GetDbName() string {
	return config.dbname
}

func (config *MysqlConfig) GetUserName() string {
	return config.username
}

func (config *MysqlConfig) GetPassWord() string {
	return config.password
}

func (config *MysqlConfig) GetCharset() string {
	return config.charset
}

func (config *MysqlConfig) GetDsnString() string {
	// 支持下面几种DSN写法，具体看mysql服务端配置，常见为第1种
	// user:password@tcp(localhost:5555)/dbname?charset=utf8
	// user@unix(/path/to/socket)/dbname?charset=utf8
	// user:password@/dbname
	// user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
	var dsnstr strings.Builder
	dsnstr.WriteString(config.GetUserName())
	dsnstr.WriteString(":")
	dsnstr.WriteString(config.GetPassWord())
	dsnstr.WriteString("@tcp(")
	dsnstr.WriteString(config.GetHost())
	dsnstr.WriteString(":")
	dsnstr.WriteString(strconv.Itoa(config.GetPort()))
	dsnstr.WriteString(")/")
	dsnstr.WriteString(config.GetDbName())
	dsnstr.WriteString("?charset=")
	dsnstr.WriteString(config.GetCharset())
	return dsnstr.String()
}
