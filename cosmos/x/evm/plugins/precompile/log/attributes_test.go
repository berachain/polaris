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

package log

import (
	"math/big"
	"strconv"

	sdkmath "cosmossdk.io/math"

	libgenerated "github.com/berachain/polaris/contracts/bindings/cosmos/lib"
	libutils "github.com/berachain/polaris/lib/utils"

	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Attributes", func() {
	Describe("Test Default Attribute Value Decoder Functions", func() {
		It("should correctly convert sdk coin strings to evm coins", func() {
			// 1 coin
			denom10 := sdk.NewCoin("denom", sdkmath.NewInt(10))
			coins, err := ConvertSdkCoins(denom10.String())
			Expect(err).ToNot(HaveOccurred())
			expectedEvmCoins := []libgenerated.CosmosCoin{
				{
					Amount: big.NewInt(10),
					Denom:  "denom",
				},
			}
			Expect(coins).To(Equal(expectedEvmCoins))

			// many coins
			manyCoins := sdk.NewCoins(
				sdk.NewCoin("abgt", sdkmath.NewInt(1)),
				sdk.NewCoin("abera", sdkmath.NewInt(2)),
				sdk.NewCoin("ahoney", sdkmath.NewInt(3)),
			)
			coins, err = ConvertSdkCoins(manyCoins.String())
			Expect(err).ToNot(HaveOccurred())
			Expect(coins).To(Equal(
				[]libgenerated.CosmosCoin{
					{
						Amount: big.NewInt(2),
						Denom:  "abera",
					},
					{
						Amount: big.NewInt(1),
						Denom:  "abgt",
					},
					{
						Amount: big.NewInt(3),
						Denom:  "ahoney",
					},
				},
			))

			// empty string
			coins, err = ConvertSdkCoins(sdk.NewCoins().String())
			Expect(err).ToNot(HaveOccurred())
			Expect(coins).To(Equal([]libgenerated.CosmosCoin{}))

			// 0 amount coins
			coins, err = ConvertSdkCoins(sdk.NewCoins(sdk.NewCoin("denom", sdkmath.NewInt(0))).String())
			Expect(err).ToNot(HaveOccurred())
			Expect(coins).To(Equal([]libgenerated.CosmosCoin{}))
		})

		It("should correctly convert creation height to int64", func() {
			creationHeightStr := strconv.FormatInt(55, 10)
			gethValue, err := ConvertInt64(creationHeightStr)
			Expect(err).ToNot(HaveOccurred())
			int64Val := libutils.MustGetAs[int64](gethValue)
			Expect(int64Val).To(Equal(int64(55)))
		})

		It("should correctly convert string to uint64", func() {
			numStr := strconv.FormatUint(1, 10)
			gethValue, err := ConvertUint64(numStr)
			Expect(err).ToNot(HaveOccurred())
			uint64Val := libutils.MustGetAs[uint64](gethValue)
			Expect(uint64Val).To(Equal(uint64(1)))
		})
	})

	Describe("Test Search Attributes for Argument", func() {
		var attributes = []abci.EventAttribute{
			{Key: "k0"},
			{Key: "k1"},
			{Key: "k2"},
			{Key: "k3"},
			{Key: "k4"},
		}

		It("should return the correct index if it contains the argument name", func() {
			Expect(searchAttributesForArg(&attributes, "k0")).To(Equal(0))
			Expect(searchAttributesForArg(&attributes, "k3")).To(Equal(3))
			Expect(searchAttributesForArg(&attributes, "k4")).To(Equal(4))
		})

		It("should return -1 if it does not contain the argument name", func() {
			Expect(searchAttributesForArg(&attributes, "")).To(Equal(-1))
			Expect(searchAttributesForArg(&attributes, "k6")).To(Equal(-1))
		})
	})
})
