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

package configuration

import (
	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/params"
	enclib "pkg.berachain.dev/polaris/lib/encoding"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Plugin", func() {
	var (
		p   *plugin
		ctx sdk.Context
	)

	BeforeEach(func() {
		ctx = testutil.NewContext()
		storeKey := storetypes.NewKVStoreKey("evm")
		p = &plugin{
			storeKey: storeKey,
		}
	})

	Describe("ChainConfig", func() {
		Context("when the params store is empty", func() {
			It("should return nil", func() {
				config := p.ChainConfig(ctx)
				Expect(config).To(BeNil())
			})
		})

		Context("when the params store contains valid params", func() {
			It("should return the chain config", func() {
				storedParams := types.Params{
					EvmDenom:    "eth",
					ExtraEIPs:   []int64{123},
					ChainConfig: string(enclib.MustMarshalJSON(params.DefaultChainConfig)),
				}
				bz, err := storedParams.Marshal()
				Expect(err).ToNot(HaveOccurred())
				ctx.KVStore(p.storeKey).Set([]byte{types.ParamsKey}, bz)

				config := p.ChainConfig(ctx)
				Expect(config).To(Equal(params.DefaultChainConfig))
			})
		})

		Context("when the params store contains invalid params", func() {
			It("should panic", func() {
				ctx.KVStore(p.storeKey).Set([]byte{types.ParamsKey}, []byte("invalid params"))
				Expect(func() { p.ChainConfig(ctx) }).To(Panic())
			})
		})
	})

	Describe("ExtraEips", func() {
		It("should return an empty slice", func() {
			eips := p.ExtraEips(ctx)
			Expect(eips).To(BeEmpty())
		})
	})
})
