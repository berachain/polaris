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

package container

import (
	"context"
	"math/big"

	"github.com/berachain/stargazer/core/vm/precompile/container/types"
	"github.com/berachain/stargazer/lib/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// `NumBytesMethodID` is the number of bytes used to represent a ABI method's ID.
const NumBytesMethodID = 4

// Compile-time assertion to ensure `StatefulContainer` is a `PrecompileContainer`.
var _ types.PrecompileContainer = (*StatefulContainer)(nil)

// `StatefulContainer` is a container for running stateful (and dynamic) precompiled contracts.
type StatefulContainer struct {
	// `idsToMethods` is a mapping of method IDs (keccak256 hash of method signatures) to native
	// precompile functions. The signature key is provided by the precompile creator and must
	// exactly match the signature in the geth abi.Method.Sig field (geth abi format). Please check
	// core/vm/precompile/container/types.go for more information.
	idsToMethods map[common.Hash]*types.PrecompileMethod
}

// `NewStatefulContainer` creates and returns a new `StatefulContainer` with the given method ids
// precompile functions map.
func NewStatefulContainer(idsToMethods map[common.Hash]*types.PrecompileMethod) *StatefulContainer {
	return &StatefulContainer{
		idsToMethods: idsToMethods,
	}
}

// `Run` loads the corresponding precompile method for given input and executes it.
//
// `Run` implements `PrecompileContainer`.
func (sc *StatefulContainer) Run(
	ctx context.Context,
	input []byte,
	caller common.Address,
	value *big.Int,
	readonly bool,
) ([]byte, error) {
	// TODO: this will change to take in `suppliedGas` and return `remainingGas`.

	if sc.idsToMethods == nil {
		return nil, types.ErrPrecompileHasNoMethods
	}

	// extract the method ID from the input and load the method.
	method, ok := sc.idsToMethods[common.BytesToHash(input[:NumBytesMethodID])]
	if !ok {
		return nil, types.ErrPrecompileMethodNotFound
	}

	// unpack the args from the input
	unpackedArgs, err := method.AbiMethod.Inputs.Unpack(input[NumBytesMethodID:])
	if err != nil {
		return nil, err
	}

	// Execute the method registered with the given signature with the given args.
	vals, err := method.Execute(
		sdk.UnwrapSDKContext(ctx),
		caller,
		value,
		readonly,
		unpackedArgs...,
	)

	// If the precompile returned an error, the error is returned to the caller.
	if err != nil {
		return nil, err
	}

	return method.AbiMethod.Outputs.PackValues(vals)
}

// `RequiredGas` checks the PrecompileMethod corresponding to input for the required gas amount.
//
// `RequiredGas` implements PrecompileContainer.
func (sc *StatefulContainer) RequiredGas(input []byte) uint64 {
	// TODO: RequiredGas will be deprecated in Geth.

	if sc.idsToMethods == nil {
		return 0
	}

	// extract the method ID from the input and load the method.
	method, ok := sc.idsToMethods[common.BytesToHash(input[:NumBytesMethodID])]
	if !ok {
		return 0
	}

	return method.RequiredGas
}
