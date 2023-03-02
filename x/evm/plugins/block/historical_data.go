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

package block

import (
	"fmt"
	"unsafe"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"

	"pkg.berachain.dev/stargazer/eth/common"
	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
	errorslib "pkg.berachain.dev/stargazer/lib/errors"
)

var (
	blockHashToNumPrefix      = []byte{0xb}
	blockHashToReceiptsPrefix = []byte{0xbb}
	txHashToTxPrefix          = []byte{0x10}
	versionKey                = []byte{0x11}
)

// `UpdateOffChainStorage` is called by the `EndBlocker` to update the off-chain storage.
func (p *plugin) UpdateOffChainStorage(block *coretypes.Block, receipts coretypes.Receipts) {
	blockHash, blockNum := block.Hash(), block.NumberU64()

	// store block hash to block number.
	numBz := sdk.Uint64ToBigEndian(blockNum)
	prefix.NewStore(p.offchainStore, blockHashToNumPrefix).Set(blockHash.Bytes(), numBz)

	// stpre block hash to receipts.
	receiptsBz, err := marshalReceipts(receipts)
	if err != nil {
		p.ctx.Logger().Error(
			"UpdateOffChainStorage: failed to marshal receipts at block number %d", blockNum,
		)
		panic(err)
	}
	prefix.NewStore(p.offchainStore, blockHashToReceiptsPrefix).Set(blockHash.Bytes(), receiptsBz)

	// store all txns in the block.
	txStore := prefix.NewStore(p.offchainStore, txHashToTxPrefix)
	for txIndex, tx := range block.Transactions() {
		txLookupEntry := &coretypes.TxLookupEntry{
			Tx:        tx,
			TxIndex:   uint64(txIndex),
			BlockHash: blockHash,
			BlockNum:  blockNum,
		}
		var tleBz []byte
		tleBz, err = txLookupEntry.MarshalBinary()
		if err != nil {
			p.ctx.Logger().Error(
				"UpdateOffChainStorage: failed to marshal tx %s at block number %d",
				tx.Hash().Hex(), blockNum,
			)
			panic(err)
		}
		txStore.Set(tx.Hash().Bytes(), tleBz)
	}

	// store the version offchain for consistency.
	if sdk.BigEndianToUint64(p.offchainStore.Get(versionKey)) != blockNum-1 {
		// TODO: resync the off-chain storage.
		panic("off-chain store's latest block number is not synced")
	}
	p.offchainStore.Set(versionKey, numBz)
	// flush the underlying buffer to disk.
	p.offchainStore.Write()
}

// `GetBlockByNumber` returns the block at the given height.
func (p *plugin) GetBlockByNumber(number int64) (*coretypes.Block, error) {
	// get header from on chain.
	header, err := p.GetHeaderByNumber(number)
	if err != nil {
		return nil, err
	}
	if int64(header.Number.Uint64()) != number {
		panic("header number is not equal to the given number")
	}

	// get receipts from off chain.
	blockHash := header.Hash()
	receiptsBz := prefix.NewStore(p.offchainStore, blockHashToReceiptsPrefix).Get(blockHash.Bytes())
	if receiptsBz == nil {
		return nil, fmt.Errorf("failed to find receipts for block hash %s", blockHash.Hex())
	}
	receipts, err := unmarshalReceipts(receiptsBz)
	if err != nil {
		return nil, errorslib.Wrapf(err, "failed to unmarshal receipts for block hash %s", blockHash.Hex())
	}

	// get txns from off chain.
	txStore := prefix.NewStore(p.offchainStore, txHashToTxPrefix)
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

// `GetBlockByHash` returns the block at the given hash.
func (p *plugin) GetBlockByHash(blockHash common.Hash) (*coretypes.Block, error) {
	// get block number from off chain.
	numBz := prefix.NewStore(p.offchainStore, blockHashToNumPrefix).Get(blockHash.Bytes())
	if numBz == nil {
		return nil, fmt.Errorf("failed to find block number for block hash %s", blockHash.Hex())
	}
	number := int64(sdk.BigEndianToUint64(numBz))
	header, err := p.GetHeaderByNumber(number)
	if err != nil {
		return nil, err
	}
	if int64(header.Number.Uint64()) != number || header.Hash() != blockHash {
		panic("header number or hash is not equal to the given number or hash")
	}

	// get receipts from off chain.
	receiptsBz := prefix.NewStore(p.offchainStore, blockHashToReceiptsPrefix).Get(blockHash.Bytes())
	if receiptsBz == nil {
		return nil, fmt.Errorf("failed to find receipts for block hash %s", blockHash.Hex())
	}
	receipts, err := unmarshalReceipts(receiptsBz)
	if err != nil {
		return nil, errorslib.Wrapf(err, "failed to unmarshal receipts for block hash %s", blockHash.Hex())
	}

	// get txns from off chain.
	txStore := prefix.NewStore(p.offchainStore, txHashToTxPrefix)
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

// `GetTransactionByHash` returns the transaction lookup entry with the given hash.
func (p *plugin) GetTransactionByHash(txHash common.Hash) (*coretypes.TxLookupEntry, error) {
	// get tx from off chain.
	tleBz := prefix.NewStore(p.offchainStore, txHashToTxPrefix).Get(txHash.Bytes())
	if tleBz == nil {
		return nil, fmt.Errorf("failed to find tx %s", txHash.Hex())
	}
	var tle *coretypes.TxLookupEntry
	err := tle.UnmarshalBinary(tleBz)
	if err != nil {
		return nil, errorslib.Wrapf(err, "failed to unmarshal tx %s", txHash.Hex())
	}
	return tle, nil
}

// `GetReceiptsByHash` returns the receipts with the given block hash.
func (p *plugin) GetReceiptsByHash(blockHash common.Hash) (coretypes.Receipts, error) {
	// get receipts from off chain.
	receiptsBz := prefix.NewStore(p.offchainStore, blockHashToReceiptsPrefix).Get(blockHash.Bytes())
	if receiptsBz == nil {
		return nil, fmt.Errorf("failed to find receipts for block hash %s", blockHash.Hex())
	}
	receipts, err := unmarshalReceipts(receiptsBz)
	if err != nil {
		return nil, errorslib.Wrapf(err, "failed to unmarshal receipts for block hash %s", blockHash.Hex())
	}
	return receipts, nil
}

func marshalReceipts(receipts coretypes.Receipts) ([]byte, error) {
	//#nosec:G103 unsafe pointer is safe here since `ReceiptForStorage` is an alias of `Receipt`.
	receiptsForStorage := *(*[]*coretypes.ReceiptForStorage)(unsafe.Pointer(&receipts))

	bz, err := rlp.EncodeToBytes(receiptsForStorage)
	if err != nil {
		return nil, err
	}
	return bz, nil
}

func unmarshalReceipts(bz []byte) (coretypes.Receipts, error) {
	var receiptsForStorage []*coretypes.ReceiptForStorage
	if err := rlp.DecodeBytes(bz, &receiptsForStorage); err != nil {
		return nil, err
	}
	//#nosec:G103 unsafe pointer is safe here since `ReceiptForStorage` is an alias of `Receipt`.
	return *(*coretypes.Receipts)(unsafe.Pointer(&receiptsForStorage)), nil
}
