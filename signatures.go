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

// SignaturesRequest fields that required for signatures request.
type SignaturesRequest struct {
	Pincode          string `json:"pincode"`
	SignatureRequest struct {
		Type       string `json:"type"`
		SecretType string `json:"secretType"`
		WalletID   string `json:"walletId"`
		Data       string `json:"data"`
	} `json:"signatureRequest"`
}

// SignaturesResponse fields that returns from signatures.
type SignaturesResponse struct {
	Type      string `json:"type"`
	R         string `json:"r"`
	S         string `json:"s"`
	V         string `json:"v"`
	Signature string `json:"signature"`
}

// Signatures calls signatures Venly api endpoint.
func (client *Client) Signatures(ctx context.Context, accessToken string, r SignaturesRequest) (response SignaturesResponse, err error) {
	jsonBody, err := json.Marshal(r)
	if err != nil {
		return SignaturesResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, client.config.DefaultURL+"signatures", bytes.NewReader(jsonBody))
	if err != nil {
		return SignaturesResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.http.Do(req.WithContext(ctx))
	if err != nil {
		return SignaturesResponse{}, err
	}

	defer func() {
		err = errs.Combine(err, resp.Body.Close())
	}()

	if resp.StatusCode != http.StatusOK {
		errorResp := ErrorResponse{}

		if err = json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
			return SignaturesResponse{}, err
		}

		return SignaturesResponse{}, errs.New(errorResp.Errors[0].Code)
	}

	var signaturesResponse SignaturesResponse
	if err = json.NewDecoder(resp.Body).Decode(&signaturesResponse); err != nil {
		return SignaturesResponse{}, err
	}

	return signaturesResponse, nil
}
