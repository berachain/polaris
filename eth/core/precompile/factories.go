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
	"errors"
	"fmt"
	"reflect"
	"unicode"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"pkg.berachain.dev/polaris/eth/core/vm"
	errorslib "pkg.berachain.dev/polaris/lib/errors"
	"pkg.berachain.dev/polaris/lib/utils"
)

const (
	// container impl names stored as constants, to be used in error messages.
	statelessContainerName = `StatelessContainerImpl`
	statefulContainerName  = `StatefulContainerImpl`
)

// AbstractFactory is an interface that all precompile container factories must adhere to.
type AbstractFactory interface {
	// Build builds and returns the precompile container for the type of container/factory.
	Build(Registrable, Plugin) (vm.PrecompileContainer, error)
}

// Compile-time assertions to ensure these container factories adhere to `AbstractFactory`.
var (
	_ AbstractFactory = (*StatelessFactory)(nil)
	_ AbstractFactory = (*StatefulFactory)(nil)
)

// ===========================================================================
// Stateless Container Factory
// ===========================================================================

// StatelessFactory is used to build stateless precompile containers.
type StatelessFactory struct{}

// NewStatelessFactory creates and returns a new `StatelessFactory`.
func NewStatelessFactory() *StatelessFactory {
	return &StatelessFactory{}
}

// Build returns a stateless precompile container for the given base contract implememntation.
// This function will return an error if the given contract is not a stateless implementation.
//
// Build implements `AbstractFactory`.
func (sf *StatelessFactory) Build(
	rp Registrable, _ Plugin,
) (vm.PrecompileContainer, error) {
	pc, ok := utils.GetAs[StatelessImpl](rp)
	if !ok {
		return nil, errorslib.Wrap(ErrWrongContainerFactory, statelessContainerName)
	}
	return pc, nil
}

// ===========================================================================
// Stateful Container Factory
// ===========================================================================

// StatefulFactory is used to build stateful precompile containers.
type StatefulFactory struct{}

// NewStatefulFactory creates and returns a new `StatefulFactory`.
func NewStatefulFactory() *StatefulFactory {
	return &StatefulFactory{}
}

// Build returns a stateful precompile container for the given base contract implementation. This
// function will return an error if the given contract is not a stateful implementation.
//
// Build implements `AbstractFactory`.
func (sf *StatefulFactory) Build(
	rp Registrable, p Plugin,
) (vm.PrecompileContainer, error) {
	si, ok := utils.GetAs[StatefulImpl](rp)
	if !ok {
		return nil, errorslib.Wrap(ErrWrongContainerFactory, statefulContainerName)
	}

	// attach the precompile plugin to the stateful contract
	si.SetPlugin(p)

	// add precompile methods to stateful container, if any exist
	idsToMethods, err := buildIdsToMethods(si.ABIMethods(), reflect.ValueOf(si))
	if err != nil {
		return nil, err
	}

	return NewStateful(si, idsToMethods)
}

// This function matches each Go implementation of the precompile to the ABI's respective function.
// It searches for the ABI function in the Go precompile contract and performs basic validation on
// the implemented function.
func buildIdsToMethods(
	pcABI map[string]abi.Method,
	contractImpl reflect.Value,
) (map[string]*Method, error) {
	contractImplType := contractImpl.Type()
	idsToMethods := make(map[string]*Method)
	for m := 0; m < contractImplType.NumMethod(); m++ {
		implMethod := contractImplType.Method(m)
		implMethodName := formatName(implMethod.Name)

		if abiMethod, found := pcABI[implMethodName]; found {
			if err := validateArgumentAndReturnTypes(implMethod, abiMethod); err != nil {
				return nil, errorslib.Wrap(err, implMethodName)
			}
			idsToMethods[utils.UnsafeBytesToStr(abiMethod.ID)] = NewMethod(
				&abiMethod,
				abiMethod.Sig,
				implMethod.Func,
			)
		}
	}

	// verify that every abi method has a corresponding precompile implementation
	for _, abiMethod := range pcABI {
		if _, found := idsToMethods[utils.UnsafeBytesToStr(abiMethod.ID)]; !found {
			return nil, errorslib.Wrap(ErrNoPrecompileMethodForABIMethod, abiMethod.Name)
		}
	}

	return idsToMethods, nil
}

// validateArgumentAndReturnTypes checks if the precompile method argument and return types match
// the abi's argument return types. this is for overloaded Solidity functions and general
// validation.
func validateArgumentAndReturnTypes(implMethod reflect.Method, abiMethod abi.Method) error {
	// reserve a return value for possible reverts/errors
	if implMethod.Type.NumOut()-1 != len(abiMethod.Outputs) {
		return errors.New("number of return types mismatch")
	}
	// receiver is first param, context is second param, check the actual function params
	j := 0
	for i := 2; i < implMethod.Type.NumIn(); i++ {
		implMethodParamType := implMethod.Type.In(i)
		abiMethodParamType := abiMethod.Inputs[j].Type.GetType()
		validate(implMethodParamType, abiMethodParamType)
		j++
	}

	for i := 0; i < implMethod.Type.NumOut()-1; i++ {
		implMethodReturnType := implMethod.Type.Out(i)
		abiMethodReturnType := abiMethod.Outputs[i].Type.GetType()
		validate(implMethodReturnType, abiMethodReturnType)
	}

	return nil
}

// helper function for validateArgumentAndReturnTypes. this function function uses reflection to see
// what types your implementation uses, and checks against the geth representation of the abi types.
func validate(implMethodVarType reflect.Type, abiMethodVarType reflect.Type) error {
	//nolint:exhaustive // nah, this is fine.
	switch implMethodVarType.Kind() {
	// we need to make sure that the struct fields match.
	case abiMethodVarType.Kind():
		return nil
	case reflect.Struct:
		if err := validateStructFields(implMethodVarType, abiMethodVarType); err != nil {
			return err
		}
	case reflect.Slice, reflect.Array:
		for j := 0; j < abiMethodVarType.Len(); j++ {
			// if it is a struct, then we need to check if the struct fields match
			if abiMethodVarType.Elem().Kind() == reflect.Struct {
				if err := validateStructFields(
					implMethodVarType.Elem(),
					abiMethodVarType.Elem(),
				); err != nil {
					return err
				}
			} else if implMethodVarType.In(j) != abiMethodVarType.In(j) {
				return fmt.Errorf("return type mismatch: %v != %v",
					implMethodVarType.Elem(),
					abiMethodVarType.Elem(),
				)
			}
		}
	default:
		return fmt.Errorf("return type mismatch: %v != %v", implMethodVarType, abiMethodVarType)
	}
	return nil
}

// this function checks to make sure that the struct fields match. if there is a nested struct, then
// we use recursion until we reach the base case of primitive types.
func validateStructFields(implMethodVarType reflect.Type,
	abiMethodVarType reflect.Type,
) error {
	if implMethodVarType == nil && abiMethodVarType == nil {
		return nil
	}
	if implMethodVarType.NumField() != abiMethodVarType.NumField() {
		return errors.New("number of return types mismatch")
	}
	for j := 0; j < implMethodVarType.NumField(); j++ {
		// if the field is a nested struct, then we recurse
		//nolint:gocritic // nah, this is fine.
		if implMethodVarType.Field(j).Type.Kind() == reflect.Struct &&
			abiMethodVarType.Field(j).Type.Kind() == reflect.Struct {
			if err := validateStructFields(
				implMethodVarType.Field(j).Type,
				abiMethodVarType.Field(j).Type,
			); err != nil {
				return err
			}
		} else if implMethodVarType.Field(j).Type != abiMethodVarType.Field(j).Type {
			return fmt.Errorf("return type mismatch: %v != %v",
				implMethodVarType.Field(j).Type,
				abiMethodVarType.Field(j).Type,
			)
		} else {
			continue
		}
	}
	return nil
}

// formatName converts to first character of name to lowercase.
func formatName(name string) string {
	if len(name) == 0 {
		return name
	}

	ret := []rune(name)
	ret[0] = unicode.ToLower(ret[0])
	return string(ret)
}
