package config

import "time"

type Config struct {
	App    AppConfig    `json:"app" mapstructure:"app"`
	Http   HttpConfig   `json:"http" mapstructure:"http"`
	Redis  RedisConfig  `json:"redis" mapstructure:"redis"`
	Remote RemoteConfig `json:"remote" mapstructure:"remote"`
}

type AppConfig struct {
	Address string `json:"address" mapstructure:"address"`
	Label   string `json:"label" mapstructure:"label"`
}

type HttpConfig struct {
	Timeout time.Duration `json:"timeout" mapstructure:"timeout"`
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
