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
	"fmt"

	"github.com/ethereum/go-ethereum/core"

	"pkg.berachain.dev/polaris/eth/core/state"
	"pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/eth/tracers"
	errorslib "pkg.berachain.dev/polaris/lib/errors"
)

// ChainResources is the interface that defines functions for code paths within the chain to acquire
// resources to use in execution such as StateDBss and EVMss.
type ChainResources interface {
	GetEVM(context.Context, vm.TxContext, vm.PolarisStateDB, *types.Header, *vm.Config) *vm.GethEVM
	GetStateByNumber(int64) (vm.PolarisStateDB, error)
	GetStateByTransaction(
		context.Context, *types.Block, int,
	) (*Message, vm.BlockContext, state.StateDBI, tracers.StateReleaseFunc, error)
}

// GetEVM returns an EVM ready to be used for executing transactions. It is used by both the StateProcessor
// to acquire a new EVM at the start of every block. As well as by the backend to acquire an EVM for running
// gas estimations, eth_call etc.
func (bc *blockchain) GetEVM(
	ctx context.Context, txContext vm.TxContext, state vm.PolarisStateDB,
	header *types.Header, vmConfig *vm.Config,
) *vm.GethEVM {
	chainCfg := bc.processor.cp.ChainConfig(ctx) // todo: get chain config at height.
	return vm.NewGethEVMWithPrecompiles(
		bc.newEVMBlockContext(ctx, header), txContext, state, chainCfg, *vmConfig, bc.processor.pp,
	)
}

// GetStateByNumber returns a statedb configured to read what the state of the blockchain is/was
// at a given block number.
func (bc *blockchain) GetStateByNumber(number int64) (vm.PolarisStateDB, error) {
	sp, err := bc.sp.GetStateByNumber(number)
	if err != nil {
		return nil, err
	}
	return state.NewStateDB(sp), nil
}

// GetStateByTransaction returns a statedb configured to read what the state of the blockchain is/was
// at a given transaction. It also returns the message, block context and a release function for the
// statedb.
func (bc *blockchain) GetStateByTransaction(_ context.Context, block *types.Block, txIndex int) (
	*Message, vm.BlockContext, state.StateDBI, tracers.StateReleaseFunc, error,
) {
	// get the statedb and state processor to execute the block
	statedb, err := bc.GetStateByNumber(block.Number().Int64())
	if err != nil {
		return nil, vm.BlockContext{}, nil, nil, err
	}
	vmConfig := &vm.Config{Debug: true}
	processor := NewStateProcessor(bc.cp, bc.pp, statedb, vmConfig)

	// prepare the state processor for executing the block
	ctx, header := statedb.GetContext(), block.Header()
	processor.Prepare(
		ctx,
		bc.GetEVM(ctx, vm.TxContext{}, statedb, header, vmConfig),
		header,
		bc.host.GetNewGasPlugin(ctx),
	)

	// execute all the transactions until we reach the txIndex
	for idx, tx := range block.Transactions() {
		txHash := tx.Hash()
		var msg *core.Message
		msg, err = TransactionToMessage(tx, processor.signer, header.BaseFee)
		if err != nil {
			return nil, vm.BlockContext{}, nil, nil,
				errorslib.Wrapf(
					err, "could not apply tx %d [%s] while tracing block %d",
					idx, txHash.Hex(), block.NumberU64(),
				)
		}
		txContext := NewEVMTxContext(msg)
		processor.evm.Reset(txContext, statedb)
		statedb.Reset(ctx)
		statedb.SetTxContext(txHash, idx)

		if idx == txIndex {
			// break condition
			return msg, processor.evm.Context, statedb, func() {}, nil
		}

		if _, err = ApplyMessage(processor.evm, msg, new(GasPool).AddGas(tx.Gas())); err != nil {
			return nil, vm.BlockContext{}, nil, nil,
				errorslib.Wrapf(
					err, "tx %d [%s] failed while tracing block %d",
					idx, txHash.Hex(), block.NumberU64(),
				)
		}

		statedb.Finalize()
	}

	return nil, vm.BlockContext{}, nil, nil,
		fmt.Errorf("txIndex %d out of range while tracing block %d", txIndex, block.NumberU64())
}

// NewEVMBlockContext creates a new block context for use in the EVM.
func (bc *blockchain) newEVMBlockContext(ctx context.Context, header *types.Header) vm.BlockContext {
	feeCollector := bc.cp.FeeCollector(ctx)
	if feeCollector == nil {
		feeCollector = &header.Coinbase
	}
	return NewEVMBlockContext(header, &chainContext{bc, ctx}, feeCollector)
}
