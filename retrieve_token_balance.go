// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package venly

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/zeebo/errs"
)

// RetrieveTokenBalanceRequest fields that required for retrieve token balance request.
type RetrieveTokenBalanceRequest struct {
	WalletID string `json:"wallet_id"`
}

// RetrieveTokenBalanceResult fields that describes result from retrieve token balance.
type RetrieveTokenBalanceResult struct {
	TokenAddress string  `json:"tokenAddress"`
	RawBalance   string  `json:"rawBalance"`
	Balance      float64 `json:"balance"`
	Decimals     int     `json:"decimals"`
	Symbol       string  `json:"symbol"`
	Logo         string  `json:"logo"`
	Type         string  `json:"type"`
	Transferable bool    `json:"transferable"`
}

// RetrieveTokenBalanceResponse fields that returns from retrieve token balance.
type RetrieveTokenBalanceResponse struct {
	Success bool                         `json:"success"`
	Result  []RetrieveTokenBalanceResult `json:"result"`
}

// RetrieveTokenBalance retrieves Venly token balance.
func (client *Client) RetrieveTokenBalance(ctx context.Context, accessToken string, r RetrieveTokenBalanceRequest) (response RetrieveTokenBalanceResponse, err error) {
	req, err := http.NewRequest(http.MethodGet, client.config.DefaultURL+"wallets/"+r.WalletID+"/balance/tokens", nil)
	if err != nil {
		return RetrieveTokenBalanceResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.http.Do(req.WithContext(ctx))
	if err != nil {
		return RetrieveTokenBalanceResponse{}, err
	}

	defer func() {
		err = errs.Combine(err, resp.Body.Close())
	}()

	if resp.StatusCode != http.StatusOK {
		errorResp := ErrorResponse{}

		if err = json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
			return RetrieveTokenBalanceResponse{}, err
		}

		if !errorResp.Success {
			return RetrieveTokenBalanceResponse{}, errs.New(errorResp.Errors[0].Code)
		}

		return RetrieveTokenBalanceResponse{}, errs.New(resp.Status)
	}

	var retrieveTokenBalanceResponse RetrieveTokenBalanceResponse
	if err = json.NewDecoder(resp.Body).Decode(&retrieveTokenBalanceResponse); err != nil {
		return RetrieveTokenBalanceResponse{}, err
	}

	return retrieveTokenBalanceResponse, nil
}
