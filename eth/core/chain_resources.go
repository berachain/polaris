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

package core

import (
	"errors"
	"fmt"

	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/core/state"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/lib/utils"
)

// ChainResources is the interface that defines functions for code paths within the chain to
// acquire resources to use in execution such as StateDBss and EVMss.
type ChainResources interface {
	StateAtBlockNumber(uint64) (state.StateDBI, error)
	StateAt(root common.Hash) (state.StateDBI, error)
	GetVMConfig() *vm.Config
}

// StateAt returns a statedb configured to read what the state of the blockchain is/was at a given.
func (bc *blockchain) StateAt(common.Hash) (state.StateDBI, error) {
	return nil, errors.New("not implemented")
}

// StateAtBlockNumber returns a statedb configured to read what the state of the blockchain is/was
// at a given block number.
func (bc *blockchain) StateAtBlockNumber(number uint64) (state.StateDBI, error) {
	sp, err := bc.sp.StateAtBlockNumber(number)
	if err != nil {
		return nil, err
	}
	return state.NewStateDB(sp, bc.pp), nil
}

// GetVMConfig returns the vm.Config for the current chain.
func (bc *blockchain) GetVMConfig() *vm.Config {
	return bc.vmConfig
}

// BuildAndRegisterPrecompiles builds the given precompiles and registers them with the precompile
// plugin.
// TODO: move precompile registration out of the state processor?
func (bc *blockchain) BuildAndRegisterPrecompiles(precompiles []precompile.Registrable) {
	for _, pc := range precompiles {
		// skip registering precompiles that are already registered.
		if bc.pp.Has(pc.RegistryKey()) {
			continue
		}

		// choose the appropriate precompile factory
		var af precompile.AbstractFactory
		switch {
		case utils.Implements[precompile.StatefulImpl](pc):
			af = precompile.NewStatefulFactory()
		case utils.Implements[precompile.StatelessImpl](pc):
			af = precompile.NewStatelessFactory()
		default:
			panic(
				fmt.Sprintf(
					"native precompile %s not properly implemented", pc.RegistryKey().Hex(),
				),
			)
		}

		// build the precompile container and register with the plugin
		container, err := af.Build(pc, bc.pp)
		if err != nil {
			panic(err)
		}
		// TODO: set code on the statedb for every precompiled contract.
		err = bc.pp.Register(container)
		if err != nil {
			panic(err)
		}
	}
}
