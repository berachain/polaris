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

package state_test

import (
	"math/big"
	"testing"

	"cosmossdk.io/log"

	testutil "github.com/berachain/polaris/cosmos/testutil"
	"github.com/berachain/polaris/cosmos/x/evm/plugins/state"
	"github.com/berachain/polaris/eth/core"

	"github.com/ethereum/go-ethereum/common"
)

var (
	numCalls           = 10000 // number of times snapshot is called
	numStoreOpsPerCall = 10    // number of read/write ops on stores during each call
	numReverts         = 2     // number of times an eth call is reverted in one tx
)

func GetNewStatePlugin() core.StatePlugin {
	ctx, ak, _, _ := testutil.SetupMinimalKeepers(log.NewTestLogger(&testing.B{}))
	sp := state.NewPlugin(ak, testutil.EvmKey, nil, nil)
	sp.Reset(ctx)
	return sp
}

func BenchmarkArbitraryStateTransition(b *testing.B) {
	sp := GetNewStatePlugin()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var snapshots []int
		for c := 0; c < numCalls; c++ {
			sp.SetNonce(testutil.Bob, uint64(c+19)) // accStore set
			sp.SetState(                            // ethStore set
				testutil.Alice,
				common.BytesToHash([]byte{byte(c + 11)}),
				common.BytesToHash([]byte{byte(c + 22)}),
			)

			snapshots = append(snapshots, sp.Snapshot())
			for s := 0; s < numStoreOpsPerCall; s++ {
				sp.GetBalance(testutil.Alice)               // bankStore read
				sp.AddBalance(testutil.Bob, big.NewInt(10)) // bankStore write
				sp.GetCode(testutil.Alice)                  // ethStore read
			}
			if c > 0 && numReverts > 0 && c%(numCalls/numReverts) == 0 {
				sp.RevertToSnapshot(snapshots[len(snapshots)/2])
			}

			sp.DeleteAccounts([]common.Address{testutil.Alice}) // will invoke a delete
		}

		// commit only once
		sp.Finalize()
	}
}
