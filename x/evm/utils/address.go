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

package utils

import (
	"github.com/berachain/stargazer/eth/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// `AccAddressToEthAddress` converts a Cosmos SDK `AccAddress` to an Ethereum `Address`.
func AccAddressToEthAddress(accAddress sdk.AccAddress) common.Address {
	return common.BytesToAddress(accAddress)
}

// `ValAddressToEthAddress` converts a Cosmos SDK `ValAddress` to an Ethereum `Address`.
func ValAddressToEthAddress(valAddress sdk.ValAddress) common.Address {
	return common.BytesToAddress(valAddress)
}

// `AddressToAccAddress` converts an Ethereum `Address` to a Cosmos SDK `AccAddress`.
func AddressToAccAddress(ethAddress common.Address) sdk.AccAddress {
	return ethAddress.Bytes()
}

// `AddressToValAddress` converts an Ethereum `Address` to a Cosmos SDK `ValAddress`.
func AddressToValAddress(ethAddress common.Address) sdk.ValAddress {
	return ethAddress.Bytes()
}
