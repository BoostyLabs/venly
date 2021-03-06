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

// TransferNativeRequest fields that required for transfer native request.
type TransferNativeRequest struct {
	TransactionRequest TxRequest `json:"transactionRequest"`
	Pincode            string    `json:"pincode"`
}

// TxRequest struct ...
type TxRequest struct {
	Type       string  `json:"type"`
	WalletID   string  `json:"walletId"`
	To         string  `json:"to"`
	SecretType string  `json:"secretType"`
	Value      float64 `json:"value"`
	Data       string  `json:"data"`
}

// TransferNativeResponse fields that returns from transfer native.
type TransferNativeResponse struct {
	Success bool `json:"success"`
	Result  struct {
		TransactionHash string `json:"transactionHash"`
	} `json:"result"`
}

// TransferType indicates that type required for transfer token.
const TransferType string = "TRANSFER"

// TransferNative transfers native token via Venly.
func (client *Client) TransferNative(ctx context.Context, accessToken string, r TransferNativeRequest) (response TransferNativeResponse, err error) {
	jsonBody, err := json.Marshal(r)
	if err != nil {
		return TransferNativeResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, client.config.DefaultURL+"transactions/execute", bytes.NewReader(jsonBody))
	if err != nil {
		return TransferNativeResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.http.Do(req.WithContext(ctx))
	if err != nil {
		return TransferNativeResponse{}, err
	}

	defer func() {
		err = errs.Combine(err, resp.Body.Close())
	}()

	if resp.StatusCode != http.StatusOK {
		errorResp := ErrorResponse{}

		if err = json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
			return TransferNativeResponse{}, err
		}

		if !errorResp.Success {
			return TransferNativeResponse{}, errs.New(errorResp.Errors[0].Message)
		}

		return TransferNativeResponse{}, errs.New(resp.Status)
	}

	var transferNativeResponse TransferNativeResponse
	if err = json.NewDecoder(resp.Body).Decode(&transferNativeResponse); err != nil {
		return TransferNativeResponse{}, err
	}

	return transferNativeResponse, nil
}
