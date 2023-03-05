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

	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/stargazer/lib/utils"
	testutil "pkg.berachain.dev/stargazer/testing/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("plugin", func() {
	var ctx sdk.Context
	var p *plugin
	var blockGasMeter storetypes.GasMeter
	var blockGasLimit = uint64(1500)

	BeforeEach(func() {
		// new block
		blockGasMeter = storetypes.NewGasMeter(blockGasLimit)
		ctx = testutil.NewContext().WithBlockGasMeter(blockGasMeter)
		p = utils.MustGetAs[*plugin](NewPlugin())
		p.Reset(ctx)
		p.Prepare(ctx)
	})

	It("correctly consume, refund, and report cumulative in the same block", func() {
		p.Reset(testutil.NewContext().WithBlockGasMeter(blockGasMeter))

		// tx 1
		p.gasMeter = storetypes.NewGasMeter(1000)
		err := p.ConsumeGas(500)
		Expect(err).ToNot(HaveOccurred())
		Expect(p.gasMeter.GasConsumed()).To(Equal(uint64(500)))
		Expect(p.gasMeter.GasRemaining()).To(Equal(uint64(500)))

		p.gasMeter.RefundGas(250, "test")
		Expect(p.gasMeter.GasConsumed()).To(Equal(uint64(250)))
		Expect(p.BlockGasConsumed()).To(Equal(uint64(0))) // shouldn't have consumed block gas,
		// as block gas is handled by the baseapp.
		blockGasMeter.ConsumeGas(250, "") // finalize tx 1

		p.Reset(testutil.NewContext().WithBlockGasMeter(blockGasMeter))

		// tx 2
		p.gasMeter = storetypes.NewGasMeter(1000)
		Expect(p.BlockGasConsumed()).To(Equal(uint64(250)))
		err = p.ConsumeGas(1000)
		Expect(err).ToNot(HaveOccurred())
		Expect(p.gasMeter.GasConsumed()).To(Equal(uint64(1000)))
		Expect(p.gasMeter.GasRemaining()).To(Equal(uint64(0)))
		Expect(p.BlockGasConsumed()).To(Equal(uint64(250))) // shouldn't have consumed.
		blockGasMeter.ConsumeGas(1000, "")                  // finalize tx 2

		p.Reset(testutil.NewContext().WithBlockGasMeter(blockGasMeter))

		// tx 3
		p.gasMeter = storetypes.NewGasMeter(1000)
		Expect(p.BlockGasConsumed()).To(Equal(uint64(1250)))
		err = p.ConsumeGas(250)
		Expect(err).ToNot(HaveOccurred())
		Expect(p.gasMeter.GasConsumed()).To(Equal(uint64(250)))
		Expect(p.gasMeter.GasRemaining()).To(Equal(uint64(750)))
		Expect(p.BlockGasConsumed()).To(Equal(blockGasLimit))
		blockGasMeter.ConsumeGas(250, "") // finalize tx 3
	})

	It("should error on overconsumption in tx", func() {
		p.gasMeter = storetypes.NewGasMeter(1000)
		p.blockGasMeter = storetypes.NewGasMeter(p.gasMeter.GasRemaining() * 2)
		err := p.ConsumeGas(p.gasMeter.GasRemaining())
		Expect(err).ToNot(HaveOccurred())
		err = p.ConsumeGas(1)
		Expect(err.Error()).To(Equal("out of gas"))
	})

	It("should error on uint64 overflow", func() {
		p.blockGasMeter = storetypes.NewInfiniteGasMeter()
		err := p.ConsumeGas(math.MaxUint64)
		Expect(err).ToNot(HaveOccurred())
		err = p.ConsumeGas(1)
		Expect(err.Error()).To(Equal("gas uint64 overflow"))
	})

	It("should error on block gas overconsumption", func() {
		Expect(p.BlockGasLimit()).To(Equal(blockGasLimit))

		p.Reset(testutil.NewContext().WithBlockGasMeter(blockGasMeter))

		// tx 1
		err := p.ConsumeGas(1000)
		Expect(err).ToNot(HaveOccurred())
		blockGasMeter.ConsumeGas(1000, "") // finalize tx 1

		p.Reset(testutil.NewContext().WithBlockGasMeter(blockGasMeter))

		// tx 2
		err = p.ConsumeGas(1000)
		Expect(err.Error()).To(Equal("block is out of gas"))
	})
})
