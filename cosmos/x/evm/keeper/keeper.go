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

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/eth"

	"pkg.berachain.dev/polaris/eth/common"
)

var stateRootKey = []byte{0x69}

type Keeper struct {
	eth      *eth.Ethereum
	storeKey storetypes.StoreKey
}

// NewKeeper creates new instances of the polaris Keeper.
func NewKeeper(
	storeKey storetypes.StoreKey,
) *Keeper {
	return &Keeper{
		storeKey: storeKey,
	}
}

// SetEthereum sets the ethereum instance to be used by the keeper.
func (k *Keeper) SetEthereum(eth *eth.Ethereum) {
	k.eth = eth
}

// Logger returns a logger for the given context.
func (k *Keeper) Logger(ctx context.Context) log.Logger {
	return sdk.UnwrapSDKContext(ctx).Logger().With("module", "polaris/evm")
}

// StoreStateRoot persists the given state root in the store.
func (k *Keeper) StoreStateRoot(ctx context.Context, stateRoot common.Hash) {
	sdk.UnwrapSDKContext(ctx).KVStore(k.storeKey).Set(
		stateRootKey, stateRoot.Bytes(),
	)
}

// GetStateRoot returns the state root of the store.
func (k *Keeper) GetStateRoot(ctx context.Context) common.Hash {
	return common.BytesToHash(
		sdk.UnwrapSDKContext(ctx).KVStore(k.storeKey).Get(
			stateRootKey,
		),
	)
}
