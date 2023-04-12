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

package precompile

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/testing/network"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ERC20", func() {
	Describe("calling the erc20 precompile directly", func() {
		When("calling read-only methods", func() {
			It("should handle empty inputs", func() {
				// nonexistent address
				denom, err := erc20Precompile.CoinDenomForERC20Address(nil, common.Address{})
				Expect(err).ToNot(HaveOccurred())
				Expect(denom).To(Equal(""))

				// invalid address
				_, err = erc20Precompile.CoinDenomForERC20Address0(nil, "")
				Expect(err).To(HaveOccurred())

				// nonexistent denom
				token, err := erc20Precompile.Erc20AddressForCoinDenom(nil, "")
				Expect(err).ToNot(HaveOccurred())
				Expect(token).To(Equal(common.Address{}))
			})

			It("should handle non-empty inputs", func() {
				token, err := erc20Precompile.Erc20AddressForCoinDenom(nil, "abera")
				Expect(err).ToNot(HaveOccurred())
				Expect(token).To(Equal(common.Address{}))

				tokenAddr := common.BytesToAddress([]byte("abera"))
				tokenBech32 := cosmlib.AddressToAccAddress(tokenAddr).String()

				denom, err := erc20Precompile.CoinDenomForERC20Address(nil, tokenAddr)
				Expect(err).ToNot(HaveOccurred())
				Expect(denom).To(Equal(""))

				denom, err = erc20Precompile.CoinDenomForERC20Address0(nil, tokenBech32)
				Expect(err).ToNot(HaveOccurred())
				Expect(denom).To(Equal(""))
			})
		})

		When("calling write methods", func() {
			It("should error on non-existent inputs", func() {
				// nonexistent address
				_, err := erc20Precompile.ConvertERC20ToCoin0(
					tf.GenerateTransactOpts(""),
					common.BytesToAddress([]byte("sUSDC")),
					network.TestAddress,
					big.NewInt(123456789),
				)
				Expect(err).To(HaveOccurred())

				// nonexistent address
				_, err = erc20Precompile.ConvertERC20ToCoin(
					tf.GenerateTransactOpts(""),
					common.BytesToAddress([]byte("sUSDC")),
					cosmlib.AddressToAccAddress(network.TestAddress).String(),
					big.NewInt(123456789),
				)
				Expect(err).To(HaveOccurred())

				// nonexistent denom
				_, err = erc20Precompile.ConvertCoinToERC20(
					tf.GenerateTransactOpts(""),
					"bOSMO",
					network.TestAddress,
					big.NewInt(123456789),
				)
				Expect(err).To(HaveOccurred())

				// nonexistent denom
				_, err = erc20Precompile.ConvertCoinToERC200(
					tf.GenerateTransactOpts(""),
					"bOSMO",
					cosmlib.AddressToAccAddress(network.TestAddress).String(),
					big.NewInt(123456789),
				)
				Expect(err).To(HaveOccurred())
			})

			It("should handle non-empty inputs", func() {
				// denom already exists
				erc20Precompile.ConvertCoinToERC200(
					tf.GenerateTransactOpts(""),
					"bATOM",
					cosmlib.AddressToAccAddress(network.TestAddress).String(),
					big.NewInt(123456789),
				)
				// Expect(err).ToNot(HaveOccurred())
			})
		})
	})

	Describe("calling the erc20 precompile via the another contract", func() {

	})
})
