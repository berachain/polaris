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
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/codec/unknownproto"
	sdk "github.com/cosmos/cosmos-sdk/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/governance"
)

// ProcessMessages takes in a slice of message bytes and returns a slice of codectypes.Any.
func (c *Contract) ProcessMessages(messagesBz [][]byte) ([]*codectypes.Any, error) {
	// Decode all the messageBz into sdk.Msg then wrap them into codectypes.Any and append to the slice.
	messages := []*codectypes.Any{}
	for _, msgBz := range messagesBz {
		// Decode the message bytes into a sdk.Msg.
		var msg sdk.Msg

		// Reject all unknown proto fields and unknown proto messages.
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

func transformToEvmProposal(p *v1.Proposal) generated.IGovernanceModuleProposal {
	// Convert the proposal messages into a slice of bytes.
	messages := [][]byte{}
	for _, msg := range p.Messages {
		messages = append(messages, msg.Value)
	}

	// Convert the proposal total deposit into evm coins.
	td := make([]generated.CosmosCoin, 0)
	for _, coin := range p.TotalDeposit {
		td = append(td, generated.CosmosCoin{
			Denom:  coin.Denom,
			Amount: coin.Amount.BigInt(),
		})
	}

	return generated.IGovernanceModuleProposal{
		Id:       p.Id,
		Messages: messages,
		Status:   int32(p.Status), // Status is an alias for int32.
		FinalTallyResult: generated.IGovernanceModuleTallyResult{
			YesCount:        p.FinalTallyResult.YesCount,
			AbstainCount:    p.FinalTallyResult.AbstainCount,
			NoCount:         p.FinalTallyResult.NoCount,
			NoWithVetoCount: p.FinalTallyResult.NoWithVetoCount,
		},
		SubmitTime:     uint64(p.SubmitTime.Unix()),
		DepositEndTime: uint64(p.DepositEndTime.Unix()),
		TotalDeposit:   td,
		Metadata:       p.Metadata,
		Title:          p.Title,
		Summary:        p.Summary,
		Proposer:       p.Proposer,
	}
}
