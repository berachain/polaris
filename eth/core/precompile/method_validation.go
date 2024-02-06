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

	"github.com/berachain/polaris/eth/accounts/abi"
)

// validateArg uses reflection to verify the implementation arg matches the ABI arg.
//
//nolint:gocognit // required for reflect.
func validateArg(implMethodVar reflect.Value, abiMethodVar reflect.Value) error {
	implMethodVarType := implMethodVar.Type()
	abiMethodVarType := abiMethodVar.Type()

	switch implMethodVarType.Kind() { //nolint:exhaustive // todo verify its okay.
	case reflect.Array, reflect.Slice:
		// abiMethodVarType is not also a slice or array
		if abiMethodVarType.Kind() != reflect.Array && abiMethodVarType.Kind() != reflect.Slice {
			return fmt.Errorf(
				"type mismatch: %v != %v", implMethodVarType, abiMethodVarType,
			)
		}

		if implMethodVarType.Elem() != abiMethodVarType.Elem() {
			// If the array is not a slice/array of structs, return an error.
			if implMethodVarType.Elem().Kind() != reflect.Struct {
				return fmt.Errorf(
					"type mismatch: %v != %v", implMethodVarType, abiMethodVarType,
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
	case reflect.Ptr:
		// The corresponding ABI type must be a struct.
		if abiMethodVarType.Kind() != reflect.Struct {
			return fmt.Errorf(
				"type mismatch: %v != %v", implMethodVarType, abiMethodVarType,
			)
		}

		// Any implementation type that is a pointer must point to a struct.
		if implMethodVarType.Elem().Kind() != reflect.Struct {
			return fmt.Errorf(
				"type mismatch: %v != %v", implMethodVarType, abiMethodVarType,
			)
		}

		// Check if the struct fields match.
		if err := validateStruct(implMethodVarType.Elem(), abiMethodVarType); err != nil {
			return err
		}
	case reflect.Interface:
		// If it's `any` (reflect.Interface), we leave it to the implementer to make sure that it is
		// used/converted correctly.
	default:
		return fmt.Errorf("type mismatch: %v != %v", implMethodVarType, abiMethodVarType)
	}

	return nil
}

// validateStruct checks to make sure that the implementation struct's fields match the ABI
// struct's fields.
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
	implMethodNumOut := implMethod.Type.NumOut()

	// The Solidity compiler requires that precompiles must return at least one value.
	// See https://github.com/berachain/polaris/issues/491 for more information.
	if len(abiMethod.Outputs) == 0 {
		//nolint:lll // error message.
		panic(fmt.Sprintf(
			"This precompile method %s must return at least one value (https://github.com/berachain/polaris/issues/491). Consider returning a boolean.",
			abiMethod.Name,
		))
	}

	// Last parameter of Go precompile implementation is an error (for reverts), so we skip that.
	if implMethodNumOut-1 != len(abiMethod.Outputs) {
		return fmt.Errorf(
			"number of return args mismatch: %v expects %v return vals, %v returns %v vals",
			abiMethod.Name,
			len(abiMethod.Outputs),
			implMethod.Name, implMethodNumOut-1,
		)
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
				"return type mismatch: %v expects %v, %v has %v",
				abiMethod.Name,
				abiMethodReturnType,
				implMethod.Name,
				implMethodReturnType,
			)
		}
	}

	return nil
}
