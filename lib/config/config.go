package config

import "time"

type Config struct {
	App     AppConfig     `json:"app" mapstructure:"app"`
	Http    HttpConfig    `json:"http" mapstructure:"http"`
	Redis   RedisConfig   `json:"redis" mapstructure:"redis"`
	Postgre PostgreConfig `json:"postgre" mapstructure:"postgre"`
	Remote  RemoteConfig  `json:"remote" mapstructure:"remote"`
}

type AppConfig struct {
	Address string `json:"address" mapstructure:"address"`
	Label   string `json:"label" mapstructure:"label"`
}

type HttpConfig struct {
	Timeout time.Duration `json:"timeout" mapstructure:"timeout"`
}

type PostgreConfig struct {
	Address            string `json:"address" mapstructure:"address"`
	Ports              string `json:"ports" mapstructure:"ports"`
	Database           string `json:"database" mapstructure:"database"`
	Username           string `json:"username" mapstructure:"username"`
	Password           string `json:"password" mapstructure:"password"`
	SSLMode            string `json:"sslmode" mapstructure:"sslmode"`
	MaxOpenConnections int    `json:"max_open_connections" mapstructure:"max_open_connections"`
	MaxIdleConnections int    `json:"max_idle_connections" mapstructure:"max_idle_connections"`
	ConnMaxIdleTime    int    `json:"connmaxidletime" mapstructure:"connmaxidletime"`
	ConnMaxLifetime    int    `json:"connmaxlifetime" mapstructure:"connmaxlifetime"`
}

type RedisConfig struct {
	Address  string `json:"address" mapstructure:"address"`
	PoolSize int    `json:"poolsize" mapstructure:"poolsize"`
}

type RemoteConfig struct {
	Url struct {
		HostApp string        `json:"hostapp" mapstructure:"hostapp"`
		Address string        `json:"address" mapstructure:"address"`
		Timeout time.Duration `json:"timeout" mapstructure:"timeout"`
	} `json:"url" mapstructure:"url"`
}
