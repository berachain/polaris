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
	"errors"

	"github.com/ethereum/go-ethereum/core"

	"pkg.berachain.dev/polaris/eth/core/state"
	"pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/log"
)

// ChainWriter defines methods that are used to perform state and block transitions.
type ChainWriter interface {
	LoadLastState(context.Context, uint64) error
	WriteGenesisBlock(block *types.Block) error
	InsertBlockAndSetHead(block *types.Block) error
	WriteBlockAndSetHead(block *types.Block, receipts []*types.Receipt, logs []*types.Log,
		state state.StateDB, emitHeadEvent bool) (status core.WriteStatus, err error)
}

// WriteGenesisBlock inserts the genesis block into the blockchain.
func (bc *blockchain) WriteGenesisBlock(block *types.Block) error {
	// TODO: add more validation here.
	if block.NumberU64() != 0 {
		return errors.New("not the genesis block")
	}
	_, err := bc.WriteBlockAndSetHead(block, nil, nil, nil, true)
	return err
}

// InsertBlockAndSetHead inserts a block into the blockchain without setting the head.
// For now, it is a huge lie. It does infact set the head.
func (bc *blockchain) InsertBlockAndSetHead(block *types.Block) error {
	// Validate that we are about to insert a valid block.
	if block.NumberU64() > 1 { // TODO DIAGNOSE
		if err := bc.validator.ValidateBody(block); err != nil {
			log.Error("invalid block body", "err", err)
			return err
		}
	}

	// Process the incoming EVM block.
	receipts, logs, usedGas, err := bc.processor.Process(block, bc.statedb, *bc.vmConfig)
	if err != nil {
		log.Error("failed to process block", "num", block.NumberU64(), "err", err)
		return err
	}

	// ValidateState validates the statedb post block processing.
	if err = bc.validator.ValidateState(block, bc.statedb, receipts, usedGas); err != nil {
		log.Error("invalid state after processing block", "num", block.NumberU64(), "err", err)
		return err
	}

	// We can just immediately finalize the block. It's okay in this context.
	if _, err = bc.WriteBlockAndSetHead(
		block, receipts, logs, nil, true); err != nil {
		log.Error("failed to write block", "num", block.NumberU64(), "err", err)
		return err
	}

	return err
}

// WriteBlockAndSetHead sets the head of the blockchain to the given block and finalizes the block.
func (bc *blockchain) WriteBlockAndSetHead(
	block *types.Block, receipts []*types.Receipt, logs []*types.Log,
	_ state.StateDB, emitHeadEvent bool,
) (core.WriteStatus, error) {
	// Write the block to the store.
	if err := bc.writeBlockWithState(block, receipts); err != nil {
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

	// TODO: this is fine to do here but not really semantically correct
	// and is very confusing.
	// For clarity reasons, we should make the cosmos chain make a separate call
	// to finalize the block.
	bc.finalizedBlock.Store(block)

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

	// Fire off the feeds.
	bc.chainFeed.Send(ChainEvent{Block: block, Hash: block.Hash(), Logs: logs})
	if len(logs) > 0 {
		bc.logsFeed.Send(logs)
	}

	// In theory, we should fire a ChainHeadEvent when we inject
	// a canonical block, but sometimes we can insert a batch of
	// canonical blocks. Avoid firing too many ChainHeadEvents,
	// we will fire an accumulated ChainHeadEvent and disable fire
	// event here.
	if emitHeadEvent {
		bc.chainHeadFeed.Send(ChainHeadEvent{Block: block})
	}

	return core.CanonStatTy, nil
}

// writeBlockWithState writes the block along with its state (receipts and logs)
// into the blockchain.
func (bc *blockchain) writeBlockWithState(
	block *types.Block, receipts []*types.Receipt,
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
	_, err = bc.statedb.Commit(block.NumberU64(), bc.config.IsEIP158(block.Number()))
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
	block *types.Block,
	receipts types.Receipts,
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
