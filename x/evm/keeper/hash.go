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

package keeper

// import (
// 	"context"
// 	"math"

// 	"github.com/berachain/stargazer/eth/common"
// 	sdk "github.com/cosmos/cosmos-sdk/types"

// )

// //===============================================================
// // BlockHash Functions
// //===============================================================z

// // `StargazerBlockHashAtHeight` returns the header at the given height.
// func (k Keeper) StargazerBlockHashAtHeight(gctx context.Context, height uint64) common.Hash {
// 	ctx := sdk.UnwrapSDKContext(gctx)
// 	if height > math.MaxInt64 {
// 		panic("height is greater than max int64")
// 	}
// 	h := int64(height)
// 	ctxHeight := ctx.BlockHeight()
// 	switch {
// 	case ctxHeight == h:
// 		return k.BlockHashFromCosmosContext(ctx)
// 	case ctxHeight > h:
// 		return k.BlockHashFromHistoricalInfo(ctx, h)
// 	default:
// 		return common.Hash{}
// 	}
// }

// // `BlockHashFromSdkContext` extracts the block hash from a Cosmos context.
// func (k *Keeper) BlockHashFromCosmosContext(ctx sdk.Context) common.Hash {
// 	headerHash := ctx.HeaderHash()
// 	if len(headerHash) != 0 {
// 		return common.BytesToHash(headerHash)
// 	}

// 	// only recompute the hash if not set (eg: checkTxState)
// 	contextBlockHeader := ctx.BlockHeader()
// 	header, err := tmtypes.HeaderFromProto(&contextBlockHeader)
// 	if err != nil {
// 		k.Logger(ctx).Error("failed to cast comet header from proto", "error", err)
// 		return common.Hash{}
// 	}

// 	return common.BytesToHash(header.Hash())
// }

// // `HashFromSdkContext` extracts the block has using historical information saved
// // in the staking keeper.
// func (k *Keeper) BlockHashFromHistoricalInfo(ctx sdk.Context, height int64) common.Hash {
// 	histInfo, found := k.stakingKeeper.GetHistoricalInfo(ctx, height)
// 	if !found {
// 		k.Logger(ctx).Debug("historical info not found", "height", height)
// 		return common.Hash{}
// 	}

// 	header, err := tmtypes.HeaderFromProto(&histInfo.Header)
// 	if err != nil {
// 		k.Logger(ctx).Error("failed to cast tendermint header from proto", "error", err)
// 		return common.Hash{}
// 	}

// 	return common.BytesToHash(header.Hash())
// }
