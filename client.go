// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package venly

import (
	"net/http"
)

var _ Venly = (*Client)(nil)

// Config is the global configuration for Venly client.
type Config struct {
	DefaultURL string `json:"venlyDefaultURL" default:"https://api.arkane.network/api/"`
	AuthURL    string `json:"venlyAuthURL" default:"https://login-staging.arkane.network/auth/realms/Arkane/protocol/openid-connect/token"`
}

// Client implementation of the Venly API.
type Client struct {
	http   http.Client
	config Config
}

// NewClient is a constructor for Venly API client.
func NewClient(config Config) *Client {
	return &Client{config: config}
}

// ErrorResponse holds fields that explaining error from Venly API.
type ErrorResponse struct {
	Success bool `json:"success"`
	Errors  []struct {
		Code      string      `json:"code"`
		TraceCode interface{} `json:"traceCode"`
		Message   interface{} `json:"message"`
	} `json:"errors"`
}
