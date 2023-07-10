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

// Build returns a stateful precompile container for the given base contract implementation.
// This function will return an error if the given contract is not a stateful implementation.
//
// Build implements `AbstractFactory`.
func (sf *StatefulFactory) Build(
	rp Registrable, p Plugin,
) (vm.PrecompileContainer, error) {
	sci, ok := utils.GetAs[StatefulImpl](rp)
	if !ok {
		return nil, errorslib.Wrap(ErrWrongContainerFactory, statefulContainerName)
	}

	// attach the precompile plugin to the stateful contract
	sci.SetPlugin(p)

	// add precompile methods to stateful container, if any exist
	var idsToMethods map[string]*Method
	idsToMethods, err := BuildIdsToMethods(sci.ABIMethods(), reflect.ValueOf(sci))
	if err != nil {
		return nil, err
	}

	return NewStateful(rp, idsToMethods)
}

// This function matches each Go implementation of the Precompile
// to the ABI's respective function.
// It first searches for the ABI function in the Go implementation. If no find, then panic.
// It then performs some basic validation on the implemented function
// Then, the implemented function's arguments are checked against the ABI's arguments' types.
func BuildIdsToMethods(pcABI map[string]abi.Method, contractImpl reflect.Value) (map[string]*Method, error) {
	contractImplType := contractImpl.Type()
	idsToMethods := make(map[string]*Method)
	for m := 0; m < contractImplType.NumMethod(); m++ {
		implMethod := contractImplType.Method(m) // grab the Impl's current method

		implMethodName := formatName(implMethod.Name)

		if abiMethod, found := pcABI[implMethodName]; found {
			if err := checkReturnTypes(implMethod); err != nil {
				return nil, errorslib.Wrap(err, implMethodName)
			}
			idsToMethods[utils.UnsafeBytesToStr(abiMethod.ID)] = &Method{
				abiMethod: &abiMethod,
				abiSig:    abiMethod.Sig,
				execute:   implMethod.Func,
			}
		}
	}

	for _, abiMethod := range pcABI {
		if _, found := idsToMethods[utils.UnsafeBytesToStr(abiMethod.ID)]; !found {
			return nil, errorslib.Wrap(ErrNoPrecompileMethodForABIMethod, abiMethod.Name)
		}
	}

	return idsToMethods, nil
}

// formatName converts to first character of name to lowercase. If the first
// three characters are "ERC" or "ABI" (which is p common), then it converts all
// three to lowercase.
func formatName(name string) string {
	ret := []rune(name)
	if name[:3] == "ERC" || name[:3] == "ABI" { // special case for ERC20, ERC721, etc.
		ret[0] = unicode.ToLower(ret[0])
		ret[1] = unicode.ToLower(ret[1])
		ret[2] = unicode.ToLower(ret[2])
	} else if len(ret) > 0 {
		ret[0] = unicode.ToLower(ret[0])
	}

	return string(ret)
}

// checkReturnTypes checks if the precompile method returns a []any and an error.
// If it does not, then an error is returned.
func checkReturnTypes(implMethod reflect.Method) error {
	if implMethod.Type.NumOut() != 2 { //nolint:gomnd // it's okay.
		return errors.New("precompile methods must return ([]any, error), but found wrong number of return types for precompile method: " + //nolint:lll // it's okay.
			implMethod.Name)
	}
	firstReturnType := implMethod.Type.Out(0)
	secondReturnType := implMethod.Type.Out(1)

	if firstReturnType.Kind() != reflect.Slice { // check if the first return type is a []any
		return errors.New("first parameter should be []any, but found " +
			firstReturnType.String() + " for precompile method: " + implMethod.Name)
	} else if firstReturnType.Elem().Kind() != reflect.Interface { // if it is but it is not an any...
		return errors.New("first parameter should be []any, but found " +
			firstReturnType.String() + " for precompile method: " + implMethod.Name)
	}

	if secondReturnType.Name() != "error" { // if the second return value is not an error
		return errors.New("second parameter should be error, but found " +
			secondReturnType.String() + " for precompile method: " + implMethod.Name)
	}
	return nil
}
