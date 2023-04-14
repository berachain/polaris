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

package types

import (
	fmt "fmt"

	"pkg.berachain.dev/polaris/eth/common"
)

const (
	// polarisDenomPrefix represents the bank module prefix all Polaris coin denominations will
	// have (ERC20 originated token).
	polarisDenomPrefix = "polaris/"

	// lenPolarisDenomPrefix is the length of the polarisDenomPrefix.
	lenPolarisDenomPrefix = 8

	// lenPolarisDenom is the length of the (polarisDenomPrefix + 20 bytes + "0x") for the address.
	lenPolarisDenom = 50
)

// NewPolarisDenomForAddress returns a new Polaris coin denomination for a given ERC20 originated
// token address.
func NewPolarisDenomForAddress(token common.Address) string {
	return fmt.Sprintf("%s%s", polarisDenomPrefix, token.Hex())
}

// IsPolarisDenom returns true if the given denom is a Polaris coin denomination.
func IsPolarisDenom(denom string) bool {
	return len(denom) == lenPolarisDenom && denom[:lenPolarisDenomPrefix] == polarisDenomPrefix
}
