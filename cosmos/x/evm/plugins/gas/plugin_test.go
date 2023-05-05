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

package gas

import (
	"math"

	storetypes "cosmossdk.io/store/types"

	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/lib/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("plugin", func() {
	var p *plugin
	var blockGasMeter storetypes.GasMeter
	var blockGasLimit = uint64(1500)

	BeforeEach(func() {
		// new block
		blockGasMeter = storetypes.NewGasMeter(blockGasLimit)
		p = utils.MustGetAs[*plugin](NewPlugin())
	})

	It("correctly consume, refund, and report cumulative in the same block", func() {
		ctx := testutil.NewContext().WithBlockGasMeter(blockGasMeter)
		// tx 1
		ctx = ctx.WithGasMeter(storetypes.NewGasMeter(1000))
		err := p.ConsumeGas(ctx, 500)
		Expect(err).ToNot(HaveOccurred())
		Expect(ctx.GasMeter().GasConsumed()).To(Equal(uint64(500)))
		Expect(ctx.GasMeter().GasRemaining()).To(Equal(uint64(500)))

		ctx.GasMeter().RefundGas(250, "test")
		Expect(ctx.GasMeter().GasConsumed()).To(Equal(uint64(250)))
		Expect(p.BlockGasConsumed(ctx)).To(Equal(uint64(0))) // shouldn't have consumed block gas,
		// as block gas is handled by the baseapp.
		blockGasMeter.ConsumeGas(250, "") // finalize tx 1
		ctx = testutil.NewContext().WithBlockGasMeter(blockGasMeter)

		// tx 2
		ctx = ctx.WithGasMeter(storetypes.NewGasMeter(1000))
		Expect(p.BlockGasConsumed(ctx)).To(Equal(uint64(250)))
		err = p.ConsumeGas(ctx, 1000)
		Expect(err).ToNot(HaveOccurred())
		Expect(ctx.GasMeter().GasConsumed()).To(Equal(uint64(1000)))
		Expect(ctx.GasMeter().GasRemaining()).To(Equal(uint64(0)))
		Expect(p.BlockGasConsumed(ctx)).To(Equal(uint64(250))) // shouldn't have consumed any additional gas yet.
		blockGasMeter.ConsumeGas(1000, "")                     // finalize tx 2
		Expect(p.BlockGasConsumed(ctx)).To(Equal(uint64(1250)))
		ctx = testutil.NewContext().WithBlockGasMeter(blockGasMeter)

		// tx 3
		ctx = ctx.WithGasMeter(storetypes.NewGasMeter(1000))
		Expect(p.BlockGasConsumed(ctx)).To(Equal(uint64(1250)))
		err = p.ConsumeGas(ctx, 250)
		Expect(err).ToNot(HaveOccurred())
		Expect(ctx.GasMeter().GasConsumed()).To(Equal(uint64(250)))
		Expect(ctx.GasMeter().GasRemaining()).To(Equal(uint64(750)))
		blockGasMeter.ConsumeGas(250, "") // finalize tx 3
		Expect(p.BlockGasConsumed(ctx)).To(Equal(blockGasLimit))
	})

	It("should error on overconsumption in tx", func() {
		ctx := testutil.NewContext().WithBlockGasMeter(blockGasMeter)
		ctx = ctx.WithGasMeter(storetypes.NewGasMeter(1000))
		ctx = ctx.WithBlockGasMeter(storetypes.NewGasMeter(ctx.GasMeter().GasRemaining() * 2))
		err := p.ConsumeGas(ctx, ctx.GasMeter().GasRemaining())
		Expect(err).ToNot(HaveOccurred())
		err = p.ConsumeGas(ctx, 1)
		Expect(err.Error()).To(Equal("out of gas"))
	})

	It("should error on uint64 overflow", func() {
		ctx := testutil.NewContext().WithBlockGasMeter(storetypes.NewInfiniteGasMeter())
		err := p.ConsumeGas(ctx, math.MaxUint64)
		Expect(err).ToNot(HaveOccurred())
		err = p.ConsumeGas(ctx, 1)
		Expect(err.Error()).To(Equal("gas uint64 overflow"))
	})

	It("should error on block gas overconsumption", func() {
		ctx := testutil.NewContext().WithBlockGasMeter(blockGasMeter)
		Expect(p.BlockGasLimit(ctx)).To(Equal(ctx.BlockGasMeter().Limit()))

		ctx = testutil.NewContext().WithBlockGasMeter(blockGasMeter)
		// tx 1
		err := p.ConsumeGas(ctx, 1000)
		Expect(err).ToNot(HaveOccurred())
		blockGasMeter.ConsumeGas(1000, "") // finalize tx 1

		ctx = testutil.NewContext().WithBlockGasMeter(blockGasMeter)
		// tx 2
		err = p.ConsumeGas(ctx, 1000)
		Expect(err.Error()).To(Equal("block is out of gas"))
	})
})
