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
	"fmt"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/cosmos/x/erc20/store"
	"pkg.berachain.dev/polaris/cosmos/x/erc20/types"
	"pkg.berachain.dev/polaris/eth/common"
)

// Keeper of this module maintains collections of erc20.
type Keeper struct {
	StoreKey   storetypes.StoreKey
	bankKeeper BankKeeper
	authority  sdk.AccAddress
}

// NewKeeper creates new instances of the erc20 Keeper.
func NewKeeper(
	storeKey storetypes.StoreKey,
	bk BankKeeper,
	authority sdk.AccAddress,
) *Keeper {
	return &Keeper{
		StoreKey:   storeKey,
		bankKeeper: bk,
		authority:  authority,
	}
}

// DenomKVStore returns a KVStore for the given denom.
func (k *Keeper) DenomKVStore(ctx sdk.Context) store.DenomKVStore {
	return store.NewDenomKVStore(ctx.KVStore(k.StoreKey))
}

// RegisterERC20CoinPair registers a new ERC20 originated token <> Polaris Coin pair and returns
// the new Polaris Coin denom.
func (k *Keeper) RegisterERC20CoinPair(ctx sdk.Context, token common.Address) string {
	// store the denomination as a Polaris coin denomination.
	polarisDenom := types.NewPolarisDenomForAddress(token)
	k.DenomKVStore(ctx).SetAddressDenomPair(token, polarisDenom)
	return polarisDenom
}

// RegisterCoinERC20Pair registers a new IBC-originated SDK Coin <> ERC20 token pair.
func (k *Keeper) RegisterCoinERC20Pair(ctx sdk.Context, denom string, token common.Address) {
	// store the new ERC20 address for the given denomination.
	k.DenomKVStore(ctx).SetAddressDenomPair(token, denom)
}

// Logger returns a module-specific logger.
func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
