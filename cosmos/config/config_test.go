// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package config_test

import (
	sgconfig "github.com/berachain/polaris/cosmos/config"
	"github.com/berachain/polaris/eth/accounts"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	sdk "github.com/cosmos/cosmos-sdk/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Configuration", func() {
	It("should set CoinType", func() {
		config := sdk.GetConfig()

		Expect(int(config.GetCoinType())).To(Equal(sdk.CoinType))
		Expect(config.GetFullBIP44Path()).To(Equal(sdk.FullFundraiserPath))

		sgconfig.SetupCosmosConfig()

		Expect(int(config.GetCoinType())).To(Equal(int(accounts.Bip44CoinType)))
		Expect(config.GetCoinType()).To(Equal(sdk.GetConfig().GetCoinType()))
		Expect(config.GetFullBIP44Path()).To(Equal(sdk.GetConfig().GetFullBIP44Path()))
	})

	It("should generate HD path", func() {
		params := *hd.NewFundraiserParams(0, accounts.Bip44CoinType, 0)
		hdPath := params.String()

		Expect(hdPath).To(Equal("m/44'/60'/0'/0/0"))
		Expect(hdPath).To(Equal(accounts.BIP44HDPath))
	})
})
