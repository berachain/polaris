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

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/berachain/stargazer/core/vm"
	"github.com/berachain/stargazer/lib/common"
)

// Compile-time assertion to ensure `cosmosRunner` adheres to `vm.PrecompileRunner`.
var _ vm.PrecompileRunner = (*cosmosRunner)(nil)

// `cosmosRunner` is a struct that runs precompile containers in a Cosmos environment.
type cosmosRunner struct {
	// `psdb` allows the `cosmosRunner` to inject a context into the execution environment of the
	// precompile container.
	psdb vm.PrecompileStateDB
}

// `NewCosmosRunner` creates and returns a `cosmosRunner` with the given `PrecompileStateDB`.
//
//nolint:revive // this will only be used as a `vm.PrecompileRunner`.
func NewCosmosRunner(psdb vm.PrecompileStateDB) *cosmosRunner {
	return &cosmosRunner{
		psdb: psdb,
	}
}

// `Run` runs the a precompile container and returns the remaining gas after execution by injecting
// a Cosmos SDK `GasMeter`. This function returns an error if insufficient gas is provided or the
// precompile execution returns an error.
//
// `Run` implements `vm.PrecompileRunner`.
func (cr *cosmosRunner) Run(
	pc vm.PrecompileContainer, input []byte, caller common.Address,
	value *big.Int, suppliedGas uint64, readonly bool,
) ([]byte, uint64, error) {
	// deterministic, static gas consumption
	gasCost := pc.RequiredGas(input)
	if suppliedGas < gasCost {
		return nil, 0, ErrOutOfGas
	}
	suppliedGas -= gasCost

	// supply context with a precompile-specific gas meter for dynamic consumption
	ctx := sdk.UnwrapSDKContext(cr.psdb.GetContext())
	ctx = ctx.WithGasMeter(sdk.NewGasMeter(suppliedGas))
	ret, err := pc.Run(ctx, input, caller, value, readonly)

	return ret, ctx.GasMeter().GasRemaining(), err
}
