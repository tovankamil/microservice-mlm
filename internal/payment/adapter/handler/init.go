package handler

import (
	"e-wallet-microservices/internal/payment/core/ports"
	"e-wallet-microservices/lib/config"
	pbHandler "e-wallet-microservices/lib/protos/v1/payment"
)

type Handler struct {
	pbHandler.UnimplementedPaymentServiceServer
	config         *config.Config
	paymentService ports.PaymentServiceAdapter
}

func NewHandler(cfg *config.Config, adapter ports.PaymentServiceAdapter) (*Handler, error) {
	return &Handler{
		config:         cfg,
		paymentService: adapter,
	}, nil
}
