package config

// Config 全局配置信息
type Config struct {
	LogConfig      LogConfig      `yaml:"log" mapstructure:"log"`
	MySQLConfig    MySQLConfig    `yaml:"mysql" mapstructure:"mysql"`
	DatabaseConfig DatabaseConfig `yaml:"database" mapstructure:"database"`
}

// LogConfig 日志配置信息
type LogConfig struct {
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}

// MySQLConfig MySQL配置信息
type MySQLConfig struct {
	MaxOpenConnections int `yaml:"max_open_connections" mapstructure:"max_open_connections"`
	MaxIdleConnections int `yaml:"max_idle_connections" mapstructure:"max_idle_connections"`
}

// DatabaseConfig 数据库配置信息
type DatabaseConfig struct {
	Type     string `yaml:"type" mapstructure:"type"`
	Name     string `yaml:"name" mapstructure:"name"`
	Host     string `yaml:"host" mapstructure:"host"`
	Port     string `yaml:"port" mapstructure:"port"`
	User     string `yaml:"user" mapstructure:"user"`
	Password string `yaml:"password" mapstructure:"password"`
}
