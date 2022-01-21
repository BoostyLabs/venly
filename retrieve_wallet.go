package venly

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/zeebo/errs"
)

// RetrieveWalletRequest fields that required for retrieve wallet request.
type RetrieveWalletRequest struct {
	WalletID string `json:"wallet_id"`
}

// RetrieveWalletResponse fields that returns from retrieve wallet.
type RetrieveWalletResponse struct {
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

// RetrieveWallet retrieves Venly wallet.
func (client *Client) RetrieveWallet(ctx context.Context, accessToken string, r RetrieveWalletRequest) (response RetrieveWalletResponse, err error) {
	req, err := http.NewRequest(http.MethodGet, client.config.DefaultURL+"wallets/"+r.WalletID, nil)
	if err != nil {
		return RetrieveWalletResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.http.Do(req.WithContext(ctx))
	if err != nil {
		return RetrieveWalletResponse{}, err
	}

	defer func() {
		err = errs.Combine(err, resp.Body.Close())
	}()

	if resp.StatusCode != http.StatusOK {
		errorResp := ErrorResponse{}

		if err = json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
			return RetrieveWalletResponse{}, err
		}

		return RetrieveWalletResponse{}, errs.New(errorResp.Errors[0].Code)
	}

	var retrieveWalletResponse RetrieveWalletResponse
	if err = json.NewDecoder(resp.Body).Decode(&retrieveWalletResponse); err != nil {
		return RetrieveWalletResponse{}, err
	}

	return retrieveWalletResponse, nil
}
