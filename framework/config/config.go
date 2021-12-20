package config

type ServerConfig struct {
	Zap   ZapConfig      `yaml:"zap"`
	Name  string         `yaml:"name"`
	Mysql []*MysqlConfig `yaml:"mysql"`
	Redis []*RedisConfig `yaml:"redis"`
}

type ZapConfig struct {
	Level         string `yaml:"level"`
	Format        string `yaml:"format"`
	Prefix        string `yaml:"prefix"`
	Director      string `yaml:"director"`
	ShowLine      bool   `yaml:"showLine"`
	EncodeLevel   string `yaml:"encode-level"`
	StacktraceKey string `yaml:"stacktrace-key"`
	LogInConsole  bool   `yaml:"log-in-console"`
}

type RedisConfig struct {
	Name     string `yaml:"name"`
	Db       int    `yaml:"db"`
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
}
