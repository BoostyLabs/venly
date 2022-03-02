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

// ResubmitTXRequest fields that required for resubmit tx request.
type ResubmitTXRequest struct {
	SecretType      string `json:"secretType"`
	TransactionHash string `json:"transactionHash"`
	Pincode         string `json:"pincode"`
}

// ResubmitTXResponse fields that returns from resubmit tx.
type ResubmitTXResponse struct {
	Success bool `json:"success"`
	Result  struct {
		TransactionHash string `json:"transactionHash"`
	} `json:"result"`
}

// ResubmitTX resubmits transaction.
func (client *Client) ResubmitTX(ctx context.Context, accessToken string, r ResubmitTXRequest) (response ResubmitTXResponse, err error) {
	jsonBody, err := json.Marshal(r)
	if err != nil {
		return ResubmitTXResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, client.config.DefaultURL+"transactions/resubmit", bytes.NewReader(jsonBody))
	if err != nil {
		return ResubmitTXResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.http.Do(req.WithContext(ctx))
	if err != nil {
		return ResubmitTXResponse{}, err
	}

	defer func() {
		err = errs.Combine(err, resp.Body.Close())
	}()

	if resp.StatusCode != http.StatusOK {
		errorResp := ErrorResponse{}

		if err = json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
			return ResubmitTXResponse{}, err
		}

		if !errorResp.Success {
			return ResubmitTXResponse{}, errs.New(errorResp.Errors[0].Code)
		}

		return ResubmitTXResponse{}, errs.New(resp.Status)
	}

	var resubmitTXResponse ResubmitTXResponse
	if err = json.NewDecoder(resp.Body).Decode(&resubmitTXResponse); err != nil {
		return ResubmitTXResponse{}, err
	}

	return resubmitTXResponse, nil
}
