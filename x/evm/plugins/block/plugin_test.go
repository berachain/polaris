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
	"time"

	storetypes "cosmossdk.io/store/types"
	dbm "github.com/cosmos/cosmos-db"
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/lib/utils"
	"pkg.berachain.dev/stargazer/store/offchain"
	"pkg.berachain.dev/stargazer/testutil"
)

var _ = Describe("Block Plugin", func() {
	var ctx sdk.Context
	var p *plugin

	BeforeEach(func() {
		ctx = testutil.NewContext().WithBlockGasMeter(storetypes.NewGasMeter(uint64(10000)))
		sk := testutil.EvmKey // testing key.
		p = utils.MustGetAs[*plugin](NewPlugin(offchain.NewFromDB(dbm.NewMemDB()), sk))
		p.Prepare(ctx)
		p.SetQueryContextFn(func(_ int64, _ bool) (sdk.Context, error) {
			return ctx, nil
		})
	})

	It("should give the constant base fee", func() {
		Expect(p.BaseFee()).To(Equal(bf))
	})

	It("should get the header at current height", func() {
		now := time.Now()
		ctx = ctx.WithBlockTime(now)
		expectedTime := uint64(ctx.BlockHeader().Time.UTC().Unix()) // what the processor will do.
		emptyHeader := types.StargazerHeader{}
		err := p.ProcessHeader(ctx, &emptyHeader) // The processor will set the values.
		Expect(err).To(BeNil())
		header, err := p.GetStargazerHeaderByNumber(ctx.BlockHeight())
		Expect(err).To(BeNil())
		Expect(header.Time).To(Equal(expectedTime))
	})

	It("should get the header at a previous height", func() {
		emptyHeader := types.StargazerHeader{}
		firstCtx := testutil.NewContext().WithBlockHeight(1)
		secondCtx := testutil.NewContext().WithBlockHeight(2)
		thirdCtx := testutil.NewContext().WithBlockHeight(3)

		// Set the query context to return the different contexts.
		p.SetQueryContextFn(func(height int64, _ bool) (sdk.Context, error) {
			if height == 0 {
				return ctx, nil
			}

			if height == 1 {
				return firstCtx, nil
			}

			if height == 2 {
				return secondCtx, nil
			}

			return thirdCtx, nil
		})

		// Set the first header and second header.
		p.Prepare(firstCtx)
		err := p.ProcessHeader(firstCtx, &emptyHeader)
		Expect(err).To(BeNil())

		// Check that the first header is set.
		header, err := p.GetStargazerHeaderByNumber(1)
		Expect(err).To(BeNil())
		Expect(header.Number.Uint64()).To(Equal(uint64(1)))

		// Set the second header.
		p.Prepare(secondCtx)
		p.ProcessHeader(secondCtx, &emptyHeader)

		// Check that the second header is set.
		header, err = p.GetStargazerHeaderByNumber(2)
		Expect(err).To(BeNil())
		Expect(header.Number.Uint64()).To(Equal(uint64(2)))

		// Set the third header.
		p.Prepare(thirdCtx)
		err = p.ProcessHeader(thirdCtx, &emptyHeader)
		Expect(err).To(BeNil())

		// Check that the third header is set.
		header, err = p.GetStargazerHeaderByNumber(3)
		Expect(err).To(BeNil())
		Expect(header.Number.Uint64()).To(Equal(uint64(3)))

		// Check that the first header is still set.
		header, err = p.GetStargazerHeaderByNumber(1)
		Expect(err).To(BeNil())
		Expect(header.Number.Uint64()).To(Equal(uint64(1)))
	})
})
