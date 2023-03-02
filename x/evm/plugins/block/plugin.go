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
	"context"
	"math/big"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/core"
	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/x/evm/plugins"
)

// TODO: change this.
const bf = uint64(1)

// `Plugin` is the interface that must be implemented by the plugin.
type Plugin interface {
	plugins.BaseCosmosStargazer
	core.BlockPlugin

	// `UpdateOffChainStorage` updates the offchain storage with the new block and receipts.
	UpdateOffChainStorage(*coretypes.Block, coretypes.Receipts)
	// `SetHeader` saves a block to the store.
	SetHeader(header *coretypes.Header) error
	// `SetQueryContextFn` sets the function used for querying historical block headers.
	SetQueryContextFn(fn func(height int64, prove bool) (sdk.Context, error))
}

// `plugin` keeps track of stargazer blocks via headers.
type plugin struct {
	// `ctx` is the current block context, used for accessing current block info and kv stores.
	ctx sdk.Context
	// `storekey` is the store key for the header store.
	storekey storetypes.StoreKey
	//  `offchainStore` is the offchain store, used for accessing offchain data.
	offchainStore storetypes.CacheKVStore
	// `getQueryContext` allows for querying block headers.
	getQueryContext func(height int64, prove bool) (sdk.Context, error)
}

// `NewPlugin` creates a new instance of the block plugin from the given context.
func NewPlugin(offchainStore storetypes.CacheKVStore, storekey storetypes.StoreKey) Plugin {
	return &plugin{
		offchainStore: offchainStore,
		storekey:      storekey,
	}
}

// `Prepare` implements core.BlockPlugin.
func (p *plugin) Prepare(ctx context.Context) {
	p.ctx = sdk.UnwrapSDKContext(ctx)
}

// `BaseFee` returns the base fee for the current block.
// TODO: implement properly with DynamicFee Module of some kind.
//
// `BaseFee` implements core.BlockPlugin.
func (p *plugin) BaseFee() uint64 {
	return bf
}

// `NewHeaderWithBlockNumber` builds an ethereum style block header from the current
// context.
func (p *plugin) NewHeaderWithBlockNumber(number int64) *coretypes.Header {
	cometHeader := p.ctx.BlockHeader()

	if cometHeader.Height != number {
		panic("block height mismatch")
	}

	// We retrieve the `TxHash` from the `DataHash` field of the `sdk.Context` opposed to deriving it
	// from solely the ethereum transaction information.
	txHash := coretypes.EmptyRootHash
	if len(cometHeader.DataHash) == 0 {
		txHash = common.BytesToHash(cometHeader.DataHash)
	}

	parentHash := common.Hash{}
	if p.ctx.BlockHeight() > 1 {
		if header, err := p.GetHeaderByNumber(p.ctx.BlockHeight() - 1); err == nil {
			parentHash = header.Hash()
		} else {
			panic("parent header not found")
		}
	}

	return &coretypes.Header{
		// `ParentHash` is set to the hash of the previous block.
		ParentHash: parentHash,
		// `UncleHash` is set empty as CometBFT does not have uncles.
		UncleHash: coretypes.EmptyUncleHash,
		// TODO: Use staking keeper to get the operator address.
		Coinbase: common.BytesToAddress(cometHeader.ProposerAddress),
		// `Root` is set to the hash of the state after the transactions are applied.
		Root: common.BytesToHash(cometHeader.AppHash),
		// `TxHash` is set to the hash of the transactions in the block. We take the
		// `DataHash` from the `sdk.Context` opposed to using DeriveSha on the StargazerBlock,
		// in order to include non-evm transactions block hash.
		TxHash: txHash,
		// We simply map the cosmos "BlockHeight" to the ethereum "BlockNumber".
		Number: big.NewInt(cometHeader.Height),
		// `GasLimit` is set to the block gas limit.
		GasLimit: blockGasLimitFromCosmosContext(p.ctx),
		// `Time` is set to the block timestamp.
		Time: uint64(cometHeader.Time.UTC().Unix()),
		// `BaseFee` is set to the block base fee.
		BaseFee: big.NewInt(int64(p.BaseFee())),
		// `ReceiptHash` set to empty. It is filled during `Finalize` in the StateProcessor.
		ReceiptHash: common.Hash{},
		// `Bloom` is set to empty. It is filled during `Finalize` in the StateProcessor.
		Bloom: coretypes.Bloom{},
		// `GasUsed` is set to 0. It is filled during `Finalize` in the StateProcessor.
		GasUsed: 0,
		// `Difficulty` is set to 0 as it is only used in PoW consensus.
		Difficulty: big.NewInt(0),
		// `MixDigest` is set empty as it is only used in PoW consensus.
		MixDigest: common.Hash{},
		// `Nonce` is set empty as it is only used in PoW consensus.
		Nonce: coretypes.BlockNonce{},
		// `Extra` is unused in Stargazer.
		Extra: []byte(nil),
	}
}

// `blockGasLimitFromCosmosContext` returns the maximum gas limit for the current block, as defined
// by either the block gas meter or the consensus parameters if the gas meter is not set or is an
// InfiniteGasMeter. If neither the gas meter nor the consensus parameters are available, it
// returns 0. This shouldn't be an issue in practice but we include this function for completeness
// defensive programming purposes.
func blockGasLimitFromCosmosContext(ctx sdk.Context) uint64 {
	blockGasMeter := ctx.BlockGasMeter()
	if blockGasMeter == nil || blockGasMeter.Limit() == 0 {
		cp := ctx.ConsensusParams()
		if cp == nil || cp.Block == nil {
			return 0
		}
		return uint64(cp.Block.MaxGas)
	}

	return blockGasMeter.Limit()
}
