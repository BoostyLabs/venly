package venly_test

import (
	"context"
	"testing"

	"amuzed/internal/venly"
)

func TestVenly(t *testing.T) {
	client := venly.NewClient(venly.Config{
		VenlyDefaultURL: "https://api-staging.arkane.network/api/",
		VenlyAuthURL:    "https://login-staging.arkane.network/auth/realms/Arkane/protocol/openid-connect/token",
	})

	auth, err := client.Auth(context.Background(), venly.AuthRequest{
		GrantType:    "client_credentials",
		ClientID:     "Testaccount-capsule",
		ClientSecret: "82c19251-1753-44f5-ae76-93438d3628de",
	})
	if err != nil {
		t.Fatal(err)
	}

	wallet, err := client.CreateWallet(context.Background(), auth.AccessToken, venly.CreateWalletRequest{
		WalletType:  "WHITE_LABEL",
		SecretType:  "ETHEREUM",
		Identifier:  "test_id4",
		Description: "my new test wallet4",
		Pincode:     "1488",
	})
	if err != nil {
		t.Fatal(err)
	}

	retrieveWallet, err := client.RetrieveWallet(context.Background(), auth.AccessToken, venly.RetrieveWalletRequest{WalletID: wallet.Result.ID})
	if err != nil {
		t.Fatal(err)
	}

	walletBalance, err := client.RetrieveWalletBalance(context.Background(), auth.AccessToken, venly.RetrieveWalletRequest{WalletID: retrieveWallet.Result.ID})
	if err != nil {
		t.Fatal(err)
	}
	print(walletBalance.Result.RawBalance)

	walletTokenBalance, err := client.RetrieveTokenBalance(context.Background(), auth.AccessToken, venly.RetrieveTokenBalanceRequest{WalletID: wallet.Result.ID})
	if err != nil {
		t.Fatal(err)
	}
	print(walletTokenBalance.Result[0].RawBalance)
}
