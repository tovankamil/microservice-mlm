package handler

import (
	"context"
	"e-wallet-microservices/lib/protos/v1/wallet"
)

// func (h *Handler) GetUserBalance(ctx context.Context, in *wallet.GetBalanceRequest) (*wallet.GetBalanceResponse, error) {
// 	resp, err := h.walletService.GetUserBalance(ctx, in.GetUserId())
// 	if err != nil {
// 		/// log err
// 		return nil, err
// 	}
// 	data := &wallet.GetBalanceResponse{
// 		UserId:  in.UserId,
// 		Balance: resp.AvailableBalance,
// 	}

// 	return data, nil
// }

// func (h *Handler) UpdateUserBalance(ctx context.Context, in *wallet.UpdateBalanceRequest) (*wallet.UpdateBalanceResponse, error) {
// 	var message = "success"
// 	amount, err := h.walletService.UpdateUserBalance(ctx, models.UpdateBalancePayload{
// 		UserID: in.GetUserId(),
// 		Amount: in.GetAmount(),
// 	})
// 	if err != nil {
// 		// log err
// 		message = err.Error()
// 	}

// 	return &wallet.UpdateBalanceResponse{
// 		Message:      message,
// 		Success:      err == nil,
// 		FinalBalance: amount,
// 	}, nil
// }

func (h *Handler) UpdateUserBalance(ctx context.Context, in *wallet.UpdateBalanceRequest) (*wallet.UpdateBalanceResponse, error) {
	return &wallet.UpdateBalanceResponse{
		Message:      "success",
		Success:      true,
		FinalBalance: in.Amount, // Mocked final balance for demonstration
	}, nil
}

func (h *Handler) GetUserBalance(ctx context.Context, in *wallet.GetBalanceRequest) (*wallet.GetBalanceResponse, error) {
	return &wallet.GetBalanceResponse{
		UserId:  in.UserId,
		Balance: 100.0, // Mocked balance for demonstration
	}, nil
}
