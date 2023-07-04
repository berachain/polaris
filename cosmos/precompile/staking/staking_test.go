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

package staking

import (
	"math/big"
	"reflect"
	"testing"

	sdkmath "cosmossdk.io/math"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/staking"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/eth/accounts/abi"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/lib/utils"
)

func TestStakingPrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/precompile/staking")
}

func createValAddrs(count int) ([]sdk.AccAddress, []sdk.ValAddress) {
	addrs := simtestutil.CreateIncrementalAccounts(count)
	valAddrs := simtestutil.ConvertAddrsToValAddrs(addrs)

	return addrs, valAddrs
}

func NewValidator(operator sdk.ValAddress, pubKey cryptotypes.PubKey) (stakingtypes.Validator, error) {
	return stakingtypes.NewValidator(operator, pubKey, stakingtypes.Description{})
}

var (
	PKs = simtestutil.CreateTestPubKeys(500)
)

var _ = Describe("Staking", func() {
	var (
		sk stakingkeeper.Keeper
		bk bankkeeper.BaseKeeper

		ctx sdk.Context

		contract *Contract
	)

	BeforeEach(func() {
		ctx, _, bk, sk = testutil.SetupMinimalKeepers()
		skPtr := &sk
		contract = utils.MustGetAs[*Contract](NewPrecompileContract(skPtr))
	})

	When("AbiMethods", func() {
		It("returns the correct methods", func() {
			var cAbi abi.ABI
			err := cAbi.UnmarshalJSON([]byte(generated.StakingModuleMetaData.ABI))
			Expect(err).ToNot(HaveOccurred())
			methods := contract.ABIMethods()
			Expect(methods).To(HaveLen(len(cAbi.Methods)))
		})
	})

	When("PrecompileMethods", func() {
		It("should return the correct methods", func() {
			Expect(ethprecompile.GeneratePrecompileMethods(contract.ABIMethods(), reflect.ValueOf(contract))).To(HaveLen(len(contract.ABIMethods())))
		})
	})

	When("ABIEvents", func() {
		It("should return the correct events", func() {
			var cAbi abi.ABI
			err := cAbi.UnmarshalJSON([]byte(generated.StakingModuleMetaData.ABI))
			Expect(err).ToNot(HaveOccurred())
			events := contract.ABIEvents()
			Expect(events).To(HaveLen(len(cAbi.Events)))
		})
	})

	When("CustomValueDecoders", func() {
		It("should be a no-op", func() {
			Expect(contract.CustomValueDecoders()).To(BeNil())
		})
	})

	When("Calling Precompile Methods", func() {
		var (
			del            sdk.AccAddress
			val            sdk.ValAddress
			validator      stakingtypes.Validator
			otherValidator stakingtypes.Validator
			otherVal       sdk.ValAddress
			caller         common.Address
		)

		BeforeEach(func() {
			delegates, validators := createValAddrs(2)
			del, val, otherVal = delegates[0], validators[0], validators[1]
			caller = cosmlib.AccAddressToEthAddress(del)

			amount, ok := new(big.Int).SetString("22000000000000000000", 10) // 22 tokens.
			Expect(ok).To(BeTrue())
			var err error

			validator, err = NewValidator(val, PKs[0])
			Expect(err).ToNot(HaveOccurred())

			otherValidator, err = NewValidator(otherVal, PKs[1])
			Expect(err).ToNot(HaveOccurred())

			validator, _ = validator.AddTokensFromDel(sdkmath.NewIntFromBigInt(amount))
			otherValidator, _ = otherValidator.AddTokensFromDel(sdkmath.NewIntFromBigInt(amount))

			validator = stakingkeeper.TestingUpdateValidator(&sk, ctx, validator, true)
			otherValidator = stakingkeeper.TestingUpdateValidator(&sk, ctx, otherValidator, true)

			delegation := stakingtypes.NewDelegation(del, val, sdkmath.LegacyNewDec(9))
			sk.SetDelegation(ctx, delegation)

			// Check that the delegation was created.
			res, found := sk.GetDelegation(ctx, del, val)
			Expect(found).To(BeTrue())
			Expect(res).To(Equal(delegation))

			// Set the denom.
			defaultParams := stakingtypes.DefaultParams()
			defaultParams.BondDenom = "stake"
			err = sk.SetParams(ctx, defaultParams)
			Expect(err).ToNot(HaveOccurred())

		})

		When("Delegate", func() {

			It("should succeed", func() {
				amountToDelegate, ok := new(big.Int).SetString("22000000000000000000", 10)
				Expect(ok).To(BeTrue())
				err := FundAccount(
					ctx,
					bk,
					del,
					sdk.NewCoins(
						sdk.NewCoin(
							"stake",
							sdkmath.NewIntFromBigInt(amountToDelegate),
						),
					),
				)
				Expect(err).ToNot(HaveOccurred())

				_, err = contract.Delegate(
					ctx, nil, caller,
					big.NewInt(0),
					true,
					cosmlib.ValAddressToEthAddress(val),
					amountToDelegate,
				)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		When("Delegate0", func() {

			It("should fail if the string is not a valid address", func() {
				res, err := contract.Delegate0(
					ctx, nil, caller,
					big.NewInt(0),
					false,
					"invalid",
					big.NewInt(0),
				)
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			})

			It("should succeed", func() {
				amountToDelegate, ok := new(big.Int).SetString("22000000000000000000", 10)
				Expect(ok).To(BeTrue())
				err := FundAccount(
					ctx,
					bk,
					del,
					sdk.NewCoins(
						sdk.NewCoin(
							"stake",
							sdkmath.NewIntFromBigInt(amountToDelegate),
						),
					),
				)
				Expect(err).ToNot(HaveOccurred())

				_, err = contract.Delegate0(
					ctx,
					nil,
					caller,
					big.NewInt(0),
					false,
					val.String(),
					amountToDelegate,
				)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		When("GetDelegation", func() {

			It("should return the correct delegation", func() {
				res, err := contract.GetDelegation(
					ctx, nil, caller,
					big.NewInt(0),
					true,
					cosmlib.AccAddressToEthAddress(del), cosmlib.ValAddressToEthAddress(val),
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(res[0]).To(Equal(big.NewInt(9))) // should have correct shares
			})
		})

		When("GetDelegation0", func() {

			It("should error if del not bech32", func() {
				res, err := contract.GetDelegation0(
					ctx, nil, caller,
					big.NewInt(0),
					true,
					"0x", val.String(),
				)
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			})

			It("should return an error if the val is not bech32", func() {
				res, err := contract.GetDelegation0(
					ctx, nil, caller,
					big.NewInt(0),
					true,
					del.String(), "0x",
				)
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			})

			It("should return the correct delegation", func() {
				res, err := contract.GetDelegation0(
					ctx, nil, caller,
					big.NewInt(0),
					true,
					del.String(), val.String(),
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(res[0]).To(Equal(big.NewInt(9))) // should have correct shares
			})
		})

		When("Undelegate", func() {

			It("should succeed", func() {
				_, err := contract.Undelegate(
					ctx, nil, caller,
					big.NewInt(0),
					true,
					cosmlib.ValAddressToEthAddress(val),
					big.NewInt(1),
				)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		When("Undelegate0", func() {
			It("should fail if the address is not of type bech32", func() {
				res, err := contract.Undelegate0(
					ctx, nil, caller,
					big.NewInt(0),
					true,
					"0x",
					big.NewInt(0),
				)
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			})

			It("should succeed", func() {
				_, err := contract.Undelegate0(
					ctx, nil, caller,
					big.NewInt(0),
					true,
					val.String(),
					big.NewInt(1),
				)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		When("BeginRedelegations", func() {

			It("should succeed", func() {
				_, err := contract.BeginRedelegate(
					ctx, nil, caller,
					big.NewInt(0),
					false,
					cosmlib.ValAddressToEthAddress(val),
					cosmlib.ValAddressToEthAddress(otherVal),
					big.NewInt(1),
				)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		When("BeginRedelegations0", func() {

			It("should fail if the srcValue is not of type bech32", func() {
				res, err := contract.BeginRedelegate0(
					ctx, nil, caller,
					big.NewInt(0),
					false,
					"0x",
					val.String(),
					big.NewInt(1),
				)
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			})

			It("should fail if the dstValue is not of type bech32", func() {
				res, err := contract.BeginRedelegate0(
					ctx, nil, caller,
					big.NewInt(0),
					false,
					val.String(),
					"0x",
					big.NewInt(1),
				)
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			})

			It("should succeed", func() {
				_, err := contract.BeginRedelegate0(
					ctx, nil, caller,
					big.NewInt(0),
					false,
					val.String(),
					otherVal.String(),
					big.NewInt(1),
				)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		When("CancelUnbondingDelegation", func() {
			It("should succeed", func() {
				creationHeight := ctx.BlockHeight()
				amount, ok := new(big.Int).SetString("1", 10)
				Expect(ok).To(BeTrue())

				// Undelegate.
				_, err := contract.Undelegate(
					ctx, nil, caller,
					big.NewInt(0),
					false,
					cosmlib.ValAddressToEthAddress(val),
					amount,
				)
				Expect(err).ToNot(HaveOccurred())

				_, err = contract.CancelUnbondingDelegation(
					ctx, nil, caller,
					big.NewInt(0),
					false,
					cosmlib.ValAddressToEthAddress(val),
					amount,
					creationHeight,
				)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		When("CancelUnbondingDelegation0", func() {

			It("should fail if the address is not a bech32 address", func() {
				res, err := contract.CancelUnbondingDelegation0(
					ctx, nil, caller,
					big.NewInt(0),
					false,
					"0x",
					big.NewInt(1),
					int64(1),
				)
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			})

			It("should succeed", func() {
				creationHeight := ctx.BlockHeight()
				amount, ok := new(big.Int).SetString("1", 10)
				Expect(ok).To(BeTrue())

				// Undelegate.
				_, err := contract.Undelegate(
					ctx, nil, caller,
					big.NewInt(0),
					false,
					cosmlib.ValAddressToEthAddress(val),
					amount,
				)
				Expect(err).ToNot(HaveOccurred())

				_, err = contract.CancelUnbondingDelegation0(
					ctx, nil, caller,
					big.NewInt(0),
					false,
					val.String(),
					amount,
					creationHeight,
				)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		When("GetUnbondingDelegation", func() {

			It("should succeed", func() {
				// Undelegate.
				amount, ok := new(big.Int).SetString("1", 10)
				Expect(ok).To(BeTrue())
				_, err := contract.Undelegate(
					ctx, nil, caller,
					big.NewInt(0),
					false,
					cosmlib.ValAddressToEthAddress(val),
					amount,
				)
				Expect(err).ToNot(HaveOccurred())

				res, err := contract.GetUnbondingDelegation(
					ctx, nil, caller,
					big.NewInt(0),
					true,
					caller,
					cosmlib.ValAddressToEthAddress(val),
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(res).ToNot(BeNil())
			})
		})

		When("GetUnbondingDelegation0", func() {

			It("should succeed", func() {
				// Undelegate.
				amount, ok := new(big.Int).SetString("1", 10)
				Expect(ok).To(BeTrue())
				_, err := contract.Undelegate(
					ctx, nil, caller,
					big.NewInt(0),
					false,
					cosmlib.ValAddressToEthAddress(val),
					amount,
				)
				Expect(err).ToNot(HaveOccurred())

				res, err := contract.GetUnbondingDelegation0(
					ctx, nil, caller,
					big.NewInt(0),
					true,
					cosmlib.AddressToAccAddress(caller).String(),
					val.String(),
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(res).ToNot(BeNil())
			})
		})

		When("GetRedelegations", func() {
			It("should succeed", func() {

				amount, ok := new(big.Int).SetString("220000000000000000000", 10)
				Expect(ok).To(BeTrue())

				err := FundAccount(
					ctx,
					bk,
					cosmlib.AddressToAccAddress(caller),
					sdk.NewCoins(
						sdk.NewCoin(
							"stake",
							sdkmath.NewIntFromBigInt(amount),
						),
					),
				)
				Expect(err).ToNot(HaveOccurred())

				validator.Status = stakingtypes.Bonded
				sk.SetValidator(ctx, validator)

				ret, err := contract.Delegate(
					ctx, nil, caller,
					big.NewInt(0),
					true,
					cosmlib.ValAddressToEthAddress(val),
					amount,
				)
				Expect(utils.MustGetAs[bool](ret[0])).To(BeTrue())
				Expect(err).ToNot(HaveOccurred())

				ret, err = contract.GetDelegation(ctx, nil, caller,
					big.NewInt(0),
					true,
					caller,
					cosmlib.ValAddressToEthAddress(val),
				)
				Expect(err).ToNot(HaveOccurred())

				Expect(utils.MustGetAs[*big.Int](ret[0]).Cmp(new(big.Int).Add(amount, big.NewInt(9)))).To(Equal(0))

				otherValidator.Status = stakingtypes.Bonded
				sk.SetValidator(ctx, otherValidator)

				ret, err = contract.BeginRedelegate(
					ctx,
					nil,
					caller,
					big.NewInt(0),
					false,
					cosmlib.ValAddressToEthAddress(val),
					cosmlib.ValAddressToEthAddress(otherVal),
					amount,
				)
				Expect(utils.MustGetAs[bool](ret[0])).To(BeTrue())
				Expect(err).ToNot(HaveOccurred())

				ret, err = contract.GetDelegation(ctx, nil, caller,
					big.NewInt(0),
					true,
					caller,
					cosmlib.ValAddressToEthAddress(val),
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(utils.MustGetAs[*big.Int](ret[0]).Cmp(big.NewInt(9))).To(Equal(0))

				ret, err = contract.GetDelegation(ctx, nil, caller,
					big.NewInt(0),
					true,
					caller,
					cosmlib.ValAddressToEthAddress(otherVal),
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(utils.MustGetAs[*big.Int](ret[0]).Cmp(amount)).To(Equal(0))

				ret, err = contract.GetRedelegations(
					ctx, nil, caller,
					big.NewInt(0),
					true,
					caller,
					cosmlib.ValAddressToEthAddress(val),
					cosmlib.ValAddressToEthAddress(otherVal),
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(ret).ToNot(BeNil())
			})
		})

		When("GetRedelegations0", func() {
			When("Calling Helper Methods", func() {
				When("delegationHelper", func() {
					It("should fail if the del address is not valid", func() {
						_, err := contract.getDelegationHelper(
							ctx,
							sdk.AccAddress(""),
							val,
						)
						Expect(err).To(HaveOccurred())
					})
					It("should fail if the val address is not valid", func() {
						_, err := contract.getDelegationHelper(
							ctx,
							del,
							sdk.ValAddress(""),
						)
						Expect(err).To(HaveOccurred())
					})
					It("should not error if there is no delegation", func() {
						vals, err := contract.getDelegationHelper(
							ctx,
							del,
							otherVal,
						)
						Expect(err).ToNot(HaveOccurred())
						del := utils.MustGetAs[*big.Int](vals[0])
						Expect(del.Cmp(big.NewInt(0))).To(Equal(0))
					})
					It("should succeed", func() {
						_, err := contract.getDelegationHelper(
							ctx,
							del,
							val,
						)
						Expect(err).ToNot(HaveOccurred())
					})
				})

				When("getUnbondingDelegationHelper", func() {
					It("should fail if caller address is wrong", func() {
						_, err := contract.getUnbondingDelegationHelper(
							ctx,
							sdk.AccAddress([]byte("")),
							val,
						)
						Expect(err).To(HaveOccurred())
					})

					It("should fail if there is no unbonding delegation", func() {
						vals, err := contract.getUnbondingDelegationHelper(
							ctx,
							cosmlib.AddressToAccAddress(caller),
							otherVal,
						)
						Expect(err).ToNot(HaveOccurred())
						_, ok := utils.GetAs[[]stakingtypes.UnbondingDelegationEntry](vals[0])
						Expect(ok).To(BeTrue())
					})

					It("should succeed", func() {
						// Undelegate.
						amount, ok := new(big.Int).SetString("1", 10)
						Expect(ok).To(BeTrue())
						_, err := contract.Undelegate(
							ctx,
							nil,
							caller,
							big.NewInt(0),
							false,
							cosmlib.ValAddressToEthAddress(val),
							amount,
						)
						Expect(err).ToNot(HaveOccurred())

						_, err = contract.getUnbondingDelegationHelper(
							ctx,
							cosmlib.AddressToAccAddress(caller),
							val,
						)
						Expect(err).ToNot(HaveOccurred())
					})
				})

				When("getRedelegationHelper", func() {
					It("should fail if caller address is wrong", func() {
						_, err := contract.getRedelegationsHelper(
							ctx,
							sdk.AccAddress([]byte("")),
							val,
							otherVal,
						)
						Expect(err).To(HaveOccurred())
					})

					It("should fail if there is no redelegation", func() {
						_, err := contract.getRedelegationsHelper(
							ctx,
							cosmlib.AddressToAccAddress(caller),
							val,
							otherVal,
						)
						Expect(err).To(HaveOccurred())
					})

					It("should succeed", func() {
						// Redelegate.
						amount, ok := new(big.Int).SetString("1", 10)
						Expect(ok).To(BeTrue())

						_, err := contract.BeginRedelegate(
							ctx,
							nil,
							caller,
							big.NewInt(0),
							false,
							cosmlib.ValAddressToEthAddress(val),
							cosmlib.ValAddressToEthAddress(otherVal),
							amount,
						)
						Expect(err).ToNot(HaveOccurred())
					})
				})
			})
		})

		When("GetActiveValidators", func() {
			It("gets active validators", func() {
				// Set the validator to be bonded.
				validator.Status = stakingtypes.Bonded
				sk.SetValidator(ctx, validator)

				// Get the active validators.
				res, err := contract.GetActiveValidators(ctx, nil, caller, big.NewInt(0), true)
				Expect(err).ToNot(HaveOccurred())
				Expect(res).To(HaveLen(1))
				addrs := utils.MustGetAs[[]common.Address](res[0])
				Expect(addrs[0]).To(Equal(cosmlib.ValAddressToEthAddress(val)))
			})
		})
	})
})

func FundAccount(ctx sdk.Context, bk bankkeeper.BaseKeeper, account sdk.AccAddress, coins sdk.Coins) error {
	if err := bk.MintCoins(ctx, stakingtypes.ModuleName, coins); err != nil {
		return err
	}
	return bk.SendCoinsFromModuleToAccount(ctx, stakingtypes.ModuleName, account, coins)
}
