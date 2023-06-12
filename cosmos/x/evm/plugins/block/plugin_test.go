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
	"pkg.berachain.dev/polaris/eth/common"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Block Plugin", func() {
	var (
		p   *plugin
		ctx sdk.Context

		header *coretypes.Header
		err    error
	)

	BeforeEach(func() {
		ctx = testutil.NewContext()
		storekey := storetypes.NewKVStoreKey("evm")
		p = &plugin{
			ctx:      ctx,
			storekey: storekey,
		}
	})

	Context("GetHeaderByNumber", func() {
		var (
			number uint64
		)

		JustBeforeEach(func() {
			header, err = p.GetHeaderByNumber(number)
		})

		It("should return the header at the given block number", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(header.Number).To(Equal(number))
		})

		It("should return the header at the genesis block number 0", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(header.Number).To(Equal(number))
		})
	})
	Context("GetHeaderByHash", func() {
		var (
			hash common.Hash
		)

		JustBeforeEach(func() {
			header, err = p.GetHeaderByHash(hash)
		})

		It("should return the header at the given block hash", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(header.Hash()).To(Equal(hash))
		})

		It("should return the header at the genesis block hash", func() {
			Expect(err).ToNot(HaveOccurred())
			Expect(header.Hash()).To(Equal(hash))
		})
	})
})
