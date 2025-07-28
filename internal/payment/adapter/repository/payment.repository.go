package repository

import (
	"context"
	"e-wallet-microservices/internal/payment/core/models"
)

func (s *PaymentRepository) ReadBalanceInfoFromDatastore(ctx context.Context, userID string) (models.DatastoreBalanceResponse, error) {

	return models.DatastoreBalanceResponse{}, nil
}

func (s *PaymentRepository) AppendBalanceInfoIntoDatastore(ctx context.Context, userID string, amount float64) error {

	return nil
}
