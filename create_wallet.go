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

// CreateWalletRequest fields that required for create wallet request.
type CreateWalletRequest struct {
	WalletType  string `json:"walletType"`
	SecretType  string `json:"secretType"`
	Identifier  string `json:"identifier"`
	Description string `json:"description"`
	Pincode     string `json:"pincode"`
}

// CreateWalletResponse fields that returns from create wallet.
type CreateWalletResponse struct {
	Success bool `json:"success"`
	Result  struct {
		ID           string `json:"id"`
		Address      string `json:"address"`
		WalletType   string `json:"walletType"`
		SecretType   string `json:"secretType"`
		CreatedAt    string `json:"createdAt"`
		Archived     bool   `json:"archived"`
		Alias        string `json:"alias"`
		Description  string `json:"description"`
		Primary      bool   `json:"primary"`
		HasCustomPin bool   `json:"hasCustomPin"`
		Identifier   string `json:"identifier"`
		Balance      struct {
			Available     bool    `json:"available"`
			SecretType    string  `json:"secretType"`
			Balance       float64 `json:"balance"`
			GasBalance    float64 `json:"gasBalance"`
			Symbol        string  `json:"symbol"`
			GasSymbol     string  `json:"gasSymbol"`
			RawBalance    string  `json:"rawBalance"`
			RawGasBalance string  `json:"rawGasBalance"`
			Decimals      int     `json:"decimals"`
		} `json:"balance"`
	} `json:"result"`
}

// CreateWallet creates Venly wallet.
func (client *Client) CreateWallet(ctx context.Context, accessToken string, wallet CreateWalletRequest) (response CreateWalletResponse, err error) {
	jsonBody, err := json.Marshal(wallet)
	if err != nil {
		return CreateWalletResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, client.config.DefaultURL+"wallets", bytes.NewReader(jsonBody))
	if err != nil {
		return CreateWalletResponse{}, err
	}

	var bearer = "Bearer " + accessToken

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	resp, err := client.http.Do(req.WithContext(ctx))
	if err != nil {
		return CreateWalletResponse{}, err
	}

	defer func() {
		err = errs.Combine(err, resp.Body.Close())
	}()

	if resp.StatusCode != http.StatusOK {
		errorResp := ErrorResponse{}

		if err = json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
			return CreateWalletResponse{}, err
		}

		if !errorResp.Success {
			return CreateWalletResponse{}, errs.New(errorResp.Errors[0].Code)
		}

		return CreateWalletResponse{}, errs.New(resp.Status)
	}

	var createWalletResponse CreateWalletResponse
	if err = json.NewDecoder(resp.Body).Decode(&createWalletResponse); err != nil {
		return CreateWalletResponse{}, err
	}

	return createWalletResponse, nil
}
