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
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"

	"pkg.berachain.dev/polaris/cosmos/x/erc20/types"
	"pkg.berachain.dev/polaris/eth/common"
)

// DenomKVStore is the store type for ERC20 token address <-> SDK Coin denominations.
type DenomKVStore interface {
	SetAddressDenomPair(address common.Address, denom string)
	GetDenomForAddress(address common.Address) string
	HasDenomForAddress(address common.Address) bool
	GetAddressForDenom(denom string) common.Address
	HasAddressForDenom(denom string) bool
}

// denomStore is a store that stores information regarding ERC20 token address <-> SDK Coin
// denominations.
type denomStore struct {
	addressToDenom storetypes.KVStore
	denomToAddress storetypes.KVStore
}

// NewDenomKVStore creates a new DenomKVStore.
func NewDenomKVStore(store storetypes.KVStore) DenomKVStore {
	return &denomStore{
		addressToDenom: prefix.NewStore(store, []byte{types.AddressToDenomKeyPrefix}),
		denomToAddress: prefix.NewStore(store, []byte{types.DenomToAddressKeyPrefix}),
	}
}

// SetAddressDenomPair sets the ERC20 address <-> SDK coin denomination pair.
func (ds *denomStore) SetAddressDenomPair(address common.Address, denom string) {
	bz := []byte(denom)
	ds.addressToDenom.Set(address.Bytes(), bz)
	ds.denomToAddress.Set(bz, address.Bytes())
}

// ==============================================================================
// ERC20 -> Denom
// ==============================================================================

// GetDenomForAddress returns the denomination correlated to a specific address.
func (ds *denomStore) GetDenomForAddress(address common.Address) string {
	bz := ds.addressToDenom.Get(address.Bytes())
	if bz == nil {
		return ""
	}
	return string(bz)
}

// HasDenomForAddress returns true if the address has a denomination.
func (ds *denomStore) HasDenomForAddress(address common.Address) bool {
	return ds.addressToDenom.Has(address.Bytes())
}

// ==============================================================================
// Denom -> ERC20
// ==============================================================================

// GetAddressForDenom returns the address correlated to a specific denomination.
func (ds *denomStore) GetAddressForDenom(denom string) common.Address {
	bz := ds.denomToAddress.Get([]byte(denom))
	if bz == nil {
		return common.Address{}
	}
	return common.BytesToAddress(bz)
}

// HasAddressForDenom returns true if the denomination has an address.
func (ds *denomStore) HasAddressForDenom(denom string) bool {
	return ds.denomToAddress.Has([]byte(denom))
}
