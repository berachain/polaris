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

// EthAdressFromAccString converts a string acc address to an Ethereum `Address`. NOTE: Do NOT use
// for val or cons address.
func EthAdressFromAccString(accAddress string) (common.Address, error) {
	addr, err := sdk.AccAddressFromBech32(accAddress)
	if err != nil {
		return common.Address{}, err
	}
	return common.BytesToAddress(addr), nil
}

// MustEthAddressFromAccString converts a string acc address to an Ethereum `Address`.
// Panics if conversion fails. NOTE: Do NOT use for val or cons address.
func MustEthAddressFromAccString(accAddress string) common.Address {
	addr, err := EthAdressFromAccString(accAddress)
	if err != nil {
		panic(err)
	}
	return addr
}

// AccStringFromEthAddress converts Ethereum `Address` to a string. NOTE: Do NOT use for val
// or cons address.
func AccStringFromEthAddress(ethAddr common.Address) (string, error) {
	accStr := sdk.AccAddress(ethAddr.Bytes()).String()
	if accStr == "" {
		return "", errors.New("empty bech32 address")
	}
	return accStr, nil
}

// MustAccStringFromEthAddress converts Ethereum `Address` to a string. Panics if the
// conversion fails. NOTE: Do NOT use for val or cons address.
func MustAccStringFromEthAddress(ethAddr common.Address) string {
	accStr, err := AccStringFromEthAddress(ethAddr)
	if err != nil {
		panic(err)
	}
	return accStr
}

///////////////////////////////////////////////////////////////////////////////
// ValAddress and ConsAddress
///////////////////////////////////////////////////////////////////////////////

// EthAddressFromString converts a Cosmos SDK (val/cons)address string to an Ethereum `Address`.
func EthAddressFromString(codec address.Codec, addr string) (common.Address, error) {
	bz, err := codec.StringToBytes(addr)
	if err != nil {
		return common.Address{}, err
	}
	return common.BytesToAddress(bz), nil
}

// MustEthAddressFromString converts a Cosmos SDK (val/cons)address string to an Ethereum
// `Address`. It panics if the conversion fails.
func MustEthAddressFromString(codec address.Codec, addr string) common.Address {
	address, err := EthAddressFromString(codec, addr)
	if err != nil {
		panic(err)
	}
	return address
}

// StringFromEthAddress converts an Ethereum `Address` to a Cosmos SDK (val/cons)address string.
func StringFromEthAddress(codec address.Codec, ethAddress common.Address) (string, error) {
	return codec.BytesToString(ethAddress.Bytes())
}

// MustStringFromEthAddress converts an Ethereum `Address` to a Cosmos SDK (val/cons)address
// string. It panics if the conversion fails.
func MustStringFromEthAddress(codec address.Codec, ethAddress common.Address) string {
	addr, err := StringFromEthAddress(codec, ethAddress)
	if err != nil {
		panic(err)
	}
	return addr
}
