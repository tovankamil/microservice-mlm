package main

import (
	"e-wallet-microservices/lib/config"
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig(configName string) (*config.Config, error) {
	cfg := &config.Config{}
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName(configName)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(cfg)
	fmt.Print("%#v\n", cfg)

	return cfg, nil
}
