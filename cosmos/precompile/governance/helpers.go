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
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/governance"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/eth/common"
)

// PrecompileAddress returns the address of the precompile contract (0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2).
func PrecompileAddress() common.Address {
	eth := cosmlib.AccAddressToEthAddress(authtypes.NewModuleAddress(govtypes.ModuleName))
	return eth
}

// UnmarshalAnyBzSlice unmarshals a slice of bytes into a slice of Any.
func UnmarshalAnyBzSlice(bzs [][]byte) ([]*codectypes.Any, error) {
	anys := []*codectypes.Any{}
	for _, bz := range bzs {
		var a codectypes.Any

		if err := a.Unmarshal(bz); err != nil {
			return nil, err
		}

		anys = append(anys, &a)
	}

	return anys, nil
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
