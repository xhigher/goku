package config

type MysqlConfig struct {
	IpPort       string `mapstructure:"ip-port" json:"ip-port" yaml:"ip-port"`
	DbName       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	Extras       string `mapstructure:"extras" json:"extras" yaml:"extras"`                         // 数据库名
	User         string `mapstructure:"user" json:"user" yaml:"user"`                               // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   // 数据库密码
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // 是否开启Gorm全局日志
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // 是否通过zap写入日志文件
}

func (m *MysqlConfig) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.IpPort + ")/" + m.DbName + "?" + m.Extras
}
