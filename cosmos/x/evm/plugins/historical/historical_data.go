// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
