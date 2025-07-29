package repository

import (
	"context"
	"e-wallet-microservices/internal/payment/core/models"
	"encoding/json"
	"net/http"
	"strings"
)

func (r *PaymentRepository) ReadBalanceInfoIntoWallet(ctx context.Context, userID string) (float64, error) {
	req, err := http.NewRequest("GET", "http://localhost:3002/balance/"+userID, nil)
	if err != nil {
		return 0, err
	}
	resp, err := r.Client.Do(req)
	if err != nil {

	}
	var data models.WalletBalanceResp
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
	}
	return float64(data.Balance), nil
}

func (r *PaymentRepository) AppendBalanceInfoIntoWallet(ctx context.Context, userID string, amount float64) error {

	url := r.config.Remote.Url.Address + "/v1/wallet/update"
	method := "POST"

	payload := models.UpdateBalancePayload{
		UserID: userID,
		Amount: amount,
	}
	payloadByte, err := json.Marshal(&payload)

	req, err := http.NewRequest(method, url, strings.NewReader(string(payloadByte)))
	if err != nil {
		// log error
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := r.Client.Do(req)
	if err != nil {
		// log error
	}
	defer res.Body.Close()

	var data models.WalletUpdateResponse
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		// log err
	}

	return err
}
