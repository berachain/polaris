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

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/trie"

	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/common"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	errorslib "pkg.berachain.dev/polaris/lib/errors"
)

// TODO: WHO WROTE THIS CODE THE FIRST TIME BLS FIX IT IS HORRIBLE.

// StoreBlock implements `core.HistoricalPlugin`.
func (p *plugin) StoreBlock(block *coretypes.Block) error {
	blockNum := block.NumberU64()

	// store block hash to block number.
	numBz := sdk.Uint64ToBigEndian(blockNum)
	store := p.ctx.KVStore(p.offchainStoreKey)
	prefix.NewStore(store, []byte{types.BlockHashKeyToNumPrefix}).Set(block.Hash().Bytes(), numBz)

	// store the version offchain for consistency.
	if sdk.BigEndianToUint64(store.Get([]byte{types.VersionKey})) != blockNum-1 {
		panic("off-chain store's latest block number is not synced")
	}
	store.Set([]byte{types.VersionKey}, numBz)
	return nil
}

// StoreReceipts implements `core.HistoricalPlugin`.
func (p *plugin) StoreReceipts(blockHash common.Hash, receipts coretypes.Receipts) error {
	// store block hash to receipts.
	receiptsBz, err := coretypes.MarshalReceipts(receipts)
	if err != nil {
		p.ctx.Logger().Error(
			"UpdateOffChainStorage: failed to marshal receipts at block hash %s", blockHash.Hex(),
		)
		return err
	}
	prefix.NewStore(p.ctx.KVStore(p.offchainStoreKey),
		[]byte{types.BlockHashKeyToReceiptsPrefix}).Set(blockHash.Bytes(), receiptsBz)

	return nil
}

// StoreTransactions implements `core.HistoricalPlugin`.
func (p *plugin) StoreTransactions(
	blockNum int64, blockHash common.Hash, txs coretypes.Transactions,
) error {
	// store all txns in the block.
	txStore := prefix.NewStore(p.ctx.KVStore(p.offchainStoreKey), []byte{types.TxHashKeyToTxPrefix})
	for txIndex, tx := range txs {
		txLookupEntry := &coretypes.TxLookupEntry{
			Tx:        tx,
			TxIndex:   uint64(txIndex),
			BlockHash: blockHash,
			BlockNum:  uint64(blockNum),
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
func (p *plugin) GetBlockByNumber(number int64) (*coretypes.Block, error) {
	// get header from on chain.
	header, err := p.bp.GetHeaderByNumber(number)
	if err != nil {
		return nil, err
	}

	// get receipts from off chain.
	blockHash := header.Hash()
	receiptsBz := prefix.NewStore(p.ctx.KVStore(p.offchainStoreKey),
		[]byte{types.BlockHashKeyToReceiptsPrefix}).Get(blockHash.Bytes())
	if receiptsBz == nil {
		return nil, fmt.Errorf("failed to find receipts for block hash %s", blockHash.Hex())
	}
	receipts, err := coretypes.UnmarshalReceipts(receiptsBz)
	if err != nil {
		return nil, errorslib.Wrapf(err, "failed to unmarshal receipts for block hash %s", blockHash.Hex())
	}

	// get txns from off chain.
	txStore := prefix.NewStore(p.ctx.KVStore(p.offchainStoreKey), []byte{types.TxHashKeyToTxPrefix})
	txs := make(coretypes.Transactions, len(receipts))
	for _, receipt := range receipts {
		tleBz := txStore.Get(receipt.TxHash.Bytes())
		if tleBz == nil {
			return nil, fmt.Errorf("failed to find tx %s", receipt.TxHash.Hex())
		}
		tle := &coretypes.TxLookupEntry{}
		err = tle.UnmarshalBinary(tleBz)
		if err != nil {
			return nil, errorslib.Wrapf(err, "failed to unmarshal tx %s", receipt.TxHash.Hex())
		}
		txs = append(txs, tle.Tx)
	}

	// build the block.
	return coretypes.NewBlock(header, txs, nil, receipts, trie.NewStackTrie(nil)), nil
}

// GetBlockByHash returns the block at the given hash.
func (p *plugin) GetBlockByHash(blockHash common.Hash) (*coretypes.Block, error) {
	// get block number from off chain.
	numBz := prefix.NewStore(p.ctx.KVStore(p.offchainStoreKey),
		[]byte{types.BlockHashKeyToNumPrefix}).Get(blockHash.Bytes())
	if numBz == nil {
		return nil, fmt.Errorf("failed to find block number for block hash %s", blockHash.Hex())
	}
	number := int64(sdk.BigEndianToUint64(numBz))
	header, err := p.bp.GetHeaderByNumber(number)
	if err != nil {
		return nil, err
	}
	if int64(header.Number.Uint64()) != number || header.Hash() != blockHash {
		panic("header number or hash is not equal to the given number or hash")
	}

	// get receipts from off chain.
	receiptsBz := prefix.NewStore(p.ctx.KVStore(p.offchainStoreKey),
		[]byte{types.BlockHashKeyToReceiptsPrefix}).Get(blockHash.Bytes())
	if receiptsBz == nil {
		return nil, fmt.Errorf("failed to find receipts for block hash %s", blockHash.Hex())
	}
	receipts, err := coretypes.UnmarshalReceipts(receiptsBz)
	if err != nil {
		return nil, errorslib.Wrapf(err, "failed to unmarshal receipts for block hash %s", blockHash.Hex())
	}

	// get txns from off chain.
	txStore := prefix.NewStore(p.ctx.KVStore(p.offchainStoreKey), []byte{types.TxHashKeyToTxPrefix})
	txs := make(coretypes.Transactions, len(receipts))
	for _, receipt := range receipts {
		tleBz := txStore.Get(receipt.TxHash.Bytes())
		if tleBz == nil {
			return nil, fmt.Errorf("failed to find tx %s", receipt.TxHash.Hex())
		}
		tle := &coretypes.TxLookupEntry{}
		err = tle.UnmarshalBinary(tleBz)
		if err != nil {
			return nil, errorslib.Wrapf(err, "failed to unmarshal tx %s", receipt.TxHash.Hex())
		}
		txs = append(txs, tle.Tx)
	}

	// build the block.
	return coretypes.NewBlock(header, txs, nil, receipts, trie.NewStackTrie(nil)), nil
}

// GetTransactionByHash returns the transaction lookup entry with the given hash.
func (p *plugin) GetTransactionByHash(txHash common.Hash) (*coretypes.TxLookupEntry, error) {
	// get tx from off chain.
	tleBz := prefix.NewStore(p.ctx.KVStore(p.offchainStoreKey), []byte{types.TxHashKeyToTxPrefix}).Get(txHash.Bytes())
	if tleBz == nil {
		return nil, fmt.Errorf("failed to find tx %s", txHash.Hex())
	}
	tle := &coretypes.TxLookupEntry{}
	err := tle.UnmarshalBinary(tleBz)
	if err != nil {
		return nil, errorslib.Wrapf(err, "failed to unmarshal tx %s", txHash.Hex())
	}
	return tle, nil
}

// GetReceiptsByHash returns the receipts with the given block hash.
func (p *plugin) GetReceiptsByHash(blockHash common.Hash) (coretypes.Receipts, error) {
	// get receipts from off chain.
	receiptsBz := prefix.NewStore(p.ctx.KVStore(p.offchainStoreKey),
		[]byte{types.BlockHashKeyToReceiptsPrefix}).Get(blockHash.Bytes())
	if receiptsBz == nil {
		return nil, fmt.Errorf("failed to find receipts for block hash %s", blockHash.Hex())
	}
	receipts, err := coretypes.UnmarshalReceipts(receiptsBz)
	if err != nil {
		return nil, errorslib.Wrapf(err, "failed to unmarshal receipts for block hash %s", blockHash.Hex())
	}
	return receipts, nil
}
