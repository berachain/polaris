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
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/log"

	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/state"
	"pkg.berachain.dev/polaris/eth/core/types"
)

// ChainWriter defines methods that are used to perform state and block transitions.
type ChainWriter interface {
	WriteGenesis(*types.Header)
	LoadLastState(context.Context, uint64) error
	InsertBlockWithoutSetHead(block *types.Block) error
	InsertBlockWithSetHead(block *types.Block) error
	WriteBlockAndSetHead(block *types.Block, receipts []*types.Receipt, logs []*types.Log,
		state state.StateDB, emitHeadEvent bool) (status core.WriteStatus, err error)
}

// InsertChain attempts to insert the given batch of blocks in to the canonical
// chain or, otherwise, create a fork. If an error is returned it will return
// the index number of the failing block as well an error describing what went
// wrong. After insertion is done, all accumulated events will be fired.
func (bc *blockchain) InsertChain(chain types.Blocks) (int, error) {
	// Sanity check that we have something meaningful to import
	if len(chain) == 0 {
		return 0, nil
	}
	// bc.blockProcFeed.Send(true)
	// defer bc.blockProcFeed.Send(false)

	// Do a sanity check that the provided chain is actually ordered and linked.
	for i := 1; i < len(chain); i++ {
		block, prev := chain[i], chain[i-1]
		if block.NumberU64() != prev.NumberU64()+1 || block.ParentHash() != prev.Hash() {
			log.Error("Non contiguous block insert",
				"number", block.Number(),
				"hash", block.Hash(),
				"parent", block.ParentHash(),
				"prevnumber", prev.Number(),
				"prevhash", prev.Hash(),
			)
			return 0, fmt.Errorf(
				"non contiguous insert: item %d is #%d [%x..], "+
					"item %d is #%d [%x..] (parent [%x..])",
				i-1, prev.NumberU64(),
				prev.Hash().Bytes()[:4],
				i, block.NumberU64(),
				block.Hash().Bytes()[:4],
				block.ParentHash().Bytes()[:4],
			)
		}
	}
	return bc.insertChain(chain, true)
}

// InsertBlockWithoutSetHead executes the block, runs the necessary verification
// upon it and then persist the block and the associate state into the database.
// The key difference between the InsertChain is it won't do the canonical chain
// updating. It relies on the additional SetCanonical call to finalize the entire
// procedure.
func (bc *blockchain) InsertBlockWithoutSetHead(block *types.Block) error {
	_, err := bc.insertChain(types.Blocks{block}, false)
	return err
}

func (bc *blockchain) InsertBlockWithSetHead(block *types.Block) error {
	val, err := bc.insertChain(types.Blocks{block}, true)
	bc.logger.Debug("InsertBlockWithSetHead", "block", block.Hash(), "val", val)
	return err
}

// insertChain attempts to insert the given batch of blocks in to the canonical
// chain.
func (bc *blockchain) insertChain(blocks types.Blocks, setHead bool) (int, error) {
	var lastCanon *types.Block

	if len(blocks) != 1 {
		return 0, errors.New("polaris only supports inserting chains of length 1")
	}

	block := blocks[0]
	start := time.Now()

	// Verify that the parent block exists.
	parent := bc.GetBlock(block.ParentHash(), block.NumberU64()-1)
	if block.Number().Cmp(common.Big0) != 0 && parent == nil {
		return 0, fmt.Errorf("parent block not found")
	}

	// Fire a single chain head event if we've progressed the chain
	defer func() {
		if lastCanon != nil && bc.CurrentBlock().Hash() == lastCanon.Hash() {
			bc.chainHeadFeed.Send(ChainHeadEvent{Block: lastCanon})
		}
	}()

	// Verify the incoming block header is valid.
	if block.Number().Cmp(common.Big0) != 0 {
		// TODO: do not skip for genesis
		if err := bc.engine.VerifyHeader(bc, block.Header()); err != nil {
			return 0, err
		}
	}

	var err error
	// Execute the state transition
	receipts, logs, usedGas, err := bc.processor.Process(block, bc.statedb, *bc.vmConfig)
	if err != nil {
		return 0, err
	}

	// Validate the state of the chain after running the state tranisitons.
	if err = bc.validator.ValidateState(block, bc.statedb, receipts, usedGas); err != nil {
		return 0, err
	}

	// Write the block to the chain and get the status.
	var (
		status core.WriteStatus
	)

	if !setHead {
		// Don't set the head, only insert the block
		err = bc.writeBlockWithState(block, receipts, bc.statedb)
	} else {
		status, err = bc.writeBlockAndSetHead(block, receipts, logs, bc.statedb, false)
	}
	if err != nil {
		return 0, err
	}

	switch status {
	case core.CanonStatTy:
		log.Debug("Inserted new block", "number", block.Number(), "hash", block.Hash(),
			"uncles", len(block.Uncles()), "txs", len(block.Transactions()), "gas", block.GasUsed(),
			"elapsed", common.PrettyDuration(time.Since(start)),
			"root", block.Root())

		lastCanon = block
	case core.SideStatTy:
		log.Debug("Inserted new side block", "number", block.Number(), "hash", block.Hash(),
			"uncles", len(block.Uncles()), "txs", len(block.Transactions()), "gas", block.GasUsed(),
			"elapsed", common.PrettyDuration(time.Since(start)),
			"root", block.Root())
	case core.NonStatTy:
	default:
		// This in theory is impossible, but lets be nice to our future selves and leave
		// a log, instead of trying to track down blocks imports that don't emit logs.
		log.Warn("Inserted block with unknown status", "number", block.Number(), "hash", block.Hash(),
			"diff", block.Difficulty(), "elapsed", common.PrettyDuration(time.Since(start)),
			"txs", len(block.Transactions()), "gas", block.GasUsed(), "uncles", len(block.Uncles()),
			"root", block.Root())
	}

	return -1, err
}

// writeBlockWithState writes block, metadata and corresponding state data to the
// database.
// func (bc *blockchain) writeBlockWithState(
// block *types.Block, receipts []*types.Receipt, state state.StateDB) error {
// 	// Calculate the total difficulty of the block
// 	ptd := bc.GetTd(block.ParentHash(), block.NumberU64()-1)
// 	if ptd == nil {
// 		return consensus.ErrUnknownAncestor
// 	}
// 	// // Make sure no inconsistent state is leaked during insertion
// 	// externTd := new(big.Int).Add(block.Difficulty(), ptd)

// 	// Irrelevant of the canonical status, write the block itself to the database.
// 	//
// 	// Note all the components of block(td, hash->number map, header, body, receipts)
// 	// // should be written atomically. BlockBatch is used for containing all components.
// 	// blockBatch := bc.db.NewBatch()
// 	// rawdb.WriteTd(blockBatch, block.Hash(), block.NumberU64(), externTd)
// 	// rawdb.WriteBlock(blockBatch, block)
// 	// rawdb.WriteReceipts(blockBatch, block.Hash(), block.NumberU64(), receipts)
// 	// rawdb.WritePreimages(blockBatch, state.Preimages())
// 	// if err := blockBatch.Write(); err != nil {
// 	// 	log.Crit("Failed to write block into disk", "err", err)
// 	// }

// 	// Commit all cached state changes into underlying memory database.
// 	_, err := state.Commit(block.NumberU64(), bc.cp.ChainConfig().IsEIP158(block.Number()))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// WriteBlockAndSetHead is a no-op in the current implementation. Potentially usable later.
func (bc *blockchain) WriteBlockAndSetHead(
	block *types.Block, receipts []*types.Receipt,
	logs []*types.Log, state state.StateDB,
	emitHeadEvent bool) (core.WriteStatus, error) {
	return bc.writeBlockAndSetHead(block, receipts, logs, state, emitHeadEvent)
}

// func (bc *blockchain) writeHeadBlock(block *types.Block) {}

// InsertBlock inserts a block into the canonical chain and updates the state of the blockchain.
func (bc *blockchain) writeBlockAndSetHead(
	block *types.Block,
	receipts types.Receipts,
	logs []*types.Log,
	_ state.StateDB,
	emitHeadEvent bool,
) (core.WriteStatus, error) {
	var status = core.CanonStatTy
	if err := bc.writeBlockWithState(block, receipts, bc.statedb); err != nil {
		return core.NonStatTy, err
	}

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

	return status, nil
}

func (bc *blockchain) writeBlockWithState(
	block *types.Block, receipts []*types.Receipt, _ state.StateDB,
) error {
	var err error
	if _, err = bc.statedb.Commit(
		block.NumberU64(),
		bc.cp.ChainConfig().IsEIP158(block.Header().Number),
	); err != nil {
		return err
	}

	// TODO: prepare historical plugin here?
	// TBH still think we should deprecate it and run in another routine as indexer.

	// ***************************************** //
	// TODO: add safety check for canonicallness //
	// ***************************************** //

	// *********************************************** //
	// TODO: restructure this function / flow it sucks //
	// *********************************************** //
	blockHash, blockNum := block.Hash(), block.Number().Uint64()
	bc.logger.Info(
		"finalizing evm block", "block_hash", blockHash.Hex(), "num_txs", len(receipts))

	// store the block header on the host chain
	err = bc.bp.StoreHeader(block.Header())
	if err != nil {
		bc.logger.Error("failed to store block header", "err", err)
		return err
	}

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

	// mark the current block, receipts, and logs
	if block != nil {
		bc.currentBlock.Store(block)
		bc.finalizedBlock.Store(block)

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
	}
	if receipts != nil {
		bc.currentReceipts.Store(receipts)
		bc.receiptsCache.Add(blockHash, receipts)
	}

	return nil
}
