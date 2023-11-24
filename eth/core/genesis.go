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

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type (
	Genesis        = core.Genesis
	GenesisAlloc   = core.GenesisAlloc
	GenesisAccount = core.GenesisAccount
)

// DefaultGenesis is the default genesis block used by Polaris.
var DefaultGenesis = &core.Genesis{
	// Genesis Block
	Nonce:     0,
	Timestamp: 0,
	ExtraData: hexutil.MustDecode(
		"0x11bbe8db4e347b4e8c937c1c8370e4b5ed33adb3db69cbdb7a38e1e50b1b82fa"),
	GasLimit:   30_000_000, //nolint:gomnd // its okay.
	Difficulty: big.NewInt(0),
	Mixhash:    common.Hash{},
	Coinbase:   common.Address{},

	// Genesis Accounts
	Alloc: core.GenesisAlloc{
		// 0xfffdbb37105441e14b0ee6330d855d8504ff39e705c3afa8f859ac9865f99306
		common.HexToAddress("0x20f33CE90A13a4b5E7697E3544c3083B8F8A51D4"): {
			Balance: big.NewInt(0).Mul(big.NewInt(5e18), big.NewInt(100)), //nolint:gomnd // its okay.
		},
	},

	// These fields are used for consensus tests. Please don't use them
	// in actual genesis blocks.
	Number:        0,
	GasUsed:       0,
	ParentHash:    common.Hash{},
	BaseFee:       nil,
	ExcessBlobGas: nil,
	BlobGasUsed:   nil,

	// These fields are used for consensus tests. Please don't use them
	// in actual genesis blocks.
}

// UnmarshalGenesisHeader sets the fields of the given header into the Genesis struct.
func UnmarshalGenesisHeader(header *ethtypes.Header, gen *Genesis) {
	// Note: cannot set the state root on the genesis.
	gen.Nonce = header.Nonce.Uint64()
	gen.Timestamp = header.Time
	gen.ParentHash = header.ParentHash
	gen.ExtraData = header.Extra
	gen.GasLimit = header.GasLimit
	gen.GasUsed = header.GasUsed
	gen.BaseFee = header.BaseFee
	gen.Difficulty = header.Difficulty
	gen.Mixhash = header.MixDigest
	gen.Coinbase = header.Coinbase
	gen.Number = header.Number.Uint64()
}
