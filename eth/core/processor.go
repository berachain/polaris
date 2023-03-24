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

	"github.com/ethereum/go-ethereum/trie"

	"pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/eth/crypto"
	"pkg.berachain.dev/polaris/lib/errors"
	"pkg.berachain.dev/polaris/lib/utils"
)

// initialTxsCapacity is the initial capacity of the transactions and receipts slice.
const initialTxsCapacity = 256

// StateProcessor is responsible for processing blocks, transactions, and updating the state.
type StateProcessor struct {
	// mtx is used to make sure we don't try to prepare a new block before finalizing the
	// current block.
	mtx sync.Mutex

	// cp provides configuration functions from the underlying chain the EVM is running on.
	cp ConfigurationPlugin
	// pp is responsible for keeping track of the stateful precompile containers that are
	// available to the EVM and executing them.
	pp PrecompilePlugin

	// signer is the signer used to verify transaction signatures. We need this in order to to
	// extract the underlying message from a transaction object in `ProcessTransaction`.
	signer types.Signer

	// evm is the EVM that is used to process transactions. We re-use a single EVM for processing
	// the entire block. This is done in order to reduce memory allocs.
	evm *vm.GethEVM
	// statedb is the state database that is used to mange state during transactions.
	statedb vm.PolarisStateDB
	// vmConfig is the configuration for the EVM.
	vmConfig *vm.Config

	// We store information about the current block being processed so that we can access it
	// during the processing of transactions. This allows us to utilize this information to
	// build the `block` and return the canonical receipts in `Finalize`.
	header   *types.Header
	txs      types.Transactions
	receipts types.Receipts
}

// NewStateProcessor creates a new state processor with the given host, statedb, vmConfig, and
// commit flag.
func NewStateProcessor(
	cp ConfigurationPlugin,
	pp PrecompilePlugin,
	statedb vm.PolarisStateDB,
	vmConfig *vm.Config,
) *StateProcessor {
	sp := &StateProcessor{
		mtx:      sync.Mutex{},
		cp:       cp,
		pp:       pp,
		vmConfig: vmConfig,
		statedb:  statedb,
	}

	if sp.pp == nil {
		sp.pp = precompile.NewDefaultPlugin()
	} else {
		sp.BuildAndRegisterPrecompiles(sp.pp.GetPrecompiles(nil))
	}

	return sp
}

// ==============================================================================
// Block, Tx Lifecycle
// ==============================================================================

// Prepare prepares the state processor for processing a block.
func (sp *StateProcessor) Prepare(
	ctx context.Context, evm *vm.GethEVM, header *types.Header, gp GasPlugin,
) {
	// We lock the state processor as a safety measure to ensure that Prepare is not called again
	// before finalize.
	sp.mtx.Lock()

	// Build a header object so we can track that status of the block as we process it.
	sp.header = header
	sp.txs = make(types.Transactions, 0, initialTxsCapacity)
	sp.receipts = make(types.Receipts, 0, initialTxsCapacity)

	// Ensure that the gas plugin and header are in sync.
	if sp.header.GasLimit != gp.BlockGasLimit() {
		panic(fmt.Sprintf("gas limit mismatch: have %d, want %d", sp.header.GasLimit, gp.BlockGasLimit()))
	}

	// We must re-create the signer since we are processing a new block and the block number has
	// increased.
	chainConfig := sp.cp.ChainConfig(ctx)
	sp.signer = types.MakeSigner(chainConfig, sp.header.Number)

	// Setup the EVM for this block.
	rules := chainConfig.Rules(sp.header.Number, true, sp.header.Time)
	// We re-register the default geth precompiles every block, this isn't optimal, but since
	// *technically* the precompiles change based on the chain config rules, to be fully correct,
	// we should check every block.
	sp.BuildAndRegisterPrecompiles(precompile.GetDefaultPrecompiles(&rules))
	sp.vmConfig.ExtraEips = sp.cp.ExtraEips(ctx)
	sp.evm = evm
}

// ProcessTransaction applies a transaction to the current state of the blockchain.
func (sp *StateProcessor) ProcessTransaction(
	ctx context.Context, tx *types.Transaction, gp GasPlugin,
) (*ExecutionResult, error) {
	txHash := tx.Hash()
	msg, err := TransactionToMessage(tx, sp.signer, sp.header.BaseFee)
	if err != nil {
		return nil, errors.Wrapf(err, "could not apply tx %d [%s]", len(sp.txs), txHash.Hex())
	}

	// Create a new context to be used in the EVM environment and tx context for the StateDB.
	txContext := NewEVMTxContext(msg)
	sp.evm.Reset(txContext, sp.statedb)
	sp.statedb.Reset(ctx)
	sp.statedb.SetTxContext(txHash, len(sp.txs))

	// Set the gasPool to have the remaining gas in the block.
	// By setting the gas pool to the delta between the block gas limit and the cumulative gas
	// used, we intrinsic handle the case where the transaction on our host chain might have
	// fully reverted, when it fact it should've been a vm error saying out of gas.
	gasPool := GasPool(gp.BlockGasLimit() - gp.BlockGasConsumed())

	// Apply the state transition.
	result, err := ApplyMessage(sp.evm, msg, &gasPool)
	if err != nil {
		return nil, errors.Wrapf(err, "could not apply message %d [%s]", len(sp.txs), txHash.Hex())
	}

	// If we used more gas than we had remaining on the gas plugin, we treat it as an out of gas error,
	// while still ensuring that we consume all the gas.
	if result.UsedGas > gp.GasRemaining() {
		result.UsedGas = gp.GasRemaining()
		result.Err = vm.ErrOutOfGas
	}

	// Consume the gas used by the state transition. In both the out of block gas as well as out of gas on
	// the plugin cases, the line below will consume the remaining gas for the block and transaction respectively.
	if err = gp.ConsumeGas(result.UsedGas); err != nil {
		return nil, errors.Wrapf(err, "could not consume gas used %d [%s]", len(sp.txs), txHash.Hex())
	}

	// Create a new receipt for the transaction.
	receipt := &types.Receipt{
		Type:              tx.Type(),
		CumulativeGasUsed: gp.BlockGasConsumed() + gp.GasConsumed(),
		TxHash:            txHash,
		GasUsed:           result.UsedGas,
		Logs:              sp.statedb.Logs(),
	}

	// If the transaction created a contract, store the creation address in the receipt.
	if msg.To == nil {
		receipt.ContractAddress = crypto.CreateAddress(txContext.Origin, tx.Nonce())
	}

	// Set the receipt status based on the execution result status.
	if result.Failed() {
		receipt.Status = types.ReceiptStatusFailed
	} else {
		receipt.Status = types.ReceiptStatusSuccessful
	}

	// Finalize the statedb to ensure that any state changes that are required are propogated.
	// We have to do this irrespective of whether the transaction failed or not, in order to
	// ensure that the sender's nonce increases as well as the transaction fees are paid.
	// The snapshotting within the EVM ensures that any reverted state changes are not reflected
	// in the finalized state.
	sp.statedb.Finalize()

	// Update the block information.
	sp.txs = append(sp.txs, tx)
	sp.receipts = append(sp.receipts, receipt)

	// Return the execution result to the caller.
	return result, nil
}

// Finalize finalizes the block in the state processor and returns the receipts and bloom filter.
func (sp *StateProcessor) Finalize(
	_ context.Context, gp GasPlugin,
) (*types.Block, types.Receipts, []*types.Log, error) {
	// We unlock the state processor to ensure that the state is consistent.
	defer sp.mtx.Unlock()

	// Now that we are done processing the block, we update the header with the consumed gas.
	sp.header.GasUsed = gp.BlockGasConsumed()

	// We iterate over all of the receipts/transactions in the block and update the receipt to
	// have the correct values. We must do this AFTER all the transactions have been processed
	// to ensure that the block hash, logs and bloom filter have the correct information.
	blockHash, blockNumber := sp.header.Hash(), sp.header.Number.Uint64()
	fmt.Println("PROCESSOR FINALIZE", "blockHash", blockHash.Hex())
	var logIndex uint
	var logs []*types.Log
	for txIndex, receipt := range sp.receipts {
		// Edit the receipts to include the block hash and bloom filter.
		for _, log := range receipt.Logs {
			log.BlockNumber = blockNumber
			log.BlockHash = blockHash
			log.Index = logIndex
			logIndex++
			logs = append(logs, log)
		}
		receipt.Bloom = types.CreateBloom(types.Receipts{receipt})
		receipt.BlockHash = blockHash
		receipt.BlockNumber = sp.header.Number
		receipt.TransactionIndex = uint(txIndex)
	}

	// We return a new block with the updated header and the receipts to the `blockchain`.
	return types.NewBlock(sp.header, sp.txs, nil, sp.receipts, trie.NewStackTrie(nil)), sp.receipts, logs, nil
}

// ===========================================================================
// Utilities
// ===========================================================================

// BuildPrecompiles builds the given precompiles and registers them with the precompile plugins.
func (sp *StateProcessor) BuildAndRegisterPrecompiles(precompiles []vm.RegistrablePrecompile) {
	for _, pc := range precompiles {
		// skip registering precompiles that are already registered.
		if sp.pp.Has(pc.RegistryKey()) {
			continue
		}

		// choose the appropriate precompile factory
		var af precompile.AbstractFactory
		switch {
		case utils.Implements[precompile.DynamicImpl](pc):
			af = precompile.NewDynamicFactory()
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
		container, err := af.Build(pc)
		if err != nil {
			panic(err)
		}
		err = sp.pp.Register(container)
		if err != nil {
			panic(err)
		}
	}
}
