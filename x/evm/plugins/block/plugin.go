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

	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/eth/core"
	coretypes "github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/x/evm/plugins"
	cbft "github.com/cometbft/cometbft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO: change this.
const bf = uint64(1e9)

// `stargazerHeaderGetter` is an interface that defines the `GetStargazerHeader` method.
type stargazerHeaderGetter interface {
	// `GetStargazerHeader` returns the stargazer header at the given height.
	GetStargazerHeader(ctx sdk.Context, height int64) (*coretypes.StargazerHeader, bool)
}

// `Plugin` is the interface that must be implemented by the plugin.
type Plugin interface {
	plugins.BaseCosmosStargazer
	core.BlockPlugin
}

// `plugin` keeps track of stargazer blocks via headers.
type plugin struct {
	// `ctx` is the current block context, used for accessing current block info and kv stores.
	ctx sdk.Context
	// `shg` is the stargazer header getter, used for accessing stargazer headers.
	shg stargazerHeaderGetter
}

// `NewPlugin` creates a new instance of the block plugin from the given context.
func NewPlugin(shg stargazerHeaderGetter) Plugin {
	return &plugin{
		shg: shg,
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

// `GetStargazerHeader` returns the stargazer header at the given height, using the plugin's
// context.
//
// `GetStargazerHeader` implements core.BlockPlugin.
func (p *plugin) GetStargazerHeaderAtHeight(height int64) *coretypes.StargazerHeader {
	// If the current block height is the same as the requested height, then we assume that the
	// block has not been written to the store yet. In this case, we build and return a header
	// from the sdk.Context.
	if p.ctx.BlockHeight() == height {
		return p.getStargazerHeaderFromCurrentContext()
	}

	// If the current block height is less than (or technically also greater than) the requested
	// height, then we assume that the block has been written to the store. In this case, we
	// return the header from the store.
	if header, found := p.shg.GetStargazerHeader(p.ctx, height); found {
		return header
	}

	return &coretypes.StargazerHeader{}
}

// `getStargazerHeaderFromCurrentContext` builds an ethereum style block header from the current
// context.
func (p *plugin) getStargazerHeaderFromCurrentContext() *coretypes.StargazerHeader {
	cometHeader := p.ctx.BlockHeader()

	// We retrieve the `TxHash` from the `DataHash` field of the `sdk.Context` opposed to deriving it
	// from solely the ethereum transaction information.
	txHash := coretypes.EmptyRootHash
	if len(cometHeader.DataHash) == 0 {
		txHash = common.BytesToHash(cometHeader.DataHash)
	}

	return coretypes.NewStargazerHeader(
		&coretypes.Header{
			// `ParentHash` is set to the hash of the previous block.
			ParentHash: common.BytesToHash(cometHeader.LastBlockId.Hash),
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
		},
		blockHashFromCosmosContext(p.ctx),
	)
}

// blockHashFromCosmosContext returns the block hash from the provided Cosmos SDK context.
// If the context contains a valid header hash, it is converted to a common.Hash and returned.
// Otherwise, if the header hash is not set (e.g., for checkTxState), the hash is computed
// from the context's block header and returned as a common.Hash. If the block header is invalid,
// the function returns an empty common.Hash and logs an error.
func blockHashFromCosmosContext(ctx sdk.Context) common.Hash {
	// Check if the context contains a header hash
	headerHash := ctx.HeaderHash()
	if len(headerHash) != 0 {
		return common.BytesToHash(headerHash)
	}

	// If the header hash is not set, compute the hash from the context's block header
	contextBlockHeader := ctx.BlockHeader()
	header, err := cbft.HeaderFromProto(&contextBlockHeader)
	if err != nil {
		// If the block header is invalid, return an empty hash
		return common.Hash{}
	}

	// Convert the computed hash to a common.Hash and return it
	return common.BytesToHash(header.Hash())
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
