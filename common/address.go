package common

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

// `AccAddressToEthAddress` converts a Cosmos SDK `AccAddress` to an Ethereum `Address`.
func AccAddressToEthAddress(accAddress sdk.AccAddress) Address {
	return common.BytesToAddress(accAddress)
}

// `EthAddressToAccAddress` converts an Ethereum `Address` to a Cosmos SDK `AccAddress`.
func EthAddressToAccAddress(ethAddress Address) sdk.AccAddress {
	return ethAddress.Bytes()
}
