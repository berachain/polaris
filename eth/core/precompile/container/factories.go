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

package container

import (
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/eth/types/abi"
	"github.com/berachain/stargazer/lib/errors"
	"github.com/berachain/stargazer/lib/utils"
)

const (
	// container impl names stored as constants, to be used in error messages.
	statelessContainerName = `StatelessContainerImpl`
	statefulContainerName  = `StatefulContainerImpl`
	dynamicContainerName   = `DynamicContainerImpl`
)

// Compile-time assertions to ensure these container factories adhere to `AbstractFactory`.
var (
	_ AbstractFactory = (*StatelessFactory)(nil)
	_ AbstractFactory = (*StatefulFactory)(nil)
	_ AbstractFactory = (*DynamicFactory)(nil)
)

// ===========================================================================
// Stateless Container Factory
// ===========================================================================

// `StatelessFactory` is used to build stateless precompile containers.
type StatelessFactory struct{}

// `NewStatelessFactory` creates and returns a new `StatelessFactory`.
func NewStatelessFactory() *StatelessFactory {
	return &StatelessFactory{}
}

// `Build` returns a stateless precompile container for the given base contract implememntation.
// This function will return an error if the given contract is not a stateless implementation.
//
// `Build` implements `AbstractFactory`.
func (sf *StatelessFactory) Build(
	rp vm.RegistrablePrecompile,
) (vm.PrecompileContainer, error) {
	pc, ok := utils.GetAs[StatelessPrecompileImpl](rp)
	if !ok {
		return nil, errors.Wrap(ErrWrongContainerFactory, statelessContainerName)
	}
	return pc, nil
}

// ===========================================================================
// Stateful Container Factory
// ===========================================================================

// `StatefulFactory` is used to build stateful precompile containers.
type StatefulFactory struct {
}

// `NewStatefulFactory` creates and returns a new `StatefulFactory`.
func NewStatefulFactory() *StatefulFactory {
	return &StatefulFactory{}
}

// `Build` returns a stateful precompile container for the given base contract implementation.
// This function will return an error if the given contract is not a stateful implementation.
//
// `Build` implements `AbstractFactory`.
func (sf *StatefulFactory) Build(
	rp vm.RegistrablePrecompile,
) (vm.PrecompileContainer, error) {
	sci, ok := utils.GetAs[StatefulPrecompileImpl](rp)
	if !ok {
		return nil, errors.Wrap(ErrWrongContainerFactory, statefulContainerName)
	}

	var err error

	// add precompile methods to stateful container, if any exist
	var idsToMethods map[string]*Method
	if precompileMethods := sci.PrecompileMethods(); precompileMethods != nil {
		idsToMethods, err = sf.buildIdsToMethods(precompileMethods, sci.ABIMethods())
		if err != nil {
			return nil, err
		}
	}

	return NewStateful(rp, idsToMethods), nil
}

// `buildIdsToMethods` builds the stateful precompile container for the given `precompileMethods`
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

// ===========================================================================
// Dynamic Container Factory
// ===========================================================================

// `DynamicFactory` is used to build dynamic precompile containers.
type DynamicFactory struct {
	*StatefulFactory
}

// `NewDynamicFactory` creates and returns a new `DynamicFactory` for the given
// log registry `lr`.
func NewDynamicFactory() *DynamicFactory {
	return &DynamicFactory{
		StatefulFactory: NewStatefulFactory(),
	}
}

// `Build` returns a dynamic precompile container for the given base contract implememntation.
// This function will return an error if the given contract is not a dyanmic implementation.
//
// `Build` implements `AbstractFactory`.
func (dcf *DynamicFactory) Build(
	rp vm.RegistrablePrecompile,
) (vm.PrecompileContainer, error) {
	dci, ok := utils.GetAs[DynamicPrecompileImpl](rp)
	if !ok {
		return nil, errors.Wrap(ErrWrongContainerFactory, dynamicContainerName)
	}

	return dcf.StatefulFactory.Build(dci)
}
