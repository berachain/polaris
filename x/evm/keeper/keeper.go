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

package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	storetypes "cosmossdk.io/store/types"
	"github.com/berachain/stargazer/eth/api"
	"github.com/berachain/stargazer/eth/core"
	"github.com/berachain/stargazer/x/evm/plugins/block"
	"github.com/berachain/stargazer/x/evm/plugins/configuration"
	"github.com/berachain/stargazer/x/evm/plugins/gas"
	"github.com/berachain/stargazer/x/evm/plugins/precompile"
	precompilelog "github.com/berachain/stargazer/x/evm/plugins/precompile/log"
	"github.com/berachain/stargazer/x/evm/plugins/state"
	"github.com/berachain/stargazer/x/evm/types"

	"github.com/cometbft/cometbft/libs/log"
)

// Compile-time interface assertions.
var _ core.StargazerHostChain = (*Keeper)(nil)

type Keeper struct {
	// The (unexposed) key used to access the store from the Context.
	storeKey storetypes.StoreKey

	ethChain api.Chain

	// sk is used to retrieve infofrmation about the current / past
	// blocks and associated validator information.
	// sk StakingKeeper

	authority string

	// plugins
	bp core.BlockPlugin
	cp core.ConfigurationPlugin
	gp core.GasPlugin
	pp core.PrecompilePlugin
	sp core.StatePlugin
}

// NewKeeper creates new instances of the stargazer Keeper.
func NewKeeper(
	authority string,
) *Keeper {
	k := &Keeper{
		authority: authority,
		storeKey:  storetypes.NewKVStoreKey(types.StoreKey),
	}
	k.ethChain = core.NewChain(k)
	return k
}

// `Logger` returns a module-specific logger.
func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", types.ModuleName)
}

func (k *Keeper) InitPlugins(ctx sdk.Context, ak state.AccountKeeper, bk state.BankKeeper) {
	k.bp = block.NewPluginFrom(ctx, k)

	k.cp = configuration.NewPluginFrom(ctx)

	k.gp = gas.NewPluginFrom(ctx)

	k.pp = precompile.NewPluginFrom(ctx)
	// TODO: register precompiles

	plf := precompilelog.NewFactory()
	// TODO: register precompile events/logs

	k.sp = state.NewPlugin(ctx, ak, bk, k.storeKey, types.ModuleName, plf)
}

func (k *Keeper) GetBlockPlugin() core.BlockPlugin {
	return k.bp
}

func (k *Keeper) GetConfigurationPlugin() core.ConfigurationPlugin {
	return k.cp
}

func (k *Keeper) GetGasPlugin() core.GasPlugin {
	return k.gp
}

func (k *Keeper) GetPrecompilePlugin() core.PrecompilePlugin {
	return k.pp
}

func (k *Keeper) GetStatePlugin() core.StatePlugin {
	return k.sp
}
