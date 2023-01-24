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

// `NewCosmosRunner` returns a `vm.RunPrecompile` that runs the a precompile container for the
// given `vm.PrecompileStateDB` and returns the remaining gas after execution. This function
// returns an error if insufficient gas is provided or the precompile execution returns an error.
func NewCosmosRunner(psdb vm.PrecompileStateDB) vm.RunPrecompile {
	return func(
		pc vm.PrecompileContainer, input []byte, caller common.Address,
		value *big.Int, suppliedGas uint64, readonly bool,
	) ([]byte, uint64, error) {
		// pre-defined, static gas consumption
		gasCost := pc.RequiredGas(input)
		if suppliedGas < gasCost {
			return nil, 0, ErrOutOfGas
		}
		suppliedGas -= gasCost

		// supply context with a precompile gas meter for dynamic consumption
		ctx := sdk.UnwrapSDKContext(psdb.GetContext())
		ctx = ctx.WithGasMeter(sdk.NewGasMeter(suppliedGas))
		ret, err := pc.Run(ctx, input, caller, value, readonly)

		return ret, ctx.GasMeter().GasRemaining(), err
	}
}
