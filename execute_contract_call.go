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

// ExecuteContractRequest fields that required for execute contract call request.
type ExecuteContractRequest struct {
	Pincode            string             `json:"pincode"`
	TransactionRequest TransactionRequest `json:"transactionRequest"`
}

// TransactionRequest struct contains in request.
type TransactionRequest struct {
	Type                string              `json:"type"`
	WalletID            string              `json:"walletId"`
	To                  string              `json:"to"`
	Alias               interface{}         `json:"alias"`
	SecretType          string              `json:"secretType"`
	FunctionName        string              `json:"functionName"`
	Value               int                 `json:"value"`
	Inputs              []Inputs            `json:"inputs"`
	ChainSpecificFields ChainSpecificFields `json:"chainSpecificFields"`
}

// ChainSpecificFields sstruct contains in TransactionRequest.
type ChainSpecificFields struct {
	GasPrice string `json:"gasPrice"`
	GasLimit string `json:"gasLimit"`
}

// Inputs struct contains in request.
type Inputs struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

// ExecuteContractResponse fields that returns from execute contract call.
type ExecuteContractResponse struct {
	TransactionHash string `json:"transactionHash"`
}

// ExecuteContract executes contract call.
func (client *Client) ExecuteContract(ctx context.Context, accessToken string, r ExecuteContractRequest) (response ExecuteContractResponse, err error) {
	jsonBody, err := json.Marshal(r)
	if err != nil {
		return ExecuteContractResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, client.config.DefaultURL+"transactions/execute", bytes.NewReader(jsonBody))
	if err != nil {
		return ExecuteContractResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.http.Do(req.WithContext(ctx))
	if err != nil {
		return ExecuteContractResponse{}, err
	}

	defer func() {
		err = errs.Combine(err, resp.Body.Close())
	}()

	if resp.StatusCode != http.StatusOK {
		errorResp := ErrorResponse{}

		if err = json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
			return ExecuteContractResponse{}, err
		}

		return ExecuteContractResponse{}, errs.New(errorResp.Errors[0].Code)
	}

	var executeContractResponse ExecuteContractResponse
	if err = json.NewDecoder(resp.Body).Decode(&executeContractResponse); err != nil {
		return ExecuteContractResponse{}, err
	}

	return executeContractResponse, nil
}
