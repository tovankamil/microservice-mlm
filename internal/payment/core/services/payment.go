package services

import (
	"context"
	"e-wallet-microservices/internal/payment/core/models"
)

func (s *PaymentService) TransferUserBalance(ctx context.Context, payload models.TransferBalancePayload) (float64, error) {
	// deduct source user saldo
	s.repository.AppendBalanceInfoIntoWallet(ctx, payload.SourceUserID, payload.Amount*-1)
	// add destination user saldo
	s.repository.AppendBalanceInfoIntoWallet(ctx, payload.TargetUserID, payload.Amount)
	// get final destination user balance
	amount, _ := s.repository.ReadBalanceInfoIntoWallet(ctx, payload.TargetUserID)
	return amount, nil
}
