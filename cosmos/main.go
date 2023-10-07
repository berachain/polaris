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

package main

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/eth/common"
)

const prefix = "polar"

// alice 0x6A1680c1Ec339657df9a4c718C8081C52daD5702
// bob 0xbBcec0f8cBAbe76879AfdfD15F3784652B9734C3
// charlie 0xacc1319Fe722A198F395F0164066ED4E309439Bf

func main() {
	if false {
		cosmosAddr := "polar1dgtgps0vxwt90hu6f3cceqypc5k664cz2kml8y"
		sdk.GetConfig().SetBech32PrefixForAccount(prefix, prefix+sdk.PrefixPublic)
		fmt.Println("0x" + common.Bytes2Hex(sdk.MustAccAddressFromBech32(cosmosAddr).Bytes()))
	} else {
		ethAddress := common.HexToAddress("0xacc1319Fe722A198F395F0164066ED4E309439Bf")
		fmt.Println(sdk.MustBech32ifyAddressBytes(prefix, ethAddress.Bytes()))
	}
}

// fmt.Print("alice", tf.Address("alice"), "\n")
// fmt.Print("bob", tf.Address("bob"), "\n")
// fmt.Print("charlie", tf.Address("charlie"), "\n")

// fmt.Println("alice", sdk.MustBech32ifyAddressBytes("polar", tf.Address("alice").Bytes()))
// fmt.Println("bob", sdk.MustBech32ifyAddressBytes("polar", tf.Address("bob").Bytes()))
// fmt.Println("charlie", sdk.MustBech32ifyAddressBytes("polar", tf.Address("charlie").Bytes()))
