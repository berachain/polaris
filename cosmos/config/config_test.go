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
