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

package mempool

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	sdk "github.com/cosmos/cosmos-sdk/types"

	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
)

func TestEthPool(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/x/evm/plugins/txpool/mempool")
}

var _ = Describe("EthTxPool", func() {
	var sp core.StatePlugin
	var etp *EthTxPool
	var addr1 = common.HexToAddress("0x1111")
	var addr2 = common.HexToAddress("0x2222")

	BeforeEach(func() {
		ctx, ak, bk, _ := testutil.SetupMinimalKeepers()
		sp = state.NewPlugin(ak, bk, testutil.EvmKey, &mockConfigurationPlugin{}, &mockPLF{})
		sp.Reset(ctx)
		sp.SetNonce(addr1, 1111)
		sp.SetNonce(addr2, 2222)
		sp.Finalize()
		sp.Reset(ctx)
		etp = NewEthTxPoolFrom(DefaultPriorityMempool())
		etp.SetNonceRetriever(sp)
	})

	Describe("All Cases", func() {
		It("should handle empty txs", func() {
			Expect(etp.Get(common.Hash{})).To(BeNil())
			emptyPending, emptyQueued := etp.Pending(false), etp.queued()
			Expect(emptyPending).To(BeEmpty())
			Expect(emptyQueued).To(BeEmpty())
			Expect(etp.Nonce(addr1)).To(Equal(uint64(1111)))
			Expect(etp.Nonce(addr2)).To(Equal(uint64(2222)))
			Expect(etp.Nonce(common.HexToAddress("0x3333"))).To(Equal(uint64(0)))
		})

		It("should return pending/queued txs with correct nonces", func() {

		})
	})
})

// MOCKS BELOW.

type mockConfigurationPlugin struct{}

func (mcp *mockConfigurationPlugin) GetEvmDenom() string {
	return "abera"
}

type mockPLF struct{}

func (mplf *mockPLF) Build(event *sdk.Event) (*coretypes.Log, error) {
	return &coretypes.Log{
		Address: common.BytesToAddress([]byte(event.Type)),
	}, nil
}
