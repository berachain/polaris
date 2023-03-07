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

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	governancekeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	governancetypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	"pkg.berachain.dev/polaris/eth/accounts/abi"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/lib/utils"
	polarisprecompile "pkg.berachain.dev/polaris/precompile"
	"pkg.berachain.dev/polaris/precompile/contracts/solidity/generated"
	evmutils "pkg.berachain.dev/polaris/x/evm/utils"
)

// `Contract` is the precompile contract for the governance module.
type Contract struct {
	contractAbi *abi.ABI

	msgServer v1.MsgServer
	querier   v1.QueryServer
}

// `NewContract` is the constructor for the governance precompile contract.
func NewContract(gk **governancekeeper.Keeper) precompile.StatefulImpl {
	var contractAbi abi.ABI
	if err := contractAbi.UnmarshalJSON([]byte(generated.GovernanceModuleMetaData.ABI)); err != nil {
		panic(err)
	}
	return &Contract{
		contractAbi: &contractAbi,
		msgServer:   governancekeeper.NewMsgServerImpl(*gk),
		querier:     *gk,
	}
}

// `RegistryKey` implements the `precompile.StatefulImpl` interface.
func (c *Contract) RegistryKey() common.Address {
	return evmutils.AccAddressToEthAddress(authtypes.NewModuleAddress(governancetypes.ModuleName))
}

// `ABIMethods` implements the `precompile.StatefulImpl` interface.
func (c *Contract) ABIMethods() map[string]abi.Method {
	return c.contractAbi.Methods
}

// `ABIEvents` implements the `precompile.StatefulImpl` interface.
func (c *Contract) ABIEvents() map[string]abi.Event {
	return c.contractAbi.Events
}

// `CustomValueDecoders` implements the `precompile.StatefulImpl` interface.
func (c *Contract) CustomValueDecoders() precompile.ValueDecoders {
	return nil
}

// `PrecompileMethods` implements the `precompile.StatefulImpl` interface.
func (c *Contract) PrecompileMethods() precompile.Methods {
	return precompile.Methods{
		{
			AbiSig:  "submitProposal(bytes,[]tuple,string,string,string,bool)",
			Execute: c.SubmitProposal,
		},
		{
			AbiSig:  "cancelProposal(uint256)",
			Execute: c.CancelProposal,
		},
		{
			AbiSig:  "vote(uint256,int32,string)",
			Execute: c.Vote,
		},
		{
			AbiSig:  "voteWeighted(uint256,[]tuple,string)",
			Execute: c.VoteWeight,
		},
	}
}

// `SubmitProposal` is the method for the `submitProposal` method of the governance precompile contract.
func (c *Contract) SubmitProposal(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	message, ok := utils.GetAs[[]*codectypes.Any](args[0]) // TODO: check if this is the correct type.
	if !ok {
		return nil, polarisprecompile.ErrInvalidAny
	}
	initialDeposit, ok := utils.GetAs[[]generated.IGovernanceModuleCoin](args[1])
	if !ok {
		return nil, polarisprecompile.ErrInvalidCoin
	}
	metadata, ok := utils.GetAs[string](args[2])
	if !ok {
		return nil, polarisprecompile.ErrInvalidString
	}
	title, ok := utils.GetAs[string](args[3])
	if !ok {
		return nil, polarisprecompile.ErrInvalidString
	}
	summary, ok := utils.GetAs[string](args[4])
	if !ok {
		return nil, polarisprecompile.ErrInvalidString
	}
	expedited, ok := utils.GetAs[bool](args[5])
	if !ok {
		return nil, polarisprecompile.ErrInvalidBool
	}

	// Caller is the proposer (msg.sender).
	proposer := sdk.AccAddress(caller.Bytes())

	return c.submitProposalHelper(ctx, message, initialDeposit, proposer, metadata, title, summary, expedited)
}

// `CancelProposal` is the method for the `cancelProposal` method of the governance precompile contract.
func (c *Contract) CancelProposal(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	id, ok := utils.GetAs[*big.Int](args[0])
	if !ok {
		return nil, polarisprecompile.ErrInvalidBigInt
	}
	proposer := sdk.AccAddress(caller.Bytes())

	return c.cancelProposalHelper(ctx, proposer, id)
}

// `Vote` is the method for the `vote` method of the governance precompile contract.
func (c *Contract) Vote(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	proposalID, ok := utils.GetAs[*big.Int](args[0])
	if !ok {
		return nil, polarisprecompile.ErrInvalidBigInt
	}
	options, ok := utils.GetAs[int32](args[1])
	if !ok {
		return nil, polarisprecompile.ErrInvalidInt32
	}
	metadata, ok := utils.GetAs[string](args[2])
	if !ok {
		return nil, polarisprecompile.ErrInvalidString
	}
	voter := sdk.AccAddress(caller.Bytes())

	return c.voteHelper(ctx, voter, proposalID, options, metadata)
}

// `VoteWeighted` is the method for the `voteWeighted` method of the governance precompile contract.
func (c *Contract) VoteWeight(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	proposalID, ok := utils.GetAs[*big.Int](args[0])
	if !ok {
		return nil, polarisprecompile.ErrInvalidBigInt
	}
	options, ok := utils.GetAs[[]generated.IGovernanceModuleWeightedVoteOption](args[1])
	if !ok {
		return nil, polarisprecompile.ErrInvalidOptions
	}
	metadata, ok := utils.GetAs[string](args[2])
	if !ok {
		return nil, polarisprecompile.ErrInvalidString
	}
	voter := sdk.AccAddress(caller.Bytes())
	return c.voteWeightedHelper(ctx, voter, proposalID, options, metadata)
}
