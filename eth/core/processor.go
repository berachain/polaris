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
	"math/big"
	"sync"

	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/eth/core/state"
	"github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/eth/crypto"
	"github.com/berachain/stargazer/lib/errors"
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
	sp.block = types.NewStargazerBlock(sp.bp.GetStargazerHeaderAtHeight(height))

	// Ensure that the gas plugin and header are in sync.
	if sp.block.GasLimit != sp.gp.BlockGasLimit() {
		panic(fmt.Sprintf("gas limit mismatch: have %d, want %d", sp.block.GasLimit, sp.gp.BlockGasLimit()))
	}

	// We must re-create the signer since we are processing a new block and the block number has increased.
	chainConfig := sp.cp.ChainConfig()
	sp.signer = types.MakeSigner(chainConfig, sp.block.Number)

	// Setup the EVM for this block.
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

	// Create a new context to be used in the EVM environment. We also must reset the StateDB and
	// precompile manager, which resets the state and precompile plugins, and gas plugin for the
	// tx.
	txContext := NewEVMTxContext(msg)
	sp.evm.SetTxContext(txContext)
	sp.statedb.Reset(ctx)
	sp.pp.Reset(ctx)
	sp.gp.Reset(ctx)

	// Apply the state transition.
	result, err := ApplyMessage(sp.evm, sp.gp, msg, sp.commit)
	if err != nil {
		return nil, errors.Wrapf(err, "could not apply message %d [%v]", sp.block.TxIndex(), tx.Hash().Hex())
	}

	receipt := &types.Receipt{
		Type:             tx.Type(),
		PostState:        nil, // TODO: Should we do something with PostState?
		TxHash:           tx.Hash(),
		GasUsed:          result.UsedGas,
		BlockHash:        sp.block.Hash(),
		BlockNumber:      sp.block.Number,
		TransactionIndex: sp.block.TxIndex(),
		// Gas from this transaction was added to the gasPlugin in `ApplyMessageAndCommit`
		// And thus CumulativeGasUsed should include gas from all prior transactions in the
		// block, plus the gas consumed during this one.
		CumulativeGasUsed: sp.gp.CumulativeGasUsed(),
	}

	// Protect the chain from getting into an invalid state.
	if (receipt.CumulativeGasUsed > sp.block.GasLimit) && (sp.block.GasLimit != 0) {
		panic(
			fmt.Sprintf(
				"cumulative gas used %d is greater than block gas limit %d",
				receipt.CumulativeGasUsed, sp.block.GasLimit,
			),
		)
	}

	// Update the receipt based on the receipt of the transaction.
	if result.Failed() {
		receipt.Status = types.ReceiptStatusFailed
	} else {
		receipt.Status = types.ReceiptStatusSuccessful
	}

	// If the transaction created a contract, store the creation address in the receipt.
	if msg.To() == nil {
		receipt.ContractAddress = crypto.CreateAddress(txContext.Origin, tx.Nonce())
	}

	// Set the receipt logs and create the bloom filter.
	receipt.Logs = sp.statedb.BuildLogsAndClear(
		receipt.TxHash, receipt.BlockHash, sp.block.TxIndex(), sp.block.LogIndex(),
	)
	receipt.Bloom = types.BytesToBloom(types.LogsBloom(receipt.Logs))

	// Update the block information.
	sp.block.AppendTx(tx, receipt)
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
	var baseFee *big.Int
	// Copy the baseFee to avoid side effects.
	if sp.block.BaseFee != nil {
		baseFee = new(big.Int).Set(sp.block.BaseFee)
	}

	return vm.BlockContext{
		CanTransfer: state.CanTransfer,
		Transfer:    state.Transfer,
		GetHash:     sp.GetHashFn(),
		Coinbase:    sp.block.Coinbase,
		BlockNumber: new(big.Int).Set(sp.block.StargazerHeader.Number),
		Time:        sp.block.StargazerHeader.Header.Time,
		Difficulty:  new(big.Int), // not used by stargazer.
		BaseFee:     baseFee,
		GasLimit:    sp.block.GasLimit,
		Random:      &common.Hash{}, // TODO: find a source of randomness
	}
}

// `GetHashFn` returns a `GetHashFunc` which retrieves header hashes by number.
func (sp *StateProcessor) GetHashFn() vm.GetHashFunc {
	return GetHashFn(sp.block.Header, &chainContext{sp})
}
