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
	"sync"

	"pkg.berachain.dev/stargazer/eth/core/precompile"
	"pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/eth/core/vm"
	"pkg.berachain.dev/stargazer/eth/crypto"
	"pkg.berachain.dev/stargazer/eth/params"
	"pkg.berachain.dev/stargazer/lib/errors"
	"pkg.berachain.dev/stargazer/lib/utils"
)

// `StateProcessor` is responsible for processing blocks, transactions, and updating the state.
type StateProcessor struct {
	// `mtx` is used to make sure we don't try to prepare a new block before finalizing the
	// previous one.
	mtx sync.Mutex

	// `bp` provides block functions from the underlying chain the EVM is running on
	bp BlockPlugin
	// `gp` provides gas functions from the underlying chain the EVM is running on
	gp GasPlugin
	// `cp` provides configuration functions from the underlying chain the EVM is running on
	cp ConfigurationPlugin
	// `pp` is responsible for keeping track of the stateful precompile containers that are
	// available to the EVM and executing them.
	pp PrecompilePlugin

	// `vmConfig` is the configuration for the EVM.
	vmConfig vm.Config
	// `statedb` is the state database that is used to mange state during transactions.
	statedb vm.StargazerStateDB
	// `commit` indicates whether the state processor should commit the state after processing a tx
	commit bool

	// `signer` is the signer used to verify transaction signatures.
	signer types.Signer
	// `evm ` is the EVM that is used to process transactions.
	evm vm.StargazerEVM
	// `block` represents the current block being processed.
	block *types.StargazerBlock
	// `finalizedBlock` represents the block that was last finalized by the state processor.
	finalizedBlock *types.StargazerBlock
}

// `NewStateProcessor` creates a new state processor with the given host, statedb, vmConfig, and
// commit flag.
func NewStateProcessor(
	host StargazerHostChain,
	statedb vm.StargazerStateDB,
	vmConfig vm.Config,
	commit bool,
) *StateProcessor {
	sp := &StateProcessor{
		mtx:      sync.Mutex{},
		bp:       host.GetBlockPlugin(),
		gp:       host.GetGasPlugin(),
		cp:       host.GetConfigurationPlugin(),
		pp:       host.GetPrecompilePlugin(),
		vmConfig: vmConfig,
		statedb:  statedb,
		commit:   commit,
	}

	if sp.pp == nil {
		sp.pp = precompile.NewDefaultPlugin()
	} else {
		// build and register the native precompile contracts
		sp.BuildPrecompiles(sp.pp.GetPrecompiles(params.Rules{}))
	}

	return sp
}

// ==============================================================================
// Block, Tx Lifecycle
// ==============================================================================

// `Prepare` prepares the state processor for processing a block.
func (sp *StateProcessor) Prepare(ctx context.Context, height int64) {
	// We lock the state processor as a safety measure to ensure that Prepare is not called again
	// before finalize.
	sp.mtx.Lock()

	// Prepare the plugins for the new block.
	sp.bp.Prepare(ctx)
	sp.cp.Prepare(ctx)
	sp.gp.Prepare(ctx)

	// Build a block object so we can track that status of the block as we process it.
	sp.block = types.NewStargazerBlock(sp.bp.GetStargazerHeaderByNumber(height))

	// Ensure that the gas plugin and header are in sync.
	if sp.block.GasLimit != sp.gp.BlockGasLimit() {
		panic(fmt.Sprintf("gas limit mismatch: have %d, want %d", sp.block.GasLimit, sp.gp.BlockGasLimit()))
	}

	// We must re-create the signer since we are processing a new block and the block number has increased.
	chainConfig := sp.cp.ChainConfig()
	sp.signer = types.MakeSigner(chainConfig, sp.block.Number)

	// Setup the EVM for this block.
	sp.BuildPrecompiles(
		precompile.NewDefaultPlugin().GetPrecompiles(
			chainConfig.Rules(sp.block.Number, true, sp.block.Time),
		),
	)
	sp.vmConfig.ExtraEips = sp.cp.ExtraEips()
	sp.evm = vm.NewStargazerEVM(
		sp.NewEVMBlockContext(),
		vm.TxContext{},
		sp.statedb,
		chainConfig,
		sp.vmConfig,
		sp.pp,
	)
}

// `ProcessTransaction` applies a transaction to the current state of the blockchain.
func (sp *StateProcessor) ProcessTransaction(
	ctx context.Context, tx *types.Transaction,
) (*types.Receipt, error) {
	msg, err := tx.AsMessage(sp.signer, sp.block.BaseFee)
	if err != nil {
		return nil, errors.Wrapf(err, "could not apply tx %d [%v]", sp.block.TxIndex(), tx.Hash().Hex())
	}

	// Create a new context to be used in the EVM environment and tx context for the StateDB.
	txContext := NewEVMTxContext(msg)
	sp.evm.SetTxContext(txContext)
	txHash := tx.Hash()
	sp.statedb.SetTxContext(txHash, sp.block.TxIndex())

	// We also must reset the StateDB and precompile and gas plugins.
	sp.statedb.Reset(ctx)
	sp.pp.Reset(ctx)
	sp.gp.Reset(ctx)

	// Set the gasPool to have the remaining gas in the block.
	// ASSUMPTION: That the host chain has not consumped the intrinsic gas yet.
	gasPool := GasPool(sp.gp.BlockGasLimit() - sp.gp.CumulativeGasUsed())
	if err = sp.gp.SetTxGasLimit(msg.Gas()); err != nil {
		return nil, errors.Wrapf(err, "could not set gas plugin limit %d [%v]", sp.block.TxIndex(), tx.Hash().Hex())
	}

	// Apply the state transition.
	result, err := ApplyMessage(sp.evm.UnderlyingEVM(), msg, &gasPool)
	if err != nil {
		return nil, errors.Wrapf(err, "could not apply message %d [%v]", sp.block.TxIndex(), tx.Hash().Hex())
	}

	// Consume the gas used by the state tranisition.
	if err = sp.gp.TxConsumeGas(result.UsedGas); err != nil {
		return nil, errors.Wrapf(err, "could not consume gas used %d [%v]", sp.block.TxIndex(), tx.Hash().Hex())
	}

	// Create a new receipt for the transaction, storing the intermediate root and gas used
	// by the tx.
	receipt := &types.Receipt{
		Type:              tx.Type(),
		CumulativeGasUsed: sp.gp.CumulativeGasUsed(),
		TxHash:            txHash,
		GasUsed:           result.UsedGas,
		BlockHash:         sp.block.Hash(),
		BlockNumber:       sp.block.Number,
		TransactionIndex:  uint(sp.block.TxIndex()),
	}

	// If the transaction created a contract, store the creation address in the receipt.
	if msg.To() == nil {
		receipt.ContractAddress = crypto.CreateAddress(txContext.Origin, tx.Nonce())
	}

	// Set the receipt logs and create the bloom filter.
	receipt.Logs = sp.statedb.GetLogs(
		receipt.TxHash, sp.block.Number.Uint64(), sp.block.Hash(),
	)
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})

	if result.Failed() {
		receipt.Status = types.ReceiptStatusFailed
	} else {
		// if the result didn't produce a consensus error then we can properly commit the state.
		if sp.commit {
			sp.evm.StateDB().Finalize()
		}
		receipt.Status = types.ReceiptStatusSuccessful
	}

	// Update the block information.
	sp.block.AppendTx(tx, receipt)

	// Return receipt to the caller.
	return receipt, nil
}

// `Finalize` finalizes the block in the state processor and returns the receipts and bloom filter.
func (sp *StateProcessor) Finalize(ctx context.Context) (*types.StargazerBlock, error) {
	// We unlock the state processor to ensure that the state is consistent.
	defer sp.mtx.Unlock()

	sp.block.Finalize(sp.gp.CumulativeGasUsed())
	return sp.block, nil
}

// ===========================================================================
// Utilities
// ===========================================================================

// `NewEVMBlockContext` creates a new block context for use in the EVM.
func (sp *StateProcessor) NewEVMBlockContext() vm.BlockContext {
	feeCollector := sp.cp.FeeCollector()
	if feeCollector == nil {
		feeCollector = &sp.block.Coinbase
	}
	return NewEVMBlockContext(sp.block.Header, &chainContext{sp}, feeCollector)
}

// `BuildPrecompiles` builds the given precompiles and registers them with the precompile plugins.
func (sp *StateProcessor) BuildPrecompiles(precompiles []vm.RegistrablePrecompile) {
	for _, pc := range precompiles {
		// choose the appropriate precompile factory
		var af precompile.AbstractFactory
		if utils.Implements[precompile.DynamicImpl](pc) {
			af = precompile.NewDynamicFactory()
		} else if utils.Implements[precompile.StatefulImpl](pc) {
			af = precompile.NewStatefulFactory()
		} else if utils.Implements[precompile.StatelessImpl](pc) {
			af = precompile.NewStatelessFactory()
		} else {
			panic(
				fmt.Sprintf(
					"native precompile %s not properly implemented", pc.RegistryKey().Hex(),
				),
			)
		}

		// build the precompile container and register with the plugin
		container, err := af.Build(pc)
		if err != nil {
			panic(err)
		}
		sp.pp.Register(container)
	}
}

// `GetHashFn` returns a `GetHashFunc` which retrieves header hashes by number.
func (sp *StateProcessor) GetHashFn() vm.GetHashFunc {
	return GetHashFn(sp.block.Header, &chainContext{sp})
}
