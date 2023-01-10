// Copyright (C) 2022, Berachain Foundation. All rights reserved.
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
	"github.com/berachain/stargazer/types/abi"
)

const (
	intBase          = 10
	supportedIntBits = 64
)

type HasEvents interface {
	CosmosEventTypes() []string
	ABIEvents() map[string]abi.Event
	AttributeKeysToValueDecoder() map[string]AttributeValueDecoder
}

// Returns a geth compatible/eth primitive type (as any) for a given Cosmos event attribute.
type AttributeValueDecoder func(attributeValue string) (ethPrimitive any, err error)

// ConvertSdkCoin is AttributeValueDecoder
// converting sdk.Coin string by extracting amount string and then converting to *big.Int.
func ConvertSdkCoin(attributeValue string) (any, error) {
	coin, err := sdk.ParseCoinNormalized(attributeValue)
	if err != nil {
		return nil, err
	}
	return coin.Amount.BigInt(), nil
}

// ConvertValAddress is AttributeValueDecoder
// convert bech32 string to ValAddress to a common.Address.
func ConvertValAddress(attributeValue string) (any, error) {
	addr, err := sdk.ValAddressFromBech32(attributeValue)
	if err != nil {
		return nil, err
	}
	return common.BytesToAddress(addr.Bytes()), nil
}

// ConvertAccAddress is AttributeValueDecoder
// convert bech32 string to ValAddress to a common.Address.
func ConvertAccAddress(attributeValue string) (any, error) {
	addr, err := sdk.AccAddressFromBech32(attributeValue)
	if err != nil {
		return nil, err
	}
	return common.BytesToAddress(addr.Bytes()), nil
}

// ConvertCreationHeight is AttributeValueDecoder
// convert creationHeight string to int64.
func ConvertCreationHeight(attributeValue string) (any, error) {
	return strconv.ParseInt(attributeValue, intBase, supportedIntBits)
}
