// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package precompile

import (
	"context"
	"math/big"

	pvm "github.com/berachain/polaris/eth/core/vm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// NumBytesMethodID is the number of bytes used to represent a ABI method's ID.
const NumBytesMethodID = 4

var _ vm.PrecompiledContract = (*statefulContainer)(nil)

// statefulContainer is a container for running statefulContainer and precompiled contracts.
type statefulContainer struct {
	// StatefulImpl is the base precompile implementation.
	StatefulImpl
	// idsToMethods is a mapping of method IDs (string of first 4 bytes of the keccak256 hash of
	// method signatures) to native precompile functions. The signature key is provided by the
	// precompile creator and must exactly match the signature in the geth abi.Method.Sig field
	// (geth abi format). Please check core/precompile/container/method.go for more information.
	idsToMethods map[methodID]*method
	// receive      *Method // TODO: implement
	// fallback     *Method // TODO: implement

}

// NewStatefulContainer creates and returns a new `statefulContainer` with the given method ids
// precompile functions map.
func NewStatefulContainer(
	si StatefulImpl, idsToMethods map[methodID]*method,
) (vm.PrecompiledContract, error) {
	if idsToMethods == nil {
		return nil, ErrContainerHasNoMethods
	}
	return &statefulContainer{
		StatefulImpl: si,
		idsToMethods: idsToMethods,
	}, nil
}

// Run loads the corresponding precompile method for given input, executes it, and handles
// output.
//
// Run implements `PrecompileContainer`.
func (sc *statefulContainer) Run(
	ctx context.Context,
	evm vm.PrecompileEVM,
	input []byte,
	caller common.Address,
	value *big.Int,
) ([]byte, error) {
	if len(input) < NumBytesMethodID {
		return nil, ErrInvalidInputToPrecompile
	}

	// Extract the method ID from the input and load the method.
	method, found := sc.idsToMethods[methodID(input)]
	if !found {
		return nil, ErrMethodNotFound
	}

	// Execute the method with the reflected ctx and raw input
	return method.Call(
		pvm.NewPolarContext(ctx, evm, caller, value),
		input,
	)
}

// RequiredGas checks the Method corresponding to input for the required gas amount. TODO: remove
// unneeded input from interface.
//
// RequiredGas implements PrecompileContainer.
func (sc *statefulContainer) RequiredGas([]byte) uint64 {
	return 0
}
