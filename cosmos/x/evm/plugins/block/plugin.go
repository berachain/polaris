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
	"fmt"
	"math/big"

	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core"
)

type Plugin interface {
	plugins.Base
	plugins.HasGenesis
	core.BlockPlugin

	// SetQueryContextFn sets the function used for querying historical block headers.
	SetQueryContextFn(fn func(height int64, prove bool) (sdk.Context, error))
}

// TODO read genesis block from EthGenesis.

type plugin struct {
	// ctx is the current block context, used for accessing current block info and kv stores.
	ctx sdk.Context
	// storekey is the store key for the header store.
	storekey storetypes.StoreKey
	// getQueryContext allows for querying block headers.
	getQueryContext func(height int64, prove bool) (sdk.Context, error)
	// sk represents the cosmos staking keeper.
	sk StakingKeeper
}

func NewPlugin(storekey storetypes.StoreKey, sk StakingKeeper) Plugin {
	return &plugin{
		storekey: storekey,
		sk:       sk,
	}
}

// Prepare implements core.BlockPlugin.
func (p *plugin) Prepare(ctx context.Context) {
	p.ctx = sdk.UnwrapSDKContext(ctx)
}

// BaseFee implements core.BlockPlugin.
func (p *plugin) BaseFee() *big.Int {
	return big.NewInt(-1) // we defer to polaris' built in eip-1559 for the base fee.
}

// GetNewBlockMetadata returns the host chain block metadata for the given block height. It returns
// the coinbase address, the timestamp of the block.
func (p *plugin) GetNewBlockMetadata(number uint64) (common.Address, uint64) {
	cometHeader := p.ctx.BlockHeader()
	if uint64(cometHeader.Height) != number {
		panic(fmt.Errorf("block height mismatch. got: %d, expected %d", cometHeader.Height, number))
	}

	val, found := p.sk.GetValidatorByConsAddr(p.ctx, cometHeader.ProposerAddress)
	if !found {
		panic(fmt.Errorf("validator not found: %s", cometHeader.ProposerAddress))
	}
	return common.BytesToAddress(val.GetOperator()), uint64(cometHeader.Time.UTC().Unix())
}

func (p *plugin) IsPlugin() {}
