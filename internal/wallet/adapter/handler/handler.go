package handler

import (
	"context"
	"e-wallet-microservices/internal/wallet/core/models"
	"e-wallet-microservices/lib/protos/v1/wallet"
)

func (h *Handler) UpdateUserBalance(ctx context.Context, in *wallet.UpdateBalanceRequest) (*wallet.UpdateBalanceResponse, error) {
	var message = "success"
	amount, err := h.walletService.UpdateUserBalance(ctx, models.UpdateBalancePayload{
		UserID: in.GetUserId(),
		Amount: in.GetAmount(),
	})
	if err != nil {
		message = err.Error()
	}

	return &wallet.UpdateBalanceResponse{
		Message:      message,
		Success:      err == nil,
		FinalBalance: amount,
	}, nil
}

func (h *Handler) GetUserBalance(ctx context.Context, in *wallet.GetBalanceRequest) (*wallet.GetBalanceResponse, error) {
	resp, err := h.walletService.GetUserBalance(ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}
	data := &wallet.GetBalanceResponse{
		UserId:  in.GetUserId(),
		Balance: resp.AvailableBalance,
	}
	return data, nil
}
