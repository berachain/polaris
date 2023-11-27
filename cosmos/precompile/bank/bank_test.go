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

package bank_test

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"

	"github.com/berachain/polaris/cosmos/precompile"
	"github.com/berachain/polaris/cosmos/precompile/bank"
	testutils "github.com/berachain/polaris/cosmos/testutil"
	pclog "github.com/berachain/polaris/cosmos/x/evm/plugins/precompile/log"
	evmtypes "github.com/berachain/polaris/cosmos/x/evm/types"
	ethprecompile "github.com/berachain/polaris/eth/core/precompile"
	"github.com/berachain/polaris/eth/core/vm"
	"github.com/berachain/polaris/lib/utils"

	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBankPrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/precompile/bank")
}

var _ = Describe("Bank Precompile Test", func() {
	var (
		contract *bank.Contract
		addr     sdk.AccAddress
		factory  *pclog.Factory
		ak       authkeeper.AccountKeeperI
		bk       bankkeeper.BaseKeeper
		ctx      context.Context
	)

	BeforeEach(func() {
		ctx, ak, bk, _ = testutils.SetupMinimalKeepers(log.NewTestLogger(GinkgoT()))

		contract = utils.MustGetAs[*bank.Contract](bank.NewPrecompileContract(
			ak, bankkeeper.NewMsgServerImpl(bk), bk),
		)
		addr = sdk.AccAddress([]byte("bank"))

		// Register the events.
		factory = pclog.NewFactory([]ethprecompile.Registrable{contract})
	})

	It("should register the send event", func() {
		event := sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeySender, addr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, sdk.NewCoin("stg", sdkmath.NewInt(100)).String()),
		)
		log, err := factory.Build(&event)
		Expect(err).ToNot(HaveOccurred())
		Expect(log.Address).To(Equal(contract.RegistryKey()))
	})

	It("should register the transfer event", func() {
		event := sdk.NewEvent(
			banktypes.EventTypeTransfer,
			sdk.NewAttribute(banktypes.AttributeKeyRecipient, addr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, sdk.NewCoin("stg", sdkmath.NewInt(100)).String()),
		)
		log, err := factory.Build(&event)
		Expect(err).ToNot(HaveOccurred())
		Expect(log.Address).To(Equal(contract.RegistryKey()))
	})

	It("should register the coin spent event", func() {
		event := sdk.NewEvent(
			banktypes.EventTypeCoinSpent,
			sdk.NewAttribute(banktypes.AttributeKeySpender, addr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, sdk.NewCoin("stg", sdkmath.NewInt(100)).String()),
		)
		log, err := factory.Build(&event)
		Expect(err).ToNot(HaveOccurred())
		Expect(log.Address).To(Equal(contract.RegistryKey()))
	})

	It("should register the coin received event", func() {
		event := sdk.NewEvent(
			banktypes.EventTypeCoinReceived,
			sdk.NewAttribute(banktypes.AttributeKeyReceiver, addr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, sdk.NewCoin("stg", sdkmath.NewInt(100)).String()),
		)
		log, err := factory.Build(&event)
		Expect(err).ToNot(HaveOccurred())
		Expect(log.Address).To(Equal(contract.RegistryKey()))
	})

	It("should register the burn event", func() {
		event := sdk.NewEvent(
			banktypes.EventTypeCoinBurn,
			sdk.NewAttribute(banktypes.AttributeKeyBurner, addr.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, sdk.NewCoin("stg", sdkmath.NewInt(100)).String()),
		)
		log, err := factory.Build(&event)
		Expect(err).ToNot(HaveOccurred())
		Expect(log.Address).To(Equal(contract.RegistryKey()))
	})

	When("Calling Precompile Methods", func() {
		var (
			acc    sdk.AccAddress
			denom  = "abera"
			denom2 = "atoken"
		)

		BeforeEach(func() {
			ctx = vm.NewPolarContext(ctx, nil, common.BytesToAddress(acc), big.NewInt(0))
		})

		When("GetBalance", func() {
			It("should fail if input denom is not a valid denom", func() {
				res, err := contract.GetBalance(
					ctx,
					common.BytesToAddress(acc),
					"_invalid_denom",
				)
				// reDnmString = `[a-zA-Z][a-zA-Z0-9/:._-]{2,127}`
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			})

			It("should succeed", func() {
				balanceAmount, ok := new(big.Int).SetString("22000000000000000000", 10)
				Expect(ok).To(BeTrue())

				acc = simtestutil.CreateRandomAccounts(1)[0]

				err := FundAccount(
					sdk.UnwrapSDKContext(vm.UnwrapPolarContext(ctx).Context()),
					bk,
					acc,
					sdk.NewCoins(
						sdk.NewCoin(
							denom,
							sdkmath.NewIntFromBigInt(balanceAmount),
						),
					),
				)
				Expect(err).ToNot(HaveOccurred())

				res, err := contract.GetBalance(
					ctx,
					common.BytesToAddress(acc),
					denom,
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(res).To(Equal(balanceAmount))
			})
		})

		When("GetAllBalance", func() {
			It("should succeed", func() {
				numOfDenoms := 3
				acc = simtestutil.CreateRandomAccounts(1)[0]
				for i := 0; i < numOfDenoms; i++ {
					balanceAmountStr := fmt.Sprintf("%d000000000000000000", i+1)
					balanceAmount, ok := new(big.Int).SetString(balanceAmountStr, 10)
					Expect(ok).To(BeTrue())

					err := FundAccount(
						sdk.UnwrapSDKContext(vm.UnwrapPolarContext(ctx).Context()),
						bk,
						acc,
						sdk.NewCoins(
							sdk.NewCoin(
								fmt.Sprintf("denom_%d", i+1),
								sdkmath.NewIntFromBigInt(balanceAmount),
							),
						),
					)
					Expect(err).ToNot(HaveOccurred())
				}

				coins, err := contract.GetAllBalances(
					ctx,
					common.BytesToAddress(acc),
				)
				Expect(err).ToNot(HaveOccurred())

				for i, coin := range coins {
					balanceAmountStr := fmt.Sprintf("%d000000000000000000", i+1)
					balanceAmount, ok2 := new(big.Int).SetString(balanceAmountStr, 10)
					Expect(ok2).To(BeTrue())

					Expect(coin.Denom).To(Equal(fmt.Sprintf("denom_%d", i+1)))
					Expect(coin.Amount).To(Equal(balanceAmount))
				}
			})
		})

		When("GetSpendableBalanceByDenom", func() {

			It("should fail if input denom is not a valid denom", func() {
				res, err := contract.GetSpendableBalance(
					ctx,
					common.BytesToAddress(acc),
					"_invalid_denom",
				)
				// reDnmString = `[a-zA-Z][a-zA-Z0-9/:._-]{2,127}`
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			})

			It("should succeed", func() {
				balanceAmount, ok := new(big.Int).SetString("22000000000000000000", 10)
				Expect(ok).To(BeTrue())
				// todo: use vesting accounts, lock some tokens
				acc = simtestutil.CreateRandomAccounts(1)[0]

				err := FundAccount(
					sdk.UnwrapSDKContext(vm.UnwrapPolarContext(ctx).Context()),
					bk,
					acc,
					sdk.NewCoins(
						sdk.NewCoin(
							denom,
							sdkmath.NewIntFromBigInt(balanceAmount),
						),
					),
				)
				Expect(err).ToNot(HaveOccurred())

				res, err := contract.GetSpendableBalance(
					ctx,
					common.BytesToAddress(acc),
					denom,
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(res).To(Equal(balanceAmount))
			})
		})

		When("GetSpendableBalances", func() {

			It("should succeed", func() {
				numOfDenoms := 3
				// todo: use vesting accounts, lock some tokens
				acc = simtestutil.CreateRandomAccounts(1)[0]
				for i := 0; i < numOfDenoms; i++ {
					balanceAmountStr := fmt.Sprintf("%d000000000000000000", i+1)
					balanceAmount, ok := new(big.Int).SetString(balanceAmountStr, 10)
					Expect(ok).To(BeTrue())

					err := FundAccount(
						sdk.UnwrapSDKContext(vm.UnwrapPolarContext(ctx).Context()),
						bk,
						acc,
						sdk.NewCoins(
							sdk.NewCoin(
								fmt.Sprintf("denom_%d", i+1),
								sdkmath.NewIntFromBigInt(balanceAmount),
							),
						),
					)
					Expect(err).ToNot(HaveOccurred())
				}

				coins, err := contract.GetAllSpendableBalances(
					ctx,
					common.BytesToAddress(acc),
				)
				Expect(err).ToNot(HaveOccurred())

				for i, coin := range coins {
					balanceAmountStr := fmt.Sprintf("%d000000000000000000", i+1)
					balanceAmount, ok2 := new(big.Int).SetString(balanceAmountStr, 10)
					Expect(ok2).To(BeTrue())

					Expect(coin.Denom).To(Equal(fmt.Sprintf("denom_%d", i+1)))
					Expect(coin.Amount).To(Equal(balanceAmount))
				}
			})
		})

		When("GetSupplyOf", func() {

			It("should fail if input denom is not a valid Denom", func() {
				res, err := contract.GetSupply(
					ctx,
					"_invalid_denom",
				)
				// fmt.Errorf("invalid denom: %s", denom)
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			})

			It("should succeed", func() {
				balanceAmount, ok := new(big.Int).SetString("22000000000000000000", 10)
				Expect(ok).To(BeTrue())
				balanceAmount3, ok := new(big.Int).SetString("66000000000000000000", 10)
				Expect(ok).To(BeTrue())

				accs := simtestutil.CreateRandomAccounts(3)

				for i := 0; i < 3; i++ {
					err := FundAccount(
						sdk.UnwrapSDKContext(vm.UnwrapPolarContext(ctx).Context()),
						bk,
						accs[i],
						sdk.NewCoins(
							sdk.NewCoin(
								denom,
								sdkmath.NewIntFromBigInt(balanceAmount),
							),
						),
					)
					Expect(err).ToNot(HaveOccurred())
				}

				res, err := contract.GetSupply(
					ctx,
					denom,
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(res).To(Equal(balanceAmount3))
			})
		})

		When("GetTotalSupply", func() {
			It("should succeed", func() {
				balanceAmount, ok := new(big.Int).SetString("22000000000000000000", 10)
				Expect(ok).To(BeTrue())
				balanceAmount3, ok := new(big.Int).SetString("66000000000000000000", 10)
				Expect(ok).To(BeTrue())

				accs := simtestutil.CreateRandomAccounts(3)
				for i := 0; i < 3; i++ {
					for j := 0; j < 3; j++ {
						err := FundAccount(
							sdk.UnwrapSDKContext(vm.UnwrapPolarContext(ctx).Context()),
							bk,
							accs[i],
							sdk.NewCoins(
								sdk.NewCoin(
									fmt.Sprintf("%s%d", denom, j),
									sdkmath.NewIntFromBigInt(balanceAmount),
								),
							),
						)
						Expect(err).ToNot(HaveOccurred())
					}
				}

				coins, err := contract.GetAllSupply(
					ctx,
				)
				Expect(err).ToNot(HaveOccurred())

				for i := 0; i < 3; i++ {
					Expect(coins[i].Denom).To(Equal(fmt.Sprintf("%s%d", denom, i)))
					Expect(coins[i].Amount).To(Equal(balanceAmount3))
				}

			})
		})

		When("Send", func() {
			It("should succeed", func() {

				balanceAmount, ok := new(big.Int).SetString("220000000000000000", 10)
				Expect(ok).To(BeTrue())

				accs := simtestutil.CreateRandomAccounts(2)
				fromAcc, toAcc := accs[0], accs[1]

				pCtx := vm.NewPolarContext(ctx, nil, common.BytesToAddress(fromAcc), new(big.Int))

				sortedSdkCoins := sdk.NewCoins(
					sdk.NewCoin(
						denom,
						sdkmath.NewIntFromBigInt(balanceAmount),
					),
					sdk.NewCoin(
						denom2,
						sdkmath.NewIntFromBigInt(balanceAmount),
					),
				)

				err := FundAccount(
					sdk.UnwrapSDKContext(ctx),
					bk,
					fromAcc,
					sortedSdkCoins,
				)
				Expect(err).ToNot(HaveOccurred())

				bk.SetSendEnabled(ctx, denom, true)
				bk.SetSendEnabled(ctx, denom2, true)

				_, err = contract.Send(
					pCtx,
					common.BytesToAddress(toAcc),
					sdkCoinsToEvmCoins(sortedSdkCoins),
				)
				Expect(err).ToNot(HaveOccurred())

				balances, err := bk.AllBalances(ctx, banktypes.NewQueryAllBalancesRequest(toAcc, nil, false))
				Expect(err).ToNot(HaveOccurred())

				Expect(balances.Balances).To(Equal(sortedSdkCoins))
			})

			It("should error when sending 0 coins", func() {
				balanceAmount, ok := new(big.Int).SetString("22000000000000000000", 10)
				Expect(ok).To(BeTrue())
				accs := simtestutil.CreateRandomAccounts(2)
				fromAcc, toAcc := accs[0], accs[1]
				coinsToMint := sdk.NewCoins(
					sdk.NewCoin(denom, sdkmath.NewIntFromBigInt(balanceAmount)),
				)
				coinsToSend := sdk.NewCoins(
					sdk.NewCoin(denom, sdkmath.NewIntFromBigInt(big.NewInt(0))),
				)
				err := FundAccount(
					sdk.UnwrapSDKContext(vm.UnwrapPolarContext(ctx).Context()),
					bk,
					fromAcc,
					coinsToMint,
				)
				Expect(err).ToNot(HaveOccurred())
				bk.SetSendEnabled(ctx, denom, true)
				_, err = contract.Send(
					ctx,
					common.BytesToAddress(toAcc),
					sdkCoinsToEvmCoins(coinsToSend),
				)
				Expect(err).To(MatchError(precompile.ErrInvalidCoin))
			})
		})
	})
})

func FundAccount(
	ctx sdk.Context,
	bk bankkeeper.BaseKeeper,
	account sdk.AccAddress,
	coins sdk.Coins,
) error {
	if err := bk.MintCoins(ctx, evmtypes.ModuleName, coins); err != nil {
		return err
	}
	return bk.SendCoinsFromModuleToAccount(ctx, evmtypes.ModuleName, account, coins)
}
func sdkCoinsToEvmCoins(sdkCoins sdk.Coins) []struct {
	Amount *big.Int `json:"amount"`
	Denom  string   `json:"denom"`
} {
	evmCoins := make([]struct {
		Amount *big.Int `json:"amount"`
		Denom  string   `json:"denom"`
	}, len(sdkCoins))
	for i, coin := range sdkCoins {
		evmCoins[i] = struct {
			Amount *big.Int `json:"amount"`
			Denom  string   `json:"denom"`
		}{
			Amount: coin.Amount.BigInt(),
			Denom:  coin.Denom,
		}
	}
	return evmCoins
}
