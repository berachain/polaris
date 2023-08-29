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

	cbindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/lib"
	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/governance"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/precompile/testutil"
	testutils "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/cosmos/types"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/core/vm"
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
		sdkCtx   sdk.Context
		bk       bankkeeper.Keeper
		gk       *governancekeeper.Keeper
		caller   sdk.AccAddress
		mockCtrl *gomock.Controller
		contract *Contract
		sf       *ethprecompile.StatefulFactory
		ctx      context.Context
	)

	BeforeEach(func() {
		t := testutil.GinkgoTestReporter{}
		mockCtrl = gomock.NewController(t)
		types.SetupCosmosConfig()
		caller = testutils.Alice.Bytes()
		sdkCtx, bk, gk = testutil.Setup(mockCtrl, caller)
		contract = utils.MustGetAs[*Contract](NewPrecompileContract(
			governancekeeper.NewMsgServerImpl(gk),
			governancekeeper.NewQueryServer(gk),
		))
		types.SetupCosmosConfig()
		sf = ethprecompile.NewStatefulFactory()
		ctx = vm.NewPolarContext(
			sdkCtx,
			nil,
			common.BytesToAddress(caller),
			big.NewInt(0),
		)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("Should have precompile tests and custom value decoders", func() {
		_, err := sf.Build(contract, nil)
		Expect(err).ToNot(HaveOccurred())
		Expect(contract.CustomValueDecoders()).To(HaveLen(1))
	})

	When("Unmarshal message and return any", func() {
		var msg banktypes.MsgSend
		BeforeEach(func() {
			msg = banktypes.MsgSend{
				FromAddress: caller.String(),
				ToAddress:   testutils.Bob.String(),
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
			_, err := contract.SubmitProposal(
				vm.NewPolarContext(sdk.Context{}, nil, common.Address{}, nil),
				[]byte("invalid"),
			)
			Expect(err).To(HaveOccurred())
		})
	})

	When("Submitting a proposal", func() {
		It("should succeed", func() {
			initDeposit := sdk.NewCoins(sdk.NewInt64Coin("abera", 100))
			govAcct := gk.GetGovernanceAccount(ctx).GetAddress()
			err := cosmlib.MintCoinsToAddress(
				sdk.UnwrapSDKContext(vm.UnwrapPolarContext(ctx).Context()),
				bk,
				governancetypes.ModuleName,
				common.BytesToAddress(govAcct),
				"abera",
				big.NewInt(100),
			)
			Expect(err).ToNot(HaveOccurred())

			// Create and marshal the proposal.
			proposal := v1.MsgSubmitProposal{
				InitialDeposit: initDeposit,
				Proposer:       caller.String(),
				Metadata:       "metadata",
				Title:          "title",
				Summary:        "summary",
				Expedited:      false,
			}
			proposalBz, err := proposal.Marshal()
			Expect(err).ToNot(HaveOccurred())
			// Submit the proposal.
			res, err := contract.SubmitProposal(
				ctx,
				proposalBz,
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).ToNot(BeNil())
		})
	})

	When("Canceling a proposal", func() {
		It("should succeed", func() {
			err := gk.SetProposal(ctx, v1.Proposal{
				Id:       1,
				Proposer: caller.String(),
				Messages: []*codectypes.Any{},
				Status:   v1.StatusVotingPeriod,
			})
			Expect(err).ToNot(HaveOccurred())
			time, height, err := contract.CancelProposal(
				ctx,
				uint64(1),
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(time).ToNot(BeZero())
			Expect(height).ToNot(BeZero())
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

		It("should fail if the proposal does not exist", func() {
			res, err := contract.Vote(
				ctx,
				uint64(1000),
				int32(1),
				"metadata",
			)
			Expect(res).To(BeFalse())
			Expect(err).To(HaveOccurred())
		})

		It("should succeed", func() {
			res, err := contract.Vote(
				ctx,
				uint64(1),
				int32(1),
				"metadata",
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).ToNot(BeNil())
		})

		When("Voting Weight", func() {
			It("should fail if the proposal does not exist", func() {
				res, err := contract.VoteWeighted(
					ctx,
					uint64(1000),
					[]generated.IGovernanceModuleWeightedVoteOption{},
					"metadata",
				)
				Expect(res).To(BeFalse())
				Expect(err).To(HaveOccurred())
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
				It("should get the proposal", func() {
					res, err := contract.GetProposal(
						ctx,
						uint64(2),
					)
					Expect(err).ToNot(HaveOccurred())
					Expect(res).ToNot(BeNil())
				})
			})
			When("GetProposals", func() {
				BeforeEach(func() {
					// Not filled proposal, hence will panic the parser.
					_, _, err := contract.CancelProposal(ctx, uint64(1))
					Expect(err).ToNot(HaveOccurred())
				})
				It("should get the proposals", func() {
					res, pageRes, err := contract.GetProposals(
						ctx,
						int32(0),
						cbindings.CosmosPageRequest{
							Key:        "test",
							Offset:     0,
							Limit:      10,
							CountTotal: true,
							Reverse:    false,
						},
					)
					Expect(err).ToNot(HaveOccurred())
					Expect(res).ToNot(BeNil())
					Expect(pageRes).ToNot(BeNil())
				})
			})
			Context("Deposits", func() {
				var coins sdk.Coins
				BeforeEach(func() {
					coins = sdk.NewCoins(sdk.NewCoin("abera", sdkmath.NewInt(100)))
					_, err := gk.AddDeposit(
						ctx, uint64(2),
						caller,
						coins,
					)
					Expect(err).ToNot(HaveOccurred())
				})
				When("GetProposalDeposits", func() {
					It("should get the proposal deposits", func() {
						res, err := contract.GetProposalDeposits(
							ctx,
							uint64(2),
						)
						Expect(err).ToNot(HaveOccurred())
						Expect(res).ToNot(BeNil())
						Expect(res).To(HaveLen(2))
						Expect(res[0].Denom).To(Equal(coins[0].Denom))
						Expect(res[0].Amount).To(Equal(coins[0].Amount.BigInt()))
					})
				})
				When("GetProposalDepositsByDepositor", func() {
					It("should get the proposal deposits by depositor", func() {
						res, err := contract.GetProposalDepositsByDepositor(
							ctx,
							uint64(2),
							common.BytesToAddress(caller),
						)
						Expect(err).ToNot(HaveOccurred())
						Expect(res).ToNot(BeNil())
						Expect(res).To(HaveLen(1))
						Expect(res[0].Denom).To(Equal(coins[0].Denom))
						Expect(res[0].Amount).To(Equal(coins[0].Amount.BigInt()))
					})
				})
			})
			Context("Votes", func() {
				var vote generated.IGovernanceModuleVote
				BeforeEach(func() {
					vote = generated.IGovernanceModuleVote{
						ProposalId: uint64(2),
						Voter:      common.BytesToAddress(caller),
						Options:    []generated.IGovernanceModuleWeightedVoteOption{},
						Metadata:   "metadata",
					}
					err := gk.AddVote(
						ctx, uint64(2),
						caller,
						v1.WeightedVoteOptions{},
						"metadata",
					)
					Expect(err).ToNot(HaveOccurred())
				})
				When("GetProposalTallyResult", func() {
					It("should get the proposal tally result", func() {
						res, err := contract.GetProposalTallyResult(
							ctx,
							uint64(2),
						)
						Expect(err).ToNot(HaveOccurred())
						Expect(res).ToNot(BeNil())
					})
				})
				When("GetProposalVotes", func() {
					It("should get the proposal votes", func() {
						res, pageRes, err := contract.GetProposalVotes(
							ctx,
							uint64(2),
							cbindings.CosmosPageRequest{
								Key:        "test",
								Offset:     0,
								Limit:      10,
								CountTotal: true,
								Reverse:    false,
							},
						)
						Expect(err).ToNot(HaveOccurred())
						Expect(res).ToNot(BeNil())
						Expect(res[0]).To(Equal(vote))
						Expect(pageRes).ToNot(BeNil())
					})
				})
				When("GetProposalVotesByVoter", func() {
					It("should get the proposal votes by voter", func() {
						res, err := contract.GetProposalVotesByVoter(
							ctx,
							uint64(2),
							common.BytesToAddress(caller),
						)
						Expect(err).ToNot(HaveOccurred())
						Expect(res).ToNot(BeNil())
						Expect(res).To(Equal(vote))
					})
				})
			})
			When("GetParams", func() {
				It("should get the params", func() {
					res, err := contract.GetParams(
						ctx,
					)
					Expect(err).ToNot(HaveOccurred())
					Expect(res).ToNot(BeNil())
				})
			})
			When("GetDepositParams", func() {
				It("should get the deposit params", func() {
					res, err := contract.GetDepositParams(
						ctx,
					)
					Expect(err).ToNot(HaveOccurred())
					Expect(res).ToNot(BeNil())
				})
			})
			When("GetVotingParams", func() {
				It("should get the voting params", func() {
					res, err := contract.GetVotingParams(
						ctx,
					)
					Expect(err).ToNot(HaveOccurred())
					Expect(res).ToNot(BeNil())
				})
			})
			When("GetTallyParams", func() {
				It("should get the tally params", func() {
					res, err := contract.GetTallyParams(
						ctx,
					)
					Expect(err).ToNot(HaveOccurred())
					Expect(res).ToNot(BeNil())
				})
			})
			When("GetConstitution", func() {
				It("should get the constitution", func() {
					err := gk.Constitution.Set(ctx, "constitution")
					Expect(err).ToNot(HaveOccurred())
					res, err := contract.GetConstitution(
						ctx,
					)
					Expect(err).ToNot(HaveOccurred())
					Expect(res).ToNot(BeNil())
					Expect(res).To(Equal("constitution"))
				})
			})
		})
	})

	It("Should be able to marshal and unmarshal porposalMsg", func() {
		// Create the send msg.
		msg := banktypes.MsgSend{
			FromAddress: sdk.AccAddress([]byte("from")).String(),
			ToAddress:   sdk.AccAddress([]byte("from")).String(),
			Amount:      sdk.NewCoins(sdk.NewCoin("abera", sdkmath.NewInt(100))),
		}
		msgAny, err := codectypes.NewAnyWithValue(&msg)
		Expect(err).ToNot(HaveOccurred())

		// Embed the send msg into a proposal.
		proposal := v1.MsgSubmitProposal{
			Messages: []*codectypes.Any{msgAny},
		}

		// Marshal the proposal.
		pBz, err := proposal.Marshal()
		Expect(err).ToNot(HaveOccurred())

		// Unmarshal the proposal.
		var p v1.MsgSubmitProposal
		err = p.Unmarshal(pBz)
		Expect(err).ToNot(HaveOccurred())
	})
})

// unmarshalMsgAndReturnAny unmarshals `[]byte` into a `codectypes.Any` message.
func unmarshalMsgAndReturnAny(bz []byte) (*codectypes.Any, error) {
	var msg banktypes.MsgSend
	if err := msg.Unmarshal(bz); err != nil {
		return nil, err
	}
	anyValue, err := codectypes.NewAnyWithValue(&msg)
	if err != nil {
		return nil, err
	}
	return anyValue, nil
}
