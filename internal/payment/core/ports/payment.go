package ports

import (
	"context"
	"e-wallet-microservices/internal/payment/core/models"
)

// PaymentServiceAdapter - port primary
type PaymentServiceAdapter interface {
}

// PaymentRepositoryAdapter - port secondary
type PaymentRepositoryAdapter interface {
	ReadBalanceInfoFromDatastore(ctx context.Context, userID string) (models.DatastoreBalanceResponse, error)
	AppendBalanceInfoIntoDatastore(ctx context.Context, userID string, amount float64) error
}
