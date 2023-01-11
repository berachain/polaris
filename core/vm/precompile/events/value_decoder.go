package events

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/berachain/stargazer/common"
)

const (
	// `intBase` is 10 because `int`s are parsed in base 10.
	intBase = 10

	// `creationHeightBits` is 64 because Cosmos `creationHeight` is stored as a `int64`.
	creationHeightBits = 64
)

// `AttributeValueDecoder` is a type of function that returns a geth compatible, eth primitive type
// (as type `any`) for a given Cosmos event attribute value (of type `string`). Cosmos event
// attribute values may require unique decodings based on their underlying string encoding.
type AttributeValueDecoder func(attributeValue string) (ethPrimitive any, err error)

// `ConvertSdkCoin` converts the string representation of an `sdk.Coin` to a `*big.Int`
//
// `ConvertSdkCoin` is a `AttributeValueDecoder`
func ConvertSdkCoin(attributeValue string) (any, error) {
	// extract the sdk.Coin from string value
	coin, err := sdk.ParseCoinNormalized(attributeValue)
	if err != nil {
		return nil, err
	}
	// convert the sdk.Coin to *big.Int
	return coin.Amount.BigInt(), nil
}

// `ConvertValAddress` converts a bech32 string representing a validator address to a
// `common.Address`
//
// `ConvertValAddress` is a `AttributeValueDecoder`
func ConvertValAddress(attributeValue string) (any, error) {
	// extract the sdk.ValAddress from string value
	valAddress, err := sdk.ValAddressFromBech32(attributeValue)
	if err != nil {
		return nil, err
	}
	// convert the sdk.ValAddress to common.Address
	return common.ValAddressToEthAddress(valAddress), nil
}

// `ConvertAccAddress` converts a bech32 string representing an account address to a
// `common.Address`
//
// `ConvertAccAddress` is a `AttributeValueDecoder`
func ConvertAccAddress(attributeValue string) (any, error) {
	// extract the sdk.AccAddress from string value
	accAddress, err := sdk.AccAddressFromBech32(attributeValue)
	if err != nil {
		return nil, err
	}
	// convert the sdk.AccAddress to common.Address
	return common.AccAddressToEthAddress(accAddress), nil
}

// `ConvertCreationHeight` converts a creationHeight `string` to an `int64`
//
// `ConvertCreationHeight` is a `AttributeValueDecoder`
func ConvertCreationHeight(attributeValue string) (any, error) {
	return strconv.ParseInt(attributeValue, intBase, creationHeightBits)
}
