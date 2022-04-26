// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package venly

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zeebo/errs"
)

// RetrieveWalletBalanceRequest fields that required for retrieve wallet balance request.
type RetrieveWalletBalanceRequest struct {
	WalletID string `json:"wallet_id"`
}

// BalanceRequest fields that required for retrieve balance by wallet address and secret type.
type BalanceRequest struct {
	SecretType    string         `json:"secretType"`
	WalletAddress common.Address `json:"walletAddress"`
}

// RetrieveWalletBalanceResponse fields that returns from retrieve wallet balance.
type RetrieveWalletBalanceResponse struct {
	Success bool `json:"success"`
	Result  struct {
		Available     bool    `json:"available"`
		SecretType    string  `json:"secretType"`
		Balance       float64 `json:"balance"`
		GasBalance    float64 `json:"gasBalance"`
		Symbol        string  `json:"symbol"`
		GasSymbol     string  `json:"gasSymbol"`
		RawBalance    string  `json:"rawBalance"`
		RawGasBalance string  `json:"rawGasBalance"`
		Decimals      int     `json:"decimals"`
	} `json:"result"`
}

// RetrieveWalletBalance retrieves Venly wallet balance.
func (client *Client) RetrieveWalletBalance(ctx context.Context, accessToken string, r RetrieveWalletRequest) (response RetrieveWalletBalanceResponse, err error) {
	req, err := http.NewRequest(http.MethodGet, client.config.DefaultURL+"wallets/"+r.WalletID+"/balance", nil)
	if err != nil {
		return RetrieveWalletBalanceResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.http.Do(req.WithContext(ctx))
	if err != nil {
		return RetrieveWalletBalanceResponse{}, err
	}

	defer func() {
		err = errs.Combine(err, resp.Body.Close())
	}()

	if resp.StatusCode != http.StatusOK {
		errorResp := ErrorResponse{}

		if err = json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
			return RetrieveWalletBalanceResponse{}, err
		}

		if !errorResp.Success {
			return RetrieveWalletBalanceResponse{}, errs.New(errorResp.Errors[0].Code)
		}

		return RetrieveWalletBalanceResponse{}, errs.New(resp.Status)
	}

	var retrieveWalletBalanceResponse RetrieveWalletBalanceResponse
	if err = json.NewDecoder(resp.Body).Decode(&retrieveWalletBalanceResponse); err != nil {
		return RetrieveWalletBalanceResponse{}, err
	}

	return retrieveWalletBalanceResponse, nil
}

// RetrieveBalanceByWalletAddress retrieves balance by wallet address and secret type.
func (client *Client) RetrieveBalanceByWalletAddress(ctx context.Context, accessToken string, r BalanceRequest) (response RetrieveWalletBalanceResponse, err error) {
	req, err := http.NewRequest(http.MethodGet, client.config.DefaultURL+"wallets/"+r.SecretType+"/"+r.WalletAddress.Hex()+"/balance", nil)
	if err != nil {
		return RetrieveWalletBalanceResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.http.Do(req.WithContext(ctx))
	if err != nil {
		return RetrieveWalletBalanceResponse{}, err
	}

	defer func() {
		err = errs.Combine(err, resp.Body.Close())
	}()

	if resp.StatusCode != http.StatusOK {
		errorResp := ErrorResponse{}

		if err = json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
			return RetrieveWalletBalanceResponse{}, err
		}

		if !errorResp.Success {
			return RetrieveWalletBalanceResponse{}, errs.New(errorResp.Errors[0].Code)
		}

		return RetrieveWalletBalanceResponse{}, errs.New(resp.Status)
	}

	var retrieveWalletBalanceResponse RetrieveWalletBalanceResponse
	if err = json.NewDecoder(resp.Body).Decode(&retrieveWalletBalanceResponse); err != nil {
		return RetrieveWalletBalanceResponse{}, err
	}

	return retrieveWalletBalanceResponse, nil
}
