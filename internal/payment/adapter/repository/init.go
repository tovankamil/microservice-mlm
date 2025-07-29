package repository

import (
	"e-wallet-microservices/internal/payment/core/ports"
	"e-wallet-microservices/lib/config"
	"net/http"
	"time"
)

type PaymentRepository struct {
	Client *http.Client
	config *config.Config
}

func NewRepository(cfg *config.Config) ports.PaymentRepositoryAdapter {
	return &PaymentRepository{
		config: cfg,
		Client: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        10,
				IdleConnTimeout:     10 * time.Second,
				MaxIdleConnsPerHost: 100,
			},
		},
	}
}
