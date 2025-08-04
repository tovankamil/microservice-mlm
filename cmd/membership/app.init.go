package main

import (
	"e-wallet-microservices/internal/payment/adapter/handler"
	"e-wallet-microservices/internal/payment/adapter/repository"
	"e-wallet-microservices/internal/payment/core/services"
	"e-wallet-microservices/lib/config"
)

func initHandler(cfg *config.Config) (*handler.Handler, error) {
	repo := repository.NewRepository(cfg)
	hdl, err := handler.NewHandler(cfg, services.NewService(cfg, repo))
	if err != nil {
		return nil, err
	}

	return hdl, nil
}
