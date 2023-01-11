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

package events

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/berachain/stargazer/common"
)

const (
	// `intBase` is the base `int`s are parsed in, 10.
	intBase = 10

	// `creationHeightBits` is the number of bits used to store `creationHeight` from the Cosmos
	// SDK staking module, which is of type `int64`.
	creationHeightBits = 64
)

// `AttributeValueDecoder` is a type of function that returns a geth compatible, eth primitive type
// (as type `any`) for a given Cosmos event attribute value (of type `string`). Cosmos event
// attribute values may require unique decodings based on their underlying string encoding.
type AttributeValueDecoder func(attributeValue string) (ethPrimitive any, err error)

// `ConvertSdkCoin` converts the string representation of an `sdk.Coin` to a `*big.Int`.
//
// `ConvertSdkCoin` is a `AttributeValueDecoder`.
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
// `common.Address`.
//
// `ConvertValAddress` is a `AttributeValueDecoder`.
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
// `common.Address`.
//
// `ConvertAccAddress` is a `AttributeValueDecoder`.
func ConvertAccAddress(attributeValue string) (any, error) {
	// extract the sdk.AccAddress from string value
	accAddress, err := sdk.AccAddressFromBech32(attributeValue)
	if err != nil {
		return nil, err
	}
	// convert the sdk.AccAddress to common.Address
	return common.AccAddressToEthAddress(accAddress), nil
}

// `ConvertCreationHeight` converts a creationHeight, from the Cosmos SDK staking module, `string`
// to an `int64`.
//
// `ConvertCreationHeight` is a `AttributeValueDecoder`.
func ConvertCreationHeight(attributeValue string) (any, error) {
	return strconv.ParseInt(attributeValue, intBase, creationHeightBits)
}
