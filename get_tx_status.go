// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package venly

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/zeebo/errs"
)

// GetTXStatusRequest fields that required for get tx status request.
type GetTXStatusRequest struct {
	SecretType string
	TXHash     string
}

// GetTXStatusResponse fields that returns from get tx status.
type GetTXStatusResponse struct {
	Success bool `json:"success"`
	Result  struct {
		Hash          string `json:"hash"`
		Status        string `json:"status"`
		Confirmations int    `json:"confirmations"`
		BlockHash     string `json:"blockHash"`
		BlockNumber   int    `json:"blockNumber"`
		Nonce         int    `json:"nonce"`
		Gas           int    `json:"gas"`
		GasUsed       int    `json:"gasUsed"`
		GasPrice      int64  `json:"gasPrice"`
		Logs          []struct {
			LogIndex int         `json:"logIndex"`
			Data     string      `json:"data"`
			Type     interface{} `json:"type"`
			Topics   []string    `json:"topics"`
		} `json:"logs"`
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"result"`
}

// GetTXStatus retrieves tx status.
func (client *Client) GetTXStatus(ctx context.Context, accessToken string, r GetTXStatusRequest) (response GetTXStatusResponse, err error) {
	req, err := http.NewRequest(http.MethodGet, client.config.DefaultURL+"transactions/"+r.SecretType+"/"+r.TXHash+"/status", nil)
	if err != nil {
		return GetTXStatusResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.http.Do(req.WithContext(ctx))
	if err != nil {
		return GetTXStatusResponse{}, err
	}

	defer func() {
		err = errs.Combine(err, resp.Body.Close())
	}()

	if resp.StatusCode != http.StatusOK {
		errorResp := ErrorResponse{}

		if err = json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
			return GetTXStatusResponse{}, err
		}

		return GetTXStatusResponse{}, errs.New(errorResp.Errors[0].Code)
	}

	var getTXStatusResponse GetTXStatusResponse
	if err = json.NewDecoder(resp.Body).Decode(&getTXStatusResponse); err != nil {
		return GetTXStatusResponse{}, err
	}

	return getTXStatusResponse, nil
}
