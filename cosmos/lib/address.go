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

import (
	"errors"

	"cosmossdk.io/core/address"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/eth/common"
)

///////////////////////////////////////////////////////////////////////////////
// AccAddress
///////////////////////////////////////////////////////////////////////////////

// EthAddressFromAccBech32 converts a Bech32 string acc address to an Ethereum `Address`. NOTE: Do
// NOT use for val or cons address.
func EthAddressFromAccBech32(accAddress string) (common.Address, error) {
	addrBech32, err := sdk.AccAddressFromBech32(accAddress)
	if err != nil {
		return common.Address{}, err
	}
	return common.BytesToAddress(addrBech32), nil
}

// MustEthAddressFromAccBech32 converts a Bech32 string acc address to an Ethereum `Address`.
// Panics if conversion fails. NOTE: Do NOT use for val or cons address.
func MustEthAddressFromAccBech32(accAddress string) common.Address {
	addr, err := EthAddressFromAccBech32(accAddress)
	if err != nil {
		panic(err)
	}
	return addr
}

// AccBech32FromEthAddress converts Ethereum `Address` to a Bech32 string. NOTE: Do NOT use for val
// or cons address.
func AccBech32FromEthAddress(ethAddr common.Address) (string, error) {
	accBech32 := sdk.AccAddress(ethAddr.Bytes()).String()
	if accBech32 == "" {
		return "", errors.New("empty bech32 address")
	}
	return accBech32, nil
}

// MustAccBech32FromEthAddress converts Ethereum `Address` to a Bech32 string. Panics if the
// conversion fails. NOTE: Do NOT use for val or cons address.
func MustAccBech32FromEthAddress(ethAddr common.Address) string {
	accBech32, err := AccBech32FromEthAddress(ethAddr)
	if err != nil {
		panic(err)
	}
	return accBech32
}

///////////////////////////////////////////////////////////////////////////////
// ValAddress and ConsAddress
///////////////////////////////////////////////////////////////////////////////

// EthAddressFromBech32 converts a Cosmos SDK (val/cons)address bech32 to an Ethereum `Address`.
func EthAddressFromBech32(codec address.Codec, bech32 string) (common.Address, error) {
	bz, err := codec.StringToBytes(bech32)
	if err != nil {
		return common.Address{}, err
	}
	return common.BytesToAddress(bz), nil
}

// MustEthAddressFromBech32 converts a Cosmos SDK (val/cons)address bech32 to an Ethereum
// `Address`. It panics if the conversion fails.
func MustEthAddressFromBech32(codec address.Codec, valAddress string) common.Address {
	address, err := EthAddressFromBech32(codec, valAddress)
	if err != nil {
		panic(err)
	}
	return address
}

// Bech32FromEthAddress converts an Ethereum `Address` to a Cosmos SDK (val/cons)address bech32.
func Bech32FromEthAddress(codec address.Codec, ethAddress common.Address) (string, error) {
	return codec.BytesToString(ethAddress.Bytes())
}

// MustBech32FromEthAddress converts an Ethereum `Address` to a Cosmos SDK (val/cons)address
// bech32. It panics if the conversion fails.
func MustBech32FromEthAddress(codec address.Codec, ethAddress common.Address) string {
	bech32, err := Bech32FromEthAddress(codec, ethAddress)
	if err != nil {
		panic(err)
	}
	return bech32
}
