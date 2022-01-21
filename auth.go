package venly

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/zeebo/errs"
)

// AuthRequest fields that required for auth request.
type AuthRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

// AuthResponse fields that returns from auth login.
type AuthResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	Scope            string `json:"scope"`
}

// Auth creates Venly auth login.
func (client *Client) Auth(ctx context.Context, auth AuthRequest) (response AuthResponse, err error) {
	data := url.Values{}
	data.Add("client_id", auth.ClientID)
	data.Add("client_secret", auth.ClientSecret)
	data.Add("grant_type", auth.GrantType)
	encodedData := data.Encode()

	req, err := http.NewRequest(http.MethodPost, client.config.AuthURL, strings.NewReader(encodedData))
	if err != nil {
		return AuthResponse{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.http.Do(req.WithContext(ctx))
	if err != nil {
		return AuthResponse{}, err
	}

	defer func() {
		err = errs.Combine(err, resp.Body.Close())
	}()

	if resp.StatusCode != http.StatusOK {
		errorResp := ErrorResponse{}

		if err = json.NewDecoder(resp.Body).Decode(&errorResp); err != nil {
			return AuthResponse{}, err
		}

		return AuthResponse{}, errs.New(errorResp.Errors[0].Code)
	}

	var authResponse AuthResponse
	if err = json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return AuthResponse{}, err
	}

	return authResponse, nil
}
