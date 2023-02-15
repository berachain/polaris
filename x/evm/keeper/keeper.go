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
	"github.com/ethereum/go-ethereum/core/vm"

	storetypes "cosmossdk.io/store/types"
	"github.com/berachain/stargazer/eth/core"
	"github.com/berachain/stargazer/x/evm/types"

	"github.com/cometbft/cometbft/libs/log"
)

// Compile-time interface assertions.
var _ core.StargazerHostChain = (*Keeper)(nil)

type Keeper struct {
	// The (unexposed) key used to access the store from the Context.
	storeKey storetypes.StoreKey

	// TODO: replace with the Chain interface.
	stateProcessor *core.StateProcessor

	// sk is used to retrieve infofrmation about the current / past
	// blocks and associated validator information.
	// sk StakingKeeper

	authority string
}

// NewKeeper creates new instances of the stargazer Keeper.
func NewKeeper(
	authority string,
) *Keeper {
	k := &Keeper{
		authority: authority,
		storeKey:  storetypes.NewKVStoreKey(types.StoreKey),
	}
	// TODO: remove state processor from here.
	k.stateProcessor = core.NewStateProcessor(k, nil, vm.Config{}, false)
	return k
}

// `Logger` returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", types.ModuleName)
}

func (k Keeper) GetBlockPlugin() core.BlockPlugin {
	return nil
}

func (k Keeper) GetPrecompilePlugin() core.PrecompilePlugin {
	return nil
}

func (k Keeper) GetStatePlugin() core.StatePlugin {
	return nil
}

func (k Keeper) GetGasPlugin() core.GasPlugin {
	return nil
}

func (k Keeper) GetConfigurationPlugin() core.ConfigurationPlugin {
	return nil
}
