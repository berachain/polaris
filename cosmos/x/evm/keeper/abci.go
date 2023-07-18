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

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/types"
)

func (k *Keeper) BeginBlocker(ctx context.Context) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	if k.lock {
		k.GetHost().GetTxPoolPlugin().SetLogger(k.Logger(sdkCtx))
		k.lock = false
	}

	// Prepare the Polaris Ethereum block.
	k.polaris.Prepare(ctx, uint64(sdkCtx.BlockHeight()))
	return nil
}

func (k *Keeper) EndBlock(ctx context.Context) error {
	// Finalize the Polaris Ethereum block.
	if err := k.polaris.Finalize(ctx); err != nil {
		return err
	}

	// Prepare the txpool for the next pending block.
	// TODO: move this.
	parent, err := k.GetHost().GetBlockPlugin().GetHeaderByNumber(
		uint64(sdk.UnwrapSDKContext(ctx).BlockHeight()),
	)
	if err != nil {
		return err
	}
	chainConfig := k.host.GetConfigurationPlugin().ChainConfig()
	k.GetHost().GetTxPoolPlugin().Prepare(
		misc.CalcBaseFee(chainConfig, parent), types.LatestSigner(chainConfig), // TODO: use MakeSigner.
	)

	return nil
}
