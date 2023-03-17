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

	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/lib/errors"
	"pkg.berachain.dev/polaris/lib/errors/debug"
	"pkg.berachain.dev/polaris/lib/utils"
)

// NumBytesMethodID is the number of bytes used to represent a ABI method's ID.
const NumBytesMethodID = 4

// stateful is a container for running stateful and dynamic precompiled contracts.
type stateful struct {
	// RegistrablePrecompile is the base precompile implementation.
	vm.RegistrablePrecompile
	// idsToMethods is a mapping of method IDs (string of first 4 bytes of the keccak256 hash of
	// method signatures) to native precompile functions. The signature key is provided by the
	// precompile creator and must exactly match the signature in the geth abi.Method.Sig field
	// (geth abi format). Please check core/precompile/container/method.go for more information.
	idsToMethods map[string]*Method
	// receive      *Method // TODO: implement
	// fallback     *Method // TODO: implement
}

// NewStateful creates and returns a new `stateful` with the given method ids precompile functions map.
func NewStateful(
	rp vm.RegistrablePrecompile, idsToMethods map[string]*Method,
) vm.PrecompileContainer {
	return &stateful{
		RegistrablePrecompile: rp,
		idsToMethods:          idsToMethods,
	}
}

// Run loads the corresponding precompile method for given input, executes it, and handles
// output.
//
// Run implements `PrecompileContainer`.
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

// RequiredGas checks the Method corresponding to input for the required gas amount.
//
// RequiredGas implements PrecompileContainer.
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
