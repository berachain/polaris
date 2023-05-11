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

package lib

// Borrowed from Osmosis
// https://github.com/osmosis-labs/osmosis/blob/5297ea4b641c0c1090cb253a34a510dd0c85b407/osmoutils/module_account.go
// Allows the safe creation of module accounts in the event that an account has already been initialized at an address

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"pkg.berachain.dev/polaris/lib/utils"
)

type AccountKeeper interface {
	NewAccount(context.Context, sdk.AccountI) sdk.AccountI
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI
	SetAccount(context.Context, sdk.AccountI)
}

// CanCreateModuleAccountAtAddr tells us if we can safely make a module account at
// a given address. By collision resistance of the address (given API safe construction),
// the only way for an account to be already be at this address is if its claimed by the same
// pre-image from the correct module,
// or some SDK command breaks assumptions and creates an account at designated address.
// This function checks if there is an account at that address, and runs some safety checks
// to be extra-sure its not a user account (e.g. non-zero sequence, pubkey, of fore-seen account types).
// If there is no account, or if we believe its not a user-spendable account, we allow module account
// creation at the address.
// else, we do not.
func CanCreateModuleAccountAtAddr(ctx sdk.Context, ak AccountKeeper, addr sdk.AccAddress) error {
	existingAcct := ak.GetAccount(ctx, addr)
	if existingAcct == nil || utils.Implements[sdk.ModuleAccountI](existingAcct) {
		// if account doesn't exist or is already a module account, we are good to go.
		return nil
	}
	if existingAcct.GetSequence() != 0 || existingAcct.GetPubKey() != nil {
		return fmt.Errorf("cannot create module account %s, "+
			"due to an account at that address already existing & having sent txs", addr)
	}
	overrideAccountTypes := map[reflect.Type]struct{}{
		reflect.TypeOf(&authtypes.BaseAccount{}):                 {},
		reflect.TypeOf(&vestingtypes.DelayedVestingAccount{}):    {},
		reflect.TypeOf(&vestingtypes.ContinuousVestingAccount{}): {},
		reflect.TypeOf(&vestingtypes.BaseVestingAccount{}):       {},
		reflect.TypeOf(&vestingtypes.PeriodicVestingAccount{}):   {},
		reflect.TypeOf(&vestingtypes.PermanentLockedAccount{}):   {},
	}

	if _, clear := overrideAccountTypes[reflect.TypeOf(existingAcct)]; clear {
		return nil
	}
	return errors.New("cannot create module account %s, " +
		"due to an account at that address already existing & not being an overrideable type")
}

// CreateModuleAccount creates a module account at the provided address.
// It overrides an account if it exists at that address, with a non-zero sequence number & pubkey
// Contract: addr is derived from `address.Module(ModuleName, key)`
func CreateModuleAccount(
	ctx sdk.Context, ak AccountKeeper, name string, permissions ...string,
) error {
	addr := authtypes.NewModuleAddress(name)
	err := CanCreateModuleAccountAtAddr(ctx, ak, addr)
	if err != nil {
		return err
	}

	acc := ak.NewAccount(
		ctx,
		authtypes.NewModuleAccount(
			authtypes.NewBaseAccountWithAddress(addr),
			name,
			permissions...,
		),
	)
	ak.SetAccount(ctx, acc)
	return nil
}
