package conf

// mysql 配置
type Databaser interface {
	GetDriver() string
	GetHost() string
	GetPort() int
	GetDbName() string
	GetUserName() string
	GetPassWord() string
	GetCharset() string
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
