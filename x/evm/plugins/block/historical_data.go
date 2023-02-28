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
	"math/big"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/stargazer/eth/common"
	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
)

var (
	blockHashKeyPrefix  = []byte{0xb}
	blockNumKeyPrefix   = []byte{0xbb}
	txHashKeyPrefix     = []byte{0x10}
	versionKey          = []byte{0x11}
	txBlockNumKeyPrefix = []byte{0x12}
)

// `UpdateOffChainStorage` is called by the `EndBlocker` to update the off-chain storage.
func (p *plugin) UpdateOffChainStorage(ctx sdk.Context, block *coretypes.StargazerBlock) {
	bz, err := block.MarshalBinary()
	if err != nil {
		panic(err)
	}
	numBz := sdk.Uint64ToBigEndian(block.Number.Uint64())
	prefix.NewStore(p.offchainStore, blockHashKeyPrefix).Set(block.Hash().Bytes(), numBz)
	prefix.NewStore(p.offchainStore, blockNumKeyPrefix).Set(numBz, bz)

	// adding txns to kv.
	txStore := prefix.NewStore(p.offchainStore, txHashKeyPrefix)

	// adding txn with block number to kv.
	txBlockNumStore := prefix.NewStore(p.offchainStore, txBlockNumKeyPrefix)

	for _, tx := range block.GetTransactions() {
		var txBytes []byte
		txBytes, err = tx.MarshalBinary()
		if err != nil {
			panic(err)
		}
		txStore.Set(tx.Hash().Bytes(), txBytes)
		txBlockNumStore.Set(tx.Hash().Bytes(), numBz)
	}

	version := block.Number
	lastVersion := p.offchainStore.Get(versionKey)
	if sdk.BigEndianToUint64(lastVersion) != version.Uint64()-1 {
		// TODO: resync the off-chain storage.
		panic("off-chain store's latest block number is not synced")
	}
	p.offchainStore.Set(versionKey, numBz)
	// flush the underlying buffer to disk.
	p.offchainStore.Write()
}

// `GetStargazerBlockByNumber` returns the stargazer header at the given height.
func (p *plugin) GetStargazerBlockByNumber(number int64) *coretypes.StargazerBlock {
	blockStore := prefix.NewStore(p.offchainStore, blockNumKeyPrefix)
	bz := blockStore.Get(sdk.Uint64ToBigEndian(uint64(number)))
	if bz == nil {
		return nil
	}
	var block coretypes.StargazerBlock
	err := block.UnmarshalBinary(bz)
	if err != nil {
		panic(err)
	}
	return &block
}

// `GetStargazerBlockByHash` returns the stargazer header at the given hash.
func (p *plugin) GetStargazerBlockByHash(hash common.Hash) *coretypes.StargazerBlock {
	blockStore := prefix.NewStore(p.offchainStore, blockHashKeyPrefix)
	bz := blockStore.Get(hash.Bytes())
	if bz == nil {
		return nil
	}
	return p.GetStargazerBlockByNumber(new(big.Int).SetBytes(bz).Int64())
}

// `GetTransactionByHash` returns the transaction with the given hash.
func (p *plugin) GetTransactionByHash(hash common.Hash) *coretypes.Transaction {
	txStore := prefix.NewStore(p.offchainStore, txHashKeyPrefix)
	bz := txStore.Get(hash.Bytes())
	if bz == nil {
		return nil
	}
	var tx coretypes.Transaction
	err := tx.UnmarshalBinary(bz)
	if err != nil {
		p.ctx.Logger().Error("failed to unmarshal transaction", "err", err)
		return nil
	}
	return &tx
}

// `GetTransactionBlockNumber` returns the block number of the transaction with the given hash.
func (p *plugin) GetTransactionBlockNumber(txHash common.Hash) *big.Int {
	txBlockNumStore := prefix.NewStore(p.offchainStore, txBlockNumKeyPrefix)
	bz := txBlockNumStore.Get(txHash.Bytes())
	if bz == nil {
		return nil
	}
	return new(big.Int).SetBytes(bz)
}

// `GetBlockHash` returns the block hash for the given block number.
func (p *plugin) GetBlockHash(blockNum *big.Int) common.Hash {
	blockNumStore := prefix.NewStore(p.offchainStore, blockNumKeyPrefix)
	data := blockNumStore.Get(sdk.Uint64ToBigEndian(blockNum.Uint64()))
	var block *coretypes.StargazerBlock
	err := block.UnmarshalBinary(data)
	if err != nil {
		panic(err)
	}
	return block.Hash()
}
