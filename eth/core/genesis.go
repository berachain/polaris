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

package core

import (
	"math/big"

	"pkg.berachain.dev/polaris/eth/common/hexutil"
	"pkg.berachain.dev/polaris/eth/params"
)

var DefaultGenesis = &Genesis{
	Config:     params.DefaultChainConfig,
	Nonce:      69, //nolint:gomnd // its okay.
	ExtraData:  hexutil.MustDecode("0x11bbe8db4e347b4e8c937c1c8370e4b5ed33adb3db69cbdb7a38e1e50b1b82fa"),
	GasLimit:   30_000_000,     //nolint:gomnd // its okay.
	Difficulty: big.NewInt(69), //nolint:gomnd // its okay.
	Alloc:      GenesisAlloc{},
	// Alloc:      decodePrealloc("mainnetAllocData"),
	// For alloc, in the startup / initGenesis, we should allow the host chain to "fill in the data"
	// i.e in Cosmos, we let the AccountKeeper/EVMKeeper/BankKeeper fill in the Bank Data into the
	// genesis and then verify the equivalency later. This is to create an invariant that the bank
	// balances from the bank keeper and the token balances in the EVM are equivalents at genesis.
}

// func decodePrealloc(data string) GenesisAlloc {
// 	var p []struct{ Addr, Balance *big.Int }
// 	if err := rlp.NewStream(strings.NewReader(data), 0).Decode(&p); err != nil {
// 		panic(err)
// 	}
// 	ga := make(GenesisAlloc, len(p))
// 	for _, account := range p {
// 		ga[common.BigToAddress(account.Addr)] = GenesisAccount{Balance: account.Balance}
// 	}
// 	return ga
// }
