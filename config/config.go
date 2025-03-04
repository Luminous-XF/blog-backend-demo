package config

import "time"

// Config 全局配置信息
type Config struct {
	ServerConfig   ServerConfig   `yaml:"server" mapstructure:"server"`
	LogConfig      LogConfig      `yaml:"log" mapstructure:"log"`
	MySQLConfig    MySQLConfig    `yaml:"mysql" mapstructure:"mysql"`
	DatabaseConfig DatabaseConfig `yaml:"database" mapstructure:"database"`
	JWTConfig      JWTConfig      `yaml:"jwt" mapstructure:"jwt"`
	RedisConfig    RedisConfig    `yaml:"redis" mapstructure:"redis"`
	EmailConfig    EmailConfig    `yaml:"email" mapstructure:"email"`
}

// LogConfig 日志配置信息
type LogConfig struct {
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}

// MySQLConfig MySQL配置信息
type MySQLConfig struct {
	MaxOpenConnections int `yaml:"maxOpenConnections" mapstructure:"maxOpenConnections"`
	MaxIdleConnections int `yaml:"maxIdleConnections" mapstructure:"maxIdleConnections"`
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

// ServerConfig 服务器配置信息
type ServerConfig struct {
	Mode         string        `yaml:"mode" mapstructure:"mode"`
	Addr         int           `yaml:"addr" mapstructure:"addr"`
	ReadTimeout  time.Duration `yaml:"readTimeout" mapstructure:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout" mapstructure:"writeTimeout"`
}

// JWTConfig JWT 配置信息
type JWTConfig struct {
	SigningKey  string `yaml:"signingKey"`
	ExpiresTime int64  `yaml:"expiresTime"`
	BufferTime  int64  `yaml:"bufferTime"`
}

// RedisConfig Redis 配置信息
type RedisConfig struct {
	Addr     string `yaml:"addr" mapstructure:"addr"`
	Password string `yaml:"password" mapstructure:"password"`
	DB       int    `yaml:"db" mapstructure:"db"`
}

// EmailConfig 邮件相关配置
type EmailConfig struct {
	Addr       string `yaml:"addr"`
	LicenseKey string `yaml:"license"`
}
