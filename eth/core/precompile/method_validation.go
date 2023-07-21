// SPDX-License-Identifier: Apache-2.0
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package precompile

import (
	"errors"
	"fmt"
	"reflect"

	"pkg.berachain.dev/polaris/eth/accounts/abi"
)

// tryMatchInputs returns true iff the argument types match between the Go implementation and the
// ABI method.
func tryMatchInputs(implMethod reflect.Method, abiMethod *abi.Method) bool {
	abiMethodNumIn := len(abiMethod.Inputs)

	// First two args of Go precompile implementation are the receiver contract and the Context, so
	// verify that the ABI method has exactly 2 fewer inputs than the implementation method.
	if implMethod.Type.NumIn()-2 != abiMethodNumIn {
		return false
	}

	// If the function does not take any inputs, no need to check.
	if abiMethodNumIn > 0 {
		// Validate that the precompile input args types match ABI input arg types, excluding the
		// first two args (receiver contract and Context).
		for i := 2; i < implMethod.Type.NumIn(); i++ {
			implMethodParamType := implMethod.Type.In(i)
			abiMethodParamType := abiMethod.Inputs[i-2].Type.GetType()
			if validateArg(
				reflect.New(implMethodParamType).Elem(), reflect.New(abiMethodParamType).Elem(),
			) != nil {
				return false
			}
		}
	}

	return true
}

// validateArg uses reflection to verify the implementation arg matches the ABI arg.
func validateArg(implMethodVar reflect.Value, abiMethodVar reflect.Value) error {
	implMethodVarType := implMethodVar.Type()
	abiMethodVarType := abiMethodVar.Type()

	//nolint:exhaustive // checking necessary conditions.
	switch implMethodVarType.Kind() {
	case reflect.Array, reflect.Slice:
		if implMethodVarType.Elem() != abiMethodVarType.Elem() {
			// If the array is not a slice/array of structs, return an error.
			if implMethodVarType.Elem().Kind() != reflect.Struct {
				return fmt.Errorf(
					"return type mismatch: %v != %v", implMethodVar.Elem(), abiMethodVar.Elem(),
				)
			}

			// If it is a slice/array of structs, check if the struct fields match.
			if err := validateStruct(implMethodVarType.Elem(), abiMethodVarType.Elem()); err != nil {
				return err
			}
		}
	case abiMethodVarType.Kind():
		// If it's a struct, check all the fields to match
		if implMethodVarType.Kind() == reflect.Struct {
			if err := validateStruct(implMethodVarType, abiMethodVarType); err != nil {
				return err
			}
		}
		// If the types (primitives) match, we're good.
	case reflect.Interface:
		// If it's `any` (reflect.Interface), we leave it to the implementer to make sure that it is
		// used/converted correctly.
	default:
		return fmt.Errorf("return type mismatch: %v != %v", implMethodVarType, abiMethodVarType)
	}

	return nil
}

// validateStruct checks to make sure that the implementation struct's fields match the ABI struct's
// fields.
func validateStruct(implMethodVarType reflect.Type, abiMethodVarType reflect.Type) error {
	if implMethodVarType.Kind() != reflect.Struct || abiMethodVarType.Kind() != reflect.Struct {
		return errors.New("validateStruct: not a struct")
	}

	if implMethodVarType.NumField() != abiMethodVarType.NumField() {
		return fmt.Errorf(
			"struct %v has %v fields, but struct %v has %v fields",
			implMethodVarType.Name(),
			implMethodVarType.NumField(),
			abiMethodVarType.Name(),
			abiMethodVarType.NumField(),
		)
	}

	// match every individual field
	for j := 0; j < implMethodVarType.NumField(); j++ {
		if err := validateArg(
			reflect.New(implMethodVarType.Field(j).Type).Elem(),
			reflect.New(abiMethodVarType.Field(j).Type).Elem(),
		); err != nil {
			return err
		}
	}
	return nil
}

// validateOutputs checks if the impl method output types match the ABI's return types.
func validateOutputs(implMethod reflect.Method, abiMethod *abi.Method) error {
	// Now verify the outputs match.
	implMethodNumOut := implMethod.Type.NumOut()

	// The Solidity compiler requires that precompiles must return at least one value.
	// See https://github.com/berachain/polaris/issues/491 for more information.
	if len(abiMethod.Outputs) == 0 {
		//nolint:lll // error message.
		panic("The Solidity compiler requires all precompile functions to return at least one value. Consider returning a boolean.")
	}

	// Last parameter of Go precompile implementation is an error (for reverts),
	// so we skip that.
	if implMethodNumOut-1 != len(abiMethod.Outputs) {
		return errors.New("number of return types mismatch")
	}

	// Validate that our implementation returns an error (revert) as the last param.
	if implMethod.Type.Out(implMethodNumOut-1) != reflect.TypeOf((*error)(nil)).Elem() {
		return fmt.Errorf(
			"last return type must be error, got %v", implMethod.Type.Out(implMethodNumOut-1),
		)
	}

	// Validate that the other implementation return types match ABI return types.
	for i := 0; i < implMethodNumOut-1; i++ {
		implMethodReturnType := implMethod.Type.Out(i)
		abiMethodReturnType := abiMethod.Outputs[i].Type.GetType()
		if err := validateArg(
			reflect.New(implMethodReturnType).Elem(), reflect.New(abiMethodReturnType).Elem(),
		); err != nil {
			return fmt.Errorf(
				"return type mismatch: %v != %v", implMethodReturnType, abiMethodReturnType,
			)
		}
	}

	return nil
}
