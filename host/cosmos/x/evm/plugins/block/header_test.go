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

package block

import (
	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/lib/utils"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Header", func() {
	var ctx sdk.Context
	var p *plugin

	BeforeEach(func() {
		ctx = testutil.NewContext().WithBlockGasMeter(storetypes.NewGasMeter(uint64(10000)))
		p = utils.MustGetAs[*plugin](NewPlugin(testutil.EvmKey))
		p.Prepare(ctx)
	})

	// It("set and get header", func() {
	// 	header := &types.Header{
	// 		ParentHash:  common.Hash{0x01},
	// 		UncleHash:   common.Hash{0x02},
	// 		Coinbase:    common.Address{0x03},
	// 		Root:        common.Hash{0x04},
	// 		TxHash:      common.Hash{0x05},
	// 		ReceiptHash: common.Hash{0x06},
	// 		Number:      big.NewInt(10),
	// 	}
	// 	err := p.SetHeader(header)
	// 	Expect(err).ToNot(HaveOccurred())

	// 	header2, found := p.GetHeaderByNumber(10)
	// 	Expect(found).To(BeTrue())
	// 	Expect(header2.Hash()).To(Equal(header.Hash()))

	// 	// get unknown header
	// 	header3, found := p.GetHeaderByNumber(11)
	// 	Expect(found).To(BeFalse())
	// 	Expect(header3).To(BeNil())
	// })

	// It("should be able to prune headers", func() {
	// 	header := &types.Header{
	// 		ParentHash:  common.Hash{0x01},
	// 		UncleHash:   common.Hash{0x02},
	// 		Coinbase:    common.Address{0x03},
	// 		Root:        common.Hash{0x04},
	// 		TxHash:      common.Hash{0x05},
	// 		ReceiptHash: common.Hash{0x06},
	// 		Number:      big.NewInt(10),
	// 	}
	// 	err := p.SetHeader(header)
	// 	Expect(err).ToNot(HaveOccurred())

	// 	// Get header.
	// 	header2, found := p.GetHeaderByNumber(10)
	// 	Expect(found).To(BeFalse())
	// 	Expect(header2).To(BeNil())
	// })

	// It("should be able to track the headers", func() {
	// 	for i := 1; i <= 260; i++ {
	// 		ctx = ctx.WithBlockHeight(int64(i))
	// 		header := &types.Header{Number: big.NewInt(int64(i))}
	// 		err := p.SetHeader(header)
	// 		Expect(err).ToNot(HaveOccurred())
	// 	}

	// 	// Run TrackHistoricalPolarisHeader on the header with height 260.
	// 	ctx = ctx.WithBlockHeight(260)
	// 	err := p.SetHeader(&types.Header{Number: big.NewInt(260)})
	// 	Expect(err).ToNot(HaveOccurred())

	// 	// Get the header with height 1.
	// 	_, found := p.GetHeaderByNumber(1)
	// 	Expect(found).To(BeFalse())

	// 	// Get the header with height 10.
	// 	_, found = p.GetHeaderByNumber(10)
	// 	Expect(found).To(BeTrue())
	// })
})
