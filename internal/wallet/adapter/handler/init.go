package handler

import (
	"e-wallet-microservices/internal/wallet/core/ports"
	"e-wallet-microservices/lib/config"
	pbHandler "e-wallet-microservices/lib/protos/v1/wallet"
)

type Handler struct {
	pbHandler.UnimplementedWalletServiceServer
	config        *config.Config
	walletService ports.WalletServiceAdapter
}

func NewHandler(cfg *config.Config, adapter ports.WalletServiceAdapter) (*Handler, error) {
	return &Handler{
		config:        cfg,
		walletService: adapter,
	}, nil
}
