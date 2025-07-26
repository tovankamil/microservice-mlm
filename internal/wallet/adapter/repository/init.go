package repository

import (
	"e-wallet-microservices/internal/wallet/core/ports"
	"e-wallet-microservices/lib/config"

	"github.com/go-redis/redis/v8"
)

type WalletRepository struct {
	config *config.Config
	client *redis.Client
}

func NewRepository(cfg *config.Config, client *redis.Client) ports.WalletRepositoryAdapter {
	return &WalletRepository{
		config: cfg,
		client: client,
	}
}
