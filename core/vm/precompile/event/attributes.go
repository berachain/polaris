// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package event

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/berachain/stargazer/common"
	"github.com/berachain/stargazer/types/abi"
)

const (
	// `intBase` is the base `int`s are parsed in, 10.
	intBase = 10

	// `creationHeightBits` is the number of bits used to store `creationHeight` from the Cosmos
	// SDK staking module, which is of type `int64`.
	creationHeightBits = 64

	// `notFound` is a default return value for searches in which an item was not found.
	notFound = -1
)

type (
	// `valueDecoder` is a type of function that returns a geth compatible, eth primitive type (as
	// type `any`) for a given Cosmos event attribute value (of type `string`). Cosmos event
	// attribute values may require unique decodings based on their underlying string encoding.
	valueDecoder func(attributeValue string) (ethPrimitive any, err error)

	// `ValueDecoders` is a type that represents a map of Cosmos event attribute keys to value
	// decoder functions.
	ValueDecoders map[string]valueDecoder
)

// ==============================================================================
// Default Attribute Value Decoder Getter
// ==============================================================================

// `getDefaultCosmosValueDecoder` returns a default Cosmos event attribute value decoder function
// for a certain Cosmos event `attributeKey`. NOTE: only the event attributes of default Cosmos SDK
// modules are supported by this function.
func getDefaultCosmosValueDecoder(attributeKey string) valueDecoder {
	return ValueDecoders{
		sdk.AttributeKeyAmount:                  convertSdkCoin,
		stakingtypes.AttributeKeyValidator:      convertValAddressFromBech32,
		stakingtypes.AttributeKeySrcValidator:   convertValAddressFromBech32,
		stakingtypes.AttributeKeyDstValidator:   convertValAddressFromBech32,
		stakingtypes.AttributeKeyCreationHeight: convertCreationHeight,
		stakingtypes.AttributeKeyDelegator:      convertAccAddressFromBech32,
		banktypes.AttributeKeySender:            convertAccAddressFromBech32,
		banktypes.AttributeKeyRecipient:         convertAccAddressFromBech32,
		banktypes.AttributeKeySpender:           convertAccAddressFromBech32,
		banktypes.AttributeKeyReceiver:          convertAccAddressFromBech32,
		banktypes.AttributeKeyMinter:            convertAccAddressFromBech32,
		banktypes.AttributeKeyBurner:            convertAccAddressFromBech32,
	}[attributeKey]
}

// ==============================================================================
// Default Attribute Value Decoder Functions
// ==============================================================================

// `convertSdkCoin` converts the string representation of an `sdk.Coin` to a `*big.Int`.
//
// `convertSdkCoin` is a `valueDecoder`.
func convertSdkCoin(attributeValue string) (any, error) {
	// extract the sdk.Coin from string value
	coin, err := sdk.ParseCoinNormalized(attributeValue)
	if err != nil {
		return nil, err
	}
	// convert the sdk.Coin to *big.Int
	return coin.Amount.BigInt(), nil
}

// `convertValAddressFromBech32` converts a bech32 string representing a validator address to a
// `common.Address`.
//
// `convertValAddressFromBech32` is a `valueDecoder`.
func convertValAddressFromBech32(attributeValue string) (any, error) {
	// extract the sdk.ValAddress from string value
	valAddress, err := sdk.ValAddressFromBech32(attributeValue)
	if err != nil {
		return nil, err
	}
	// convert the sdk.ValAddress to common.Address
	return common.ValAddressToEthAddress(valAddress), nil
}

// `convertAccAddressFromBech32` converts a bech32 string representing an account address to a
// `common.Address`.
//
// `convertAccAddressFromBech32` is a `valueDecoder`.
func convertAccAddressFromBech32(attributeValue string) (any, error) {
	// extract the sdk.AccAddress from string value
	accAddress, err := sdk.AccAddressFromBech32(attributeValue)
	if err != nil {
		return nil, err
	}
	// convert the sdk.AccAddress to common.Address
	return common.AccAddressToEthAddress(accAddress), nil
}

// `convertCreationHeight` converts a creation height (from the Cosmos SDK staking module) `string`
// to an `int64`.
//
// `convertCreationHeight` is a `valueDecoder`.
func convertCreationHeight(attributeValue string) (any, error) {
	return strconv.ParseInt(attributeValue, intBase, creationHeightBits)
}

// ==============================================================================
// Helpers
// ==============================================================================

// `searchAttributesForArg` searches through the given slice `attributes` for any attribute having
// a key that matches `argName`. This function returns the index where `argName` was found or -1 if
// `argName` was not found.
func searchAttributesForArg(attributes *[]abci.EventAttribute, argName string) int {
	for i, attribute := range *attributes {
		if abi.ToMixedCase(attribute.Key) == argName {
			return i
		}
	}
	// reached end of loop, `argName` not found
	return notFound
}
