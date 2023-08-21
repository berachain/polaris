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

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/governance"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/precompile/log"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/core/vm"
)

// Contract is the precompile contract for the governance module.
type Contract struct {
	ethprecompile.BaseContract

	msgServer v1.MsgServer
	querier   v1.QueryServer
}

// NewPrecompileContract creates a new precompile contract for the governance module.
func NewPrecompileContract(m v1.MsgServer, q v1.QueryServer) *Contract {
	return &Contract{
		BaseContract: ethprecompile.NewBaseContract(
			generated.GovernanceModuleMetaData.ABI,
			// Precompile Address: 0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2
			cosmlib.AccAddressToEthAddress(authtypes.NewModuleAddress(govtypes.ModuleName)),
		),
		msgServer: m,
		querier:   q,
	}
}

// CustomValueDecoders implements the `ethprecompile.StatefulImpl` interface.
func (c *Contract) CustomValueDecoders() ethprecompile.ValueDecoders {
	return ethprecompile.ValueDecoders{
		AttributeProposalSender: log.ConvertCommonHexAddress,
	}
}

// SubmitProposal is the method for the `submitProposal` method of the governance precompile contract.
func (c *Contract) SubmitProposal(
	ctx context.Context,
	proposalBz []byte,
	messageBz []byte,
) (uint64, error) {
	message, err := unmarshalMsgAndReturnAny(messageBz)
	if err != nil {
		return 0, err
	}
	return c.submitProposalHelper(ctx, proposalBz, []*codectypes.Any{message})
}

// CancelProposal is the method for the `cancelProposal` method of the governance precompile contract.
func (c *Contract) CancelProposal(
	ctx context.Context,
	id uint64,
) (uint64, uint64, error) {
	proposer := sdk.AccAddress(vm.UnwrapPolarContext(ctx).MsgSender().Bytes())

	return c.cancelProposalHelper(ctx, proposer, id)
}

// Vote is the method for the `vote` method of the governance precompile contract.
func (c *Contract) Vote(
	ctx context.Context,
	proposalID uint64,
	options int32,
	metadata string,
) (bool, error) {
	voter := sdk.AccAddress(vm.UnwrapPolarContext(ctx).MsgSender().Bytes())

	return c.voteHelper(ctx, voter, proposalID, options, metadata)
}

// VoteWeighted is the method for the `voteWeighted` method of the governance precompile contract.
func (c *Contract) VoteWeighted(
	ctx context.Context,
	proposalID uint64,
	options []generated.IGovernanceModuleWeightedVoteOption,
	metadata string,
) (bool, error) {
	voter := sdk.AccAddress(vm.UnwrapPolarContext(ctx).MsgSender().Bytes())
	return c.voteWeightedHelper(ctx, voter, proposalID, options, metadata)
}

// GetProposal is the method for the `getProposal` method of the governance precompile contract.
func (c *Contract) GetProposal(
	ctx context.Context,
	proposalID uint64,
) (generated.IGovernanceModuleProposal, error) {
	return c.getProposalHelper(ctx, proposalID)
}

// GetProposals is the method for the `getProposal` method of the governance precompile contract.
func (c *Contract) GetProposals(
	ctx context.Context,
	proposalStatus int32,
) ([]generated.IGovernanceModuleProposal, error) {
	return c.getProposalsHelper(ctx, proposalStatus)
}

// GetProposalDeposits is the method for the `getProposalDeposits` method of the governance precompile contract.
func (c *Contract) GetProposalDeposits(
	ctx context.Context,
	proposalID uint64,
) ([]generated.CosmosCoin, error) {
	return c.getProposalDepositsHelper(ctx, proposalID)
}

// GetProposalDepositsByDepositor is the method for the `getProposalDepositsByDepositor` method
// of the governance precompile contract.
func (c *Contract) GetProposalDepositsByDepositor(
	ctx context.Context,
	proposalID uint64,
	depositor common.Address,
) ([]generated.CosmosCoin, error) {
	return c.getProposalDepositsByDepositorHelper(ctx, depositor, proposalID)
}

// GetProposalVotes is the method for the `getProposalVotes` method of the governance precompile contract.
func (c *Contract) GetProposalTallyResult(
	ctx context.Context,
	proposalID uint64,
) (generated.IGovernanceModuleTallyResult, error) {
	return c.getProposalTallyResultHelper(ctx, proposalID)
}

// GetProposalVotes is the method for the `getProposalVotes` method of the governance precompile contract.
func (c *Contract) GetProposalVotes(
	ctx context.Context,
	proposalID uint64,
) ([]generated.IGovernanceModuleVote, error) {
	return c.getProposalVotesHelper(ctx, proposalID)
}

// GetProposalVotesByVoter is the method for the `getProposalVotesByVoter` method of the governance
// precompile contract.
func (c *Contract) GetProposalVotesByVoter(
	ctx context.Context,
	proposalID uint64,
	voter common.Address,
) (generated.IGovernanceModuleVote, error) {
	return c.getProposalVotesByVoterHelper(ctx, proposalID, voter)
}

// GetProposalVoteByVoter is the method for the `getProposalVoteByVoter` method of the governance
// precompile contract.
func (c *Contract) GetParams(
	ctx context.Context,
) (generated.IGovernanceModuleParams, error) {
	return c.getParamsHelper(ctx)
}

// GetDepositParams is the method for the `getDepositParams` method of the governance precompile contract.
func (c *Contract) GetDepositParams(
	ctx context.Context,
) (generated.IGovernanceModuleDepositParams, error) {
	return c.getDepositParamsHelper(ctx)
}

// GetVotingParams is the method for the `getVotingParams` method of the governance precompile contract.
func (c *Contract) GetVotingParams(
	ctx context.Context,
) (generated.IGovernanceModuleVotingParams, error) {
	return c.getVotingParamsHelper(ctx)
}

// GetTallyParams is the method for the `getTallyParams` method of the governance precompile contract.
func (c *Contract) GetTallyParams(
	ctx context.Context,
) (generated.IGovernanceModuleTallyParams, error) {
	return c.getTallyParamsHelper(ctx)
}

// GetConstitution is the method for the `getConstitution` method of the governance precompile contract.
func (c *Contract) GetConstitution(
	ctx context.Context,
) (string, error) {
	return c.getConstitutionHelper(ctx)
}

// unmarshalMsgAndReturnAny unmarshals `[]byte` into a `codectypes.Any` message.
// TODO: This is a temporary solution until we have a better way to unmarshal messages.
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
