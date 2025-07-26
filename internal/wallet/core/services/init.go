package services

import (
	"e-wallet-microservices/internal/wallet/core/ports"
	"e-wallet-microservices/lib/config"
)

type WalletService struct {
	config     *config.Config
	repository ports.WalletRepositoryAdapter
}

func NewService(cfg *config.Config, repository ports.WalletRepositoryAdapter) ports.WalletServiceAdapter {
	return &WalletService{
		config:     cfg,
		repository: repository,
	}
}
