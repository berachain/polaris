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
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/lib/errors"
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
		return nil, errors.Wrap(ErrWrongContainerFactory, statelessContainerName)
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

// Build returns a stateful precompile container for the given base contract implementation.
// This function will return an error if the given contract is not a stateful implementation.
//
// Build implements `AbstractFactory`.
func (sf *StatefulFactory) Build(
	rp Registrable, p Plugin,
) (vm.PrecompileContainer, error) {
	sci, ok := utils.GetAs[StatefulImpl](rp)
	if !ok {
		return nil, errors.Wrap(ErrWrongContainerFactory, statefulContainerName)
	}

	// attach the precompile plugin to the stateful contract
	sci.SetPlugin(p)

	var err error

	// add precompile methods to stateful container, if any exist
	var idsToMethods map[string]*Method
	if precompileMethods := GeneratePrecompileMethods(sci.ABIMethods(), reflect.ValueOf(sci)); precompileMethods != nil {
		idsToMethods, err = sf.buildIdsToMethods(precompileMethods, sci.ABIMethods())
		if err != nil {
			return nil, err
		}
	}

	return NewStateful(rp, idsToMethods), nil
}

// buildIdsToMethods builds the stateful precompile container for the given `precompileMethods`
// and `abiMethods`. This function will return an error if every method in `abiMethods` does not
// have a valid, corresponding `Method`.
func (sf *StatefulFactory) buildIdsToMethods(
	precompileMethods Methods,
	abiMethods map[string]abi.Method,
) (map[string]*Method, error) {
	// validate precompile methods
	for _, pm := range precompileMethods {
		if err := pm.ValidateBasic(); err != nil {
			return nil, err
		}
	}

	// match every ABI method to corresponding precompile method
	idsToMethods := make(map[string]*Method)
	for name := range abiMethods {
		abiMethod := abiMethods[name]

		// find the corresponding precompile method for abiMethod based on signature
		var precompileMethod *Method
		i := 0
		for ; i < len(precompileMethods); i++ {
			if precompileMethods[i].AbiSig == abiMethod.Sig {
				precompileMethod = precompileMethods[i]
				break
			}
		}
		if i == len(precompileMethods) {
			return nil, errors.Wrap(ErrNoPrecompileMethodForABIMethod, abiMethod.Sig)
		}

		// attach the ABI method to the precompile method for stateful container to handle
		precompileMethod.AbiMethod = &abiMethod
		idsToMethods[utils.UnsafeBytesToStr(abiMethod.ID)] = precompileMethod
	}
	return idsToMethods, nil
}

// GeneratePrecompileMethods generates the methods for the given Precompile's ABI.
func GeneratePrecompileMethods(ABI map[string]abi.Method, contractImpl reflect.Value) Methods {
	return suitableMethods(ABI, contractImpl.Type())
}

// This function matches each Go implementation of the Precompile
// to the ABI's respective function.
// It first searches for the ABI function in the Go implementation. If no find, then panic.
// It then performs some basic validation on the implemented function
// Then, the implemented function's arguments are checked against the ABI's arguments' types.
func suitableMethods(pcABI map[string]abi.Method, contractImpl reflect.Type) Methods {
	var output Methods

	// populate allMethods with all of the contract implementation's methods
	allMethods := make([]reflect.Method, 0, contractImpl.NumMethod())
	for m := 0; m < contractImpl.NumMethod(); m++ {
		allMethods[m] = contractImpl.Method(m)
	}

	// match every method from ABI to allMethods
	for _, abiMethod := range pcABI {
		method, err := matchMethod(abiMethod, allMethods)
		if err != nil {
			panic(err)
		}
		output = append(output,
			&Method{
				AbiMethod: &abiMethod,
				AbiSig:    abiMethod.Sig,
				Execute:   method,
			})
	}
	return nil
}

// matchMethod matches the given ABI method to the corresponding method in the given list of methods.
// This function will return an error if the given ABI method does not have a corresponding method
// in the given list of methods.
func matchMethod(toCheck abi.Method, allMethods []reflect.Method) (reflect.Value, error) {
	name := formatName(toCheck.Name)

	// create methods subset by prefix matching
	var matchedMethods []reflect.Method
	for _, method := range allMethods {
		if strings.HasPrefix(method.Name, name) {
			matchedMethods = append(matchedMethods, method)
		}
	}

	// for method in subset, check if params match
	for _, method := range matchedMethods {
		if matchParams(toCheck, method) {
			return method.Func, nil
		}
	}
	return reflect.Value{}, errors.Wrap(ErrNoPrecompileMethodForABIMethod, toCheck.Sig)
}

// matchParams matches the given ABI method's parameters to the given method's parameters.
// This function will return true if the given ABI method's parameters match the given method's
// parameters and false otherwise.
func matchParams(abiMethod abi.Method, implMethod reflect.Method) bool {
	// uncomment when we change all the function signatures to match the abi method params
	implMethodIdx := 2 // start at 2 as 0th params should be a receiver, and 1 is the PolarContext
	for i := 0; i < len(abiMethod.Inputs); i++ {
		if implMethod.Type.In(implMethodIdx) != abiMethod.Inputs[i].Type.GetType() {
			// hold on, this doesn't indicate a failure due to overloaded functions
			fmt.Println("implMethod.Type.In(implMethodIdx): ", implMethod.Type.In(implMethodIdx), "abiMethod.Inputs[i].Type.GetType(): ", abiMethod.Inputs[i].Type.GetType())
			// panic("does not match types")
			return false
		}
		implMethodIdx++
	}
	return true
}

// this is a helper function that checks three things:
// 1. the first parameter is a context.Context.
// 2. the number of arguments match.
// 3. the types of the arguments match.
func basicValidation(_ reflect.Method, _ abi.Method) error {
	// if implMethod.Type.In(1) != reflect.TypeOf((*PolarContext)(nil)).Elem() {
	// return errors.Wrap(ErrNoContext, abiMethod.Sig)
	// } else if implMethod.Type.NumIn()-1 != len(abiMethod.Inputs) {
	// // return errors.Wrap(ErrNoPrecompileMethodForABIMethod, abiMethod.Sig)
	// fmt.Println("TODO: fix later")
	// }
	return nil
}

func newExecute(fn reflect.Method) reflect.Value {
	return fn.Func
}

// formatName converts to first character of name to uppercase.
// If the first three characters are "erc" (which is p common), then it converts all three to uppercase.
// the code below has been inspired by Geth.
func formatName(name string) string {
	ret := []rune(name)
	if name[:3] == "erc" { // special case for erc20, erc721, etc.
		ret[0] = unicode.ToUpper(ret[0])
		ret[1] = unicode.ToUpper(ret[1])
		ret[2] = unicode.ToUpper(ret[2])
	} else if len(ret) > 0 {
		ret[0] = unicode.ToUpper(ret[0])
	}

	return string(ret)
}
