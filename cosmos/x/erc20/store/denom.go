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

package store

import (
	"errors"
	"fmt"

	storetypes "cosmossdk.io/store/types"

	"pkg.berachain.dev/polaris/cosmos/x/erc20/types"
	"pkg.berachain.dev/polaris/eth/common"
)

var ErrDenomNotFound = errors.New("denom not found")
var ErrAddressNotFound = errors.New("address not found")

// DenomERC20 is the store type for ERC20 <-> x/bank module denominations.
type DenomKVStore interface {
	SetDenomForAddress(address common.Address, denom string)
	GetDenomForAddress(address common.Address) (string, error)
	HasDenomForAddress(address common.Address) bool
	SetAddressForDenom(denom string, address common.Address)
	GetAddressForDenom(denom string) (common.Address, error)
	HasAddressForDenom(denom string) bool
}

// denomStore is a store that stores information regarding ERC20 <-> x/bank module denominations.
type denomStore struct {
	storetypes.KVStore
}

// NewDenomKVStore creates a new denomStore.
func NewDenomKVStore(store storetypes.KVStore) DenomKVStore {
	return &denomStore{store}
}

// ==============================================================================
// ERC20 -> Denom
// ==============================================================================

// SetDenomForAddress sets the address correlated to a specific denomStore.
func (ds denomStore) SetDenomForAddress(address common.Address, denom string) {
	ds.Set(DenomToAddressKey(address), []byte(types.DenomForAddress(address)))
}

// GetDenomForAddress returns the denomStore correlated to a specific address.
func (ds denomStore) GetDenomForAddress(address common.Address) (string, error) {
	bz := ds.Get(DenomToAddressKey(address))
	if bz == nil {
		return "", ErrDenomNotFound
	}
	return string(bz), nil
}

// HasDenomForAddress returns true if the address has a denomStore.
func (ds denomStore) HasDenomForAddress(address common.Address) bool {
	return ds.Has(DenomToAddressKey(address))
}

// ==============================================================================
// Denom -> ERC20
// ==============================================================================

// SetAddressForDenom sets the denomStore correlated to a specific address.
func (ds denomStore) SetAddressForDenom(denom string, address common.Address) {
	ds.Set(AddressForDenomKey(denom), address.Bytes())
}

// GetAddressForDenom returns the address correlated to a specific denomStore.
func (ds denomStore) GetAddressForDenom(denom string) (common.Address, error) {
	bz := ds.Get(AddressForDenomKey(denom))
	if bz == nil {
		return common.Address{}, fmt.Errorf("no address for denom %s", denom)
	}
	return common.BytesToAddress(bz), nil
}

// HasAddressForDenom returns true if the denomStore has an address.
func (ds denomStore) HasAddressForDenom(denom string) bool {
	return ds.Has(AddressForDenomKey(denom))
}
