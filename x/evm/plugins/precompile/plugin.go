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

package precompile

import (
	"context"
	"math/big"

	storetypes "cosmossdk.io/store/types"
	"github.com/berachain/stargazer/x/evm/plugins/state"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/eth/core"
	"github.com/berachain/stargazer/eth/core/precompile"
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/lib/utils"
)

// `plugin` runs precompile containers in the Cosmos environment with the context gas configs.
type plugin struct {
	sdk.Context
}

// `NewPluginFrom` creates and returns a `plugin` with the given context.
func NewPluginFrom(ctx sdk.Context) core.PrecompilePlugin {
	return &plugin{
		Context: ctx,
	}
}

// `Reset` implements `core.PrecompilePlugin`.
func (p *plugin) Reset(ctx context.Context) {
	p.Context = sdk.UnwrapSDKContext(ctx)
}

// `Run` runs the a precompile container and returns the remaining gas after execution by injecting
// a Cosmos SDK `GasMeter`. This function returns an error if the precompile execution returns an
// error or insufficient gas is provided.
//
// `Run` implements `core.PrecompilePlugin`.
func (p *plugin) Run(
	ldb precompile.LogsDB, pc vm.PrecompileContainer, input []byte,
	caller common.Address, value *big.Int, suppliedGas uint64, readonly bool,
) ([]byte, uint64, error) {
	// use a precompile-specific gas meter for dynamic consumption
	gm := storetypes.NewInfiniteGasMeter()
	// consume static gas from RequiredGas
	gm.ConsumeGas(pc.RequiredGas(input), "RequiredGas")

	// begin precompile execution => begin emitting Cosmos event as Eth logs
	cem := utils.MustGetAs[state.ControllableEventManager](p.Context.EventManager())
	cem.BeginPrecompileExecution(ldb)

	// run precompile container
	ret, err := pc.Run(
		p.Context.WithGasMeter(gm),
		input,
		caller,
		value,
		readonly,
	)

	// end precompile execution => stop emitting Cosmos event as Eth logs
	cem.EndPrecompileExecution()

	// handle overconsumption of gas
	if gm.GasConsumed() > suppliedGas {
		return nil, 0, vm.ErrOutOfGas
	}

	// valid precompile gas consumption => return supplied gas
	return ret, suppliedGas - gm.GasConsumed(), err
}
