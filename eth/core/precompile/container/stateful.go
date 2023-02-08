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

	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/errors"
	"github.com/berachain/stargazer/lib/errors/debug"
	"github.com/berachain/stargazer/lib/utils"
)

// `NumBytesMethodID` is the number of bytes used to represent a ABI method's ID.
const NumBytesMethodID = 4

// `stateful` is a container for running stateful and dynamic precompiled contracts.
type stateful struct {
	// `RegistrablePrecompile` is the base precompile implementation.
	vm.RegistrablePrecompile
	// `idsToMethods` is a mapping of method IDs (string of first 4 bytes of the keccak256 hash of
	// method signatures) to native precompile functions. The signature key is provided by the
	// precompile creator and must exactly match the signature in the geth abi.Method.Sig field
	// (geth abi format). Please check core/precompile/container/method.go for more information.
	idsToMethods map[string]*Method
	// receive      *Method // TODO: implement
	// fallback     *Method // TODO: implement
}

// `NewStateful` creates and returns a new `stateful` with the given method ids precompile functions map.
func NewStateful(
	rj vm.RegistrablePrecompile, idsToMethods map[string]*Method,
) vm.PrecompileContainer {
	return &stateful{
		RegistrablePrecompile: rj,
		idsToMethods:          idsToMethods,
	}
}

// `Run` loads the corresponding precompile method for given input, executes it, and handles
// output.
//
// `Run` implements `PrecompileContainer`.
func (sc *stateful) Run(
	ctx context.Context,
	input []byte,
	caller common.Address,
	value *big.Int,
	readonly bool,
) ([]byte, error) {
	if sc.idsToMethods == nil {
		return nil, ErrContainerHasNoMethods
	}
	if len(input) < NumBytesMethodID {
		return nil, ErrInvalidInputToPrecompile
	}

	// Extract the method ID from the input and load the method.
	method, found := sc.idsToMethods[utils.UnsafeBytesToStr(input[:NumBytesMethodID])]
	if !found {
		return nil, ErrMethodNotFound
	}

	// Unpack the args from the input, if any exist.
	unpackedArgs, err := method.AbiMethod.Inputs.Unpack(input[NumBytesMethodID:])
	if err != nil {
		return nil, err
	}

	// Execute the method registered with the given signature with the given args.
	vals, err := method.Execute(
		ctx,
		caller,
		value,
		readonly,
		unpackedArgs...,
	)

	// If the precompile returned an error, the error is returned to the caller.
	if err != nil {
		return nil, errors.Wrap(err, debug.GetFnName(method.Execute))
	}

	// Pack the return values and return, if any exist.
	ret, err := method.AbiMethod.Outputs.Pack(vals...)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// `RequiredGas` checks the Method corresponding to input for the required gas amount.
//
// `RequiredGas` implements PrecompileContainer.
func (sc *stateful) RequiredGas(input []byte) uint64 {
	if sc.idsToMethods == nil || len(input) < NumBytesMethodID {
		return 0
	}

	// Extract the method ID from the input and load the method.
	method, found := sc.idsToMethods[utils.UnsafeBytesToStr(input[:NumBytesMethodID])]
	if !found {
		return 0
	}

	return method.RequiredGas
}
