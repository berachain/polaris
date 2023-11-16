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

	"github.com/berachain/polaris/eth/common"
	"github.com/berachain/polaris/eth/common/hexutil"
	"github.com/berachain/polaris/eth/core/types"
	"github.com/berachain/polaris/eth/params"

	"github.com/ethereum/go-ethereum/core"
)

type (
	Genesis        = core.Genesis
	GenesisAlloc   = core.GenesisAlloc
	GenesisAccount = core.GenesisAccount
)

// DefaultGenesis is the default genesis block used by Polaris.
var DefaultGenesis = &core.Genesis{
	// Chain Config
	Config: params.DefaultChainConfig,

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
		// 0xbac
		common.HexToAddress("0xFE94Cc9f0dfbb657a6C6850701aBF6356227F8c3"): {
			Balance: big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(1e7)), //nolint:gomnd // its okay.
		},
		common.HexToAddress("0x11e2E77c864BAcCF47E8D70dA82f15426BEc7816"): {
			Balance: big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(1e7)), //nolint:gomnd // its okay.
		},
		common.HexToAddress("0x318D5326BBbabaBb208531cAC6B29aB116497179"): {
			Balance: big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(1e7)), //nolint:gomnd // its okay.
		},
		common.HexToAddress("0x7B856C6D250eED55D2D7543ae2169a1cd7f034Ad"): {
			Balance: big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(1e7)), //nolint:gomnd // its okay.
		},
		common.HexToAddress("0x08D9255C2922528da6e8853319bcc85A1f6e283c"): {
			Balance: big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(1e7)), //nolint:gomnd // its okay.
		},
		common.HexToAddress("0xF6581Da6b4e27A6eA0aD60C2b31FDD0B34b04FF7"): {
			Balance: big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(1e7)), //nolint:gomnd // its okay.
		},
		common.HexToAddress("0xD8F62DB27ae97a22914b01BAA229502124A4597b"): {
			Balance: big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(1e7)), //nolint:gomnd // its okay.
		},
		common.HexToAddress("0xfeC42ac9FB61185f43697194fBcA8ff726cCaE7B"): {
			Balance: big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(1e7)), //nolint:gomnd // its okay.
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
func UnmarshalGenesisHeader(header *types.Header, gen *Genesis) {
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
