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

package state

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// "github.com/berachain/stargazer/params".
	"github.com/berachain/stargazer/core/state/types"
)

type StateDBFactory struct { //nolint:revive // the vibes are good.
	// Cosmos Keeper References
	ak types.AccountKeeper
	bk types.BankKeeper

	// evmStoreKey is the store key for the EVM store.
	evmStoreKey storetypes.StoreKey

	// evmDenom is the denom used for EVM transactions.
	// evmDenom params.Retriever[params.EVMDenom]

	lf types.EthereumLogFactory
}

// NewStateDBFactory returns a new StateDBFactory instance.
func NewStateDBFactory(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	evmStoreKey storetypes.StoreKey,
	// evmDenom params.Retriever[params.EVMDenom],
	logFactory types.EthereumLogFactory,
) *StateDBFactory {
	return &StateDBFactory{
		ak:          ak,
		bk:          bk,
		lf:          logFactory,
		evmStoreKey: evmStoreKey,
		// evmDenom:    evmDenom,
		// er:          er,
	}
}

// BuildNewStateDB returns a new StateDB instance.
func (sdf *StateDBFactory) BuildStateDB(ctx sdk.Context) *StateDB {
	return NewStateDB(ctx, sdf.ak, sdf.bk, sdf.lf, sdf.evmStoreKey, "abera") // sdf.evmDenom.Get(ctx), sdf.er)
}
