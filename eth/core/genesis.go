// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
