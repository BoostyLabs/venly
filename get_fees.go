package venly

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/zeebo/errs"
)

// GetFeesRequest fields that required for get fees request.
type GetFeesRequest struct {
	SecretType string
}

// GetFeesResponse fields that returns from get fees.
type GetFeesResponse struct {
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

// GetFees retrieves fees.
func (client *Client) GetFees(ctx context.Context, accessToken string, r GetFeesRequest) (response GetFeesResponse, err error) {
	req, err := http.NewRequest(http.MethodGet, client.config.DefaultURL+"transactions/"+r.SecretType+"/fees", nil)
	if err != nil {
		return GetFeesResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.http.Do(req.WithContext(ctx))
	if err != nil {
		return GetFeesResponse{}, err
	}

	defer func() {
		err = errs.Combine(err, resp.Body.Close())
	}()

	if resp.StatusCode != http.StatusOK {
		errorResp := ErrorResponse{}

		if err = json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
			return GetFeesResponse{}, err
		}

		return GetFeesResponse{}, errs.New(errorResp.Errors[0].Code)
	}

	var getFeesResponse GetFeesResponse
	if err = json.NewDecoder(resp.Body).Decode(&getFeesResponse); err != nil {
		return GetFeesResponse{}, err
	}

	return getFeesResponse, nil
}
