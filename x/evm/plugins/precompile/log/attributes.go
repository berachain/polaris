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

package log

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/berachain/stargazer/eth/types/abi"
	"github.com/berachain/stargazer/x/evm/utils"
)

const (
	// `intBase` is the base `int`s are parsed in, 10.
	intBase = 10
	// `int64Bits` is the number of bits stored in a variabe of `int64` type.
	int64Bits = 64
	// `notFound` is a default return value for searches in which an item was not found.
	notFound = -1
)

type (
	// `valueDecoder` is a type of function that returns a geth compatible, eth primitive type (as
	// type `any`) for a given event attribute value (of type `string`). Event attribute values may
	// require unique decodings based on their underlying string encoding.
	valueDecoder func(attributeValue string) (ethPrimitive any, err error)
	// `ValueDecoders` is a type that represents a map of event attribute keys to value decoder
	// functions.
	ValueDecoders map[string]valueDecoder
)

// ==============================================================================
// Default Attribute Value Decoder Getter
// ==============================================================================

// `defaultCosmosValueDecoders` is a map of default Cosmos event attribute value decoder functions
// for the default Cosmos SDK event `attributeKey`s. NOTE: only the event attributes of default
// Cosmos SDK modules (bank, staking) are supported by this function.
var defaultCosmosValueDecoders = ValueDecoders{
	sdk.AttributeKeyAmount:                  ConvertSdkCoin,
	stakingtypes.AttributeKeyValidator:      ConvertValAddressFromBech32,
	stakingtypes.AttributeKeySrcValidator:   ConvertValAddressFromBech32,
	stakingtypes.AttributeKeyDstValidator:   ConvertValAddressFromBech32,
	stakingtypes.AttributeKeyCreationHeight: ConvertInt64,
	stakingtypes.AttributeKeyDelegator:      ConvertAccAddressFromBech32,
	banktypes.AttributeKeySender:            ConvertAccAddressFromBech32,
	banktypes.AttributeKeyRecipient:         ConvertAccAddressFromBech32,
	banktypes.AttributeKeySpender:           ConvertAccAddressFromBech32,
	banktypes.AttributeKeyReceiver:          ConvertAccAddressFromBech32,
	banktypes.AttributeKeyMinter:            ConvertAccAddressFromBech32,
	banktypes.AttributeKeyBurner:            ConvertAccAddressFromBech32,
}

// ==============================================================================
// Default Attribute Value Decoder Functions
// ==============================================================================

// Compile-time assertions to ensure that the default attribute value decoder functions are
// valueDecoders.
var (
	_ valueDecoder = ConvertSdkCoin
	_ valueDecoder = ConvertValAddressFromBech32
	_ valueDecoder = ConvertAccAddressFromBech32
	_ valueDecoder = ConvertInt64
)

// `ConvertSdkCoin` converts the string representation of an `sdk.Coin` to a `*big.Int`.
//
// `ConvertSdkCoin` is a `valueDecoder`.
func ConvertSdkCoin(attributeValue string) (any, error) {
	// extract the sdk.Coin from string value
	coin, err := sdk.ParseCoinNormalized(attributeValue)
	if err != nil {
		return nil, err
	}
	// convert the sdk.Coin to *big.Int
	return coin.Amount.BigInt(), nil
}

// `ConvertValAddressFromBech32` converts a bech32 string representing a validator address to a
// `common.Address`.
//
// `ConvertValAddressFromBech32` is a `valueDecoder`.
func ConvertValAddressFromBech32(attributeValue string) (any, error) {
	// extract the sdk.ValAddress from string value
	valAddress, err := sdk.ValAddressFromBech32(attributeValue)
	if err != nil {
		return nil, err
	}
	// convert the sdk.ValAddress to common.Address
	return utils.ValAddressToEthAddress(valAddress), nil
}

// `ConvertAccAddressFromBech32` converts a bech32 string representing an account address to a
// `common.Address`.
//
// `ConvertAccAddressFromBech32` is a `valueDecoder`.
func ConvertAccAddressFromBech32(attributeValue string) (any, error) {
	// extract the sdk.AccAddress from string value
	accAddress, err := sdk.AccAddressFromBech32(attributeValue)
	if err != nil {
		return nil, err
	}
	// convert the sdk.AccAddress to common.Address
	return utils.AccAddressToEthAddress(accAddress), nil
}

// `ConvertInt64` converts a creation height (from the Cosmos SDK staking module) `string`
// to an `int64`.
//
// `ConvertInt64` is a `valueDecoder`.
func ConvertInt64(attributeValue string) (any, error) {
	return strconv.ParseInt(attributeValue, intBase, int64Bits)
}

// ==============================================================================
// Helpers
// ==============================================================================

// `validateAttributes` validates an incoming Cosmos `event`. Specifically, it verifies that the
// number of attributes provided by the Cosmos `event` are adequate for it's corresponding
// Ethereum events.
func validateAttributes(pl *precompileLog, event *sdk.Event) error {
	if len(event.Attributes) < len(pl.indexedInputs)+len(pl.nonIndexedInputs) {
		return ErrNotEnoughAttributes
	}
	return nil
}

// `searchAttributesForArg` does a linear search through the given slice `attributes` for any
// attribute having a key that matches an Ethereum input `argName`. This function returns the index
// where `argName` was found or -1 if `argName` was not found.
// Complexity: O(N*M), N = len(`attributes`), M = average length of attribute key strings.
func searchAttributesForArg(attributes *[]abci.EventAttribute, argName string) int {
	for i, attribute := range *attributes {
		if abi.ToMixedCase(attribute.Key) == argName {
			return i
		}
	}
	// reached end of loop, `argName` not found
	return notFound
}
