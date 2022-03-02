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

// ReadContractRequest fields that required for read contract request.
type ReadContractRequest struct {
	SecretType      string `json:"secretType"`
	WalletAddress   string `json:"walletAddress"`
	ContractAddress string `json:"contractAddress"`
	FunctionName    string `json:"functionName"`
	Inputs          []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"inputs"`
	Outputs []struct {
		Type string `json:"type"`
	} `json:"outputs"`
}

// ReadContractResponse struct contains response from read contract.
type ReadContractResponse struct {
	Success bool `json:"success"`
	Result  []struct {
		Type  string `json:"type"`
		Value int    `json:"value"`
	} `json:"result"`
}

// ReadContract reads contract.
func (client *Client) ReadContract(ctx context.Context, accessToken string, r ReadContractRequest) (response ReadContractResponse, err error) {
	jsonBody, err := json.Marshal(r)
	if err != nil {
		return ReadContractResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, client.config.DefaultURL+"contracts/read", bytes.NewReader(jsonBody))
	if err != nil {
		return ReadContractResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.http.Do(req.WithContext(ctx))
	if err != nil {
		return ReadContractResponse{}, err
	}

	defer func() {
		err = errs.Combine(err, resp.Body.Close())
	}()

	if resp.StatusCode != http.StatusOK {
		errorResp := ErrorResponse{}

		if err = json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
			return ReadContractResponse{}, err
		}

		if !errorResp.Success {
			return ReadContractResponse{}, errs.New(errorResp.Errors[0].Code)
		}

		return ReadContractResponse{}, errs.New(resp.Status)
	}

	var readContractResponse ReadContractResponse
	if err = json.NewDecoder(resp.Body).Decode(&readContractResponse); err != nil {
		return ReadContractResponse{}, err
	}

	return readContractResponse, nil
}
