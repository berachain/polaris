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
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state"
	"pkg.berachain.dev/polaris/eth/common"
)

// TODO: Remove these, because they're hacky af.
// Required temporarily for BGT plugin in the Berachain repo, these REALLY
// REALLY REALLY REALLY NEED TO GO, IF YOU ARE BUILDING YOUR OWN POLARIS CHAIN
// PLEASE DO NOT USE THIS FUNCTIONS DIRECTLY.

func (k *Keeper) GetBalance(ctx sdk.Context, addr sdk.AccAddress) *big.Int {
	ethAddr := common.BytesToAddress(addr)
	return new(big.Int).SetBytes(ctx.KVStore(k.storeKey).Get(state.BalanceKeyFor(ethAddr)))
}

func (k *Keeper) SetBalance(ctx sdk.Context, addr sdk.AccAddress, amount *big.Int) {
	ethAddr := common.BytesToAddress(addr)
	ctx.KVStore(k.storeKey).Set(state.BalanceKeyFor(ethAddr), amount.Bytes())
}

func (k *Keeper) AddBalance(ctx sdk.Context, addr sdk.AccAddress, amount *big.Int) {
	if amount.Sign() == 0 {
		return
	}
	ethAddr := common.BytesToAddress(addr)
	ctx.KVStore(k.storeKey).Set(
		state.BalanceKeyFor(ethAddr), new(big.Int).Add(k.GetBalance(ctx, addr), amount).Bytes())
}

func (k *Keeper) SubBalance(ctx sdk.Context, addr sdk.AccAddress, amount *big.Int) {
	if amount.Sign() == 0 {
		return
	}
	ethAddr := common.BytesToAddress(addr)
	ctx.KVStore(k.storeKey).Set(
		state.BalanceKeyFor(ethAddr), new(big.Int).Sub(k.GetBalance(ctx, addr), amount).Bytes())
}
