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
	"reflect"

	"pkg.berachain.dev/polaris/eth/accounts/abi"
	"pkg.berachain.dev/polaris/eth/core/vm"
	errorslib "pkg.berachain.dev/polaris/lib/errors"
	"pkg.berachain.dev/polaris/lib/errors/debug"
)

/**
 * 	Welcome to Stateful Precompiled Contracts! To build a stateful precompile, you must implement
 *  the `StatefulImpl` interface in `interfaces.go`; below are the suggested steps to
 *  follow:
 *	  1) Define a Solidity interface with the methods that you want implemented via a precompile.
 *	  2) Build a Go precompile contract, which implements the interface's methods.
 *       A) This precompile contract should expose the ABI's `Methods`, which can be generated via
 *          Go-Ethereum's abi package. These methods are of type `abi.Method`.
 *   	 B) This precompile contract should also expose the `Method`s. A `Method` includes the
 *          `executable`, which is the direct implementation of a corresponding ABI method, the
 *          `executable`'s `RequiredGas`, and the ABI signature. Do NOT provide the `AbiMethod` as
 *          this field will be automatically populated.
 * 		 C) If implementing an overloaded function, suffix the overloaded methods' names with starting with
 *  		0, then 1, 2, etc.  for every overloaded function. For example, if you have two functions named `foo` in
 *			your smart contract, then name the first function `foo` and the second `foo0`.
 *			This is because Go does not allow overloaded functions, and is very similar to how Geth handles it.
 **/

// Method is a struct that contains the required information for the EVM to execute a stateful
// precompiled contract method.
type Method struct {
	// AbiMethod is the ABI `Methods` struct corresponding to this precompile's executable. NOTE:
	// this field should be left empty (as nil) as this will automatically be populated by the
	// corresponding interface's ABI.
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
func (m *Method) Call(ctx []reflect.Value, input []byte) ([]byte, error) {
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
	results := m.execute.Call(append(ctx, reflectedUnpackedArgs...))

	// If the precompile returned an error, the error is returned to the caller.
	if !results[1].IsNil() {
		return nil, errorslib.Wrapf(
			vm.ErrExecutionReverted,
			"vm error [%v] occurred during precompile execution of [%s]",
			results[1].Interface().(error), debug.GetFnName(m.execute.Interface()),
		)
	}

	// Pack the return values and return, if any exist.
	retVal := results[0]
	ret, err := m.abiMethod.Outputs.PackValues(retVal.Interface().([]any)) // 1) What
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// Methods is a type that represents a list of precompile methods. This is what a stateful
// precompiled contract implementation should expose.
type Methods []*Method
