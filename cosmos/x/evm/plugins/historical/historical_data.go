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

package historical

import (
	"fmt"

	"cosmossdk.io/store/prefix"

	"github.com/berachain/polaris/cosmos/x/evm/types"
	"github.com/berachain/polaris/eth/core"
	coretypes "github.com/berachain/polaris/eth/core/types"
	errorslib "github.com/berachain/polaris/lib/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

// TODO: WHO WROTE THIS CODE THE FIRST TIME BLS FIX IT IS HORRIBLE.

// StoreBlock implements `core.HistoricalPlugin`.
func (p *plugin) StoreBlock(block *ethtypes.Block) error {
	blockNum := block.NumberU64()

	// store block hash to block number.
	numBz := sdk.Uint64ToBigEndian(blockNum)
	store := p.ctx.MultiStore().GetKVStore(p.storeKey)

	// store block num to block
	blockBz, err := rlp.EncodeToBytes(block)
	if err != nil {
		return err
	}
	prefix.NewStore(store, []byte{types.BlockNumKeyToBlockPrefix}).
		Set(numBz, blockBz)

	// store block hash to block number.
	prefix.NewStore(store, []byte{types.BlockHashKeyToNumPrefix}).
		Set(block.Hash().Bytes(), numBz)

	// store the version offchain for consistency.
	offChainNum := sdk.BigEndianToUint64(store.Get([]byte{types.VersionKey}))
	if blockNum > 0 && offChainNum != blockNum-1 {
		panic(
			fmt.Errorf(
				"off-chain store's latest block number %d not synced with prev block number %d",
				offChainNum,
				blockNum-1,
			),
		)
	}
	store.Set([]byte{types.VersionKey}, numBz)
	return nil
}

// StoreReceipts implements `core.HistoricalPlugin`.
func (p *plugin) StoreReceipts(blockHash common.Hash, receipts ethtypes.Receipts) error {
	// store block hash to receipts.
	receiptsBz, err := coretypes.MarshalReceipts(receipts)
	if err != nil {
		p.ctx.Logger().Error(
			"UpdateOffChainStorage: failed to marshal receipts at block hash %s", blockHash.Hex(),
		)
		return err
	}
	prefix.NewStore(p.ctx.MultiStore().GetKVStore(p.storeKey),
		[]byte{types.BlockHashKeyToReceiptsPrefix}).Set(blockHash.Bytes(), receiptsBz)

	return nil
}

// StoreTransactions implements `core.HistoricalPlugin`.
func (p *plugin) StoreTransactions(
	blockNum uint64, blockHash common.Hash, txs ethtypes.Transactions,
) error {
	// store all txns in the block.
	txStore := prefix.NewStore(
		p.ctx.MultiStore().GetKVStore(p.storeKey), []byte{types.TxHashKeyToTxPrefix},
	)
	for txIndex, tx := range txs {
		txLookupEntry := &coretypes.TxLookupEntry{
			Tx:        tx,
			TxIndex:   uint64(txIndex),
			BlockHash: blockHash,
			BlockNum:  blockNum,
		}
		var tleBz []byte
		tleBz, err := txLookupEntry.MarshalBinary()
		if err != nil {
			p.ctx.Logger().Error(
				"UpdateOffChainStorage: failed to marshal tx %s at block number %d",
				tx.Hash().Hex(), blockNum,
			)
			return err
		}
		txStore.Set(tx.Hash().Bytes(), tleBz)
	}

	return nil
}

// GetBlockByNumber returns the block at the given height.
func (p *plugin) GetBlockByNumber(number uint64) (*ethtypes.Block, error) {
	store := p.ctx.MultiStore().GetKVStore(p.storeKey)
	numBz := sdk.Uint64ToBigEndian(number)
	blockBz := prefix.NewStore(store, []byte{types.BlockNumKeyToBlockPrefix}).Get(numBz)
	block := &ethtypes.Block{}
	err := rlp.DecodeBytes(blockBz, block)
	if err != nil {
		return nil, err
	}
	return block, nil
}

// GetBlockByHash returns the block at the given hash.
func (p *plugin) GetBlockByHash(blockHash common.Hash) (*ethtypes.Block, error) {
	store := p.ctx.MultiStore().GetKVStore(p.storeKey)
	numBz := prefix.NewStore(
		store, []byte{types.BlockHashKeyToNumPrefix}).Get(blockHash.Bytes())
	if numBz == nil {
		return nil, core.ErrBlockNotFound
	}

	blockBz := prefix.NewStore(store, []byte{types.BlockNumKeyToBlockPrefix}).Get(numBz)
	block := &ethtypes.Block{}

	err := rlp.DecodeBytes(blockBz, block)
	if err != nil {
		return nil, err
	}
	return block, nil
}

// GetTransactionByHash returns the transaction lookup entry with the given hash.
func (p *plugin) GetTransactionByHash(txHash common.Hash) (*coretypes.TxLookupEntry, error) {
	// get tx from off chain.
	tleBz := prefix.NewStore(
		p.ctx.MultiStore().GetKVStore(p.storeKey), []byte{types.TxHashKeyToTxPrefix}).Get(txHash.Bytes())
	if tleBz == nil {
		return nil, core.ErrTxNotFound
	}
	tle := &coretypes.TxLookupEntry{}
	err := tle.UnmarshalBinary(tleBz)
	if err != nil {
		return nil, errorslib.Wrapf(err, "failed to unmarshal tx %s", txHash.Hex())
	}
	return tle, nil
}

// GetReceiptsByHash returns the receipts with the given block hash.
func (p *plugin) GetReceiptsByHash(blockHash common.Hash) (ethtypes.Receipts, error) {
	// get receipts from off chain.
	receiptsBz := prefix.NewStore(p.ctx.MultiStore().GetKVStore(p.storeKey),
		[]byte{types.BlockHashKeyToReceiptsPrefix}).Get(blockHash.Bytes())
	if receiptsBz == nil {
		return nil, fmt.Errorf("failed to find receipts for block hash %s", blockHash.Hex())
	}
	receipts, err := coretypes.UnmarshalReceipts(receiptsBz)
	if err != nil {
		return nil, errorslib.Wrapf(
			err, "failed to unmarshal receipts for block hash %s", blockHash.Hex())
	}

	// get block to derive fields on receipts
	block, err := p.GetBlockByHash(blockHash)
	if err != nil {
		return nil, err
	}

	return coretypes.DeriveReceiptsFromBlock(p.chainConfig, receipts, block)
}
