// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package venly

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/zeebo/errs"
)

// TransferFungibleRequest fields that required for transfer fungible request.
type TransferFungibleRequest struct {
	TransactionRequest struct {
		Type         string      `json:"type"`
		WalletID     string      `json:"walletId"`
		To           string      `json:"to"`
		SecretType   string      `json:"secretType"`
		TokenAddress string      `json:"tokenAddress"`
		Value        float64     `json:"value"`
		TokenID      interface{} `json:"tokenId"`
	} `json:"transactionRequest"`
	Pincode string `json:"pincode"`
}

// TransferFungibleResponse fields that returns from transfer fungible.
type TransferFungibleResponse struct {
	Success bool `json:"success"`
	Result  struct {
		TransactionHash string `json:"transactionHash"`
	} `json:"result"`
}

// TransferFungible transfers fungible token via Venly.
func (client *Client) TransferFungible(ctx context.Context, accessToken string, r TransferFungibleRequest) (response TransferFungibleResponse, err error) {
	jsonBody, err := json.Marshal(r)
	if err != nil {
		return TransferFungibleResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, client.config.DefaultURL+"transactions/execute", bytes.NewReader(jsonBody))
	if err != nil {
		return TransferFungibleResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.http.Do(req.WithContext(ctx))
	if err != nil {
		return TransferFungibleResponse{}, err
	}

	defer func() {
		err = errs.Combine(err, resp.Body.Close())
	}()

	if resp.StatusCode != http.StatusOK {
		errorResp := ErrorResponse{}

		if err = json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
			return TransferFungibleResponse{}, err
		}

		if !errorResp.Success {
			return TransferFungibleResponse{}, errs.New(errorResp.Errors[0].Code)
		}

		return TransferFungibleResponse{}, errs.New(resp.Status)
	}

	var transferFungibleResponse TransferFungibleResponse
	if err = json.NewDecoder(resp.Body).Decode(&transferFungibleResponse); err != nil {
		return TransferFungibleResponse{}, err
	}

	return transferFungibleResponse, nil
}
