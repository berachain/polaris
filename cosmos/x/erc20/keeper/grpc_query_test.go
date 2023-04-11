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

package keeper_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/cosmos/x/erc20/keeper"
	"pkg.berachain.dev/polaris/cosmos/x/erc20/types"
	"pkg.berachain.dev/polaris/eth/common"
)

var _ = Describe("GRPC Query Server", func() {
	var k *keeper.Keeper
	var qs types.QueryServiceServer
	var ctx sdk.Context
	var bk keeper.BankKeeper

	BeforeEach(func() {
		ctx, _, bk, _ = utils.SetupMinimalKeepers()
		k = keeper.NewKeeper(
			storetypes.NewKVStoreKey("erc20"), bk, authtypes.NewModuleAddress(govtypes.ModuleName),
		)
		qs = k
	})

	It("should correctly empty inputs", func() {
		_, err := qs.CoinDenomForERC20Address(ctx, &types.CoinDenomForERC20AddressRequest{})
		Expect(err).To(HaveOccurred())

		resp, err := qs.ERC20AddressForCoinDenom(ctx, &types.ERC20AddressForCoinDenomRequest{})
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.Token).To(BeEmpty())
	})

	It("should correctly handle erc20 --> coin", func() {
		tokenAddr := common.BytesToAddress([]byte("USDC"))

		// check the denom doesn't exist
		resp, err := qs.CoinDenomForERC20Address(ctx, &types.CoinDenomForERC20AddressRequest{
			Token: cosmlib.AddressToAccAddress(tokenAddr).String(),
		})
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.Denom).To(BeEmpty())

		// register the new denom for token
		k.RegisterERC20CoinPair(ctx, tokenAddr)

		// check the denom exists
		resp, err = qs.CoinDenomForERC20Address(ctx, &types.CoinDenomForERC20AddressRequest{
			Token: cosmlib.AddressToAccAddress(tokenAddr).String(),
		})
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.Denom).To(Equal(types.NewPolarisDenomForAddress(tokenAddr)))
	})

	It("should correctly handle coin --> erc20", func() {
		denom := "osmo"
		tokenAddr := common.BytesToAddress([]byte(denom))

		// check the denom doesn't exist
		resp, err := qs.ERC20AddressForCoinDenom(ctx, &types.ERC20AddressForCoinDenomRequest{
			Denom: denom,
		})
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.Token).To(BeEmpty())

		// register the new denom for token
		k.RegisterCoinERC20Pair(ctx, denom, tokenAddr)

		// check the denom exists
		resp, err = qs.ERC20AddressForCoinDenom(ctx, &types.ERC20AddressForCoinDenomRequest{
			Denom: denom,
		})
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.Token).To(Equal(cosmlib.AddressToAccAddress(tokenAddr).String()))
	})
})
