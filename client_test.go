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
		ClientID:     "Amuzed-capsule",
		ClientSecret: "1e5fe27a-c0c7-4b59-856c-5dd519723e20",
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
		Pincode: "14328",
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

	dd, err := client.RetrieveWalletBalance(context.Background(), resp.AccessToken, venly.RetrieveWalletRequest{WalletID: "aed10c77-a7a5-4260-b865-7ade0c63a300"})
	if err != nil {
		t.Fatal(err)
	}
	println(dd.Result.Balance)

	tr, err := client.TransferNative(context.Background(), resp.AccessToken, venly.TransferNativeRequest{
		TransactionRequest: venly.TxRequest{
			Type:       "TRANSFER",
			WalletID:   "aed10c77-a7a5-4260-b865-7ade0c63a300",
			To:         "0x2346b33F2E379dDA22b2563B009382a0Fc9aA926",
			SecretType: "ETHEREUM",
			Value:      0.1,
			Data:       "test",
		},
		Pincode: "1489",
	})
	if err != nil {
		t.Fatal(err)
	}
	println(tr.Result.TransactionHash)

	dd, err = client.RetrieveWalletBalance(context.Background(), resp.AccessToken, venly.RetrieveWalletRequest{WalletID: "aed10c77-a7a5-4260-b865-7ade0c63a300"})
	if err != nil {
		t.Fatal(err)
	}
	println(dd.Result.Balance)
}
