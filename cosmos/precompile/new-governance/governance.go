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

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/codec/unknownproto"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/governance"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
)

// Contract is the governance precompile contract.
type Contract struct {
	ethprecompile.BaseContract

	msgServer   v1.MsgServer
	queryServer v1.QueryServer
	cdc         codec.ProtoCodecMarshaler
}

// NewPrecompileContract creates a new governance precompile contract.
func NewPrecompileContract(m v1.MsgServer, q v1.QueryServer, cdc codec.ProtoCodecMarshaler) *Contract {
	return &Contract{
		BaseContract: ethprecompile.NewBaseContract(generated.GovernanceModuleMetaData.ABI, PrecompileAddress()),

		msgServer:   m,
		queryServer: q,
		cdc:         cdc,
	}
}

// PrecompileAddress returns the address of the precompile contract (0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2).
func PrecompileAddress() common.Address {
	eth := cosmlib.AccAddressToEthAddress(authtypes.NewModuleAddress(govtypes.ModuleName))
	return eth
}

// SubmitProposal is the method for the `submitProposal` function.
func (c *Contract) SubmitProposal(ctx context.Context, proposalBz []byte, messageBz [][]byte) (uint64, error) {
	// Decode the proposal bytes into  v1.MsgSubmitProposal.
	var proposal v1.MsgSubmitProposal
	if err := proposal.Unmarshal(proposalBz); err != nil {
		return 0, err
	}

	// Decode all the messages into sdk.Msg then wrap them into codectypes.Any.
	messages, err := c.ProcessMessages(messageBz)
	if err != nil {
		return 0, err
	}
	proposal.Messages = messages

	// Create the proposal.
	res, err := c.msgServer.SubmitProposal(ctx, &proposal)
	if err != nil {
		return 0, err
	}

	return res.ProposalId, nil

}

// ProcessMessages takes in a slice of message bytes and returns a slice of codectypes.Any.
func (c *Contract) ProcessMessages(messagesBz [][]byte) ([]*codectypes.Any, error) {
	// Decode all the messageBz into sdk.Msg then wrap them into codectypes.Any and append to the slice.
	var messages []*codectypes.Any
	for _, msgBz := range messagesBz {
		// Decode the message bytes into a sdk.Msg.
		var msg sdk.Msg

		// Reject all unkown proto fields and unknown proto messages.
		if err := unknownproto.RejectUnknownFieldsStrict(msgBz, msg, c.cdc.InterfaceRegistry()); err != nil {
			return nil, err
		}

		// Unmarshal the message bytes into a sdk.Msg.
		if err := c.cdc.Unmarshal(msgBz, msg); err != nil {
			return nil, err
		}

		// Wrap the msg into any proto message and add to the proposal msg.
		msgAny, err := codectypes.NewAnyWithValue(msg)
		if err != nil {
			return nil, err
		}

		messages = append(messages, msgAny)
	}

	return messages, nil
}
