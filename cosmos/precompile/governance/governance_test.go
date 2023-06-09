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
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	sdkmath "cosmossdk.io/math"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	governancekeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	governancetypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/governance"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/precompile"
	precomtest "pkg.berachain.dev/polaris/cosmos/precompile/test"
	"pkg.berachain.dev/polaris/cosmos/types"
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
		t := precomtest.GinkgoTestReporter{}
		mockCtrl = gomock.NewController(t)
		config.SetupCosmosConfig()
		caller = cosmlib.AddressToAccAddress(testutil.Alice)
		ctx, bk, gk = precomtest.Setup(mockCtrl, caller)
		contract = utils.MustGetAs[*Contract](NewPrecompileContract(
			governancekeeper.NewMsgServerImpl(gk),
			governancekeeper.NewQueryServer(gk),
		))
		config.SetupCosmosConfig()
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("Should have precompile tests and custom value decoders", func() {
		Expect(contract.PrecompileMethods()).To(HaveLen(6))
		Expect(contract.CustomValueDecoders()).ToNot(BeNil())
	})

	When("Unmarshal message and return any", func() {
		var msg banktypes.MsgSend
		BeforeEach(func() {
			msg = banktypes.MsgSend{
				FromAddress: caller.String(),
				ToAddress:   testutil.Bob.String(),
				Amount:      sdk.NewCoins(sdk.NewInt64Coin("abera", 100)),
			}
		})

		It("should fail if the message is wrong type", func() {
			bz := []byte("invalid")
			_, err := unmarshalMsgAndReturnAny(bz)
			Expect(err).To(HaveOccurred())
		})
		It("should succeed if the message is correct types", func() {
			bz, err := msg.Marshal()
			Expect(err).ToNot(HaveOccurred())
			a, err := unmarshalMsgAndReturnAny(bz)
			Expect(err).ToNot(HaveOccurred())
			Expect(a).ToNot(BeNil())
		})
	})

	When("submitting proposal handler", func() {
		It("should fail if the proposal cant be unmarshalled", func() {
			_, err := contract.submitProposalHelper(
				context.TODO(), []byte("invalid"), nil,
			)
			Expect(err).To(HaveOccurred())
		})
	})

	When("Submitting a proposal", func() {
		It("Should fail if proposal is of wrong type", func() {
			res, err := contract.SubmitProposal(
				ctx,
				nil,
				cosmlib.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				"invalid",
			)
			Expect(err).To(MatchError(precompile.ErrInvalidBytes))
			Expect(res).To(BeNil())
		})
		It("Should fail if the message is of wrong type", func() {
			res, err := contract.SubmitProposal(
				ctx,
				nil,
				cosmlib.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				[]byte{},
				"invalid",
			)
			Expect(err).To(MatchError(precompile.ErrInvalidBytes))
			Expect(res).To(BeNil())
		})
		It("should succeed", func() {
			initDeposit := sdk.NewCoins(sdk.NewInt64Coin("abera", 100))
			govAcct := gk.GetGovernanceAccount(ctx).GetAddress()
			err := cosmlib.MintCoinsToAddress(
				ctx, bk, governancetypes.ModuleName,
				cosmlib.AccAddressToEthAddress(govAcct), "abera", big.NewInt(100),
			)
			Expect(err).ToNot(HaveOccurred())
			message := &banktypes.MsgSend{
				FromAddress: govAcct.String(),
				ToAddress:   caller.String(),
				Amount:      initDeposit,
			}
			metadata := "metadata"
			title := "title"
			summary := "summary "
			msgBz, err := message.Marshal()
			Expect(err).ToNot(HaveOccurred())
			// Create and marshal the proposal.
			proposal := v1.MsgSubmitProposal{
				InitialDeposit: initDeposit,
				Proposer:       caller.String(),
				Metadata:       metadata,
				Title:          title,
				Summary:        summary,
				Expedited:      false,
			}
			proposalBz, err := proposal.Marshal()
			Expect(err).ToNot(HaveOccurred())
			// Submit the proposal.
			res, err := contract.SubmitProposal(
				ctx,
				nil,
				cosmlib.AccAddressToEthAddress(caller),
				big.NewInt(0),
				false,
				proposalBz,
				msgBz,
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
			err := gk.SetProposal(ctx, v1.Proposal{
				Id:       1,
				Proposer: caller.String(),
				Messages: []*codectypes.Any{},
				Status:   v1.StatusVotingPeriod,
			})
			Expect(err).ToNot(HaveOccurred())
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
			err := gk.SetProposal(ctx, v1.Proposal{
				Id:       1,
				Proposer: caller.String(),
				Messages: []*codectypes.Any{},
				Status:   v1.StatusVotingPeriod,
			})
			Expect(err).ToNot(HaveOccurred())
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
			Expect(res).To(HaveLen(1))
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
				Expect(res).To(HaveLen(1))
			})
			It("should succeed", func() {
				weight, err := sdkmath.LegacyNewDecFromStr("1")
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
				err := gk.SetProposal(ctx, v1.Proposal{
					Id:               2,
					Proposer:         caller.String(),
					Messages:         []*codectypes.Any{},
					Status:           v1.StatusVotingPeriod,
					FinalTallyResult: &v1.TallyResult{},
					SubmitTime:       &time.Time{},
					DepositEndTime:   &time.Time{},
					TotalDeposit:     sdk.NewCoins(sdk.NewCoin("stake", sdkmath.NewInt(100))),
					VotingStartTime:  &time.Time{},
					VotingEndTime:    &time.Time{},
					Metadata:         "metadata",
					Title:            "title",
					Summary:          "summary",
					Expedited:        false,
				})
				Expect(err).ToNot(HaveOccurred())
				err = gk.SetProposal(ctx, v1.Proposal{
					Id:               3,
					Proposer:         caller.String(),
					Messages:         []*codectypes.Any{},
					Status:           v1.StatusVotingPeriod,
					FinalTallyResult: &v1.TallyResult{},
					SubmitTime:       &time.Time{},
					DepositEndTime:   &time.Time{},
					TotalDeposit:     sdk.NewCoins(sdk.NewCoin("stake", sdkmath.NewInt(100))),
					VotingStartTime:  &time.Time{},
					VotingEndTime:    &time.Time{},
					Metadata:         "metadata",
					Title:            "title",
					Summary:          "summary",
					Expedited:        false,
				})
				Expect(err).ToNot(HaveOccurred())
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
