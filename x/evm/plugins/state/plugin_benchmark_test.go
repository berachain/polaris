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

package state_test

import (
	"math/big"
	"testing"

	ethstate "github.com/berachain/stargazer/eth/core/state"
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/testutil"
	"github.com/berachain/stargazer/x/evm/plugins/state"
	"github.com/ethereum/go-ethereum/common"
)

var (
	numCalls           = 10000 // number of times snapshot is called
	numStoreOpsPerCall = 10    // number of read/write ops on stores during each call
	numReverts         = 2     // number of times an eth call is reverted in one tx
)

func GetNewStatePlugin() ethstate.Plugin {
	ctx, ak, bk, _ := testutil.SetupMinimalKeepers()
	return state.NewPlugin(ctx, ak, bk, testutil.EvmKey, "abera")
}

func GetNewStateDB() vm.StargazerStateDB {
	ctx, ak, bk, _ := testutil.SetupMinimalKeepers()
	return ethstate.NewStateDB(state.NewPlugin(ctx, ak, bk, testutil.EvmKey, "abera"))
}

func BenchmarkArbitraryStateTransition(b *testing.B) {
	sdb := GetNewStateDB()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var snapshots []int
		for c := 0; c < numCalls; c++ {
			sdb.SetNonce(testutil.Bob, uint64(c+19)) // accStore set
			sdb.SetState(                            // ethStore set
				testutil.Alice,
				common.BytesToHash([]byte{byte(c + 11)}),
				common.BytesToHash([]byte{byte(c + 22)}),
			)

			snapshots = append(snapshots, sdb.Snapshot())
			for s := 0; s < numStoreOpsPerCall; s++ {
				sdb.GetBalance(testutil.Alice)               // bankStore read
				sdb.AddBalance(testutil.Bob, big.NewInt(10)) // bankStore write
				sdb.GetCode(testutil.Alice)                  // ethStore read
			}
			if c > 0 && numReverts > 0 && c%(numCalls/numReverts) == 0 {
				sdb.RevertToSnapshot(snapshots[len(snapshots)/2])
			}

			sdb.Suicide(testutil.Alice) // will invoke a delete
		}

		// commit only once
		sdb.Finalize()
	}
}
