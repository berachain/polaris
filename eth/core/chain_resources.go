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
	"context"

	"pkg.berachain.dev/stargazer/eth/core/state"
	"pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/eth/core/vm"
)

// `ChainResources` is the interface that defines functions for code paths within the chain to acquire
// resources to use in execution such as StateDBss and EVMss.
type ChainResources interface {
	GetStateByNumber(int64) (vm.GethStateDB, error)
	GetEVM(context.Context, vm.TxContext, vm.StargazerStateDB, *types.Header, *vm.Config) *vm.GethEVM
}

// `GetStateByNumber` returns a statedb configured to read what the state of the blockchain is/was
// at a given block number.
func (bc *blockchain) GetStateByNumber(number int64) (vm.GethStateDB, error) {
	sp, err := bc.host.GetStatePlugin().GetStateByNumber(number)
	if err != nil {
		return nil, err
	}
	return state.NewStateDB(sp), nil
}

// `GetEVM` returns an EVM ready to be used for executing transactions. It is used by both the StateProcessor
// to acquire a new EVM at the start of every block. As well as by the backend to acquire an EVM for running
// gas estimations, eth_call etc.
func (bc *blockchain) GetEVM(
	_ context.Context, txContext vm.TxContext, state vm.StargazerStateDB,
	header *types.Header, vmConfig *vm.Config,
) *vm.GethEVM {
	chainCfg := bc.processor.cp.ChainConfig() // todo: get chain config at height.
	return vm.NewGethEVMWithPrecompiles(
		bc.NewEVMBlockContext(header), txContext, state, chainCfg, *vmConfig, bc.processor.pp,
	)
}

// `NewEVMBlockContext` creates a new block context for use in the EVM.
func (bc *blockchain) NewEVMBlockContext(header *types.Header) vm.BlockContext {
	feeCollector := bc.host.GetConfigurationPlugin().FeeCollector()
	if feeCollector == nil {
		feeCollector = &header.Coinbase
	}
	return NewEVMBlockContext(header, &chainContext{bc}, feeCollector)
}
