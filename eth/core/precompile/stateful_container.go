// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

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
