package venly

import (
	"context"
)

type Venly interface {
	// CreateWallet mocks method to create wallet.
	CreateWallet(ctx context.Context, accessToken string, wallet CreateWalletRequest) (CreateWalletResponse, error)
	// RetrieveWallet mocks method to retrieve wallet.
	RetrieveWallet(ctx context.Context, accessToken string, r RetrieveWalletRequest) (RetrieveWalletResponse, error)
	// RetrieveWalletBalance mocks method to retrieve wallet balance.
	RetrieveWalletBalance(ctx context.Context, accessToken string, r RetrieveWalletRequest) (RetrieveWalletBalanceResponse, error)
	// RetrieveTokenBalance mocks method to retrieve token balance.
	RetrieveTokenBalance(ctx context.Context, accessToken string, r RetrieveTokenBalanceRequest) (RetrieveTokenBalanceResponse, error)
	// TransferNative mocks method to transfer native tokens.
	TransferNative(ctx context.Context, accessToken string, r TransferNativeRequest) (TransferNativeResponse, error)
	// TransferFungible mocks method to transfer fungible tokens.
	TransferFungible(ctx context.Context, accessToken string, r TransferFungibleRequest) (TransferFungibleResponse, error)
	// TransferNonFungible transfers non-fungible token via Venly.
	TransferNonFungible(ctx context.Context, accessToken string, r TransferNonFungibleRequest) (TransferNonFungibleResponse, error)
	// ExecuteContract executes contract call.
	ExecuteContract(ctx context.Context, accessToken string, r ExecuteContractRequest) (response ExecuteContractResponse, err error)
	// GetFees retrieves fees.
	GetFees(ctx context.Context, accessToken string, r GetFeesRequest) (response GetFeesResponse, err error)
	// GetTXStatus retrieves tx status.
	GetTXStatus(ctx context.Context, accessToken string, r GetTXStatusRequest) (response GetTXStatusResponse, err error)
	// ReadContract reads contract.
	ReadContract(ctx context.Context, accessToken string, r ReadContractRequest) (response ReadContractResponse, err error)
	// ResubmitTX resubmits transaction.
	ResubmitTX(ctx context.Context, accessToken string, r ResubmitTXRequest) (response ResubmitTXResponse, err error)
	// Signatures calls signatures Venly api endpoint.
	Signatures(ctx context.Context, accessToken string, r SignaturesRequest) (response SignaturesResponse, err error)
	// RetrieveBalanceByWalletAddress retrieves balance by wallet address and secret type.
	RetrieveBalanceByWalletAddress(ctx context.Context, accessToken string, r BalanceRequest) (response RetrieveWalletBalanceResponse, err error)
	// RetrieveWallets retrieves Venly wallets.
	RetrieveWallets(ctx context.Context, accessToken string) (response RetrieveWalletResponse, err error)
}
