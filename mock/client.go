// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package mock

import (
	"context"

	"github.com/BoostyLabs/venly"
)

// Mock mocked struct for venly wallet.
type Mock struct{}

// CreateWallet mocks method to create wallet.
func (mock *Mock) CreateWallet(ctx context.Context, accessToken string, wallet venly.CreateWalletRequest) (venly.CreateWalletResponse, error) {
	return venly.CreateWalletResponse{
		Success: true,
	}, nil
}

// RetrieveWallet mocks method to retrieve wallet.
func (mock *Mock) RetrieveWallet(ctx context.Context, accessToken string, r venly.RetrieveWalletRequest) (venly.RetrieveWalletResponse, error) {
	return venly.RetrieveWalletResponse{
		Success: true,
	}, nil
}

// RetrieveWalletBalance mocks method to retrieve wallet balance.
func (mock *Mock) RetrieveWalletBalance(ctx context.Context, accessToken string, r venly.RetrieveWalletRequest) (venly.RetrieveWalletBalanceResponse, error) {
	return venly.RetrieveWalletBalanceResponse{
		Success: true,
	}, nil
}

// RetrieveTokenBalance mocks method to retrieve token balance.
func (mock *Mock) RetrieveTokenBalance(ctx context.Context, accessToken string, r venly.RetrieveTokenBalanceRequest) (venly.RetrieveTokenBalanceResponse, error) {
	return venly.RetrieveTokenBalanceResponse{
		Success: true,
	}, nil
}

// TransferNative mocks method to transfer native tokens.
func (mock *Mock) TransferNative(ctx context.Context, accessToken string, r venly.TransferNativeRequest) (venly.TransferNativeResponse, error) {
	return venly.TransferNativeResponse{
		Success: true,
	}, nil
}

// TransferFungible mocks method to transfer fungible tokens.
func (mock *Mock) TransferFungible(ctx context.Context, accessToken string, r venly.TransferFungibleRequest) (venly.TransferFungibleResponse, error) {
	return venly.TransferFungibleResponse{
		Success: true,
	}, nil
}
