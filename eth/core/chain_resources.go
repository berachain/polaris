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
	"math/big"

	"github.com/ethereum/go-ethereum/consensus/misc"

	"pkg.berachain.dev/polaris/eth/core/state"
	"pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/eth/params"
)

// ChainResources is the interface that defines functions for code paths within the chain to acquire
// resources to use in execution such as StateDBss and EVMss.
type ChainResources interface {
	StateAtBlockNumber(uint64) (vm.GethStateDB, error)
	GetVMConfig() *vm.Config
	GetEVM(context.Context, vm.TxContext, vm.PolarisStateDB, *types.Header, *vm.Config) *vm.GethEVM
	NewEVMBlockContext(header *types.Header) *vm.BlockContext
}

// StateAtBlockNumber returns a statedb configured to read what the state of the blockchain is/was
// at a given block number.
func (bc *blockchain) StateAtBlockNumber(number uint64) (vm.GethStateDB, error) {
	sp, err := bc.sp.StateAtBlockNumber(number)
	if err != nil {
		return nil, err
	}
	return state.NewStateDB(sp), nil
}

// GetEVM returns an EVM ready to be used for executing transactions. It is used by both the
// StateProcessor to acquire a new EVM at the start of every block. As well as by the backend to
// acquire an EVM for running gas estimations, eth_call etc.
func (bc *blockchain) GetEVM(
	_ context.Context, txContext vm.TxContext, state vm.PolarisStateDB,
	header *types.Header, vmConfig *vm.Config,
) *vm.GethEVM {
	chainCfg := bc.processor.cp.ChainConfig() // TODO: get chain config at height.
	return vm.NewGethEVMWithPrecompiles(
		*bc.NewEVMBlockContext(header), txContext, state, chainCfg, *vmConfig, bc.processor.pp,
	)
}

// NewEVMBlockContext creates a new block context for use in the EVM.
func (bc *blockchain) NewEVMBlockContext(header *types.Header) *vm.BlockContext {
	feeCollector := bc.cp.FeeCollector()
	if feeCollector == nil {
		feeCollector = &header.Coinbase
	}
	blockContext := NewEVMBlockContext(header, bc, feeCollector)
	return &blockContext
}

// GetVMConfig returns the vm.Config for the current chain.
func (bc *blockchain) GetVMConfig() *vm.Config {
	return bc.vmConfig
}

// CalculateBaseFee calculates the base fee for the next block based on the finalized block or the
// plugin's base fee.
func (bc *blockchain) CalculateNextBaseFee(number uint64) *big.Int {
	// TODO: ONCE WE ARE STORING GENESIS HEADER CORRECTLY, WE SHOULD REMOVE THIS HACK
	// TODO: THE BELOW GETHEADERBYNUMBER SHOULD WORK ON 0
	if number == 0 {
		return big.NewInt(int64(params.InitialBaseFee))
	}

	if pluginBaseFee := bc.bp.BaseFee(); pluginBaseFee.Cmp(new(big.Int)) >= 0 /* non-negative */ {
		return pluginBaseFee
	}

	// If the base fee supplied by the plugins is negative, then we assume that the host chain
	// wants to use the built-in EIP-1559 math.
	if parent := bc.finalizedBlock.Load(); parent != nil {
		// If the base fee supplied by the plugins is non-negative, then we assume that the host
		// chain wants to use the base fee supplied by the plugin.
		misc.CalcBaseFee(bc.Config(), parent.Header())
	}

	// We check the block plugin to see if it has a base fee for the previous block.
	if parent, err := bc.bp.GetHeaderByNumber(number - 1); err == nil {
		return misc.CalcBaseFee(bc.Config(), parent)
	}

	// TODO: we should probably panic here to be safe, but tbd.
	return big.NewInt(int64(params.InitialBaseFee))
}
