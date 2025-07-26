package main

import (
	"e-wallet-microservices/internal/wallet/adapter/handler"
	"e-wallet-microservices/internal/wallet/adapter/repository"
	"e-wallet-microservices/internal/wallet/core/services"
	"e-wallet-microservices/lib/config"
	"fmt"

	"github.com/go-redis/redis/v8"

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

func initHandler(cfg *config.Config, redisClient *redis.Client) (*handler.Handler, error) {
	repo := repository.NewRepository(cfg, redisClient)
	hdl, err := handler.NewHandler(cfg, services.NewService(cfg, repo))
	if err != nil {
		return nil, err
	}

	return hdl, nil
}

func InitRedis(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		PoolSize: cfg.Redis.PoolSize,
	})
	return client
}
