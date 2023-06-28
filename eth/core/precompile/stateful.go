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
	"fmt"
	"math/big"
	"reflect"

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
	// Registrable is the base precompile implementation.
	Registrable
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
	rp Registrable, idsToMethods map[string]*Method,
) vm.PrecompileContainer {
	return &stateful{
		Registrable:  rp,
		idsToMethods: idsToMethods,
	}
}

// Run loads the corresponding precompile method for given input, executes it, and handles
// output.
//
// Run implements `PrecompileContainer`.
func (sc *stateful) Run(
	ctx context.Context,
	evm EVM,
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

	// Get args ready for precompile call.
	// TODO, remove most of these args. In the future , we should only need the arguments from the method according to the ABI and a context rather than all of these.
	var fullargs []reflect.Value
	if reflect.ValueOf(sc.Registrable).IsValid() {
		fullargs = append(fullargs, reflect.ValueOf(sc.Registrable))
	}
	fullargs = append(fullargs, reflect.ValueOf(ctx))
	fullargs = append(fullargs, reflect.ValueOf(evm))
	fullargs = append(fullargs, reflect.ValueOf(caller))
	fullargs = append(fullargs, reflect.ValueOf(value))
	fullargs = append(fullargs, reflect.ValueOf(readonly))

	var reflectedUnpackedArgs []reflect.Value // needed for .Call(...)

	for _, unpacked := range unpackedArgs {
		reflectedUnpackedArgs = append(reflectedUnpackedArgs, reflect.ValueOf(unpacked))
		fmt.Println("type of unpacked", reflect.TypeOf(unpacked).String())
	}
	fullargs = append(fullargs, reflectedUnpackedArgs...)

	// Execute the method registered with the given signature with the given args.
	results := method.Execute.Call(fullargs)
	fmt.Println("results: ", results)
	// If the precompile returned an error, the error is returned to the caller.
	if !results[1].IsNil() {
		if err = results[1].Interface().(error); err != nil {
			fmt.Println("errored: ", err)
			return nil, errors.Wrapf(
				vm.ErrExecutionReverted,
				"vm error [%v] occurred during precompile execution of [%s]",
				err, debug.GetFnName(method.Execute.Interface()),
			)
		}
	}

	// Pack the return values and return, if any exist.
	retVal := results[0]
	fmt.Println("retVal: ", retVal)
	ret, err := method.AbiMethod.Outputs.PackValues(retVal.Interface().([]interface{})) // 1) What
	if err != nil {
		return nil, err
	}
	fmt.Println("ret: ", ret)
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
