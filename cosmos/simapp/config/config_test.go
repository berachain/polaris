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
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	sdk "github.com/cosmos/cosmos-sdk/types"

	sgconfig "pkg.berachain.dev/polaris/cosmos/simapp/config"
	"pkg.berachain.dev/polaris/eth/accounts"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/simapp/config")
}

var _ = Describe("Configuration", func() {
	It("should set Bech32 prefixes", func() {
		config := sdk.GetConfig()

		Expect(config.GetBech32AccountAddrPrefix()).To(Equal(sdk.Bech32PrefixAccAddr))
		Expect(config.GetBech32AccountPubPrefix()).To(Equal(sdk.Bech32PrefixAccPub))
		Expect(config.GetBech32ValidatorAddrPrefix()).To(Equal(sdk.Bech32PrefixValAddr))
		Expect(config.GetBech32ValidatorPubPrefix()).To(Equal(sdk.Bech32PrefixValPub))
		Expect(config.GetBech32ConsensusAddrPrefix()).To(Equal(sdk.Bech32PrefixConsAddr))
		Expect(config.GetBech32ConsensusPubPrefix()).To(Equal(sdk.Bech32PrefixConsPub))

		sgconfig.SetBech32Prefixes(config)

		Expect(config.GetBech32AccountAddrPrefix()).To(Equal(sgconfig.Bech32PrefixAccAddr))
		Expect(config.GetBech32AccountPubPrefix()).To(Equal(sgconfig.Bech32PrefixAccPub))
		Expect(config.GetBech32ValidatorAddrPrefix()).To(Equal(sgconfig.Bech32PrefixValAddr))
		Expect(config.GetBech32ValidatorPubPrefix()).To(Equal(sgconfig.Bech32PrefixValPub))
		Expect(config.GetBech32ConsensusAddrPrefix()).To(Equal(sgconfig.Bech32PrefixConsAddr))
		Expect(config.GetBech32ConsensusPubPrefix()).To(Equal(sgconfig.Bech32PrefixConsPub))

		Expect(config.GetBech32AccountAddrPrefix()).To(Equal(sdk.GetConfig().GetBech32AccountAddrPrefix()))
		Expect(config.GetBech32AccountPubPrefix()).To(Equal(sdk.GetConfig().GetBech32AccountPubPrefix()))
		Expect(config.GetBech32ValidatorAddrPrefix()).To(Equal(sdk.GetConfig().GetBech32ValidatorAddrPrefix()))
		Expect(config.GetBech32ValidatorPubPrefix()).To(Equal(sdk.GetConfig().GetBech32ValidatorPubPrefix()))
		Expect(config.GetBech32ConsensusAddrPrefix()).To(Equal(sdk.GetConfig().GetBech32ConsensusAddrPrefix()))
		Expect(config.GetBech32ConsensusPubPrefix()).To(Equal(sdk.GetConfig().GetBech32ConsensusPubPrefix()))
	})

	It("should set CoinType", func() {
		config := sdk.GetConfig()

		Expect(int(config.GetCoinType())).To(Equal(sdk.CoinType))
		Expect(config.GetFullBIP44Path()).To(Equal(sdk.FullFundraiserPath))

		sgconfig.SetBip44CoinType(config)

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

var _ = Describe("RegisterDenoms", func() {
	It("should register the base and display denominations", func() {
		sgconfig.RegisterDenoms()

		// Check if the base denomination was registered correctly
		baseDenom, err := sdk.GetBaseDenom()
		Expect(baseDenom).To(Equal("abera"))
		Expect(err).ToNot(HaveOccurred())

		denomUnit, found := sdk.GetDenomUnit(baseDenom)
		Expect(denomUnit).To(Equal(sdk.NewDecWithPrec(1, 18)))
		Expect(found).To(BeTrue())
	})
})
