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

	"github.com/berachain/stargazer/common"
)

const (
	// `intBase` is the base `int`s are parsed in, 10.
	intBase = 10

	// `creationHeightBits` is the number of bits used to store `creationHeight` from the Cosmos
	// SDK staking module, which is of type `int64`.
	creationHeightBits = 64
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

// `getDefaultCosmosValueDecoder` returns a default Cosmos event attribute value decoder function
// for a certain Cosmos event `attributeKey`.ß NOTE: only the event attributes of default Cosmos SDK
// modules are supported by this function.
func getDefaultCosmosValueDecoder(attributeKey string) valueDecoder {
	return ValueDecoders{
		sdk.AttributeKeyAmount:                  convertSdkCoin,
		stakingtypes.AttributeKeyValidator:      convertValAddress,
		stakingtypes.AttributeKeySrcValidator:   convertValAddress,
		stakingtypes.AttributeKeyDstValidator:   convertValAddress,
		stakingtypes.AttributeKeyCreationHeight: convertCreationHeight,
		stakingtypes.AttributeKeyDelegator:      convertAccAddress,
		banktypes.AttributeKeySender:            convertAccAddress,
		banktypes.AttributeKeyRecipient:         convertAccAddress,
		banktypes.AttributeKeySpender:           convertAccAddress,
		banktypes.AttributeKeyReceiver:          convertAccAddress,
		banktypes.AttributeKeyMinter:            convertAccAddress,
		banktypes.AttributeKeyBurner:            convertAccAddress,
	}[attributeKey]
}

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

// `convertValAddress` converts a bech32 string representing a validator address to a
// `common.Address`.
//
// `convertValAddress` is a `valueDecoder`.
func convertValAddress(attributeValue string) (any, error) {
	// extract the sdk.ValAddress from string value
	valAddress, err := sdk.ValAddressFromBech32(attributeValue)
	if err != nil {
		return nil, err
	}
	// convert the sdk.ValAddress to common.Address
	return common.ValAddressToEthAddress(valAddress), nil
}

// `convertAccAddress` converts a bech32 string representing an account address to a
// `common.Address`.
//
// `convertAccAddress` is a `valueDecoder`.
func convertAccAddress(attributeValue string) (any, error) {
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
