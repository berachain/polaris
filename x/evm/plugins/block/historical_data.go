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
	"github.com/berachain/stargazer/eth/common"
	coretypes "github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/store/ethrlp"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// `UpdateOffChainStorage` is called by the `EndBlocker` to update the off-chain storage.
func (p *plugin) UpdateOffChainStorage(ctx sdk.Context, block *coretypes.StargazerBlock) {
	parent := p.offchainStore

	blockStore := ethrlp.NewRlpEncodedStore[*coretypes.StargazerBlock](parent, []byte("blocks"))
	blockStore.Set(block)

	// adding txns to kv.
	txStore := ethrlp.NewRlpEncodedStore[*coretypes.Transaction](parent, []byte("tx"))
	for _, tx := range block.GetTransactions() {
		txStore.Set(tx)
	}

	version := block.Number
	lastVersion := parent.Get(versionKeyPrefix)
	if sdk.BigEndianToUint64(lastVersion) != version.Uint64()-1 {
		panic("REEE")
	}
	parent.Set(versionKeyPrefix, sdk.Uint64ToBigEndian(uint64(version.Int64())))
	// flush the underlying buffer to disk.
	parent.Write()
}

var versionKeyPrefix = []byte("version")

// `GetStargazerHeader` returns the stargazer header at the given height.
func (p *plugin) GetStargazerBlockByNumber(height int64) *coretypes.StargazerBlock {
	// Get the stargazer header at the given height.
	// header, ok := p.shg.GetStargazerHeader(p.ctx, height)
	// if !ok {
	// 	return nil, fmt.Errorf("stargazer header not found at height %d", height)
	// }
	// // Get the stargazer block at the given height.
	// block, ok := p.shg.GetStargazerBlock(p.ctx, height)
	// if !ok {
	// 	return nil, fmt.Errorf("stargazer block not found at height %d", height)
	// }
	// // Return the stargazer block.
	// return &StargazerBlock{
	// 	Header: header,
	// 	Block:  block,
	// }, nil
	return nil
}

func (p *plugin) GetStargazerBlockByHash(hash common.Hash) *coretypes.StargazerBlock {
	// // Get the stargazer header at the given height.
	// header, ok := p.shg.GetStargazerHeaderByHash(p.ctx, hash)
	// if !ok {
	// 	return nil
	// }
	// Return the stargazer block.
	return nil
}
