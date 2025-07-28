package services

import (
	"e-wallet-microservices/internal/payment/core/ports"
	"e-wallet-microservices/lib/config"
)

type PaymentService struct {
	config     *config.Config
	repository ports.PaymentRepositoryAdapter
}

func NewService(cfg *config.Config, repository ports.PaymentRepositoryAdapter) ports.PaymentServiceAdapter {
	return &PaymentService{
		config:     cfg,
		repository: repository,
	}
}
