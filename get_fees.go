// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

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
	Result  []struct {
		GasPrice     int64 `json:"gasPrice"`
		DefaultPrice bool  `json:"defaultPrice"`
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
