package main

import (
	"e-wallet-microservices/internal/payment/adapter/handler"
	"e-wallet-microservices/internal/payment/adapter/repository"
	"e-wallet-microservices/internal/payment/core/services"
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
	fmt.Printf("%#v\n", cfg)
	return cfg, err
}

func initHandler(cfg *config.Config) (*handler.Handler, error) {
	repo := repository.NewRepository(cfg)
	hdl, err := handler.NewHandler(cfg, services.NewService(cfg, repo))
	if err != nil {
		return nil, err
	}

	return hdl, nil
}
