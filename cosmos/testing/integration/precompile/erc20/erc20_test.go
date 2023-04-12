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

package erc20

import (
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	bindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/testing/integration"
	"pkg.berachain.dev/polaris/cosmos/testing/network"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "pkg.berachain.dev/polaris/cosmos/testing/integration/utils"
)

func TestERC20Precompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/testing/integration/precompile/erc20")
}

var (
	etf             *integration.TestFixture
	erc20Precompile *bindings.ERC20Module
)

var _ = SynchronizedBeforeSuite(func() []byte {
	// Setup the network and clients here.
	etf = integration.NewTestFixture(GinkgoT())
	erc20Precompile, _ = bindings.NewERC20Module(
		// cosmlib.AccAddressToEthAddress(
		// 	authtypes.NewModuleAddress(erc20types.ModuleName),
		// ),
		common.HexToAddress("0x696969"),
		etf.EthClient,
	)
	return nil
}, func(data []byte) {})

var _ = SynchronizedAfterSuite(func() {
	// Local AfterSuite actions.
}, func() {
	// Global AfterSuite actions.
	os.RemoveAll("data")
})

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
			It("should error on non-existent coin denoms", func() {
				_, err := erc20Precompile.ConvertCoinToERC20(
					etf.GenerateTransactOpts(""),
					"bOSMO",
					network.TestAddress,
					big.NewInt(123456789),
				)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("insufficient funds"))

				_, err = erc20Precompile.ConvertCoinToERC200(
					etf.GenerateTransactOpts(""),
					"bOSMO",
					cosmlib.AddressToAccAddress(network.TestAddress).String(),
					big.NewInt(123456789),
				)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("insufficient funds"))
			})

			It("should handle non-empty inputs", func() {
				// denom already exists
				tx, err := erc20Precompile.ConvertCoinToERC20(
					etf.GenerateTransactOpts(""),
					"bATOM",
					network.TestAddress,
					big.NewInt(123456789),
				)
				Expect(err).ToNot(HaveOccurred())
				ExpectMined(etf.EthClient, tx)
				ExpectSuccessReceipt(etf.EthClient, tx)

				// // denom already exists
				// tx, err = erc20Precompile.ConvertCoinToERC200(
				// 	etf.GenerateTransactOpts(""),
				// 	"bATOM",
				// 	cosmlib.AddressToAccAddress(network.TestAddress).String(),
				// 	big.NewInt(123456789),
				// )
				// Expect(err).ToNot(HaveOccurred())
				// ExpectMined(etf.EthClient, tx)
				// ExpectSuccessReceipt(etf.EthClient, tx)

				// // nonexistent address
				// tx, err := erc20Precompile.ConvertERC20ToCoin0(
				// 	etf.GenerateTransactOpts(""),
				// 	common.BytesToAddress([]byte("sUSDC")),
				// 	network.TestAddress,
				// 	big.NewInt(123456789),
				// )

				// // nonexistent address
				// _, err = erc20Precompile.ConvertERC20ToCoin(
				// 	etf.GenerateTransactOpts(""),
				// 	common.BytesToAddress([]byte("sUSDC")),
				// 	cosmlib.AddressToAccAddress(network.TestAddress).String(),
				// 	big.NewInt(123456789),
				// )
				// Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("calling the erc20 precompile via the another contract", func() {

	})
})
