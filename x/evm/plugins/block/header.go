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
	"errors"
	"fmt"
	"math/big"

	cbft "github.com/cometbft/cometbft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"pkg.berachain.dev/stargazer/eth/common"
	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/eth/rpc"
)

var (
	SGHeaderKey = []byte("SGHeaderKey")
)

// `SetQueryContextFn` sets the query context func for the plugin.
func (p *plugin) SetQueryContextFn(gqc func(height int64, prove bool) (sdk.Context, error)) {
	p.getQueryContext = gqc
}

// `ProcessHeader` takes in the header and process it using the `ctx` and stores it in the context store.
func (p *plugin) ProcessHeader(ctx sdk.Context, header *coretypes.StargazerHeader) error {
	header = p.fillHeader(ctx, header)
	bz, err := header.MarshalBinary()
	if err != nil {
		return err
	}
	ctx.KVStore(p.storekey).Set(SGHeaderKey, bz)
	return nil
}

// `GetStargazerHeaderByNumber` returns the stargazer header for the given block number.
func (p *plugin) GetStargazerHeaderByNumber(number int64) (*coretypes.StargazerHeader, error) {
	if p.getQueryContext == nil {
		return nil, fmt.Errorf("query context not set")
	}

	iavlHeight, err := p.getIAVLHeight(number)
	if err != nil {
		return nil, err
	}

	ctx, err := p.getQueryContext(iavlHeight, false)
	if err != nil {
		return nil, err
	}

	// Unmarshal the header from the context kv store.
	var header coretypes.StargazerHeader
	bz := ctx.KVStore(p.storekey).Get(SGHeaderKey)
	if bz == nil {
		return nil, errors.New("stargazer header not found")
	}
	if err := header.UnmarshalBinary(bz); err != nil {
		return nil, err
	}
	return &header, nil
}

// `getIAVLHeight` returns the IAVL height for the given block number.
func (p *plugin) getIAVLHeight(number int64) (int64, error) {
	var iavlHeight int64
	switch rpc.BlockNumber(number) {
	case rpc.SafeBlockNumber:
	case rpc.FinalizedBlockNumber:
		iavlHeight = p.ctx.BlockHeight() - 1
	case rpc.PendingBlockNumber:
	case rpc.LatestBlockNumber:
		iavlHeight = p.ctx.BlockHeight()
	case rpc.EarliestBlockNumber:
		iavlHeight = 1
	default:
		iavlHeight = number
	}

	if iavlHeight < 0 {
		return 1, fmt.Errorf("invalid block number %d", number)
	}

	return iavlHeight, nil
}

// `fillHeader` takes in a `coretypes.StargazerHeader` and returns a `coretypes.StargazerHeader` with the
// Fields set to the correct values from the `sdk.Context`.
func (p *plugin) fillHeader(ctx sdk.Context, header *coretypes.StargazerHeader) *coretypes.StargazerHeader {
	cometHeader := ctx.BlockHeader()

	// We retrieve the `TxHash` from the `DataHash` field of the `sdk.Context` opposed to deriving it
	// from solely the ethereum transaction information.
	txHash := coretypes.EmptyRootHash
	if len(cometHeader.DataHash) == 0 {
		txHash = common.BytesToHash(cometHeader.DataHash)
	}

	parentHash := common.Hash{}
	if ctx.BlockHeight() > 1 {
		header, err := p.GetStargazerHeaderByNumber(ctx.BlockHeight() - 1)
		if err != nil || header == nil {
			panic("failed to get parent stargazer header")
		}
		parentHash = header.Hash()
	}

	return coretypes.NewStargazerHeader(
		&coretypes.Header{
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
			GasLimit: blockGasLimitFromCosmosContext(ctx),
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
		blockHashFromCosmosContext(ctx),
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
