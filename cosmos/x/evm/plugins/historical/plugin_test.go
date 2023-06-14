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

package historical

import (
	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/core/mock"
	"pkg.berachain.dev/polaris/lib/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Historical Data", func() {
	var (
		p   *plugin
		ctx sdk.Context
	)

	BeforeEach(func() {
		ctx = testutil.NewContext().WithBlockHeight(0)
		cp := mock.NewConfigurationPluginMock()
		bp := mock.NewBlockPluginMock()

		p = utils.MustGetAs[*plugin](NewPlugin(cp, bp, nil, testutil.EvmKey))
		p.InitGenesis(ctx, core.DefaultGenesis)
	})

	When("Genesis block", func() {
		It("should return the header without error", func() {
			block, err := p.GetBlockByNumber(0)
			Expect(err).ToNot(HaveOccurred())
			header := block.Header()
			Expect(header).ToNot(BeNil())
			blockByHash, err := p.GetBlockByHash(block.Hash())
			Expect(err).ToNot(HaveOccurred())
			Expect(blockByHash).ToNot(BeNil())
			Expect(blockByHash.Hash()).To(Equal(block.Hash()))
		})
	})

	When("Other blocks", func() {

		// It("should get the header at current height", func() {
		// 	header, err := p.GetHeaderByNumber(ctx.BlockHeight())
		// 	Expect(err).ToNot(HaveOccurred())
		// 	Expect(header.TxHash).To(Equal(common.BytesToHash(ctx.BlockHeader().DataHash)))
		// })

		// It("should return empty header for non-existent height", func() {
		// 	header, err := p.GetHeaderByNumber(100000)
		// 	Expect(err).ToNot(HaveOccurred())
		// 	Expect(*header).To(Equal(types.Header{}))
		// })
	})

})
