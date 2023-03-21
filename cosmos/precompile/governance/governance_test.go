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

package governance

import (
	"math/big"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"cosmossdk.io/math"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	governancekeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/precompile"
	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/lib/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGovernancePrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/precompile/governance")
}

var _ = Describe("Governance Precompile", func() {
	var (
		ctx      sdk.Context
		bk       bankkeeper.Keeper
		gk       *governancekeeper.Keeper
		caller   sdk.AccAddress
		mockCtrl *gomock.Controller
		contract *Contract
	)

	BeforeEach(func() {
		t := GinkgoTestReporter{}
		mockCtrl = gomock.NewController(t)
		caller = cosmlib.AddressToAccAddress(testutil.Alice)
		ctx, bk, gk = setup(mockCtrl, caller)
		contract = utils.MustGetAs[*Contract](
			NewPrecompileContract(gk, governancekeeper.NewMsgServerImpl(gk)),
		)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	When("Submitting a proposal", func() {
		It("should fail if the message is not of type", func() {
			res, err := contract.SubmitProposal(
				ctx,
				nil,
				cosmlib.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				"invalid",
			)
			Expect(err).To(MatchError(precompile.ErrInvalidAny))
			Expect(res).To(BeNil())
		})
		It("should fail if the initial deposit is wrong type", func() {
			res, err := contract.SubmitProposal(
				ctx,
				nil,
				cosmlib.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				[]*codectypes.Any{},
				"invalid",
			)
			Expect(err).To(MatchError(precompile.ErrInvalidCoin))
			Expect(res).To(BeNil())
		})
		It("should fail if metadata is of wrong type", func() {
			res, err := contract.SubmitProposal(
				ctx,
				nil,
				cosmlib.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				[]*codectypes.Any{},
				[]generated.IGovernanceModuleCoin{},
				123,
			)
			Expect(err).To(MatchError(precompile.ErrInvalidString))
			Expect(res).To(BeNil())
		})
		It("should fail if title is of wrong type", func() {
			res, err := contract.SubmitProposal(
				ctx,
				nil,
				cosmlib.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				[]*codectypes.Any{},
				[]generated.IGovernanceModuleCoin{},
				"metadata",
				123,
			)
			Expect(err).To(MatchError(precompile.ErrInvalidString))
			Expect(res).To(BeNil())
		})
		It("should fail if summary is of wrong type", func() {
			res, err := contract.SubmitProposal(
				ctx,
				nil,
				cosmlib.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				[]*codectypes.Any{},
				[]generated.IGovernanceModuleCoin{},
				"metadata",
				"title",
				123,
			)
			Expect(err).To(MatchError(precompile.ErrInvalidString))
			Expect(res).To(BeNil())
		})
		It("should fail if expedited is of wrong type", func() {
			res, err := contract.SubmitProposal(
				ctx,
				nil,
				cosmlib.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				[]*codectypes.Any{},
				[]generated.IGovernanceModuleCoin{},
				"metadata",
				"title",
				"summary",
				123,
			)
			Expect(err).To(MatchError(precompile.ErrInvalidBool))
			Expect(res).To(BeNil())
		})
		It("should succeed", func() {
			initDeposit := sdk.NewCoins(sdk.NewInt64Coin("usdc", 100))
			govAcct := gk.GetGovernanceAccount(ctx).GetAddress()
			fundAccount(ctx, bk, govAcct, initDeposit)
			message := &banktypes.MsgSend{
				FromAddress: govAcct.String(),
				ToAddress:   caller.String(),
				Amount:      initDeposit,
			}

			metadata := "metadata"
			title := "title"
			summary := "summary"

			msg, err := codectypes.NewAnyWithValue(message)
			Expect(err).ToNot(HaveOccurred())

			res, err := contract.SubmitProposal(
				ctx,
				nil,
				cosmlib.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				[]*codectypes.Any{msg},
				[]generated.IGovernanceModuleCoin{
					{
						Amount: 100,
						Denom:  "usdc",
					},
				},
				metadata,
				title,
				summary,
				false,
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).ToNot(BeNil())
		})
	})

	When("Canceling a proposal", func() {
		It("should fail if the proposal ID is invalid", func() {
			res, err := contract.CancelProposal(
				ctx,
				nil,
				cosmlib.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				"invalid",
			)
			Expect(err).To(MatchError(precompile.ErrInvalidUint64))
			Expect(res).To(BeNil())
		})
		It("should fail if the proposal does not exist", func() {
			res, err := contract.CancelProposal(
				ctx,
				nil,
				cosmlib.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				big.NewInt(1),
			)
			Expect(err).To(HaveOccurred())
			Expect(res).To(BeNil())
		})
		It("should succeed", func() {
			gk.SetProposal(ctx, v1.Proposal{
				Id:       1,
				Proposer: caller.String(),
				Messages: []*codectypes.Any{},
				Status:   v1.StatusVotingPeriod,
			})
			res, err := contract.CancelProposal(
				ctx,
				nil,
				cosmlib.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				uint64(1),
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).ToNot(BeNil())
		})

	})

	When("Voting on a proposal", func() {
		BeforeEach(func() {
			gk.SetProposal(ctx, v1.Proposal{
				Id:       1,
				Proposer: caller.String(),
				Messages: []*codectypes.Any{},
				Status:   v1.StatusVotingPeriod,
			})
		})
		It("should fail if the proposal ID is of invalid type", func() {
			res, err := contract.Vote(
				ctx,
				nil,
				cosmlib.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				"invalid",
				int32(1),
				"metadata",
			)
			Expect(err).To(MatchError(precompile.ErrInvalidUint64))
			Expect(res).To(BeNil())
		})
		It("should fail if the vote option is of invalid type", func() {
			res, err := contract.Vote(
				ctx,
				nil,
				cosmlib.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				uint64(1),
				"invalid",
				"metadata",
			)
			Expect(err).To(MatchError(precompile.ErrInvalidInt32))
			Expect(res).To(BeNil())
		})
		It("should fail if the metadata is of invalid type", func() {
			res, err := contract.Vote(
				ctx,
				nil,
				cosmlib.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				uint64(1),
				int32(1),
				123,
			)
			Expect(err).To(MatchError(precompile.ErrInvalidString))
			Expect(res).To(BeNil())
		})
		It("should fail if the proposal does not exist", func() {
			res, err := contract.Vote(
				ctx,
				nil,
				cosmlib.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				uint64(1000),
				int32(1),
				"metadata",
			)
			Expect(err).To(HaveOccurred())
			Expect(res).To(BeNil())
		})
		It("should succeed", func() {
			res, err := contract.Vote(
				ctx,
				nil,
				cosmlib.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				uint64(1),
				int32(1),
				"metadata",
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).ToNot(BeNil())
		})

		When("Voting Weight", func() {
			It("should fail if the proposal ID is of invalid type", func() {
				res, err := contract.VoteWeighted(
					ctx,
					nil,
					cosmlib.AccAddressToEthAddress(caller),
					big.NewInt(0),
					false,
					"invalid",
					[]generated.IGovernanceModuleWeightedVoteOption{},
					"metadata",
				)
				Expect(err).To(MatchError(precompile.ErrInvalidBigInt))
				Expect(res).To(BeNil())
			})
			It("should fail if the vote option is of invalid type", func() {
				res, err := contract.VoteWeighted(
					ctx,
					nil,
					cosmlib.AccAddressToEthAddress(caller),
					big.NewInt(0),
					false,
					uint64(1),
					12,
					"metadata",
				)
				Expect(err).To(MatchError(precompile.ErrInvalidOptions))
				Expect(res).To(BeNil())
			})
			It("should fail if the metadata is of invalid type", func() {
				res, err := contract.VoteWeighted(
					ctx,
					nil,
					cosmlib.AccAddressToEthAddress(caller),
					big.NewInt(0),
					false,
					uint64(1),
					[]generated.IGovernanceModuleWeightedVoteOption{},
					123,
				)
				Expect(err).To(MatchError(precompile.ErrInvalidString))
				Expect(res).To(BeNil())
			})
			It("should fail if the proposal does not exist", func() {
				res, err := contract.VoteWeighted(
					ctx,
					nil,
					cosmlib.AccAddressToEthAddress(caller),
					big.NewInt(0),
					false,
					uint64(1000),
					[]generated.IGovernanceModuleWeightedVoteOption{},
					"metadata",
				)
				Expect(err).To(HaveOccurred())
				Expect(res).To(BeNil())
			})
			It("should succeed", func() {
				weight, err := math.LegacyNewDecFromStr("0.4")
				Expect(err).ToNot(HaveOccurred())
				options := []generated.IGovernanceModuleWeightedVoteOption{
					{
						VoteOption: int32(1),
						Weight:     weight.String(),
					},
				}
				res, err := contract.VoteWeighted(
					ctx,
					nil,
					cosmlib.AccAddressToEthAddress(caller),
					big.NewInt(0),
					false,
					uint64(1),
					options,
					"metadata",
				)
				Expect(err).ToNot(HaveOccurred())
				Expect(res).ToNot(BeNil())
			})
		})

		When("Reading Methods", func() {
			BeforeEach(func() {
				gk.SetProposal(ctx, v1.Proposal{
					Id:               2,
					Proposer:         caller.String(),
					Messages:         []*codectypes.Any{},
					Status:           v1.StatusVotingPeriod,
					FinalTallyResult: &v1.TallyResult{},
					SubmitTime:       &time.Time{},
					DepositEndTime:   &time.Time{},
					TotalDeposit:     sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(100))),
					VotingStartTime:  &time.Time{},
					VotingEndTime:    &time.Time{},
					Metadata:         "metadata",
					Title:            "title",
					Summary:          "summary",
					Expedited:        false,
				})
				gk.SetProposal(ctx, v1.Proposal{
					Id:               3,
					Proposer:         caller.String(),
					Messages:         []*codectypes.Any{},
					Status:           v1.StatusVotingPeriod,
					FinalTallyResult: &v1.TallyResult{},
					SubmitTime:       &time.Time{},
					DepositEndTime:   &time.Time{},
					TotalDeposit:     sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(100))),
					VotingStartTime:  &time.Time{},
					VotingEndTime:    &time.Time{},
					Metadata:         "metadata",
					Title:            "title",
					Summary:          "summary",
					Expedited:        false,
				})
			})

			When("GetProposal", func() {
				It("should fail if the proposal ID is of invalid type", func() {
					res, err := contract.GetProposal(
						ctx,
						nil,
						cosmlib.AccAddressToEthAddress(caller),
						big.NewInt(0),
						false,
						"invalid",
					)
					Expect(err).To(MatchError(precompile.ErrInvalidUint64))
					Expect(res).To(BeNil())
				})
				It("should get the proposal", func() {
					res, err := contract.GetProposal(
						ctx,
						nil,
						cosmlib.AccAddressToEthAddress(caller),
						big.NewInt(0),
						false,
						uint64(2),
					)
					Expect(err).ToNot(HaveOccurred())
					Expect(res).ToNot(BeNil())
					Expect(res).To(HaveLen(1))
				})
			})
			When("GetProposals", func() {
				BeforeEach(func() {
					// Not filled proposal, hence will panic the parser.
					_, err := contract.CancelProposal(ctx, nil, common.Address(caller), big.NewInt(0), false, uint64(1))
					Expect(err).ToNot(HaveOccurred())
				})
				It("should fail if the status is of invalid type", func() {
					res, err := contract.GetProposals(
						ctx,
						nil,
						cosmlib.AccAddressToEthAddress(caller),
						big.NewInt(0),
						false,
						"",
					)
					Expect(err).To(MatchError(precompile.ErrInvalidInt32))
					Expect(res).To(BeNil())
				})
				It("should get the proposals", func() {
					res, err := contract.GetProposals(
						ctx,
						nil,
						cosmlib.AccAddressToEthAddress(caller),
						big.NewInt(0),
						false,
						int32(0),
					)
					Expect(err).ToNot(HaveOccurred())
					Expect(res).ToNot(BeNil())
				})
			})
		})
	})
})
