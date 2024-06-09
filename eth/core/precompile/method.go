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
	"errors"
	"reflect"

	"github.com/berachain/polaris/eth/accounts/abi"
	errorslib "github.com/berachain/polaris/lib/errors"
	"github.com/berachain/polaris/lib/utils"

	"github.com/ethereum/go-ethereum/core/vm"
)

// methodID is a fixed length byte array that represents the method ID of a precompile method.
type methodID [NumBytesMethodID]byte

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

// method is a struct that contains the required information for the EVM to execute a stateful
// precompiled contract method.
type method struct {
	// rcvr is the receiver of the method's executable. This is the stateful precompile
	// that implements the respective precompile method.
	rcvr StatefulImpl

	// abiMethod is the ABI `Methods` struct corresponding to this precompile's executable.
	abiMethod abi.Method

	// execute is the precompile's executable which will execute the logic of the implemented
	// ABI method.
	execute reflect.Method
}

// newMethod creates and returns a new `method` with the given abiMethod, abiSig, and executable.
func newMethod(
	rcvr StatefulImpl, abiMethod abi.Method, execute reflect.Method,
) *method {
	return &method{
		rcvr:      rcvr,
		abiMethod: abiMethod,
		execute:   execute,
	}
}

// Call executes the precompile's executable with the given context and input arguments.
func (m *method) Call(ctx context.Context, input []byte) ([]byte, error) {
	// Unpack the args from the input, if any exist.
	unpackedArgs, err := m.abiMethod.Inputs.Unpack(input[NumBytesMethodID:])
	if err != nil {
		return nil, err
	}

	// Convert the unpacked args to reflect values.
	reflectedUnpackedArgs := make([]reflect.Value, 0, len(unpackedArgs))
	for _, unpacked := range unpackedArgs {
		reflectedUnpackedArgs = append(reflectedUnpackedArgs, reflect.ValueOf(unpacked))
	}

	// TODO: convert any unnamed structs into their corresponding named struct.

	// Call the executable the reflected values.
	results := m.execute.Func.Call(
		append(
			[]reflect.Value{
				reflect.ValueOf(m.rcvr),
				reflect.ValueOf(ctx),
			},
			reflectedUnpackedArgs...,
		),
	)

	// If the precompile returned an error, the error is returned to the caller.
	if revert := results[len(results)-1].Interface(); revert != nil {
		err = utils.MustGetAs[error](revert)
	}
	if err != nil {
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
	retVals := make([]any, 0, len(results)-1)
	for _, val := range results[0 : len(results)-1] {
		retVals = append(retVals, val.Interface())
	}
	ret, err := m.abiMethod.Outputs.PackValues(retVals)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
