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

package auth_test

import (
	"context"
	"math/big"
	"testing"

	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/auth"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/precompile/auth"
	"pkg.berachain.dev/polaris/cosmos/precompile/auth/mock"
	"pkg.berachain.dev/polaris/cosmos/precompile/testutil"
	testutils "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/eth/accounts/abi"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/lib/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAddressPrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/precompile/auth")
}

var _ = Describe("Address Precompile", func() {
	var contract *auth.Contract
	var sf *ethprecompile.StatefulFactory
	BeforeEach(func() {
		_, ak, _, _ := testutils.SetupMinimalKeepers()
		k := authzkeeper.NewKeeper(
			runtime.NewKVStoreService(storetypes.NewKVStoreKey(authtypes.StoreKey)),
			testutils.GetEncodingConfig().Codec,
			MsgRouterMockWithSend(),
			ak,
		)
		contract = utils.MustGetAs[*auth.Contract](
			auth.NewPrecompileContract(
				authkeeper.NewQueryServer(ak), k, k,
			),
		)
		sf = ethprecompile.NewStatefulFactory()
	})

	It("should have static registry key", func() {
		Expect(contract.RegistryKey()).To(Equal(
			cosmlib.AccAddressToEthAddress(authtypes.NewModuleAddress(authtypes.ModuleName))),
		)
	})

	It("should have correct ABI methods", func() {
		var cAbi abi.ABI
		err := cAbi.UnmarshalJSON([]byte(generated.AuthModuleMetaData.ABI))
		Expect(err).ToNot(HaveOccurred())
		Expect(contract.ABIMethods()).To(Equal(cAbi.Methods))
	})

	It("should match the precompile methods", func() {
		_, err := sf.Build(contract, nil)
		Expect(err).ToNot(HaveOccurred())
	})

	It("custom value decoder should be no-op", func() {
		Expect(contract.CustomValueDecoders()).To(BeNil())
	})

	When("When Calling ConvertHexToBech32", func() {
		// should probably put something here. utACK
	})

	When("SendGrant", func() {
		var (
			evm              *mock.PrecompileEVMMock
			granter, grantee common.Address
			limit            sdk.Coins
			ctx              context.Context
		)

		BeforeEach(func() {
			// Genereate an evm where the block time is 100.
			sdkCtx, _, _, _ := testutils.SetupMinimalKeepers()
			evm = mock.NewPrecompileEVMMock()
			evm.GetContextFunc = func() *vm.BlockContext {
				blockCtx := vm.BlockContext{}
				blockCtx.Time = 100
				return &blockCtx
			}

			// Generate a granter and grantee address.
			granterAcc := sdk.AccAddress([]byte("granter"))
			granteeAcc := sdk.AccAddress([]byte("grantee"))
			granter = cosmlib.AccAddressToEthAddress(granterAcc)
			grantee = cosmlib.AccAddressToEthAddress(granteeAcc)

			ctx = vm.NewPolarContext(
				sdkCtx,
				evm,
				granter,
				new(big.Int),
			)

			// Generate a limit.
			limit = sdk.NewCoins(sdk.NewInt64Coin("test", 100))

			// expiredTime = big.NewInt(200)
		})

		It("should error if the expiration is before the current block time", func() {

			_, err := contract.SetSendAllowance(
				ctx,
				grantee,
				testutil.SdkCoinsToEvmCoins(limit),
				big.NewInt(1),
			)
			Expect(err).To(HaveOccurred())
		})

		It("should succeed with expiration", func() {
			_, err := contract.SetSendAllowance(
				ctx,
				grantee,
				testutil.SdkCoinsToEvmCoins(limit),
				big.NewInt(110),
			)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should succeed without expiration", func() {
			_, err := contract.SetSendAllowance(
				ctx,
				grantee,
				testutil.SdkCoinsToEvmCoins(limit),
				new(big.Int),
			)
			Expect(err).ToNot(HaveOccurred())
		})

		When("Get Send Allowance: ", func() {
			BeforeEach(func() {
				// Set up a spend limit grant.
				_, err := contract.SetSendAllowance(
					ctx,
					grantee,
					testutil.SdkCoinsToEvmCoins(limit),
					new(big.Int),
				)
				Expect(err).ToNot(HaveOccurred())
			})

			It("should get the spend allowance", func() {
				res, err := contract.GetSendAllowance(
					ctx,
					granter,
					grantee,
					"test",
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(res).To(Equal(big.NewInt(100)))
			})
		})
	})

})

func MsgRouterMockWithSend() *mock.MessageRouterMock {
	router := mock.NewMsgRouterMock()
	router.HandlerByTypeURLFunc = func(typeURL string) func(ctx sdk.Context, req sdk.Msg) (*sdk.Result, error) {
		return func(ctx sdk.Context, req sdk.Msg) (*sdk.Result, error) {
			return &sdk.Result{}, nil
		}
	}

	return router
}
