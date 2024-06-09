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
	"reflect"

	errorslib "github.com/berachain/polaris/lib/errors"
	"github.com/berachain/polaris/lib/utils"

	"github.com/ethereum/go-ethereum/core/vm"
)

const (
	// impl names stored as constants, to be used in error messages.
	statelessContainerName = `StatelessImpl`
	statefulContainerName  = `StatefulImpl`
)

// AbstractFactory is an interface that all precompile container factories must adhere to.
type AbstractFactory interface {
	// Build builds and returns the precompile container for the type of container/factory.
	Build(Registrable, Plugin) (vm.PrecompiledContract, error)
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

// Build returns a stateless precompile container for the given base contract implementation.
// This function will return an error if the given contract is not a stateless implementation.
//
// Build implements `AbstractFactory`.
func (sf *StatelessFactory) Build(
	rp Registrable, _ Plugin,
) (vm.PrecompiledContract, error) {
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
) (vm.PrecompiledContract, error) {
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

	return NewStatefulContainer(si, idsToMethods)
}

// This function matches each Go implementation of the precompile to the ABI's respective function.
// It searches for the ABI function in the Go precompile contract and performs basic validation on
// the implemented function.
func buildIdsToMethods(si StatefulImpl, contractImpl reflect.Value) (map[methodID]*method, error) {
	precompileABI := si.ABIMethods()
	contractImplType := contractImpl.Type()
	idsToMethods := make(map[methodID]*method)
	for m := 0; m < contractImplType.NumMethod(); m++ {
		implMethod := contractImplType.Method(m)

		methodName, err := findMatchingABIMethod(implMethod, precompileABI)
		if err != nil {
			return nil, err
		}
		if methodName == "" {
			continue // nothing in the abi matches our go method.
		}

		method := newMethod(si, precompileABI[methodName], implMethod)
		idsToMethods[methodID(precompileABI[methodName].ID)] = method
	}

	// verify that every abi method has a corresponding precompile implementation
	for _, abiMethod := range precompileABI {
		if _, found := idsToMethods[methodID(abiMethod.ID)]; !found {
			return nil, errorslib.Wrap(ErrNoPrecompileMethodForABIMethod, abiMethod.Name)
		}
	}

	return idsToMethods, nil
}
