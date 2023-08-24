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
	"cosmossdk.io/core/address"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/eth/common"
)

// AccAddressToEthAddress converts a Cosmos SDK `AccAddress` to an Ethereum `Address`.
func AccAddressToEthAddress(accAddress sdk.AccAddress) common.Address {
	return common.BytesToAddress(accAddress)
}

// EthAddressFromBech32 converts a Bech32 string to an Ethereum `Address`.
// Note: Do NOT use for val or cons address.
func EthAddressFromBech32(bech32Str string) common.Address {
	addrBech32, _ := sdk.AccAddressFromBech32(bech32Str)
	return AccAddressToEthAddress(addrBech32)
}

// Bech32FromEthAddress converts Ethereum `Address` to a Bech32 string.
// Note: Do NOT use for val or cons address.
func Bech32FromEthAddress(ethAddr common.Address) string {
	return AddressToAccAddress(ethAddr).String()
}

// ConsAddressToEthAddress converts a Cosmos SDK `ConsAddress` to an Ethereum `Address`.
func ConsAddressToEthAddress(codec address.Codec, consAddress string) (common.Address, error) {
	consBz, err := codec.StringToBytes(consAddress)
	if err != nil {
		return common.Address{}, err
	}
	return common.BytesToAddress(consBz), nil
}

// MustConsAddressToEthAddress converts a Cosmos SDK `ConsAddress` to an Ethereum `Address`.
// It panics if the conversion fails.
func MustConsAddressToEthAddress(codec address.Codec, consAddress string) common.Address {
	address, err := ConsAddressToEthAddress(codec, consAddress)
	if err != nil {
		panic(err)
	}
	return address
}

// ValAddressToEthAddress converts a Cosmos SDK `ValAddress` to an Ethereum `Address`.
func ValAddressToEthAddress(codec address.Codec, valAddress string) (common.Address, error) {
	valBz, err := codec.StringToBytes(valAddress)
	if err != nil {
		return common.Address{}, err
	}
	return common.BytesToAddress(valBz), nil
}

// MustValAddressToEthAddress converts a Cosmos SDK `ValAddress` to an Ethereum `Address`.
// It panics if the conversion fails.
func MustValAddressToEthAddress(codec address.Codec, valAddress string) common.Address {
	address, err := ValAddressToEthAddress(codec, valAddress)
	if err != nil {
		panic(err)
	}
	return address
}

// AddressToAccAddress converts an Ethereum `Address` to a Cosmos SDK `AccAddress`.
func AddressToAccAddress(ethAddress common.Address) sdk.AccAddress {
	return ethAddress.Bytes()
}

// AddressToConsAddress converts an Ethereum `Address` to a Cosmos SDK `ConsAddress`.
func AddressToConsAddress(codec address.Codec, ethAddress common.Address) (sdk.ConsAddress, error) {
	ethStr, err := codec.BytesToString(ethAddress.Bytes())
	if err != nil {
		return sdk.ConsAddress{}, err
	}
	return sdk.ConsAddressFromBech32(ethStr)
}

// MustAddressToConsAddress converts an Ethereum `Address` to a Cosmos SDK `ConsAddress`.
// It panics if the conversion fails.
func MustAddressToConsAddress(codec address.Codec, ethAddress common.Address) sdk.ConsAddress {
	consAddr, err := AddressToConsAddress(codec, ethAddress)
	if err != nil {
		panic(err)
	}
	return consAddr
}

// AddressToValAddress converts an Ethereum `Address` to a Cosmos SDK `ValAddress`.
func AddressToValAddress(codec address.Codec, ethAddress common.Address) (sdk.ValAddress, error) {
	ethStr, err := codec.BytesToString(ethAddress.Bytes())
	if err != nil {
		return sdk.ValAddress{}, err
	}
	return sdk.ValAddressFromBech32(ethStr)
}

// MustAddressToValAddress converts an Ethereum `Address` to a Cosmos SDK `ValAddress`.
// It panics if the conversion fails.
func MustAddressToValAddress(codec address.Codec, ethAddress common.Address) sdk.ValAddress {
	valAddr, err := AddressToValAddress(codec, ethAddress)
	if err != nil {
		panic(err)
	}
	return valAddr
}
