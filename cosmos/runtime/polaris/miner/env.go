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

package miner

import (
	"github.com/ethereum/go-ethereum/core/types"

	"pkg.berachain.dev/polaris/eth/common"
)

// environment is the worker's current environment and holds all
// information of the sealing block generation.
type environment struct {
	signer   types.Signer
	header   *types.Header
	coinbase common.Address
	// TODO: Use more fields
	// ancestors mapset.Set[common.Hash] // ancestor set (used for checking uncle parent validity)
	// family    mapset.Set[common.Hash] // family set (used for checking uncle invalidity)
	// tcount   int           // tx count in cycle
	// gasPool  *core.GasPool // available gas used to pack transactions

	// txs      []*types.Transaction
	// receipts []*types.Receipt
	// uncles   map[common.Hash]*types.Header
}

// // copy creates a deep copy of environment.
// func (env *environment) copy() *environment {
// 	cpy := &environment{
// 		signer: env.signer,
// 		header: types.CopyHeader(env.header),
// 		// tcount:   env.tcount,
// 		// coinbase: env.coinbase,

// 		// receipts: copyReceipts(env.receipts),
// 	}
// 	// if env.gasPool != nil {
// 	// 	gasPool := *env.gasPool
// 	// 	cpy.gasPool = &gasPool
// 	// }
// 	// // The content of txs and uncles are immutable, unnecessary
// 	// // to do the expensive deep copy for them.
// 	// // cpy.txs = make([]*types.Transaction, len(env.txs))
// 	// copy(cpy.txs, env.txs)
// 	// cpy.uncles = make(map[common.Hash]*types.Header)
// 	// for hash, uncle := range env.uncles {
// 	// 	cpy.uncles[hash] = uncle
// 	// }
// 	return cpy
// }

// copyReceipts makes a deep copy of the given receipts.
// func copyReceipts(receipts []*types.Receipt) []*types.Receipt {
// 	result := make([]*types.Receipt, len(receipts))
// 	for i, l := range receipts {
// 		cpy := *l
// 		result[i] = &cpy
// 	}
// 	return result
// }
