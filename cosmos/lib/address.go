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
	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/eth/common"
)

// AccAddressToEthAddress converts a Cosmos SDK `AccAddress` to an Ethereum `Address`.
func AccAddressToEthAddress(accAddress sdk.AccAddress) common.Address {
	return common.BytesToAddress(accAddress)
}

// EthAddressFromBEch32 converts a Bech32 string to an Ethereum `Address`.
func EthAddressFromBech32(bech32Str string) common.Address {
	addrBech32, _ := sdk.AccAddressFromBech32(bech32Str)
	return AccAddressToEthAddress(addrBech32)
}

// EthAddressFromBEch32 converts Ethereum `Address` to a Bech32 string.
func Bech32FromEthAddress(ethAddr common.Address) string {
	return AddressToAccAddress(ethAddr).String()
}

// ConsAddressToEthAddress converts a Cosmos SDK `ConsAddress` to an Ethereum `Address`.
func ConsAddressToEthAddress(consAddress sdk.ConsAddress) common.Address {
	return common.BytesToAddress(consAddress)
}

// ValAddressToEthAddress converts a Cosmos SDK `ValAddress` to an Ethereum `Address`.
func ValAddressToEthAddress(valAddress sdk.ValAddress) common.Address {
	return common.BytesToAddress(valAddress)
}

// AddressToAccAddress converts an Ethereum `Address` to a Cosmos SDK `AccAddress`.
func AddressToAccAddress(ethAddress common.Address) sdk.AccAddress {
	return ethAddress.Bytes()
}

// AddressToConsAddress converts an Ethereum `Address` to a Cosmos SDK `ConsAddress`.
func AddressToConsAddress(ethAddress common.Address) sdk.ConsAddress {
	return ethAddress.Bytes()
}

// AddressToValAddress converts an Ethereum `Address` to a Cosmos SDK `ValAddress`.
func AddressToValAddress(ethAddress common.Address) sdk.ValAddress {
	return ethAddress.Bytes()
}
