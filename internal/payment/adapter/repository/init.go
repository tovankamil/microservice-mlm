package repository

import (
	"e-wallet-microservices/internal/payment/core/ports"
	"e-wallet-microservices/lib/config"
)

type PaymentRepository struct {
	config *config.Config
}

func NewRepository(cfg *config.Config) ports.PaymentRepositoryAdapter {
	return &PaymentRepository{
		config: cfg,
	}
}
