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
	"math/big"

	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	evmtypes "pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/lib/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Header", func() {
	var ctx sdk.Context
	var p *plugin

	BeforeEach(func() {
		_, _, _, sk := testutil.SetupMinimalKeepers()
		ctx = testutil.NewContext().WithBlockGasMeter(storetypes.NewGasMeter(uint64(10000)))
		p = utils.MustGetAs[*plugin](NewPlugin(testutil.EvmKey, sk))
		p.SetQueryContextFn(mockQueryContext)
		p.Prepare(ctx)
	})

	FIt("set and get header", func() {
		// we are on block 10
		p.Prepare(ctx.WithBlockHeight(10))

		// just finished processing block 10
		header := &types.Header{
			ParentHash:  common.Hash{0x01},
			UncleHash:   common.Hash{0x02},
			Coinbase:    common.Address{0x03},
			Root:        common.Hash{0x04},
			TxHash:      common.Hash{0x05},
			ReceiptHash: common.Hash{0x06},
			Number:      big.NewInt(10),
		}
		err := p.StoreHeader(header)
		Expect(err).ToNot(HaveOccurred())

		// now querying
		header2, err := p.GetHeaderByNumber(10)
		Expect(err).To(BeNil())
		Expect(header2.Hash()).To(Equal(header.Hash()))

		// get unknown header should return the latest header (10)
		header3, found := p.GetHeaderByNumber(11)
		Expect(found).To(BeNil())
		Expect(header3.Hash()).To(Equal(header.Hash()))
	})

	FIt("should be able to prune headers", func() {
		p.Prepare(ctx.WithBlockHeight(1))
		header := &types.Header{
			ParentHash:  common.Hash{0x01},
			UncleHash:   common.Hash{0x02},
			Coinbase:    common.Address{0x03},
			Root:        common.Hash{0x04},
			TxHash:      common.Hash{0x05},
			ReceiptHash: common.Hash{0x06},
			Number:      big.NewInt(1),
		}
		err := p.StoreHeader(header)
		Expect(err).To(BeNil())

		for i := 1; i <= 260; i++ {
			p.Prepare(ctx.WithBlockHeight(int64(i)))
			ctx = ctx.WithBlockHeight(int64(i))
			header := &types.Header{Number: big.NewInt(int64(i))}
			err := p.StoreHeader(header)
			Expect(err).ToNot(HaveOccurred())
		}

		// Run TrackHistoricalPolarisHeader on the header with height 260.
		ctx = ctx.WithBlockHeight(260)
		err = p.StoreHeader(&types.Header{Number: big.NewInt(260)})
		Expect(err).ToNot(HaveOccurred())

		// Get the header with height 1.
		_, err = p.GetHeaderByNumber(1)
		Expect(err).To(BeNil())

		// Get the header with height 10.
		_, err = p.GetHeaderByNumber(10)
		Expect(err).To(BeNil())
	})
})

func mockQueryContext(height int64, prove bool) (sdk.Context, error) {
	ctx := testutil.NewContext().WithBlockHeight(height)
	header := &types.Header{
		ParentHash:  common.Hash{0x01},
		UncleHash:   common.Hash{0x02},
		Coinbase:    common.Address{0x03},
		Root:        common.Hash{0x04},
		TxHash:      common.Hash{0x05},
		ReceiptHash: common.Hash{0x06},
		Number:      big.NewInt(height),
	}
	headerBz, err := types.MarshalHeader(header)
	if err != nil {
		return sdk.Context{}, err
	}
	ctx.KVStore(testutil.EvmKey).Set([]byte{evmtypes.HeaderKey}, headerBz)
	ctx.KVStore(testutil.EvmKey).Set(header.Hash().Bytes(), header.Number.Bytes())
	return ctx, nil
}
