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
	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"pkg.berachain.dev/stargazer/eth/common"
	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
)

// `UpdateOffChainStorage` is called by the `EndBlocker` to update the off-chain storage.
func (p *plugin) UpdateOffChainStorage(ctx sdk.Context, block *coretypes.StargazerBlock) {
	blockStore := prefix.NewStore(p.offchainStore, []byte("blocks"))
	bz, err := block.MarshalBinary()
	if err != nil {
		panic(err)
	}
	numBz := sdk.Uint64ToBigEndian(uint64(block.Number.Int64()))
	blockStore.Set(block.Hash().Bytes(), numBz)
	blockStore.Set(numBz, bz)

	// adding txns to kv.
	txStore := prefix.NewStore(p.offchainStore, []byte("tx"))
	for _, tx := range block.GetTransactions() {
		bz, err = tx.MarshalBinary()
		if err != nil {
			panic(err)
		}
		txStore.Set(tx.Hash().Bytes(), bz)
	}

	version := block.Number
	lastVersion := p.offchainStore.Get(versionKeyPrefix)
	if sdk.BigEndianToUint64(lastVersion) != version.Uint64()-1 {
		panic(err)
	}
	p.offchainStore.Set(versionKeyPrefix, sdk.Uint64ToBigEndian(uint64(version.Int64())))
	// flush the underlying buffer to disk.
	p.offchainStore.Write()
}

var versionKeyPrefix = []byte("version")

// `GetStargazerBlockByNumber` returns the stargazer header at the given height.
func (p *plugin) GetStargazerBlockByNumber(number int64) *coretypes.StargazerBlock {
	blockStore := prefix.NewStore(p.offchainStore, []byte("blocks"))
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
	blockStore := prefix.NewStore(p.offchainStore, []byte("blocks"))
	bz := blockStore.Get(hash.Bytes())
	if bz == nil {
		return nil
	}
	return p.GetStargazerBlockByNumber(int64(sdk.BigEndianToUint64(bz)))
}
