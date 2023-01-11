package common

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

func AccAddressToEthAddress(accAddress sdk.AccAddress) Address {
	return common.BytesToAddress(accAddress)
}

func EthAddressToAccAddress(ethAddress Address) sdk.AccAddress {
	return ethAddress.Bytes()
}
