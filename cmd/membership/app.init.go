package main

import (
	"e-wallet-microservices/internal/payment/adapter/handler"
	"e-wallet-microservices/internal/payment/adapter/repository"
	"e-wallet-microservices/internal/payment/core/services"
	"e-wallet-microservices/lib/config"
	"e-wallet-microservices/lib/datastore/postgre"
	"log"

	"github.com/go-redis/redis/v8"
)

func initHandler(cfg *config.Config) (*handler.Handler, error) {
	repo := repository.NewRepository(cfg)
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

func InitDB(cfg *config.Config) (postgre.DBInterface, error) {
	dbConn := postgre.InitDbConnection(cfg)
	dbinterface, err := dbConn.InitiateConnection()
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL database:", err)

	}
	return dbinterface, nil

}
