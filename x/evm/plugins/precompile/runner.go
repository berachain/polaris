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
	"math/big"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/lib/common"
)

// Compile-time assertion to ensure `CosmosRunner` adheres to `vm.PrecompileRunner`.
var _ vm.PrecompileRunner = (*CosmosRunner)(nil)

// `CosmosRunner` runs precompile containers in a Cosmos environment for a given context.
type CosmosRunner struct {
	// `kvGasConfig` is the gas config for execution of kv store operations in native precompiles.
	kvGasConfig *sdk.GasConfig

	// `transientKVGasConfig` is the gas config for execution transient kv store operations in
	// native precompiles.
	transientKVGasConfig *sdk.GasConfig
}

// `NewCosmosRunner` creates and returns a `CosmosRunner` with the SDK default gas configs.
func NewCosmosRunner() *CosmosRunner {
	defaultKVGasConfig := storetypes.KVGasConfig()
	defaultTransientKVGasConfig := storetypes.TransientGasConfig()

	return &CosmosRunner{
		kvGasConfig:          &defaultKVGasConfig,
		transientKVGasConfig: &defaultTransientKVGasConfig,
	}
}

// `WithKVGasConfig` returns a `CosmosRunner` with `kvGasConfig` attached.
func (cr CosmosRunner) WithKVGasConfig(kvGasConfig *sdk.GasConfig) CosmosRunner {
	cr.kvGasConfig = kvGasConfig
	return cr
}

// `WithTransientKVGasConfig` returns a `CosmosRunner` with `transientKVGasConfig` attached.
func (cr CosmosRunner) WithTransientKVGasConfig(transientKVGasConfig *sdk.GasConfig) CosmosRunner {
	cr.transientKVGasConfig = transientKVGasConfig
	return cr
}

// `Run` runs the a precompile container and returns the remaining gas after execution by injecting
// a Cosmos SDK `GasMeter`. This function returns an error if insufficient gas is provided or the
// precompile execution returns an error.
//
// `Run` implements `vm.PrecompileRunner`.
func (cr *CosmosRunner) Run(
	pc vm.PrecompileContainer, ssdb vm.StargazerStateDB, input []byte,
	caller common.Address, value *big.Int, suppliedGas uint64, readonly bool,
) ([]byte, uint64, error) {
	// use a precompile-specific gas meter for dynamic consumption
	gm := sdk.NewInfiniteGasMeter()
	// consume static gas from RequiredGas
	gm.ConsumeGas(pc.RequiredGas(input), "RequiredGas")

	// run precompile container
	ret, err := pc.Run(
		sdk.UnwrapSDKContext(ssdb.GetContext()).
			WithGasMeter(gm).
			WithKVGasConfig(*cr.kvGasConfig).
			WithTransientKVGasConfig(*cr.transientKVGasConfig),
		ssdb,
		input,
		caller,
		value,
		readonly,
	)

	// handle overconsumption of gas
	if gm.GasConsumed() > suppliedGas {
		return nil, 0, ErrOutOfGas
	}

	return ret, suppliedGas - gm.GasConsumed(), err
}
