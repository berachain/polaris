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
	"context"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ethstate "github.com/berachain/stargazer/eth/core/state"
	"github.com/berachain/stargazer/x/evm/plugins/state/types"
)

type PluginFactory struct { 
	// Cosmos Keeper References
	ak types.AccountKeeper
	bk types.BankKeeper

	// evmStoreKey is the store key for the EVM store.
	evmStoreKey storetypes.StoreKey

	// evmDenom is the denom used for EVM transactions.
	// evmDenom params.Retriever[params.EVMDenom]
}

// NewPluginFactory returns a new PluginFactory instance.
func NewPluginFactory(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	evmStoreKey storetypes.StoreKey,
	logFactory types.EthereumLogFactory,
) *PluginFactory {
	return &PluginFactory{
		ak:          ak,
		bk:          bk,
		evmStoreKey: evmStoreKey,
	}
}

// Build returns a new StateDB instance.
func (pf *PluginFactory) Build(ctx context.Context) ethstate.StatePlugin {
	// TODO: handle error? / ignore it completely?
	sp, _ := NewPlugin(sdk.UnwrapSDKContext(ctx), pf.ak, pf.bk, pf.evmStoreKey, "abera")
	return sp
}
