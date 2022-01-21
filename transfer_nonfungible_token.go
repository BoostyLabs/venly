package venly

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/zeebo/errs"
)

// TransferNonFungibleRequest fields that required for transfer non-fungible request.
type TransferNonFungibleRequest struct {
	TransactionRequest struct {
		Type         string      `json:"type"`
		WalletID     string      `json:"walletId"`
		To           string      `json:"to"`
		SecretType   string      `json:"secretType"`
		TokenAddress string      `json:"tokenAddress"`
		TokenID      interface{} `json:"tokenId"`
	} `json:"transactionRequest"`
	Pincode string `json:"pincode"`
}

// TransferNonFungibleResponse fields that returns from transfer non-fungible.
type TransferNonFungibleResponse struct {
	Success bool `json:"success"`
	Result  struct {
		TransactionHash string `json:"transactionHash"`
	} `json:"result"`
}

// TransferNonFungible transfers non-fungible token via Venly.
func (client *Client) TransferNonFungible(ctx context.Context, accessToken string, r TransferNonFungibleRequest) (response TransferNonFungibleResponse, err error) {
	jsonBody, err := json.Marshal(r)
	if err != nil {
		return TransferNonFungibleResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, client.config.DefaultURL+"transactions/execute", bytes.NewReader(jsonBody))
	if err != nil {
		return TransferNonFungibleResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.http.Do(req.WithContext(ctx))
	if err != nil {
		return TransferNonFungibleResponse{}, err
	}

	defer func() {
		err = errs.Combine(err, resp.Body.Close())
	}()

	if resp.StatusCode != http.StatusOK {
		errorResp := ErrorResponse{}

		if err = json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
			return TransferNonFungibleResponse{}, err
		}

		return TransferNonFungibleResponse{}, errs.New(errorResp.Errors[0].Code)
	}

	var transferNonFungibleResponse TransferNonFungibleResponse
	if err = json.NewDecoder(resp.Body).Decode(&transferNonFungibleResponse); err != nil {
		return TransferNonFungibleResponse{}, err
	}

	return transferNonFungibleResponse, nil
}
