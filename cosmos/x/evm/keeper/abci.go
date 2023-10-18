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

import (
	"context"
)

// Precommit runs on the Cosmos-SDK lifecycle Precommit().
func (k *Keeper) Precommit(_ context.Context) error {
	// Verify that the EVM block was written.
	// 	// TODO: Set/GetHead to set and get the canonical head.
	// 	blockNum := uint64(sdk.UnwrapSDKContext(ctx).BlockHeight())
	// 	block := k.executionClient.Eth.GetBlockByNumber(blockNum)
	// // The code block is performing a verification check on the EVM block.
	// It first checks if the EVM block
	// // is nil, which means that the block was not successfully written. If it is
	//  nil, it throws a panic
	// // with an error message indicating the failure at the specific block number.
	// 	if block == nil {
	// 		panic(
	// 			fmt.Sprintf("EVM BLOCK FAILURE AT BLOCK %d", blockNum),
	// 		)
	// 	} else if block.NumberU64() != blockNum {
	// 		panic(
	// 			fmt.Sprintf(
	// 				"EVM BLOCK [%d] DOES NOT MATCH COMET BLOCK [%d]", block.NumberU64(), blockNum,
	// 			),
	// 		)
	// 	}
	return nil
}

// PrepareCheckState runs on the Cosmos-SDK lifecycle PrepareCheckState().
func (k *Keeper) PrepareCheckState(_ context.Context) error {
	return nil
}
