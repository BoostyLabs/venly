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

	client2 := venly.NewClient(venly.Config{
		DefaultURL: "https://api-staging.arkane.network/api/",
		AuthURL:    "https://login-staging.arkane.network/auth/realms/Arkane/protocol/openid-connect/token",
	})

	rr, err := client2.CreateWallet(context.Background(), resp.AccessToken, venly.CreateWalletRequest{
		WalletType:  "WHITE_LABEL",
		SecretType:  "ETHEREUM",
		Identifier:  "type=unrecoverable",
		Description: "wtf1233",
		Pincode:     "14328",
	})
	if err != nil {
		t.Fatal(err)
	}
	println(rr.Result.ID)

	r2, err := client2.Signatures(context.Background(), resp.AccessToken, venly.SignaturesRequest{
		Pincode:          "14328",
		SignatureRequest: venly.SignatureRequest{
			Type:       "MESSAGE",
			SecretType: "ETHEREUM",
			WalletID:   rr.Result.ID,
			Data:       "I agree with terms and conditions",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	println(r2.Result.Type)
}
