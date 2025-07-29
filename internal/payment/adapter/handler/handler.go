package handler

import (
	"context"
	"e-wallet-microservices/internal/payment/core/models"
	"e-wallet-microservices/lib/protos/v1/payment"
)

func (h *Handler) TransferBalanceService(ctx context.Context, in *payment.TransferBalanceRequest) (*payment.TransferBalanceResponse, error) {
	balance, err := h.paymentService.TransferUserBalance(ctx, models.TransferBalancePayload{
		SourceUserID: in.GetSourceUserId(),
		TargetUserID: in.GetDestination(),
		Amount:       in.GetAmount(),
	})
	if err != nil {
		return nil, nil
	}

	return &payment.TransferBalanceResponse{
		Success:          err == nil,
		DestinationAmout: balance,
	}, nil

}
