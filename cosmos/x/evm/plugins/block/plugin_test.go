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
	"log"
	"math/big"

	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/eth/common"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Block Plugin", func() {
	var (
		p   *plugin
		ctx sdk.Context
		sk  stakingkeeper.Keeper

		header *coretypes.Header
		err    error

		header1 = &coretypes.Header{Number: big.NewInt(0)}
		header2 = &coretypes.Header{Number: big.NewInt(1)}
	)

	BeforeEach(func() {
		ctx, _, _, sk = testutil.SetupMinimalKeepers()
		storekey := storetypes.NewKVStoreKey("evm")

		// TODO: setup query context func.
		p = &plugin{
			storekey: storekey,
			sk:       sk,
			getQueryContext: func(height int64, prove bool) (sdk.Context, error) {
				return p.ctx, nil
			},
		}

		p.Prepare(ctx)

		// Add two blocks to the block store.
		if err = p.StoreHeader(header1); err != nil {
			log.Panic("failed to store header", "err", err)
		}

		if err = p.StoreHeader(header2); err != nil {
			log.Panic("failed to store header", "err", err)
		}
	})

	Context("GetHeaderByNumber", func() {
		var (
			number uint64
		)

		JustBeforeEach(func() {
			header, err = p.GetHeaderByNumber(number)
		})

		When("the number is 0", func() {
			BeforeEach(func() {
				number = 0
			})

			It("should return the header at the genesis block number 0", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(header.Number.Uint64()).To(Equal(number))
			})
		})

		When("the number is not 0", func() {
			BeforeEach(func() {
				number = 1
			})

			It("should return the header at the given block number", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(header.Number.Uint64()).To(Equal(number))
			})
		})
	})
	Context("GetHeaderByHash", func() {
		var (
			hash common.Hash
		)

		JustBeforeEach(func() {
			header, err = p.GetHeaderByHash(hash)
		})

		When("the hash refers to the genesis block", func() {
			BeforeEach(func() {
				hash = header1.Hash()
			})

			It("should return the header at the genesis block hash", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(header.Hash()).To(Equal(hash))
			})
		})

		When("the hash does not refer to the genesis block", func() {
			BeforeEach(func() {
				hash = header2.Hash()
			})

			It("should return the header at the given block hash", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(header.Hash()).To(Equal(hash))
			})
		})
	})
})
