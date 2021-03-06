// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package mock

import (
	"context"

	"github.com/BoostyLabs/venly"
)

var _ venly.Venly = (*Mock)(nil)

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
		Result: []venly.RetrieveTokenBalanceResult{{
			TokenAddress: "",
			RawBalance:   "",
			Balance:      0,
			Decimals:     0,
			Symbol:       "WETH",
			Logo:         "",
			Type:         "",
			Transferable: true,
		}},
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

// ExecuteContract executes contract call.
func (mock *Mock) ExecuteContract(ctx context.Context, accessToken string, r venly.ExecuteContractRequest) (response venly.ExecuteContractResponse, err error) {
	return venly.ExecuteContractResponse{
		TransactionHash: "true",
	}, nil
}

// GetFees retrieves fees.
func (mock *Mock) GetFees(ctx context.Context, accessToken string, r venly.GetFeesRequest) (response venly.GetFeesResponse, err error) {
	return venly.GetFeesResponse{
		Success: true,
	}, nil
}

// GetTXStatus retrieves tx status.
func (mock *Mock) GetTXStatus(ctx context.Context, accessToken string, r venly.GetTXStatusRequest) (response venly.GetTXStatusResponse, err error) {
	return venly.GetTXStatusResponse{
		Success: true,
	}, nil
}

// ReadContract reads contract.
func (mock *Mock) ReadContract(ctx context.Context, accessToken string, r venly.ReadContractRequest) (response venly.ReadContractResponse, err error) {
	return venly.ReadContractResponse{
		Success: true,
	}, nil
}

// ResubmitTX resubmits transaction.
func (mock *Mock) ResubmitTX(ctx context.Context, accessToken string, r venly.ResubmitTXRequest) (response venly.ResubmitTXResponse, err error) {
	return venly.ResubmitTXResponse{
		Success: true,
	}, nil
}

// Signatures calls signatures Venly api endpoint.
func (mock *Mock) Signatures(ctx context.Context, accessToken string, r venly.SignaturesRequest) (response venly.SignaturesResponse, err error) {
	return venly.SignaturesResponse{
		Success: true,
	}, nil
}

// TransferNonFungible transfers non-fungible token via Venly.
func (mock *Mock) TransferNonFungible(ctx context.Context, accessToken string, r venly.TransferNonFungibleRequest) (response venly.TransferNonFungibleResponse, err error) {
	return venly.TransferNonFungibleResponse{
		Success: true,
	}, nil
}

// RetrieveBalanceByWalletAddress mocks method to retrieve balance by wallet address and secret type.
func (mock *Mock) RetrieveBalanceByWalletAddress(ctx context.Context, accessToken string, r venly.BalanceRequest) (response venly.RetrieveWalletBalanceResponse, err error) {
	return venly.RetrieveWalletBalanceResponse{
		Success: true,
	}, nil
}

// RetrieveWallets mocks method to retrieves Venly wallets.
func (mock *Mock) RetrieveWallets(ctx context.Context, accessToken string) (response venly.RetrieveWalletResponse, err error) {
	return venly.RetrieveWalletResponse{
		Success: true,
	}, nil
}
