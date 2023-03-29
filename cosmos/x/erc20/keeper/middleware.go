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
	"errors"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/x/erc20/store"
	"pkg.berachain.dev/polaris/cosmos/x/erc20/types"
	"pkg.berachain.dev/polaris/eth/common"
)

// HandleIncomingERC20 handles an incoming ERC20 transfer from the EVM to Cosmos, using
// the token address contract as the lookup key.
func (k *Keeper) HandleIncomingERC20(
	ctx sdk.Context, recipient common.Address, token common.Address, amount *big.Int,
) (int8, string, error) {
	// For safety, we check the deployLock.
	if !k.deployLock.TryLock() {
		panic("deploy lock is already locked")
	}
	defer k.deployLock.Unlock()

	// Get the denom corresponding to a given ERC20 token.
	denom, err := k.DenomKVStore(ctx).GetDenomForAddress(token)

	// If the denom is not found, we need to register it, this means that the token
	// began it's life as an ERC20.
	if errors.Is(err, store.ErrDenomNotFound) {
		k.RegisterDenomTokenPair(ctx, token)
	}

	// Mint x/bank coins to the given address.
	return types.ShimHandlerType(denom), denom, lib.MintCoinsToAddress(ctx, k.bankKeeper, recipient, denom, amount)
}

// HandleIncomingERC20 handles an incoming ERC20 transfer.
func (k *Keeper) HandleOutgoingDenom(
	ctx sdk.Context, sender, recipient common.Address, denom string, amount *big.Int,
) (int8, bool, error) {
	// For safety we check the deploy lock.
	if !k.deployLock.TryLock() {
		panic("deploy lock is already locked")
	}

	// Get the ERC20 token corresponding to a given denom.
	if _, err := k.DenomKVStore(ctx).GetAddressForDenom(denom); err != nil {
		// If the token was found, that means we don't have to deploy it. So we need to unlock the lock.
		defer k.deployLock.Unlock()

		// If the denom is found, we need to burn the corresponding bank denom.
		if err = lib.BurnCoinsFromAddress(ctx, k.bankKeeper, sender, denom, amount); err != nil {
			return -1, false, err
		}

		// We return false to signify to the shim that we don't need to trigger the deployment.
		return types.ShimHandlerType(denom), false, nil
	}

	// If the token was not found, it means the smart contract shim needs to follow up with a second call back into
	// the keeper after deploying the contract in order to register it and burn the coins.
	return types.ShimHandlerType(denom), true, nil
}

// HandleDeployed handles the deployment of the ERC20 contract. If HandleOutgoingDenom returns true,
// this means that the smart contract shim needs to handle the deployment of a new ERC20 contract.
// Once the contract is deployed, the shim will call this function to register the new token.
func (k *Keeper) HandleDeployed(ctx sdk.Context, token common.Address) {
	// Once we know that the token was deployed, we can unlock the lock.
	defer k.deployLock.Unlock()
	k.RegisterDenomTokenPair(ctx, token)

	// // For defensive programming practices, we also check the StateDB to ensure that the code exists.
	// if code := k.GetCodeSize(token); codeSize == 0 {
	// 	panic("code not found in state db")
	// }
}
