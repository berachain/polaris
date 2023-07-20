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

	"pkg.berachain.dev/polaris/eth/accounts/abi"
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
	idsToMethods, err := buildIdsToMethods(si, reflect.ValueOf(si))
	if err != nil {
		return nil, err
	}

	return NewStateful(si, idsToMethods)
}

// This function matches each Go implementation of the precompile to the ABI's respective function.
// It searches for the ABI function in the Go precompile contract and performs basic validation on
// the implemented function.
func buildIdsToMethods(
	si StatefulImpl,
	contractImpl reflect.Value,
) (map[string]*Method, error) {
	precompileABI := si.ABIMethods()
	contractImplType := contractImpl.Type()
	idsToMethods := make(map[string]*Method)
	for m := 0; m < contractImplType.NumMethod(); m++ {
		implMethod := contractImplType.Method(m)
		// This only runs when we have overloaded functions. In most cases, it's an O(1) operation.

		abiMethod := findInABI(implMethod, precompileABI)
		// we need to make sure that this is actually mapping the right implMethod by checking the
		// arg types due to the overloaded function edge case.
		if abiMethod.Name != "" {
			method := NewMethod(
				si,
				&abiMethod,
				abiMethod.Sig,
				implMethod,
			)

			if err := method.ValidateBasic(); err != nil {
				return nil, err
			} else {
				idsToMethods[utils.UnsafeBytesToStr(abiMethod.ID)] = method
			}
		}
	}

	// verify that every abi method has a corresponding precompile implementation
	for _, abiMethod := range precompileABI {
		if _, found := idsToMethods[utils.UnsafeBytesToStr(abiMethod.ID)]; !found {
			return nil, errorslib.Wrap(ErrNoPrecompileMethodForABIMethod, abiMethod.Name)
		}
	}

	return idsToMethods, nil
}

// Find returns the longest substring of `implMethodName` that is a key in `precompileABI`.
// This function is used to find the ABI method that corresponds to the Go implementation.
// and provides safeguarding against the overloaded function edge case.
func findInABI(implMethod reflect.Method, precompileABI map[string]abi.Method) abi.Method {
	implMethodName := formatName(implMethod.Name)

	for i := len(implMethodName); i > 0; i-- {
		for _, abiMethod := range precompileABI {
			// this could be the function
			if implMethodName == abiMethod.RawName {
				// same function name and same amount of arguments. if this case fails then
				// we either don't have the right function or we have an overloaded function.
				if err := validateInputs(implMethod, abiMethod); err == nil {
					return abiMethod
				}
			}
			implMethodName = implMethodName[:i]
		}
	}

	return abi.Method{}
}

func validateInputs(implMethod reflect.Method, abiMethod abi.Method) error {
	abiMethodNumIn := len(abiMethod.Inputs)

	// First two args of Go precompile implementation are the receiver contract and the Context, so
	// verify that the ABI method has exactly 2 fewer inputs than the implementation method.
	if implMethod.Type.NumIn()-2 != abiMethodNumIn {
		return errors.New("number of arguments mismatch")
	}

	// If the function does not take any inputs, no need to check.
	if abiMethodNumIn > 0 {
		// Validate that the precompile input args types match ABI input arg types, excluding the
		// first two args (receiver contract and Context).
		for i := 2; i < implMethod.Type.NumIn(); i++ {
			implMethodParamType := implMethod.Type.In(i)
			abiMethodParamType := abiMethod.Inputs[i-2].Type.GetType()
			if err := validateArg(implMethodParamType, abiMethodParamType); err != nil {
				fmt.Println("we error", err)
				return err
			}
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
