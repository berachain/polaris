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
	"errors"
	"reflect"

	"pkg.berachain.dev/polaris/eth/accounts/abi"
	"pkg.berachain.dev/polaris/eth/core/vm"
	errorslib "pkg.berachain.dev/polaris/lib/errors"
	"pkg.berachain.dev/polaris/lib/utils"
)

/**
 * 	Welcome to Stateful Precompiled Contracts! To build a stateful precompile, you must implement
 *  the `StatefulImpl` interface in `interfaces.go`; below are the suggested steps to
 *  follow:
 *	  1) Define a Solidity interface with the methods that you want implemented via a precompile.
 *	  2) Build a Go precompile contract, which implements the interface's methods.
 *       A) This precompile contract should expose the ABI's `Methods`, which can be generated via
 *          Go-Ethereum's abi package. These methods are of type `abi.Method`.
 * 		 B) If implementing an overloaded function, suffix the overloaded methods' names starting
 *          with 0, 1, 2, ... for every overloaded function. For example, if you have two functions
 *          named `foo` in your smart contract, then name the first function `foo` and the second
 *          `foo0`. We enforce the same overloading scheme that geth's abi package uses.
 **/

// Method is a struct that contains the required information for the EVM to execute a stateful
// precompiled contract method.
type Method struct {
	// AbiMethod is the ABI `Methods` struct corresponding to this precompile's executable.
	abiMethod *abi.Method

	// AbiSig returns the method's string signature according to the ABI spec.
	// e.g.		function foo(uint32 a, int b) = "foo(uint32,int256)"
	// Note that there are no spaces and variable names in the signature.
	// Also note that "int" is substitute for its canonical representation "int256".
	abiSig string

	// Execute is the precompile's executable which will execute the logic of the implemented
	// ABI method.
	execute reflect.Value
}

// NewMethod creates and returns a new `Method` with the given abiMethod, abiSig, and executable.
func NewMethod(
	abiMethod *abi.Method, abiSig string, execute reflect.Value,
) *Method {
	return &Method{
		abiMethod: abiMethod,
		abiSig:    abiSig,
		execute:   execute,
	}
}

// Call executes the precompile's executable with the given context and input arguments.
//
//nolint:revive // needed for reflection.
func (m *Method) Call(si StatefulImpl, ctx context.Context, input []byte) ([]byte, error) {
	// Unpack the args from the input, if any exist.
	unpackedArgs, err := m.abiMethod.Inputs.Unpack(input[NumBytesMethodID:])
	if err != nil {
		return nil, err
	}

	// Build argument list
	reflectedUnpackedArgs := make([]reflect.Value, 0, len(unpackedArgs))
	for _, unpacked := range unpackedArgs {
		reflectedUnpackedArgs = append(reflectedUnpackedArgs, reflect.ValueOf(unpacked))
	}

	// Call the executable
	results := m.execute.Call(append(
		[]reflect.Value{
			reflect.ValueOf(sc),
			reflect.ValueOf(ctx),
		}, reflectedUnpackedArgs...))

	// If the precompile returned an error, the error is returned to the caller.
	if !results[1].IsNil() {
		err = utils.MustGetAs[error](results[1].Interface())
		if !errors.Is(err, vm.ErrWriteProtection) {
			err = errorslib.Wrapf(
				vm.ErrExecutionReverted,
				"vm error [%v] occurred during precompile execution of [%s]",
				err, m.abiMethod.Name,
			)
		}
		return nil, err
	}

	// Pack the return values and return, if any exist.
	ret, err := m.abiMethod.Outputs.PackValues(utils.MustGetAs[[]any](results[0].Interface()))
	if err != nil {
		return nil, err
	}

	return ret, nil
}
