package staking_test

import (
	"math/big"
	"testing"

	"cosmossdk.io/math"
	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/eth/types/abi"
	"github.com/berachain/stargazer/precompile/contracts/solidity/generated"
	"github.com/berachain/stargazer/precompile/staking"
	"github.com/berachain/stargazer/testutil"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingkeepertypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStakingPrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "precompile/staking")
}

func createValAddrs(count int) ([]sdk.AccAddress, []sdk.ValAddress) {
	addrs := simtestutil.CreateIncrementalAccounts(count)
	valAddrs := simtestutil.ConvertAddrsToValAddrs(addrs)

	return addrs, valAddrs
}

func NewValidator(operator sdk.ValAddress, pubKey cryptotypes.PubKey) (stakingkeepertypes.Validator, error) {
	return stakingkeepertypes.NewValidator(operator, pubKey, stakingkeepertypes.Description{})
}

var (
	PKs = simtestutil.CreateTestPubKeys(500)
)

var _ = Describe("Staking", func() {
	var (
		sk stakingkeeper.Keeper
		bk bankkeeper.BaseKeeper

		ctx sdk.Context

		contract *staking.Contract
	)

	BeforeEach(func() {
		ctx, _, bk, sk = testutil.SetupMinimalKeepers()
		contract = staking.NewContract(&sk)
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
			Expect(contract.PrecompileMethods()).To(HaveLen(len(contract.ABIMethods())))
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
			del       sdk.AccAddress
			val       sdk.ValAddress
			validator stakingkeepertypes.Validator
			otherVal  sdk.ValAddress
			caller    common.Address
		)

		BeforeEach(func() {
			delegates, validators := createValAddrs(2)
			del, val, otherVal = delegates[0], validators[0], validators[1]
			caller = common.BytesToAddress(del.Bytes())

			amount, ok := new(big.Int).SetString("22000000000000000000", 10) // 22 tokens.
			Expect(ok).To(BeTrue())
			var err error
			validator, err = NewValidator(val, PKs[0])
			Expect(err).ToNot(HaveOccurred())
			otherValidator, err := NewValidator(otherVal, PKs[1])
			Expect(err).ToNot(HaveOccurred())
			validator, _ = validator.AddTokensFromDel(sdk.NewIntFromBigInt(amount))
			otherValidator, _ = otherValidator.AddTokensFromDel(sdk.NewIntFromBigInt(amount))
			validator = stakingkeeper.TestingUpdateValidator(&sk, ctx, validator, true)
			otherValidator = stakingkeeper.TestingUpdateValidator(&sk, ctx, otherValidator, true)

			delegation := stakingkeepertypes.NewDelegation(del, val, math.LegacyNewDec(9))
			sk.SetDelegation(ctx, delegation)

			// Check that the delegation was created.
			res, found := sk.GetDelegation(ctx, del, val)
			Expect(found).To(BeTrue())
			Expect(res).To(Equal(delegation))

			// Set the denom.
			defaultParams := stakingkeepertypes.DefaultParams()
			defaultParams.BondDenom = "stake"
			sk.SetParams(ctx, defaultParams)

		})

		When("DelegateAddrInput", func() {
			It("should fail if input is not a common.Address", func() {
				res, err := contract.DelegateAddrInput(
					ctx,
					caller,
					big.NewInt(0),
					true,
					"0x",
				)
				Expect(err).To(MatchError(staking.ErrInvalidValidatorAddr))
				Expect(res).To(BeNil())
			})

			It("should fail if the amount is not a *big.Int", func() {
				res, err := contract.DelegateAddrInput(
					ctx,
					caller,
					big.NewInt(0),
					true,
					common.BytesToAddress(val.Bytes()),
					"amount",
				)
				Expect(err).To(MatchError(staking.ErrInvalidBigInt))
				Expect(res).To(BeNil())
			})

			It("should succeed", func() {
				amountToDelegate, ok := new(big.Int).SetString("22000000000000000000", 10)
				Expect(ok).To(BeTrue())
				FundAccount(
					ctx,
					bk,
					del,
					sdk.NewCoins(
						sdk.NewCoin(
							"stake",
							sdk.NewIntFromBigInt(amountToDelegate),
						),
					),
				)

				_, err := contract.DelegateAddrInput(
					ctx,
					caller,
					big.NewInt(0),
					true,
					common.BytesToAddress(val.Bytes()),
					amountToDelegate,
				)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		When("DelegateStringInput", func() {
			It("should fail if input is not a string", func() {
				res, err := contract.DelegateStringInput(
					ctx,
					caller,
					big.NewInt(0),
					false,
					90909,
				)
				Expect(err).To(MatchError(staking.ErrInvalidString))
				Expect(res).To(BeNil())
			})

			It("should fail if the amount is not a *big.Int", func() {
				res, err := contract.DelegateStringInput(
					ctx,
					caller,
					big.NewInt(0),
					false,
					val.String(),
					"amount",
				)
				Expect(err).To(MatchError(staking.ErrInvalidBigInt))
				Expect(res).To(BeNil())
			})

			It("should fail if the string is not a valid address", func() {
				res, err := contract.DelegateStringInput(
					ctx,
					caller,
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
				FundAccount(
					ctx,
					bk,
					del,
					sdk.NewCoins(
						sdk.NewCoin(
							"stake",
							sdk.NewIntFromBigInt(amountToDelegate),
						),
					),
				)

				_, err := contract.DelegateStringInput(
					ctx,
					caller,
					big.NewInt(0),
					false,
					val.String(),
					amountToDelegate,
				)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		When("GetDelegationAddrInput", func() {
			It("should return an error if the input is not a common.Address", func() {
				res, err := contract.GetDelegationAddrInput(
					ctx,
					caller,
					big.NewInt(0),
					true,
					"0x",
				)
				Expect(err).To(MatchError(staking.ErrInvalidValidatorAddr))
				Expect(res).To(BeNil())
			})

			It("should return the correct delegation", func() {
				res, err := contract.GetDelegationAddrInput(
					ctx,
					caller,
					big.NewInt(0),
					true,
					common.BytesToAddress(val.Bytes()),
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(res[0]).To(Equal(big.NewInt(9))) // should have correct shares
			})
		})

		When("GetDelegationStringInput", func() {
			It("should error if not string", func() {
				res, err := contract.GetDelegationStringInput(
					ctx,
					caller,
					big.NewInt(0),
					true,
					10,
				)
				Expect(err).To(MatchError(staking.ErrInvalidString))
				Expect(res).To(BeNil())
			})

			It("should error if not bech32", func() {
				res, err := contract.GetDelegationStringInput(
					ctx,
					caller,
					big.NewInt(0),
					true,
					"0x",
				)
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			})

			It("should return the correct delegation", func() {
				res, err := contract.GetDelegationStringInput(
					ctx,
					caller,
					big.NewInt(0),
					true,
					val.String(),
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(res[0]).To(Equal(big.NewInt(9))) // should have correct shares
			})
		})

		When("UndelegateAddrInput", func() {
			It("should fail if the input is not a common.Address", func() {
				res, err := contract.UndelegateAddrInput(
					ctx,
					caller,
					big.NewInt(0),
					true,
					"0x",
					big.NewInt(0),
				)
				Expect(err).To(MatchError(staking.ErrInvalidValidatorAddr))
				Expect(res).To(BeNil())
			})

			It("should fail if the amount is not a *big.Int", func() {
				res, err := contract.UndelegateAddrInput(
					ctx,
					caller,
					big.NewInt(0),
					true,
					common.BytesToAddress(val.Bytes()),
					"amount",
				)
				Expect(err).To(MatchError(staking.ErrInvalidBigInt))
				Expect(res).To(BeNil())
			})

			It("should succeed", func() {
				_, err := contract.UndelegateAddrInput(
					ctx,
					caller,
					big.NewInt(0),
					true,
					common.BytesToAddress(val.Bytes()),
					big.NewInt(1),
				)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		When("UndelegateStringInput", func() {
			It("should fail if the input is not a string", func() {
				res, err := contract.UndelegateStringInput(
					ctx,
					caller,
					big.NewInt(0),
					true,
					90909,
					big.NewInt(0),
				)
				Expect(err).To(MatchError(staking.ErrInvalidString))
				Expect(res).To(BeNil())
			})

			It("should fail if the amount is not a *big.Int", func() {
				res, err := contract.UndelegateStringInput(
					ctx,
					caller,
					big.NewInt(0),
					true,
					val.String(),
					"amount",
				)
				Expect(err).To(MatchError(staking.ErrInvalidBigInt))
				Expect(res).To(BeNil())
			})

			It("should fail if the address is not of type bech32", func() {
				res, err := contract.UndelegateStringInput(
					ctx,
					caller,
					big.NewInt(0),
					true,
					"0x",
					big.NewInt(0),
				)
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			})

			It("should succeed", func() {
				_, err := contract.UndelegateStringInput(
					ctx,
					caller,
					big.NewInt(0),
					true,
					val.String(),
					big.NewInt(1),
				)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		When("BeginRedelegationsAddrInput", func() {
			It("should fail if the srcValue is not a common.Address", func() {
				res, err := contract.BeginRedelegateAddrInput(
					ctx,
					caller,
					big.NewInt(0),
					false,
					10,
					common.BytesToAddress(val.Bytes()),
					big.NewInt(1),
				)
				Expect(err).To(MatchError(staking.ErrInvalidValidatorAddr))
				Expect(res).To(BeNil())
			})

			It("should fail if the dstValue is not a common.Address", func() {
				res, err := contract.BeginRedelegateAddrInput(
					ctx,
					caller,
					big.NewInt(0),
					false,
					common.BytesToAddress(val.Bytes()),
					10,
					big.NewInt(1),
				)
				Expect(err).To(MatchError(staking.ErrInvalidValidatorAddr))
				Expect(res).To(BeNil())
			})

			It("should fail if the amount is not a *big.Int", func() {
				res, err := contract.BeginRedelegateAddrInput(
					ctx,
					caller,
					big.NewInt(0),
					false,
					common.BytesToAddress(val.Bytes()),
					common.BytesToAddress(val.Bytes()),
					"amount",
				)
				Expect(err).To(MatchError(staking.ErrInvalidBigInt))
				Expect(res).To(BeNil())
			})

			It("should succeed", func() {
				_, err := contract.BeginRedelegateAddrInput(
					ctx,
					caller,
					big.NewInt(0),
					false,
					common.BytesToAddress(val.Bytes()),
					common.BytesToAddress(otherVal.Bytes()),
					big.NewInt(1),
				)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		When("BeginRedelegationsStringInput", func() {
			It("should fail if the srcValue is not a string", func() {
				res, err := contract.BeginRedelegateStringInput(
					ctx,
					caller,
					big.NewInt(0),
					false,
					10,
					val.String(),
					big.NewInt(1),
				)
				Expect(err).To(MatchError(staking.ErrInvalidString))
				Expect(res).To(BeNil())
			})

			It("should fail if the dstValue is not a string", func() {
				res, err := contract.BeginRedelegateStringInput(
					ctx,
					caller,
					big.NewInt(0),
					false,
					val.String(),
					10,
					big.NewInt(1),
				)
				Expect(err).To(MatchError(staking.ErrInvalidString))
				Expect(res).To(BeNil())
			})

			It("should fail if the amount is not a *big.Int", func() {
				res, err := contract.BeginRedelegateStringInput(
					ctx,
					caller,
					big.NewInt(0),
					false,
					val.String(),
					otherVal.String(),
					"amount",
				)
				Expect(err).To(MatchError(staking.ErrInvalidBigInt))
				Expect(res).To(BeNil())
			})

			It("should fail if the srcValue is not of type bech32", func() {
				res, err := contract.BeginRedelegateStringInput(
					ctx,
					caller,
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
				res, err := contract.BeginRedelegateStringInput(
					ctx,
					caller,
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
				_, err := contract.BeginRedelegateStringInput(
					ctx,
					caller,
					big.NewInt(0),
					false,
					val.String(),
					otherVal.String(),
					big.NewInt(1),
				)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		When("CancelUnbondingDelegationAddrInput", func() {
			It("should fail if the address is not a common.Address", func() {
				res, err := contract.CancelUnbondingDelegationAddrInput(
					ctx,
					caller,
					big.NewInt(0),
					false,
					"common.BytesToAddress(val.Bytes())",
					big.NewInt(1),
					int64(1),
				)
				Expect(err).To(MatchError(staking.ErrInvalidValidatorAddr))
				Expect(res).To(BeNil())
			})

			It("should fail if the amount is not a *big.Int", func() {
				res, err := contract.CancelUnbondingDelegationAddrInput(
					ctx,
					caller,
					big.NewInt(0),
					false,
					common.BytesToAddress(val.Bytes()),
					"amount",
					int64(1),
				)
				Expect(err).To(MatchError(staking.ErrInvalidBigInt))
				Expect(res).To(BeNil())
			})

			It("should fail if creation height is not an int64", func() {
				res, err := contract.CancelUnbondingDelegationAddrInput(
					ctx,
					caller,
					big.NewInt(0),
					false,
					common.BytesToAddress(val.Bytes()),
					big.NewInt(1),
					"height",
				)
				Expect(err).To(MatchError(staking.ErrInvalidInt64))
				Expect(res).To(BeNil())
			})

			It("should succeed", func() {
				creationHeight := ctx.BlockHeight()
				amount, ok := new(big.Int).SetString("1", 10)
				Expect(ok).To(BeTrue())

				// Undelegate.
				_, err := contract.UndelegateAddrInput(
					ctx,
					caller,
					big.NewInt(0),
					false,
					common.BytesToAddress(val.Bytes()),
					amount,
				)
				Expect(err).ToNot(HaveOccurred())

				_, err = contract.CancelUnbondingDelegationAddrInput(
					ctx,
					caller,
					big.NewInt(0),
					false,
					common.BytesToAddress(val.Bytes()),
					amount,
					creationHeight,
				)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		When("CancelUnbondingDelegationStringInput", func() {
			It("should fail if the address is not a string", func() {
				res, err := contract.CancelUnbondingDelegationStringInput(
					ctx,
					caller,
					big.NewInt(0),
					false,
					10,
					big.NewInt(1),
					int64(1),
				)
				Expect(err).To(MatchError(staking.ErrInvalidString))
				Expect(res).To(BeNil())
			})

			It("should fail if the amount is not a *big.Int", func() {
				res, err := contract.CancelUnbondingDelegationStringInput(
					ctx,
					caller,
					big.NewInt(0),
					false,
					val.String(),
					"amount",
					int64(1),
				)
				Expect(err).To(MatchError(staking.ErrInvalidBigInt))
				Expect(res).To(BeNil())
			})

			It("should fail if creation height is not an int64", func() {
				res, err := contract.CancelUnbondingDelegationStringInput(
					ctx,
					caller,
					big.NewInt(0),
					false,
					val.String(),
					big.NewInt(1),
					"height",
				)
				Expect(err).To(MatchError(staking.ErrInvalidInt64))
				Expect(res).To(BeNil())
			})

			It("should succeed", func() {
				creationHeight := ctx.BlockHeight()
				amount, ok := new(big.Int).SetString("1", 10)
				Expect(ok).To(BeTrue())

				// Undelegate.
				_, err := contract.UndelegateAddrInput(
					ctx,
					caller,
					big.NewInt(0),
					false,
					common.BytesToAddress(val.Bytes()),
					amount,
				)
				Expect(err).ToNot(HaveOccurred())

				_, err = contract.CancelUnbondingDelegationStringInput(
					ctx,
					caller,
					big.NewInt(0),
					false,
					val.String(),
					amount,
					creationHeight,
				)
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})
})

func FundAccount(ctx sdk.Context, bk bankkeeper.BaseKeeper, account sdk.AccAddress, coins sdk.Coins) error {
	bk.MintCoins(ctx, stakingkeepertypes.ModuleName, coins)
	return bk.SendCoinsFromModuleToAccount(ctx, stakingkeepertypes.ModuleName, account, coins)
}
