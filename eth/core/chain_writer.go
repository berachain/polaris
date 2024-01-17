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

	"github.com/berachain/polaris/eth/core/state"
	"github.com/berachain/polaris/eth/core/types"

	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

// ChainWriter defines methods that are used to perform state and block transitions.
type ChainWriter interface {
	LoadLastState(uint64) error
	WriteGenesisBlock(block *ethtypes.Block) error
	InsertBlock(block *ethtypes.Block) error
	InsertBlockAndSetHead(block *ethtypes.Block) error
	SetFinalizedBlock() error
	WriteBlockAndSetHead(block *ethtypes.Block, receipts []*ethtypes.Receipt, logs []*ethtypes.Log,
		state state.StateDB, emitHeadEvent bool) (status core.WriteStatus, err error)
}

// InsertBlock inserts a block into the blockchain without setting it as the head.
func (bc *blockchain) InsertBlock(block *ethtypes.Block) error {
	// Get the state with the latest insert chain context.
	sp := bc.spf.NewPluginWithMode(state.Insert)
	state := state.NewStateDB(sp, bc.pp)

	// Call the private method to insert the block and setting it as the head.
	_, _, err := bc.insertBlock(block, state)
	// Return any error that might have occurred.
	return err
}

// insertBlock inserts a block into the blockchain by running the state processor and
// validating whether its okay.
func (bc *blockchain) insertBlock(
	block *ethtypes.Block, state state.StateDB,
) ([]*ethtypes.Receipt, []*ethtypes.Log, error) {
	// Validate that we are about to insert a valid block.
	// If the block number is greater than 1,
	// it means it's not the genesis block and needs to be validated. TODO kinda hood.
	if block.NumberU64() > 1 {
		if err := bc.validator.ValidateBody(block); err != nil {
			log.Error("invalid block body", "err", err)
			return nil, nil, err
		}
	}

	// Process the incoming EVM block.
	receipts, logs, usedGas, err := bc.processor.Process(block, state, *bc.vmConfig)
	if err != nil {
		log.Error("failed to process block", "num", block.NumberU64(), "err", err)
		return nil, nil, err
	}

	// ValidateState validates the statedb post block processing.
	if err = bc.validator.ValidateState(block, state, receipts, usedGas); err != nil {
		log.Error("invalid state after processing block", "num", block.NumberU64(), "err", err)
		return nil, nil, err
	}

	return receipts, logs, nil
}

// InsertBlockAndSetHead inserts a block into the blockchain and sets the head.
func (bc *blockchain) InsertBlockAndSetHead(block *ethtypes.Block) error {
	// Get the state with the latest finalize block context.
	sp := bc.spf.NewPluginWithMode(state.Finalize)
	state := state.NewStateDB(sp, bc.pp)

	receipts, logs, err := bc.insertBlock(block, state)
	if err != nil {
		return err
	}
	// We can just immediately finalize the block. It's okay in this context.
	if _, err = bc.WriteBlockAndSetHead(
		block, receipts, logs, state, true); err != nil {
		log.Error("failed to write block", "num", block.NumberU64(), "err", err)
		return err
	}
	return err
}

// WriteBlockAndSetHead sets the head of the blockchain to the given block and finalizes the block.
func (bc *blockchain) WriteBlockAndSetHead(
	block *ethtypes.Block, receipts []*ethtypes.Receipt, logs []*ethtypes.Log,
	state state.StateDB, emitHeadEvent bool,
) (core.WriteStatus, error) {
	// Write the block to the store.
	if err := bc.writeBlockWithState(block, receipts, state); err != nil {
		return core.NonStatTy, err
	}
	currentBlock := bc.currentBlock.Load()

	// We need to error if the parent is not the head block.
	if block.NumberU64() > 0 && block.ParentHash() != currentBlock.Hash() {
		log.Error("canonical chain broken",
			"block-number", block.NumberU64(), "block-hash", block.ParentHash().Hex())
		return core.NonStatTy, errors.New("canonical chain broken")
	}

	// Set the current block.
	bc.currentBlock.Store(block)

	// Store txLookup entries for all transactions in the block.
	blockNum := block.NumberU64()
	blockHash := block.Hash()
	bc.blockNumCache.Add(blockNum, block)
	bc.blockHashCache.Add(blockHash, block)
	for txIndex, tx := range block.Transactions() {
		bc.txLookupCache.Add(
			tx.Hash(),
			&types.TxLookupEntry{
				Tx:        tx,
				TxIndex:   uint64(txIndex),
				BlockNum:  blockNum,
				BlockHash: blockHash,
			},
		)
	}

	// Write the receipts cache.
	// TODO deprecate this cache?
	if receipts != nil {
		bc.receiptsCache.Add(block.Hash(), receipts)
	}

	// In theory, we should fire a ChainHeadEvent when we inject
	// a canonical block, but sometimes we can insert a batch of
	// canonical blocks. Avoid firing too many ChainHeadEvents,
	// we will fire an accumulated ChainHeadEvent and disable fire
	// event here.
	if emitHeadEvent {
		// Fire off the feeds.
		bc.chainFeed.Send(core.ChainEvent{Block: block, Hash: block.Hash(), Logs: logs})
		if len(logs) > 0 {
			bc.logsFeed.Send(logs)
		}
		bc.chainHeadFeed.Send(core.ChainHeadEvent{Block: block})
	}

	return core.CanonStatTy, nil
}

// writeBlockWithState writes the block along with its state (receipts and logs)
// into the blockchain.
func (bc *blockchain) writeBlockWithState(
	block *ethtypes.Block, receipts []*ethtypes.Receipt, state state.StateDB,
) error {
	// In Polaris since we are using single block finality.
	// Finalized == Current == Safe. All are the same.
	// Store the header as well as update all the finalized stuff.
	err := bc.bp.StoreHeader(block.Header())
	if err != nil {
		bc.logger.Error("failed to store block header", "err", err)
		return err
	}

	// Irrelevant of the canonical status, write the block itself to the database.
	// TODO THIS NEEDS TO WRITE TO EXTERNAL DB.
	if err = bc.writeHistoricalData(block, receipts); err != nil {
		return err
	}

	// Commit all cached state changes into underlying memory database.
	// In Polaris this is a no-op.
	_, err = state.Commit(block.NumberU64(), bc.config.IsEIP158(block.Number()))
	if err != nil {
		return err
	}

	bc.logger.Info(
		"finalizing evm block", "hash", block.Hash().Hex(), "num_txs", len(receipts))

	return nil
}

// InsertBlock inserts a block into the canonical chain and updates the state of the blockchain.
// TODO: WRITE TO EXTERNAL STORE
func (bc *blockchain) writeHistoricalData(
	block *ethtypes.Block,
	receipts ethtypes.Receipts,
) error {
	var err error
	blockHash, blockNum := block.Hash(), block.Number().Uint64()

	// store the block, receipts, and txs on the host chain if historical plugin is supported
	if bc.hp != nil {
		if err = bc.hp.StoreBlock(block); err != nil {
			bc.logger.Error("failed to store block", "err", err)
			return err
		}
		if err = bc.hp.StoreReceipts(blockHash, receipts); err != nil {
			bc.logger.Error("failed to store receipts", "err", err)
			return err
		}
		if err = bc.hp.StoreTransactions(blockNum, blockHash, block.Transactions()); err != nil {
			bc.logger.Error("failed to store transactions", "err", err)
			return err
		}
	}

	return nil
}

// For clarity reasons, the host chain makes a separate call to finalize the block. Only called
// once it is known the current block is the finalized block.
func (bc *blockchain) SetFinalizedBlock() error {
	if currBlock := bc.currentBlock.Load(); currBlock != nil {
		bc.finalizedBlock.Store(currBlock)
	}
	return nil
}
