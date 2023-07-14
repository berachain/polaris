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
		govtypes.AttributeKeyProposalID:       log.ConvertUint64,
		govtypes.AttributeKeyProposalMessages: log.ReturnStringAsIs,
		govtypes.AttributeKeyOption:           log.ReturnStringAsIs,
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
) ([]uint64, error) {
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
