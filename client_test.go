package venly_test

import (
	"context"
	"testing"

	"github.com/BoostyLabs/venly"
)

func TestNew(t *testing.T) {
	client := venly.NewClient(venly.Config{
		DefaultURL: "https://api-staging.arkane.network/api/",
		AuthURL:    "https://login-staging.arkane.network/auth/realms/Arkane/protocol/openid-connect/token",
	})

	resp, err := client.Auth(context.Background(), venly.AuthRequest{
		ClientID:     "Testaccount-capsule",
		ClientSecret: "82c19251-1753-44f5-ae76-93438d3628de",
		GrantType:    "client_credentials",
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.CreateWallet(context.Background(), resp.AccessToken, venly.CreateWalletRequest{
		WalletType:  "WHITE_LABEL",
		SecretType:  "ETHEREUM",
		Identifier:  "type=unrecoverable",
		Description: "wtf123",
		Pincode:     "1488",
	})
	if err != nil {
		t.Fatal(err)
	}
}
