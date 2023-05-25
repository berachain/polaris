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

	"cosmossdk.io/core/address"
	"pkg.berachain.dev/polaris/eth/common"
)

var _ address.Codec

type EIP55AddressEncoder struct {
	// contains filtered or unexported fields
}

func NewEIP55AddressEncoder() *EIP55AddressEncoder {
	return &EIP55AddressEncoder{}
}

func (e EIP55AddressEncoder) BytesToString(address []byte) (string, error) {
	fmt.Println("LETS GOOOO 1")
	return common.Bytes2Hex(address), nil
}

func (e EIP55AddressEncoder) StringToBytes(address string) ([]byte, error) {
	fmt.Println("LETS GOOOO 2")
	return common.Hex2Bytes(address), nil
}
