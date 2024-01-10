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
