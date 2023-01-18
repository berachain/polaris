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

package vm

import (
	"math/big"

	"github.com/berachain/stargazer/core/state"
	"github.com/berachain/stargazer/core/vm/precompile"
	"github.com/berachain/stargazer/core/vm/precompile/container/types"
	"github.com/berachain/stargazer/lib/common"
)

// Compile-time assertion to ensure `PrecompileHost` adheres to `precompile.Host`.
var _ precompile.Host = (*PrecompileHost)(nil)

// `PrecompileHost` is gets and executes a precompiled container at a given address.
type PrecompileHost struct {
	pr *PrecompileRegistry
}

// `NewPrecompileHost` creates and returns a new `PrecompileHost` for the given precompile
// registry `pr`.
func NewPrecompileHost(pr *PrecompileRegistry) *PrecompileHost {
	return &PrecompileHost{
		pr: pr,
	}
}

// `Exists` gets a precompile container at the given `addr` from the precompile registry.
//
// `Exists` implements `precompile.Host`.
func (ph *PrecompileHost) Exists(addr common.Address) (types.PrecompileContainer, bool) {
	return ph.pr.Get(addr)
}

// `Run` runs the given precompile container and returns the remaining gas after execution. This
// function returns an error if the given statedb is not compatible with precompiles, insufficient
// gas is provided, or the precompile execution returns an error.
//
// `Run` implements `precompile.Host`.
func (ph *PrecompileHost) Run(
	pc types.PrecompileContainer,
	sdb state.GethStateDB,
	input []byte,
	caller common.Address,
	value *big.Int,
	suppliedGas uint64,
	readonly bool,
) ([]byte, uint64, error) {
	psdb, ok := sdb.(state.PrecompileStateDB)
	if !ok {
		return nil, 0, ErrStateDBNotSupported
	}

	// TODO: move gas calculation to precompile container using gas meter.
	gasCost := pc.RequiredGas(input)
	if suppliedGas < gasCost {
		return nil, 0, ErrOutOfGas
	}
	suppliedGas -= gasCost

	psdb.EnableEventLogging()
	ret, err := pc.Run(psdb.GetContext(), input, caller, value, readonly)
	psdb.DisableEventLogging()
	if err != nil {
		return nil, suppliedGas, err
	}

	return ret, suppliedGas, nil
}
