package ports

import (
	"context"
	"e-wallet-microservices/internal/payment/core/models"
)

// PaymentServiceAdapter - port primary
type PaymentServiceAdapter interface {
	TransferUserBalance(ctx context.Context, payload models.TransferBalancePayload) (float64, error)
}

// PaymentRepositoryAdapter - port secondary
type PaymentRepositoryAdapter interface {
	ReadBalanceInfoIntoWallet(ctx context.Context, userID string) (float64, error)
	AppendBalanceInfoIntoWallet(ctx context.Context, userID string, amount float64) error
}
